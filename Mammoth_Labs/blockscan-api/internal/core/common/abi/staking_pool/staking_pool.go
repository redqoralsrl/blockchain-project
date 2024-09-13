// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking_pool

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

// StakingPoolMetaData contains all meta data concerning the StakingPool contract.
var StakingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimableRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolMetaData.ABI instead.
var StakingPoolABI = StakingPoolMetaData.ABI

// StakingPool is an auto generated Go binding around an Ethereum contract.
type StakingPool struct {
	StakingPoolCaller     // Read-only binding to the contract
	StakingPoolTransactor // Write-only binding to the contract
	StakingPoolFilterer   // Log filterer for contract events
}

// StakingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolSession struct {
	Contract     *StakingPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolCallerSession struct {
	Contract *StakingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StakingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolTransactorSession struct {
	Contract     *StakingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StakingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolRaw struct {
	Contract *StakingPool // Generic contract binding to access the raw methods on
}

// StakingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolCallerRaw struct {
	Contract *StakingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolTransactorRaw struct {
	Contract *StakingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPool creates a new instance of StakingPool, bound to a specific deployed contract.
func NewStakingPool(address common.Address, backend bind.ContractBackend) (*StakingPool, error) {
	contract, err := bindStakingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPool{StakingPoolCaller: StakingPoolCaller{contract: contract}, StakingPoolTransactor: StakingPoolTransactor{contract: contract}, StakingPoolFilterer: StakingPoolFilterer{contract: contract}}, nil
}

// NewStakingPoolCaller creates a new read-only instance of StakingPool, bound to a specific deployed contract.
func NewStakingPoolCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolCaller, error) {
	contract, err := bindStakingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolCaller{contract: contract}, nil
}

// NewStakingPoolTransactor creates a new write-only instance of StakingPool, bound to a specific deployed contract.
func NewStakingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolTransactor, error) {
	contract, err := bindStakingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolTransactor{contract: contract}, nil
}

// NewStakingPoolFilterer creates a new log filterer instance of StakingPool, bound to a specific deployed contract.
func NewStakingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolFilterer, error) {
	contract, err := bindStakingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFilterer{contract: contract}, nil
}

// bindStakingPool binds a generic wrapper to an already deployed contract.
func bindStakingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPool *StakingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPool.Contract.StakingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPool *StakingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPool.Contract.StakingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPool *StakingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPool.Contract.StakingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPool *StakingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPool *StakingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPool *StakingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPool.Contract.contract.Transact(opts, method, params...)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0x6be9dcce.
//
// Solidity: function claimableRewards(address validator, address staker) view returns(uint256)
func (_StakingPool *StakingPoolCaller) ClaimableRewards(opts *bind.CallOpts, validator common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPool.contract.Call(opts, &out, "claimableRewards", validator, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableRewards is a free data retrieval call binding the contract method 0x6be9dcce.
//
// Solidity: function claimableRewards(address validator, address staker) view returns(uint256)
func (_StakingPool *StakingPoolSession) ClaimableRewards(validator common.Address, staker common.Address) (*big.Int, error) {
	return _StakingPool.Contract.ClaimableRewards(&_StakingPool.CallOpts, validator, staker)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0x6be9dcce.
//
// Solidity: function claimableRewards(address validator, address staker) view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) ClaimableRewards(validator common.Address, staker common.Address) (*big.Int, error) {
	return _StakingPool.Contract.ClaimableRewards(&_StakingPool.CallOpts, validator, staker)
}

// GetStakedAmount is a free data retrieval call binding the contract method 0x0db14e95.
//
// Solidity: function getStakedAmount(address validator, address staker) view returns(uint256)
func (_StakingPool *StakingPoolCaller) GetStakedAmount(opts *bind.CallOpts, validator common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPool.contract.Call(opts, &out, "getStakedAmount", validator, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedAmount is a free data retrieval call binding the contract method 0x0db14e95.
//
// Solidity: function getStakedAmount(address validator, address staker) view returns(uint256)
func (_StakingPool *StakingPoolSession) GetStakedAmount(validator common.Address, staker common.Address) (*big.Int, error) {
	return _StakingPool.Contract.GetStakedAmount(&_StakingPool.CallOpts, validator, staker)
}

// GetStakedAmount is a free data retrieval call binding the contract method 0x0db14e95.
//
// Solidity: function getStakedAmount(address validator, address staker) view returns(uint256)
func (_StakingPool *StakingPoolCallerSession) GetStakedAmount(validator common.Address, staker common.Address) (*big.Int, error) {
	return _StakingPool.Contract.GetStakedAmount(&_StakingPool.CallOpts, validator, staker)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address validator) returns()
func (_StakingPool *StakingPoolTransactor) Claim(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _StakingPool.contract.Transact(opts, "claim", validator)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address validator) returns()
func (_StakingPool *StakingPoolSession) Claim(validator common.Address) (*types.Transaction, error) {
	return _StakingPool.Contract.Claim(&_StakingPool.TransactOpts, validator)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address validator) returns()
func (_StakingPool *StakingPoolTransactorSession) Claim(validator common.Address) (*types.Transaction, error) {
	return _StakingPool.Contract.Claim(&_StakingPool.TransactOpts, validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns()
func (_StakingPool *StakingPoolTransactor) Stake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _StakingPool.contract.Transact(opts, "stake", validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns()
func (_StakingPool *StakingPoolSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _StakingPool.Contract.Stake(&_StakingPool.TransactOpts, validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns()
func (_StakingPool *StakingPoolTransactorSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _StakingPool.Contract.Stake(&_StakingPool.TransactOpts, validator)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_StakingPool *StakingPoolTransactor) Unstake(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.contract.Transact(opts, "unstake", validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_StakingPool *StakingPoolSession) Unstake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.Contract.Unstake(&_StakingPool.TransactOpts, validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_StakingPool *StakingPoolTransactorSession) Unstake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingPool.Contract.Unstake(&_StakingPool.TransactOpts, validator, amount)
}
