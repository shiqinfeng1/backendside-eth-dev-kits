package eth

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
)

//ErrorMsgReponse 消息错误格式
type ErrorMsgReponse struct {
	Error    int32  `json:"error"`     //  错误码 32	not nil 	0:代表成功,其它:代表异常
	ErrorMsg string `json:"error_msg"` //  异常信息	否
}

//UserLoginRequest 春雨用户登录消息格式
type UserLoginRequest struct {
	UserID   string `json:"user_id"` //用户名 <32	not nil	用户唯一标识,合作方定义
	Password string `json:"password"`
	Lon      string `json:"lon"`
	Lat      string `json:"lat"`
	Partner  string `json:"partner"` //合作方标识 len<32	not nil
	Sign     string `json:"sign"`    //签名 <32	not nil
	Atime    int64  `json:"atime"`   //签名时间戳 <64	not nil 	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
}

//UserLoginReponse 春雨用户登录相应消息格式
type UserLoginReponse struct {
	ErrorMsgReponse
}

//PendingPoolParas pendingpool的输入参数
type PendingPoolParas struct {
	ChainType   string
	UserID      string
	TxHash      ethcmn.Hash
	From        ethcmn.Address
	To          ethcmn.Address
	Nonce       uint64
	Description string
}

//PointsParas 参数
type PointsParas struct {
	ChainType      string
	UserID         string
	UserAddress    ethcmn.Address
	TxnType        string
	TxHash         ethcmn.Hash
	PreBalance     uint64
	ExpectBalance  uint64
	IncurredAmount uint64
	CurrentStatus  string
}

//RawData 离线交易参数
type RawData struct {
	SignedData string
	ChainType  string
}
