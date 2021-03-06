package eth

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/gorequest"
)

//ResponseBase 响应
type ResponseBase struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int64           `json:"id"`
	Error   *ObjectError    `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

//ObjectError 错误对象
type ObjectError struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ObjectError) Error() string {

	var jsonrpc2ErrorMessages = map[int64]string{
		-32700: "Parse error",
		-32600: "Invalid Request",
		-32601: "Method not found",
		-32602: "Invalid params",
		-32603: "Internal error",
		-32000: "Server error",
	}
	fmt.Printf("%d (%s) %s\n%v", e.Code, jsonrpc2ErrorMessages[e.Code], e.Message, e.Data)

	return e.Message
}

// Client rpc客户端
type Client struct {
	url        string
	httpClient *gorequest.SuperAgent
	id         int64
	idLock     sync.Mutex
}

//Request rpc请求结构
type Request struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	ID      int64         `json:"id"`
	Params  []interface{} `json:"params"`
}

//BlockNumber 区块号
type BlockNumber big.Int

func newRequest() *gorequest.SuperAgent {
	request := gorequest.New()
	request.SetDebug(cmn.Config().GetBool("ethereum.debug"))
	//request.SetLogger(cmn.Logger)
	return request
}

//NewClient 创建新的连接
func NewClient(url string) *Client {
	return &Client{
		url:        url,
		httpClient: newRequest(),
	}
}

//CallMethod 封装rpc调用接口
func (c *Client) CallMethod(v interface{}, method string, params ...interface{}) error {
	c.idLock.Lock()

	c.id++

	req := Request{
		JSONRPC: "2.0",
		ID:      c.id,
		Method:  method,
		Params:  params,
	}

	c.idLock.Unlock()

	_, bodybytes, errs := c.httpClient.Post(c.url).
		Send(req).
		EndBytes()
	if errs != nil {
		err := fmt.Errorf("CallMethod %s error: %q", method, errs)
		return err
	}

	var parsed ResponseBase
	err := json.Unmarshal(bodybytes, &parsed)
	if err != nil {
		return err
	}

	if parsed.Error != nil {
		return parsed.Error
	}

	if req.ID != parsed.ID || parsed.JSONRPC != "2.0" {
		return errors.New("Error: JSONRPC 2.0 Specification error")
	}

	return json.Unmarshal(parsed.Result, v)
}

//Web3ClientVersion 实现web3.clientVersion
func (c *Client) Web3ClientVersion() (string, error) {
	var v string

	err := c.CallMethod(&v, "web3_clientVersion")
	if err != nil {
		return "", err
	}

	return v, nil
}

//Web3Sha3 实现web3.sha3
func (c *Client) Web3Sha3(data []byte) (common.Hash, error) {
	var v common.Hash

	err := c.CallMethod(&v, "web3_sha3", data)
	if err != nil {
		return common.Hash{}, err
	}

	return v, nil
}

//NetVersion 实现web3.version
func (c *Client) NetVersion() (string, error) {
	var v string

	err := c.CallMethod(&v, "net_version")
	if err != nil {
		return "", err
	}

	return v, nil
}

//NetListening 实现web3.version
func (c *Client) NetListening() (bool, error) {
	var v bool

	err := c.CallMethod(&v, "net_Listening")
	if err != nil {
		return false, err
	}

	return v, nil
}

//NetPeerCount 实现web3.peerCount
func (c *Client) NetPeerCount() (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "net_peerCount")
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthProtocolVersion 实现web3.protocolVersion
func (c *Client) EthProtocolVersion() (string, error) {
	var v string

	err := c.CallMethod(&v, "eth_protocolVersion")
	if err != nil {
		return "", err
	}

	return v, nil
}

//EthSyncing 实现web3.syncing
func (c *Client) EthSyncing() (bool, error) {
	var v bool

	err := c.CallMethod(&v, "eth_syncing")
	if err != nil {
		return false, err
	}

	return v, nil
}

//EthCoinbase 实现web3.coinbase
func (c *Client) EthCoinbase() (common.Address, error) {
	var v common.Address

	err := c.CallMethod(&v, "eth_coinbase")
	if err != nil {
		return common.Address{}, err
	}

	return v, nil
}

//EthMining 实现web3.mining
func (c *Client) EthMining() (bool, error) {
	var v bool

	err := c.CallMethod(&v, "eth_mining")
	if err != nil {
		return false, err
	}

	return v, nil
}

//EthHashrate 实现web3.hashrate
func (c *Client) EthHashrate() (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_hashrate")
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGasPrice 实现web3.gasPrice
func (c *Client) EthGasPrice() (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_gasPrice")
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthAccounts 实现web3.accounts
func (c *Client) EthAccounts() ([]common.Address, error) {
	var v []common.Address

	err := c.CallMethod(&v, "eth_accounts")
	if err != nil {
		return nil, err
	}

	return v, nil
}

//EthBlockNumber 实现web3.blockNumber
func (c *Client) EthBlockNumber() (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_blockNumber")
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetBalance 实现web3.getBalance
func (c *Client) EthGetBalance(addr common.Address, block *hexutil.Big) (*hexutil.Big, error) {
	var v hexutil.Big
	if block == nil {
		err := c.CallMethod(&v, "eth_getBalance", addr, "latest")
		if err != nil {
			return nil, err
		}
	} else {
		err := c.CallMethod(&v, "eth_getBalance", addr, block)
		if err != nil {
			return nil, err
		}
	}

	return &v, nil
}

//EthGetStorageAt 实现web3.getStorageAt
func (c *Client) EthGetStorageAt(addr common.Address, pos, block *hexutil.Big) ([]byte, error) {
	var v hexutil.Bytes

	err := c.CallMethod(&v, "eth_getStorageAt", addr, pos, block)
	if err != nil {
		return nil, err
	}

	return v, nil
}

//TxpoolGetInspect 实现web3.txpool_inspect
func (c *Client) TxpoolGetInspect() (*TxpoolInspectResponse, error) {
	var v TxpoolInspectResponse

	err := c.CallMethod(&v, "txpool_inspect")
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetNonce 实现web3.getTransactionCount
func (c *Client) EthGetNonce(addr common.Address) (*hexutil.Big, error) {
	return c.EthGetTransactionCount(addr, "pending")
}

//EthGetTransactionCount 实现web3.getTransactionCount
func (c *Client) EthGetTransactionCount(addr common.Address, stage string) (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_getTransactionCount", addr, stage)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetTransactionCountByBlock 实现web3.getTransactionCount
func (c *Client) EthGetTransactionCountByBlock(addr common.Address, block *hexutil.Big) (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_getTransactionCount", addr, block)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetBlockTransactionCountByHash 实现web3.getBlockTransactionCountByHash
func (c *Client) EthGetBlockTransactionCountByHash(hash common.Hash) (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_getBlockTransactionCountByHash", hash)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetBlockTransactionCountByNumber 实现web3.getBlockTransactionCountByNumber
func (c *Client) EthGetBlockTransactionCountByNumber(block *hexutil.Big) (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_getBlockTransactionCountByNumber", block)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetUncleCountByHash 实现web3.getUncleCountByHash
func (c *Client) EthGetUncleCountByHash(hash common.Hash) (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_getUncleCountByHash", hash)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetUncleCountByNumber 实现web3.getUncleCountByNumber
func (c *Client) EthGetUncleCountByNumber(block *hexutil.Big) (*hexutil.Big, error) {
	var v hexutil.Big

	err := c.CallMethod(&v, "eth_getUncleCountByNumber", block)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

//EthGetCode 实现web3.getCode
func (c *Client) EthGetCode(addr common.Address, block *hexutil.Big) ([]byte, error) {
	var v hexutil.Bytes

	err := c.CallMethod(&v, "eth_getCode", addr, block)
	if err != nil {
		return nil, err
	}

	return v, nil
}

//EthSign 实现web3.getCode
func (c *Client) EthSign(addr common.Address, msg []byte) ([]byte, error) {
	var v hexutil.Bytes

	err := c.CallMethod(&v, "eth_getCode", addr, hexutil.Bytes(msg))
	if err != nil {
		return nil, err
	}

	return v, nil
}

//EthSendTransaction 实现web3.sendTransaction
func (c *Client) EthSendTransaction(req *cmn.TransactionRequest) (common.Hash, error) {
	var v common.Hash

	err := c.CallMethod(&v, "eth_sendTransaction", req)
	if err != nil {
		return common.Hash{}, err
	}

	return v, nil
}

//EthSendRawTransaction 实现web3.sendRawTransaction
func (c *Client) EthSendRawTransaction(data []byte) (common.Hash, error) {
	var v common.Hash

	err := c.CallMethod(&v, "eth_sendRawTransaction", hexutil.Bytes(data))
	if err != nil {
		return common.Hash{}, err
	}

	return v, nil
}

// func (c *Client) EthCall() ([]byte, error)                                               {}
// func (c *Client) EthEstimateGas() (big.Int, error)

//EthGetBlockByHash 实现web3.getBlockByHash
func (c *Client) EthGetBlockByHash(hash common.Hash, full bool) (common.Hash, error) {
	var v common.Hash

	err := c.CallMethod(&v, "eth_getBlockByHash", hash, full)
	if err != nil {
		return common.Hash{}, err
	}

	return v, nil
}

// func (c *Client) EthGetBlockByNumber(block BlockNumber, full bool) (*Block, error)       {}

//EthGetTransactionByHash 查询交易详情
func (c *Client) EthGetTransactionByHash(hash common.Hash) (*Transaction, error) {
	var v Transaction

	err := c.CallMethod(&v, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// func (c *Client) EthGetTransactionByBlockHashAndIndex(hash [32]byte, idx big.Int) (*Transaction, error) {}
// func (c *Client) EthGetTransactionByBlockHashAndIndex(block BlockNumber, idx big.Int) (*Transaction, error) {}

//EthGetTransactionReceipt 实现web3.getTransactionReceipt
func (c *Client) EthGetTransactionReceipt(hash common.Hash) (*TransactionReceipt, error) {
	var v TransactionReceipt

	err := c.CallMethod(&v, "eth_getTransactionReceipt", hash)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

// func (c *Client) EthGetUncleByBlockHashAndIndex(hash [32]byte, idx big.Int) (*Block, error)       {}
// func (c *Client) EthGetUncleByBlockNumberAndIndex(block BlockNumber, idx big.Int) (*Block, error) {}

//EthGetCompilers 实现web3.getCompilers
func (c *Client) EthGetCompilers() ([]string, error) {
	var v []string

	err := c.CallMethod(&v, "eth_getCompilers")
	if err != nil {
		return nil, err
	}

	return v, nil
}

//EthCompileLLL 实现web3.compileLLL
func (c *Client) EthCompileLLL(code string) ([]byte, error) {
	var v hexutil.Bytes

	err := c.CallMethod(&v, "eth_compileLLL", code)
	if err != nil {
		return nil, err
	}

	return v, nil
}

//EthCompileSolidity 实现web3.compileSolidity
func (c *Client) EthCompileSolidity(code string) ([]byte, error) {
	var v hexutil.Bytes

	err := c.CallMethod(&v, "eth_compileSolidity", code)
	if err != nil {
		return nil, err
	}

	return v, nil
}

//EthCompileSerpent 实现web3.compileSerpent
func (c *Client) EthCompileSerpent(code string) ([]byte, error) {
	var v hexutil.Bytes

	err := c.CallMethod(&v, "eth_compileSerpent", code)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// func (c *Client) EthNewFilter(fromBlock BlockNumber, toBlock BlockNumber, addrs [][20]byte, topics []byte) (big.Int, error) {}
// func (c *Client) EthNewBlockFilter() (big.Int, error)                             {}
// func (c *Client) EthNewPendingTransactionFilter() (big.Int, error)                {}
// func (c *Client) EthUninstallFilter(id big.Int) (bool, error)                     {}
// func (c *Client) EthGetFilterChanges(id big.Int) ([]*Filter, error)               {}
// func (c *Client) EthGetFilterLogs(id big.Int) ([]*Filter, error)                  {}
// func (c *Client) EthGetLogs()                                                     {}

//EthGetLogs 实现web3.work
func (c *Client) EthGetLogs() ([3]common.Hash, error) {
	var v [3]common.Hash

	err := c.CallMethod(&v, "eth_getLogs")
	if err != nil {
		return [3]common.Hash{}, err
	}

	return v, nil
}

//EthWork 实现web3.work
func (c *Client) EthWork() ([3]common.Hash, error) {
	var v [3]common.Hash

	err := c.CallMethod(&v, "eth_work")
	if err != nil {
		return [3]common.Hash{}, err
	}

	return v, nil
}

//EthSubmitWork 实现web3.submitWork
func (c *Client) EthSubmitWork(nonce [8]byte, header, mix common.Hash) (bool, error) {
	var v bool

	err := c.CallMethod(&v, "eth_submitWork", hexutil.Bytes(nonce[:]), header, mix)
	if err != nil {
		return false, err
	}

	return v, nil
}

//EthSubmitHashrate 实现web3.submitHashrate
func (c *Client) EthSubmitHashrate(hashrate, id common.Hash) (bool, error) {
	var v bool

	err := c.CallMethod(&v, "eth_submitHashrate", hashrate, id)
	if err != nil {
		return false, err
	}

	return v, nil
}

// func (c *Client) ShhVersion() (string, error) {}
// func (c *Client) ShhPost() (bool, error) {}
// func (c *Client) ShhNewIdentity() ([60]byte, error){}
// func (c *Client) ShhHasIdentity(identity [60]byte) (bool, error) {}
// func (c *Client) ShhNewGroup() ([60]byte, error) {}
// func (c *Client) ShhAddToGroup(identity [60]byte) (bool, error) {}
// func (c *Client) ShhNewFilter(identity [60]byte, topics []byte) (big.Int, error) {}
// func (c *Client) ShhUninstallFilter(filterId big.Int) (bool, error) {}
// func (c *Client) ShhGetFilterChanges(filterId big.Int) (error) {}
// func (c *Client) ShhGetMessages(filterId big.Int) ( error) {}

// TransactionReceipt 交易凭证
type TransactionReceipt struct {
	// Consensus fields
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	From              *common.Address `json:"from"`
	To                *common.Address `json:"to"`
	Status            *hexutil.Big    `json:"status"`
	CumulativeGasUsed *hexutil.Big    `json:"cumulativeGasUsed"`
	Bloom             types.Bloom     `json:"logsBloom"`
	Logs              []*Log          `json:"logs"`

	// Implementation fields (don't reorder!)
	TxHash          common.Hash     `json:"transactionHash"`
	ContractAddress *common.Address `json:"contractAddress"`
	GasUsed         *hexutil.Big    `json:"gasUsed"`
	TxIndex         *hexutil.Big    `json:"transactionIndex"`
}

// Transaction 交易凭证
type Transaction struct {
	Hash             common.Hash     `json:"hash"`
	Nonce            *hexutil.Big    `json:"nonce"`
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	TransactionIndex *hexutil.Big    `json:"transactionIndex"`
	From             *common.Address `json:"from"`
	To               *common.Address `json:"to"`
	Value            *hexutil.Big    `json:"value"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Gas              *hexutil.Big    `json:"gas"`
	Input            json.RawMessage `json:"input"`
}

// TxpoolInspectResponse 交易池中的交易
type TxpoolInspectResponse struct {
	Pending map[string]map[string][]string `json:"pending"`
	Queued  map[string]map[string][]string `json:"queued"`
}

// Log 事件存储
type Log struct {
	// Consensus fields:
	// address of the contract that generated the event
	Address *common.Address `json:"address" `
	// list of topics provided by the contract.
	Topics []common.Hash `json:"topics"`
	// supplied by the contract, usually ABI-encoded
	Data json.RawMessage `json:"data"`

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber *hexutil.Big `json:"blockNumber"`
	// hash of the transaction
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`
	// index of the transaction in the block
	TxIndex *hexutil.Big `json:"transactionIndex" gencodec:"required"`
	// hash of the block in which the transaction was included
	BlockHash common.Hash `json:"blockHash"`
	// index of the log in the receipt
	Index *hexutil.Big `json:"logIndex" gencodec:"required"`

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool `json:"removed"`
}
