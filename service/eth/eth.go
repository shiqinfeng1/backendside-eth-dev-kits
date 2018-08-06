package eth

import (
	"os/exec"

	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EthConn 一个ethnode的连接
var EthConn *ethclient.Client

// PoAConn 一个ethnode的连接
var PoAConn *ethclient.Client

//AttachEthNode 连接以太坊节点
func AttachEthNode() error {
	con1, err := ethclient.Dial(cmn.Config().GetString("ethereum.endpoints")) //也可以是https地址,websocket地址
	if err != nil {
		cmn.Logger.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	EthConn = con1
	con2, err := ethclient.Dial(cmn.Config().GetString("poa.endpoints")) //也可以是https地址,websocket地址
	if err != nil {
		cmn.Logger.Fatalf("Failed to connect to the PoA client: %v", err)
	}
	PoAConn = con2
	return err
}

//CompileSolidity 编译智能合约
func CompileSolidity() error {
	for _, solcFile := range cmn.Config().GetStringSlice("soliditySource.endpoints") {
		cmd := exec.Command(
			"solc",
			"--sol ./contracts/"+solcFile+".sol",
			"--pkg contracts",
			"--out ./service/contracts/"+solcFile+".go") // 实际可以直接写成-alh
		if b, err := cmd.CombinedOutput(); err != nil {
			return err
		} else {
			cmn.Logger.Printf("compile solidity successfully: %v", b)
		}
	}
	return nil
}
