package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper implements the module data storaging.
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
	ps       paramtypes.Subspace

	accountKeeper auth.AccountKeeperI
	bankKeeper    bank.Keeper
}

// NewKeeper creates a multisig keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper auth.AccountKeeperI,
	bankKeeper bank.Keeper,
) *Keeper {
	return &Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		ps:            ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

////////////////////////////////////////////////////////////////
// Wallet
////////////////////////////////////////////////////////////////

// GetWallet returns multisig wallet metadata struct with specified address.
func (k Keeper) GetWallet(ctx sdk.Context, address string) (wallet types.Wallet, err error) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.KeyPrefixWallet, []byte(address)...)
	value := store.Get(key)
	if len(value) == 0 {
		err = fmt.Errorf("wallet %s is not found in the key-value store", address)
		return
	}
	err = k.cdc.UnmarshalLengthPrefixed(value, &wallet)
	return
}

// SetWallet sets the entire wallet metadata struct for a multisig wallet.
func (k Keeper) SetWallet(ctx sdk.Context, wallet types.Wallet) {
	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshalLengthPrefixed(&wallet)
	key := append(types.KeyPrefixWallet, []byte(wallet.Address)...)
	store.Set(key, value)
}

// GetWallets returns multisig wallets metadata struct for specified owner.
func (k Keeper) GetWallets(ctx sdk.Context, owner string) (wallets []types.Wallet, err error) {
	store := ctx.KVStore(k.storeKey)

	for iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixWallet)); iterator.Valid(); iterator.Next() {
		var wallet types.Wallet
		value := iterator.Value()
		if len(value) == 0 {
			err = fmt.Errorf("empty value in the key-value store")
			return
		}
		err = k.cdc.UnmarshalLengthPrefixed(value, &wallet)
		if err != nil {
			return
		}
		for _, o := range wallet.Owners {
			if o == owner {
				wallets = append(wallets, wallet)
				break
			}
		}
	}

	return
}

////////////////////////////////////////////////////////////////
// Transaction
////////////////////////////////////////////////////////////////

// GetTransaction returns multisig wallet transaction metadata with specified address transaction ID.
func (k Keeper) GetTransaction(ctx sdk.Context, txID string) (transaction types.Transaction, err error) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.KeyPrefixTransaction, []byte(txID)...)
	value := store.Get(key)
	if len(value) == 0 {
		err = fmt.Errorf("transaction %s is not found in the key-value store", txID)
		return
	}
	err = k.cdc.UnmarshalLengthPrefixed(value, &transaction)
	return
}

// SetTransaction sets the entire multisig wallet transaction metadata struct for a multisig wallet.
func (k Keeper) SetTransaction(ctx sdk.Context, transaction types.Transaction) {
	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshalLengthPrefixed(&transaction)
	key := append(types.KeyPrefixTransaction, []byte(transaction.Id)...)
	store.Set(key, value)
}

// GetTransactions returns transactions for specified multisig wallet.
func (k Keeper) GetTransactions(ctx sdk.Context, wallet string) (transactions []types.Transaction, err error) {
	store := ctx.KVStore(k.storeKey)

	for iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixTransaction)); iterator.Valid(); iterator.Next() {
		var tx types.Transaction
		value := iterator.Value()
		if len(value) == 0 {
			err = fmt.Errorf("empty value in the key-value store")
			return
		}
		err = k.cdc.UnmarshalLengthPrefixed(value, &tx)
		if err != nil {
			return
		}
		if tx.Wallet == wallet {
			transactions = append(transactions, tx)
		}
	}

	return
}