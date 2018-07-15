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
type EmergencyGraphCreateRequest struct {
	FreeProblemCreateRequest
	PartnerOrderId string `json:"partner_order_id"` //唯一标识本次支付行为
	Price          int    `json:"price"`            //价格必须与实时查询到的价格一致
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

type FastPhoneOrderRequest struct {
	//科室编号,一次查询只能提交一个科室的对应编号  not nil
	Phone          string `json:"phone"` //购买的医生列表 使用#进行连接多个医生，不能有空格
	ClinicNo       string `json:"clinic_no"`
	Partner        string `json:"partner"`          //合作方标识 len<32	not nil
	PartnerOrderId string `json:"partner_order_id"` //唯一标识本次支付行为
	Sign           string `json:"sign"`             //签名 <32	not nil
	UserId         string `json:"user_id"`          //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime          int64  `json:"atime"`            //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type PaidProblemRefundRequest struct {
	ProblemId int    `json:"problem_id"`
	UserId    string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义

	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
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
type AskHistoryRequest struct {
	Partner  string `json:"partner"`   //合作方标识 len<32	not nil
	Sign     string `json:"sign"`      //签名 <32	not nil
	UserId   string `json:"user_id"`   //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime    int64  `json:"atime"`     //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	StartNum uint32 `json:"start_num"` //开始数 <32	not nil	用于支持翻页功能，从0开始计数
	Count    uint32 `json:"count"`     //每次取的问题数 <32	not nil	最大200
}
type RecommendedDoctorRequest struct {
	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	UserId  string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type DoctorDetailRequest struct {
	DoctorId string `json:"doctor_id"`
	UserId   string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Partner  string `json:"partner"` //合作方标识 len<32	not nil
	Sign     string `json:"sign"`    //签名 <32	not nil
	Atime    int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type ProblemDetailRequest struct {
	ProblemId     int64  `json:"problem_id"`
	UserId        string `json:"user_id"`
	LastContentId uint64 `json:"last_content_id"` //最后一个回复编号,会返回所有大于此编号的回复列表
	Partner       string `json:"partner"`         //合作方标识 len<32	not nil
	Sign          string `json:"sign"`            //签名 <32	not nil
	Atime         int64  `json:"atime"`           //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type EmergencyGraphRequest struct {
	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	UserId  string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	Content string `json:"content"` //提问的内容
}
type FastPhoneInfoRequest struct {
	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	UserId  string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}
type AppendProblemRequest struct {
	ProblemId int    `json:"problem_id"`
	Partner   string `json:"partner"` //合作方标识 len<32	not nil
	Sign      string `json:"sign"`    //签名 <32	not nil
	UserId    string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Atime     int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	Content   string `json:"content"` //提问的内容
}
type AssessProblemRequest struct {
	ProblemId  int    `json:"problem_id"`  //问题编号                                           //
	UserId     string `json:"user_id"`     //用户名 最大长度=32	not nil	用户唯一标识,合作方定义
	Content    string `json:"content"`     //问题内容
	AssessInfo string `json:"assess_info"` //如:'{"level": "best", "tag_keys":["3201", "3102"]}'

	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
	Atime   int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)

}

type DeleteProblemRequest struct {
	ProblemId int    `json:"problem_id"` //问题编号                                           //
	UserId    string `json:"user_id"`    //用户名 最大长度=32	not nil	用户唯一标识,合作方定义

	Partner string `json:"partner"` //合作方标识 len<32	not nil
	Sign    string `json:"sign"`    //签名 <32	not nil
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
type DoctorInfoForHistory struct {
	Id         string `json:"id"`          //医生id not nil
	Name       string `json:"name"`        //医生姓名 not nil
	Image      string `json:"image"`       //医生照片的 URL <200 not nil
	Title      string `json:"title"`       //职称 not nil
	LevelTitle string `json:"level_title"` //带医院级别的医生职称
	Clinic     string `json:"clinic"`      //科室名称 not nil
	Hospital   string `json:"hospital"`    //医院名称 not nil
}
type ContentInfo struct {
	Id            string `json:"id"`              //回复编号
	CreatedTimeMs uint64 `json:"created_time_ms"` //创建问题时间戳

	Type    string `json:"type"`    //p是用户回复,d是医生回复
	Content string `json:"content"` //同问题追问的 content
}
type ClinicInfo struct {
	ClinicName string `json:"clinic_name"` //科室名称 not nil
	ClinicNo   string `json:"clinic_no"`   //'1':妇科, '15':眼科, '21':产科, 'fa' :小儿科，'ha':皮肤科
	Begin      string `json:"begin"`       //服务开始时间
	End        string `json:"end"`         //服务结束时间
	Price      uint32 `json:"price"`       //价格 not nil	单位为元
	Disabled   bool   `json:"disabled"`    //购买数量 not nil
}
type ClinicInfoLite struct {
	ClinicName string `json:"clinic_name"` //科室名称 not nil
	ClinicNo   string `json:"clinic_no"`   //'1':妇科, '15':眼科, '21':产科, 'fa' :小儿科，'ha':皮肤科
	Icon       string `json:"icon"`        //科室对应的图标
}
type ProblemInfo struct {
	Id int32 `json:"id"` //问题id not nil
	/*
		i 初始化,空白问题或未付款问题---空白问题
		n 新问题
		a 已认领---医生认领,等待医生回答
		s 已回复
		c 已关闭---用户待评价
		v 回复已查看---用户看过医生的回复
		p 系统举报---因为含有违禁词等原因被举报
		d 已评价
	*/
	Status        string `json:"status"`          //问题状态
	Price         uint32 `json:"price"`           //价格 not nil	单位为元
	ToDoc         bool   `json:"to_doc"`          //是否是针对医生的定向提问
	Title         string `json:"title"`           //问题标题
	Ask           string `json:"ask"`             //提问内容
	ClinicNo      string `json:"clinic_no"`       //问题所在的科室号
	ClinicName    string `json:"hospital_name"`   //问题所在的科室名字
	IsViewed      bool   `json:"is_viewed"`       //用户是否查看过该问题
	HasAnswer     bool   `json:"has_answer"`      //是否被医生答复
	NeedAssess    bool   `json:"need_assess"`     //问题是否需要被评价
	CreatedTimeMs uint64 `json:"created_time_ms"` //创建时间的毫秒数
	CreatedTime   string `json:"created_time"`    //创建时间:'%Y-%m-%d %H: %M:%S'
	Star          int    `json:"star"`            //问题星级
}
type ProblemAndDoctorMap struct {
	DoctorId  string `json:"doctor_id"`  //科室名称 not nil
	ProblemId string `json:"problem_id"` //擅长 not nil
}
type ProblemIDReponse struct {
	ProblemID int32 `json:"problem_id"` //	 问题ID
	ErrorMsgReponse
}
type FastPhoneOrderReponse struct {
	ServiceId int `json:"service_id"` //	 服务ID
	ErrorMsgReponse
}
type ClinicDoctorReponse struct {
	Doctors []DoctorInfo `json:"doctors"`
	ErrorMsgReponse
}
type AskHistoryReponse struct {
	Problem ProblemInfo          `json:"problem"`
	Doctor  DoctorInfoForHistory `json:"doctor"`
	ErrorMsgReponse
}
type DoctorDetailReponse struct {
	ClinicName     string `json:"clinic_name"`      //科室名称 not nil
	GoodAt         string `json:"good_at"`          //擅长 not nil
	Hospital       string `json:"hospital"`         //医院名称 not nil
	HospitalGrade  string `json:"hospital_grade"`   //医院级别 not nil
	Image          string `json:"image"`            //医生照片的 URL <200 not nil
	Id             string `json:"id"`               //医生id not nil
	Name           string `json:"name"`             //医生姓名 not nil
	Title          string `json:"title"`            //职称 not nil
	Price          uint32 `json:"price"`            //价格 not nil	单位为分
	SolutionScore  uint32 `json:"solution_score"`   //专业度指数
	RecommendRate  string `json:"recommend_rate"`   //推荐指数
	IsFamousDoctor bool   `json:"is_famous_doctor"` //是否是名医咨询 是	名医咨询10次交互/48h后问题关闭；普通定向问题50次交互/48h后问题关闭
	Description    string `json:"description"`      //专家简介
	GoodRate       string `json:"good_rate"`        //好评率
	RewardNum      int    `json:"reward_num"`       //送心意数量
	ReplyNum       int    `json:"reply_num"`        //咨询数量
	FansNum        int    `json:"fans_num"`         //粉丝数量
	ErrorMsgReponse
}
type ProblemDetailReponse struct {
	Problem ProblemInfo          `json:"problem"`
	Doctor  DoctorInfoForHistory `json:"doctor"`
	Content []ContentInfo        `json:"content"`
	ErrorMsgReponse
}
type EmergencyGraphReponse struct {
	ClinicInfo []ClinicInfo `json:"clinic_info"` //	 医生list not nil
	Err        int          `json:"err"`         //0 成功,1 失败
}
type FastPhoneInfoReponse struct {
	ClinicInfo []ClinicInfoLite `json:"clinic_info"` //	 医生list not nil
	Err        int              `json:"err"`         //0 成功,1 失败
}
type ClinicNoReponse struct {
	ClinicNo int `json:"clinic_no"` //	 科室号
}
type ProblemAndDoctorReponse struct {
	Problems []ProblemAndDoctorMap `json:"problems"` //	 医生list not nil
	ErrorMsgReponse
}
