package eth

import (
	"errors"
	"strings"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

//SignedDataToTransaction 解析离线数据
func SignedDataToTransaction(signedData string) (*types.Transaction, *ethcmn.Address, error) {
	// 解析离线交易的交易数据
	transaction := new(types.Transaction)
	if err := rlp.DecodeBytes([]byte(signedData), transaction); err != nil {
		return nil, nil, err
	}
	common.Logger.Debug("\nDecodeRLP transaction:\n", transaction)

	// 恢复得到交易签名者
	from, err := types.Sender(types.HomesteadSigner{}, transaction)
	if err != nil {
		return nil, nil, err
	}
	return transaction, &from, nil
}

func truncationZero(str string) string {
	begin := 0
	for idx, item := range str {
		if item != '0' {
			begin = idx
			break
		}
	}
	return str[begin:]

}

// PraseERC20Transfer generate erc20 txn
// from to value
func PraseERC20Transfer(txn *types.Transaction) (*[]string, error) {

	var functionCall = make([]string, 3)

	if len(txn.Data()) < ERC20TransferLength {
		common.Logger.Debugf("transaction input:%v is not Transfer call.", txn.Data())
		return nil, errors.New("not a erc20 Transfer transaction:length mismatch")
	}
	r := strings.HasPrefix(string(txn.Data()), HexPrefix+ERC20MethodTransfer)
	if r == false {
		return nil, errors.New("not a erc20 Transfer transaction:functionName mismatch")
	}
	// 恢复得到交易签名者
	from, err := types.Sender(types.HomesteadSigner{}, txn)
	if err != nil {
		return nil, err
	}
	functionCall[0] = from.String()
	functionCall[1] = HexPrefix + string(txn.Data()[34:74])

	i256, ok := math.ParseBig256(HexPrefix + truncationZero(string(txn.Data()[74:])))
	if ok {
		functionCall[2] = i256.Text(10)
	}
	return &functionCall, nil
}

// PraseBuyPoints 购买积分
// from to value
func PraseBuyPoints(txn *types.Transaction) (*[]string, error) {

	var functionCall = make([]string, 3)

	if len(txn.Data()) != ERC20TransferLength {
		common.Logger.Debugf("transaction input:%v is not buypoints call.", string(txn.Data()))
		return nil, errors.New("not a buypoints transaction:length mismatch")
	}
	r := strings.HasPrefix(string(txn.Data()), HexPrefix+ERC20MethodTransfer)
	if r == false {
		return nil, errors.New("not a buypoints transaction:functionName mismatch")
	}
	// 恢复得到交易签名者
	from, err := types.Sender(types.HomesteadSigner{}, txn)
	if err != nil {
		return nil, err
	}
	functionCall[0] = from.String()
	functionCall[1] = HexPrefix + string(txn.Data()[34:74])

	i256, ok := math.ParseBig256(HexPrefix + truncationZero(string(txn.Data()[74:])))
	if ok {
		functionCall[2] = i256.Text(10)
	}
	return &functionCall, nil
}
