// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// PreimageOracleMetaData contains all meta data concerning the PreimageOracle contract.
var PreimageOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"part\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"cheat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_preimage\",\"type\":\"bytes\"}],\"name\":\"computeKeccak256PreimageKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_partOffset\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_preimage\",\"type\":\"bytes\"}],\"name\":\"loadKeccak256PreimagePart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bootInfo\",\"type\":\"bytes\"}],\"name\":\"loadLocalData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"preimageLengths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preimagePartOk\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preimageParts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"}],\"name\":\"readPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"dat_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"datLen_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610afd806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063e03110e11161005b578063e03110e114610131578063e159261114610159578063fe4ac08e1461016c578063fef2b4ed146101e157600080fd5b80633eba58ce1461008d57806361238bde146100a25780638542cf50146100e0578063dfb776671461011e575b600080fd5b6100a061009b3660046107c7565b610201565b005b6100cd6100b0366004610847565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b61010e6100ee366004610847565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100d7565b6100cd61012c3660046108b2565b6104b3565b61014461013f366004610847565b610512565b604080519283526020830191909152016100d7565b6100a06101673660046108f4565b610603565b6100a061017a366004610940565b6000838152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558684528252808320968352958152858220939093559283529082905291902055565b6100cd6101ef366004610972565b60006020819052908152604090205481565b6000806000806000808680602001905181019061021e91906109fa565b9550955095509550955095506102e4565b81517f0100000000000000000000000000000000000000000000000000000000000000808317600090815233602052604090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff16176020840160005b838110156102dc57808201516000848152600260209081526040808320858452825280832060019081905587845282528083208584528252808320939093558582528181529190208590550161028b565b505050505050565b7f0100000000000000000000000000000000000000000000000000000000000000600081815233602081815260408084207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811686178086526002808552838720878052855283872060019081905582885280865284882088805286528488208f90559187528685528387208590557f0100000000000000000000000000000000000000000000000000000000000001875285855283872083168817808852818652848820888052865284882083905580885282865284882088805286528488208e905587528685528387208590557f0100000000000000000000000000000000000000000000000000000000000002875285855283872083168817808852818652848820888052865284882083905580885282865284882088805286528488208d905587528685528387208590557f01000000000000000000000000000000000000000000000000000000000000038752948452828620909116909517808552928252808420848052825280842085905582845293815283832083805281528383208790559082528181529190205561049f82600461022f565b6104aa81600561022f565b50505050505050565b60243560c081901b608052600090608881858237207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0200000000000000000000000000000000000000000000000000000000000000179392505050565b6000828152600260209081526040808320848452909152812054819060ff1661059b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b50600083815260208181526040909120546105b7816008610ac1565b6105c2856020610ac1565b106105e057836105d3826008610ac1565b6105dd9190610ad9565b91505b506000938452600160209081526040808620948652939052919092205492909150565b6044356000806008830186111561061957600080fd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561077957610779610703565b604052919050565b600067ffffffffffffffff82111561079b5761079b610703565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000602082840312156107d957600080fd5b813567ffffffffffffffff8111156107f057600080fd5b8201601f8101841361080157600080fd5b803561081461080f82610781565b610732565b81815285602083850101111561082957600080fd5b81602084016020830137600091810160200191909152949350505050565b6000806040838503121561085a57600080fd5b50508035926020909101359150565b60008083601f84011261087b57600080fd5b50813567ffffffffffffffff81111561089357600080fd5b6020830191508360208285010111156108ab57600080fd5b9250929050565b600080602083850312156108c557600080fd5b823567ffffffffffffffff8111156108dc57600080fd5b6108e885828601610869565b90969095509350505050565b60008060006040848603121561090957600080fd5b83359250602084013567ffffffffffffffff81111561092757600080fd5b61093386828701610869565b9497909650939450505050565b6000806000806080858703121561095657600080fd5b5050823594602084013594506040840135936060013592509050565b60006020828403121561098457600080fd5b5035919050565b600082601f83011261099c57600080fd5b81516109aa61080f82610781565b818152602085818487010111156109c057600080fd5b60005b838110156109de5785810182015183820183015281016109c3565b838111156109ef5760008285850101525b509095945050505050565b60008060008060008060c08789031215610a1357600080fd5b865195506020870151945060408701519350606087015167ffffffffffffffff8082168214610a4157600080fd5b608089015191945080821115610a5657600080fd5b610a628a838b0161098b565b935060a0890151915080821115610a7857600080fd5b50610a8589828a0161098b565b9150509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610ad457610ad4610a92565b500190565b600082821015610aeb57610aeb610a92565b50039056fea164736f6c634300080f000a",
}

// PreimageOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PreimageOracleMetaData.ABI instead.
var PreimageOracleABI = PreimageOracleMetaData.ABI

// PreimageOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PreimageOracleMetaData.Bin instead.
var PreimageOracleBin = PreimageOracleMetaData.Bin

// DeployPreimageOracle deploys a new Ethereum contract, binding an instance of PreimageOracle to it.
func DeployPreimageOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PreimageOracle, error) {
	parsed, err := PreimageOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PreimageOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PreimageOracle{PreimageOracleCaller: PreimageOracleCaller{contract: contract}, PreimageOracleTransactor: PreimageOracleTransactor{contract: contract}, PreimageOracleFilterer: PreimageOracleFilterer{contract: contract}}, nil
}

// PreimageOracle is an auto generated Go binding around an Ethereum contract.
type PreimageOracle struct {
	PreimageOracleCaller     // Read-only binding to the contract
	PreimageOracleTransactor // Write-only binding to the contract
	PreimageOracleFilterer   // Log filterer for contract events
}

// PreimageOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreimageOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreimageOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreimageOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreimageOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreimageOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreimageOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreimageOracleSession struct {
	Contract     *PreimageOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreimageOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreimageOracleCallerSession struct {
	Contract *PreimageOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PreimageOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreimageOracleTransactorSession struct {
	Contract     *PreimageOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PreimageOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreimageOracleRaw struct {
	Contract *PreimageOracle // Generic contract binding to access the raw methods on
}

// PreimageOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreimageOracleCallerRaw struct {
	Contract *PreimageOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PreimageOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreimageOracleTransactorRaw struct {
	Contract *PreimageOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreimageOracle creates a new instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracle(address common.Address, backend bind.ContractBackend) (*PreimageOracle, error) {
	contract, err := bindPreimageOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreimageOracle{PreimageOracleCaller: PreimageOracleCaller{contract: contract}, PreimageOracleTransactor: PreimageOracleTransactor{contract: contract}, PreimageOracleFilterer: PreimageOracleFilterer{contract: contract}}, nil
}

// NewPreimageOracleCaller creates a new read-only instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracleCaller(address common.Address, caller bind.ContractCaller) (*PreimageOracleCaller, error) {
	contract, err := bindPreimageOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreimageOracleCaller{contract: contract}, nil
}

// NewPreimageOracleTransactor creates a new write-only instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PreimageOracleTransactor, error) {
	contract, err := bindPreimageOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreimageOracleTransactor{contract: contract}, nil
}

// NewPreimageOracleFilterer creates a new log filterer instance of PreimageOracle, bound to a specific deployed contract.
func NewPreimageOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PreimageOracleFilterer, error) {
	contract, err := bindPreimageOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreimageOracleFilterer{contract: contract}, nil
}

// bindPreimageOracle binds a generic wrapper to an already deployed contract.
func bindPreimageOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreimageOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreimageOracle *PreimageOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreimageOracle.Contract.PreimageOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreimageOracle *PreimageOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreimageOracle.Contract.PreimageOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreimageOracle *PreimageOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreimageOracle.Contract.PreimageOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreimageOracle *PreimageOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreimageOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreimageOracle *PreimageOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreimageOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreimageOracle *PreimageOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreimageOracle.Contract.contract.Transact(opts, method, params...)
}

// ComputeKeccak256PreimageKey is a free data retrieval call binding the contract method 0xdfb77667.
//
// Solidity: function computeKeccak256PreimageKey(bytes _preimage) pure returns(bytes32 key_)
func (_PreimageOracle *PreimageOracleCaller) ComputeKeccak256PreimageKey(opts *bind.CallOpts, _preimage []byte) ([32]byte, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "computeKeccak256PreimageKey", _preimage)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeKeccak256PreimageKey is a free data retrieval call binding the contract method 0xdfb77667.
//
// Solidity: function computeKeccak256PreimageKey(bytes _preimage) pure returns(bytes32 key_)
func (_PreimageOracle *PreimageOracleSession) ComputeKeccak256PreimageKey(_preimage []byte) ([32]byte, error) {
	return _PreimageOracle.Contract.ComputeKeccak256PreimageKey(&_PreimageOracle.CallOpts, _preimage)
}

// ComputeKeccak256PreimageKey is a free data retrieval call binding the contract method 0xdfb77667.
//
// Solidity: function computeKeccak256PreimageKey(bytes _preimage) pure returns(bytes32 key_)
func (_PreimageOracle *PreimageOracleCallerSession) ComputeKeccak256PreimageKey(_preimage []byte) ([32]byte, error) {
	return _PreimageOracle.Contract.ComputeKeccak256PreimageKey(&_PreimageOracle.CallOpts, _preimage)
}

// PreimageLengths is a free data retrieval call binding the contract method 0xfef2b4ed.
//
// Solidity: function preimageLengths(bytes32 ) view returns(uint256)
func (_PreimageOracle *PreimageOracleCaller) PreimageLengths(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "preimageLengths", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreimageLengths is a free data retrieval call binding the contract method 0xfef2b4ed.
//
// Solidity: function preimageLengths(bytes32 ) view returns(uint256)
func (_PreimageOracle *PreimageOracleSession) PreimageLengths(arg0 [32]byte) (*big.Int, error) {
	return _PreimageOracle.Contract.PreimageLengths(&_PreimageOracle.CallOpts, arg0)
}

// PreimageLengths is a free data retrieval call binding the contract method 0xfef2b4ed.
//
// Solidity: function preimageLengths(bytes32 ) view returns(uint256)
func (_PreimageOracle *PreimageOracleCallerSession) PreimageLengths(arg0 [32]byte) (*big.Int, error) {
	return _PreimageOracle.Contract.PreimageLengths(&_PreimageOracle.CallOpts, arg0)
}

// PreimagePartOk is a free data retrieval call binding the contract method 0x8542cf50.
//
// Solidity: function preimagePartOk(bytes32 , uint256 ) view returns(bool)
func (_PreimageOracle *PreimageOracleCaller) PreimagePartOk(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "preimagePartOk", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PreimagePartOk is a free data retrieval call binding the contract method 0x8542cf50.
//
// Solidity: function preimagePartOk(bytes32 , uint256 ) view returns(bool)
func (_PreimageOracle *PreimageOracleSession) PreimagePartOk(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _PreimageOracle.Contract.PreimagePartOk(&_PreimageOracle.CallOpts, arg0, arg1)
}

// PreimagePartOk is a free data retrieval call binding the contract method 0x8542cf50.
//
// Solidity: function preimagePartOk(bytes32 , uint256 ) view returns(bool)
func (_PreimageOracle *PreimageOracleCallerSession) PreimagePartOk(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _PreimageOracle.Contract.PreimagePartOk(&_PreimageOracle.CallOpts, arg0, arg1)
}

// PreimageParts is a free data retrieval call binding the contract method 0x61238bde.
//
// Solidity: function preimageParts(bytes32 , uint256 ) view returns(bytes32)
func (_PreimageOracle *PreimageOracleCaller) PreimageParts(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "preimageParts", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PreimageParts is a free data retrieval call binding the contract method 0x61238bde.
//
// Solidity: function preimageParts(bytes32 , uint256 ) view returns(bytes32)
func (_PreimageOracle *PreimageOracleSession) PreimageParts(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _PreimageOracle.Contract.PreimageParts(&_PreimageOracle.CallOpts, arg0, arg1)
}

// PreimageParts is a free data retrieval call binding the contract method 0x61238bde.
//
// Solidity: function preimageParts(bytes32 , uint256 ) view returns(bytes32)
func (_PreimageOracle *PreimageOracleCallerSession) PreimageParts(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _PreimageOracle.Contract.PreimageParts(&_PreimageOracle.CallOpts, arg0, arg1)
}

// ReadPreimage is a free data retrieval call binding the contract method 0xe03110e1.
//
// Solidity: function readPreimage(bytes32 _key, uint256 _offset) view returns(bytes32 dat_, uint256 datLen_)
func (_PreimageOracle *PreimageOracleCaller) ReadPreimage(opts *bind.CallOpts, _key [32]byte, _offset *big.Int) (struct {
	Dat    [32]byte
	DatLen *big.Int
}, error) {
	var out []interface{}
	err := _PreimageOracle.contract.Call(opts, &out, "readPreimage", _key, _offset)

	outstruct := new(struct {
		Dat    [32]byte
		DatLen *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dat = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DatLen = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReadPreimage is a free data retrieval call binding the contract method 0xe03110e1.
//
// Solidity: function readPreimage(bytes32 _key, uint256 _offset) view returns(bytes32 dat_, uint256 datLen_)
func (_PreimageOracle *PreimageOracleSession) ReadPreimage(_key [32]byte, _offset *big.Int) (struct {
	Dat    [32]byte
	DatLen *big.Int
}, error) {
	return _PreimageOracle.Contract.ReadPreimage(&_PreimageOracle.CallOpts, _key, _offset)
}

// ReadPreimage is a free data retrieval call binding the contract method 0xe03110e1.
//
// Solidity: function readPreimage(bytes32 _key, uint256 _offset) view returns(bytes32 dat_, uint256 datLen_)
func (_PreimageOracle *PreimageOracleCallerSession) ReadPreimage(_key [32]byte, _offset *big.Int) (struct {
	Dat    [32]byte
	DatLen *big.Int
}, error) {
	return _PreimageOracle.Contract.ReadPreimage(&_PreimageOracle.CallOpts, _key, _offset)
}

// Cheat is a paid mutator transaction binding the contract method 0xfe4ac08e.
//
// Solidity: function cheat(uint256 partOffset, bytes32 key, bytes32 part, uint256 size) returns()
func (_PreimageOracle *PreimageOracleTransactor) Cheat(opts *bind.TransactOpts, partOffset *big.Int, key [32]byte, part [32]byte, size *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.contract.Transact(opts, "cheat", partOffset, key, part, size)
}

// Cheat is a paid mutator transaction binding the contract method 0xfe4ac08e.
//
// Solidity: function cheat(uint256 partOffset, bytes32 key, bytes32 part, uint256 size) returns()
func (_PreimageOracle *PreimageOracleSession) Cheat(partOffset *big.Int, key [32]byte, part [32]byte, size *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.Contract.Cheat(&_PreimageOracle.TransactOpts, partOffset, key, part, size)
}

// Cheat is a paid mutator transaction binding the contract method 0xfe4ac08e.
//
// Solidity: function cheat(uint256 partOffset, bytes32 key, bytes32 part, uint256 size) returns()
func (_PreimageOracle *PreimageOracleTransactorSession) Cheat(partOffset *big.Int, key [32]byte, part [32]byte, size *big.Int) (*types.Transaction, error) {
	return _PreimageOracle.Contract.Cheat(&_PreimageOracle.TransactOpts, partOffset, key, part, size)
}

// LoadKeccak256PreimagePart is a paid mutator transaction binding the contract method 0xe1592611.
//
// Solidity: function loadKeccak256PreimagePart(uint256 _partOffset, bytes _preimage) returns()
func (_PreimageOracle *PreimageOracleTransactor) LoadKeccak256PreimagePart(opts *bind.TransactOpts, _partOffset *big.Int, _preimage []byte) (*types.Transaction, error) {
	return _PreimageOracle.contract.Transact(opts, "loadKeccak256PreimagePart", _partOffset, _preimage)
}

// LoadKeccak256PreimagePart is a paid mutator transaction binding the contract method 0xe1592611.
//
// Solidity: function loadKeccak256PreimagePart(uint256 _partOffset, bytes _preimage) returns()
func (_PreimageOracle *PreimageOracleSession) LoadKeccak256PreimagePart(_partOffset *big.Int, _preimage []byte) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadKeccak256PreimagePart(&_PreimageOracle.TransactOpts, _partOffset, _preimage)
}

// LoadKeccak256PreimagePart is a paid mutator transaction binding the contract method 0xe1592611.
//
// Solidity: function loadKeccak256PreimagePart(uint256 _partOffset, bytes _preimage) returns()
func (_PreimageOracle *PreimageOracleTransactorSession) LoadKeccak256PreimagePart(_partOffset *big.Int, _preimage []byte) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadKeccak256PreimagePart(&_PreimageOracle.TransactOpts, _partOffset, _preimage)
}

// LoadLocalData is a paid mutator transaction binding the contract method 0x3eba58ce.
//
// Solidity: function loadLocalData(bytes _bootInfo) returns()
func (_PreimageOracle *PreimageOracleTransactor) LoadLocalData(opts *bind.TransactOpts, _bootInfo []byte) (*types.Transaction, error) {
	return _PreimageOracle.contract.Transact(opts, "loadLocalData", _bootInfo)
}

// LoadLocalData is a paid mutator transaction binding the contract method 0x3eba58ce.
//
// Solidity: function loadLocalData(bytes _bootInfo) returns()
func (_PreimageOracle *PreimageOracleSession) LoadLocalData(_bootInfo []byte) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadLocalData(&_PreimageOracle.TransactOpts, _bootInfo)
}

// LoadLocalData is a paid mutator transaction binding the contract method 0x3eba58ce.
//
// Solidity: function loadLocalData(bytes _bootInfo) returns()
func (_PreimageOracle *PreimageOracleTransactorSession) LoadLocalData(_bootInfo []byte) (*types.Transaction, error) {
	return _PreimageOracle.Contract.LoadLocalData(&_PreimageOracle.TransactOpts, _bootInfo)
}
