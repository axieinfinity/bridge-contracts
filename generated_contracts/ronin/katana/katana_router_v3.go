// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package katana

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

// RouterParameters is an auto generated low-level Go binding around an user-defined struct.
type RouterParameters struct {
	Permit2          common.Address
	Weth9            common.Address
	Governance       common.Address
	V2Factory        common.Address
	V3Factory        common.Address
	PairInitCodeHash [32]byte
	PoolInitCodeHash [32]byte
}

// KatanaRouterV3MetaData contains all meta data concerning the KatanaRouterV3 contract.
var KatanaRouterV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3Factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pairInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRouterParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FromAddressIsNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBips\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2UnauthorizedPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidSwap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3UnauthorizedSwap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"katanaV3SwapCallback\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KatanaRouterV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use KatanaRouterV3MetaData.ABI instead.
var KatanaRouterV3ABI = KatanaRouterV3MetaData.ABI

// KatanaRouterV3 is an auto generated Go binding around an Ethereum contract.
type KatanaRouterV3 struct {
	KatanaRouterV3Caller     // Read-only binding to the contract
	KatanaRouterV3Transactor // Write-only binding to the contract
	KatanaRouterV3Filterer   // Log filterer for contract events
}

// KatanaRouterV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type KatanaRouterV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaRouterV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type KatanaRouterV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaRouterV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KatanaRouterV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaRouterV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KatanaRouterV3Session struct {
	Contract     *KatanaRouterV3   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KatanaRouterV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KatanaRouterV3CallerSession struct {
	Contract *KatanaRouterV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KatanaRouterV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KatanaRouterV3TransactorSession struct {
	Contract     *KatanaRouterV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KatanaRouterV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type KatanaRouterV3Raw struct {
	Contract *KatanaRouterV3 // Generic contract binding to access the raw methods on
}

// KatanaRouterV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KatanaRouterV3CallerRaw struct {
	Contract *KatanaRouterV3Caller // Generic read-only contract binding to access the raw methods on
}

// KatanaRouterV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KatanaRouterV3TransactorRaw struct {
	Contract *KatanaRouterV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKatanaRouterV3 creates a new instance of KatanaRouterV3, bound to a specific deployed contract.
func NewKatanaRouterV3(address common.Address, backend bind.ContractBackend) (*KatanaRouterV3, error) {
	contract, err := bindKatanaRouterV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterV3{KatanaRouterV3Caller: KatanaRouterV3Caller{contract: contract}, KatanaRouterV3Transactor: KatanaRouterV3Transactor{contract: contract}, KatanaRouterV3Filterer: KatanaRouterV3Filterer{contract: contract}}, nil
}

// NewKatanaRouterV3Caller creates a new read-only instance of KatanaRouterV3, bound to a specific deployed contract.
func NewKatanaRouterV3Caller(address common.Address, caller bind.ContractCaller) (*KatanaRouterV3Caller, error) {
	contract, err := bindKatanaRouterV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterV3Caller{contract: contract}, nil
}

// NewKatanaRouterV3Transactor creates a new write-only instance of KatanaRouterV3, bound to a specific deployed contract.
func NewKatanaRouterV3Transactor(address common.Address, transactor bind.ContractTransactor) (*KatanaRouterV3Transactor, error) {
	contract, err := bindKatanaRouterV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterV3Transactor{contract: contract}, nil
}

// NewKatanaRouterV3Filterer creates a new log filterer instance of KatanaRouterV3, bound to a specific deployed contract.
func NewKatanaRouterV3Filterer(address common.Address, filterer bind.ContractFilterer) (*KatanaRouterV3Filterer, error) {
	contract, err := bindKatanaRouterV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterV3Filterer{contract: contract}, nil
}

// bindKatanaRouterV3 binds a generic wrapper to an already deployed contract.
func bindKatanaRouterV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KatanaRouterV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaRouterV3 *KatanaRouterV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaRouterV3.Contract.KatanaRouterV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaRouterV3 *KatanaRouterV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.KatanaRouterV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaRouterV3 *KatanaRouterV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.KatanaRouterV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaRouterV3 *KatanaRouterV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaRouterV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaRouterV3 *KatanaRouterV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaRouterV3 *KatanaRouterV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_KatanaRouterV3 *KatanaRouterV3Transactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _KatanaRouterV3.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_KatanaRouterV3 *KatanaRouterV3Session) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.Execute(&_KatanaRouterV3.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_KatanaRouterV3 *KatanaRouterV3TransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.Execute(&_KatanaRouterV3.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_KatanaRouterV3 *KatanaRouterV3Transactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouterV3.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_KatanaRouterV3 *KatanaRouterV3Session) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.Execute0(&_KatanaRouterV3.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_KatanaRouterV3 *KatanaRouterV3TransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.Execute0(&_KatanaRouterV3.TransactOpts, commands, inputs, deadline)
}

// KatanaV3SwapCallback is a paid mutator transaction binding the contract method 0x9c65a7d2.
//
// Solidity: function katanaV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_KatanaRouterV3 *KatanaRouterV3Transactor) KatanaV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaRouterV3.contract.Transact(opts, "katanaV3SwapCallback", amount0Delta, amount1Delta, data)
}

// KatanaV3SwapCallback is a paid mutator transaction binding the contract method 0x9c65a7d2.
//
// Solidity: function katanaV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_KatanaRouterV3 *KatanaRouterV3Session) KatanaV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.KatanaV3SwapCallback(&_KatanaRouterV3.TransactOpts, amount0Delta, amount1Delta, data)
}

// KatanaV3SwapCallback is a paid mutator transaction binding the contract method 0x9c65a7d2.
//
// Solidity: function katanaV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_KatanaRouterV3 *KatanaRouterV3TransactorSession) KatanaV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.KatanaV3SwapCallback(&_KatanaRouterV3.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KatanaRouterV3 *KatanaRouterV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaRouterV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KatanaRouterV3 *KatanaRouterV3Session) Receive() (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.Receive(&_KatanaRouterV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KatanaRouterV3 *KatanaRouterV3TransactorSession) Receive() (*types.Transaction, error) {
	return _KatanaRouterV3.Contract.Receive(&_KatanaRouterV3.TransactOpts)
}
