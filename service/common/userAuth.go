package common

import (
	"github.com/shiqinfeng1/gorequest"
)

//UserAuthReponse 春雨用户验证有效性相应消息格式
type UserAuthReponse struct {
	Status int `json:"status"`
}

func newRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	request.SetDebug(Config().GetBool("common.debug"))
	request.SetLogger(Logger)
	return request
}

//UserAuth :验证用户的有效性
func UserAuth(token string) bool {

	var resp UserAuthReponse
	url := "/appapi/getaccesstokenvalid.html?accessToken=" + token
	_, _, errs := newRequest().Get(
		Config().GetString("common.userVerifyDomain") + url).
		EndStruct(&resp)
	if errs != nil {
		Logger.Debug(" user Auth error:", errs)
		return false
	}
	if resp.Status == 1 {
		return true
	}
	Logger.Debug(" Auth token: " + token + ": false. ")
	return true
}
