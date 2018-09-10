package common

import (
	"sync"
	"time"
)

type uvTimeout struct {
	UserID      string
	Expires     int
	AccessToken string
}

//UserVerifyReponse 春雨用户验证有效性相应消息格式
type UserVerifyReponse struct {
	AccessToken interface{} `json:"accessToken"`
	Expires     int         `json:"expires"`
}

var userVerifyTimeoutChan = make(chan *uvTimeout)
var userVerifyTimeout = make(map[string]uvTimeout)
var userVerifyTimeoutLock sync.Mutex

func init() {
	go func() {
		for arg := range userVerifyTimeoutChan {
			func() {
				userVerifyTimeoutLock.Lock()
				defer userVerifyTimeoutLock.Unlock()
				Logger.Debug("user verify timeout: %s:%s.\n", arg.UserID, userVerifyTimeout[arg.UserID].AccessToken)
				delete(userVerifyTimeout, arg.UserID)

			}()
		}
	}()
}

//UserVerify :验证用户的有效性, 支持本地缓存,超时失效
func UserVerify(userID, key string) bool {

	//检查本地缓存是否存在该userID
	userVerifyTimeoutLock.Lock()
	if v, ok := userVerifyTimeout[userID]; ok {
		userVerifyTimeoutLock.Unlock()
		Logger.Debug("user verify ok from local: %s AccessToken: %s", v.UserID, v.AccessToken)
		return true
	}
	userVerifyTimeoutLock.Unlock()

	//本地缓存不存在,重新查询用户有效性
	var resp UserVerifyReponse
	url := "/appapi/token.html?appid=" + userID + "&appsecret=" + key
	_, _, errs := newRequest().Get(
		Config().GetString("common.userVerifyDomain") + url).
		EndStruct(&resp)
	if errs != nil {
		Logger.Debug(" user Verify error: %q", errs)
		return false
	}
	switch v := resp.AccessToken.(type) {
	case bool:
		var s bool
		s = v
		if s == false {
			Logger.Debug(" Verify user: %s : false.", userID)
			return false
		}
	case string:
		//为了保证精确度, 如果用户有效期时间不足10秒,不在本地缓存, 每次都需要实时查询
		if resp.Expires <= 10 {
			Logger.Debug(" Verify user: %s : less than 10s.", userID)
			return true
		}
		var accessToken string
		accessToken = v
		vc := &uvTimeout{UserID: userID, Expires: resp.Expires, AccessToken: accessToken}
		userVerifyTimeoutLock.Lock()
		userVerifyTimeout[userID] = *vc
		userVerifyTimeoutLock.Unlock()

		go func(vc *uvTimeout) {
			for {
				//只缓存比有效期少10秒的本地缓存
				dur := (time.Duration(resp.Expires) - 10) * time.Second
				Logger.Debug(" Verify user %s : true. duration: %s", userID, dur.String())
				<-time.After(dur) //等待超时
				userVerifyTimeoutChan <- vc
				break
			}
		}(vc)
		Logger.Debug(" Verify user: %s : true.", userID)
		return true
	}
	Logger.Debug(" Verify user: %s : false. response error.", userID)
	return false
}
