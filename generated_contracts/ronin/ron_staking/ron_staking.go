// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronStaking

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

// RonStakingMetaData contains all meta data concerning the RonStaking contract.
var RonStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minSecs\",\"type\":\"uint256\"}],\"name\":\"CooldownSecsToUndelegateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensuAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MinValidatorStakingAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"PoolApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"PoolSharesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validator\",\"type\":\"address[]\"}],\"name\":\"PoolsDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"poolAddrs\",\"type\":\"address[]\"}],\"name\":\"PoolsUpdateConflicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"poolAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"name\":\"PoolsUpdateFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"poolAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"aRps\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"name\":\"PoolsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensuAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"StakingAmountTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensuAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensuAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debited\",\"type\":\"uint256\"}],\"name\":\"UserRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"WaitingSecsToRevokeUpdated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidateAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"applyValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_pools\",\"type\":\"address[]\"}],\"name\":\"bulkSelfStaking\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_selfStakings\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_poolAddrs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_userList\",\"type\":\"address[]\"}],\"name\":\"bulkStakingAmountOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_poolList\",\"type\":\"address[]\"}],\"name\":\"bulkStakingTotal\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrList\",\"type\":\"address[]\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownSecsToUndelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deductStakingAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrList\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_consensusAddrDst\",\"type\":\"address\"}],\"name\":\"delegateRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_pools\",\"type\":\"address[]\"}],\"name\":\"deprecatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_poolAddrList\",\"type\":\"address[]\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"}],\"name\":\"getStakingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__minValidatorStakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__cooldownSecsToUndelegate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__waitingSecsToRevoke\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minValidatorStakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"recordRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddrSrc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddrDst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"requestRenounce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cooldownSecs\",\"type\":\"uint256\"}],\"name\":\"setCooldownSecsToUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setMinValidatorStakingAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_secs\",\"type\":\"uint256\"}],\"name\":\"setWaitingSecsToRevoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"stakingAmountOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"}],\"name\":\"stakingTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"waitingSecsToRevoke\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RonStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use RonStakingMetaData.ABI instead.
var RonStakingABI = RonStakingMetaData.ABI

// RonStaking is an auto generated Go binding around an Ethereum contract.
type RonStaking struct {
	RonStakingCaller     // Read-only binding to the contract
	RonStakingTransactor // Write-only binding to the contract
	RonStakingFilterer   // Log filterer for contract events
}

// RonStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RonStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RonStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RonStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RonStakingSession struct {
	Contract     *RonStaking       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RonStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RonStakingCallerSession struct {
	Contract *RonStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RonStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RonStakingTransactorSession struct {
	Contract     *RonStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RonStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RonStakingRaw struct {
	Contract *RonStaking // Generic contract binding to access the raw methods on
}

// RonStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RonStakingCallerRaw struct {
	Contract *RonStakingCaller // Generic read-only contract binding to access the raw methods on
}

// RonStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RonStakingTransactorRaw struct {
	Contract *RonStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRonStaking creates a new instance of RonStaking, bound to a specific deployed contract.
func NewRonStaking(address common.Address, backend bind.ContractBackend) (*RonStaking, error) {
	contract, err := bindRonStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RonStaking{RonStakingCaller: RonStakingCaller{contract: contract}, RonStakingTransactor: RonStakingTransactor{contract: contract}, RonStakingFilterer: RonStakingFilterer{contract: contract}}, nil
}

// NewRonStakingCaller creates a new read-only instance of RonStaking, bound to a specific deployed contract.
func NewRonStakingCaller(address common.Address, caller bind.ContractCaller) (*RonStakingCaller, error) {
	contract, err := bindRonStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RonStakingCaller{contract: contract}, nil
}

// NewRonStakingTransactor creates a new write-only instance of RonStaking, bound to a specific deployed contract.
func NewRonStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*RonStakingTransactor, error) {
	contract, err := bindRonStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RonStakingTransactor{contract: contract}, nil
}

// NewRonStakingFilterer creates a new log filterer instance of RonStaking, bound to a specific deployed contract.
func NewRonStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*RonStakingFilterer, error) {
	contract, err := bindRonStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RonStakingFilterer{contract: contract}, nil
}

// bindRonStaking binds a generic wrapper to an already deployed contract.
func bindRonStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RonStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RonStaking *RonStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RonStaking.Contract.RonStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RonStaking *RonStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RonStaking.Contract.RonStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RonStaking *RonStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RonStaking.Contract.RonStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RonStaking *RonStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RonStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RonStaking *RonStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RonStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RonStaking *RonStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RonStaking.Contract.contract.Transact(opts, method, params...)
}

// BulkSelfStaking is a free data retrieval call binding the contract method 0x017dd950.
//
// Solidity: function bulkSelfStaking(address[] _pools) view returns(uint256[] _selfStakings)
func (_RonStaking *RonStakingCaller) BulkSelfStaking(opts *bind.CallOpts, _pools []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "bulkSelfStaking", _pools)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BulkSelfStaking is a free data retrieval call binding the contract method 0x017dd950.
//
// Solidity: function bulkSelfStaking(address[] _pools) view returns(uint256[] _selfStakings)
func (_RonStaking *RonStakingSession) BulkSelfStaking(_pools []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.BulkSelfStaking(&_RonStaking.CallOpts, _pools)
}

// BulkSelfStaking is a free data retrieval call binding the contract method 0x017dd950.
//
// Solidity: function bulkSelfStaking(address[] _pools) view returns(uint256[] _selfStakings)
func (_RonStaking *RonStakingCallerSession) BulkSelfStaking(_pools []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.BulkSelfStaking(&_RonStaking.CallOpts, _pools)
}

// BulkStakingAmountOf is a free data retrieval call binding the contract method 0xe793d5bc.
//
// Solidity: function bulkStakingAmountOf(address[] _poolAddrs, address[] _userList) view returns(uint256[] _stakingAmounts)
func (_RonStaking *RonStakingCaller) BulkStakingAmountOf(opts *bind.CallOpts, _poolAddrs []common.Address, _userList []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "bulkStakingAmountOf", _poolAddrs, _userList)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BulkStakingAmountOf is a free data retrieval call binding the contract method 0xe793d5bc.
//
// Solidity: function bulkStakingAmountOf(address[] _poolAddrs, address[] _userList) view returns(uint256[] _stakingAmounts)
func (_RonStaking *RonStakingSession) BulkStakingAmountOf(_poolAddrs []common.Address, _userList []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.BulkStakingAmountOf(&_RonStaking.CallOpts, _poolAddrs, _userList)
}

// BulkStakingAmountOf is a free data retrieval call binding the contract method 0xe793d5bc.
//
// Solidity: function bulkStakingAmountOf(address[] _poolAddrs, address[] _userList) view returns(uint256[] _stakingAmounts)
func (_RonStaking *RonStakingCallerSession) BulkStakingAmountOf(_poolAddrs []common.Address, _userList []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.BulkStakingAmountOf(&_RonStaking.CallOpts, _poolAddrs, _userList)
}

// BulkStakingTotal is a free data retrieval call binding the contract method 0xafcb2e02.
//
// Solidity: function bulkStakingTotal(address[] _poolList) view returns(uint256[] _stakingAmounts)
func (_RonStaking *RonStakingCaller) BulkStakingTotal(opts *bind.CallOpts, _poolList []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "bulkStakingTotal", _poolList)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BulkStakingTotal is a free data retrieval call binding the contract method 0xafcb2e02.
//
// Solidity: function bulkStakingTotal(address[] _poolList) view returns(uint256[] _stakingAmounts)
func (_RonStaking *RonStakingSession) BulkStakingTotal(_poolList []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.BulkStakingTotal(&_RonStaking.CallOpts, _poolList)
}

// BulkStakingTotal is a free data retrieval call binding the contract method 0xafcb2e02.
//
// Solidity: function bulkStakingTotal(address[] _poolList) view returns(uint256[] _stakingAmounts)
func (_RonStaking *RonStakingCallerSession) BulkStakingTotal(_poolList []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.BulkStakingTotal(&_RonStaking.CallOpts, _poolList)
}

// CooldownSecsToUndelegate is a free data retrieval call binding the contract method 0x0682e8fa.
//
// Solidity: function cooldownSecsToUndelegate() view returns(uint256)
func (_RonStaking *RonStakingCaller) CooldownSecsToUndelegate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "cooldownSecsToUndelegate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownSecsToUndelegate is a free data retrieval call binding the contract method 0x0682e8fa.
//
// Solidity: function cooldownSecsToUndelegate() view returns(uint256)
func (_RonStaking *RonStakingSession) CooldownSecsToUndelegate() (*big.Int, error) {
	return _RonStaking.Contract.CooldownSecsToUndelegate(&_RonStaking.CallOpts)
}

// CooldownSecsToUndelegate is a free data retrieval call binding the contract method 0x0682e8fa.
//
// Solidity: function cooldownSecsToUndelegate() view returns(uint256)
func (_RonStaking *RonStakingCallerSession) CooldownSecsToUndelegate() (*big.Int, error) {
	return _RonStaking.Contract.CooldownSecsToUndelegate(&_RonStaking.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _poolAddr, address _user) view returns(uint256)
func (_RonStaking *RonStakingCaller) GetReward(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "getReward", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _poolAddr, address _user) view returns(uint256)
func (_RonStaking *RonStakingSession) GetReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _RonStaking.Contract.GetReward(&_RonStaking.CallOpts, _poolAddr, _user)
}

// GetReward is a free data retrieval call binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _poolAddr, address _user) view returns(uint256)
func (_RonStaking *RonStakingCallerSession) GetReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _RonStaking.Contract.GetReward(&_RonStaking.CallOpts, _poolAddr, _user)
}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _rewards)
func (_RonStaking *RonStakingCaller) GetRewards(opts *bind.CallOpts, _user common.Address, _poolAddrList []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "getRewards", _user, _poolAddrList)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _rewards)
func (_RonStaking *RonStakingSession) GetRewards(_user common.Address, _poolAddrList []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.GetRewards(&_RonStaking.CallOpts, _user, _poolAddrList)
}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _rewards)
func (_RonStaking *RonStakingCallerSession) GetRewards(_user common.Address, _poolAddrList []common.Address) ([]*big.Int, error) {
	return _RonStaking.Contract.GetRewards(&_RonStaking.CallOpts, _user, _poolAddrList)
}

// GetStakingPool is a free data retrieval call binding the contract method 0x9e614e0e.
//
// Solidity: function getStakingPool(address _poolAddr) view returns(address _admin, uint256 _stakingAmount, uint256 _stakingTotal)
func (_RonStaking *RonStakingCaller) GetStakingPool(opts *bind.CallOpts, _poolAddr common.Address) (struct {
	Admin         common.Address
	StakingAmount *big.Int
	StakingTotal  *big.Int
}, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "getStakingPool", _poolAddr)

	outstruct := new(struct {
		Admin         common.Address
		StakingAmount *big.Int
		StakingTotal  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Admin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakingAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakingTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakingPool is a free data retrieval call binding the contract method 0x9e614e0e.
//
// Solidity: function getStakingPool(address _poolAddr) view returns(address _admin, uint256 _stakingAmount, uint256 _stakingTotal)
func (_RonStaking *RonStakingSession) GetStakingPool(_poolAddr common.Address) (struct {
	Admin         common.Address
	StakingAmount *big.Int
	StakingTotal  *big.Int
}, error) {
	return _RonStaking.Contract.GetStakingPool(&_RonStaking.CallOpts, _poolAddr)
}

// GetStakingPool is a free data retrieval call binding the contract method 0x9e614e0e.
//
// Solidity: function getStakingPool(address _poolAddr) view returns(address _admin, uint256 _stakingAmount, uint256 _stakingTotal)
func (_RonStaking *RonStakingCallerSession) GetStakingPool(_poolAddr common.Address) (struct {
	Admin         common.Address
	StakingAmount *big.Int
	StakingTotal  *big.Int
}, error) {
	return _RonStaking.Contract.GetStakingPool(&_RonStaking.CallOpts, _poolAddr)
}

// MinValidatorStakingAmount is a free data retrieval call binding the contract method 0x909791dd.
//
// Solidity: function minValidatorStakingAmount() view returns(uint256)
func (_RonStaking *RonStakingCaller) MinValidatorStakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "minValidatorStakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorStakingAmount is a free data retrieval call binding the contract method 0x909791dd.
//
// Solidity: function minValidatorStakingAmount() view returns(uint256)
func (_RonStaking *RonStakingSession) MinValidatorStakingAmount() (*big.Int, error) {
	return _RonStaking.Contract.MinValidatorStakingAmount(&_RonStaking.CallOpts)
}

// MinValidatorStakingAmount is a free data retrieval call binding the contract method 0x909791dd.
//
// Solidity: function minValidatorStakingAmount() view returns(uint256)
func (_RonStaking *RonStakingCallerSession) MinValidatorStakingAmount() (*big.Int, error) {
	return _RonStaking.Contract.MinValidatorStakingAmount(&_RonStaking.CallOpts)
}

// StakingAmountOf is a free data retrieval call binding the contract method 0x161a1daf.
//
// Solidity: function stakingAmountOf(address _poolAddr, address _user) view returns(uint256)
func (_RonStaking *RonStakingCaller) StakingAmountOf(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "stakingAmountOf", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingAmountOf is a free data retrieval call binding the contract method 0x161a1daf.
//
// Solidity: function stakingAmountOf(address _poolAddr, address _user) view returns(uint256)
func (_RonStaking *RonStakingSession) StakingAmountOf(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _RonStaking.Contract.StakingAmountOf(&_RonStaking.CallOpts, _poolAddr, _user)
}

// StakingAmountOf is a free data retrieval call binding the contract method 0x161a1daf.
//
// Solidity: function stakingAmountOf(address _poolAddr, address _user) view returns(uint256)
func (_RonStaking *RonStakingCallerSession) StakingAmountOf(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _RonStaking.Contract.StakingAmountOf(&_RonStaking.CallOpts, _poolAddr, _user)
}

// StakingTotal is a free data retrieval call binding the contract method 0x82c47bce.
//
// Solidity: function stakingTotal(address _poolAddr) view returns(uint256)
func (_RonStaking *RonStakingCaller) StakingTotal(opts *bind.CallOpts, _poolAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "stakingTotal", _poolAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingTotal is a free data retrieval call binding the contract method 0x82c47bce.
//
// Solidity: function stakingTotal(address _poolAddr) view returns(uint256)
func (_RonStaking *RonStakingSession) StakingTotal(_poolAddr common.Address) (*big.Int, error) {
	return _RonStaking.Contract.StakingTotal(&_RonStaking.CallOpts, _poolAddr)
}

// StakingTotal is a free data retrieval call binding the contract method 0x82c47bce.
//
// Solidity: function stakingTotal(address _poolAddr) view returns(uint256)
func (_RonStaking *RonStakingCallerSession) StakingTotal(_poolAddr common.Address) (*big.Int, error) {
	return _RonStaking.Contract.StakingTotal(&_RonStaking.CallOpts, _poolAddr)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_RonStaking *RonStakingCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_RonStaking *RonStakingSession) ValidatorContract() (common.Address, error) {
	return _RonStaking.Contract.ValidatorContract(&_RonStaking.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_RonStaking *RonStakingCallerSession) ValidatorContract() (common.Address, error) {
	return _RonStaking.Contract.ValidatorContract(&_RonStaking.CallOpts)
}

// WaitingSecsToRevoke is a free data retrieval call binding the contract method 0xaf245429.
//
// Solidity: function waitingSecsToRevoke() view returns(uint256)
func (_RonStaking *RonStakingCaller) WaitingSecsToRevoke(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RonStaking.contract.Call(opts, &out, "waitingSecsToRevoke")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WaitingSecsToRevoke is a free data retrieval call binding the contract method 0xaf245429.
//
// Solidity: function waitingSecsToRevoke() view returns(uint256)
func (_RonStaking *RonStakingSession) WaitingSecsToRevoke() (*big.Int, error) {
	return _RonStaking.Contract.WaitingSecsToRevoke(&_RonStaking.CallOpts)
}

// WaitingSecsToRevoke is a free data retrieval call binding the contract method 0xaf245429.
//
// Solidity: function waitingSecsToRevoke() view returns(uint256)
func (_RonStaking *RonStakingCallerSession) WaitingSecsToRevoke() (*big.Int, error) {
	return _RonStaking.Contract.WaitingSecsToRevoke(&_RonStaking.CallOpts)
}

// ApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xe5376f54.
//
// Solidity: function applyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) payable returns()
func (_RonStaking *RonStakingTransactor) ApplyValidatorCandidate(opts *bind.TransactOpts, _candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "applyValidatorCandidate", _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xe5376f54.
//
// Solidity: function applyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) payable returns()
func (_RonStaking *RonStakingSession) ApplyValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.ApplyValidatorCandidate(&_RonStaking.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xe5376f54.
//
// Solidity: function applyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) payable returns()
func (_RonStaking *RonStakingTransactorSession) ApplyValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.ApplyValidatorCandidate(&_RonStaking.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// BulkUndelegate is a paid mutator transaction binding the contract method 0x9488e4e9.
//
// Solidity: function bulkUndelegate(address[] _consensusAddrs, uint256[] _amounts) returns()
func (_RonStaking *RonStakingTransactor) BulkUndelegate(opts *bind.TransactOpts, _consensusAddrs []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "bulkUndelegate", _consensusAddrs, _amounts)
}

// BulkUndelegate is a paid mutator transaction binding the contract method 0x9488e4e9.
//
// Solidity: function bulkUndelegate(address[] _consensusAddrs, uint256[] _amounts) returns()
func (_RonStaking *RonStakingSession) BulkUndelegate(_consensusAddrs []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.BulkUndelegate(&_RonStaking.TransactOpts, _consensusAddrs, _amounts)
}

// BulkUndelegate is a paid mutator transaction binding the contract method 0x9488e4e9.
//
// Solidity: function bulkUndelegate(address[] _consensusAddrs, uint256[] _amounts) returns()
func (_RonStaking *RonStakingTransactorSession) BulkUndelegate(_consensusAddrs []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.BulkUndelegate(&_RonStaking.TransactOpts, _consensusAddrs, _amounts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_RonStaking *RonStakingTransactor) ClaimRewards(opts *bind.TransactOpts, _consensusAddrList []common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "claimRewards", _consensusAddrList)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_RonStaking *RonStakingSession) ClaimRewards(_consensusAddrList []common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.ClaimRewards(&_RonStaking.TransactOpts, _consensusAddrList)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_RonStaking *RonStakingTransactorSession) ClaimRewards(_consensusAddrList []common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.ClaimRewards(&_RonStaking.TransactOpts, _consensusAddrList)
}

// DeductStakingAmount is a paid mutator transaction binding the contract method 0xc905bb35.
//
// Solidity: function deductStakingAmount(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactor) DeductStakingAmount(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "deductStakingAmount", _consensusAddr, _amount)
}

// DeductStakingAmount is a paid mutator transaction binding the contract method 0xc905bb35.
//
// Solidity: function deductStakingAmount(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingSession) DeductStakingAmount(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.DeductStakingAmount(&_RonStaking.TransactOpts, _consensusAddr, _amount)
}

// DeductStakingAmount is a paid mutator transaction binding the contract method 0xc905bb35.
//
// Solidity: function deductStakingAmount(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactorSession) DeductStakingAmount(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.DeductStakingAmount(&_RonStaking.TransactOpts, _consensusAddr, _amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_RonStaking *RonStakingTransactor) Delegate(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "delegate", _consensusAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_RonStaking *RonStakingSession) Delegate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.Delegate(&_RonStaking.TransactOpts, _consensusAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_RonStaking *RonStakingTransactorSession) Delegate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.Delegate(&_RonStaking.TransactOpts, _consensusAddr)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_RonStaking *RonStakingTransactor) DelegateRewards(opts *bind.TransactOpts, _consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "delegateRewards", _consensusAddrList, _consensusAddrDst)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_RonStaking *RonStakingSession) DelegateRewards(_consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.DelegateRewards(&_RonStaking.TransactOpts, _consensusAddrList, _consensusAddrDst)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_RonStaking *RonStakingTransactorSession) DelegateRewards(_consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.DelegateRewards(&_RonStaking.TransactOpts, _consensusAddrList, _consensusAddrDst)
}

// DeprecatePools is a paid mutator transaction binding the contract method 0xe9ac5e06.
//
// Solidity: function deprecatePools(address[] _pools) returns()
func (_RonStaking *RonStakingTransactor) DeprecatePools(opts *bind.TransactOpts, _pools []common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "deprecatePools", _pools)
}

// DeprecatePools is a paid mutator transaction binding the contract method 0xe9ac5e06.
//
// Solidity: function deprecatePools(address[] _pools) returns()
func (_RonStaking *RonStakingSession) DeprecatePools(_pools []common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.DeprecatePools(&_RonStaking.TransactOpts, _pools)
}

// DeprecatePools is a paid mutator transaction binding the contract method 0xe9ac5e06.
//
// Solidity: function deprecatePools(address[] _pools) returns()
func (_RonStaking *RonStakingTransactorSession) DeprecatePools(_pools []common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.DeprecatePools(&_RonStaking.TransactOpts, _pools)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address __validatorContract, uint256 __minValidatorStakingAmount, uint256 __cooldownSecsToUndelegate, uint256 __waitingSecsToRevoke) returns()
func (_RonStaking *RonStakingTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __minValidatorStakingAmount *big.Int, __cooldownSecsToUndelegate *big.Int, __waitingSecsToRevoke *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "initialize", __validatorContract, __minValidatorStakingAmount, __cooldownSecsToUndelegate, __waitingSecsToRevoke)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address __validatorContract, uint256 __minValidatorStakingAmount, uint256 __cooldownSecsToUndelegate, uint256 __waitingSecsToRevoke) returns()
func (_RonStaking *RonStakingSession) Initialize(__validatorContract common.Address, __minValidatorStakingAmount *big.Int, __cooldownSecsToUndelegate *big.Int, __waitingSecsToRevoke *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Initialize(&_RonStaking.TransactOpts, __validatorContract, __minValidatorStakingAmount, __cooldownSecsToUndelegate, __waitingSecsToRevoke)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address __validatorContract, uint256 __minValidatorStakingAmount, uint256 __cooldownSecsToUndelegate, uint256 __waitingSecsToRevoke) returns()
func (_RonStaking *RonStakingTransactorSession) Initialize(__validatorContract common.Address, __minValidatorStakingAmount *big.Int, __cooldownSecsToUndelegate *big.Int, __waitingSecsToRevoke *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Initialize(&_RonStaking.TransactOpts, __validatorContract, __minValidatorStakingAmount, __cooldownSecsToUndelegate, __waitingSecsToRevoke)
}

// RecordRewards is a paid mutator transaction binding the contract method 0x3b8cb16b.
//
// Solidity: function recordRewards(address[] _consensusAddrs, uint256[] _rewards, uint256 _period) payable returns()
func (_RonStaking *RonStakingTransactor) RecordRewards(opts *bind.TransactOpts, _consensusAddrs []common.Address, _rewards []*big.Int, _period *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "recordRewards", _consensusAddrs, _rewards, _period)
}

// RecordRewards is a paid mutator transaction binding the contract method 0x3b8cb16b.
//
// Solidity: function recordRewards(address[] _consensusAddrs, uint256[] _rewards, uint256 _period) payable returns()
func (_RonStaking *RonStakingSession) RecordRewards(_consensusAddrs []common.Address, _rewards []*big.Int, _period *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.RecordRewards(&_RonStaking.TransactOpts, _consensusAddrs, _rewards, _period)
}

// RecordRewards is a paid mutator transaction binding the contract method 0x3b8cb16b.
//
// Solidity: function recordRewards(address[] _consensusAddrs, uint256[] _rewards, uint256 _period) payable returns()
func (_RonStaking *RonStakingTransactorSession) RecordRewards(_consensusAddrs []common.Address, _rewards []*big.Int, _period *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.RecordRewards(&_RonStaking.TransactOpts, _consensusAddrs, _rewards, _period)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactor) Redelegate(opts *bind.TransactOpts, _consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "redelegate", _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_RonStaking *RonStakingSession) Redelegate(_consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Redelegate(&_RonStaking.TransactOpts, _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactorSession) Redelegate(_consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Redelegate(&_RonStaking.TransactOpts, _consensusAddrSrc, _consensusAddrDst, _amount)
}

// RequestRenounce is a paid mutator transaction binding the contract method 0x1658c86e.
//
// Solidity: function requestRenounce(address _consensusAddr) returns()
func (_RonStaking *RonStakingTransactor) RequestRenounce(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "requestRenounce", _consensusAddr)
}

// RequestRenounce is a paid mutator transaction binding the contract method 0x1658c86e.
//
// Solidity: function requestRenounce(address _consensusAddr) returns()
func (_RonStaking *RonStakingSession) RequestRenounce(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.RequestRenounce(&_RonStaking.TransactOpts, _consensusAddr)
}

// RequestRenounce is a paid mutator transaction binding the contract method 0x1658c86e.
//
// Solidity: function requestRenounce(address _consensusAddr) returns()
func (_RonStaking *RonStakingTransactorSession) RequestRenounce(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.RequestRenounce(&_RonStaking.TransactOpts, _consensusAddr)
}

// SetCooldownSecsToUndelegate is a paid mutator transaction binding the contract method 0x888b9ae9.
//
// Solidity: function setCooldownSecsToUndelegate(uint256 _cooldownSecs) returns()
func (_RonStaking *RonStakingTransactor) SetCooldownSecsToUndelegate(opts *bind.TransactOpts, _cooldownSecs *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "setCooldownSecsToUndelegate", _cooldownSecs)
}

// SetCooldownSecsToUndelegate is a paid mutator transaction binding the contract method 0x888b9ae9.
//
// Solidity: function setCooldownSecsToUndelegate(uint256 _cooldownSecs) returns()
func (_RonStaking *RonStakingSession) SetCooldownSecsToUndelegate(_cooldownSecs *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.SetCooldownSecsToUndelegate(&_RonStaking.TransactOpts, _cooldownSecs)
}

// SetCooldownSecsToUndelegate is a paid mutator transaction binding the contract method 0x888b9ae9.
//
// Solidity: function setCooldownSecsToUndelegate(uint256 _cooldownSecs) returns()
func (_RonStaking *RonStakingTransactorSession) SetCooldownSecsToUndelegate(_cooldownSecs *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.SetCooldownSecsToUndelegate(&_RonStaking.TransactOpts, _cooldownSecs)
}

// SetMinValidatorStakingAmount is a paid mutator transaction binding the contract method 0x679a6e43.
//
// Solidity: function setMinValidatorStakingAmount(uint256 _threshold) returns()
func (_RonStaking *RonStakingTransactor) SetMinValidatorStakingAmount(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "setMinValidatorStakingAmount", _threshold)
}

// SetMinValidatorStakingAmount is a paid mutator transaction binding the contract method 0x679a6e43.
//
// Solidity: function setMinValidatorStakingAmount(uint256 _threshold) returns()
func (_RonStaking *RonStakingSession) SetMinValidatorStakingAmount(_threshold *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.SetMinValidatorStakingAmount(&_RonStaking.TransactOpts, _threshold)
}

// SetMinValidatorStakingAmount is a paid mutator transaction binding the contract method 0x679a6e43.
//
// Solidity: function setMinValidatorStakingAmount(uint256 _threshold) returns()
func (_RonStaking *RonStakingTransactorSession) SetMinValidatorStakingAmount(_threshold *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.SetMinValidatorStakingAmount(&_RonStaking.TransactOpts, _threshold)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_RonStaking *RonStakingTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_RonStaking *RonStakingSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.SetValidatorContract(&_RonStaking.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_RonStaking *RonStakingTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.SetValidatorContract(&_RonStaking.TransactOpts, _addr)
}

// SetWaitingSecsToRevoke is a paid mutator transaction binding the contract method 0x969ffc14.
//
// Solidity: function setWaitingSecsToRevoke(uint256 _secs) returns()
func (_RonStaking *RonStakingTransactor) SetWaitingSecsToRevoke(opts *bind.TransactOpts, _secs *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "setWaitingSecsToRevoke", _secs)
}

// SetWaitingSecsToRevoke is a paid mutator transaction binding the contract method 0x969ffc14.
//
// Solidity: function setWaitingSecsToRevoke(uint256 _secs) returns()
func (_RonStaking *RonStakingSession) SetWaitingSecsToRevoke(_secs *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.SetWaitingSecsToRevoke(&_RonStaking.TransactOpts, _secs)
}

// SetWaitingSecsToRevoke is a paid mutator transaction binding the contract method 0x969ffc14.
//
// Solidity: function setWaitingSecsToRevoke(uint256 _secs) returns()
func (_RonStaking *RonStakingTransactorSession) SetWaitingSecsToRevoke(_secs *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.SetWaitingSecsToRevoke(&_RonStaking.TransactOpts, _secs)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_RonStaking *RonStakingTransactor) Stake(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "stake", _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_RonStaking *RonStakingSession) Stake(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.Stake(&_RonStaking.TransactOpts, _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_RonStaking *RonStakingTransactorSession) Stake(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RonStaking.Contract.Stake(&_RonStaking.TransactOpts, _consensusAddr)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactor) Undelegate(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "undelegate", _consensusAddr, _amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingSession) Undelegate(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Undelegate(&_RonStaking.TransactOpts, _consensusAddr, _amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactorSession) Undelegate(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Undelegate(&_RonStaking.TransactOpts, _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactor) Unstake(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.contract.Transact(opts, "unstake", _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingSession) Unstake(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Unstake(&_RonStaking.TransactOpts, _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_RonStaking *RonStakingTransactorSession) Unstake(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RonStaking.Contract.Unstake(&_RonStaking.TransactOpts, _consensusAddr, _amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RonStaking *RonStakingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RonStaking.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RonStaking *RonStakingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RonStaking.Contract.Fallback(&_RonStaking.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RonStaking *RonStakingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RonStaking.Contract.Fallback(&_RonStaking.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RonStaking *RonStakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RonStaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RonStaking *RonStakingSession) Receive() (*types.Transaction, error) {
	return _RonStaking.Contract.Receive(&_RonStaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RonStaking *RonStakingTransactorSession) Receive() (*types.Transaction, error) {
	return _RonStaking.Contract.Receive(&_RonStaking.TransactOpts)
}

// RonStakingCooldownSecsToUndelegateUpdatedIterator is returned from FilterCooldownSecsToUndelegateUpdated and is used to iterate over the raw logs and unpacked data for CooldownSecsToUndelegateUpdated events raised by the RonStaking contract.
type RonStakingCooldownSecsToUndelegateUpdatedIterator struct {
	Event *RonStakingCooldownSecsToUndelegateUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingCooldownSecsToUndelegateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingCooldownSecsToUndelegateUpdated)
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
		it.Event = new(RonStakingCooldownSecsToUndelegateUpdated)
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
func (it *RonStakingCooldownSecsToUndelegateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingCooldownSecsToUndelegateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingCooldownSecsToUndelegateUpdated represents a CooldownSecsToUndelegateUpdated event raised by the RonStaking contract.
type RonStakingCooldownSecsToUndelegateUpdated struct {
	MinSecs *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCooldownSecsToUndelegateUpdated is a free log retrieval operation binding the contract event 0x4956b65267b8f1e642284bcb5037116c69a9c78d9ca576beeae0974737a4872a.
//
// Solidity: event CooldownSecsToUndelegateUpdated(uint256 minSecs)
func (_RonStaking *RonStakingFilterer) FilterCooldownSecsToUndelegateUpdated(opts *bind.FilterOpts) (*RonStakingCooldownSecsToUndelegateUpdatedIterator, error) {

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "CooldownSecsToUndelegateUpdated")
	if err != nil {
		return nil, err
	}
	return &RonStakingCooldownSecsToUndelegateUpdatedIterator{contract: _RonStaking.contract, event: "CooldownSecsToUndelegateUpdated", logs: logs, sub: sub}, nil
}

// WatchCooldownSecsToUndelegateUpdated is a free log subscription operation binding the contract event 0x4956b65267b8f1e642284bcb5037116c69a9c78d9ca576beeae0974737a4872a.
//
// Solidity: event CooldownSecsToUndelegateUpdated(uint256 minSecs)
func (_RonStaking *RonStakingFilterer) WatchCooldownSecsToUndelegateUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingCooldownSecsToUndelegateUpdated) (event.Subscription, error) {

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "CooldownSecsToUndelegateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingCooldownSecsToUndelegateUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "CooldownSecsToUndelegateUpdated", log); err != nil {
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

// ParseCooldownSecsToUndelegateUpdated is a log parse operation binding the contract event 0x4956b65267b8f1e642284bcb5037116c69a9c78d9ca576beeae0974737a4872a.
//
// Solidity: event CooldownSecsToUndelegateUpdated(uint256 minSecs)
func (_RonStaking *RonStakingFilterer) ParseCooldownSecsToUndelegateUpdated(log types.Log) (*RonStakingCooldownSecsToUndelegateUpdated, error) {
	event := new(RonStakingCooldownSecsToUndelegateUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "CooldownSecsToUndelegateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the RonStaking contract.
type RonStakingDelegatedIterator struct {
	Event *RonStakingDelegated // Event containing the contract specifics and raw log

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
func (it *RonStakingDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingDelegated)
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
		it.Event = new(RonStakingDelegated)
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
func (it *RonStakingDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingDelegated represents a Delegated event raised by the RonStaking contract.
type RonStakingDelegated struct {
	Delegator    common.Address
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, consensuAddr []common.Address) (*RonStakingDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "Delegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingDelegatedIterator{contract: _RonStaking.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *RonStakingDelegated, delegator []common.Address, consensuAddr []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "Delegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingDelegated)
				if err := _RonStaking.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) ParseDelegated(log types.Log) (*RonStakingDelegated, error) {
	event := new(RonStakingDelegated)
	if err := _RonStaking.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RonStaking contract.
type RonStakingInitializedIterator struct {
	Event *RonStakingInitialized // Event containing the contract specifics and raw log

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
func (it *RonStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingInitialized)
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
		it.Event = new(RonStakingInitialized)
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
func (it *RonStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingInitialized represents a Initialized event raised by the RonStaking contract.
type RonStakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RonStaking *RonStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*RonStakingInitializedIterator, error) {

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RonStakingInitializedIterator{contract: _RonStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RonStaking *RonStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RonStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingInitialized)
				if err := _RonStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RonStaking *RonStakingFilterer) ParseInitialized(log types.Log) (*RonStakingInitialized, error) {
	event := new(RonStakingInitialized)
	if err := _RonStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingMinValidatorStakingAmountUpdatedIterator is returned from FilterMinValidatorStakingAmountUpdated and is used to iterate over the raw logs and unpacked data for MinValidatorStakingAmountUpdated events raised by the RonStaking contract.
type RonStakingMinValidatorStakingAmountUpdatedIterator struct {
	Event *RonStakingMinValidatorStakingAmountUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingMinValidatorStakingAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingMinValidatorStakingAmountUpdated)
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
		it.Event = new(RonStakingMinValidatorStakingAmountUpdated)
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
func (it *RonStakingMinValidatorStakingAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingMinValidatorStakingAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingMinValidatorStakingAmountUpdated represents a MinValidatorStakingAmountUpdated event raised by the RonStaking contract.
type RonStakingMinValidatorStakingAmountUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinValidatorStakingAmountUpdated is a free log retrieval operation binding the contract event 0x372bbdb8d72373b0012f84ee5a11671e5fb72b8bea902ebca93a19cb45d32be2.
//
// Solidity: event MinValidatorStakingAmountUpdated(uint256 threshold)
func (_RonStaking *RonStakingFilterer) FilterMinValidatorStakingAmountUpdated(opts *bind.FilterOpts) (*RonStakingMinValidatorStakingAmountUpdatedIterator, error) {

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "MinValidatorStakingAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &RonStakingMinValidatorStakingAmountUpdatedIterator{contract: _RonStaking.contract, event: "MinValidatorStakingAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchMinValidatorStakingAmountUpdated is a free log subscription operation binding the contract event 0x372bbdb8d72373b0012f84ee5a11671e5fb72b8bea902ebca93a19cb45d32be2.
//
// Solidity: event MinValidatorStakingAmountUpdated(uint256 threshold)
func (_RonStaking *RonStakingFilterer) WatchMinValidatorStakingAmountUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingMinValidatorStakingAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "MinValidatorStakingAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingMinValidatorStakingAmountUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "MinValidatorStakingAmountUpdated", log); err != nil {
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

// ParseMinValidatorStakingAmountUpdated is a log parse operation binding the contract event 0x372bbdb8d72373b0012f84ee5a11671e5fb72b8bea902ebca93a19cb45d32be2.
//
// Solidity: event MinValidatorStakingAmountUpdated(uint256 threshold)
func (_RonStaking *RonStakingFilterer) ParseMinValidatorStakingAmountUpdated(log types.Log) (*RonStakingMinValidatorStakingAmountUpdated, error) {
	event := new(RonStakingMinValidatorStakingAmountUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "MinValidatorStakingAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingPoolApprovedIterator is returned from FilterPoolApproved and is used to iterate over the raw logs and unpacked data for PoolApproved events raised by the RonStaking contract.
type RonStakingPoolApprovedIterator struct {
	Event *RonStakingPoolApproved // Event containing the contract specifics and raw log

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
func (it *RonStakingPoolApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingPoolApproved)
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
		it.Event = new(RonStakingPoolApproved)
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
func (it *RonStakingPoolApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingPoolApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingPoolApproved represents a PoolApproved event raised by the RonStaking contract.
type RonStakingPoolApproved struct {
	Validator common.Address
	Admin     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolApproved is a free log retrieval operation binding the contract event 0xfc1f1e73948cbc47c5b7f90e5601b7daccd9ad7173218486ccc74bdd051d05e8.
//
// Solidity: event PoolApproved(address indexed validator, address indexed admin)
func (_RonStaking *RonStakingFilterer) FilterPoolApproved(opts *bind.FilterOpts, validator []common.Address, admin []common.Address) (*RonStakingPoolApprovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "PoolApproved", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingPoolApprovedIterator{contract: _RonStaking.contract, event: "PoolApproved", logs: logs, sub: sub}, nil
}

// WatchPoolApproved is a free log subscription operation binding the contract event 0xfc1f1e73948cbc47c5b7f90e5601b7daccd9ad7173218486ccc74bdd051d05e8.
//
// Solidity: event PoolApproved(address indexed validator, address indexed admin)
func (_RonStaking *RonStakingFilterer) WatchPoolApproved(opts *bind.WatchOpts, sink chan<- *RonStakingPoolApproved, validator []common.Address, admin []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "PoolApproved", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingPoolApproved)
				if err := _RonStaking.contract.UnpackLog(event, "PoolApproved", log); err != nil {
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

// ParsePoolApproved is a log parse operation binding the contract event 0xfc1f1e73948cbc47c5b7f90e5601b7daccd9ad7173218486ccc74bdd051d05e8.
//
// Solidity: event PoolApproved(address indexed validator, address indexed admin)
func (_RonStaking *RonStakingFilterer) ParsePoolApproved(log types.Log) (*RonStakingPoolApproved, error) {
	event := new(RonStakingPoolApproved)
	if err := _RonStaking.contract.UnpackLog(event, "PoolApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingPoolSharesUpdatedIterator is returned from FilterPoolSharesUpdated and is used to iterate over the raw logs and unpacked data for PoolSharesUpdated events raised by the RonStaking contract.
type RonStakingPoolSharesUpdatedIterator struct {
	Event *RonStakingPoolSharesUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingPoolSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingPoolSharesUpdated)
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
		it.Event = new(RonStakingPoolSharesUpdated)
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
func (it *RonStakingPoolSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingPoolSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingPoolSharesUpdated represents a PoolSharesUpdated event raised by the RonStaking contract.
type RonStakingPoolSharesUpdated struct {
	Period   *big.Int
	PoolAddr common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolSharesUpdated is a free log retrieval operation binding the contract event 0x81faf50e2aaf52eaba2ab841071efb9f6f0850a3e7d008b1336e6001d3d4963c.
//
// Solidity: event PoolSharesUpdated(uint256 indexed period, address indexed poolAddr, uint256 shares)
func (_RonStaking *RonStakingFilterer) FilterPoolSharesUpdated(opts *bind.FilterOpts, period []*big.Int, poolAddr []common.Address) (*RonStakingPoolSharesUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "PoolSharesUpdated", periodRule, poolAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingPoolSharesUpdatedIterator{contract: _RonStaking.contract, event: "PoolSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolSharesUpdated is a free log subscription operation binding the contract event 0x81faf50e2aaf52eaba2ab841071efb9f6f0850a3e7d008b1336e6001d3d4963c.
//
// Solidity: event PoolSharesUpdated(uint256 indexed period, address indexed poolAddr, uint256 shares)
func (_RonStaking *RonStakingFilterer) WatchPoolSharesUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingPoolSharesUpdated, period []*big.Int, poolAddr []common.Address) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "PoolSharesUpdated", periodRule, poolAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingPoolSharesUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "PoolSharesUpdated", log); err != nil {
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

// ParsePoolSharesUpdated is a log parse operation binding the contract event 0x81faf50e2aaf52eaba2ab841071efb9f6f0850a3e7d008b1336e6001d3d4963c.
//
// Solidity: event PoolSharesUpdated(uint256 indexed period, address indexed poolAddr, uint256 shares)
func (_RonStaking *RonStakingFilterer) ParsePoolSharesUpdated(log types.Log) (*RonStakingPoolSharesUpdated, error) {
	event := new(RonStakingPoolSharesUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "PoolSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingPoolsDeprecatedIterator is returned from FilterPoolsDeprecated and is used to iterate over the raw logs and unpacked data for PoolsDeprecated events raised by the RonStaking contract.
type RonStakingPoolsDeprecatedIterator struct {
	Event *RonStakingPoolsDeprecated // Event containing the contract specifics and raw log

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
func (it *RonStakingPoolsDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingPoolsDeprecated)
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
		it.Event = new(RonStakingPoolsDeprecated)
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
func (it *RonStakingPoolsDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingPoolsDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingPoolsDeprecated represents a PoolsDeprecated event raised by the RonStaking contract.
type RonStakingPoolsDeprecated struct {
	Validator []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsDeprecated is a free log retrieval operation binding the contract event 0x4f257d3ba23679d338f1d94296086bba5724af341b7fa31aa0ff297bfcdc62d8.
//
// Solidity: event PoolsDeprecated(address[] validator)
func (_RonStaking *RonStakingFilterer) FilterPoolsDeprecated(opts *bind.FilterOpts) (*RonStakingPoolsDeprecatedIterator, error) {

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "PoolsDeprecated")
	if err != nil {
		return nil, err
	}
	return &RonStakingPoolsDeprecatedIterator{contract: _RonStaking.contract, event: "PoolsDeprecated", logs: logs, sub: sub}, nil
}

// WatchPoolsDeprecated is a free log subscription operation binding the contract event 0x4f257d3ba23679d338f1d94296086bba5724af341b7fa31aa0ff297bfcdc62d8.
//
// Solidity: event PoolsDeprecated(address[] validator)
func (_RonStaking *RonStakingFilterer) WatchPoolsDeprecated(opts *bind.WatchOpts, sink chan<- *RonStakingPoolsDeprecated) (event.Subscription, error) {

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "PoolsDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingPoolsDeprecated)
				if err := _RonStaking.contract.UnpackLog(event, "PoolsDeprecated", log); err != nil {
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

// ParsePoolsDeprecated is a log parse operation binding the contract event 0x4f257d3ba23679d338f1d94296086bba5724af341b7fa31aa0ff297bfcdc62d8.
//
// Solidity: event PoolsDeprecated(address[] validator)
func (_RonStaking *RonStakingFilterer) ParsePoolsDeprecated(log types.Log) (*RonStakingPoolsDeprecated, error) {
	event := new(RonStakingPoolsDeprecated)
	if err := _RonStaking.contract.UnpackLog(event, "PoolsDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingPoolsUpdateConflictedIterator is returned from FilterPoolsUpdateConflicted and is used to iterate over the raw logs and unpacked data for PoolsUpdateConflicted events raised by the RonStaking contract.
type RonStakingPoolsUpdateConflictedIterator struct {
	Event *RonStakingPoolsUpdateConflicted // Event containing the contract specifics and raw log

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
func (it *RonStakingPoolsUpdateConflictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingPoolsUpdateConflicted)
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
		it.Event = new(RonStakingPoolsUpdateConflicted)
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
func (it *RonStakingPoolsUpdateConflictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingPoolsUpdateConflictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingPoolsUpdateConflicted represents a PoolsUpdateConflicted event raised by the RonStaking contract.
type RonStakingPoolsUpdateConflicted struct {
	Period    *big.Int
	PoolAddrs []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsUpdateConflicted is a free log retrieval operation binding the contract event 0xee74f10cc50bf4b7e57fd36be7d46288795f3a9151dae97505b718b392ba14a3.
//
// Solidity: event PoolsUpdateConflicted(uint256 indexed period, address[] poolAddrs)
func (_RonStaking *RonStakingFilterer) FilterPoolsUpdateConflicted(opts *bind.FilterOpts, period []*big.Int) (*RonStakingPoolsUpdateConflictedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "PoolsUpdateConflicted", periodRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingPoolsUpdateConflictedIterator{contract: _RonStaking.contract, event: "PoolsUpdateConflicted", logs: logs, sub: sub}, nil
}

// WatchPoolsUpdateConflicted is a free log subscription operation binding the contract event 0xee74f10cc50bf4b7e57fd36be7d46288795f3a9151dae97505b718b392ba14a3.
//
// Solidity: event PoolsUpdateConflicted(uint256 indexed period, address[] poolAddrs)
func (_RonStaking *RonStakingFilterer) WatchPoolsUpdateConflicted(opts *bind.WatchOpts, sink chan<- *RonStakingPoolsUpdateConflicted, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "PoolsUpdateConflicted", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingPoolsUpdateConflicted)
				if err := _RonStaking.contract.UnpackLog(event, "PoolsUpdateConflicted", log); err != nil {
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

// ParsePoolsUpdateConflicted is a log parse operation binding the contract event 0xee74f10cc50bf4b7e57fd36be7d46288795f3a9151dae97505b718b392ba14a3.
//
// Solidity: event PoolsUpdateConflicted(uint256 indexed period, address[] poolAddrs)
func (_RonStaking *RonStakingFilterer) ParsePoolsUpdateConflicted(log types.Log) (*RonStakingPoolsUpdateConflicted, error) {
	event := new(RonStakingPoolsUpdateConflicted)
	if err := _RonStaking.contract.UnpackLog(event, "PoolsUpdateConflicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingPoolsUpdateFailedIterator is returned from FilterPoolsUpdateFailed and is used to iterate over the raw logs and unpacked data for PoolsUpdateFailed events raised by the RonStaking contract.
type RonStakingPoolsUpdateFailedIterator struct {
	Event *RonStakingPoolsUpdateFailed // Event containing the contract specifics and raw log

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
func (it *RonStakingPoolsUpdateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingPoolsUpdateFailed)
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
		it.Event = new(RonStakingPoolsUpdateFailed)
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
func (it *RonStakingPoolsUpdateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingPoolsUpdateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingPoolsUpdateFailed represents a PoolsUpdateFailed event raised by the RonStaking contract.
type RonStakingPoolsUpdateFailed struct {
	Period    *big.Int
	PoolAddrs []common.Address
	Rewards   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsUpdateFailed is a free log retrieval operation binding the contract event 0xae52c603227f64e4c6101dde593aa9790a16b3ac77546bd746d758511e9560a5.
//
// Solidity: event PoolsUpdateFailed(uint256 indexed period, address[] poolAddrs, uint256[] rewards)
func (_RonStaking *RonStakingFilterer) FilterPoolsUpdateFailed(opts *bind.FilterOpts, period []*big.Int) (*RonStakingPoolsUpdateFailedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "PoolsUpdateFailed", periodRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingPoolsUpdateFailedIterator{contract: _RonStaking.contract, event: "PoolsUpdateFailed", logs: logs, sub: sub}, nil
}

// WatchPoolsUpdateFailed is a free log subscription operation binding the contract event 0xae52c603227f64e4c6101dde593aa9790a16b3ac77546bd746d758511e9560a5.
//
// Solidity: event PoolsUpdateFailed(uint256 indexed period, address[] poolAddrs, uint256[] rewards)
func (_RonStaking *RonStakingFilterer) WatchPoolsUpdateFailed(opts *bind.WatchOpts, sink chan<- *RonStakingPoolsUpdateFailed, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "PoolsUpdateFailed", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingPoolsUpdateFailed)
				if err := _RonStaking.contract.UnpackLog(event, "PoolsUpdateFailed", log); err != nil {
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

// ParsePoolsUpdateFailed is a log parse operation binding the contract event 0xae52c603227f64e4c6101dde593aa9790a16b3ac77546bd746d758511e9560a5.
//
// Solidity: event PoolsUpdateFailed(uint256 indexed period, address[] poolAddrs, uint256[] rewards)
func (_RonStaking *RonStakingFilterer) ParsePoolsUpdateFailed(log types.Log) (*RonStakingPoolsUpdateFailed, error) {
	event := new(RonStakingPoolsUpdateFailed)
	if err := _RonStaking.contract.UnpackLog(event, "PoolsUpdateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingPoolsUpdatedIterator is returned from FilterPoolsUpdated and is used to iterate over the raw logs and unpacked data for PoolsUpdated events raised by the RonStaking contract.
type RonStakingPoolsUpdatedIterator struct {
	Event *RonStakingPoolsUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingPoolsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingPoolsUpdated)
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
		it.Event = new(RonStakingPoolsUpdated)
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
func (it *RonStakingPoolsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingPoolsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingPoolsUpdated represents a PoolsUpdated event raised by the RonStaking contract.
type RonStakingPoolsUpdated struct {
	Period    *big.Int
	PoolAddrs []common.Address
	ARps      []*big.Int
	Shares    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsUpdated is a free log retrieval operation binding the contract event 0x0e54e0485f70f0f63bc25889ddbf01ce1269ad6f07fdb2df573a0fbdb4d66f88.
//
// Solidity: event PoolsUpdated(uint256 indexed period, address[] poolAddrs, uint256[] aRps, uint256[] shares)
func (_RonStaking *RonStakingFilterer) FilterPoolsUpdated(opts *bind.FilterOpts, period []*big.Int) (*RonStakingPoolsUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "PoolsUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingPoolsUpdatedIterator{contract: _RonStaking.contract, event: "PoolsUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolsUpdated is a free log subscription operation binding the contract event 0x0e54e0485f70f0f63bc25889ddbf01ce1269ad6f07fdb2df573a0fbdb4d66f88.
//
// Solidity: event PoolsUpdated(uint256 indexed period, address[] poolAddrs, uint256[] aRps, uint256[] shares)
func (_RonStaking *RonStakingFilterer) WatchPoolsUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingPoolsUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "PoolsUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingPoolsUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "PoolsUpdated", log); err != nil {
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

// ParsePoolsUpdated is a log parse operation binding the contract event 0x0e54e0485f70f0f63bc25889ddbf01ce1269ad6f07fdb2df573a0fbdb4d66f88.
//
// Solidity: event PoolsUpdated(uint256 indexed period, address[] poolAddrs, uint256[] aRps, uint256[] shares)
func (_RonStaking *RonStakingFilterer) ParsePoolsUpdated(log types.Log) (*RonStakingPoolsUpdated, error) {
	event := new(RonStakingPoolsUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "PoolsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the RonStaking contract.
type RonStakingRewardClaimedIterator struct {
	Event *RonStakingRewardClaimed // Event containing the contract specifics and raw log

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
func (it *RonStakingRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingRewardClaimed)
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
		it.Event = new(RonStakingRewardClaimed)
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
func (it *RonStakingRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingRewardClaimed represents a RewardClaimed event raised by the RonStaking contract.
type RonStakingRewardClaimed struct {
	PoolAddr common.Address
	User     common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed poolAddr, address indexed user, uint256 amount)
func (_RonStaking *RonStakingFilterer) FilterRewardClaimed(opts *bind.FilterOpts, poolAddr []common.Address, user []common.Address) (*RonStakingRewardClaimedIterator, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "RewardClaimed", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingRewardClaimedIterator{contract: _RonStaking.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed poolAddr, address indexed user, uint256 amount)
func (_RonStaking *RonStakingFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *RonStakingRewardClaimed, poolAddr []common.Address, user []common.Address) (event.Subscription, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "RewardClaimed", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingRewardClaimed)
				if err := _RonStaking.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed poolAddr, address indexed user, uint256 amount)
func (_RonStaking *RonStakingFilterer) ParseRewardClaimed(log types.Log) (*RonStakingRewardClaimed, error) {
	event := new(RonStakingRewardClaimed)
	if err := _RonStaking.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the RonStaking contract.
type RonStakingStakedIterator struct {
	Event *RonStakingStaked // Event containing the contract specifics and raw log

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
func (it *RonStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingStaked)
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
		it.Event = new(RonStakingStaked)
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
func (it *RonStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingStaked represents a Staked event raised by the RonStaking contract.
type RonStakingStaked struct {
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) FilterStaked(opts *bind.FilterOpts, consensuAddr []common.Address) (*RonStakingStakedIterator, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "Staked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingStakedIterator{contract: _RonStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *RonStakingStaked, consensuAddr []common.Address) (event.Subscription, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "Staked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingStaked)
				if err := _RonStaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) ParseStaked(log types.Log) (*RonStakingStaked, error) {
	event := new(RonStakingStaked)
	if err := _RonStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingStakingAmountTransferFailedIterator is returned from FilterStakingAmountTransferFailed and is used to iterate over the raw logs and unpacked data for StakingAmountTransferFailed events raised by the RonStaking contract.
type RonStakingStakingAmountTransferFailedIterator struct {
	Event *RonStakingStakingAmountTransferFailed // Event containing the contract specifics and raw log

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
func (it *RonStakingStakingAmountTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingStakingAmountTransferFailed)
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
		it.Event = new(RonStakingStakingAmountTransferFailed)
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
func (it *RonStakingStakingAmountTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingStakingAmountTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingStakingAmountTransferFailed represents a StakingAmountTransferFailed event raised by the RonStaking contract.
type RonStakingStakingAmountTransferFailed struct {
	Validator       common.Address
	Admin           common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingAmountTransferFailed is a free log retrieval operation binding the contract event 0x7dc5115a5aba081f5a174f56a3d02eea582824783322a4ac03f7bd388f444194.
//
// Solidity: event StakingAmountTransferFailed(address indexed validator, address indexed admin, uint256 amount, uint256 contractBalance)
func (_RonStaking *RonStakingFilterer) FilterStakingAmountTransferFailed(opts *bind.FilterOpts, validator []common.Address, admin []common.Address) (*RonStakingStakingAmountTransferFailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "StakingAmountTransferFailed", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingStakingAmountTransferFailedIterator{contract: _RonStaking.contract, event: "StakingAmountTransferFailed", logs: logs, sub: sub}, nil
}

// WatchStakingAmountTransferFailed is a free log subscription operation binding the contract event 0x7dc5115a5aba081f5a174f56a3d02eea582824783322a4ac03f7bd388f444194.
//
// Solidity: event StakingAmountTransferFailed(address indexed validator, address indexed admin, uint256 amount, uint256 contractBalance)
func (_RonStaking *RonStakingFilterer) WatchStakingAmountTransferFailed(opts *bind.WatchOpts, sink chan<- *RonStakingStakingAmountTransferFailed, validator []common.Address, admin []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "StakingAmountTransferFailed", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingStakingAmountTransferFailed)
				if err := _RonStaking.contract.UnpackLog(event, "StakingAmountTransferFailed", log); err != nil {
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

// ParseStakingAmountTransferFailed is a log parse operation binding the contract event 0x7dc5115a5aba081f5a174f56a3d02eea582824783322a4ac03f7bd388f444194.
//
// Solidity: event StakingAmountTransferFailed(address indexed validator, address indexed admin, uint256 amount, uint256 contractBalance)
func (_RonStaking *RonStakingFilterer) ParseStakingAmountTransferFailed(log types.Log) (*RonStakingStakingAmountTransferFailed, error) {
	event := new(RonStakingStakingAmountTransferFailed)
	if err := _RonStaking.contract.UnpackLog(event, "StakingAmountTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the RonStaking contract.
type RonStakingUndelegatedIterator struct {
	Event *RonStakingUndelegated // Event containing the contract specifics and raw log

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
func (it *RonStakingUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingUndelegated)
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
		it.Event = new(RonStakingUndelegated)
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
func (it *RonStakingUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingUndelegated represents a Undelegated event raised by the RonStaking contract.
type RonStakingUndelegated struct {
	Delegator    common.Address
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, consensuAddr []common.Address) (*RonStakingUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "Undelegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingUndelegatedIterator{contract: _RonStaking.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *RonStakingUndelegated, delegator []common.Address, consensuAddr []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "Undelegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingUndelegated)
				if err := _RonStaking.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) ParseUndelegated(log types.Log) (*RonStakingUndelegated, error) {
	event := new(RonStakingUndelegated)
	if err := _RonStaking.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the RonStaking contract.
type RonStakingUnstakedIterator struct {
	Event *RonStakingUnstaked // Event containing the contract specifics and raw log

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
func (it *RonStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingUnstaked)
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
		it.Event = new(RonStakingUnstaked)
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
func (it *RonStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingUnstaked represents a Unstaked event raised by the RonStaking contract.
type RonStakingUnstaked struct {
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, consensuAddr []common.Address) (*RonStakingUnstakedIterator, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "Unstaked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingUnstakedIterator{contract: _RonStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *RonStakingUnstaked, consensuAddr []common.Address) (event.Subscription, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "Unstaked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingUnstaked)
				if err := _RonStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed consensuAddr, uint256 amount)
func (_RonStaking *RonStakingFilterer) ParseUnstaked(log types.Log) (*RonStakingUnstaked, error) {
	event := new(RonStakingUnstaked)
	if err := _RonStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingUserRewardUpdatedIterator is returned from FilterUserRewardUpdated and is used to iterate over the raw logs and unpacked data for UserRewardUpdated events raised by the RonStaking contract.
type RonStakingUserRewardUpdatedIterator struct {
	Event *RonStakingUserRewardUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingUserRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingUserRewardUpdated)
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
		it.Event = new(RonStakingUserRewardUpdated)
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
func (it *RonStakingUserRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingUserRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingUserRewardUpdated represents a UserRewardUpdated event raised by the RonStaking contract.
type RonStakingUserRewardUpdated struct {
	PoolAddr common.Address
	User     common.Address
	Debited  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUserRewardUpdated is a free log retrieval operation binding the contract event 0xaa7c29611027fd4be148712bb54960253b7a7d5998c17769bfc424c2f5f185ad.
//
// Solidity: event UserRewardUpdated(address indexed poolAddr, address indexed user, uint256 debited)
func (_RonStaking *RonStakingFilterer) FilterUserRewardUpdated(opts *bind.FilterOpts, poolAddr []common.Address, user []common.Address) (*RonStakingUserRewardUpdatedIterator, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "UserRewardUpdated", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingUserRewardUpdatedIterator{contract: _RonStaking.contract, event: "UserRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchUserRewardUpdated is a free log subscription operation binding the contract event 0xaa7c29611027fd4be148712bb54960253b7a7d5998c17769bfc424c2f5f185ad.
//
// Solidity: event UserRewardUpdated(address indexed poolAddr, address indexed user, uint256 debited)
func (_RonStaking *RonStakingFilterer) WatchUserRewardUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingUserRewardUpdated, poolAddr []common.Address, user []common.Address) (event.Subscription, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "UserRewardUpdated", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingUserRewardUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "UserRewardUpdated", log); err != nil {
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

// ParseUserRewardUpdated is a log parse operation binding the contract event 0xaa7c29611027fd4be148712bb54960253b7a7d5998c17769bfc424c2f5f185ad.
//
// Solidity: event UserRewardUpdated(address indexed poolAddr, address indexed user, uint256 debited)
func (_RonStaking *RonStakingFilterer) ParseUserRewardUpdated(log types.Log) (*RonStakingUserRewardUpdated, error) {
	event := new(RonStakingUserRewardUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "UserRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the RonStaking contract.
type RonStakingValidatorContractUpdatedIterator struct {
	Event *RonStakingValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingValidatorContractUpdated)
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
		it.Event = new(RonStakingValidatorContractUpdated)
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
func (it *RonStakingValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingValidatorContractUpdated represents a ValidatorContractUpdated event raised by the RonStaking contract.
type RonStakingValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_RonStaking *RonStakingFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*RonStakingValidatorContractUpdatedIterator, error) {

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RonStakingValidatorContractUpdatedIterator{contract: _RonStaking.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_RonStaking *RonStakingFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingValidatorContractUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_RonStaking *RonStakingFilterer) ParseValidatorContractUpdated(log types.Log) (*RonStakingValidatorContractUpdated, error) {
	event := new(RonStakingValidatorContractUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingWaitingSecsToRevokeUpdatedIterator is returned from FilterWaitingSecsToRevokeUpdated and is used to iterate over the raw logs and unpacked data for WaitingSecsToRevokeUpdated events raised by the RonStaking contract.
type RonStakingWaitingSecsToRevokeUpdatedIterator struct {
	Event *RonStakingWaitingSecsToRevokeUpdated // Event containing the contract specifics and raw log

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
func (it *RonStakingWaitingSecsToRevokeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingWaitingSecsToRevokeUpdated)
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
		it.Event = new(RonStakingWaitingSecsToRevokeUpdated)
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
func (it *RonStakingWaitingSecsToRevokeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingWaitingSecsToRevokeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingWaitingSecsToRevokeUpdated represents a WaitingSecsToRevokeUpdated event raised by the RonStaking contract.
type RonStakingWaitingSecsToRevokeUpdated struct {
	Secs *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWaitingSecsToRevokeUpdated is a free log retrieval operation binding the contract event 0x02be0b73b597f2c0f138aebee162b3b0e25d5b5a26854c15dcf79176e9a1c678.
//
// Solidity: event WaitingSecsToRevokeUpdated(uint256 secs)
func (_RonStaking *RonStakingFilterer) FilterWaitingSecsToRevokeUpdated(opts *bind.FilterOpts) (*RonStakingWaitingSecsToRevokeUpdatedIterator, error) {

	logs, sub, err := _RonStaking.contract.FilterLogs(opts, "WaitingSecsToRevokeUpdated")
	if err != nil {
		return nil, err
	}
	return &RonStakingWaitingSecsToRevokeUpdatedIterator{contract: _RonStaking.contract, event: "WaitingSecsToRevokeUpdated", logs: logs, sub: sub}, nil
}

// WatchWaitingSecsToRevokeUpdated is a free log subscription operation binding the contract event 0x02be0b73b597f2c0f138aebee162b3b0e25d5b5a26854c15dcf79176e9a1c678.
//
// Solidity: event WaitingSecsToRevokeUpdated(uint256 secs)
func (_RonStaking *RonStakingFilterer) WatchWaitingSecsToRevokeUpdated(opts *bind.WatchOpts, sink chan<- *RonStakingWaitingSecsToRevokeUpdated) (event.Subscription, error) {

	logs, sub, err := _RonStaking.contract.WatchLogs(opts, "WaitingSecsToRevokeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingWaitingSecsToRevokeUpdated)
				if err := _RonStaking.contract.UnpackLog(event, "WaitingSecsToRevokeUpdated", log); err != nil {
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

// ParseWaitingSecsToRevokeUpdated is a log parse operation binding the contract event 0x02be0b73b597f2c0f138aebee162b3b0e25d5b5a26854c15dcf79176e9a1c678.
//
// Solidity: event WaitingSecsToRevokeUpdated(uint256 secs)
func (_RonStaking *RonStakingFilterer) ParseWaitingSecsToRevokeUpdated(log types.Log) (*RonStakingWaitingSecsToRevokeUpdated, error) {
	event := new(RonStakingWaitingSecsToRevokeUpdated)
	if err := _RonStaking.contract.UnpackLog(event, "WaitingSecsToRevokeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
