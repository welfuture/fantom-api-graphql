// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// SfcV400ConstantsMetaData contains all meta data concerning the SfcV400Constants contract.
var SfcV400ConstantsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueTooSmall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"averageUptimeEpochWindow\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burntFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasPriceBalancingCounterweight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAverageUptime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offlinePenaltyThresholdBlocksNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offlinePenaltyThresholdTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetGasPowerPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"v\",\"type\":\"uint32\"}],\"name\":\"updateAverageUptimeEpochWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateBurntFeeShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateGasPriceBalancingCounterweight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateMaxDelegatedRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"v\",\"type\":\"uint64\"}],\"name\":\"updateMinAverageUptime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateOfflinePenaltyThresholdBlocksNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateOfflinePenaltyThresholdTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateTargetGasPowerPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateTreasuryFeeShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateValidatorCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawalPeriodEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawalPeriodTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SfcV400ConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use SfcV400ConstantsMetaData.ABI instead.
var SfcV400ConstantsABI = SfcV400ConstantsMetaData.ABI

// SfcV400Constants is an auto generated Go binding around an Ethereum contract.
type SfcV400Constants struct {
	SfcV400ConstantsCaller     // Read-only binding to the contract
	SfcV400ConstantsTransactor // Write-only binding to the contract
	SfcV400ConstantsFilterer   // Log filterer for contract events
}

// SfcV400ConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SfcV400ConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV400ConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SfcV400ConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV400ConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SfcV400ConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV400ConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SfcV400ConstantsSession struct {
	Contract     *SfcV400Constants // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SfcV400ConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SfcV400ConstantsCallerSession struct {
	Contract *SfcV400ConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SfcV400ConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SfcV400ConstantsTransactorSession struct {
	Contract     *SfcV400ConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SfcV400ConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SfcV400ConstantsRaw struct {
	Contract *SfcV400Constants // Generic contract binding to access the raw methods on
}

// SfcV400ConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SfcV400ConstantsCallerRaw struct {
	Contract *SfcV400ConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// SfcV400ConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SfcV400ConstantsTransactorRaw struct {
	Contract *SfcV400ConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSfcV400Constants creates a new instance of SfcV400Constants, bound to a specific deployed contract.
func NewSfcV400Constants(address common.Address, backend bind.ContractBackend) (*SfcV400Constants, error) {
	contract, err := bindSfcV400Constants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SfcV400Constants{SfcV400ConstantsCaller: SfcV400ConstantsCaller{contract: contract}, SfcV400ConstantsTransactor: SfcV400ConstantsTransactor{contract: contract}, SfcV400ConstantsFilterer: SfcV400ConstantsFilterer{contract: contract}}, nil
}

// NewSfcV400ConstantsCaller creates a new read-only instance of SfcV400Constants, bound to a specific deployed contract.
func NewSfcV400ConstantsCaller(address common.Address, caller bind.ContractCaller) (*SfcV400ConstantsCaller, error) {
	contract, err := bindSfcV400Constants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SfcV400ConstantsCaller{contract: contract}, nil
}

// NewSfcV400ConstantsTransactor creates a new write-only instance of SfcV400Constants, bound to a specific deployed contract.
func NewSfcV400ConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*SfcV400ConstantsTransactor, error) {
	contract, err := bindSfcV400Constants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SfcV400ConstantsTransactor{contract: contract}, nil
}

// NewSfcV400ConstantsFilterer creates a new log filterer instance of SfcV400Constants, bound to a specific deployed contract.
func NewSfcV400ConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*SfcV400ConstantsFilterer, error) {
	contract, err := bindSfcV400Constants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SfcV400ConstantsFilterer{contract: contract}, nil
}

// bindSfcV400Constants binds a generic wrapper to an already deployed contract.
func bindSfcV400Constants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SfcV400ConstantsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcV400Constants *SfcV400ConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcV400Constants.Contract.SfcV400ConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcV400Constants *SfcV400ConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.SfcV400ConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcV400Constants *SfcV400ConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.SfcV400ConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcV400Constants *SfcV400ConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcV400Constants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcV400Constants *SfcV400ConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcV400Constants *SfcV400ConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.contract.Transact(opts, method, params...)
}

// AverageUptimeEpochWindow is a free data retrieval call binding the contract method 0x3fa22548.
//
// Solidity: function averageUptimeEpochWindow() view returns(uint32)
func (_SfcV400Constants *SfcV400ConstantsCaller) AverageUptimeEpochWindow(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "averageUptimeEpochWindow")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AverageUptimeEpochWindow is a free data retrieval call binding the contract method 0x3fa22548.
//
// Solidity: function averageUptimeEpochWindow() view returns(uint32)
func (_SfcV400Constants *SfcV400ConstantsSession) AverageUptimeEpochWindow() (uint32, error) {
	return _SfcV400Constants.Contract.AverageUptimeEpochWindow(&_SfcV400Constants.CallOpts)
}

// AverageUptimeEpochWindow is a free data retrieval call binding the contract method 0x3fa22548.
//
// Solidity: function averageUptimeEpochWindow() view returns(uint32)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) AverageUptimeEpochWindow() (uint32, error) {
	return _SfcV400Constants.Contract.AverageUptimeEpochWindow(&_SfcV400Constants.CallOpts)
}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) BaseRewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "baseRewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) BaseRewardPerSecond() (*big.Int, error) {
	return _SfcV400Constants.Contract.BaseRewardPerSecond(&_SfcV400Constants.CallOpts)
}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) BaseRewardPerSecond() (*big.Int, error) {
	return _SfcV400Constants.Contract.BaseRewardPerSecond(&_SfcV400Constants.CallOpts)
}

// BurntFeeShare is a free data retrieval call binding the contract method 0xc74dd621.
//
// Solidity: function burntFeeShare() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) BurntFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "burntFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurntFeeShare is a free data retrieval call binding the contract method 0xc74dd621.
//
// Solidity: function burntFeeShare() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) BurntFeeShare() (*big.Int, error) {
	return _SfcV400Constants.Contract.BurntFeeShare(&_SfcV400Constants.CallOpts)
}

// BurntFeeShare is a free data retrieval call binding the contract method 0xc74dd621.
//
// Solidity: function burntFeeShare() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) BurntFeeShare() (*big.Int, error) {
	return _SfcV400Constants.Contract.BurntFeeShare(&_SfcV400Constants.CallOpts)
}

// GasPriceBalancingCounterweight is a free data retrieval call binding the contract method 0x2c8c36a5.
//
// Solidity: function gasPriceBalancingCounterweight() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) GasPriceBalancingCounterweight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "gasPriceBalancingCounterweight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceBalancingCounterweight is a free data retrieval call binding the contract method 0x2c8c36a5.
//
// Solidity: function gasPriceBalancingCounterweight() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) GasPriceBalancingCounterweight() (*big.Int, error) {
	return _SfcV400Constants.Contract.GasPriceBalancingCounterweight(&_SfcV400Constants.CallOpts)
}

// GasPriceBalancingCounterweight is a free data retrieval call binding the contract method 0x2c8c36a5.
//
// Solidity: function gasPriceBalancingCounterweight() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) GasPriceBalancingCounterweight() (*big.Int, error) {
	return _SfcV400Constants.Contract.GasPriceBalancingCounterweight(&_SfcV400Constants.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "maxDelegatedRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) MaxDelegatedRatio() (*big.Int, error) {
	return _SfcV400Constants.Contract.MaxDelegatedRatio(&_SfcV400Constants.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _SfcV400Constants.Contract.MaxDelegatedRatio(&_SfcV400Constants.CallOpts)
}

// MinAverageUptime is a free data retrieval call binding the contract method 0x1c254337.
//
// Solidity: function minAverageUptime() view returns(uint64)
func (_SfcV400Constants *SfcV400ConstantsCaller) MinAverageUptime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "minAverageUptime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinAverageUptime is a free data retrieval call binding the contract method 0x1c254337.
//
// Solidity: function minAverageUptime() view returns(uint64)
func (_SfcV400Constants *SfcV400ConstantsSession) MinAverageUptime() (uint64, error) {
	return _SfcV400Constants.Contract.MinAverageUptime(&_SfcV400Constants.CallOpts)
}

// MinAverageUptime is a free data retrieval call binding the contract method 0x1c254337.
//
// Solidity: function minAverageUptime() view returns(uint64)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) MinAverageUptime() (uint64, error) {
	return _SfcV400Constants.Contract.MinAverageUptime(&_SfcV400Constants.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "minSelfStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) MinSelfStake() (*big.Int, error) {
	return _SfcV400Constants.Contract.MinSelfStake(&_SfcV400Constants.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) MinSelfStake() (*big.Int, error) {
	return _SfcV400Constants.Contract.MinSelfStake(&_SfcV400Constants.CallOpts)
}

// OfflinePenaltyThresholdBlocksNum is a free data retrieval call binding the contract method 0x5a68f01a.
//
// Solidity: function offlinePenaltyThresholdBlocksNum() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) OfflinePenaltyThresholdBlocksNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "offlinePenaltyThresholdBlocksNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfflinePenaltyThresholdBlocksNum is a free data retrieval call binding the contract method 0x5a68f01a.
//
// Solidity: function offlinePenaltyThresholdBlocksNum() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) OfflinePenaltyThresholdBlocksNum() (*big.Int, error) {
	return _SfcV400Constants.Contract.OfflinePenaltyThresholdBlocksNum(&_SfcV400Constants.CallOpts)
}

// OfflinePenaltyThresholdBlocksNum is a free data retrieval call binding the contract method 0x5a68f01a.
//
// Solidity: function offlinePenaltyThresholdBlocksNum() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) OfflinePenaltyThresholdBlocksNum() (*big.Int, error) {
	return _SfcV400Constants.Contract.OfflinePenaltyThresholdBlocksNum(&_SfcV400Constants.CallOpts)
}

// OfflinePenaltyThresholdTime is a free data retrieval call binding the contract method 0x00cc7f83.
//
// Solidity: function offlinePenaltyThresholdTime() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) OfflinePenaltyThresholdTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "offlinePenaltyThresholdTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfflinePenaltyThresholdTime is a free data retrieval call binding the contract method 0x00cc7f83.
//
// Solidity: function offlinePenaltyThresholdTime() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) OfflinePenaltyThresholdTime() (*big.Int, error) {
	return _SfcV400Constants.Contract.OfflinePenaltyThresholdTime(&_SfcV400Constants.CallOpts)
}

// OfflinePenaltyThresholdTime is a free data retrieval call binding the contract method 0x00cc7f83.
//
// Solidity: function offlinePenaltyThresholdTime() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) OfflinePenaltyThresholdTime() (*big.Int, error) {
	return _SfcV400Constants.Contract.OfflinePenaltyThresholdTime(&_SfcV400Constants.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV400Constants *SfcV400ConstantsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV400Constants *SfcV400ConstantsSession) Owner() (common.Address, error) {
	return _SfcV400Constants.Contract.Owner(&_SfcV400Constants.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) Owner() (common.Address, error) {
	return _SfcV400Constants.Contract.Owner(&_SfcV400Constants.CallOpts)
}

// TargetGasPowerPerSecond is a free data retrieval call binding the contract method 0x3a3ef66c.
//
// Solidity: function targetGasPowerPerSecond() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) TargetGasPowerPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "targetGasPowerPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetGasPowerPerSecond is a free data retrieval call binding the contract method 0x3a3ef66c.
//
// Solidity: function targetGasPowerPerSecond() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) TargetGasPowerPerSecond() (*big.Int, error) {
	return _SfcV400Constants.Contract.TargetGasPowerPerSecond(&_SfcV400Constants.CallOpts)
}

// TargetGasPowerPerSecond is a free data retrieval call binding the contract method 0x3a3ef66c.
//
// Solidity: function targetGasPowerPerSecond() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) TargetGasPowerPerSecond() (*big.Int, error) {
	return _SfcV400Constants.Contract.TargetGasPowerPerSecond(&_SfcV400Constants.CallOpts)
}

// TreasuryFeeShare is a free data retrieval call binding the contract method 0x94c3e914.
//
// Solidity: function treasuryFeeShare() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) TreasuryFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "treasuryFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFeeShare is a free data retrieval call binding the contract method 0x94c3e914.
//
// Solidity: function treasuryFeeShare() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) TreasuryFeeShare() (*big.Int, error) {
	return _SfcV400Constants.Contract.TreasuryFeeShare(&_SfcV400Constants.CallOpts)
}

// TreasuryFeeShare is a free data retrieval call binding the contract method 0x94c3e914.
//
// Solidity: function treasuryFeeShare() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) TreasuryFeeShare() (*big.Int, error) {
	return _SfcV400Constants.Contract.TreasuryFeeShare(&_SfcV400Constants.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "validatorCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) ValidatorCommission() (*big.Int, error) {
	return _SfcV400Constants.Contract.ValidatorCommission(&_SfcV400Constants.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) ValidatorCommission() (*big.Int, error) {
	return _SfcV400Constants.Contract.ValidatorCommission(&_SfcV400Constants.CallOpts)
}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) WithdrawalPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "withdrawalPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) WithdrawalPeriodEpochs() (*big.Int, error) {
	return _SfcV400Constants.Contract.WithdrawalPeriodEpochs(&_SfcV400Constants.CallOpts)
}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) WithdrawalPeriodEpochs() (*big.Int, error) {
	return _SfcV400Constants.Contract.WithdrawalPeriodEpochs(&_SfcV400Constants.CallOpts)
}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCaller) WithdrawalPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Constants.contract.Call(opts, &out, "withdrawalPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsSession) WithdrawalPeriodTime() (*big.Int, error) {
	return _SfcV400Constants.Contract.WithdrawalPeriodTime(&_SfcV400Constants.CallOpts)
}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() view returns(uint256)
func (_SfcV400Constants *SfcV400ConstantsCallerSession) WithdrawalPeriodTime() (*big.Int, error) {
	return _SfcV400Constants.Contract.WithdrawalPeriodTime(&_SfcV400Constants.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV400Constants *SfcV400ConstantsSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcV400Constants.Contract.RenounceOwnership(&_SfcV400Constants.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcV400Constants.Contract.RenounceOwnership(&_SfcV400Constants.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.TransferOwnership(&_SfcV400Constants.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.TransferOwnership(&_SfcV400Constants.TransactOpts, newOwner)
}

// UpdateAverageUptimeEpochWindow is a paid mutator transaction binding the contract method 0x256dc572.
//
// Solidity: function updateAverageUptimeEpochWindow(uint32 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateAverageUptimeEpochWindow(opts *bind.TransactOpts, v uint32) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateAverageUptimeEpochWindow", v)
}

// UpdateAverageUptimeEpochWindow is a paid mutator transaction binding the contract method 0x256dc572.
//
// Solidity: function updateAverageUptimeEpochWindow(uint32 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateAverageUptimeEpochWindow(v uint32) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateAverageUptimeEpochWindow(&_SfcV400Constants.TransactOpts, v)
}

// UpdateAverageUptimeEpochWindow is a paid mutator transaction binding the contract method 0x256dc572.
//
// Solidity: function updateAverageUptimeEpochWindow(uint32 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateAverageUptimeEpochWindow(v uint32) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateAverageUptimeEpochWindow(&_SfcV400Constants.TransactOpts, v)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateBaseRewardPerSecond(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateBaseRewardPerSecond", v)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateBaseRewardPerSecond(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateBaseRewardPerSecond(&_SfcV400Constants.TransactOpts, v)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateBaseRewardPerSecond(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateBaseRewardPerSecond(&_SfcV400Constants.TransactOpts, v)
}

// UpdateBurntFeeShare is a paid mutator transaction binding the contract method 0x2bb9fe8d.
//
// Solidity: function updateBurntFeeShare(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateBurntFeeShare(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateBurntFeeShare", v)
}

// UpdateBurntFeeShare is a paid mutator transaction binding the contract method 0x2bb9fe8d.
//
// Solidity: function updateBurntFeeShare(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateBurntFeeShare(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateBurntFeeShare(&_SfcV400Constants.TransactOpts, v)
}

// UpdateBurntFeeShare is a paid mutator transaction binding the contract method 0x2bb9fe8d.
//
// Solidity: function updateBurntFeeShare(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateBurntFeeShare(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateBurntFeeShare(&_SfcV400Constants.TransactOpts, v)
}

// UpdateGasPriceBalancingCounterweight is a paid mutator transaction binding the contract method 0xd3f48dbe.
//
// Solidity: function updateGasPriceBalancingCounterweight(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateGasPriceBalancingCounterweight(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateGasPriceBalancingCounterweight", v)
}

// UpdateGasPriceBalancingCounterweight is a paid mutator transaction binding the contract method 0xd3f48dbe.
//
// Solidity: function updateGasPriceBalancingCounterweight(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateGasPriceBalancingCounterweight(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateGasPriceBalancingCounterweight(&_SfcV400Constants.TransactOpts, v)
}

// UpdateGasPriceBalancingCounterweight is a paid mutator transaction binding the contract method 0xd3f48dbe.
//
// Solidity: function updateGasPriceBalancingCounterweight(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateGasPriceBalancingCounterweight(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateGasPriceBalancingCounterweight(&_SfcV400Constants.TransactOpts, v)
}

// UpdateMaxDelegatedRatio is a paid mutator transaction binding the contract method 0x81ffcdf1.
//
// Solidity: function updateMaxDelegatedRatio(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateMaxDelegatedRatio(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateMaxDelegatedRatio", v)
}

// UpdateMaxDelegatedRatio is a paid mutator transaction binding the contract method 0x81ffcdf1.
//
// Solidity: function updateMaxDelegatedRatio(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateMaxDelegatedRatio(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateMaxDelegatedRatio(&_SfcV400Constants.TransactOpts, v)
}

// UpdateMaxDelegatedRatio is a paid mutator transaction binding the contract method 0x81ffcdf1.
//
// Solidity: function updateMaxDelegatedRatio(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateMaxDelegatedRatio(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateMaxDelegatedRatio(&_SfcV400Constants.TransactOpts, v)
}

// UpdateMinAverageUptime is a paid mutator transaction binding the contract method 0x165e2639.
//
// Solidity: function updateMinAverageUptime(uint64 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateMinAverageUptime(opts *bind.TransactOpts, v uint64) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateMinAverageUptime", v)
}

// UpdateMinAverageUptime is a paid mutator transaction binding the contract method 0x165e2639.
//
// Solidity: function updateMinAverageUptime(uint64 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateMinAverageUptime(v uint64) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateMinAverageUptime(&_SfcV400Constants.TransactOpts, v)
}

// UpdateMinAverageUptime is a paid mutator transaction binding the contract method 0x165e2639.
//
// Solidity: function updateMinAverageUptime(uint64 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateMinAverageUptime(v uint64) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateMinAverageUptime(&_SfcV400Constants.TransactOpts, v)
}

// UpdateMinSelfStake is a paid mutator transaction binding the contract method 0x866c4b17.
//
// Solidity: function updateMinSelfStake(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateMinSelfStake(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateMinSelfStake", v)
}

// UpdateMinSelfStake is a paid mutator transaction binding the contract method 0x866c4b17.
//
// Solidity: function updateMinSelfStake(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateMinSelfStake(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateMinSelfStake(&_SfcV400Constants.TransactOpts, v)
}

// UpdateMinSelfStake is a paid mutator transaction binding the contract method 0x866c4b17.
//
// Solidity: function updateMinSelfStake(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateMinSelfStake(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateMinSelfStake(&_SfcV400Constants.TransactOpts, v)
}

// UpdateOfflinePenaltyThresholdBlocksNum is a paid mutator transaction binding the contract method 0x2e84e8e6.
//
// Solidity: function updateOfflinePenaltyThresholdBlocksNum(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateOfflinePenaltyThresholdBlocksNum(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateOfflinePenaltyThresholdBlocksNum", v)
}

// UpdateOfflinePenaltyThresholdBlocksNum is a paid mutator transaction binding the contract method 0x2e84e8e6.
//
// Solidity: function updateOfflinePenaltyThresholdBlocksNum(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateOfflinePenaltyThresholdBlocksNum(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateOfflinePenaltyThresholdBlocksNum(&_SfcV400Constants.TransactOpts, v)
}

// UpdateOfflinePenaltyThresholdBlocksNum is a paid mutator transaction binding the contract method 0x2e84e8e6.
//
// Solidity: function updateOfflinePenaltyThresholdBlocksNum(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateOfflinePenaltyThresholdBlocksNum(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateOfflinePenaltyThresholdBlocksNum(&_SfcV400Constants.TransactOpts, v)
}

// UpdateOfflinePenaltyThresholdTime is a paid mutator transaction binding the contract method 0x6348ebb8.
//
// Solidity: function updateOfflinePenaltyThresholdTime(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateOfflinePenaltyThresholdTime(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateOfflinePenaltyThresholdTime", v)
}

// UpdateOfflinePenaltyThresholdTime is a paid mutator transaction binding the contract method 0x6348ebb8.
//
// Solidity: function updateOfflinePenaltyThresholdTime(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateOfflinePenaltyThresholdTime(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateOfflinePenaltyThresholdTime(&_SfcV400Constants.TransactOpts, v)
}

// UpdateOfflinePenaltyThresholdTime is a paid mutator transaction binding the contract method 0x6348ebb8.
//
// Solidity: function updateOfflinePenaltyThresholdTime(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateOfflinePenaltyThresholdTime(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateOfflinePenaltyThresholdTime(&_SfcV400Constants.TransactOpts, v)
}

// UpdateTargetGasPowerPerSecond is a paid mutator transaction binding the contract method 0x43326867.
//
// Solidity: function updateTargetGasPowerPerSecond(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateTargetGasPowerPerSecond(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateTargetGasPowerPerSecond", v)
}

// UpdateTargetGasPowerPerSecond is a paid mutator transaction binding the contract method 0x43326867.
//
// Solidity: function updateTargetGasPowerPerSecond(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateTargetGasPowerPerSecond(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateTargetGasPowerPerSecond(&_SfcV400Constants.TransactOpts, v)
}

// UpdateTargetGasPowerPerSecond is a paid mutator transaction binding the contract method 0x43326867.
//
// Solidity: function updateTargetGasPowerPerSecond(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateTargetGasPowerPerSecond(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateTargetGasPowerPerSecond(&_SfcV400Constants.TransactOpts, v)
}

// UpdateTreasuryFeeShare is a paid mutator transaction binding the contract method 0xf8d5177e.
//
// Solidity: function updateTreasuryFeeShare(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateTreasuryFeeShare(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateTreasuryFeeShare", v)
}

// UpdateTreasuryFeeShare is a paid mutator transaction binding the contract method 0xf8d5177e.
//
// Solidity: function updateTreasuryFeeShare(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateTreasuryFeeShare(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateTreasuryFeeShare(&_SfcV400Constants.TransactOpts, v)
}

// UpdateTreasuryFeeShare is a paid mutator transaction binding the contract method 0xf8d5177e.
//
// Solidity: function updateTreasuryFeeShare(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateTreasuryFeeShare(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateTreasuryFeeShare(&_SfcV400Constants.TransactOpts, v)
}

// UpdateValidatorCommission is a paid mutator transaction binding the contract method 0x2ee71132.
//
// Solidity: function updateValidatorCommission(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateValidatorCommission(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateValidatorCommission", v)
}

// UpdateValidatorCommission is a paid mutator transaction binding the contract method 0x2ee71132.
//
// Solidity: function updateValidatorCommission(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateValidatorCommission(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateValidatorCommission(&_SfcV400Constants.TransactOpts, v)
}

// UpdateValidatorCommission is a paid mutator transaction binding the contract method 0x2ee71132.
//
// Solidity: function updateValidatorCommission(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateValidatorCommission(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateValidatorCommission(&_SfcV400Constants.TransactOpts, v)
}

// UpdateWithdrawalPeriodEpochs is a paid mutator transaction binding the contract method 0x8f078bfa.
//
// Solidity: function updateWithdrawalPeriodEpochs(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateWithdrawalPeriodEpochs(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateWithdrawalPeriodEpochs", v)
}

// UpdateWithdrawalPeriodEpochs is a paid mutator transaction binding the contract method 0x8f078bfa.
//
// Solidity: function updateWithdrawalPeriodEpochs(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateWithdrawalPeriodEpochs(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateWithdrawalPeriodEpochs(&_SfcV400Constants.TransactOpts, v)
}

// UpdateWithdrawalPeriodEpochs is a paid mutator transaction binding the contract method 0x8f078bfa.
//
// Solidity: function updateWithdrawalPeriodEpochs(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateWithdrawalPeriodEpochs(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateWithdrawalPeriodEpochs(&_SfcV400Constants.TransactOpts, v)
}

// UpdateWithdrawalPeriodTime is a paid mutator transaction binding the contract method 0x455366a4.
//
// Solidity: function updateWithdrawalPeriodTime(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactor) UpdateWithdrawalPeriodTime(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.contract.Transact(opts, "updateWithdrawalPeriodTime", v)
}

// UpdateWithdrawalPeriodTime is a paid mutator transaction binding the contract method 0x455366a4.
//
// Solidity: function updateWithdrawalPeriodTime(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsSession) UpdateWithdrawalPeriodTime(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateWithdrawalPeriodTime(&_SfcV400Constants.TransactOpts, v)
}

// UpdateWithdrawalPeriodTime is a paid mutator transaction binding the contract method 0x455366a4.
//
// Solidity: function updateWithdrawalPeriodTime(uint256 v) returns()
func (_SfcV400Constants *SfcV400ConstantsTransactorSession) UpdateWithdrawalPeriodTime(v *big.Int) (*types.Transaction, error) {
	return _SfcV400Constants.Contract.UpdateWithdrawalPeriodTime(&_SfcV400Constants.TransactOpts, v)
}

// SfcV400ConstantsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SfcV400Constants contract.
type SfcV400ConstantsInitializedIterator struct {
	Event *SfcV400ConstantsInitialized // Event containing the contract specifics and raw log

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
func (it *SfcV400ConstantsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ConstantsInitialized)
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
		it.Event = new(SfcV400ConstantsInitialized)
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
func (it *SfcV400ConstantsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ConstantsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ConstantsInitialized represents a Initialized event raised by the SfcV400Constants contract.
type SfcV400ConstantsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SfcV400Constants *SfcV400ConstantsFilterer) FilterInitialized(opts *bind.FilterOpts) (*SfcV400ConstantsInitializedIterator, error) {

	logs, sub, err := _SfcV400Constants.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SfcV400ConstantsInitializedIterator{contract: _SfcV400Constants.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SfcV400Constants *SfcV400ConstantsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SfcV400ConstantsInitialized) (event.Subscription, error) {

	logs, sub, err := _SfcV400Constants.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ConstantsInitialized)
				if err := _SfcV400Constants.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SfcV400Constants *SfcV400ConstantsFilterer) ParseInitialized(log types.Log) (*SfcV400ConstantsInitialized, error) {
	event := new(SfcV400ConstantsInitialized)
	if err := _SfcV400Constants.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ConstantsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SfcV400Constants contract.
type SfcV400ConstantsOwnershipTransferredIterator struct {
	Event *SfcV400ConstantsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SfcV400ConstantsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ConstantsOwnershipTransferred)
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
		it.Event = new(SfcV400ConstantsOwnershipTransferred)
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
func (it *SfcV400ConstantsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ConstantsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ConstantsOwnershipTransferred represents a OwnershipTransferred event raised by the SfcV400Constants contract.
type SfcV400ConstantsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV400Constants *SfcV400ConstantsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SfcV400ConstantsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcV400Constants.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ConstantsOwnershipTransferredIterator{contract: _SfcV400Constants.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV400Constants *SfcV400ConstantsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SfcV400ConstantsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcV400Constants.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ConstantsOwnershipTransferred)
				if err := _SfcV400Constants.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV400Constants *SfcV400ConstantsFilterer) ParseOwnershipTransferred(log types.Log) (*SfcV400ConstantsOwnershipTransferred, error) {
	event := new(SfcV400ConstantsOwnershipTransferred)
	if err := _SfcV400Constants.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
