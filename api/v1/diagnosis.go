package v1

import (
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/chunyuyisheng/service/chunyu"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/guoyi"
)

//医生回复后的通知接口
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

//问题关闭后的通知接口
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

//患者的登录信息同步到第三方平台
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

//众包问题创建
func FreeProblemCreate(c echo.Context) error {
	p := common.FreeProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		problemid, err := chunyu.FreeProblemCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, problemid)
	}

	return nil
}

//众包问题创建
func PaidProblemCreate(c echo.Context) error {
	p := common.PaidProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		problemid, err := chunyu.PaidProblemCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, problemid)
	}

	return nil
}

//众包升级问题未回复的退款
func PaidProblemRefund(c echo.Context) error {
	p := common.PaidProblemRefundPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		problemid, err := chunyu.PaidProblemRefund(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, problemid)
	}

	return nil
}

//众包升级问题查询分配的科室
func PaidProblemQueryClinicNo(c echo.Context) error {
	p := common.PaidProblemQueryClinicNoPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		problemid, err := chunyu.PaidProblemQueryClinicNo(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, problemid)
	}
	return nil
}

//获取指定科室和指定城市的医生列表
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

//获取推荐的医生列表
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

//定向问题创建
func OrientedProblemCreate(c echo.Context) error {
	p := common.OrientedProblemPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		problemid, err := chunyu.OrientedProblemCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, problemid)
	}

	return nil
}

//定向问题的付费退款,注意:用户主动发起付费问题退款，退款只能在医生未回答的情况下才能成功。
func OrientedProblemRefund(c echo.Context) error {
	return PaidProblemRefund(c)
}

//获取指定科室和指定城市的医生列表
func GetEmergencyGraph(c echo.Context) error {
	p := common.EmergencyGraphPayload{}
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

//众包问题创建
func EmergencyGraphCreate(c echo.Context) error {
	p := common.EmergencyGraphPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	if p.Platform == "chunyu" {
		problemid, err := chunyu.EmergencyGraphCreate(p)
		if err != nil {
			return common.BizError1002
		}
		return common.JSONReturns(c, problemid)
	}

	return nil
}
