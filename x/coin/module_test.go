package coin_test

import (
	testkeeper "bitbucket.org/decimalteam/go-smart-node/testutil/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	_, app, ctx := testkeeper.GetTestAppWithCoinKeeper(t)

	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.BondedPoolName))
	require.NotNil(t, acc)

	acc = app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.NotBondedPoolName))
	require.NotNil(t, acc)
}
