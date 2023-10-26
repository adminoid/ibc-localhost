package testvalidator

import (
	"bitbucket.org/decimalteam/go-smart-node/x/validator/types"
	sdk "github.com/adminoid/cosmos-sdk/types"
	bankkeeper "github.com/adminoid/cosmos-sdk/x/bank/keeper"
)

func FundModuleAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, types.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, recipientMod, amounts)
}
