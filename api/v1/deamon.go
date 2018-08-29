package v1

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
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

//BuyPoints : 发送离线交易 购买积分
func BuyPoints(c echo.Context) error {
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

	//for test

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
	// 执行交易之前获取blocknumber,监听事件时从该block开始检查
	blockNum, err := eth.ConnectEthNodeForWeb3(p.ChainType).EthBlockNumber()
	if err != nil {
		common.Logger.Errorf("Failed to EthBlockNumber: %v", err)
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}

	//执行交易
	txhash, err := eth.BuyPointsOffline(p)
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
