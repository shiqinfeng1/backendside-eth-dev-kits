package contracts

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
func DeployOMCToken(chainName, userID string, auth *bind.TransactOpts) (*ERC20.OMC, error) {

	//设置gaslimit 和 gasgprice
	v, _ := math.ParseBig256(cmn.Config().GetString("ethereum.gas"))
	gas := v.Uint64() * 20
	g, _ := math.ParseBig256(cmn.Config().GetString("ethereum.price"))
	price := hexutil.Big(*g)
	auth.GasPrice = price.ToInt()
	auth.GasLimit = gas

	//部署合约
	client := eth.ConnectEthNodeForContract(chainName)
	addr, txn, token, err := ERC20.DeployOMC(auth, client)
	//dump部署信息
	cmn.PrintDeployContactInfo(addr, txn, err)
	if err == nil {
		//更新合约地址到配置文件, 这样下次启动后不会重复部署
		updateOmcAddressToConfig(addr.Hex())
		//将会交易加入监听池
		var para = &eth.PendingPoolParas{
			ChainType: chainName,
			UserID:    userID,
			TxHash:    txn.Hash(),
			From:      auth.From,
			To:        common.Address{},
			Nonce:     txn.Nonce(),
		}
		eth.AppendToPendingPool(para)
		//等待交易上链
		mined, success, timeout := waitMinedSync(txn.Hash().Hex())
		cmn.Logger.Noticef("transaction: %v mined:%v success:%v timeout:%v", txn.Hash().Hex(), mined, success, timeout)
	}

	return token, err
}

// AttachOMCToken 部署合约
func AttachOMCToken(chainName string) (*ERC20.OMC, *ethclient.Client, error) {
	client := eth.ConnectEthNodeForContract(chainName)
	if client == nil {
		return nil, nil, errors.New("no eth client")
	}
	token, err := ERC20.NewOMC(common.HexToAddress(cmn.Config().GetString(chainName+".omcaddress")), client)
	if err != nil {
		cmn.Logger.Errorf("Failed to instantiate a Token contract: %v", err)
	}
	return token, client, err
}

// OMCTokenTransfer 执行transfer
func OMCTokenTransfer(chainName, userID string, auth *bind.TransactOpts, receiver string, amount uint64) (*types.Transaction, error) {

	//执行交易之前获取blocknumber,监听事件时从该block开始检查
	blockNum, err := eth.ConnectEthNodeForWeb3(chainName).EthBlockNumber()
	if err != nil {
		cmn.Logger.Errorf("Failed to EthBlockNumber: %v", err)
		return nil, err
	}

	nb, _ := OMCTokenBalanceOf(chainName, auth.From)
	if amount > nb.Uint64() {
		cmn.Logger.Errorf("Insufficient balance: has %v. need %v", nb.Uint64(), amount)
		return nil, err
	}
	// cmn.Logger.Infof("balance:%v", nb)
	// 手动指定gaslimit和gasprice
	// gas, _ := math.ParseBig256(cmn.Config().GetString("ethereum.gas"))
	// auth.GasLimit = gas.Uint64()
	omc, conn, err := AttachOMCToken(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachOMCToken: %v", err)
		return nil, err
	}
	defer conn.Close()
	txn, err := omc.Transfer(auth, common.HexToAddress(receiver), big.NewInt(0).SetUint64(amount))
	if err != nil {
		cmn.Logger.Errorf("Failed to Transfer: %v", err)
		return nil, err
	}

	var para = &eth.PendingPoolParas{
		ChainType: chainName,
		UserID:    userID,
		TxHash:    txn.Hash(),
		From:      auth.From,
		To:        *txn.To(),
		Nonce:     txn.Nonce(),
	}
	eth.AppendToPendingPool(para)

	//同步等待交易上链
	PollEventTransfer(
		chainName,
		txn.Hash().Hex(),
		blockNum.ToInt().Uint64(),
		auth.From, common.HexToAddress(receiver))

	return txn, err
}

//OMCTokenBalanceOf 查询余额
func OMCTokenBalanceOf(chainName string, addr common.Address) (*big.Int, error) {
	omc, conn, err := AttachOMCToken(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachOMCToken: %v", err)
		return nil, err
	}
	defer conn.Close()
	balance, err := omc.BalanceOf(&bind.CallOpts{Pending: true}, addr)
	if err != nil {
		cmn.Logger.Errorf("Get BalanceOf: %v fail: %v", addr, err)
		return nil, err
	}
	return balance, err
}

//PollEventTransfer 等待交易上链,如果执行成功,捕获Transfer事件
func PollEventTransfer(chainName, txHash string, startBlock uint64, from, to common.Address) {
	omc, conn, err := AttachOMCToken(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachOMCToken: %v", err)
		return
	}
	defer conn.Close()
	mined, success, timeout := waitMinedSync(txHash)
	cmn.Logger.Noticef("Transaction: %v Mined:%v Success:%v Timeout:%v", txHash, mined, success, timeout)

	//如果交易失败,则不会有事件触发,无需监听
	if success == true {
		catchEventTransfer(
			omc,
			startBlock,
			[]common.Address{from},
			[]common.Address{to})
	}
}

func catchEventTransfer(omc *ERC20.OMC, startBlock uint64, from, to []common.Address) {
	//TODO: 记录捕获Transfer事件
	history, err := omc.FilterTransfer(&bind.FilterOpts{Start: startBlock}, from, to)
	if err != nil {
		cmn.Logger.Errorf("fail to FilterTransfer: %v", err)
		return
	}
	for history.Next() {
		e := history.Event
		cmn.Logger.Infof("%s transfer to %s value=%s, at %d", e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
	}
}

// DeployPointCoin 部署合约
func DeployPointCoin(chainName, userID string, auth *bind.TransactOpts) (*ERC20.PointCoin, error) {

	//设置gaslimit 和 gasgprice
	v, _ := math.ParseBig256(cmn.Config().GetString("ethereum.gas"))
	gas := v.Uint64() * 20
	g, _ := math.ParseBig256(cmn.Config().GetString("ethereum.price"))
	price := hexutil.Big(*g)
	auth.GasPrice = price.ToInt()
	auth.GasLimit = gas

	//部署合约
	client := eth.ConnectEthNodeForContract(chainName)
	addr, txn, points, err := ERC20.DeployPointCoin(auth, client)
	//dump部署信息
	cmn.PrintDeployContactInfo(addr, txn, err)
	if err == nil {
		//更新合约地址到配置文件, 这样下次启动后不会重复部署
		updateOmcAddressToConfig(addr.Hex())
		//将会交易加入监听池
		var para = &eth.PendingPoolParas{
			ChainType: chainName,
			UserID:    userID,
			TxHash:    txn.Hash(),
			From:      auth.From,
			To:        common.Address{},
			Nonce:     txn.Nonce(),
		}
		eth.AppendToPendingPool(para)
		//等待交易上链
		mined, success, timeout := waitMinedSync(txn.Hash().Hex())
		cmn.Logger.Noticef("[deploy points]transaction: %v mined:%v success:%v timeout:%v", txn.Hash().Hex(), mined, success, timeout)
	}

	return points, err
}

// AttachPointCoin 关联合约
func AttachPointCoin(chainName string) (*ERC20.PointCoin, *ethclient.Client, error) {
	client := eth.ConnectEthNodeForContract(chainName)
	if client == nil {
		return nil, nil, errors.New("no eth client")
	}
	points, err := ERC20.NewPointCoin(common.HexToAddress(cmn.Config().GetString(chainName+".pointsaddress")), client)
	if err != nil {
		cmn.Logger.Errorf("Failed to instantiate a Token contract: %v", err)
	}
	return points, client, err
}

//PollEventMint 等待交易上链,如果执行成功,捕获Mint事件
func PollEventMint(chainName, txHash string, startBlock uint64, from, to common.Address) {
	points, conn, err := AttachPointCoin(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachPointCoin: %v", err)
		return
	}
	defer conn.Close()
	mined, success, timeout := waitMinedSync(txHash)
	cmn.Logger.Noticef("Transaction: %v Mined:%v Success:%v Timeout:%v", txHash, mined, success, timeout)

	//如果交易失败,则不会有事件触发,无需监听
	if success == true {
		catchEventMint(
			points,
			startBlock,
			[]common.Address{to})
	}
}
func catchEventMint(points *ERC20.PointCoin, startBlock uint64, to []common.Address) {
	//TODO: 记录捕获Transfer事件
	history, err := points.FilterMint(&bind.FilterOpts{Start: startBlock}, to)
	if err != nil {
		cmn.Logger.Errorf("fail to FilterTransfer: %v", err)
		return
	}
	for history.Next() {
		e := history.Event
		cmn.Logger.Infof("%s buy points to %s value=%s, at %d", e.To.String(), e.Amount, e.Raw.BlockNumber)
	}
}

func waitMinedSync(txHash string) (mined bool, success bool, timeout bool, minedBlock uint64, comfired int) {
	var count int
	var err error
	timeout = false
	for {
		time.Sleep(time.Second * 2)
		if mined, success, minedBlock, comfired, err = eth.IsMined(txHash); err != nil {
			cmn.Logger.Errorf("transaction %v is mined fail: %v", txHash, err)
			return
		}
		if mined == true {
			return
		}
		count++
		if count > cmn.Config().GetInt("ethereum.txtimeout")/2 {
			break
		}
	}
	timeout = true
	return
}
