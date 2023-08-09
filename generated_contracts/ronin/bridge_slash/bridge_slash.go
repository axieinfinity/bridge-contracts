// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeSlash

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

// BridgeSlashMetaData contains all meta data concerning the BridgeSlash contract.
var BridgeSlashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"ErrContractTypeNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumRoleAccess\",\"name\":\"expectedRole\",\"type\":\"uint8\"}],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumContractType\",\"name\":\"expectedContractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ErrUnexpectedInternalCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BridgeTrackingIncorrectlyResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"RemovalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumIBridgeSlashEvents.Tier\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashUntilPeriod\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINIMUM_VOTE_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REMOVE_DURATION_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIER_1_PENALTY_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIER_2_PENALTY_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allBridgeOperators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ballots\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"totalBallot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"execSlashBridgeOperators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"slashed\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"getAddedPeriodOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"addedPeriods\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPenaltyDurations\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"penaltyDurations\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ballot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVote\",\"type\":\"uint256\"}],\"name\":\"getSlashTier\",\"outputs\":[{\"internalType\":\"enumIBridgeSlashEvents.Tier\",\"name\":\"tier\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"getSlashUntilPeriodOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"untilPeriods\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeTrackingContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentBridgeOperator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newBridgeOperator\",\"type\":\"address\"}],\"name\":\"onBridgeOperatorUpdated\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"addeds\",\"type\":\"bool[]\"}],\"name\":\"onBridgeOperatorsAdded\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"name\":\"onBridgeOperatorsRemoved\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BridgeSlashABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeSlashMetaData.ABI instead.
var BridgeSlashABI = BridgeSlashMetaData.ABI

// BridgeSlash is an auto generated Go binding around an Ethereum contract.
type BridgeSlash struct {
	BridgeSlashCaller     // Read-only binding to the contract
	BridgeSlashTransactor // Write-only binding to the contract
	BridgeSlashFilterer   // Log filterer for contract events
}

// BridgeSlashCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeSlashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSlashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeSlashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSlashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeSlashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSlashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSlashSession struct {
	Contract     *BridgeSlash      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeSlashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeSlashCallerSession struct {
	Contract *BridgeSlashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeSlashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeSlashTransactorSession struct {
	Contract     *BridgeSlashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeSlashRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeSlashRaw struct {
	Contract *BridgeSlash // Generic contract binding to access the raw methods on
}

// BridgeSlashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeSlashCallerRaw struct {
	Contract *BridgeSlashCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeSlashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeSlashTransactorRaw struct {
	Contract *BridgeSlashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeSlash creates a new instance of BridgeSlash, bound to a specific deployed contract.
func NewBridgeSlash(address common.Address, backend bind.ContractBackend) (*BridgeSlash, error) {
	contract, err := bindBridgeSlash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeSlash{BridgeSlashCaller: BridgeSlashCaller{contract: contract}, BridgeSlashTransactor: BridgeSlashTransactor{contract: contract}, BridgeSlashFilterer: BridgeSlashFilterer{contract: contract}}, nil
}

// NewBridgeSlashCaller creates a new read-only instance of BridgeSlash, bound to a specific deployed contract.
func NewBridgeSlashCaller(address common.Address, caller bind.ContractCaller) (*BridgeSlashCaller, error) {
	contract, err := bindBridgeSlash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeSlashCaller{contract: contract}, nil
}

// NewBridgeSlashTransactor creates a new write-only instance of BridgeSlash, bound to a specific deployed contract.
func NewBridgeSlashTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeSlashTransactor, error) {
	contract, err := bindBridgeSlash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeSlashTransactor{contract: contract}, nil
}

// NewBridgeSlashFilterer creates a new log filterer instance of BridgeSlash, bound to a specific deployed contract.
func NewBridgeSlashFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeSlashFilterer, error) {
	contract, err := bindBridgeSlash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeSlashFilterer{contract: contract}, nil
}

// bindBridgeSlash binds a generic wrapper to an already deployed contract.
func bindBridgeSlash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeSlashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeSlash *BridgeSlashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeSlash.Contract.BridgeSlashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeSlash *BridgeSlashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeSlash.Contract.BridgeSlashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeSlash *BridgeSlashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeSlash.Contract.BridgeSlashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeSlash *BridgeSlashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeSlash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeSlash *BridgeSlashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeSlash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeSlash *BridgeSlashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeSlash.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMVOTETHRESHOLD is a free data retrieval call binding the contract method 0x9c2f4459.
//
// Solidity: function MINIMUM_VOTE_THRESHOLD() view returns(uint256)
func (_BridgeSlash *BridgeSlashCaller) MINIMUMVOTETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "MINIMUM_VOTE_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMVOTETHRESHOLD is a free data retrieval call binding the contract method 0x9c2f4459.
//
// Solidity: function MINIMUM_VOTE_THRESHOLD() view returns(uint256)
func (_BridgeSlash *BridgeSlashSession) MINIMUMVOTETHRESHOLD() (*big.Int, error) {
	return _BridgeSlash.Contract.MINIMUMVOTETHRESHOLD(&_BridgeSlash.CallOpts)
}

// MINIMUMVOTETHRESHOLD is a free data retrieval call binding the contract method 0x9c2f4459.
//
// Solidity: function MINIMUM_VOTE_THRESHOLD() view returns(uint256)
func (_BridgeSlash *BridgeSlashCallerSession) MINIMUMVOTETHRESHOLD() (*big.Int, error) {
	return _BridgeSlash.Contract.MINIMUMVOTETHRESHOLD(&_BridgeSlash.CallOpts)
}

// REMOVEDURATIONTHRESHOLD is a free data retrieval call binding the contract method 0xddc3f7f6.
//
// Solidity: function REMOVE_DURATION_THRESHOLD() view returns(uint256)
func (_BridgeSlash *BridgeSlashCaller) REMOVEDURATIONTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "REMOVE_DURATION_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REMOVEDURATIONTHRESHOLD is a free data retrieval call binding the contract method 0xddc3f7f6.
//
// Solidity: function REMOVE_DURATION_THRESHOLD() view returns(uint256)
func (_BridgeSlash *BridgeSlashSession) REMOVEDURATIONTHRESHOLD() (*big.Int, error) {
	return _BridgeSlash.Contract.REMOVEDURATIONTHRESHOLD(&_BridgeSlash.CallOpts)
}

// REMOVEDURATIONTHRESHOLD is a free data retrieval call binding the contract method 0xddc3f7f6.
//
// Solidity: function REMOVE_DURATION_THRESHOLD() view returns(uint256)
func (_BridgeSlash *BridgeSlashCallerSession) REMOVEDURATIONTHRESHOLD() (*big.Int, error) {
	return _BridgeSlash.Contract.REMOVEDURATIONTHRESHOLD(&_BridgeSlash.CallOpts)
}

// TIER1PENALTYDURATION is a free data retrieval call binding the contract method 0x6dda4408.
//
// Solidity: function TIER_1_PENALTY_DURATION() view returns(uint256)
func (_BridgeSlash *BridgeSlashCaller) TIER1PENALTYDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "TIER_1_PENALTY_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIER1PENALTYDURATION is a free data retrieval call binding the contract method 0x6dda4408.
//
// Solidity: function TIER_1_PENALTY_DURATION() view returns(uint256)
func (_BridgeSlash *BridgeSlashSession) TIER1PENALTYDURATION() (*big.Int, error) {
	return _BridgeSlash.Contract.TIER1PENALTYDURATION(&_BridgeSlash.CallOpts)
}

// TIER1PENALTYDURATION is a free data retrieval call binding the contract method 0x6dda4408.
//
// Solidity: function TIER_1_PENALTY_DURATION() view returns(uint256)
func (_BridgeSlash *BridgeSlashCallerSession) TIER1PENALTYDURATION() (*big.Int, error) {
	return _BridgeSlash.Contract.TIER1PENALTYDURATION(&_BridgeSlash.CallOpts)
}

// TIER2PENALTYDURATION is a free data retrieval call binding the contract method 0xd1e1f2f8.
//
// Solidity: function TIER_2_PENALTY_DURATION() view returns(uint256)
func (_BridgeSlash *BridgeSlashCaller) TIER2PENALTYDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "TIER_2_PENALTY_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIER2PENALTYDURATION is a free data retrieval call binding the contract method 0xd1e1f2f8.
//
// Solidity: function TIER_2_PENALTY_DURATION() view returns(uint256)
func (_BridgeSlash *BridgeSlashSession) TIER2PENALTYDURATION() (*big.Int, error) {
	return _BridgeSlash.Contract.TIER2PENALTYDURATION(&_BridgeSlash.CallOpts)
}

// TIER2PENALTYDURATION is a free data retrieval call binding the contract method 0xd1e1f2f8.
//
// Solidity: function TIER_2_PENALTY_DURATION() view returns(uint256)
func (_BridgeSlash *BridgeSlashCallerSession) TIER2PENALTYDURATION() (*big.Int, error) {
	return _BridgeSlash.Contract.TIER2PENALTYDURATION(&_BridgeSlash.CallOpts)
}

// GetAddedPeriodOf is a free data retrieval call binding the contract method 0x1288810a.
//
// Solidity: function getAddedPeriodOf(address[] bridgeOperators) view returns(uint256[] addedPeriods)
func (_BridgeSlash *BridgeSlashCaller) GetAddedPeriodOf(opts *bind.CallOpts, bridgeOperators []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "getAddedPeriodOf", bridgeOperators)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAddedPeriodOf is a free data retrieval call binding the contract method 0x1288810a.
//
// Solidity: function getAddedPeriodOf(address[] bridgeOperators) view returns(uint256[] addedPeriods)
func (_BridgeSlash *BridgeSlashSession) GetAddedPeriodOf(bridgeOperators []common.Address) ([]*big.Int, error) {
	return _BridgeSlash.Contract.GetAddedPeriodOf(&_BridgeSlash.CallOpts, bridgeOperators)
}

// GetAddedPeriodOf is a free data retrieval call binding the contract method 0x1288810a.
//
// Solidity: function getAddedPeriodOf(address[] bridgeOperators) view returns(uint256[] addedPeriods)
func (_BridgeSlash *BridgeSlashCallerSession) GetAddedPeriodOf(bridgeOperators []common.Address) ([]*big.Int, error) {
	return _BridgeSlash.Contract.GetAddedPeriodOf(&_BridgeSlash.CallOpts, bridgeOperators)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeSlash *BridgeSlashCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeSlash *BridgeSlashSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeSlash.Contract.GetContract(&_BridgeSlash.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeSlash *BridgeSlashCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeSlash.Contract.GetContract(&_BridgeSlash.CallOpts, contractType)
}

// GetPenaltyDurations is a free data retrieval call binding the contract method 0xf9f60873.
//
// Solidity: function getPenaltyDurations() pure returns(uint256[] penaltyDurations)
func (_BridgeSlash *BridgeSlashCaller) GetPenaltyDurations(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "getPenaltyDurations")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPenaltyDurations is a free data retrieval call binding the contract method 0xf9f60873.
//
// Solidity: function getPenaltyDurations() pure returns(uint256[] penaltyDurations)
func (_BridgeSlash *BridgeSlashSession) GetPenaltyDurations() ([]*big.Int, error) {
	return _BridgeSlash.Contract.GetPenaltyDurations(&_BridgeSlash.CallOpts)
}

// GetPenaltyDurations is a free data retrieval call binding the contract method 0xf9f60873.
//
// Solidity: function getPenaltyDurations() pure returns(uint256[] penaltyDurations)
func (_BridgeSlash *BridgeSlashCallerSession) GetPenaltyDurations() ([]*big.Int, error) {
	return _BridgeSlash.Contract.GetPenaltyDurations(&_BridgeSlash.CallOpts)
}

// GetSlashTier is a free data retrieval call binding the contract method 0xfbb2f194.
//
// Solidity: function getSlashTier(uint256 ballot, uint256 totalVote) pure returns(uint8 tier)
func (_BridgeSlash *BridgeSlashCaller) GetSlashTier(opts *bind.CallOpts, ballot *big.Int, totalVote *big.Int) (uint8, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "getSlashTier", ballot, totalVote)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSlashTier is a free data retrieval call binding the contract method 0xfbb2f194.
//
// Solidity: function getSlashTier(uint256 ballot, uint256 totalVote) pure returns(uint8 tier)
func (_BridgeSlash *BridgeSlashSession) GetSlashTier(ballot *big.Int, totalVote *big.Int) (uint8, error) {
	return _BridgeSlash.Contract.GetSlashTier(&_BridgeSlash.CallOpts, ballot, totalVote)
}

// GetSlashTier is a free data retrieval call binding the contract method 0xfbb2f194.
//
// Solidity: function getSlashTier(uint256 ballot, uint256 totalVote) pure returns(uint8 tier)
func (_BridgeSlash *BridgeSlashCallerSession) GetSlashTier(ballot *big.Int, totalVote *big.Int) (uint8, error) {
	return _BridgeSlash.Contract.GetSlashTier(&_BridgeSlash.CallOpts, ballot, totalVote)
}

// GetSlashUntilPeriodOf is a free data retrieval call binding the contract method 0x5311153b.
//
// Solidity: function getSlashUntilPeriodOf(address[] bridgeOperators) view returns(uint256[] untilPeriods)
func (_BridgeSlash *BridgeSlashCaller) GetSlashUntilPeriodOf(opts *bind.CallOpts, bridgeOperators []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "getSlashUntilPeriodOf", bridgeOperators)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSlashUntilPeriodOf is a free data retrieval call binding the contract method 0x5311153b.
//
// Solidity: function getSlashUntilPeriodOf(address[] bridgeOperators) view returns(uint256[] untilPeriods)
func (_BridgeSlash *BridgeSlashSession) GetSlashUntilPeriodOf(bridgeOperators []common.Address) ([]*big.Int, error) {
	return _BridgeSlash.Contract.GetSlashUntilPeriodOf(&_BridgeSlash.CallOpts, bridgeOperators)
}

// GetSlashUntilPeriodOf is a free data retrieval call binding the contract method 0x5311153b.
//
// Solidity: function getSlashUntilPeriodOf(address[] bridgeOperators) view returns(uint256[] untilPeriods)
func (_BridgeSlash *BridgeSlashCallerSession) GetSlashUntilPeriodOf(bridgeOperators []common.Address) ([]*big.Int, error) {
	return _BridgeSlash.Contract.GetSlashUntilPeriodOf(&_BridgeSlash.CallOpts, bridgeOperators)
}

// OnBridgeOperatorsRemoved is a free data retrieval call binding the contract method 0xc48549de.
//
// Solidity: function onBridgeOperatorsRemoved(address[] , bool[] ) view returns(bytes4)
func (_BridgeSlash *BridgeSlashCaller) OnBridgeOperatorsRemoved(opts *bind.CallOpts, arg0 []common.Address, arg1 []bool) ([4]byte, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "onBridgeOperatorsRemoved", arg0, arg1)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnBridgeOperatorsRemoved is a free data retrieval call binding the contract method 0xc48549de.
//
// Solidity: function onBridgeOperatorsRemoved(address[] , bool[] ) view returns(bytes4)
func (_BridgeSlash *BridgeSlashSession) OnBridgeOperatorsRemoved(arg0 []common.Address, arg1 []bool) ([4]byte, error) {
	return _BridgeSlash.Contract.OnBridgeOperatorsRemoved(&_BridgeSlash.CallOpts, arg0, arg1)
}

// OnBridgeOperatorsRemoved is a free data retrieval call binding the contract method 0xc48549de.
//
// Solidity: function onBridgeOperatorsRemoved(address[] , bool[] ) view returns(bytes4)
func (_BridgeSlash *BridgeSlashCallerSession) OnBridgeOperatorsRemoved(arg0 []common.Address, arg1 []bool) ([4]byte, error) {
	return _BridgeSlash.Contract.OnBridgeOperatorsRemoved(&_BridgeSlash.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_BridgeSlash *BridgeSlashCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeSlash.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_BridgeSlash *BridgeSlashSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeSlash.Contract.SupportsInterface(&_BridgeSlash.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_BridgeSlash *BridgeSlashCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeSlash.Contract.SupportsInterface(&_BridgeSlash.CallOpts, interfaceId)
}

// ExecSlashBridgeOperators is a paid mutator transaction binding the contract method 0x4dca5925.
//
// Solidity: function execSlashBridgeOperators(address[] allBridgeOperators, uint256[] ballots, uint256 totalBallot, uint256 totalVote, uint256 period) returns(bool slashed)
func (_BridgeSlash *BridgeSlashTransactor) ExecSlashBridgeOperators(opts *bind.TransactOpts, allBridgeOperators []common.Address, ballots []*big.Int, totalBallot *big.Int, totalVote *big.Int, period *big.Int) (*types.Transaction, error) {
	return _BridgeSlash.contract.Transact(opts, "execSlashBridgeOperators", allBridgeOperators, ballots, totalBallot, totalVote, period)
}

// ExecSlashBridgeOperators is a paid mutator transaction binding the contract method 0x4dca5925.
//
// Solidity: function execSlashBridgeOperators(address[] allBridgeOperators, uint256[] ballots, uint256 totalBallot, uint256 totalVote, uint256 period) returns(bool slashed)
func (_BridgeSlash *BridgeSlashSession) ExecSlashBridgeOperators(allBridgeOperators []common.Address, ballots []*big.Int, totalBallot *big.Int, totalVote *big.Int, period *big.Int) (*types.Transaction, error) {
	return _BridgeSlash.Contract.ExecSlashBridgeOperators(&_BridgeSlash.TransactOpts, allBridgeOperators, ballots, totalBallot, totalVote, period)
}

// ExecSlashBridgeOperators is a paid mutator transaction binding the contract method 0x4dca5925.
//
// Solidity: function execSlashBridgeOperators(address[] allBridgeOperators, uint256[] ballots, uint256 totalBallot, uint256 totalVote, uint256 period) returns(bool slashed)
func (_BridgeSlash *BridgeSlashTransactorSession) ExecSlashBridgeOperators(allBridgeOperators []common.Address, ballots []*big.Int, totalBallot *big.Int, totalVote *big.Int, period *big.Int) (*types.Transaction, error) {
	return _BridgeSlash.Contract.ExecSlashBridgeOperators(&_BridgeSlash.TransactOpts, allBridgeOperators, ballots, totalBallot, totalVote, period)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address validatorContract, address bridgeManagerContract, address bridgeTrackingContract) returns()
func (_BridgeSlash *BridgeSlashTransactor) Initialize(opts *bind.TransactOpts, validatorContract common.Address, bridgeManagerContract common.Address, bridgeTrackingContract common.Address) (*types.Transaction, error) {
	return _BridgeSlash.contract.Transact(opts, "initialize", validatorContract, bridgeManagerContract, bridgeTrackingContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address validatorContract, address bridgeManagerContract, address bridgeTrackingContract) returns()
func (_BridgeSlash *BridgeSlashSession) Initialize(validatorContract common.Address, bridgeManagerContract common.Address, bridgeTrackingContract common.Address) (*types.Transaction, error) {
	return _BridgeSlash.Contract.Initialize(&_BridgeSlash.TransactOpts, validatorContract, bridgeManagerContract, bridgeTrackingContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address validatorContract, address bridgeManagerContract, address bridgeTrackingContract) returns()
func (_BridgeSlash *BridgeSlashTransactorSession) Initialize(validatorContract common.Address, bridgeManagerContract common.Address, bridgeTrackingContract common.Address) (*types.Transaction, error) {
	return _BridgeSlash.Contract.Initialize(&_BridgeSlash.TransactOpts, validatorContract, bridgeManagerContract, bridgeTrackingContract)
}

// OnBridgeOperatorUpdated is a paid mutator transaction binding the contract method 0xc9631a12.
//
// Solidity: function onBridgeOperatorUpdated(address currentBridgeOperator, address newBridgeOperator) returns(bytes4)
func (_BridgeSlash *BridgeSlashTransactor) OnBridgeOperatorUpdated(opts *bind.TransactOpts, currentBridgeOperator common.Address, newBridgeOperator common.Address) (*types.Transaction, error) {
	return _BridgeSlash.contract.Transact(opts, "onBridgeOperatorUpdated", currentBridgeOperator, newBridgeOperator)
}

// OnBridgeOperatorUpdated is a paid mutator transaction binding the contract method 0xc9631a12.
//
// Solidity: function onBridgeOperatorUpdated(address currentBridgeOperator, address newBridgeOperator) returns(bytes4)
func (_BridgeSlash *BridgeSlashSession) OnBridgeOperatorUpdated(currentBridgeOperator common.Address, newBridgeOperator common.Address) (*types.Transaction, error) {
	return _BridgeSlash.Contract.OnBridgeOperatorUpdated(&_BridgeSlash.TransactOpts, currentBridgeOperator, newBridgeOperator)
}

// OnBridgeOperatorUpdated is a paid mutator transaction binding the contract method 0xc9631a12.
//
// Solidity: function onBridgeOperatorUpdated(address currentBridgeOperator, address newBridgeOperator) returns(bytes4)
func (_BridgeSlash *BridgeSlashTransactorSession) OnBridgeOperatorUpdated(currentBridgeOperator common.Address, newBridgeOperator common.Address) (*types.Transaction, error) {
	return _BridgeSlash.Contract.OnBridgeOperatorUpdated(&_BridgeSlash.TransactOpts, currentBridgeOperator, newBridgeOperator)
}

// OnBridgeOperatorsAdded is a paid mutator transaction binding the contract method 0x5ebae8a0.
//
// Solidity: function onBridgeOperatorsAdded(address[] bridgeOperators, bool[] addeds) returns(bytes4)
func (_BridgeSlash *BridgeSlashTransactor) OnBridgeOperatorsAdded(opts *bind.TransactOpts, bridgeOperators []common.Address, addeds []bool) (*types.Transaction, error) {
	return _BridgeSlash.contract.Transact(opts, "onBridgeOperatorsAdded", bridgeOperators, addeds)
}

// OnBridgeOperatorsAdded is a paid mutator transaction binding the contract method 0x5ebae8a0.
//
// Solidity: function onBridgeOperatorsAdded(address[] bridgeOperators, bool[] addeds) returns(bytes4)
func (_BridgeSlash *BridgeSlashSession) OnBridgeOperatorsAdded(bridgeOperators []common.Address, addeds []bool) (*types.Transaction, error) {
	return _BridgeSlash.Contract.OnBridgeOperatorsAdded(&_BridgeSlash.TransactOpts, bridgeOperators, addeds)
}

// OnBridgeOperatorsAdded is a paid mutator transaction binding the contract method 0x5ebae8a0.
//
// Solidity: function onBridgeOperatorsAdded(address[] bridgeOperators, bool[] addeds) returns(bytes4)
func (_BridgeSlash *BridgeSlashTransactorSession) OnBridgeOperatorsAdded(bridgeOperators []common.Address, addeds []bool) (*types.Transaction, error) {
	return _BridgeSlash.Contract.OnBridgeOperatorsAdded(&_BridgeSlash.TransactOpts, bridgeOperators, addeds)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeSlash *BridgeSlashTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeSlash.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeSlash *BridgeSlashSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeSlash.Contract.SetContract(&_BridgeSlash.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeSlash *BridgeSlashTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeSlash.Contract.SetContract(&_BridgeSlash.TransactOpts, contractType, addr)
}

// BridgeSlashBridgeTrackingIncorrectlyRespondedIterator is returned from FilterBridgeTrackingIncorrectlyResponded and is used to iterate over the raw logs and unpacked data for BridgeTrackingIncorrectlyResponded events raised by the BridgeSlash contract.
type BridgeSlashBridgeTrackingIncorrectlyRespondedIterator struct {
	Event *BridgeSlashBridgeTrackingIncorrectlyResponded // Event containing the contract specifics and raw log

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
func (it *BridgeSlashBridgeTrackingIncorrectlyRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSlashBridgeTrackingIncorrectlyResponded)
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
		it.Event = new(BridgeSlashBridgeTrackingIncorrectlyResponded)
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
func (it *BridgeSlashBridgeTrackingIncorrectlyRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSlashBridgeTrackingIncorrectlyRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSlashBridgeTrackingIncorrectlyResponded represents a BridgeTrackingIncorrectlyResponded event raised by the BridgeSlash contract.
type BridgeSlashBridgeTrackingIncorrectlyResponded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBridgeTrackingIncorrectlyResponded is a free log retrieval operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_BridgeSlash *BridgeSlashFilterer) FilterBridgeTrackingIncorrectlyResponded(opts *bind.FilterOpts) (*BridgeSlashBridgeTrackingIncorrectlyRespondedIterator, error) {

	logs, sub, err := _BridgeSlash.contract.FilterLogs(opts, "BridgeTrackingIncorrectlyResponded")
	if err != nil {
		return nil, err
	}
	return &BridgeSlashBridgeTrackingIncorrectlyRespondedIterator{contract: _BridgeSlash.contract, event: "BridgeTrackingIncorrectlyResponded", logs: logs, sub: sub}, nil
}

// WatchBridgeTrackingIncorrectlyResponded is a free log subscription operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_BridgeSlash *BridgeSlashFilterer) WatchBridgeTrackingIncorrectlyResponded(opts *bind.WatchOpts, sink chan<- *BridgeSlashBridgeTrackingIncorrectlyResponded) (event.Subscription, error) {

	logs, sub, err := _BridgeSlash.contract.WatchLogs(opts, "BridgeTrackingIncorrectlyResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSlashBridgeTrackingIncorrectlyResponded)
				if err := _BridgeSlash.contract.UnpackLog(event, "BridgeTrackingIncorrectlyResponded", log); err != nil {
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

// ParseBridgeTrackingIncorrectlyResponded is a log parse operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_BridgeSlash *BridgeSlashFilterer) ParseBridgeTrackingIncorrectlyResponded(log types.Log) (*BridgeSlashBridgeTrackingIncorrectlyResponded, error) {
	event := new(BridgeSlashBridgeTrackingIncorrectlyResponded)
	if err := _BridgeSlash.contract.UnpackLog(event, "BridgeTrackingIncorrectlyResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSlashContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the BridgeSlash contract.
type BridgeSlashContractUpdatedIterator struct {
	Event *BridgeSlashContractUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeSlashContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSlashContractUpdated)
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
		it.Event = new(BridgeSlashContractUpdated)
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
func (it *BridgeSlashContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSlashContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSlashContractUpdated represents a ContractUpdated event raised by the BridgeSlash contract.
type BridgeSlashContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeSlash *BridgeSlashFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*BridgeSlashContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeSlash.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSlashContractUpdatedIterator{contract: _BridgeSlash.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeSlash *BridgeSlashFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *BridgeSlashContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeSlash.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSlashContractUpdated)
				if err := _BridgeSlash.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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
func (_BridgeSlash *BridgeSlashFilterer) ParseContractUpdated(log types.Log) (*BridgeSlashContractUpdated, error) {
	event := new(BridgeSlashContractUpdated)
	if err := _BridgeSlash.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSlashInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeSlash contract.
type BridgeSlashInitializedIterator struct {
	Event *BridgeSlashInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeSlashInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSlashInitialized)
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
		it.Event = new(BridgeSlashInitialized)
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
func (it *BridgeSlashInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSlashInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSlashInitialized represents a Initialized event raised by the BridgeSlash contract.
type BridgeSlashInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeSlash *BridgeSlashFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeSlashInitializedIterator, error) {

	logs, sub, err := _BridgeSlash.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeSlashInitializedIterator{contract: _BridgeSlash.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeSlash *BridgeSlashFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeSlashInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeSlash.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSlashInitialized)
				if err := _BridgeSlash.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeSlash *BridgeSlashFilterer) ParseInitialized(log types.Log) (*BridgeSlashInitialized, error) {
	event := new(BridgeSlashInitialized)
	if err := _BridgeSlash.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSlashRemovalRequestedIterator is returned from FilterRemovalRequested and is used to iterate over the raw logs and unpacked data for RemovalRequested events raised by the BridgeSlash contract.
type BridgeSlashRemovalRequestedIterator struct {
	Event *BridgeSlashRemovalRequested // Event containing the contract specifics and raw log

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
func (it *BridgeSlashRemovalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSlashRemovalRequested)
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
		it.Event = new(BridgeSlashRemovalRequested)
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
func (it *BridgeSlashRemovalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSlashRemovalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSlashRemovalRequested represents a RemovalRequested event raised by the BridgeSlash contract.
type BridgeSlashRemovalRequested struct {
	Period         *big.Int
	BridgeOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemovalRequested is a free log retrieval operation binding the contract event 0xb32a150b9737190a456d8b2b81dd7d592a799ab2933ea494e44351acd41f835d.
//
// Solidity: event RemovalRequested(uint256 indexed period, address indexed bridgeOperator)
func (_BridgeSlash *BridgeSlashFilterer) FilterRemovalRequested(opts *bind.FilterOpts, period []*big.Int, bridgeOperator []common.Address) (*BridgeSlashRemovalRequestedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}

	logs, sub, err := _BridgeSlash.contract.FilterLogs(opts, "RemovalRequested", periodRule, bridgeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSlashRemovalRequestedIterator{contract: _BridgeSlash.contract, event: "RemovalRequested", logs: logs, sub: sub}, nil
}

// WatchRemovalRequested is a free log subscription operation binding the contract event 0xb32a150b9737190a456d8b2b81dd7d592a799ab2933ea494e44351acd41f835d.
//
// Solidity: event RemovalRequested(uint256 indexed period, address indexed bridgeOperator)
func (_BridgeSlash *BridgeSlashFilterer) WatchRemovalRequested(opts *bind.WatchOpts, sink chan<- *BridgeSlashRemovalRequested, period []*big.Int, bridgeOperator []common.Address) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}

	logs, sub, err := _BridgeSlash.contract.WatchLogs(opts, "RemovalRequested", periodRule, bridgeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSlashRemovalRequested)
				if err := _BridgeSlash.contract.UnpackLog(event, "RemovalRequested", log); err != nil {
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

// ParseRemovalRequested is a log parse operation binding the contract event 0xb32a150b9737190a456d8b2b81dd7d592a799ab2933ea494e44351acd41f835d.
//
// Solidity: event RemovalRequested(uint256 indexed period, address indexed bridgeOperator)
func (_BridgeSlash *BridgeSlashFilterer) ParseRemovalRequested(log types.Log) (*BridgeSlashRemovalRequested, error) {
	event := new(BridgeSlashRemovalRequested)
	if err := _BridgeSlash.contract.UnpackLog(event, "RemovalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSlashSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the BridgeSlash contract.
type BridgeSlashSlashedIterator struct {
	Event *BridgeSlashSlashed // Event containing the contract specifics and raw log

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
func (it *BridgeSlashSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSlashSlashed)
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
		it.Event = new(BridgeSlashSlashed)
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
func (it *BridgeSlashSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSlashSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSlashSlashed represents a Slashed event raised by the BridgeSlash contract.
type BridgeSlashSlashed struct {
	Tier             uint8
	BridgeOperator   common.Address
	Period           *big.Int
	SlashUntilPeriod *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x14441e950b7f9ed959e16b2405dd1a9d163efd5d85027b222dcf78b902a00d75.
//
// Solidity: event Slashed(uint8 indexed tier, address indexed bridgeOperator, uint256 indexed period, uint256 slashUntilPeriod)
func (_BridgeSlash *BridgeSlashFilterer) FilterSlashed(opts *bind.FilterOpts, tier []uint8, bridgeOperator []common.Address, period []*big.Int) (*BridgeSlashSlashedIterator, error) {

	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeSlash.contract.FilterLogs(opts, "Slashed", tierRule, bridgeOperatorRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &BridgeSlashSlashedIterator{contract: _BridgeSlash.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x14441e950b7f9ed959e16b2405dd1a9d163efd5d85027b222dcf78b902a00d75.
//
// Solidity: event Slashed(uint8 indexed tier, address indexed bridgeOperator, uint256 indexed period, uint256 slashUntilPeriod)
func (_BridgeSlash *BridgeSlashFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *BridgeSlashSlashed, tier []uint8, bridgeOperator []common.Address, period []*big.Int) (event.Subscription, error) {

	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeSlash.contract.WatchLogs(opts, "Slashed", tierRule, bridgeOperatorRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSlashSlashed)
				if err := _BridgeSlash.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x14441e950b7f9ed959e16b2405dd1a9d163efd5d85027b222dcf78b902a00d75.
//
// Solidity: event Slashed(uint8 indexed tier, address indexed bridgeOperator, uint256 indexed period, uint256 slashUntilPeriod)
func (_BridgeSlash *BridgeSlashFilterer) ParseSlashed(log types.Log) (*BridgeSlashSlashed, error) {
	event := new(BridgeSlashSlashed)
	if err := _BridgeSlash.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
