// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// OMCABI is the input ABI used to generate the binding from.
const OMCABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// OMCBin is the compiled bytecode used for deploying new contracts.
const OMCBin = `0x60806040526005805461010060a860020a031916741dcef12e93b0abf2d36f723e8b59cc762775d5130017905560068054731dcef12e93b0abf2d36f723e8b59cc762775d513600160a060020a03199182168117909255600780548216831790556008805482168317905560098054821683179055600a8054821683179055600b8054909116909117905534801561009657600080fd5b50604080518082018252600d81527f4f4d436861696e20546f6b656e000000000000000000000000000000000000006020808301918252835180850190945260038085527f4f4d43000000000000000000000000000000000000000000000000000000000091850191909152825192939260069261011492916101e7565b5081516101289060049060208501906101e7565b506005805460ff191660ff92909216919091179081905566038d7ea4c680006001556101009004600160a060020a03908116600090815260208190526040808220655af3107a40009081905560065484168352818320652d79883d20009055600754841683528183206512309ce54000905560085484168352818320651b48eb57e00090556009548416835281832066016bcc41e900009055600a548416835281832055600b54909216815220660110d9316ec0009055506102829050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061022857805160ff1916838001178555610255565b82800160010185558215610255579182015b8281111561025557825182559160200191906001019061023a565b50610261929150610265565b5090565b61027f91905b80821115610261576000815560010161026b565b90565b610a88806102916000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100c9578063095ea7b31461015357806318160ddd1461018b57806323b872dd146101b2578063313ce567146101dc57806342966c6814610207578063661884631461022157806370a082311461024557806379cc67901461026657806395d89b411461028a578063a9059cbb1461029f578063d73dd623146102c3578063dd62ed3e146102e7575b600080fd5b3480156100d557600080fd5b506100de61030e565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610118578181015183820152602001610100565b50505050905090810190601f1680156101455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015f57600080fd5b50610177600160a060020a036004351660243561039c565b604080519115158252519081900360200190f35b34801561019757600080fd5b506101a0610402565b60408051918252519081900360200190f35b3480156101be57600080fd5b50610177600160a060020a0360043581169060243516604435610408565b3480156101e857600080fd5b506101f161057f565b6040805160ff9092168252519081900360200190f35b34801561021357600080fd5b5061021f600435610588565b005b34801561022d57600080fd5b50610177600160a060020a0360043516602435610595565b34801561025157600080fd5b506101a0600160a060020a0360043516610685565b34801561027257600080fd5b5061021f600160a060020a03600435166024356106a0565b34801561029657600080fd5b506100de610736565b3480156102ab57600080fd5b50610177600160a060020a0360043516602435610791565b3480156102cf57600080fd5b50610177600160a060020a0360043516602435610872565b3480156102f357600080fd5b506101a0600160a060020a036004358116906024351661090b565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103945780601f1061036957610100808354040283529160200191610394565b820191906000526020600020905b81548152906001019060200180831161037757829003601f168201915b505050505081565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b6000600160a060020a038316151561041f57600080fd5b600160a060020a03841660009081526020819052604090205482111561044457600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561047457600080fd5b600160a060020a03841660009081526020819052604090205461049d908363ffffffff61093616565b600160a060020a0380861660009081526020819052604080822093909355908516815220546104d2908363ffffffff61094816565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610514908363ffffffff61093616565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b60055460ff1681565b610592338261095b565b50565b336000908152600260209081526040808320600160a060020a0386168452909152812054808311156105ea57336000908152600260209081526040808320600160a060020a038816845290915281205561061f565b6105fa818463ffffffff61093616565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b600160a060020a03821660009081526002602090815260408083203384529091529020548111156106d057600080fd5b600160a060020a0382166000908152600260209081526040808320338452909152902054610704908263ffffffff61093616565b600160a060020a0383166000908152600260209081526040808320338452909152902055610732828261095b565b5050565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103945780601f1061036957610100808354040283529160200191610394565b6000600160a060020a03831615156107a857600080fd5b336000908152602081905260409020548211156107c457600080fd5b336000908152602081905260409020546107e4908363ffffffff61093616565b3360009081526020819052604080822092909255600160a060020a03851681522054610816908363ffffffff61094816565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546108a6908363ffffffff61094816565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561094257fe5b50900390565b8181018281101561095557fe5b92915050565b600160a060020a03821660009081526020819052604090205481111561098057600080fd5b600160a060020a0382166000908152602081905260409020546109a9908263ffffffff61093616565b600160a060020a0383166000908152602081905260409020556001546109d5908263ffffffff61093616565b600155604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350505600a165627a7a723058200ccedb442f4369ec1adfa6132b41249c48bef0337c8c60dbe9f2baa49ff3640c0029`

// DeployOMC deploys a new Ethereum contract, binding an instance of OMC to it.
func DeployOMC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OMC, error) {
	parsed, err := abi.JSON(strings.NewReader(OMCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OMCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OMC{OMCCaller: OMCCaller{contract: contract}, OMCTransactor: OMCTransactor{contract: contract}, OMCFilterer: OMCFilterer{contract: contract}}, nil
}

// OMC is an auto generated Go binding around an Ethereum contract.
type OMC struct {
	OMCCaller     // Read-only binding to the contract
	OMCTransactor // Write-only binding to the contract
	OMCFilterer   // Log filterer for contract events
}

// OMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type OMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OMCSession struct {
	Contract     *OMC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OMCCallerSession struct {
	Contract *OMCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OMCTransactorSession struct {
	Contract     *OMCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type OMCRaw struct {
	Contract *OMC // Generic contract binding to access the raw methods on
}

// OMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OMCCallerRaw struct {
	Contract *OMCCaller // Generic read-only contract binding to access the raw methods on
}

// OMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OMCTransactorRaw struct {
	Contract *OMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOMC creates a new instance of OMC, bound to a specific deployed contract.
func NewOMC(address common.Address, backend bind.ContractBackend) (*OMC, error) {
	contract, err := bindOMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OMC{OMCCaller: OMCCaller{contract: contract}, OMCTransactor: OMCTransactor{contract: contract}, OMCFilterer: OMCFilterer{contract: contract}}, nil
}

// NewOMCCaller creates a new read-only instance of OMC, bound to a specific deployed contract.
func NewOMCCaller(address common.Address, caller bind.ContractCaller) (*OMCCaller, error) {
	contract, err := bindOMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OMCCaller{contract: contract}, nil
}

// NewOMCTransactor creates a new write-only instance of OMC, bound to a specific deployed contract.
func NewOMCTransactor(address common.Address, transactor bind.ContractTransactor) (*OMCTransactor, error) {
	contract, err := bindOMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OMCTransactor{contract: contract}, nil
}

// NewOMCFilterer creates a new log filterer instance of OMC, bound to a specific deployed contract.
func NewOMCFilterer(address common.Address, filterer bind.ContractFilterer) (*OMCFilterer, error) {
	contract, err := bindOMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OMCFilterer{contract: contract}, nil
}

// bindOMC binds a generic wrapper to an already deployed contract.
func bindOMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OMC *OMCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OMC.Contract.OMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OMC *OMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OMC.Contract.OMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OMC *OMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OMC.Contract.OMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OMC *OMCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OMC *OMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OMC *OMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OMC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_OMC *OMCCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OMC.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_OMC *OMCSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OMC.Contract.Allowance(&_OMC.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_OMC *OMCCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _OMC.Contract.Allowance(&_OMC.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_OMC *OMCCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OMC.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_OMC *OMCSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _OMC.Contract.BalanceOf(&_OMC.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_OMC *OMCCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _OMC.Contract.BalanceOf(&_OMC.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_OMC *OMCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OMC.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_OMC *OMCSession) Decimals() (uint8, error) {
	return _OMC.Contract.Decimals(&_OMC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_OMC *OMCCallerSession) Decimals() (uint8, error) {
	return _OMC.Contract.Decimals(&_OMC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_OMC *OMCCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OMC.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_OMC *OMCSession) Name() (string, error) {
	return _OMC.Contract.Name(&_OMC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_OMC *OMCCallerSession) Name() (string, error) {
	return _OMC.Contract.Name(&_OMC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_OMC *OMCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OMC.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_OMC *OMCSession) Symbol() (string, error) {
	return _OMC.Contract.Symbol(&_OMC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_OMC *OMCCallerSession) Symbol() (string, error) {
	return _OMC.Contract.Symbol(&_OMC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_OMC *OMCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OMC.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_OMC *OMCSession) TotalSupply() (*big.Int, error) {
	return _OMC.Contract.TotalSupply(&_OMC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_OMC *OMCCallerSession) TotalSupply() (*big.Int, error) {
	return _OMC.Contract.TotalSupply(&_OMC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_OMC *OMCTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_OMC *OMCSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.Approve(&_OMC.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_OMC *OMCTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.Approve(&_OMC.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_OMC *OMCTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_OMC *OMCSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.Burn(&_OMC.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_OMC *OMCTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.Burn(&_OMC.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns()
func (_OMC *OMCTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns()
func (_OMC *OMCSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.BurnFrom(&_OMC.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns()
func (_OMC *OMCTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.BurnFrom(&_OMC.TransactOpts, _from, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_OMC *OMCTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_OMC *OMCSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.DecreaseApproval(&_OMC.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_OMC *OMCTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.DecreaseApproval(&_OMC.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_OMC *OMCTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_OMC *OMCSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.IncreaseApproval(&_OMC.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_OMC *OMCTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.IncreaseApproval(&_OMC.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_OMC *OMCTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_OMC *OMCSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.Transfer(&_OMC.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_OMC *OMCTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.Transfer(&_OMC.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_OMC *OMCTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_OMC *OMCSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.TransferFrom(&_OMC.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_OMC *OMCTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _OMC.Contract.TransferFrom(&_OMC.TransactOpts, _from, _to, _value)
}

// OMCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OMC contract.
type OMCApprovalIterator struct {
	Event *OMCApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OMCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OMCApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OMCApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OMCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OMCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OMCApproval represents a Approval event raised by the OMC contract.
type OMCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_OMC *OMCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OMCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OMC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OMCApprovalIterator{contract: _OMC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_OMC *OMCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OMCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OMC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OMCApproval)
				if err := _OMC.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OMCBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the OMC contract.
type OMCBurnIterator struct {
	Event *OMCBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OMCBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OMCBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OMCBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OMCBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OMCBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OMCBurn represents a Burn event raised by the OMC contract.
type OMCBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_OMC *OMCFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*OMCBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _OMC.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &OMCBurnIterator{contract: _OMC.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_OMC *OMCFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *OMCBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _OMC.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OMCBurn)
				if err := _OMC.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OMCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OMC contract.
type OMCTransferIterator struct {
	Event *OMCTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OMCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OMCTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OMCTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OMCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OMCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OMCTransfer represents a Transfer event raised by the OMC contract.
type OMCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_OMC *OMCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OMCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OMC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OMCTransferIterator{contract: _OMC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_OMC *OMCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OMCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OMC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OMCTransfer)
				if err := _OMC.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
