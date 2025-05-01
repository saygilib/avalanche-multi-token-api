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

// ERC1155MetaData contains all meta data concerning the ERC1155 contract.
var ERC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155MetaData.ABI instead.
var ERC1155ABI = ERC1155MetaData.ABI

// ERC1155 is an auto generated Go binding around an Ethereum contract.
type ERC1155 struct {
	ERC1155Caller     // Read-only binding to the contract
	ERC1155Transactor // Write-only binding to the contract
	ERC1155Filterer   // Log filterer for contract events
}

// ERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155Session struct {
	Contract     *ERC1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155CallerSession struct {
	Contract *ERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TransactorSession struct {
	Contract     *ERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155Raw struct {
	Contract *ERC1155 // Generic contract binding to access the raw methods on
}

// ERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155CallerRaw struct {
	Contract *ERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TransactorRaw struct {
	Contract *ERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155 creates a new instance of ERC1155, bound to a specific deployed contract.
func NewERC1155(address common.Address, backend bind.ContractBackend) (*ERC1155, error) {
	contract, err := bindERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// NewERC1155Caller creates a new read-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Caller(address common.Address, caller bind.ContractCaller) (*ERC1155Caller, error) {
	contract, err := bindERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Caller{contract: contract}, nil
}

// NewERC1155Transactor creates a new write-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155Transactor, error) {
	contract, err := bindERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Transactor{contract: contract}, nil
}

// NewERC1155Filterer creates a new log filterer instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155Filterer, error) {
	contract, err := bindERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155Filterer{contract: contract}, nil
}

// bindERC1155 binds a generic wrapper to an already deployed contract.
func bindERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.ERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155 *ERC1155Caller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155 *ERC1155Session) Uri(tokenId *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Uri(tokenId *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, string uri) returns()
func (_ERC1155 *ERC1155Transactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, uri string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "mint", to, id, amount, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, string uri) returns()
func (_ERC1155 *ERC1155Session) Mint(to common.Address, id *big.Int, amount *big.Int, uri string) (*types.Transaction, error) {
	return _ERC1155.Contract.Mint(&_ERC1155.TransactOpts, to, id, amount, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, string uri) returns()
func (_ERC1155 *ERC1155TransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int, uri string) (*types.Transaction, error) {
	return _ERC1155.Contract.Mint(&_ERC1155.TransactOpts, to, id, amount, uri)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, value, data)
}
