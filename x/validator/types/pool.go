package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	BondedPoolName    = "bonded_tokens_pool"
	NotBondedPoolName = "not_bonded_tokens_pool"
)

var BondedPoolAddress = authtypes.NewModuleAddress(BondedPoolName).String()
var NotBondedPoolAddress = authtypes.NewModuleAddress(NotBondedPoolName).String()

// NewPool creates a new Pool instance used for queries.
func NewPool(bonded sdk.Coins, notBonded sdk.Coins) Pool {
	return Pool{
		Bonded:    bonded,
		NotBonded: notBonded,
	}
}
