// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingVesting

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

// StakingVestingMetaData contains all meta data concerning the StakingVesting contract.
var StakingVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BlockProducerBonusPerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockProducerAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeOperatorAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"BonusTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockProducerAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgeOperatorAmount\",\"type\":\"uint256\"}],\"name\":\"BonusTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorBonusPerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockProducerBlockBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeOperatorBlockBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__blockProducerBonusPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__bridgeOperatorBonusPerBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlockSendingBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveRON\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_forBlockProducer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_forBridgeOperator\",\"type\":\"bool\"}],\"name\":\"requestBonus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_blockProducerBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bridgeOperatorBonus\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBlockProducerBonusPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBridgeOperatorBonusPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingVestingMetaData.ABI instead.
var StakingVestingABI = StakingVestingMetaData.ABI

// StakingVesting is an auto generated Go binding around an Ethereum contract.
type StakingVesting struct {
	StakingVestingCaller     // Read-only binding to the contract
	StakingVestingTransactor // Write-only binding to the contract
	StakingVestingFilterer   // Log filterer for contract events
}

// StakingVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingVestingSession struct {
	Contract     *StakingVesting   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingVestingCallerSession struct {
	Contract *StakingVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingVestingTransactorSession struct {
	Contract     *StakingVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingVestingRaw struct {
	Contract *StakingVesting // Generic contract binding to access the raw methods on
}

// StakingVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingVestingCallerRaw struct {
	Contract *StakingVestingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingVestingTransactorRaw struct {
	Contract *StakingVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingVesting creates a new instance of StakingVesting, bound to a specific deployed contract.
func NewStakingVesting(address common.Address, backend bind.ContractBackend) (*StakingVesting, error) {
	contract, err := bindStakingVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingVesting{StakingVestingCaller: StakingVestingCaller{contract: contract}, StakingVestingTransactor: StakingVestingTransactor{contract: contract}, StakingVestingFilterer: StakingVestingFilterer{contract: contract}}, nil
}

// NewStakingVestingCaller creates a new read-only instance of StakingVesting, bound to a specific deployed contract.
func NewStakingVestingCaller(address common.Address, caller bind.ContractCaller) (*StakingVestingCaller, error) {
	contract, err := bindStakingVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingVestingCaller{contract: contract}, nil
}

// NewStakingVestingTransactor creates a new write-only instance of StakingVesting, bound to a specific deployed contract.
func NewStakingVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingVestingTransactor, error) {
	contract, err := bindStakingVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingVestingTransactor{contract: contract}, nil
}

// NewStakingVestingFilterer creates a new log filterer instance of StakingVesting, bound to a specific deployed contract.
func NewStakingVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingVestingFilterer, error) {
	contract, err := bindStakingVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingVestingFilterer{contract: contract}, nil
}

// bindStakingVesting binds a generic wrapper to an already deployed contract.
func bindStakingVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingVesting *StakingVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingVesting.Contract.StakingVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingVesting *StakingVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVesting.Contract.StakingVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingVesting *StakingVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingVesting.Contract.StakingVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingVesting *StakingVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingVesting *StakingVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingVesting *StakingVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingVesting.Contract.contract.Transact(opts, method, params...)
}

// BlockProducerBlockBonus is a free data retrieval call binding the contract method 0xd8209d07.
//
// Solidity: function blockProducerBlockBonus(uint256 ) view returns(uint256)
func (_StakingVesting *StakingVestingCaller) BlockProducerBlockBonus(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingVesting.contract.Call(opts, &out, "blockProducerBlockBonus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockProducerBlockBonus is a free data retrieval call binding the contract method 0xd8209d07.
//
// Solidity: function blockProducerBlockBonus(uint256 ) view returns(uint256)
func (_StakingVesting *StakingVestingSession) BlockProducerBlockBonus(arg0 *big.Int) (*big.Int, error) {
	return _StakingVesting.Contract.BlockProducerBlockBonus(&_StakingVesting.CallOpts, arg0)
}

// BlockProducerBlockBonus is a free data retrieval call binding the contract method 0xd8209d07.
//
// Solidity: function blockProducerBlockBonus(uint256 ) view returns(uint256)
func (_StakingVesting *StakingVestingCallerSession) BlockProducerBlockBonus(arg0 *big.Int) (*big.Int, error) {
	return _StakingVesting.Contract.BlockProducerBlockBonus(&_StakingVesting.CallOpts, arg0)
}

// BridgeOperatorBlockBonus is a free data retrieval call binding the contract method 0xfa8674a1.
//
// Solidity: function bridgeOperatorBlockBonus(uint256 ) view returns(uint256)
func (_StakingVesting *StakingVestingCaller) BridgeOperatorBlockBonus(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingVesting.contract.Call(opts, &out, "bridgeOperatorBlockBonus", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeOperatorBlockBonus is a free data retrieval call binding the contract method 0xfa8674a1.
//
// Solidity: function bridgeOperatorBlockBonus(uint256 ) view returns(uint256)
func (_StakingVesting *StakingVestingSession) BridgeOperatorBlockBonus(arg0 *big.Int) (*big.Int, error) {
	return _StakingVesting.Contract.BridgeOperatorBlockBonus(&_StakingVesting.CallOpts, arg0)
}

// BridgeOperatorBlockBonus is a free data retrieval call binding the contract method 0xfa8674a1.
//
// Solidity: function bridgeOperatorBlockBonus(uint256 ) view returns(uint256)
func (_StakingVesting *StakingVestingCallerSession) BridgeOperatorBlockBonus(arg0 *big.Int) (*big.Int, error) {
	return _StakingVesting.Contract.BridgeOperatorBlockBonus(&_StakingVesting.CallOpts, arg0)
}

// LastBlockSendingBonus is a free data retrieval call binding the contract method 0x0d9160e7.
//
// Solidity: function lastBlockSendingBonus() view returns(uint256)
func (_StakingVesting *StakingVestingCaller) LastBlockSendingBonus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVesting.contract.Call(opts, &out, "lastBlockSendingBonus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlockSendingBonus is a free data retrieval call binding the contract method 0x0d9160e7.
//
// Solidity: function lastBlockSendingBonus() view returns(uint256)
func (_StakingVesting *StakingVestingSession) LastBlockSendingBonus() (*big.Int, error) {
	return _StakingVesting.Contract.LastBlockSendingBonus(&_StakingVesting.CallOpts)
}

// LastBlockSendingBonus is a free data retrieval call binding the contract method 0x0d9160e7.
//
// Solidity: function lastBlockSendingBonus() view returns(uint256)
func (_StakingVesting *StakingVestingCallerSession) LastBlockSendingBonus() (*big.Int, error) {
	return _StakingVesting.Contract.LastBlockSendingBonus(&_StakingVesting.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_StakingVesting *StakingVestingCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingVesting.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_StakingVesting *StakingVestingSession) ValidatorContract() (common.Address, error) {
	return _StakingVesting.Contract.ValidatorContract(&_StakingVesting.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_StakingVesting *StakingVestingCallerSession) ValidatorContract() (common.Address, error) {
	return _StakingVesting.Contract.ValidatorContract(&_StakingVesting.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address __validatorContract, uint256 __blockProducerBonusPerBlock, uint256 __bridgeOperatorBonusPerBlock) payable returns()
func (_StakingVesting *StakingVestingTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __blockProducerBonusPerBlock *big.Int, __bridgeOperatorBonusPerBlock *big.Int) (*types.Transaction, error) {
	return _StakingVesting.contract.Transact(opts, "initialize", __validatorContract, __blockProducerBonusPerBlock, __bridgeOperatorBonusPerBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address __validatorContract, uint256 __blockProducerBonusPerBlock, uint256 __bridgeOperatorBonusPerBlock) payable returns()
func (_StakingVesting *StakingVestingSession) Initialize(__validatorContract common.Address, __blockProducerBonusPerBlock *big.Int, __bridgeOperatorBonusPerBlock *big.Int) (*types.Transaction, error) {
	return _StakingVesting.Contract.Initialize(&_StakingVesting.TransactOpts, __validatorContract, __blockProducerBonusPerBlock, __bridgeOperatorBonusPerBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address __validatorContract, uint256 __blockProducerBonusPerBlock, uint256 __bridgeOperatorBonusPerBlock) payable returns()
func (_StakingVesting *StakingVestingTransactorSession) Initialize(__validatorContract common.Address, __blockProducerBonusPerBlock *big.Int, __bridgeOperatorBonusPerBlock *big.Int) (*types.Transaction, error) {
	return _StakingVesting.Contract.Initialize(&_StakingVesting.TransactOpts, __validatorContract, __blockProducerBonusPerBlock, __bridgeOperatorBonusPerBlock)
}

// ReceiveRON is a paid mutator transaction binding the contract method 0x59f778df.
//
// Solidity: function receiveRON() payable returns()
func (_StakingVesting *StakingVestingTransactor) ReceiveRON(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVesting.contract.Transact(opts, "receiveRON")
}

// ReceiveRON is a paid mutator transaction binding the contract method 0x59f778df.
//
// Solidity: function receiveRON() payable returns()
func (_StakingVesting *StakingVestingSession) ReceiveRON() (*types.Transaction, error) {
	return _StakingVesting.Contract.ReceiveRON(&_StakingVesting.TransactOpts)
}

// ReceiveRON is a paid mutator transaction binding the contract method 0x59f778df.
//
// Solidity: function receiveRON() payable returns()
func (_StakingVesting *StakingVestingTransactorSession) ReceiveRON() (*types.Transaction, error) {
	return _StakingVesting.Contract.ReceiveRON(&_StakingVesting.TransactOpts)
}

// RequestBonus is a paid mutator transaction binding the contract method 0x0634f5b9.
//
// Solidity: function requestBonus(bool _forBlockProducer, bool _forBridgeOperator) returns(bool _success, uint256 _blockProducerBonus, uint256 _bridgeOperatorBonus)
func (_StakingVesting *StakingVestingTransactor) RequestBonus(opts *bind.TransactOpts, _forBlockProducer bool, _forBridgeOperator bool) (*types.Transaction, error) {
	return _StakingVesting.contract.Transact(opts, "requestBonus", _forBlockProducer, _forBridgeOperator)
}

// RequestBonus is a paid mutator transaction binding the contract method 0x0634f5b9.
//
// Solidity: function requestBonus(bool _forBlockProducer, bool _forBridgeOperator) returns(bool _success, uint256 _blockProducerBonus, uint256 _bridgeOperatorBonus)
func (_StakingVesting *StakingVestingSession) RequestBonus(_forBlockProducer bool, _forBridgeOperator bool) (*types.Transaction, error) {
	return _StakingVesting.Contract.RequestBonus(&_StakingVesting.TransactOpts, _forBlockProducer, _forBridgeOperator)
}

// RequestBonus is a paid mutator transaction binding the contract method 0x0634f5b9.
//
// Solidity: function requestBonus(bool _forBlockProducer, bool _forBridgeOperator) returns(bool _success, uint256 _blockProducerBonus, uint256 _bridgeOperatorBonus)
func (_StakingVesting *StakingVestingTransactorSession) RequestBonus(_forBlockProducer bool, _forBridgeOperator bool) (*types.Transaction, error) {
	return _StakingVesting.Contract.RequestBonus(&_StakingVesting.TransactOpts, _forBlockProducer, _forBridgeOperator)
}

// SetBlockProducerBonusPerBlock is a paid mutator transaction binding the contract method 0xf13ba644.
//
// Solidity: function setBlockProducerBonusPerBlock(uint256 _amount) returns()
func (_StakingVesting *StakingVestingTransactor) SetBlockProducerBonusPerBlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StakingVesting.contract.Transact(opts, "setBlockProducerBonusPerBlock", _amount)
}

// SetBlockProducerBonusPerBlock is a paid mutator transaction binding the contract method 0xf13ba644.
//
// Solidity: function setBlockProducerBonusPerBlock(uint256 _amount) returns()
func (_StakingVesting *StakingVestingSession) SetBlockProducerBonusPerBlock(_amount *big.Int) (*types.Transaction, error) {
	return _StakingVesting.Contract.SetBlockProducerBonusPerBlock(&_StakingVesting.TransactOpts, _amount)
}

// SetBlockProducerBonusPerBlock is a paid mutator transaction binding the contract method 0xf13ba644.
//
// Solidity: function setBlockProducerBonusPerBlock(uint256 _amount) returns()
func (_StakingVesting *StakingVestingTransactorSession) SetBlockProducerBonusPerBlock(_amount *big.Int) (*types.Transaction, error) {
	return _StakingVesting.Contract.SetBlockProducerBonusPerBlock(&_StakingVesting.TransactOpts, _amount)
}

// SetBridgeOperatorBonusPerBlock is a paid mutator transaction binding the contract method 0x67e9941c.
//
// Solidity: function setBridgeOperatorBonusPerBlock(uint256 _amount) returns()
func (_StakingVesting *StakingVestingTransactor) SetBridgeOperatorBonusPerBlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StakingVesting.contract.Transact(opts, "setBridgeOperatorBonusPerBlock", _amount)
}

// SetBridgeOperatorBonusPerBlock is a paid mutator transaction binding the contract method 0x67e9941c.
//
// Solidity: function setBridgeOperatorBonusPerBlock(uint256 _amount) returns()
func (_StakingVesting *StakingVestingSession) SetBridgeOperatorBonusPerBlock(_amount *big.Int) (*types.Transaction, error) {
	return _StakingVesting.Contract.SetBridgeOperatorBonusPerBlock(&_StakingVesting.TransactOpts, _amount)
}

// SetBridgeOperatorBonusPerBlock is a paid mutator transaction binding the contract method 0x67e9941c.
//
// Solidity: function setBridgeOperatorBonusPerBlock(uint256 _amount) returns()
func (_StakingVesting *StakingVestingTransactorSession) SetBridgeOperatorBonusPerBlock(_amount *big.Int) (*types.Transaction, error) {
	return _StakingVesting.Contract.SetBridgeOperatorBonusPerBlock(&_StakingVesting.TransactOpts, _amount)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_StakingVesting *StakingVestingTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _StakingVesting.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_StakingVesting *StakingVestingSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _StakingVesting.Contract.SetValidatorContract(&_StakingVesting.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_StakingVesting *StakingVestingTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _StakingVesting.Contract.SetValidatorContract(&_StakingVesting.TransactOpts, _addr)
}

// StakingVestingBlockProducerBonusPerBlockUpdatedIterator is returned from FilterBlockProducerBonusPerBlockUpdated and is used to iterate over the raw logs and unpacked data for BlockProducerBonusPerBlockUpdated events raised by the StakingVesting contract.
type StakingVestingBlockProducerBonusPerBlockUpdatedIterator struct {
	Event *StakingVestingBlockProducerBonusPerBlockUpdated // Event containing the contract specifics and raw log

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
func (it *StakingVestingBlockProducerBonusPerBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVestingBlockProducerBonusPerBlockUpdated)
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
		it.Event = new(StakingVestingBlockProducerBonusPerBlockUpdated)
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
func (it *StakingVestingBlockProducerBonusPerBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVestingBlockProducerBonusPerBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVestingBlockProducerBonusPerBlockUpdated represents a BlockProducerBonusPerBlockUpdated event raised by the StakingVesting contract.
type StakingVestingBlockProducerBonusPerBlockUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlockProducerBonusPerBlockUpdated is a free log retrieval operation binding the contract event 0x861f03c645467325a586235bb3155834f1dddf12413d0a802f416eb6d4035e6d.
//
// Solidity: event BlockProducerBonusPerBlockUpdated(uint256 arg0)
func (_StakingVesting *StakingVestingFilterer) FilterBlockProducerBonusPerBlockUpdated(opts *bind.FilterOpts) (*StakingVestingBlockProducerBonusPerBlockUpdatedIterator, error) {

	logs, sub, err := _StakingVesting.contract.FilterLogs(opts, "BlockProducerBonusPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingVestingBlockProducerBonusPerBlockUpdatedIterator{contract: _StakingVesting.contract, event: "BlockProducerBonusPerBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockProducerBonusPerBlockUpdated is a free log subscription operation binding the contract event 0x861f03c645467325a586235bb3155834f1dddf12413d0a802f416eb6d4035e6d.
//
// Solidity: event BlockProducerBonusPerBlockUpdated(uint256 arg0)
func (_StakingVesting *StakingVestingFilterer) WatchBlockProducerBonusPerBlockUpdated(opts *bind.WatchOpts, sink chan<- *StakingVestingBlockProducerBonusPerBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingVesting.contract.WatchLogs(opts, "BlockProducerBonusPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVestingBlockProducerBonusPerBlockUpdated)
				if err := _StakingVesting.contract.UnpackLog(event, "BlockProducerBonusPerBlockUpdated", log); err != nil {
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

// ParseBlockProducerBonusPerBlockUpdated is a log parse operation binding the contract event 0x861f03c645467325a586235bb3155834f1dddf12413d0a802f416eb6d4035e6d.
//
// Solidity: event BlockProducerBonusPerBlockUpdated(uint256 arg0)
func (_StakingVesting *StakingVestingFilterer) ParseBlockProducerBonusPerBlockUpdated(log types.Log) (*StakingVestingBlockProducerBonusPerBlockUpdated, error) {
	event := new(StakingVestingBlockProducerBonusPerBlockUpdated)
	if err := _StakingVesting.contract.UnpackLog(event, "BlockProducerBonusPerBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVestingBonusTransferFailedIterator is returned from FilterBonusTransferFailed and is used to iterate over the raw logs and unpacked data for BonusTransferFailed events raised by the StakingVesting contract.
type StakingVestingBonusTransferFailedIterator struct {
	Event *StakingVestingBonusTransferFailed // Event containing the contract specifics and raw log

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
func (it *StakingVestingBonusTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVestingBonusTransferFailed)
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
		it.Event = new(StakingVestingBonusTransferFailed)
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
func (it *StakingVestingBonusTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVestingBonusTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVestingBonusTransferFailed represents a BonusTransferFailed event raised by the StakingVesting contract.
type StakingVestingBonusTransferFailed struct {
	BlockNumber          *big.Int
	Recipient            common.Address
	BlockProducerAmount  *big.Int
	BridgeOperatorAmount *big.Int
	ContractBalance      *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBonusTransferFailed is a free log retrieval operation binding the contract event 0x137e697384eeada9cf7614b88e4ac940aeff18d0fef7e86bce1abdc812b95e09.
//
// Solidity: event BonusTransferFailed(uint256 indexed blockNumber, address indexed recipient, uint256 blockProducerAmount, uint256 bridgeOperatorAmount, uint256 contractBalance)
func (_StakingVesting *StakingVestingFilterer) FilterBonusTransferFailed(opts *bind.FilterOpts, blockNumber []*big.Int, recipient []common.Address) (*StakingVestingBonusTransferFailedIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakingVesting.contract.FilterLogs(opts, "BonusTransferFailed", blockNumberRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingVestingBonusTransferFailedIterator{contract: _StakingVesting.contract, event: "BonusTransferFailed", logs: logs, sub: sub}, nil
}

// WatchBonusTransferFailed is a free log subscription operation binding the contract event 0x137e697384eeada9cf7614b88e4ac940aeff18d0fef7e86bce1abdc812b95e09.
//
// Solidity: event BonusTransferFailed(uint256 indexed blockNumber, address indexed recipient, uint256 blockProducerAmount, uint256 bridgeOperatorAmount, uint256 contractBalance)
func (_StakingVesting *StakingVestingFilterer) WatchBonusTransferFailed(opts *bind.WatchOpts, sink chan<- *StakingVestingBonusTransferFailed, blockNumber []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakingVesting.contract.WatchLogs(opts, "BonusTransferFailed", blockNumberRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVestingBonusTransferFailed)
				if err := _StakingVesting.contract.UnpackLog(event, "BonusTransferFailed", log); err != nil {
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

// ParseBonusTransferFailed is a log parse operation binding the contract event 0x137e697384eeada9cf7614b88e4ac940aeff18d0fef7e86bce1abdc812b95e09.
//
// Solidity: event BonusTransferFailed(uint256 indexed blockNumber, address indexed recipient, uint256 blockProducerAmount, uint256 bridgeOperatorAmount, uint256 contractBalance)
func (_StakingVesting *StakingVestingFilterer) ParseBonusTransferFailed(log types.Log) (*StakingVestingBonusTransferFailed, error) {
	event := new(StakingVestingBonusTransferFailed)
	if err := _StakingVesting.contract.UnpackLog(event, "BonusTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVestingBonusTransferredIterator is returned from FilterBonusTransferred and is used to iterate over the raw logs and unpacked data for BonusTransferred events raised by the StakingVesting contract.
type StakingVestingBonusTransferredIterator struct {
	Event *StakingVestingBonusTransferred // Event containing the contract specifics and raw log

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
func (it *StakingVestingBonusTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVestingBonusTransferred)
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
		it.Event = new(StakingVestingBonusTransferred)
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
func (it *StakingVestingBonusTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVestingBonusTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVestingBonusTransferred represents a BonusTransferred event raised by the StakingVesting contract.
type StakingVestingBonusTransferred struct {
	BlockNumber          *big.Int
	Recipient            common.Address
	BlockProducerAmount  *big.Int
	BridgeOperatorAmount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBonusTransferred is a free log retrieval operation binding the contract event 0x60200441f885b45b3b7f1fdc45a47bb0d0a0884a6a17722f8dd7232830de9bd2.
//
// Solidity: event BonusTransferred(uint256 indexed blockNumber, address indexed recipient, uint256 blockProducerAmount, uint256 bridgeOperatorAmount)
func (_StakingVesting *StakingVestingFilterer) FilterBonusTransferred(opts *bind.FilterOpts, blockNumber []*big.Int, recipient []common.Address) (*StakingVestingBonusTransferredIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakingVesting.contract.FilterLogs(opts, "BonusTransferred", blockNumberRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingVestingBonusTransferredIterator{contract: _StakingVesting.contract, event: "BonusTransferred", logs: logs, sub: sub}, nil
}

// WatchBonusTransferred is a free log subscription operation binding the contract event 0x60200441f885b45b3b7f1fdc45a47bb0d0a0884a6a17722f8dd7232830de9bd2.
//
// Solidity: event BonusTransferred(uint256 indexed blockNumber, address indexed recipient, uint256 blockProducerAmount, uint256 bridgeOperatorAmount)
func (_StakingVesting *StakingVestingFilterer) WatchBonusTransferred(opts *bind.WatchOpts, sink chan<- *StakingVestingBonusTransferred, blockNumber []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakingVesting.contract.WatchLogs(opts, "BonusTransferred", blockNumberRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVestingBonusTransferred)
				if err := _StakingVesting.contract.UnpackLog(event, "BonusTransferred", log); err != nil {
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

// ParseBonusTransferred is a log parse operation binding the contract event 0x60200441f885b45b3b7f1fdc45a47bb0d0a0884a6a17722f8dd7232830de9bd2.
//
// Solidity: event BonusTransferred(uint256 indexed blockNumber, address indexed recipient, uint256 blockProducerAmount, uint256 bridgeOperatorAmount)
func (_StakingVesting *StakingVestingFilterer) ParseBonusTransferred(log types.Log) (*StakingVestingBonusTransferred, error) {
	event := new(StakingVestingBonusTransferred)
	if err := _StakingVesting.contract.UnpackLog(event, "BonusTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator is returned from FilterBridgeOperatorBonusPerBlockUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorBonusPerBlockUpdated events raised by the StakingVesting contract.
type StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator struct {
	Event *StakingVestingBridgeOperatorBonusPerBlockUpdated // Event containing the contract specifics and raw log

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
func (it *StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVestingBridgeOperatorBonusPerBlockUpdated)
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
		it.Event = new(StakingVestingBridgeOperatorBonusPerBlockUpdated)
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
func (it *StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVestingBridgeOperatorBonusPerBlockUpdated represents a BridgeOperatorBonusPerBlockUpdated event raised by the StakingVesting contract.
type StakingVestingBridgeOperatorBonusPerBlockUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorBonusPerBlockUpdated is a free log retrieval operation binding the contract event 0x57a23b4b11f619fb9dea21a5a8115bb90103c1043eb3318d773558829d25f12c.
//
// Solidity: event BridgeOperatorBonusPerBlockUpdated(uint256 arg0)
func (_StakingVesting *StakingVestingFilterer) FilterBridgeOperatorBonusPerBlockUpdated(opts *bind.FilterOpts) (*StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator, error) {

	logs, sub, err := _StakingVesting.contract.FilterLogs(opts, "BridgeOperatorBonusPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingVestingBridgeOperatorBonusPerBlockUpdatedIterator{contract: _StakingVesting.contract, event: "BridgeOperatorBonusPerBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorBonusPerBlockUpdated is a free log subscription operation binding the contract event 0x57a23b4b11f619fb9dea21a5a8115bb90103c1043eb3318d773558829d25f12c.
//
// Solidity: event BridgeOperatorBonusPerBlockUpdated(uint256 arg0)
func (_StakingVesting *StakingVestingFilterer) WatchBridgeOperatorBonusPerBlockUpdated(opts *bind.WatchOpts, sink chan<- *StakingVestingBridgeOperatorBonusPerBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingVesting.contract.WatchLogs(opts, "BridgeOperatorBonusPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVestingBridgeOperatorBonusPerBlockUpdated)
				if err := _StakingVesting.contract.UnpackLog(event, "BridgeOperatorBonusPerBlockUpdated", log); err != nil {
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

// ParseBridgeOperatorBonusPerBlockUpdated is a log parse operation binding the contract event 0x57a23b4b11f619fb9dea21a5a8115bb90103c1043eb3318d773558829d25f12c.
//
// Solidity: event BridgeOperatorBonusPerBlockUpdated(uint256 arg0)
func (_StakingVesting *StakingVestingFilterer) ParseBridgeOperatorBonusPerBlockUpdated(log types.Log) (*StakingVestingBridgeOperatorBonusPerBlockUpdated, error) {
	event := new(StakingVestingBridgeOperatorBonusPerBlockUpdated)
	if err := _StakingVesting.contract.UnpackLog(event, "BridgeOperatorBonusPerBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVestingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingVesting contract.
type StakingVestingInitializedIterator struct {
	Event *StakingVestingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingVestingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVestingInitialized)
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
		it.Event = new(StakingVestingInitialized)
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
func (it *StakingVestingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVestingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVestingInitialized represents a Initialized event raised by the StakingVesting contract.
type StakingVestingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingVesting *StakingVestingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingVestingInitializedIterator, error) {

	logs, sub, err := _StakingVesting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingVestingInitializedIterator{contract: _StakingVesting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingVesting *StakingVestingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingVestingInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingVesting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVestingInitialized)
				if err := _StakingVesting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingVesting *StakingVestingFilterer) ParseInitialized(log types.Log) (*StakingVestingInitialized, error) {
	event := new(StakingVestingInitialized)
	if err := _StakingVesting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVestingValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the StakingVesting contract.
type StakingVestingValidatorContractUpdatedIterator struct {
	Event *StakingVestingValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *StakingVestingValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVestingValidatorContractUpdated)
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
		it.Event = new(StakingVestingValidatorContractUpdated)
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
func (it *StakingVestingValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVestingValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVestingValidatorContractUpdated represents a ValidatorContractUpdated event raised by the StakingVesting contract.
type StakingVestingValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_StakingVesting *StakingVestingFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*StakingVestingValidatorContractUpdatedIterator, error) {

	logs, sub, err := _StakingVesting.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingVestingValidatorContractUpdatedIterator{contract: _StakingVesting.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_StakingVesting *StakingVestingFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *StakingVestingValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingVesting.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVestingValidatorContractUpdated)
				if err := _StakingVesting.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_StakingVesting *StakingVestingFilterer) ParseValidatorContractUpdated(log types.Log) (*StakingVestingValidatorContractUpdated, error) {
	event := new(StakingVestingValidatorContractUpdated)
	if err := _StakingVesting.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
