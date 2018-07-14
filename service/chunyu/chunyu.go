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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/login").
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/free_problem/create").
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/create_paid_problem").
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/refund").
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/get_problem_clinic_no").
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/get_clinic_doctors").
		Send(reqArgs).
		EndStruct(&doctorsList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetClinicDoctors error: %q", errs)
		return nil, err
	}
	return &doctorsList, nil
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/get_recommended_doctors").
		Send(reqArgs).
		EndStruct(&doctorsList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetRecommendedDoctors error: %q", errs)
		return nil, err
	}
	return &doctorsList, nil
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
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))
	request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/problem/create_oriented_problem").
		Send(reqArgs).
		EndStruct(&Problemlist)
	if errs != nil {
		err := fmt.Errorf("chunyu.OrientedProblemCreate error: %q", errs)
		return nil, err
	}
	return &Problemlist, nil
}
