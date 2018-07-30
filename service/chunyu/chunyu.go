package chunyu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	cmn "github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/db"
	"github.com/shiqinfeng1/chunyuyisheng/service/nsqs"
	"github.com/shiqinfeng1/gorequest"
)

/* getSign 签名
partner_key: 合作方 partner_key,注意不是 partner
atime: UNIX TIMESTAMP 最小单位为秒
user_id: 第三方用户唯一标识，可以为字母与数字组合的字符串
*/
func getSign(partnerKey, atime, UserID string) string {
	md5Ctx := md5.New()
	data := partnerKey + atime + UserID
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)

	//获得签名: md5的32位结果取中间16位
	return hex.EncodeToString(cipherStr)[8:24]
}

func newRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	return request
}

//DoctorResponseCallback 医生回复回调函数
func DoctorResponseCallback(payload cmn.ChunyuDoctorResponsePayload) (err error) {
	/*
		if payload.Sign != getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(payload.Atime, 10), strconv.Itoa(payload.ProblemID)) {
			return
		}
	*/
	err = nsqs.PostTopic(cmn.TopicChunyuDoctorResponse, payload)
	return
}

//QuestionCloseCallback 问题关闭回调
func QuestionCloseCallback(payload cmn.ChunyuQuestionClosePayload) (err error) {
	err = nsqs.PostTopic(cmn.TopicChunyuQuestionClose, payload)
	return
}

func createUserSyncToDB(logininfo UserLoginRequest) error {
	userinfo := &db.UserInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()
	dbconn.Model(&db.UserInfo{}).Where("user_id = ?", logininfo.UserID).Find(&userinfo)
	if userinfo.IsSynced == true {
		return nil
	}
	userinfo.IsSynced = true
	userinfo.UserID = logininfo.UserID

	err := dbconn.Create(userinfo).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

func userIsSynced(userID string) bool {
	userinfo := &db.UserInfo{}
	dbconn := db.MysqlBegin()
	dbconn.Model(&db.UserInfo{}).Where("user_id = ?", userID).Find(&userinfo)
	dbconn.MysqlRollback()
	return userinfo.IsSynced
}

//UserLogin 患者登录信息同步到春雨平台
func UserLogin(payload cmn.PatientLoginPayload) (*UserLoginReponse, error) {
	now := time.Now().Unix()

	// if cmn.UserVerify(payload.UserID) == false {
	// 	return nil, nil
	// }
	reqArgs := UserLoginRequest{
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:   payload.UserID,
		Atime:    now,
		Lon:      payload.Lon,
		Lat:      payload.Lat,
		Password: payload.Password,
	}
	var resp UserLoginReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/login").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.UserLogin error: %q", errs)
		return nil, err
	}

	err := createUserSyncToDB(reqArgs)

	return &resp, err

}

//FreeProblemCreate 创建众包问题
func FreeProblemCreate(payload cmn.FreeProblemPayload) (*ProblemIDReponse, error) {
	var ProblemID ProblemIDReponse
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		cmn.Logger.Error(err)
		return &ProblemID, err
	}
	content, _ := json.Marshal(payload.Content)
	reqArgs := FreeProblemCreateRequest{
		ClinicNo: payload.ClinicNo,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:   payload.UserID,
		Atime:    now,
		Content:  string(content),
	}

	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/free_problem/create").
		Send(reqArgs).
		EndStruct(&ProblemID)
	if errs != nil {
		err := fmt.Errorf("chunyu.FreeProblemCreate error: %q", errs)
		return nil, err
	}
	return &ProblemID, nil
}

func createPaidProblemInfoToDB(req PaidProblemCreateRequest, rep ProblemIDReponse) error {
	probleminfo := &db.ProblemInfo{}
	probleminfo.UserID = req.UserID
	probleminfo.PartnerOrderID = req.PartnerOrderID
	probleminfo.ProblemID = rep.ProblemID
	probleminfo.Content = req.Content
	probleminfo.PaidType = "PaidProblem"

	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	err := dbconn.Create(probleminfo).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

func createOrientedProblemInfoToDB(req OrientedProblemCreateRequest, rep ProblemAndDoctorReponse) error {
	probleminfo := &db.ProblemInfo{}
	probleminfo.UserID = req.UserID
	probleminfo.PartnerOrderID = req.PartnerOrderID
	probleminfo.Content = req.Content
	probleminfo.Price = req.Price
	probleminfo.PaidType = "OrientedProblem"

	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	for _, v := range rep.Problems {
		probleminfo.ProblemID = v.ProblemID
		probleminfo.DoctorID = v.DoctorID
		err := dbconn.Create(probleminfo).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	}

	dbconn.MysqlCommit()
	return nil
}

func updateProblemInfoToDB(userid string, problemid int, item string, value interface{}) error {

	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	err := dbconn.Model(&db.ProblemInfo{}).
		Where("problem_id = ? AND user_id = ?", problemid, userid).
		Update(item, value).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

func appendProblemInfoToDB(req AppendProblemRequest) error {
	appendprobleminfo := &db.AppendProblemInfo{}
	appendprobleminfo.Content = req.Content
	appendprobleminfo.ProblemID = req.ProblemID
	appendprobleminfo.UserID = req.UserID

	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	err := dbconn.Create(appendprobleminfo).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

func assessProblemInfoToDB(req AssessProblemRequest) error {
	assessprobleminfo := &db.AssessProblemInfo{}
	assessprobleminfo.AssessContent = req.Content
	assessprobleminfo.ProblemID = req.ProblemID
	assessprobleminfo.UserID = req.UserID
	assessprobleminfo.AssessInfo = req.AssessInfo

	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	err := dbconn.Create(assessprobleminfo).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

//PaidProblemCreate 创建众包升级问题
func PaidProblemCreate(payload cmn.PaidProblemPayload) (*ProblemIDReponse, error) {
	now := time.Now().Unix()
	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}
	content, _ := json.Marshal(payload.Content)
	reqArgs := PaidProblemCreateRequest{
		FreeProblemCreateRequest{
			ClinicNo: payload.ClinicNo,
			Partner:  cmn.Config().GetString("chunyu.partner"),
			Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
			UserID:   payload.UserID,
			Atime:    now,
			Content:  string(content),
		},
		payload.PartnerOrderID,
		payload.PayType,
	}

	var ProblemID ProblemIDReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/create_paid_problem/").
		Send(reqArgs).
		EndStruct(&ProblemID)
	if errs != nil {
		err := fmt.Errorf("chunyu.PaidProblemCreate error: %q", errs)
		return nil, err
	}
	err := createPaidProblemInfoToDB(reqArgs, ProblemID)
	return &ProblemID, err
}

//PaidProblemRefund 众包升级问题未回复后退款
func PaidProblemRefund(payload cmn.PaidProblemRefundPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	reqArgs := PaidProblemRefundRequest{
		ProblemID: payload.ProblemID,
		Partner:   cmn.Config().GetString("chunyu.partner"),
		Sign:      getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:    payload.UserID,
		Atime:     now,
	}

	var resp ErrorMsgReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/refund").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.PaidProblemRefund error: %q", errs)
		return nil, err
	}
	return &resp, nil
}

//PaidProblemQueryClinicNo 众包升级问题查询科室
func PaidProblemQueryClinicNo(payload cmn.PaidProblemQueryClinicNoPayload) (*ClinicNoReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	content, _ := json.Marshal(payload.Ask)
	reqArgs := PaidProblemQueryClinicNoRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:  payload.UserID,
		Atime:   now,
		Ask:     string(content),
	}

	var clinicno ClinicNoReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/get_problem_clinic_no/").
		Send(reqArgs).
		EndStruct(&clinicno)
	if errs != nil {
		err := fmt.Errorf("chunyu.PaidProblemQueryClinicNo error: %q", errs)
		return nil, err
	}
	return &clinicno, nil
}

//GetClinicDoctors 获取指定科室指定城市的医生列表
func GetClinicDoctors(payload cmn.ClinicDoctorsPayload) (*ClinicDoctorReponse, error) {
	now := time.Now().Unix()

	reqArgs := ClinicDoctorRequest{
		ClinicNo: payload.ClinicNo,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:   payload.UserID,
		Atime:    now,
		StartNum: payload.Page,
		Count:    payload.PerPage,
		Province: payload.Province,
		City:     payload.City,
	}

	var doctorsList ClinicDoctorReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/get_clinic_doctors").
		Send(reqArgs).
		EndStruct(&doctorsList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetClinicDoctors error: %q", errs)
		return nil, err
	}
	return &doctorsList, nil
}

//GetAskHistory 获取提问历史
func GetAskHistory(payload cmn.AskHistoryPayload) (*[]AskHistoryReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	reqArgs := AskHistoryRequest{
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:   payload.UserID,
		Atime:    now,
		StartNum: payload.Page,
		Count:    payload.PerPage,
	}

	var history []AskHistoryReponse

	_, bodyBytes, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/list/my").
		Send(reqArgs).
		EndBytes()
	if errs != nil {
		err := fmt.Errorf("chunyu.GetAskHistory error: %q", errs)
		return nil, err
	}
	json.Unmarshal(bodyBytes, &history)
	return &history, nil
}

//GetRecommendedDoctors 获取推荐的医生列表
func GetRecommendedDoctors(payload cmn.RecommendedDoctorsPayload) (*ClinicDoctorReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	content, _ := json.Marshal(payload.Ask)
	reqArgs := RecommendedDoctorRequest{
		Ask:     string(content),
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:  payload.UserID,
		Atime:   now,
	}

	var doctorsList ClinicDoctorReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/get_recommended_doctors").
		Send(reqArgs).
		EndStruct(&doctorsList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetRecommendedDoctors error: %q", errs)
		return nil, err
	}
	return &doctorsList, nil
}

//GetDoctorDetail 获取医生详情
func GetDoctorDetail(payload cmn.DoctorDetailPayload) (*DoctorDetailReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	reqArgs := DoctorDetailRequest{
		UserID:   payload.UserID,
		DoctorID: payload.DoctorID,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		Atime:    now,
	}

	var doctorsdetail DoctorDetailReponse
	_, bodyBytes, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/detail").
		Send(reqArgs).
		EndBytes()
	if errs != nil {
		err := fmt.Errorf("chunyu.GetDoctorDetail error: %q", errs)
		return nil, err
	}
	json.Unmarshal(bodyBytes, &doctorsdetail)
	return &doctorsdetail, nil
}

//GetProblemDetail 获取医生详情
func GetProblemDetail(payload cmn.ProblemDetailPayload) (*ProblemDetailReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	reqArgs := ProblemDetailRequest{
		ProblemID:     payload.ProblemID,
		LastContentID: payload.LastContentID,
		UserID:        payload.UserID,
		Partner:       cmn.Config().GetString("chunyu.partner"),
		Sign:          getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		Atime:         now,
	}

	var problem ProblemDetailReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/detail").
		Send(reqArgs).
		EndStruct(&problem)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetDoctorDetail error: %q", errs)
		return nil, err
	}
	return &problem, nil
}

//OrientedProblemCreate 创建定向问题
func OrientedProblemCreate(payload cmn.OrientedProblemPayload) (*ProblemAndDoctorReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	content, _ := json.Marshal(payload.Content)
	reqArgs := OrientedProblemCreateRequest{
		DoctorIDs:      payload.DoctorIDs,
		Partner:        cmn.Config().GetString("chunyu.partner"),
		Sign:           getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:         payload.UserID,
		Atime:          now,
		PartnerOrderID: payload.PartnerOrderID,
		Content:        string(content),
		Price:          payload.Price,
	}

	var Problemlist ProblemAndDoctorReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/create_oriented_problem/").
		Send(reqArgs).
		EndStruct(&Problemlist)
	if errs != nil {
		err := fmt.Errorf("chunyu.OrientedProblemCreate error: %q", errs)
		return nil, err
	}

	err := createOrientedProblemInfoToDB(reqArgs, Problemlist)
	return &Problemlist, err
}

//GetEmergencyGraph 获取急诊图文信息
func GetEmergencyGraph(payload cmn.EmergencyGraphPayload) (*EmergencyGraphReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	content, _ := json.Marshal(payload.Content)
	reqArgs := EmergencyGraphRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:  payload.UserID,
		Atime:   now,
		Content: string(content),
	}

	var clinicList EmergencyGraphReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/get_emergency_graph_info/").
		Send(reqArgs).
		EndStruct(&clinicList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetEmergencyGraph error: %q", errs)
		return nil, err
	}
	return &clinicList, nil
}

//EmergencyGraphCreate 创建急诊图文问题
func EmergencyGraphCreate(payload cmn.EmergencyGraphCreatePayload) (*ProblemIDReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	content, _ := json.Marshal(payload.Content)
	reqArgs := EmergencyGraphCreateRequest{
		FreeProblemCreateRequest{
			ClinicNo: payload.ClinicNo,
			Partner:  cmn.Config().GetString("chunyu.partner"),
			Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
			UserID:   payload.UserID,
			Atime:    now,
			Content:  string(content),
		},
		payload.PartnerOrderID,
		payload.Price,
	}

	var ProblemID ProblemIDReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/create_emergency_graph/").
		Send(reqArgs).
		EndStruct(&ProblemID)
	if errs != nil {
		err := fmt.Errorf("chunyu.EmergencyGraphCreate error: %q", errs)
		return nil, err
	}
	return &ProblemID, nil
}

//GetFastPhoneInfo 获取急诊图文信息
func GetFastPhoneInfo(payload cmn.FastPhoneInfoPayload) (*FastPhoneInfoReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	reqArgs := FastPhoneInfoRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:  payload.UserID,
		Atime:   now,
	}

	var clinicList FastPhoneInfoReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/phone/get_fast_phone_info/").
		Send(reqArgs).
		EndStruct(&clinicList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetFastPhoneInfo error: %q", errs)
		return nil, err
	}
	return &clinicList, nil
}

//FastPhoneOrderCreate 创建急诊图文问题
func FastPhoneOrderCreate(payload cmn.FastPhoneOrderCreatePayload) (*FastPhoneOrderReponse, error) {
	now := time.Now().Unix()

	if userIsSynced(payload.UserID) == false {
		err := fmt.Errorf("user: %s is not login", payload.UserID)
		return nil, err
	}

	reqArgs := FastPhoneOrderRequest{
		ClinicNo:       payload.ClinicNo,
		Partner:        cmn.Config().GetString("chunyu.partner"),
		Sign:           getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:         payload.UserID,
		Atime:          now,
		PartnerOrderID: payload.PartnerOrderID,
		Phone:          payload.Phone,
	}

	var sercviceid FastPhoneOrderReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/phone/create_fast_phone_order/").
		Send(reqArgs).
		EndStruct(&sercviceid)
	if errs != nil {
		err := fmt.Errorf("chunyu.FastPhoneCreate error: %q", errs)
		return nil, err
	}
	return &sercviceid, nil
}

//AppendProblem 追加问题
func AppendProblem(payload cmn.AppendProblemPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()
	content, _ := json.Marshal(payload.Content)
	reqArgs := AppendProblemRequest{
		ProblemID: payload.ProblemID,
		Partner:   cmn.Config().GetString("chunyu.partner"),
		Sign:      getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		UserID:    payload.UserID,
		Atime:     now,
		Content:   string(content),
	}

	var resp ErrorMsgReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem_content/create").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.AppendProblem error: %q", errs)
		return nil, err
	}
	err := appendProblemInfoToDB(reqArgs)
	return &resp, err
}

//AssessProblem 评价问题
func AssessProblem(payload cmn.AssessProblemPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	content, _ := json.Marshal(payload.Content)
	assess, _ := json.Marshal(payload.AssessInfo)
	reqArgs := AssessProblemRequest{
		ProblemID:  payload.ProblemID,
		Content:    string(content),
		UserID:     payload.UserID,
		AssessInfo: string(assess),
		Partner:    cmn.Config().GetString("chunyu.partner"),
		Sign:       getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		Atime:      now,
	}

	var resp ErrorMsgReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/assess").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.AssessProblem error: %q", errs)
		return nil, err
	}
	err := assessProblemInfoToDB(reqArgs)
	return &resp, err
}

//DeleteProblem 删除问题
func DeleteProblem(payload cmn.DeleteProblemPayload, flag string) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	reqArgs := DeleteProblemRequest{
		ProblemID: payload.ProblemID,
		UserID:    payload.UserID,
		Partner:   cmn.Config().GetString("chunyu.partner"),
		Sign:      getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserID),
		Atime:     now,
	}

	var resp ErrorMsgReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/assess").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.AssessProblem error: %q", errs)
		return nil, err
	}
	updateProblemInfoToDB(reqArgs.UserID, reqArgs.ProblemID, flag, true)
	return &resp, nil
}

//CloseProblem 关闭问题
func CloseProblem(payload cmn.DeleteProblemPayload, flag string) (*ErrorMsgReponse, error) {
	return DeleteProblem(payload, flag)
}
