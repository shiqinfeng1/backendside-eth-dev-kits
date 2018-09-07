package eth

import (
	"errors"
	"strings"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

func hexToBytes(signedHex string) ([]byte, error) {
	//如果传入的是16进制数字字符串,需要先转换为字符数组
	var signed []byte
	if signedHex[:2] == "0x" {
		signedData := new(hexutil.Bytes)
		if err := signedData.UnmarshalText([]byte(signedHex)); err != nil {
			common.Logger.Errorf("@@@@ Convert Signed Data:%s to Bytes Fail.", signedHex)
			return signed, err
		}
		signed = []byte(*signedData)
		common.Logger.Debugf("@@@@ Convert Signed Data:%s to Bytes.", signedHex)
		return signed, nil
	}
	common.Logger.Errorf("@@@@ SignedHex:%s Is Not Start With '0x'.", signedHex)
	return signed, errors.New("SignedHex Is Not Start With '0x'")
}

//SignedDataToTransaction 解析离线数据
func SignedDataToTransaction(signedData string) (*types.Transaction, *ethcmn.Address, error) {
	var err error
	var signed []byte
	if signedData[:2] == "0x" {
		if signed, err = hexToBytes(signedData); err != nil {
			return nil, nil, err
		}
	} else {
		signed = []byte(signedData)
	}
	// 解析离线交易的交易数据
	transaction := new(types.Transaction)
	if err = rlp.DecodeBytes(signed, transaction); err != nil {
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

func getSinger(txn *types.Transaction) (ethcmn.Address, error) {
	from, err := types.Sender(types.HomesteadSigner{}, txn)
	if err != nil {
		return ethcmn.Address{}, err
	}
	return from, nil
}

func checkInputData(input []byte, length int, funcHash string) (bool, error) {
	if len(input) < length {
		common.Logger.Errorf("transaction input length:%d mismatch. expect:%d", len(input), length)
		return false, errors.New("contract function call input length mismatch")
	}
	r := strings.HasPrefix(string(input), HexPrefix+funcHash)
	if r == false {
		common.Logger.Errorf("transaction input functionName:%s mismatch. expect:%s", string(input)[2:6], funcHash)
		return false, errors.New("contract function call funcHash mismatch")
	}
	return true, nil
}

// PraseERC20Transfer generate erc20 txn
// from to value
func PraseERC20Transfer(txn *types.Transaction) (*[]string, error) {

	var functionCall = make([]string, 3)

	if ok, err := checkInputData(txn.Data(), ERC20TransferLength, ERC20MethodTransfer); ok == false {
		return nil, err
	}

	// 恢复得到交易签名者
	from, err := getSinger(txn)
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

	//检查输入数据的合法性
	if ok, err := checkInputData(txn.Data(), PointsBuyPointsLength, PointsMethodBuyPoints); ok == false {
		return nil, err
	}
	// 恢复得到交易签名者
	from, err := getSinger(txn)
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

// PraseConsumePoints 消费积分
// from to value
func PraseConsumePoints(txn *types.Transaction) ([]string, error) {

	var functionCall = make([]string, 2)

	if ok, err := checkInputData(txn.Data(), PointsConsumePointsLength, PointsMethodConsumePoints); ok == false {
		return nil, err
	}
	// 恢复得到交易签名者
	from, err := getSinger(txn)
	if err != nil {
		return nil, err
	}
	functionCall[0] = from.String()
	i256, ok := math.ParseBig256(HexPrefix + truncationZero(string(txn.Data()[10:])))
	if ok {
		functionCall[1] = i256.Text(10)
	}
	return functionCall, nil
}

// PraseRefundPoints 消费积分
// from to value
func PraseRefundPoints(txn *types.Transaction) ([]string, error) {

	var functionCall = make([]string, 1)

	if ok, err := checkInputData(txn.Data(), PointsRefundPointsLength, PointsMethodRefundPoints); ok == false {
		return nil, err
	}
	// 恢复得到交易签名者
	from, err := getSinger(txn)
	if err != nil {
		return nil, err
	}
	functionCall[0] = from.String()

	return functionCall, nil
}
