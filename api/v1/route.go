package v1

import (
	"github.com/labstack/echo"
)

//RegisterDevKitsAPI :注册api
func RegisterDevKitsAPI(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	apiv1.POST("/transfer_eth", UserTransferETH)
	apiv1.Static("/images", "images")
}
