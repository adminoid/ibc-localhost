package keeper_test

import (
	commonTypes "bitbucket.org/decimalteam/go-smart-node/types"
	"bitbucket.org/decimalteam/go-smart-node/x/legacy/types"
	multisigtypes "bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	nfttypes "bitbucket.org/decimalteam/go-smart-node/x/nft/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	"testing"
)

var (
	invalidAccPk         = secp256k1.GenPrivKey().PubKey()
	invalidOldAddress, _ = commonTypes.GetLegacyAddressFromPubKey(invalidAccPk.Bytes())
	invalidActualAddress = sdk.AccAddress(ethsecp256k1.PubKey{Key: invalidAccPk.Bytes()}.Address())
	defaultTokenID       = "test_record_token_ID"
	defaultRecord        = types.Record{
		LegacyAddress: oldAddress,
		Coins:         sdk.NewCoins(sdk.NewCoin("test1", sdk.NewInt(100)), sdk.NewCoin("test2", sdk.NewInt(102))),
		Wallets:       []string{defaultMultisigWalletBefore.Address},
		NFTs:          []string{defaultTokenID},
	}
	subTokenReserve        = sdk.NewCoin(baseDenom, sdk.NewInt(10))
	defaultSubTokensBefore = []nfttypes.SubToken{
		{
			ID:      1,
			Owner:   oldAddress,
			Reserve: &subTokenReserve,
		},
	}
	defaultSubTokensAfter = []nfttypes.SubToken{
		{
			ID:      1,
			Owner:   actualAddress,
			Reserve: &subTokenReserve,
		},
	}
	defaultMultisigWalletBefore = multisigtypes.Wallet{
		Address: "wallet",
		Owners:  []string{oldAddress, invalidActualAddress.String()},
	}
	defaultMultisigWalletAfter = multisigtypes.Wallet{
		Address: "wallet",
		Owners:  []string{actualAddress, invalidActualAddress.String()},
	}
)

func (s *KeeperTestSuite) TestMsgReturnLegacy() {
	ctx, k, msgServer := s.ctx, s.legacyKeeper, s.msgServer
	require := s.Require()

	k.SetLegacyRecord(ctx, defaultRecord)
	k.RestoreCache(ctx)

	testCases := []struct {
		name   string
		input  *types.MsgReturnLegacy
		expErr bool
	}{
		{
			"valid request",
			types.NewMsgReturnLegacy(newAddress, publicKey),
			false,
		},
		{
			"not have record",
			types.NewMsgReturnLegacy(invalidActualAddress, invalidAccPk.Bytes()),
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.name, func(t *testing.T) {
			_, err := msgServer.ReturnLegacy(ctx, tc.input)
			if tc.expErr {
				require.Error(err)
			} else {
				require.NoError(err)
			}
		})
	}
}
