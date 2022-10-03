package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Wallet(c context.Context, req *types.QueryWalletRequest) (*types.QueryWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	wallet, err := k.GetWallet(ctx, req.Wallet)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryWalletResponse{Wallet: wallet}, nil
}

func (k Keeper) Wallets(c context.Context, req *types.QueryWalletsRequest) (*types.QueryWalletsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixWallet)
	wallets := make([]types.Wallet, 0)
	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(_, value []byte) error {
			var wallet types.Wallet

			if err := k.cdc.UnmarshalLengthPrefixed(value, &wallet); err != nil {
				return err
			}
			for _, o := range wallet.Owners {
				if req.Owner == o {
					wallets = append(wallets, wallet)
					break
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryWalletsResponse{Wallets: wallets, Pagination: pageRes}, nil
}

func (k Keeper) Transaction(c context.Context, req *types.QueryTransactionRequest) (*types.QueryTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	tx, err := k.GetTransaction(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTransactionResponse{Transaction: tx}, nil
}

func (k Keeper) Transactions(c context.Context, req *types.QueryTransactionsRequest) (*types.QueryTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTransaction)
	transactions := make([]types.Transaction, 0)
	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(_, value []byte) error {
			var tx types.Transaction

			if err := k.cdc.UnmarshalLengthPrefixed(value, &tx); err != nil {
				return err
			}

			transactions = append(transactions, tx)
			return nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTransactionsResponse{Transactions: transactions, Pagination: pageRes}, nil
}

func (k Keeper) UniversalTransactions(c context.Context, req *types.QueryUniversalTransactionsRequest) (*types.QueryUniversalTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	wallet, err := k.GetWallet(ctx, req.Wallet)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUniversalTransaction)
	txResponses := make([]types.UniversalTransactionResponse, 0)
	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(_, value []byte) error {
			var tx types.UniversalTransaction
			var txres types.UniversalTransactionResponse

			if err := k.cdc.UnmarshalLengthPrefixed(value, &tx); err != nil {
				return err
			}
			if tx.Wallet == req.Wallet {
				txres.Transaction = tx
				txres.Completed = k.IsCompleted(ctx, tx.Id)
				for _, owner := range wallet.Owners {
					if k.IsSigned(ctx, tx.Id, owner) {
						txres.Signers = append(txres.Signers, owner)
					}
				}
				txResponses = append(txResponses, txres)
			}
			return nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUniversalTransactionsResponse{Transactions: txResponses, Pagination: pageRes}, nil
}

func (k Keeper) UniversalTransaction(c context.Context, req *types.QueryUniversalTransactionRequest) (*types.QueryUniversalTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	tx, err := k.GetUniversalTransaction(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	wallet, err := k.GetWallet(ctx, tx.Wallet)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var txres types.UniversalTransactionResponse

	txres.Transaction = tx
	txres.Completed = k.IsCompleted(ctx, tx.Id)
	for _, owner := range wallet.Owners {
		if k.IsSigned(ctx, tx.Id, owner) {
			txres.Signers = append(txres.Signers, owner)
		}
	}

	return &types.QueryUniversalTransactionResponse{Transaction: txres}, nil
}
