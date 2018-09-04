package contracts

import (
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/contracts/ERC20"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
)

//PollEventTransfer 等待交易上链,如果执行成功,捕获Transfer事件
func PollEventTransfer(chainName, txHash string, startBlock uint64, from, to common.Address) {
	omc, conn, err := AttachOMCToken(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachOMCToken: %v", err)
		return
	}
	defer conn.Close()
	_, success, _, _, _ := waitMinedSync(txHash)

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
		cmn.Logger.Noticef("%s transfer to %s value=%s, at %d", e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
	}
}

//PollEventMint 等待交易上链,如果执行成功,捕获Mint事件
func PollEventMint(chainName, txHash string, startBlock uint64, to common.Address) {
	points, conn, err := AttachPointCoin(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachPointCoin: %v", err)
		return
	}
	defer conn.Close()
	_, success, _, _, _ := waitMinedSync(txHash)

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
		cmn.Logger.Errorf("fail to FilterMint: %v", err)
		return
	}
	for history.Next() {
		e := history.Event
		PointsBuyComfiredToDB(e.Raw.TxHash.String(), e.To.String(), e.Amount.Uint64())
		cmn.Logger.Noticef("%s Buy Points %v at block %d", e.To.String(), e.Amount, e.Raw.BlockNumber)
	}
}

//PollEventBurn 等待交易上链,如果执行成功,捕获Burn事件
func PollEventBurn(txnType, chainName, txHash string, startBlock uint64, burner common.Address) {
	points, conn, err := AttachPointCoin(chainName)
	if err != nil {
		cmn.Logger.Errorf("Failed to AttachPointCoin: %v", err)
		return
	}
	defer conn.Close()
	_, success, _, _, _ := waitMinedSync(txHash)
	//如果交易失败,则不会有事件触发,无需监听
	if success == true {
		catchEventBurn(
			txnType,
			points,
			startBlock,
			[]common.Address{burner})
	}
}
func catchEventBurn(txnType string, points *ERC20.PointCoin, startBlock uint64, burner []common.Address) {
	//TODO: 记录捕获Transfer事件
	history, err := points.FilterBurn(&bind.FilterOpts{Start: startBlock}, burner)
	if err != nil {
		cmn.Logger.Errorf("fail to FilterBurn: %v", err)
		return
	}
	for history.Next() {
		e := history.Event
		if txnType == "consume" {
			PointsConsumeComfiredToDB(e.Raw.TxHash.Hex(), e.Burner.String(), e.Value.Uint64())
		}
		if txnType == "refund" {
			PointsRefundComfiredToDB(e.Raw.TxHash.Hex(), e.Burner.String(), e.Value.Uint64())
		}
		cmn.Logger.Noticef("%s Burn Points %v at block %d", e.Burner.String(), e.Value, e.Raw.BlockNumber)
	}
}
func waitMinedSync(txHash string) (mined bool, success bool, timeout bool, minedBlock uint64, comfired int) {
	var (
		count int
		desc  string
		err   error
	)
	timeout = false
	defer func() {
		cmn.Logger.Noticef("[waitMinedSync]Pending Desc:%v Txn: %v Status: Mined:%v Success:%v Timeout:%v minedBlock:%v comfired:%v",
			desc, txHash, mined, success, timeout, minedBlock, comfired)
	}()
	for {
		time.Sleep(time.Second * 2)
		if mined, success, minedBlock, comfired, desc, err = eth.IsMined(txHash); err != nil {
			cmn.Logger.Errorf("[waitMinedSync]transaction %v is mined fail: %v", txHash, err)
			return
		}
		if mined == true {
			return
		}
		count++
		if count > cmn.Config().GetInt("ethereum.txtimeout")/2 {
			timeout = true
			return
		}
	}
}
