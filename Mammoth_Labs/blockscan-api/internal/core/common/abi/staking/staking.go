// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"ValidatorDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"ValidatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"}],\"name\":\"ValidatorModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"ValidatorOwnerClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"ValidatorReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"slashes\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ECOSYSTEM_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"activateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"calcAvailableForRedelegateAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsDust\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"}],\"name\":\"changeValidatorCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeValidatorOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"claimDelegatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"beforeEpoch\",\"type\":\"uint64\"}],\"name\":\"claimDelegatorFeeAtEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"claimValidatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"beforeEpoch\",\"type\":\"uint64\"}],\"name\":\"claimValidatorFeeAtEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"initialStakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"}],\"name\":\"ctor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"disableValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"}],\"name\":\"getDelegatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"}],\"name\":\"getPendingDelegatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"getPendingValidatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getValidatorByOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getValidatorDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"atEpoch\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"getValidatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalDelegated\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"slashesCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"changedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"jailedBefore\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalRewards\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"getValidatorStatusAtEpoch\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalDelegated\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"slashesCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"changedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"jailedBefore\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalRewards\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEpoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"redelegateDelegatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commissionRate\",\"type\":\"uint16\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"releaseValidatorFromJail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setEcosystemReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// ECOSYSTEMREWARD is a free data retrieval call binding the contract method 0x8aadbbf3.
//
// Solidity: function ECOSYSTEM_REWARD() view returns(uint256)
func (_Staking *StakingCaller) ECOSYSTEMREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "ECOSYSTEM_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ECOSYSTEMREWARD is a free data retrieval call binding the contract method 0x8aadbbf3.
//
// Solidity: function ECOSYSTEM_REWARD() view returns(uint256)
func (_Staking *StakingSession) ECOSYSTEMREWARD() (*big.Int, error) {
	return _Staking.Contract.ECOSYSTEMREWARD(&_Staking.CallOpts)
}

// ECOSYSTEMREWARD is a free data retrieval call binding the contract method 0x8aadbbf3.
//
// Solidity: function ECOSYSTEM_REWARD() view returns(uint256)
func (_Staking *StakingCallerSession) ECOSYSTEMREWARD() (*big.Int, error) {
	return _Staking.Contract.ECOSYSTEMREWARD(&_Staking.CallOpts)
}

// CalcAvailableForRedelegateAmount is a free data retrieval call binding the contract method 0x5ef9e8c6.
//
// Solidity: function calcAvailableForRedelegateAmount(address validator, address delegator) view returns(uint256 amountToStake, uint256 rewardsDust)
func (_Staking *StakingCaller) CalcAvailableForRedelegateAmount(opts *bind.CallOpts, validator common.Address, delegator common.Address) (struct {
	AmountToStake *big.Int
	RewardsDust   *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "calcAvailableForRedelegateAmount", validator, delegator)

	outstruct := new(struct {
		AmountToStake *big.Int
		RewardsDust   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountToStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardsDust = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalcAvailableForRedelegateAmount is a free data retrieval call binding the contract method 0x5ef9e8c6.
//
// Solidity: function calcAvailableForRedelegateAmount(address validator, address delegator) view returns(uint256 amountToStake, uint256 rewardsDust)
func (_Staking *StakingSession) CalcAvailableForRedelegateAmount(validator common.Address, delegator common.Address) (struct {
	AmountToStake *big.Int
	RewardsDust   *big.Int
}, error) {
	return _Staking.Contract.CalcAvailableForRedelegateAmount(&_Staking.CallOpts, validator, delegator)
}

// CalcAvailableForRedelegateAmount is a free data retrieval call binding the contract method 0x5ef9e8c6.
//
// Solidity: function calcAvailableForRedelegateAmount(address validator, address delegator) view returns(uint256 amountToStake, uint256 rewardsDust)
func (_Staking *StakingCallerSession) CalcAvailableForRedelegateAmount(validator common.Address, delegator common.Address) (struct {
	AmountToStake *big.Int
	RewardsDust   *big.Int
}, error) {
	return _Staking.Contract.CalcAvailableForRedelegateAmount(&_Staking.CallOpts, validator, delegator)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint64)
func (_Staking *StakingCaller) CurrentEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint64)
func (_Staking *StakingSession) CurrentEpoch() (uint64, error) {
	return _Staking.Contract.CurrentEpoch(&_Staking.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint64)
func (_Staking *StakingCallerSession) CurrentEpoch() (uint64, error) {
	return _Staking.Contract.CurrentEpoch(&_Staking.CallOpts)
}

// GetDelegatorFee is a free data retrieval call binding the contract method 0x52b7bea2.
//
// Solidity: function getDelegatorFee(address validatorAddress, address delegatorAddress) view returns(uint256)
func (_Staking *StakingCaller) GetDelegatorFee(opts *bind.CallOpts, validatorAddress common.Address, delegatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegatorFee", validatorAddress, delegatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatorFee is a free data retrieval call binding the contract method 0x52b7bea2.
//
// Solidity: function getDelegatorFee(address validatorAddress, address delegatorAddress) view returns(uint256)
func (_Staking *StakingSession) GetDelegatorFee(validatorAddress common.Address, delegatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetDelegatorFee(&_Staking.CallOpts, validatorAddress, delegatorAddress)
}

// GetDelegatorFee is a free data retrieval call binding the contract method 0x52b7bea2.
//
// Solidity: function getDelegatorFee(address validatorAddress, address delegatorAddress) view returns(uint256)
func (_Staking *StakingCallerSession) GetDelegatorFee(validatorAddress common.Address, delegatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetDelegatorFee(&_Staking.CallOpts, validatorAddress, delegatorAddress)
}

// GetPendingDelegatorFee is a free data retrieval call binding the contract method 0xc2fd58fc.
//
// Solidity: function getPendingDelegatorFee(address validatorAddress, address delegatorAddress) view returns(uint256)
func (_Staking *StakingCaller) GetPendingDelegatorFee(opts *bind.CallOpts, validatorAddress common.Address, delegatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getPendingDelegatorFee", validatorAddress, delegatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingDelegatorFee is a free data retrieval call binding the contract method 0xc2fd58fc.
//
// Solidity: function getPendingDelegatorFee(address validatorAddress, address delegatorAddress) view returns(uint256)
func (_Staking *StakingSession) GetPendingDelegatorFee(validatorAddress common.Address, delegatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetPendingDelegatorFee(&_Staking.CallOpts, validatorAddress, delegatorAddress)
}

// GetPendingDelegatorFee is a free data retrieval call binding the contract method 0xc2fd58fc.
//
// Solidity: function getPendingDelegatorFee(address validatorAddress, address delegatorAddress) view returns(uint256)
func (_Staking *StakingCallerSession) GetPendingDelegatorFee(validatorAddress common.Address, delegatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetPendingDelegatorFee(&_Staking.CallOpts, validatorAddress, delegatorAddress)
}

// GetPendingValidatorFee is a free data retrieval call binding the contract method 0xc6fb9065.
//
// Solidity: function getPendingValidatorFee(address validatorAddress) view returns(uint256)
func (_Staking *StakingCaller) GetPendingValidatorFee(opts *bind.CallOpts, validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getPendingValidatorFee", validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingValidatorFee is a free data retrieval call binding the contract method 0xc6fb9065.
//
// Solidity: function getPendingValidatorFee(address validatorAddress) view returns(uint256)
func (_Staking *StakingSession) GetPendingValidatorFee(validatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetPendingValidatorFee(&_Staking.CallOpts, validatorAddress)
}

// GetPendingValidatorFee is a free data retrieval call binding the contract method 0xc6fb9065.
//
// Solidity: function getPendingValidatorFee(address validatorAddress) view returns(uint256)
func (_Staking *StakingCallerSession) GetPendingValidatorFee(validatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetPendingValidatorFee(&_Staking.CallOpts, validatorAddress)
}

// GetValidatorByOwner is a free data retrieval call binding the contract method 0x30108c22.
//
// Solidity: function getValidatorByOwner(address owner) view returns(address)
func (_Staking *StakingCaller) GetValidatorByOwner(opts *bind.CallOpts, owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorByOwner", owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorByOwner is a free data retrieval call binding the contract method 0x30108c22.
//
// Solidity: function getValidatorByOwner(address owner) view returns(address)
func (_Staking *StakingSession) GetValidatorByOwner(owner common.Address) (common.Address, error) {
	return _Staking.Contract.GetValidatorByOwner(&_Staking.CallOpts, owner)
}

// GetValidatorByOwner is a free data retrieval call binding the contract method 0x30108c22.
//
// Solidity: function getValidatorByOwner(address owner) view returns(address)
func (_Staking *StakingCallerSession) GetValidatorByOwner(owner common.Address) (common.Address, error) {
	return _Staking.Contract.GetValidatorByOwner(&_Staking.CallOpts, owner)
}

// GetValidatorDelegation is a free data retrieval call binding the contract method 0xd951e186.
//
// Solidity: function getValidatorDelegation(address validatorAddress, address delegator) view returns(uint256 delegatedAmount, uint64 atEpoch)
func (_Staking *StakingCaller) GetValidatorDelegation(opts *bind.CallOpts, validatorAddress common.Address, delegator common.Address) (struct {
	DelegatedAmount *big.Int
	AtEpoch         uint64
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorDelegation", validatorAddress, delegator)

	outstruct := new(struct {
		DelegatedAmount *big.Int
		AtEpoch         uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegatedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AtEpoch = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetValidatorDelegation is a free data retrieval call binding the contract method 0xd951e186.
//
// Solidity: function getValidatorDelegation(address validatorAddress, address delegator) view returns(uint256 delegatedAmount, uint64 atEpoch)
func (_Staking *StakingSession) GetValidatorDelegation(validatorAddress common.Address, delegator common.Address) (struct {
	DelegatedAmount *big.Int
	AtEpoch         uint64
}, error) {
	return _Staking.Contract.GetValidatorDelegation(&_Staking.CallOpts, validatorAddress, delegator)
}

// GetValidatorDelegation is a free data retrieval call binding the contract method 0xd951e186.
//
// Solidity: function getValidatorDelegation(address validatorAddress, address delegator) view returns(uint256 delegatedAmount, uint64 atEpoch)
func (_Staking *StakingCallerSession) GetValidatorDelegation(validatorAddress common.Address, delegator common.Address) (struct {
	DelegatedAmount *big.Int
	AtEpoch         uint64
}, error) {
	return _Staking.Contract.GetValidatorDelegation(&_Staking.CallOpts, validatorAddress, delegator)
}

// GetValidatorFee is a free data retrieval call binding the contract method 0x457179fd.
//
// Solidity: function getValidatorFee(address validatorAddress) view returns(uint256)
func (_Staking *StakingCaller) GetValidatorFee(opts *bind.CallOpts, validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorFee", validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorFee is a free data retrieval call binding the contract method 0x457179fd.
//
// Solidity: function getValidatorFee(address validatorAddress) view returns(uint256)
func (_Staking *StakingSession) GetValidatorFee(validatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetValidatorFee(&_Staking.CallOpts, validatorAddress)
}

// GetValidatorFee is a free data retrieval call binding the contract method 0x457179fd.
//
// Solidity: function getValidatorFee(address validatorAddress) view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorFee(validatorAddress common.Address) (*big.Int, error) {
	return _Staking.Contract.GetValidatorFee(&_Staking.CallOpts, validatorAddress)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address validatorAddress) view returns(address ownerAddress, uint8 status, uint256 totalDelegated, uint32 slashesCount, uint64 changedAt, uint64 jailedBefore, uint64 claimedAt, uint16 commissionRate, uint96 totalRewards)
func (_Staking *StakingCaller) GetValidatorStatus(opts *bind.CallOpts, validatorAddress common.Address) (struct {
	OwnerAddress   common.Address
	Status         uint8
	TotalDelegated *big.Int
	SlashesCount   uint32
	ChangedAt      uint64
	JailedBefore   uint64
	ClaimedAt      uint64
	CommissionRate uint16
	TotalRewards   *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorStatus", validatorAddress)

	outstruct := new(struct {
		OwnerAddress   common.Address
		Status         uint8
		TotalDelegated *big.Int
		SlashesCount   uint32
		ChangedAt      uint64
		JailedBefore   uint64
		ClaimedAt      uint64
		CommissionRate uint16
		TotalRewards   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OwnerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.TotalDelegated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SlashesCount = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.ChangedAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.JailedBefore = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.ClaimedAt = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.CommissionRate = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.TotalRewards = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address validatorAddress) view returns(address ownerAddress, uint8 status, uint256 totalDelegated, uint32 slashesCount, uint64 changedAt, uint64 jailedBefore, uint64 claimedAt, uint16 commissionRate, uint96 totalRewards)
func (_Staking *StakingSession) GetValidatorStatus(validatorAddress common.Address) (struct {
	OwnerAddress   common.Address
	Status         uint8
	TotalDelegated *big.Int
	SlashesCount   uint32
	ChangedAt      uint64
	JailedBefore   uint64
	ClaimedAt      uint64
	CommissionRate uint16
	TotalRewards   *big.Int
}, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, validatorAddress)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address validatorAddress) view returns(address ownerAddress, uint8 status, uint256 totalDelegated, uint32 slashesCount, uint64 changedAt, uint64 jailedBefore, uint64 claimedAt, uint16 commissionRate, uint96 totalRewards)
func (_Staking *StakingCallerSession) GetValidatorStatus(validatorAddress common.Address) (struct {
	OwnerAddress   common.Address
	Status         uint8
	TotalDelegated *big.Int
	SlashesCount   uint32
	ChangedAt      uint64
	JailedBefore   uint64
	ClaimedAt      uint64
	CommissionRate uint16
	TotalRewards   *big.Int
}, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, validatorAddress)
}

// GetValidatorStatusAtEpoch is a free data retrieval call binding the contract method 0x48124d26.
//
// Solidity: function getValidatorStatusAtEpoch(address validatorAddress, uint64 epoch) view returns(address ownerAddress, uint8 status, uint256 totalDelegated, uint32 slashesCount, uint64 changedAt, uint64 jailedBefore, uint64 claimedAt, uint16 commissionRate, uint96 totalRewards)
func (_Staking *StakingCaller) GetValidatorStatusAtEpoch(opts *bind.CallOpts, validatorAddress common.Address, epoch uint64) (struct {
	OwnerAddress   common.Address
	Status         uint8
	TotalDelegated *big.Int
	SlashesCount   uint32
	ChangedAt      uint64
	JailedBefore   uint64
	ClaimedAt      uint64
	CommissionRate uint16
	TotalRewards   *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorStatusAtEpoch", validatorAddress, epoch)

	outstruct := new(struct {
		OwnerAddress   common.Address
		Status         uint8
		TotalDelegated *big.Int
		SlashesCount   uint32
		ChangedAt      uint64
		JailedBefore   uint64
		ClaimedAt      uint64
		CommissionRate uint16
		TotalRewards   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OwnerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.TotalDelegated = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SlashesCount = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.ChangedAt = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.JailedBefore = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.ClaimedAt = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.CommissionRate = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.TotalRewards = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorStatusAtEpoch is a free data retrieval call binding the contract method 0x48124d26.
//
// Solidity: function getValidatorStatusAtEpoch(address validatorAddress, uint64 epoch) view returns(address ownerAddress, uint8 status, uint256 totalDelegated, uint32 slashesCount, uint64 changedAt, uint64 jailedBefore, uint64 claimedAt, uint16 commissionRate, uint96 totalRewards)
func (_Staking *StakingSession) GetValidatorStatusAtEpoch(validatorAddress common.Address, epoch uint64) (struct {
	OwnerAddress   common.Address
	Status         uint8
	TotalDelegated *big.Int
	SlashesCount   uint32
	ChangedAt      uint64
	JailedBefore   uint64
	ClaimedAt      uint64
	CommissionRate uint16
	TotalRewards   *big.Int
}, error) {
	return _Staking.Contract.GetValidatorStatusAtEpoch(&_Staking.CallOpts, validatorAddress, epoch)
}

// GetValidatorStatusAtEpoch is a free data retrieval call binding the contract method 0x48124d26.
//
// Solidity: function getValidatorStatusAtEpoch(address validatorAddress, uint64 epoch) view returns(address ownerAddress, uint8 status, uint256 totalDelegated, uint32 slashesCount, uint64 changedAt, uint64 jailedBefore, uint64 claimedAt, uint16 commissionRate, uint96 totalRewards)
func (_Staking *StakingCallerSession) GetValidatorStatusAtEpoch(validatorAddress common.Address, epoch uint64) (struct {
	OwnerAddress   common.Address
	Status         uint8
	TotalDelegated *big.Int
	SlashesCount   uint32
	ChangedAt      uint64
	JailedBefore   uint64
	ClaimedAt      uint64
	CommissionRate uint16
	TotalRewards   *big.Int
}, error) {
	return _Staking.Contract.GetValidatorStatusAtEpoch(&_Staking.CallOpts, validatorAddress, epoch)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingSession) GetValidators() ([]common.Address, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingCallerSession) GetValidators() ([]common.Address, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) view returns(bool)
func (_Staking *StakingCaller) IsValidator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) view returns(bool)
func (_Staking *StakingSession) IsValidator(account common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) view returns(bool)
func (_Staking *StakingCallerSession) IsValidator(account common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, account)
}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address account) view returns(bool)
func (_Staking *StakingCaller) IsValidatorActive(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidatorActive", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address account) view returns(bool)
func (_Staking *StakingSession) IsValidatorActive(account common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorActive(&_Staking.CallOpts, account)
}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address account) view returns(bool)
func (_Staking *StakingCallerSession) IsValidatorActive(account common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorActive(&_Staking.CallOpts, account)
}

// NextEpoch is a free data retrieval call binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() view returns(uint64)
func (_Staking *StakingCaller) NextEpoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nextEpoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextEpoch is a free data retrieval call binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() view returns(uint64)
func (_Staking *StakingSession) NextEpoch() (uint64, error) {
	return _Staking.Contract.NextEpoch(&_Staking.CallOpts)
}

// NextEpoch is a free data retrieval call binding the contract method 0xaea0e78b.
//
// Solidity: function nextEpoch() view returns(uint64)
func (_Staking *StakingCallerSession) NextEpoch() (uint64, error) {
	return _Staking.Contract.NextEpoch(&_Staking.CallOpts)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address validator) returns()
func (_Staking *StakingTransactor) ActivateValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "activateValidator", validator)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address validator) returns()
func (_Staking *StakingSession) ActivateValidator(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ActivateValidator(&_Staking.TransactOpts, validator)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address validator) returns()
func (_Staking *StakingTransactorSession) ActivateValidator(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ActivateValidator(&_Staking.TransactOpts, validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address account) returns()
func (_Staking *StakingTransactor) AddValidator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addValidator", account)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address account) returns()
func (_Staking *StakingSession) AddValidator(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddValidator(&_Staking.TransactOpts, account)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address account) returns()
func (_Staking *StakingTransactorSession) AddValidator(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddValidator(&_Staking.TransactOpts, account)
}

// ChangeValidatorCommissionRate is a paid mutator transaction binding the contract method 0x14f8649f.
//
// Solidity: function changeValidatorCommissionRate(address validatorAddress, uint16 commissionRate) returns()
func (_Staking *StakingTransactor) ChangeValidatorCommissionRate(opts *bind.TransactOpts, validatorAddress common.Address, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "changeValidatorCommissionRate", validatorAddress, commissionRate)
}

// ChangeValidatorCommissionRate is a paid mutator transaction binding the contract method 0x14f8649f.
//
// Solidity: function changeValidatorCommissionRate(address validatorAddress, uint16 commissionRate) returns()
func (_Staking *StakingSession) ChangeValidatorCommissionRate(validatorAddress common.Address, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.Contract.ChangeValidatorCommissionRate(&_Staking.TransactOpts, validatorAddress, commissionRate)
}

// ChangeValidatorCommissionRate is a paid mutator transaction binding the contract method 0x14f8649f.
//
// Solidity: function changeValidatorCommissionRate(address validatorAddress, uint16 commissionRate) returns()
func (_Staking *StakingTransactorSession) ChangeValidatorCommissionRate(validatorAddress common.Address, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.Contract.ChangeValidatorCommissionRate(&_Staking.TransactOpts, validatorAddress, commissionRate)
}

// ChangeValidatorOwner is a paid mutator transaction binding the contract method 0x0052c9e1.
//
// Solidity: function changeValidatorOwner(address validatorAddress, address newOwner) returns()
func (_Staking *StakingTransactor) ChangeValidatorOwner(opts *bind.TransactOpts, validatorAddress common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "changeValidatorOwner", validatorAddress, newOwner)
}

// ChangeValidatorOwner is a paid mutator transaction binding the contract method 0x0052c9e1.
//
// Solidity: function changeValidatorOwner(address validatorAddress, address newOwner) returns()
func (_Staking *StakingSession) ChangeValidatorOwner(validatorAddress common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ChangeValidatorOwner(&_Staking.TransactOpts, validatorAddress, newOwner)
}

// ChangeValidatorOwner is a paid mutator transaction binding the contract method 0x0052c9e1.
//
// Solidity: function changeValidatorOwner(address validatorAddress, address newOwner) returns()
func (_Staking *StakingTransactorSession) ChangeValidatorOwner(validatorAddress common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ChangeValidatorOwner(&_Staking.TransactOpts, validatorAddress, newOwner)
}

// ClaimDelegatorFee is a paid mutator transaction binding the contract method 0x426594b1.
//
// Solidity: function claimDelegatorFee(address validatorAddress) returns()
func (_Staking *StakingTransactor) ClaimDelegatorFee(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimDelegatorFee", validatorAddress)
}

// ClaimDelegatorFee is a paid mutator transaction binding the contract method 0x426594b1.
//
// Solidity: function claimDelegatorFee(address validatorAddress) returns()
func (_Staking *StakingSession) ClaimDelegatorFee(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ClaimDelegatorFee(&_Staking.TransactOpts, validatorAddress)
}

// ClaimDelegatorFee is a paid mutator transaction binding the contract method 0x426594b1.
//
// Solidity: function claimDelegatorFee(address validatorAddress) returns()
func (_Staking *StakingTransactorSession) ClaimDelegatorFee(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ClaimDelegatorFee(&_Staking.TransactOpts, validatorAddress)
}

// ClaimDelegatorFeeAtEpoch is a paid mutator transaction binding the contract method 0xfe38ebef.
//
// Solidity: function claimDelegatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) returns()
func (_Staking *StakingTransactor) ClaimDelegatorFeeAtEpoch(opts *bind.TransactOpts, validatorAddress common.Address, beforeEpoch uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimDelegatorFeeAtEpoch", validatorAddress, beforeEpoch)
}

// ClaimDelegatorFeeAtEpoch is a paid mutator transaction binding the contract method 0xfe38ebef.
//
// Solidity: function claimDelegatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) returns()
func (_Staking *StakingSession) ClaimDelegatorFeeAtEpoch(validatorAddress common.Address, beforeEpoch uint64) (*types.Transaction, error) {
	return _Staking.Contract.ClaimDelegatorFeeAtEpoch(&_Staking.TransactOpts, validatorAddress, beforeEpoch)
}

// ClaimDelegatorFeeAtEpoch is a paid mutator transaction binding the contract method 0xfe38ebef.
//
// Solidity: function claimDelegatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) returns()
func (_Staking *StakingTransactorSession) ClaimDelegatorFeeAtEpoch(validatorAddress common.Address, beforeEpoch uint64) (*types.Transaction, error) {
	return _Staking.Contract.ClaimDelegatorFeeAtEpoch(&_Staking.TransactOpts, validatorAddress, beforeEpoch)
}

// ClaimValidatorFee is a paid mutator transaction binding the contract method 0xff4794fc.
//
// Solidity: function claimValidatorFee(address validatorAddress) returns()
func (_Staking *StakingTransactor) ClaimValidatorFee(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimValidatorFee", validatorAddress)
}

// ClaimValidatorFee is a paid mutator transaction binding the contract method 0xff4794fc.
//
// Solidity: function claimValidatorFee(address validatorAddress) returns()
func (_Staking *StakingSession) ClaimValidatorFee(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ClaimValidatorFee(&_Staking.TransactOpts, validatorAddress)
}

// ClaimValidatorFee is a paid mutator transaction binding the contract method 0xff4794fc.
//
// Solidity: function claimValidatorFee(address validatorAddress) returns()
func (_Staking *StakingTransactorSession) ClaimValidatorFee(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ClaimValidatorFee(&_Staking.TransactOpts, validatorAddress)
}

// ClaimValidatorFeeAtEpoch is a paid mutator transaction binding the contract method 0xadf2a79c.
//
// Solidity: function claimValidatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) returns()
func (_Staking *StakingTransactor) ClaimValidatorFeeAtEpoch(opts *bind.TransactOpts, validatorAddress common.Address, beforeEpoch uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimValidatorFeeAtEpoch", validatorAddress, beforeEpoch)
}

// ClaimValidatorFeeAtEpoch is a paid mutator transaction binding the contract method 0xadf2a79c.
//
// Solidity: function claimValidatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) returns()
func (_Staking *StakingSession) ClaimValidatorFeeAtEpoch(validatorAddress common.Address, beforeEpoch uint64) (*types.Transaction, error) {
	return _Staking.Contract.ClaimValidatorFeeAtEpoch(&_Staking.TransactOpts, validatorAddress, beforeEpoch)
}

// ClaimValidatorFeeAtEpoch is a paid mutator transaction binding the contract method 0xadf2a79c.
//
// Solidity: function claimValidatorFeeAtEpoch(address validatorAddress, uint64 beforeEpoch) returns()
func (_Staking *StakingTransactorSession) ClaimValidatorFeeAtEpoch(validatorAddress common.Address, beforeEpoch uint64) (*types.Transaction, error) {
	return _Staking.Contract.ClaimValidatorFeeAtEpoch(&_Staking.TransactOpts, validatorAddress, beforeEpoch)
}

// Ctor is a paid mutator transaction binding the contract method 0xc15f28c4.
//
// Solidity: function ctor(address[] validators, uint256[] initialStakes, uint16 commissionRate) returns()
func (_Staking *StakingTransactor) Ctor(opts *bind.TransactOpts, validators []common.Address, initialStakes []*big.Int, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "ctor", validators, initialStakes, commissionRate)
}

// Ctor is a paid mutator transaction binding the contract method 0xc15f28c4.
//
// Solidity: function ctor(address[] validators, uint256[] initialStakes, uint16 commissionRate) returns()
func (_Staking *StakingSession) Ctor(validators []common.Address, initialStakes []*big.Int, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.Contract.Ctor(&_Staking.TransactOpts, validators, initialStakes, commissionRate)
}

// Ctor is a paid mutator transaction binding the contract method 0xc15f28c4.
//
// Solidity: function ctor(address[] validators, uint256[] initialStakes, uint16 commissionRate) returns()
func (_Staking *StakingTransactorSession) Ctor(validators []common.Address, initialStakes []*big.Int, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.Contract.Ctor(&_Staking.TransactOpts, validators, initialStakes, commissionRate)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validatorAddress) payable returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", validatorAddress)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validatorAddress) payable returns()
func (_Staking *StakingSession) Delegate(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validatorAddress)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validatorAddress) payable returns()
func (_Staking *StakingTransactorSession) Delegate(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validatorAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address validatorAddress) payable returns()
func (_Staking *StakingTransactor) Deposit(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "deposit", validatorAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address validatorAddress) payable returns()
func (_Staking *StakingSession) Deposit(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, validatorAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address validatorAddress) payable returns()
func (_Staking *StakingTransactorSession) Deposit(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, validatorAddress)
}

// DisableValidator is a paid mutator transaction binding the contract method 0x1fe97684.
//
// Solidity: function disableValidator(address validator) returns()
func (_Staking *StakingTransactor) DisableValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "disableValidator", validator)
}

// DisableValidator is a paid mutator transaction binding the contract method 0x1fe97684.
//
// Solidity: function disableValidator(address validator) returns()
func (_Staking *StakingSession) DisableValidator(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DisableValidator(&_Staking.TransactOpts, validator)
}

// DisableValidator is a paid mutator transaction binding the contract method 0x1fe97684.
//
// Solidity: function disableValidator(address validator) returns()
func (_Staking *StakingTransactorSession) DisableValidator(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DisableValidator(&_Staking.TransactOpts, validator)
}

// RedelegateDelegatorFee is a paid mutator transaction binding the contract method 0x8ecb3fc9.
//
// Solidity: function redelegateDelegatorFee(address validator) returns()
func (_Staking *StakingTransactor) RedelegateDelegatorFee(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "redelegateDelegatorFee", validator)
}

// RedelegateDelegatorFee is a paid mutator transaction binding the contract method 0x8ecb3fc9.
//
// Solidity: function redelegateDelegatorFee(address validator) returns()
func (_Staking *StakingSession) RedelegateDelegatorFee(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RedelegateDelegatorFee(&_Staking.TransactOpts, validator)
}

// RedelegateDelegatorFee is a paid mutator transaction binding the contract method 0x8ecb3fc9.
//
// Solidity: function redelegateDelegatorFee(address validator) returns()
func (_Staking *StakingTransactorSession) RedelegateDelegatorFee(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RedelegateDelegatorFee(&_Staking.TransactOpts, validator)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x61cadbf4.
//
// Solidity: function registerValidator(address validatorAddress, uint16 commissionRate) payable returns()
func (_Staking *StakingTransactor) RegisterValidator(opts *bind.TransactOpts, validatorAddress common.Address, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "registerValidator", validatorAddress, commissionRate)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x61cadbf4.
//
// Solidity: function registerValidator(address validatorAddress, uint16 commissionRate) payable returns()
func (_Staking *StakingSession) RegisterValidator(validatorAddress common.Address, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.Contract.RegisterValidator(&_Staking.TransactOpts, validatorAddress, commissionRate)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x61cadbf4.
//
// Solidity: function registerValidator(address validatorAddress, uint16 commissionRate) payable returns()
func (_Staking *StakingTransactorSession) RegisterValidator(validatorAddress common.Address, commissionRate uint16) (*types.Transaction, error) {
	return _Staking.Contract.RegisterValidator(&_Staking.TransactOpts, validatorAddress, commissionRate)
}

// ReleaseValidatorFromJail is a paid mutator transaction binding the contract method 0x73a3dda6.
//
// Solidity: function releaseValidatorFromJail(address validatorAddress) returns()
func (_Staking *StakingTransactor) ReleaseValidatorFromJail(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "releaseValidatorFromJail", validatorAddress)
}

// ReleaseValidatorFromJail is a paid mutator transaction binding the contract method 0x73a3dda6.
//
// Solidity: function releaseValidatorFromJail(address validatorAddress) returns()
func (_Staking *StakingSession) ReleaseValidatorFromJail(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ReleaseValidatorFromJail(&_Staking.TransactOpts, validatorAddress)
}

// ReleaseValidatorFromJail is a paid mutator transaction binding the contract method 0x73a3dda6.
//
// Solidity: function releaseValidatorFromJail(address validatorAddress) returns()
func (_Staking *StakingTransactorSession) ReleaseValidatorFromJail(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ReleaseValidatorFromJail(&_Staking.TransactOpts, validatorAddress)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address account) returns()
func (_Staking *StakingTransactor) RemoveValidator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "removeValidator", account)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address account) returns()
func (_Staking *StakingSession) RemoveValidator(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveValidator(&_Staking.TransactOpts, account)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address account) returns()
func (_Staking *StakingTransactorSession) RemoveValidator(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveValidator(&_Staking.TransactOpts, account)
}

// SetEcosystemReward is a paid mutator transaction binding the contract method 0xb92f0e80.
//
// Solidity: function setEcosystemReward(uint256 _amount) returns()
func (_Staking *StakingTransactor) SetEcosystemReward(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setEcosystemReward", _amount)
}

// SetEcosystemReward is a paid mutator transaction binding the contract method 0xb92f0e80.
//
// Solidity: function setEcosystemReward(uint256 _amount) returns()
func (_Staking *StakingSession) SetEcosystemReward(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetEcosystemReward(&_Staking.TransactOpts, _amount)
}

// SetEcosystemReward is a paid mutator transaction binding the contract method 0xb92f0e80.
//
// Solidity: function setEcosystemReward(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) SetEcosystemReward(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetEcosystemReward(&_Staking.TransactOpts, _amount)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validatorAddress) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", validatorAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validatorAddress) returns()
func (_Staking *StakingSession) Slash(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, validatorAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validatorAddress) returns()
func (_Staking *StakingTransactorSession) Slash(validatorAddress common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, validatorAddress)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns()
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns()
func (_Staking *StakingSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns()
func (_Staking *StakingTransactorSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, validatorAddress, amount)
}

// StakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Staking contract.
type StakingClaimedIterator struct {
	Event *StakingClaimed // Event containing the contract specifics and raw log

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
func (it *StakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimed)
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
		it.Event = new(StakingClaimed)
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
func (it *StakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimed represents a Claimed event raised by the Staking contract.
type StakingClaimed struct {
	Validator common.Address
	Staker    common.Address
	Amount    *big.Int
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xb22dec804803f8b1c5333f626cdbfdfb1bd629f1e1bb45dcfb22b5f74ed46b1c.
//
// Solidity: event Claimed(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) FilterClaimed(opts *bind.FilterOpts, validator []common.Address, staker []common.Address) (*StakingClaimedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Claimed", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingClaimedIterator{contract: _Staking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xb22dec804803f8b1c5333f626cdbfdfb1bd629f1e1bb45dcfb22b5f74ed46b1c.
//
// Solidity: event Claimed(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *StakingClaimed, validator []common.Address, staker []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Claimed", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimed)
				if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xb22dec804803f8b1c5333f626cdbfdfb1bd629f1e1bb45dcfb22b5f74ed46b1c.
//
// Solidity: event Claimed(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) ParseClaimed(log types.Log) (*StakingClaimed, error) {
	event := new(StakingClaimed)
	if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Staking contract.
type StakingDelegatedIterator struct {
	Event *StakingDelegated // Event containing the contract specifics and raw log

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
func (it *StakingDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegated)
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
		it.Event = new(StakingDelegated)
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
func (it *StakingDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegated represents a Delegated event raised by the Staking contract.
type StakingDelegated struct {
	Validator common.Address
	Staker    common.Address
	Amount    *big.Int
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x30bcda2f188b532c7644e632473e83a6fb3c5c79717650d0ac790d141bb1b177.
//
// Solidity: event Delegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) FilterDelegated(opts *bind.FilterOpts, validator []common.Address, staker []common.Address) (*StakingDelegatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Delegated", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegatedIterator{contract: _Staking.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x30bcda2f188b532c7644e632473e83a6fb3c5c79717650d0ac790d141bb1b177.
//
// Solidity: event Delegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakingDelegated, validator []common.Address, staker []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Delegated", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegated)
				if err := _Staking.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x30bcda2f188b532c7644e632473e83a6fb3c5c79717650d0ac790d141bb1b177.
//
// Solidity: event Delegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) ParseDelegated(log types.Log) (*StakingDelegated, error) {
	event := new(StakingDelegated)
	if err := _Staking.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the Staking contract.
type StakingRedelegatedIterator struct {
	Event *StakingRedelegated // Event containing the contract specifics and raw log

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
func (it *StakingRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRedelegated)
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
		it.Event = new(StakingRedelegated)
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
func (it *StakingRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRedelegated represents a Redelegated event raised by the Staking contract.
type StakingRedelegated struct {
	Validator common.Address
	Staker    common.Address
	Amount    *big.Int
	Dust      *big.Int
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0xa82f74002b6639f6cfc2cfd4f3ade1998108eda0f484d9064e3098c211e81d6e.
//
// Solidity: event Redelegated(address indexed validator, address indexed staker, uint256 amount, uint256 dust, uint64 epoch)
func (_Staking *StakingFilterer) FilterRedelegated(opts *bind.FilterOpts, validator []common.Address, staker []common.Address) (*StakingRedelegatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Redelegated", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingRedelegatedIterator{contract: _Staking.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0xa82f74002b6639f6cfc2cfd4f3ade1998108eda0f484d9064e3098c211e81d6e.
//
// Solidity: event Redelegated(address indexed validator, address indexed staker, uint256 amount, uint256 dust, uint64 epoch)
func (_Staking *StakingFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *StakingRedelegated, validator []common.Address, staker []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Redelegated", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRedelegated)
				if err := _Staking.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0xa82f74002b6639f6cfc2cfd4f3ade1998108eda0f484d9064e3098c211e81d6e.
//
// Solidity: event Redelegated(address indexed validator, address indexed staker, uint256 amount, uint256 dust, uint64 epoch)
func (_Staking *StakingFilterer) ParseRedelegated(log types.Log) (*StakingRedelegated, error) {
	event := new(StakingRedelegated)
	if err := _Staking.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Staking contract.
type StakingUndelegatedIterator struct {
	Event *StakingUndelegated // Event containing the contract specifics and raw log

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
func (it *StakingUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUndelegated)
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
		it.Event = new(StakingUndelegated)
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
func (it *StakingUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUndelegated represents a Undelegated event raised by the Staking contract.
type StakingUndelegated struct {
	Validator common.Address
	Staker    common.Address
	Amount    *big.Int
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0xa410e32157a44414a502bb47d775234de1aa7da123f5adfe426898f1601883fd.
//
// Solidity: event Undelegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) FilterUndelegated(opts *bind.FilterOpts, validator []common.Address, staker []common.Address) (*StakingUndelegatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Undelegated", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingUndelegatedIterator{contract: _Staking.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0xa410e32157a44414a502bb47d775234de1aa7da123f5adfe426898f1601883fd.
//
// Solidity: event Undelegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *StakingUndelegated, validator []common.Address, staker []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Undelegated", validatorRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUndelegated)
				if err := _Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0xa410e32157a44414a502bb47d775234de1aa7da123f5adfe426898f1601883fd.
//
// Solidity: event Undelegated(address indexed validator, address indexed staker, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) ParseUndelegated(log types.Log) (*StakingUndelegated, error) {
	event := new(StakingUndelegated)
	if err := _Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Staking contract.
type StakingValidatorAddedIterator struct {
	Event *StakingValidatorAdded // Event containing the contract specifics and raw log

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
func (it *StakingValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorAdded)
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
		it.Event = new(StakingValidatorAdded)
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
func (it *StakingValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorAdded represents a ValidatorAdded event raised by the Staking contract.
type StakingValidatorAdded struct {
	Validator      common.Address
	Owner          common.Address
	Status         uint8
	CommissionRate uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d.
//
// Solidity: event ValidatorAdded(address indexed validator, address owner, uint8 status, uint16 commissionRate)
func (_Staking *StakingFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorAddedIterator{contract: _Staking.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d.
//
// Solidity: event ValidatorAdded(address indexed validator, address owner, uint8 status, uint16 commissionRate)
func (_Staking *StakingFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *StakingValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorAdded)
				if err := _Staking.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0x42449fd19d367b0177da9082fe6da7d4da41af7573e3a3c1750ecffeffe26f9d.
//
// Solidity: event ValidatorAdded(address indexed validator, address owner, uint8 status, uint16 commissionRate)
func (_Staking *StakingFilterer) ParseValidatorAdded(log types.Log) (*StakingValidatorAdded, error) {
	event := new(StakingValidatorAdded)
	if err := _Staking.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorDepositedIterator is returned from FilterValidatorDeposited and is used to iterate over the raw logs and unpacked data for ValidatorDeposited events raised by the Staking contract.
type StakingValidatorDepositedIterator struct {
	Event *StakingValidatorDeposited // Event containing the contract specifics and raw log

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
func (it *StakingValidatorDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorDeposited)
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
		it.Event = new(StakingValidatorDeposited)
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
func (it *StakingValidatorDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorDeposited represents a ValidatorDeposited event raised by the Staking contract.
type StakingValidatorDeposited struct {
	Validator common.Address
	Amount    *big.Int
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposited is a free log retrieval operation binding the contract event 0x9fce3fb7dc05df2879421f9d9c06784dad3d39aba44c0c1ee0b9d094e7655005.
//
// Solidity: event ValidatorDeposited(address indexed validator, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) FilterValidatorDeposited(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorDepositedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorDeposited", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorDepositedIterator{contract: _Staking.contract, event: "ValidatorDeposited", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposited is a free log subscription operation binding the contract event 0x9fce3fb7dc05df2879421f9d9c06784dad3d39aba44c0c1ee0b9d094e7655005.
//
// Solidity: event ValidatorDeposited(address indexed validator, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) WatchValidatorDeposited(opts *bind.WatchOpts, sink chan<- *StakingValidatorDeposited, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorDeposited", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorDeposited)
				if err := _Staking.contract.UnpackLog(event, "ValidatorDeposited", log); err != nil {
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

// ParseValidatorDeposited is a log parse operation binding the contract event 0x9fce3fb7dc05df2879421f9d9c06784dad3d39aba44c0c1ee0b9d094e7655005.
//
// Solidity: event ValidatorDeposited(address indexed validator, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) ParseValidatorDeposited(log types.Log) (*StakingValidatorDeposited, error) {
	event := new(StakingValidatorDeposited)
	if err := _Staking.contract.UnpackLog(event, "ValidatorDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Staking contract.
type StakingValidatorJailedIterator struct {
	Event *StakingValidatorJailed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorJailed)
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
		it.Event = new(StakingValidatorJailed)
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
func (it *StakingValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorJailed represents a ValidatorJailed event raised by the Staking contract.
type StakingValidatorJailed struct {
	Validator common.Address
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0x6acc1079130362b805ce4a8874346d4ee8893d55f223e378ddc73031ce427eab.
//
// Solidity: event ValidatorJailed(address indexed validator, uint64 epoch)
func (_Staking *StakingFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorJailedIterator{contract: _Staking.contract, event: "ValidatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0x6acc1079130362b805ce4a8874346d4ee8893d55f223e378ddc73031ce427eab.
//
// Solidity: event ValidatorJailed(address indexed validator, uint64 epoch)
func (_Staking *StakingFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *StakingValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorJailed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0x6acc1079130362b805ce4a8874346d4ee8893d55f223e378ddc73031ce427eab.
//
// Solidity: event ValidatorJailed(address indexed validator, uint64 epoch)
func (_Staking *StakingFilterer) ParseValidatorJailed(log types.Log) (*StakingValidatorJailed, error) {
	event := new(StakingValidatorJailed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorModifiedIterator is returned from FilterValidatorModified and is used to iterate over the raw logs and unpacked data for ValidatorModified events raised by the Staking contract.
type StakingValidatorModifiedIterator struct {
	Event *StakingValidatorModified // Event containing the contract specifics and raw log

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
func (it *StakingValidatorModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorModified)
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
		it.Event = new(StakingValidatorModified)
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
func (it *StakingValidatorModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorModified represents a ValidatorModified event raised by the Staking contract.
type StakingValidatorModified struct {
	Validator      common.Address
	Owner          common.Address
	Status         uint8
	CommissionRate uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorModified is a free log retrieval operation binding the contract event 0xc00107e0d011ac7c8e4dfa18e3dd3623ff151f8bfcc3821cd39bc114bd6504d9.
//
// Solidity: event ValidatorModified(address indexed validator, address owner, uint8 status, uint16 commissionRate)
func (_Staking *StakingFilterer) FilterValidatorModified(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorModifiedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorModified", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorModifiedIterator{contract: _Staking.contract, event: "ValidatorModified", logs: logs, sub: sub}, nil
}

// WatchValidatorModified is a free log subscription operation binding the contract event 0xc00107e0d011ac7c8e4dfa18e3dd3623ff151f8bfcc3821cd39bc114bd6504d9.
//
// Solidity: event ValidatorModified(address indexed validator, address owner, uint8 status, uint16 commissionRate)
func (_Staking *StakingFilterer) WatchValidatorModified(opts *bind.WatchOpts, sink chan<- *StakingValidatorModified, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorModified", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorModified)
				if err := _Staking.contract.UnpackLog(event, "ValidatorModified", log); err != nil {
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

// ParseValidatorModified is a log parse operation binding the contract event 0xc00107e0d011ac7c8e4dfa18e3dd3623ff151f8bfcc3821cd39bc114bd6504d9.
//
// Solidity: event ValidatorModified(address indexed validator, address owner, uint8 status, uint16 commissionRate)
func (_Staking *StakingFilterer) ParseValidatorModified(log types.Log) (*StakingValidatorModified, error) {
	event := new(StakingValidatorModified)
	if err := _Staking.contract.UnpackLog(event, "ValidatorModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorOwnerClaimedIterator is returned from FilterValidatorOwnerClaimed and is used to iterate over the raw logs and unpacked data for ValidatorOwnerClaimed events raised by the Staking contract.
type StakingValidatorOwnerClaimedIterator struct {
	Event *StakingValidatorOwnerClaimed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorOwnerClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorOwnerClaimed)
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
		it.Event = new(StakingValidatorOwnerClaimed)
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
func (it *StakingValidatorOwnerClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorOwnerClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorOwnerClaimed represents a ValidatorOwnerClaimed event raised by the Staking contract.
type StakingValidatorOwnerClaimed struct {
	Validator common.Address
	Amount    *big.Int
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorOwnerClaimed is a free log retrieval operation binding the contract event 0xb947d7b49cedaf132fd7a9592099c21170864455405d51b48250324415324100.
//
// Solidity: event ValidatorOwnerClaimed(address indexed validator, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) FilterValidatorOwnerClaimed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorOwnerClaimedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorOwnerClaimed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorOwnerClaimedIterator{contract: _Staking.contract, event: "ValidatorOwnerClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorOwnerClaimed is a free log subscription operation binding the contract event 0xb947d7b49cedaf132fd7a9592099c21170864455405d51b48250324415324100.
//
// Solidity: event ValidatorOwnerClaimed(address indexed validator, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) WatchValidatorOwnerClaimed(opts *bind.WatchOpts, sink chan<- *StakingValidatorOwnerClaimed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorOwnerClaimed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorOwnerClaimed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorOwnerClaimed", log); err != nil {
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

// ParseValidatorOwnerClaimed is a log parse operation binding the contract event 0xb947d7b49cedaf132fd7a9592099c21170864455405d51b48250324415324100.
//
// Solidity: event ValidatorOwnerClaimed(address indexed validator, uint256 amount, uint64 epoch)
func (_Staking *StakingFilterer) ParseValidatorOwnerClaimed(log types.Log) (*StakingValidatorOwnerClaimed, error) {
	event := new(StakingValidatorOwnerClaimed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorOwnerClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorReleasedIterator is returned from FilterValidatorReleased and is used to iterate over the raw logs and unpacked data for ValidatorReleased events raised by the Staking contract.
type StakingValidatorReleasedIterator struct {
	Event *StakingValidatorReleased // Event containing the contract specifics and raw log

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
func (it *StakingValidatorReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorReleased)
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
		it.Event = new(StakingValidatorReleased)
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
func (it *StakingValidatorReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorReleased represents a ValidatorReleased event raised by the Staking contract.
type StakingValidatorReleased struct {
	Validator common.Address
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorReleased is a free log retrieval operation binding the contract event 0x7bd9ec00705b7640a93da0fcf84dc0bfce1b2c9a5146d548ab1dabd23322c9ac.
//
// Solidity: event ValidatorReleased(address indexed validator, uint64 epoch)
func (_Staking *StakingFilterer) FilterValidatorReleased(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorReleasedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorReleased", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorReleasedIterator{contract: _Staking.contract, event: "ValidatorReleased", logs: logs, sub: sub}, nil
}

// WatchValidatorReleased is a free log subscription operation binding the contract event 0x7bd9ec00705b7640a93da0fcf84dc0bfce1b2c9a5146d548ab1dabd23322c9ac.
//
// Solidity: event ValidatorReleased(address indexed validator, uint64 epoch)
func (_Staking *StakingFilterer) WatchValidatorReleased(opts *bind.WatchOpts, sink chan<- *StakingValidatorReleased, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorReleased", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorReleased)
				if err := _Staking.contract.UnpackLog(event, "ValidatorReleased", log); err != nil {
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

// ParseValidatorReleased is a log parse operation binding the contract event 0x7bd9ec00705b7640a93da0fcf84dc0bfce1b2c9a5146d548ab1dabd23322c9ac.
//
// Solidity: event ValidatorReleased(address indexed validator, uint64 epoch)
func (_Staking *StakingFilterer) ParseValidatorReleased(log types.Log) (*StakingValidatorReleased, error) {
	event := new(StakingValidatorReleased)
	if err := _Staking.contract.UnpackLog(event, "ValidatorReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Staking contract.
type StakingValidatorRemovedIterator struct {
	Event *StakingValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *StakingValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorRemoved)
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
		it.Event = new(StakingValidatorRemoved)
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
func (it *StakingValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorRemoved represents a ValidatorRemoved event raised by the Staking contract.
type StakingValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_Staking *StakingFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorRemovedIterator{contract: _Staking.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_Staking *StakingFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *StakingValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorRemoved)
				if err := _Staking.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_Staking *StakingFilterer) ParseValidatorRemoved(log types.Log) (*StakingValidatorRemoved, error) {
	event := new(StakingValidatorRemoved)
	if err := _Staking.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Staking contract.
type StakingValidatorSlashedIterator struct {
	Event *StakingValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorSlashed)
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
		it.Event = new(StakingValidatorSlashed)
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
func (it *StakingValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorSlashed represents a ValidatorSlashed event raised by the Staking contract.
type StakingValidatorSlashed struct {
	Validator common.Address
	Slashes   uint32
	Epoch     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0xf30f0392d5346908bf200bc0924d0c9860e22d46fc606e979a2e14fff24c62b9.
//
// Solidity: event ValidatorSlashed(address indexed validator, uint32 slashes, uint64 epoch)
func (_Staking *StakingFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorSlashedIterator{contract: _Staking.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0xf30f0392d5346908bf200bc0924d0c9860e22d46fc606e979a2e14fff24c62b9.
//
// Solidity: event ValidatorSlashed(address indexed validator, uint32 slashes, uint64 epoch)
func (_Staking *StakingFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *StakingValidatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorSlashed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0xf30f0392d5346908bf200bc0924d0c9860e22d46fc606e979a2e14fff24c62b9.
//
// Solidity: event ValidatorSlashed(address indexed validator, uint32 slashes, uint64 epoch)
func (_Staking *StakingFilterer) ParseValidatorSlashed(log types.Log) (*StakingValidatorSlashed, error) {
	event := new(StakingValidatorSlashed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
