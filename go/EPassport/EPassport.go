// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EPassport

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

// EPassportMetaData contains all meta data concerning the EPassport contract.
var EPassportMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"enumPassport.PassportType\",\"name\":\"id_type\",\"type\":\"uint8\"}],\"name\":\"CheckUserHavePassedTTP_ByAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_address\",\"type\":\"address\"},{\"internalType\":\"enumPassport.PassportType\",\"name\":\"id_type\",\"type\":\"uint8\"}],\"name\":\"CheckUserHaveTypeIDByAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"text\",\"type\":\"string\"}],\"name\":\"GetKeccakHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumPassport.PassportType\",\"name\":\"id_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"mrz_uid_hash\",\"type\":\"bytes32\"}],\"name\":\"ProoveUserByTTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"T3P\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"service_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ENS_address\",\"type\":\"address\"}],\"name\":\"addNewTTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mrz_uid_hash\",\"type\":\"bytes32\"}],\"name\":\"checkUserHavePassedTTP_ByHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moderator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"service_address\",\"type\":\"address\"}],\"name\":\"pauseTTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"service_address\",\"type\":\"address\"}],\"name\":\"switchTTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EPassportABI is the input ABI used to generate the binding from.
// Deprecated: Use EPassportMetaData.ABI instead.
var EPassportABI = EPassportMetaData.ABI

// EPassport is an auto generated Go binding around an Ethereum contract.
type EPassport struct {
	EPassportCaller     // Read-only binding to the contract
	EPassportTransactor // Write-only binding to the contract
	EPassportFilterer   // Log filterer for contract events
}

// EPassportCaller is an auto generated read-only Go binding around an Ethereum contract.
type EPassportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPassportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EPassportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPassportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EPassportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPassportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EPassportSession struct {
	Contract     *EPassport        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EPassportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EPassportCallerSession struct {
	Contract *EPassportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EPassportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EPassportTransactorSession struct {
	Contract     *EPassportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EPassportRaw is an auto generated low-level Go binding around an Ethereum contract.
type EPassportRaw struct {
	Contract *EPassport // Generic contract binding to access the raw methods on
}

// EPassportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EPassportCallerRaw struct {
	Contract *EPassportCaller // Generic read-only contract binding to access the raw methods on
}

// EPassportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EPassportTransactorRaw struct {
	Contract *EPassportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEPassport creates a new instance of EPassport, bound to a specific deployed contract.
func NewEPassport(address common.Address, backend bind.ContractBackend) (*EPassport, error) {
	contract, err := bindEPassport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EPassport{EPassportCaller: EPassportCaller{contract: contract}, EPassportTransactor: EPassportTransactor{contract: contract}, EPassportFilterer: EPassportFilterer{contract: contract}}, nil
}

// NewEPassportCaller creates a new read-only instance of EPassport, bound to a specific deployed contract.
func NewEPassportCaller(address common.Address, caller bind.ContractCaller) (*EPassportCaller, error) {
	contract, err := bindEPassport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EPassportCaller{contract: contract}, nil
}

// NewEPassportTransactor creates a new write-only instance of EPassport, bound to a specific deployed contract.
func NewEPassportTransactor(address common.Address, transactor bind.ContractTransactor) (*EPassportTransactor, error) {
	contract, err := bindEPassport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EPassportTransactor{contract: contract}, nil
}

// NewEPassportFilterer creates a new log filterer instance of EPassport, bound to a specific deployed contract.
func NewEPassportFilterer(address common.Address, filterer bind.ContractFilterer) (*EPassportFilterer, error) {
	contract, err := bindEPassport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EPassportFilterer{contract: contract}, nil
}

// bindEPassport binds a generic wrapper to an already deployed contract.
func bindEPassport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EPassportMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPassport *EPassportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPassport.Contract.EPassportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPassport *EPassportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPassport.Contract.EPassportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPassport *EPassportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPassport.Contract.EPassportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPassport *EPassportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPassport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPassport *EPassportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPassport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPassport *EPassportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPassport.Contract.contract.Transact(opts, method, params...)
}

// CheckUserHavePassedTTPByAddr is a free data retrieval call binding the contract method 0x64ec7aa3.
//
// Solidity: function CheckUserHavePassedTTP_ByAddr(address user_address, uint8 id_type) view returns(bool)
func (_EPassport *EPassportCaller) CheckUserHavePassedTTPByAddr(opts *bind.CallOpts, user_address common.Address, id_type uint8) (bool, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "CheckUserHavePassedTTP_ByAddr", user_address, id_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserHavePassedTTPByAddr is a free data retrieval call binding the contract method 0x64ec7aa3.
//
// Solidity: function CheckUserHavePassedTTP_ByAddr(address user_address, uint8 id_type) view returns(bool)
func (_EPassport *EPassportSession) CheckUserHavePassedTTPByAddr(user_address common.Address, id_type uint8) (bool, error) {
	return _EPassport.Contract.CheckUserHavePassedTTPByAddr(&_EPassport.CallOpts, user_address, id_type)
}

// CheckUserHavePassedTTPByAddr is a free data retrieval call binding the contract method 0x64ec7aa3.
//
// Solidity: function CheckUserHavePassedTTP_ByAddr(address user_address, uint8 id_type) view returns(bool)
func (_EPassport *EPassportCallerSession) CheckUserHavePassedTTPByAddr(user_address common.Address, id_type uint8) (bool, error) {
	return _EPassport.Contract.CheckUserHavePassedTTPByAddr(&_EPassport.CallOpts, user_address, id_type)
}

// CheckUserHaveTypeIDByAddr is a free data retrieval call binding the contract method 0x59f463c2.
//
// Solidity: function CheckUserHaveTypeIDByAddr(address user_address, uint8 id_type) view returns(bool)
func (_EPassport *EPassportCaller) CheckUserHaveTypeIDByAddr(opts *bind.CallOpts, user_address common.Address, id_type uint8) (bool, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "CheckUserHaveTypeIDByAddr", user_address, id_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserHaveTypeIDByAddr is a free data retrieval call binding the contract method 0x59f463c2.
//
// Solidity: function CheckUserHaveTypeIDByAddr(address user_address, uint8 id_type) view returns(bool)
func (_EPassport *EPassportSession) CheckUserHaveTypeIDByAddr(user_address common.Address, id_type uint8) (bool, error) {
	return _EPassport.Contract.CheckUserHaveTypeIDByAddr(&_EPassport.CallOpts, user_address, id_type)
}

// CheckUserHaveTypeIDByAddr is a free data retrieval call binding the contract method 0x59f463c2.
//
// Solidity: function CheckUserHaveTypeIDByAddr(address user_address, uint8 id_type) view returns(bool)
func (_EPassport *EPassportCallerSession) CheckUserHaveTypeIDByAddr(user_address common.Address, id_type uint8) (bool, error) {
	return _EPassport.Contract.CheckUserHaveTypeIDByAddr(&_EPassport.CallOpts, user_address, id_type)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EPassport *EPassportCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EPassport *EPassportSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EPassport.Contract.DEFAULTADMINROLE(&_EPassport.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EPassport *EPassportCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EPassport.Contract.DEFAULTADMINROLE(&_EPassport.CallOpts)
}

// GetKeccakHash is a free data retrieval call binding the contract method 0x9e026a3b.
//
// Solidity: function GetKeccakHash(string text) pure returns(bytes32)
func (_EPassport *EPassportCaller) GetKeccakHash(opts *bind.CallOpts, text string) ([32]byte, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "GetKeccakHash", text)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeccakHash is a free data retrieval call binding the contract method 0x9e026a3b.
//
// Solidity: function GetKeccakHash(string text) pure returns(bytes32)
func (_EPassport *EPassportSession) GetKeccakHash(text string) ([32]byte, error) {
	return _EPassport.Contract.GetKeccakHash(&_EPassport.CallOpts, text)
}

// GetKeccakHash is a free data retrieval call binding the contract method 0x9e026a3b.
//
// Solidity: function GetKeccakHash(string text) pure returns(bytes32)
func (_EPassport *EPassportCallerSession) GetKeccakHash(text string) ([32]byte, error) {
	return _EPassport.Contract.GetKeccakHash(&_EPassport.CallOpts, text)
}

// T3P is a free data retrieval call binding the contract method 0xc0d1bb94.
//
// Solidity: function T3P() view returns(bytes32)
func (_EPassport *EPassportCaller) T3P(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "T3P")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// T3P is a free data retrieval call binding the contract method 0xc0d1bb94.
//
// Solidity: function T3P() view returns(bytes32)
func (_EPassport *EPassportSession) T3P() ([32]byte, error) {
	return _EPassport.Contract.T3P(&_EPassport.CallOpts)
}

// T3P is a free data retrieval call binding the contract method 0xc0d1bb94.
//
// Solidity: function T3P() view returns(bytes32)
func (_EPassport *EPassportCallerSession) T3P() ([32]byte, error) {
	return _EPassport.Contract.T3P(&_EPassport.CallOpts)
}

// CheckUserHavePassedTTPByHash is a free data retrieval call binding the contract method 0x39c1d6b2.
//
// Solidity: function checkUserHavePassedTTP_ByHash(bytes32 mrz_uid_hash) view returns(bool)
func (_EPassport *EPassportCaller) CheckUserHavePassedTTPByHash(opts *bind.CallOpts, mrz_uid_hash [32]byte) (bool, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "checkUserHavePassedTTP_ByHash", mrz_uid_hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUserHavePassedTTPByHash is a free data retrieval call binding the contract method 0x39c1d6b2.
//
// Solidity: function checkUserHavePassedTTP_ByHash(bytes32 mrz_uid_hash) view returns(bool)
func (_EPassport *EPassportSession) CheckUserHavePassedTTPByHash(mrz_uid_hash [32]byte) (bool, error) {
	return _EPassport.Contract.CheckUserHavePassedTTPByHash(&_EPassport.CallOpts, mrz_uid_hash)
}

// CheckUserHavePassedTTPByHash is a free data retrieval call binding the contract method 0x39c1d6b2.
//
// Solidity: function checkUserHavePassedTTP_ByHash(bytes32 mrz_uid_hash) view returns(bool)
func (_EPassport *EPassportCallerSession) CheckUserHavePassedTTPByHash(mrz_uid_hash [32]byte) (bool, error) {
	return _EPassport.Contract.CheckUserHavePassedTTPByHash(&_EPassport.CallOpts, mrz_uid_hash)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EPassport *EPassportCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EPassport *EPassportSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EPassport.Contract.GetRoleAdmin(&_EPassport.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EPassport *EPassportCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EPassport.Contract.GetRoleAdmin(&_EPassport.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EPassport *EPassportCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EPassport *EPassportSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EPassport.Contract.HasRole(&_EPassport.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EPassport *EPassportCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EPassport.Contract.HasRole(&_EPassport.CallOpts, role, account)
}

// Moderator is a free data retrieval call binding the contract method 0x38743904.
//
// Solidity: function moderator() view returns(bytes32)
func (_EPassport *EPassportCaller) Moderator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "moderator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Moderator is a free data retrieval call binding the contract method 0x38743904.
//
// Solidity: function moderator() view returns(bytes32)
func (_EPassport *EPassportSession) Moderator() ([32]byte, error) {
	return _EPassport.Contract.Moderator(&_EPassport.CallOpts)
}

// Moderator is a free data retrieval call binding the contract method 0x38743904.
//
// Solidity: function moderator() view returns(bytes32)
func (_EPassport *EPassportCallerSession) Moderator() ([32]byte, error) {
	return _EPassport.Contract.Moderator(&_EPassport.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPassport *EPassportCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPassport *EPassportSession) Owner() (common.Address, error) {
	return _EPassport.Contract.Owner(&_EPassport.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPassport *EPassportCallerSession) Owner() (common.Address, error) {
	return _EPassport.Contract.Owner(&_EPassport.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EPassport *EPassportCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EPassport.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EPassport *EPassportSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EPassport.Contract.SupportsInterface(&_EPassport.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EPassport *EPassportCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EPassport.Contract.SupportsInterface(&_EPassport.CallOpts, interfaceId)
}

// ProoveUserByTTP is a paid mutator transaction binding the contract method 0xd32cb95f.
//
// Solidity: function ProoveUserByTTP(address user, uint8 id_type, bytes32 mrz_uid_hash) returns()
func (_EPassport *EPassportTransactor) ProoveUserByTTP(opts *bind.TransactOpts, user common.Address, id_type uint8, mrz_uid_hash [32]byte) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "ProoveUserByTTP", user, id_type, mrz_uid_hash)
}

// ProoveUserByTTP is a paid mutator transaction binding the contract method 0xd32cb95f.
//
// Solidity: function ProoveUserByTTP(address user, uint8 id_type, bytes32 mrz_uid_hash) returns()
func (_EPassport *EPassportSession) ProoveUserByTTP(user common.Address, id_type uint8, mrz_uid_hash [32]byte) (*types.Transaction, error) {
	return _EPassport.Contract.ProoveUserByTTP(&_EPassport.TransactOpts, user, id_type, mrz_uid_hash)
}

// ProoveUserByTTP is a paid mutator transaction binding the contract method 0xd32cb95f.
//
// Solidity: function ProoveUserByTTP(address user, uint8 id_type, bytes32 mrz_uid_hash) returns()
func (_EPassport *EPassportTransactorSession) ProoveUserByTTP(user common.Address, id_type uint8, mrz_uid_hash [32]byte) (*types.Transaction, error) {
	return _EPassport.Contract.ProoveUserByTTP(&_EPassport.TransactOpts, user, id_type, mrz_uid_hash)
}

// AddNewTTP is a paid mutator transaction binding the contract method 0x67bde528.
//
// Solidity: function addNewTTP(address service_address, address ENS_address) returns()
func (_EPassport *EPassportTransactor) AddNewTTP(opts *bind.TransactOpts, service_address common.Address, ENS_address common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "addNewTTP", service_address, ENS_address)
}

// AddNewTTP is a paid mutator transaction binding the contract method 0x67bde528.
//
// Solidity: function addNewTTP(address service_address, address ENS_address) returns()
func (_EPassport *EPassportSession) AddNewTTP(service_address common.Address, ENS_address common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.AddNewTTP(&_EPassport.TransactOpts, service_address, ENS_address)
}

// AddNewTTP is a paid mutator transaction binding the contract method 0x67bde528.
//
// Solidity: function addNewTTP(address service_address, address ENS_address) returns()
func (_EPassport *EPassportTransactorSession) AddNewTTP(service_address common.Address, ENS_address common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.AddNewTTP(&_EPassport.TransactOpts, service_address, ENS_address)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EPassport *EPassportTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EPassport *EPassportSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.GrantRole(&_EPassport.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EPassport *EPassportTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.GrantRole(&_EPassport.TransactOpts, role, account)
}

// PauseTTP is a paid mutator transaction binding the contract method 0x74767ecc.
//
// Solidity: function pauseTTP(address service_address) returns()
func (_EPassport *EPassportTransactor) PauseTTP(opts *bind.TransactOpts, service_address common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "pauseTTP", service_address)
}

// PauseTTP is a paid mutator transaction binding the contract method 0x74767ecc.
//
// Solidity: function pauseTTP(address service_address) returns()
func (_EPassport *EPassportSession) PauseTTP(service_address common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.PauseTTP(&_EPassport.TransactOpts, service_address)
}

// PauseTTP is a paid mutator transaction binding the contract method 0x74767ecc.
//
// Solidity: function pauseTTP(address service_address) returns()
func (_EPassport *EPassportTransactorSession) PauseTTP(service_address common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.PauseTTP(&_EPassport.TransactOpts, service_address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPassport *EPassportTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPassport *EPassportSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPassport.Contract.RenounceOwnership(&_EPassport.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPassport *EPassportTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPassport.Contract.RenounceOwnership(&_EPassport.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EPassport *EPassportTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EPassport *EPassportSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.RenounceRole(&_EPassport.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EPassport *EPassportTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.RenounceRole(&_EPassport.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EPassport *EPassportTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EPassport *EPassportSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.RevokeRole(&_EPassport.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EPassport *EPassportTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.RevokeRole(&_EPassport.TransactOpts, role, account)
}

// SwitchTTP is a paid mutator transaction binding the contract method 0x79072169.
//
// Solidity: function switchTTP(address service_address) returns()
func (_EPassport *EPassportTransactor) SwitchTTP(opts *bind.TransactOpts, service_address common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "switchTTP", service_address)
}

// SwitchTTP is a paid mutator transaction binding the contract method 0x79072169.
//
// Solidity: function switchTTP(address service_address) returns()
func (_EPassport *EPassportSession) SwitchTTP(service_address common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.SwitchTTP(&_EPassport.TransactOpts, service_address)
}

// SwitchTTP is a paid mutator transaction binding the contract method 0x79072169.
//
// Solidity: function switchTTP(address service_address) returns()
func (_EPassport *EPassportTransactorSession) SwitchTTP(service_address common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.SwitchTTP(&_EPassport.TransactOpts, service_address)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPassport *EPassportTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EPassport.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPassport *EPassportSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.TransferOwnership(&_EPassport.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPassport *EPassportTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPassport.Contract.TransferOwnership(&_EPassport.TransactOpts, newOwner)
}

// EPassportOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EPassport contract.
type EPassportOwnershipTransferredIterator struct {
	Event *EPassportOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EPassportOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPassportOwnershipTransferred)
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
		it.Event = new(EPassportOwnershipTransferred)
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
func (it *EPassportOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPassportOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPassportOwnershipTransferred represents a OwnershipTransferred event raised by the EPassport contract.
type EPassportOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPassport *EPassportFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EPassportOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPassport.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EPassportOwnershipTransferredIterator{contract: _EPassport.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPassport *EPassportFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EPassportOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPassport.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPassportOwnershipTransferred)
				if err := _EPassport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EPassport *EPassportFilterer) ParseOwnershipTransferred(log types.Log) (*EPassportOwnershipTransferred, error) {
	event := new(EPassportOwnershipTransferred)
	if err := _EPassport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EPassportRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EPassport contract.
type EPassportRoleAdminChangedIterator struct {
	Event *EPassportRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EPassportRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPassportRoleAdminChanged)
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
		it.Event = new(EPassportRoleAdminChanged)
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
func (it *EPassportRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPassportRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPassportRoleAdminChanged represents a RoleAdminChanged event raised by the EPassport contract.
type EPassportRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EPassport *EPassportFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EPassportRoleAdminChangedIterator, error) {

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

	logs, sub, err := _EPassport.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EPassportRoleAdminChangedIterator{contract: _EPassport.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EPassport *EPassportFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EPassportRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _EPassport.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPassportRoleAdminChanged)
				if err := _EPassport.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_EPassport *EPassportFilterer) ParseRoleAdminChanged(log types.Log) (*EPassportRoleAdminChanged, error) {
	event := new(EPassportRoleAdminChanged)
	if err := _EPassport.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EPassportRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EPassport contract.
type EPassportRoleGrantedIterator struct {
	Event *EPassportRoleGranted // Event containing the contract specifics and raw log

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
func (it *EPassportRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPassportRoleGranted)
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
		it.Event = new(EPassportRoleGranted)
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
func (it *EPassportRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPassportRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPassportRoleGranted represents a RoleGranted event raised by the EPassport contract.
type EPassportRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EPassport *EPassportFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EPassportRoleGrantedIterator, error) {

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

	logs, sub, err := _EPassport.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EPassportRoleGrantedIterator{contract: _EPassport.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EPassport *EPassportFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EPassportRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EPassport.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPassportRoleGranted)
				if err := _EPassport.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_EPassport *EPassportFilterer) ParseRoleGranted(log types.Log) (*EPassportRoleGranted, error) {
	event := new(EPassportRoleGranted)
	if err := _EPassport.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EPassportRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EPassport contract.
type EPassportRoleRevokedIterator struct {
	Event *EPassportRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EPassportRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPassportRoleRevoked)
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
		it.Event = new(EPassportRoleRevoked)
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
func (it *EPassportRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPassportRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPassportRoleRevoked represents a RoleRevoked event raised by the EPassport contract.
type EPassportRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EPassport *EPassportFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EPassportRoleRevokedIterator, error) {

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

	logs, sub, err := _EPassport.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EPassportRoleRevokedIterator{contract: _EPassport.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EPassport *EPassportFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EPassportRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EPassport.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPassportRoleRevoked)
				if err := _EPassport.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_EPassport *EPassportFilterer) ParseRoleRevoked(log types.Log) (*EPassportRoleRevoked, error) {
	event := new(EPassportRoleRevoked)
	if err := _EPassport.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
