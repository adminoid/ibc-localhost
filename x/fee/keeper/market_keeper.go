package keeper

import (
	"math/big"

	"bitbucket.org/decimalteam/go-smart-node/cmd/config"
	feeconfig "bitbucket.org/decimalteam/go-smart-node/x/fee/config"
	"bitbucket.org/decimalteam/go-smart-node/x/fee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
)

// implementation of interface FeeMarketKeeper
// for evm module
var _ types.FeeMarketKeeper = Keeper{}

var defaultBase = sdk.NewInt(1000000000)

func (k Keeper) GetBaseFee(ctx sdk.Context) *big.Int {
	price, err := k.GetPrice(ctx, config.BaseDenom, feeconfig.DefaultQuote)
	if err != nil {
		// fallback to default price
		return defaultBase.BigInt()
	}
	fee := sdk.OneDec().MulInt(defaultBase).Quo(price.Price).RoundInt()
	return fee.BigInt()
}

func (k Keeper) GetParams(ctx sdk.Context) feemarkettypes.Params {
	// TODO: watch for new params
	return feemarkettypes.NewParams(
		false,                                  //noBaseFee bool,
		1,                                      //baseFeeChangeDenom,
		1,                                      //elasticityMultiplier uint32,
		k.GetBaseFee(ctx).Uint64(),             //baseFee uint64,
		0,                                      //enableHeight int64,
		feemarkettypes.DefaultMinGasPrice,      //minGasPrice sdk.Dec,
		feemarkettypes.DefaultMinGasMultiplier, //minGasPriceMultiplier sdk.Dec,
	)
}

func (k Keeper) AddTransientGasWanted(ctx sdk.Context, gasWanted uint64) (uint64, error) {
	// TODO: this function is used in NewGasWantedDecorator
	// Do we need implement?
	return 0, nil
}
