// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vote

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

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"passport_contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uit_event\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"promt\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"promt_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidate_total_votes\",\"type\":\"uint256\"}],\"name\":\"ENSVoteCommited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uit_event\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"promt\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"promt_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidate_total_votes\",\"type\":\"uint256\"}],\"name\":\"ENS_T3P_VoteCommited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"promt\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"promt_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidate_total_votes\",\"type\":\"uint256\"}],\"name\":\"FreeVoteCommited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumVote.VoteType\",\"name\":\"vote_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumPassport.PassportType\",\"name\":\"id_type_required\",\"type\":\"uint8\"}],\"name\":\"VoteCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"promt_choice\",\"type\":\"string\"}],\"name\":\"CommitChoiceENSValid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"promt_choice\",\"type\":\"string\"}],\"name\":\"CommitChoiceFreePromt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"promt_choice\",\"type\":\"string\"}],\"name\":\"CommitChoice_ENS_and_T3P\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"T3P\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_vote\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"option\",\"type\":\"bytes32\"}],\"name\":\"VoteResultsFreePromt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"checkENS_User_by_string\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"}],\"name\":\"checkVoteExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"}],\"name\":\"checkVotePhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orginiser_or_ens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start_date_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vote_hours\",\"type\":\"uint256\"},{\"internalType\":\"enumPassport.PassportType\",\"name\":\"id_type_required\",\"type\":\"uint8\"},{\"internalType\":\"enumVote.VoteType\",\"name\":\"vote_type\",\"type\":\"uint8\"}],\"name\":\"createNewVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"uid_event\",\"type\":\"uint256\"}],\"name\":\"getVoteStatus\",\"outputs\":[{\"internalType\":\"enumVote.Phase\",\"name\":\"ph\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moderator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"uid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ens_domain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time_to_vote_hours\",\"type\":\"uint256\"},{\"internalType\":\"enumPassport.PassportType\",\"name\":\"id_type_required\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"votes_total\",\"type\":\"uint256\"},{\"internalType\":\"enumVote.VoteType\",\"name\":\"vote_type\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vote *VoteCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vote *VoteSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vote.Contract.DEFAULTADMINROLE(&_Vote.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vote *VoteCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vote.Contract.DEFAULTADMINROLE(&_Vote.CallOpts)
}

// T3P is a free data retrieval call binding the contract method 0xc0d1bb94.
//
// Solidity: function T3P() view returns(bytes32)
func (_Vote *VoteCaller) T3P(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "T3P")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// T3P is a free data retrieval call binding the contract method 0xc0d1bb94.
//
// Solidity: function T3P() view returns(bytes32)
func (_Vote *VoteSession) T3P() ([32]byte, error) {
	return _Vote.Contract.T3P(&_Vote.CallOpts)
}

// T3P is a free data retrieval call binding the contract method 0xc0d1bb94.
//
// Solidity: function T3P() view returns(bytes32)
func (_Vote *VoteCallerSession) T3P() ([32]byte, error) {
	return _Vote.Contract.T3P(&_Vote.CallOpts)
}

// VoteResultsFreePromt is a free data retrieval call binding the contract method 0x5883e6d8.
//
// Solidity: function VoteResultsFreePromt(uint256 uid_vote, bytes32 option) view returns(uint256 count)
func (_Vote *VoteCaller) VoteResultsFreePromt(opts *bind.CallOpts, uid_vote *big.Int, option [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "VoteResultsFreePromt", uid_vote, option)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteResultsFreePromt is a free data retrieval call binding the contract method 0x5883e6d8.
//
// Solidity: function VoteResultsFreePromt(uint256 uid_vote, bytes32 option) view returns(uint256 count)
func (_Vote *VoteSession) VoteResultsFreePromt(uid_vote *big.Int, option [32]byte) (*big.Int, error) {
	return _Vote.Contract.VoteResultsFreePromt(&_Vote.CallOpts, uid_vote, option)
}

// VoteResultsFreePromt is a free data retrieval call binding the contract method 0x5883e6d8.
//
// Solidity: function VoteResultsFreePromt(uint256 uid_vote, bytes32 option) view returns(uint256 count)
func (_Vote *VoteCallerSession) VoteResultsFreePromt(uid_vote *big.Int, option [32]byte) (*big.Int, error) {
	return _Vote.Contract.VoteResultsFreePromt(&_Vote.CallOpts, uid_vote, option)
}

// CheckENSUserByString is a free data retrieval call binding the contract method 0x8fbfc0ab.
//
// Solidity: function checkENS_User_by_string(string name) view returns(address)
func (_Vote *VoteCaller) CheckENSUserByString(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "checkENS_User_by_string", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckENSUserByString is a free data retrieval call binding the contract method 0x8fbfc0ab.
//
// Solidity: function checkENS_User_by_string(string name) view returns(address)
func (_Vote *VoteSession) CheckENSUserByString(name string) (common.Address, error) {
	return _Vote.Contract.CheckENSUserByString(&_Vote.CallOpts, name)
}

// CheckENSUserByString is a free data retrieval call binding the contract method 0x8fbfc0ab.
//
// Solidity: function checkENS_User_by_string(string name) view returns(address)
func (_Vote *VoteCallerSession) CheckENSUserByString(name string) (common.Address, error) {
	return _Vote.Contract.CheckENSUserByString(&_Vote.CallOpts, name)
}

// CheckVoteExist is a free data retrieval call binding the contract method 0xfd49165b.
//
// Solidity: function checkVoteExist(uint256 uid_event) view returns(bool)
func (_Vote *VoteCaller) CheckVoteExist(opts *bind.CallOpts, uid_event *big.Int) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "checkVoteExist", uid_event)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckVoteExist is a free data retrieval call binding the contract method 0xfd49165b.
//
// Solidity: function checkVoteExist(uint256 uid_event) view returns(bool)
func (_Vote *VoteSession) CheckVoteExist(uid_event *big.Int) (bool, error) {
	return _Vote.Contract.CheckVoteExist(&_Vote.CallOpts, uid_event)
}

// CheckVoteExist is a free data retrieval call binding the contract method 0xfd49165b.
//
// Solidity: function checkVoteExist(uint256 uid_event) view returns(bool)
func (_Vote *VoteCallerSession) CheckVoteExist(uid_event *big.Int) (bool, error) {
	return _Vote.Contract.CheckVoteExist(&_Vote.CallOpts, uid_event)
}

// CheckVotePhase is a free data retrieval call binding the contract method 0x48778484.
//
// Solidity: function checkVotePhase(uint256 uid_event) view returns(bool)
func (_Vote *VoteCaller) CheckVotePhase(opts *bind.CallOpts, uid_event *big.Int) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "checkVotePhase", uid_event)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckVotePhase is a free data retrieval call binding the contract method 0x48778484.
//
// Solidity: function checkVotePhase(uint256 uid_event) view returns(bool)
func (_Vote *VoteSession) CheckVotePhase(uid_event *big.Int) (bool, error) {
	return _Vote.Contract.CheckVotePhase(&_Vote.CallOpts, uid_event)
}

// CheckVotePhase is a free data retrieval call binding the contract method 0x48778484.
//
// Solidity: function checkVotePhase(uint256 uid_event) view returns(bool)
func (_Vote *VoteCallerSession) CheckVotePhase(uid_event *big.Int) (bool, error) {
	return _Vote.Contract.CheckVotePhase(&_Vote.CallOpts, uid_event)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vote *VoteCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vote *VoteSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vote.Contract.GetRoleAdmin(&_Vote.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vote *VoteCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vote.Contract.GetRoleAdmin(&_Vote.CallOpts, role)
}

// GetVoteStatus is a free data retrieval call binding the contract method 0x0519bb83.
//
// Solidity: function getVoteStatus(uint256 uid_event) view returns(uint8 ph)
func (_Vote *VoteCaller) GetVoteStatus(opts *bind.CallOpts, uid_event *big.Int) (uint8, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getVoteStatus", uid_event)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetVoteStatus is a free data retrieval call binding the contract method 0x0519bb83.
//
// Solidity: function getVoteStatus(uint256 uid_event) view returns(uint8 ph)
func (_Vote *VoteSession) GetVoteStatus(uid_event *big.Int) (uint8, error) {
	return _Vote.Contract.GetVoteStatus(&_Vote.CallOpts, uid_event)
}

// GetVoteStatus is a free data retrieval call binding the contract method 0x0519bb83.
//
// Solidity: function getVoteStatus(uint256 uid_event) view returns(uint8 ph)
func (_Vote *VoteCallerSession) GetVoteStatus(uid_event *big.Int) (uint8, error) {
	return _Vote.Contract.GetVoteStatus(&_Vote.CallOpts, uid_event)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vote *VoteCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vote *VoteSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vote.Contract.HasRole(&_Vote.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vote *VoteCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vote.Contract.HasRole(&_Vote.CallOpts, role, account)
}

// Moderator is a free data retrieval call binding the contract method 0x38743904.
//
// Solidity: function moderator() view returns(bytes32)
func (_Vote *VoteCaller) Moderator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "moderator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Moderator is a free data retrieval call binding the contract method 0x38743904.
//
// Solidity: function moderator() view returns(bytes32)
func (_Vote *VoteSession) Moderator() ([32]byte, error) {
	return _Vote.Contract.Moderator(&_Vote.CallOpts)
}

// Moderator is a free data retrieval call binding the contract method 0x38743904.
//
// Solidity: function moderator() view returns(bytes32)
func (_Vote *VoteCallerSession) Moderator() ([32]byte, error) {
	return _Vote.Contract.Moderator(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCallerSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_Vote *VoteCaller) Resolve(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "resolve", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_Vote *VoteSession) Resolve(node [32]byte) (common.Address, error) {
	return _Vote.Contract.Resolve(&_Vote.CallOpts, node)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_Vote *VoteCallerSession) Resolve(node [32]byte) (common.Address, error) {
	return _Vote.Contract.Resolve(&_Vote.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vote *VoteCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vote *VoteSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vote.Contract.SupportsInterface(&_Vote.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vote *VoteCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vote.Contract.SupportsInterface(&_Vote.CallOpts, interfaceId)
}

// Votings is a free data retrieval call binding the contract method 0xa598d03c.
//
// Solidity: function votings(uint256 ) view returns(uint256 uid, address ens_domain, address operator, string desc, uint256 start_date, uint256 time_to_vote_hours, uint8 id_type_required, uint256 votes_total, uint8 vote_type)
func (_Vote *VoteCaller) Votings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Uid             *big.Int
	EnsDomain       common.Address
	Operator        common.Address
	Desc            string
	StartDate       *big.Int
	TimeToVoteHours *big.Int
	IdTypeRequired  uint8
	VotesTotal      *big.Int
	VoteType        uint8
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "votings", arg0)

	outstruct := new(struct {
		Uid             *big.Int
		EnsDomain       common.Address
		Operator        common.Address
		Desc            string
		StartDate       *big.Int
		TimeToVoteHours *big.Int
		IdTypeRequired  uint8
		VotesTotal      *big.Int
		VoteType        uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EnsDomain = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Desc = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.StartDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TimeToVoteHours = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IdTypeRequired = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.VotesTotal = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VoteType = *abi.ConvertType(out[8], new(uint8)).(*uint8)

	return *outstruct, err

}

// Votings is a free data retrieval call binding the contract method 0xa598d03c.
//
// Solidity: function votings(uint256 ) view returns(uint256 uid, address ens_domain, address operator, string desc, uint256 start_date, uint256 time_to_vote_hours, uint8 id_type_required, uint256 votes_total, uint8 vote_type)
func (_Vote *VoteSession) Votings(arg0 *big.Int) (struct {
	Uid             *big.Int
	EnsDomain       common.Address
	Operator        common.Address
	Desc            string
	StartDate       *big.Int
	TimeToVoteHours *big.Int
	IdTypeRequired  uint8
	VotesTotal      *big.Int
	VoteType        uint8
}, error) {
	return _Vote.Contract.Votings(&_Vote.CallOpts, arg0)
}

// Votings is a free data retrieval call binding the contract method 0xa598d03c.
//
// Solidity: function votings(uint256 ) view returns(uint256 uid, address ens_domain, address operator, string desc, uint256 start_date, uint256 time_to_vote_hours, uint8 id_type_required, uint256 votes_total, uint8 vote_type)
func (_Vote *VoteCallerSession) Votings(arg0 *big.Int) (struct {
	Uid             *big.Int
	EnsDomain       common.Address
	Operator        common.Address
	Desc            string
	StartDate       *big.Int
	TimeToVoteHours *big.Int
	IdTypeRequired  uint8
	VotesTotal      *big.Int
	VoteType        uint8
}, error) {
	return _Vote.Contract.Votings(&_Vote.CallOpts, arg0)
}

// CommitChoiceENSValid is a paid mutator transaction binding the contract method 0x652913e2.
//
// Solidity: function CommitChoiceENSValid(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteTransactor) CommitChoiceENSValid(opts *bind.TransactOpts, uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "CommitChoiceENSValid", uid_event, promt_choice)
}

// CommitChoiceENSValid is a paid mutator transaction binding the contract method 0x652913e2.
//
// Solidity: function CommitChoiceENSValid(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteSession) CommitChoiceENSValid(uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.Contract.CommitChoiceENSValid(&_Vote.TransactOpts, uid_event, promt_choice)
}

// CommitChoiceENSValid is a paid mutator transaction binding the contract method 0x652913e2.
//
// Solidity: function CommitChoiceENSValid(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteTransactorSession) CommitChoiceENSValid(uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.Contract.CommitChoiceENSValid(&_Vote.TransactOpts, uid_event, promt_choice)
}

// CommitChoiceFreePromt is a paid mutator transaction binding the contract method 0x40400081.
//
// Solidity: function CommitChoiceFreePromt(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteTransactor) CommitChoiceFreePromt(opts *bind.TransactOpts, uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "CommitChoiceFreePromt", uid_event, promt_choice)
}

// CommitChoiceFreePromt is a paid mutator transaction binding the contract method 0x40400081.
//
// Solidity: function CommitChoiceFreePromt(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteSession) CommitChoiceFreePromt(uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.Contract.CommitChoiceFreePromt(&_Vote.TransactOpts, uid_event, promt_choice)
}

// CommitChoiceFreePromt is a paid mutator transaction binding the contract method 0x40400081.
//
// Solidity: function CommitChoiceFreePromt(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteTransactorSession) CommitChoiceFreePromt(uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.Contract.CommitChoiceFreePromt(&_Vote.TransactOpts, uid_event, promt_choice)
}

// CommitChoiceENSAndT3P is a paid mutator transaction binding the contract method 0x236b3ad9.
//
// Solidity: function CommitChoice_ENS_and_T3P(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteTransactor) CommitChoiceENSAndT3P(opts *bind.TransactOpts, uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "CommitChoice_ENS_and_T3P", uid_event, promt_choice)
}

// CommitChoiceENSAndT3P is a paid mutator transaction binding the contract method 0x236b3ad9.
//
// Solidity: function CommitChoice_ENS_and_T3P(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteSession) CommitChoiceENSAndT3P(uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.Contract.CommitChoiceENSAndT3P(&_Vote.TransactOpts, uid_event, promt_choice)
}

// CommitChoiceENSAndT3P is a paid mutator transaction binding the contract method 0x236b3ad9.
//
// Solidity: function CommitChoice_ENS_and_T3P(uint256 uid_event, string promt_choice) returns()
func (_Vote *VoteTransactorSession) CommitChoiceENSAndT3P(uid_event *big.Int, promt_choice string) (*types.Transaction, error) {
	return _Vote.Contract.CommitChoiceENSAndT3P(&_Vote.TransactOpts, uid_event, promt_choice)
}

// CreateNewVote is a paid mutator transaction binding the contract method 0x138f3cfd.
//
// Solidity: function createNewVote(address orginiser_or_ens, address operator, uint256 start_date_timestamp, uint256 vote_hours, uint8 id_type_required, uint8 vote_type) returns(uint256)
func (_Vote *VoteTransactor) CreateNewVote(opts *bind.TransactOpts, orginiser_or_ens common.Address, operator common.Address, start_date_timestamp *big.Int, vote_hours *big.Int, id_type_required uint8, vote_type uint8) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "createNewVote", orginiser_or_ens, operator, start_date_timestamp, vote_hours, id_type_required, vote_type)
}

// CreateNewVote is a paid mutator transaction binding the contract method 0x138f3cfd.
//
// Solidity: function createNewVote(address orginiser_or_ens, address operator, uint256 start_date_timestamp, uint256 vote_hours, uint8 id_type_required, uint8 vote_type) returns(uint256)
func (_Vote *VoteSession) CreateNewVote(orginiser_or_ens common.Address, operator common.Address, start_date_timestamp *big.Int, vote_hours *big.Int, id_type_required uint8, vote_type uint8) (*types.Transaction, error) {
	return _Vote.Contract.CreateNewVote(&_Vote.TransactOpts, orginiser_or_ens, operator, start_date_timestamp, vote_hours, id_type_required, vote_type)
}

// CreateNewVote is a paid mutator transaction binding the contract method 0x138f3cfd.
//
// Solidity: function createNewVote(address orginiser_or_ens, address operator, uint256 start_date_timestamp, uint256 vote_hours, uint8 id_type_required, uint8 vote_type) returns(uint256)
func (_Vote *VoteTransactorSession) CreateNewVote(orginiser_or_ens common.Address, operator common.Address, start_date_timestamp *big.Int, vote_hours *big.Int, id_type_required uint8, vote_type uint8) (*types.Transaction, error) {
	return _Vote.Contract.CreateNewVote(&_Vote.TransactOpts, orginiser_or_ens, operator, start_date_timestamp, vote_hours, id_type_required, vote_type)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vote *VoteTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vote *VoteSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GrantRole(&_Vote.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vote *VoteTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vote.Contract.GrantRole(&_Vote.TransactOpts, role, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vote *VoteTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vote *VoteSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vote.Contract.RenounceOwnership(&_Vote.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vote *VoteTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vote.Contract.RenounceOwnership(&_Vote.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Vote *VoteTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Vote *VoteSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Vote.Contract.RenounceRole(&_Vote.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Vote *VoteTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Vote.Contract.RenounceRole(&_Vote.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vote *VoteTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vote *VoteSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vote.Contract.RevokeRole(&_Vote.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vote *VoteTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vote.Contract.RevokeRole(&_Vote.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vote *VoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vote *VoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.TransferOwnership(&_Vote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vote *VoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.TransferOwnership(&_Vote.TransactOpts, newOwner)
}

// VoteENSVoteCommitedIterator is returned from FilterENSVoteCommited and is used to iterate over the raw logs and unpacked data for ENSVoteCommited events raised by the Vote contract.
type VoteENSVoteCommitedIterator struct {
	Event *VoteENSVoteCommited // Event containing the contract specifics and raw log

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
func (it *VoteENSVoteCommitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteENSVoteCommited)
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
		it.Event = new(VoteENSVoteCommited)
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
func (it *VoteENSVoteCommitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteENSVoteCommitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteENSVoteCommited represents a ENSVoteCommited event raised by the Vote contract.
type VoteENSVoteCommited struct {
	UitEvent            *big.Int
	Promt               string
	PromtHash           [32]byte
	CandidateTotalVotes *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterENSVoteCommited is a free log retrieval operation binding the contract event 0xbfdc205034116aadc8301a4690f813afc2105a39fc790bc22c74f36524af74af.
//
// Solidity: event ENSVoteCommited(uint256 indexed uit_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) FilterENSVoteCommited(opts *bind.FilterOpts, uit_event []*big.Int, promt_hash [][32]byte) (*VoteENSVoteCommitedIterator, error) {

	var uit_eventRule []interface{}
	for _, uit_eventItem := range uit_event {
		uit_eventRule = append(uit_eventRule, uit_eventItem)
	}

	var promt_hashRule []interface{}
	for _, promt_hashItem := range promt_hash {
		promt_hashRule = append(promt_hashRule, promt_hashItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "ENSVoteCommited", uit_eventRule, promt_hashRule)
	if err != nil {
		return nil, err
	}
	return &VoteENSVoteCommitedIterator{contract: _Vote.contract, event: "ENSVoteCommited", logs: logs, sub: sub}, nil
}

// WatchENSVoteCommited is a free log subscription operation binding the contract event 0xbfdc205034116aadc8301a4690f813afc2105a39fc790bc22c74f36524af74af.
//
// Solidity: event ENSVoteCommited(uint256 indexed uit_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) WatchENSVoteCommited(opts *bind.WatchOpts, sink chan<- *VoteENSVoteCommited, uit_event []*big.Int, promt_hash [][32]byte) (event.Subscription, error) {

	var uit_eventRule []interface{}
	for _, uit_eventItem := range uit_event {
		uit_eventRule = append(uit_eventRule, uit_eventItem)
	}

	var promt_hashRule []interface{}
	for _, promt_hashItem := range promt_hash {
		promt_hashRule = append(promt_hashRule, promt_hashItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "ENSVoteCommited", uit_eventRule, promt_hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteENSVoteCommited)
				if err := _Vote.contract.UnpackLog(event, "ENSVoteCommited", log); err != nil {
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

// ParseENSVoteCommited is a log parse operation binding the contract event 0xbfdc205034116aadc8301a4690f813afc2105a39fc790bc22c74f36524af74af.
//
// Solidity: event ENSVoteCommited(uint256 indexed uit_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) ParseENSVoteCommited(log types.Log) (*VoteENSVoteCommited, error) {
	event := new(VoteENSVoteCommited)
	if err := _Vote.contract.UnpackLog(event, "ENSVoteCommited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteENST3PVoteCommitedIterator is returned from FilterENST3PVoteCommited and is used to iterate over the raw logs and unpacked data for ENST3PVoteCommited events raised by the Vote contract.
type VoteENST3PVoteCommitedIterator struct {
	Event *VoteENST3PVoteCommited // Event containing the contract specifics and raw log

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
func (it *VoteENST3PVoteCommitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteENST3PVoteCommited)
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
		it.Event = new(VoteENST3PVoteCommited)
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
func (it *VoteENST3PVoteCommitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteENST3PVoteCommitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteENST3PVoteCommited represents a ENST3PVoteCommited event raised by the Vote contract.
type VoteENST3PVoteCommited struct {
	UitEvent            *big.Int
	Promt               string
	PromtHash           [32]byte
	CandidateTotalVotes *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterENST3PVoteCommited is a free log retrieval operation binding the contract event 0x34a30b725ce5926dad1fb8ea9b85937b6cdc418faef7eb4b1295d6071e4b6212.
//
// Solidity: event ENS_T3P_VoteCommited(uint256 indexed uit_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) FilterENST3PVoteCommited(opts *bind.FilterOpts, uit_event []*big.Int, promt_hash [][32]byte) (*VoteENST3PVoteCommitedIterator, error) {

	var uit_eventRule []interface{}
	for _, uit_eventItem := range uit_event {
		uit_eventRule = append(uit_eventRule, uit_eventItem)
	}

	var promt_hashRule []interface{}
	for _, promt_hashItem := range promt_hash {
		promt_hashRule = append(promt_hashRule, promt_hashItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "ENS_T3P_VoteCommited", uit_eventRule, promt_hashRule)
	if err != nil {
		return nil, err
	}
	return &VoteENST3PVoteCommitedIterator{contract: _Vote.contract, event: "ENS_T3P_VoteCommited", logs: logs, sub: sub}, nil
}

// WatchENST3PVoteCommited is a free log subscription operation binding the contract event 0x34a30b725ce5926dad1fb8ea9b85937b6cdc418faef7eb4b1295d6071e4b6212.
//
// Solidity: event ENS_T3P_VoteCommited(uint256 indexed uit_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) WatchENST3PVoteCommited(opts *bind.WatchOpts, sink chan<- *VoteENST3PVoteCommited, uit_event []*big.Int, promt_hash [][32]byte) (event.Subscription, error) {

	var uit_eventRule []interface{}
	for _, uit_eventItem := range uit_event {
		uit_eventRule = append(uit_eventRule, uit_eventItem)
	}

	var promt_hashRule []interface{}
	for _, promt_hashItem := range promt_hash {
		promt_hashRule = append(promt_hashRule, promt_hashItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "ENS_T3P_VoteCommited", uit_eventRule, promt_hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteENST3PVoteCommited)
				if err := _Vote.contract.UnpackLog(event, "ENS_T3P_VoteCommited", log); err != nil {
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

// ParseENST3PVoteCommited is a log parse operation binding the contract event 0x34a30b725ce5926dad1fb8ea9b85937b6cdc418faef7eb4b1295d6071e4b6212.
//
// Solidity: event ENS_T3P_VoteCommited(uint256 indexed uit_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) ParseENST3PVoteCommited(log types.Log) (*VoteENST3PVoteCommited, error) {
	event := new(VoteENST3PVoteCommited)
	if err := _Vote.contract.UnpackLog(event, "ENS_T3P_VoteCommited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteFreeVoteCommitedIterator is returned from FilterFreeVoteCommited and is used to iterate over the raw logs and unpacked data for FreeVoteCommited events raised by the Vote contract.
type VoteFreeVoteCommitedIterator struct {
	Event *VoteFreeVoteCommited // Event containing the contract specifics and raw log

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
func (it *VoteFreeVoteCommitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteFreeVoteCommited)
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
		it.Event = new(VoteFreeVoteCommited)
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
func (it *VoteFreeVoteCommitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteFreeVoteCommitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteFreeVoteCommited represents a FreeVoteCommited event raised by the Vote contract.
type VoteFreeVoteCommited struct {
	UidEvent            *big.Int
	Promt               string
	PromtHash           [32]byte
	CandidateTotalVotes *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFreeVoteCommited is a free log retrieval operation binding the contract event 0x25478a6e22c5e96409b84d17c38a6a747d52a504443b11294f187fb1692079b0.
//
// Solidity: event FreeVoteCommited(uint256 indexed uid_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) FilterFreeVoteCommited(opts *bind.FilterOpts, uid_event []*big.Int, promt_hash [][32]byte) (*VoteFreeVoteCommitedIterator, error) {

	var uid_eventRule []interface{}
	for _, uid_eventItem := range uid_event {
		uid_eventRule = append(uid_eventRule, uid_eventItem)
	}

	var promt_hashRule []interface{}
	for _, promt_hashItem := range promt_hash {
		promt_hashRule = append(promt_hashRule, promt_hashItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "FreeVoteCommited", uid_eventRule, promt_hashRule)
	if err != nil {
		return nil, err
	}
	return &VoteFreeVoteCommitedIterator{contract: _Vote.contract, event: "FreeVoteCommited", logs: logs, sub: sub}, nil
}

// WatchFreeVoteCommited is a free log subscription operation binding the contract event 0x25478a6e22c5e96409b84d17c38a6a747d52a504443b11294f187fb1692079b0.
//
// Solidity: event FreeVoteCommited(uint256 indexed uid_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) WatchFreeVoteCommited(opts *bind.WatchOpts, sink chan<- *VoteFreeVoteCommited, uid_event []*big.Int, promt_hash [][32]byte) (event.Subscription, error) {

	var uid_eventRule []interface{}
	for _, uid_eventItem := range uid_event {
		uid_eventRule = append(uid_eventRule, uid_eventItem)
	}

	var promt_hashRule []interface{}
	for _, promt_hashItem := range promt_hash {
		promt_hashRule = append(promt_hashRule, promt_hashItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "FreeVoteCommited", uid_eventRule, promt_hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteFreeVoteCommited)
				if err := _Vote.contract.UnpackLog(event, "FreeVoteCommited", log); err != nil {
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

// ParseFreeVoteCommited is a log parse operation binding the contract event 0x25478a6e22c5e96409b84d17c38a6a747d52a504443b11294f187fb1692079b0.
//
// Solidity: event FreeVoteCommited(uint256 indexed uid_event, string promt, bytes32 indexed promt_hash, uint256 candidate_total_votes)
func (_Vote *VoteFilterer) ParseFreeVoteCommited(log types.Log) (*VoteFreeVoteCommited, error) {
	event := new(VoteFreeVoteCommited)
	if err := _Vote.contract.UnpackLog(event, "FreeVoteCommited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vote contract.
type VoteOwnershipTransferredIterator struct {
	Event *VoteOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VoteOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteOwnershipTransferred)
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
		it.Event = new(VoteOwnershipTransferred)
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
func (it *VoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteOwnershipTransferred represents a OwnershipTransferred event raised by the Vote contract.
type VoteOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vote *VoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VoteOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoteOwnershipTransferredIterator{contract: _Vote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vote *VoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VoteOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteOwnershipTransferred)
				if err := _Vote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Vote *VoteFilterer) ParseOwnershipTransferred(log types.Log) (*VoteOwnershipTransferred, error) {
	event := new(VoteOwnershipTransferred)
	if err := _Vote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vote contract.
type VoteRoleAdminChangedIterator struct {
	Event *VoteRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoteRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteRoleAdminChanged)
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
		it.Event = new(VoteRoleAdminChanged)
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
func (it *VoteRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteRoleAdminChanged represents a RoleAdminChanged event raised by the Vote contract.
type VoteRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vote *VoteFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VoteRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Vote.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VoteRoleAdminChangedIterator{contract: _Vote.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vote *VoteFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VoteRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Vote.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteRoleAdminChanged)
				if err := _Vote.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Vote *VoteFilterer) ParseRoleAdminChanged(log types.Log) (*VoteRoleAdminChanged, error) {
	event := new(VoteRoleAdminChanged)
	if err := _Vote.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vote contract.
type VoteRoleGrantedIterator struct {
	Event *VoteRoleGranted // Event containing the contract specifics and raw log

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
func (it *VoteRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteRoleGranted)
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
		it.Event = new(VoteRoleGranted)
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
func (it *VoteRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteRoleGranted represents a RoleGranted event raised by the Vote contract.
type VoteRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vote *VoteFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoteRoleGrantedIterator, error) {

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

	logs, sub, err := _Vote.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoteRoleGrantedIterator{contract: _Vote.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vote *VoteFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VoteRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vote.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteRoleGranted)
				if err := _Vote.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Vote *VoteFilterer) ParseRoleGranted(log types.Log) (*VoteRoleGranted, error) {
	event := new(VoteRoleGranted)
	if err := _Vote.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vote contract.
type VoteRoleRevokedIterator struct {
	Event *VoteRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VoteRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteRoleRevoked)
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
		it.Event = new(VoteRoleRevoked)
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
func (it *VoteRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteRoleRevoked represents a RoleRevoked event raised by the Vote contract.
type VoteRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vote *VoteFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoteRoleRevokedIterator, error) {

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

	logs, sub, err := _Vote.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoteRoleRevokedIterator{contract: _Vote.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vote *VoteFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VoteRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vote.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteRoleRevoked)
				if err := _Vote.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Vote *VoteFilterer) ParseRoleRevoked(log types.Log) (*VoteRoleRevoked, error) {
	event := new(VoteRoleRevoked)
	if err := _Vote.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteVoteCreatedIterator is returned from FilterVoteCreated and is used to iterate over the raw logs and unpacked data for VoteCreated events raised by the Vote contract.
type VoteVoteCreatedIterator struct {
	Event *VoteVoteCreated // Event containing the contract specifics and raw log

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
func (it *VoteVoteCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteVoteCreated)
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
		it.Event = new(VoteVoteCreated)
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
func (it *VoteVoteCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteVoteCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteVoteCreated represents a VoteCreated event raised by the Vote contract.
type VoteVoteCreated struct {
	UidEvent       *big.Int
	VoteType       uint8
	IdTypeRequired uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVoteCreated is a free log retrieval operation binding the contract event 0x0a480e9461032fc4b7d6311035853883db94fd7654aec8b0a9e81b3f886457b3.
//
// Solidity: event VoteCreated(uint256 uid_event, uint8 vote_type, uint8 id_type_required)
func (_Vote *VoteFilterer) FilterVoteCreated(opts *bind.FilterOpts) (*VoteVoteCreatedIterator, error) {

	logs, sub, err := _Vote.contract.FilterLogs(opts, "VoteCreated")
	if err != nil {
		return nil, err
	}
	return &VoteVoteCreatedIterator{contract: _Vote.contract, event: "VoteCreated", logs: logs, sub: sub}, nil
}

// WatchVoteCreated is a free log subscription operation binding the contract event 0x0a480e9461032fc4b7d6311035853883db94fd7654aec8b0a9e81b3f886457b3.
//
// Solidity: event VoteCreated(uint256 uid_event, uint8 vote_type, uint8 id_type_required)
func (_Vote *VoteFilterer) WatchVoteCreated(opts *bind.WatchOpts, sink chan<- *VoteVoteCreated) (event.Subscription, error) {

	logs, sub, err := _Vote.contract.WatchLogs(opts, "VoteCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteVoteCreated)
				if err := _Vote.contract.UnpackLog(event, "VoteCreated", log); err != nil {
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

// ParseVoteCreated is a log parse operation binding the contract event 0x0a480e9461032fc4b7d6311035853883db94fd7654aec8b0a9e81b3f886457b3.
//
// Solidity: event VoteCreated(uint256 uid_event, uint8 vote_type, uint8 id_type_required)
func (_Vote *VoteFilterer) ParseVoteCreated(log types.Log) (*VoteVoteCreated, error) {
	event := new(VoteVoteCreated)
	if err := _Vote.contract.UnpackLog(event, "VoteCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
