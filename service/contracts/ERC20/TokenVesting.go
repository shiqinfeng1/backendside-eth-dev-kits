// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TokenVestingABI is the input ABI used to generate the binding from.
const TokenVestingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cliff\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"releasableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"vestedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"revocable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"revoked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_start\",\"type\":\"uint256\"},{\"name\":\"_cliff\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_revocable\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenVestingBin is the compiled bytecode used for deploying new contracts.
const TokenVestingBin = `0x608060405234801561001057600080fd5b5060405160a08061098a8339810160409081528151602083015191830151606084015160809094015160008054600160a060020a0319163317905591939091600160a060020a038516151561006457600080fd5b8183111561007157600080fd5b60018054600160a060020a031916600160a060020a0387161790556005805460ff191682151517905560048290556100b684846401000000006100c5810261078d1704565b600255505050600355506100d8565b818101828110156100d257fe5b92915050565b6108a3806100e76000396000f3006080604052600436106100b65763ffffffff60e060020a6000350416630fb5a6b481146100bb57806313d033c0146100e25780631726cbc8146100f75780631916558714610118578063384711cc1461013b57806338af3eed1461015c578063715018a61461018d57806374a8f103146101a2578063872a7810146101c35780638da5cb5b146101ec5780639852595c14610201578063be9a655514610222578063f2fde38b14610237578063fa01dc0614610258575b600080fd5b3480156100c757600080fd5b506100d0610279565b60408051918252519081900360200190f35b3480156100ee57600080fd5b506100d061027f565b34801561010357600080fd5b506100d0600160a060020a0360043516610285565b34801561012457600080fd5b50610139600160a060020a03600435166102bd565b005b34801561014757600080fd5b506100d0600160a060020a0360043516610369565b34801561016857600080fd5b506101716104c0565b60408051600160a060020a039092168252519081900360200190f35b34801561019957600080fd5b506101396104cf565b3480156101ae57600080fd5b50610139600160a060020a036004351661053b565b3480156101cf57600080fd5b506101d86106a2565b604080519115158252519081900360200190f35b3480156101f857600080fd5b506101716106ab565b34801561020d57600080fd5b506100d0600160a060020a03600435166106ba565b34801561022e57600080fd5b506100d06106cc565b34801561024357600080fd5b50610139600160a060020a03600435166106d2565b34801561026457600080fd5b506101d8600160a060020a0360043516610766565b60045481565b60025481565b600160a060020a0381166000908152600660205260408120546102b7906102ab84610369565b9063ffffffff61077b16565b92915050565b60006102c882610285565b9050600081116102d757600080fd5b600160a060020a038216600090815260066020526040902054610300908263ffffffff61078d16565b600160a060020a038084166000818152600660205260409020929092556001546103329291168363ffffffff61079a16565b6040805182815290517ffb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c5659181900360200190a15050565b600080600083600160a060020a03166370a08231306040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600087803b1580156103c957600080fd5b505af11580156103dd573d6000803e3d6000fd5b505050506040513d60208110156103f357600080fd5b5051600160a060020a03851660009081526006602052604090205490925061042290839063ffffffff61078d16565b905060025442101561043757600092506104b9565b60045460035461044c9163ffffffff61078d16565b421015806104725750600160a060020a03841660009081526007602052604090205460ff165b1561047f578092506104b9565b6104b66004546104aa61049d6003544261077b90919063ffffffff16565b849063ffffffff61083916565b9063ffffffff61086216565b92505b5050919050565b600154600160a060020a031681565b600054600160a060020a031633146104e657600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000805481908190600160a060020a0316331461055757600080fd5b60055460ff16151561056857600080fd5b600160a060020a03841660009081526007602052604090205460ff161561058e57600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038616916370a082319160248083019260209291908290030181600087803b1580156105ef57600080fd5b505af1158015610603573d6000803e3d6000fd5b505050506040513d602081101561061957600080fd5b5051925061062684610285565b9150610638838363ffffffff61077b16565b600160a060020a038086166000818152600760205260408120805460ff1916600117905554929350610673929091168363ffffffff61079a16565b6040517f44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee690600090a150505050565b60055460ff1681565b600054600160a060020a031681565b60066020526000908152604090205481565b60035481565b600054600160a060020a031633146106e957600080fd5b600160a060020a03811615156106fe57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60076020526000908152604090205460ff1681565b60008282111561078757fe5b50900390565b818101828110156102b757fe5b82600160a060020a031663a9059cbb83836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156107fd57600080fd5b505af1158015610811573d6000803e3d6000fd5b505050506040513d602081101561082757600080fd5b5051151561083457600080fd5b505050565b600082151561084a575060006102b7565b5081810281838281151561085a57fe5b04146102b757fe5b6000818381151561086f57fe5b0493925050505600a165627a7a723058200fd5d6fd9ff0c2d6fe5e6a8f2234aff095790f267a106174b434e14abafd96360029`

// DeployTokenVesting deploys a new Ethereum contract, binding an instance of TokenVesting to it.
func DeployTokenVesting(auth *bind.TransactOpts, backend bind.ContractBackend, _beneficiary common.Address, _start *big.Int, _cliff *big.Int, _duration *big.Int, _revocable bool) (common.Address, *types.Transaction, *TokenVesting, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenVestingBin), backend, _beneficiary, _start, _cliff, _duration, _revocable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// TokenVesting is an auto generated Go binding around an Ethereum contract.
type TokenVesting struct {
	TokenVestingCaller     // Read-only binding to the contract
	TokenVestingTransactor // Write-only binding to the contract
	TokenVestingFilterer   // Log filterer for contract events
}

// TokenVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenVestingSession struct {
	Contract     *TokenVesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenVestingCallerSession struct {
	Contract *TokenVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenVestingTransactorSession struct {
	Contract     *TokenVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenVestingRaw struct {
	Contract *TokenVesting // Generic contract binding to access the raw methods on
}

// TokenVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenVestingCallerRaw struct {
	Contract *TokenVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokenVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenVestingTransactorRaw struct {
	Contract *TokenVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenVesting creates a new instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVesting(address common.Address, backend bind.ContractBackend) (*TokenVesting, error) {
	contract, err := bindTokenVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// NewTokenVestingCaller creates a new read-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingCaller(address common.Address, caller bind.ContractCaller) (*TokenVestingCaller, error) {
	contract, err := bindTokenVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingCaller{contract: contract}, nil
}

// NewTokenVestingTransactor creates a new write-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenVestingTransactor, error) {
	contract, err := bindTokenVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingTransactor{contract: contract}, nil
}

// NewTokenVestingFilterer creates a new log filterer instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenVestingFilterer, error) {
	contract, err := bindTokenVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenVestingFilterer{contract: contract}, nil
}

// bindTokenVesting binds a generic wrapper to an already deployed contract.
func bindTokenVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.TokenVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingSession) Beneficiary() (common.Address, error) {
	return _TokenVesting.Contract.Beneficiary(&_TokenVesting.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Beneficiary() (common.Address, error) {
	return _TokenVesting.Contract.Beneficiary(&_TokenVesting.CallOpts)
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Cliff(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "cliff")
	return *ret0, err
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Cliff() (*big.Int, error) {
	return _TokenVesting.Contract.Cliff(&_TokenVesting.CallOpts)
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Cliff() (*big.Int, error) {
	return _TokenVesting.Contract.Cliff(&_TokenVesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Duration() (*big.Int, error) {
	return _TokenVesting.Contract.Duration(&_TokenVesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Duration() (*big.Int, error) {
	return _TokenVesting.Contract.Duration(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// ReleasableAmount is a free data retrieval call binding the contract method 0x1726cbc8.
//
// Solidity: function releasableAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) ReleasableAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "releasableAmount", token)
	return *ret0, err
}

// ReleasableAmount is a free data retrieval call binding the contract method 0x1726cbc8.
//
// Solidity: function releasableAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) ReleasableAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleasableAmount(&_TokenVesting.CallOpts, token)
}

// ReleasableAmount is a free data retrieval call binding the contract method 0x1726cbc8.
//
// Solidity: function releasableAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) ReleasableAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleasableAmount(&_TokenVesting.CallOpts, token)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Released(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "released", arg0)
	return *ret0, err
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Released(arg0 common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Released(&_TokenVesting.CallOpts, arg0)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released( address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Released(arg0 common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Released(&_TokenVesting.CallOpts, arg0)
}

// Revocable is a free data retrieval call binding the contract method 0x872a7810.
//
// Solidity: function revocable() constant returns(bool)
func (_TokenVesting *TokenVestingCaller) Revocable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "revocable")
	return *ret0, err
}

// Revocable is a free data retrieval call binding the contract method 0x872a7810.
//
// Solidity: function revocable() constant returns(bool)
func (_TokenVesting *TokenVestingSession) Revocable() (bool, error) {
	return _TokenVesting.Contract.Revocable(&_TokenVesting.CallOpts)
}

// Revocable is a free data retrieval call binding the contract method 0x872a7810.
//
// Solidity: function revocable() constant returns(bool)
func (_TokenVesting *TokenVestingCallerSession) Revocable() (bool, error) {
	return _TokenVesting.Contract.Revocable(&_TokenVesting.CallOpts)
}

// Revoked is a free data retrieval call binding the contract method 0xfa01dc06.
//
// Solidity: function revoked( address) constant returns(bool)
func (_TokenVesting *TokenVestingCaller) Revoked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "revoked", arg0)
	return *ret0, err
}

// Revoked is a free data retrieval call binding the contract method 0xfa01dc06.
//
// Solidity: function revoked( address) constant returns(bool)
func (_TokenVesting *TokenVestingSession) Revoked(arg0 common.Address) (bool, error) {
	return _TokenVesting.Contract.Revoked(&_TokenVesting.CallOpts, arg0)
}

// Revoked is a free data retrieval call binding the contract method 0xfa01dc06.
//
// Solidity: function revoked( address) constant returns(bool)
func (_TokenVesting *TokenVestingCallerSession) Revoked(arg0 common.Address) (bool, error) {
	return _TokenVesting.Contract.Revoked(&_TokenVesting.CallOpts, arg0)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "start")
	return *ret0, err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Start() (*big.Int, error) {
	return _TokenVesting.Contract.Start(&_TokenVesting.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Start() (*big.Int, error) {
	return _TokenVesting.Contract.Start(&_TokenVesting.CallOpts)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) VestedAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "vestedAmount", token)
	return *ret0, err
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) VestedAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, token)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(token address) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) VestedAmount(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(token address) returns()
func (_TokenVesting *TokenVestingTransactor) Release(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "release", token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(token address) returns()
func (_TokenVesting *TokenVestingSession) Release(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(token address) returns()
func (_TokenVesting *TokenVestingTransactorSession) Release(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenVesting.Contract.RenounceOwnership(&_TokenVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenVesting.Contract.RenounceOwnership(&_TokenVesting.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(token address) returns()
func (_TokenVesting *TokenVestingTransactor) Revoke(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "revoke", token)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(token address) returns()
func (_TokenVesting *TokenVestingSession) Revoke(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Revoke(&_TokenVesting.TransactOpts, token)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(token address) returns()
func (_TokenVesting *TokenVestingTransactorSession) Revoke(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Revoke(&_TokenVesting.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenVesting *TokenVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TokenVestingOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the TokenVesting contract.
type TokenVestingOwnershipRenouncedIterator struct {
	Event *TokenVestingOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *TokenVestingOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingOwnershipRenounced)
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
		it.Event = new(TokenVestingOwnershipRenounced)
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
func (it *TokenVestingOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingOwnershipRenounced represents a OwnershipRenounced event raised by the TokenVesting contract.
type TokenVestingOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*TokenVestingOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingOwnershipRenouncedIterator{contract: _TokenVesting.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *TokenVestingOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingOwnershipRenounced)
				if err := _TokenVesting.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// TokenVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenVesting contract.
type TokenVestingOwnershipTransferredIterator struct {
	Event *TokenVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingOwnershipTransferred)
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
		it.Event = new(TokenVestingOwnershipTransferred)
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
func (it *TokenVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TokenVesting contract.
type TokenVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingOwnershipTransferredIterator{contract: _TokenVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenVesting *TokenVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingOwnershipTransferred)
				if err := _TokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenVestingReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the TokenVesting contract.
type TokenVestingReleasedIterator struct {
	Event *TokenVestingReleased // Event containing the contract specifics and raw log

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
func (it *TokenVestingReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingReleased)
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
		it.Event = new(TokenVestingReleased)
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
func (it *TokenVestingReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingReleased represents a Released event raised by the TokenVesting contract.
type TokenVestingReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: e Released(amount uint256)
func (_TokenVesting *TokenVestingFilterer) FilterReleased(opts *bind.FilterOpts) (*TokenVestingReleasedIterator, error) {

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return &TokenVestingReleasedIterator{contract: _TokenVesting.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0xfb81f9b30d73d830c3544b34d827c08142579ee75710b490bab0b3995468c565.
//
// Solidity: e Released(amount uint256)
func (_TokenVesting *TokenVestingFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *TokenVestingReleased) (event.Subscription, error) {

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "Released")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingReleased)
				if err := _TokenVesting.contract.UnpackLog(event, "Released", log); err != nil {
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

// TokenVestingRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the TokenVesting contract.
type TokenVestingRevokedIterator struct {
	Event *TokenVestingRevoked // Event containing the contract specifics and raw log

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
func (it *TokenVestingRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingRevoked)
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
		it.Event = new(TokenVestingRevoked)
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
func (it *TokenVestingRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingRevoked represents a Revoked event raised by the TokenVesting contract.
type TokenVestingRevoked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: e Revoked()
func (_TokenVesting *TokenVestingFilterer) FilterRevoked(opts *bind.FilterOpts) (*TokenVestingRevokedIterator, error) {

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return &TokenVestingRevokedIterator{contract: _TokenVesting.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x44825a4b2df8acb19ce4e1afba9aa850c8b65cdb7942e2078f27d0b0960efee6.
//
// Solidity: e Revoked()
func (_TokenVesting *TokenVestingFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *TokenVestingRevoked) (event.Subscription, error) {

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "Revoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingRevoked)
				if err := _TokenVesting.contract.UnpackLog(event, "Revoked", log); err != nil {
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
