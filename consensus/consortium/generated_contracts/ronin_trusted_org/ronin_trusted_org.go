// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package roninTrustedOrg

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

// IRoninTrustedOrganizationTrustedOrganization is an auto generated low-level Go binding around an user-defined struct.
type IRoninTrustedOrganizationTrustedOrganization struct {
	ConsensusAddr         common.Address
	Governor              common.Address
	DeprecatedBridgeVoter common.Address
	Weight                *big.Int
	AddedBlock            *big.Int
}

// RoninTrustedOrgMetaData contains all meta data concerning the RoninTrustedOrg contract.
var RoninTrustedOrgMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrConsensusAddressIsAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrConsensusAddressIsNotAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"ErrContractTypeNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrDuplicated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrEmptyArray\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrGovernorAddressIsAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidVoteWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrQueryForDupplicated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrQueryForNonExistentConsensusAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumRoleAccess\",\"name\":\"expectedRole\",\"type\":\"uint8\"}],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumContractType\",\"name\":\"expectedContractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ErrUnexpectedInternalCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization\",\"name\":\"orgAfterChanged\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"TConsensus\",\"name\":\"oldConsensus\",\"type\":\"address\"}],\"name\":\"ConsensusAddressOfTrustedOrgChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"orgs\",\"type\":\"tuple[]\"}],\"name\":\"TrustedOrganizationsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"TConsensus[]\",\"name\":\"orgs\",\"type\":\"address[]\"}],\"name\":\"TrustedOrganizationsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"orgs\",\"type\":\"tuple[]\"}],\"name\":\"TrustedOrganizationsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"name\":\"addTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countTrustedOrganization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"oldAddr\",\"type\":\"address\"},{\"internalType\":\"TConsensus\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"execChangeConsensusAddressForTrustedOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTrustedOrganizations\",\"outputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"}],\"name\":\"getConsensusWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"}],\"name\":\"getConsensusWeightById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus[]\",\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"getConsensusWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cids\",\"type\":\"address[]\"}],\"name\":\"getConsensusWeightsById\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"getGovernorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"getGovernorWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_res\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"getTrustedOrganization\",\"outputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"getTrustedOrganizationAt\",\"outputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"trustedOrgs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileContract\",\"type\":\"address\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus[]\",\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"removeTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"sumConsensusWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"sumGovernorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__deprecatedBridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"name\":\"updateTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RoninTrustedOrgABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninTrustedOrgMetaData.ABI instead.
var RoninTrustedOrgABI = RoninTrustedOrgMetaData.ABI

// RoninTrustedOrg is an auto generated Go binding around an Ethereum contract.
type RoninTrustedOrg struct {
	RoninTrustedOrgCaller     // Read-only binding to the contract
	RoninTrustedOrgTransactor // Write-only binding to the contract
	RoninTrustedOrgFilterer   // Log filterer for contract events
}

// RoninTrustedOrgCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninTrustedOrgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTrustedOrgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninTrustedOrgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTrustedOrgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninTrustedOrgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTrustedOrgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninTrustedOrgSession struct {
	Contract     *RoninTrustedOrg  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoninTrustedOrgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninTrustedOrgCallerSession struct {
	Contract *RoninTrustedOrgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RoninTrustedOrgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninTrustedOrgTransactorSession struct {
	Contract     *RoninTrustedOrgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RoninTrustedOrgRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninTrustedOrgRaw struct {
	Contract *RoninTrustedOrg // Generic contract binding to access the raw methods on
}

// RoninTrustedOrgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninTrustedOrgCallerRaw struct {
	Contract *RoninTrustedOrgCaller // Generic read-only contract binding to access the raw methods on
}

// RoninTrustedOrgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninTrustedOrgTransactorRaw struct {
	Contract *RoninTrustedOrgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninTrustedOrg creates a new instance of RoninTrustedOrg, bound to a specific deployed contract.
func NewRoninTrustedOrg(address common.Address, backend bind.ContractBackend) (*RoninTrustedOrg, error) {
	contract, err := bindRoninTrustedOrg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrg{RoninTrustedOrgCaller: RoninTrustedOrgCaller{contract: contract}, RoninTrustedOrgTransactor: RoninTrustedOrgTransactor{contract: contract}, RoninTrustedOrgFilterer: RoninTrustedOrgFilterer{contract: contract}}, nil
}

// NewRoninTrustedOrgCaller creates a new read-only instance of RoninTrustedOrg, bound to a specific deployed contract.
func NewRoninTrustedOrgCaller(address common.Address, caller bind.ContractCaller) (*RoninTrustedOrgCaller, error) {
	contract, err := bindRoninTrustedOrg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgCaller{contract: contract}, nil
}

// NewRoninTrustedOrgTransactor creates a new write-only instance of RoninTrustedOrg, bound to a specific deployed contract.
func NewRoninTrustedOrgTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninTrustedOrgTransactor, error) {
	contract, err := bindRoninTrustedOrg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgTransactor{contract: contract}, nil
}

// NewRoninTrustedOrgFilterer creates a new log filterer instance of RoninTrustedOrg, bound to a specific deployed contract.
func NewRoninTrustedOrgFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninTrustedOrgFilterer, error) {
	contract, err := bindRoninTrustedOrg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgFilterer{contract: contract}, nil
}

// bindRoninTrustedOrg binds a generic wrapper to an already deployed contract.
func bindRoninTrustedOrg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninTrustedOrgABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninTrustedOrg *RoninTrustedOrgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninTrustedOrg.Contract.RoninTrustedOrgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninTrustedOrg *RoninTrustedOrgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.RoninTrustedOrgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninTrustedOrg *RoninTrustedOrgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.RoninTrustedOrgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninTrustedOrg *RoninTrustedOrgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninTrustedOrg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninTrustedOrg *RoninTrustedOrgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninTrustedOrg *RoninTrustedOrgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.contract.Transact(opts, method, params...)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_RoninTrustedOrg *RoninTrustedOrgSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _RoninTrustedOrg.Contract.CheckThreshold(&_RoninTrustedOrg.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _RoninTrustedOrg.Contract.CheckThreshold(&_RoninTrustedOrg.CallOpts, _voteWeight)
}

// CountTrustedOrganization is a free data retrieval call binding the contract method 0xb7f67e97.
//
// Solidity: function countTrustedOrganization() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) CountTrustedOrganization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "countTrustedOrganization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountTrustedOrganization is a free data retrieval call binding the contract method 0xb7f67e97.
//
// Solidity: function countTrustedOrganization() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) CountTrustedOrganization() (*big.Int, error) {
	return _RoninTrustedOrg.Contract.CountTrustedOrganization(&_RoninTrustedOrg.CallOpts)
}

// CountTrustedOrganization is a free data retrieval call binding the contract method 0xb7f67e97.
//
// Solidity: function countTrustedOrganization() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) CountTrustedOrganization() (*big.Int, error) {
	return _RoninTrustedOrg.Contract.CountTrustedOrganization(&_RoninTrustedOrg.CallOpts)
}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address,address,address,uint256,uint256)[] list)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetAllTrustedOrganizations(opts *bind.CallOpts) ([]IRoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getAllTrustedOrganizations")

	if err != nil {
		return *new([]IRoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRoninTrustedOrganizationTrustedOrganization)).(*[]IRoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address,address,address,uint256,uint256)[] list)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetAllTrustedOrganizations() ([]IRoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrg.Contract.GetAllTrustedOrganizations(&_RoninTrustedOrg.CallOpts)
}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address,address,address,uint256,uint256)[] list)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetAllTrustedOrganizations() ([]IRoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrg.Contract.GetAllTrustedOrganizations(&_RoninTrustedOrg.CallOpts)
}

// GetConsensusWeight is a free data retrieval call binding the contract method 0x41feed1c.
//
// Solidity: function getConsensusWeight(address consensusAddr) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetConsensusWeight(opts *bind.CallOpts, consensusAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getConsensusWeight", consensusAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusWeight is a free data retrieval call binding the contract method 0x41feed1c.
//
// Solidity: function getConsensusWeight(address consensusAddr) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetConsensusWeight(consensusAddr common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeight(&_RoninTrustedOrg.CallOpts, consensusAddr)
}

// GetConsensusWeight is a free data retrieval call binding the contract method 0x41feed1c.
//
// Solidity: function getConsensusWeight(address consensusAddr) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetConsensusWeight(consensusAddr common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeight(&_RoninTrustedOrg.CallOpts, consensusAddr)
}

// GetConsensusWeightById is a free data retrieval call binding the contract method 0xb8cc3a50.
//
// Solidity: function getConsensusWeightById(address cid) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetConsensusWeightById(opts *bind.CallOpts, cid common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getConsensusWeightById", cid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusWeightById is a free data retrieval call binding the contract method 0xb8cc3a50.
//
// Solidity: function getConsensusWeightById(address cid) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetConsensusWeightById(cid common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeightById(&_RoninTrustedOrg.CallOpts, cid)
}

// GetConsensusWeightById is a free data retrieval call binding the contract method 0xb8cc3a50.
//
// Solidity: function getConsensusWeightById(address cid) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetConsensusWeightById(cid common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeightById(&_RoninTrustedOrg.CallOpts, cid)
}

// GetConsensusWeights is a free data retrieval call binding the contract method 0x520fce62.
//
// Solidity: function getConsensusWeights(address[] list) view returns(uint256[])
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetConsensusWeights(opts *bind.CallOpts, list []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getConsensusWeights", list)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetConsensusWeights is a free data retrieval call binding the contract method 0x520fce62.
//
// Solidity: function getConsensusWeights(address[] list) view returns(uint256[])
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetConsensusWeights(list []common.Address) ([]*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeights(&_RoninTrustedOrg.CallOpts, list)
}

// GetConsensusWeights is a free data retrieval call binding the contract method 0x520fce62.
//
// Solidity: function getConsensusWeights(address[] list) view returns(uint256[])
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetConsensusWeights(list []common.Address) ([]*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeights(&_RoninTrustedOrg.CallOpts, list)
}

// GetConsensusWeightsById is a free data retrieval call binding the contract method 0x2cccb53c.
//
// Solidity: function getConsensusWeightsById(address[] cids) view returns(uint256[])
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetConsensusWeightsById(opts *bind.CallOpts, cids []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getConsensusWeightsById", cids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetConsensusWeightsById is a free data retrieval call binding the contract method 0x2cccb53c.
//
// Solidity: function getConsensusWeightsById(address[] cids) view returns(uint256[])
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetConsensusWeightsById(cids []common.Address) ([]*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeightsById(&_RoninTrustedOrg.CallOpts, cids)
}

// GetConsensusWeightsById is a free data retrieval call binding the contract method 0x2cccb53c.
//
// Solidity: function getConsensusWeightsById(address[] cids) view returns(uint256[])
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetConsensusWeightsById(cids []common.Address) ([]*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetConsensusWeightsById(&_RoninTrustedOrg.CallOpts, cids)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetContract(contractType uint8) (common.Address, error) {
	return _RoninTrustedOrg.Contract.GetContract(&_RoninTrustedOrg.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _RoninTrustedOrg.Contract.GetContract(&_RoninTrustedOrg.CallOpts, contractType)
}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address _governor) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetGovernorWeight(opts *bind.CallOpts, _governor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getGovernorWeight", _governor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address _governor) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetGovernorWeight(_governor common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetGovernorWeight(&_RoninTrustedOrg.CallOpts, _governor)
}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address _governor) view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetGovernorWeight(_governor common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetGovernorWeight(&_RoninTrustedOrg.CallOpts, _governor)
}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] _list) view returns(uint256[] _res)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetGovernorWeights(opts *bind.CallOpts, _list []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getGovernorWeights", _list)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] _list) view returns(uint256[] _res)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetGovernorWeights(_list []common.Address) ([]*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetGovernorWeights(&_RoninTrustedOrg.CallOpts, _list)
}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] _list) view returns(uint256[] _res)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetGovernorWeights(_list []common.Address) ([]*big.Int, error) {
	return _RoninTrustedOrg.Contract.GetGovernorWeights(&_RoninTrustedOrg.CallOpts, _list)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetThreshold(opts *bind.CallOpts) (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getThreshold")

	outstruct := new(struct {
		Num   *big.Int
		Denom *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Num = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Denom = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetThreshold() (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	return _RoninTrustedOrg.Contract.GetThreshold(&_RoninTrustedOrg.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetThreshold() (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	return _RoninTrustedOrg.Contract.GetThreshold(&_RoninTrustedOrg.CallOpts)
}

// GetTrustedOrganization is a free data retrieval call binding the contract method 0xdb6693a2.
//
// Solidity: function getTrustedOrganization(address _consensusAddr) view returns((address,address,address,uint256,uint256))
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetTrustedOrganization(opts *bind.CallOpts, _consensusAddr common.Address) (IRoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getTrustedOrganization", _consensusAddr)

	if err != nil {
		return *new(IRoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new(IRoninTrustedOrganizationTrustedOrganization)).(*IRoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetTrustedOrganization is a free data retrieval call binding the contract method 0xdb6693a2.
//
// Solidity: function getTrustedOrganization(address _consensusAddr) view returns((address,address,address,uint256,uint256))
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetTrustedOrganization(_consensusAddr common.Address) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrg.Contract.GetTrustedOrganization(&_RoninTrustedOrg.CallOpts, _consensusAddr)
}

// GetTrustedOrganization is a free data retrieval call binding the contract method 0xdb6693a2.
//
// Solidity: function getTrustedOrganization(address _consensusAddr) view returns((address,address,address,uint256,uint256))
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetTrustedOrganization(_consensusAddr common.Address) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrg.Contract.GetTrustedOrganization(&_RoninTrustedOrg.CallOpts, _consensusAddr)
}

// GetTrustedOrganizationAt is a free data retrieval call binding the contract method 0xf09267c2.
//
// Solidity: function getTrustedOrganizationAt(uint256 _idx) view returns((address,address,address,uint256,uint256))
func (_RoninTrustedOrg *RoninTrustedOrgCaller) GetTrustedOrganizationAt(opts *bind.CallOpts, _idx *big.Int) (IRoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "getTrustedOrganizationAt", _idx)

	if err != nil {
		return *new(IRoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new(IRoninTrustedOrganizationTrustedOrganization)).(*IRoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetTrustedOrganizationAt is a free data retrieval call binding the contract method 0xf09267c2.
//
// Solidity: function getTrustedOrganizationAt(uint256 _idx) view returns((address,address,address,uint256,uint256))
func (_RoninTrustedOrg *RoninTrustedOrgSession) GetTrustedOrganizationAt(_idx *big.Int) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrg.Contract.GetTrustedOrganizationAt(&_RoninTrustedOrg.CallOpts, _idx)
}

// GetTrustedOrganizationAt is a free data retrieval call binding the contract method 0xf09267c2.
//
// Solidity: function getTrustedOrganizationAt(uint256 _idx) view returns((address,address,address,uint256,uint256))
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) GetTrustedOrganizationAt(_idx *big.Int) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrg.Contract.GetTrustedOrganizationAt(&_RoninTrustedOrg.CallOpts, _idx)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) MinimumVoteWeight() (*big.Int, error) {
	return _RoninTrustedOrg.Contract.MinimumVoteWeight(&_RoninTrustedOrg.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _RoninTrustedOrg.Contract.MinimumVoteWeight(&_RoninTrustedOrg.CallOpts)
}

// SumConsensusWeight is a free data retrieval call binding the contract method 0x691845a9.
//
// Solidity: function sumConsensusWeight(address[] _list) view returns(uint256 _res)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) SumConsensusWeight(opts *bind.CallOpts, _list []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "sumConsensusWeight", _list)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumConsensusWeight is a free data retrieval call binding the contract method 0x691845a9.
//
// Solidity: function sumConsensusWeight(address[] _list) view returns(uint256 _res)
func (_RoninTrustedOrg *RoninTrustedOrgSession) SumConsensusWeight(_list []common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.SumConsensusWeight(&_RoninTrustedOrg.CallOpts, _list)
}

// SumConsensusWeight is a free data retrieval call binding the contract method 0x691845a9.
//
// Solidity: function sumConsensusWeight(address[] _list) view returns(uint256 _res)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) SumConsensusWeight(_list []common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.SumConsensusWeight(&_RoninTrustedOrg.CallOpts, _list)
}

// SumGovernorWeight is a free data retrieval call binding the contract method 0x903bb3c5.
//
// Solidity: function sumGovernorWeight(address[] _list) view returns(uint256 _res)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) SumGovernorWeight(opts *bind.CallOpts, _list []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "sumGovernorWeight", _list)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumGovernorWeight is a free data retrieval call binding the contract method 0x903bb3c5.
//
// Solidity: function sumGovernorWeight(address[] _list) view returns(uint256 _res)
func (_RoninTrustedOrg *RoninTrustedOrgSession) SumGovernorWeight(_list []common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.SumGovernorWeight(&_RoninTrustedOrg.CallOpts, _list)
}

// SumGovernorWeight is a free data retrieval call binding the contract method 0x903bb3c5.
//
// Solidity: function sumGovernorWeight(address[] _list) view returns(uint256 _res)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) SumGovernorWeight(_list []common.Address) (*big.Int, error) {
	return _RoninTrustedOrg.Contract.SumGovernorWeight(&_RoninTrustedOrg.CallOpts, _list)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninTrustedOrg.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) TotalWeight() (*big.Int, error) {
	return _RoninTrustedOrg.Contract.TotalWeight(&_RoninTrustedOrg.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_RoninTrustedOrg *RoninTrustedOrgCallerSession) TotalWeight() (*big.Int, error) {
	return _RoninTrustedOrg.Contract.TotalWeight(&_RoninTrustedOrg.CallOpts)
}

// AddTrustedOrganizations is a paid mutator transaction binding the contract method 0x0ed285df.
//
// Solidity: function addTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) AddTrustedOrganizations(opts *bind.TransactOpts, _list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "addTrustedOrganizations", _list)
}

// AddTrustedOrganizations is a paid mutator transaction binding the contract method 0x0ed285df.
//
// Solidity: function addTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) AddTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.AddTrustedOrganizations(&_RoninTrustedOrg.TransactOpts, _list)
}

// AddTrustedOrganizations is a paid mutator transaction binding the contract method 0x0ed285df.
//
// Solidity: function addTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) AddTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.AddTrustedOrganizations(&_RoninTrustedOrg.TransactOpts, _list)
}

// ExecChangeConsensusAddressForTrustedOrg is a paid mutator transaction binding the contract method 0xa0c302a7.
//
// Solidity: function execChangeConsensusAddressForTrustedOrg(address oldAddr, address newAddr) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) ExecChangeConsensusAddressForTrustedOrg(opts *bind.TransactOpts, oldAddr common.Address, newAddr common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "execChangeConsensusAddressForTrustedOrg", oldAddr, newAddr)
}

// ExecChangeConsensusAddressForTrustedOrg is a paid mutator transaction binding the contract method 0xa0c302a7.
//
// Solidity: function execChangeConsensusAddressForTrustedOrg(address oldAddr, address newAddr) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) ExecChangeConsensusAddressForTrustedOrg(oldAddr common.Address, newAddr common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.ExecChangeConsensusAddressForTrustedOrg(&_RoninTrustedOrg.TransactOpts, oldAddr, newAddr)
}

// ExecChangeConsensusAddressForTrustedOrg is a paid mutator transaction binding the contract method 0xa0c302a7.
//
// Solidity: function execChangeConsensusAddressForTrustedOrg(address oldAddr, address newAddr) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) ExecChangeConsensusAddressForTrustedOrg(oldAddr common.Address, newAddr common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.ExecChangeConsensusAddressForTrustedOrg(&_RoninTrustedOrg.TransactOpts, oldAddr, newAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x7c37103c.
//
// Solidity: function initialize((address,address,address,uint256,uint256)[] trustedOrgs, uint256 num, uint256 denom) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) Initialize(opts *bind.TransactOpts, trustedOrgs []IRoninTrustedOrganizationTrustedOrganization, num *big.Int, denom *big.Int) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "initialize", trustedOrgs, num, denom)
}

// Initialize is a paid mutator transaction binding the contract method 0x7c37103c.
//
// Solidity: function initialize((address,address,address,uint256,uint256)[] trustedOrgs, uint256 num, uint256 denom) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) Initialize(trustedOrgs []IRoninTrustedOrganizationTrustedOrganization, num *big.Int, denom *big.Int) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.Initialize(&_RoninTrustedOrg.TransactOpts, trustedOrgs, num, denom)
}

// Initialize is a paid mutator transaction binding the contract method 0x7c37103c.
//
// Solidity: function initialize((address,address,address,uint256,uint256)[] trustedOrgs, uint256 num, uint256 denom) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) Initialize(trustedOrgs []IRoninTrustedOrganizationTrustedOrganization, num *big.Int, denom *big.Int) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.Initialize(&_RoninTrustedOrg.TransactOpts, trustedOrgs, num, denom)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address profileContract) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) InitializeV2(opts *bind.TransactOpts, profileContract common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "initializeV2", profileContract)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address profileContract) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) InitializeV2(profileContract common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.InitializeV2(&_RoninTrustedOrg.TransactOpts, profileContract)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address profileContract) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) InitializeV2(profileContract common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.InitializeV2(&_RoninTrustedOrg.TransactOpts, profileContract)
}

// RemoveTrustedOrganizations is a paid mutator transaction binding the contract method 0xa85c7d6e.
//
// Solidity: function removeTrustedOrganizations(address[] list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) RemoveTrustedOrganizations(opts *bind.TransactOpts, list []common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "removeTrustedOrganizations", list)
}

// RemoveTrustedOrganizations is a paid mutator transaction binding the contract method 0xa85c7d6e.
//
// Solidity: function removeTrustedOrganizations(address[] list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) RemoveTrustedOrganizations(list []common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.RemoveTrustedOrganizations(&_RoninTrustedOrg.TransactOpts, list)
}

// RemoveTrustedOrganizations is a paid mutator transaction binding the contract method 0xa85c7d6e.
//
// Solidity: function removeTrustedOrganizations(address[] list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) RemoveTrustedOrganizations(list []common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.RemoveTrustedOrganizations(&_RoninTrustedOrg.TransactOpts, list)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.SetContract(&_RoninTrustedOrg.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.SetContract(&_RoninTrustedOrg.TransactOpts, contractType, addr)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_RoninTrustedOrg *RoninTrustedOrgSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.SetThreshold(&_RoninTrustedOrg.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.SetThreshold(&_RoninTrustedOrg.TransactOpts, _numerator, _denominator)
}

// UpdateTrustedOrganizations is a paid mutator transaction binding the contract method 0xb505a07c.
//
// Solidity: function updateTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactor) UpdateTrustedOrganizations(opts *bind.TransactOpts, _list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _RoninTrustedOrg.contract.Transact(opts, "updateTrustedOrganizations", _list)
}

// UpdateTrustedOrganizations is a paid mutator transaction binding the contract method 0xb505a07c.
//
// Solidity: function updateTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgSession) UpdateTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.UpdateTrustedOrganizations(&_RoninTrustedOrg.TransactOpts, _list)
}

// UpdateTrustedOrganizations is a paid mutator transaction binding the contract method 0xb505a07c.
//
// Solidity: function updateTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_RoninTrustedOrg *RoninTrustedOrgTransactorSession) UpdateTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _RoninTrustedOrg.Contract.UpdateTrustedOrganizations(&_RoninTrustedOrg.TransactOpts, _list)
}

// RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator is returned from FilterConsensusAddressOfTrustedOrgChanged and is used to iterate over the raw logs and unpacked data for ConsensusAddressOfTrustedOrgChanged events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator struct {
	Event *RoninTrustedOrgConsensusAddressOfTrustedOrgChanged // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgConsensusAddressOfTrustedOrgChanged)
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
		it.Event = new(RoninTrustedOrgConsensusAddressOfTrustedOrgChanged)
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
func (it *RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgConsensusAddressOfTrustedOrgChanged represents a ConsensusAddressOfTrustedOrgChanged event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgConsensusAddressOfTrustedOrgChanged struct {
	OrgAfterChanged IRoninTrustedOrganizationTrustedOrganization
	OldConsensus    common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsensusAddressOfTrustedOrgChanged is a free log retrieval operation binding the contract event 0xaf495cd00cf8595e5fc29a082a18b71cad844d97ab40237b7405693a910c50a2.
//
// Solidity: event ConsensusAddressOfTrustedOrgChanged((address,address,address,uint256,uint256) orgAfterChanged, address oldConsensus)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterConsensusAddressOfTrustedOrgChanged(opts *bind.FilterOpts) (*RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator, error) {

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "ConsensusAddressOfTrustedOrgChanged")
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgConsensusAddressOfTrustedOrgChangedIterator{contract: _RoninTrustedOrg.contract, event: "ConsensusAddressOfTrustedOrgChanged", logs: logs, sub: sub}, nil
}

// WatchConsensusAddressOfTrustedOrgChanged is a free log subscription operation binding the contract event 0xaf495cd00cf8595e5fc29a082a18b71cad844d97ab40237b7405693a910c50a2.
//
// Solidity: event ConsensusAddressOfTrustedOrgChanged((address,address,address,uint256,uint256) orgAfterChanged, address oldConsensus)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchConsensusAddressOfTrustedOrgChanged(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgConsensusAddressOfTrustedOrgChanged) (event.Subscription, error) {

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "ConsensusAddressOfTrustedOrgChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgConsensusAddressOfTrustedOrgChanged)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "ConsensusAddressOfTrustedOrgChanged", log); err != nil {
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

// ParseConsensusAddressOfTrustedOrgChanged is a log parse operation binding the contract event 0xaf495cd00cf8595e5fc29a082a18b71cad844d97ab40237b7405693a910c50a2.
//
// Solidity: event ConsensusAddressOfTrustedOrgChanged((address,address,address,uint256,uint256) orgAfterChanged, address oldConsensus)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseConsensusAddressOfTrustedOrgChanged(log types.Log) (*RoninTrustedOrgConsensusAddressOfTrustedOrgChanged, error) {
	event := new(RoninTrustedOrgConsensusAddressOfTrustedOrgChanged)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "ConsensusAddressOfTrustedOrgChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTrustedOrgContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgContractUpdatedIterator struct {
	Event *RoninTrustedOrgContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgContractUpdated)
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
		it.Event = new(RoninTrustedOrgContractUpdated)
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
func (it *RoninTrustedOrgContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgContractUpdated represents a ContractUpdated event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*RoninTrustedOrgContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgContractUpdatedIterator{contract: _RoninTrustedOrg.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgContractUpdated)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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

// ParseContractUpdated is a log parse operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseContractUpdated(log types.Log) (*RoninTrustedOrgContractUpdated, error) {
	event := new(RoninTrustedOrgContractUpdated)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTrustedOrgInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgInitializedIterator struct {
	Event *RoninTrustedOrgInitialized // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgInitialized)
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
		it.Event = new(RoninTrustedOrgInitialized)
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
func (it *RoninTrustedOrgInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgInitialized represents a Initialized event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterInitialized(opts *bind.FilterOpts) (*RoninTrustedOrgInitializedIterator, error) {

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgInitializedIterator{contract: _RoninTrustedOrg.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgInitialized) (event.Subscription, error) {

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgInitialized)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseInitialized(log types.Log) (*RoninTrustedOrgInitialized, error) {
	event := new(RoninTrustedOrgInitialized)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTrustedOrgThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgThresholdUpdatedIterator struct {
	Event *RoninTrustedOrgThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgThresholdUpdated)
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
		it.Event = new(RoninTrustedOrgThresholdUpdated)
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
func (it *RoninTrustedOrgThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgThresholdUpdated represents a ThresholdUpdated event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*RoninTrustedOrgThresholdUpdatedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgThresholdUpdatedIterator{contract: _RoninTrustedOrg.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgThresholdUpdated)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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

// ParseThresholdUpdated is a log parse operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseThresholdUpdated(log types.Log) (*RoninTrustedOrgThresholdUpdated, error) {
	event := new(RoninTrustedOrgThresholdUpdated)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTrustedOrgTrustedOrganizationsAddedIterator is returned from FilterTrustedOrganizationsAdded and is used to iterate over the raw logs and unpacked data for TrustedOrganizationsAdded events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgTrustedOrganizationsAddedIterator struct {
	Event *RoninTrustedOrgTrustedOrganizationsAdded // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgTrustedOrganizationsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgTrustedOrganizationsAdded)
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
		it.Event = new(RoninTrustedOrgTrustedOrganizationsAdded)
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
func (it *RoninTrustedOrgTrustedOrganizationsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgTrustedOrganizationsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgTrustedOrganizationsAdded represents a TrustedOrganizationsAdded event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgTrustedOrganizationsAdded struct {
	Orgs []IRoninTrustedOrganizationTrustedOrganization
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedOrganizationsAdded is a free log retrieval operation binding the contract event 0xc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c.
//
// Solidity: event TrustedOrganizationsAdded((address,address,address,uint256,uint256)[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterTrustedOrganizationsAdded(opts *bind.FilterOpts) (*RoninTrustedOrgTrustedOrganizationsAddedIterator, error) {

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "TrustedOrganizationsAdded")
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgTrustedOrganizationsAddedIterator{contract: _RoninTrustedOrg.contract, event: "TrustedOrganizationsAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedOrganizationsAdded is a free log subscription operation binding the contract event 0xc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c.
//
// Solidity: event TrustedOrganizationsAdded((address,address,address,uint256,uint256)[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchTrustedOrganizationsAdded(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgTrustedOrganizationsAdded) (event.Subscription, error) {

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "TrustedOrganizationsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgTrustedOrganizationsAdded)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "TrustedOrganizationsAdded", log); err != nil {
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

// ParseTrustedOrganizationsAdded is a log parse operation binding the contract event 0xc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c.
//
// Solidity: event TrustedOrganizationsAdded((address,address,address,uint256,uint256)[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseTrustedOrganizationsAdded(log types.Log) (*RoninTrustedOrgTrustedOrganizationsAdded, error) {
	event := new(RoninTrustedOrgTrustedOrganizationsAdded)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "TrustedOrganizationsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTrustedOrgTrustedOrganizationsRemovedIterator is returned from FilterTrustedOrganizationsRemoved and is used to iterate over the raw logs and unpacked data for TrustedOrganizationsRemoved events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgTrustedOrganizationsRemovedIterator struct {
	Event *RoninTrustedOrgTrustedOrganizationsRemoved // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgTrustedOrganizationsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgTrustedOrganizationsRemoved)
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
		it.Event = new(RoninTrustedOrgTrustedOrganizationsRemoved)
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
func (it *RoninTrustedOrgTrustedOrganizationsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgTrustedOrganizationsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgTrustedOrganizationsRemoved represents a TrustedOrganizationsRemoved event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgTrustedOrganizationsRemoved struct {
	Orgs []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedOrganizationsRemoved is a free log retrieval operation binding the contract event 0x121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a7.
//
// Solidity: event TrustedOrganizationsRemoved(address[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterTrustedOrganizationsRemoved(opts *bind.FilterOpts) (*RoninTrustedOrgTrustedOrganizationsRemovedIterator, error) {

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "TrustedOrganizationsRemoved")
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgTrustedOrganizationsRemovedIterator{contract: _RoninTrustedOrg.contract, event: "TrustedOrganizationsRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedOrganizationsRemoved is a free log subscription operation binding the contract event 0x121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a7.
//
// Solidity: event TrustedOrganizationsRemoved(address[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchTrustedOrganizationsRemoved(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgTrustedOrganizationsRemoved) (event.Subscription, error) {

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "TrustedOrganizationsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgTrustedOrganizationsRemoved)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "TrustedOrganizationsRemoved", log); err != nil {
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

// ParseTrustedOrganizationsRemoved is a log parse operation binding the contract event 0x121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a7.
//
// Solidity: event TrustedOrganizationsRemoved(address[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseTrustedOrganizationsRemoved(log types.Log) (*RoninTrustedOrgTrustedOrganizationsRemoved, error) {
	event := new(RoninTrustedOrgTrustedOrganizationsRemoved)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "TrustedOrganizationsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTrustedOrgTrustedOrganizationsUpdatedIterator is returned from FilterTrustedOrganizationsUpdated and is used to iterate over the raw logs and unpacked data for TrustedOrganizationsUpdated events raised by the RoninTrustedOrg contract.
type RoninTrustedOrgTrustedOrganizationsUpdatedIterator struct {
	Event *RoninTrustedOrgTrustedOrganizationsUpdated // Event containing the contract specifics and raw log

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
func (it *RoninTrustedOrgTrustedOrganizationsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTrustedOrgTrustedOrganizationsUpdated)
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
		it.Event = new(RoninTrustedOrgTrustedOrganizationsUpdated)
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
func (it *RoninTrustedOrgTrustedOrganizationsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTrustedOrgTrustedOrganizationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTrustedOrgTrustedOrganizationsUpdated represents a TrustedOrganizationsUpdated event raised by the RoninTrustedOrg contract.
type RoninTrustedOrgTrustedOrganizationsUpdated struct {
	Orgs []IRoninTrustedOrganizationTrustedOrganization
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedOrganizationsUpdated is a free log retrieval operation binding the contract event 0xe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda.
//
// Solidity: event TrustedOrganizationsUpdated((address,address,address,uint256,uint256)[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) FilterTrustedOrganizationsUpdated(opts *bind.FilterOpts) (*RoninTrustedOrgTrustedOrganizationsUpdatedIterator, error) {

	logs, sub, err := _RoninTrustedOrg.contract.FilterLogs(opts, "TrustedOrganizationsUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrgTrustedOrganizationsUpdatedIterator{contract: _RoninTrustedOrg.contract, event: "TrustedOrganizationsUpdated", logs: logs, sub: sub}, nil
}

// WatchTrustedOrganizationsUpdated is a free log subscription operation binding the contract event 0xe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda.
//
// Solidity: event TrustedOrganizationsUpdated((address,address,address,uint256,uint256)[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) WatchTrustedOrganizationsUpdated(opts *bind.WatchOpts, sink chan<- *RoninTrustedOrgTrustedOrganizationsUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninTrustedOrg.contract.WatchLogs(opts, "TrustedOrganizationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTrustedOrgTrustedOrganizationsUpdated)
				if err := _RoninTrustedOrg.contract.UnpackLog(event, "TrustedOrganizationsUpdated", log); err != nil {
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

// ParseTrustedOrganizationsUpdated is a log parse operation binding the contract event 0xe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda.
//
// Solidity: event TrustedOrganizationsUpdated((address,address,address,uint256,uint256)[] orgs)
func (_RoninTrustedOrg *RoninTrustedOrgFilterer) ParseTrustedOrganizationsUpdated(log types.Log) (*RoninTrustedOrgTrustedOrganizationsUpdated, error) {
	event := new(RoninTrustedOrgTrustedOrganizationsUpdated)
	if err := _RoninTrustedOrg.contract.UnpackLog(event, "TrustedOrganizationsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
