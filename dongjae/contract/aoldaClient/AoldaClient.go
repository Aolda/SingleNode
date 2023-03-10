// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aoldaClient

import (
	"errors"
	"fmt"
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

// AoldaClientMetaData contains all meta data concerning the AoldaClient contract.
var AoldaClientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fn\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"arguments\",\"type\":\"string[]\"}],\"name\":\"CallAolda\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fn\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"arguments\",\"type\":\"string[]\"}],\"name\":\"callAolda\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fn\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"arguments\",\"type\":\"string[]\"}],\"name\":\"makeSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c1b806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80636984394014610051578063bb9a274d14610081578063c7bfebd61461009d578063f9b1d733146100b9575b600080fd5b61006b600480360381019061006691906102bc565b6100e9565b6040516100789190610379565b60405180910390f35b61009b600480360381019061009691906105b6565b61018d565b005b6100b760048036038101906100b2919061062e565b6101ca565b005b6100d360048036038101906100ce91906105b6565b6101ee565b6040516100e09190610699565b60405180910390f35b60606000808381526020019081526020016000208054610108906106e3565b80601f0160208091040260200160405190810160405280929190818152602001828054610134906106e3565b80156101815780601f1061015657610100808354040283529160200191610181565b820191906000526020600020905b81548152906001019060200180831161016457829003601f168201915b50505050509050919050565b7f8d397347d14d1ad0861879661d8b750cec64835b214eee49d80b6bf6bb5950de82826040516101be929190610820565b60405180910390a15050565b8060008084815260200190815260200160002090816101e99190610a0d565b505050565b60008083905060005b8351811015610251578184828151811061021457610213610adf565b5b602002602001015160405160200161022d929190610b4a565b6040516020818303038152906040529150808061024990610b9d565b9150506101f7565b5061025b81610264565b91505092915050565b600060208201519050919050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61029981610286565b81146102a457600080fd5b50565b6000813590506102b681610290565b92915050565b6000602082840312156102d2576102d161027c565b5b60006102e0848285016102a7565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610323578082015181840152602081019050610308565b60008484015250505050565b6000601f19601f8301169050919050565b600061034b826102e9565b61035581856102f4565b9350610365818560208601610305565b61036e8161032f565b840191505092915050565b600060208201905081810360008301526103938184610340565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103dd8261032f565b810181811067ffffffffffffffff821117156103fc576103fb6103a5565b5b80604052505050565b600061040f610272565b905061041b82826103d4565b919050565b600067ffffffffffffffff82111561043b5761043a6103a5565b5b6104448261032f565b9050602081019050919050565b82818337600083830152505050565b600061047361046e84610420565b610405565b90508281526020810184848401111561048f5761048e6103a0565b5b61049a848285610451565b509392505050565b600082601f8301126104b7576104b661039b565b5b81356104c7848260208601610460565b91505092915050565b600067ffffffffffffffff8211156104eb576104ea6103a5565b5b602082029050602081019050919050565b600080fd5b600061051461050f846104d0565b610405565b90508083825260208201905060208402830185811115610537576105366104fc565b5b835b8181101561057e57803567ffffffffffffffff81111561055c5761055b61039b565b5b80860161056989826104a2565b85526020850194505050602081019050610539565b5050509392505050565b600082601f83011261059d5761059c61039b565b5b81356105ad848260208601610501565b91505092915050565b600080604083850312156105cd576105cc61027c565b5b600083013567ffffffffffffffff8111156105eb576105ea610281565b5b6105f7858286016104a2565b925050602083013567ffffffffffffffff81111561061857610617610281565b5b61062485828601610588565b9150509250929050565b600080604083850312156106455761064461027c565b5b6000610653858286016102a7565b925050602083013567ffffffffffffffff81111561067457610673610281565b5b610680858286016104a2565b9150509250929050565b61069381610286565b82525050565b60006020820190506106ae600083018461068a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806106fb57607f821691505b60208210810361070e5761070d6106b4565b5b50919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061075c826102e9565b6107668185610740565b9350610776818560208601610305565b61077f8161032f565b840191505092915050565b60006107968383610751565b905092915050565b6000602082019050919050565b60006107b682610714565b6107c0818561071f565b9350836020820285016107d285610730565b8060005b8581101561080e57848403895281516107ef858261078a565b94506107fa8361079e565b925060208a019950506001810190506107d6565b50829750879550505050505092915050565b6000604082019050818103600083015261083a8185610340565b9050818103602083015261084e81846107ab565b90509392505050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108b97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261087c565b6108c3868361087c565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061090a610905610900846108db565b6108e5565b6108db565b9050919050565b6000819050919050565b610924836108ef565b61093861093082610911565b848454610889565b825550505050565b600090565b61094d610940565b61095881848461091b565b505050565b5b8181101561097c57610971600082610945565b60018101905061095e565b5050565b601f8211156109c15761099281610857565b61099b8461086c565b810160208510156109aa578190505b6109be6109b68561086c565b83018261095d565b50505b505050565b600082821c905092915050565b60006109e4600019846008026109c6565b1980831691505092915050565b60006109fd83836109d3565b9150826002028217905092915050565b610a16826102e9565b67ffffffffffffffff811115610a2f57610a2e6103a5565b5b610a3982546106e3565b610a44828285610980565b600060209050601f831160018114610a775760008415610a65578287015190505b610a6f85826109f1565b865550610ad7565b601f198416610a8586610857565b60005b82811015610aad57848901518255600182019150602085019450602081019050610a88565b86831015610aca5784890151610ac6601f8916826109d3565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b6000610b24826102e9565b610b2e8185610b0e565b9350610b3e818560208601610305565b80840191505092915050565b6000610b568285610b19565b9150610b628284610b19565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ba8826108db565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bda57610bd9610b6e565b5b60018201905091905056fea2646970667358221220ade88dfa20cfa3b7a1b1a55ad4f7bbee423bc112379175d1e7d5cb3e5f9b415a64736f6c63430008110033",
}

// AoldaClientABI is the input ABI used to generate the binding from.
// Deprecated: Use AoldaClientMetaData.ABI instead.
var AoldaClientABI = AoldaClientMetaData.ABI

// AoldaClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AoldaClientMetaData.Bin instead.
var AoldaClientBin = AoldaClientMetaData.Bin

// DeployAoldaClient deploys a new Ethereum contract, binding an instance of AoldaClient to it.
func DeployAoldaClient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AoldaClient, error) {
	parsed, err := AoldaClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AoldaClientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AoldaClient{AoldaClientCaller: AoldaClientCaller{contract: contract}, AoldaClientTransactor: AoldaClientTransactor{contract: contract}, AoldaClientFilterer: AoldaClientFilterer{contract: contract}}, nil
}

// AoldaClient is an auto generated Go binding around an Ethereum contract.
type AoldaClient struct {
	AoldaClientCaller     // Read-only binding to the contract
	AoldaClientTransactor // Write-only binding to the contract
	AoldaClientFilterer   // Log filterer for contract events
}

// AoldaClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type AoldaClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoldaClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AoldaClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoldaClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AoldaClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoldaClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AoldaClientSession struct {
	Contract     *AoldaClient      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AoldaClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AoldaClientCallerSession struct {
	Contract *AoldaClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AoldaClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AoldaClientTransactorSession struct {
	Contract     *AoldaClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AoldaClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type AoldaClientRaw struct {
	Contract *AoldaClient // Generic contract binding to access the raw methods on
}

// AoldaClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AoldaClientCallerRaw struct {
	Contract *AoldaClientCaller // Generic read-only contract binding to access the raw methods on
}

// AoldaClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AoldaClientTransactorRaw struct {
	Contract *AoldaClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAoldaClient creates a new instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClient(address common.Address, backend bind.ContractBackend) (*AoldaClient, error) {
	contract, err := bindAoldaClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AoldaClient{AoldaClientCaller: AoldaClientCaller{contract: contract}, AoldaClientTransactor: AoldaClientTransactor{contract: contract}, AoldaClientFilterer: AoldaClientFilterer{contract: contract}}, nil
}

// NewAoldaClientCaller creates a new read-only instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClientCaller(address common.Address, caller bind.ContractCaller) (*AoldaClientCaller, error) {
	contract, err := bindAoldaClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AoldaClientCaller{contract: contract}, nil
}

// NewAoldaClientTransactor creates a new write-only instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClientTransactor(address common.Address, transactor bind.ContractTransactor) (*AoldaClientTransactor, error) {
	contract, err := bindAoldaClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AoldaClientTransactor{contract: contract}, nil
}

// NewAoldaClientFilterer creates a new log filterer instance of AoldaClient, bound to a specific deployed contract.
func NewAoldaClientFilterer(address common.Address, filterer bind.ContractFilterer) (*AoldaClientFilterer, error) {
	contract, err := bindAoldaClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AoldaClientFilterer{contract: contract}, nil
}

// bindAoldaClient binds a generic wrapper to an already deployed contract.
func bindAoldaClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AoldaClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AoldaClient *AoldaClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AoldaClient.Contract.AoldaClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AoldaClient *AoldaClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AoldaClient.Contract.AoldaClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AoldaClient *AoldaClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AoldaClient.Contract.AoldaClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AoldaClient *AoldaClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AoldaClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AoldaClient *AoldaClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AoldaClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AoldaClient *AoldaClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AoldaClient.Contract.contract.Transact(opts, method, params...)
}

// GetValue is a free data retrieval call binding the contract method 0x69843940.
//
// Solidity: function getValue(bytes32 signature) view returns(string)
func (_AoldaClient *AoldaClientCaller) GetValue(opts *bind.CallOpts, signature [32]byte) (string, error) {
	var out []interface{}
	err := _AoldaClient.contract.Call(opts, &out, "getValue", signature)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x69843940.
//
// Solidity: function getValue(bytes32 signature) view returns(string)
func (_AoldaClient *AoldaClientSession) GetValue(signature [32]byte) (string, error) {
	return _AoldaClient.Contract.GetValue(&_AoldaClient.CallOpts, signature)
}

// GetValue is a free data retrieval call binding the contract method 0x69843940.
//
// Solidity: function getValue(bytes32 signature) view returns(string)
func (_AoldaClient *AoldaClientCallerSession) GetValue(signature [32]byte) (string, error) {
	return _AoldaClient.Contract.GetValue(&_AoldaClient.CallOpts, signature)
}

// MakeSignature is a free data retrieval call binding the contract method 0xf9b1d733.
//
// Solidity: function makeSignature(string fn, string[] arguments) pure returns(bytes32)
func (_AoldaClient *AoldaClientCaller) MakeSignature(opts *bind.CallOpts, fn string, arguments []string) ([32]byte, error) {
	var out []interface{}
	err := _AoldaClient.contract.Call(opts, &out, "makeSignature", fn, arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_AoldaClient *AoldaClientCaller) Test(opts *bind.CallOpts, fn string, arguments []string){
	fmt.Println(fn);
}

// MakeSignature is a free data retrieval call binding the contract method 0xf9b1d733.
//
// Solidity: function makeSignature(string fn, string[] arguments) pure returns(bytes32)
func (_AoldaClient *AoldaClientSession) MakeSignature(fn string, arguments []string) ([32]byte, error) {
	return _AoldaClient.Contract.MakeSignature(&_AoldaClient.CallOpts, fn, arguments)
}

// MakeSignature is a free data retrieval call binding the contract method 0xf9b1d733.
//
// Solidity: function makeSignature(string fn, string[] arguments) pure returns(bytes32)
func (_AoldaClient *AoldaClientCallerSession) MakeSignature(fn string, arguments []string) ([32]byte, error) {
	return _AoldaClient.Contract.MakeSignature(&_AoldaClient.CallOpts, fn, arguments)
}

// CallAolda is a paid mutator transaction binding the contract method 0xbb9a274d.
//
// Solidity: function callAolda(string fn, string[] arguments) returns()
func (_AoldaClient *AoldaClientTransactor) CallAolda(opts *bind.TransactOpts, fn string, arguments []string) (*types.Transaction, error) {
	return _AoldaClient.contract.Transact(opts, "callAolda", fn, arguments)
}

// CallAolda is a paid mutator transaction binding the contract method 0xbb9a274d.
//
// Solidity: function callAolda(string fn, string[] arguments) returns()
func (_AoldaClient *AoldaClientSession) CallAolda(fn string, arguments []string) (*types.Transaction, error) {
	return _AoldaClient.Contract.CallAolda(&_AoldaClient.TransactOpts, fn, arguments)
}

// CallAolda is a paid mutator transaction binding the contract method 0xbb9a274d.
//
// Solidity: function callAolda(string fn, string[] arguments) returns()
func (_AoldaClient *AoldaClientTransactorSession) CallAolda(fn string, arguments []string) (*types.Transaction, error) {
	return _AoldaClient.Contract.CallAolda(&_AoldaClient.TransactOpts, fn, arguments)
}

// SetValue is a paid mutator transaction binding the contract method 0xc7bfebd6.
//
// Solidity: function setValue(bytes32 signature, string value) returns()
func (_AoldaClient *AoldaClientTransactor) SetValue(opts *bind.TransactOpts, signature [32]byte, value string) (*types.Transaction, error) {
	return _AoldaClient.contract.Transact(opts, "setValue", signature, value)
}

// SetValue is a paid mutator transaction binding the contract method 0xc7bfebd6.
//
// Solidity: function setValue(bytes32 signature, string value) returns()
func (_AoldaClient *AoldaClientSession) SetValue(signature [32]byte, value string) (*types.Transaction, error) {
	return _AoldaClient.Contract.SetValue(&_AoldaClient.TransactOpts, signature, value)
}

// SetValue is a paid mutator transaction binding the contract method 0xc7bfebd6.
//
// Solidity: function setValue(bytes32 signature, string value) returns()
func (_AoldaClient *AoldaClientTransactorSession) SetValue(signature [32]byte, value string) (*types.Transaction, error) {
	return _AoldaClient.Contract.SetValue(&_AoldaClient.TransactOpts, signature, value)
}

// AoldaClientCallAoldaIterator is returned from FilterCallAolda and is used to iterate over the raw logs and unpacked data for CallAolda events raised by the AoldaClient contract.
type AoldaClientCallAoldaIterator struct {
	Event *AoldaClientCallAolda // Event containing the contract specifics and raw log

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
func (it *AoldaClientCallAoldaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoldaClientCallAolda)
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
		it.Event = new(AoldaClientCallAolda)
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
func (it *AoldaClientCallAoldaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoldaClientCallAoldaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoldaClientCallAolda represents a CallAolda event raised by the AoldaClient contract.
type AoldaClientCallAolda struct {
	Fn        string
	Arguments []string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCallAolda is a free log retrieval operation binding the contract event 0x8d397347d14d1ad0861879661d8b750cec64835b214eee49d80b6bf6bb5950de.
//
// Solidity: event CallAolda(string fn, string[] arguments)
func (_AoldaClient *AoldaClientFilterer) FilterCallAolda(opts *bind.FilterOpts) (*AoldaClientCallAoldaIterator, error) {

	logs, sub, err := _AoldaClient.contract.FilterLogs(opts, "CallAolda")
	if err != nil {
		return nil, err
	}
	return &AoldaClientCallAoldaIterator{contract: _AoldaClient.contract, event: "CallAolda", logs: logs, sub: sub}, nil
}

// WatchCallAolda is a free log subscription operation binding the contract event 0x8d397347d14d1ad0861879661d8b750cec64835b214eee49d80b6bf6bb5950de.
//
// Solidity: event CallAolda(string fn, string[] arguments)
func (_AoldaClient *AoldaClientFilterer) WatchCallAolda(opts *bind.WatchOpts, sink chan<- *AoldaClientCallAolda) (event.Subscription, error) {

	logs, sub, err := _AoldaClient.contract.WatchLogs(opts, "CallAolda")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoldaClientCallAolda)
				if err := _AoldaClient.contract.UnpackLog(event, "CallAolda", log); err != nil {
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

// ParseCallAolda is a log parse operation binding the contract event 0x8d397347d14d1ad0861879661d8b750cec64835b214eee49d80b6bf6bb5950de.
//
// Solidity: event CallAolda(string fn, string[] arguments)
func (_AoldaClient *AoldaClientFilterer) ParseCallAolda(log types.Log) (*AoldaClientCallAolda, error) {
	event := new(AoldaClientCallAolda)
	if err := _AoldaClient.contract.UnpackLog(event, "CallAolda", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
