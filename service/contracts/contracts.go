package contracts

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/contracts/ERC20"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
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
	updateAddressToConfig("pointsaddress", pointAddress)
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

	cmn.Logger.Debugf("[DeployOMCToken] chainName:%v userID: %v", chainName, userID)
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
			ChainType:   chainName,
			UserID:      userID,
			TxHash:      txn.Hash(),
			From:        auth.From,
			To:          ethcmn.Address{},
			Nonce:       txn.Nonce(),
			Description: fmt.Sprintf("%v.%v.%v.%v", chainName, "DeployOMCToken", userID, auth.From.Hex()),
		}
		eth.AppendToPendingPool(para)
		//等待交易上链
		waitMinedSync(txn.Hash().Hex())
	}

	return token, err
}

// AttachOMCToken 部署合约
func AttachOMCToken(chainName string) (*ERC20.OMC, *ethclient.Client, error) {
	client := eth.ConnectEthNodeForContract(chainName)
	if client == nil {
		return nil, nil, errors.New("no eth client")
	}
	token, err := ERC20.NewOMC(ethcmn.HexToAddress(cmn.Config().GetString(chainName+".omcaddress")), client)
	if err != nil {
		cmn.Logger.Errorf("Failed to instantiate a Token contract: %v", err)
	}
	return token, client, err
}

// OMCTokenTransfer 执行transfer
func OMCTokenTransfer(chainName, userID string, auth *bind.TransactOpts, receiver string, amount uint64) (*types.Transaction, error) {

	cmn.Logger.Debugf("[OMCTokenTransfer] chainName:%v userID: %v receiver:%v amount:%v", chainName, userID, receiver, amount)
	//执行交易之前获取blocknumber,监听事件时从该block开始检查
	blockNum, err := eth.ConnectEthNodeForWeb3(chainName).EthBlockNumber()
	if err != nil {
		cmn.Logger.Errorf("Failed to EthBlockNumber: %v", err)
		return nil, err
	}

	nb, _ := OMCTokenBalanceOf(chainName, auth.From)
	if amount > nb.Uint64() {
		cmn.Logger.Errorf("Insufficient balance: has %v. need %v", nb.Uint64(), amount)
		return nil, errors.New("Insufficient balance")
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
	txn, err := omc.Transfer(auth, ethcmn.HexToAddress(receiver), big.NewInt(0).SetUint64(amount))
	if err != nil {
		cmn.Logger.Errorf("Failed to Transfer: %v", err)
		return nil, err
	}

	//交易加入pending监听队列
	var para = &eth.PendingPoolParas{
		ChainType:   chainName,
		UserID:      userID,
		TxHash:      txn.Hash(),
		From:        auth.From,
		To:          *txn.To(),
		Nonce:       txn.Nonce(),
		Description: fmt.Sprintf("%v.%v.%v.%v:%v.%v", chainName, "OMCToken.transfer", userID, auth.From.Hex(), receiver, amount),
	}
	eth.AppendToPendingPool(para)

	//等待交易上链,并捕获Transfer事件
	go PollEventTransfer(
		chainName,
		txn.Hash().Hex(),
		blockNum.ToInt().Uint64(),
		auth.From, ethcmn.HexToAddress(receiver))

	return txn, err
}

//OMCTokenBalanceOf 查询余额
func OMCTokenBalanceOf(chainName string, addr ethcmn.Address) (*big.Int, error) {
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

// DeployPointCoin 部署合约
func DeployPointCoin(chainName, userID string, auth *bind.TransactOpts) (*ERC20.PointCoin, error) {

	cmn.Logger.Debugf("[DeployPointCoin] chainName:%v userID: %v", chainName, userID)
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
		updatePointAddressToConfig(addr.Hex())
		//将会交易加入监听池
		var para = &eth.PendingPoolParas{
			ChainType:   chainName,
			UserID:      userID,
			TxHash:      txn.Hash(),
			From:        auth.From,
			To:          ethcmn.Address{},
			Nonce:       txn.Nonce(),
			Description: fmt.Sprintf("%v.%v.%v.%v", chainName, "DeployPointCoin", userID, auth.From.Hex()),
		}
		eth.AppendToPendingPool(para)
		//等待交易上链
		waitMinedSync(txn.Hash().Hex())
	}

	return points, err
}

// AttachPointCoin 关联合约
func AttachPointCoin(chainName string) (*ERC20.PointCoin, *ethclient.Client, error) {
	client := eth.ConnectEthNodeForContract(chainName)
	if client == nil {
		return nil, nil, errors.New("no eth client")
	}
	points, err := ERC20.NewPointCoin(ethcmn.HexToAddress(cmn.Config().GetString(chainName+".pointsaddress")), client)
	if err != nil {
		cmn.Logger.Errorf("Failed to instantiate a Token contract: %v", err)
	}
	return points, client, err
}

//PointCoinBalanceOf 查询余额
func PointCoinBalanceOf(chainName string, addr ethcmn.Address) (*big.Int, error) {
	points, conn, err := AttachPointCoin(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachPointCion: %v", err)
		return nil, err
	}
	defer conn.Close()
	balance, err := points.BalanceOf(&bind.CallOpts{Pending: true}, addr)
	if err != nil {
		cmn.Logger.Errorf("Get BalanceOf: %v fail: %v", addr, err)
		return nil, err
	}
	return balance, err
}

// PointsBuy 购买积分
func PointsBuy(chainName, userID string, auth *bind.TransactOpts, receiver string, amount uint64) (*types.Transaction, error) {

	cmn.Logger.Debugf("[PointsBuy] chainName:%v userID: %v receiver:%v amount:%v", chainName, userID, receiver, amount)
	//执行交易之前获取blocknumber,监听事件时从该block开始检查
	blockNum, err := eth.ConnectEthNodeForWeb3(chainName).EthBlockNumber()
	if err != nil {
		cmn.Logger.Errorf("Failed to EthBlockNumber: %v", err)
		return nil, err
	}
	//关联合约
	points, conn, err := AttachPointCoin(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachPointCoin: %v", err)
		return nil, err
	}
	defer conn.Close()
	txn, err := points.Buy(auth, ethcmn.HexToAddress(receiver), big.NewInt(0).SetUint64(amount))
	if err != nil {
		cmn.Logger.Errorf("Failed to Buy points: %v", err)
		return nil, err
	}
	// 记录消费积分的信息
	balance, err := PointCoinBalanceOf(chainName, ethcmn.HexToAddress(receiver))
	if err != nil {
		return nil, err
	}
	var ppara = &eth.PointsParas{
		ChainType:      chainName,
		UserID:         userID,
		TxHash:         txn.Hash(),
		UserAddress:    ethcmn.HexToAddress(receiver),
		TxnType:        "buy",
		PreBalance:     balance.Uint64(),
		ExpectBalance:  balance.Uint64() + amount,
		IncurredAmount: uint64(amount),
		CurrentStatus:  "apply",
	}
	if err := PointsBuyRequireToDB(ppara); err != nil {
		return nil, err
	}

	var para = &eth.PendingPoolParas{
		ChainType:   chainName,
		UserID:      userID,
		TxHash:      txn.Hash(),
		From:        auth.From,
		To:          *txn.To(),
		Nonce:       txn.Nonce(),
		Description: fmt.Sprintf("%v.%v.%v.%v:%v", chainName, "PointCoin.buy", auth.From.Hex(), receiver, amount),
	}
	eth.AppendToPendingPool(para)
	//等待交易上链,并捕获Transfer事件
	go PollEventMint(
		chainName,
		txn.Hash().Hex(),
		blockNum.ToInt().Uint64(),
		ethcmn.HexToAddress(receiver))
	return txn, err
}

// PointsConsume 消费积分,keystore和密码进行离线签名并发送交易
func PointsConsume(chainName, consumerID, consumer, passphrase string, amount int64) (*types.Transaction, error) {

	cmn.Logger.Debugf("[PointsConsume] chainName:%v consumer:%v amount:%v", chainName, consumer, amount)
	userAddress := ethcmn.HexToAddress(consumer)

	nb, _ := PointCoinBalanceOf(chainName, userAddress)
	if uint64(amount) > nb.Uint64() {
		cmn.Logger.Errorf("Insufficient balance: has %v. need %v", nb.Uint64(), amount)
		return nil, errors.New("Insufficient balance")
	}

	//组装函数调用rlp编码数据
	abi, err := abi.JSON(strings.NewReader(ERC20.PointCoinABI))
	if err != nil {
		cmn.Logger.Errorf("Prase ABI Fail: %v", err)
		return nil, err
	}

	input, err := abi.Pack("consume", big.NewInt(amount))
	if err != nil {
		cmn.Logger.Errorf("Pack Input Fail: %v", err)
		return nil, err
	}

	//设置gas和gaslimit
	v, _ := math.ParseBig256(cmn.Config().GetString("ethereum.gas"))
	gas := hexutil.Big(*v)
	g, _ := math.ParseBig256(cmn.Config().GetString("ethereum.price"))
	price := hexutil.Big(*g)

	//获取指定地址的nonce
	con := eth.ConnectEthNodeForWeb3(chainName)
	if con == nil {
		cmn.Logger.Errorf("Connect Eth Node Fail: %v", err)
		return nil, err
	}
	nonce, err := con.EthGetNonce(userAddress)
	if err != nil {
		cmn.Logger.Errorf("Get address Nonce Fail: %v", err)
		return nil, err
	}
	//构造交易数据
	rawTx := types.NewTransaction(
		nonce.ToInt().Uint64(),
		ethcmn.HexToAddress(cmn.Config().GetString(chainName+".pointsaddress")),
		nil, //value=0
		gas.ToInt().Uint64(),
		price.ToInt(),
		input)

	//签名交易数据
	key, err := accounts.GetAccountFromKeystore(userAddress.Hex(), passphrase)
	if err != nil {
		cmn.Logger.Errorf("Get address Nonce Fail: %v", err)
		return nil, err
	}
	var signedData bytes.Buffer
	signedTx, err := types.SignTx(rawTx, types.HomesteadSigner{}, key.PrivateKey)
	if err != nil {
		cmn.Logger.Errorf("SignTx Fail: %v", err)
		return nil, err
	}
	signedTx.EncodeRLP(&signedData)

	//发送离线交易
	var rawSigned = &eth.RawData{SignedData: signedData.String(), ChainType: chainName}

	txHash, blockNum, err := eth.SendRawTxn(rawSigned)
	if err != nil {
		cmn.Logger.Errorf("SendRawTxn Fail: %v", err)
		return nil, err
	}

	var ppara = &eth.PointsParas{
		ChainType:      chainName,
		UserID:         consumerID,
		TxHash:         ethcmn.HexToHash(txHash),
		UserAddress:    userAddress,
		TxnType:        "consume",
		PreBalance:     nb.Uint64(),
		ExpectBalance:  uint64(nb.Int64() - amount),
		IncurredAmount: uint64(amount),
		CurrentStatus:  "apply",
	}
	if err := PointsConsumeRequireToDB(ppara); err != nil {
		return nil, err
	}

	var para = &eth.PendingPoolParas{
		ChainType:   chainName,
		UserID:      consumerID,
		TxHash:      ethcmn.HexToHash(txHash),
		From:        userAddress,
		To:          *rawTx.To(),
		Nonce:       rawTx.Nonce(),
		Description: fmt.Sprintf("%s.%s.%s:%v", chainName, "PointCoin.comsume", consumer, amount),
	}
	eth.AppendToPendingPool(para)
	//等待交易上链,并捕获Burn事件
	go PollEventBurn(
		"consume",
		chainName,
		txHash,
		blockNum.ToInt().Uint64(),
		userAddress)
	return rawTx, err
}

//PointsBuyRequireToDB 记录积分购买
func PointsBuyRequireToDB(para *eth.PointsParas) error {
	return addPointsRequireRecord(para)
}

//PointsConsumeRequireToDB 记录积分消费
func PointsConsumeRequireToDB(para *eth.PointsParas) error {
	return addPointsRequireRecord(para)
}

//PointsRefundRequireToDB 记录积分退还
func PointsRefundRequireToDB(para *eth.PointsParas) error {
	return addPointsRequireRecord(para)
}

func addPointsRequireRecord(para *eth.PointsParas) error {

	pInfo := &db.PointsInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notfound := dbconn.Model(&db.PointsInfo{}).
		Where("txn_hash = ?", para.TxHash.String()).
		Find(pInfo).
		RecordNotFound()
	if notfound {
		pInfo.ChainType = para.ChainType
		pInfo.UserID = para.UserID
		pInfo.UserAddress = para.UserAddress.String()
		pInfo.TxnType = para.TxnType
		pInfo.TxnHash = para.TxHash.String()
		pInfo.PreBalance = para.PreBalance
		pInfo.ExpectBalance = para.ExpectBalance
		pInfo.IncurredAmount = para.IncurredAmount
		pInfo.CurrentStatus = para.CurrentStatus

		err := dbconn.Create(pInfo).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	} else {
		err := fmt.Errorf("txHash:%s is already added", para.TxHash.String())
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

//PointsBuyComfiredToDB 购买积分交易确认
func PointsBuyComfiredToDB(txHash string, userAddress string, amount uint64) error {

	pInfo := &db.PointsInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notfound := dbconn.Model(&db.PointsInfo{}).
		Where("txn_hash = ?", txHash).
		Find(pInfo).
		RecordNotFound()
	if !notfound &&
		pInfo.UserAddress == userAddress &&
		pInfo.TxnType == "buy" &&
		pInfo.IncurredAmount == amount {

		err := dbconn.Model(pInfo).Update("current_status", "comfired").Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	} else {
		err := fmt.Errorf("PointsBuyComfiredToDB. txHash:%s not found", txHash)
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}

//PointsConsumeComfiredToDB 消费积分交易确认
func PointsConsumeComfiredToDB(txHash string, userAddress string, amount uint64) error {
	return addPointsComfiredRecord("consume", txHash, userAddress, amount)
}

//PointsRefundComfiredToDB 消费积分交易确认
func PointsRefundComfiredToDB(txHash string, userAddress string, amount uint64) error {
	return addPointsComfiredRecord("refund", txHash, userAddress, amount)
}
func addPointsComfiredRecord(txnType, txHash string, userAddress string, amount uint64) error {

	pInfo := &db.PointsInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notfound := dbconn.Model(&db.PointsInfo{}).
		Where("txn_hash = ?", txHash).
		Find(pInfo).
		RecordNotFound()
	if !notfound &&
		pInfo.UserAddress == userAddress &&
		pInfo.TxnType == txnType &&
		pInfo.IncurredAmount == amount {

		err := dbconn.Model(pInfo).Update("current_status", "comfired").Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	} else {
		err := fmt.Errorf("addPointsComfiredRecord.%s txHash:%s not found", txnType, txHash)
		cmn.Logger.Error(err)
		return err
	}
	dbconn.MysqlCommit()
	return nil
}
