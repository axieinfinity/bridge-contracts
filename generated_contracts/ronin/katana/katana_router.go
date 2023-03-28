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

// KatanaRouterMetaData contains all meta data concerning the KatanaRouter contract.
var KatanaRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WRON\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WRON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRONMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityRON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRON\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountB\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountB\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRONMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityRON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRON\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRONMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityRONSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountRON\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRONMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityRONWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRON\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountRONMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityRONWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountRON\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountB\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactRONForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactRONForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForRON\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForRONSupportingFeeOnTransferTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapRONForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactRON\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KatanaRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use KatanaRouterMetaData.ABI instead.
var KatanaRouterABI = KatanaRouterMetaData.ABI

// KatanaRouter is an auto generated Go binding around an Ethereum contract.
type KatanaRouter struct {
	KatanaRouterCaller     // Read-only binding to the contract
	KatanaRouterTransactor // Write-only binding to the contract
	KatanaRouterFilterer   // Log filterer for contract events
}

// KatanaRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type KatanaRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KatanaRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KatanaRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KatanaRouterSession struct {
	Contract     *KatanaRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KatanaRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KatanaRouterCallerSession struct {
	Contract *KatanaRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KatanaRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KatanaRouterTransactorSession struct {
	Contract     *KatanaRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KatanaRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type KatanaRouterRaw struct {
	Contract *KatanaRouter // Generic contract binding to access the raw methods on
}

// KatanaRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KatanaRouterCallerRaw struct {
	Contract *KatanaRouterCaller // Generic read-only contract binding to access the raw methods on
}

// KatanaRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KatanaRouterTransactorRaw struct {
	Contract *KatanaRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKatanaRouter creates a new instance of KatanaRouter, bound to a specific deployed contract.
func NewKatanaRouter(address common.Address, backend bind.ContractBackend) (*KatanaRouter, error) {
	contract, err := bindKatanaRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KatanaRouter{KatanaRouterCaller: KatanaRouterCaller{contract: contract}, KatanaRouterTransactor: KatanaRouterTransactor{contract: contract}, KatanaRouterFilterer: KatanaRouterFilterer{contract: contract}}, nil
}

// NewKatanaRouterCaller creates a new read-only instance of KatanaRouter, bound to a specific deployed contract.
func NewKatanaRouterCaller(address common.Address, caller bind.ContractCaller) (*KatanaRouterCaller, error) {
	contract, err := bindKatanaRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterCaller{contract: contract}, nil
}

// NewKatanaRouterTransactor creates a new write-only instance of KatanaRouter, bound to a specific deployed contract.
func NewKatanaRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*KatanaRouterTransactor, error) {
	contract, err := bindKatanaRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterTransactor{contract: contract}, nil
}

// NewKatanaRouterFilterer creates a new log filterer instance of KatanaRouter, bound to a specific deployed contract.
func NewKatanaRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*KatanaRouterFilterer, error) {
	contract, err := bindKatanaRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KatanaRouterFilterer{contract: contract}, nil
}

// bindKatanaRouter binds a generic wrapper to an already deployed contract.
func bindKatanaRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KatanaRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaRouter *KatanaRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaRouter.Contract.KatanaRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaRouter *KatanaRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaRouter.Contract.KatanaRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaRouter *KatanaRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaRouter.Contract.KatanaRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaRouter *KatanaRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaRouter *KatanaRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaRouter *KatanaRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaRouter.Contract.contract.Transact(opts, method, params...)
}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_KatanaRouter *KatanaRouterCaller) WRON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "WRON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_KatanaRouter *KatanaRouterSession) WRON() (common.Address, error) {
	return _KatanaRouter.Contract.WRON(&_KatanaRouter.CallOpts)
}

// WRON is a free data retrieval call binding the contract method 0x281388a8.
//
// Solidity: function WRON() view returns(address)
func (_KatanaRouter *KatanaRouterCallerSession) WRON() (common.Address, error) {
	return _KatanaRouter.Contract.WRON(&_KatanaRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaRouter *KatanaRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaRouter *KatanaRouterSession) Factory() (common.Address, error) {
	return _KatanaRouter.Contract.Factory(&_KatanaRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaRouter *KatanaRouterCallerSession) Factory() (common.Address, error) {
	return _KatanaRouter.Contract.Factory(&_KatanaRouter.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 _amountOut, uint256 _reserveIn, uint256 _reserveOut) pure returns(uint256 _amountIn)
func (_KatanaRouter *KatanaRouterCaller) GetAmountIn(opts *bind.CallOpts, _amountOut *big.Int, _reserveIn *big.Int, _reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "getAmountIn", _amountOut, _reserveIn, _reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 _amountOut, uint256 _reserveIn, uint256 _reserveOut) pure returns(uint256 _amountIn)
func (_KatanaRouter *KatanaRouterSession) GetAmountIn(_amountOut *big.Int, _reserveIn *big.Int, _reserveOut *big.Int) (*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountIn(&_KatanaRouter.CallOpts, _amountOut, _reserveIn, _reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 _amountOut, uint256 _reserveIn, uint256 _reserveOut) pure returns(uint256 _amountIn)
func (_KatanaRouter *KatanaRouterCallerSession) GetAmountIn(_amountOut *big.Int, _reserveIn *big.Int, _reserveOut *big.Int) (*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountIn(&_KatanaRouter.CallOpts, _amountOut, _reserveIn, _reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 _amountIn, uint256 _reserveIn, uint256 _reserveOut) pure returns(uint256 _amountOut)
func (_KatanaRouter *KatanaRouterCaller) GetAmountOut(opts *bind.CallOpts, _amountIn *big.Int, _reserveIn *big.Int, _reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "getAmountOut", _amountIn, _reserveIn, _reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 _amountIn, uint256 _reserveIn, uint256 _reserveOut) pure returns(uint256 _amountOut)
func (_KatanaRouter *KatanaRouterSession) GetAmountOut(_amountIn *big.Int, _reserveIn *big.Int, _reserveOut *big.Int) (*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountOut(&_KatanaRouter.CallOpts, _amountIn, _reserveIn, _reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 _amountIn, uint256 _reserveIn, uint256 _reserveOut) pure returns(uint256 _amountOut)
func (_KatanaRouter *KatanaRouterCallerSession) GetAmountOut(_amountIn *big.Int, _reserveIn *big.Int, _reserveOut *big.Int) (*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountOut(&_KatanaRouter.CallOpts, _amountIn, _reserveIn, _reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 _amountOut, address[] _path) view returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterCaller) GetAmountsIn(opts *bind.CallOpts, _amountOut *big.Int, _path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "getAmountsIn", _amountOut, _path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 _amountOut, address[] _path) view returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) GetAmountsIn(_amountOut *big.Int, _path []common.Address) ([]*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountsIn(&_KatanaRouter.CallOpts, _amountOut, _path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 _amountOut, address[] _path) view returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterCallerSession) GetAmountsIn(_amountOut *big.Int, _path []common.Address) ([]*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountsIn(&_KatanaRouter.CallOpts, _amountOut, _path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 _amountIn, address[] _path) view returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterCaller) GetAmountsOut(opts *bind.CallOpts, _amountIn *big.Int, _path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "getAmountsOut", _amountIn, _path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 _amountIn, address[] _path) view returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) GetAmountsOut(_amountIn *big.Int, _path []common.Address) ([]*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountsOut(&_KatanaRouter.CallOpts, _amountIn, _path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 _amountIn, address[] _path) view returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterCallerSession) GetAmountsOut(_amountIn *big.Int, _path []common.Address) ([]*big.Int, error) {
	return _KatanaRouter.Contract.GetAmountsOut(&_KatanaRouter.CallOpts, _amountIn, _path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 _amountA, uint256 _reserveA, uint256 _reserveB) pure returns(uint256 _amountB)
func (_KatanaRouter *KatanaRouterCaller) Quote(opts *bind.CallOpts, _amountA *big.Int, _reserveA *big.Int, _reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KatanaRouter.contract.Call(opts, &out, "quote", _amountA, _reserveA, _reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 _amountA, uint256 _reserveA, uint256 _reserveB) pure returns(uint256 _amountB)
func (_KatanaRouter *KatanaRouterSession) Quote(_amountA *big.Int, _reserveA *big.Int, _reserveB *big.Int) (*big.Int, error) {
	return _KatanaRouter.Contract.Quote(&_KatanaRouter.CallOpts, _amountA, _reserveA, _reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 _amountA, uint256 _reserveA, uint256 _reserveB) pure returns(uint256 _amountB)
func (_KatanaRouter *KatanaRouterCallerSession) Quote(_amountA *big.Int, _reserveA *big.Int, _reserveB *big.Int) (*big.Int, error) {
	return _KatanaRouter.Contract.Quote(&_KatanaRouter.CallOpts, _amountA, _reserveA, _reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _amountADesired, uint256 _amountBDesired, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline) returns(uint256 _amountA, uint256 _amountB, uint256 _liquidity)
func (_KatanaRouter *KatanaRouterTransactor) AddLiquidity(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _amountADesired *big.Int, _amountBDesired *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "addLiquidity", _tokenA, _tokenB, _amountADesired, _amountBDesired, _amountAMin, _amountBMin, _to, _deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _amountADesired, uint256 _amountBDesired, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline) returns(uint256 _amountA, uint256 _amountB, uint256 _liquidity)
func (_KatanaRouter *KatanaRouterSession) AddLiquidity(_tokenA common.Address, _tokenB common.Address, _amountADesired *big.Int, _amountBDesired *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.AddLiquidity(&_KatanaRouter.TransactOpts, _tokenA, _tokenB, _amountADesired, _amountBDesired, _amountAMin, _amountBMin, _to, _deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address _tokenA, address _tokenB, uint256 _amountADesired, uint256 _amountBDesired, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline) returns(uint256 _amountA, uint256 _amountB, uint256 _liquidity)
func (_KatanaRouter *KatanaRouterTransactorSession) AddLiquidity(_tokenA common.Address, _tokenB common.Address, _amountADesired *big.Int, _amountBDesired *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.AddLiquidity(&_KatanaRouter.TransactOpts, _tokenA, _tokenB, _amountADesired, _amountBDesired, _amountAMin, _amountBMin, _to, _deadline)
}

// AddLiquidityRON is a paid mutator transaction binding the contract method 0xf8f712c1.
//
// Solidity: function addLiquidityRON(address _token, uint256 _amountTokenDesired, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) payable returns(uint256 _amountToken, uint256 _amountRON, uint256 _liquidity)
func (_KatanaRouter *KatanaRouterTransactor) AddLiquidityRON(opts *bind.TransactOpts, _token common.Address, _amountTokenDesired *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "addLiquidityRON", _token, _amountTokenDesired, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// AddLiquidityRON is a paid mutator transaction binding the contract method 0xf8f712c1.
//
// Solidity: function addLiquidityRON(address _token, uint256 _amountTokenDesired, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) payable returns(uint256 _amountToken, uint256 _amountRON, uint256 _liquidity)
func (_KatanaRouter *KatanaRouterSession) AddLiquidityRON(_token common.Address, _amountTokenDesired *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.AddLiquidityRON(&_KatanaRouter.TransactOpts, _token, _amountTokenDesired, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// AddLiquidityRON is a paid mutator transaction binding the contract method 0xf8f712c1.
//
// Solidity: function addLiquidityRON(address _token, uint256 _amountTokenDesired, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) payable returns(uint256 _amountToken, uint256 _amountRON, uint256 _liquidity)
func (_KatanaRouter *KatanaRouterTransactorSession) AddLiquidityRON(_token common.Address, _amountTokenDesired *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.AddLiquidityRON(&_KatanaRouter.TransactOpts, _token, _amountTokenDesired, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _liquidity, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline) returns(uint256 _amountA, uint256 _amountB)
func (_KatanaRouter *KatanaRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _liquidity *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "removeLiquidity", _tokenA, _tokenB, _liquidity, _amountAMin, _amountBMin, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _liquidity, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline) returns(uint256 _amountA, uint256 _amountB)
func (_KatanaRouter *KatanaRouterSession) RemoveLiquidity(_tokenA common.Address, _tokenB common.Address, _liquidity *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidity(&_KatanaRouter.TransactOpts, _tokenA, _tokenB, _liquidity, _amountAMin, _amountBMin, _to, _deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address _tokenA, address _tokenB, uint256 _liquidity, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline) returns(uint256 _amountA, uint256 _amountB)
func (_KatanaRouter *KatanaRouterTransactorSession) RemoveLiquidity(_tokenA common.Address, _tokenB common.Address, _liquidity *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidity(&_KatanaRouter.TransactOpts, _tokenA, _tokenB, _liquidity, _amountAMin, _amountBMin, _to, _deadline)
}

// RemoveLiquidityRON is a paid mutator transaction binding the contract method 0xfb536a43.
//
// Solidity: function removeLiquidityRON(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) returns(uint256 _amountToken, uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactor) RemoveLiquidityRON(opts *bind.TransactOpts, _token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "removeLiquidityRON", _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidityRON is a paid mutator transaction binding the contract method 0xfb536a43.
//
// Solidity: function removeLiquidityRON(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) returns(uint256 _amountToken, uint256 _amountRON)
func (_KatanaRouter *KatanaRouterSession) RemoveLiquidityRON(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRON(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidityRON is a paid mutator transaction binding the contract method 0xfb536a43.
//
// Solidity: function removeLiquidityRON(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) returns(uint256 _amountToken, uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactorSession) RemoveLiquidityRON(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRON(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidityRONSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x77c33cd8.
//
// Solidity: function removeLiquidityRONSupportingFeeOnTransferTokens(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) returns(uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactor) RemoveLiquidityRONSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "removeLiquidityRONSupportingFeeOnTransferTokens", _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidityRONSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x77c33cd8.
//
// Solidity: function removeLiquidityRONSupportingFeeOnTransferTokens(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) returns(uint256 _amountRON)
func (_KatanaRouter *KatanaRouterSession) RemoveLiquidityRONSupportingFeeOnTransferTokens(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRONSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidityRONSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x77c33cd8.
//
// Solidity: function removeLiquidityRONSupportingFeeOnTransferTokens(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline) returns(uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactorSession) RemoveLiquidityRONSupportingFeeOnTransferTokens(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRONSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline)
}

// RemoveLiquidityRONWithPermit is a paid mutator transaction binding the contract method 0x76e58e8f.
//
// Solidity: function removeLiquidityRONWithPermit(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountToken, uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactor) RemoveLiquidityRONWithPermit(opts *bind.TransactOpts, _token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "removeLiquidityRONWithPermit", _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityRONWithPermit is a paid mutator transaction binding the contract method 0x76e58e8f.
//
// Solidity: function removeLiquidityRONWithPermit(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountToken, uint256 _amountRON)
func (_KatanaRouter *KatanaRouterSession) RemoveLiquidityRONWithPermit(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRONWithPermit(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityRONWithPermit is a paid mutator transaction binding the contract method 0x76e58e8f.
//
// Solidity: function removeLiquidityRONWithPermit(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountToken, uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactorSession) RemoveLiquidityRONWithPermit(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRONWithPermit(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb18446a0.
//
// Solidity: function removeLiquidityRONWithPermitSupportingFeeOnTransferTokens(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactor) RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "removeLiquidityRONWithPermitSupportingFeeOnTransferTokens", _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb18446a0.
//
// Solidity: function removeLiquidityRONWithPermitSupportingFeeOnTransferTokens(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountRON)
func (_KatanaRouter *KatanaRouterSession) RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb18446a0.
//
// Solidity: function removeLiquidityRONWithPermitSupportingFeeOnTransferTokens(address _token, uint256 _liquidity, uint256 _amountTokenMin, uint256 _amountRONMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountRON)
func (_KatanaRouter *KatanaRouterTransactorSession) RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens(_token common.Address, _liquidity *big.Int, _amountTokenMin *big.Int, _amountRONMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityRONWithPermitSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _token, _liquidity, _amountTokenMin, _amountRONMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address _tokenA, address _tokenB, uint256 _liquidity, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountA, uint256 _amountB)
func (_KatanaRouter *KatanaRouterTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address, _liquidity *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "removeLiquidityWithPermit", _tokenA, _tokenB, _liquidity, _amountAMin, _amountBMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address _tokenA, address _tokenB, uint256 _liquidity, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountA, uint256 _amountB)
func (_KatanaRouter *KatanaRouterSession) RemoveLiquidityWithPermit(_tokenA common.Address, _tokenB common.Address, _liquidity *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityWithPermit(&_KatanaRouter.TransactOpts, _tokenA, _tokenB, _liquidity, _amountAMin, _amountBMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address _tokenA, address _tokenB, uint256 _liquidity, uint256 _amountAMin, uint256 _amountBMin, address _to, uint256 _deadline, bool _approveMax, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256 _amountA, uint256 _amountB)
func (_KatanaRouter *KatanaRouterTransactorSession) RemoveLiquidityWithPermit(_tokenA common.Address, _tokenB common.Address, _liquidity *big.Int, _amountAMin *big.Int, _amountBMin *big.Int, _to common.Address, _deadline *big.Int, _approveMax bool, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.RemoveLiquidityWithPermit(&_KatanaRouter.TransactOpts, _tokenA, _tokenB, _liquidity, _amountAMin, _amountBMin, _to, _deadline, _approveMax, _v, _r, _s)
}

// SwapExactRONForTokens is a paid mutator transaction binding the contract method 0x7da5cd66.
//
// Solidity: function swapExactRONForTokens(uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) payable returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactor) SwapExactRONForTokens(opts *bind.TransactOpts, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapExactRONForTokens", _amountOutMin, _path, _to, _deadline)
}

// SwapExactRONForTokens is a paid mutator transaction binding the contract method 0x7da5cd66.
//
// Solidity: function swapExactRONForTokens(uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) payable returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) SwapExactRONForTokens(_amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactRONForTokens(&_KatanaRouter.TransactOpts, _amountOutMin, _path, _to, _deadline)
}

// SwapExactRONForTokens is a paid mutator transaction binding the contract method 0x7da5cd66.
//
// Solidity: function swapExactRONForTokens(uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) payable returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactorSession) SwapExactRONForTokens(_amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactRONForTokens(&_KatanaRouter.TransactOpts, _amountOutMin, _path, _to, _deadline)
}

// SwapExactRONForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9b010b20.
//
// Solidity: function swapExactRONForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) payable returns()
func (_KatanaRouter *KatanaRouterTransactor) SwapExactRONForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapExactRONForTokensSupportingFeeOnTransferTokens", _amountOutMin, _path, _to, _deadline)
}

// SwapExactRONForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9b010b20.
//
// Solidity: function swapExactRONForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) payable returns()
func (_KatanaRouter *KatanaRouterSession) SwapExactRONForTokensSupportingFeeOnTransferTokens(_amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactRONForTokensSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _amountOutMin, _path, _to, _deadline)
}

// SwapExactRONForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9b010b20.
//
// Solidity: function swapExactRONForTokensSupportingFeeOnTransferTokens(uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) payable returns()
func (_KatanaRouter *KatanaRouterTransactorSession) SwapExactRONForTokensSupportingFeeOnTransferTokens(_amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactRONForTokensSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForRON is a paid mutator transaction binding the contract method 0xf05b8c56.
//
// Solidity: function swapExactTokensForRON(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactor) SwapExactTokensForRON(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapExactTokensForRON", _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForRON is a paid mutator transaction binding the contract method 0xf05b8c56.
//
// Solidity: function swapExactTokensForRON(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) SwapExactTokensForRON(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForRON(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForRON is a paid mutator transaction binding the contract method 0xf05b8c56.
//
// Solidity: function swapExactTokensForRON(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactorSession) SwapExactTokensForRON(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForRON(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForRONSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x54cd9182.
//
// Solidity: function swapExactTokensForRONSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns()
func (_KatanaRouter *KatanaRouterTransactor) SwapExactTokensForRONSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapExactTokensForRONSupportingFeeOnTransferTokens", _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForRONSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x54cd9182.
//
// Solidity: function swapExactTokensForRONSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns()
func (_KatanaRouter *KatanaRouterSession) SwapExactTokensForRONSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForRONSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForRONSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x54cd9182.
//
// Solidity: function swapExactTokensForRONSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns()
func (_KatanaRouter *KatanaRouterTransactorSession) SwapExactTokensForRONSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForRONSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapExactTokensForTokens", _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) SwapExactTokensForTokens(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForTokens(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactorSession) SwapExactTokensForTokens(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForTokens(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns()
func (_KatanaRouter *KatanaRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns()
func (_KatanaRouter *KatanaRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 _amountIn, uint256 _amountOutMin, address[] _path, address _to, uint256 _deadline) returns()
func (_KatanaRouter *KatanaRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(_amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_KatanaRouter.TransactOpts, _amountIn, _amountOutMin, _path, _to, _deadline)
}

// SwapRONForExactTokens is a paid mutator transaction binding the contract method 0x54bc383f.
//
// Solidity: function swapRONForExactTokens(uint256 _amountOut, address[] _path, address _to, uint256 _deadline) payable returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactor) SwapRONForExactTokens(opts *bind.TransactOpts, _amountOut *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapRONForExactTokens", _amountOut, _path, _to, _deadline)
}

// SwapRONForExactTokens is a paid mutator transaction binding the contract method 0x54bc383f.
//
// Solidity: function swapRONForExactTokens(uint256 _amountOut, address[] _path, address _to, uint256 _deadline) payable returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) SwapRONForExactTokens(_amountOut *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapRONForExactTokens(&_KatanaRouter.TransactOpts, _amountOut, _path, _to, _deadline)
}

// SwapRONForExactTokens is a paid mutator transaction binding the contract method 0x54bc383f.
//
// Solidity: function swapRONForExactTokens(uint256 _amountOut, address[] _path, address _to, uint256 _deadline) payable returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactorSession) SwapRONForExactTokens(_amountOut *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapRONForExactTokens(&_KatanaRouter.TransactOpts, _amountOut, _path, _to, _deadline)
}

// SwapTokensForExactRON is a paid mutator transaction binding the contract method 0xa5089b8f.
//
// Solidity: function swapTokensForExactRON(uint256 _amountOut, uint256 _amountInMax, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactor) SwapTokensForExactRON(opts *bind.TransactOpts, _amountOut *big.Int, _amountInMax *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapTokensForExactRON", _amountOut, _amountInMax, _path, _to, _deadline)
}

// SwapTokensForExactRON is a paid mutator transaction binding the contract method 0xa5089b8f.
//
// Solidity: function swapTokensForExactRON(uint256 _amountOut, uint256 _amountInMax, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) SwapTokensForExactRON(_amountOut *big.Int, _amountInMax *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapTokensForExactRON(&_KatanaRouter.TransactOpts, _amountOut, _amountInMax, _path, _to, _deadline)
}

// SwapTokensForExactRON is a paid mutator transaction binding the contract method 0xa5089b8f.
//
// Solidity: function swapTokensForExactRON(uint256 _amountOut, uint256 _amountInMax, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactorSession) SwapTokensForExactRON(_amountOut *big.Int, _amountInMax *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapTokensForExactRON(&_KatanaRouter.TransactOpts, _amountOut, _amountInMax, _path, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, _amountOut *big.Int, _amountInMax *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.contract.Transact(opts, "swapTokensForExactTokens", _amountOut, _amountInMax, _path, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterSession) SwapTokensForExactTokens(_amountOut *big.Int, _amountInMax *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapTokensForExactTokens(&_KatanaRouter.TransactOpts, _amountOut, _amountInMax, _path, _to, _deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 _amountOut, uint256 _amountInMax, address[] _path, address _to, uint256 _deadline) returns(uint256[] _amounts)
func (_KatanaRouter *KatanaRouterTransactorSession) SwapTokensForExactTokens(_amountOut *big.Int, _amountInMax *big.Int, _path []common.Address, _to common.Address, _deadline *big.Int) (*types.Transaction, error) {
	return _KatanaRouter.Contract.SwapTokensForExactTokens(&_KatanaRouter.TransactOpts, _amountOut, _amountInMax, _path, _to, _deadline)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_KatanaRouter *KatanaRouterTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _KatanaRouter.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_KatanaRouter *KatanaRouterSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.Fallback(&_KatanaRouter.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_KatanaRouter *KatanaRouterTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _KatanaRouter.Contract.Fallback(&_KatanaRouter.TransactOpts, calldata)
}
