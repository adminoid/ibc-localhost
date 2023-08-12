package drc20cosmos

import (
	"bytes"
	"embed"
	"encoding/hex"
	"encoding/json"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ethereum/go-ethereum/core/vm"
	"math/big"

	"bitbucket.org/decimalteam/go-smart-node/x/coin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethante "github.com/evmos/ethermint/app/ante"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	evm "github.com/evmos/ethermint/x/evm/vm"
)

const (
	addressForContractOwner = "0x2941b073ad6b59b1de4fc70c69e39a9e130b25ce"
	firstReward             = 50

	firstIncrease = 5
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type Drc20Cosmos struct {
	abi.ABI
	ctx        sdk.Context
	evmKeeper  ethante.EVMKeeper
	bankKeeper bankkeeper.Keeper
	stateDB    *statedb.StateDB
	evm        evm.EVM
	msg        *evmtypes.MsgEthereumTx
	cfg        *evmtypes.EVMConfig
	coin       types.Coin
}

// NewDrc20Cosmos create instance of contract
func NewDrc20Cosmos(ctx sdk.Context,
	evmKeeper ethante.EVMKeeper,
	bankKeeper bankkeeper.Keeper,
	msgEthTx *evmtypes.MsgEthereumTx,
	coinAction types.Coin,
) (*Drc20Cosmos, error) {
	abiBz, err := f.ReadFile("abi.json")
	if err != nil {
		return nil, err
	}

	newAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		return nil, err
	}

	params := evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig(evmKeeper.ChainID())
	baseFee := evmKeeper.GetBaseFee(ctx, ethCfg)

	signer := ethtypes.MakeSigner(ethCfg, big.NewInt(ctx.BlockHeight()))

	coreMsg, err := msgEthTx.AsMessage(signer, baseFee)
	if err != nil {
		return nil, err
	}

	cfg := &evmtypes.EVMConfig{
		ChainConfig: ethCfg,
		Params:      params,
		CoinBase:    common.Address{},
		BaseFee:     baseFee,
	}

	stateNewDB := statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash().Bytes())))
	evmNew := evmKeeper.NewEVM(ctx, coreMsg, cfg, evmtypes.NewNoOpTracer(), stateNewDB)

	// nonce := stateDB.GetNonce(common.HexToAddress("0x2941b073ad6b59b1de4fc70c69e39a9e130b25ce"))

	// stateDB.SetNonce(common.HexToAddress("0x2941b073ad6b59b1de4fc70c69e39a9e130b25ce"), nonce)
	// ret, _, leftoverGas, vmErr = evm.Create(sender, msg.Data(), leftoverGas, msg.Value())
	// stateDB.SetNonce(sender.Address(), msg.Nonce()+1)

	return &Drc20Cosmos{
		ABI:        newAbi,
		ctx:        ctx,
		evmKeeper:  evmKeeper,
		bankKeeper: bankKeeper,
		stateDB:    stateNewDB,
		evm:        evmNew,
		msg:        msgEthTx,
		cfg:        cfg,
		coin:       coinAction,
	}, nil
}

// CreateContractIfNotSet creation contract if not not to coin
func (drc Drc20Cosmos) CreateContractIfNotSet() (bool, error) {

	sender := vm.AccountRef(common.HexToAddress(addressForContractOwner))

	if drc.coin.Drc20Address == "" {
		drc.ctx.Logger().Info(drc.coin.Title)
	}
	drc.ctx.Logger().Info(drc.coin.Drc20Address)
	drc.ctx.Logger().Info(drc.coin.Denom)

	msgTx := drc.msg.AsTransaction()

	// access list preparation is moved from ante handler to here, because it's needed when `ApplyMessage` is called
	// under contexts where ante handlers are not run, for example `eth_call` and `eth_estimateGas`.
	if rules := drc.cfg.ChainConfig.Rules(big.NewInt(drc.ctx.BlockHeight()), drc.cfg.ChainConfig.MergeNetsplitBlock != nil); rules.IsBerlin {
		drc.stateDB.PrepareAccessList(common.HexToAddress(addressForContractOwner), nil, drc.evm.ActivePrecompiles(rules), msgTx.AccessList())
	}

	// receive nonce for owner address for new contract
	nonce := drc.stateDB.GetNonce(common.HexToAddress(addressForContractOwner))
	drc.stateDB.SetNonce(common.HexToAddress(addressForContractOwner), nonce)

	contractCode := "60806040523480156200001157600080fd5b506040516200134c3803806200134c83398101604081905262000034916200024f565b600662000042848262000351565b50600562000051838262000351565b506004805460ff1916601217905560038190553360008181526001602052604090208290556200008190620000c5565b60035460405190815233906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050506200041d565b6001600160a01b0381166200012f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840160405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001b257600080fd5b81516001600160401b0380821115620001cf57620001cf6200018a565b604051601f8301601f19908116603f01168101908282118183101715620001fa57620001fa6200018a565b816040528381526020925086838588010111156200021757600080fd5b600091505b838210156200023b57858201830151818301840152908201906200021c565b600093810190920192909252949350505050565b6000806000606084860312156200026557600080fd5b83516001600160401b03808211156200027d57600080fd5b6200028b87838801620001a0565b94506020860151915080821115620002a257600080fd5b50620002b186828701620001a0565b925050604084015190509250925092565b600181811c90821680620002d757607f821691505b602082108103620002f857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200034c57600081815260208120601f850160051c81016020861015620003275750805b601f850160051c820191505b81811015620003485782815560010162000333565b5050505b505050565b81516001600160401b038111156200036d576200036d6200018a565b62000385816200037e8454620002c2565b84620002fe565b602080601f831160018114620003bd5760008415620003a45750858301515b600019600386901b1c1916600185901b17855562000348565b600085815260208120601f198616915b82811015620003ee57888601518255948401946001909101908401620003cd565b50858210156200040d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610f1f806200042d6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063893d20e8116100ad578063a9059cbb11610071578063a9059cbb1461027a578063b09f12661461028d578063d28d885214610295578063dd62ed3e1461029d578063f2fde38b146102d657600080fd5b8063893d20e8146102165780638da5cb5b1461023b57806395d89b411461024c5780639dc29fac14610254578063a457c2d71461026757600080fd5b806332424aa3116100f457806332424aa3146101b057806339509351146101bd57806340c10f19146101d057806370a08231146101e3578063715018a61461020c57600080fd5b806306fdde0314610131578063095ea7b31461014f57806318160ddd1461017257806323b872dd14610184578063313ce56714610197575b600080fd5b6101396102e9565b6040516101469190610c8b565b60405180910390f35b61016261015d366004610cf5565b61037b565b6040519015158152602001610146565b6003545b604051908152602001610146565b610162610192366004610d1f565b610392565b60045460ff165b60405160ff9091168152602001610146565b60045461019e9060ff1681565b6101626101cb366004610cf5565b6103fb565b6101626101de366004610cf5565b610431565b6101766101f1366004610d5b565b6001600160a01b031660009081526001602052604090205490565b61021461046f565b005b6000546001600160a01b03165b6040516001600160a01b039091168152602001610146565b6000546001600160a01b0316610223565b6101396104e3565b610162610262366004610cf5565b6104f2565b610162610275366004610cf5565b610527565b610162610288366004610cf5565b610576565b610139610583565b610139610611565b6101766102ab366004610d76565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6102146102e4366004610d5b565b61061e565b6060600680546102f890610da9565b80601f016020809104026020016040519081016040528092919081815260200182805461032490610da9565b80156103715780601f1061034657610100808354040283529160200191610371565b820191906000526020600020905b81548152906001019060200180831161035457829003601f168201915b5050505050905090565b6000610388338484610654565b5060015b92915050565b600061039f848484610779565b6103f184336103ec85604051806060016040528060288152602001610e55602891396001600160a01b038a16600090815260026020908152604080832033845290915290205491906108ff565b610654565b5060019392505050565b3360008181526002602090815260408083206001600160a01b038716845290915281205490916103889185906103ec9086610939565b600080546001600160a01b031633146104655760405162461bcd60e51b815260040161045c90610de3565b60405180910390fd5b610388838361099f565b6000546001600160a01b031633146104995760405162461bcd60e51b815260040161045c90610de3565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6060600580546102f890610da9565b600080546001600160a01b0316331461051d5760405162461bcd60e51b815260040161045c90610de3565b6103888383610a85565b600061038833846103ec85604051806060016040528060258152602001610ea3602591393360009081526002602090815260408083206001600160a01b038d16845290915290205491906108ff565b6000610388338484610779565b6005805461059090610da9565b80601f01602080910402602001604051908101604052809291908181526020018280546105bc90610da9565b80156106095780601f106105de57610100808354040283529160200191610609565b820191906000526020600020905b8154815290600101906020018083116105ec57829003601f168201915b505050505081565b6006805461059090610da9565b6000546001600160a01b031633146106485760405162461bcd60e51b815260040161045c90610de3565b61065181610b89565b50565b6001600160a01b0383166106b65760405162461bcd60e51b8152602060048201526024808201527f42455032303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161045c565b6001600160a01b0382166107175760405162461bcd60e51b815260206004820152602260248201527f42455032303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161045c565b6001600160a01b0383811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0383166107dd5760405162461bcd60e51b815260206004820152602560248201527f42455032303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161045c565b6001600160a01b03821661083f5760405162461bcd60e51b815260206004820152602360248201527f42455032303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161045c565b61087c81604051806060016040528060268152602001610e7d602691396001600160a01b03861660009081526001602052604090205491906108ff565b6001600160a01b0380851660009081526001602052604080822093909355908416815220546108ab9082610939565b6001600160a01b0380841660008181526001602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061076c9085815260200190565b600081848411156109235760405162461bcd60e51b815260040161045c9190610c8b565b5060006109308486610e2e565b95945050505050565b6000806109468385610e41565b9050838110156109985760405162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015260640161045c565b9392505050565b6001600160a01b0382166109f55760405162461bcd60e51b815260206004820152601f60248201527f42455032303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161045c565b600354610a029082610939565b6003556001600160a01b038216600090815260016020526040902054610a289082610939565b6001600160a01b0383166000818152600160205260408082209390935591519091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610a799085815260200190565b60405180910390a35050565b6001600160a01b038216610ae55760405162461bcd60e51b815260206004820152602160248201527f42455032303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161045c565b610b2281604051806060016040528060228152602001610ec8602291396001600160a01b03851660009081526001602052604090205491906108ff565b6001600160a01b038316600090815260016020526040902055600354610b489082610c49565b6003556040518181526000906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610a79565b6001600160a01b038116610bee5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161045c565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600061099883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506108ff565b600060208083528351808285015260005b81811015610cb857858101830151858201604001528201610c9c565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610cf057600080fd5b919050565b60008060408385031215610d0857600080fd5b610d1183610cd9565b946020939093013593505050565b600080600060608486031215610d3457600080fd5b610d3d84610cd9565b9250610d4b60208501610cd9565b9150604084013590509250925092565b600060208284031215610d6d57600080fd5b61099882610cd9565b60008060408385031215610d8957600080fd5b610d9283610cd9565b9150610da060208401610cd9565b90509250929050565b600181811c90821680610dbd57607f821691505b602082108103610ddd57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b8181038181111561038c5761038c610e18565b8082018082111561038c5761038c610e1856fe42455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636542455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636542455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f42455032303a206275726e20616d6f756e7420657863656564732062616c616e6365a264697066735822122049eb5f485e208960ea8c98da009bc8d779b38c32ba4536fc245494d8a5efbf4764736f6c63430008110033"

	inputContract, err := hex.DecodeString(contractCode)
	if err != nil {
		return false, err
	}
	ret, contractAddr, _, vmErr := drc.evm.Create(sender, inputContract, 1000000, big.NewInt(100))
	drc.stateDB.SetNonce(sender.Address(), nonce+1)

	drc.ctx.Logger().Info("Result create contract %T", contractAddr.Hex())
	drc.ctx.Logger().With(ret).Info("Result create contract")

	if vmErr != nil {
		drc.ctx.Logger().Info(vmErr.Error())
		//return false, sdkerrors.ErrUnknownRequest.Wrapf("failed to encode log vmErr %T", vmErr)
	}

	txLogAttrs := make([]sdk.Attribute, len(evmtypes.NewLogsFromEth(drc.stateDB.Logs())))
	for i, log := range drc.stateDB.Logs() {
		value, err := json.Marshal(log)
		if err != nil {
			drc.ctx.Logger().Info(drc.coin.Denom)
			return false, sdkerrors.ErrUnknownRequest.Wrapf("failed to encode log %T", err)
		}
		txLogAttrs[i] = sdk.NewAttribute(evmtypes.AttributeKeyTxLog, string(value))
	}

	// emit events
	drc.ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			evmtypes.EventTypeTxLog,
			txLogAttrs...,
		),
	})

	// The dirty states in `StateDB` is either committed or discarded after return
	if err := drc.stateDB.Commit(); err != nil {
		return false, sdkerrors.ErrUnknownRequest.Wrapf("failed to encode log %T", err)
	}

	return true, nil
}
