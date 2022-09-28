package keeper

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"bitbucket.org/decimalteam/go-smart-node/utils/events"
	"bitbucket.org/decimalteam/go-smart-node/x/swap/errors"

	"bitbucket.org/decimalteam/go-smart-node/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = &Keeper{}

func (k Keeper) InitializeSwap(goCtx context.Context, msg *types.MsgInitializeSwap) (*types.MsgInitializeSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasChain(ctx, msg.DestChain) {
		return nil, errors.ChainDoesNotExists
	}
	if !k.HasChain(ctx, msg.FromChain) {
		return nil, errors.ChainDoesNotExists
	}

	funds := sdk.NewCoins(sdk.NewCoin(strings.ToLower(msg.TokenSymbol), msg.Amount))

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, errors.InvalidSenderAddress
	}

	ok, err := k.CheckBalance(ctx, sender, funds)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.InsufficientAccountFunds
	}

	err = k.LockFunds(ctx, sender, funds)
	if err != nil {
		return nil, err
	}

	err = events.EmitTypedEvent(ctx, &types.EventInitializeSwap{
		Sender:            msg.Sender,
		From:              msg.Sender,
		DestChain:         msg.DestChain,
		Recipient:         msg.Recipient,
		Amount:            msg.Amount.String(),
		TransactionNumber: msg.TransactionNumber,
		TokenSymbol:       msg.TokenSymbol,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	return &types.MsgInitializeSwapResponse{}, nil

}

func (k Keeper) RedeemSwap(goCtx context.Context, msg *types.MsgRedeemSwap) (*types.MsgRedeemSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	transactionNumber, ok := sdk.NewIntFromString(msg.TransactionNumber)
	if !ok {
		return nil, errors.InvalidTransactionNumber
	}

	hash, err := types.GetHash(transactionNumber, msg.TokenSymbol, msg.Amount, msg.Recipient, msg.FromChain, msg.DestChain)
	if err != nil {
		return nil, err
	}

	if k.HasSwap(ctx, hash) {
		return nil, errors.AlreadyRedeemed
	}

	R := big.NewInt(0)
	R.SetBytes(msg.R[:])

	S := big.NewInt(0)
	S.SetBytes(msg.S[:])

	address, err := types.Ecrecover(hash, R, S, sdk.NewInt(int64(msg.V)).BigInt())
	if err != nil {
		return nil, err
	}

	params := k.GetParams(ctx)

	if hex.EncodeToString(address.Bytes()) != params.CheckingAddress {
		return nil, errors.InvalidServiceAddress
	}

	k.SetSwap(ctx, hash)

	funds := sdk.NewCoins(sdk.NewCoin(strings.ToLower(msg.TokenSymbol), msg.Amount))

	if !k.CheckPoolFunds(ctx, funds) {
		return nil, errors.InsufficientPoolFunds
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, errors.InvalidSenderAddress
	}

	err = k.UnlockFunds(ctx, recipient, funds)
	if err != nil {
		return nil, err
	}

	err = events.EmitTypedEvent(ctx, &types.EventRedeemSwap{
		Sender:            msg.Sender,
		From:              msg.From,
		DestChain:         msg.DestChain,
		Recipient:         msg.Recipient,
		Amount:            msg.Amount.String(),
		TransactionNumber: msg.TransactionNumber,
		TokenSymbol:       msg.TokenSymbol,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	return &types.MsgRedeemSwapResponse{}, nil

}

func (k Keeper) ActivateChain(goCtx context.Context, msg *types.MsgActivateChain) (*types.MsgActivateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if msg.Sender != params.ServiceAddress {
		return nil, errors.SenderIsNotSwapService
	}

	chain, found := k.GetChain(ctx, msg.ID)
	if found {
		chain.Active = true
	} else {
		chain = types.NewChain(msg.ID, msg.Name, true)
	}

	k.SetChain(ctx, &chain)

	err := events.EmitTypedEvent(ctx, &types.EventActivateChain{
		ID:   msg.ID,
		Name: msg.Name,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	return &types.MsgActivateChainResponse{}, nil
}

func (k Keeper) DeactivateChain(goCtx context.Context, msg *types.MsgDeactivateChain) (*types.MsgDeactivateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if msg.Sender != params.ServiceAddress {
		return nil, errors.SenderIsNotSwapService
	}

	chain, found := k.GetChain(ctx, msg.ID)
	if !found {
		return nil, errors.ChainDoesNotExists
	}

	chain.Active = false
	k.SetChain(ctx, &chain)

	err := events.EmitTypedEvent(ctx, &types.EventDeactivateChain{
		ID: msg.ID,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	return &types.MsgDeactivateChainResponse{}, nil
}
