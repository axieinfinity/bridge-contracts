// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wron

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

// WRONMetaData contains all meta data concerning the WRON contract.
var WRONMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"SpenderUnwhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"SpenderWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"unwhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WRONABI is the input ABI used to generate the binding from.
// Deprecated: Use WRONMetaData.ABI instead.
var WRONABI = WRONMetaData.ABI

// WRON is an auto generated Go binding around an Ethereum contract.
type WRON struct {
	WRONCaller     // Read-only binding to the contract
	WRONTransactor // Write-only binding to the contract
	WRONFilterer   // Log filterer for contract events
}

// WRONCaller is an auto generated read-only Go binding around an Ethereum contract.
type WRONCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRONTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WRONTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRONFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WRONFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRONSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WRONSession struct {
	Contract     *WRON             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WRONCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WRONCallerSession struct {
	Contract *WRONCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WRONTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WRONTransactorSession struct {
	Contract     *WRONTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WRONRaw is an auto generated low-level Go binding around an Ethereum contract.
type WRONRaw struct {
	Contract *WRON // Generic contract binding to access the raw methods on
}

// WRONCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WRONCallerRaw struct {
	Contract *WRONCaller // Generic read-only contract binding to access the raw methods on
}

// WRONTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WRONTransactorRaw struct {
	Contract *WRONTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWRON creates a new instance of WRON, bound to a specific deployed contract.
func NewWRON(address common.Address, backend bind.ContractBackend) (*WRON, error) {
	contract, err := bindWRON(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WRON{WRONCaller: WRONCaller{contract: contract}, WRONTransactor: WRONTransactor{contract: contract}, WRONFilterer: WRONFilterer{contract: contract}}, nil
}

// NewWRONCaller creates a new read-only instance of WRON, bound to a specific deployed contract.
func NewWRONCaller(address common.Address, caller bind.ContractCaller) (*WRONCaller, error) {
	contract, err := bindWRON(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WRONCaller{contract: contract}, nil
}

// NewWRONTransactor creates a new write-only instance of WRON, bound to a specific deployed contract.
func NewWRONTransactor(address common.Address, transactor bind.ContractTransactor) (*WRONTransactor, error) {
	contract, err := bindWRON(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WRONTransactor{contract: contract}, nil
}

// NewWRONFilterer creates a new log filterer instance of WRON, bound to a specific deployed contract.
func NewWRONFilterer(address common.Address, filterer bind.ContractFilterer) (*WRONFilterer, error) {
	contract, err := bindWRON(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WRONFilterer{contract: contract}, nil
}

// bindWRON binds a generic wrapper to an already deployed contract.
func bindWRON(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WRONABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRON *WRONRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WRON.Contract.WRONCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRON *WRONRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRON.Contract.WRONTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRON *WRONRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRON.Contract.WRONTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRON *WRONCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WRON.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRON *WRONTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRON.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRON *WRONTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRON.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_WRON *WRONCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_WRON *WRONSession) Admin() (common.Address, error) {
	return _WRON.Contract.Admin(&_WRON.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_WRON *WRONCallerSession) Admin() (common.Address, error) {
	return _WRON.Contract.Admin(&_WRON.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 _value)
func (_WRON *WRONCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 _value)
func (_WRON *WRONSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WRON.Contract.Allowance(&_WRON.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 _value)
func (_WRON *WRONCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WRON.Contract.Allowance(&_WRON.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WRON *WRONCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WRON *WRONSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WRON.Contract.BalanceOf(&_WRON.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WRON *WRONCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WRON.Contract.BalanceOf(&_WRON.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WRON *WRONCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WRON *WRONSession) Decimals() (uint8, error) {
	return _WRON.Contract.Decimals(&_WRON.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WRON *WRONCallerSession) Decimals() (uint8, error) {
	return _WRON.Contract.Decimals(&_WRON.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WRON *WRONCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WRON *WRONSession) Name() (string, error) {
	return _WRON.Contract.Name(&_WRON.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WRON *WRONCallerSession) Name() (string, error) {
	return _WRON.Contract.Name(&_WRON.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WRON *WRONCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WRON *WRONSession) Symbol() (string, error) {
	return _WRON.Contract.Symbol(&_WRON.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WRON *WRONCallerSession) Symbol() (string, error) {
	return _WRON.Contract.Symbol(&_WRON.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WRON *WRONCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WRON *WRONSession) TotalSupply() (*big.Int, error) {
	return _WRON.Contract.TotalSupply(&_WRON.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WRON *WRONCallerSession) TotalSupply() (*big.Int, error) {
	return _WRON.Contract.TotalSupply(&_WRON.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_WRON *WRONCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WRON.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_WRON *WRONSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _WRON.Contract.Whitelisted(&_WRON.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_WRON *WRONCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _WRON.Contract.Whitelisted(&_WRON.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool _success)
func (_WRON *WRONTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool _success)
func (_WRON *WRONSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.Approve(&_WRON.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool _success)
func (_WRON *WRONTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.Approve(&_WRON.TransactOpts, _spender, _value)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_WRON *WRONTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_WRON *WRONSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _WRON.Contract.ChangeAdmin(&_WRON.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_WRON *WRONTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _WRON.Contract.ChangeAdmin(&_WRON.TransactOpts, _newAdmin)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WRON *WRONTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WRON *WRONSession) Deposit() (*types.Transaction, error) {
	return _WRON.Contract.Deposit(&_WRON.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WRON *WRONTransactorSession) Deposit() (*types.Transaction, error) {
	return _WRON.Contract.Deposit(&_WRON.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_WRON *WRONTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_WRON *WRONSession) RemoveAdmin() (*types.Transaction, error) {
	return _WRON.Contract.RemoveAdmin(&_WRON.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_WRON *WRONTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _WRON.Contract.RemoveAdmin(&_WRON.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_WRON *WRONTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_WRON *WRONSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.Transfer(&_WRON.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_WRON *WRONTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.Transfer(&_WRON.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool _success)
func (_WRON *WRONTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool _success)
func (_WRON *WRONSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.TransferFrom(&_WRON.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool _success)
func (_WRON *WRONTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.TransferFrom(&_WRON.TransactOpts, _from, _to, _value)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_WRON *WRONTransactor) Unwhitelist(opts *bind.TransactOpts, _spender common.Address) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "unwhitelist", _spender)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_WRON *WRONSession) Unwhitelist(_spender common.Address) (*types.Transaction, error) {
	return _WRON.Contract.Unwhitelist(&_WRON.TransactOpts, _spender)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_WRON *WRONTransactorSession) Unwhitelist(_spender common.Address) (*types.Transaction, error) {
	return _WRON.Contract.Unwhitelist(&_WRON.TransactOpts, _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_WRON *WRONTransactor) Whitelist(opts *bind.TransactOpts, _spender common.Address) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "whitelist", _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_WRON *WRONSession) Whitelist(_spender common.Address) (*types.Transaction, error) {
	return _WRON.Contract.Whitelist(&_WRON.TransactOpts, _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_WRON *WRONTransactorSession) Whitelist(_spender common.Address) (*types.Transaction, error) {
	return _WRON.Contract.Whitelist(&_WRON.TransactOpts, _spender)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WRON *WRONTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WRON.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WRON *WRONSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.Withdraw(&_WRON.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WRON *WRONTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WRON.Contract.Withdraw(&_WRON.TransactOpts, wad)
}

// WRONAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WRON contract.
type WRONAdminChangedIterator struct {
	Event *WRONAdminChanged // Event containing the contract specifics and raw log

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
func (it *WRONAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONAdminChanged)
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
		it.Event = new(WRONAdminChanged)
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
func (it *WRONAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONAdminChanged represents a AdminChanged event raised by the WRON contract.
type WRONAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_WRON *WRONFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*WRONAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _WRON.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &WRONAdminChangedIterator{contract: _WRON.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_WRON *WRONFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WRONAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _WRON.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONAdminChanged)
				if err := _WRON.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_WRON *WRONFilterer) ParseAdminChanged(log types.Log) (*WRONAdminChanged, error) {
	event := new(WRONAdminChanged)
	if err := _WRON.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the WRON contract.
type WRONAdminRemovedIterator struct {
	Event *WRONAdminRemoved // Event containing the contract specifics and raw log

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
func (it *WRONAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONAdminRemoved)
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
		it.Event = new(WRONAdminRemoved)
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
func (it *WRONAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONAdminRemoved represents a AdminRemoved event raised by the WRON contract.
type WRONAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_WRON *WRONFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*WRONAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _WRON.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &WRONAdminRemovedIterator{contract: _WRON.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_WRON *WRONFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *WRONAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _WRON.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONAdminRemoved)
				if err := _WRON.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_WRON *WRONFilterer) ParseAdminRemoved(log types.Log) (*WRONAdminRemoved, error) {
	event := new(WRONAdminRemoved)
	if err := _WRON.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WRON contract.
type WRONApprovalIterator struct {
	Event *WRONApproval // Event containing the contract specifics and raw log

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
func (it *WRONApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONApproval)
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
		it.Event = new(WRONApproval)
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
func (it *WRONApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONApproval represents a Approval event raised by the WRON contract.
type WRONApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_WRON *WRONFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*WRONApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _WRON.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &WRONApprovalIterator{contract: _WRON.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_WRON *WRONFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WRONApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _WRON.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONApproval)
				if err := _WRON.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_WRON *WRONFilterer) ParseApproval(log types.Log) (*WRONApproval, error) {
	event := new(WRONApproval)
	if err := _WRON.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WRON contract.
type WRONDepositIterator struct {
	Event *WRONDeposit // Event containing the contract specifics and raw log

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
func (it *WRONDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONDeposit)
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
		it.Event = new(WRONDeposit)
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
func (it *WRONDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONDeposit represents a Deposit event raised by the WRON contract.
type WRONDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 value)
func (_WRON *WRONFilterer) FilterDeposit(opts *bind.FilterOpts) (*WRONDepositIterator, error) {

	logs, sub, err := _WRON.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &WRONDepositIterator{contract: _WRON.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 value)
func (_WRON *WRONFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WRONDeposit) (event.Subscription, error) {

	logs, sub, err := _WRON.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONDeposit)
				if err := _WRON.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 value)
func (_WRON *WRONFilterer) ParseDeposit(log types.Log) (*WRONDeposit, error) {
	event := new(WRONDeposit)
	if err := _WRON.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONSpenderUnwhitelistedIterator is returned from FilterSpenderUnwhitelisted and is used to iterate over the raw logs and unpacked data for SpenderUnwhitelisted events raised by the WRON contract.
type WRONSpenderUnwhitelistedIterator struct {
	Event *WRONSpenderUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *WRONSpenderUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONSpenderUnwhitelisted)
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
		it.Event = new(WRONSpenderUnwhitelisted)
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
func (it *WRONSpenderUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONSpenderUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONSpenderUnwhitelisted represents a SpenderUnwhitelisted event raised by the WRON contract.
type WRONSpenderUnwhitelisted struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderUnwhitelisted is a free log retrieval operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_WRON *WRONFilterer) FilterSpenderUnwhitelisted(opts *bind.FilterOpts, _spender []common.Address) (*WRONSpenderUnwhitelistedIterator, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _WRON.contract.FilterLogs(opts, "SpenderUnwhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return &WRONSpenderUnwhitelistedIterator{contract: _WRON.contract, event: "SpenderUnwhitelisted", logs: logs, sub: sub}, nil
}

// WatchSpenderUnwhitelisted is a free log subscription operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_WRON *WRONFilterer) WatchSpenderUnwhitelisted(opts *bind.WatchOpts, sink chan<- *WRONSpenderUnwhitelisted, _spender []common.Address) (event.Subscription, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _WRON.contract.WatchLogs(opts, "SpenderUnwhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONSpenderUnwhitelisted)
				if err := _WRON.contract.UnpackLog(event, "SpenderUnwhitelisted", log); err != nil {
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

// ParseSpenderUnwhitelisted is a log parse operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_WRON *WRONFilterer) ParseSpenderUnwhitelisted(log types.Log) (*WRONSpenderUnwhitelisted, error) {
	event := new(WRONSpenderUnwhitelisted)
	if err := _WRON.contract.UnpackLog(event, "SpenderUnwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONSpenderWhitelistedIterator is returned from FilterSpenderWhitelisted and is used to iterate over the raw logs and unpacked data for SpenderWhitelisted events raised by the WRON contract.
type WRONSpenderWhitelistedIterator struct {
	Event *WRONSpenderWhitelisted // Event containing the contract specifics and raw log

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
func (it *WRONSpenderWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONSpenderWhitelisted)
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
		it.Event = new(WRONSpenderWhitelisted)
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
func (it *WRONSpenderWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONSpenderWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONSpenderWhitelisted represents a SpenderWhitelisted event raised by the WRON contract.
type WRONSpenderWhitelisted struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderWhitelisted is a free log retrieval operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_WRON *WRONFilterer) FilterSpenderWhitelisted(opts *bind.FilterOpts, _spender []common.Address) (*WRONSpenderWhitelistedIterator, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _WRON.contract.FilterLogs(opts, "SpenderWhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return &WRONSpenderWhitelistedIterator{contract: _WRON.contract, event: "SpenderWhitelisted", logs: logs, sub: sub}, nil
}

// WatchSpenderWhitelisted is a free log subscription operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_WRON *WRONFilterer) WatchSpenderWhitelisted(opts *bind.WatchOpts, sink chan<- *WRONSpenderWhitelisted, _spender []common.Address) (event.Subscription, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _WRON.contract.WatchLogs(opts, "SpenderWhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONSpenderWhitelisted)
				if err := _WRON.contract.UnpackLog(event, "SpenderWhitelisted", log); err != nil {
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

// ParseSpenderWhitelisted is a log parse operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_WRON *WRONFilterer) ParseSpenderWhitelisted(log types.Log) (*WRONSpenderWhitelisted, error) {
	event := new(WRONSpenderWhitelisted)
	if err := _WRON.contract.UnpackLog(event, "SpenderWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WRON contract.
type WRONTransferIterator struct {
	Event *WRONTransfer // Event containing the contract specifics and raw log

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
func (it *WRONTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONTransfer)
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
		it.Event = new(WRONTransfer)
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
func (it *WRONTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONTransfer represents a Transfer event raised by the WRON contract.
type WRONTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_WRON *WRONFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*WRONTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _WRON.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &WRONTransferIterator{contract: _WRON.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_WRON *WRONFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WRONTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _WRON.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONTransfer)
				if err := _WRON.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_WRON *WRONFilterer) ParseTransfer(log types.Log) (*WRONTransfer, error) {
	event := new(WRONTransfer)
	if err := _WRON.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRONWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WRON contract.
type WRONWithdrawalIterator struct {
	Event *WRONWithdrawal // Event containing the contract specifics and raw log

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
func (it *WRONWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRONWithdrawal)
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
		it.Event = new(WRONWithdrawal)
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
func (it *WRONWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRONWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRONWithdrawal represents a Withdrawal event raised by the WRON contract.
type WRONWithdrawal struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address sender, uint256 value)
func (_WRON *WRONFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*WRONWithdrawalIterator, error) {

	logs, sub, err := _WRON.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &WRONWithdrawalIterator{contract: _WRON.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address sender, uint256 value)
func (_WRON *WRONFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WRONWithdrawal) (event.Subscription, error) {

	logs, sub, err := _WRON.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRONWithdrawal)
				if err := _WRON.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address sender, uint256 value)
func (_WRON *WRONFilterer) ParseWithdrawal(log types.Log) (*WRONWithdrawal, error) {
	event := new(WRONWithdrawal)
	if err := _WRON.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
