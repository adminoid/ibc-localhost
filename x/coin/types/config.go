package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"bitbucket.org/decimalteam/go-smart-node/utils/helpers"
)

const maxCoinNameBytes = 64
const allowedCoinSymbols = "^[a-zA-Z][a-zA-Z0-9]{2,9}$"

var coinSymbolValidator = regexp.MustCompile(allowedCoinSymbols)

var MinCoinSupply = helpers.EtherToWei(sdk.NewInt(1))
var MaxCoinSupply = helpers.EtherToWei(sdk.NewInt(1000000000000000))
var MinCoinReserve = helpers.EtherToWei(sdk.NewInt(1000))
