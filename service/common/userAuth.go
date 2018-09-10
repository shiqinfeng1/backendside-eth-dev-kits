package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shiqinfeng1/gorequest"
)

//ResponseCodeOk The response code which stands for a sms is sent successfully.
const ResponseCodeOk = "OK"

//AliyunResponse @see https://help.aliyun.com/document_detail/55284.html#出参列表
// The Response of sending sms API.
type AliyunResponse struct {
	// The raw response from server.
	RawResponse []byte `json:"-"`
	/* Response body */
	RequestID string `json:"RequestID"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	BizID     string `json:"BizID"`
}

//IsSuccessful 检查返回结果是否成功
func (m *AliyunResponse) IsSuccessful() bool {
	return m.Code == ResponseCodeOk
}

//VCTimeout Verify Code 超时管理结构
type VCTimeout struct {
	id      string
	timeout time.Time //验证码预计超时时间
	vcode   string
}

//UserAuthReponse 春雨用户验证有效性相应消息格式
type UserAuthReponse struct {
	Status int `json:"status"`
}

//User 标志一个用户
type User struct {
	UserID string `json:"user_id" form:"user_id"`
}

//VerifyCodeManage 验证码管理
type VerifyCodeManage struct {
	timeoutMap  map[string]VCTimeout
	timeoutLock sync.Mutex
}

const sortQueryStringFmt string = "AccessKeyID=%s" +
	"&Action=SendSms" +
	"&Format=JSON" +
	"&OutId=123" +
	"&PhoneNumbers=%s" +
	"&RegionId=cn-hangzhou" +
	"&SignName=%s" +
	"&SignatureMethod=HMAC-SHA1" +
	"&SignatureNonce=%s" +
	"&SignatureVersion=1.0" +
	"&TemplateCode=%s" +
	"&TemplateParam=%s" +
	"&Timestamp=%s" +
	"&Version=2017-05-25"

var vcManage VerifyCodeManage

//UpdateTimeout 更新验证码及超时时间
func (vcm *VerifyCodeManage) UpdateTimeout(userID, vcode string) {
	vc := &VCTimeout{
		id:      userID,
		timeout: time.Now().Add(time.Duration(Config().GetInt64("common.verifycodetimeout"))),
		vcode:   vcode,
	}

	vcm.timeoutLock.Lock()
	vcm.timeoutMap[userID] = *vc //注意:这里会覆盖旧验证码,不管旧验证码是否超时
	vcm.timeoutLock.Unlock()
	Logger.Debugf(" %s: Verify Code=%s ExpectTimeout At %s.\n", userID, vcm.timeoutMap[userID], vc.timeout)

}
func newRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	request.SetDebug(Config().GetBool("common.debug"))
	//request.SetLogger(Logger)
	return request
}

//UserAuth :验证用户的有效性.通过独立的用户管理服务验证用户合法性
func UserAuth(userID, vcode string) bool {

	var resp UserAuthReponse
	url := "/appapi/getaccesstokenvalid.html?accessToken=" + userID
	_, _, errs := newRequest().Get(
		Config().GetString("common.userVerifyDomain") + url).
		EndStruct(&resp)
	if errs != nil {
		Logger.Error(" user Auth error:", errs)
		return false
	}
	if resp.Status == 1 {
		return true
	}

	if vcode != "" {
		valid, err := IsVerifyCodeValid(userID, vcode)
		if err != nil {
			Logger.Errorf("UserID: %s VerifyCode: %s Invalid:%v", userID, vcode, err)
			return false
		}
		Logger.Debugf("UserID: %s VerifyCode: %s valid: %v", userID, vcode)
		return valid
	}
	return false
}

//NewVerifyCodeManage 验证码超时管理
func NewVerifyCodeManage() {
	vcManage.timeoutMap = make(map[string]VCTimeout)
}

func encodeLocal(encodeStr string) string {
	urlencode := url.QueryEscape(encodeStr)
	urlencode = strings.Replace(urlencode, "+", "%%20", -1)
	urlencode = strings.Replace(urlencode, "*", "%2A", -1)
	urlencode = strings.Replace(urlencode, "%%7E", "~", -1)
	urlencode = strings.Replace(urlencode, "/", "%%2F", -1)
	return urlencode
}

//SendByAliyun 通过阿里云服务发送验证码
func SendByAliyun(phone, vCode string) error {
	const accessSecret string = "xxxxx&" // 阿里云 accessSecret 注意这个地方要添加一个 &
	AccessKeyID := "xxxx"                // 自己的阿里云 AccessKeyID
	PhoneNumbers := phone                // 发送目标的手机号
	SignName := url.QueryEscape("链创中医")
	SignatureNonce, _ := uuid.NewV4()
	TemplateCode := "SMS_xxx"                                        //短信模板ID，需要到阿里云帐号上申请，通过后会生成ID (模板示例："亲，你的验证码是${code},
	TemplateParam := url.QueryEscape("{\"code\":\"" + vCode + "\"}") //传入模板的参数(参数示例："{\"code\":\"1234\"}" )
	Timestamp := url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))

	sortQueryString := fmt.Sprintf(sortQueryStringFmt,
		AccessKeyID,
		PhoneNumbers,
		SignName,
		SignatureNonce,
		TemplateCode,
		TemplateParam,
		Timestamp,
	)

	mac := hmac.New(sha1.New, []byte(accessSecret))
	mac.Write([]byte(fmt.Sprintf("GET&%%2F&%s", encodeLocal(sortQueryString))))
	signture := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signture = encodeLocal(signture)

	var resp AliyunResponse
	request := fmt.Sprintf("http://dysmsapi.aliyuncs.com/?Signature=%s&%s\n", signture, sortQueryString)
	_, _, errs := newRequest().Get(request).EndStruct(&resp)
	if errs != nil {
		Logger.Errorf("Verify Code Send By Aliyun Fail:%v", errs)
		return errs[0] //取第一个错误
	}
	return nil
}

//SendVerifyCode 发送验证码
func SendVerifyCode(userID string) error {

	//生成验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//发送验证码
	if err := SendByAliyun(userID, vcode); err != nil {
		return err
	}

	//记录验证码超时
	vcManage.UpdateTimeout(userID, vcode)

	return nil
}

//IsVerifyCodeValid 验证验证码
func IsVerifyCodeValid(userID, vcode string) (bool, error) {
	vcManage.timeoutLock.Lock()
	defer vcManage.timeoutLock.Unlock()
	if _, ok := vcManage.timeoutMap[userID]; !ok { //不存在
		return false, fmt.Errorf("No Such UserID:%s", userID)
	}

	if vcManage.timeoutMap[userID].vcode != vcode {
		return false, fmt.Errorf("UserID:%s vcode:%s Invalid", userID, vcode)
	}

	if vcManage.timeoutMap[userID].timeout.After(time.Now()) {
		return false, fmt.Errorf("UserID:%s vcode:%s Is Timeout", userID, vcode)
	}
	return true, nil
}
