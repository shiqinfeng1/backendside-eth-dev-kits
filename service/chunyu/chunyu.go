package chunyu

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	cmn "github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/nsqs"
	"github.com/shiqinfeng1/gorequest"
)

/*
partner_key: 合作方 partner_key,注意不是 partner
atime: UNIX TIMESTAMP 最小单位为秒
user_id: 第三方用户唯一标识，可以为字母与数字组合的字符串
*/
func getSign(partner_key, atime, user_id string) string {
	md5Ctx := md5.New()
	data := partner_key + atime + user_id
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)

	//获得签名: md5的32位结果取中间16位
	return hex.EncodeToString(cipherStr)[8:24]
}

func DoctorResponseCallback(payload cmn.ChunyuDoctorResponsePayload) (err error) {
	err = nsqs.PostTopic(cmn.TopicChunyuDoctorResponse, payload)
	return
}
func QuestionCloseCallback(payload cmn.ChunyuQuestionClosePayload) (err error) {
	err = nsqs.PostTopic(cmn.TopicChunyuQuestionClose, payload)
	return
}

func newRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	return request
}

//患者登录
func UserLogin(payload cmn.PatientLoginPayload) (*UserLoginReponse, error) {
	now := time.Now().Unix()

	reqArgs := UserLoginRequest{
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:   payload.UserId,
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
	return &resp, nil

}

//创建众包问题
func FreeProblemCreate(payload cmn.FreeProblemPayload) (*ProblemIDReponse, error) {
	now := time.Now().Unix()

	reqArgs := FreeProblemCreateRequest{
		ClinicNo: payload.ClinicNo,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:   payload.UserId,
		Atime:    now,
		Content:  payload.Content,
	}

	var Problemid ProblemIDReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/free_problem/create").
		Send(reqArgs).
		EndStruct(&Problemid)
	if errs != nil {
		err := fmt.Errorf("chunyu.FreeProblemCreate error: %q", errs)
		return nil, err
	}
	return &Problemid, nil
}

//创建众包升级问题
func PaidProblemCreate(payload cmn.PaidProblemPayload) (*ProblemIDReponse, error) {
	now := time.Now().Unix()

	reqArgs := PaidProblemCreateRequest{
		FreeProblemCreateRequest{
			ClinicNo: payload.ClinicNo,
			Partner:  cmn.Config().GetString("chunyu.partner"),
			Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
			UserId:   payload.UserId,
			Atime:    now,
			Content:  payload.Content,
		},
		payload.PartnerOrderId,
		payload.PayType,
	}

	var Problemid ProblemIDReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/create_paid_problem").
		Send(reqArgs).
		EndStruct(&Problemid)
	if errs != nil {
		err := fmt.Errorf("chunyu.PaidProblemCreate error: %q", errs)
		return nil, err
	}
	return &Problemid, nil
}

//众包升级问题未回复后退款
func PaidProblemRefund(payload cmn.PaidProblemRefundPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	reqArgs := PaidProblemRefundRequest{
		ProblemId: payload.ProblemId,
		Partner:   cmn.Config().GetString("chunyu.partner"),
		Sign:      getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:    payload.UserId,
		Atime:     now,
	}

	var resp ErrorMsgReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/refund").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.PaidProblemRefund error: %q", errs)
		return nil, err
	}
	return &resp, nil
}

//众包升级问题查询科室
func PaidProblemQueryClinicNo(payload cmn.PaidProblemQueryClinicNoPayload) (*ClinicNoReponse, error) {
	now := time.Now().Unix()

	reqArgs := PaidProblemQueryClinicNoRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:  payload.UserId,
		Atime:   now,
		Ask:     payload.Ask,
	}

	var clinicno ClinicNoReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/get_problem_clinic_no").
		Send(reqArgs).
		EndStruct(&clinicno)
	if errs != nil {
		err := fmt.Errorf("chunyu.PaidProblemQueryClinicNo error: %q", errs)
		return nil, err
	}
	return &clinicno, nil
}

//获取指定科室指定城市的医生列表
func GetClinicDoctors(payload cmn.ClinicDoctorsPayload) (*ClinicDoctorReponse, error) {
	now := time.Now().Unix()

	reqArgs := ClinicDoctorRequest{
		ClinicNo: payload.ClinicNo,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:   payload.UserId,
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

//获取提问历史
func GetAskHistory(payload cmn.AskHistoryPayload) (*AskHistoryReponse, error) {
	now := time.Now().Unix()

	reqArgs := AskHistoryRequest{
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:   payload.UserId,
		Atime:    now,
		StartNum: payload.Page,
		Count:    payload.PerPage,
	}

	var history AskHistoryReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/list/my").
		Send(reqArgs).
		EndStruct(&history)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetAskHistory error: %q", errs)
		return nil, err
	}
	return &history, nil
}

//获取推荐的医生列表
func GetRecommendedDoctors(payload cmn.RecommendedDoctorsPayload) (*ClinicDoctorReponse, error) {
	now := time.Now().Unix()

	reqArgs := RecommendedDoctorRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:  payload.UserId,
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

//获取医生详情
func GetDoctorDetail(payload cmn.DoctorDetailPayload) (*DoctorDetailReponse, error) {
	now := time.Now().Unix()

	reqArgs := DoctorDetailRequest{
		UserId:   payload.UserId,
		DoctorId: payload.DoctorId,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		Atime:    now,
	}

	var doctorsdetail DoctorDetailReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/detail").
		Send(reqArgs).
		EndStruct(&doctorsdetail)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetDoctorDetail error: %q", errs)
		return nil, err
	}
	return &doctorsdetail, nil
}

//获取医生详情
func GetProblemDetail(payload cmn.ProblemDetailPayload) (*ProblemDetailReponse, error) {
	now := time.Now().Unix()

	reqArgs := ProblemDetailRequest{
		ProblemId:     payload.ProblemId,
		LastContentId: payload.LastContentId,
		UserId:        payload.UserId,
		Partner:       cmn.Config().GetString("chunyu.partner"),
		Sign:          getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		Atime:         now,
	}

	var problem ProblemDetailReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/detail").
		Send(reqArgs).
		EndStruct(&problem)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetDoctorDetail error: %q", errs)
		return nil, err
	}
	return &problem, nil
}

//创建定向问题
func OrientedProblemCreate(payload cmn.OrientedProblemPayload) (*ProblemAndDoctorReponse, error) {
	now := time.Now().Unix()

	reqArgs := OrientedProblemCreateRequest{
		DoctorIds:      payload.DoctorIds,
		Partner:        cmn.Config().GetString("chunyu.partner"),
		Sign:           getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:         payload.UserId,
		Atime:          now,
		PartnerOrderId: payload.PartnerOrderId,
		Content:        payload.Content,
		Price:          payload.Price,
	}

	var Problemlist ProblemAndDoctorReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/create_oriented_problem").
		Send(reqArgs).
		EndStruct(&Problemlist)
	if errs != nil {
		err := fmt.Errorf("chunyu.OrientedProblemCreate error: %q", errs)
		return nil, err
	}
	return &Problemlist, nil
}

//获取急诊图文信息
func GetEmergencyGraph(payload cmn.EmergencyGraphPayload) (*EmergencyGraphReponse, error) {
	now := time.Now().Unix()

	reqArgs := EmergencyGraphRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:  payload.UserId,
		Atime:   now,
		Content: payload.Content,
	}

	var clinicList EmergencyGraphReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem/get_emergency_graph_info").
		Send(reqArgs).
		EndStruct(&clinicList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetEmergencyGraph error: %q", errs)
		return nil, err
	}
	return &clinicList, nil
}

//创建急诊图文问题
func EmergencyGraphCreate(payload cmn.EmergencyGraphCreatePayload) (*ProblemIDReponse, error) {
	now := time.Now().Unix()

	reqArgs := EmergencyGraphCreateRequest{
		FreeProblemCreateRequest{
			ClinicNo: payload.ClinicNo,
			Partner:  cmn.Config().GetString("chunyu.partner"),
			Sign:     getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
			UserId:   payload.UserId,
			Atime:    now,
			Content:  payload.Content,
		},
		payload.PartnerOrderId,
		payload.Price,
	}

	var Problemid ProblemIDReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/create_emergency_graph").
		Send(reqArgs).
		EndStruct(&Problemid)
	if errs != nil {
		err := fmt.Errorf("chunyu.EmergencyGraphCreate error: %q", errs)
		return nil, err
	}
	return &Problemid, nil
}

//获取急诊图文信息
func GetFastPhoneInfo(payload cmn.FastPhoneInfoPayload) (*FastPhoneInfoReponse, error) {
	now := time.Now().Unix()

	reqArgs := FastPhoneInfoRequest{
		Partner: cmn.Config().GetString("chunyu.partner"),
		Sign:    getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:  payload.UserId,
		Atime:   now,
	}

	var clinicList FastPhoneInfoReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/phone/get_fast_phone_info").
		Send(reqArgs).
		EndStruct(&clinicList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetFastPhoneInfo error: %q", errs)
		return nil, err
	}
	return &clinicList, nil
}

//创建急诊图文问题
func FastPhoneOrderCreate(payload cmn.FastPhoneOrderCreatePayload) (*FastPhoneOrderReponse, error) {
	now := time.Now().Unix()

	reqArgs := FastPhoneOrderRequest{
		ClinicNo:       payload.ClinicNo,
		Partner:        cmn.Config().GetString("chunyu.partner"),
		Sign:           getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:         payload.UserId,
		Atime:          now,
		PartnerOrderId: payload.PartnerOrderId,
		Phone:          payload.Phone,
	}

	var sercviceid FastPhoneOrderReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/phone/create_fast_phone_order").
		Send(reqArgs).
		EndStruct(&sercviceid)
	if errs != nil {
		err := fmt.Errorf("chunyu.FastPhoneCreate error: %q", errs)
		return nil, err
	}
	return &sercviceid, nil
}

//追加问题
func AppendProblem(payload cmn.AppendProblemPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	reqArgs := AppendProblemRequest{
		ProblemId: payload.ProblemId,
		Partner:   cmn.Config().GetString("chunyu.partner"),
		Sign:      getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:    payload.UserId,
		Atime:     now,
		Content:   payload.Content,
	}

	var resp ErrorMsgReponse
	_, _, errs := newRequest().Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/problem_content/create").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.AppendProblem error: %q", errs)
		return nil, err
	}
	return &resp, nil
}

//评价问题
func AssessProblem(payload cmn.AssessProblemPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	reqArgs := AssessProblemRequest{
		ProblemId:  payload.ProblemId,
		Content:    payload.Content,
		UserId:     payload.UserId,
		AssessInfo: payload.AssessInfo,
		Partner:    cmn.Config().GetString("chunyu.partner"),
		Sign:       getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
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
	return &resp, nil
}

//删除问题
func DeleteProblem(payload cmn.DeleteProblemPayload) (*ErrorMsgReponse, error) {
	now := time.Now().Unix()

	reqArgs := DeleteProblemRequest{
		ProblemId: payload.ProblemId,
		UserId:    payload.UserId,
		Partner:   cmn.Config().GetString("chunyu.partner"),
		Sign:      getSign(cmn.Config().GetString("chunyu.partnerKey"), strconv.FormatInt(now, 10), payload.UserId),
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
	return &resp, nil
}
