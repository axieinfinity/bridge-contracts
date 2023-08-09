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
	_ = abi.ConvertType
)

// BridgeTrackingMetaData contains all meta data concerning the BridgeTracking contract.
var BridgeTrackingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"ErrContractTypeNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumRoleAccess\",\"name\":\"expectedRole\",\"type\":\"uint8\"}],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumContractType\",\"name\":\"expectedContractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ErrUnexpectedInternalCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExternalCallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"name\":\"getManyTotalBallots\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_res\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIBridgeTracking.VoteKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"handleVoteApproved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startedAtBlock_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeSlash\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeReward\",\"type\":\"address\"}],\"name\":\"initializeV3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIBridgeTracking.VoteKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"recordVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startedAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"totalBallot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalBallot_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"totalBallotOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalVote_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := BridgeTrackingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeTracking *BridgeTrackingCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeTracking *BridgeTrackingSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeTracking.Contract.GetContract(&_BridgeTracking.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeTracking *BridgeTrackingCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeTracking.Contract.GetContract(&_BridgeTracking.CallOpts, contractType)
}

// GetManyTotalBallots is a free data retrieval call binding the contract method 0xf67e8152.
//
// Solidity: function getManyTotalBallots(uint256 period, address[] operators) view returns(uint256[] _res)
func (_BridgeTracking *BridgeTrackingCaller) GetManyTotalBallots(opts *bind.CallOpts, period *big.Int, operators []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "getManyTotalBallots", period, operators)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetManyTotalBallots is a free data retrieval call binding the contract method 0xf67e8152.
//
// Solidity: function getManyTotalBallots(uint256 period, address[] operators) view returns(uint256[] _res)
func (_BridgeTracking *BridgeTrackingSession) GetManyTotalBallots(period *big.Int, operators []common.Address) ([]*big.Int, error) {
	return _BridgeTracking.Contract.GetManyTotalBallots(&_BridgeTracking.CallOpts, period, operators)
}

// GetManyTotalBallots is a free data retrieval call binding the contract method 0xf67e8152.
//
// Solidity: function getManyTotalBallots(uint256 period, address[] operators) view returns(uint256[] _res)
func (_BridgeTracking *BridgeTrackingCallerSession) GetManyTotalBallots(period *big.Int, operators []common.Address) ([]*big.Int, error) {
	return _BridgeTracking.Contract.GetManyTotalBallots(&_BridgeTracking.CallOpts, period, operators)
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

// TotalBallot is a free data retrieval call binding the contract method 0xd25ed4c6.
//
// Solidity: function totalBallot(uint256 period) view returns(uint256 totalBallot_)
func (_BridgeTracking *BridgeTrackingCaller) TotalBallot(opts *bind.CallOpts, period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "totalBallot", period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBallot is a free data retrieval call binding the contract method 0xd25ed4c6.
//
// Solidity: function totalBallot(uint256 period) view returns(uint256 totalBallot_)
func (_BridgeTracking *BridgeTrackingSession) TotalBallot(period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallot(&_BridgeTracking.CallOpts, period)
}

// TotalBallot is a free data retrieval call binding the contract method 0xd25ed4c6.
//
// Solidity: function totalBallot(uint256 period) view returns(uint256 totalBallot_)
func (_BridgeTracking *BridgeTrackingCallerSession) TotalBallot(period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallot(&_BridgeTracking.CallOpts, period)
}

// TotalBallotOf is a free data retrieval call binding the contract method 0x4ac0bcda.
//
// Solidity: function totalBallotOf(uint256 period, address bridgeOperator) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCaller) TotalBallotOf(opts *bind.CallOpts, period *big.Int, bridgeOperator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "totalBallotOf", period, bridgeOperator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBallotOf is a free data retrieval call binding the contract method 0x4ac0bcda.
//
// Solidity: function totalBallotOf(uint256 period, address bridgeOperator) view returns(uint256)
func (_BridgeTracking *BridgeTrackingSession) TotalBallotOf(period *big.Int, bridgeOperator common.Address) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallotOf(&_BridgeTracking.CallOpts, period, bridgeOperator)
}

// TotalBallotOf is a free data retrieval call binding the contract method 0x4ac0bcda.
//
// Solidity: function totalBallotOf(uint256 period, address bridgeOperator) view returns(uint256)
func (_BridgeTracking *BridgeTrackingCallerSession) TotalBallotOf(period *big.Int, bridgeOperator common.Address) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalBallotOf(&_BridgeTracking.CallOpts, period, bridgeOperator)
}

// TotalVote is a free data retrieval call binding the contract method 0xe2a75f36.
//
// Solidity: function totalVote(uint256 period) view returns(uint256 totalVote_)
func (_BridgeTracking *BridgeTrackingCaller) TotalVote(opts *bind.CallOpts, period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTracking.contract.Call(opts, &out, "totalVote", period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVote is a free data retrieval call binding the contract method 0xe2a75f36.
//
// Solidity: function totalVote(uint256 period) view returns(uint256 totalVote_)
func (_BridgeTracking *BridgeTrackingSession) TotalVote(period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalVote(&_BridgeTracking.CallOpts, period)
}

// TotalVote is a free data retrieval call binding the contract method 0xe2a75f36.
//
// Solidity: function totalVote(uint256 period) view returns(uint256 totalVote_)
func (_BridgeTracking *BridgeTrackingCallerSession) TotalVote(period *big.Int) (*big.Int, error) {
	return _BridgeTracking.Contract.TotalVote(&_BridgeTracking.CallOpts, period)
}

// HandleVoteApproved is a paid mutator transaction binding the contract method 0x229f88ea.
//
// Solidity: function handleVoteApproved(uint8 kind, uint256 requestId) returns()
func (_BridgeTracking *BridgeTrackingTransactor) HandleVoteApproved(opts *bind.TransactOpts, kind uint8, requestId *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "handleVoteApproved", kind, requestId)
}

// HandleVoteApproved is a paid mutator transaction binding the contract method 0x229f88ea.
//
// Solidity: function handleVoteApproved(uint8 kind, uint256 requestId) returns()
func (_BridgeTracking *BridgeTrackingSession) HandleVoteApproved(kind uint8, requestId *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.HandleVoteApproved(&_BridgeTracking.TransactOpts, kind, requestId)
}

// HandleVoteApproved is a paid mutator transaction binding the contract method 0x229f88ea.
//
// Solidity: function handleVoteApproved(uint8 kind, uint256 requestId) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) HandleVoteApproved(kind uint8, requestId *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.HandleVoteApproved(&_BridgeTracking.TransactOpts, kind, requestId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address bridgeContract, address validatorContract, uint256 startedAtBlock_) returns()
func (_BridgeTracking *BridgeTrackingTransactor) Initialize(opts *bind.TransactOpts, bridgeContract common.Address, validatorContract common.Address, startedAtBlock_ *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "initialize", bridgeContract, validatorContract, startedAtBlock_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address bridgeContract, address validatorContract, uint256 startedAtBlock_) returns()
func (_BridgeTracking *BridgeTrackingSession) Initialize(bridgeContract common.Address, validatorContract common.Address, startedAtBlock_ *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.Initialize(&_BridgeTracking.TransactOpts, bridgeContract, validatorContract, startedAtBlock_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address bridgeContract, address validatorContract, uint256 startedAtBlock_) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) Initialize(bridgeContract common.Address, validatorContract common.Address, startedAtBlock_ *big.Int) (*types.Transaction, error) {
	return _BridgeTracking.Contract.Initialize(&_BridgeTracking.TransactOpts, bridgeContract, validatorContract, startedAtBlock_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_BridgeTracking *BridgeTrackingTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_BridgeTracking *BridgeTrackingSession) InitializeV2() (*types.Transaction, error) {
	return _BridgeTracking.Contract.InitializeV2(&_BridgeTracking.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _BridgeTracking.Contract.InitializeV2(&_BridgeTracking.TransactOpts)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0xe7ec7b39.
//
// Solidity: function initializeV3(address bridgeManager, address bridgeSlash, address bridgeReward) returns()
func (_BridgeTracking *BridgeTrackingTransactor) InitializeV3(opts *bind.TransactOpts, bridgeManager common.Address, bridgeSlash common.Address, bridgeReward common.Address) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "initializeV3", bridgeManager, bridgeSlash, bridgeReward)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0xe7ec7b39.
//
// Solidity: function initializeV3(address bridgeManager, address bridgeSlash, address bridgeReward) returns()
func (_BridgeTracking *BridgeTrackingSession) InitializeV3(bridgeManager common.Address, bridgeSlash common.Address, bridgeReward common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.InitializeV3(&_BridgeTracking.TransactOpts, bridgeManager, bridgeSlash, bridgeReward)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0xe7ec7b39.
//
// Solidity: function initializeV3(address bridgeManager, address bridgeSlash, address bridgeReward) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) InitializeV3(bridgeManager common.Address, bridgeSlash common.Address, bridgeReward common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.InitializeV3(&_BridgeTracking.TransactOpts, bridgeManager, bridgeSlash, bridgeReward)
}

// RecordVote is a paid mutator transaction binding the contract method 0xc7c4fea9.
//
// Solidity: function recordVote(uint8 kind, uint256 requestId, address operator) returns()
func (_BridgeTracking *BridgeTrackingTransactor) RecordVote(opts *bind.TransactOpts, kind uint8, requestId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "recordVote", kind, requestId, operator)
}

// RecordVote is a paid mutator transaction binding the contract method 0xc7c4fea9.
//
// Solidity: function recordVote(uint8 kind, uint256 requestId, address operator) returns()
func (_BridgeTracking *BridgeTrackingSession) RecordVote(kind uint8, requestId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.RecordVote(&_BridgeTracking.TransactOpts, kind, requestId, operator)
}

// RecordVote is a paid mutator transaction binding the contract method 0xc7c4fea9.
//
// Solidity: function recordVote(uint8 kind, uint256 requestId, address operator) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) RecordVote(kind uint8, requestId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.RecordVote(&_BridgeTracking.TransactOpts, kind, requestId, operator)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeTracking *BridgeTrackingTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeTracking *BridgeTrackingSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.SetContract(&_BridgeTracking.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeTracking *BridgeTrackingTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeTracking.Contract.SetContract(&_BridgeTracking.TransactOpts, contractType, addr)
}

// BridgeTrackingContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the BridgeTracking contract.
type BridgeTrackingContractUpdatedIterator struct {
	Event *BridgeTrackingContractUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTrackingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTrackingContractUpdated)
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
		it.Event = new(BridgeTrackingContractUpdated)
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
func (it *BridgeTrackingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTrackingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTrackingContractUpdated represents a ContractUpdated event raised by the BridgeTracking contract.
type BridgeTrackingContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeTracking *BridgeTrackingFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*BridgeTrackingContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeTracking.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingContractUpdatedIterator{contract: _BridgeTracking.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeTracking *BridgeTrackingFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTrackingContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeTracking.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTrackingContractUpdated)
				if err := _BridgeTracking.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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

// ParseContractUpdated is a log parse operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeTracking *BridgeTrackingFilterer) ParseContractUpdated(log types.Log) (*BridgeTrackingContractUpdated, error) {
	event := new(BridgeTrackingContractUpdated)
	if err := _BridgeTracking.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTrackingExternalCallFailedIterator is returned from FilterExternalCallFailed and is used to iterate over the raw logs and unpacked data for ExternalCallFailed events raised by the BridgeTracking contract.
type BridgeTrackingExternalCallFailedIterator struct {
	Event *BridgeTrackingExternalCallFailed // Event containing the contract specifics and raw log

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
func (it *BridgeTrackingExternalCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTrackingExternalCallFailed)
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
		it.Event = new(BridgeTrackingExternalCallFailed)
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
func (it *BridgeTrackingExternalCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTrackingExternalCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTrackingExternalCallFailed represents a ExternalCallFailed event raised by the BridgeTracking contract.
type BridgeTrackingExternalCallFailed struct {
	To     common.Address
	MsgSig [4]byte
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExternalCallFailed is a free log retrieval operation binding the contract event 0xeaa424ccc38ebcf22402729592dedf8315790e0128cb577cdeff1a3ee627f827.
//
// Solidity: event ExternalCallFailed(address indexed to, bytes4 indexed msgSig, bytes reason)
func (_BridgeTracking *BridgeTrackingFilterer) FilterExternalCallFailed(opts *bind.FilterOpts, to []common.Address, msgSig [][4]byte) (*BridgeTrackingExternalCallFailedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var msgSigRule []interface{}
	for _, msgSigItem := range msgSig {
		msgSigRule = append(msgSigRule, msgSigItem)
	}

	logs, sub, err := _BridgeTracking.contract.FilterLogs(opts, "ExternalCallFailed", toRule, msgSigRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTrackingExternalCallFailedIterator{contract: _BridgeTracking.contract, event: "ExternalCallFailed", logs: logs, sub: sub}, nil
}

// WatchExternalCallFailed is a free log subscription operation binding the contract event 0xeaa424ccc38ebcf22402729592dedf8315790e0128cb577cdeff1a3ee627f827.
//
// Solidity: event ExternalCallFailed(address indexed to, bytes4 indexed msgSig, bytes reason)
func (_BridgeTracking *BridgeTrackingFilterer) WatchExternalCallFailed(opts *bind.WatchOpts, sink chan<- *BridgeTrackingExternalCallFailed, to []common.Address, msgSig [][4]byte) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var msgSigRule []interface{}
	for _, msgSigItem := range msgSig {
		msgSigRule = append(msgSigRule, msgSigItem)
	}

	logs, sub, err := _BridgeTracking.contract.WatchLogs(opts, "ExternalCallFailed", toRule, msgSigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTrackingExternalCallFailed)
				if err := _BridgeTracking.contract.UnpackLog(event, "ExternalCallFailed", log); err != nil {
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

// ParseExternalCallFailed is a log parse operation binding the contract event 0xeaa424ccc38ebcf22402729592dedf8315790e0128cb577cdeff1a3ee627f827.
//
// Solidity: event ExternalCallFailed(address indexed to, bytes4 indexed msgSig, bytes reason)
func (_BridgeTracking *BridgeTrackingFilterer) ParseExternalCallFailed(log types.Log) (*BridgeTrackingExternalCallFailed, error) {
	event := new(BridgeTrackingExternalCallFailed)
	if err := _BridgeTracking.contract.UnpackLog(event, "ExternalCallFailed", log); err != nil {
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
