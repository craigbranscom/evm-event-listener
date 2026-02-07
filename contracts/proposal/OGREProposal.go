// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proposal

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

// StructsAction is an auto generated low-level Go binding around an user-defined struct.
type StructsAction struct {
	Target common.Address
	Value  *big.Int
	Sig    string
	Data   []byte
	Ready  *big.Int
}

// ProposalMetaData contains all meta data concerning the Proposal contract.
var ProposalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"proposalTitle_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"daoAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumEnums.ProposalStatus\",\"name\":\"currentStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumEnums.ProposalStatus\",\"name\":\"requiredStatus\",\"type\":\"uint8\"}],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"quorumPassed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportPassed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumVotesThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supportVotesThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalResults\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newStatus\",\"type\":\"string\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sig\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isRevotable\",\"type\":\"bool\"}],\"name\":\"configureProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sig\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"ready\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.Action\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revotable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"readyTime\",\"type\":\"uint256\"}],\"name\":\"setActionReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newProposalTitle\",\"type\":\"string\"}],\"name\":\"setProposalTitle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newEndTime\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumEnums.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalMetaData.ABI instead.
var ProposalABI = ProposalMetaData.ABI

// Proposal is an auto generated Go binding around an Ethereum contract.
type Proposal struct {
	ProposalCaller     // Read-only binding to the contract
	ProposalTransactor // Write-only binding to the contract
	ProposalFilterer   // Log filterer for contract events
}

// ProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalSession struct {
	Contract     *Proposal         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalCallerSession struct {
	Contract *ProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalTransactorSession struct {
	Contract     *ProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalRaw struct {
	Contract *Proposal // Generic contract binding to access the raw methods on
}

// ProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalCallerRaw struct {
	Contract *ProposalCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalTransactorRaw struct {
	Contract *ProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposal creates a new instance of Proposal, bound to a specific deployed contract.
func NewProposal(address common.Address, backend bind.ContractBackend) (*Proposal, error) {
	contract, err := bindProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proposal{ProposalCaller: ProposalCaller{contract: contract}, ProposalTransactor: ProposalTransactor{contract: contract}, ProposalFilterer: ProposalFilterer{contract: contract}}, nil
}

// NewProposalCaller creates a new read-only instance of Proposal, bound to a specific deployed contract.
func NewProposalCaller(address common.Address, caller bind.ContractCaller) (*ProposalCaller, error) {
	contract, err := bindProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalCaller{contract: contract}, nil
}

// NewProposalTransactor creates a new write-only instance of Proposal, bound to a specific deployed contract.
func NewProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalTransactor, error) {
	contract, err := bindProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalTransactor{contract: contract}, nil
}

// NewProposalFilterer creates a new log filterer instance of Proposal, bound to a specific deployed contract.
func NewProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalFilterer, error) {
	contract, err := bindProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalFilterer{contract: contract}, nil
}

// bindProposal binds a generic wrapper to an already deployed contract.
func bindProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.ProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transact(opts, method, params...)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Proposal *ProposalCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Proposal *ProposalSession) DaoAddress() (common.Address, error) {
	return _Proposal.Contract.DaoAddress(&_Proposal.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Proposal *ProposalCallerSession) DaoAddress() (common.Address, error) {
	return _Proposal.Contract.DaoAddress(&_Proposal.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Proposal *ProposalCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Proposal *ProposalSession) EndTime() (*big.Int, error) {
	return _Proposal.Contract.EndTime(&_Proposal.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Proposal *ProposalCallerSession) EndTime() (*big.Int, error) {
	return _Proposal.Contract.EndTime(&_Proposal.CallOpts)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(uint256 index) view returns((address,uint256,string,bytes,uint256))
func (_Proposal *ProposalCaller) GetAction(opts *bind.CallOpts, index *big.Int) (StructsAction, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "getAction", index)

	if err != nil {
		return *new(StructsAction), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsAction)).(*StructsAction)

	return out0, err

}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(uint256 index) view returns((address,uint256,string,bytes,uint256))
func (_Proposal *ProposalSession) GetAction(index *big.Int) (StructsAction, error) {
	return _Proposal.Contract.GetAction(&_Proposal.CallOpts, index)
}

// GetAction is a free data retrieval call binding the contract method 0xb6e76873.
//
// Solidity: function getAction(uint256 index) view returns((address,uint256,string,bytes,uint256))
func (_Proposal *ProposalCallerSession) GetAction(index *big.Int) (StructsAction, error) {
	return _Proposal.Contract.GetAction(&_Proposal.CallOpts, index)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_Proposal *ProposalCaller) GetActionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "getActionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_Proposal *ProposalSession) GetActionCount() (*big.Int, error) {
	return _Proposal.Contract.GetActionCount(&_Proposal.CallOpts)
}

// GetActionCount is a free data retrieval call binding the contract method 0x5eecd218.
//
// Solidity: function getActionCount() view returns(uint256)
func (_Proposal *ProposalCallerSession) GetActionCount() (*big.Int, error) {
	return _Proposal.Contract.GetActionCount(&_Proposal.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 tokenId) view returns(uint8)
func (_Proposal *ProposalCaller) GetVote(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "getVote", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 tokenId) view returns(uint8)
func (_Proposal *ProposalSession) GetVote(tokenId *big.Int) (uint8, error) {
	return _Proposal.Contract.GetVote(&_Proposal.CallOpts, tokenId)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 tokenId) view returns(uint8)
func (_Proposal *ProposalCallerSession) GetVote(tokenId *big.Int) (uint8, error) {
	return _Proposal.Contract.GetVote(&_Proposal.CallOpts, tokenId)
}

// HasVoted is a free data retrieval call binding the contract method 0xecca031f.
//
// Solidity: function hasVoted(uint256 tokenId) view returns(bool)
func (_Proposal *ProposalCaller) HasVoted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "hasVoted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0xecca031f.
//
// Solidity: function hasVoted(uint256 tokenId) view returns(bool)
func (_Proposal *ProposalSession) HasVoted(tokenId *big.Int) (bool, error) {
	return _Proposal.Contract.HasVoted(&_Proposal.CallOpts, tokenId)
}

// HasVoted is a free data retrieval call binding the contract method 0xecca031f.
//
// Solidity: function hasVoted(uint256 tokenId) view returns(bool)
func (_Proposal *ProposalCallerSession) HasVoted(tokenId *big.Int) (bool, error) {
	return _Proposal.Contract.HasVoted(&_Proposal.CallOpts, tokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proposal *ProposalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proposal *ProposalSession) Owner() (common.Address, error) {
	return _Proposal.Contract.Owner(&_Proposal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proposal *ProposalCallerSession) Owner() (common.Address, error) {
	return _Proposal.Contract.Owner(&_Proposal.CallOpts)
}

// ProposalTitle is a free data retrieval call binding the contract method 0xe5045752.
//
// Solidity: function proposalTitle() view returns(string)
func (_Proposal *ProposalCaller) ProposalTitle(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "proposalTitle")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProposalTitle is a free data retrieval call binding the contract method 0xe5045752.
//
// Solidity: function proposalTitle() view returns(string)
func (_Proposal *ProposalSession) ProposalTitle() (string, error) {
	return _Proposal.Contract.ProposalTitle(&_Proposal.CallOpts)
}

// ProposalTitle is a free data retrieval call binding the contract method 0xe5045752.
//
// Solidity: function proposalTitle() view returns(string)
func (_Proposal *ProposalCallerSession) ProposalTitle() (string, error) {
	return _Proposal.Contract.ProposalTitle(&_Proposal.CallOpts)
}

// Revotable is a free data retrieval call binding the contract method 0x0c8ba758.
//
// Solidity: function revotable() view returns(bool)
func (_Proposal *ProposalCaller) Revotable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "revotable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Revotable is a free data retrieval call binding the contract method 0x0c8ba758.
//
// Solidity: function revotable() view returns(bool)
func (_Proposal *ProposalSession) Revotable() (bool, error) {
	return _Proposal.Contract.Revotable(&_Proposal.CallOpts)
}

// Revotable is a free data retrieval call binding the contract method 0x0c8ba758.
//
// Solidity: function revotable() view returns(bool)
func (_Proposal *ProposalCallerSession) Revotable() (bool, error) {
	return _Proposal.Contract.Revotable(&_Proposal.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Proposal *ProposalCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Proposal *ProposalSession) StartTime() (*big.Int, error) {
	return _Proposal.Contract.StartTime(&_Proposal.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Proposal *ProposalCallerSession) StartTime() (*big.Int, error) {
	return _Proposal.Contract.StartTime(&_Proposal.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Proposal *ProposalCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Proposal *ProposalSession) Status() (uint8, error) {
	return _Proposal.Contract.Status(&_Proposal.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Proposal *ProposalCallerSession) Status() (uint8, error) {
	return _Proposal.Contract.Status(&_Proposal.CallOpts)
}

// VoteCount is a free data retrieval call binding the contract method 0xc6384071.
//
// Solidity: function voteCount() view returns(uint256)
func (_Proposal *ProposalCaller) VoteCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "voteCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteCount is a free data retrieval call binding the contract method 0xc6384071.
//
// Solidity: function voteCount() view returns(uint256)
func (_Proposal *ProposalSession) VoteCount() (*big.Int, error) {
	return _Proposal.Contract.VoteCount(&_Proposal.CallOpts)
}

// VoteCount is a free data retrieval call binding the contract method 0xc6384071.
//
// Solidity: function voteCount() view returns(uint256)
func (_Proposal *ProposalCallerSession) VoteCount() (*big.Int, error) {
	return _Proposal.Contract.VoteCount(&_Proposal.CallOpts)
}

// VoteTotals is a free data retrieval call binding the contract method 0x814a1d8e.
//
// Solidity: function voteTotals(uint256 ) view returns(uint256)
func (_Proposal *ProposalCaller) VoteTotals(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "voteTotals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteTotals is a free data retrieval call binding the contract method 0x814a1d8e.
//
// Solidity: function voteTotals(uint256 ) view returns(uint256)
func (_Proposal *ProposalSession) VoteTotals(arg0 *big.Int) (*big.Int, error) {
	return _Proposal.Contract.VoteTotals(&_Proposal.CallOpts, arg0)
}

// VoteTotals is a free data retrieval call binding the contract method 0x814a1d8e.
//
// Solidity: function voteTotals(uint256 ) view returns(uint256)
func (_Proposal *ProposalCallerSession) VoteTotals(arg0 *big.Int) (*big.Int, error) {
	return _Proposal.Contract.VoteTotals(&_Proposal.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) view returns(uint8 direction, bool voted)
func (_Proposal *ProposalCaller) Votes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Direction uint8
	Voted     bool
}, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "votes", arg0)

	outstruct := new(struct {
		Direction uint8
		Voted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Direction = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) view returns(uint8 direction, bool voted)
func (_Proposal *ProposalSession) Votes(arg0 *big.Int) (struct {
	Direction uint8
	Voted     bool
}, error) {
	return _Proposal.Contract.Votes(&_Proposal.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) view returns(uint8 direction, bool voted)
func (_Proposal *ProposalCallerSession) Votes(arg0 *big.Int) (struct {
	Direction uint8
	Voted     bool
}, error) {
	return _Proposal.Contract.Votes(&_Proposal.CallOpts, arg0)
}

// AddAction is a paid mutator transaction binding the contract method 0x54eb23a9.
//
// Solidity: function addAction(address target, uint256 value, string sig, bytes data) returns()
func (_Proposal *ProposalTransactor) AddAction(opts *bind.TransactOpts, target common.Address, value *big.Int, sig string, data []byte) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "addAction", target, value, sig, data)
}

// AddAction is a paid mutator transaction binding the contract method 0x54eb23a9.
//
// Solidity: function addAction(address target, uint256 value, string sig, bytes data) returns()
func (_Proposal *ProposalSession) AddAction(target common.Address, value *big.Int, sig string, data []byte) (*types.Transaction, error) {
	return _Proposal.Contract.AddAction(&_Proposal.TransactOpts, target, value, sig, data)
}

// AddAction is a paid mutator transaction binding the contract method 0x54eb23a9.
//
// Solidity: function addAction(address target, uint256 value, string sig, bytes data) returns()
func (_Proposal *ProposalTransactorSession) AddAction(target common.Address, value *big.Int, sig string, data []byte) (*types.Transaction, error) {
	return _Proposal.Contract.AddAction(&_Proposal.TransactOpts, target, value, sig, data)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x9070577f.
//
// Solidity: function cancelProposal() returns()
func (_Proposal *ProposalTransactor) CancelProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "cancelProposal")
}

// CancelProposal is a paid mutator transaction binding the contract method 0x9070577f.
//
// Solidity: function cancelProposal() returns()
func (_Proposal *ProposalSession) CancelProposal() (*types.Transaction, error) {
	return _Proposal.Contract.CancelProposal(&_Proposal.TransactOpts)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x9070577f.
//
// Solidity: function cancelProposal() returns()
func (_Proposal *ProposalTransactorSession) CancelProposal() (*types.Transaction, error) {
	return _Proposal.Contract.CancelProposal(&_Proposal.TransactOpts)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 tokenId, uint8 vote) returns()
func (_Proposal *ProposalTransactor) CastVote(opts *bind.TransactOpts, tokenId *big.Int, vote uint8) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "castVote", tokenId, vote)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 tokenId, uint8 vote) returns()
func (_Proposal *ProposalSession) CastVote(tokenId *big.Int, vote uint8) (*types.Transaction, error) {
	return _Proposal.Contract.CastVote(&_Proposal.TransactOpts, tokenId, vote)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 tokenId, uint8 vote) returns()
func (_Proposal *ProposalTransactorSession) CastVote(tokenId *big.Int, vote uint8) (*types.Transaction, error) {
	return _Proposal.Contract.CastVote(&_Proposal.TransactOpts, tokenId, vote)
}

// ConfigureProposal is a paid mutator transaction binding the contract method 0x5b4d2f85.
//
// Solidity: function configureProposal(bool isRevotable) returns()
func (_Proposal *ProposalTransactor) ConfigureProposal(opts *bind.TransactOpts, isRevotable bool) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "configureProposal", isRevotable)
}

// ConfigureProposal is a paid mutator transaction binding the contract method 0x5b4d2f85.
//
// Solidity: function configureProposal(bool isRevotable) returns()
func (_Proposal *ProposalSession) ConfigureProposal(isRevotable bool) (*types.Transaction, error) {
	return _Proposal.Contract.ConfigureProposal(&_Proposal.TransactOpts, isRevotable)
}

// ConfigureProposal is a paid mutator transaction binding the contract method 0x5b4d2f85.
//
// Solidity: function configureProposal(bool isRevotable) returns()
func (_Proposal *ProposalTransactorSession) ConfigureProposal(isRevotable bool) (*types.Transaction, error) {
	return _Proposal.Contract.ConfigureProposal(&_Proposal.TransactOpts, isRevotable)
}

// RemoveAction is a paid mutator transaction binding the contract method 0x421e4444.
//
// Solidity: function removeAction() returns()
func (_Proposal *ProposalTransactor) RemoveAction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "removeAction")
}

// RemoveAction is a paid mutator transaction binding the contract method 0x421e4444.
//
// Solidity: function removeAction() returns()
func (_Proposal *ProposalSession) RemoveAction() (*types.Transaction, error) {
	return _Proposal.Contract.RemoveAction(&_Proposal.TransactOpts)
}

// RemoveAction is a paid mutator transaction binding the contract method 0x421e4444.
//
// Solidity: function removeAction() returns()
func (_Proposal *ProposalTransactorSession) RemoveAction() (*types.Transaction, error) {
	return _Proposal.Contract.RemoveAction(&_Proposal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proposal *ProposalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proposal *ProposalSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proposal.Contract.RenounceOwnership(&_Proposal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proposal *ProposalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proposal.Contract.RenounceOwnership(&_Proposal.TransactOpts)
}

// SetActionReady is a paid mutator transaction binding the contract method 0x280b0c87.
//
// Solidity: function setActionReady(uint256 index, uint256 readyTime) returns()
func (_Proposal *ProposalTransactor) SetActionReady(opts *bind.TransactOpts, index *big.Int, readyTime *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setActionReady", index, readyTime)
}

// SetActionReady is a paid mutator transaction binding the contract method 0x280b0c87.
//
// Solidity: function setActionReady(uint256 index, uint256 readyTime) returns()
func (_Proposal *ProposalSession) SetActionReady(index *big.Int, readyTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetActionReady(&_Proposal.TransactOpts, index, readyTime)
}

// SetActionReady is a paid mutator transaction binding the contract method 0x280b0c87.
//
// Solidity: function setActionReady(uint256 index, uint256 readyTime) returns()
func (_Proposal *ProposalTransactorSession) SetActionReady(index *big.Int, readyTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetActionReady(&_Proposal.TransactOpts, index, readyTime)
}

// SetProposalTitle is a paid mutator transaction binding the contract method 0xbaceec7a.
//
// Solidity: function setProposalTitle(string newProposalTitle) returns()
func (_Proposal *ProposalTransactor) SetProposalTitle(opts *bind.TransactOpts, newProposalTitle string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setProposalTitle", newProposalTitle)
}

// SetProposalTitle is a paid mutator transaction binding the contract method 0xbaceec7a.
//
// Solidity: function setProposalTitle(string newProposalTitle) returns()
func (_Proposal *ProposalSession) SetProposalTitle(newProposalTitle string) (*types.Transaction, error) {
	return _Proposal.Contract.SetProposalTitle(&_Proposal.TransactOpts, newProposalTitle)
}

// SetProposalTitle is a paid mutator transaction binding the contract method 0xbaceec7a.
//
// Solidity: function setProposalTitle(string newProposalTitle) returns()
func (_Proposal *ProposalTransactorSession) SetProposalTitle(newProposalTitle string) (*types.Transaction, error) {
	return _Proposal.Contract.SetProposalTitle(&_Proposal.TransactOpts, newProposalTitle)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xa15a36c7.
//
// Solidity: function setVotingPeriod(uint256 newStartTime, uint256 newEndTime) returns()
func (_Proposal *ProposalTransactor) SetVotingPeriod(opts *bind.TransactOpts, newStartTime *big.Int, newEndTime *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setVotingPeriod", newStartTime, newEndTime)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xa15a36c7.
//
// Solidity: function setVotingPeriod(uint256 newStartTime, uint256 newEndTime) returns()
func (_Proposal *ProposalSession) SetVotingPeriod(newStartTime *big.Int, newEndTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetVotingPeriod(&_Proposal.TransactOpts, newStartTime, newEndTime)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xa15a36c7.
//
// Solidity: function setVotingPeriod(uint256 newStartTime, uint256 newEndTime) returns()
func (_Proposal *ProposalTransactorSession) SetVotingPeriod(newStartTime *big.Int, newEndTime *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.SetVotingPeriod(&_Proposal.TransactOpts, newStartTime, newEndTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proposal *ProposalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proposal *ProposalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.TransferOwnership(&_Proposal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proposal *ProposalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.TransferOwnership(&_Proposal.TransactOpts, newOwner)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x0b3af7f9.
//
// Solidity: function updateStatus(uint8 newStatus) returns()
func (_Proposal *ProposalTransactor) UpdateStatus(opts *bind.TransactOpts, newStatus uint8) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "updateStatus", newStatus)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x0b3af7f9.
//
// Solidity: function updateStatus(uint8 newStatus) returns()
func (_Proposal *ProposalSession) UpdateStatus(newStatus uint8) (*types.Transaction, error) {
	return _Proposal.Contract.UpdateStatus(&_Proposal.TransactOpts, newStatus)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x0b3af7f9.
//
// Solidity: function updateStatus(uint8 newStatus) returns()
func (_Proposal *ProposalTransactorSession) UpdateStatus(newStatus uint8) (*types.Transaction, error) {
	return _Proposal.Contract.UpdateStatus(&_Proposal.TransactOpts, newStatus)
}

// ProposalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Proposal contract.
type ProposalOwnershipTransferredIterator struct {
	Event *ProposalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProposalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalOwnershipTransferred)
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
		it.Event = new(ProposalOwnershipTransferred)
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
func (it *ProposalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalOwnershipTransferred represents a OwnershipTransferred event raised by the Proposal contract.
type ProposalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proposal *ProposalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProposalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProposalOwnershipTransferredIterator{contract: _Proposal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proposal *ProposalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProposalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalOwnershipTransferred)
				if err := _Proposal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Proposal *ProposalFilterer) ParseOwnershipTransferred(log types.Log) (*ProposalOwnershipTransferred, error) {
	event := new(ProposalOwnershipTransferred)
	if err := _Proposal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalProposalResultsIterator is returned from FilterProposalResults and is used to iterate over the raw logs and unpacked data for ProposalResults events raised by the Proposal contract.
type ProposalProposalResultsIterator struct {
	Event *ProposalProposalResults // Event containing the contract specifics and raw log

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
func (it *ProposalProposalResultsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalProposalResults)
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
		it.Event = new(ProposalProposalResults)
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
func (it *ProposalProposalResultsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalProposalResultsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalProposalResults represents a ProposalResults event raised by the Proposal contract.
type ProposalProposalResults struct {
	QuorumPassed          bool
	SupportPassed         bool
	TotalVotes            *big.Int
	QuorumVotesThreshold  *big.Int
	SupportVotesThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProposalResults is a free log retrieval operation binding the contract event 0xc757c414d8a1f46722ffdfdf111db96f47d6b2fb327d0cb8869bf7e69125716f.
//
// Solidity: event ProposalResults(bool quorumPassed, bool supportPassed, uint256 totalVotes, uint256 quorumVotesThreshold, uint256 supportVotesThreshold)
func (_Proposal *ProposalFilterer) FilterProposalResults(opts *bind.FilterOpts) (*ProposalProposalResultsIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "ProposalResults")
	if err != nil {
		return nil, err
	}
	return &ProposalProposalResultsIterator{contract: _Proposal.contract, event: "ProposalResults", logs: logs, sub: sub}, nil
}

// WatchProposalResults is a free log subscription operation binding the contract event 0xc757c414d8a1f46722ffdfdf111db96f47d6b2fb327d0cb8869bf7e69125716f.
//
// Solidity: event ProposalResults(bool quorumPassed, bool supportPassed, uint256 totalVotes, uint256 quorumVotesThreshold, uint256 supportVotesThreshold)
func (_Proposal *ProposalFilterer) WatchProposalResults(opts *bind.WatchOpts, sink chan<- *ProposalProposalResults) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "ProposalResults")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalProposalResults)
				if err := _Proposal.contract.UnpackLog(event, "ProposalResults", log); err != nil {
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

// ParseProposalResults is a log parse operation binding the contract event 0xc757c414d8a1f46722ffdfdf111db96f47d6b2fb327d0cb8869bf7e69125716f.
//
// Solidity: event ProposalResults(bool quorumPassed, bool supportPassed, uint256 totalVotes, uint256 quorumVotesThreshold, uint256 supportVotesThreshold)
func (_Proposal *ProposalFilterer) ParseProposalResults(log types.Log) (*ProposalProposalResults, error) {
	event := new(ProposalProposalResults)
	if err := _Proposal.contract.UnpackLog(event, "ProposalResults", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the Proposal contract.
type ProposalStatusUpdatedIterator struct {
	Event *ProposalStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ProposalStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalStatusUpdated)
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
		it.Event = new(ProposalStatusUpdated)
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
func (it *ProposalStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalStatusUpdated represents a StatusUpdated event raised by the Proposal contract.
type ProposalStatusUpdated struct {
	NewStatus string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x69f380a4c820112d1c3e49f6e19e1d33c66d98a47f9c8ad3221a82b79b0bfcbc.
//
// Solidity: event StatusUpdated(string newStatus)
func (_Proposal *ProposalFilterer) FilterStatusUpdated(opts *bind.FilterOpts) (*ProposalStatusUpdatedIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "StatusUpdated")
	if err != nil {
		return nil, err
	}
	return &ProposalStatusUpdatedIterator{contract: _Proposal.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x69f380a4c820112d1c3e49f6e19e1d33c66d98a47f9c8ad3221a82b79b0bfcbc.
//
// Solidity: event StatusUpdated(string newStatus)
func (_Proposal *ProposalFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *ProposalStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "StatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalStatusUpdated)
				if err := _Proposal.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
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

// ParseStatusUpdated is a log parse operation binding the contract event 0x69f380a4c820112d1c3e49f6e19e1d33c66d98a47f9c8ad3221a82b79b0bfcbc.
//
// Solidity: event StatusUpdated(string newStatus)
func (_Proposal *ProposalFilterer) ParseStatusUpdated(log types.Log) (*ProposalStatusUpdated, error) {
	event := new(ProposalStatusUpdated)
	if err := _Proposal.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Proposal contract.
type ProposalVoteCastIterator struct {
	Event *ProposalVoteCast // Event containing the contract specifics and raw log

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
func (it *ProposalVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalVoteCast)
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
		it.Event = new(ProposalVoteCast)
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
func (it *ProposalVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalVoteCast represents a VoteCast event raised by the Proposal contract.
type ProposalVoteCast struct {
	Voter   common.Address
	TokenId *big.Int
	Vote    uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x251f168c3d37dca4f3922f4aa12852da0c3de8ef70710e0cf58eba85805168a2.
//
// Solidity: event VoteCast(address voter, uint256 tokenId, uint8 vote)
func (_Proposal *ProposalFilterer) FilterVoteCast(opts *bind.FilterOpts) (*ProposalVoteCastIterator, error) {

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return &ProposalVoteCastIterator{contract: _Proposal.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x251f168c3d37dca4f3922f4aa12852da0c3de8ef70710e0cf58eba85805168a2.
//
// Solidity: event VoteCast(address voter, uint256 tokenId, uint8 vote)
func (_Proposal *ProposalFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *ProposalVoteCast) (event.Subscription, error) {

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalVoteCast)
				if err := _Proposal.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0x251f168c3d37dca4f3922f4aa12852da0c3de8ef70710e0cf58eba85805168a2.
//
// Solidity: event VoteCast(address voter, uint256 tokenId, uint8 vote)
func (_Proposal *ProposalFilterer) ParseVoteCast(log types.Log) (*ProposalVoteCast, error) {
	event := new(ProposalVoteCast)
	if err := _Proposal.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
