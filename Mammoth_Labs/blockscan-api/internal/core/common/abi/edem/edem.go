// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package edem

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

// MarketTypesMarketPlace is an auto generated low-level Go binding around an user-defined struct.
type MarketTypesMarketPlace struct {
	IsDealer          bool
	EdemSigner        common.Address
	Collection        common.Address
	TokenId           *big.Int
	NftAmount         *big.Int
	Price             *big.Int
	ProtocolAddress   common.Address
	TradeTokenAddress common.Address
	Nonce             *big.Int
	StartTime         *big.Int
	EndTime           *big.Int
	V                 uint8
	R                 [32]byte
	S                 [32]byte
}

// MarketTypesUser is an auto generated low-level Go binding around an user-defined struct.
type MarketTypesUser struct {
	IsDealer     bool
	TakerAddress common.Address
	Price        *big.Int
	TokenId      *big.Int
}

// EdemMetaData contains all meta data concerning the Edem contract.
var EdemMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinNonce\",\"type\":\"uint256\"}],\"name\":\"CancelAllOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"CancelMultipleOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolExecutionManager\",\"type\":\"address\"}],\"name\":\"NewProtocolExecutionManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocoleFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyManager\",\"type\":\"address\"}],\"name\":\"NewRoyaltyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tradeTokenManager\",\"type\":\"address\"}],\"name\":\"NewTradeTokenManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferManager\",\"type\":\"address\"}],\"name\":\"NewTransferManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RoyaltyPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"marketPlaceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userProposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketSeller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"UserProposer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"marketPlaceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userSeller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketProposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"UserSeller\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INTERFACE_ID_ERC1155\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INTERFACE_ID_ERC721\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNonce\",\"type\":\"uint256\"}],\"name\":\"cancelAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelSellAndSuggest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderNonces\",\"type\":\"uint256\"}],\"name\":\"isUserOrderNonceExecutedOrCancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketTypes.User\",\"name\":\"userProposer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"edemSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structMarketTypes.MarketPlace\",\"name\":\"marketSeller\",\"type\":\"tuple\"}],\"name\":\"proposerPay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketTypes.User\",\"name\":\"userProposer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"edemSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structMarketTypes.MarketPlace\",\"name\":\"marketSeller\",\"type\":\"tuple\"}],\"name\":\"proposerPayETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketTypes.User\",\"name\":\"userProposer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"edemSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structMarketTypes.MarketPlace\",\"name\":\"marketSeller\",\"type\":\"tuple\"}],\"name\":\"proposerPayETHAndWETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocoleFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structMarketTypes.User\",\"name\":\"userSeller\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isDealer\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"edemSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structMarketTypes.MarketPlace\",\"name\":\"marketProposer\",\"type\":\"tuple\"}],\"name\":\"suggestApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EdemABI is the input ABI used to generate the binding from.
// Deprecated: Use EdemMetaData.ABI instead.
var EdemABI = EdemMetaData.ABI

// Edem is an auto generated Go binding around an Ethereum contract.
type Edem struct {
	EdemCaller     // Read-only binding to the contract
	EdemTransactor // Write-only binding to the contract
	EdemFilterer   // Log filterer for contract events
}

// EdemCaller is an auto generated read-only Go binding around an Ethereum contract.
type EdemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EdemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EdemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EdemSession struct {
	Contract     *Edem             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EdemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EdemCallerSession struct {
	Contract *EdemCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EdemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EdemTransactorSession struct {
	Contract     *EdemTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EdemRaw is an auto generated low-level Go binding around an Ethereum contract.
type EdemRaw struct {
	Contract *Edem // Generic contract binding to access the raw methods on
}

// EdemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EdemCallerRaw struct {
	Contract *EdemCaller // Generic read-only contract binding to access the raw methods on
}

// EdemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EdemTransactorRaw struct {
	Contract *EdemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEdem creates a new instance of Edem, bound to a specific deployed contract.
func NewEdem(address common.Address, backend bind.ContractBackend) (*Edem, error) {
	contract, err := bindEdem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Edem{EdemCaller: EdemCaller{contract: contract}, EdemTransactor: EdemTransactor{contract: contract}, EdemFilterer: EdemFilterer{contract: contract}}, nil
}

// NewEdemCaller creates a new read-only instance of Edem, bound to a specific deployed contract.
func NewEdemCaller(address common.Address, caller bind.ContractCaller) (*EdemCaller, error) {
	contract, err := bindEdem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EdemCaller{contract: contract}, nil
}

// NewEdemTransactor creates a new write-only instance of Edem, bound to a specific deployed contract.
func NewEdemTransactor(address common.Address, transactor bind.ContractTransactor) (*EdemTransactor, error) {
	contract, err := bindEdem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EdemTransactor{contract: contract}, nil
}

// NewEdemFilterer creates a new log filterer instance of Edem, bound to a specific deployed contract.
func NewEdemFilterer(address common.Address, filterer bind.ContractFilterer) (*EdemFilterer, error) {
	contract, err := bindEdem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EdemFilterer{contract: contract}, nil
}

// bindEdem binds a generic wrapper to an already deployed contract.
func bindEdem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EdemMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Edem *EdemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Edem.Contract.EdemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Edem *EdemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Edem.Contract.EdemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Edem *EdemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Edem.Contract.EdemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Edem *EdemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Edem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Edem *EdemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Edem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Edem *EdemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Edem.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Edem *EdemCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Edem.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Edem *EdemSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Edem.Contract.DOMAINSEPARATOR(&_Edem.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Edem *EdemCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Edem.Contract.DOMAINSEPARATOR(&_Edem.CallOpts)
}

// INTERFACEIDERC1155 is a free data retrieval call binding the contract method 0x33bf6156.
//
// Solidity: function INTERFACE_ID_ERC1155() view returns(bytes4)
func (_Edem *EdemCaller) INTERFACEIDERC1155(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Edem.contract.Call(opts, &out, "INTERFACE_ID_ERC1155")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// INTERFACEIDERC1155 is a free data retrieval call binding the contract method 0x33bf6156.
//
// Solidity: function INTERFACE_ID_ERC1155() view returns(bytes4)
func (_Edem *EdemSession) INTERFACEIDERC1155() ([4]byte, error) {
	return _Edem.Contract.INTERFACEIDERC1155(&_Edem.CallOpts)
}

// INTERFACEIDERC1155 is a free data retrieval call binding the contract method 0x33bf6156.
//
// Solidity: function INTERFACE_ID_ERC1155() view returns(bytes4)
func (_Edem *EdemCallerSession) INTERFACEIDERC1155() ([4]byte, error) {
	return _Edem.Contract.INTERFACEIDERC1155(&_Edem.CallOpts)
}

// INTERFACEIDERC721 is a free data retrieval call binding the contract method 0xbc6bc0cd.
//
// Solidity: function INTERFACE_ID_ERC721() view returns(bytes4)
func (_Edem *EdemCaller) INTERFACEIDERC721(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Edem.contract.Call(opts, &out, "INTERFACE_ID_ERC721")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// INTERFACEIDERC721 is a free data retrieval call binding the contract method 0xbc6bc0cd.
//
// Solidity: function INTERFACE_ID_ERC721() view returns(bytes4)
func (_Edem *EdemSession) INTERFACEIDERC721() ([4]byte, error) {
	return _Edem.Contract.INTERFACEIDERC721(&_Edem.CallOpts)
}

// INTERFACEIDERC721 is a free data retrieval call binding the contract method 0xbc6bc0cd.
//
// Solidity: function INTERFACE_ID_ERC721() view returns(bytes4)
func (_Edem *EdemCallerSession) INTERFACEIDERC721() ([4]byte, error) {
	return _Edem.Contract.INTERFACEIDERC721(&_Edem.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Edem *EdemCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Edem.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Edem *EdemSession) WETH() (common.Address, error) {
	return _Edem.Contract.WETH(&_Edem.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Edem *EdemCallerSession) WETH() (common.Address, error) {
	return _Edem.Contract.WETH(&_Edem.CallOpts)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonces) view returns(bool)
func (_Edem *EdemCaller) IsUserOrderNonceExecutedOrCancelled(opts *bind.CallOpts, user common.Address, orderNonces *big.Int) (bool, error) {
	var out []interface{}
	err := _Edem.contract.Call(opts, &out, "isUserOrderNonceExecutedOrCancelled", user, orderNonces)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonces) view returns(bool)
func (_Edem *EdemSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonces *big.Int) (bool, error) {
	return _Edem.Contract.IsUserOrderNonceExecutedOrCancelled(&_Edem.CallOpts, user, orderNonces)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonces) view returns(bool)
func (_Edem *EdemCallerSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonces *big.Int) (bool, error) {
	return _Edem.Contract.IsUserOrderNonceExecutedOrCancelled(&_Edem.CallOpts, user, orderNonces)
}

// ProtocoleFeeRecipient is a free data retrieval call binding the contract method 0x8b655ee0.
//
// Solidity: function protocoleFeeRecipient() view returns(address)
func (_Edem *EdemCaller) ProtocoleFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Edem.contract.Call(opts, &out, "protocoleFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocoleFeeRecipient is a free data retrieval call binding the contract method 0x8b655ee0.
//
// Solidity: function protocoleFeeRecipient() view returns(address)
func (_Edem *EdemSession) ProtocoleFeeRecipient() (common.Address, error) {
	return _Edem.Contract.ProtocoleFeeRecipient(&_Edem.CallOpts)
}

// ProtocoleFeeRecipient is a free data retrieval call binding the contract method 0x8b655ee0.
//
// Solidity: function protocoleFeeRecipient() view returns(address)
func (_Edem *EdemCallerSession) ProtocoleFeeRecipient() (common.Address, error) {
	return _Edem.Contract.ProtocoleFeeRecipient(&_Edem.CallOpts)
}

// CancelAll is a paid mutator transaction binding the contract method 0x146fce32.
//
// Solidity: function cancelAll(uint256 minNonce) returns()
func (_Edem *EdemTransactor) CancelAll(opts *bind.TransactOpts, minNonce *big.Int) (*types.Transaction, error) {
	return _Edem.contract.Transact(opts, "cancelAll", minNonce)
}

// CancelAll is a paid mutator transaction binding the contract method 0x146fce32.
//
// Solidity: function cancelAll(uint256 minNonce) returns()
func (_Edem *EdemSession) CancelAll(minNonce *big.Int) (*types.Transaction, error) {
	return _Edem.Contract.CancelAll(&_Edem.TransactOpts, minNonce)
}

// CancelAll is a paid mutator transaction binding the contract method 0x146fce32.
//
// Solidity: function cancelAll(uint256 minNonce) returns()
func (_Edem *EdemTransactorSession) CancelAll(minNonce *big.Int) (*types.Transaction, error) {
	return _Edem.Contract.CancelAll(&_Edem.TransactOpts, minNonce)
}

// CancelSellAndSuggest is a paid mutator transaction binding the contract method 0xaa747d3a.
//
// Solidity: function cancelSellAndSuggest(uint256[] orderNonces) returns()
func (_Edem *EdemTransactor) CancelSellAndSuggest(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _Edem.contract.Transact(opts, "cancelSellAndSuggest", orderNonces)
}

// CancelSellAndSuggest is a paid mutator transaction binding the contract method 0xaa747d3a.
//
// Solidity: function cancelSellAndSuggest(uint256[] orderNonces) returns()
func (_Edem *EdemSession) CancelSellAndSuggest(orderNonces []*big.Int) (*types.Transaction, error) {
	return _Edem.Contract.CancelSellAndSuggest(&_Edem.TransactOpts, orderNonces)
}

// CancelSellAndSuggest is a paid mutator transaction binding the contract method 0xaa747d3a.
//
// Solidity: function cancelSellAndSuggest(uint256[] orderNonces) returns()
func (_Edem *EdemTransactorSession) CancelSellAndSuggest(orderNonces []*big.Int) (*types.Transaction, error) {
	return _Edem.Contract.CancelSellAndSuggest(&_Edem.TransactOpts, orderNonces)
}

// ProposerPay is a paid mutator transaction binding the contract method 0x93a4be75.
//
// Solidity: function proposerPay((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) returns()
func (_Edem *EdemTransactor) ProposerPay(opts *bind.TransactOpts, userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.contract.Transact(opts, "proposerPay", userProposer, marketSeller)
}

// ProposerPay is a paid mutator transaction binding the contract method 0x93a4be75.
//
// Solidity: function proposerPay((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) returns()
func (_Edem *EdemSession) ProposerPay(userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.ProposerPay(&_Edem.TransactOpts, userProposer, marketSeller)
}

// ProposerPay is a paid mutator transaction binding the contract method 0x93a4be75.
//
// Solidity: function proposerPay((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) returns()
func (_Edem *EdemTransactorSession) ProposerPay(userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.ProposerPay(&_Edem.TransactOpts, userProposer, marketSeller)
}

// ProposerPayETH is a paid mutator transaction binding the contract method 0xf4195742.
//
// Solidity: function proposerPayETH((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) payable returns()
func (_Edem *EdemTransactor) ProposerPayETH(opts *bind.TransactOpts, userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.contract.Transact(opts, "proposerPayETH", userProposer, marketSeller)
}

// ProposerPayETH is a paid mutator transaction binding the contract method 0xf4195742.
//
// Solidity: function proposerPayETH((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) payable returns()
func (_Edem *EdemSession) ProposerPayETH(userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.ProposerPayETH(&_Edem.TransactOpts, userProposer, marketSeller)
}

// ProposerPayETH is a paid mutator transaction binding the contract method 0xf4195742.
//
// Solidity: function proposerPayETH((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) payable returns()
func (_Edem *EdemTransactorSession) ProposerPayETH(userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.ProposerPayETH(&_Edem.TransactOpts, userProposer, marketSeller)
}

// ProposerPayETHAndWETH is a paid mutator transaction binding the contract method 0x076e5718.
//
// Solidity: function proposerPayETHAndWETH((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) payable returns()
func (_Edem *EdemTransactor) ProposerPayETHAndWETH(opts *bind.TransactOpts, userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.contract.Transact(opts, "proposerPayETHAndWETH", userProposer, marketSeller)
}

// ProposerPayETHAndWETH is a paid mutator transaction binding the contract method 0x076e5718.
//
// Solidity: function proposerPayETHAndWETH((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) payable returns()
func (_Edem *EdemSession) ProposerPayETHAndWETH(userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.ProposerPayETHAndWETH(&_Edem.TransactOpts, userProposer, marketSeller)
}

// ProposerPayETHAndWETH is a paid mutator transaction binding the contract method 0x076e5718.
//
// Solidity: function proposerPayETHAndWETH((bool,address,uint256,uint256) userProposer, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketSeller) payable returns()
func (_Edem *EdemTransactorSession) ProposerPayETHAndWETH(userProposer MarketTypesUser, marketSeller MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.ProposerPayETHAndWETH(&_Edem.TransactOpts, userProposer, marketSeller)
}

// SuggestApprove is a paid mutator transaction binding the contract method 0x2ef19e0d.
//
// Solidity: function suggestApprove((bool,address,uint256,uint256) userSeller, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketProposer) returns()
func (_Edem *EdemTransactor) SuggestApprove(opts *bind.TransactOpts, userSeller MarketTypesUser, marketProposer MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.contract.Transact(opts, "suggestApprove", userSeller, marketProposer)
}

// SuggestApprove is a paid mutator transaction binding the contract method 0x2ef19e0d.
//
// Solidity: function suggestApprove((bool,address,uint256,uint256) userSeller, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketProposer) returns()
func (_Edem *EdemSession) SuggestApprove(userSeller MarketTypesUser, marketProposer MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.SuggestApprove(&_Edem.TransactOpts, userSeller, marketProposer)
}

// SuggestApprove is a paid mutator transaction binding the contract method 0x2ef19e0d.
//
// Solidity: function suggestApprove((bool,address,uint256,uint256) userSeller, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8,bytes32,bytes32) marketProposer) returns()
func (_Edem *EdemTransactorSession) SuggestApprove(userSeller MarketTypesUser, marketProposer MarketTypesMarketPlace) (*types.Transaction, error) {
	return _Edem.Contract.SuggestApprove(&_Edem.TransactOpts, userSeller, marketProposer)
}

// EdemCancelAllOrdersIterator is returned from FilterCancelAllOrders and is used to iterate over the raw logs and unpacked data for CancelAllOrders events raised by the Edem contract.
type EdemCancelAllOrdersIterator struct {
	Event *EdemCancelAllOrders // Event containing the contract specifics and raw log

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
func (it *EdemCancelAllOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemCancelAllOrders)
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
		it.Event = new(EdemCancelAllOrders)
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
func (it *EdemCancelAllOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemCancelAllOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemCancelAllOrders represents a CancelAllOrders event raised by the Edem contract.
type EdemCancelAllOrders struct {
	User        common.Address
	NewMinNonce *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelAllOrders is a free log retrieval operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Edem *EdemFilterer) FilterCancelAllOrders(opts *bind.FilterOpts, user []common.Address) (*EdemCancelAllOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &EdemCancelAllOrdersIterator{contract: _Edem.contract, event: "CancelAllOrders", logs: logs, sub: sub}, nil
}

// WatchCancelAllOrders is a free log subscription operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Edem *EdemFilterer) WatchCancelAllOrders(opts *bind.WatchOpts, sink chan<- *EdemCancelAllOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemCancelAllOrders)
				if err := _Edem.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
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

// ParseCancelAllOrders is a log parse operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Edem *EdemFilterer) ParseCancelAllOrders(log types.Log) (*EdemCancelAllOrders, error) {
	event := new(EdemCancelAllOrders)
	if err := _Edem.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemCancelMultipleOrdersIterator is returned from FilterCancelMultipleOrders and is used to iterate over the raw logs and unpacked data for CancelMultipleOrders events raised by the Edem contract.
type EdemCancelMultipleOrdersIterator struct {
	Event *EdemCancelMultipleOrders // Event containing the contract specifics and raw log

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
func (it *EdemCancelMultipleOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemCancelMultipleOrders)
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
		it.Event = new(EdemCancelMultipleOrders)
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
func (it *EdemCancelMultipleOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemCancelMultipleOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemCancelMultipleOrders represents a CancelMultipleOrders event raised by the Edem contract.
type EdemCancelMultipleOrders struct {
	User        common.Address
	OrderNonces []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelMultipleOrders is a free log retrieval operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Edem *EdemFilterer) FilterCancelMultipleOrders(opts *bind.FilterOpts, user []common.Address) (*EdemCancelMultipleOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &EdemCancelMultipleOrdersIterator{contract: _Edem.contract, event: "CancelMultipleOrders", logs: logs, sub: sub}, nil
}

// WatchCancelMultipleOrders is a free log subscription operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Edem *EdemFilterer) WatchCancelMultipleOrders(opts *bind.WatchOpts, sink chan<- *EdemCancelMultipleOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemCancelMultipleOrders)
				if err := _Edem.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
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

// ParseCancelMultipleOrders is a log parse operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Edem *EdemFilterer) ParseCancelMultipleOrders(log types.Log) (*EdemCancelMultipleOrders, error) {
	event := new(EdemCancelMultipleOrders)
	if err := _Edem.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemNewProtocolExecutionManagerIterator is returned from FilterNewProtocolExecutionManager and is used to iterate over the raw logs and unpacked data for NewProtocolExecutionManager events raised by the Edem contract.
type EdemNewProtocolExecutionManagerIterator struct {
	Event *EdemNewProtocolExecutionManager // Event containing the contract specifics and raw log

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
func (it *EdemNewProtocolExecutionManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemNewProtocolExecutionManager)
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
		it.Event = new(EdemNewProtocolExecutionManager)
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
func (it *EdemNewProtocolExecutionManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemNewProtocolExecutionManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemNewProtocolExecutionManager represents a NewProtocolExecutionManager event raised by the Edem contract.
type EdemNewProtocolExecutionManager struct {
	ProtocolExecutionManager common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolExecutionManager is a free log retrieval operation binding the contract event 0x418a3785ac5fd0d0b92e92d1e814233310678c13b024a3741d41c6db3c803fc3.
//
// Solidity: event NewProtocolExecutionManager(address indexed protocolExecutionManager)
func (_Edem *EdemFilterer) FilterNewProtocolExecutionManager(opts *bind.FilterOpts, protocolExecutionManager []common.Address) (*EdemNewProtocolExecutionManagerIterator, error) {

	var protocolExecutionManagerRule []interface{}
	for _, protocolExecutionManagerItem := range protocolExecutionManager {
		protocolExecutionManagerRule = append(protocolExecutionManagerRule, protocolExecutionManagerItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "NewProtocolExecutionManager", protocolExecutionManagerRule)
	if err != nil {
		return nil, err
	}
	return &EdemNewProtocolExecutionManagerIterator{contract: _Edem.contract, event: "NewProtocolExecutionManager", logs: logs, sub: sub}, nil
}

// WatchNewProtocolExecutionManager is a free log subscription operation binding the contract event 0x418a3785ac5fd0d0b92e92d1e814233310678c13b024a3741d41c6db3c803fc3.
//
// Solidity: event NewProtocolExecutionManager(address indexed protocolExecutionManager)
func (_Edem *EdemFilterer) WatchNewProtocolExecutionManager(opts *bind.WatchOpts, sink chan<- *EdemNewProtocolExecutionManager, protocolExecutionManager []common.Address) (event.Subscription, error) {

	var protocolExecutionManagerRule []interface{}
	for _, protocolExecutionManagerItem := range protocolExecutionManager {
		protocolExecutionManagerRule = append(protocolExecutionManagerRule, protocolExecutionManagerItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "NewProtocolExecutionManager", protocolExecutionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemNewProtocolExecutionManager)
				if err := _Edem.contract.UnpackLog(event, "NewProtocolExecutionManager", log); err != nil {
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

// ParseNewProtocolExecutionManager is a log parse operation binding the contract event 0x418a3785ac5fd0d0b92e92d1e814233310678c13b024a3741d41c6db3c803fc3.
//
// Solidity: event NewProtocolExecutionManager(address indexed protocolExecutionManager)
func (_Edem *EdemFilterer) ParseNewProtocolExecutionManager(log types.Log) (*EdemNewProtocolExecutionManager, error) {
	event := new(EdemNewProtocolExecutionManager)
	if err := _Edem.contract.UnpackLog(event, "NewProtocolExecutionManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemNewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the Edem contract.
type EdemNewProtocolFeeRecipientIterator struct {
	Event *EdemNewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *EdemNewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemNewProtocolFeeRecipient)
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
		it.Event = new(EdemNewProtocolFeeRecipient)
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
func (it *EdemNewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemNewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemNewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the Edem contract.
type EdemNewProtocolFeeRecipient struct {
	ProtocoleFeeRecipient common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocoleFeeRecipient)
func (_Edem *EdemFilterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts, protocoleFeeRecipient []common.Address) (*EdemNewProtocolFeeRecipientIterator, error) {

	var protocoleFeeRecipientRule []interface{}
	for _, protocoleFeeRecipientItem := range protocoleFeeRecipient {
		protocoleFeeRecipientRule = append(protocoleFeeRecipientRule, protocoleFeeRecipientItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "NewProtocolFeeRecipient", protocoleFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &EdemNewProtocolFeeRecipientIterator{contract: _Edem.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocoleFeeRecipient)
func (_Edem *EdemFilterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *EdemNewProtocolFeeRecipient, protocoleFeeRecipient []common.Address) (event.Subscription, error) {

	var protocoleFeeRecipientRule []interface{}
	for _, protocoleFeeRecipientItem := range protocoleFeeRecipient {
		protocoleFeeRecipientRule = append(protocoleFeeRecipientRule, protocoleFeeRecipientItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "NewProtocolFeeRecipient", protocoleFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemNewProtocolFeeRecipient)
				if err := _Edem.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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

// ParseNewProtocolFeeRecipient is a log parse operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocoleFeeRecipient)
func (_Edem *EdemFilterer) ParseNewProtocolFeeRecipient(log types.Log) (*EdemNewProtocolFeeRecipient, error) {
	event := new(EdemNewProtocolFeeRecipient)
	if err := _Edem.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemNewRoyaltyManagerIterator is returned from FilterNewRoyaltyManager and is used to iterate over the raw logs and unpacked data for NewRoyaltyManager events raised by the Edem contract.
type EdemNewRoyaltyManagerIterator struct {
	Event *EdemNewRoyaltyManager // Event containing the contract specifics and raw log

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
func (it *EdemNewRoyaltyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemNewRoyaltyManager)
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
		it.Event = new(EdemNewRoyaltyManager)
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
func (it *EdemNewRoyaltyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemNewRoyaltyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemNewRoyaltyManager represents a NewRoyaltyManager event raised by the Edem contract.
type EdemNewRoyaltyManager struct {
	RoyaltyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRoyaltyManager is a free log retrieval operation binding the contract event 0x23e4135208439ee9f05d104daf57025ac32a6842326405d76852f0434085e933.
//
// Solidity: event NewRoyaltyManager(address indexed royaltyManager)
func (_Edem *EdemFilterer) FilterNewRoyaltyManager(opts *bind.FilterOpts, royaltyManager []common.Address) (*EdemNewRoyaltyManagerIterator, error) {

	var royaltyManagerRule []interface{}
	for _, royaltyManagerItem := range royaltyManager {
		royaltyManagerRule = append(royaltyManagerRule, royaltyManagerItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "NewRoyaltyManager", royaltyManagerRule)
	if err != nil {
		return nil, err
	}
	return &EdemNewRoyaltyManagerIterator{contract: _Edem.contract, event: "NewRoyaltyManager", logs: logs, sub: sub}, nil
}

// WatchNewRoyaltyManager is a free log subscription operation binding the contract event 0x23e4135208439ee9f05d104daf57025ac32a6842326405d76852f0434085e933.
//
// Solidity: event NewRoyaltyManager(address indexed royaltyManager)
func (_Edem *EdemFilterer) WatchNewRoyaltyManager(opts *bind.WatchOpts, sink chan<- *EdemNewRoyaltyManager, royaltyManager []common.Address) (event.Subscription, error) {

	var royaltyManagerRule []interface{}
	for _, royaltyManagerItem := range royaltyManager {
		royaltyManagerRule = append(royaltyManagerRule, royaltyManagerItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "NewRoyaltyManager", royaltyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemNewRoyaltyManager)
				if err := _Edem.contract.UnpackLog(event, "NewRoyaltyManager", log); err != nil {
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

// ParseNewRoyaltyManager is a log parse operation binding the contract event 0x23e4135208439ee9f05d104daf57025ac32a6842326405d76852f0434085e933.
//
// Solidity: event NewRoyaltyManager(address indexed royaltyManager)
func (_Edem *EdemFilterer) ParseNewRoyaltyManager(log types.Log) (*EdemNewRoyaltyManager, error) {
	event := new(EdemNewRoyaltyManager)
	if err := _Edem.contract.UnpackLog(event, "NewRoyaltyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemNewTradeTokenManagerIterator is returned from FilterNewTradeTokenManager and is used to iterate over the raw logs and unpacked data for NewTradeTokenManager events raised by the Edem contract.
type EdemNewTradeTokenManagerIterator struct {
	Event *EdemNewTradeTokenManager // Event containing the contract specifics and raw log

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
func (it *EdemNewTradeTokenManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemNewTradeTokenManager)
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
		it.Event = new(EdemNewTradeTokenManager)
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
func (it *EdemNewTradeTokenManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemNewTradeTokenManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemNewTradeTokenManager represents a NewTradeTokenManager event raised by the Edem contract.
type EdemNewTradeTokenManager struct {
	TradeTokenManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewTradeTokenManager is a free log retrieval operation binding the contract event 0x2c3e29f8b5d22d1cfa8497a6c45621ed49f95178416c7ef3ee1c2193c8812eb4.
//
// Solidity: event NewTradeTokenManager(address indexed tradeTokenManager)
func (_Edem *EdemFilterer) FilterNewTradeTokenManager(opts *bind.FilterOpts, tradeTokenManager []common.Address) (*EdemNewTradeTokenManagerIterator, error) {

	var tradeTokenManagerRule []interface{}
	for _, tradeTokenManagerItem := range tradeTokenManager {
		tradeTokenManagerRule = append(tradeTokenManagerRule, tradeTokenManagerItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "NewTradeTokenManager", tradeTokenManagerRule)
	if err != nil {
		return nil, err
	}
	return &EdemNewTradeTokenManagerIterator{contract: _Edem.contract, event: "NewTradeTokenManager", logs: logs, sub: sub}, nil
}

// WatchNewTradeTokenManager is a free log subscription operation binding the contract event 0x2c3e29f8b5d22d1cfa8497a6c45621ed49f95178416c7ef3ee1c2193c8812eb4.
//
// Solidity: event NewTradeTokenManager(address indexed tradeTokenManager)
func (_Edem *EdemFilterer) WatchNewTradeTokenManager(opts *bind.WatchOpts, sink chan<- *EdemNewTradeTokenManager, tradeTokenManager []common.Address) (event.Subscription, error) {

	var tradeTokenManagerRule []interface{}
	for _, tradeTokenManagerItem := range tradeTokenManager {
		tradeTokenManagerRule = append(tradeTokenManagerRule, tradeTokenManagerItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "NewTradeTokenManager", tradeTokenManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemNewTradeTokenManager)
				if err := _Edem.contract.UnpackLog(event, "NewTradeTokenManager", log); err != nil {
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

// ParseNewTradeTokenManager is a log parse operation binding the contract event 0x2c3e29f8b5d22d1cfa8497a6c45621ed49f95178416c7ef3ee1c2193c8812eb4.
//
// Solidity: event NewTradeTokenManager(address indexed tradeTokenManager)
func (_Edem *EdemFilterer) ParseNewTradeTokenManager(log types.Log) (*EdemNewTradeTokenManager, error) {
	event := new(EdemNewTradeTokenManager)
	if err := _Edem.contract.UnpackLog(event, "NewTradeTokenManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemNewTransferManagerIterator is returned from FilterNewTransferManager and is used to iterate over the raw logs and unpacked data for NewTransferManager events raised by the Edem contract.
type EdemNewTransferManagerIterator struct {
	Event *EdemNewTransferManager // Event containing the contract specifics and raw log

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
func (it *EdemNewTransferManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemNewTransferManager)
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
		it.Event = new(EdemNewTransferManager)
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
func (it *EdemNewTransferManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemNewTransferManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemNewTransferManager represents a NewTransferManager event raised by the Edem contract.
type EdemNewTransferManager struct {
	TransferManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewTransferManager is a free log retrieval operation binding the contract event 0xbb34fee0d3bb7c602ecfe1c8bed62b67a720f718e0eaf72246d3457fcb25be17.
//
// Solidity: event NewTransferManager(address indexed transferManager)
func (_Edem *EdemFilterer) FilterNewTransferManager(opts *bind.FilterOpts, transferManager []common.Address) (*EdemNewTransferManagerIterator, error) {

	var transferManagerRule []interface{}
	for _, transferManagerItem := range transferManager {
		transferManagerRule = append(transferManagerRule, transferManagerItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "NewTransferManager", transferManagerRule)
	if err != nil {
		return nil, err
	}
	return &EdemNewTransferManagerIterator{contract: _Edem.contract, event: "NewTransferManager", logs: logs, sub: sub}, nil
}

// WatchNewTransferManager is a free log subscription operation binding the contract event 0xbb34fee0d3bb7c602ecfe1c8bed62b67a720f718e0eaf72246d3457fcb25be17.
//
// Solidity: event NewTransferManager(address indexed transferManager)
func (_Edem *EdemFilterer) WatchNewTransferManager(opts *bind.WatchOpts, sink chan<- *EdemNewTransferManager, transferManager []common.Address) (event.Subscription, error) {

	var transferManagerRule []interface{}
	for _, transferManagerItem := range transferManager {
		transferManagerRule = append(transferManagerRule, transferManagerItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "NewTransferManager", transferManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemNewTransferManager)
				if err := _Edem.contract.UnpackLog(event, "NewTransferManager", log); err != nil {
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

// ParseNewTransferManager is a log parse operation binding the contract event 0xbb34fee0d3bb7c602ecfe1c8bed62b67a720f718e0eaf72246d3457fcb25be17.
//
// Solidity: event NewTransferManager(address indexed transferManager)
func (_Edem *EdemFilterer) ParseNewTransferManager(log types.Log) (*EdemNewTransferManager, error) {
	event := new(EdemNewTransferManager)
	if err := _Edem.contract.UnpackLog(event, "NewTransferManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemRoyaltyPaymentIterator is returned from FilterRoyaltyPayment and is used to iterate over the raw logs and unpacked data for RoyaltyPayment events raised by the Edem contract.
type EdemRoyaltyPaymentIterator struct {
	Event *EdemRoyaltyPayment // Event containing the contract specifics and raw log

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
func (it *EdemRoyaltyPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemRoyaltyPayment)
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
		it.Event = new(EdemRoyaltyPayment)
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
func (it *EdemRoyaltyPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemRoyaltyPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemRoyaltyPayment represents a RoyaltyPayment event raised by the Edem contract.
type EdemRoyaltyPayment struct {
	Collection       common.Address
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	Currency         common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPayment is a free log retrieval operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Edem *EdemFilterer) FilterRoyaltyPayment(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (*EdemRoyaltyPaymentIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &EdemRoyaltyPaymentIterator{contract: _Edem.contract, event: "RoyaltyPayment", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPayment is a free log subscription operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Edem *EdemFilterer) WatchRoyaltyPayment(opts *bind.WatchOpts, sink chan<- *EdemRoyaltyPayment, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemRoyaltyPayment)
				if err := _Edem.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
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

// ParseRoyaltyPayment is a log parse operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Edem *EdemFilterer) ParseRoyaltyPayment(log types.Log) (*EdemRoyaltyPayment, error) {
	event := new(EdemRoyaltyPayment)
	if err := _Edem.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemUserProposerIterator is returned from FilterUserProposer and is used to iterate over the raw logs and unpacked data for UserProposer events raised by the Edem contract.
type EdemUserProposerIterator struct {
	Event *EdemUserProposer // Event containing the contract specifics and raw log

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
func (it *EdemUserProposerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemUserProposer)
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
		it.Event = new(EdemUserProposer)
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
func (it *EdemUserProposerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemUserProposerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemUserProposer represents a UserProposer event raised by the Edem contract.
type EdemUserProposer struct {
	MarketPlaceHash   [32]byte
	OrderNonce        *big.Int
	UserProposer      common.Address
	MarketSeller      common.Address
	ProtocolAddress   common.Address
	TradeTokenAddress common.Address
	Collection        common.Address
	TokenId           *big.Int
	Amount            *big.Int
	Price             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserProposer is a free log retrieval operation binding the contract event 0x50d3c79a72acf972e873ec1c16c55eecf26a8bb67b60b136dd0ca2e1c6565ad7.
//
// Solidity: event UserProposer(bytes32 marketPlaceHash, uint256 orderNonce, address indexed userProposer, address indexed marketSeller, address indexed protocolAddress, address tradeTokenAddress, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Edem *EdemFilterer) FilterUserProposer(opts *bind.FilterOpts, userProposer []common.Address, marketSeller []common.Address, protocolAddress []common.Address) (*EdemUserProposerIterator, error) {

	var userProposerRule []interface{}
	for _, userProposerItem := range userProposer {
		userProposerRule = append(userProposerRule, userProposerItem)
	}
	var marketSellerRule []interface{}
	for _, marketSellerItem := range marketSeller {
		marketSellerRule = append(marketSellerRule, marketSellerItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "UserProposer", userProposerRule, marketSellerRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &EdemUserProposerIterator{contract: _Edem.contract, event: "UserProposer", logs: logs, sub: sub}, nil
}

// WatchUserProposer is a free log subscription operation binding the contract event 0x50d3c79a72acf972e873ec1c16c55eecf26a8bb67b60b136dd0ca2e1c6565ad7.
//
// Solidity: event UserProposer(bytes32 marketPlaceHash, uint256 orderNonce, address indexed userProposer, address indexed marketSeller, address indexed protocolAddress, address tradeTokenAddress, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Edem *EdemFilterer) WatchUserProposer(opts *bind.WatchOpts, sink chan<- *EdemUserProposer, userProposer []common.Address, marketSeller []common.Address, protocolAddress []common.Address) (event.Subscription, error) {

	var userProposerRule []interface{}
	for _, userProposerItem := range userProposer {
		userProposerRule = append(userProposerRule, userProposerItem)
	}
	var marketSellerRule []interface{}
	for _, marketSellerItem := range marketSeller {
		marketSellerRule = append(marketSellerRule, marketSellerItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "UserProposer", userProposerRule, marketSellerRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemUserProposer)
				if err := _Edem.contract.UnpackLog(event, "UserProposer", log); err != nil {
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

// ParseUserProposer is a log parse operation binding the contract event 0x50d3c79a72acf972e873ec1c16c55eecf26a8bb67b60b136dd0ca2e1c6565ad7.
//
// Solidity: event UserProposer(bytes32 marketPlaceHash, uint256 orderNonce, address indexed userProposer, address indexed marketSeller, address indexed protocolAddress, address tradeTokenAddress, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Edem *EdemFilterer) ParseUserProposer(log types.Log) (*EdemUserProposer, error) {
	event := new(EdemUserProposer)
	if err := _Edem.contract.UnpackLog(event, "UserProposer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdemUserSellerIterator is returned from FilterUserSeller and is used to iterate over the raw logs and unpacked data for UserSeller events raised by the Edem contract.
type EdemUserSellerIterator struct {
	Event *EdemUserSeller // Event containing the contract specifics and raw log

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
func (it *EdemUserSellerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdemUserSeller)
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
		it.Event = new(EdemUserSeller)
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
func (it *EdemUserSellerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdemUserSellerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdemUserSeller represents a UserSeller event raised by the Edem contract.
type EdemUserSeller struct {
	MarketPlaceHash   [32]byte
	OrderNonce        *big.Int
	UserSeller        common.Address
	MarketProposer    common.Address
	ProtocolAddress   common.Address
	TradeTokenAddress common.Address
	Collection        common.Address
	TokenId           *big.Int
	NftAmount         *big.Int
	Price             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserSeller is a free log retrieval operation binding the contract event 0x8ce1a009c26eecc60e31b457e624e94f61c3d5f4149f76e8f43f73ebc27b9300.
//
// Solidity: event UserSeller(bytes32 marketPlaceHash, uint256 orderNonce, address indexed userSeller, address indexed marketProposer, address indexed protocolAddress, address tradeTokenAddress, address collection, uint256 tokenId, uint256 nftAmount, uint256 price)
func (_Edem *EdemFilterer) FilterUserSeller(opts *bind.FilterOpts, userSeller []common.Address, marketProposer []common.Address, protocolAddress []common.Address) (*EdemUserSellerIterator, error) {

	var userSellerRule []interface{}
	for _, userSellerItem := range userSeller {
		userSellerRule = append(userSellerRule, userSellerItem)
	}
	var marketProposerRule []interface{}
	for _, marketProposerItem := range marketProposer {
		marketProposerRule = append(marketProposerRule, marketProposerItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Edem.contract.FilterLogs(opts, "UserSeller", userSellerRule, marketProposerRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &EdemUserSellerIterator{contract: _Edem.contract, event: "UserSeller", logs: logs, sub: sub}, nil
}

// WatchUserSeller is a free log subscription operation binding the contract event 0x8ce1a009c26eecc60e31b457e624e94f61c3d5f4149f76e8f43f73ebc27b9300.
//
// Solidity: event UserSeller(bytes32 marketPlaceHash, uint256 orderNonce, address indexed userSeller, address indexed marketProposer, address indexed protocolAddress, address tradeTokenAddress, address collection, uint256 tokenId, uint256 nftAmount, uint256 price)
func (_Edem *EdemFilterer) WatchUserSeller(opts *bind.WatchOpts, sink chan<- *EdemUserSeller, userSeller []common.Address, marketProposer []common.Address, protocolAddress []common.Address) (event.Subscription, error) {

	var userSellerRule []interface{}
	for _, userSellerItem := range userSeller {
		userSellerRule = append(userSellerRule, userSellerItem)
	}
	var marketProposerRule []interface{}
	for _, marketProposerItem := range marketProposer {
		marketProposerRule = append(marketProposerRule, marketProposerItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _Edem.contract.WatchLogs(opts, "UserSeller", userSellerRule, marketProposerRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdemUserSeller)
				if err := _Edem.contract.UnpackLog(event, "UserSeller", log); err != nil {
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

// ParseUserSeller is a log parse operation binding the contract event 0x8ce1a009c26eecc60e31b457e624e94f61c3d5f4149f76e8f43f73ebc27b9300.
//
// Solidity: event UserSeller(bytes32 marketPlaceHash, uint256 orderNonce, address indexed userSeller, address indexed marketProposer, address indexed protocolAddress, address tradeTokenAddress, address collection, uint256 tokenId, uint256 nftAmount, uint256 price)
func (_Edem *EdemFilterer) ParseUserSeller(log types.Log) (*EdemUserSeller, error) {
	event := new(EdemUserSeller)
	if err := _Edem.contract.UnpackLog(event, "UserSeller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
