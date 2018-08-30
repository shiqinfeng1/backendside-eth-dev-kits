package v1

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/contracts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
)

//UserTransferETH : 用户转账,非离线签名转账,仅限于后台hd钱包生成维护管理的用户地址
func UserTransferETH(c echo.Context) error {
	p := httpservice.TransferPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证
	if common.UserAuth(p.UserID) == false {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	}

	//签名并发送交易
	txhash, err := eth.TransferEth(p)
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
	if common.UserAuth(p.UserID) == false {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	}

	txhash, err := eth.SendRawTransaction(p)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	return httpservice.JSONReturns(c, txhash)
}

//BuyPoints : 发送离线交易 购买积分
func BuyPoints(c echo.Context) error {
	p := httpservice.BuyPointsysPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证和短信验证码验证
	if common.UserAuth(p.UserID) == false {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	}

	//解析参数
	a, err := math.ParseBig256(p.Amount)
	if err == false {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "invalid transfer amount")
	}
	amount := hexutil.Big(*a)

	//获取铸币的账户
	adminAddress, err2 := accounts.GetadminAddress()
	if err2 != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "no admin. err:"+err2.Error())
	}
	auth, err3 := contracts.GetUserAuth(adminAddress.Hex())
	if err3 != nil {
		common.Logger.Error("GetUserAuth: 15422339579 fail.")
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err3.Error())
	}

	txn, err5 := contracts.PointsBuy(p.ChainType, adminAddress.Hex(), auth, p.Buyer, amount.ToInt().Uint64())
	if err5 != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err5.Error())
	}
	return httpservice.JSONReturns(c, txn.Hash().String())
}

//ConsumePoints : 发送离线交易 消费积分
func ConsumePoints(c echo.Context) error {
	p := httpservice.RawTransactionPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证和短信验证码验证
	if common.UserAuth(p.UserID) == false {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	}

	//解析交易
	transaction, from, err := eth.SignedDataToTransaction(p.SignedData)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	//解析交易执行的输入数据
	_, err = eth.PraseConsumePoints(transaction)
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
	go contracts.PollEventBurn(p.ChainType, txhash, blockNum.ToInt().Uint64(), *transaction.To())

	return httpservice.JSONReturns(c, txhash)
}

//RefundPoints : 发送离线交易 退还积分
func RefundPoints(c echo.Context) error {
	p := httpservice.RawTransactionPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	//TODO: 用户验证和短信验证码验证
	if common.UserAuth(p.UserID) == false {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, "token auth failed")
	}

	//解析交易
	transaction, from, err := eth.SignedDataToTransaction(p.SignedData)
	if err != nil {
		return httpservice.ErrorReturns(c, httpservice.ErrorCode1, err.Error())
	}
	//解析交易执行的输入数据
	_, err = eth.PraseRefundPoints(transaction)
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
	go contracts.PollEventBurn(p.ChainType, txhash, blockNum.ToInt().Uint64(), *transaction.To())

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
