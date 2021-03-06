package eth

import (
	"errors"
	"fmt"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/endpoints"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
)

//ConnectEthNodeForContract 连接以太坊节点
func ConnectEthNodeForContract(nodeType string) *ethclient.Client {
	for _, endpoint := range endpoints.GetEndPointsManager().AliveEndpoints {
		if nodeType == endpoint.NodeType {
			con1, err := ethclient.Dial(endpoint.URL) //也可以是https地址,websocket地址
			if err != nil {
				continue
			}
			return con1
		}
	}
	return nil
}

//ConnectEthNodeForWeb3 连接以太坊节点
func ConnectEthNodeForWeb3(nodeType string) *Client {
	for _, endpoint := range endpoints.GetEndPointsManager().AliveEndpoints {
		if nodeType == endpoint.NodeType {
			con := NewClient(endpoint.URL)
			// addr := common.HexToAddress("0x1dcef12e93b0abf2d36f723e8b59cc762775d513")
			// v, err := con.EthGetBalance(addr, nil)
			// if err != nil {
			// 	fmt.Println(v, err)
			// 	return nil
			// }
			// fmt.Println(v)
			return con
		}
	}
	return nil
}

// func getCurrPath() string {
// 	file, _ := exec.LookPath(os.Args[0])
// 	path, _ := filepath.Abs(file)
// 	index := strings.LastIndex(path, string(os.PathSeparator))
// 	ret := path[:(index - len("build/bin/"))]
// 	return ret
// }

//TransferEth 发送转账交易
func TransferEth(p httpservice.TransferPayload) (string, error) {
	a, err := math.ParseBig256(p.Amount)
	if err == false {
		return "", errors.New("invalid transfer amount")
	}
	amount := hexutil.Big(*a)

	v, _ := math.ParseBig256(cmn.Config().GetString("ethereum.gas"))
	gas := hexutil.Big(*v)

	g, _ := math.ParseBig256(cmn.Config().GetString("ethereum.price"))
	price := hexutil.Big(*g)

	userAddress, err2 := accounts.GetUserAddress(p.UserID)
	if err2 != nil {
		return "", errors.New("no such userID: " + p.UserID + ". err:" + err2.Error())
	}

	nonce, err3 := GetNonce(userAddress.Hex())
	if err3 != nil {
		return "", errors.New("get nounce fail addr: " + userAddress.Hex() + ". err:" + err3.Error())
	}
	n := hexutil.Big(*nonce)
	transaction := &cmn.TransactionRequest{
		From:     userAddress,
		To:       ethcmn.HexToAddress(cmn.Config().GetString("ethereum.adminaccount")),
		Gas:      &gas,
		GasPrice: &price,
		Value:    &amount,
		Data:     *new(hexutil.Bytes),
		Nonce:    &n,
	}
	signedTx, err4 := accounts.SignTx(p.UserID, transaction)
	if err4 != nil {
		cmn.Logger.Error("\nsignedTx:", signedTx, "\nerror: ", err4, "\n")
		return "", nil
	}

	con := ConnectEthNodeForWeb3(p.ChainType)
	if con == nil {
		return "", errors.New("no valid endpoint")
	}
	txHash, err5 := con.EthSendRawTransaction(signedTx)
	if err5 != nil {
		cmn.Logger.Error("\ntxHash:", txHash.String(), "\nerror: ", err5, "\n")
		return "", err5
	}
	var para = &PendingPoolParas{
		ChainType:   p.ChainType,
		UserID:      p.UserID,
		TxHash:      txHash,
		From:        userAddress,
		To:          transaction.To,
		Nonce:       transaction.Nonce.ToInt().Uint64(),
		Description: fmt.Sprintf("%v.%v.%v.%v:%v.%v", p.ChainType, "OMCToken.transfer", p.UserID, userAddress.Hex(), transaction.To.Hex(), amount),
	}
	AppendToPendingPool(para)
	return txHash.String(), nil
}

//SendRawTransaction 发送离线交易
func SendRawTransaction(p httpservice.RawTransactionPayload) (string, error) {

	_, err2 := accounts.GetUserAddress(p.UserID)
	if err2 != nil {
		return "", errors.New("no such userID: " + p.UserID + ". err:" + err2.Error())
	}

	var raw = &RawData{SignedData: p.SignedData, ChainType: p.ChainType}
	txHash, _, err := SendRawTxn(raw)
	if err != nil {
		cmn.Logger.Error("\ntxHash:", txHash, "\nerror: ", err, "\n")
		return "", err
	}
	transaction, from, err := SignedDataToTransaction(p.SignedData)
	if err != nil {
		return ethcmn.Hash{}.String(), err
	}
	cmn.Logger.Debug("\nDecodeRLP \nsigner:", from, "\ntxhash:", txHash)
	var para = &PendingPoolParas{
		ChainType:   p.ChainType,
		UserID:      p.UserID,
		TxHash:      ethcmn.HexToHash(txHash),
		From:        *from,
		To:          *transaction.To(),
		Nonce:       transaction.Nonce(),
		Description: fmt.Sprintf("%v.%v.%v.%v.%v", p.ChainType, "RawTxn", p.UserID, from.Hex(), (*transaction.To()).Hex()),
	}
	AppendToPendingPool(para)
	return txHash, nil
}

//TransactionIsMined 发送离线交易
func TransactionIsMined(p httpservice.QueryTransactionPayload) (resp *httpservice.QueryTransactionReponse, err error) {
	var desc string
	addr, err := accounts.GetUserAddress(p.UserID)
	if err != nil {
		return nil, errors.New("no such userID: " + p.UserID + ". err:" + err.Error())
	}

	resp.Mined, resp.Success, resp.MinedBlock, resp.Comfired, desc, err = IsMined(p.TxHash)
	if err != nil {
		return nil, errors.New("desc: " + desc + " addr: " + addr.String() + " txhash: " + p.TxHash + "check mined err:" + err.Error())
	}
	return
}

//SendRawTxn 发送离线交易
func SendRawTxn(p *RawData) (string, *hexutil.Big, error) {

	// 获取节点连接
	con := ConnectEthNodeForWeb3(p.ChainType)
	if con == nil {
		return "", nil, errors.New("no valid endpoint")
	}
	// 执行交易之前获取blocknumber,监听事件时从该block开始检查
	blockNum, err := con.EthBlockNumber()
	if err != nil {
		cmn.Logger.Errorf("Failed to EthBlockNumber: %v", err)
		return "", nil, err
	}

	//如果传入的是16进制数字字符串,需要先转换为字符数组
	var signed []byte
	if p.SignedData[:2] == "0x" {
		if signed, err = hexToBytes(p.SignedData); err != nil {
			return "", nil, err
		}
	} else {
		signed = []byte(p.SignedData)
	}
	// 发送离线交易
	txHash, err5 := con.EthSendRawTransaction(signed)
	if err5 != nil {
		cmn.Logger.Error("\ntxHash:", txHash.String(), "\nerror: ", err5, "\n")
		return "", nil, err5
	}

	return txHash.String(), blockNum, nil
}
