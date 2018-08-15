package eth

import (
	"errors"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/kr/pretty"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
)

// PendingTransactionManager endpoints of ethereum
type PendingTransactionManager struct {
	currentPending map[string][]string
	txHashTimeout  chan string
	exit           chan bool
	closed         chan bool
}

var ptm *PendingTransactionManager

// NewPendingTransactionManager create a pending transaction manager
func NewPendingTransactionManager() *PendingTransactionManager {
	ptm = &PendingTransactionManager{
		currentPending: make(map[string][]string),
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

func (e *PendingTransactionManager) importAllPendingTransactions() {
	var ptInfo []*db.PendingTransactionInfo
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()
	dbconn.Model(&db.PendingTransactionInfo{}).Where("mined = ? and listen_timeout = ?", false, false).Find(&ptInfo)

	var currentPendingEth, currentPendingPoa []string
	for _, item := range ptInfo {
		if item.ChainType == "ethereum" {
			currentPendingEth = append(currentPendingEth, item.TxHash)
		} else if item.ChainType == "poa" {
			currentPendingPoa = append(currentPendingPoa, item.TxHash)
		} else {
			cmn.Logger.Error("Unknow transaction chaintype:", item.ChainType)
		}
	}
	e.currentPending["ethereum"] = currentPendingEth
	e.currentPending["poa"] = currentPendingPoa
}

// 更新交易的状态 mined
func (e *PendingTransactionManager) updateTransactionStatus(txHash string, status bool) error {
	var ptInfo = &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).Where("tx_hash = ?", txHash).Find(ptInfo).RecordNotFound()
	if notFound {
		cmn.Logger.Error("not found txHash:" + txHash + "in PendingTransactionInfo")
		return errors.New("not found txHash:" + txHash + "in PendingTransactionInfo")
	}
	err := dbconn.Model(ptInfo).Update("mined", status).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}

	cmn.Logger.Debugf("txHash %s update status to %v", txHash, status)
	dbconn.MysqlCommit()
	return nil
}

// 更新交易的超时 listentimeout
func (e *PendingTransactionManager) updateTransactionTimeout(txHash string, listenTimeoutAt time.Time) error {
	var ptInfo = &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).Where("tx_hash = ?", txHash).Find(ptInfo).RecordNotFound()
	if notFound {
		cmn.Logger.Error("not found txHash:" + txHash + "in PendingTransactionInfo")
		return errors.New("not found txHash:" + txHash + "in PendingTransactionInfo")
	}

	err := dbconn.Model(ptInfo).Update("listen_timeout_at", listenTimeoutAt).Update("listen_timeout", true).Error
	if err != nil {
		cmn.Logger.Error(err)
		return err
	}

	cmn.Logger.Infof("txHash %s update listenTimeoutAt %v", txHash, listenTimeoutAt)
	dbconn.MysqlCommit()
	return nil
}

//监视的所有pending交易
func (e *PendingTransactionManager) watchPendingTransaction() error {

	conEth := ConnectEthNodeForWeb3("ethereum")
	conPoa := ConnectEthNodeForWeb3("poa")

	for _, txHash := range e.currentPending["ethereum"] {
		if conEth == nil {
			break
		}
		transaction, err := conEth.EthGetTransactionByHash(ethcmn.HexToHash(txHash))
		if err != nil {
			cmn.Logger.Debug("query in ethereum txhash: " + txHash + ". error : " + err.Error())
			continue
		}
		pretty.Println("query ethereum transaction:\n", transaction)
		if transaction.BlockNumber.ToInt().Uint64() > 0 {
			e.updateTransactionStatus(txHash, true)
			break
		}
	}
	for _, txHash := range e.currentPending["poa"] {
		if conPoa == nil {
			break
		}
		transaction, err := conPoa.EthGetTransactionByHash(ethcmn.HexToHash(txHash))
		if err != nil {
			cmn.Logger.Debug("query in poa txhash: " + txHash + ". error :" + err.Error())
			continue
		}
		pretty.Println("query poa transaction:\n", transaction)
		if transaction.BlockNumber.ToInt().Uint64() > 0 {
			e.updateTransactionStatus(txHash, true)
			break
		}
	}

	e.importAllPendingTransactions()
	return nil
}

// StartListeningPending 启动超时
func (e *PendingTransactionManager) StartListeningPending(txHash string) {
	ticker := time.NewTicker(time.Second * 60) // time.Hour * 24 * 7
	defer ticker.Stop()

	select {
	case <-ticker.C:
		e.txHashTimeout <- txHash
	}
}

// Run pending交易管理
func (e *PendingTransactionManager) Run() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	e.importAllPendingTransactions()
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
func AppendToPendingPool(chainType, userID string, txHash ethcmn.Hash, from, to ethcmn.Address, nonce uint64) error {

	ptInfo := &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).
		Where("tx_hash = ?", txHash.String()).
		Find(ptInfo).
		RecordNotFound()
	if notFound {
		ptInfo.UserID = userID
		ptInfo.From = from.String()
		ptInfo.To = to.String()
		ptInfo.TxHash = txHash.String()
		ptInfo.Nonce = nonce
		ptInfo.Mined = false
		ptInfo.ListenTimeout = false
		ptInfo.ListenTimeoutAt = time.Now()
		ptInfo.ChainType = chainType

		err := dbconn.Create(ptInfo).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	} else {
		err := dbconn.Model(ptInfo).Update("mined", false).Update("nonce", nonce).Update("listen_timeout", false).Error
		if err != nil {
			cmn.Logger.Error(err)
			return err
		}
	}
	dbconn.MysqlCommit()
	//启动超时
	go GetPendingTransactionManager().StartListeningPending(txHash.String())
	return nil
}

//IsMined 查询交易是否上链
func IsMined(txHash string) (bool, error) {
	ptInfo := &db.PendingTransactionInfo{}
	dbconn := db.MysqlBegin()
	defer dbconn.MysqlRollback()

	notFound := dbconn.Model(&db.PendingTransactionInfo{}).
		Where("tx_hash = ?", txHash).
		Find(ptInfo).
		RecordNotFound()

	if notFound {
		cmn.Logger.Error("[query transaction is mined]no such txHash:", txHash)
		return false, errors.New("no such txHash: " + txHash)
	}

	if ptInfo.ListenTimeout == true {
		cmn.Logger.Error("[query transaction is mined]txhash:", txHash, ". timeout at:", ptInfo.ListenTimeoutAt)
		return false, errors.New("no more listen this pending transaction: " + txHash)
	}
	return ptInfo.Mined, nil
}
