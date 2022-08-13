package keeper

import (
	"bitbucket.org/decimalteam/go-smart-node/x/nft/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryCollectionSupply(c context.Context, req *types.QueryCollectionSupplyRequest) (*types.QueryCollectionSupplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	collection, found := k.GetCollection(ctx, req.Denom)
	if !found {
		return nil, types.ErrUnknownCollection(req.Denom)
	}

	return &types.QueryCollectionSupplyResponse{Supply: int64(collection.Supply())}, nil
}

func (k Keeper) QueryOwnerCollections(c context.Context, req *types.QueryOwnerCollectionsRequest) (*types.QueryOwnerCollectionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ownerAddress, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	owner := types.Owner{
		Address: req.GetOwner(),
	}
	if req.Denom == "" {
		owner.Collections = k.GetOwnerCollections(ctx, ownerAddress)
	} else {
		collection, found := k.GetOwnerCollectionByDenom(ctx, ownerAddress, req.Denom)
		if !found {
			collection = types.NewOwnerCollection(req.Denom, []string{})
		}
		owner.Collections = append(owner.Collections, collection)
	}

	return &types.QueryOwnerCollectionsResponse{Owner: owner}, nil
}

func (k Keeper) QueryCollection(c context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	collection, found := k.GetCollection(ctx, req.Denom)
	if !found {
		return nil, types.ErrUnknownCollection(req.Denom)
	}

	return &types.QueryCollectionResponse{
		Collection: collection,
	}, nil
}

func (k Keeper) QueryDenoms(c context.Context, _ *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	denoms := k.GetDenoms(ctx)

	return &types.QueryDenomsResponse{
		Denoms: denoms,
	}, nil
}

func (k Keeper) QueryNFT(c context.Context, req *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nft, err := k.GetNFT(ctx, req.Denom, req.TokenID)
	if err != nil {
		return nil, types.ErrUnknownNFT(req.Denom, req.TokenID)
	}

	return &types.QueryNFTResponse{
		Nft: nft,
	}, nil
}

func (k Keeper) QuerySubTokens(c context.Context, req *types.QuerySubTokensRequest) (*types.QuerySubTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var subTokens []types.SubToken
	for _, id := range req.SubTokenIDs {
		subToken, ok := k.GetSubToken(ctx, req.TokenID, id)
		if !ok {
			continue
		}
		subTokens = append(subTokens, subToken)
	}

	return &types.QuerySubTokensResponse{
		SubTokens: subTokens,
	}, nil
}