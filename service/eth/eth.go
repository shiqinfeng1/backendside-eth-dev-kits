package eth

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kr/pretty"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/endpoints"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
)

func getCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:(index - len("build/bin/"))]
	return ret
}

//检查目录是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(filename + " not exist\n")
		exist = false
	}
	return exist
}

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

//CompileContracts 编译智能合约
func CompileContracts() error {

	for _, solcFile := range cmn.Config().GetStringSlice("solidity.source") {
		fmt.Printf("get solcFile: %v ", solcFile)
		path := strings.Split(solcFile, string(os.PathSeparator))
		path = path[:len(path)-1]
		dir := "./service/contracts"
		for _, name := range path {
			dir = dir + "/" + name
			if !checkFileIsExist(dir) {
				err := os.Mkdir(dir, os.ModePerm) //创建文件夹
				if err != nil {
					cmn.Logger.Error(err)
					return err
				}
			}
		}
		if err := CompileSolidity(
			"./contracts/"+solcFile+".sol",
			"./service/contracts/"+solcFile+".go",
			cmn.Config().GetString("solidity.exclude")); err != nil {
			return err
		}
	}
	return nil
}

//CompileSolidity 编译
func CompileSolidity(source, dest, exclude string) error {
	cmd := exec.Command(
		"./abigen",
		"-sol="+source,
		"-pkg=contracts",
		"-exc="+exclude,
		"-out="+dest)
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("\n\n*****Compile solidity fail********************")
		fmt.Println("output:", string(b))
		fmt.Print(err)
		fmt.Printf("**************************************************\n\n")
		return err
	}
	if b != nil && len(b) == 0 {
		cmn.Logger.Printf("compile solidity %s ok.", source)
	} else {
		fmt.Println("\n\n*****Compile solidity output********************")
		fmt.Print(string(b))
		fmt.Printf("**************************************************\n\n")
	}
	return nil
}

//Transfer 发送交易
func Transfer(p httpservice.PadPayload) (string, error) {
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

	con := ConnectEthNodeForWeb3(p.ChainType)
	nonce, err3 := con.EthGetNonce(userAddress)
	if err3 != nil {
		return "", errors.New("get nounce fail addr: " + userAddress.Hex() + ". err:" + err3.Error())
	}

	transaction := &cmn.TransactionRequest{
		From:     userAddress,
		To:       ethcmn.HexToAddress(cmn.Config().GetString("ethereum.adminaccount")),
		Gas:      &gas,
		GasPrice: &price,
		Value:    &amount,
		Data:     *new(hexutil.Bytes),
		Nonce:    nonce,
	}
	signedTx, err4 := accounts.SignTx(p.UserID, transaction)
	if err4 != nil {
		pretty.Println("\nsignedTx:", signedTx, "\nerror: ", err4, "\n")
		return "", nil
	}
	txHash, err5 := con.EthSendRawTransaction(signedTx)
	if err5 != nil {
		pretty.Println("\ntxHash:", txHash.Hex(), "\nerror: ", err5, "\n")
		return "", nil
	}
	return txHash.String(), nil
}
