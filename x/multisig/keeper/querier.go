package keeper

import (
	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryWallet:
			return queryWallet(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryWallets:
			return queryWallets(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryTransaction:
			return queryTransaction(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryTransactions:
			return queryTransactions(ctx, path[1:], req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.ErrUnknownRequest
		}
	}
}

func queryWallet(ctx sdk.Context, path []string, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	address := path[0]

	wallet, err := k.GetWallet(ctx, address)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, wallet)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func queryWallets(ctx sdk.Context, path []string, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	owner := path[0]

	wallets, err := k.GetWallets(ctx, owner)
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, wallets)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func queryTransaction(ctx sdk.Context, path []string, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	txID := path[0]

	tx, err := k.GetTransaction(ctx, txID)
	if err != nil {
		return nil, err
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, tx)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func queryTransactions(ctx sdk.Context, path []string, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	wallet := path[0]

	transactions, err := k.GetTransactions(ctx, wallet)
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, transactions)
	if err != nil {
		return nil, err
	}

	return bz, nil
}
