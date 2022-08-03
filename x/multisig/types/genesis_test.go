package types_test

import (
	"testing"

	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisStateValidate(t *testing.T) {
	var testCases = []struct {
		description  string
		genesisState *types.GenesisState
		valid        bool
	}{
		{
			"valid genesis",
			types.DefaultGenesisState(),
			true,
		},
		{
			"valid legacy owners",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners:  []string{"", ""},
						LegacyOwners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			true,
		},
		{
			"double wallet",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
		{
			"invalid wallet address",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String() + "1",
						Owners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
		{
			"double owner",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
		{
			"empty owner",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners: []string{
							"",
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
		{
			"double legacy owners",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners:  []string{"", ""},
						LegacyOwners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
		{
			"threshold over sum of weights",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 3,
					},
				},
			},
			false,
		},
		{
			"invalid owner/legacy combination (filled)",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners:  []string{sdk.AccAddress([]byte{1, 2, 3, 4}).String(), ""},
						LegacyOwners: []string{
							sdk.AccAddress([]byte{1, 2, 3, 1}).String(),
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
		{
			"invalid owner/legacy combination (empty)",
			&types.GenesisState{
				Wallets: []types.Wallet{
					{
						Address: sdk.AccAddress([]byte{1, 2, 3}).String(),
						Owners:  []string{"", ""},
						LegacyOwners: []string{
							"",
							sdk.AccAddress([]byte{1, 2, 3, 2}).String(),
						},
						Weights:   []uint64{1, 1},
						Threshold: 2,
					},
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.genesisState.Validate()
		if tc.valid {
			require.NoError(t, err, tc.description)
		} else {
			require.Error(t, err, tc.description)
		}
	}
}
