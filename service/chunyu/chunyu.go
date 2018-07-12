package chunyu

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/parnurzeal/gorequest"
	cmn "github.com/shiqinfeng1/chunyuyisheng/service/common"
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

func DoctorResponseCallback(payload cmn.ChunyuDoctorResponsePayload) {

}
func UserLogin(payload cmn.PatientLoginPayload) (*UserLoginReponse, error) {
	now := time.Now().Unix()

	reqArgs := UserLoginRequest{
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partner_key"), strconv.FormatInt(now, 10), payload.UserId),
		UserId:   payload.UserId,
		Atime:    now,
		Lon:      payload.Lon,
		Lat:      payload.Lat,
		Password: payload.Password,
	}
	var resp UserLoginReponse
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("chunyu.debug"))

	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/login").
		Send(reqArgs).
		EndStruct(&resp)
	if errs != nil {
		err := fmt.Errorf("chunyu.UserLogin error: %q", errs)
		return nil, err
	}
	return &resp, nil

}

func GetClinicDoctors(payload cmn.ClinicDoctorsPayload) (*ClinicDoctorReponse, error) {
	now := time.Now().Unix()

	reqArgs := ClinicDoctorRequest{
		ClinicNo: payload.ClinicNo,
		Partner:  cmn.Config().GetString("chunyu.partner"),
		Sign:     getSign(cmn.Config().GetString("chunyu.partner_key"), strconv.FormatInt(now, 10), payload.UserId),
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

	_, _, errs := request.Post(cmn.Config().GetString("chunyu.domain") + "/cooperation/server/doctor/get_clinic_doctors").
		Send(reqArgs).
		EndStruct(&doctorsList)
	if errs != nil {
		err := fmt.Errorf("chunyu.GetClinicDoctors error: %q", errs)
		return nil, err
	}
	return &doctorsList, nil

}
