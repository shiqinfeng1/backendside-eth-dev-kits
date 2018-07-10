package v1

import (
	"errors"

	"github.com/shiqinfeng1/chunyuyisheng/service/chunyu"
	"github.com/shiqinfeng1/chunyuyisheng/service/guoyi"
)

func GetClinicDoctors(c echo.Context) error {
	chunyu.GetClinicDoctors()
	guoyi.GetClinicDoctors()

	return errors.New("to be continue.")
}
