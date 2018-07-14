package chunyu

type ErrorMsgReponse struct {
	Error    int32  `json:"error"`     //  错误码 32	not nil 	0:代表成功,其它:代表异常
	ErrorMsg string `json:"error_msg"` //  异常信息	否
}

type UserLoginRequest struct {
	UserId   string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Password string `json:"password"`
	Lon      string `json:"lon"`
	Lat      string `json:"lat"`
	Partner  string `json:"partner"` //合作方标识 len<32	not nil
	Sign     string `json:"sign"`    //签名 <32	not nil
	Atime    int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type UserLoginReponse struct {
	ErrorMsgReponse
}

type FreeProblemCreateRequest struct {
	//科室编号,一次查询只能提交一个科室的对应编号  not nil
	ClinicNo string `json:"clinic_no"`
	Partner  string `json:"partner"` //合作方标识 len<32	not nil
	Sign     string `json:"sign"`    //签名 <32	not nil
	UserId   string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime    int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	Content  string `json:"content"` //提问的内容
}

type PaidProblemCreateRequest struct {
	FreeProblemCreateRequest
	PartnerOrderId string `json:"partner_order_id"` //唯一标识本次支付行为
	PayType        string `json:"pay_type"`         //付费方式 二甲医生：qc_hospital_common 三甲医生：qc_hospital_upgrade
}
type OrientedProblemCreateRequest struct {
	//科室编号,一次查询只能提交一个科室的对应编号  not nil
	DoctorIds      string `json:"doctor_ids"`       //购买的医生列表 使用#进行连接多个医生，不能有空格
	Content        string `json:"content"`          //提问的内容
	Partner        string `json:"partner"`          //合作方标识 len<32	not nil
	PartnerOrderId string `json:"partner_order_id"` //唯一标识本次支付行为
	Price          int    `json:"price"`            //价格 not nil	单位为分
	Sign           string `json:"sign"`             //签名 <32	not nil
	UserId         string `json:"user_id"`          //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime          int64  `json:"atime"`            //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)

}
type PaidProblemRefundRequest struct {
	ProblemId int    `json:"problem_id"`
	Partner   string `json:"partner"` //合作方标识 len<32	not nil
	Sign      string `json:"sign"`    //签名 <32	not nil
	UserId    string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime     int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type PaidProblemQueryClinicNoRequest struct {
	Ask     string `json:"ask"`     //首次提问的问题文本
	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	UserId  string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type ClinicDoctorRequest struct {
	//科室编号,一次查询只能提交一个科室的对应编号  not nil
	//'1':妇科,  '2':儿科,  '3':内科,  '4':皮肤性病科,
	//'6':营养科,  '7':骨伤科,  '8':男科,  '9':外科,
	//'11':肿瘤及防治科,  '12':中医科,  '13':口腔颌面科,  '14':耳鼻咽喉科,
	//'15':眼科,  '16':整形美容科,  '17':精神心理科,  '21':产科,
	ClinicNo string `json:"clinic_no"`
	Partner  string `json:"partner"`   //合作方标识 len<32	not nil
	Sign     string `json:"sign"`      //签名 <32	not nil
	UserId   string `json:"user_id"`   //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime    int64  `json:"atime"`     //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	StartNum uint32 `json:"start_num"` //开始数 <32	not nil	用于支持翻页功能，从0开始计数
	Count    uint32 `json:"count"`     //每次取的问题数 <32	not nil	最大200
	Province string `json:"province"`  //省份 <32
	City     string `json:"city"`      //城市 <32
}
type RecommendedDoctorRequest struct {
	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	UserId  string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}

type DoctorInfo struct {
	ClinicName     string `json:"clinic_name"`      //科室名称 not nil
	GoodAt         string `json:"good_at"`          //擅长 not nil
	HospitalName   string `json:"hospital_name"`    //医院名称 not nil
	Image          string `json:"image"`            //医生照片的 URL <200 not nil
	Id             string `json:"id"`               //医生id not nil
	Name           string `json:"name"`             //医生姓名 not nil
	Price          uint32 `json:"price"`            //价格 not nil	单位为分
	PurchaseNum    uint32 `json:"purchase_num"`     //购买数量 not nil
	Title          string `json:"title"`            //职称 not nil
	IsFamousDoctor bool   `json:"is_famous_doctor"` //是否是名医咨询 是	名医咨询10次交互/48h后问题关闭；普通定向问题50次交互/48h后问题关闭
}
type ProblemAndDoctorMap struct {
	DoctorId  string `json:"doctor_id"`  //科室名称 not nil
	ProblemId string `json:"problem_id"` //擅长 not nil
}
type ProblemIDReponse struct {
	ProblemID int32 `json:"problem_id"` //	 问题ID
	ErrorMsgReponse
}
type ClinicDoctorReponse struct {
	Doctors []DoctorInfo `json:"doctors"` //	 医生list not nil
	ErrorMsgReponse
}
type ClinicNoReponse struct {
	ClinicNo int `json:"clinic_no"` //	 科室号
}
type ProblemAndDoctorReponse struct {
	Problems []ProblemAndDoctorMap `json:"problems"` //	 医生list not nil
	ErrorMsgReponse
}
