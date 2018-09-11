package eth

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

//NonceData nonce管理数据
type NonceData struct {
	NonceAvailableValue *big.Int //可用的nonce值
	NonceLastUpdate     time.Time
	NonceLock           *sync.Mutex
}

// NonceManage nonce管理结构
type NonceManage struct {
	Nonces              map[string]NonceData
	NonceUpdateInterval time.Duration
	ChainName           string
}

var nonceManage *NonceManage

// NewNonceManage 新建NewNonceManage
func NewNonceManage(chainName string, interval uint64) *NonceManage {
	nonceManage = &NonceManage{
		Nonces: make(map[string]NonceData),
	}
	nonceManage.NonceUpdateInterval = time.Duration(interval) * time.Second
	nonceManage.ChainName = chainName
	return nonceManage
}

//NonceNeedsUpdate 检查是否更新nonce
func (n *NonceManage) NonceNeedsUpdate(addr string) bool {
	if _, ok := n.Nonces[addr]; !ok { //不存在
		return true
	}
	//超过更新周期
	return n.Nonces[addr].NonceLastUpdate.Add(n.NonceUpdateInterval).After(time.Now())
}

//NonceUpdateFromNode 更新nonce
func (n *NonceManage) NonceUpdateFromNode(addr string) error {
	var newNonce NonceData
	var nonce = new(big.Int)
	nonce.SetInt64(-2)
	if _, ok := n.Nonces[addr]; !ok { //不存在
		newNonce.NonceAvailableValue = new(big.Int).SetInt64(-1)
		newNonce.NonceLastUpdate = time.Now()
		newNonce.NonceLock = new(sync.Mutex)
	} else {
		newNonce = n.Nonces[addr]
	}

	con := ConnectEthNodeForWeb3(n.ChainName)
	if con == nil {
		err := fmt.Errorf("Connect Eth Node Fail")
		cmn.Logger.Errorf(err.Error())
		return err
	}
	// Wait until all tx are registered as pending
	for {
		if nonce.Cmp(newNonce.NonceAvailableValue) == -1 { // nonce < newNonce.NonceAvailableValue)
			nn, _ := con.EthGetNonce(ethcmn.HexToAddress(addr)) //获取指定地址的nonce
			nonce = nn.ToInt()
		} else {
			break
		}
		if newNonce.NonceLastUpdate.Add(n.NonceUpdateInterval * 5).Before(time.Now()) {
			cmn.Logger.Error("Get Nonce Fail: timeout")
			break
		}
	}
	if queuedNonce, err := getMinQueueedNonce(con, addr); err == nil {
		cmn.Logger.Debugf("Get Pending Nonce:%v Queued Nonce:%v", nonce, queuedNonce)
	}
	newNonce.NonceLastUpdate = time.Now()
	newNonce.NonceAvailableValue = nonce
	n.Nonces[addr] = newNonce
	return nil
}

func getMinQueueedNonce(con *Client, userAddress string) (*[]int, error) {
	queuedTxns, err := inpsectTxpool(con, userAddress)
	if err != nil {
		return nil, err
	}
	var keys []int
	for K := range *queuedTxns {
		if iK, err := strconv.Atoi(K); err == nil {
			keys = append(keys, iK)
		}
	}
	sort.Ints(keys)

	return &keys, nil

}
func inpsectTxpool(con *Client, userAddress string) (*map[string][]string, error) {
	inspect, err := con.TxpoolGetInspect()
	if err != nil {
		return nil, err
	}
	var queuedTxn *map[string][]string
	for addr, txns := range inspect.Queued {
		if addr == userAddress {
			queuedTxn = &txns
			return queuedTxn, nil
		}
	}
	return nil, errors.New("no queued transactions")

}

//GetNonce 获取nonce并自增
func GetNonce(addr string) (*big.Int, error) {
	//如果没有保存该地址的nonce, 从节点获取nonce
	if _, ok := nonceManage.Nonces[addr]; !ok {
		if err := nonceManage.NonceUpdateFromNode(addr); err != nil {
			return new(big.Int).SetInt64(-1), err
		}
		return nonceManage.Nonces[addr].NonceAvailableValue, nil
	}
	nm := nonceManage.Nonces[addr]
	nm.NonceLock.Lock()
	defer nm.NonceLock.Unlock()
	//已保存地址的nonce,判断是否需要更新
	if nonceManage.NonceNeedsUpdate(addr) == true {
		if err := nonceManage.NonceUpdateFromNode(addr); err != nil {
			return new(big.Int).SetInt64(-1), err
		}
	}
	n := nm.NonceAvailableValue
	nm.NonceAvailableValue.Add(nm.NonceAvailableValue, new(big.Int).SetInt64(1))
	return n, nil
}

//CurrentNonce 获取当前nonce
func CurrentNonce(chainName, addr string) (int64, error) {
	//如果没有保存该地址的nonce, 从节点获取nonce
	if _, ok := nonceManage.Nonces[addr]; !ok {
		n, err := ConnectEthNodeForWeb3(chainName).EthGetNonce(ethcmn.HexToAddress(addr))
		if err != nil {
			return -1, err
		}
		return n.ToInt().Int64(), nil
	}
	nm := nonceManage.Nonces[addr]
	return nm.NonceAvailableValue.Int64() - 1, nil
}
