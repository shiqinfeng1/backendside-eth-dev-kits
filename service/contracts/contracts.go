package contracts

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/contracts/ERC20"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
)

//检查目录是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(filename + " not exist\n")
		exist = false
	}
	return exist
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
			strings.Split(solcFile, "/")[0],
			cmn.Config().GetString("solidity.exclude")); err != nil {
			return err
		}
	}
	return nil
}

//CompileSolidity 编译
func CompileSolidity(source, dest, pkgname, exclude string) error {
	cmd := exec.Command(
		"./abigen",
		"-sol="+source,
		"-pkg="+pkgname,
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
		fmt.Printf("compile solidity %s ok.", source)
	} else {
		fmt.Println("\n\n*****Compile solidity output********************")
		fmt.Print(string(b))
		fmt.Printf("**************************************************\n\n")
	}
	return nil
}

//=========================

func updateAddressToConfig(match, address string) {

	cfg, _ := filepath.Abs("./myConfig.yaml")
	input, err := ioutil.ReadFile(cfg)
	if err != nil {
		cmn.Logger.Errorf(err.Error())
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, match) {
			lines[i] = "  " + match + ": " + address
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./myConfig.yaml", []byte(output), 0644)
	if err != nil {
		cmn.Logger.Errorf(err.Error())
	}
	return
}
func updatePointAddressToConfig(pointAddress string) {
	updateAddressToConfig("pointaddress", pointAddress)
}
func updateOmcAddressToConfig(omcAddress string) {
	updateAddressToConfig("omcaddress", omcAddress)
}

//GetUserAuth 根据用户名获取auth
func GetUserAuth(userID string) (*bind.TransactOpts, error) {
	auth, err := accounts.GetTransactOptsFromHDWallet(userID)
	if err != nil {
		cmn.Logger.Errorf("Failed to getUserAuth: %v", err)
		return nil, err
	}
	return auth, err
}

//GetUserAuthWithPassword 根据以太地址和密码获取
func GetUserAuthWithPassword(userAddr, pwd string) (*bind.TransactOpts, error) {
	auth, err := accounts.GetTransactOptsFromKeystore(userAddr, pwd)
	if err != nil {
		cmn.Logger.Errorf("Failed to getUserAuthWithPassword: %v", err)
		return nil, err
	}
	return auth, err
}

// DeployOMCToken 部署合约
func DeployOMCToken(auth *bind.TransactOpts) (*ERC20.OMC, error) {

	v, _ := math.ParseBig256(cmn.Config().GetString("ethereum.gas"))
	gas := v.Uint64() * 20

	g, _ := math.ParseBig256(cmn.Config().GetString("ethereum.price"))
	price := hexutil.Big(*g)

	auth.GasPrice = price.ToInt()
	auth.GasLimit = gas

	client := eth.ConnectEthNodeForContract("poa")
	addr, txn, token, err := ERC20.DeployOMC(auth, client)
	if err == nil {
		updateOmcAddressToConfig(addr.Hex())
	}
	cmn.PrintDeployContactInfo(addr, txn, err)

	return token, err
}

// AttachOMCToken 部署合约
func AttachOMCToken() (*ERC20.OMC, error) {
	client := eth.ConnectEthNodeForContract("ethereum")
	token, err := ERC20.NewOMC(common.HexToAddress(cmn.Config().GetString("ethereum.omcaddress")), client)
	if err != nil {
		cmn.Logger.Errorf("Failed to instantiate a Token contract: %v", err)
	}
	return token, err
}

// OMCTokenTransfer 执行transfer
func OMCTokenTransfer(auth *bind.TransactOpts, receiver string, amount int64) (*types.Transaction, error) {

	omc, err := AttachOMCToken()
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachOMCToken: %v", err)
		return nil, err
	}

	txHash, err := omc.Transfer(auth, common.HexToAddress(receiver), big.NewInt(amount))

	return txHash, err
}
