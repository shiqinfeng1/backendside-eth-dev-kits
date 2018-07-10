package v1

import (
	"github.com/labstack/echo"
)

func RegisterDiagAPI(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	apiv1.POST("/get_clinic_doctors", GetClinicDoctors)

}
