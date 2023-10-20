package multisig

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "cosmossdk.io/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	multisigsim "bitbucket.org/decimalteam/go-smart-node/x/multisig/simulation"
	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
)

var (
	_ = multisigsim.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accounts := make([]string, len(simState.Accounts))
	for i, account := range simState.Accounts {
		accounts[i] = account.Address.String()
	}
	genesis := types.DefaultGenesisState()
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	//
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	return operations
}
