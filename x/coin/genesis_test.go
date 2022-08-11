package coin_test

import (
	"bitbucket.org/decimalteam/go-smart-node/app"
	"bitbucket.org/decimalteam/go-smart-node/utils/helpers"
	"bitbucket.org/decimalteam/go-smart-node/x/coin"
	"bitbucket.org/decimalteam/go-smart-node/x/coin/testcoin"
	"bitbucket.org/decimalteam/go-smart-node/x/coin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	//tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func bootstrapGenesisTest() (*app.DSC, sdk.Context) {
	_, dsc, ctx := getBaseAppWithCustomKeeper()

	return dsc, ctx
}

var (
	atomCoin = types.Coin{
		Title:       "Cosmos Hub Atom",
		Symbol:      "ATOM",
		CRR:         50,
		Reserve:     sdk.NewInt(1_000_000_000),
		Volume:      helpers.EtherToWei(sdk.NewInt(1000000000000)),
		LimitVolume: sdk.NewInt(1_000_000_000_000_000),
		Creator:     "uatom",
		Identity:    "dx1hs2wdrm87c92rzhq0vgmgrxr6u57xpr2lcygc2",
	}
	tstCoin = types.Coin{
		Title:       "Test Suite Token",
		Symbol:      "TST",
		CRR:         100,
		Reserve:     sdk.NewInt(1_000_000_000),
		Volume:      sdk.NewInt(1_000_000_000_0),
		LimitVolume: sdk.NewInt(1_000_000_000_000_000_000),
		Creator:     "uatom",
		Identity:    "dx1hs2wdrm87c92rzhq0vgmgrxr6u57xpr2lcygc2",
	}

	check1 types.Check
	check2 types.Check
)

func TestAppModuleBasic_InitGenesis(t *testing.T) {
	app, ctx := bootstrapGenesisTest()

	// write genesis
	params := app.CoinKeeper.GetParams(ctx)

	coins := []types.Coin{
		atomCoin,
		tstCoin,
	}
	check1 = testcoin.CreateNewCheck(ctx.ChainID(), "100000del", "9", "", 1)
	check2 = testcoin.CreateNewCheck(ctx.ChainID(), "100000del", "10", "", 1)
	checks := []types.Check{
		check1,
		check2,
	}

	genesisState := types.NewGenesisState(params, coins, checks)
	coin.InitGenesis(ctx, app.CoinKeeper, genesisState)

	// export genesis

	coins = append(coins, types.Coin{
		Title:       params.BaseTitle,
		Symbol:      params.BaseSymbol,
		Volume:      params.BaseInitialVolume,
		CRR:         0,
		Reserve:     sdk.NewInt(0),
		Creator:     "",
		LimitVolume: sdk.NewInt(0),
		Identity:    "",
	})

	exportedGenesis := coin.ExportGenesis(ctx, app.CoinKeeper)
	require.Equal(t, params, exportedGenesis.Params)
	require.True(t, testcoin.CoinsEqual(coins, exportedGenesis.Coins))
	require.True(t, testcoin.ChecksEqual(checks, exportedGenesis.Checks))
}

func TestAppModuleBasic_ValidateGenesis(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*types.GenesisState)
		wantErr bool
	}{
		{"default", func(*types.GenesisState) {}, false},
		// validate genesis validators
		{"params coin title > 64 symbols", func(data *types.GenesisState) {
			data.Params.BaseTitle = "vsafa;jkdfndsj;anf;asdnf;dsjfkldasfkmsdkalmf;alkdsmflmasl;dkmf;lds"
		}, true},
		{"params symbol is regexp", func(data *types.GenesisState) {
			data.Params.BaseSymbol = "laSK;DM"
		}, true},
		{"params init volume is < min", func(data *types.GenesisState) {
			data.Params.BaseInitialVolume = sdk.NewInt(0)
		}, true},
		{"params init volume is > max", func(data *types.GenesisState) {
			data.Params.BaseInitialVolume = helpers.EtherToWei(sdk.NewInt(1000000000000002))
		}, true},
		{"valid coins is repeated", func(data *types.GenesisState) {
			data.Coins = append(data.Coins, data.Coins[0])
		}, true},
		{"valid checks is repeated", func(data *types.GenesisState) {
			data.Checks = append(data.Checks, data.Checks[0])
		}, true},
		//{"invalid coin", func(data *types.GenesisState) {
		//	data.Coins = append(data.Coins, types.Coin{
		//		Title:       "vsafa;jkdfndsj;anf;asdnf;dsjfkldasfkmsdkalmf;alkdsmflmasl;dkmf;lds",
		//		Symbol:      "sdjfn",
		//		Volume:      helpers.EtherToWei(sdk.NewInt(1000000000000002)),
		//		CRR:         0,
		//		Reserve:     sdk.NewInt(0),
		//		Creator:     "",
		//		LimitVolume: sdk.NewInt(0),
		//		Identity:    "",
		//	})
		//}, true},
		//{"invalid check", func(data *types.GenesisState) {
		//	data.Checks = append(data.Checks, types.Check{})
		//}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			genesisState := types.DefaultGenesisState()
			genesisState.Coins = append(genesisState.Coins, atomCoin)
			genesisState.Checks = append(genesisState.Checks, check1)

			tt.mutate(genesisState)
			if tt.wantErr {
				require.Error(t, genesisState.Validate())
			} else {
				require.NoError(t, genesisState.Validate())
			}
		})
	}
}
