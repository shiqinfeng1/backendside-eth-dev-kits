package common

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TransactionRequest 交易请求
type TransactionRequest struct {
	From     ethcmn.Address `json:"from"`
	To       ethcmn.Address `json:"to"`
	Gas      *hexutil.Big   `json:"gas"`
	GasPrice *hexutil.Big   `json:"gasPrice"`
	Value    *hexutil.Big   `json:"value"`
	Data     hexutil.Bytes  `json:"data"`
	Nonce    *hexutil.Big   `json:"nonce"`
}
