package v1

import (
	"github.com/labstack/echo"
)

//RegisterDiagAPI :注册api
func RegisterDiagAPI(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	apiv1.POST("/get_clinic_doctors", GetClinicDoctors)
	apiv1.POST("/patient_login", PatientLogin)
	apiv1.POST("/chunyu_doctor_response", ChunyuDoctorResponseCallback)
	apiv1.POST("/chunyu_question_close", ChunyuQuestionCloseCallback)
	apiv1.POST("/get_ask_history", GetAskHistory)
	apiv1.POST("/create_oriented_problem", OrientedProblemCreate)
	apiv1.POST("/get_recommended_doctors", GetrecommendedDoctors)
	apiv1.POST("/create_paid_problem", PaidProblemCreate)
	apiv1.POST("/create_free_problem", FreeProblemCreate)

}
