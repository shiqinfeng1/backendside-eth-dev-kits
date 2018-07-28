package v1

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/shiqinfeng1/chunyuyisheng/service/chunyu"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/guoyi"
)

//ChunyuDoctorResponseCallback :医生回复后的通知接口
func ChunyuDoctorResponseCallback(c echo.Context) error {
	p := common.ChunyuDoctorResponsePayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}

	err := chunyu.DoctorResponseCallback(p)
	if err != nil {
		return common.ChunyuJSONReturns(c, 1, err.Error())
	}
	return common.ChunyuJSONReturns(c, 0, "")
}

//ChunyuQuestionCloseCallback :问题关闭后的通知接口
func ChunyuQuestionCloseCallback(c echo.Context) error {
	p := common.ChunyuQuestionClosePayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}

	err := chunyu.QuestionCloseCallback(p)
	if err != nil {
		return common.ChunyuJSONReturns(c, 1, err.Error())
	}
	return common.ChunyuJSONReturns(c, 0, "")
}

//PatientLogin : 患者的登录信息同步到第三方平台
func PatientLogin(c echo.Context) error {
	p := common.PatientLoginPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}

	resp, err := chunyu.UserLogin(p)
	if err != nil {
		return common.BizError1002
	}
	return common.JSONReturns(c, resp)

}

//FreeProblemCreate :众包问题创建
func FreeProblemCreate(c echo.Context) error {
	p := common.FreeProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.FreeProblemCreate(p)
		if err != nil {
			return common.ErrorReturns(c, common.ErrorCode1, err.Error())
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//PaidProblemCreate :众包问题创建
func PaidProblemCreate(c echo.Context) error {
	p := common.PaidProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.PaidProblemCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//PaidProblemRefund : 众包升级问题未回复的退款
func PaidProblemRefund(c echo.Context) error {
	p := common.PaidProblemRefundPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.PaidProblemRefund(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//PaidProblemQueryClinicNo :众包升级问题查询分配的科室
func PaidProblemQueryClinicNo(c echo.Context) error {
	p := common.PaidProblemQueryClinicNoPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.PaidProblemQueryClinicNo(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}
	return nil
}

//GetClinicDoctors :获取指定科室和指定城市的医生列表
func GetClinicDoctors(c echo.Context) error {
	p := common.ClinicDoctorsPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.PerPage == 0 {
		p.PerPage = 20
	}
	doctorlist, err := chunyu.GetClinicDoctors(p)
	if err != nil {
		return common.BizError1002
	}
	//to be continue...
	guoyi.GetClinicDoctors()

	return common.JSONReturns(c, doctorlist, p.Page, 0, p.PerPage)
}

//GetAskHistory :获取指定科室和指定城市的医生列表
func GetAskHistory(c echo.Context) error {
	p := common.AskHistoryPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.PerPage == 0 {
		p.PerPage = 20
	}
	asklist, err := chunyu.GetAskHistory(p)
	if err != nil {
		return common.BizError1002
	}
	return common.JSONReturns(c, asklist, p.Page, 0, p.PerPage)
}

//GetrecommendedDoctors :获取推荐的医生列表
func GetrecommendedDoctors(c echo.Context) error {
	p := common.RecommendedDoctorsPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	doctorlist, err := chunyu.GetRecommendedDoctors(p)
	if err != nil {
		return common.BizError1002
	}
	//to be continue...
	guoyi.GetClinicDoctors()

	return common.JSONReturns(c, doctorlist)
}

//GetDoctorDetail :获取推荐的医生列表
func GetDoctorDetail(c echo.Context) error {
	p := common.DoctorDetailPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	doctorlist, err := chunyu.GetDoctorDetail(p)
	if err != nil {
		return common.BizError1002
	}
	//to be continue...
	guoyi.GetClinicDoctors()

	return common.JSONReturns(c, doctorlist)
}

//GetProblemDetail :获取推荐的医生列表
func GetProblemDetail(c echo.Context) error {
	p := common.ProblemDetailPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	doctorlist, err := chunyu.GetProblemDetail(p)
	if err != nil {
		return common.BizError1002
	}
	if doctorlist.Error > 0 {
		return common.JSONReturns(c, &chunyu.ErrorMsgReponse{Error: doctorlist.Error, ErrorMsg: doctorlist.ErrorMsg})
	}
	return common.JSONReturns(c, doctorlist)
}

//OrientedProblemCreate  : 定向问题创建
func OrientedProblemCreate(c echo.Context) error {
	p := common.OrientedProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.OrientedProblemCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//OrientedProblemRefund :定向问题的付费退款,注意:用户主动发起付费问题退款，退款只能在医生未回答的情况下才能成功。
func OrientedProblemRefund(c echo.Context) error {
	return PaidProblemRefund(c)
}

//GetEmergencyGraph :获取指定科室和指定城市的医生列表
func GetEmergencyGraph(c echo.Context) error {
	p := common.EmergencyGraphPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	doctorlist, err := chunyu.GetEmergencyGraph(p)
	if err != nil {
		return common.BizError1002
	}
	//to be continue...
	guoyi.GetClinicDoctors()

	return common.JSONReturns(c, doctorlist)
}

//EmergencyGraphCreate :众包问题创建
func EmergencyGraphCreate(c echo.Context) error {
	p := common.EmergencyGraphCreatePayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.EmergencyGraphCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//GetFastPhoneInfo :插叙急诊图文信息
func GetFastPhoneInfo(c echo.Context) error {
	p := common.FastPhoneInfoPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.GetFastPhoneInfo(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//FastPhoneOrderCreate :插叙急诊图文信息
func FastPhoneOrderCreate(c echo.Context) error {
	p := common.FastPhoneOrderCreatePayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.FastPhoneOrderCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//ProblemAppend :问题追加
func ProblemAppend(c echo.Context) error {
	p := common.AppendProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.AppendProblem(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//ProblemAssess :问题追加
func ProblemAssess(c echo.Context) error {
	p := common.AssessProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.AssessProblem(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//ProblemDelete :问题删除
func ProblemDelete(c echo.Context) error {
	p := common.DeleteProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.DeleteProblem(p, "is_deleted")
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}
	return nil
}

//ProblemClose :问题关闭
func ProblemClose(c echo.Context) error {
	p := common.DeleteProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		ProblemID, err := chunyu.DeleteProblem(p, "is_closed")
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}
	return nil
}

//检查目录是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(filename + " not exist\n")
		exist = false
	}
	return exist
}

//保存文件
func saveMediaFile(dir, userid, index string, src multipart.File) map[string]interface{} {
	result := make(map[string]interface{})
	year := strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	day := strconv.Itoa(time.Now().Day())

	//检查文件存储路径是否存在
	boolImagesexist := checkFileIsExist(dir)
	if !boolImagesexist {
		err1 := os.Mkdir(dir, os.ModePerm) //创建文件夹
		if err1 != nil {
			common.Logger.Error(err1)
			result["result"] = "file create fail"
			return result
		}
	}
	dir = dir + "/" + year
	boolYearexist := checkFileIsExist(dir)
	if !boolYearexist {
		err1 := os.Mkdir(dir, os.ModePerm) //创建文件夹
		if err1 != nil {
			common.Logger.Error(err1)
			result["result"] = "file create fail"
			return result
		}
	}
	dir = dir + "/" + month
	boolMonthexist := checkFileIsExist(dir)
	if !boolMonthexist {
		err1 := os.Mkdir(dir, os.ModePerm) //创建文件夹
		if err1 != nil {
			common.Logger.Error(err1)
			result["result"] = "file create fail"
			return result
		}
	}
	dir = dir + "/" + day
	boolDayexist := checkFileIsExist(dir)
	if !boolDayexist {
		err1 := os.Mkdir(dir, os.ModePerm) //创建文件夹
		if err1 != nil {
			common.Logger.Error(err1)
			result["result"] = "file create fail"
			return result
		}
	}
	dir = dir + "/" + userid + "_" + index + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	f, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		common.Logger.Error(err)
		result["result"] = "file open fail"
		return result
	}
	defer f.Close()
	if _, err = io.Copy(f, src); err != nil {
		result["result"] = err.Error()
		return result
	}
	common.Logger.Printf("file:%s save ok.\n", dir)
	result["result"] = "ok"
	result["url"] = dir[2:]
	return result
}

//UploadQuestionImage 上传图片
func UploadQuestionImage(c echo.Context) error {

	//sessionid := c.FormValue("sessionid")
	index := c.FormValue("index")   //文件索引
	userid := c.FormValue("userid") //用户名
	file, err := c.FormFile("file") //文件内容
	if err != nil {
		return common.BizError1002
	}
	//读取上传文件数据
	src, err := file.Open()
	if err != nil {
		return common.BizError1002
	}
	defer src.Close()

	//TODO:用户和会话合法性验证

	result := saveMediaFile("./images", userid, index, src)
	return common.JSONReturns(c, result)
}

//UploadQuestionAudio 上传图片
func UploadQuestionAudio(c echo.Context) error {

	//sessionid := c.FormValue("sessionid")
	index := c.FormValue("index")   //文件索引
	userid := c.FormValue("userid") //用户名
	file, err := c.FormFile("file") //文件内容
	if err != nil {
		return common.BizError1002
	}
	//读取上传文件数据
	src, err := file.Open()
	if err != nil {
		return common.BizError1002
	}
	defer src.Close()

	//TODO:用户和会话合法性验证

	result := saveMediaFile("./audio", userid, index, src)
	return common.JSONReturns(c, result)
}
