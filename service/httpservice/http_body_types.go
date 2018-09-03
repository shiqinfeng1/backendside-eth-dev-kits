package httpservice

import (
	"github.com/labstack/echo"
)

// PageBody 分页结果
type PageBody struct {
	Current int `json:"current"`
	Total   int `json:"total,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// ReturnBody 返回值封装
type ReturnBody struct {
	Errcode string      `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Page    PageBody    `json:"page"`
}

// PageParams 分页参数，用于mixin到请求对象中
type PageParams struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page" validate:"gte=0,lte=200"`
}

// JSONReturns API返回值的统一封装，直接做json返回。
// `data`为需要返回的数据
// `pages`为翻页数据，不是必须要有。顺序为: page, total, per_page，其中per_page如果不设置则默认为20。
// 如果使用了这个参数，则 page, total必须有
func JSONReturns(c echo.Context, data interface{}, pages ...int) error {
	var page PageBody
	if len(pages) > 0 {
		current := pages[0]
		total := pages[1]
		perPage := 20
		if len(pages) > 2 {
			perPage = pages[2]
		}
		page = PageBody{
			Current: current,
			Total:   total,
			PerPage: perPage,
		}
	}
	returns := &ReturnBody{
		Errcode: ErrorCode0,
		Data:    data,
		Page:    page,
	}

	return c.JSON(200, returns)
}

// ErrorReturns 发生错误的时候的返回值封装
func ErrorReturns(c echo.Context, errcode string, msg string) error {
	returns := &ReturnBody{
		Errcode: errcode,
		Msg:     msg,
	}
	return c.JSON(200, returns)
}

// ErrorReturnsStruct 发生错误的时候的返回值封装
func ErrorReturnsStruct(c echo.Context, errcode string, msg string) *ReturnBody {
	returns := &ReturnBody{
		Errcode: errcode,
		Msg:     msg,
		Page:    PageBody{},
	}
	return returns
}

//UserAuthPayload 转账原生币参数
type UserAuthPayload struct {
	Sign       string `json:"sign" validate:"max=32"`    //签名
	Atime      int64  `json:"atime" validate:"required"` //签名时间戳	当前UNIX TIMESTAMP签名时间戳 (如:137322417)
	VerifyCode string `json:"verify_code" validate:"required"`
	UserID     string `json:"user_id" validate:"required"`
	ChainType  string `json:"chain_type" validate:"required,oneof=ethereum poa"`
}

//TransferPayload 转账原生币参数
type TransferPayload struct {
	UserAuthPayload
	Amount string `json:"amount" validate:"required"`
}

//BuyPointsysPayload 转账原生币参数
type BuyPointsysPayload struct {
	UserAuthPayload
	Amount string `json:"amount" validate:"required"`
	Buyer  string `json:"buyer" validate:"required"`
}

//QueryPointsPayload 转账原生币参数
type QueryPointsPayload struct {
	UserAuthPayload
	Address string `json:"address" validate:"required"`
}

//RawTransactionPayload 离线交易参数
type RawTransactionPayload struct {
	UserAuthPayload
	SignedData string `json:"signed_data" validate:"required"`
}

//BuyPointsPayload 购买交易参数
type BuyPointsPayload struct {
	UserAuthPayload
	Amount string `json:"amount" validate:"required"`
}

//QueryTransactionPayload 离线交易参数
type QueryTransactionPayload struct {
	UserAuthPayload
	TxHash string `json:"tx_hash" validate:"required"`
}

//QueryTransactionReponse 返回交易状态结构
type QueryTransactionReponse struct {
	Mined      bool   `json:"mined"`
	Success    bool   `json:"success"`
	MinedBlock uint64 `json:"mined_block"`
	Comfired   int    `json:"comfired"`
}
