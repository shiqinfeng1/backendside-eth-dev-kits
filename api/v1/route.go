package v1

import (
	"github.com/labstack/echo"
)

//RegisterDevKitsAPI :注册api
func RegisterDevKitsAPI(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	apiv1.POST("/transfer_eth", UserTransferETH)
	apiv1.POST("/send_raw_transaction", SendRawTransaction)
	apiv1.POST("/buy_points", BuyPoints)
	//apiv1.POST("/consume_points", ConsumePoints)
	//apiv1.POST("/refund_points", RefundPoints)
	apiv1.POST("/query_txn_mined", QueryTxnMined)
	apiv1.Static("/images", "images")
}
