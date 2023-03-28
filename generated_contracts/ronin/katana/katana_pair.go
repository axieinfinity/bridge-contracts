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
)

// KatanaPairMetaData contains all meta data concerning the KatanaPair contract.
var KatanaPairMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RON_RELEASE_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WRON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addedMinters\",\"type\":\"address[]\"}],\"name\":\"addMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIKatanaFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removedMinters\",\"type\":\"address[]\"}],\"name\":\"removeMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KatanaPairABI is the input ABI used to generate the binding from.
// Deprecated: Use KatanaPairMetaData.ABI instead.
var KatanaPairABI = KatanaPairMetaData.ABI

// KatanaPair is an auto generated Go binding around an Ethereum contract.
type KatanaPair struct {
	KatanaPairCaller     // Read-only binding to the contract
	KatanaPairTransactor // Write-only binding to the contract
	KatanaPairFilterer   // Log filterer for contract events
}

// KatanaPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type KatanaPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KatanaPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KatanaPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KatanaPairSession struct {
	Contract     *KatanaPair       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KatanaPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KatanaPairCallerSession struct {
	Contract *KatanaPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// KatanaPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KatanaPairTransactorSession struct {
	Contract     *KatanaPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KatanaPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type KatanaPairRaw struct {
	Contract *KatanaPair // Generic contract binding to access the raw methods on
}

// KatanaPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KatanaPairCallerRaw struct {
	Contract *KatanaPairCaller // Generic read-only contract binding to access the raw methods on
}

// KatanaPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KatanaPairTransactorRaw struct {
	Contract *KatanaPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKatanaPair creates a new instance of KatanaPair, bound to a specific deployed contract.
func NewKatanaPair(address common.Address, backend bind.ContractBackend) (*KatanaPair, error) {
	contract, err := bindKatanaPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KatanaPair{KatanaPairCaller: KatanaPairCaller{contract: contract}, KatanaPairTransactor: KatanaPairTransactor{contract: contract}, KatanaPairFilterer: KatanaPairFilterer{contract: contract}}, nil
}

// NewKatanaPairCaller creates a new read-only instance of KatanaPair, bound to a specific deployed contract.
func NewKatanaPairCaller(address common.Address, caller bind.ContractCaller) (*KatanaPairCaller, error) {
	contract, err := bindKatanaPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaPairCaller{contract: contract}, nil
}

// NewKatanaPairTransactor creates a new write-only instance of KatanaPair, bound to a specific deployed contract.
func NewKatanaPairTransactor(address common.Address, transactor bind.ContractTransactor) (*KatanaPairTransactor, error) {
	contract, err := bindKatanaPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaPairTransactor{contract: contract}, nil
}

// NewKatanaPairFilterer creates a new log filterer instance of KatanaPair, bound to a specific deployed contract.
func NewKatanaPairFilterer(address common.Address, filterer bind.ContractFilterer) (*KatanaPairFilterer, error) {
	contract, err := bindKatanaPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KatanaPairFilterer{contract: contract}, nil
}

// bindKatanaPair binds a generic wrapper to an already deployed contract.
func bindKatanaPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KatanaPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaPair *KatanaPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaPair.Contract.KatanaPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaPair *KatanaPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaPair.Contract.KatanaPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaPair *KatanaPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaPair.Contract.KatanaPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaPair *KatanaPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaPair *KatanaPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaPair *KatanaPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaPair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_KatanaPair *KatanaPairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_KatanaPair *KatanaPairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _KatanaPair.Contract.DOMAINSEPARATOR(&_KatanaPair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_KatanaPair *KatanaPairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _KatanaPair.Contract.DOMAINSEPARATOR(&_KatanaPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_KatanaPair *KatanaPairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_KatanaPair *KatanaPairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _KatanaPair.Contract.MINIMUMLIQUIDITY(&_KatanaPair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _KatanaPair.Contract.MINIMUMLIQUIDITY(&_KatanaPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_KatanaPair *KatanaPairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_KatanaPair *KatanaPairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _KatanaPair.Contract.PERMITTYPEHASH(&_KatanaPair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_KatanaPair *KatanaPairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _KatanaPair.Contract.PERMITTYPEHASH(&_KatanaPair.CallOpts)
}

// RONRELEASEBLOCK is a free data retrieval call binding the contract method 0xd5466f34.
//
// Solidity: function RON_RELEASE_BLOCK() view returns(uint256)
func (_KatanaPair *KatanaPairCaller) RONRELEASEBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "RON_RELEASE_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RONRELEASEBLOCK is a free data retrieval call binding the contract method 0xd5466f34.
//
// Solidity: function RON_RELEASE_BLOCK() view returns(uint256)
func (_KatanaPair *KatanaPairSession) RONRELEASEBLOCK() (*big.Int, error) {
	return _KatanaPair.Contract.RONRELEASEBLOCK(&_KatanaPair.CallOpts)
}

// RONRELEASEBLOCK is a free data retrieval call binding the contract method 0xd5466f34.
//
// Solidity: function RON_RELEASE_BLOCK() view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) RONRELEASEBLOCK() (*big.Int, error) {
	return _KatanaPair.Contract.RONRELEASEBLOCK(&_KatanaPair.CallOpts)
}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_KatanaPair *KatanaPairCaller) WRON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "WRON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_KatanaPair *KatanaPairSession) WRON() (common.Address, error) {
	return _KatanaPair.Contract.WRON(&_KatanaPair.CallOpts)
}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_KatanaPair *KatanaPairCallerSession) WRON() (common.Address, error) {
	return _KatanaPair.Contract.WRON(&_KatanaPair.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KatanaPair *KatanaPairCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KatanaPair *KatanaPairSession) Admin() (common.Address, error) {
	return _KatanaPair.Contract.Admin(&_KatanaPair.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_KatanaPair *KatanaPairCallerSession) Admin() (common.Address, error) {
	return _KatanaPair.Contract.Admin(&_KatanaPair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 _value)
func (_KatanaPair *KatanaPairCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 _value)
func (_KatanaPair *KatanaPairSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _KatanaPair.Contract.Allowance(&_KatanaPair.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 _value)
func (_KatanaPair *KatanaPairCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _KatanaPair.Contract.Allowance(&_KatanaPair.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_KatanaPair *KatanaPairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_KatanaPair *KatanaPairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _KatanaPair.Contract.BalanceOf(&_KatanaPair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _KatanaPair.Contract.BalanceOf(&_KatanaPair.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KatanaPair *KatanaPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KatanaPair *KatanaPairSession) Decimals() (uint8, error) {
	return _KatanaPair.Contract.Decimals(&_KatanaPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KatanaPair *KatanaPairCallerSession) Decimals() (uint8, error) {
	return _KatanaPair.Contract.Decimals(&_KatanaPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaPair *KatanaPairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaPair *KatanaPairSession) Factory() (common.Address, error) {
	return _KatanaPair.Contract.Factory(&_KatanaPair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaPair *KatanaPairCallerSession) Factory() (common.Address, error) {
	return _KatanaPair.Contract.Factory(&_KatanaPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112, uint112, uint32)
func (_KatanaPair *KatanaPairCaller) GetReserves(opts *bind.CallOpts) (*big.Int, *big.Int, uint32, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return out0, out1, out2, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112, uint112, uint32)
func (_KatanaPair *KatanaPairSession) GetReserves() (*big.Int, *big.Int, uint32, error) {
	return _KatanaPair.Contract.GetReserves(&_KatanaPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112, uint112, uint32)
func (_KatanaPair *KatanaPairCallerSession) GetReserves() (*big.Int, *big.Int, uint32, error) {
	return _KatanaPair.Contract.GetReserves(&_KatanaPair.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_KatanaPair *KatanaPairCaller) IsMinter(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "isMinter", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_KatanaPair *KatanaPairSession) IsMinter(_addr common.Address) (bool, error) {
	return _KatanaPair.Contract.IsMinter(&_KatanaPair.CallOpts, _addr)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_KatanaPair *KatanaPairCallerSession) IsMinter(_addr common.Address) (bool, error) {
	return _KatanaPair.Contract.IsMinter(&_KatanaPair.CallOpts, _addr)
}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_KatanaPair *KatanaPairCaller) Minter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "minter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_KatanaPair *KatanaPairSession) Minter(arg0 common.Address) (bool, error) {
	return _KatanaPair.Contract.Minter(&_KatanaPair.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_KatanaPair *KatanaPairCallerSession) Minter(arg0 common.Address) (bool, error) {
	return _KatanaPair.Contract.Minter(&_KatanaPair.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_KatanaPair *KatanaPairCaller) Minters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_KatanaPair *KatanaPairSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _KatanaPair.Contract.Minters(&_KatanaPair.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_KatanaPair *KatanaPairCallerSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _KatanaPair.Contract.Minters(&_KatanaPair.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KatanaPair *KatanaPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KatanaPair *KatanaPairSession) Name() (string, error) {
	return _KatanaPair.Contract.Name(&_KatanaPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KatanaPair *KatanaPairCallerSession) Name() (string, error) {
	return _KatanaPair.Contract.Name(&_KatanaPair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_KatanaPair *KatanaPairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_KatanaPair *KatanaPairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _KatanaPair.Contract.Nonces(&_KatanaPair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _KatanaPair.Contract.Nonces(&_KatanaPair.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_KatanaPair *KatanaPairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_KatanaPair *KatanaPairSession) Price0CumulativeLast() (*big.Int, error) {
	return _KatanaPair.Contract.Price0CumulativeLast(&_KatanaPair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _KatanaPair.Contract.Price0CumulativeLast(&_KatanaPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_KatanaPair *KatanaPairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_KatanaPair *KatanaPairSession) Price1CumulativeLast() (*big.Int, error) {
	return _KatanaPair.Contract.Price1CumulativeLast(&_KatanaPair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _KatanaPair.Contract.Price1CumulativeLast(&_KatanaPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KatanaPair *KatanaPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KatanaPair *KatanaPairSession) Symbol() (string, error) {
	return _KatanaPair.Contract.Symbol(&_KatanaPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KatanaPair *KatanaPairCallerSession) Symbol() (string, error) {
	return _KatanaPair.Contract.Symbol(&_KatanaPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KatanaPair *KatanaPairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KatanaPair *KatanaPairSession) Token0() (common.Address, error) {
	return _KatanaPair.Contract.Token0(&_KatanaPair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KatanaPair *KatanaPairCallerSession) Token0() (common.Address, error) {
	return _KatanaPair.Contract.Token0(&_KatanaPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KatanaPair *KatanaPairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KatanaPair *KatanaPairSession) Token1() (common.Address, error) {
	return _KatanaPair.Contract.Token1(&_KatanaPair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KatanaPair *KatanaPairCallerSession) Token1() (common.Address, error) {
	return _KatanaPair.Contract.Token1(&_KatanaPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KatanaPair *KatanaPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KatanaPair *KatanaPairSession) TotalSupply() (*big.Int, error) {
	return _KatanaPair.Contract.TotalSupply(&_KatanaPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KatanaPair *KatanaPairCallerSession) TotalSupply() (*big.Int, error) {
	return _KatanaPair.Contract.TotalSupply(&_KatanaPair.CallOpts)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_KatanaPair *KatanaPairTransactor) AddMinters(opts *bind.TransactOpts, _addedMinters []common.Address) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "addMinters", _addedMinters)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_KatanaPair *KatanaPairSession) AddMinters(_addedMinters []common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.AddMinters(&_KatanaPair.TransactOpts, _addedMinters)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_KatanaPair *KatanaPairTransactorSession) AddMinters(_addedMinters []common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.AddMinters(&_KatanaPair.TransactOpts, _addedMinters)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Approve(&_KatanaPair.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Approve(&_KatanaPair.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Burn(&_KatanaPair.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Burn(&_KatanaPair.TransactOpts, _value)
}

// Burn0 is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _to) returns(uint256 _amount0, uint256 _amount1)
func (_KatanaPair *KatanaPairTransactor) Burn0(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "burn0", _to)
}

// Burn0 is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _to) returns(uint256 _amount0, uint256 _amount1)
func (_KatanaPair *KatanaPairSession) Burn0(_to common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.Burn0(&_KatanaPair.TransactOpts, _to)
}

// Burn0 is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _to) returns(uint256 _amount0, uint256 _amount1)
func (_KatanaPair *KatanaPairTransactorSession) Burn0(_to common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.Burn0(&_KatanaPair.TransactOpts, _to)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.BurnFrom(&_KatanaPair.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.BurnFrom(&_KatanaPair.TransactOpts, _from, _value)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_KatanaPair *KatanaPairTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_KatanaPair *KatanaPairSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.ChangeAdmin(&_KatanaPair.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_KatanaPair *KatanaPairTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.ChangeAdmin(&_KatanaPair.TransactOpts, _newAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _token0, address _token1, address _admin, string _name, string _symbol) returns()
func (_KatanaPair *KatanaPairTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address, _admin common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "initialize", _token0, _token1, _admin, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _token0, address _token1, address _admin, string _name, string _symbol) returns()
func (_KatanaPair *KatanaPairSession) Initialize(_token0 common.Address, _token1 common.Address, _admin common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KatanaPair.Contract.Initialize(&_KatanaPair.TransactOpts, _token0, _token1, _admin, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x83b43589.
//
// Solidity: function initialize(address _token0, address _token1, address _admin, string _name, string _symbol) returns()
func (_KatanaPair *KatanaPairTransactorSession) Initialize(_token0 common.Address, _token1 common.Address, _admin common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KatanaPair.Contract.Initialize(&_KatanaPair.TransactOpts, _token0, _token1, _admin, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Mint(&_KatanaPair.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Mint(&_KatanaPair.TransactOpts, _to, _value)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 _liquidity)
func (_KatanaPair *KatanaPairTransactor) Mint0(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "mint0", _to)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 _liquidity)
func (_KatanaPair *KatanaPairSession) Mint0(_to common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.Mint0(&_KatanaPair.TransactOpts, _to)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256 _liquidity)
func (_KatanaPair *KatanaPairTransactorSession) Mint0(_to common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.Mint0(&_KatanaPair.TransactOpts, _to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_KatanaPair *KatanaPairTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_KatanaPair *KatanaPairSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaPair.Contract.Permit(&_KatanaPair.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_KatanaPair *KatanaPairTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaPair.Contract.Permit(&_KatanaPair.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_KatanaPair *KatanaPairTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_KatanaPair *KatanaPairSession) RemoveAdmin() (*types.Transaction, error) {
	return _KatanaPair.Contract.RemoveAdmin(&_KatanaPair.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_KatanaPair *KatanaPairTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _KatanaPair.Contract.RemoveAdmin(&_KatanaPair.TransactOpts)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_KatanaPair *KatanaPairTransactor) RemoveMinters(opts *bind.TransactOpts, _removedMinters []common.Address) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "removeMinters", _removedMinters)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_KatanaPair *KatanaPairSession) RemoveMinters(_removedMinters []common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.RemoveMinters(&_KatanaPair.TransactOpts, _removedMinters)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_KatanaPair *KatanaPairTransactorSession) RemoveMinters(_removedMinters []common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.RemoveMinters(&_KatanaPair.TransactOpts, _removedMinters)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address _to) returns()
func (_KatanaPair *KatanaPairTransactor) Skim(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "skim", _to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address _to) returns()
func (_KatanaPair *KatanaPairSession) Skim(_to common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.Skim(&_KatanaPair.TransactOpts, _to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address _to) returns()
func (_KatanaPair *KatanaPairTransactorSession) Skim(_to common.Address) (*types.Transaction, error) {
	return _KatanaPair.Contract.Skim(&_KatanaPair.TransactOpts, _to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 _amount0Out, uint256 _amount1Out, address _to, bytes ) returns()
func (_KatanaPair *KatanaPairTransactor) Swap(opts *bind.TransactOpts, _amount0Out *big.Int, _amount1Out *big.Int, _to common.Address, arg3 []byte) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "swap", _amount0Out, _amount1Out, _to, arg3)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 _amount0Out, uint256 _amount1Out, address _to, bytes ) returns()
func (_KatanaPair *KatanaPairSession) Swap(_amount0Out *big.Int, _amount1Out *big.Int, _to common.Address, arg3 []byte) (*types.Transaction, error) {
	return _KatanaPair.Contract.Swap(&_KatanaPair.TransactOpts, _amount0Out, _amount1Out, _to, arg3)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 _amount0Out, uint256 _amount1Out, address _to, bytes ) returns()
func (_KatanaPair *KatanaPairTransactorSession) Swap(_amount0Out *big.Int, _amount1Out *big.Int, _to common.Address, arg3 []byte) (*types.Transaction, error) {
	return _KatanaPair.Contract.Swap(&_KatanaPair.TransactOpts, _amount0Out, _amount1Out, _to, arg3)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_KatanaPair *KatanaPairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_KatanaPair *KatanaPairSession) Sync() (*types.Transaction, error) {
	return _KatanaPair.Contract.Sync(&_KatanaPair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_KatanaPair *KatanaPairTransactorSession) Sync() (*types.Transaction, error) {
	return _KatanaPair.Contract.Sync(&_KatanaPair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Transfer(&_KatanaPair.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.Transfer(&_KatanaPair.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.TransferFrom(&_KatanaPair.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool _success)
func (_KatanaPair *KatanaPairTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KatanaPair.Contract.TransferFrom(&_KatanaPair.TransactOpts, _from, _to, _value)
}

// KatanaPairAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the KatanaPair contract.
type KatanaPairAdminChangedIterator struct {
	Event *KatanaPairAdminChanged // Event containing the contract specifics and raw log

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
func (it *KatanaPairAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairAdminChanged)
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
		it.Event = new(KatanaPairAdminChanged)
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
func (it *KatanaPairAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairAdminChanged represents a AdminChanged event raised by the KatanaPair contract.
type KatanaPairAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_KatanaPair *KatanaPairFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*KatanaPairAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairAdminChangedIterator{contract: _KatanaPair.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_KatanaPair *KatanaPairFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *KatanaPairAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairAdminChanged)
				if err := _KatanaPair.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_KatanaPair *KatanaPairFilterer) ParseAdminChanged(log types.Log) (*KatanaPairAdminChanged, error) {
	event := new(KatanaPairAdminChanged)
	if err := _KatanaPair.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the KatanaPair contract.
type KatanaPairAdminRemovedIterator struct {
	Event *KatanaPairAdminRemoved // Event containing the contract specifics and raw log

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
func (it *KatanaPairAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairAdminRemoved)
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
		it.Event = new(KatanaPairAdminRemoved)
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
func (it *KatanaPairAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairAdminRemoved represents a AdminRemoved event raised by the KatanaPair contract.
type KatanaPairAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_KatanaPair *KatanaPairFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*KatanaPairAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairAdminRemovedIterator{contract: _KatanaPair.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_KatanaPair *KatanaPairFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *KatanaPairAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairAdminRemoved)
				if err := _KatanaPair.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_KatanaPair *KatanaPairFilterer) ParseAdminRemoved(log types.Log) (*KatanaPairAdminRemoved, error) {
	event := new(KatanaPairAdminRemoved)
	if err := _KatanaPair.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KatanaPair contract.
type KatanaPairApprovalIterator struct {
	Event *KatanaPairApproval // Event containing the contract specifics and raw log

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
func (it *KatanaPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairApproval)
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
		it.Event = new(KatanaPairApproval)
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
func (it *KatanaPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairApproval represents a Approval event raised by the KatanaPair contract.
type KatanaPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_KatanaPair *KatanaPairFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*KatanaPairApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairApprovalIterator{contract: _KatanaPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_KatanaPair *KatanaPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KatanaPairApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairApproval)
				if err := _KatanaPair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KatanaPair *KatanaPairFilterer) ParseApproval(log types.Log) (*KatanaPairApproval, error) {
	event := new(KatanaPairApproval)
	if err := _KatanaPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the KatanaPair contract.
type KatanaPairBurnIterator struct {
	Event *KatanaPairBurn // Event containing the contract specifics and raw log

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
func (it *KatanaPairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairBurn)
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
		it.Event = new(KatanaPairBurn)
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
func (it *KatanaPairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairBurn represents a Burn event raised by the KatanaPair contract.
type KatanaPairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed _sender, uint256 _amount0, uint256 _amount1, address indexed _to)
func (_KatanaPair *KatanaPairFilterer) FilterBurn(opts *bind.FilterOpts, _sender []common.Address, _to []common.Address) (*KatanaPairBurnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "Burn", _senderRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairBurnIterator{contract: _KatanaPair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed _sender, uint256 _amount0, uint256 _amount1, address indexed _to)
func (_KatanaPair *KatanaPairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *KatanaPairBurn, _sender []common.Address, _to []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "Burn", _senderRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairBurn)
				if err := _KatanaPair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed _sender, uint256 _amount0, uint256 _amount1, address indexed _to)
func (_KatanaPair *KatanaPairFilterer) ParseBurn(log types.Log) (*KatanaPairBurn, error) {
	event := new(KatanaPairBurn)
	if err := _KatanaPair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the KatanaPair contract.
type KatanaPairMintIterator struct {
	Event *KatanaPairMint // Event containing the contract specifics and raw log

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
func (it *KatanaPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairMint)
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
		it.Event = new(KatanaPairMint)
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
func (it *KatanaPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairMint represents a Mint event raised by the KatanaPair contract.
type KatanaPairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed _sender, uint256 _amount0, uint256 _amount1)
func (_KatanaPair *KatanaPairFilterer) FilterMint(opts *bind.FilterOpts, _sender []common.Address) (*KatanaPairMintIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "Mint", _senderRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairMintIterator{contract: _KatanaPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed _sender, uint256 _amount0, uint256 _amount1)
func (_KatanaPair *KatanaPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *KatanaPairMint, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "Mint", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairMint)
				if err := _KatanaPair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed _sender, uint256 _amount0, uint256 _amount1)
func (_KatanaPair *KatanaPairFilterer) ParseMint(log types.Log) (*KatanaPairMint, error) {
	event := new(KatanaPairMint)
	if err := _KatanaPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the KatanaPair contract.
type KatanaPairMinterAddedIterator struct {
	Event *KatanaPairMinterAdded // Event containing the contract specifics and raw log

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
func (it *KatanaPairMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairMinterAdded)
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
		it.Event = new(KatanaPairMinterAdded)
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
func (it *KatanaPairMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairMinterAdded represents a MinterAdded event raised by the KatanaPair contract.
type KatanaPairMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_KatanaPair *KatanaPairFilterer) FilterMinterAdded(opts *bind.FilterOpts, _minter []common.Address) (*KatanaPairMinterAddedIterator, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "MinterAdded", _minterRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairMinterAddedIterator{contract: _KatanaPair.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_KatanaPair *KatanaPairFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *KatanaPairMinterAdded, _minter []common.Address) (event.Subscription, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "MinterAdded", _minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairMinterAdded)
				if err := _KatanaPair.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_KatanaPair *KatanaPairFilterer) ParseMinterAdded(log types.Log) (*KatanaPairMinterAdded, error) {
	event := new(KatanaPairMinterAdded)
	if err := _KatanaPair.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the KatanaPair contract.
type KatanaPairMinterRemovedIterator struct {
	Event *KatanaPairMinterRemoved // Event containing the contract specifics and raw log

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
func (it *KatanaPairMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairMinterRemoved)
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
		it.Event = new(KatanaPairMinterRemoved)
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
func (it *KatanaPairMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairMinterRemoved represents a MinterRemoved event raised by the KatanaPair contract.
type KatanaPairMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_KatanaPair *KatanaPairFilterer) FilterMinterRemoved(opts *bind.FilterOpts, _minter []common.Address) (*KatanaPairMinterRemovedIterator, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "MinterRemoved", _minterRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairMinterRemovedIterator{contract: _KatanaPair.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_KatanaPair *KatanaPairFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *KatanaPairMinterRemoved, _minter []common.Address) (event.Subscription, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "MinterRemoved", _minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairMinterRemoved)
				if err := _KatanaPair.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_KatanaPair *KatanaPairFilterer) ParseMinterRemoved(log types.Log) (*KatanaPairMinterRemoved, error) {
	event := new(KatanaPairMinterRemoved)
	if err := _KatanaPair.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the KatanaPair contract.
type KatanaPairSwapIterator struct {
	Event *KatanaPairSwap // Event containing the contract specifics and raw log

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
func (it *KatanaPairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairSwap)
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
		it.Event = new(KatanaPairSwap)
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
func (it *KatanaPairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairSwap represents a Swap event raised by the KatanaPair contract.
type KatanaPairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed _sender, uint256 _amount0In, uint256 _amount1In, uint256 _amount0Out, uint256 _amount1Out, address indexed _to)
func (_KatanaPair *KatanaPairFilterer) FilterSwap(opts *bind.FilterOpts, _sender []common.Address, _to []common.Address) (*KatanaPairSwapIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "Swap", _senderRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairSwapIterator{contract: _KatanaPair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed _sender, uint256 _amount0In, uint256 _amount1In, uint256 _amount0Out, uint256 _amount1Out, address indexed _to)
func (_KatanaPair *KatanaPairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *KatanaPairSwap, _sender []common.Address, _to []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "Swap", _senderRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairSwap)
				if err := _KatanaPair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed _sender, uint256 _amount0In, uint256 _amount1In, uint256 _amount0Out, uint256 _amount1Out, address indexed _to)
func (_KatanaPair *KatanaPairFilterer) ParseSwap(log types.Log) (*KatanaPairSwap, error) {
	event := new(KatanaPairSwap)
	if err := _KatanaPair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the KatanaPair contract.
type KatanaPairSyncIterator struct {
	Event *KatanaPairSync // Event containing the contract specifics and raw log

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
func (it *KatanaPairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairSync)
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
		it.Event = new(KatanaPairSync)
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
func (it *KatanaPairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairSync represents a Sync event raised by the KatanaPair contract.
type KatanaPairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 _reserve0, uint112 _reserve1)
func (_KatanaPair *KatanaPairFilterer) FilterSync(opts *bind.FilterOpts) (*KatanaPairSyncIterator, error) {

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &KatanaPairSyncIterator{contract: _KatanaPair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 _reserve0, uint112 _reserve1)
func (_KatanaPair *KatanaPairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *KatanaPairSync) (event.Subscription, error) {

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairSync)
				if err := _KatanaPair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 _reserve0, uint112 _reserve1)
func (_KatanaPair *KatanaPairFilterer) ParseSync(log types.Log) (*KatanaPairSync, error) {
	event := new(KatanaPairSync)
	if err := _KatanaPair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KatanaPair contract.
type KatanaPairTransferIterator struct {
	Event *KatanaPairTransfer // Event containing the contract specifics and raw log

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
func (it *KatanaPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairTransfer)
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
		it.Event = new(KatanaPairTransfer)
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
func (it *KatanaPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairTransfer represents a Transfer event raised by the KatanaPair contract.
type KatanaPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_KatanaPair *KatanaPairFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*KatanaPairTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KatanaPair.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairTransferIterator{contract: _KatanaPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_KatanaPair *KatanaPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KatanaPairTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _KatanaPair.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairTransfer)
				if err := _KatanaPair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KatanaPair *KatanaPairFilterer) ParseTransfer(log types.Log) (*KatanaPairTransfer, error) {
	event := new(KatanaPairTransfer)
	if err := _KatanaPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
