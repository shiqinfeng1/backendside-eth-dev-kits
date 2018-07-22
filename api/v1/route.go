package v1

import (
	"github.com/labstack/echo"
)

//RegisterDiagAPI :注册api
func RegisterDiagAPI(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	apiv1.POST("/patient_login", PatientLogin)
	apiv1.POST("/get_clinic_doctors", GetClinicDoctors)
	apiv1.POST("/get_recommended_doctors", GetrecommendedDoctors)
	apiv1.POST("/get_ask_history", GetAskHistory)
	apiv1.POST("/get_doctor_detail", GetDoctorDetail)
	apiv1.POST("/get_problem_detail", GetProblemDetail)
	apiv1.POST("/get_paid_problem_clinicno", PaidProblemQueryClinicNo)
	apiv1.POST("/get_emergency_graph", GetEmergencyGraph)
	apiv1.POST("/chunyu_doctor_response", ChunyuDoctorResponseCallback)
	apiv1.POST("/chunyu_question_close", ChunyuQuestionCloseCallback)
	apiv1.POST("/create_oriented_problem", OrientedProblemCreate)
	apiv1.POST("/create_paid_problem", PaidProblemCreate)
	apiv1.POST("/create_free_problem", FreeProblemCreate)
	apiv1.POST("/create_emergency_graph", EmergencyGraphCreate)

	apiv1.POST("/refund_oriented_problem", OrientedProblemRefund)
	apiv1.POST("/refund_paid_problem", PaidProblemRefund)
	apiv1.POST("/append_problem", ProblemAppend)
	apiv1.POST("/assess_problem", ProblemAssess)
	apiv1.POST("/delete_problem", ProblemDelete)
	apiv1.POST("/close_problem", ProblemClose)

	apiv1.POST("/upload_question_image", UploadQuestionImage)
	apiv1.POST("/upload_question_audio", UploadQuestionAudio)
	apiv1.POST("/create_free_problem", FreeProblemCreate)
	apiv1.Static("/images", "images")
	apiv1.Static("/audio", "audio")
}
