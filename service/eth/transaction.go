package eth

import (
	"errors"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
)

//Comfired 等待确认的交易
type Comfired struct {
	TxHash     string
	MinedBlock uint64
}

// PendingTransactionManager endpoints of ethereum
type PendingTransactionManager struct {
	currentPending map[string][]string
	waitComfired   map[string][]Comfired
	txHashTimeout  chan string
	exit           chan bool
	closed         chan bool
}

var ptm *PendingTransactionManager

// NewPendingTransactionManager create a pending transaction manager
func NewPendingTransactionManager() *PendingTransactionManager {
	ptm = &PendingTransactionManager{
		currentPending: make(map[string][]string),
		waitComfired:   make(map[string][]Comfired),
		txHashTimeout:  make(chan string, 256),
		exit:           make(chan bool),
		closed:         make(chan bool),
	}
	return ptm
}

// GetPendingTransactionManager return a pending transaction manager
func GetPendingTransactionManager() *PendingTransactionManager {
	return ptm
}

func (e *PendingTransactionManager) importAllPendingTransactions(chainName string) {
	var ptInfo []*db.PendingTransactionInfo
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()
	dbconn.Model(&db.PendingTransactionInfo{}).
		Where("mined = ? and listen_timeout = ? and chain_type = ?", false, false, chainName).Find(&ptInfo)

	var currentPending []string
	for _, item := range ptInfo {
		currentPending = append(currentPending, item.TxHash)
	}
	e.currentPending[chainName] = currentPending
}

// 更新交易的状态 mined
func (e *PendingTransactionManager) updateTransactionStatus(txHash string, status, success bool, minedblock uint64) error {
	var ptInfo = &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).Where("tx_hash = ?", txHash).Find(ptInfo).RecordNotFound()
	if notFound {
		err := "not found txHash:" + txHash + "in PendingTransactionInfo"
		cmn.Logger.Error(err)
		return errors.New(err)
	}
	err := dbconn.Model(ptInfo).Update("mined", status).Update("success", success).Update("minedblock", minedblock).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}

	cmn.Logger.Debugf("txHash %s update status to %v", txHash, status)
	dbconn.MysqlCommit()
	return nil
}

// 更新交易的确认数
func (e *PendingTransactionManager) updateTransactionComfired(txHash string, comfired int) error {
	var ptInfo = &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).Where("tx_hash = ?", txHash).Find(ptInfo).RecordNotFound()
	if notFound {
		err := "not found txHash:" + txHash + "in PendingTransactionInfo"
		cmn.Logger.Error(err)
		return errors.New(err)
	}
	err := dbconn.Model(ptInfo).Update("comfired", comfired).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}

	cmn.Logger.Debugf("txHash %s update comfired block to %v", txHash, comfired)
	dbconn.MysqlCommit()
	return nil
}

// 监听交易是否上链超时后,更新交易的状态记录 listentimeout
func (e *PendingTransactionManager) updateTransactionTimeout(txHash string, listenTimeoutAt time.Time) error {
	var ptInfo = &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).Where("tx_hash = ?", txHash).Find(ptInfo).RecordNotFound()
	if notFound {
		err := "not found txHash:" + txHash + "in PendingTransactionInfo"
		cmn.Logger.Error(err)
		return errors.New(err)
	}
	//在超时之前已经上链
	if ptInfo.Mined == true {
		cmn.Logger.Debug("transaction is already be mined: " + txHash)
		return nil
	}
	//更新记录超时时间
	err := dbconn.Model(ptInfo).Update("listen_timeout_at", listenTimeoutAt).Update("listen_timeout", true).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}

	cmn.Logger.Infof("txHash %s update listenTimeoutAt %v", txHash, listenTimeoutAt)
	dbconn.MysqlCommit()
	return nil
}

//监听并更新所有发送单位被确认上链的交易
func (e *PendingTransactionManager) watch(chainName string) {
	//和节点建立连接
	con := ConnectEthNodeForWeb3(chainName)
	if con == nil {
		return
	}
	//遍历等待上链的交易,检查是否上链
	for _, txHash := range e.currentPending[chainName] {

		transaction, err := con.EthGetTransactionByHash(ethcmn.HexToHash(txHash))
		if err != nil {
			cmn.Logger.Debug("query in [" + chainName + "] txhash: " + txHash + ". error :" + err.Error())
			continue
		}
		cmn.Logger.Debugf("query %s transaction: %v\n", chainName, transaction)
		if transaction.BlockNumber != nil && transaction.BlockNumber.ToInt().Uint64() > 0 {
			//检查交易是否成功
			receipt, err := con.EthGetTransactionReceipt(ethcmn.HexToHash(txHash))
			if err != nil {
				cmn.Logger.Debugf("query receipt in [%s] txhash: %s. error : %v", chainName, txHash, err.Error())
				continue
			}
			cmn.Logger.Debugf("[poa]receipt.GasUsed=%v transaction.Gas=%v", receipt.GasUsed.ToInt().Uint64(), transaction.Gas.ToInt().Uint64())
			//交易成功的标志是status==1,或者消耗的gas小于设置的gaslimit
			success := ((receipt.Status.ToInt().Uint64() == 1) ||
				(receipt.GasUsed.ToInt().Uint64() < transaction.Gas.ToInt().Uint64()))
			//更新交易状态到数据库
			e.updateTransactionStatus(txHash, true, success, transaction.BlockNumber.ToInt().Uint64())
		}
	}
	//遍历已经上链的交易,更新确认的块数
	for _, comfired := range e.waitComfired[chainName] {
		blockNum, err := con.EthBlockNumber()
		if err != nil {
			cmn.Logger.Errorf("[watch]Failed to EthBlockNumber: %v", err)
			return
		}
		if blockNum.ToInt().Uint64() < comfired.MinedBlock {
			cmn.Logger.Errorf("[watch]txhash:%s MinedBlock: %v > current:%v",
				comfired.TxHash, comfired.MinedBlock, blockNum.ToInt().Uint64())
			continue
		}
		c := int(blockNum.ToInt().Uint64() - comfired.MinedBlock)
		if c > cmn.Config().GetInt(chainName+".txcomfired") {
			c = cmn.Config().GetInt(chainName + ".txcomfired")
		}
		//更新交易确认数到数据库
		e.updateTransactionComfired(comfired.TxHash, c)
	}
	return
}

//监视的所有pending交易
func (e *PendingTransactionManager) watchPendingTransaction() error {
	e.watch("ethereum")
	e.watch("poa")
	e.importAllPendingTransactions("ethereum")
	e.importAllPendingTransactions("poa")
	return nil
}

// StartListeningPending 每一笔交易加入监听队列时,启动监听超时定时器
func (e *PendingTransactionManager) StartListeningPending(chainType string, txHash string) {

	ticker := time.NewTicker(time.Second * time.Duration(cmn.Config().GetInt(chainType+".txtimeout")))
	defer ticker.Stop()

	select {
	case <-ticker.C:
		e.txHashTimeout <- txHash
	}
}

// Run pending交易管理
func (e *PendingTransactionManager) Run() {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	e.importAllPendingTransactions("ethereum")
	e.importAllPendingTransactions("poa")
	for {
		select {
		case <-ticker.C:
			e.watchPendingTransaction()
		case txHash := <-e.txHashTimeout:
			e.updateTransactionTimeout(txHash, time.Now())
		case <-e.exit:
			close(e.closed)
			cmn.Logger.Info("watch alive endpoint service done!!!")
			return
		}
	}
}

// Stop Stop
func (e *PendingTransactionManager) Stop() {
	close(e.exit)
	// wait for stop
	<-e.closed
}

//AppendToPendingPool 记录待检查的交易
func AppendToPendingPool(para *PendingPoolParas) error {

	ptInfo := &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).
		Where("tx_hash = ?", para.TxHash.String()).
		Find(ptInfo).
		RecordNotFound()
	if notFound {
		ptInfo.UserID = para.UserID
		ptInfo.From = para.From.String()
		ptInfo.To = para.To.String()
		ptInfo.TxHash = para.TxHash.String()
		ptInfo.Nonce = para.Nonce
		ptInfo.Mined = false
		ptInfo.ListenTimeout = false
		ptInfo.ListenTimeoutAt = time.Now()
		ptInfo.ChainType = para.ChainType

		err := dbconn.Create(ptInfo).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	} else {
		err := dbconn.Model(ptInfo).Update("mined", false).Update("nonce", para.Nonce).Update("listen_timeout", false).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	}
	dbconn.MysqlCommit()
	//启动超时计时器
	go GetPendingTransactionManager().StartListeningPending(para.ChainType, para.TxHash.String())
	return nil
}

//IsMined 查询交易是否上链
func IsMined(txHash string) (bool, bool, uint64, int, error) {
	ptInfo := &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).
		Where("tx_hash = ?", txHash).
		Find(ptInfo).
		RecordNotFound()

	if notFound {
		cmn.Logger.Error("[query transaction is mined]no such txHash:", txHash)
		return false, false, 0, 0, errors.New("no such txHash: " + txHash)
	}

	if ptInfo.ListenTimeout == true {
		cmn.Logger.Error("[query transaction is mined]txhash:", txHash, ". timeout at:", ptInfo.ListenTimeoutAt)
		return false, false, 0, 0, errors.New("no more listen this pending transaction: " + txHash)
	}
	return ptInfo.Mined, ptInfo.Success, ptInfo.MinedBlock, ptInfo.Comfired, nil
}
