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

// PointCoinABI is the input ABI used to generate the binding from.
const PointCoinABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PointCoinBin is the compiled bytecode used for deploying new contracts.
const PointCoinBin = `0x608060405260038054600160a860020a03191633179055610c6d806100256000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100df578063095ea7b31461010857806318160ddd1461012c57806323b872dd1461015357806340c10f191461017d57806342966c68146101a157806366188463146101bb57806370a08231146101df578063715018a6146102005780637d64bcb4146102155780638da5cb5b1461022a578063a9059cbb1461025b578063d73dd6231461027f578063dd62ed3e146102a3578063f2fde38b146102ca575b600080fd5b3480156100eb57600080fd5b506100f46102eb565b604080519115158252519081900360200190f35b34801561011457600080fd5b506100f4600160a060020a036004351660243561030c565b34801561013857600080fd5b50610141610372565b60408051918252519081900360200190f35b34801561015f57600080fd5b506100f4600160a060020a0360043581169060243516604435610378565b34801561018957600080fd5b506100f4600160a060020a03600435166024356104dd565b3480156101ad57600080fd5b506101b96004356105e6565b005b3480156101c757600080fd5b506100f4600160a060020a03600435166024356105f3565b3480156101eb57600080fd5b50610141600160a060020a03600435166106e3565b34801561020c57600080fd5b506101b96106fe565b34801561022157600080fd5b506100f461076c565b34801561023657600080fd5b5061023f610812565b60408051600160a060020a039092168252519081900360200190f35b34801561026757600080fd5b506100f4600160a060020a0360043516602435610821565b34801561028b57600080fd5b506100f4600160a060020a03600435166024356109b4565b3480156102af57600080fd5b50610141600160a060020a0360043581169060243516610a4d565b3480156102d657600080fd5b506101b9600160a060020a0360043516610a78565b60035474010000000000000000000000000000000000000000900460ff1681565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b6000600160a060020a038316151561038f57600080fd5b600160a060020a0384166000908152602081905260409020548211156103b457600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156103e457600080fd5b600160a060020a03841660009081526020819052604090205461040d908363ffffffff610b0d16565b600160a060020a038086166000908152602081905260408082209390935590851681522054610442908363ffffffff610b1f16565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610484908363ffffffff610b0d16565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610c22833981519152929181900390910190a35060019392505050565b600354600090600160a060020a031633146104f757600080fd5b60035474010000000000000000000000000000000000000000900460ff161561051f57600080fd5b600154610532908363ffffffff610b1f16565b600155600160a060020a03831660009081526020819052604090205461055e908363ffffffff610b1f16565b600160a060020a03841660008181526020818152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a03851691600091600080516020610c228339815191529181900360200190a350600192915050565b6105f03382610b32565b50565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561064857336000908152600260209081526040808320600160a060020a038816845290915281205561067d565b610658818463ffffffff610b0d16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b600354600160a060020a0316331461071557600080fd5b600354604051600160a060020a03909116907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26003805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354600090600160a060020a0316331461078657600080fd5b60035474010000000000000000000000000000000000000000900460ff16156107ae57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b600354600160a060020a031681565b6000600160a060020a038316151561089a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f5f746f206164647265737320697320302e000000000000000000000000000000604482015290519081900360640190fd5b3360009081526020819052604090205482111561091857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f496e73756666696369656e742062616c616e63652e0000000000000000000000604482015290519081900360640190fd5b33600090815260208190526040902054610938908363ffffffff610b0d16565b3360009081526020819052604080822092909255600160a060020a0385168152205461096a908363ffffffff610b1f16565b600160a060020a03841660008181526020818152604091829020939093558051858152905191923392600080516020610c228339815191529281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546109e8908363ffffffff610b1f16565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a03163314610a8f57600080fd5b600160a060020a0381161515610aa457600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610b1957fe5b50900390565b81810182811015610b2c57fe5b92915050565b600160a060020a038216600090815260208190526040902054811115610b5757600080fd5b600160a060020a038216600090815260208190526040902054610b80908263ffffffff610b0d16565b600160a060020a038316600090815260208190526040902055600154610bac908263ffffffff610b0d16565b600155604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020610c228339815191529181900360200190a350505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a7230582067962ae395b8ca346ce478035d09339b0d6309c45a28b18a35e3bc10a2961a130029`

// DeployPointCoin deploys a new Ethereum contract, binding an instance of PointCoin to it.
func DeployPointCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PointCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(PointCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PointCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PointCoin{PointCoinCaller: PointCoinCaller{contract: contract}, PointCoinTransactor: PointCoinTransactor{contract: contract}, PointCoinFilterer: PointCoinFilterer{contract: contract}}, nil
}

// PointCoin is an auto generated Go binding around an Ethereum contract.
type PointCoin struct {
	PointCoinCaller     // Read-only binding to the contract
	PointCoinTransactor // Write-only binding to the contract
	PointCoinFilterer   // Log filterer for contract events
}

// PointCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type PointCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PointCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PointCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PointCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PointCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PointCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PointCoinSession struct {
	Contract     *PointCoin        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PointCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PointCoinCallerSession struct {
	Contract *PointCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PointCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PointCoinTransactorSession struct {
	Contract     *PointCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PointCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type PointCoinRaw struct {
	Contract *PointCoin // Generic contract binding to access the raw methods on
}

// PointCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PointCoinCallerRaw struct {
	Contract *PointCoinCaller // Generic read-only contract binding to access the raw methods on
}

// PointCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PointCoinTransactorRaw struct {
	Contract *PointCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPointCoin creates a new instance of PointCoin, bound to a specific deployed contract.
func NewPointCoin(address common.Address, backend bind.ContractBackend) (*PointCoin, error) {
	contract, err := bindPointCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PointCoin{PointCoinCaller: PointCoinCaller{contract: contract}, PointCoinTransactor: PointCoinTransactor{contract: contract}, PointCoinFilterer: PointCoinFilterer{contract: contract}}, nil
}

// NewPointCoinCaller creates a new read-only instance of PointCoin, bound to a specific deployed contract.
func NewPointCoinCaller(address common.Address, caller bind.ContractCaller) (*PointCoinCaller, error) {
	contract, err := bindPointCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PointCoinCaller{contract: contract}, nil
}

// NewPointCoinTransactor creates a new write-only instance of PointCoin, bound to a specific deployed contract.
func NewPointCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*PointCoinTransactor, error) {
	contract, err := bindPointCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PointCoinTransactor{contract: contract}, nil
}

// NewPointCoinFilterer creates a new log filterer instance of PointCoin, bound to a specific deployed contract.
func NewPointCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*PointCoinFilterer, error) {
	contract, err := bindPointCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PointCoinFilterer{contract: contract}, nil
}

// bindPointCoin binds a generic wrapper to an already deployed contract.
func bindPointCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PointCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PointCoin *PointCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PointCoin.Contract.PointCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PointCoin *PointCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PointCoin.Contract.PointCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PointCoin *PointCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PointCoin.Contract.PointCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PointCoin *PointCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PointCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PointCoin *PointCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PointCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PointCoin *PointCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PointCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PointCoin *PointCoinCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PointCoin.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PointCoin *PointCoinSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PointCoin.Contract.Allowance(&_PointCoin.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PointCoin *PointCoinCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PointCoin.Contract.Allowance(&_PointCoin.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PointCoin *PointCoinCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PointCoin.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PointCoin *PointCoinSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PointCoin.Contract.BalanceOf(&_PointCoin.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PointCoin *PointCoinCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PointCoin.Contract.BalanceOf(&_PointCoin.CallOpts, _owner)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_PointCoin *PointCoinCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PointCoin.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_PointCoin *PointCoinSession) MintingFinished() (bool, error) {
	return _PointCoin.Contract.MintingFinished(&_PointCoin.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_PointCoin *PointCoinCallerSession) MintingFinished() (bool, error) {
	return _PointCoin.Contract.MintingFinished(&_PointCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PointCoin *PointCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PointCoin.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PointCoin *PointCoinSession) Owner() (common.Address, error) {
	return _PointCoin.Contract.Owner(&_PointCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PointCoin *PointCoinCallerSession) Owner() (common.Address, error) {
	return _PointCoin.Contract.Owner(&_PointCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PointCoin *PointCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PointCoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PointCoin *PointCoinSession) TotalSupply() (*big.Int, error) {
	return _PointCoin.Contract.TotalSupply(&_PointCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PointCoin *PointCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _PointCoin.Contract.TotalSupply(&_PointCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PointCoin *PointCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PointCoin *PointCoinSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Approve(&_PointCoin.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PointCoin *PointCoinTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Approve(&_PointCoin.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_PointCoin *PointCoinTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_PointCoin *PointCoinSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Burn(&_PointCoin.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_PointCoin *PointCoinTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Burn(&_PointCoin.TransactOpts, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_PointCoin *PointCoinTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_PointCoin *PointCoinSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.DecreaseApproval(&_PointCoin.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_PointCoin *PointCoinTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.DecreaseApproval(&_PointCoin.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_PointCoin *PointCoinTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_PointCoin *PointCoinSession) FinishMinting() (*types.Transaction, error) {
	return _PointCoin.Contract.FinishMinting(&_PointCoin.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_PointCoin *PointCoinTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _PointCoin.Contract.FinishMinting(&_PointCoin.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_PointCoin *PointCoinTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_PointCoin *PointCoinSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.IncreaseApproval(&_PointCoin.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_PointCoin *PointCoinTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.IncreaseApproval(&_PointCoin.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_PointCoin *PointCoinTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_PointCoin *PointCoinSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Mint(&_PointCoin.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_PointCoin *PointCoinTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Mint(&_PointCoin.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PointCoin *PointCoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PointCoin *PointCoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _PointCoin.Contract.RenounceOwnership(&_PointCoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PointCoin *PointCoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PointCoin.Contract.RenounceOwnership(&_PointCoin.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PointCoin *PointCoinTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PointCoin *PointCoinSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Transfer(&_PointCoin.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PointCoin *PointCoinTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.Transfer(&_PointCoin.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PointCoin *PointCoinTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PointCoin *PointCoinSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.TransferFrom(&_PointCoin.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PointCoin *PointCoinTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PointCoin.Contract.TransferFrom(&_PointCoin.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PointCoin *PointCoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PointCoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PointCoin *PointCoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PointCoin.Contract.TransferOwnership(&_PointCoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PointCoin *PointCoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PointCoin.Contract.TransferOwnership(&_PointCoin.TransactOpts, newOwner)
}

// PointCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PointCoin contract.
type PointCoinApprovalIterator struct {
	Event *PointCoinApproval // Event containing the contract specifics and raw log

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
func (it *PointCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinApproval)
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
		it.Event = new(PointCoinApproval)
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
func (it *PointCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinApproval represents a Approval event raised by the PointCoin contract.
type PointCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PointCoin *PointCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PointCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PointCoinApprovalIterator{contract: _PointCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PointCoin *PointCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PointCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinApproval)
				if err := _PointCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PointCoinBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the PointCoin contract.
type PointCoinBurnIterator struct {
	Event *PointCoinBurn // Event containing the contract specifics and raw log

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
func (it *PointCoinBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinBurn)
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
		it.Event = new(PointCoinBurn)
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
func (it *PointCoinBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinBurn represents a Burn event raised by the PointCoin contract.
type PointCoinBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_PointCoin *PointCoinFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*PointCoinBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &PointCoinBurnIterator{contract: _PointCoin.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_PointCoin *PointCoinFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PointCoinBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinBurn)
				if err := _PointCoin.contract.UnpackLog(event, "Burn", log); err != nil {
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

// PointCoinMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PointCoin contract.
type PointCoinMintIterator struct {
	Event *PointCoinMint // Event containing the contract specifics and raw log

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
func (it *PointCoinMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinMint)
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
		it.Event = new(PointCoinMint)
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
func (it *PointCoinMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinMint represents a Mint event raised by the PointCoin contract.
type PointCoinMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_PointCoin *PointCoinFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*PointCoinMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &PointCoinMintIterator{contract: _PointCoin.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_PointCoin *PointCoinFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PointCoinMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinMint)
				if err := _PointCoin.contract.UnpackLog(event, "Mint", log); err != nil {
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

// PointCoinMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the PointCoin contract.
type PointCoinMintFinishedIterator struct {
	Event *PointCoinMintFinished // Event containing the contract specifics and raw log

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
func (it *PointCoinMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinMintFinished)
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
		it.Event = new(PointCoinMintFinished)
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
func (it *PointCoinMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinMintFinished represents a MintFinished event raised by the PointCoin contract.
type PointCoinMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_PointCoin *PointCoinFilterer) FilterMintFinished(opts *bind.FilterOpts) (*PointCoinMintFinishedIterator, error) {

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &PointCoinMintFinishedIterator{contract: _PointCoin.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_PointCoin *PointCoinFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *PointCoinMintFinished) (event.Subscription, error) {

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinMintFinished)
				if err := _PointCoin.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// PointCoinOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PointCoin contract.
type PointCoinOwnershipRenouncedIterator struct {
	Event *PointCoinOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PointCoinOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinOwnershipRenounced)
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
		it.Event = new(PointCoinOwnershipRenounced)
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
func (it *PointCoinOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinOwnershipRenounced represents a OwnershipRenounced event raised by the PointCoin contract.
type PointCoinOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PointCoin *PointCoinFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PointCoinOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PointCoinOwnershipRenouncedIterator{contract: _PointCoin.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PointCoin *PointCoinFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PointCoinOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinOwnershipRenounced)
				if err := _PointCoin.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PointCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PointCoin contract.
type PointCoinOwnershipTransferredIterator struct {
	Event *PointCoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PointCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinOwnershipTransferred)
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
		it.Event = new(PointCoinOwnershipTransferred)
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
func (it *PointCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinOwnershipTransferred represents a OwnershipTransferred event raised by the PointCoin contract.
type PointCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PointCoin *PointCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PointCoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PointCoinOwnershipTransferredIterator{contract: _PointCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PointCoin *PointCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PointCoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinOwnershipTransferred)
				if err := _PointCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PointCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PointCoin contract.
type PointCoinTransferIterator struct {
	Event *PointCoinTransfer // Event containing the contract specifics and raw log

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
func (it *PointCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointCoinTransfer)
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
		it.Event = new(PointCoinTransfer)
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
func (it *PointCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointCoinTransfer represents a Transfer event raised by the PointCoin contract.
type PointCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PointCoin *PointCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PointCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PointCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PointCoinTransferIterator{contract: _PointCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PointCoin *PointCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PointCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PointCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointCoinTransfer)
				if err := _PointCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
