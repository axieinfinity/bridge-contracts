// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeTracking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeTrackingMetaData contains all meta data concerning the BridgeTracking contract.
var BridgeTrackingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"bulkTotalBallotsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_res\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIBridgeTracking.VoteKind\",\"name\":\"_kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"handleVoteApproved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startedAtBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIBridgeTracking.VoteKind\",\"name\":\"_kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"recordVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startedAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"totalBallots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_bridgeOperator\",\"type\":\"address\"}],\"name\":\"totalBallotsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"totalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeTrackingABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTrackingMetaData.ABI instead.
var BridgeTrackingABI = BridgeTrackingMetaData.ABI

// BridgeTracking is an auto generated Go binding around an Ethereum contract.
type BridgeTracking struct {
	BridgeTrackingCaller     // Read-only binding to the contract
	BridgeTrackingTransactor // Write-only binding to the contract
	BridgeTrackingFilterer   // Log filterer for contract events
}

// BridgeTrackingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTrackingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTrackingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTrackingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTrackingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTrackingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTrackingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTrackingSession struct {
	Contract     *BridgeTracking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTrackingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTrackingCallerSession struct {
	Contract *BridgeTrackingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeTrackingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTrackingTransactorSession struct {
	Contract     *BridgeTrackingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeTrackingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTrackingRaw struct {
	Contract *BridgeTracking // Generic contract binding to access the raw methods on
}

// BridgeTrackingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTrackingCallerRaw struct {
	Contract *BridgeTrackingCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTrackingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTrackingTransactorRaw struct {
	Contract *BridgeTrackingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTracking creates a new instance of BridgeTracking, bound to a specific deployed contract.
func NewBridgeTracking(address common.Address, backend bind.ContractBackend) (*BridgeTracking, error) {
	contract, err := bindBridgeTracking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTracking{BridgeTrackingCaller: BridgeTrackingCaller{contract: contract}, BridgeTrackingTransactor: BridgeTrackingTransactor{contract: contract}, BridgeTrackingFilterer: BridgeTrackingFilterer{contract: contract}}, nil
}

// NewBridgeTrackingCaller creates a new read-only instance of BridgeTracking, bound to a specific deployed contract.
func NewBridgeTrackingCaller(address common.Address, caller bind.ContractCaller) (*BridgeTrackingCaller, error) {
	contract, err := bindBridgeTracking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingCaller{contract: contract}, nil
}

// NewBridgeTrackingTransactor creates a new write-only instance of BridgeTracking, bound to a specific deployed contract.
func NewBridgeTrackingTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTrackingTransactor, error) {
	contract, err := bindBridgeTracking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingTransactor{contract: contract}, nil
}

// NewBridgeTrackingFilterer creates a new log filterer instance of BridgeTracking, bound to a specific deployed contract.
func NewBridgeTrackingFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTrackingFilterer, error) {
	contract, err := bindBridgeTracking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingFilterer{contract: contract}, nil
}

// bindBridgeTracking binds a generic wrapper to an already deployed contract.
func bindBridgeTracking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTrackingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTracking *BridgeTrackingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTracking.Contract.BridgeTrackingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTracking *BridgeTrackingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTracking.Contract.BridgeTrackingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTracking *BridgeTrackingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTracking.Contract.BridgeTrackingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTracking *BridgeTrackingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTracking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTracking *BridgeTrackingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTracking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTracking *BridgeTrackingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTracking.Contract.contract.Transact(opts, method, params...)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_BridgeTracking *BridgeTrackingCaller) BridgeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "bridgeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_BridgeTracking *BridgeTrackingSession) BridgeContract() (common.Address, error) {
	return _BridgeTracking.Contract.BridgeContract(&_BridgeTracking.CallOpts)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_BridgeTracking *BridgeTrackingCallerSession) BridgeContract() (common.Address, error) {
	return _BridgeTracking.Contract.BridgeContract(&_BridgeTracking.CallOpts)
}

// BulkTotalBallotsOf is a free data retrieval call binding the contract method 0x57daa170.
//
// Solidity: function bulkTotalBallotsOf(uint256 _period, address[] _bridgeOperators) view returns(uint256[] _res)
func (_BridgeTracking *BridgeTrackingCaller) BulkTotalBallotsOf(opts *bind.CallOpts, _period *big.Int, _bridgeOperators []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "bulkTotalBallotsOf", _period, _bridgeOperators)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BulkTotalBallotsOf is a free data retrieval call binding the contract method 0x57daa170.
//
// Solidity: function bulkTotalBallotsOf(uint256 _period, address[] _bridgeOperators) view returns(uint256[] _res)
func (_BridgeTracking *BridgeTrackingSession) BulkTotalBallotsOf(_period *big.Int, _bridgeOperators []common.Address) ([]*big.Int, error) {
	return _BridgeTracking.Contract.BulkTotalBallotsOf(&_BridgeTracking.CallOpts, _period, _bridgeOperators)
}

// BulkTotalBallotsOf is a free data retrieval call binding the contract method 0x57daa170.
//
// Solidity: function bulkTotalBallotsOf(uint256 _period, address[] _bridgeOperators) view returns(uint256[] _res)
func (_BridgeTracking *BridgeTrackingCallerSession) BulkTotalBallotsOf(_period *big.Int, _bridgeOperators []common.Address) ([]*big.Int, error) {
	return _BridgeTracking.Contract.BulkTotalBallotsOf(&_BridgeTracking.CallOpts, _period, _bridgeOperators)
}

// StartedAtBlock is a free data retrieval call binding the contract method 0xf84bd121.
//
// Solidity: function startedAtBlock() view returns(uint256)
func (_BridgeTracking *BridgeTrackingCaller) StartedAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "startedAtBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartedAtBlock is a free data retrieval call binding the contract method 0xf84bd121.
//
// Solidity: function startedAtBlock() view returns(uint256)
func (_BridgeTracking *BridgeTrackingSession) StartedAtBlock() (*big.Int, error) {
	return _BridgeTracking.Contract.StartedAtBlock(&_BridgeTracking.CallOpts)
}

// StartedAtBlock is a free data retrieval call binding the contract method 0xf84bd121.
//
// Solidity: function startedAtBlock() view returns(uint256)
func (_BridgeTracking *BridgeTrackingCallerSession) StartedAtBlock() (*big.Int, error) {
	return _BridgeTracking.Contract.StartedAtBlock(&_BridgeTracking.CallOpts)
}

// TotalBallots is a free data retrieval call binding the contract method 0x889998ef.
//
// Solidity: function totalBallots(uint256 _period) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCaller) TotalBallots(opts *bind.CallOpts, _period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "totalBallots", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBallots is a free data retrieval call binding the contract method 0x889998ef.
//
// Solidity: function totalBallots(uint256 _period) view returns(uint256)
func (_BridgeTracking *BridgeTrackingSession) TotalBallots(_period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallots(&_BridgeTracking.CallOpts, _period)
}

// TotalBallots is a free data retrieval call binding the contract method 0x889998ef.
//
// Solidity: function totalBallots(uint256 _period) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCallerSession) TotalBallots(_period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallots(&_BridgeTracking.CallOpts, _period)
}

// TotalBallotsOf is a free data retrieval call binding the contract method 0x04375dcf.
//
// Solidity: function totalBallotsOf(uint256 _period, address _bridgeOperator) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCaller) TotalBallotsOf(opts *bind.CallOpts, _period *big.Int, _bridgeOperator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "totalBallotsOf", _period, _bridgeOperator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBallotsOf is a free data retrieval call binding the contract method 0x04375dcf.
//
// Solidity: function totalBallotsOf(uint256 _period, address _bridgeOperator) view returns(uint256)
func (_BridgeTracking *BridgeTrackingSession) TotalBallotsOf(_period *big.Int, _bridgeOperator common.Address) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallotsOf(&_BridgeTracking.CallOpts, _period, _bridgeOperator)
}

// TotalBallotsOf is a free data retrieval call binding the contract method 0x04375dcf.
//
// Solidity: function totalBallotsOf(uint256 _period, address _bridgeOperator) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCallerSession) TotalBallotsOf(_period *big.Int, _bridgeOperator common.Address) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallotsOf(&_BridgeTracking.CallOpts, _period, _bridgeOperator)
}

// TotalVotes is a free data retrieval call binding the contract method 0x19e6e158.
//
// Solidity: function totalVotes(uint256 _period) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCaller) TotalVotes(opts *bind.CallOpts, _period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "totalVotes", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotes is a free data retrieval call binding the contract method 0x19e6e158.
//
// Solidity: function totalVotes(uint256 _period) view returns(uint256)
func (_BridgeTracking *BridgeTrackingSession) TotalVotes(_period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalVotes(&_BridgeTracking.CallOpts, _period)
}

// TotalVotes is a free data retrieval call binding the contract method 0x19e6e158.
//
// Solidity: function totalVotes(uint256 _period) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCallerSession) TotalVotes(_period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalVotes(&_BridgeTracking.CallOpts, _period)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_BridgeTracking *BridgeTrackingCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_BridgeTracking *BridgeTrackingSession) ValidatorContract() (common.Address, error) {
	return _BridgeTracking.Contract.ValidatorContract(&_BridgeTracking.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_BridgeTracking *BridgeTrackingCallerSession) ValidatorContract() (common.Address, error) {
	return _BridgeTracking.Contract.ValidatorContract(&_BridgeTracking.CallOpts)
}

// HandleVoteApproved is a paid mutator transaction binding the contract method 0x229f88ea.
//
// Solidity: function handleVoteApproved(uint8 _kind, uint256 _requestId) returns()
func (_BridgeTracking *BridgeTrackingTransactor) HandleVoteApproved(opts *bind.TransactOpts, _kind uint8, _requestId *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "handleVoteApproved", _kind, _requestId)
}

// HandleVoteApproved is a paid mutator transaction binding the contract method 0x229f88ea.
//
// Solidity: function handleVoteApproved(uint8 _kind, uint256 _requestId) returns()
func (_BridgeTracking *BridgeTrackingSession) HandleVoteApproved(_kind uint8, _requestId *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.HandleVoteApproved(&_BridgeTracking.TransactOpts, _kind, _requestId)
}

// HandleVoteApproved is a paid mutator transaction binding the contract method 0x229f88ea.
//
// Solidity: function handleVoteApproved(uint8 _kind, uint256 _requestId) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) HandleVoteApproved(_kind uint8, _requestId *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.HandleVoteApproved(&_BridgeTracking.TransactOpts, _kind, _requestId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _bridgeContract, address _validatorContract, uint256 _startedAtBlock) returns()
func (_BridgeTracking *BridgeTrackingTransactor) Initialize(opts *bind.TransactOpts, _bridgeContract common.Address, _validatorContract common.Address, _startedAtBlock *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "initialize", _bridgeContract, _validatorContract, _startedAtBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _bridgeContract, address _validatorContract, uint256 _startedAtBlock) returns()
func (_BridgeTracking *BridgeTrackingSession) Initialize(_bridgeContract common.Address, _validatorContract common.Address, _startedAtBlock *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.Initialize(&_BridgeTracking.TransactOpts, _bridgeContract, _validatorContract, _startedAtBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _bridgeContract, address _validatorContract, uint256 _startedAtBlock) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) Initialize(_bridgeContract common.Address, _validatorContract common.Address, _startedAtBlock *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.Initialize(&_BridgeTracking.TransactOpts, _bridgeContract, _validatorContract, _startedAtBlock)
}

// RecordVote is a paid mutator transaction binding the contract method 0xc7c4fea9.
//
// Solidity: function recordVote(uint8 _kind, uint256 _requestId, address _operator) returns()
func (_BridgeTracking *BridgeTrackingTransactor) RecordVote(opts *bind.TransactOpts, _kind uint8, _requestId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "recordVote", _kind, _requestId, _operator)
}

// RecordVote is a paid mutator transaction binding the contract method 0xc7c4fea9.
//
// Solidity: function recordVote(uint8 _kind, uint256 _requestId, address _operator) returns()
func (_BridgeTracking *BridgeTrackingSession) RecordVote(_kind uint8, _requestId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.RecordVote(&_BridgeTracking.TransactOpts, _kind, _requestId, _operator)
}

// RecordVote is a paid mutator transaction binding the contract method 0xc7c4fea9.
//
// Solidity: function recordVote(uint8 _kind, uint256 _requestId, address _operator) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) RecordVote(_kind uint8, _requestId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.RecordVote(&_BridgeTracking.TransactOpts, _kind, _requestId, _operator)
}

// SetBridgeContract is a paid mutator transaction binding the contract method 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _addr) returns()
func (_BridgeTracking *BridgeTrackingTransactor) SetBridgeContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "setBridgeContract", _addr)
}

// SetBridgeContract is a paid mutator transaction binding the contract method 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _addr) returns()
func (_BridgeTracking *BridgeTrackingSession) SetBridgeContract(_addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.SetBridgeContract(&_BridgeTracking.TransactOpts, _addr)
}

// SetBridgeContract is a paid mutator transaction binding the contract method 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _addr) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) SetBridgeContract(_addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.SetBridgeContract(&_BridgeTracking.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_BridgeTracking *BridgeTrackingTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_BridgeTracking *BridgeTrackingSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.SetValidatorContract(&_BridgeTracking.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.SetValidatorContract(&_BridgeTracking.TransactOpts, _addr)
}

// BridgeTrackingBridgeContractUpdatedIterator is returned from FilterBridgeContractUpdated and is used to iterate over the raw logs and unpacked data for BridgeContractUpdated events raised by the BridgeTracking contract.
type BridgeTrackingBridgeContractUpdatedIterator struct {
	Event *BridgeTrackingBridgeContractUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTrackingBridgeContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTrackingBridgeContractUpdated)
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
		it.Event = new(BridgeTrackingBridgeContractUpdated)
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
func (it *BridgeTrackingBridgeContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTrackingBridgeContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTrackingBridgeContractUpdated represents a BridgeContractUpdated event raised by the BridgeTracking contract.
type BridgeTrackingBridgeContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeContractUpdated is a free log retrieval operation binding the contract event 0x5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048.
//
// Solidity: event BridgeContractUpdated(address arg0)
func (_BridgeTracking *BridgeTrackingFilterer) FilterBridgeContractUpdated(opts *bind.FilterOpts) (*BridgeTrackingBridgeContractUpdatedIterator, error) {

	logs, sub, err := _BridgeTracking.contract.FilterLogs(opts, "BridgeContractUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingBridgeContractUpdatedIterator{contract: _BridgeTracking.contract, event: "BridgeContractUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeContractUpdated is a free log subscription operation binding the contract event 0x5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048.
//
// Solidity: event BridgeContractUpdated(address arg0)
func (_BridgeTracking *BridgeTrackingFilterer) WatchBridgeContractUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTrackingBridgeContractUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTracking.contract.WatchLogs(opts, "BridgeContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTrackingBridgeContractUpdated)
				if err := _BridgeTracking.contract.UnpackLog(event, "BridgeContractUpdated", log); err != nil {
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

// ParseBridgeContractUpdated is a log parse operation binding the contract event 0x5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048.
//
// Solidity: event BridgeContractUpdated(address arg0)
func (_BridgeTracking *BridgeTrackingFilterer) ParseBridgeContractUpdated(log types.Log) (*BridgeTrackingBridgeContractUpdated, error) {
	event := new(BridgeTrackingBridgeContractUpdated)
	if err := _BridgeTracking.contract.UnpackLog(event, "BridgeContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTrackingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeTracking contract.
type BridgeTrackingInitializedIterator struct {
	Event *BridgeTrackingInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeTrackingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTrackingInitialized)
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
		it.Event = new(BridgeTrackingInitialized)
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
func (it *BridgeTrackingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTrackingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTrackingInitialized represents a Initialized event raised by the BridgeTracking contract.
type BridgeTrackingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTracking *BridgeTrackingFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeTrackingInitializedIterator, error) {

	logs, sub, err := _BridgeTracking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingInitializedIterator{contract: _BridgeTracking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTracking *BridgeTrackingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeTrackingInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeTracking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTrackingInitialized)
				if err := _BridgeTracking.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTracking *BridgeTrackingFilterer) ParseInitialized(log types.Log) (*BridgeTrackingInitialized, error) {
	event := new(BridgeTrackingInitialized)
	if err := _BridgeTracking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTrackingValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the BridgeTracking contract.
type BridgeTrackingValidatorContractUpdatedIterator struct {
	Event *BridgeTrackingValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTrackingValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTrackingValidatorContractUpdated)
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
		it.Event = new(BridgeTrackingValidatorContractUpdated)
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
func (it *BridgeTrackingValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTrackingValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTrackingValidatorContractUpdated represents a ValidatorContractUpdated event raised by the BridgeTracking contract.
type BridgeTrackingValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_BridgeTracking *BridgeTrackingFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*BridgeTrackingValidatorContractUpdatedIterator, error) {

	logs, sub, err := _BridgeTracking.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingValidatorContractUpdatedIterator{contract: _BridgeTracking.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_BridgeTracking *BridgeTrackingFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTrackingValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTracking.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTrackingValidatorContractUpdated)
				if err := _BridgeTracking.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_BridgeTracking *BridgeTrackingFilterer) ParseValidatorContractUpdated(log types.Log) (*BridgeTrackingValidatorContractUpdated, error) {
	event := new(BridgeTrackingValidatorContractUpdated)
	if err := _BridgeTracking.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
