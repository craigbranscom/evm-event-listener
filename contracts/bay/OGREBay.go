// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bay

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

// BayMetaData contains all meta data concerning the Bay contract.
var BayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daoAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderFee_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowlistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"AskCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"BayCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"BidCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOrderFee\",\"type\":\"uint256\"}],\"name\":\"OrderFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"asks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currencyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calcOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"cancelAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"cancelBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createAsk\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setContractAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newOrderFee\",\"type\":\"uint256\"}],\"name\":\"setOrderFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BayABI is the input ABI used to generate the binding from.
// Deprecated: Use BayMetaData.ABI instead.
var BayABI = BayMetaData.ABI

// Bay is an auto generated Go binding around an Ethereum contract.
type Bay struct {
	BayCaller     // Read-only binding to the contract
	BayTransactor // Write-only binding to the contract
	BayFilterer   // Log filterer for contract events
}

// BayCaller is an auto generated read-only Go binding around an Ethereum contract.
type BayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaySession struct {
	Contract     *Bay              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BayCallerSession struct {
	Contract *BayCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BayTransactorSession struct {
	Contract     *BayTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BayRaw is an auto generated low-level Go binding around an Ethereum contract.
type BayRaw struct {
	Contract *Bay // Generic contract binding to access the raw methods on
}

// BayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BayCallerRaw struct {
	Contract *BayCaller // Generic read-only contract binding to access the raw methods on
}

// BayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BayTransactorRaw struct {
	Contract *BayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBay creates a new instance of Bay, bound to a specific deployed contract.
func NewBay(address common.Address, backend bind.ContractBackend) (*Bay, error) {
	contract, err := bindBay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bay{BayCaller: BayCaller{contract: contract}, BayTransactor: BayTransactor{contract: contract}, BayFilterer: BayFilterer{contract: contract}}, nil
}

// NewBayCaller creates a new read-only instance of Bay, bound to a specific deployed contract.
func NewBayCaller(address common.Address, caller bind.ContractCaller) (*BayCaller, error) {
	contract, err := bindBay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BayCaller{contract: contract}, nil
}

// NewBayTransactor creates a new write-only instance of Bay, bound to a specific deployed contract.
func NewBayTransactor(address common.Address, transactor bind.ContractTransactor) (*BayTransactor, error) {
	contract, err := bindBay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BayTransactor{contract: contract}, nil
}

// NewBayFilterer creates a new log filterer instance of Bay, bound to a specific deployed contract.
func NewBayFilterer(address common.Address, filterer bind.ContractFilterer) (*BayFilterer, error) {
	contract, err := bindBay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BayFilterer{contract: contract}, nil
}

// bindBay binds a generic wrapper to an already deployed contract.
func bindBay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bay *BayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bay.Contract.BayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bay *BayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bay.Contract.BayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bay *BayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bay.Contract.BayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bay *BayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bay *BayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bay *BayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bay.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bay *BayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bay *BaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bay.Contract.DEFAULTADMINROLE(&_Bay.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bay *BayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bay.Contract.DEFAULTADMINROLE(&_Bay.CallOpts)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_Bay *BayCaller) AllowedContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "allowedContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_Bay *BaySession) AllowedContracts(arg0 common.Address) (bool, error) {
	return _Bay.Contract.AllowedContracts(&_Bay.CallOpts, arg0)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x51e0e26b.
//
// Solidity: function allowedContracts(address ) view returns(bool)
func (_Bay *BayCallerSession) AllowedContracts(arg0 common.Address) (bool, error) {
	return _Bay.Contract.AllowedContracts(&_Bay.CallOpts, arg0)
}

// Asks is a free data retrieval call binding the contract method 0xfe2e8c9c.
//
// Solidity: function asks(bytes32 ) view returns(address creator, address nftAddress, uint256 tokenId, address currencyAddress, uint256 amount)
func (_Bay *BayCaller) Asks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Creator         common.Address
	NftAddress      common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "asks", arg0)

	outstruct := new(struct {
		Creator         common.Address
		NftAddress      common.Address
		TokenId         *big.Int
		CurrencyAddress common.Address
		Amount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrencyAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Asks is a free data retrieval call binding the contract method 0xfe2e8c9c.
//
// Solidity: function asks(bytes32 ) view returns(address creator, address nftAddress, uint256 tokenId, address currencyAddress, uint256 amount)
func (_Bay *BaySession) Asks(arg0 [32]byte) (struct {
	Creator         common.Address
	NftAddress      common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	return _Bay.Contract.Asks(&_Bay.CallOpts, arg0)
}

// Asks is a free data retrieval call binding the contract method 0xfe2e8c9c.
//
// Solidity: function asks(bytes32 ) view returns(address creator, address nftAddress, uint256 tokenId, address currencyAddress, uint256 amount)
func (_Bay *BayCallerSession) Asks(arg0 [32]byte) (struct {
	Creator         common.Address
	NftAddress      common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	return _Bay.Contract.Asks(&_Bay.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x8f98eeda.
//
// Solidity: function bids(bytes32 ) view returns(address creator, address nftAddress, uint256 tokenId, address currencyAddress, uint256 amount)
func (_Bay *BayCaller) Bids(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Creator         common.Address
	NftAddress      common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "bids", arg0)

	outstruct := new(struct {
		Creator         common.Address
		NftAddress      common.Address
		TokenId         *big.Int
		CurrencyAddress common.Address
		Amount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrencyAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x8f98eeda.
//
// Solidity: function bids(bytes32 ) view returns(address creator, address nftAddress, uint256 tokenId, address currencyAddress, uint256 amount)
func (_Bay *BaySession) Bids(arg0 [32]byte) (struct {
	Creator         common.Address
	NftAddress      common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	return _Bay.Contract.Bids(&_Bay.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x8f98eeda.
//
// Solidity: function bids(bytes32 ) view returns(address creator, address nftAddress, uint256 tokenId, address currencyAddress, uint256 amount)
func (_Bay *BayCallerSession) Bids(arg0 [32]byte) (struct {
	Creator         common.Address
	NftAddress      common.Address
	TokenId         *big.Int
	CurrencyAddress common.Address
	Amount          *big.Int
}, error) {
	return _Bay.Contract.Bids(&_Bay.CallOpts, arg0)
}

// CalcOrderHash is a free data retrieval call binding the contract method 0xff97be5c.
//
// Solidity: function calcOrderHash(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) pure returns(bytes32)
func (_Bay *BayCaller) CalcOrderHash(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "calcOrderHash", erc721Address, tokenId, erc20Address, amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalcOrderHash is a free data retrieval call binding the contract method 0xff97be5c.
//
// Solidity: function calcOrderHash(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) pure returns(bytes32)
func (_Bay *BaySession) CalcOrderHash(erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) ([32]byte, error) {
	return _Bay.Contract.CalcOrderHash(&_Bay.CallOpts, erc721Address, tokenId, erc20Address, amount)
}

// CalcOrderHash is a free data retrieval call binding the contract method 0xff97be5c.
//
// Solidity: function calcOrderHash(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) pure returns(bytes32)
func (_Bay *BayCallerSession) CalcOrderHash(erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) ([32]byte, error) {
	return _Bay.Contract.CalcOrderHash(&_Bay.CallOpts, erc721Address, tokenId, erc20Address, amount)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Bay *BayCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Bay *BaySession) DaoAddress() (common.Address, error) {
	return _Bay.Contract.DaoAddress(&_Bay.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Bay *BayCallerSession) DaoAddress() (common.Address, error) {
	return _Bay.Contract.DaoAddress(&_Bay.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bay *BayCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bay *BaySession) FeeRecipient() (common.Address, error) {
	return _Bay.Contract.FeeRecipient(&_Bay.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bay *BayCallerSession) FeeRecipient() (common.Address, error) {
	return _Bay.Contract.FeeRecipient(&_Bay.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bay *BayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bay *BaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bay.Contract.GetRoleAdmin(&_Bay.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bay *BayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bay.Contract.GetRoleAdmin(&_Bay.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bay *BayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bay *BaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bay.Contract.HasRole(&_Bay.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bay *BayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bay.Contract.HasRole(&_Bay.CallOpts, role, account)
}

// OrderFee is a free data retrieval call binding the contract method 0x1392fb3e.
//
// Solidity: function orderFee() view returns(uint256)
func (_Bay *BayCaller) OrderFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "orderFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderFee is a free data retrieval call binding the contract method 0x1392fb3e.
//
// Solidity: function orderFee() view returns(uint256)
func (_Bay *BaySession) OrderFee() (*big.Int, error) {
	return _Bay.Contract.OrderFee(&_Bay.CallOpts)
}

// OrderFee is a free data retrieval call binding the contract method 0x1392fb3e.
//
// Solidity: function orderFee() view returns(uint256)
func (_Bay *BayCallerSession) OrderFee() (*big.Int, error) {
	return _Bay.Contract.OrderFee(&_Bay.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bay *BayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bay.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bay *BaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bay.Contract.SupportsInterface(&_Bay.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bay *BayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bay.Contract.SupportsInterface(&_Bay.CallOpts, interfaceId)
}

// CancelAsk is a paid mutator transaction binding the contract method 0x6a0fd985.
//
// Solidity: function cancelAsk(bytes32 orderHash) returns()
func (_Bay *BayTransactor) CancelAsk(opts *bind.TransactOpts, orderHash [32]byte) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "cancelAsk", orderHash)
}

// CancelAsk is a paid mutator transaction binding the contract method 0x6a0fd985.
//
// Solidity: function cancelAsk(bytes32 orderHash) returns()
func (_Bay *BaySession) CancelAsk(orderHash [32]byte) (*types.Transaction, error) {
	return _Bay.Contract.CancelAsk(&_Bay.TransactOpts, orderHash)
}

// CancelAsk is a paid mutator transaction binding the contract method 0x6a0fd985.
//
// Solidity: function cancelAsk(bytes32 orderHash) returns()
func (_Bay *BayTransactorSession) CancelAsk(orderHash [32]byte) (*types.Transaction, error) {
	return _Bay.Contract.CancelAsk(&_Bay.TransactOpts, orderHash)
}

// CancelBid is a paid mutator transaction binding the contract method 0xdf7cec28.
//
// Solidity: function cancelBid(bytes32 orderHash) returns()
func (_Bay *BayTransactor) CancelBid(opts *bind.TransactOpts, orderHash [32]byte) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "cancelBid", orderHash)
}

// CancelBid is a paid mutator transaction binding the contract method 0xdf7cec28.
//
// Solidity: function cancelBid(bytes32 orderHash) returns()
func (_Bay *BaySession) CancelBid(orderHash [32]byte) (*types.Transaction, error) {
	return _Bay.Contract.CancelBid(&_Bay.TransactOpts, orderHash)
}

// CancelBid is a paid mutator transaction binding the contract method 0xdf7cec28.
//
// Solidity: function cancelBid(bytes32 orderHash) returns()
func (_Bay *BayTransactorSession) CancelBid(orderHash [32]byte) (*types.Transaction, error) {
	return _Bay.Contract.CancelBid(&_Bay.TransactOpts, orderHash)
}

// CreateAsk is a paid mutator transaction binding the contract method 0x006795d5.
//
// Solidity: function createAsk(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) payable returns()
func (_Bay *BayTransactor) CreateAsk(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "createAsk", erc721Address, tokenId, erc20Address, amount)
}

// CreateAsk is a paid mutator transaction binding the contract method 0x006795d5.
//
// Solidity: function createAsk(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) payable returns()
func (_Bay *BaySession) CreateAsk(erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bay.Contract.CreateAsk(&_Bay.TransactOpts, erc721Address, tokenId, erc20Address, amount)
}

// CreateAsk is a paid mutator transaction binding the contract method 0x006795d5.
//
// Solidity: function createAsk(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) payable returns()
func (_Bay *BayTransactorSession) CreateAsk(erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bay.Contract.CreateAsk(&_Bay.TransactOpts, erc721Address, tokenId, erc20Address, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xbd3a27b0.
//
// Solidity: function createBid(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) payable returns()
func (_Bay *BayTransactor) CreateBid(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "createBid", erc721Address, tokenId, erc20Address, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xbd3a27b0.
//
// Solidity: function createBid(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) payable returns()
func (_Bay *BaySession) CreateBid(erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bay.Contract.CreateBid(&_Bay.TransactOpts, erc721Address, tokenId, erc20Address, amount)
}

// CreateBid is a paid mutator transaction binding the contract method 0xbd3a27b0.
//
// Solidity: function createBid(address erc721Address, uint256 tokenId, address erc20Address, uint256 amount) payable returns()
func (_Bay *BayTransactorSession) CreateBid(erc721Address common.Address, tokenId *big.Int, erc20Address common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bay.Contract.CreateBid(&_Bay.TransactOpts, erc721Address, tokenId, erc20Address, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bay *BayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bay *BaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.Contract.GrantRole(&_Bay.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bay *BayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.Contract.GrantRole(&_Bay.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bay *BayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bay *BaySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.Contract.RenounceRole(&_Bay.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bay *BayTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.Contract.RenounceRole(&_Bay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bay *BayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bay *BaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.Contract.RevokeRole(&_Bay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bay *BayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bay.Contract.RevokeRole(&_Bay.TransactOpts, role, account)
}

// SetContractAllowed is a paid mutator transaction binding the contract method 0x7965d76b.
//
// Solidity: function setContractAllowed(address contractAddress, bool allowed) returns()
func (_Bay *BayTransactor) SetContractAllowed(opts *bind.TransactOpts, contractAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "setContractAllowed", contractAddress, allowed)
}

// SetContractAllowed is a paid mutator transaction binding the contract method 0x7965d76b.
//
// Solidity: function setContractAllowed(address contractAddress, bool allowed) returns()
func (_Bay *BaySession) SetContractAllowed(contractAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _Bay.Contract.SetContractAllowed(&_Bay.TransactOpts, contractAddress, allowed)
}

// SetContractAllowed is a paid mutator transaction binding the contract method 0x7965d76b.
//
// Solidity: function setContractAllowed(address contractAddress, bool allowed) returns()
func (_Bay *BayTransactorSession) SetContractAllowed(contractAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _Bay.Contract.SetContractAllowed(&_Bay.TransactOpts, contractAddress, allowed)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Bay *BayTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Bay *BaySession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Bay.Contract.SetFeeRecipient(&_Bay.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_Bay *BayTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Bay.Contract.SetFeeRecipient(&_Bay.TransactOpts, newFeeRecipient)
}

// SetOrderFee is a paid mutator transaction binding the contract method 0xc014930f.
//
// Solidity: function setOrderFee(uint256 newOrderFee) returns()
func (_Bay *BayTransactor) SetOrderFee(opts *bind.TransactOpts, newOrderFee *big.Int) (*types.Transaction, error) {
	return _Bay.contract.Transact(opts, "setOrderFee", newOrderFee)
}

// SetOrderFee is a paid mutator transaction binding the contract method 0xc014930f.
//
// Solidity: function setOrderFee(uint256 newOrderFee) returns()
func (_Bay *BaySession) SetOrderFee(newOrderFee *big.Int) (*types.Transaction, error) {
	return _Bay.Contract.SetOrderFee(&_Bay.TransactOpts, newOrderFee)
}

// SetOrderFee is a paid mutator transaction binding the contract method 0xc014930f.
//
// Solidity: function setOrderFee(uint256 newOrderFee) returns()
func (_Bay *BayTransactorSession) SetOrderFee(newOrderFee *big.Int) (*types.Transaction, error) {
	return _Bay.Contract.SetOrderFee(&_Bay.TransactOpts, newOrderFee)
}

// BayAllowlistUpdatedIterator is returned from FilterAllowlistUpdated and is used to iterate over the raw logs and unpacked data for AllowlistUpdated events raised by the Bay contract.
type BayAllowlistUpdatedIterator struct {
	Event *BayAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *BayAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayAllowlistUpdated)
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
		it.Event = new(BayAllowlistUpdated)
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
func (it *BayAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayAllowlistUpdated represents a AllowlistUpdated event raised by the Bay contract.
type BayAllowlistUpdated struct {
	ContractAddress common.Address
	Allowed         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAllowlistUpdated is a free log retrieval operation binding the contract event 0x13518841ff4d3053cb7703afaa39b145c6331829b982d42f4d4fd7568b2e8e24.
//
// Solidity: event AllowlistUpdated(address contractAddress, bool allowed)
func (_Bay *BayFilterer) FilterAllowlistUpdated(opts *bind.FilterOpts) (*BayAllowlistUpdatedIterator, error) {

	logs, sub, err := _Bay.contract.FilterLogs(opts, "AllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &BayAllowlistUpdatedIterator{contract: _Bay.contract, event: "AllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowlistUpdated is a free log subscription operation binding the contract event 0x13518841ff4d3053cb7703afaa39b145c6331829b982d42f4d4fd7568b2e8e24.
//
// Solidity: event AllowlistUpdated(address contractAddress, bool allowed)
func (_Bay *BayFilterer) WatchAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *BayAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _Bay.contract.WatchLogs(opts, "AllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayAllowlistUpdated)
				if err := _Bay.contract.UnpackLog(event, "AllowlistUpdated", log); err != nil {
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

// ParseAllowlistUpdated is a log parse operation binding the contract event 0x13518841ff4d3053cb7703afaa39b145c6331829b982d42f4d4fd7568b2e8e24.
//
// Solidity: event AllowlistUpdated(address contractAddress, bool allowed)
func (_Bay *BayFilterer) ParseAllowlistUpdated(log types.Log) (*BayAllowlistUpdated, error) {
	event := new(BayAllowlistUpdated)
	if err := _Bay.contract.UnpackLog(event, "AllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayAskCancelledIterator is returned from FilterAskCancelled and is used to iterate over the raw logs and unpacked data for AskCancelled events raised by the Bay contract.
type BayAskCancelledIterator struct {
	Event *BayAskCancelled // Event containing the contract specifics and raw log

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
func (it *BayAskCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayAskCancelled)
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
		it.Event = new(BayAskCancelled)
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
func (it *BayAskCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayAskCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayAskCancelled represents a AskCancelled event raised by the Bay contract.
type BayAskCancelled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAskCancelled is a free log retrieval operation binding the contract event 0xa7e78ab2d839285f692f95318e21738ed76fe005b13097399f19eb043f46cbdd.
//
// Solidity: event AskCancelled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) FilterAskCancelled(opts *bind.FilterOpts, orderHash [][32]byte) (*BayAskCancelledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "AskCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &BayAskCancelledIterator{contract: _Bay.contract, event: "AskCancelled", logs: logs, sub: sub}, nil
}

// WatchAskCancelled is a free log subscription operation binding the contract event 0xa7e78ab2d839285f692f95318e21738ed76fe005b13097399f19eb043f46cbdd.
//
// Solidity: event AskCancelled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) WatchAskCancelled(opts *bind.WatchOpts, sink chan<- *BayAskCancelled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "AskCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayAskCancelled)
				if err := _Bay.contract.UnpackLog(event, "AskCancelled", log); err != nil {
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

// ParseAskCancelled is a log parse operation binding the contract event 0xa7e78ab2d839285f692f95318e21738ed76fe005b13097399f19eb043f46cbdd.
//
// Solidity: event AskCancelled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) ParseAskCancelled(log types.Log) (*BayAskCancelled, error) {
	event := new(BayAskCancelled)
	if err := _Bay.contract.UnpackLog(event, "AskCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayAskCreatedIterator is returned from FilterAskCreated and is used to iterate over the raw logs and unpacked data for AskCreated events raised by the Bay contract.
type BayAskCreatedIterator struct {
	Event *BayAskCreated // Event containing the contract specifics and raw log

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
func (it *BayAskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayAskCreated)
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
		it.Event = new(BayAskCreated)
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
func (it *BayAskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayAskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayAskCreated represents a AskCreated event raised by the Bay contract.
type BayAskCreated struct {
	OrderHash     [32]byte
	Creator       common.Address
	Erc721Address common.Address
	TokenId       *big.Int
	Erc20Address  common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAskCreated is a free log retrieval operation binding the contract event 0x22649edb01e89b0ff40ea9cbf2302950808ee082b372ffed1423e6e6cef8f49a.
//
// Solidity: event AskCreated(bytes32 indexed orderHash, address indexed creator, address indexed erc721Address, uint256 tokenId, address erc20Address, uint256 amount)
func (_Bay *BayFilterer) FilterAskCreated(opts *bind.FilterOpts, orderHash [][32]byte, creator []common.Address, erc721Address []common.Address) (*BayAskCreatedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "AskCreated", orderHashRule, creatorRule, erc721AddressRule)
	if err != nil {
		return nil, err
	}
	return &BayAskCreatedIterator{contract: _Bay.contract, event: "AskCreated", logs: logs, sub: sub}, nil
}

// WatchAskCreated is a free log subscription operation binding the contract event 0x22649edb01e89b0ff40ea9cbf2302950808ee082b372ffed1423e6e6cef8f49a.
//
// Solidity: event AskCreated(bytes32 indexed orderHash, address indexed creator, address indexed erc721Address, uint256 tokenId, address erc20Address, uint256 amount)
func (_Bay *BayFilterer) WatchAskCreated(opts *bind.WatchOpts, sink chan<- *BayAskCreated, orderHash [][32]byte, creator []common.Address, erc721Address []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "AskCreated", orderHashRule, creatorRule, erc721AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayAskCreated)
				if err := _Bay.contract.UnpackLog(event, "AskCreated", log); err != nil {
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

// ParseAskCreated is a log parse operation binding the contract event 0x22649edb01e89b0ff40ea9cbf2302950808ee082b372ffed1423e6e6cef8f49a.
//
// Solidity: event AskCreated(bytes32 indexed orderHash, address indexed creator, address indexed erc721Address, uint256 tokenId, address erc20Address, uint256 amount)
func (_Bay *BayFilterer) ParseAskCreated(log types.Log) (*BayAskCreated, error) {
	event := new(BayAskCreated)
	if err := _Bay.contract.UnpackLog(event, "AskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayBayCreatedIterator is returned from FilterBayCreated and is used to iterate over the raw logs and unpacked data for BayCreated events raised by the Bay contract.
type BayBayCreatedIterator struct {
	Event *BayBayCreated // Event containing the contract specifics and raw log

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
func (it *BayBayCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayBayCreated)
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
		it.Event = new(BayBayCreated)
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
func (it *BayBayCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayBayCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayBayCreated represents a BayCreated event raised by the Bay contract.
type BayBayCreated struct {
	DaoAddress common.Address
	Admin      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBayCreated is a free log retrieval operation binding the contract event 0xafcc4615a5c6ce81e1ee1ecc9f150c249104ed97f69f737f6eb8fec74b8a91e4.
//
// Solidity: event BayCreated(address daoAddress, address admin)
func (_Bay *BayFilterer) FilterBayCreated(opts *bind.FilterOpts) (*BayBayCreatedIterator, error) {

	logs, sub, err := _Bay.contract.FilterLogs(opts, "BayCreated")
	if err != nil {
		return nil, err
	}
	return &BayBayCreatedIterator{contract: _Bay.contract, event: "BayCreated", logs: logs, sub: sub}, nil
}

// WatchBayCreated is a free log subscription operation binding the contract event 0xafcc4615a5c6ce81e1ee1ecc9f150c249104ed97f69f737f6eb8fec74b8a91e4.
//
// Solidity: event BayCreated(address daoAddress, address admin)
func (_Bay *BayFilterer) WatchBayCreated(opts *bind.WatchOpts, sink chan<- *BayBayCreated) (event.Subscription, error) {

	logs, sub, err := _Bay.contract.WatchLogs(opts, "BayCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayBayCreated)
				if err := _Bay.contract.UnpackLog(event, "BayCreated", log); err != nil {
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

// ParseBayCreated is a log parse operation binding the contract event 0xafcc4615a5c6ce81e1ee1ecc9f150c249104ed97f69f737f6eb8fec74b8a91e4.
//
// Solidity: event BayCreated(address daoAddress, address admin)
func (_Bay *BayFilterer) ParseBayCreated(log types.Log) (*BayBayCreated, error) {
	event := new(BayBayCreated)
	if err := _Bay.contract.UnpackLog(event, "BayCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayBidCancelledIterator is returned from FilterBidCancelled and is used to iterate over the raw logs and unpacked data for BidCancelled events raised by the Bay contract.
type BayBidCancelledIterator struct {
	Event *BayBidCancelled // Event containing the contract specifics and raw log

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
func (it *BayBidCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayBidCancelled)
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
		it.Event = new(BayBidCancelled)
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
func (it *BayBidCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayBidCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayBidCancelled represents a BidCancelled event raised by the Bay contract.
type BayBidCancelled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidCancelled is a free log retrieval operation binding the contract event 0xb56dc4096011ba5fd2e46e5c3e7b04dec423b5e7b5fce9a17a419d77c832177c.
//
// Solidity: event BidCancelled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) FilterBidCancelled(opts *bind.FilterOpts, orderHash [][32]byte) (*BayBidCancelledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "BidCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &BayBidCancelledIterator{contract: _Bay.contract, event: "BidCancelled", logs: logs, sub: sub}, nil
}

// WatchBidCancelled is a free log subscription operation binding the contract event 0xb56dc4096011ba5fd2e46e5c3e7b04dec423b5e7b5fce9a17a419d77c832177c.
//
// Solidity: event BidCancelled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) WatchBidCancelled(opts *bind.WatchOpts, sink chan<- *BayBidCancelled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "BidCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayBidCancelled)
				if err := _Bay.contract.UnpackLog(event, "BidCancelled", log); err != nil {
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

// ParseBidCancelled is a log parse operation binding the contract event 0xb56dc4096011ba5fd2e46e5c3e7b04dec423b5e7b5fce9a17a419d77c832177c.
//
// Solidity: event BidCancelled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) ParseBidCancelled(log types.Log) (*BayBidCancelled, error) {
	event := new(BayBidCancelled)
	if err := _Bay.contract.UnpackLog(event, "BidCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayBidCreatedIterator is returned from FilterBidCreated and is used to iterate over the raw logs and unpacked data for BidCreated events raised by the Bay contract.
type BayBidCreatedIterator struct {
	Event *BayBidCreated // Event containing the contract specifics and raw log

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
func (it *BayBidCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayBidCreated)
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
		it.Event = new(BayBidCreated)
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
func (it *BayBidCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayBidCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayBidCreated represents a BidCreated event raised by the Bay contract.
type BayBidCreated struct {
	OrderHash     [32]byte
	Creator       common.Address
	Erc721Address common.Address
	TokenId       *big.Int
	Erc20Address  common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBidCreated is a free log retrieval operation binding the contract event 0xb70335faa97c8bf1fd6727d01143b2c03fd75096ed19e11a7cd4764a89ae21ef.
//
// Solidity: event BidCreated(bytes32 indexed orderHash, address indexed creator, address indexed erc721Address, uint256 tokenId, address erc20Address, uint256 amount)
func (_Bay *BayFilterer) FilterBidCreated(opts *bind.FilterOpts, orderHash [][32]byte, creator []common.Address, erc721Address []common.Address) (*BayBidCreatedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "BidCreated", orderHashRule, creatorRule, erc721AddressRule)
	if err != nil {
		return nil, err
	}
	return &BayBidCreatedIterator{contract: _Bay.contract, event: "BidCreated", logs: logs, sub: sub}, nil
}

// WatchBidCreated is a free log subscription operation binding the contract event 0xb70335faa97c8bf1fd6727d01143b2c03fd75096ed19e11a7cd4764a89ae21ef.
//
// Solidity: event BidCreated(bytes32 indexed orderHash, address indexed creator, address indexed erc721Address, uint256 tokenId, address erc20Address, uint256 amount)
func (_Bay *BayFilterer) WatchBidCreated(opts *bind.WatchOpts, sink chan<- *BayBidCreated, orderHash [][32]byte, creator []common.Address, erc721Address []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "BidCreated", orderHashRule, creatorRule, erc721AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayBidCreated)
				if err := _Bay.contract.UnpackLog(event, "BidCreated", log); err != nil {
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

// ParseBidCreated is a log parse operation binding the contract event 0xb70335faa97c8bf1fd6727d01143b2c03fd75096ed19e11a7cd4764a89ae21ef.
//
// Solidity: event BidCreated(bytes32 indexed orderHash, address indexed creator, address indexed erc721Address, uint256 tokenId, address erc20Address, uint256 amount)
func (_Bay *BayFilterer) ParseBidCreated(log types.Log) (*BayBidCreated, error) {
	event := new(BayBidCreated)
	if err := _Bay.contract.UnpackLog(event, "BidCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Bay contract.
type BayFeeRecipientUpdatedIterator struct {
	Event *BayFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *BayFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayFeeRecipientUpdated)
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
		it.Event = new(BayFeeRecipientUpdated)
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
func (it *BayFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Bay contract.
type BayFeeRecipientUpdated struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address newFeeRecipient)
func (_Bay *BayFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts) (*BayFeeRecipientUpdatedIterator, error) {

	logs, sub, err := _Bay.contract.FilterLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return &BayFeeRecipientUpdatedIterator{contract: _Bay.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address newFeeRecipient)
func (_Bay *BayFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *BayFeeRecipientUpdated) (event.Subscription, error) {

	logs, sub, err := _Bay.contract.WatchLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayFeeRecipientUpdated)
				if err := _Bay.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0x7a7b5a0a132f9e0581eb8527f66eae9ee89c2a3e79d4ac7e41a1f1f4d48a7fc2.
//
// Solidity: event FeeRecipientUpdated(address newFeeRecipient)
func (_Bay *BayFilterer) ParseFeeRecipientUpdated(log types.Log) (*BayFeeRecipientUpdated, error) {
	event := new(BayFeeRecipientUpdated)
	if err := _Bay.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayOrderFeeUpdatedIterator is returned from FilterOrderFeeUpdated and is used to iterate over the raw logs and unpacked data for OrderFeeUpdated events raised by the Bay contract.
type BayOrderFeeUpdatedIterator struct {
	Event *BayOrderFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BayOrderFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayOrderFeeUpdated)
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
		it.Event = new(BayOrderFeeUpdated)
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
func (it *BayOrderFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayOrderFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayOrderFeeUpdated represents a OrderFeeUpdated event raised by the Bay contract.
type BayOrderFeeUpdated struct {
	NewOrderFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderFeeUpdated is a free log retrieval operation binding the contract event 0xafb0bb88fecd164b1de860279e72dadcc296d3b91d41cef931856f3416cdd1db.
//
// Solidity: event OrderFeeUpdated(uint256 newOrderFee)
func (_Bay *BayFilterer) FilterOrderFeeUpdated(opts *bind.FilterOpts) (*BayOrderFeeUpdatedIterator, error) {

	logs, sub, err := _Bay.contract.FilterLogs(opts, "OrderFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BayOrderFeeUpdatedIterator{contract: _Bay.contract, event: "OrderFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchOrderFeeUpdated is a free log subscription operation binding the contract event 0xafb0bb88fecd164b1de860279e72dadcc296d3b91d41cef931856f3416cdd1db.
//
// Solidity: event OrderFeeUpdated(uint256 newOrderFee)
func (_Bay *BayFilterer) WatchOrderFeeUpdated(opts *bind.WatchOpts, sink chan<- *BayOrderFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Bay.contract.WatchLogs(opts, "OrderFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayOrderFeeUpdated)
				if err := _Bay.contract.UnpackLog(event, "OrderFeeUpdated", log); err != nil {
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

// ParseOrderFeeUpdated is a log parse operation binding the contract event 0xafb0bb88fecd164b1de860279e72dadcc296d3b91d41cef931856f3416cdd1db.
//
// Solidity: event OrderFeeUpdated(uint256 newOrderFee)
func (_Bay *BayFilterer) ParseOrderFeeUpdated(log types.Log) (*BayOrderFeeUpdated, error) {
	event := new(BayOrderFeeUpdated)
	if err := _Bay.contract.UnpackLog(event, "OrderFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Bay contract.
type BayOrderFulfilledIterator struct {
	Event *BayOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *BayOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayOrderFulfilled)
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
		it.Event = new(BayOrderFulfilled)
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
func (it *BayOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayOrderFulfilled represents a OrderFulfilled event raised by the Bay contract.
type BayOrderFulfilled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, orderHash [][32]byte) (*BayOrderFulfilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "OrderFulfilled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &BayOrderFulfilledIterator{contract: _Bay.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *BayOrderFulfilled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "OrderFulfilled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayOrderFulfilled)
				if err := _Bay.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// ParseOrderFulfilled is a log parse operation binding the contract event 0x48224e98144457f56f46e15959589d8155a9e29ca79439ab9ba37f60561b5c56.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderHash)
func (_Bay *BayFilterer) ParseOrderFulfilled(log types.Log) (*BayOrderFulfilled, error) {
	event := new(BayOrderFulfilled)
	if err := _Bay.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bay contract.
type BayRoleAdminChangedIterator struct {
	Event *BayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayRoleAdminChanged)
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
		it.Event = new(BayRoleAdminChanged)
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
func (it *BayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayRoleAdminChanged represents a RoleAdminChanged event raised by the Bay contract.
type BayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bay *BayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BayRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BayRoleAdminChangedIterator{contract: _Bay.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bay *BayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayRoleAdminChanged)
				if err := _Bay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bay *BayFilterer) ParseRoleAdminChanged(log types.Log) (*BayRoleAdminChanged, error) {
	event := new(BayRoleAdminChanged)
	if err := _Bay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bay contract.
type BayRoleGrantedIterator struct {
	Event *BayRoleGranted // Event containing the contract specifics and raw log

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
func (it *BayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayRoleGranted)
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
		it.Event = new(BayRoleGranted)
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
func (it *BayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayRoleGranted represents a RoleGranted event raised by the Bay contract.
type BayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bay *BayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BayRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BayRoleGrantedIterator{contract: _Bay.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bay *BayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayRoleGranted)
				if err := _Bay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bay *BayFilterer) ParseRoleGranted(log types.Log) (*BayRoleGranted, error) {
	event := new(BayRoleGranted)
	if err := _Bay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bay contract.
type BayRoleRevokedIterator struct {
	Event *BayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BayRoleRevoked)
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
		it.Event = new(BayRoleRevoked)
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
func (it *BayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BayRoleRevoked represents a RoleRevoked event raised by the Bay contract.
type BayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bay *BayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BayRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bay.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BayRoleRevokedIterator{contract: _Bay.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bay *BayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bay.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BayRoleRevoked)
				if err := _Bay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bay *BayFilterer) ParseRoleRevoked(log types.Log) (*BayRoleRevoked, error) {
	event := new(BayRoleRevoked)
	if err := _Bay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
