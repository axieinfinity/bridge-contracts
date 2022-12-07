// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maintenance

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	From             *big.Int
	To               *big.Int
	LastUpdatedBlock *big.Int
}

// MaintenanceMetaData contains all meta data concerning the Maintenance contract.
var MaintenanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minMaintenanceDurationInBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMaintenanceDurationInBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minOffsetToStartSchedule\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxOffsetToStartSchedule\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSchedules\",\"type\":\"uint256\"}],\"name\":\"MaintenanceConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMaintenance.Schedule\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"MaintenanceScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"checkMaintained\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toBlock\",\"type\":\"uint256\"}],\"name\":\"checkMaintainedInBlockRange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"checkManyMaintained\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_resList\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toBlock\",\"type\":\"uint256\"}],\"name\":\"checkManyMaintainedInBlockRange\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_resList\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"checkScheduled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"getSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIMaintenance.Schedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minMaintenanceDurationInBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMaintenanceDurationInBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOffsetToStartSchedule\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxOffsetToStartSchedule\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSchedules\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMaintenanceDurationInBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxOffsetToStartSchedule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSchedules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMaintenanceDurationInBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minOffsetToStartSchedule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endedAtBlock\",\"type\":\"uint256\"}],\"name\":\"schedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minMaintenanceDurationInBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMaintenanceDurationInBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOffsetToStartSchedule\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxOffsetToStartSchedule\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSchedules\",\"type\":\"uint256\"}],\"name\":\"setMaintenanceConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSchedules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MaintenanceABI is the input ABI used to generate the binding from.
// Deprecated: Use MaintenanceMetaData.ABI instead.
var MaintenanceABI = MaintenanceMetaData.ABI

// Maintenance is an auto generated Go binding around an Ethereum contract.
type Maintenance struct {
	MaintenanceCaller     // Read-only binding to the contract
	MaintenanceTransactor // Write-only binding to the contract
	MaintenanceFilterer   // Log filterer for contract events
}

// MaintenanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaintenanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaintenanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaintenanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaintenanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaintenanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaintenanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaintenanceSession struct {
	Contract     *Maintenance      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaintenanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaintenanceCallerSession struct {
	Contract *MaintenanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MaintenanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaintenanceTransactorSession struct {
	Contract     *MaintenanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MaintenanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaintenanceRaw struct {
	Contract *Maintenance // Generic contract binding to access the raw methods on
}

// MaintenanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaintenanceCallerRaw struct {
	Contract *MaintenanceCaller // Generic read-only contract binding to access the raw methods on
}

// MaintenanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaintenanceTransactorRaw struct {
	Contract *MaintenanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaintenance creates a new instance of Maintenance, bound to a specific deployed contract.
func NewMaintenance(address common.Address, backend bind.ContractBackend) (*Maintenance, error) {
	contract, err := bindMaintenance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Maintenance{MaintenanceCaller: MaintenanceCaller{contract: contract}, MaintenanceTransactor: MaintenanceTransactor{contract: contract}, MaintenanceFilterer: MaintenanceFilterer{contract: contract}}, nil
}

// NewMaintenanceCaller creates a new read-only instance of Maintenance, bound to a specific deployed contract.
func NewMaintenanceCaller(address common.Address, caller bind.ContractCaller) (*MaintenanceCaller, error) {
	contract, err := bindMaintenance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaintenanceCaller{contract: contract}, nil
}

// NewMaintenanceTransactor creates a new write-only instance of Maintenance, bound to a specific deployed contract.
func NewMaintenanceTransactor(address common.Address, transactor bind.ContractTransactor) (*MaintenanceTransactor, error) {
	contract, err := bindMaintenance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaintenanceTransactor{contract: contract}, nil
}

// NewMaintenanceFilterer creates a new log filterer instance of Maintenance, bound to a specific deployed contract.
func NewMaintenanceFilterer(address common.Address, filterer bind.ContractFilterer) (*MaintenanceFilterer, error) {
	contract, err := bindMaintenance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaintenanceFilterer{contract: contract}, nil
}

// bindMaintenance binds a generic wrapper to an already deployed contract.
func bindMaintenance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MaintenanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Maintenance *MaintenanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Maintenance.Contract.MaintenanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Maintenance *MaintenanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maintenance.Contract.MaintenanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Maintenance *MaintenanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Maintenance.Contract.MaintenanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Maintenance *MaintenanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Maintenance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Maintenance *MaintenanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Maintenance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Maintenance *MaintenanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Maintenance.Contract.contract.Transact(opts, method, params...)
}

// CheckMaintained is a free data retrieval call binding the contract method 0x0fbeb37f.
//
// Solidity: function checkMaintained(address _consensusAddr, uint256 _block) view returns(bool)
func (_Maintenance *MaintenanceCaller) CheckMaintained(opts *bind.CallOpts, _consensusAddr common.Address, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "checkMaintained", _consensusAddr, _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMaintained is a free data retrieval call binding the contract method 0x0fbeb37f.
//
// Solidity: function checkMaintained(address _consensusAddr, uint256 _block) view returns(bool)
func (_Maintenance *MaintenanceSession) CheckMaintained(_consensusAddr common.Address, _block *big.Int) (bool, error) {
	return _Maintenance.Contract.CheckMaintained(&_Maintenance.CallOpts, _consensusAddr, _block)
}

// CheckMaintained is a free data retrieval call binding the contract method 0x0fbeb37f.
//
// Solidity: function checkMaintained(address _consensusAddr, uint256 _block) view returns(bool)
func (_Maintenance *MaintenanceCallerSession) CheckMaintained(_consensusAddr common.Address, _block *big.Int) (bool, error) {
	return _Maintenance.Contract.CheckMaintained(&_Maintenance.CallOpts, _consensusAddr, _block)
}

// CheckMaintainedInBlockRange is a free data retrieval call binding the contract method 0x088e8de7.
//
// Solidity: function checkMaintainedInBlockRange(address _consensusAddr, uint256 _fromBlock, uint256 _toBlock) view returns(bool)
func (_Maintenance *MaintenanceCaller) CheckMaintainedInBlockRange(opts *bind.CallOpts, _consensusAddr common.Address, _fromBlock *big.Int, _toBlock *big.Int) (bool, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "checkMaintainedInBlockRange", _consensusAddr, _fromBlock, _toBlock)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMaintainedInBlockRange is a free data retrieval call binding the contract method 0x088e8de7.
//
// Solidity: function checkMaintainedInBlockRange(address _consensusAddr, uint256 _fromBlock, uint256 _toBlock) view returns(bool)
func (_Maintenance *MaintenanceSession) CheckMaintainedInBlockRange(_consensusAddr common.Address, _fromBlock *big.Int, _toBlock *big.Int) (bool, error) {
	return _Maintenance.Contract.CheckMaintainedInBlockRange(&_Maintenance.CallOpts, _consensusAddr, _fromBlock, _toBlock)
}

// CheckMaintainedInBlockRange is a free data retrieval call binding the contract method 0x088e8de7.
//
// Solidity: function checkMaintainedInBlockRange(address _consensusAddr, uint256 _fromBlock, uint256 _toBlock) view returns(bool)
func (_Maintenance *MaintenanceCallerSession) CheckMaintainedInBlockRange(_consensusAddr common.Address, _fromBlock *big.Int, _toBlock *big.Int) (bool, error) {
	return _Maintenance.Contract.CheckMaintainedInBlockRange(&_Maintenance.CallOpts, _consensusAddr, _fromBlock, _toBlock)
}

// CheckManyMaintained is a free data retrieval call binding the contract method 0xfdadda81.
//
// Solidity: function checkManyMaintained(address[] _addrList, uint256 _block) view returns(bool[] _resList)
func (_Maintenance *MaintenanceCaller) CheckManyMaintained(opts *bind.CallOpts, _addrList []common.Address, _block *big.Int) ([]bool, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "checkManyMaintained", _addrList, _block)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckManyMaintained is a free data retrieval call binding the contract method 0xfdadda81.
//
// Solidity: function checkManyMaintained(address[] _addrList, uint256 _block) view returns(bool[] _resList)
func (_Maintenance *MaintenanceSession) CheckManyMaintained(_addrList []common.Address, _block *big.Int) ([]bool, error) {
	return _Maintenance.Contract.CheckManyMaintained(&_Maintenance.CallOpts, _addrList, _block)
}

// CheckManyMaintained is a free data retrieval call binding the contract method 0xfdadda81.
//
// Solidity: function checkManyMaintained(address[] _addrList, uint256 _block) view returns(bool[] _resList)
func (_Maintenance *MaintenanceCallerSession) CheckManyMaintained(_addrList []common.Address, _block *big.Int) ([]bool, error) {
	return _Maintenance.Contract.CheckManyMaintained(&_Maintenance.CallOpts, _addrList, _block)
}

// CheckManyMaintainedInBlockRange is a free data retrieval call binding the contract method 0xba303755.
//
// Solidity: function checkManyMaintainedInBlockRange(address[] _addrList, uint256 _fromBlock, uint256 _toBlock) view returns(bool[] _resList)
func (_Maintenance *MaintenanceCaller) CheckManyMaintainedInBlockRange(opts *bind.CallOpts, _addrList []common.Address, _fromBlock *big.Int, _toBlock *big.Int) ([]bool, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "checkManyMaintainedInBlockRange", _addrList, _fromBlock, _toBlock)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckManyMaintainedInBlockRange is a free data retrieval call binding the contract method 0xba303755.
//
// Solidity: function checkManyMaintainedInBlockRange(address[] _addrList, uint256 _fromBlock, uint256 _toBlock) view returns(bool[] _resList)
func (_Maintenance *MaintenanceSession) CheckManyMaintainedInBlockRange(_addrList []common.Address, _fromBlock *big.Int, _toBlock *big.Int) ([]bool, error) {
	return _Maintenance.Contract.CheckManyMaintainedInBlockRange(&_Maintenance.CallOpts, _addrList, _fromBlock, _toBlock)
}

// CheckManyMaintainedInBlockRange is a free data retrieval call binding the contract method 0xba303755.
//
// Solidity: function checkManyMaintainedInBlockRange(address[] _addrList, uint256 _fromBlock, uint256 _toBlock) view returns(bool[] _resList)
func (_Maintenance *MaintenanceCallerSession) CheckManyMaintainedInBlockRange(_addrList []common.Address, _fromBlock *big.Int, _toBlock *big.Int) ([]bool, error) {
	return _Maintenance.Contract.CheckManyMaintainedInBlockRange(&_Maintenance.CallOpts, _addrList, _fromBlock, _toBlock)
}

// CheckScheduled is a free data retrieval call binding the contract method 0x2ddc08a2.
//
// Solidity: function checkScheduled(address _consensusAddr) view returns(bool)
func (_Maintenance *MaintenanceCaller) CheckScheduled(opts *bind.CallOpts, _consensusAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "checkScheduled", _consensusAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckScheduled is a free data retrieval call binding the contract method 0x2ddc08a2.
//
// Solidity: function checkScheduled(address _consensusAddr) view returns(bool)
func (_Maintenance *MaintenanceSession) CheckScheduled(_consensusAddr common.Address) (bool, error) {
	return _Maintenance.Contract.CheckScheduled(&_Maintenance.CallOpts, _consensusAddr)
}

// CheckScheduled is a free data retrieval call binding the contract method 0x2ddc08a2.
//
// Solidity: function checkScheduled(address _consensusAddr) view returns(bool)
func (_Maintenance *MaintenanceCallerSession) CheckScheduled(_consensusAddr common.Address) (bool, error) {
	return _Maintenance.Contract.CheckScheduled(&_Maintenance.CallOpts, _consensusAddr)
}

// GetSchedule is a free data retrieval call binding the contract method 0xd39fee34.
//
// Solidity: function getSchedule(address _consensusAddr) view returns((uint256,uint256,uint256))
func (_Maintenance *MaintenanceCaller) GetSchedule(opts *bind.CallOpts, _consensusAddr common.Address) (Struct0, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "getSchedule", _consensusAddr)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetSchedule is a free data retrieval call binding the contract method 0xd39fee34.
//
// Solidity: function getSchedule(address _consensusAddr) view returns((uint256,uint256,uint256))
func (_Maintenance *MaintenanceSession) GetSchedule(_consensusAddr common.Address) (Struct0, error) {
	return _Maintenance.Contract.GetSchedule(&_Maintenance.CallOpts, _consensusAddr)
}

// GetSchedule is a free data retrieval call binding the contract method 0xd39fee34.
//
// Solidity: function getSchedule(address _consensusAddr) view returns((uint256,uint256,uint256))
func (_Maintenance *MaintenanceCallerSession) GetSchedule(_consensusAddr common.Address) (Struct0, error) {
	return _Maintenance.Contract.GetSchedule(&_Maintenance.CallOpts, _consensusAddr)
}

// MaxMaintenanceDurationInBlock is a free data retrieval call binding the contract method 0xbfa89b9b.
//
// Solidity: function maxMaintenanceDurationInBlock() view returns(uint256)
func (_Maintenance *MaintenanceCaller) MaxMaintenanceDurationInBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "maxMaintenanceDurationInBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMaintenanceDurationInBlock is a free data retrieval call binding the contract method 0xbfa89b9b.
//
// Solidity: function maxMaintenanceDurationInBlock() view returns(uint256)
func (_Maintenance *MaintenanceSession) MaxMaintenanceDurationInBlock() (*big.Int, error) {
	return _Maintenance.Contract.MaxMaintenanceDurationInBlock(&_Maintenance.CallOpts)
}

// MaxMaintenanceDurationInBlock is a free data retrieval call binding the contract method 0xbfa89b9b.
//
// Solidity: function maxMaintenanceDurationInBlock() view returns(uint256)
func (_Maintenance *MaintenanceCallerSession) MaxMaintenanceDurationInBlock() (*big.Int, error) {
	return _Maintenance.Contract.MaxMaintenanceDurationInBlock(&_Maintenance.CallOpts)
}

// MaxOffsetToStartSchedule is a free data retrieval call binding the contract method 0x7a50802d.
//
// Solidity: function maxOffsetToStartSchedule() view returns(uint256)
func (_Maintenance *MaintenanceCaller) MaxOffsetToStartSchedule(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "maxOffsetToStartSchedule")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOffsetToStartSchedule is a free data retrieval call binding the contract method 0x7a50802d.
//
// Solidity: function maxOffsetToStartSchedule() view returns(uint256)
func (_Maintenance *MaintenanceSession) MaxOffsetToStartSchedule() (*big.Int, error) {
	return _Maintenance.Contract.MaxOffsetToStartSchedule(&_Maintenance.CallOpts)
}

// MaxOffsetToStartSchedule is a free data retrieval call binding the contract method 0x7a50802d.
//
// Solidity: function maxOffsetToStartSchedule() view returns(uint256)
func (_Maintenance *MaintenanceCallerSession) MaxOffsetToStartSchedule() (*big.Int, error) {
	return _Maintenance.Contract.MaxOffsetToStartSchedule(&_Maintenance.CallOpts)
}

// MaxSchedules is a free data retrieval call binding the contract method 0xc44cb233.
//
// Solidity: function maxSchedules() view returns(uint256)
func (_Maintenance *MaintenanceCaller) MaxSchedules(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "maxSchedules")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSchedules is a free data retrieval call binding the contract method 0xc44cb233.
//
// Solidity: function maxSchedules() view returns(uint256)
func (_Maintenance *MaintenanceSession) MaxSchedules() (*big.Int, error) {
	return _Maintenance.Contract.MaxSchedules(&_Maintenance.CallOpts)
}

// MaxSchedules is a free data retrieval call binding the contract method 0xc44cb233.
//
// Solidity: function maxSchedules() view returns(uint256)
func (_Maintenance *MaintenanceCallerSession) MaxSchedules() (*big.Int, error) {
	return _Maintenance.Contract.MaxSchedules(&_Maintenance.CallOpts)
}

// MinMaintenanceDurationInBlock is a free data retrieval call binding the contract method 0x09e34c38.
//
// Solidity: function minMaintenanceDurationInBlock() view returns(uint256)
func (_Maintenance *MaintenanceCaller) MinMaintenanceDurationInBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "minMaintenanceDurationInBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMaintenanceDurationInBlock is a free data retrieval call binding the contract method 0x09e34c38.
//
// Solidity: function minMaintenanceDurationInBlock() view returns(uint256)
func (_Maintenance *MaintenanceSession) MinMaintenanceDurationInBlock() (*big.Int, error) {
	return _Maintenance.Contract.MinMaintenanceDurationInBlock(&_Maintenance.CallOpts)
}

// MinMaintenanceDurationInBlock is a free data retrieval call binding the contract method 0x09e34c38.
//
// Solidity: function minMaintenanceDurationInBlock() view returns(uint256)
func (_Maintenance *MaintenanceCallerSession) MinMaintenanceDurationInBlock() (*big.Int, error) {
	return _Maintenance.Contract.MinMaintenanceDurationInBlock(&_Maintenance.CallOpts)
}

// MinOffsetToStartSchedule is a free data retrieval call binding the contract method 0xbc1710e9.
//
// Solidity: function minOffsetToStartSchedule() view returns(uint256)
func (_Maintenance *MaintenanceCaller) MinOffsetToStartSchedule(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "minOffsetToStartSchedule")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinOffsetToStartSchedule is a free data retrieval call binding the contract method 0xbc1710e9.
//
// Solidity: function minOffsetToStartSchedule() view returns(uint256)
func (_Maintenance *MaintenanceSession) MinOffsetToStartSchedule() (*big.Int, error) {
	return _Maintenance.Contract.MinOffsetToStartSchedule(&_Maintenance.CallOpts)
}

// MinOffsetToStartSchedule is a free data retrieval call binding the contract method 0xbc1710e9.
//
// Solidity: function minOffsetToStartSchedule() view returns(uint256)
func (_Maintenance *MaintenanceCallerSession) MinOffsetToStartSchedule() (*big.Int, error) {
	return _Maintenance.Contract.MinOffsetToStartSchedule(&_Maintenance.CallOpts)
}

// TotalSchedules is a free data retrieval call binding the contract method 0x965720af.
//
// Solidity: function totalSchedules() view returns(uint256 _count)
func (_Maintenance *MaintenanceCaller) TotalSchedules(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "totalSchedules")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSchedules is a free data retrieval call binding the contract method 0x965720af.
//
// Solidity: function totalSchedules() view returns(uint256 _count)
func (_Maintenance *MaintenanceSession) TotalSchedules() (*big.Int, error) {
	return _Maintenance.Contract.TotalSchedules(&_Maintenance.CallOpts)
}

// TotalSchedules is a free data retrieval call binding the contract method 0x965720af.
//
// Solidity: function totalSchedules() view returns(uint256 _count)
func (_Maintenance *MaintenanceCallerSession) TotalSchedules() (*big.Int, error) {
	return _Maintenance.Contract.TotalSchedules(&_Maintenance.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Maintenance *MaintenanceCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Maintenance.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Maintenance *MaintenanceSession) ValidatorContract() (common.Address, error) {
	return _Maintenance.Contract.ValidatorContract(&_Maintenance.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Maintenance *MaintenanceCallerSession) ValidatorContract() (common.Address, error) {
	return _Maintenance.Contract.ValidatorContract(&_Maintenance.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x45ff4c80.
//
// Solidity: function initialize(address __validatorContract, uint256 _minMaintenanceDurationInBlock, uint256 _maxMaintenanceDurationInBlock, uint256 _minOffsetToStartSchedule, uint256 _maxOffsetToStartSchedule, uint256 _maxSchedules) returns()
func (_Maintenance *MaintenanceTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, _minMaintenanceDurationInBlock *big.Int, _maxMaintenanceDurationInBlock *big.Int, _minOffsetToStartSchedule *big.Int, _maxOffsetToStartSchedule *big.Int, _maxSchedules *big.Int) (*types.Transaction, error) {
	return _Maintenance.contract.Transact(opts, "initialize", __validatorContract, _minMaintenanceDurationInBlock, _maxMaintenanceDurationInBlock, _minOffsetToStartSchedule, _maxOffsetToStartSchedule, _maxSchedules)
}

// Initialize is a paid mutator transaction binding the contract method 0x45ff4c80.
//
// Solidity: function initialize(address __validatorContract, uint256 _minMaintenanceDurationInBlock, uint256 _maxMaintenanceDurationInBlock, uint256 _minOffsetToStartSchedule, uint256 _maxOffsetToStartSchedule, uint256 _maxSchedules) returns()
func (_Maintenance *MaintenanceSession) Initialize(__validatorContract common.Address, _minMaintenanceDurationInBlock *big.Int, _maxMaintenanceDurationInBlock *big.Int, _minOffsetToStartSchedule *big.Int, _maxOffsetToStartSchedule *big.Int, _maxSchedules *big.Int) (*types.Transaction, error) {
	return _Maintenance.Contract.Initialize(&_Maintenance.TransactOpts, __validatorContract, _minMaintenanceDurationInBlock, _maxMaintenanceDurationInBlock, _minOffsetToStartSchedule, _maxOffsetToStartSchedule, _maxSchedules)
}

// Initialize is a paid mutator transaction binding the contract method 0x45ff4c80.
//
// Solidity: function initialize(address __validatorContract, uint256 _minMaintenanceDurationInBlock, uint256 _maxMaintenanceDurationInBlock, uint256 _minOffsetToStartSchedule, uint256 _maxOffsetToStartSchedule, uint256 _maxSchedules) returns()
func (_Maintenance *MaintenanceTransactorSession) Initialize(__validatorContract common.Address, _minMaintenanceDurationInBlock *big.Int, _maxMaintenanceDurationInBlock *big.Int, _minOffsetToStartSchedule *big.Int, _maxOffsetToStartSchedule *big.Int, _maxSchedules *big.Int) (*types.Transaction, error) {
	return _Maintenance.Contract.Initialize(&_Maintenance.TransactOpts, __validatorContract, _minMaintenanceDurationInBlock, _maxMaintenanceDurationInBlock, _minOffsetToStartSchedule, _maxOffsetToStartSchedule, _maxSchedules)
}

// Schedule is a paid mutator transaction binding the contract method 0x2d538c2c.
//
// Solidity: function schedule(address _consensusAddr, uint256 _startedAtBlock, uint256 _endedAtBlock) returns()
func (_Maintenance *MaintenanceTransactor) Schedule(opts *bind.TransactOpts, _consensusAddr common.Address, _startedAtBlock *big.Int, _endedAtBlock *big.Int) (*types.Transaction, error) {
	return _Maintenance.contract.Transact(opts, "schedule", _consensusAddr, _startedAtBlock, _endedAtBlock)
}

// Schedule is a paid mutator transaction binding the contract method 0x2d538c2c.
//
// Solidity: function schedule(address _consensusAddr, uint256 _startedAtBlock, uint256 _endedAtBlock) returns()
func (_Maintenance *MaintenanceSession) Schedule(_consensusAddr common.Address, _startedAtBlock *big.Int, _endedAtBlock *big.Int) (*types.Transaction, error) {
	return _Maintenance.Contract.Schedule(&_Maintenance.TransactOpts, _consensusAddr, _startedAtBlock, _endedAtBlock)
}

// Schedule is a paid mutator transaction binding the contract method 0x2d538c2c.
//
// Solidity: function schedule(address _consensusAddr, uint256 _startedAtBlock, uint256 _endedAtBlock) returns()
func (_Maintenance *MaintenanceTransactorSession) Schedule(_consensusAddr common.Address, _startedAtBlock *big.Int, _endedAtBlock *big.Int) (*types.Transaction, error) {
	return _Maintenance.Contract.Schedule(&_Maintenance.TransactOpts, _consensusAddr, _startedAtBlock, _endedAtBlock)
}

// SetMaintenanceConfig is a paid mutator transaction binding the contract method 0x0c4cea59.
//
// Solidity: function setMaintenanceConfig(uint256 _minMaintenanceDurationInBlock, uint256 _maxMaintenanceDurationInBlock, uint256 _minOffsetToStartSchedule, uint256 _maxOffsetToStartSchedule, uint256 _maxSchedules) returns()
func (_Maintenance *MaintenanceTransactor) SetMaintenanceConfig(opts *bind.TransactOpts, _minMaintenanceDurationInBlock *big.Int, _maxMaintenanceDurationInBlock *big.Int, _minOffsetToStartSchedule *big.Int, _maxOffsetToStartSchedule *big.Int, _maxSchedules *big.Int) (*types.Transaction, error) {
	return _Maintenance.contract.Transact(opts, "setMaintenanceConfig", _minMaintenanceDurationInBlock, _maxMaintenanceDurationInBlock, _minOffsetToStartSchedule, _maxOffsetToStartSchedule, _maxSchedules)
}

// SetMaintenanceConfig is a paid mutator transaction binding the contract method 0x0c4cea59.
//
// Solidity: function setMaintenanceConfig(uint256 _minMaintenanceDurationInBlock, uint256 _maxMaintenanceDurationInBlock, uint256 _minOffsetToStartSchedule, uint256 _maxOffsetToStartSchedule, uint256 _maxSchedules) returns()
func (_Maintenance *MaintenanceSession) SetMaintenanceConfig(_minMaintenanceDurationInBlock *big.Int, _maxMaintenanceDurationInBlock *big.Int, _minOffsetToStartSchedule *big.Int, _maxOffsetToStartSchedule *big.Int, _maxSchedules *big.Int) (*types.Transaction, error) {
	return _Maintenance.Contract.SetMaintenanceConfig(&_Maintenance.TransactOpts, _minMaintenanceDurationInBlock, _maxMaintenanceDurationInBlock, _minOffsetToStartSchedule, _maxOffsetToStartSchedule, _maxSchedules)
}

// SetMaintenanceConfig is a paid mutator transaction binding the contract method 0x0c4cea59.
//
// Solidity: function setMaintenanceConfig(uint256 _minMaintenanceDurationInBlock, uint256 _maxMaintenanceDurationInBlock, uint256 _minOffsetToStartSchedule, uint256 _maxOffsetToStartSchedule, uint256 _maxSchedules) returns()
func (_Maintenance *MaintenanceTransactorSession) SetMaintenanceConfig(_minMaintenanceDurationInBlock *big.Int, _maxMaintenanceDurationInBlock *big.Int, _minOffsetToStartSchedule *big.Int, _maxOffsetToStartSchedule *big.Int, _maxSchedules *big.Int) (*types.Transaction, error) {
	return _Maintenance.Contract.SetMaintenanceConfig(&_Maintenance.TransactOpts, _minMaintenanceDurationInBlock, _maxMaintenanceDurationInBlock, _minOffsetToStartSchedule, _maxOffsetToStartSchedule, _maxSchedules)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Maintenance *MaintenanceTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Maintenance.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Maintenance *MaintenanceSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Maintenance.Contract.SetValidatorContract(&_Maintenance.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Maintenance *MaintenanceTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Maintenance.Contract.SetValidatorContract(&_Maintenance.TransactOpts, _addr)
}

// MaintenanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Maintenance contract.
type MaintenanceInitializedIterator struct {
	Event *MaintenanceInitialized // Event containing the contract specifics and raw log

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
func (it *MaintenanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaintenanceInitialized)
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
		it.Event = new(MaintenanceInitialized)
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
func (it *MaintenanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaintenanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaintenanceInitialized represents a Initialized event raised by the Maintenance contract.
type MaintenanceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Maintenance *MaintenanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*MaintenanceInitializedIterator, error) {

	logs, sub, err := _Maintenance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MaintenanceInitializedIterator{contract: _Maintenance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Maintenance *MaintenanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MaintenanceInitialized) (event.Subscription, error) {

	logs, sub, err := _Maintenance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaintenanceInitialized)
				if err := _Maintenance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Maintenance *MaintenanceFilterer) ParseInitialized(log types.Log) (*MaintenanceInitialized, error) {
	event := new(MaintenanceInitialized)
	if err := _Maintenance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaintenanceMaintenanceConfigUpdatedIterator is returned from FilterMaintenanceConfigUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceConfigUpdated events raised by the Maintenance contract.
type MaintenanceMaintenanceConfigUpdatedIterator struct {
	Event *MaintenanceMaintenanceConfigUpdated // Event containing the contract specifics and raw log

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
func (it *MaintenanceMaintenanceConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaintenanceMaintenanceConfigUpdated)
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
		it.Event = new(MaintenanceMaintenanceConfigUpdated)
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
func (it *MaintenanceMaintenanceConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaintenanceMaintenanceConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaintenanceMaintenanceConfigUpdated represents a MaintenanceConfigUpdated event raised by the Maintenance contract.
type MaintenanceMaintenanceConfigUpdated struct {
	MinMaintenanceDurationInBlock *big.Int
	MaxMaintenanceDurationInBlock *big.Int
	MinOffsetToStartSchedule      *big.Int
	MaxOffsetToStartSchedule      *big.Int
	MaxSchedules                  *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceConfigUpdated is a free log retrieval operation binding the contract event 0xa95659ec997f7f4e9eb2cc9d46f0c1a257b876d8ebbb8938ff0bd4d5c7cce4fa.
//
// Solidity: event MaintenanceConfigUpdated(uint256 minMaintenanceDurationInBlock, uint256 maxMaintenanceDurationInBlock, uint256 minOffsetToStartSchedule, uint256 maxOffsetToStartSchedule, uint256 maxSchedules)
func (_Maintenance *MaintenanceFilterer) FilterMaintenanceConfigUpdated(opts *bind.FilterOpts) (*MaintenanceMaintenanceConfigUpdatedIterator, error) {

	logs, sub, err := _Maintenance.contract.FilterLogs(opts, "MaintenanceConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &MaintenanceMaintenanceConfigUpdatedIterator{contract: _Maintenance.contract, event: "MaintenanceConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceConfigUpdated is a free log subscription operation binding the contract event 0xa95659ec997f7f4e9eb2cc9d46f0c1a257b876d8ebbb8938ff0bd4d5c7cce4fa.
//
// Solidity: event MaintenanceConfigUpdated(uint256 minMaintenanceDurationInBlock, uint256 maxMaintenanceDurationInBlock, uint256 minOffsetToStartSchedule, uint256 maxOffsetToStartSchedule, uint256 maxSchedules)
func (_Maintenance *MaintenanceFilterer) WatchMaintenanceConfigUpdated(opts *bind.WatchOpts, sink chan<- *MaintenanceMaintenanceConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _Maintenance.contract.WatchLogs(opts, "MaintenanceConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaintenanceMaintenanceConfigUpdated)
				if err := _Maintenance.contract.UnpackLog(event, "MaintenanceConfigUpdated", log); err != nil {
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

// ParseMaintenanceConfigUpdated is a log parse operation binding the contract event 0xa95659ec997f7f4e9eb2cc9d46f0c1a257b876d8ebbb8938ff0bd4d5c7cce4fa.
//
// Solidity: event MaintenanceConfigUpdated(uint256 minMaintenanceDurationInBlock, uint256 maxMaintenanceDurationInBlock, uint256 minOffsetToStartSchedule, uint256 maxOffsetToStartSchedule, uint256 maxSchedules)
func (_Maintenance *MaintenanceFilterer) ParseMaintenanceConfigUpdated(log types.Log) (*MaintenanceMaintenanceConfigUpdated, error) {
	event := new(MaintenanceMaintenanceConfigUpdated)
	if err := _Maintenance.contract.UnpackLog(event, "MaintenanceConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaintenanceMaintenanceScheduledIterator is returned from FilterMaintenanceScheduled and is used to iterate over the raw logs and unpacked data for MaintenanceScheduled events raised by the Maintenance contract.
type MaintenanceMaintenanceScheduledIterator struct {
	Event *MaintenanceMaintenanceScheduled // Event containing the contract specifics and raw log

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
func (it *MaintenanceMaintenanceScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaintenanceMaintenanceScheduled)
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
		it.Event = new(MaintenanceMaintenanceScheduled)
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
func (it *MaintenanceMaintenanceScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaintenanceMaintenanceScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaintenanceMaintenanceScheduled represents a MaintenanceScheduled event raised by the Maintenance contract.
type MaintenanceMaintenanceScheduled struct {
	ConsensusAddr common.Address
	Arg1          Struct0
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceScheduled is a free log retrieval operation binding the contract event 0x3136fc9202ed1c3f8bc7ea3afb46704719833536485624cea83d7076f8785f43.
//
// Solidity: event MaintenanceScheduled(address indexed consensusAddr, (uint256,uint256,uint256) arg1)
func (_Maintenance *MaintenanceFilterer) FilterMaintenanceScheduled(opts *bind.FilterOpts, consensusAddr []common.Address) (*MaintenanceMaintenanceScheduledIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Maintenance.contract.FilterLogs(opts, "MaintenanceScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &MaintenanceMaintenanceScheduledIterator{contract: _Maintenance.contract, event: "MaintenanceScheduled", logs: logs, sub: sub}, nil
}

// WatchMaintenanceScheduled is a free log subscription operation binding the contract event 0x3136fc9202ed1c3f8bc7ea3afb46704719833536485624cea83d7076f8785f43.
//
// Solidity: event MaintenanceScheduled(address indexed consensusAddr, (uint256,uint256,uint256) arg1)
func (_Maintenance *MaintenanceFilterer) WatchMaintenanceScheduled(opts *bind.WatchOpts, sink chan<- *MaintenanceMaintenanceScheduled, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Maintenance.contract.WatchLogs(opts, "MaintenanceScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaintenanceMaintenanceScheduled)
				if err := _Maintenance.contract.UnpackLog(event, "MaintenanceScheduled", log); err != nil {
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

// ParseMaintenanceScheduled is a log parse operation binding the contract event 0x3136fc9202ed1c3f8bc7ea3afb46704719833536485624cea83d7076f8785f43.
//
// Solidity: event MaintenanceScheduled(address indexed consensusAddr, (uint256,uint256,uint256) arg1)
func (_Maintenance *MaintenanceFilterer) ParseMaintenanceScheduled(log types.Log) (*MaintenanceMaintenanceScheduled, error) {
	event := new(MaintenanceMaintenanceScheduled)
	if err := _Maintenance.contract.UnpackLog(event, "MaintenanceScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaintenanceValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Maintenance contract.
type MaintenanceValidatorContractUpdatedIterator struct {
	Event *MaintenanceValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *MaintenanceValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaintenanceValidatorContractUpdated)
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
		it.Event = new(MaintenanceValidatorContractUpdated)
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
func (it *MaintenanceValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaintenanceValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaintenanceValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Maintenance contract.
type MaintenanceValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Maintenance *MaintenanceFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*MaintenanceValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Maintenance.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &MaintenanceValidatorContractUpdatedIterator{contract: _Maintenance.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Maintenance *MaintenanceFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *MaintenanceValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Maintenance.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaintenanceValidatorContractUpdated)
				if err := _Maintenance.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_Maintenance *MaintenanceFilterer) ParseValidatorContractUpdated(log types.Log) (*MaintenanceValidatorContractUpdated, error) {
	event := new(MaintenanceValidatorContractUpdated)
	if err := _Maintenance.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
