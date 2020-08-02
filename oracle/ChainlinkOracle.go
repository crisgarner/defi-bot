// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChainlinkOracleABI is the input ABI used to generate the binding from.
const ChainlinkOracleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLatestAnswer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_back\",\"type\":\"uint256\"}],\"name\":\"getPreviousAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_back\",\"type\":\"uint256\"}],\"name\":\"getPreviousTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setReferenceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChainlinkOracle is an auto generated Go binding around an Ethereum contract.
type ChainlinkOracle struct {
	ChainlinkOracleCaller     // Read-only binding to the contract
	ChainlinkOracleTransactor // Write-only binding to the contract
	ChainlinkOracleFilterer   // Log filterer for contract events
}

// ChainlinkOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkOracleSession struct {
	Contract     *ChainlinkOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainlinkOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkOracleCallerSession struct {
	Contract *ChainlinkOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ChainlinkOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkOracleTransactorSession struct {
	Contract     *ChainlinkOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChainlinkOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkOracleRaw struct {
	Contract *ChainlinkOracle // Generic contract binding to access the raw methods on
}

// ChainlinkOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkOracleCallerRaw struct {
	Contract *ChainlinkOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkOracleTransactorRaw struct {
	Contract *ChainlinkOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkOracle creates a new instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracle(address common.Address, backend bind.ContractBackend) (*ChainlinkOracle, error) {
	contract, err := bindChainlinkOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracle{ChainlinkOracleCaller: ChainlinkOracleCaller{contract: contract}, ChainlinkOracleTransactor: ChainlinkOracleTransactor{contract: contract}, ChainlinkOracleFilterer: ChainlinkOracleFilterer{contract: contract}}, nil
}

// NewChainlinkOracleCaller creates a new read-only instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracleCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkOracleCaller, error) {
	contract, err := bindChainlinkOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleCaller{contract: contract}, nil
}

// NewChainlinkOracleTransactor creates a new write-only instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkOracleTransactor, error) {
	contract, err := bindChainlinkOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleTransactor{contract: contract}, nil
}

// NewChainlinkOracleFilterer creates a new log filterer instance of ChainlinkOracle, bound to a specific deployed contract.
func NewChainlinkOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkOracleFilterer, error) {
	contract, err := bindChainlinkOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleFilterer{contract: contract}, nil
}

// bindChainlinkOracle binds a generic wrapper to an already deployed contract.
func bindChainlinkOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkOracle *ChainlinkOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkOracle.Contract.ChainlinkOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkOracle *ChainlinkOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ChainlinkOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkOracle *ChainlinkOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ChainlinkOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkOracle *ChainlinkOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChainlinkOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkOracle *ChainlinkOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkOracle *ChainlinkOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.contract.Transact(opts, method, params...)
}

// GetLatestAnswer is a free data retrieval call binding the contract method 0x96237c02.
//
// Solidity: function getLatestAnswer() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetLatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "getLatestAnswer")
	return *ret0, err
}

// GetLatestAnswer is a free data retrieval call binding the contract method 0x96237c02.
//
// Solidity: function getLatestAnswer() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetLatestAnswer() (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetLatestAnswer(&_ChainlinkOracle.CallOpts)
}

// GetLatestAnswer is a free data retrieval call binding the contract method 0x96237c02.
//
// Solidity: function getLatestAnswer() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetLatestAnswer() (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetLatestAnswer(&_ChainlinkOracle.CallOpts)
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0xf43b52cb.
//
// Solidity: function getLatestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetLatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "getLatestTimestamp")
	return *ret0, err
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0xf43b52cb.
//
// Solidity: function getLatestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetLatestTimestamp() (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetLatestTimestamp(&_ChainlinkOracle.CallOpts)
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0xf43b52cb.
//
// Solidity: function getLatestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetLatestTimestamp() (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetLatestTimestamp(&_ChainlinkOracle.CallOpts)
}

// GetPreviousAnswer is a free data retrieval call binding the contract method 0xd8d7e950.
//
// Solidity: function getPreviousAnswer(uint256 _back) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetPreviousAnswer(opts *bind.CallOpts, _back *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "getPreviousAnswer", _back)
	return *ret0, err
}

// GetPreviousAnswer is a free data retrieval call binding the contract method 0xd8d7e950.
//
// Solidity: function getPreviousAnswer(uint256 _back) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetPreviousAnswer(_back *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetPreviousAnswer(&_ChainlinkOracle.CallOpts, _back)
}

// GetPreviousAnswer is a free data retrieval call binding the contract method 0xd8d7e950.
//
// Solidity: function getPreviousAnswer(uint256 _back) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetPreviousAnswer(_back *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetPreviousAnswer(&_ChainlinkOracle.CallOpts, _back)
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x03892125.
//
// Solidity: function getPreviousTimestamp(uint256 _back) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetPreviousTimestamp(opts *bind.CallOpts, _back *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "getPreviousTimestamp", _back)
	return *ret0, err
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x03892125.
//
// Solidity: function getPreviousTimestamp(uint256 _back) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetPreviousTimestamp(_back *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetPreviousTimestamp(&_ChainlinkOracle.CallOpts, _back)
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x03892125.
//
// Solidity: function getPreviousTimestamp(uint256 _back) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetPreviousTimestamp(_back *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetPreviousTimestamp(&_ChainlinkOracle.CallOpts, _back)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) Owner() (common.Address, error) {
	return _ChainlinkOracle.Contract.Owner(&_ChainlinkOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Owner() (common.Address, error) {
	return _ChainlinkOracle.Contract.Owner(&_ChainlinkOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkOracle *ChainlinkOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.RenounceOwnership(&_ChainlinkOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.RenounceOwnership(&_ChainlinkOracle.TransactOpts)
}

// SetReferenceContract is a paid mutator transaction binding the contract method 0x3cf8a697.
//
// Solidity: function setReferenceContract(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) SetReferenceContract(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "setReferenceContract", _aggregator)
}

// SetReferenceContract is a paid mutator transaction binding the contract method 0x3cf8a697.
//
// Solidity: function setReferenceContract(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) SetReferenceContract(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.SetReferenceContract(&_ChainlinkOracle.TransactOpts, _aggregator)
}

// SetReferenceContract is a paid mutator transaction binding the contract method 0x3cf8a697.
//
// Solidity: function setReferenceContract(address _aggregator) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) SetReferenceContract(_aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.SetReferenceContract(&_ChainlinkOracle.TransactOpts, _aggregator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferOwnership(&_ChainlinkOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferOwnership(&_ChainlinkOracle.TransactOpts, newOwner)
}

// ChainlinkOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainlinkOracle contract.
type ChainlinkOracleOwnershipTransferredIterator struct {
	Event *ChainlinkOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleOwnershipTransferred)
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
		it.Event = new(ChainlinkOracleOwnershipTransferred)
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
func (it *ChainlinkOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleOwnershipTransferred represents a OwnershipTransferred event raised by the ChainlinkOracle contract.
type ChainlinkOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChainlinkOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleOwnershipTransferredIterator{contract: _ChainlinkOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleOwnershipTransferred)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseOwnershipTransferred(log types.Log) (*ChainlinkOracleOwnershipTransferred, error) {
	event := new(ChainlinkOracleOwnershipTransferred)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
