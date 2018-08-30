package v1

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/contracts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
)

//UserTransferETH : 用户转账,限于后台管理的用户以太地址
func UserTransferETH(c echo.Context) error {
	p := httpservice.TransferPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证
	// if common.UserAuth(p.UserID) == false {
	// 	return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	// }

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
	// if common.UserAuth(p.UserID) == false {
	// 	return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	// }

	txhash, err := eth.SendRawTransaction(p)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	return httpservice.JSONReturns(c, txhash)
}

//ConsumePoints : 发送离线交易 购买积分, 该交易需要管理员账户签署
func ConsumePoints(c echo.Context) error {
	p := httpservice.RawTransactionPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证
	// if common.UserAuth(p.UserID) == false {
	// 	return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	// }

	//检查购买积分用户的合法性
	_, err2 := accounts.GetUserAddress(p.UserID)
	if err2 != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "no such userID: "+p.UserID+". err:"+err2.Error())
	}
	//解析交易
	transaction, from, err := eth.SignedDataToTransaction(p.SignedData)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	//解析交易执行的输入数据
	_, err = eth.PraseBuyPoints(transaction)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}

	//执行交易
	var raw = &eth.RawData{SignedData: p.SignedData, ChainType: p.ChainType}
	txhash, blockNum, err := eth.SendRawTxn(raw)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}

	//加入交易监听队列, 如果交易上链,会同步更新数据库
	var para = &eth.PendingPoolParas{
		ChainType: p.ChainType,
		UserID:    p.UserID,
		TxHash:    ethcmn.HexToHash(txhash),
		From:      *from,
		To:        *transaction.To(),
		Nonce:     transaction.Nonce(),
	}
	eth.AppendToPendingPool(para)
	//等待上链,如果执行成功, 则捕获buy事件,保存到数据库
	go contracts.PollEventMint(p.ChainType, txhash, blockNum.ToInt().Uint64(), *transaction.To())

	return httpservice.JSONReturns(c, txhash)
}

//QueryTxnMined : 查询交易状态
func QueryTxnMined(c echo.Context) error {
	p := httpservice.QueryTransactionPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	resp, err := eth.TransactionIsMined(p)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	return httpservice.JSONReturns(c, resp)
}
