package eth

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EthConn 一个ethnode的连接
var EthConn *ethclient.Client

// PoAConn 一个ethnode的连接
var PoAConn *ethclient.Client

//GetCurrPath 获取路径
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:(index - len("build/bin/"))]
	return ret
}

//AttachEthNode 连接以太坊节点
func AttachEthNode() error {
	con1, err := ethclient.Dial(cmn.Config().GetString("ethereum.endpoints")) //也可以是https地址,websocket地址
	if err != nil {
		cmn.Logger.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return err
	}
	EthConn = con1

	con2, err := ethclient.Dial(cmn.Config().GetString("poa.endpoints")) //也可以是https地址,websocket地址
	if err != nil {
		cmn.Logger.Fatalf("Failed to connect to the PoA client: %v", err)
		return err
	}
	PoAConn = con2
	return err
}

//CompileSolidity 编译智能合约
func CompileSolidity() error {
	for _, solcFile := range cmn.Config().GetStringSlice("soliditySource") {
		cmn.Logger.Printf("get solcFile: %v ", solcFile)
		cmd := exec.Command(
			"./abigen",
			"-sol=./contracts/"+solcFile+".sol",
			"-pkg=contracts",
			"-exc=ERC20,ERC20Basic,SafeERC20,SafeMath,Ownable",
			"-out=./service/contracts/"+solcFile[strings.LastIndex(solcFile, string(os.PathSeparator))+1:]+".go") // 实际可以直接写成-alh
		b, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("\n*****Compile solidity fail********************")
			fmt.Println("output:", string(b))
			fmt.Println(err)
			fmt.Printf("**************************************************\n\n")
			return err
		}
		if b != nil && len(b) == 0 {
			cmn.Logger.Printf("compile solidity %s ok.", solcFile)
		} else {
			fmt.Println("\n*****Compile solidity output********************")
			fmt.Println(string(b))
			fmt.Printf("**************************************************\n\n")
		}
	}
	return nil
}
