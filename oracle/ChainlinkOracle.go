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
const ChainlinkOracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"name\":\"_payment\",\"type\":\"uint256\"},{\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedRequesters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumResponses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oracles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferLINK\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_clRequestId\",\"type\":\"bytes32\"},{\"name\":\"_response\",\"type\":\"int256\"}],\"name\":\"chainlinkCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"name\":\"_minimumResponses\",\"type\":\"uint128\"},{\"name\":\"_oracles\",\"type\":\"address[]\"},{\"name\":\"_jobIds\",\"type\":\"bytes32[]\"}],\"name\":\"updateRequestDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paymentAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestRateUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requester\",\"type\":\"address\"},{\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setAuthorization\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_link\",\"type\":\"address\"},{\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"name\":\"_minimumResponses\",\"type\":\"uint128\"},{\"name\":\"_oracles\",\"type\":\"address[]\"},{\"name\":\"_jobIds\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"response\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"answerId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ResponseReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"startedBy\",\"type\":\"address\"}],\"name\":\"NewRound\",\"type\":\"event\"}]"

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

// AuthorizedRequesters is a free data retrieval call binding the contract method 0x3ea478aa.
//
// Solidity: function authorizedRequesters(address ) view returns(bool)
func (_ChainlinkOracle *ChainlinkOracleCaller) AuthorizedRequesters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "authorizedRequesters", arg0)
	return *ret0, err
}

// AuthorizedRequesters is a free data retrieval call binding the contract method 0x3ea478aa.
//
// Solidity: function authorizedRequesters(address ) view returns(bool)
func (_ChainlinkOracle *ChainlinkOracleSession) AuthorizedRequesters(arg0 common.Address) (bool, error) {
	return _ChainlinkOracle.Contract.AuthorizedRequesters(&_ChainlinkOracle.CallOpts, arg0)
}

// AuthorizedRequesters is a free data retrieval call binding the contract method 0x3ea478aa.
//
// Solidity: function authorizedRequesters(address ) view returns(bool)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) AuthorizedRequesters(arg0 common.Address) (bool, error) {
	return _ChainlinkOracle.Contract.AuthorizedRequesters(&_ChainlinkOracle.CallOpts, arg0)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "getAnswer", _roundId)
	return *ret0, err
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetAnswer(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetAnswer(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "getTimestamp", _roundId)
	return *ret0, err
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetTimestamp(&_ChainlinkOracle.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _ChainlinkOracle.Contract.GetTimestamp(&_ChainlinkOracle.CallOpts, _roundId)
}

// JobIds is a free data retrieval call binding the contract method 0x4162cc88.
//
// Solidity: function jobIds(uint256 ) view returns(bytes32)
func (_ChainlinkOracle *ChainlinkOracleCaller) JobIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "jobIds", arg0)
	return *ret0, err
}

// JobIds is a free data retrieval call binding the contract method 0x4162cc88.
//
// Solidity: function jobIds(uint256 ) view returns(bytes32)
func (_ChainlinkOracle *ChainlinkOracleSession) JobIds(arg0 *big.Int) ([32]byte, error) {
	return _ChainlinkOracle.Contract.JobIds(&_ChainlinkOracle.CallOpts, arg0)
}

// JobIds is a free data retrieval call binding the contract method 0x4162cc88.
//
// Solidity: function jobIds(uint256 ) view returns(bytes32)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) JobIds(arg0 *big.Int) ([32]byte, error) {
	return _ChainlinkOracle.Contract.JobIds(&_ChainlinkOracle.CallOpts, arg0)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "latestAnswer")
	return *ret0, err
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestAnswer(&_ChainlinkOracle.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestAnswer() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestAnswer(&_ChainlinkOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "latestRound")
	return *ret0, err
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestRound() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestRound(&_ChainlinkOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestRound() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestRound(&_ChainlinkOracle.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "latestTimestamp")
	return *ret0, err
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestTimestamp(&_ChainlinkOracle.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) LatestTimestamp() (*big.Int, error) {
	return _ChainlinkOracle.Contract.LatestTimestamp(&_ChainlinkOracle.CallOpts)
}

// MinimumResponses is a free data retrieval call binding the contract method 0x54bcd7ff.
//
// Solidity: function minimumResponses() view returns(uint128)
func (_ChainlinkOracle *ChainlinkOracleCaller) MinimumResponses(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "minimumResponses")
	return *ret0, err
}

// MinimumResponses is a free data retrieval call binding the contract method 0x54bcd7ff.
//
// Solidity: function minimumResponses() view returns(uint128)
func (_ChainlinkOracle *ChainlinkOracleSession) MinimumResponses() (*big.Int, error) {
	return _ChainlinkOracle.Contract.MinimumResponses(&_ChainlinkOracle.CallOpts)
}

// MinimumResponses is a free data retrieval call binding the contract method 0x54bcd7ff.
//
// Solidity: function minimumResponses() view returns(uint128)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) MinimumResponses() (*big.Int, error) {
	return _ChainlinkOracle.Contract.MinimumResponses(&_ChainlinkOracle.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCaller) Oracles(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "oracles", arg0)
	return *ret0, err
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_ChainlinkOracle *ChainlinkOracleSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _ChainlinkOracle.Contract.Oracles(&_ChainlinkOracle.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0x5b69a7d8.
//
// Solidity: function oracles(uint256 ) view returns(address)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) Oracles(arg0 *big.Int) (common.Address, error) {
	return _ChainlinkOracle.Contract.Oracles(&_ChainlinkOracle.CallOpts, arg0)
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

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_ChainlinkOracle *ChainlinkOracleCaller) PaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChainlinkOracle.contract.Call(opts, out, "paymentAmount")
	return *ret0, err
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_ChainlinkOracle *ChainlinkOracleSession) PaymentAmount() (*big.Int, error) {
	return _ChainlinkOracle.Contract.PaymentAmount(&_ChainlinkOracle.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_ChainlinkOracle *ChainlinkOracleCallerSession) PaymentAmount() (*big.Int, error) {
	return _ChainlinkOracle.Contract.PaymentAmount(&_ChainlinkOracle.CallOpts)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x33bfcdd8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, uint256 _expiration) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) CancelRequest(opts *bind.TransactOpts, _requestId [32]byte, _payment *big.Int, _expiration *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "cancelRequest", _requestId, _payment, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x33bfcdd8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, uint256 _expiration) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _expiration *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.CancelRequest(&_ChainlinkOracle.TransactOpts, _requestId, _payment, _expiration)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x33bfcdd8.
//
// Solidity: function cancelRequest(bytes32 _requestId, uint256 _payment, uint256 _expiration) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) CancelRequest(_requestId [32]byte, _payment *big.Int, _expiration *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.CancelRequest(&_ChainlinkOracle.TransactOpts, _requestId, _payment, _expiration)
}

// ChainlinkCallback is a paid mutator transaction binding the contract method 0x6a9705b4.
//
// Solidity: function chainlinkCallback(bytes32 _clRequestId, int256 _response) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) ChainlinkCallback(opts *bind.TransactOpts, _clRequestId [32]byte, _response *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "chainlinkCallback", _clRequestId, _response)
}

// ChainlinkCallback is a paid mutator transaction binding the contract method 0x6a9705b4.
//
// Solidity: function chainlinkCallback(bytes32 _clRequestId, int256 _response) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) ChainlinkCallback(_clRequestId [32]byte, _response *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ChainlinkCallback(&_ChainlinkOracle.TransactOpts, _clRequestId, _response)
}

// ChainlinkCallback is a paid mutator transaction binding the contract method 0x6a9705b4.
//
// Solidity: function chainlinkCallback(bytes32 _clRequestId, int256 _response) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) ChainlinkCallback(_clRequestId [32]byte, _response *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.ChainlinkCallback(&_ChainlinkOracle.TransactOpts, _clRequestId, _response)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ChainlinkOracle *ChainlinkOracleSession) Destroy() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.Destroy(&_ChainlinkOracle.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) Destroy() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.Destroy(&_ChainlinkOracle.TransactOpts)
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

// RequestRateUpdate is a paid mutator transaction binding the contract method 0xdaa6d556.
//
// Solidity: function requestRateUpdate() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) RequestRateUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "requestRateUpdate")
}

// RequestRateUpdate is a paid mutator transaction binding the contract method 0xdaa6d556.
//
// Solidity: function requestRateUpdate() returns()
func (_ChainlinkOracle *ChainlinkOracleSession) RequestRateUpdate() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.RequestRateUpdate(&_ChainlinkOracle.TransactOpts)
}

// RequestRateUpdate is a paid mutator transaction binding the contract method 0xdaa6d556.
//
// Solidity: function requestRateUpdate() returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) RequestRateUpdate() (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.RequestRateUpdate(&_ChainlinkOracle.TransactOpts)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address _requester, bool _allowed) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) SetAuthorization(opts *bind.TransactOpts, _requester common.Address, _allowed bool) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "setAuthorization", _requester, _allowed)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address _requester, bool _allowed) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) SetAuthorization(_requester common.Address, _allowed bool) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.SetAuthorization(&_ChainlinkOracle.TransactOpts, _requester, _allowed)
}

// SetAuthorization is a paid mutator transaction binding the contract method 0xeecea000.
//
// Solidity: function setAuthorization(address _requester, bool _allowed) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) SetAuthorization(_requester common.Address, _allowed bool) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.SetAuthorization(&_ChainlinkOracle.TransactOpts, _requester, _allowed)
}

// TransferLINK is a paid mutator transaction binding the contract method 0x5cd9b90b.
//
// Solidity: function transferLINK(address _recipient, uint256 _amount) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) TransferLINK(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "transferLINK", _recipient, _amount)
}

// TransferLINK is a paid mutator transaction binding the contract method 0x5cd9b90b.
//
// Solidity: function transferLINK(address _recipient, uint256 _amount) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) TransferLINK(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferLINK(&_ChainlinkOracle.TransactOpts, _recipient, _amount)
}

// TransferLINK is a paid mutator transaction binding the contract method 0x5cd9b90b.
//
// Solidity: function transferLINK(address _recipient, uint256 _amount) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) TransferLINK(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferLINK(&_ChainlinkOracle.TransactOpts, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferOwnership(&_ChainlinkOracle.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.TransferOwnership(&_ChainlinkOracle.TransactOpts, _newOwner)
}

// UpdateRequestDetails is a paid mutator transaction binding the contract method 0x78a66674.
//
// Solidity: function updateRequestDetails(uint128 _paymentAmount, uint128 _minimumResponses, address[] _oracles, bytes32[] _jobIds) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactor) UpdateRequestDetails(opts *bind.TransactOpts, _paymentAmount *big.Int, _minimumResponses *big.Int, _oracles []common.Address, _jobIds [][32]byte) (*types.Transaction, error) {
	return _ChainlinkOracle.contract.Transact(opts, "updateRequestDetails", _paymentAmount, _minimumResponses, _oracles, _jobIds)
}

// UpdateRequestDetails is a paid mutator transaction binding the contract method 0x78a66674.
//
// Solidity: function updateRequestDetails(uint128 _paymentAmount, uint128 _minimumResponses, address[] _oracles, bytes32[] _jobIds) returns()
func (_ChainlinkOracle *ChainlinkOracleSession) UpdateRequestDetails(_paymentAmount *big.Int, _minimumResponses *big.Int, _oracles []common.Address, _jobIds [][32]byte) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.UpdateRequestDetails(&_ChainlinkOracle.TransactOpts, _paymentAmount, _minimumResponses, _oracles, _jobIds)
}

// UpdateRequestDetails is a paid mutator transaction binding the contract method 0x78a66674.
//
// Solidity: function updateRequestDetails(uint128 _paymentAmount, uint128 _minimumResponses, address[] _oracles, bytes32[] _jobIds) returns()
func (_ChainlinkOracle *ChainlinkOracleTransactorSession) UpdateRequestDetails(_paymentAmount *big.Int, _minimumResponses *big.Int, _oracles []common.Address, _jobIds [][32]byte) (*types.Transaction, error) {
	return _ChainlinkOracle.Contract.UpdateRequestDetails(&_ChainlinkOracle.TransactOpts, _paymentAmount, _minimumResponses, _oracles, _jobIds)
}

// ChainlinkOracleAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the ChainlinkOracle contract.
type ChainlinkOracleAnswerUpdatedIterator struct {
	Event *ChainlinkOracleAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleAnswerUpdated)
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
		it.Event = new(ChainlinkOracleAnswerUpdated)
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
func (it *ChainlinkOracleAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleAnswerUpdated represents a AnswerUpdated event raised by the ChainlinkOracle contract.
type ChainlinkOracleAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ChainlinkOracleAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleAnswerUpdatedIterator{contract: _ChainlinkOracle.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleAnswerUpdated)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseAnswerUpdated(log types.Log) (*ChainlinkOracleAnswerUpdated, error) {
	event := new(ChainlinkOracleAnswerUpdated)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkOracleChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the ChainlinkOracle contract.
type ChainlinkOracleChainlinkCancelledIterator struct {
	Event *ChainlinkOracleChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleChainlinkCancelled)
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
		it.Event = new(ChainlinkOracleChainlinkCancelled)
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
func (it *ChainlinkOracleChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleChainlinkCancelled represents a ChainlinkCancelled event raised by the ChainlinkOracle contract.
type ChainlinkOracleChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*ChainlinkOracleChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleChainlinkCancelledIterator{contract: _ChainlinkOracle.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleChainlinkCancelled)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseChainlinkCancelled(log types.Log) (*ChainlinkOracleChainlinkCancelled, error) {
	event := new(ChainlinkOracleChainlinkCancelled)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkOracleChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the ChainlinkOracle contract.
type ChainlinkOracleChainlinkFulfilledIterator struct {
	Event *ChainlinkOracleChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleChainlinkFulfilled)
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
		it.Event = new(ChainlinkOracleChainlinkFulfilled)
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
func (it *ChainlinkOracleChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleChainlinkFulfilled represents a ChainlinkFulfilled event raised by the ChainlinkOracle contract.
type ChainlinkOracleChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*ChainlinkOracleChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleChainlinkFulfilledIterator{contract: _ChainlinkOracle.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleChainlinkFulfilled)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseChainlinkFulfilled(log types.Log) (*ChainlinkOracleChainlinkFulfilled, error) {
	event := new(ChainlinkOracleChainlinkFulfilled)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkOracleChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the ChainlinkOracle contract.
type ChainlinkOracleChainlinkRequestedIterator struct {
	Event *ChainlinkOracleChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleChainlinkRequested)
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
		it.Event = new(ChainlinkOracleChainlinkRequested)
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
func (it *ChainlinkOracleChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleChainlinkRequested represents a ChainlinkRequested event raised by the ChainlinkOracle contract.
type ChainlinkOracleChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*ChainlinkOracleChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleChainlinkRequestedIterator{contract: _ChainlinkOracle.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleChainlinkRequested)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseChainlinkRequested(log types.Log) (*ChainlinkOracleChainlinkRequested, error) {
	event := new(ChainlinkOracleChainlinkRequested)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkOracleNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the ChainlinkOracle contract.
type ChainlinkOracleNewRoundIterator struct {
	Event *ChainlinkOracleNewRound // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleNewRound)
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
		it.Event = new(ChainlinkOracleNewRound)
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
func (it *ChainlinkOracleNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleNewRound represents a NewRound event raised by the ChainlinkOracle contract.
type ChainlinkOracleNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0xc3c45d1924f55369653f407ee9f095309d1e687b2c0011b1f709042d4f457e17.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ChainlinkOracleNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleNewRoundIterator{contract: _ChainlinkOracle.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0xc3c45d1924f55369653f407ee9f095309d1e687b2c0011b1f709042d4f457e17.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleNewRound)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0xc3c45d1924f55369653f407ee9f095309d1e687b2c0011b1f709042d4f457e17.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseNewRound(log types.Log) (*ChainlinkOracleNewRound, error) {
	event := new(ChainlinkOracleNewRound)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChainlinkOracleOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ChainlinkOracle contract.
type ChainlinkOracleOwnershipRenouncedIterator struct {
	Event *ChainlinkOracleOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleOwnershipRenounced)
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
		it.Event = new(ChainlinkOracleOwnershipRenounced)
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
func (it *ChainlinkOracleOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleOwnershipRenounced represents a OwnershipRenounced event raised by the ChainlinkOracle contract.
type ChainlinkOracleOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ChainlinkOracleOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleOwnershipRenouncedIterator{contract: _ChainlinkOracle.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleOwnershipRenounced)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseOwnershipRenounced(log types.Log) (*ChainlinkOracleOwnershipRenounced, error) {
	event := new(ChainlinkOracleOwnershipRenounced)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ChainlinkOracleResponseReceivedIterator is returned from FilterResponseReceived and is used to iterate over the raw logs and unpacked data for ResponseReceived events raised by the ChainlinkOracle contract.
type ChainlinkOracleResponseReceivedIterator struct {
	Event *ChainlinkOracleResponseReceived // Event containing the contract specifics and raw log

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
func (it *ChainlinkOracleResponseReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkOracleResponseReceived)
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
		it.Event = new(ChainlinkOracleResponseReceived)
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
func (it *ChainlinkOracleResponseReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkOracleResponseReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkOracleResponseReceived represents a ResponseReceived event raised by the ChainlinkOracle contract.
type ChainlinkOracleResponseReceived struct {
	Response *big.Int
	AnswerId *big.Int
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResponseReceived is a free log retrieval operation binding the contract event 0xb51168059c83c860caf5b830c5d2e64c2172c6fb2fe9f25447d9838e18d93b60.
//
// Solidity: event ResponseReceived(int256 indexed response, uint256 indexed answerId, address indexed sender)
func (_ChainlinkOracle *ChainlinkOracleFilterer) FilterResponseReceived(opts *bind.FilterOpts, response []*big.Int, answerId []*big.Int, sender []common.Address) (*ChainlinkOracleResponseReceivedIterator, error) {

	var responseRule []interface{}
	for _, responseItem := range response {
		responseRule = append(responseRule, responseItem)
	}
	var answerIdRule []interface{}
	for _, answerIdItem := range answerId {
		answerIdRule = append(answerIdRule, answerIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.FilterLogs(opts, "ResponseReceived", responseRule, answerIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkOracleResponseReceivedIterator{contract: _ChainlinkOracle.contract, event: "ResponseReceived", logs: logs, sub: sub}, nil
}

// WatchResponseReceived is a free log subscription operation binding the contract event 0xb51168059c83c860caf5b830c5d2e64c2172c6fb2fe9f25447d9838e18d93b60.
//
// Solidity: event ResponseReceived(int256 indexed response, uint256 indexed answerId, address indexed sender)
func (_ChainlinkOracle *ChainlinkOracleFilterer) WatchResponseReceived(opts *bind.WatchOpts, sink chan<- *ChainlinkOracleResponseReceived, response []*big.Int, answerId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var responseRule []interface{}
	for _, responseItem := range response {
		responseRule = append(responseRule, responseItem)
	}
	var answerIdRule []interface{}
	for _, answerIdItem := range answerId {
		answerIdRule = append(answerIdRule, answerIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ChainlinkOracle.contract.WatchLogs(opts, "ResponseReceived", responseRule, answerIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkOracleResponseReceived)
				if err := _ChainlinkOracle.contract.UnpackLog(event, "ResponseReceived", log); err != nil {
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

// ParseResponseReceived is a log parse operation binding the contract event 0xb51168059c83c860caf5b830c5d2e64c2172c6fb2fe9f25447d9838e18d93b60.
//
// Solidity: event ResponseReceived(int256 indexed response, uint256 indexed answerId, address indexed sender)
func (_ChainlinkOracle *ChainlinkOracleFilterer) ParseResponseReceived(log types.Log) (*ChainlinkOracleResponseReceived, error) {
	event := new(ChainlinkOracleResponseReceived)
	if err := _ChainlinkOracle.contract.UnpackLog(event, "ResponseReceived", log); err != nil {
		return nil, err
	}
	return event, nil
}
