// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proposalfactory

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

// ProposalfactoryMetaData contains all meta data concerning the Proposalfactory contract.
var ProposalfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factoryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"ContractProduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalFactoryCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"produceOGREProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProposalfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalfactoryMetaData.ABI instead.
var ProposalfactoryABI = ProposalfactoryMetaData.ABI

// Proposalfactory is an auto generated Go binding around an Ethereum contract.
type Proposalfactory struct {
	ProposalfactoryCaller     // Read-only binding to the contract
	ProposalfactoryTransactor // Write-only binding to the contract
	ProposalfactoryFilterer   // Log filterer for contract events
}

// ProposalfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalfactorySession struct {
	Contract     *Proposalfactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalfactoryCallerSession struct {
	Contract *ProposalfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProposalfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalfactoryTransactorSession struct {
	Contract     *ProposalfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProposalfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalfactoryRaw struct {
	Contract *Proposalfactory // Generic contract binding to access the raw methods on
}

// ProposalfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalfactoryCallerRaw struct {
	Contract *ProposalfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalfactoryTransactorRaw struct {
	Contract *ProposalfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalfactory creates a new instance of Proposalfactory, bound to a specific deployed contract.
func NewProposalfactory(address common.Address, backend bind.ContractBackend) (*Proposalfactory, error) {
	contract, err := bindProposalfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proposalfactory{ProposalfactoryCaller: ProposalfactoryCaller{contract: contract}, ProposalfactoryTransactor: ProposalfactoryTransactor{contract: contract}, ProposalfactoryFilterer: ProposalfactoryFilterer{contract: contract}}, nil
}

// NewProposalfactoryCaller creates a new read-only instance of Proposalfactory, bound to a specific deployed contract.
func NewProposalfactoryCaller(address common.Address, caller bind.ContractCaller) (*ProposalfactoryCaller, error) {
	contract, err := bindProposalfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalfactoryCaller{contract: contract}, nil
}

// NewProposalfactoryTransactor creates a new write-only instance of Proposalfactory, bound to a specific deployed contract.
func NewProposalfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalfactoryTransactor, error) {
	contract, err := bindProposalfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalfactoryTransactor{contract: contract}, nil
}

// NewProposalfactoryFilterer creates a new log filterer instance of Proposalfactory, bound to a specific deployed contract.
func NewProposalfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalfactoryFilterer, error) {
	contract, err := bindProposalfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalfactoryFilterer{contract: contract}, nil
}

// bindProposalfactory binds a generic wrapper to an already deployed contract.
func bindProposalfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposalfactory *ProposalfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposalfactory.Contract.ProposalfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposalfactory *ProposalfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposalfactory.Contract.ProposalfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposalfactory *ProposalfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposalfactory.Contract.ProposalfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposalfactory *ProposalfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposalfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposalfactory *ProposalfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposalfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposalfactory *ProposalfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposalfactory.Contract.contract.Transact(opts, method, params...)
}

// ProductionCount is a free data retrieval call binding the contract method 0x17350165.
//
// Solidity: function productionCount() view returns(uint256)
func (_Proposalfactory *ProposalfactoryCaller) ProductionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposalfactory.contract.Call(opts, &out, "productionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProductionCount is a free data retrieval call binding the contract method 0x17350165.
//
// Solidity: function productionCount() view returns(uint256)
func (_Proposalfactory *ProposalfactorySession) ProductionCount() (*big.Int, error) {
	return _Proposalfactory.Contract.ProductionCount(&_Proposalfactory.CallOpts)
}

// ProductionCount is a free data retrieval call binding the contract method 0x17350165.
//
// Solidity: function productionCount() view returns(uint256)
func (_Proposalfactory *ProposalfactoryCallerSession) ProductionCount() (*big.Int, error) {
	return _Proposalfactory.Contract.ProductionCount(&_Proposalfactory.CallOpts)
}

// ProduceOGREProposal is a paid mutator transaction binding the contract method 0xcec70da6.
//
// Solidity: function produceOGREProposal(string title, address daoAddress, address owner) returns(address)
func (_Proposalfactory *ProposalfactoryTransactor) ProduceOGREProposal(opts *bind.TransactOpts, title string, daoAddress common.Address, owner common.Address) (*types.Transaction, error) {
	return _Proposalfactory.contract.Transact(opts, "produceOGREProposal", title, daoAddress, owner)
}

// ProduceOGREProposal is a paid mutator transaction binding the contract method 0xcec70da6.
//
// Solidity: function produceOGREProposal(string title, address daoAddress, address owner) returns(address)
func (_Proposalfactory *ProposalfactorySession) ProduceOGREProposal(title string, daoAddress common.Address, owner common.Address) (*types.Transaction, error) {
	return _Proposalfactory.Contract.ProduceOGREProposal(&_Proposalfactory.TransactOpts, title, daoAddress, owner)
}

// ProduceOGREProposal is a paid mutator transaction binding the contract method 0xcec70da6.
//
// Solidity: function produceOGREProposal(string title, address daoAddress, address owner) returns(address)
func (_Proposalfactory *ProposalfactoryTransactorSession) ProduceOGREProposal(title string, daoAddress common.Address, owner common.Address) (*types.Transaction, error) {
	return _Proposalfactory.Contract.ProduceOGREProposal(&_Proposalfactory.TransactOpts, title, daoAddress, owner)
}

// ProposalfactoryContractProducedIterator is returned from FilterContractProduced and is used to iterate over the raw logs and unpacked data for ContractProduced events raised by the Proposalfactory contract.
type ProposalfactoryContractProducedIterator struct {
	Event *ProposalfactoryContractProduced // Event containing the contract specifics and raw log

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
func (it *ProposalfactoryContractProducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalfactoryContractProduced)
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
		it.Event = new(ProposalfactoryContractProduced)
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
func (it *ProposalfactoryContractProducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalfactoryContractProducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalfactoryContractProduced represents a ContractProduced event raised by the Proposalfactory contract.
type ProposalfactoryContractProduced struct {
	ContractAddress common.Address
	FactoryAddress  common.Address
	Producer        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractProduced is a free log retrieval operation binding the contract event 0x5db1f34edd85623861e21337b01e1c2035c3383b68b101d961bc2337c580a962.
//
// Solidity: event ContractProduced(address contractAddress, address factoryAddress, address producer)
func (_Proposalfactory *ProposalfactoryFilterer) FilterContractProduced(opts *bind.FilterOpts) (*ProposalfactoryContractProducedIterator, error) {

	logs, sub, err := _Proposalfactory.contract.FilterLogs(opts, "ContractProduced")
	if err != nil {
		return nil, err
	}
	return &ProposalfactoryContractProducedIterator{contract: _Proposalfactory.contract, event: "ContractProduced", logs: logs, sub: sub}, nil
}

// WatchContractProduced is a free log subscription operation binding the contract event 0x5db1f34edd85623861e21337b01e1c2035c3383b68b101d961bc2337c580a962.
//
// Solidity: event ContractProduced(address contractAddress, address factoryAddress, address producer)
func (_Proposalfactory *ProposalfactoryFilterer) WatchContractProduced(opts *bind.WatchOpts, sink chan<- *ProposalfactoryContractProduced) (event.Subscription, error) {

	logs, sub, err := _Proposalfactory.contract.WatchLogs(opts, "ContractProduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalfactoryContractProduced)
				if err := _Proposalfactory.contract.UnpackLog(event, "ContractProduced", log); err != nil {
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

// ParseContractProduced is a log parse operation binding the contract event 0x5db1f34edd85623861e21337b01e1c2035c3383b68b101d961bc2337c580a962.
//
// Solidity: event ContractProduced(address contractAddress, address factoryAddress, address producer)
func (_Proposalfactory *ProposalfactoryFilterer) ParseContractProduced(log types.Log) (*ProposalfactoryContractProduced, error) {
	event := new(ProposalfactoryContractProduced)
	if err := _Proposalfactory.contract.UnpackLog(event, "ContractProduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalfactoryProposalFactoryCreatedIterator is returned from FilterProposalFactoryCreated and is used to iterate over the raw logs and unpacked data for ProposalFactoryCreated events raised by the Proposalfactory contract.
type ProposalfactoryProposalFactoryCreatedIterator struct {
	Event *ProposalfactoryProposalFactoryCreated // Event containing the contract specifics and raw log

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
func (it *ProposalfactoryProposalFactoryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalfactoryProposalFactoryCreated)
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
		it.Event = new(ProposalfactoryProposalFactoryCreated)
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
func (it *ProposalfactoryProposalFactoryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalfactoryProposalFactoryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalfactoryProposalFactoryCreated represents a ProposalFactoryCreated event raised by the Proposalfactory contract.
type ProposalfactoryProposalFactoryCreated struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProposalFactoryCreated is a free log retrieval operation binding the contract event 0x35d71b99bdb343ca2dd2814592bbf63b5a379cf9b151185dcd862a0199dd8753.
//
// Solidity: event ProposalFactoryCreated(address creator)
func (_Proposalfactory *ProposalfactoryFilterer) FilterProposalFactoryCreated(opts *bind.FilterOpts) (*ProposalfactoryProposalFactoryCreatedIterator, error) {

	logs, sub, err := _Proposalfactory.contract.FilterLogs(opts, "ProposalFactoryCreated")
	if err != nil {
		return nil, err
	}
	return &ProposalfactoryProposalFactoryCreatedIterator{contract: _Proposalfactory.contract, event: "ProposalFactoryCreated", logs: logs, sub: sub}, nil
}

// WatchProposalFactoryCreated is a free log subscription operation binding the contract event 0x35d71b99bdb343ca2dd2814592bbf63b5a379cf9b151185dcd862a0199dd8753.
//
// Solidity: event ProposalFactoryCreated(address creator)
func (_Proposalfactory *ProposalfactoryFilterer) WatchProposalFactoryCreated(opts *bind.WatchOpts, sink chan<- *ProposalfactoryProposalFactoryCreated) (event.Subscription, error) {

	logs, sub, err := _Proposalfactory.contract.WatchLogs(opts, "ProposalFactoryCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalfactoryProposalFactoryCreated)
				if err := _Proposalfactory.contract.UnpackLog(event, "ProposalFactoryCreated", log); err != nil {
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

// ParseProposalFactoryCreated is a log parse operation binding the contract event 0x35d71b99bdb343ca2dd2814592bbf63b5a379cf9b151185dcd862a0199dd8753.
//
// Solidity: event ProposalFactoryCreated(address creator)
func (_Proposalfactory *ProposalfactoryFilterer) ParseProposalFactoryCreated(log types.Log) (*ProposalfactoryProposalFactoryCreated, error) {
	event := new(ProposalfactoryProposalFactoryCreated)
	if err := _Proposalfactory.contract.UnpackLog(event, "ProposalFactoryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
