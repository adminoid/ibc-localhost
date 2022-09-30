package keeper

import (
	"context"
	"strings"

	"bitbucket.org/decimalteam/go-smart-node/utils/events"
	"bitbucket.org/decimalteam/go-smart-node/x/multisig/errors"

	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateWallet(goCtx context.Context, msg *types.MsgCreateWallet) (*types.MsgCreateWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create new multisig wallet
	wallet, err := types.NewWallet(msg.Owners, msg.Weights, msg.Threshold, ctx.TxBytes())
	if err != nil {
		return nil, errors.UnableToCreateWallet
	}

	// Ensure multisig wallet with the address does not exist
	_, err = k.GetWallet(ctx, wallet.Address)
	if err == nil {
		return nil, errors.WalletAlreadyExists
	}

	adr, err := sdk.AccAddressFromBech32(wallet.Address)
	if err != nil {
		return nil, errors.InvalidWallet
	}
	// Ensure account with multisig address does not exist
	existingAccount := k.accountKeeper.GetAccount(ctx, adr)
	if existingAccount != nil && !existingAccount.GetAddress().Empty() {
		return nil, errors.AccountAlreadyExists
	}

	// Save created multisig wallet to the KVStore
	k.SetWallet(ctx, *wallet)

	// Emit transaction events
	err = events.EmitTypedEvent(ctx, &types.EventCreateWallet{
		Sender:    msg.Sender,
		Wallet:    wallet.Address,
		Owners:    msg.Owners,
		Weights:   msg.Weights,
		Threshold: msg.Threshold,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	return &types.MsgCreateWalletResponse{
		Wallet: wallet.Address,
	}, nil
}

func (k Keeper) CreateTransaction(goCtx context.Context, msg *types.MsgCreateTransaction) (*types.MsgCreateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve multisig wallet from the KVStore
	wallet, err := k.GetWallet(ctx, msg.Wallet)
	if err != nil {
		return nil, err
	}
	adr, err := sdk.AccAddressFromBech32(wallet.Address)
	if err != nil {
		return nil, errors.InvalidWallet
	}
	// Retrieve coins hold on the multisig wallet
	walletCoins := k.bankKeeper.GetAllBalances(ctx, adr)

	// Ensure there are enough coins on the multisig wallet
	for _, coin := range msg.Coins {
		coinName := strings.ToLower(coin.Denom)
		if walletCoins.AmountOf(coinName).LT(coin.Amount) {
			return nil, errors.InsufficientFunds
		}
	}

	// Create new multisig transaction
	transaction, err := types.NewTransaction(
		msg.Wallet,
		msg.Receiver,
		msg.Coins,
		len(wallet.Owners),
		ctx.BlockHeight(),
		ctx.TxBytes(),
	)
	if err != nil {
		return nil, err
	}

	// Save created multisig transaction to the KVStore
	k.SetTransaction(ctx, *transaction)

	// Emit transaction events
	err = events.EmitTypedEvent(ctx, &types.EventCreateTransaction{
		Sender:      msg.Sender,
		Wallet:      msg.Wallet,
		Receiver:    msg.Receiver,
		Coins:       msg.Coins.String(),
		Transaction: transaction.Id,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	// Sign created multisig transaction by the creator
	_, err = k.SignTransaction(goCtx, &types.MsgSignTransaction{
		Sender: msg.Sender,
		ID:     transaction.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateTransactionResponse{
		ID: transaction.Id,
	}, nil
}

func (k Keeper) SignTransaction(goCtx context.Context, msg *types.MsgSignTransaction) (*types.MsgSignTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve multisig transaction from the KVStore
	transaction, err := k.GetTransaction(ctx, msg.ID)
	if err != nil {
		return nil, err
	}

	// Retrieve multisig wallet from the KVStore
	wallet, err := k.GetWallet(ctx, transaction.Wallet)
	if err != nil {
		return nil, err
	}

	// Calculate current weight of signatures
	confirmations := uint32(0)
	for i := 0; i < len(wallet.Owners); i++ {
		if transaction.Signers[i] != "" {
			confirmations += wallet.Weights[i]
		}
	}

	// Ensure current weight of signatures is not enough
	if confirmations >= wallet.Threshold {
		return nil, errors.AlreadyEnoughSignatures
	}

	// Append the signature to the multisig transaction
	weight := uint32(0)
	signed := false
	for i := 0; i < len(wallet.Owners); i++ {
		if wallet.Owners[i] != msg.Sender {
			continue
		}
		if transaction.Signers[i] != "" {
			return nil, errors.TransactionAlreadySigned
		}
		signed = true
		weight = wallet.Weights[i]
		confirmations += weight
		transaction.Signers[i] = msg.Sender
		break
	}
	if !signed {
		return nil, errors.SignerIsNotOwner
	}

	// Save updated multisig transaction to the KVStore
	k.SetTransaction(ctx, transaction)

	// Check if new weight of signatures is enough to perform multisig transaction
	confirmed := confirmations >= wallet.Threshold
	if confirmed {
		wAdr, err := sdk.AccAddressFromBech32(wallet.Address)
		if err != nil {
			return nil, errors.InvalidWallet
		}
		rAdr, err := sdk.AccAddressFromBech32(transaction.Receiver)
		if err != nil {
			return nil, errors.InvalidReceiver
		}
		// Perform transaction
		err = k.bankKeeper.SendCoins(ctx, wAdr, rAdr, transaction.Coins)
		if err != nil {
			return nil, err
		}
	}

	// Emit transaction events
	err = events.EmitTypedEvent(ctx, &types.EventSignTransaction{
		Sender:        msg.Sender,
		Wallet:        wallet.Address,
		Transaction:   transaction.Id,
		SignerWeight:  weight,
		Confirmations: confirmations,
		Confirmed:     confirmed,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	return &types.MsgSignTransactionResponse{}, nil
}

func (k Keeper) CreateUniversalTransaction(goCtx context.Context, msg *types.MsgCreateUniversalTransaction) (*types.MsgCreateUniversalTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve multisig wallet from the KVStore
	wallet, err := k.GetWallet(ctx, msg.Wallet)
	if err != nil {
		return nil, err
	}
	_, err = sdk.AccAddressFromBech32(wallet.Address)
	if err != nil {
		return nil, errors.InvalidWallet
	}

	// Create new multisig transaction
	transaction, err := types.NewUniversalTransaction(
		k.cdc,
		msg.Wallet,
		*msg.Content,
		len(wallet.Owners),
		ctx.BlockHeight(),
		ctx.TxBytes(),
	)
	if err != nil {
		return nil, err
	}

	// Check internal transaction
	var internal sdk.Msg
	err = k.cdc.UnpackAny(msg.Content, &internal)
	if err != nil {
		return nil, err
	}
	err = internal.ValidateBasic()
	if err != nil {
		return nil, err
	}
	handler := k.router.Handler(internal)
	if handler == nil {
		return nil, errors.NoHandlerForInternal
	}

	// Save created multisig transaction to the KVStore
	k.SetUniversalTransaction(ctx, *transaction)

	// Emit transaction events
	err = events.EmitTypedEvent(ctx, &types.EventCreateUniversalTransaction{
		Sender:      msg.Sender,
		Wallet:      msg.Wallet,
		Transaction: transaction.Id,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	// Sign created multisig transaction by the creator
	_, err = k.SignUniversalTransaction(goCtx, &types.MsgSignUniversalTransaction{
		Sender: msg.Sender,
		ID:     transaction.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateUniversalTransactionResponse{
		ID: transaction.Id,
	}, nil

}

func (k Keeper) SignUniversalTransaction(goCtx context.Context, msg *types.MsgSignUniversalTransaction) (*types.MsgSignUniversalTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve multisig transaction from the KVStore
	transaction, err := k.GetUniversalTransaction(ctx, msg.ID)
	if err != nil {
		return nil, err
	}

	// Retrieve multisig wallet from the KVStore
	wallet, err := k.GetWallet(ctx, transaction.Wallet)
	if err != nil {
		return nil, err
	}

	if k.IsSigned(ctx, msg.ID, msg.Sender) {
		return nil, errors.TransactionAlreadySigned
	}

	// Calculate current weight of signatures and check sender
	confirmations := uint32(0)
	senderIsOwner := false
	senderWeight := uint32(0)
	for i, owner := range wallet.Owners {
		if k.IsSigned(ctx, msg.ID, owner) {
			confirmations += wallet.Weights[i]
		}
		if owner == msg.Sender {
			senderIsOwner = true
			senderWeight = wallet.Weights[i]
		}
	}

	if !senderIsOwner {
		return nil, errors.SignerIsNotOwner
	}
	// Ensure current weight of signatures is not enough
	if confirmations >= wallet.Threshold {
		return nil, errors.AlreadyEnoughSignatures
	}

	// Append the signature to the multisig transaction
	k.SetUniversalSign(ctx, msg.ID, msg.Sender)

	confirmations += senderWeight

	// Check if new weight of signatures is enough to perform multisig transaction
	confirmed := confirmations >= wallet.Threshold

	// Emit transaction events
	err = events.EmitTypedEvent(ctx, &types.EventSignUniversalTransaction{
		Sender:        msg.Sender,
		Wallet:        wallet.Address,
		Transaction:   transaction.Id,
		SignerWeight:  senderWeight,
		Confirmations: confirmations,
		Confirmed:     confirmed,
	})
	if err != nil {
		return nil, errors.Internal.Wrapf("err: %s", err.Error())
	}

	if confirmed {
		var msg sdk.Msg
		err = k.cdc.UnpackAny(&transaction.Message, &msg)
		if err != nil {
			return nil, err
		}
		handler := k.router.Handler(msg)
		if handler == nil {
			return nil, errors.NoHandlerForInternal
		}
		// TODO: do we need result?
		_, err = handler(ctx, msg)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgSignUniversalTransactionResponse{}, nil
}
