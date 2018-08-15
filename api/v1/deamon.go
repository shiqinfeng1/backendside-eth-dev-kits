package v1

import (
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
)

//UserTransferETH : 用户转账
func UserTransferETH(c echo.Context) error {
	p := httpservice.TransferPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证
	//
	//

	txhash, err := eth.Transfer(p)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	return httpservice.JSONReturns(c, txhash)
}

//SendRawTransaction : 发送离线交易
func SendRawTransaction(c echo.Context) error {
	p := httpservice.RawTransactionPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证
	//
	//

	txhash, err := eth.SendRawTransaction(p)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	return httpservice.JSONReturns(c, txhash)
}
