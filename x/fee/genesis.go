package fee

import (
	"bitbucket.org/decimalteam/go-smart-node/x/fee/keeper"
	"bitbucket.org/decimalteam/go-smart-node/x/fee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets nft information for genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	k.SetModuleParams(ctx, gs.Params)
	err := k.SavePrice(ctx, gs.InitialPrice)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	price, err := k.GetPrice(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params:       k.GetModuleParams(ctx),
		InitialPrice: price,
	}
}
