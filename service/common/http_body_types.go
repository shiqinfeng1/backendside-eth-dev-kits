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

// ReturnBody 返回值封装
type ChunyuReturnBody struct {
	Error    int    `json:"error"`
	ErrorMsg string `json:"error_msg"`
}

// PageParams 分页参数，用于mixin到请求对象中
type PageParams struct {
	PerPage uint32 `json:"per_page"`
	Page    uint32 `json:"page" validate:"gte=0,lte=200"`
}

func ChunyuJSONReturns(c echo.Context, err int, err_msg string) error {
	returns := &ChunyuReturnBody{
		Error:    err,
		ErrorMsg: err_msg,
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
func ErrorReturns(errcode string, msg string) *ReturnBody {
	return &ReturnBody{
		Errcode: errcode,
		Msg:     msg,
		Page:    PageBody{},
	}
}

type DoctorInfo struct {
	Id            string `json:"id" validate:"required"`                 //医生 ID
	Name          string `json:"name" validate:"required,max=200"`       //医生姓名
	Image         string `json:"image" validate:"max=200"`               //医生头像	医生照片的 url
	Title         string `json:"title" validate:"required,max=32"`       //医生职称
	Level_title   string `json:"level_title" validate:"required,max=32"` //带医院级别的医生职称
	Clinic        string `json:"clinic" validate:"max=20"`               //科室名称
	Clinic_no     string `json:"clinic_no" validate:"max=20"`            //科室号
	Hospital      string `json:"hospital" validate:"max=100"`            //医院名字
	HospitalGrade string `json:"hospital_grade" validate:"required"`     //医院级别
	GoodAt        string `json:"good_at" validate:"required"`            //擅长领域（医生回复接口里的医生信息是简版的信息，建议通过医生详情接口获取医生的详细信息）
}

type ChunyuDoctorResponsePayload struct {
	ProblemId int        `json:"problem_id" validate:"required"`     //问题编号
	UserId    string     `json:"user_id" validate:"required,max=32"` //用户名 用户唯一标识,合作方定义
	Content   string     `json:"content" validate:"required,max=32"` //医生答复内容 可以包含除 patient_meta 之外的三种类型。
	Sign      string     `json:"sign" validate:"required,max=32"`    //签名 将生成方法中user_id换成problem_id,其他不变
	Atime     int64      `json:"atime" validate:"required"`          //签名时间戳	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	IsSummary bool       `json:"is_summary" `                        //是否是医生总结
	Doctor    DoctorInfo `json:"doctor" validate:"required"`
}
type ChunyuQuestionClosePayload struct {
	ProblemId int    `json:"problem_id" validate:"required"`                //问题编号
	UserId    string `json:"user_id" validate:"required,max=32"`            //用户名 用户唯一标识,合作方定义
	Msg       string `json:"msg" validate:"required"`                       //消息内容
	Status    string `json:"status" validate:"required,oneof=close refund"` //close 回答完毕后关闭 refund 问题退款
	Price     int    `json:"price"`                                         //单位为分
	Sign      string `json:"sign" validate:"required,max=32"`               //签名
	Atime     int64  `json:"atime" validate:"required"`                     //签名时间戳
}
type PatientLoginPayload struct {
	Password string `json:"password" validate:"required,max=32"`             //
	UserId   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Lon      string `json:"lon"`                                             //用户地址经度
	Lat      string `json:"lat"`                                             //用户地址维度
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

type ClinicDoctorsPayload struct {
	ClinicNo string `json:"clinic_no" validate:"required,numeric"` //
	UserId   string `json:"user_id" validate:"required,max=32"`    //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Province string `json:"province" validate:"max=32"`            //省份 <32
	City     string `json:"city" validate:"max=32"`                //城市 <32
	PageParams
}
type RecommendedDoctorsPayload struct {
	Ask    string `json:"ask" validate:"required"`            //
	UserId string `json:"user_id" validate:"required,max=32"` //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
}

type FreeProblemPayload struct {
	ClinicNo string `json:"clinic_no" validate:"numeric"`                    //首次提问内容文本
	UserId   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content  string `json:"content" validate:"required,max=5120"`            //用户提问内容列表
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

type PaidProblemPayload struct {
	ClinicNo       string `json:"clinic_no" validate:"numeric"`                                              //
	UserId         string `json:"user_id" validate:"required,max=32"`                                        //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content        string `json:"content" validate:"required,max=5120"`                                      //用户提问内容列表
	PartnerOrderId string `json:"partner_order_id" validate:"required"`                                      //唯一标识本次支付行为
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"`                           //问诊平台
	PayType        string `json:"pay_type" validate:"required,oneof=qc_hospital_common qc_hospital_upgrade"` //付费方式 二甲医生： qc_hospital_common 三甲医生： qc_hospital_upgrade
}
type OrientedProblemPayload struct {
	DoctorIds      string `json:"doctor_ids" validate:"required"`                  //购买的医生列表,使用#进行连接多个医生，不能有空格                                           //
	UserId         string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content        string `json:"content" validate:"required,max=5120"`            //首次提问内容                                  //用户提问内容列表
	PartnerOrderId string `json:"partner_order_id" validate:"required"`            //合作方支付ID
	Price          int    `json:"price" validate:"required"`                       //订单价格	单位为分
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}
type PaidProblemRefundPayload struct {
	UserId    string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	ProblemId int    `json:"problem_id" validate:"required"`                  //用户提问内容列表
	Platform  string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

type PaidProblemQueryClinicNoPayload struct {
	UserId   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Ask      string `json:"ask" validate:"required"`                         //首次提问的问题文本
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}

type EmergencyGraphPayload struct {
	UserId   string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content  string `json:"content" validate:"required,max=5120"`            //用户提问内容列表
	Platform string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}
type EmergencyGraphCreatePayload struct {
	ClinicNo       string `json:"clinic_no" validate:"required"`                   //购买的医生列表,使用#进行连接多个医生，不能有空格                                           //
	UserId         string `json:"user_id" validate:"required,max=32"`              //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content        string `json:"content" validate:"required,max=5120"`            //首次提问内容                                  //用户提问内容列表
	PartnerOrderId string `json:"partner_order_id" validate:"required"`            //合作方支付ID
	Price          int    `json:"price" validate:"required"`                       //价格必须与实时查询到的价格一致	单位为元
	Platform       string `json:"platform" validate:"required,oneof=chunyu guoyi"` //问诊平台
}
