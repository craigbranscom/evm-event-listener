// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dao

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

// DaoMetaData contains all meta data concerning the Dao contract.
var DaoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"daoName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"daoMetadata_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposalFactoryAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalCost_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ActionExecutionFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"}],\"name\":\"ActionNotLoaded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"readyTime\",\"type\":\"uint256\"}],\"name\":\"ActionNotReady\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sig\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ready\",\"type\":\"uint256\"}],\"name\":\"ActionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sig\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ready\",\"type\":\"uint256\"}],\"name\":\"ActionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sig\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ready\",\"type\":\"uint256\"}],\"name\":\"ActionLoaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposalFactoryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"DAOCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Contract\",\"type\":\"address\"}],\"name\":\"ERC721Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Contract\",\"type\":\"address\"}],\"name\":\"ERC721Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"MemberInvited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"memberAddress\",\"type\":\"address\"}],\"name\":\"MemberRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"memberAddress\",\"type\":\"address\"}],\"name\":\"MemberUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposal\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"quorumPassed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportPassed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumVotesThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supportVotesThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalEvaluated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalTitle\",\"type\":\"string\"}],\"name\":\"draftProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"evaluateProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMemberStatus\",\"outputs\":[{\"internalType\":\"enumEnums.MemberStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"inviteMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sig\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"ready\",\"type\":\"uint256\"}],\"name\":\"isActionLoaded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposal\",\"type\":\"address\"}],\"name\":\"isProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"isTokenOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"loadedActions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVotePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"registerMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newDAOMetadata\",\"type\":\"string\"}],\"name\":\"setDAOMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newDAOName\",\"type\":\"string\"}],\"name\":\"setDAOName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinVotePeriod\",\"type\":\"uint256\"}],\"name\":\"setMinVotePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumThreshold\",\"type\":\"uint256\"}],\"name\":\"setQuorumThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSupportThreshold\",\"type\":\"uint256\"}],\"name\":\"setSupportThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unregisterMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DaoABI is the input ABI used to generate the binding from.
// Deprecated: Use DaoMetaData.ABI instead.
var DaoABI = DaoMetaData.ABI

// Dao is an auto generated Go binding around an Ethereum contract.
type Dao struct {
	DaoCaller     // Read-only binding to the contract
	DaoTransactor // Write-only binding to the contract
	DaoFilterer   // Log filterer for contract events
}

// DaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaoSession struct {
	Contract     *Dao              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaoCallerSession struct {
	Contract *DaoCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaoTransactorSession struct {
	Contract     *DaoTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaoRaw struct {
	Contract *Dao // Generic contract binding to access the raw methods on
}

// DaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaoCallerRaw struct {
	Contract *DaoCaller // Generic read-only contract binding to access the raw methods on
}

// DaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaoTransactorRaw struct {
	Contract *DaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDao creates a new instance of Dao, bound to a specific deployed contract.
func NewDao(address common.Address, backend bind.ContractBackend) (*Dao, error) {
	contract, err := bindDao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dao{DaoCaller: DaoCaller{contract: contract}, DaoTransactor: DaoTransactor{contract: contract}, DaoFilterer: DaoFilterer{contract: contract}}, nil
}

// NewDaoCaller creates a new read-only instance of Dao, bound to a specific deployed contract.
func NewDaoCaller(address common.Address, caller bind.ContractCaller) (*DaoCaller, error) {
	contract, err := bindDao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoCaller{contract: contract}, nil
}

// NewDaoTransactor creates a new write-only instance of Dao, bound to a specific deployed contract.
func NewDaoTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoTransactor, error) {
	contract, err := bindDao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoTransactor{contract: contract}, nil
}

// NewDaoFilterer creates a new log filterer instance of Dao, bound to a specific deployed contract.
func NewDaoFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoFilterer, error) {
	contract, err := bindDao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoFilterer{contract: contract}, nil
}

// bindDao binds a generic wrapper to an already deployed contract.
func bindDao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dao *DaoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dao.Contract.DaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dao *DaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.Contract.DaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dao *DaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dao.Contract.DaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dao *DaoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dao *DaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dao *DaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dao.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Dao *DaoCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Dao *DaoSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Dao.Contract.DEFAULTADMINROLE(&_Dao.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Dao *DaoCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Dao.Contract.DEFAULTADMINROLE(&_Dao.CallOpts)
}

// DaoMetadata is a free data retrieval call binding the contract method 0x0787e15c.
//
// Solidity: function daoMetadata() view returns(string)
func (_Dao *DaoCaller) DaoMetadata(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "daoMetadata")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DaoMetadata is a free data retrieval call binding the contract method 0x0787e15c.
//
// Solidity: function daoMetadata() view returns(string)
func (_Dao *DaoSession) DaoMetadata() (string, error) {
	return _Dao.Contract.DaoMetadata(&_Dao.CallOpts)
}

// DaoMetadata is a free data retrieval call binding the contract method 0x0787e15c.
//
// Solidity: function daoMetadata() view returns(string)
func (_Dao *DaoCallerSession) DaoMetadata() (string, error) {
	return _Dao.Contract.DaoMetadata(&_Dao.CallOpts)
}

// DaoName is a free data retrieval call binding the contract method 0x48976936.
//
// Solidity: function daoName() view returns(string)
func (_Dao *DaoCaller) DaoName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "daoName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DaoName is a free data retrieval call binding the contract method 0x48976936.
//
// Solidity: function daoName() view returns(string)
func (_Dao *DaoSession) DaoName() (string, error) {
	return _Dao.Contract.DaoName(&_Dao.CallOpts)
}

// DaoName is a free data retrieval call binding the contract method 0x48976936.
//
// Solidity: function daoName() view returns(string)
func (_Dao *DaoCallerSession) DaoName() (string, error) {
	return _Dao.Contract.DaoName(&_Dao.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Dao *DaoCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Dao *DaoSession) Delay() (*big.Int, error) {
	return _Dao.Contract.Delay(&_Dao.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Dao *DaoCallerSession) Delay() (*big.Int, error) {
	return _Dao.Contract.Delay(&_Dao.CallOpts)
}

// GetMemberStatus is a free data retrieval call binding the contract method 0x7d0d719b.
//
// Solidity: function getMemberStatus(uint256 tokenId) view returns(uint8)
func (_Dao *DaoCaller) GetMemberStatus(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "getMemberStatus", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMemberStatus is a free data retrieval call binding the contract method 0x7d0d719b.
//
// Solidity: function getMemberStatus(uint256 tokenId) view returns(uint8)
func (_Dao *DaoSession) GetMemberStatus(tokenId *big.Int) (uint8, error) {
	return _Dao.Contract.GetMemberStatus(&_Dao.CallOpts, tokenId)
}

// GetMemberStatus is a free data retrieval call binding the contract method 0x7d0d719b.
//
// Solidity: function getMemberStatus(uint256 tokenId) view returns(uint8)
func (_Dao *DaoCallerSession) GetMemberStatus(tokenId *big.Int) (uint8, error) {
	return _Dao.Contract.GetMemberStatus(&_Dao.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Dao *DaoCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Dao *DaoSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Dao.Contract.GetRoleAdmin(&_Dao.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Dao *DaoCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Dao.Contract.GetRoleAdmin(&_Dao.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Dao *DaoCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Dao *DaoSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Dao.Contract.HasRole(&_Dao.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Dao *DaoCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Dao.Contract.HasRole(&_Dao.CallOpts, role, account)
}

// IsActionLoaded is a free data retrieval call binding the contract method 0x51ff2b74.
//
// Solidity: function isActionLoaded(address target, uint256 value, string sig, bytes data, uint256 ready) view returns(bool)
func (_Dao *DaoCaller) IsActionLoaded(opts *bind.CallOpts, target common.Address, value *big.Int, sig string, data []byte, ready *big.Int) (bool, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "isActionLoaded", target, value, sig, data, ready)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActionLoaded is a free data retrieval call binding the contract method 0x51ff2b74.
//
// Solidity: function isActionLoaded(address target, uint256 value, string sig, bytes data, uint256 ready) view returns(bool)
func (_Dao *DaoSession) IsActionLoaded(target common.Address, value *big.Int, sig string, data []byte, ready *big.Int) (bool, error) {
	return _Dao.Contract.IsActionLoaded(&_Dao.CallOpts, target, value, sig, data, ready)
}

// IsActionLoaded is a free data retrieval call binding the contract method 0x51ff2b74.
//
// Solidity: function isActionLoaded(address target, uint256 value, string sig, bytes data, uint256 ready) view returns(bool)
func (_Dao *DaoCallerSession) IsActionLoaded(target common.Address, value *big.Int, sig string, data []byte, ready *big.Int) (bool, error) {
	return _Dao.Contract.IsActionLoaded(&_Dao.CallOpts, target, value, sig, data, ready)
}

// IsProposal is a free data retrieval call binding the contract method 0x52050003.
//
// Solidity: function isProposal(address proposal) view returns(bool)
func (_Dao *DaoCaller) IsProposal(opts *bind.CallOpts, proposal common.Address) (bool, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "isProposal", proposal)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposal is a free data retrieval call binding the contract method 0x52050003.
//
// Solidity: function isProposal(address proposal) view returns(bool)
func (_Dao *DaoSession) IsProposal(proposal common.Address) (bool, error) {
	return _Dao.Contract.IsProposal(&_Dao.CallOpts, proposal)
}

// IsProposal is a free data retrieval call binding the contract method 0x52050003.
//
// Solidity: function isProposal(address proposal) view returns(bool)
func (_Dao *DaoCallerSession) IsProposal(proposal common.Address) (bool, error) {
	return _Dao.Contract.IsProposal(&_Dao.CallOpts, proposal)
}

// IsTokenOwner is a free data retrieval call binding the contract method 0xd59f2827.
//
// Solidity: function isTokenOwner(uint256 tokenId, address member) view returns(bool)
func (_Dao *DaoCaller) IsTokenOwner(opts *bind.CallOpts, tokenId *big.Int, member common.Address) (bool, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "isTokenOwner", tokenId, member)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenOwner is a free data retrieval call binding the contract method 0xd59f2827.
//
// Solidity: function isTokenOwner(uint256 tokenId, address member) view returns(bool)
func (_Dao *DaoSession) IsTokenOwner(tokenId *big.Int, member common.Address) (bool, error) {
	return _Dao.Contract.IsTokenOwner(&_Dao.CallOpts, tokenId, member)
}

// IsTokenOwner is a free data retrieval call binding the contract method 0xd59f2827.
//
// Solidity: function isTokenOwner(uint256 tokenId, address member) view returns(bool)
func (_Dao *DaoCallerSession) IsTokenOwner(tokenId *big.Int, member common.Address) (bool, error) {
	return _Dao.Contract.IsTokenOwner(&_Dao.CallOpts, tokenId, member)
}

// LoadedActions is a free data retrieval call binding the contract method 0x6e8c7ab9.
//
// Solidity: function loadedActions(bytes32 ) view returns(bool)
func (_Dao *DaoCaller) LoadedActions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "loadedActions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LoadedActions is a free data retrieval call binding the contract method 0x6e8c7ab9.
//
// Solidity: function loadedActions(bytes32 ) view returns(bool)
func (_Dao *DaoSession) LoadedActions(arg0 [32]byte) (bool, error) {
	return _Dao.Contract.LoadedActions(&_Dao.CallOpts, arg0)
}

// LoadedActions is a free data retrieval call binding the contract method 0x6e8c7ab9.
//
// Solidity: function loadedActions(bytes32 ) view returns(bool)
func (_Dao *DaoCallerSession) LoadedActions(arg0 [32]byte) (bool, error) {
	return _Dao.Contract.LoadedActions(&_Dao.CallOpts, arg0)
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() view returns(uint256)
func (_Dao *DaoCaller) MemberCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "memberCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() view returns(uint256)
func (_Dao *DaoSession) MemberCount() (*big.Int, error) {
	return _Dao.Contract.MemberCount(&_Dao.CallOpts)
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() view returns(uint256)
func (_Dao *DaoCallerSession) MemberCount() (*big.Int, error) {
	return _Dao.Contract.MemberCount(&_Dao.CallOpts)
}

// MinVotePeriod is a free data retrieval call binding the contract method 0x4d6561c1.
//
// Solidity: function minVotePeriod() view returns(uint256)
func (_Dao *DaoCaller) MinVotePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "minVotePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotePeriod is a free data retrieval call binding the contract method 0x4d6561c1.
//
// Solidity: function minVotePeriod() view returns(uint256)
func (_Dao *DaoSession) MinVotePeriod() (*big.Int, error) {
	return _Dao.Contract.MinVotePeriod(&_Dao.CallOpts)
}

// MinVotePeriod is a free data retrieval call binding the contract method 0x4d6561c1.
//
// Solidity: function minVotePeriod() view returns(uint256)
func (_Dao *DaoCallerSession) MinVotePeriod() (*big.Int, error) {
	return _Dao.Contract.MinVotePeriod(&_Dao.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_Dao *DaoCaller) NftAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "nftAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_Dao *DaoSession) NftAddress() (common.Address, error) {
	return _Dao.Contract.NftAddress(&_Dao.CallOpts)
}

// NftAddress is a free data retrieval call binding the contract method 0x5bf8633a.
//
// Solidity: function nftAddress() view returns(address)
func (_Dao *DaoCallerSession) NftAddress() (common.Address, error) {
	return _Dao.Contract.NftAddress(&_Dao.CallOpts)
}

// ProposalCost is a free data retrieval call binding the contract method 0xe664f3b2.
//
// Solidity: function proposalCost() view returns(uint256)
func (_Dao *DaoCaller) ProposalCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "proposalCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCost is a free data retrieval call binding the contract method 0xe664f3b2.
//
// Solidity: function proposalCost() view returns(uint256)
func (_Dao *DaoSession) ProposalCost() (*big.Int, error) {
	return _Dao.Contract.ProposalCost(&_Dao.CallOpts)
}

// ProposalCost is a free data retrieval call binding the contract method 0xe664f3b2.
//
// Solidity: function proposalCost() view returns(uint256)
func (_Dao *DaoCallerSession) ProposalCost() (*big.Int, error) {
	return _Dao.Contract.ProposalCost(&_Dao.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Dao *DaoCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Dao *DaoSession) ProposalCount() (*big.Int, error) {
	return _Dao.Contract.ProposalCount(&_Dao.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Dao *DaoCallerSession) ProposalCount() (*big.Int, error) {
	return _Dao.Contract.ProposalCount(&_Dao.CallOpts)
}

// ProposalFactoryAddress is a free data retrieval call binding the contract method 0xd94f22ee.
//
// Solidity: function proposalFactoryAddress() view returns(address)
func (_Dao *DaoCaller) ProposalFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "proposalFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalFactoryAddress is a free data retrieval call binding the contract method 0xd94f22ee.
//
// Solidity: function proposalFactoryAddress() view returns(address)
func (_Dao *DaoSession) ProposalFactoryAddress() (common.Address, error) {
	return _Dao.Contract.ProposalFactoryAddress(&_Dao.CallOpts)
}

// ProposalFactoryAddress is a free data retrieval call binding the contract method 0xd94f22ee.
//
// Solidity: function proposalFactoryAddress() view returns(address)
func (_Dao *DaoCallerSession) ProposalFactoryAddress() (common.Address, error) {
	return _Dao.Contract.ProposalFactoryAddress(&_Dao.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address)
func (_Dao *DaoCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "proposals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address)
func (_Dao *DaoSession) Proposals(arg0 *big.Int) (common.Address, error) {
	return _Dao.Contract.Proposals(&_Dao.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address)
func (_Dao *DaoCallerSession) Proposals(arg0 *big.Int) (common.Address, error) {
	return _Dao.Contract.Proposals(&_Dao.CallOpts, arg0)
}

// QuorumThreshold is a free data retrieval call binding the contract method 0x7b7a91dd.
//
// Solidity: function quorumThreshold() view returns(uint256)
func (_Dao *DaoCaller) QuorumThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "quorumThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumThreshold is a free data retrieval call binding the contract method 0x7b7a91dd.
//
// Solidity: function quorumThreshold() view returns(uint256)
func (_Dao *DaoSession) QuorumThreshold() (*big.Int, error) {
	return _Dao.Contract.QuorumThreshold(&_Dao.CallOpts)
}

// QuorumThreshold is a free data retrieval call binding the contract method 0x7b7a91dd.
//
// Solidity: function quorumThreshold() view returns(uint256)
func (_Dao *DaoCallerSession) QuorumThreshold() (*big.Int, error) {
	return _Dao.Contract.QuorumThreshold(&_Dao.CallOpts)
}

// SupportThreshold is a free data retrieval call binding the contract method 0x7c36e8e8.
//
// Solidity: function supportThreshold() view returns(uint256)
func (_Dao *DaoCaller) SupportThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "supportThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupportThreshold is a free data retrieval call binding the contract method 0x7c36e8e8.
//
// Solidity: function supportThreshold() view returns(uint256)
func (_Dao *DaoSession) SupportThreshold() (*big.Int, error) {
	return _Dao.Contract.SupportThreshold(&_Dao.CallOpts)
}

// SupportThreshold is a free data retrieval call binding the contract method 0x7c36e8e8.
//
// Solidity: function supportThreshold() view returns(uint256)
func (_Dao *DaoCallerSession) SupportThreshold() (*big.Int, error) {
	return _Dao.Contract.SupportThreshold(&_Dao.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dao *DaoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dao.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dao *DaoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dao.Contract.SupportsInterface(&_Dao.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dao *DaoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dao.Contract.SupportsInterface(&_Dao.CallOpts, interfaceId)
}

// DraftProposal is a paid mutator transaction binding the contract method 0x9f367e08.
//
// Solidity: function draftProposal(string proposalTitle) payable returns(address)
func (_Dao *DaoTransactor) DraftProposal(opts *bind.TransactOpts, proposalTitle string) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "draftProposal", proposalTitle)
}

// DraftProposal is a paid mutator transaction binding the contract method 0x9f367e08.
//
// Solidity: function draftProposal(string proposalTitle) payable returns(address)
func (_Dao *DaoSession) DraftProposal(proposalTitle string) (*types.Transaction, error) {
	return _Dao.Contract.DraftProposal(&_Dao.TransactOpts, proposalTitle)
}

// DraftProposal is a paid mutator transaction binding the contract method 0x9f367e08.
//
// Solidity: function draftProposal(string proposalTitle) payable returns(address)
func (_Dao *DaoTransactorSession) DraftProposal(proposalTitle string) (*types.Transaction, error) {
	return _Dao.Contract.DraftProposal(&_Dao.TransactOpts, proposalTitle)
}

// EvaluateProposal is a paid mutator transaction binding the contract method 0x2df57dd6.
//
// Solidity: function evaluateProposal(address proposal) returns(bool)
func (_Dao *DaoTransactor) EvaluateProposal(opts *bind.TransactOpts, proposal common.Address) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "evaluateProposal", proposal)
}

// EvaluateProposal is a paid mutator transaction binding the contract method 0x2df57dd6.
//
// Solidity: function evaluateProposal(address proposal) returns(bool)
func (_Dao *DaoSession) EvaluateProposal(proposal common.Address) (*types.Transaction, error) {
	return _Dao.Contract.EvaluateProposal(&_Dao.TransactOpts, proposal)
}

// EvaluateProposal is a paid mutator transaction binding the contract method 0x2df57dd6.
//
// Solidity: function evaluateProposal(address proposal) returns(bool)
func (_Dao *DaoTransactorSession) EvaluateProposal(proposal common.Address) (*types.Transaction, error) {
	return _Dao.Contract.EvaluateProposal(&_Dao.TransactOpts, proposal)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xa67a03ab.
//
// Solidity: function executeProposal(address proposal) returns()
func (_Dao *DaoTransactor) ExecuteProposal(opts *bind.TransactOpts, proposal common.Address) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "executeProposal", proposal)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xa67a03ab.
//
// Solidity: function executeProposal(address proposal) returns()
func (_Dao *DaoSession) ExecuteProposal(proposal common.Address) (*types.Transaction, error) {
	return _Dao.Contract.ExecuteProposal(&_Dao.TransactOpts, proposal)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xa67a03ab.
//
// Solidity: function executeProposal(address proposal) returns()
func (_Dao *DaoTransactorSession) ExecuteProposal(proposal common.Address) (*types.Transaction, error) {
	return _Dao.Contract.ExecuteProposal(&_Dao.TransactOpts, proposal)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Dao *DaoTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Dao *DaoSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.Contract.GrantRole(&_Dao.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Dao *DaoTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.Contract.GrantRole(&_Dao.TransactOpts, role, account)
}

// InviteMember is a paid mutator transaction binding the contract method 0xbc5862a4.
//
// Solidity: function inviteMember(uint256 tokenId) returns()
func (_Dao *DaoTransactor) InviteMember(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "inviteMember", tokenId)
}

// InviteMember is a paid mutator transaction binding the contract method 0xbc5862a4.
//
// Solidity: function inviteMember(uint256 tokenId) returns()
func (_Dao *DaoSession) InviteMember(tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.InviteMember(&_Dao.TransactOpts, tokenId)
}

// InviteMember is a paid mutator transaction binding the contract method 0xbc5862a4.
//
// Solidity: function inviteMember(uint256 tokenId) returns()
func (_Dao *DaoTransactorSession) InviteMember(tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.InviteMember(&_Dao.TransactOpts, tokenId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Dao *DaoTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Dao *DaoSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dao.Contract.OnERC721Received(&_Dao.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Dao *DaoTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dao.Contract.OnERC721Received(&_Dao.TransactOpts, operator, from, tokenId, data)
}

// RegisterMember is a paid mutator transaction binding the contract method 0x22e994f3.
//
// Solidity: function registerMember(uint256 tokenId) returns()
func (_Dao *DaoTransactor) RegisterMember(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "registerMember", tokenId)
}

// RegisterMember is a paid mutator transaction binding the contract method 0x22e994f3.
//
// Solidity: function registerMember(uint256 tokenId) returns()
func (_Dao *DaoSession) RegisterMember(tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.RegisterMember(&_Dao.TransactOpts, tokenId)
}

// RegisterMember is a paid mutator transaction binding the contract method 0x22e994f3.
//
// Solidity: function registerMember(uint256 tokenId) returns()
func (_Dao *DaoTransactorSession) RegisterMember(tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.RegisterMember(&_Dao.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Dao *DaoTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Dao *DaoSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.Contract.RenounceRole(&_Dao.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Dao *DaoTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.Contract.RenounceRole(&_Dao.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Dao *DaoTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Dao *DaoSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.Contract.RevokeRole(&_Dao.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Dao *DaoTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Dao.Contract.RevokeRole(&_Dao.TransactOpts, role, account)
}

// SetDAOMetadata is a paid mutator transaction binding the contract method 0x977d1e18.
//
// Solidity: function setDAOMetadata(string newDAOMetadata) returns()
func (_Dao *DaoTransactor) SetDAOMetadata(opts *bind.TransactOpts, newDAOMetadata string) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "setDAOMetadata", newDAOMetadata)
}

// SetDAOMetadata is a paid mutator transaction binding the contract method 0x977d1e18.
//
// Solidity: function setDAOMetadata(string newDAOMetadata) returns()
func (_Dao *DaoSession) SetDAOMetadata(newDAOMetadata string) (*types.Transaction, error) {
	return _Dao.Contract.SetDAOMetadata(&_Dao.TransactOpts, newDAOMetadata)
}

// SetDAOMetadata is a paid mutator transaction binding the contract method 0x977d1e18.
//
// Solidity: function setDAOMetadata(string newDAOMetadata) returns()
func (_Dao *DaoTransactorSession) SetDAOMetadata(newDAOMetadata string) (*types.Transaction, error) {
	return _Dao.Contract.SetDAOMetadata(&_Dao.TransactOpts, newDAOMetadata)
}

// SetDAOName is a paid mutator transaction binding the contract method 0xa1de78e2.
//
// Solidity: function setDAOName(string newDAOName) returns()
func (_Dao *DaoTransactor) SetDAOName(opts *bind.TransactOpts, newDAOName string) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "setDAOName", newDAOName)
}

// SetDAOName is a paid mutator transaction binding the contract method 0xa1de78e2.
//
// Solidity: function setDAOName(string newDAOName) returns()
func (_Dao *DaoSession) SetDAOName(newDAOName string) (*types.Transaction, error) {
	return _Dao.Contract.SetDAOName(&_Dao.TransactOpts, newDAOName)
}

// SetDAOName is a paid mutator transaction binding the contract method 0xa1de78e2.
//
// Solidity: function setDAOName(string newDAOName) returns()
func (_Dao *DaoTransactorSession) SetDAOName(newDAOName string) (*types.Transaction, error) {
	return _Dao.Contract.SetDAOName(&_Dao.TransactOpts, newDAOName)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 newDelay) returns()
func (_Dao *DaoTransactor) SetDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "setDelay", newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 newDelay) returns()
func (_Dao *DaoSession) SetDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetDelay(&_Dao.TransactOpts, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 newDelay) returns()
func (_Dao *DaoTransactorSession) SetDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetDelay(&_Dao.TransactOpts, newDelay)
}

// SetMinVotePeriod is a paid mutator transaction binding the contract method 0x6c78a1c2.
//
// Solidity: function setMinVotePeriod(uint256 newMinVotePeriod) returns()
func (_Dao *DaoTransactor) SetMinVotePeriod(opts *bind.TransactOpts, newMinVotePeriod *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "setMinVotePeriod", newMinVotePeriod)
}

// SetMinVotePeriod is a paid mutator transaction binding the contract method 0x6c78a1c2.
//
// Solidity: function setMinVotePeriod(uint256 newMinVotePeriod) returns()
func (_Dao *DaoSession) SetMinVotePeriod(newMinVotePeriod *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetMinVotePeriod(&_Dao.TransactOpts, newMinVotePeriod)
}

// SetMinVotePeriod is a paid mutator transaction binding the contract method 0x6c78a1c2.
//
// Solidity: function setMinVotePeriod(uint256 newMinVotePeriod) returns()
func (_Dao *DaoTransactorSession) SetMinVotePeriod(newMinVotePeriod *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetMinVotePeriod(&_Dao.TransactOpts, newMinVotePeriod)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x4524c7e1.
//
// Solidity: function setQuorumThreshold(uint256 newQuorumThreshold) returns()
func (_Dao *DaoTransactor) SetQuorumThreshold(opts *bind.TransactOpts, newQuorumThreshold *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "setQuorumThreshold", newQuorumThreshold)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x4524c7e1.
//
// Solidity: function setQuorumThreshold(uint256 newQuorumThreshold) returns()
func (_Dao *DaoSession) SetQuorumThreshold(newQuorumThreshold *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetQuorumThreshold(&_Dao.TransactOpts, newQuorumThreshold)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x4524c7e1.
//
// Solidity: function setQuorumThreshold(uint256 newQuorumThreshold) returns()
func (_Dao *DaoTransactorSession) SetQuorumThreshold(newQuorumThreshold *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetQuorumThreshold(&_Dao.TransactOpts, newQuorumThreshold)
}

// SetSupportThreshold is a paid mutator transaction binding the contract method 0xfefd778e.
//
// Solidity: function setSupportThreshold(uint256 newSupportThreshold) returns()
func (_Dao *DaoTransactor) SetSupportThreshold(opts *bind.TransactOpts, newSupportThreshold *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "setSupportThreshold", newSupportThreshold)
}

// SetSupportThreshold is a paid mutator transaction binding the contract method 0xfefd778e.
//
// Solidity: function setSupportThreshold(uint256 newSupportThreshold) returns()
func (_Dao *DaoSession) SetSupportThreshold(newSupportThreshold *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetSupportThreshold(&_Dao.TransactOpts, newSupportThreshold)
}

// SetSupportThreshold is a paid mutator transaction binding the contract method 0xfefd778e.
//
// Solidity: function setSupportThreshold(uint256 newSupportThreshold) returns()
func (_Dao *DaoTransactorSession) SetSupportThreshold(newSupportThreshold *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.SetSupportThreshold(&_Dao.TransactOpts, newSupportThreshold)
}

// UnregisterMember is a paid mutator transaction binding the contract method 0x23dd523d.
//
// Solidity: function unregisterMember(uint256 tokenId) returns()
func (_Dao *DaoTransactor) UnregisterMember(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.contract.Transact(opts, "unregisterMember", tokenId)
}

// UnregisterMember is a paid mutator transaction binding the contract method 0x23dd523d.
//
// Solidity: function unregisterMember(uint256 tokenId) returns()
func (_Dao *DaoSession) UnregisterMember(tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.UnregisterMember(&_Dao.TransactOpts, tokenId)
}

// UnregisterMember is a paid mutator transaction binding the contract method 0x23dd523d.
//
// Solidity: function unregisterMember(uint256 tokenId) returns()
func (_Dao *DaoTransactorSession) UnregisterMember(tokenId *big.Int) (*types.Transaction, error) {
	return _Dao.Contract.UnregisterMember(&_Dao.TransactOpts, tokenId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Dao *DaoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Dao.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Dao *DaoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Dao.Contract.Fallback(&_Dao.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Dao *DaoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Dao.Contract.Fallback(&_Dao.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dao *DaoTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dao.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dao *DaoSession) Receive() (*types.Transaction, error) {
	return _Dao.Contract.Receive(&_Dao.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Dao *DaoTransactorSession) Receive() (*types.Transaction, error) {
	return _Dao.Contract.Receive(&_Dao.TransactOpts)
}

// DaoActionCancelledIterator is returned from FilterActionCancelled and is used to iterate over the raw logs and unpacked data for ActionCancelled events raised by the Dao contract.
type DaoActionCancelledIterator struct {
	Event *DaoActionCancelled // Event containing the contract specifics and raw log

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
func (it *DaoActionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoActionCancelled)
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
		it.Event = new(DaoActionCancelled)
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
func (it *DaoActionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoActionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoActionCancelled represents a ActionCancelled event raised by the Dao contract.
type DaoActionCancelled struct {
	TrxHash [32]byte
	Target  common.Address
	Value   *big.Int
	Sig     string
	Data    []byte
	Ready   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActionCancelled is a free log retrieval operation binding the contract event 0x65fc52b902965c1a53e05e64a9d7823fe42907aa394f4ed34111f0fbeea39d32.
//
// Solidity: event ActionCancelled(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) FilterActionCancelled(opts *bind.FilterOpts) (*DaoActionCancelledIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ActionCancelled")
	if err != nil {
		return nil, err
	}
	return &DaoActionCancelledIterator{contract: _Dao.contract, event: "ActionCancelled", logs: logs, sub: sub}, nil
}

// WatchActionCancelled is a free log subscription operation binding the contract event 0x65fc52b902965c1a53e05e64a9d7823fe42907aa394f4ed34111f0fbeea39d32.
//
// Solidity: event ActionCancelled(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) WatchActionCancelled(opts *bind.WatchOpts, sink chan<- *DaoActionCancelled) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ActionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoActionCancelled)
				if err := _Dao.contract.UnpackLog(event, "ActionCancelled", log); err != nil {
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

// ParseActionCancelled is a log parse operation binding the contract event 0x65fc52b902965c1a53e05e64a9d7823fe42907aa394f4ed34111f0fbeea39d32.
//
// Solidity: event ActionCancelled(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) ParseActionCancelled(log types.Log) (*DaoActionCancelled, error) {
	event := new(DaoActionCancelled)
	if err := _Dao.contract.UnpackLog(event, "ActionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoActionExecutedIterator is returned from FilterActionExecuted and is used to iterate over the raw logs and unpacked data for ActionExecuted events raised by the Dao contract.
type DaoActionExecutedIterator struct {
	Event *DaoActionExecuted // Event containing the contract specifics and raw log

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
func (it *DaoActionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoActionExecuted)
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
		it.Event = new(DaoActionExecuted)
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
func (it *DaoActionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoActionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoActionExecuted represents a ActionExecuted event raised by the Dao contract.
type DaoActionExecuted struct {
	TrxHash [32]byte
	Target  common.Address
	Value   *big.Int
	Sig     string
	Data    []byte
	Ready   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActionExecuted is a free log retrieval operation binding the contract event 0xb9c70c38e172c4486313c999fbf9d4e951ee83219d0ec61c57547959f089ad37.
//
// Solidity: event ActionExecuted(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) FilterActionExecuted(opts *bind.FilterOpts) (*DaoActionExecutedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ActionExecuted")
	if err != nil {
		return nil, err
	}
	return &DaoActionExecutedIterator{contract: _Dao.contract, event: "ActionExecuted", logs: logs, sub: sub}, nil
}

// WatchActionExecuted is a free log subscription operation binding the contract event 0xb9c70c38e172c4486313c999fbf9d4e951ee83219d0ec61c57547959f089ad37.
//
// Solidity: event ActionExecuted(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) WatchActionExecuted(opts *bind.WatchOpts, sink chan<- *DaoActionExecuted) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ActionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoActionExecuted)
				if err := _Dao.contract.UnpackLog(event, "ActionExecuted", log); err != nil {
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

// ParseActionExecuted is a log parse operation binding the contract event 0xb9c70c38e172c4486313c999fbf9d4e951ee83219d0ec61c57547959f089ad37.
//
// Solidity: event ActionExecuted(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) ParseActionExecuted(log types.Log) (*DaoActionExecuted, error) {
	event := new(DaoActionExecuted)
	if err := _Dao.contract.UnpackLog(event, "ActionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoActionLoadedIterator is returned from FilterActionLoaded and is used to iterate over the raw logs and unpacked data for ActionLoaded events raised by the Dao contract.
type DaoActionLoadedIterator struct {
	Event *DaoActionLoaded // Event containing the contract specifics and raw log

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
func (it *DaoActionLoadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoActionLoaded)
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
		it.Event = new(DaoActionLoaded)
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
func (it *DaoActionLoadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoActionLoadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoActionLoaded represents a ActionLoaded event raised by the Dao contract.
type DaoActionLoaded struct {
	TrxHash [32]byte
	Target  common.Address
	Value   *big.Int
	Sig     string
	Data    []byte
	Ready   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActionLoaded is a free log retrieval operation binding the contract event 0xc033a419f7d89fe56731c5063af781a0385d4bc7cd66e2ce68fda5c52abb9efd.
//
// Solidity: event ActionLoaded(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) FilterActionLoaded(opts *bind.FilterOpts) (*DaoActionLoadedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ActionLoaded")
	if err != nil {
		return nil, err
	}
	return &DaoActionLoadedIterator{contract: _Dao.contract, event: "ActionLoaded", logs: logs, sub: sub}, nil
}

// WatchActionLoaded is a free log subscription operation binding the contract event 0xc033a419f7d89fe56731c5063af781a0385d4bc7cd66e2ce68fda5c52abb9efd.
//
// Solidity: event ActionLoaded(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) WatchActionLoaded(opts *bind.WatchOpts, sink chan<- *DaoActionLoaded) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ActionLoaded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoActionLoaded)
				if err := _Dao.contract.UnpackLog(event, "ActionLoaded", log); err != nil {
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

// ParseActionLoaded is a log parse operation binding the contract event 0xc033a419f7d89fe56731c5063af781a0385d4bc7cd66e2ce68fda5c52abb9efd.
//
// Solidity: event ActionLoaded(bytes32 trxHash, address target, uint256 value, string sig, bytes data, uint256 ready)
func (_Dao *DaoFilterer) ParseActionLoaded(log types.Log) (*DaoActionLoaded, error) {
	event := new(DaoActionLoaded)
	if err := _Dao.contract.UnpackLog(event, "ActionLoaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoDAOCreatedIterator is returned from FilterDAOCreated and is used to iterate over the raw logs and unpacked data for DAOCreated events raised by the Dao contract.
type DaoDAOCreatedIterator struct {
	Event *DaoDAOCreated // Event containing the contract specifics and raw log

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
func (it *DaoDAOCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoDAOCreated)
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
		it.Event = new(DaoDAOCreated)
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
func (it *DaoDAOCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoDAOCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoDAOCreated represents a DAOCreated event raised by the Dao contract.
type DaoDAOCreated struct {
	NftAddress             common.Address
	ProposalFactoryAddress common.Address
	Admin                  common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDAOCreated is a free log retrieval operation binding the contract event 0x552b78c92c2932581cee967b6925eea60f5d345efa257d0432a333a6b55dc2b1.
//
// Solidity: event DAOCreated(address nftAddress, address proposalFactoryAddress, address admin)
func (_Dao *DaoFilterer) FilterDAOCreated(opts *bind.FilterOpts) (*DaoDAOCreatedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "DAOCreated")
	if err != nil {
		return nil, err
	}
	return &DaoDAOCreatedIterator{contract: _Dao.contract, event: "DAOCreated", logs: logs, sub: sub}, nil
}

// WatchDAOCreated is a free log subscription operation binding the contract event 0x552b78c92c2932581cee967b6925eea60f5d345efa257d0432a333a6b55dc2b1.
//
// Solidity: event DAOCreated(address nftAddress, address proposalFactoryAddress, address admin)
func (_Dao *DaoFilterer) WatchDAOCreated(opts *bind.WatchOpts, sink chan<- *DaoDAOCreated) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "DAOCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoDAOCreated)
				if err := _Dao.contract.UnpackLog(event, "DAOCreated", log); err != nil {
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

// ParseDAOCreated is a log parse operation binding the contract event 0x552b78c92c2932581cee967b6925eea60f5d345efa257d0432a333a6b55dc2b1.
//
// Solidity: event DAOCreated(address nftAddress, address proposalFactoryAddress, address admin)
func (_Dao *DaoFilterer) ParseDAOCreated(log types.Log) (*DaoDAOCreated, error) {
	event := new(DaoDAOCreated)
	if err := _Dao.contract.UnpackLog(event, "DAOCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoERC721ReceivedIterator is returned from FilterERC721Received and is used to iterate over the raw logs and unpacked data for ERC721Received events raised by the Dao contract.
type DaoERC721ReceivedIterator struct {
	Event *DaoERC721Received // Event containing the contract specifics and raw log

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
func (it *DaoERC721ReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoERC721Received)
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
		it.Event = new(DaoERC721Received)
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
func (it *DaoERC721ReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoERC721ReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoERC721Received represents a ERC721Received event raised by the Dao contract.
type DaoERC721Received struct {
	From           common.Address
	TokenId        *big.Int
	Erc721Contract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterERC721Received is a free log retrieval operation binding the contract event 0x53f9fb1a779fe0d4eee06280249fc20441cca6949207450cad7c5ef85de6ce23.
//
// Solidity: event ERC721Received(address from, uint256 tokenId, address erc721Contract)
func (_Dao *DaoFilterer) FilterERC721Received(opts *bind.FilterOpts) (*DaoERC721ReceivedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ERC721Received")
	if err != nil {
		return nil, err
	}
	return &DaoERC721ReceivedIterator{contract: _Dao.contract, event: "ERC721Received", logs: logs, sub: sub}, nil
}

// WatchERC721Received is a free log subscription operation binding the contract event 0x53f9fb1a779fe0d4eee06280249fc20441cca6949207450cad7c5ef85de6ce23.
//
// Solidity: event ERC721Received(address from, uint256 tokenId, address erc721Contract)
func (_Dao *DaoFilterer) WatchERC721Received(opts *bind.WatchOpts, sink chan<- *DaoERC721Received) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ERC721Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoERC721Received)
				if err := _Dao.contract.UnpackLog(event, "ERC721Received", log); err != nil {
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

// ParseERC721Received is a log parse operation binding the contract event 0x53f9fb1a779fe0d4eee06280249fc20441cca6949207450cad7c5ef85de6ce23.
//
// Solidity: event ERC721Received(address from, uint256 tokenId, address erc721Contract)
func (_Dao *DaoFilterer) ParseERC721Received(log types.Log) (*DaoERC721Received, error) {
	event := new(DaoERC721Received)
	if err := _Dao.contract.UnpackLog(event, "ERC721Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoERC721SentIterator is returned from FilterERC721Sent and is used to iterate over the raw logs and unpacked data for ERC721Sent events raised by the Dao contract.
type DaoERC721SentIterator struct {
	Event *DaoERC721Sent // Event containing the contract specifics and raw log

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
func (it *DaoERC721SentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoERC721Sent)
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
		it.Event = new(DaoERC721Sent)
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
func (it *DaoERC721SentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoERC721SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoERC721Sent represents a ERC721Sent event raised by the Dao contract.
type DaoERC721Sent struct {
	To             common.Address
	TokenId        *big.Int
	Erc721Contract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterERC721Sent is a free log retrieval operation binding the contract event 0xe766171301ed524b0850d38c1eeced46e559aeccfbfd994fc19bbb80ac8b1881.
//
// Solidity: event ERC721Sent(address to, uint256 tokenId, address erc721Contract)
func (_Dao *DaoFilterer) FilterERC721Sent(opts *bind.FilterOpts) (*DaoERC721SentIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ERC721Sent")
	if err != nil {
		return nil, err
	}
	return &DaoERC721SentIterator{contract: _Dao.contract, event: "ERC721Sent", logs: logs, sub: sub}, nil
}

// WatchERC721Sent is a free log subscription operation binding the contract event 0xe766171301ed524b0850d38c1eeced46e559aeccfbfd994fc19bbb80ac8b1881.
//
// Solidity: event ERC721Sent(address to, uint256 tokenId, address erc721Contract)
func (_Dao *DaoFilterer) WatchERC721Sent(opts *bind.WatchOpts, sink chan<- *DaoERC721Sent) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ERC721Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoERC721Sent)
				if err := _Dao.contract.UnpackLog(event, "ERC721Sent", log); err != nil {
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

// ParseERC721Sent is a log parse operation binding the contract event 0xe766171301ed524b0850d38c1eeced46e559aeccfbfd994fc19bbb80ac8b1881.
//
// Solidity: event ERC721Sent(address to, uint256 tokenId, address erc721Contract)
func (_Dao *DaoFilterer) ParseERC721Sent(log types.Log) (*DaoERC721Sent, error) {
	event := new(DaoERC721Sent)
	if err := _Dao.contract.UnpackLog(event, "ERC721Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoMemberInvitedIterator is returned from FilterMemberInvited and is used to iterate over the raw logs and unpacked data for MemberInvited events raised by the Dao contract.
type DaoMemberInvitedIterator struct {
	Event *DaoMemberInvited // Event containing the contract specifics and raw log

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
func (it *DaoMemberInvitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoMemberInvited)
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
		it.Event = new(DaoMemberInvited)
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
func (it *DaoMemberInvitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoMemberInvitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoMemberInvited represents a MemberInvited event raised by the Dao contract.
type DaoMemberInvited struct {
	DaoAddress common.Address
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMemberInvited is a free log retrieval operation binding the contract event 0x57ff8c3065de342518178781cf8cd82078e0253b4e031002aedd00047e1ff18f.
//
// Solidity: event MemberInvited(address daoAddress, address nftAddress, uint256 tokenId)
func (_Dao *DaoFilterer) FilterMemberInvited(opts *bind.FilterOpts) (*DaoMemberInvitedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "MemberInvited")
	if err != nil {
		return nil, err
	}
	return &DaoMemberInvitedIterator{contract: _Dao.contract, event: "MemberInvited", logs: logs, sub: sub}, nil
}

// WatchMemberInvited is a free log subscription operation binding the contract event 0x57ff8c3065de342518178781cf8cd82078e0253b4e031002aedd00047e1ff18f.
//
// Solidity: event MemberInvited(address daoAddress, address nftAddress, uint256 tokenId)
func (_Dao *DaoFilterer) WatchMemberInvited(opts *bind.WatchOpts, sink chan<- *DaoMemberInvited) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "MemberInvited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoMemberInvited)
				if err := _Dao.contract.UnpackLog(event, "MemberInvited", log); err != nil {
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

// ParseMemberInvited is a log parse operation binding the contract event 0x57ff8c3065de342518178781cf8cd82078e0253b4e031002aedd00047e1ff18f.
//
// Solidity: event MemberInvited(address daoAddress, address nftAddress, uint256 tokenId)
func (_Dao *DaoFilterer) ParseMemberInvited(log types.Log) (*DaoMemberInvited, error) {
	event := new(DaoMemberInvited)
	if err := _Dao.contract.UnpackLog(event, "MemberInvited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoMemberRegisteredIterator is returned from FilterMemberRegistered and is used to iterate over the raw logs and unpacked data for MemberRegistered events raised by the Dao contract.
type DaoMemberRegisteredIterator struct {
	Event *DaoMemberRegistered // Event containing the contract specifics and raw log

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
func (it *DaoMemberRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoMemberRegistered)
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
		it.Event = new(DaoMemberRegistered)
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
func (it *DaoMemberRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoMemberRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoMemberRegistered represents a MemberRegistered event raised by the Dao contract.
type DaoMemberRegistered struct {
	DaoAddress    common.Address
	NftAddress    common.Address
	TokenId       *big.Int
	MemberAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMemberRegistered is a free log retrieval operation binding the contract event 0x6d6a6e32530d4895a149a76159bcea283d7bbd32170c5e5a95b7c02a9098c38a.
//
// Solidity: event MemberRegistered(address daoAddress, address nftAddress, uint256 tokenId, address memberAddress)
func (_Dao *DaoFilterer) FilterMemberRegistered(opts *bind.FilterOpts) (*DaoMemberRegisteredIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return &DaoMemberRegisteredIterator{contract: _Dao.contract, event: "MemberRegistered", logs: logs, sub: sub}, nil
}

// WatchMemberRegistered is a free log subscription operation binding the contract event 0x6d6a6e32530d4895a149a76159bcea283d7bbd32170c5e5a95b7c02a9098c38a.
//
// Solidity: event MemberRegistered(address daoAddress, address nftAddress, uint256 tokenId, address memberAddress)
func (_Dao *DaoFilterer) WatchMemberRegistered(opts *bind.WatchOpts, sink chan<- *DaoMemberRegistered) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "MemberRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoMemberRegistered)
				if err := _Dao.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
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

// ParseMemberRegistered is a log parse operation binding the contract event 0x6d6a6e32530d4895a149a76159bcea283d7bbd32170c5e5a95b7c02a9098c38a.
//
// Solidity: event MemberRegistered(address daoAddress, address nftAddress, uint256 tokenId, address memberAddress)
func (_Dao *DaoFilterer) ParseMemberRegistered(log types.Log) (*DaoMemberRegistered, error) {
	event := new(DaoMemberRegistered)
	if err := _Dao.contract.UnpackLog(event, "MemberRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoMemberUnregisteredIterator is returned from FilterMemberUnregistered and is used to iterate over the raw logs and unpacked data for MemberUnregistered events raised by the Dao contract.
type DaoMemberUnregisteredIterator struct {
	Event *DaoMemberUnregistered // Event containing the contract specifics and raw log

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
func (it *DaoMemberUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoMemberUnregistered)
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
		it.Event = new(DaoMemberUnregistered)
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
func (it *DaoMemberUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoMemberUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoMemberUnregistered represents a MemberUnregistered event raised by the Dao contract.
type DaoMemberUnregistered struct {
	DaoAddress    common.Address
	NftAddress    common.Address
	TokenId       *big.Int
	MemberAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMemberUnregistered is a free log retrieval operation binding the contract event 0x83052bbd673d540f623d24c36b94005890bac88efc6fafce2010850d9b02ba46.
//
// Solidity: event MemberUnregistered(address daoAddress, address nftAddress, uint256 tokenId, address memberAddress)
func (_Dao *DaoFilterer) FilterMemberUnregistered(opts *bind.FilterOpts) (*DaoMemberUnregisteredIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "MemberUnregistered")
	if err != nil {
		return nil, err
	}
	return &DaoMemberUnregisteredIterator{contract: _Dao.contract, event: "MemberUnregistered", logs: logs, sub: sub}, nil
}

// WatchMemberUnregistered is a free log subscription operation binding the contract event 0x83052bbd673d540f623d24c36b94005890bac88efc6fafce2010850d9b02ba46.
//
// Solidity: event MemberUnregistered(address daoAddress, address nftAddress, uint256 tokenId, address memberAddress)
func (_Dao *DaoFilterer) WatchMemberUnregistered(opts *bind.WatchOpts, sink chan<- *DaoMemberUnregistered) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "MemberUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoMemberUnregistered)
				if err := _Dao.contract.UnpackLog(event, "MemberUnregistered", log); err != nil {
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

// ParseMemberUnregistered is a log parse operation binding the contract event 0x83052bbd673d540f623d24c36b94005890bac88efc6fafce2010850d9b02ba46.
//
// Solidity: event MemberUnregistered(address daoAddress, address nftAddress, uint256 tokenId, address memberAddress)
func (_Dao *DaoFilterer) ParseMemberUnregistered(log types.Log) (*DaoMemberUnregistered, error) {
	event := new(DaoMemberUnregistered)
	if err := _Dao.contract.UnpackLog(event, "MemberUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Dao contract.
type DaoProposalCreatedIterator struct {
	Event *DaoProposalCreated // Event containing the contract specifics and raw log

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
func (it *DaoProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoProposalCreated)
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
		it.Event = new(DaoProposalCreated)
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
func (it *DaoProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoProposalCreated represents a ProposalCreated event raised by the Dao contract.
type DaoProposalCreated struct {
	DaoAddress common.Address
	Proposal   common.Address
	ProposalId *big.Int
	Creator    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa33abb4d0226d76a56e89b18b3c8f3dec914fce06449b6ff646962c19ddc971e.
//
// Solidity: event ProposalCreated(address daoAddress, address proposal, uint256 proposalId, address creator)
func (_Dao *DaoFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*DaoProposalCreatedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &DaoProposalCreatedIterator{contract: _Dao.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa33abb4d0226d76a56e89b18b3c8f3dec914fce06449b6ff646962c19ddc971e.
//
// Solidity: event ProposalCreated(address daoAddress, address proposal, uint256 proposalId, address creator)
func (_Dao *DaoFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *DaoProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoProposalCreated)
				if err := _Dao.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xa33abb4d0226d76a56e89b18b3c8f3dec914fce06449b6ff646962c19ddc971e.
//
// Solidity: event ProposalCreated(address daoAddress, address proposal, uint256 proposalId, address creator)
func (_Dao *DaoFilterer) ParseProposalCreated(log types.Log) (*DaoProposalCreated, error) {
	event := new(DaoProposalCreated)
	if err := _Dao.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoProposalEvaluatedIterator is returned from FilterProposalEvaluated and is used to iterate over the raw logs and unpacked data for ProposalEvaluated events raised by the Dao contract.
type DaoProposalEvaluatedIterator struct {
	Event *DaoProposalEvaluated // Event containing the contract specifics and raw log

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
func (it *DaoProposalEvaluatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoProposalEvaluated)
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
		it.Event = new(DaoProposalEvaluated)
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
func (it *DaoProposalEvaluatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoProposalEvaluatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoProposalEvaluated represents a ProposalEvaluated event raised by the Dao contract.
type DaoProposalEvaluated struct {
	QuorumPassed          bool
	SupportPassed         bool
	TotalVotes            *big.Int
	QuorumVotesThreshold  *big.Int
	SupportVotesThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProposalEvaluated is a free log retrieval operation binding the contract event 0x835dbb3bf252734271b555e79649ca7817703fb8fbeb9f7a38f9d36b04a01786.
//
// Solidity: event ProposalEvaluated(bool quorumPassed, bool supportPassed, uint256 totalVotes, uint256 quorumVotesThreshold, uint256 supportVotesThreshold)
func (_Dao *DaoFilterer) FilterProposalEvaluated(opts *bind.FilterOpts) (*DaoProposalEvaluatedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ProposalEvaluated")
	if err != nil {
		return nil, err
	}
	return &DaoProposalEvaluatedIterator{contract: _Dao.contract, event: "ProposalEvaluated", logs: logs, sub: sub}, nil
}

// WatchProposalEvaluated is a free log subscription operation binding the contract event 0x835dbb3bf252734271b555e79649ca7817703fb8fbeb9f7a38f9d36b04a01786.
//
// Solidity: event ProposalEvaluated(bool quorumPassed, bool supportPassed, uint256 totalVotes, uint256 quorumVotesThreshold, uint256 supportVotesThreshold)
func (_Dao *DaoFilterer) WatchProposalEvaluated(opts *bind.WatchOpts, sink chan<- *DaoProposalEvaluated) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ProposalEvaluated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoProposalEvaluated)
				if err := _Dao.contract.UnpackLog(event, "ProposalEvaluated", log); err != nil {
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

// ParseProposalEvaluated is a log parse operation binding the contract event 0x835dbb3bf252734271b555e79649ca7817703fb8fbeb9f7a38f9d36b04a01786.
//
// Solidity: event ProposalEvaluated(bool quorumPassed, bool supportPassed, uint256 totalVotes, uint256 quorumVotesThreshold, uint256 supportVotesThreshold)
func (_Dao *DaoFilterer) ParseProposalEvaluated(log types.Log) (*DaoProposalEvaluated, error) {
	event := new(DaoProposalEvaluated)
	if err := _Dao.contract.UnpackLog(event, "ProposalEvaluated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Dao contract.
type DaoProposalExecutedIterator struct {
	Event *DaoProposalExecuted // Event containing the contract specifics and raw log

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
func (it *DaoProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoProposalExecuted)
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
		it.Event = new(DaoProposalExecuted)
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
func (it *DaoProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoProposalExecuted represents a ProposalExecuted event raised by the Dao contract.
type DaoProposalExecuted struct {
	Proposal common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x077e00e56e5d36ae097a95cb003427416253dd9bb4491393d3eb4cf8f63fd41e.
//
// Solidity: event ProposalExecuted(address proposal)
func (_Dao *DaoFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*DaoProposalExecutedIterator, error) {

	logs, sub, err := _Dao.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &DaoProposalExecutedIterator{contract: _Dao.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x077e00e56e5d36ae097a95cb003427416253dd9bb4491393d3eb4cf8f63fd41e.
//
// Solidity: event ProposalExecuted(address proposal)
func (_Dao *DaoFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *DaoProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Dao.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoProposalExecuted)
				if err := _Dao.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x077e00e56e5d36ae097a95cb003427416253dd9bb4491393d3eb4cf8f63fd41e.
//
// Solidity: event ProposalExecuted(address proposal)
func (_Dao *DaoFilterer) ParseProposalExecuted(log types.Log) (*DaoProposalExecuted, error) {
	event := new(DaoProposalExecuted)
	if err := _Dao.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Dao contract.
type DaoRoleAdminChangedIterator struct {
	Event *DaoRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DaoRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRoleAdminChanged)
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
		it.Event = new(DaoRoleAdminChanged)
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
func (it *DaoRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRoleAdminChanged represents a RoleAdminChanged event raised by the Dao contract.
type DaoRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Dao *DaoFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DaoRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Dao.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DaoRoleAdminChangedIterator{contract: _Dao.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Dao *DaoFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DaoRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Dao.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRoleAdminChanged)
				if err := _Dao.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Dao *DaoFilterer) ParseRoleAdminChanged(log types.Log) (*DaoRoleAdminChanged, error) {
	event := new(DaoRoleAdminChanged)
	if err := _Dao.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Dao contract.
type DaoRoleGrantedIterator struct {
	Event *DaoRoleGranted // Event containing the contract specifics and raw log

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
func (it *DaoRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRoleGranted)
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
		it.Event = new(DaoRoleGranted)
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
func (it *DaoRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRoleGranted represents a RoleGranted event raised by the Dao contract.
type DaoRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Dao *DaoFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DaoRoleGrantedIterator, error) {

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

	logs, sub, err := _Dao.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DaoRoleGrantedIterator{contract: _Dao.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Dao *DaoFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DaoRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dao.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRoleGranted)
				if err := _Dao.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Dao *DaoFilterer) ParseRoleGranted(log types.Log) (*DaoRoleGranted, error) {
	event := new(DaoRoleGranted)
	if err := _Dao.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Dao contract.
type DaoRoleRevokedIterator struct {
	Event *DaoRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DaoRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoRoleRevoked)
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
		it.Event = new(DaoRoleRevoked)
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
func (it *DaoRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoRoleRevoked represents a RoleRevoked event raised by the Dao contract.
type DaoRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Dao *DaoFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DaoRoleRevokedIterator, error) {

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

	logs, sub, err := _Dao.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DaoRoleRevokedIterator{contract: _Dao.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Dao *DaoFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DaoRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Dao.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoRoleRevoked)
				if err := _Dao.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Dao *DaoFilterer) ParseRoleRevoked(log types.Log) (*DaoRoleRevoked, error) {
	event := new(DaoRoleRevoked)
	if err := _Dao.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
