package validator_test

import (
	"testing"
)

func TestValidateGenesis(t *testing.T) {
	//genValidators1 := make([]types.Validator, 1, 5)
	//pk := ed25519.GenPrivKey().PubKey()
	//genValidators1[0] = testvalidator.NewValidator(t, sdk.ValAddress(pk.Address()), pk)
	//genValidators1[0].Tokens = sdk.OneInt()
	//genValidators1[0].DelegatorShares = sdk.OneDec()
	//
	//tests := []struct {
	//	name    string
	//	mutate  func(*types.GenesisState)
	//	wantErr bool
	//}{
	//	{"default", func(*types.GenesisState) {}, false},
	//	// validate genesis validators
	//	{"duplicate validator", func(data *types.GenesisState) {
	//		data.Validators = genValidators1
	//		data.Validators = append(data.Validators, genValidators1[0])
	//	}, true},
	//	{"no delegator shares", func(data *types.GenesisState) {
	//		data.Validators = genValidators1
	//		data.Validators[0].DelegatorShares = sdk.ZeroDec()
	//	}, true},
	//	{"jailed and bonded validator", func(data *types.GenesisState) {
	//		data.Validators = genValidators1
	//		data.Validators[0].Jailed = true
	//		data.Validators[0].Status = types.Bonded
	//	}, true},
	//}
	//
	//for _, tt := range tests {
	//	tt := tt
	//
	//	t.Run(tt.name, func(t *testing.T) {
	//		genesisState := types.DefaultGenesisState()
	//		tt.mutate(genesisState)
	//
	//		if tt.wantErr {
	//			assert.Error(t, validator.ValidateGenesis(genesisState))
	//		} else {
	//			assert.NoError(t, validator.ValidateGenesis(genesisState))
	//		}
	//	})
	//}
}