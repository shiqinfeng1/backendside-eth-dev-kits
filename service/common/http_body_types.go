package common

import (
	"github.com/labstack/echo"
)

// PageBody 分页结果
type PageBody struct {
	Current int `json:"current"`
	Total   int `json:"total,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// ReturnBody 返回值封装
type ReturnBody struct {
	Errcode string      `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Page    PageBody    `json:"page"`
}

// ChunyuReturnBody 返回值封装
type ChunyuReturnBody struct {
	Error    int    `json:"error"`
	ErrorMsg string `json:"error_msg"`
}

// PageParams 分页参数，用于mixin到请求对象中
type PageParams struct {
	PerPage uint32 `json:"per_page"`
	Page    uint32 `json:"page" validate:"gte=0,lte=200"`
}

//ChunyuJSONReturns 返回春雨规范的响应格式
func ChunyuJSONReturns(c echo.Context, err int, errMsg string) error {
	returns := &ChunyuReturnBody{
		Error:    err,
		ErrorMsg: errMsg,
	}

	return c.JSON(200, returns)
}

// JSONReturns API返回值的统一封装，直接做json返回。
// `data`为需要返回的数据
// `pages`为翻页数据，不是必须要有。顺序为: page, total, per_page，其中per_page如果不设置则默认为20。
// 如果使用了这个参数，则 page, total必须有
func JSONReturns(c echo.Context, data interface{}, pages ...int) error {
	var page PageBody
	if len(pages) > 0 {
		current := pages[0]
		total := pages[1]
		perPage := 20
		if len(pages) > 2 {
			perPage = pages[2]
		}
		page = PageBody{
			Current: current,
			Total:   total,
			PerPage: perPage,
		}
	}
	returns := &ReturnBody{
		Errcode: ErrorCode0,
		Data:    data,
		Page:    page,
	}

	return c.JSON(200, returns)
}

// ErrorReturns 发生错误的时候的返回值封装
func ErrorReturns(c echo.Context, errcode string, msg string) error {
	returns := &ReturnBody{
		Errcode: errcode,
		Msg:     msg,
		Page:    PageBody{},
	}
	return c.JSON(200, returns)
}

//DoctorInfo 医生信息
type DoctorInfo struct {
	ID            string `json:"id" validate:"required"`                 //医生 ID
	Name          string `json:"name" validate:"required,max=200"`       //医生姓名
	Image         string `json:"image" validate:"max=200"`               //医生头像	医生照片的 url
	Title         string `json:"title" validate:"required,max=32"`       //医生职称
	LevelTitle    string `json:"level_title" validate:"required,max=32"` //带医院级别的医生职称
	Clinic        string `json:"clinic" validate:"max=20"`               //科室名称
	ClinicNo      string `json:"clinic_no" validate:"max=20"`            //科室号
	Hospital      string `json:"hospital" validate:"max=100"`            //医院名字
	HospitalGrade string `json:"hospital_grade" validate:"required"`     //医院级别
	GoodAt        string `json:"good_at" validate:"required"`            //擅长领域（医生回复接口里的医生信息是简版的信息，建议通过医生详情接口获取医生的详细信息）
}

//ChunyuDoctorResponsePayload 春雨回调输入
type ChunyuDoctorResponsePayload struct {
	ProblemID int        `json:"problem_id" validate:"required"`     //问题编号
	UserID    string     `json:"user_id" validate:"required,max=32"` //用户名 用户唯一标识,合作方定义
	Content   string     `json:"content" validate:"required"`        //医生答复内容 可以包含除 patient_meta 之外的三种类型。
	Sign      string     `json:"sign" validate:"required,max=32"`    //签名 将生成方法中user_id换成problem_id,其他不变
	Atime     int64      `json:"atime" validate:"required"`          //签名时间戳	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	IsSummary bool       `json:"is_summary" `                        //是否是医生总结
	Doctor    DoctorInfo `json:"doctor" validate:"required"`
}

//ChunyuQuestionClosePayload 春雨回调输入
type ChunyuQuestionClosePayload struct {
	ProblemID int    `json:"problem_id" validate:"required"`                //问题编号
	UserID    string `json:"user_id" validate:"required,max=32"`            //用户名 用户唯一标识,合作方定义
	Msg       string `json:"msg" validate:"required"`                       //消息内容
	Status    string `json:"status" validate:"required,oneof=close refund"` //close 回答完毕后关闭 refund 问题退款
	Price     int    `json:"price"`                                         //单位为分
	Sign      string `json:"sign" validate:"required,max=32"`               //签名
	Atime     int64  `json:"atime" validate:"required"`                     //签名时间戳
}

//PatientLoginPayload 用户登录参数
type PatientLoginPayload struct {
	Password string `json:"password" validate:"required,max=32"`             //
	UserID   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Lon      string `json:"lon"`                                             //用户地址经度
	Lat      string `json:"lat"`                                             //用户地址维度
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//ClinicDoctorsPayload 查询医生接口参数
type ClinicDoctorsPayload struct {
	ClinicNo string `json:"clinic_no" validate:"required,numeric"` //
	UserID   string `json:"user_id" validate:"required,max=32"`    //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Province string `json:"province" validate:"max=32"`            //省份 <32
	City     string `json:"city" validate:"max=32"`                //城市 <32
	PageParams
}

//AskHistoryPayload 提问历史查询接口参数
type AskHistoryPayload struct {
	UserID string `json:"user_id" validate:"required,max=32"` //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	PageParams
}

//DoctorDetailPayload 查询医生详情接口参数
type DoctorDetailPayload struct {
	DoctorID string `json:"doctor_id" validate:"required"`
	UserID   string `json:"user_id" validate:"required,max=32"`
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"`
}

//ProblemDetailPayload 医生详情
type ProblemDetailPayload struct {
	ProblemID     int64  `json:"problem_id" validate:"required"`
	UserID        string `json:"user_id" validate:"required,max=32"`
	LastContentID uint64 `json:"last_content_id"` //最后一个回复编号,会返回所有大于此编号的回复列表
	Platform      string `json:"platform" validate:"required,oneof=chunyu guoyi"`
}

//RecommendedDoctorsPayload 推荐医生
type RecommendedDoctorsPayload struct {
	Ask    string `json:"ask" validate:"required"`            //
	UserID string `json:"user_id" validate:"required,max=32"` //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
}

//FreeProblemPayload 众包问题
type FreeProblemPayload struct {
	ClinicNo string `json:"clinic_no" validate:"numeric"`                    //首次提问内容文本
	UserID   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content  string `json:"content" validate:"required,max=5120"`            //用户提问内容列表
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//PaidProblemPayload 付费问题
type PaidProblemPayload struct {
	ClinicNo       string `json:"clinic_no" validate:"numeric"`                                              //
	UserID         string `json:"user_id" validate:"required,max=32"`                                        //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content        string `json:"content" validate:"required,max=5120"`                                      //用户提问内容列表
	PartnerOrderID string `json:"partner_order_id" validate:"required"`                                      //唯一标识本次支付行为
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"`                           //问诊平台
	PayType        string `json:"pay_type" validate:"required,oneof=qc_hospital_common qc_hospital_upgrade"` //付费方式 二甲医生： qc_hospital_common 三甲医生： qc_hospital_upgrade
}

//AppendProblemPayload 追问
type AppendProblemPayload struct {
	ProblemID int    `json:"problem_id" validate:"required"`                  //问题编号                                           //
	UserID    string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content   string `json:"content" validate:"required,max=5120"`            //唯一标识本次支付行为
	Platform  string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//OrientedProblemPayload 定向问题创建
type OrientedProblemPayload struct {
	DoctorIDs      string `json:"doctor_ids" validate:"required"`                  //购买的医生列表,使用#进行连接多个医生，不能有空格                                           //
	UserID         string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content        string `json:"content" validate:"required,max=5120"`            //首次提问内容                                  //用户提问内容列表
	PartnerOrderID string `json:"partner_order_id" validate:"required"`            //合作方支付ID
	Price          int32  `json:"price" validate:"required"`                       //订单价格	单位为分
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//PaidProblemRefundPayload 付费问题
type PaidProblemRefundPayload struct {
	UserID    string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	ProblemID int    `json:"problem_id" validate:"required"`                  //用户提问内容列表
	Platform  string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//PaidProblemQueryClinicNoPayload 付费问题查询科室
type PaidProblemQueryClinicNoPayload struct {
	UserID   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Ask      string `json:"ask" validate:"required"`                         //首次提问的问题文本
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//EmergencyGraphPayload 急诊查询
type EmergencyGraphPayload struct {
	UserID   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content  string `json:"content" validate:"required,max=5120"`            //用户提问内容列表
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//EmergencyGraphCreatePayload 急诊创建
type EmergencyGraphCreatePayload struct {
	ClinicNo       string `json:"clinic_no" validate:"required"`                   //必须是春雨开通的科室                                          //
	UserID         string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content        string `json:"content" validate:"required,max=5120"`            //首次提问内容                                  //用户提问内容列表
	PartnerOrderID string `json:"partner_order_id" validate:"required"`            //合作方支付ID
	Price          int    `json:"price" validate:"required"`                       //价格必须与实时查询到的价格一致	单位为元
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//FastPhoneInfoPayload 一对会议电话查询
type FastPhoneInfoPayload struct {
	UserID   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//FastPhoneOrderCreatePayload 一对一电话创建
type FastPhoneOrderCreatePayload struct {
	ClinicNo       string `json:"clinic_no" validate:"required"`        //必须是春雨开通的科室                                          //
	UserID         string `json:"user_id" validate:"required,max=32"`   //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	PartnerOrderID string `json:"partner_order_id" validate:"required"` //合作方支付ID
	Phone          string `json:"phone" validate:"required"`
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//AssessProblemPayload 问题评价
type AssessProblemPayload struct {
	ProblemID  int    `json:"problem_id" validate:"required"`                  //问题编号                                           //
	UserID     string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content    string `json:"content" validate:"required,max=5120"`            //问题内容
	AssessInfo string `json:"assess_info" validate:"required,max=32"`          //如:'{"level": "best", "tag_keys":["3201", "3102"]}'
	Platform   string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

//DeleteProblemPayload 问题删除接口参数
type DeleteProblemPayload struct {
	UserID    string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	ProblemID int    `json:"problem_id" validate:"required"`                  //用户提问内容列表
	Platform  string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}
