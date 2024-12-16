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

// KatanaPairV3MetaData contains all meta data concerning the KatanaPairV3 contract.
var KatanaPairV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocolNumeratorOld\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocolDenominatorOld\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocolNumeratorNew\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocolDenominatorNew\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal0X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGrowthGlobal1X128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"}],\"name\":\"increaseObservationCardinalityNext\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1_\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee_\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing_\",\"type\":\"int24\"}],\"name\":\"initializeImmutables\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationCardinalityNext\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocolNum\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feeProtocolDen\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"snapshotCumulativesInside\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"tickCumulativeInside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsInside\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"tickBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"tickCumulativeOutside\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"secondsOutside\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KatanaPairV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use KatanaPairV3MetaData.ABI instead.
var KatanaPairV3ABI = KatanaPairV3MetaData.ABI

// KatanaPairV3 is an auto generated Go binding around an Ethereum contract.
type KatanaPairV3 struct {
	KatanaPairV3Caller     // Read-only binding to the contract
	KatanaPairV3Transactor // Write-only binding to the contract
	KatanaPairV3Filterer   // Log filterer for contract events
}

// KatanaPairV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type KatanaPairV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaPairV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type KatanaPairV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaPairV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KatanaPairV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KatanaPairV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KatanaPairV3Session struct {
	Contract     *KatanaPairV3     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KatanaPairV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KatanaPairV3CallerSession struct {
	Contract *KatanaPairV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KatanaPairV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KatanaPairV3TransactorSession struct {
	Contract     *KatanaPairV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KatanaPairV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type KatanaPairV3Raw struct {
	Contract *KatanaPairV3 // Generic contract binding to access the raw methods on
}

// KatanaPairV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KatanaPairV3CallerRaw struct {
	Contract *KatanaPairV3Caller // Generic read-only contract binding to access the raw methods on
}

// KatanaPairV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KatanaPairV3TransactorRaw struct {
	Contract *KatanaPairV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKatanaPairV3 creates a new instance of KatanaPairV3, bound to a specific deployed contract.
func NewKatanaPairV3(address common.Address, backend bind.ContractBackend) (*KatanaPairV3, error) {
	contract, err := bindKatanaPairV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3{KatanaPairV3Caller: KatanaPairV3Caller{contract: contract}, KatanaPairV3Transactor: KatanaPairV3Transactor{contract: contract}, KatanaPairV3Filterer: KatanaPairV3Filterer{contract: contract}}, nil
}

// NewKatanaPairV3Caller creates a new read-only instance of KatanaPairV3, bound to a specific deployed contract.
func NewKatanaPairV3Caller(address common.Address, caller bind.ContractCaller) (*KatanaPairV3Caller, error) {
	contract, err := bindKatanaPairV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3Caller{contract: contract}, nil
}

// NewKatanaPairV3Transactor creates a new write-only instance of KatanaPairV3, bound to a specific deployed contract.
func NewKatanaPairV3Transactor(address common.Address, transactor bind.ContractTransactor) (*KatanaPairV3Transactor, error) {
	contract, err := bindKatanaPairV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3Transactor{contract: contract}, nil
}

// NewKatanaPairV3Filterer creates a new log filterer instance of KatanaPairV3, bound to a specific deployed contract.
func NewKatanaPairV3Filterer(address common.Address, filterer bind.ContractFilterer) (*KatanaPairV3Filterer, error) {
	contract, err := bindKatanaPairV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3Filterer{contract: contract}, nil
}

// bindKatanaPairV3 binds a generic wrapper to an already deployed contract.
func bindKatanaPairV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KatanaPairV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaPairV3 *KatanaPairV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaPairV3.Contract.KatanaPairV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaPairV3 *KatanaPairV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.KatanaPairV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaPairV3 *KatanaPairV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.KatanaPairV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KatanaPairV3 *KatanaPairV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KatanaPairV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KatanaPairV3 *KatanaPairV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KatanaPairV3 *KatanaPairV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Session) Factory() (common.Address, error) {
	return _KatanaPairV3.Contract.Factory(&_KatanaPairV3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Factory() (common.Address, error) {
	return _KatanaPairV3.Contract.Factory(&_KatanaPairV3.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_KatanaPairV3 *KatanaPairV3Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_KatanaPairV3 *KatanaPairV3Session) Fee() (*big.Int, error) {
	return _KatanaPairV3.Contract.Fee(&_KatanaPairV3.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Fee() (*big.Int, error) {
	return _KatanaPairV3.Contract.Fee(&_KatanaPairV3.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3Caller) FeeGrowthGlobal0X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "feeGrowthGlobal0X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3Session) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _KatanaPairV3.Contract.FeeGrowthGlobal0X128(&_KatanaPairV3.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3CallerSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _KatanaPairV3.Contract.FeeGrowthGlobal0X128(&_KatanaPairV3.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3Caller) FeeGrowthGlobal1X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "feeGrowthGlobal1X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3Session) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _KatanaPairV3.Contract.FeeGrowthGlobal1X128(&_KatanaPairV3.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3CallerSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _KatanaPairV3.Contract.FeeGrowthGlobal1X128(&_KatanaPairV3.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Session) Governance() (common.Address, error) {
	return _KatanaPairV3.Contract.Governance(&_KatanaPairV3.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Governance() (common.Address, error) {
	return _KatanaPairV3.Contract.Governance(&_KatanaPairV3.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_KatanaPairV3 *KatanaPairV3Caller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_KatanaPairV3 *KatanaPairV3Session) Liquidity() (*big.Int, error) {
	return _KatanaPairV3.Contract.Liquidity(&_KatanaPairV3.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Liquidity() (*big.Int, error) {
	return _KatanaPairV3.Contract.Liquidity(&_KatanaPairV3.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_KatanaPairV3 *KatanaPairV3Caller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_KatanaPairV3 *KatanaPairV3Session) MaxLiquidityPerTick() (*big.Int, error) {
	return _KatanaPairV3.Contract.MaxLiquidityPerTick(&_KatanaPairV3.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_KatanaPairV3 *KatanaPairV3CallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _KatanaPairV3.Contract.MaxLiquidityPerTick(&_KatanaPairV3.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_KatanaPairV3 *KatanaPairV3Caller) Observations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "observations", arg0)

	outstruct := new(struct {
		BlockTimestamp                    uint32
		TickCumulative                    *big.Int
		SecondsPerLiquidityCumulativeX128 *big.Int
		Initialized                       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_KatanaPairV3 *KatanaPairV3Session) Observations(arg0 *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _KatanaPairV3.Contract.Observations(&_KatanaPairV3.CallOpts, arg0)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Observations(arg0 *big.Int) (struct {
	BlockTimestamp                    uint32
	TickCumulative                    *big.Int
	SecondsPerLiquidityCumulativeX128 *big.Int
	Initialized                       bool
}, error) {
	return _KatanaPairV3.Contract.Observations(&_KatanaPairV3.CallOpts, arg0)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_KatanaPairV3 *KatanaPairV3Caller) Observe(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "observe", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                    []*big.Int
		SecondsPerLiquidityCumulativeX128s []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128s = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_KatanaPairV3 *KatanaPairV3Session) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _KatanaPairV3.Contract.Observe(&_KatanaPairV3.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _KatanaPairV3.Contract.Observe(&_KatanaPairV3.CallOpts, secondsAgos)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Caller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Session) PositionManager() (common.Address, error) {
	return _KatanaPairV3.Contract.PositionManager(&_KatanaPairV3.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_KatanaPairV3 *KatanaPairV3CallerSession) PositionManager() (common.Address, error) {
	return _KatanaPairV3.Contract.PositionManager(&_KatanaPairV3.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_KatanaPairV3 *KatanaPairV3Caller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_KatanaPairV3 *KatanaPairV3Session) Positions(arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _KatanaPairV3.Contract.Positions(&_KatanaPairV3.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Positions(arg0 [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _KatanaPairV3.Contract.Positions(&_KatanaPairV3.CallOpts, arg0)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocolNum, uint8 feeProtocolDen, bool unlocked)
func (_KatanaPairV3 *KatanaPairV3Caller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocolNum             uint8
	FeeProtocolDen             uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocolNum             uint8
		FeeProtocolDen             uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocolNum = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.FeeProtocolDen = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocolNum, uint8 feeProtocolDen, bool unlocked)
func (_KatanaPairV3 *KatanaPairV3Session) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocolNum             uint8
	FeeProtocolDen             uint8
	Unlocked                   bool
}, error) {
	return _KatanaPairV3.Contract.Slot0(&_KatanaPairV3.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocolNum, uint8 feeProtocolDen, bool unlocked)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocolNum             uint8
	FeeProtocolDen             uint8
	Unlocked                   bool
}, error) {
	return _KatanaPairV3.Contract.Slot0(&_KatanaPairV3.CallOpts)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_KatanaPairV3 *KatanaPairV3Caller) SnapshotCumulativesInside(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "snapshotCumulativesInside", tickLower, tickUpper)

	outstruct := new(struct {
		TickCumulativeInside          *big.Int
		SecondsPerLiquidityInsideX128 *big.Int
		SecondsInside                 uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulativeInside = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityInsideX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsInside = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_KatanaPairV3 *KatanaPairV3Session) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _KatanaPairV3.Contract.SnapshotCumulativesInside(&_KatanaPairV3.CallOpts, tickLower, tickUpper)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint32 secondsInside)
func (_KatanaPairV3 *KatanaPairV3CallerSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside          *big.Int
	SecondsPerLiquidityInsideX128 *big.Int
	SecondsInside                 uint32
}, error) {
	return _KatanaPairV3.Contract.SnapshotCumulativesInside(&_KatanaPairV3.CallOpts, tickLower, tickUpper)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 ) view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3Caller) TickBitmap(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "tickBitmap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 ) view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3Session) TickBitmap(arg0 int16) (*big.Int, error) {
	return _KatanaPairV3.Contract.TickBitmap(&_KatanaPairV3.CallOpts, arg0)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 ) view returns(uint256)
func (_KatanaPairV3 *KatanaPairV3CallerSession) TickBitmap(arg0 int16) (*big.Int, error) {
	return _KatanaPairV3.Contract.TickBitmap(&_KatanaPairV3.CallOpts, arg0)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_KatanaPairV3 *KatanaPairV3Caller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_KatanaPairV3 *KatanaPairV3Session) TickSpacing() (*big.Int, error) {
	return _KatanaPairV3.Contract.TickSpacing(&_KatanaPairV3.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_KatanaPairV3 *KatanaPairV3CallerSession) TickSpacing() (*big.Int, error) {
	return _KatanaPairV3.Contract.TickSpacing(&_KatanaPairV3.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_KatanaPairV3 *KatanaPairV3Caller) Ticks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "ticks", arg0)

	outstruct := new(struct {
		LiquidityGross                 *big.Int
		LiquidityNet                   *big.Int
		FeeGrowthOutside0X128          *big.Int
		FeeGrowthOutside1X128          *big.Int
		TickCumulativeOutside          *big.Int
		SecondsPerLiquidityOutsideX128 *big.Int
		SecondsOutside                 uint32
		Initialized                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityGross = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityNet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside0X128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside1X128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickCumulativeOutside = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityOutsideX128 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SecondsOutside = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_KatanaPairV3 *KatanaPairV3Session) Ticks(arg0 *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _KatanaPairV3.Contract.Ticks(&_KatanaPairV3.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityGross, int128 liquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Ticks(arg0 *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _KatanaPairV3.Contract.Ticks(&_KatanaPairV3.CallOpts, arg0)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Caller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Session) Token0() (common.Address, error) {
	return _KatanaPairV3.Contract.Token0(&_KatanaPairV3.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Token0() (common.Address, error) {
	return _KatanaPairV3.Contract.Token0(&_KatanaPairV3.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Caller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KatanaPairV3.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KatanaPairV3 *KatanaPairV3Session) Token1() (common.Address, error) {
	return _KatanaPairV3.Contract.Token1(&_KatanaPairV3.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KatanaPairV3 *KatanaPairV3CallerSession) Token1() (common.Address, error) {
	return _KatanaPairV3.Contract.Token1(&_KatanaPairV3.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Transactor) Burn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "burn", tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Session) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Burn(&_KatanaPairV3.TransactOpts, tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3TransactorSession) Burn(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Burn(&_KatanaPairV3.TransactOpts, tickLower, tickUpper, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Transactor) Collect(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "collect", recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Session) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Collect(&_KatanaPairV3.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3TransactorSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Collect(&_KatanaPairV3.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_KatanaPairV3 *KatanaPairV3Transactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_KatanaPairV3 *KatanaPairV3Session) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Flash(&_KatanaPairV3.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_KatanaPairV3 *KatanaPairV3TransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Flash(&_KatanaPairV3.TransactOpts, recipient, amount0, amount1, data)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_KatanaPairV3 *KatanaPairV3Transactor) IncreaseObservationCardinalityNext(opts *bind.TransactOpts, observationCardinalityNext uint16) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "increaseObservationCardinalityNext", observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_KatanaPairV3 *KatanaPairV3Session) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.IncreaseObservationCardinalityNext(&_KatanaPairV3.TransactOpts, observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_KatanaPairV3 *KatanaPairV3TransactorSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.IncreaseObservationCardinalityNext(&_KatanaPairV3.TransactOpts, observationCardinalityNext)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_KatanaPairV3 *KatanaPairV3Transactor) Initialize(opts *bind.TransactOpts, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "initialize", sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_KatanaPairV3 *KatanaPairV3Session) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Initialize(&_KatanaPairV3.TransactOpts, sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_KatanaPairV3 *KatanaPairV3TransactorSession) Initialize(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Initialize(&_KatanaPairV3.TransactOpts, sqrtPriceX96)
}

// InitializeImmutables is a paid mutator transaction binding the contract method 0x077bc178.
//
// Solidity: function initializeImmutables(address factory_, address token0_, address token1_, uint24 fee_, int24 tickSpacing_) returns()
func (_KatanaPairV3 *KatanaPairV3Transactor) InitializeImmutables(opts *bind.TransactOpts, factory_ common.Address, token0_ common.Address, token1_ common.Address, fee_ *big.Int, tickSpacing_ *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "initializeImmutables", factory_, token0_, token1_, fee_, tickSpacing_)
}

// InitializeImmutables is a paid mutator transaction binding the contract method 0x077bc178.
//
// Solidity: function initializeImmutables(address factory_, address token0_, address token1_, uint24 fee_, int24 tickSpacing_) returns()
func (_KatanaPairV3 *KatanaPairV3Session) InitializeImmutables(factory_ common.Address, token0_ common.Address, token1_ common.Address, fee_ *big.Int, tickSpacing_ *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.InitializeImmutables(&_KatanaPairV3.TransactOpts, factory_, token0_, token1_, fee_, tickSpacing_)
}

// InitializeImmutables is a paid mutator transaction binding the contract method 0x077bc178.
//
// Solidity: function initializeImmutables(address factory_, address token0_, address token1_, uint24 fee_, int24 tickSpacing_) returns()
func (_KatanaPairV3 *KatanaPairV3TransactorSession) InitializeImmutables(factory_ common.Address, token0_ common.Address, token1_ common.Address, fee_ *big.Int, tickSpacing_ *big.Int) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.InitializeImmutables(&_KatanaPairV3.TransactOpts, factory_, token0_, token1_, fee_, tickSpacing_)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Transactor) Mint(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "mint", recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Session) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Mint(&_KatanaPairV3.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3TransactorSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Mint(&_KatanaPairV3.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_KatanaPairV3 *KatanaPairV3Transactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_KatanaPairV3 *KatanaPairV3Session) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Swap(&_KatanaPairV3.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_KatanaPairV3 *KatanaPairV3TransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _KatanaPairV3.Contract.Swap(&_KatanaPairV3.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// KatanaPairV3BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the KatanaPairV3 contract.
type KatanaPairV3BurnIterator struct {
	Event *KatanaPairV3Burn // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3Burn)
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
		it.Event = new(KatanaPairV3Burn)
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
func (it *KatanaPairV3BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3Burn represents a Burn event raised by the KatanaPairV3 contract.
type KatanaPairV3Burn struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*KatanaPairV3BurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3BurnIterator{contract: _KatanaPairV3.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *KatanaPairV3Burn, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3Burn)
				if err := _KatanaPairV3.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseBurn(log types.Log) (*KatanaPairV3Burn, error) {
	event := new(KatanaPairV3Burn)
	if err := _KatanaPairV3.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3CollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the KatanaPairV3 contract.
type KatanaPairV3CollectIterator struct {
	Event *KatanaPairV3Collect // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3CollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3Collect)
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
		it.Event = new(KatanaPairV3Collect)
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
func (it *KatanaPairV3CollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3CollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3Collect represents a Collect event raised by the KatanaPairV3 contract.
type KatanaPairV3Collect struct {
	Owner     common.Address
	Recipient common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*KatanaPairV3CollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3CollectIterator{contract: _KatanaPairV3.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *KatanaPairV3Collect, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3Collect)
				if err := _KatanaPairV3.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseCollect(log types.Log) (*KatanaPairV3Collect, error) {
	event := new(KatanaPairV3Collect)
	if err := _KatanaPairV3.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3CollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the KatanaPairV3 contract.
type KatanaPairV3CollectProtocolIterator struct {
	Event *KatanaPairV3CollectProtocol // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3CollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3CollectProtocol)
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
		it.Event = new(KatanaPairV3CollectProtocol)
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
func (it *KatanaPairV3CollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3CollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3CollectProtocol represents a CollectProtocol event raised by the KatanaPairV3 contract.
type KatanaPairV3CollectProtocol struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterCollectProtocol(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*KatanaPairV3CollectProtocolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3CollectProtocolIterator{contract: _KatanaPairV3.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *KatanaPairV3CollectProtocol, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3CollectProtocol)
				if err := _KatanaPairV3.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
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

// ParseCollectProtocol is a log parse operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseCollectProtocol(log types.Log) (*KatanaPairV3CollectProtocol, error) {
	event := new(KatanaPairV3CollectProtocol)
	if err := _KatanaPairV3.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3FlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the KatanaPairV3 contract.
type KatanaPairV3FlashIterator struct {
	Event *KatanaPairV3Flash // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3FlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3Flash)
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
		it.Event = new(KatanaPairV3Flash)
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
func (it *KatanaPairV3FlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3FlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3Flash represents a Flash event raised by the KatanaPairV3 contract.
type KatanaPairV3Flash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*KatanaPairV3FlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3FlashIterator{contract: _KatanaPairV3.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *KatanaPairV3Flash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3Flash)
				if err := _KatanaPairV3.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseFlash(log types.Log) (*KatanaPairV3Flash, error) {
	event := new(KatanaPairV3Flash)
	if err := _KatanaPairV3.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3IncreaseObservationCardinalityNextIterator is returned from FilterIncreaseObservationCardinalityNext and is used to iterate over the raw logs and unpacked data for IncreaseObservationCardinalityNext events raised by the KatanaPairV3 contract.
type KatanaPairV3IncreaseObservationCardinalityNextIterator struct {
	Event *KatanaPairV3IncreaseObservationCardinalityNext // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3IncreaseObservationCardinalityNextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3IncreaseObservationCardinalityNext)
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
		it.Event = new(KatanaPairV3IncreaseObservationCardinalityNext)
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
func (it *KatanaPairV3IncreaseObservationCardinalityNextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3IncreaseObservationCardinalityNextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3IncreaseObservationCardinalityNext represents a IncreaseObservationCardinalityNext event raised by the KatanaPairV3 contract.
type KatanaPairV3IncreaseObservationCardinalityNext struct {
	ObservationCardinalityNextOld uint16
	ObservationCardinalityNextNew uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseObservationCardinalityNext is a free log retrieval operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterIncreaseObservationCardinalityNext(opts *bind.FilterOpts) (*KatanaPairV3IncreaseObservationCardinalityNextIterator, error) {

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3IncreaseObservationCardinalityNextIterator{contract: _KatanaPairV3.contract, event: "IncreaseObservationCardinalityNext", logs: logs, sub: sub}, nil
}

// WatchIncreaseObservationCardinalityNext is a free log subscription operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchIncreaseObservationCardinalityNext(opts *bind.WatchOpts, sink chan<- *KatanaPairV3IncreaseObservationCardinalityNext) (event.Subscription, error) {

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3IncreaseObservationCardinalityNext)
				if err := _KatanaPairV3.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
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

// ParseIncreaseObservationCardinalityNext is a log parse operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseIncreaseObservationCardinalityNext(log types.Log) (*KatanaPairV3IncreaseObservationCardinalityNext, error) {
	event := new(KatanaPairV3IncreaseObservationCardinalityNext)
	if err := _KatanaPairV3.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3InitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the KatanaPairV3 contract.
type KatanaPairV3InitializeIterator struct {
	Event *KatanaPairV3Initialize // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3InitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3Initialize)
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
		it.Event = new(KatanaPairV3Initialize)
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
func (it *KatanaPairV3InitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3InitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3Initialize represents a Initialize event raised by the KatanaPairV3 contract.
type KatanaPairV3Initialize struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterInitialize(opts *bind.FilterOpts) (*KatanaPairV3InitializeIterator, error) {

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3InitializeIterator{contract: _KatanaPairV3.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *KatanaPairV3Initialize) (event.Subscription, error) {

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3Initialize)
				if err := _KatanaPairV3.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseInitialize(log types.Log) (*KatanaPairV3Initialize, error) {
	event := new(KatanaPairV3Initialize)
	if err := _KatanaPairV3.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the KatanaPairV3 contract.
type KatanaPairV3MintIterator struct {
	Event *KatanaPairV3Mint // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3Mint)
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
		it.Event = new(KatanaPairV3Mint)
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
func (it *KatanaPairV3MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3Mint represents a Mint event raised by the KatanaPairV3 contract.
type KatanaPairV3Mint struct {
	Sender    common.Address
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*KatanaPairV3MintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3MintIterator{contract: _KatanaPairV3.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *KatanaPairV3Mint, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3Mint)
				if err := _KatanaPairV3.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseMint(log types.Log) (*KatanaPairV3Mint, error) {
	event := new(KatanaPairV3Mint)
	if err := _KatanaPairV3.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3SetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the KatanaPairV3 contract.
type KatanaPairV3SetFeeProtocolIterator struct {
	Event *KatanaPairV3SetFeeProtocol // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3SetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3SetFeeProtocol)
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
		it.Event = new(KatanaPairV3SetFeeProtocol)
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
func (it *KatanaPairV3SetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3SetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3SetFeeProtocol represents a SetFeeProtocol event raised by the KatanaPairV3 contract.
type KatanaPairV3SetFeeProtocol struct {
	FeeProtocolNumeratorOld   uint8
	FeeProtocolDenominatorOld uint8
	FeeProtocolNumeratorNew   uint8
	FeeProtocolDenominatorNew uint8
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocolNumeratorOld, uint8 feeProtocolDenominatorOld, uint8 feeProtocolNumeratorNew, uint8 feeProtocolDenominatorNew)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*KatanaPairV3SetFeeProtocolIterator, error) {

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3SetFeeProtocolIterator{contract: _KatanaPairV3.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocolNumeratorOld, uint8 feeProtocolDenominatorOld, uint8 feeProtocolNumeratorNew, uint8 feeProtocolDenominatorNew)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *KatanaPairV3SetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3SetFeeProtocol)
				if err := _KatanaPairV3.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
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

// ParseSetFeeProtocol is a log parse operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocolNumeratorOld, uint8 feeProtocolDenominatorOld, uint8 feeProtocolNumeratorNew, uint8 feeProtocolDenominatorNew)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseSetFeeProtocol(log types.Log) (*KatanaPairV3SetFeeProtocol, error) {
	event := new(KatanaPairV3SetFeeProtocol)
	if err := _KatanaPairV3.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KatanaPairV3SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the KatanaPairV3 contract.
type KatanaPairV3SwapIterator struct {
	Event *KatanaPairV3Swap // Event containing the contract specifics and raw log

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
func (it *KatanaPairV3SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KatanaPairV3Swap)
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
		it.Event = new(KatanaPairV3Swap)
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
func (it *KatanaPairV3SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KatanaPairV3SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KatanaPairV3Swap represents a Swap event raised by the KatanaPairV3 contract.
type KatanaPairV3Swap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_KatanaPairV3 *KatanaPairV3Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*KatanaPairV3SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _KatanaPairV3.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &KatanaPairV3SwapIterator{contract: _KatanaPairV3.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_KatanaPairV3 *KatanaPairV3Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *KatanaPairV3Swap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _KatanaPairV3.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KatanaPairV3Swap)
				if err := _KatanaPairV3.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_KatanaPairV3 *KatanaPairV3Filterer) ParseSwap(log types.Log) (*KatanaPairV3Swap, error) {
	event := new(KatanaPairV3Swap)
	if err := _KatanaPairV3.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
