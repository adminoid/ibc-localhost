package keeper

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"bitbucket.org/decimalteam/go-smart-node/app"
	cmdcfg "bitbucket.org/decimalteam/go-smart-node/cmd/config"
	"bitbucket.org/decimalteam/go-smart-node/x/nft/keeper"
	"bitbucket.org/decimalteam/go-smart-node/x/nft/types"
)

func GetBaseAppWithCustomKeeper(t *testing.T) (*app.DSC, sdk.Context) {
	dsc := app.Setup(t, false, feemarkettypes.DefaultGenesisState())
	ctx := dsc.BaseApp.NewContext(false, tmproto.Header{})

	appCodec := dsc.AppCodec()

	dsc.NFTKeeper = *keeper.NewKeeper(
		appCodec,
		dsc.GetKey(types.StoreKey),
		dsc.BankKeeper,
		cmdcfg.BaseDenom,
	)

	return dsc, ctx
}
