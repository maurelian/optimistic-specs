// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sro

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

// StateRootOracleMetaData contains all meta data concerning the StateRootOracle contract.
var StateRootOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionFrequency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_historicalTotalBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_stateRoot\",\"type\":\"bytes32\"}],\"name\":\"StateRootAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_stateRoot\",\"type\":\"bytes32\"}],\"name\":\"StateRootDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"appendStateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"computeL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"deleteStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSubmittedStateRoots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalTotalBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"insideFraudProofWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_inside\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionFrequency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyStateCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_verified\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161094a38038061094a8339818101604052608081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050600084116100bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806109256025913960400191505060405180910390fd5b60008311610131576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6c32426c6f636b54696d65206d75737420626520706f7369746976650000000081525060200191505060405180910390fd5b8360008190555082600181905550816002600042815260200190815260200160002081905550426003819055504260048190555080600581905550505050506107a68061017f6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806397dc6b281161008c578063c90ec2da11610066578063c90ec2da14610222578063d8b91d7714610240578063f163481314610284578063f4cac30d146102b2576100cf565b806397dc6b28146101c6578063b210dc21146101e6578063c5095d6814610204576100cf565b806302e51345146100d45780630c1952d3146101165780632b75a87c14610134578063357e951f146101525780635e5ec9ae1461017057806393991af3146101a8575b600080fd5b610100600480360360208110156100ea57600080fd5b81019080803590602001909291905050506102f4565b6040518082815260200191505060405180910390f35b61011e610389565b6040518082815260200191505060405180910390f35b61013c61038f565b6040518082815260200191505060405180910390f35b61015a6103a7565b6040518082815260200191505060405180910390f35b6101a66004803603604081101561018657600080fd5b8101908080359060200190929190803590602001909291905050506103b4565b005b6101b061053e565b6040518082815260200191505060405180910390f35b6101ce610544565b60405180821515815260200191505060405180910390f35b6101ee610549565b6040518082815260200191505060405180910390f35b61020c61054f565b6040518082815260200191505060405180910390f35b61022a610555565b6040518082815260200191505060405180910390f35b61026c6004803603602081101561025657600080fd5b810190808035906020019092919050505061055b565b60405180821515815260200191505060405180910390f35b6102b06004803603602081101561029a57600080fd5b81019080803590602001909291905050506105f2565b005b6102de600480360360208110156102c857600080fd5b81019080803590602001909291905050506106f5565b6040518082815260200191505060405180910390f35b600060045482101561036e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f74696d657374616d70206265666f72652073746172740000000000000000000081525060200191505060405180910390fd5b60015460045483038161037d57fe5b04600554019050919050565b60035481565b6000805460045460035403816103a157fe5b04905090565b6000805460035401905090565b80421161040c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260388152602001806107396038913960400191505060405180910390fd5b6000801b821415610485576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f43616e6e6f74207375626d697420656d70747920737461746520726f6f74000081525060200191505060405180910390fd5b6000546003540181146104e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b81526020018061070e602b913960400191505060405180910390fd5b81600260008381526020019081526020016000208190555080600381905550807fd8cd688ff2c3eeace7aacdfeebdc550248887b0af3bc4d79a984b3fe791b4219836040518082815260200191505060405180910390a25050565b60015481565b600090565b60005481565b60045481565b60055481565b60008060001b600260008481526020019081526020016000205414156105e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f537461746520726f6f74206e6f74207375626d6974746564207965742e00000081525060200191505060405180910390fd5b60009050919050565b60008054820190506000801b600260008381526020019081526020016000205414610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4d7573742064656c6574652074697020737461746520726f6f742e000000000081525060200191505060405180910390fd5b6000600260008481526020019081526020016000205490506000801b6002600085815260200190815260200160002081905550827f37cbca86a576a39ea213b9486a41b0291e628d7f9ab47b8694301c2ae3e55d49826040518082815260200191505060405180910390a2505050565b6002602052806000526040600020600091509050548156fe4d757374207375626d697420737461746520726f6f7420666f72206576657279203235206d696e7574657343616e6e6f7420617070656e6420737461746520726f6f74732066726f6d20746865206675747572652e20436f6d65206f6e20627275682ea26469706673582212205564e7661565ce6ec85ed75e88ded5d4133da5dca1767d9459c7a3a252411ac564736f6c634300070600337375626d697373696f6e206672657175656e6379206d75737420626520706f736974697665",
}

// StateRootOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use StateRootOracleMetaData.ABI instead.
var StateRootOracleABI = StateRootOracleMetaData.ABI

// StateRootOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateRootOracleMetaData.Bin instead.
var StateRootOracleBin = StateRootOracleMetaData.Bin

// DeployStateRootOracle deploys a new Ethereum contract, binding an instance of StateRootOracle to it.
func DeployStateRootOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _submissionFrequency *big.Int, _l2BlockTime *big.Int, _genesisRoot [32]byte, _historicalTotalBlocks *big.Int) (common.Address, *types.Transaction, *StateRootOracle, error) {
	parsed, err := StateRootOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateRootOracleBin), backend, _submissionFrequency, _l2BlockTime, _genesisRoot, _historicalTotalBlocks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateRootOracle{StateRootOracleCaller: StateRootOracleCaller{contract: contract}, StateRootOracleTransactor: StateRootOracleTransactor{contract: contract}, StateRootOracleFilterer: StateRootOracleFilterer{contract: contract}}, nil
}

// StateRootOracle is an auto generated Go binding around an Ethereum contract.
type StateRootOracle struct {
	StateRootOracleCaller     // Read-only binding to the contract
	StateRootOracleTransactor // Write-only binding to the contract
	StateRootOracleFilterer   // Log filterer for contract events
}

// StateRootOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateRootOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateRootOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateRootOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateRootOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateRootOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateRootOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateRootOracleSession struct {
	Contract     *StateRootOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRootOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateRootOracleCallerSession struct {
	Contract *StateRootOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StateRootOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateRootOracleTransactorSession struct {
	Contract     *StateRootOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StateRootOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateRootOracleRaw struct {
	Contract *StateRootOracle // Generic contract binding to access the raw methods on
}

// StateRootOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateRootOracleCallerRaw struct {
	Contract *StateRootOracleCaller // Generic read-only contract binding to access the raw methods on
}

// StateRootOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateRootOracleTransactorRaw struct {
	Contract *StateRootOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateRootOracle creates a new instance of StateRootOracle, bound to a specific deployed contract.
func NewStateRootOracle(address common.Address, backend bind.ContractBackend) (*StateRootOracle, error) {
	contract, err := bindStateRootOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateRootOracle{StateRootOracleCaller: StateRootOracleCaller{contract: contract}, StateRootOracleTransactor: StateRootOracleTransactor{contract: contract}, StateRootOracleFilterer: StateRootOracleFilterer{contract: contract}}, nil
}

// NewStateRootOracleCaller creates a new read-only instance of StateRootOracle, bound to a specific deployed contract.
func NewStateRootOracleCaller(address common.Address, caller bind.ContractCaller) (*StateRootOracleCaller, error) {
	contract, err := bindStateRootOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateRootOracleCaller{contract: contract}, nil
}

// NewStateRootOracleTransactor creates a new write-only instance of StateRootOracle, bound to a specific deployed contract.
func NewStateRootOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*StateRootOracleTransactor, error) {
	contract, err := bindStateRootOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateRootOracleTransactor{contract: contract}, nil
}

// NewStateRootOracleFilterer creates a new log filterer instance of StateRootOracle, bound to a specific deployed contract.
func NewStateRootOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*StateRootOracleFilterer, error) {
	contract, err := bindStateRootOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateRootOracleFilterer{contract: contract}, nil
}

// bindStateRootOracle binds a generic wrapper to an already deployed contract.
func bindStateRootOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateRootOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateRootOracle *StateRootOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateRootOracle.Contract.StateRootOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateRootOracle *StateRootOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateRootOracle.Contract.StateRootOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateRootOracle *StateRootOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateRootOracle.Contract.StateRootOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateRootOracle *StateRootOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateRootOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateRootOracle *StateRootOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateRootOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateRootOracle *StateRootOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateRootOracle.Contract.contract.Transact(opts, method, params...)
}

// ComputeL2BlockNumber is a free data retrieval call binding the contract method 0x02e51345.
//
// Solidity: function computeL2BlockNumber(uint256 _timestamp) view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) ComputeL2BlockNumber(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "computeL2BlockNumber", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2BlockNumber is a free data retrieval call binding the contract method 0x02e51345.
//
// Solidity: function computeL2BlockNumber(uint256 _timestamp) view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) ComputeL2BlockNumber(_timestamp *big.Int) (*big.Int, error) {
	return _StateRootOracle.Contract.ComputeL2BlockNumber(&_StateRootOracle.CallOpts, _timestamp)
}

// ComputeL2BlockNumber is a free data retrieval call binding the contract method 0x02e51345.
//
// Solidity: function computeL2BlockNumber(uint256 _timestamp) view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) ComputeL2BlockNumber(_timestamp *big.Int) (*big.Int, error) {
	return _StateRootOracle.Contract.ComputeL2BlockNumber(&_StateRootOracle.CallOpts, _timestamp)
}

// GetTotalSubmittedStateRoots is a free data retrieval call binding the contract method 0x2b75a87c.
//
// Solidity: function getTotalSubmittedStateRoots() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) GetTotalSubmittedStateRoots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "getTotalSubmittedStateRoots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSubmittedStateRoots is a free data retrieval call binding the contract method 0x2b75a87c.
//
// Solidity: function getTotalSubmittedStateRoots() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) GetTotalSubmittedStateRoots() (*big.Int, error) {
	return _StateRootOracle.Contract.GetTotalSubmittedStateRoots(&_StateRootOracle.CallOpts)
}

// GetTotalSubmittedStateRoots is a free data retrieval call binding the contract method 0x2b75a87c.
//
// Solidity: function getTotalSubmittedStateRoots() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) GetTotalSubmittedStateRoots() (*big.Int, error) {
	return _StateRootOracle.Contract.GetTotalSubmittedStateRoots(&_StateRootOracle.CallOpts)
}

// HistoricalTotalBlocks is a free data retrieval call binding the contract method 0xc90ec2da.
//
// Solidity: function historicalTotalBlocks() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) HistoricalTotalBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "historicalTotalBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalTotalBlocks is a free data retrieval call binding the contract method 0xc90ec2da.
//
// Solidity: function historicalTotalBlocks() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) HistoricalTotalBlocks() (*big.Int, error) {
	return _StateRootOracle.Contract.HistoricalTotalBlocks(&_StateRootOracle.CallOpts)
}

// HistoricalTotalBlocks is a free data retrieval call binding the contract method 0xc90ec2da.
//
// Solidity: function historicalTotalBlocks() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) HistoricalTotalBlocks() (*big.Int, error) {
	return _StateRootOracle.Contract.HistoricalTotalBlocks(&_StateRootOracle.CallOpts)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0xd8b91d77.
//
// Solidity: function insideFraudProofWindow(uint256 _timestamp) view returns(bool _inside)
func (_StateRootOracle *StateRootOracleCaller) InsideFraudProofWindow(opts *bind.CallOpts, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "insideFraudProofWindow", _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0xd8b91d77.
//
// Solidity: function insideFraudProofWindow(uint256 _timestamp) view returns(bool _inside)
func (_StateRootOracle *StateRootOracleSession) InsideFraudProofWindow(_timestamp *big.Int) (bool, error) {
	return _StateRootOracle.Contract.InsideFraudProofWindow(&_StateRootOracle.CallOpts, _timestamp)
}

// InsideFraudProofWindow is a free data retrieval call binding the contract method 0xd8b91d77.
//
// Solidity: function insideFraudProofWindow(uint256 _timestamp) view returns(bool _inside)
func (_StateRootOracle *StateRootOracleCallerSession) InsideFraudProofWindow(_timestamp *big.Int) (bool, error) {
	return _StateRootOracle.Contract.InsideFraudProofWindow(&_StateRootOracle.CallOpts, _timestamp)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) L2BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "l2BlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) L2BlockTime() (*big.Int, error) {
	return _StateRootOracle.Contract.L2BlockTime(&_StateRootOracle.CallOpts)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) L2BlockTime() (*big.Int, error) {
	return _StateRootOracle.Contract.L2BlockTime(&_StateRootOracle.CallOpts)
}

// LatestBlockTimestamp is a free data retrieval call binding the contract method 0x0c1952d3.
//
// Solidity: function latestBlockTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) LatestBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "latestBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockTimestamp is a free data retrieval call binding the contract method 0x0c1952d3.
//
// Solidity: function latestBlockTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) LatestBlockTimestamp() (*big.Int, error) {
	return _StateRootOracle.Contract.LatestBlockTimestamp(&_StateRootOracle.CallOpts)
}

// LatestBlockTimestamp is a free data retrieval call binding the contract method 0x0c1952d3.
//
// Solidity: function latestBlockTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) LatestBlockTimestamp() (*big.Int, error) {
	return _StateRootOracle.Contract.LatestBlockTimestamp(&_StateRootOracle.CallOpts)
}

// NextTimestamp is a free data retrieval call binding the contract method 0x357e951f.
//
// Solidity: function nextTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) NextTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "nextTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTimestamp is a free data retrieval call binding the contract method 0x357e951f.
//
// Solidity: function nextTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) NextTimestamp() (*big.Int, error) {
	return _StateRootOracle.Contract.NextTimestamp(&_StateRootOracle.CallOpts)
}

// NextTimestamp is a free data retrieval call binding the contract method 0x357e951f.
//
// Solidity: function nextTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) NextTimestamp() (*big.Int, error) {
	return _StateRootOracle.Contract.NextTimestamp(&_StateRootOracle.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) StartingBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "startingBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) StartingBlockTimestamp() (*big.Int, error) {
	return _StateRootOracle.Contract.StartingBlockTimestamp(&_StateRootOracle.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) StartingBlockTimestamp() (*big.Int, error) {
	return _StateRootOracle.Contract.StartingBlockTimestamp(&_StateRootOracle.CallOpts)
}

// StateRoots is a free data retrieval call binding the contract method 0xf4cac30d.
//
// Solidity: function stateRoots(uint256 ) view returns(bytes32)
func (_StateRootOracle *StateRootOracleCaller) StateRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "stateRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRoots is a free data retrieval call binding the contract method 0xf4cac30d.
//
// Solidity: function stateRoots(uint256 ) view returns(bytes32)
func (_StateRootOracle *StateRootOracleSession) StateRoots(arg0 *big.Int) ([32]byte, error) {
	return _StateRootOracle.Contract.StateRoots(&_StateRootOracle.CallOpts, arg0)
}

// StateRoots is a free data retrieval call binding the contract method 0xf4cac30d.
//
// Solidity: function stateRoots(uint256 ) view returns(bytes32)
func (_StateRootOracle *StateRootOracleCallerSession) StateRoots(arg0 *big.Int) ([32]byte, error) {
	return _StateRootOracle.Contract.StateRoots(&_StateRootOracle.CallOpts, arg0)
}

// SubmissionFrequency is a free data retrieval call binding the contract method 0xb210dc21.
//
// Solidity: function submissionFrequency() view returns(uint256)
func (_StateRootOracle *StateRootOracleCaller) SubmissionFrequency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "submissionFrequency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionFrequency is a free data retrieval call binding the contract method 0xb210dc21.
//
// Solidity: function submissionFrequency() view returns(uint256)
func (_StateRootOracle *StateRootOracleSession) SubmissionFrequency() (*big.Int, error) {
	return _StateRootOracle.Contract.SubmissionFrequency(&_StateRootOracle.CallOpts)
}

// SubmissionFrequency is a free data retrieval call binding the contract method 0xb210dc21.
//
// Solidity: function submissionFrequency() view returns(uint256)
func (_StateRootOracle *StateRootOracleCallerSession) SubmissionFrequency() (*big.Int, error) {
	return _StateRootOracle.Contract.SubmissionFrequency(&_StateRootOracle.CallOpts)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0x97dc6b28.
//
// Solidity: function verifyStateCommitment() pure returns(bool _verified)
func (_StateRootOracle *StateRootOracleCaller) VerifyStateCommitment(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StateRootOracle.contract.Call(opts, &out, "verifyStateCommitment")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0x97dc6b28.
//
// Solidity: function verifyStateCommitment() pure returns(bool _verified)
func (_StateRootOracle *StateRootOracleSession) VerifyStateCommitment() (bool, error) {
	return _StateRootOracle.Contract.VerifyStateCommitment(&_StateRootOracle.CallOpts)
}

// VerifyStateCommitment is a free data retrieval call binding the contract method 0x97dc6b28.
//
// Solidity: function verifyStateCommitment() pure returns(bool _verified)
func (_StateRootOracle *StateRootOracleCallerSession) VerifyStateCommitment() (bool, error) {
	return _StateRootOracle.Contract.VerifyStateCommitment(&_StateRootOracle.CallOpts)
}

// AppendStateRoot is a paid mutator transaction binding the contract method 0x5e5ec9ae.
//
// Solidity: function appendStateRoot(bytes32 _stateRoot, uint256 _timestamp) returns()
func (_StateRootOracle *StateRootOracleTransactor) AppendStateRoot(opts *bind.TransactOpts, _stateRoot [32]byte, _timestamp *big.Int) (*types.Transaction, error) {
	return _StateRootOracle.contract.Transact(opts, "appendStateRoot", _stateRoot, _timestamp)
}

// AppendStateRoot is a paid mutator transaction binding the contract method 0x5e5ec9ae.
//
// Solidity: function appendStateRoot(bytes32 _stateRoot, uint256 _timestamp) returns()
func (_StateRootOracle *StateRootOracleSession) AppendStateRoot(_stateRoot [32]byte, _timestamp *big.Int) (*types.Transaction, error) {
	return _StateRootOracle.Contract.AppendStateRoot(&_StateRootOracle.TransactOpts, _stateRoot, _timestamp)
}

// AppendStateRoot is a paid mutator transaction binding the contract method 0x5e5ec9ae.
//
// Solidity: function appendStateRoot(bytes32 _stateRoot, uint256 _timestamp) returns()
func (_StateRootOracle *StateRootOracleTransactorSession) AppendStateRoot(_stateRoot [32]byte, _timestamp *big.Int) (*types.Transaction, error) {
	return _StateRootOracle.Contract.AppendStateRoot(&_StateRootOracle.TransactOpts, _stateRoot, _timestamp)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0xf1634813.
//
// Solidity: function deleteStateBatch(uint256 _timestamp) returns()
func (_StateRootOracle *StateRootOracleTransactor) DeleteStateBatch(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _StateRootOracle.contract.Transact(opts, "deleteStateBatch", _timestamp)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0xf1634813.
//
// Solidity: function deleteStateBatch(uint256 _timestamp) returns()
func (_StateRootOracle *StateRootOracleSession) DeleteStateBatch(_timestamp *big.Int) (*types.Transaction, error) {
	return _StateRootOracle.Contract.DeleteStateBatch(&_StateRootOracle.TransactOpts, _timestamp)
}

// DeleteStateBatch is a paid mutator transaction binding the contract method 0xf1634813.
//
// Solidity: function deleteStateBatch(uint256 _timestamp) returns()
func (_StateRootOracle *StateRootOracleTransactorSession) DeleteStateBatch(_timestamp *big.Int) (*types.Transaction, error) {
	return _StateRootOracle.Contract.DeleteStateBatch(&_StateRootOracle.TransactOpts, _timestamp)
}

// StateRootOracleStateRootAppendedIterator is returned from FilterStateRootAppended and is used to iterate over the raw logs and unpacked data for StateRootAppended events raised by the StateRootOracle contract.
type StateRootOracleStateRootAppendedIterator struct {
	Event *StateRootOracleStateRootAppended // Event containing the contract specifics and raw log

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
func (it *StateRootOracleStateRootAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateRootOracleStateRootAppended)
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
		it.Event = new(StateRootOracleStateRootAppended)
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
func (it *StateRootOracleStateRootAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateRootOracleStateRootAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateRootOracleStateRootAppended represents a StateRootAppended event raised by the StateRootOracle contract.
type StateRootOracleStateRootAppended struct {
	Timestamp *big.Int
	StateRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateRootAppended is a free log retrieval operation binding the contract event 0xd8cd688ff2c3eeace7aacdfeebdc550248887b0af3bc4d79a984b3fe791b4219.
//
// Solidity: event StateRootAppended(uint256 indexed _timestamp, bytes32 _stateRoot)
func (_StateRootOracle *StateRootOracleFilterer) FilterStateRootAppended(opts *bind.FilterOpts, _timestamp []*big.Int) (*StateRootOracleStateRootAppendedIterator, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _StateRootOracle.contract.FilterLogs(opts, "StateRootAppended", _timestampRule)
	if err != nil {
		return nil, err
	}
	return &StateRootOracleStateRootAppendedIterator{contract: _StateRootOracle.contract, event: "StateRootAppended", logs: logs, sub: sub}, nil
}

// WatchStateRootAppended is a free log subscription operation binding the contract event 0xd8cd688ff2c3eeace7aacdfeebdc550248887b0af3bc4d79a984b3fe791b4219.
//
// Solidity: event StateRootAppended(uint256 indexed _timestamp, bytes32 _stateRoot)
func (_StateRootOracle *StateRootOracleFilterer) WatchStateRootAppended(opts *bind.WatchOpts, sink chan<- *StateRootOracleStateRootAppended, _timestamp []*big.Int) (event.Subscription, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _StateRootOracle.contract.WatchLogs(opts, "StateRootAppended", _timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateRootOracleStateRootAppended)
				if err := _StateRootOracle.contract.UnpackLog(event, "StateRootAppended", log); err != nil {
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

// ParseStateRootAppended is a log parse operation binding the contract event 0xd8cd688ff2c3eeace7aacdfeebdc550248887b0af3bc4d79a984b3fe791b4219.
//
// Solidity: event StateRootAppended(uint256 indexed _timestamp, bytes32 _stateRoot)
func (_StateRootOracle *StateRootOracleFilterer) ParseStateRootAppended(log types.Log) (*StateRootOracleStateRootAppended, error) {
	event := new(StateRootOracleStateRootAppended)
	if err := _StateRootOracle.contract.UnpackLog(event, "StateRootAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateRootOracleStateRootDeletedIterator is returned from FilterStateRootDeleted and is used to iterate over the raw logs and unpacked data for StateRootDeleted events raised by the StateRootOracle contract.
type StateRootOracleStateRootDeletedIterator struct {
	Event *StateRootOracleStateRootDeleted // Event containing the contract specifics and raw log

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
func (it *StateRootOracleStateRootDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateRootOracleStateRootDeleted)
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
		it.Event = new(StateRootOracleStateRootDeleted)
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
func (it *StateRootOracleStateRootDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateRootOracleStateRootDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateRootOracleStateRootDeleted represents a StateRootDeleted event raised by the StateRootOracle contract.
type StateRootOracleStateRootDeleted struct {
	Timestamp *big.Int
	StateRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateRootDeleted is a free log retrieval operation binding the contract event 0x37cbca86a576a39ea213b9486a41b0291e628d7f9ab47b8694301c2ae3e55d49.
//
// Solidity: event StateRootDeleted(uint256 indexed _timestamp, bytes32 _stateRoot)
func (_StateRootOracle *StateRootOracleFilterer) FilterStateRootDeleted(opts *bind.FilterOpts, _timestamp []*big.Int) (*StateRootOracleStateRootDeletedIterator, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _StateRootOracle.contract.FilterLogs(opts, "StateRootDeleted", _timestampRule)
	if err != nil {
		return nil, err
	}
	return &StateRootOracleStateRootDeletedIterator{contract: _StateRootOracle.contract, event: "StateRootDeleted", logs: logs, sub: sub}, nil
}

// WatchStateRootDeleted is a free log subscription operation binding the contract event 0x37cbca86a576a39ea213b9486a41b0291e628d7f9ab47b8694301c2ae3e55d49.
//
// Solidity: event StateRootDeleted(uint256 indexed _timestamp, bytes32 _stateRoot)
func (_StateRootOracle *StateRootOracleFilterer) WatchStateRootDeleted(opts *bind.WatchOpts, sink chan<- *StateRootOracleStateRootDeleted, _timestamp []*big.Int) (event.Subscription, error) {

	var _timestampRule []interface{}
	for _, _timestampItem := range _timestamp {
		_timestampRule = append(_timestampRule, _timestampItem)
	}

	logs, sub, err := _StateRootOracle.contract.WatchLogs(opts, "StateRootDeleted", _timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateRootOracleStateRootDeleted)
				if err := _StateRootOracle.contract.UnpackLog(event, "StateRootDeleted", log); err != nil {
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

// ParseStateRootDeleted is a log parse operation binding the contract event 0x37cbca86a576a39ea213b9486a41b0291e628d7f9ab47b8694301c2ae3e55d49.
//
// Solidity: event StateRootDeleted(uint256 indexed _timestamp, bytes32 _stateRoot)
func (_StateRootOracle *StateRootOracleFilterer) ParseStateRootDeleted(log types.Log) (*StateRootOracleStateRootDeleted, error) {
	event := new(StateRootOracleStateRootDeleted)
	if err := _StateRootOracle.contract.UnpackLog(event, "StateRootDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
