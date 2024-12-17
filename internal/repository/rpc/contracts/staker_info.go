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

// StakerInfoMetaData contains all meta data concerning the StakerInfo contract.
var StakerInfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakerContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"InfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakerID\",\"type\":\"uint256\"}],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerInfos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_configUrl\",\"type\":\"string\"}],\"name\":\"updateInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakerContractAddress\",\"type\":\"address\"}],\"name\":\"updateStakerContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakerInfoABI is the input ABI used to generate the binding from.
// Deprecated: Use StakerInfoMetaData.ABI instead.
var StakerInfoABI = StakerInfoMetaData.ABI

// StakerInfo is an auto generated Go binding around an Ethereum contract.
type StakerInfo struct {
	StakerInfoCaller     // Read-only binding to the contract
	StakerInfoTransactor // Write-only binding to the contract
	StakerInfoFilterer   // Log filterer for contract events
}

// StakerInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakerInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakerInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakerInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakerInfoSession struct {
	Contract     *StakerInfo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakerInfoCallerSession struct {
	Contract *StakerInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakerInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakerInfoTransactorSession struct {
	Contract     *StakerInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakerInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakerInfoRaw struct {
	Contract *StakerInfo // Generic contract binding to access the raw methods on
}

// StakerInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakerInfoCallerRaw struct {
	Contract *StakerInfoCaller // Generic read-only contract binding to access the raw methods on
}

// StakerInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakerInfoTransactorRaw struct {
	Contract *StakerInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakerInfo creates a new instance of StakerInfo, bound to a specific deployed contract.
func NewStakerInfo(address common.Address, backend bind.ContractBackend) (*StakerInfo, error) {
	contract, err := bindStakerInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakerInfo{StakerInfoCaller: StakerInfoCaller{contract: contract}, StakerInfoTransactor: StakerInfoTransactor{contract: contract}, StakerInfoFilterer: StakerInfoFilterer{contract: contract}}, nil
}

// NewStakerInfoCaller creates a new read-only instance of StakerInfo, bound to a specific deployed contract.
func NewStakerInfoCaller(address common.Address, caller bind.ContractCaller) (*StakerInfoCaller, error) {
	contract, err := bindStakerInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakerInfoCaller{contract: contract}, nil
}

// NewStakerInfoTransactor creates a new write-only instance of StakerInfo, bound to a specific deployed contract.
func NewStakerInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*StakerInfoTransactor, error) {
	contract, err := bindStakerInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakerInfoTransactor{contract: contract}, nil
}

// NewStakerInfoFilterer creates a new log filterer instance of StakerInfo, bound to a specific deployed contract.
func NewStakerInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*StakerInfoFilterer, error) {
	contract, err := bindStakerInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakerInfoFilterer{contract: contract}, nil
}

// bindStakerInfo binds a generic wrapper to an already deployed contract.
func bindStakerInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakerInfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerInfo *StakerInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakerInfo.Contract.StakerInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerInfo *StakerInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerInfo.Contract.StakerInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerInfo *StakerInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerInfo.Contract.StakerInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerInfo *StakerInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakerInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerInfo *StakerInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerInfo *StakerInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerInfo.Contract.contract.Transact(opts, method, params...)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _stakerID) view returns(string)
func (_StakerInfo *StakerInfoCaller) GetInfo(opts *bind.CallOpts, _stakerID *big.Int) (string, error) {
	var out []interface{}
	err := _StakerInfo.contract.Call(opts, &out, "getInfo", _stakerID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _stakerID) view returns(string)
func (_StakerInfo *StakerInfoSession) GetInfo(_stakerID *big.Int) (string, error) {
	return _StakerInfo.Contract.GetInfo(&_StakerInfo.CallOpts, _stakerID)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _stakerID) view returns(string)
func (_StakerInfo *StakerInfoCallerSession) GetInfo(_stakerID *big.Int) (string, error) {
	return _StakerInfo.Contract.GetInfo(&_StakerInfo.CallOpts, _stakerID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakerInfo *StakerInfoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakerInfo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakerInfo *StakerInfoSession) Owner() (common.Address, error) {
	return _StakerInfo.Contract.Owner(&_StakerInfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakerInfo *StakerInfoCallerSession) Owner() (common.Address, error) {
	return _StakerInfo.Contract.Owner(&_StakerInfo.CallOpts)
}

// StakerInfos is a free data retrieval call binding the contract method 0x33470433.
//
// Solidity: function stakerInfos(uint256 ) view returns(string)
func (_StakerInfo *StakerInfoCaller) StakerInfos(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _StakerInfo.contract.Call(opts, &out, "stakerInfos", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StakerInfos is a free data retrieval call binding the contract method 0x33470433.
//
// Solidity: function stakerInfos(uint256 ) view returns(string)
func (_StakerInfo *StakerInfoSession) StakerInfos(arg0 *big.Int) (string, error) {
	return _StakerInfo.Contract.StakerInfos(&_StakerInfo.CallOpts, arg0)
}

// StakerInfos is a free data retrieval call binding the contract method 0x33470433.
//
// Solidity: function stakerInfos(uint256 ) view returns(string)
func (_StakerInfo *StakerInfoCallerSession) StakerInfos(arg0 *big.Int) (string, error) {
	return _StakerInfo.Contract.StakerInfos(&_StakerInfo.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakerInfo *StakerInfoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerInfo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakerInfo *StakerInfoSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakerInfo.Contract.RenounceOwnership(&_StakerInfo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakerInfo *StakerInfoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakerInfo.Contract.RenounceOwnership(&_StakerInfo.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakerInfo *StakerInfoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakerInfo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakerInfo *StakerInfoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakerInfo.Contract.TransferOwnership(&_StakerInfo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakerInfo *StakerInfoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakerInfo.Contract.TransferOwnership(&_StakerInfo.TransactOpts, newOwner)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xab29511b.
//
// Solidity: function updateInfo(string _configUrl) returns()
func (_StakerInfo *StakerInfoTransactor) UpdateInfo(opts *bind.TransactOpts, _configUrl string) (*types.Transaction, error) {
	return _StakerInfo.contract.Transact(opts, "updateInfo", _configUrl)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xab29511b.
//
// Solidity: function updateInfo(string _configUrl) returns()
func (_StakerInfo *StakerInfoSession) UpdateInfo(_configUrl string) (*types.Transaction, error) {
	return _StakerInfo.Contract.UpdateInfo(&_StakerInfo.TransactOpts, _configUrl)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xab29511b.
//
// Solidity: function updateInfo(string _configUrl) returns()
func (_StakerInfo *StakerInfoTransactorSession) UpdateInfo(_configUrl string) (*types.Transaction, error) {
	return _StakerInfo.Contract.UpdateInfo(&_StakerInfo.TransactOpts, _configUrl)
}

// UpdateStakerContractAddress is a paid mutator transaction binding the contract method 0xfb9cba30.
//
// Solidity: function updateStakerContractAddress(address _stakerContractAddress) returns()
func (_StakerInfo *StakerInfoTransactor) UpdateStakerContractAddress(opts *bind.TransactOpts, _stakerContractAddress common.Address) (*types.Transaction, error) {
	return _StakerInfo.contract.Transact(opts, "updateStakerContractAddress", _stakerContractAddress)
}

// UpdateStakerContractAddress is a paid mutator transaction binding the contract method 0xfb9cba30.
//
// Solidity: function updateStakerContractAddress(address _stakerContractAddress) returns()
func (_StakerInfo *StakerInfoSession) UpdateStakerContractAddress(_stakerContractAddress common.Address) (*types.Transaction, error) {
	return _StakerInfo.Contract.UpdateStakerContractAddress(&_StakerInfo.TransactOpts, _stakerContractAddress)
}

// UpdateStakerContractAddress is a paid mutator transaction binding the contract method 0xfb9cba30.
//
// Solidity: function updateStakerContractAddress(address _stakerContractAddress) returns()
func (_StakerInfo *StakerInfoTransactorSession) UpdateStakerContractAddress(_stakerContractAddress common.Address) (*types.Transaction, error) {
	return _StakerInfo.Contract.UpdateStakerContractAddress(&_StakerInfo.TransactOpts, _stakerContractAddress)
}

// StakerInfoInfoUpdatedIterator is returned from FilterInfoUpdated and is used to iterate over the raw logs and unpacked data for InfoUpdated events raised by the StakerInfo contract.
type StakerInfoInfoUpdatedIterator struct {
	Event *StakerInfoInfoUpdated // Event containing the contract specifics and raw log

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
func (it *StakerInfoInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerInfoInfoUpdated)
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
		it.Event = new(StakerInfoInfoUpdated)
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
func (it *StakerInfoInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerInfoInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerInfoInfoUpdated represents a InfoUpdated event raised by the StakerInfo contract.
type StakerInfoInfoUpdated struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInfoUpdated is a free log retrieval operation binding the contract event 0x3a668b70276c6b5af986be90ab9921c67bbef483987bb44cd5145c4984e59f24.
//
// Solidity: event InfoUpdated(uint256 stakerID)
func (_StakerInfo *StakerInfoFilterer) FilterInfoUpdated(opts *bind.FilterOpts) (*StakerInfoInfoUpdatedIterator, error) {

	logs, sub, err := _StakerInfo.contract.FilterLogs(opts, "InfoUpdated")
	if err != nil {
		return nil, err
	}
	return &StakerInfoInfoUpdatedIterator{contract: _StakerInfo.contract, event: "InfoUpdated", logs: logs, sub: sub}, nil
}

// WatchInfoUpdated is a free log subscription operation binding the contract event 0x3a668b70276c6b5af986be90ab9921c67bbef483987bb44cd5145c4984e59f24.
//
// Solidity: event InfoUpdated(uint256 stakerID)
func (_StakerInfo *StakerInfoFilterer) WatchInfoUpdated(opts *bind.WatchOpts, sink chan<- *StakerInfoInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _StakerInfo.contract.WatchLogs(opts, "InfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerInfoInfoUpdated)
				if err := _StakerInfo.contract.UnpackLog(event, "InfoUpdated", log); err != nil {
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

// ParseInfoUpdated is a log parse operation binding the contract event 0x3a668b70276c6b5af986be90ab9921c67bbef483987bb44cd5145c4984e59f24.
//
// Solidity: event InfoUpdated(uint256 stakerID)
func (_StakerInfo *StakerInfoFilterer) ParseInfoUpdated(log types.Log) (*StakerInfoInfoUpdated, error) {
	event := new(StakerInfoInfoUpdated)
	if err := _StakerInfo.contract.UnpackLog(event, "InfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerInfoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakerInfo contract.
type StakerInfoOwnershipTransferredIterator struct {
	Event *StakerInfoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakerInfoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerInfoOwnershipTransferred)
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
		it.Event = new(StakerInfoOwnershipTransferred)
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
func (it *StakerInfoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerInfoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerInfoOwnershipTransferred represents a OwnershipTransferred event raised by the StakerInfo contract.
type StakerInfoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakerInfo *StakerInfoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakerInfoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakerInfo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakerInfoOwnershipTransferredIterator{contract: _StakerInfo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakerInfo *StakerInfoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakerInfoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakerInfo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerInfoOwnershipTransferred)
				if err := _StakerInfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakerInfo *StakerInfoFilterer) ParseOwnershipTransferred(log types.Log) (*StakerInfoOwnershipTransferred, error) {
	event := new(StakerInfoOwnershipTransferred)
	if err := _StakerInfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
