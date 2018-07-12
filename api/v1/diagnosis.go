package v1

import (
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/chunyuyisheng/service/chunyu"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/guoyi"
)

func ChunyuDoctorResponseCallback(c echo.Context) error {
	p := common.ChunyuDoctorResponsePayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}

	resp, err := chunyu.DoctorResponseCallback(p)
	if err != nil {
		return common.ChunyuJSONReturns(c, 1, err.Error())
	}
	return common.ChunyuJSONReturns(c, 0, "")

}
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
