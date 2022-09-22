package keeper_test

import (
	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

var (
	defaultDenom = "del"
	defaultCoins = sdk.NewCoins(sdk.NewCoin(defaultDenom, sdk.NewInt(1)))
)

func (s *KeeperTestSuite) TestMsgCreateWallet() {
	ctx, _, msgServer := s.ctx, s.msKeeper, s.msgServer
	require := s.Require()

	testCases := []struct {
		name   string
		input  *types.MsgCreateWallet
		expErr bool
		ctx    sdk.Context
	}{
		{
			"valid request",
			types.NewMsgCreateWallet(user4, defaultOwners, defaultWeights, defaultThreeshold),
			false,
			ctx,
		},
		{
			"wallet exists",
			types.NewMsgCreateWallet(user4, defaultOwners, defaultWeights, defaultThreeshold),
			true,
			ctx,
		},
		{
			"account address exists",
			types.NewMsgCreateWallet(user4, defaultOwners, defaultWeights, defaultThreeshold),
			true,
			ctx.WithTxBytes([]byte{1}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := msgServer.CreateWallet(ctx, tc.input)
			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgCreateTransaction() {
	ctx, k, msgServer := s.ctx, s.msKeeper, s.msgServer
	require := s.Require()

	k.SetWallet(ctx, *defaultWallet)
	createTestCases := []struct {
		do     func()
		name   string
		input  *types.MsgCreateTransaction
		expErr bool
	}{
		{
			func() {},
			"valid request -- first confirm",
			types.NewMsgCreateTransaction(user3, defaultWalletAddress.String(), user1.String(), defaultCoins),
			false,
		},
		{
			func() {},
			"wallet not exists",
			types.NewMsgCreateTransaction(user3, existsWalletAddress.String(), user4.String(), defaultCoins),
			true,
		},
		{
			func() { k.SetWallet(ctx, *existsWallet) },
			"insufuccient funds",
			types.NewMsgCreateTransaction(user3, existsWalletAddress.String(), user4.String(), defaultCoins),
			true,
		},
	}

	createResults := make([]string, 0)
	for _, tc := range createTestCases {
		tc := tc
		tc.do()
		s.T().Run(tc.name, func(t *testing.T) {
			res, err := msgServer.CreateTransaction(ctx, tc.input)
			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
				createResults = append(createResults, res.ID)
			}
		})
	}

	signTestCases := []struct {
		name   string
		input  *types.MsgSignTransaction
		expErr bool
	}{
		{
			"valid request -- second confirm",
			types.NewMsgSignTransaction(user2, createResults[0]),
			false,
		},
		{
			"third confirm -- enough conformations",
			types.NewMsgSignTransaction(user1, createResults[0]),
			true,
		},
		{
			"not exists tx with current ID",
			types.NewMsgSignTransaction(user2, "2"),
			true,
		},
	}

	for _, tc := range signTestCases {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := msgServer.SignTransaction(ctx, tc.input)
			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
			}
		})
	}
}
