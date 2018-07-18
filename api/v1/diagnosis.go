package v1

import (
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

	if p.Platform == "chunyu" {
		resp, err := chunyu.UserLogin(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, resp)
	}
	return nil
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
	doctorlist, err := chunyu.GetClinicDoctors(p)
	if err != nil {
		return common.BizError1002
	}
	//to be continue...
	guoyi.GetClinicDoctors()

	return common.JSONReturns(c, doctorlist)
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
	asklist, err := chunyu.GetAskHistory(p)
	if err != nil {
		return common.BizError1002
	}
	return common.JSONReturns(c, asklist)
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
		ProblemID, err := chunyu.DeleteProblem(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, ProblemID)
	}

	return nil
}

//ProblemClose :问题关闭
func ProblemClose(c echo.Context) error {
	return ProblemDelete(c)
}
