package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"bitbucket.org/decimalteam/go-smart-node/x/validator/types"
)

func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetValidatorKey(addr))
	if value == nil {
		return validator, false
	}
	validator = types.MustUnmarshalValidator(k.cdc, value)
	k.MustGetValidatorRewards(ctx, &validator)

	return validator, true
}

// compatible for StakingKeeper for Ethermint
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	addr := store.Get(types.GetValidatorByConsAddrIndexKey(consAddr))
	if addr == nil {
		return validator, false
	}
	val, found := k.GetValidator(ctx, addr)
	if !found {
		return stakingtypes.Validator{}, false
	}
	// TODO: make right conversion
	pk, err := val.ConsPubKey()
	if err != nil {
		return stakingtypes.Validator{}, false
	}
	stVal, err := stakingtypes.NewValidator(
		val.GetOperator(),
		pk,
		stakingtypes.Description(val.Description),
	)
	if err != nil {
		return stakingtypes.Validator{}, false
	}
	return stVal, true
}

func (k Keeper) GetValidatorByConsAddrDecimal(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	addr := store.Get(types.GetValidatorByConsAddrIndexKey(consAddr))
	if addr == nil {
		return validator, false
	}
	// TODO: convert our validator for staking types
	return k.GetValidator(ctx, addr)
}

func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, &validator)
	store.Set(types.GetValidatorKey(validator.GetOperator()), bz)
}

func (k Keeper) SetValidatorRewards(ctx sdk.Context, valAddr sdk.ValAddress, rewards types.ValidatorRewards) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidatorRewards(k.cdc, &rewards)
	store.Set(types.GetValidatorRewards(valAddr), bz)
}

func (k Keeper) GetValidatorRewards(ctx sdk.Context, valAddr sdk.ValAddress) (rewards types.ValidatorRewards, err error) {
	/*
		store := ctx.KVStore(k.storeKey)
		value := store.Get(types.GetValidatorKey(valAddr))
		if value == nil {
			return rewards, fmt.Errorf("not found rewards for validator")
		}
		rewards = types.MustUnmarshalValidatorRewards(k.cdc, value)
		return rewards, nil
	*/
	return types.ValidatorRewards{
		sdkmath.ZeroInt(),
		sdkmath.ZeroInt(),
	}, nil
}

func (k Keeper) MustGetValidatorRewards(ctx sdk.Context, validator *types.Validator) {
	rewards, err := k.GetValidatorRewards(ctx, validator.GetOperator())
	if err != nil {
		panic(err)
	}
	validator.Rewards = rewards.Rewards
	validator.TotalRewards = rewards.TotalRewards
}

// validator index
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) error {
	consPk, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorByConsAddrIndexKey(consPk), validator.GetOperator())
	return nil
}

// validator index
func (k Keeper) SetValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	// jailed validators are not kept in the power index
	if validator.Jailed {
		return
	}
	total, err := k.TotalStakeInBaseCoin(ctx, validator.GetOperator())
	if err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(k.GetValidatorByPowerIndexKey(validator, TokensToConsensusPower(total)), validator.GetOperator())
}

// validator index
func (k Keeper) SetNewValidatorByPowerIndex(ctx sdk.Context, validator types.Validator, power int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(k.GetValidatorByPowerIndexKey(validator, power), validator.GetOperator())
}

func (k Keeper) SetValidatorByPowerIndexWithoutCalc(ctx sdk.Context, validator types.Validator, power int64) {
	// jailed validators are not kept in the power index
	if validator.Jailed {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(k.GetValidatorByPowerIndexKey(validator, power), validator.GetOperator())
}

// validator index
func (k Keeper) DeleteValidatorByPowerIndex(ctx sdk.Context, validator types.Validator, power int64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(k.GetValidatorByPowerIndexKey(validator, power))
}

func (k Keeper) GetAllValidatorsByPowerIndex(ctx sdk.Context) (types.Validators, []int64, sdkmath.Int) {
	validators := make([]types.Validator, 0)
	powers := make([]int64, 0)
	totalPower := sdk.ZeroInt()
	iterator := k.ValidatorsPowerStoreIterator(ctx)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		powerBytes := iterator.Key()[1:9]
		power := binary.BigEndian.Uint64(powerBytes)
		totalPower.Add(sdk.NewIntFromUint64(power))

		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		k.MustGetValidatorRewards(ctx, &validator)

		validators = append(validators, validator)
		powers = append(powers, int64(power))
	}

	return validators, powers, totalPower
}

// GetValidatorByPowerIndexKey creates the validator by power index.
// Power index is the key used in the power-store, and represents the relative power ranking of the validator.
func (k Keeper) GetValidatorByPowerIndexKey(validator types.Validator, power int64) []byte {
	// NOTE the address doesn't need to be stored because counter bytes must always be different
	// NOTE the larger values are of higher value

	//key := types.GetValidatorsByPowerIndexKey()
	consensusPowerBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(consensusPowerBytes, uint64(power))

	powerBytes := consensusPowerBytes
	powerBytesLen := len(powerBytes) // 8

	operAddrInvr := sdk.CopyBytes(validator.GetOperator())
	addrLen := len(operAddrInvr)

	for i, b := range operAddrInvr {
		operAddrInvr[i] = ^b
	}

	// key is of format prefix || powerbytes || addrLen (1byte) || addrBytes
	key := make([]byte, 1+powerBytesLen+1+addrLen)

	key[0] = types.GetValidatorsByPowerIndexKey()[0]
	copy(key[1:powerBytesLen+1], powerBytes)
	key[powerBytesLen+1] = byte(addrLen)
	copy(key[powerBytesLen+2:], operAddrInvr)

	return key
}

//// validator index
//func (k Keeper) SetNewValidatorByPowerIndex(ctx sdk.Context, validator types.Validator, power int64) {
//	store := ctx.KVStore(k.storeKey)
//	store.Set(k.GetValidatorByPowerIndexKey(ctx, validator, power), validator.GetOperator())
//}

//// Update the tokens of an existing validator, update the validators power index key
//func (k Keeper) AddValidatorTokensAndShares(ctx sdk.Context, validator types.Validator, tokensToAdd sdkmath.Int) (valOut types.Validator, addedShares sdk.Dec) {
//	k.DeleteValidatorByPowerIndex(ctx, validator)
//	validator, addedShares = validator.AddTokensFromDel(tokensToAdd)
//	k.SetValidator(ctx, validator)
//	k.SetValidatorByPowerIndex(ctx, validator)
//	return validator, addedShares
//}
//
//// Update the tokens of an existing validator, update the validators power index key
//func (k Keeper) RemoveValidatorTokensAndShares(ctx sdk.Context, validator types.Validator, sharesToRemove sdk.Dec) (valOut types.Validator, removedTokens sdkmath.Int) {
//	k.DeleteValidatorByPowerIndex(ctx, validator)
//	validator, removedTokens = validator.RemoveDelShares(sharesToRemove)
//	k.SetValidator(ctx, validator)
//	k.SetValidatorByPowerIndex(ctx, validator)
//	return validator, removedTokens
//}
//
//// Update the tokens of an existing validator, update the validators power index key
//func (k Keeper) RemoveValidatorTokens(ctx sdk.Context, validator types.Validator, tokensToRemove sdkmath.Int) types.Validator {
//	k.DeleteValidatorByPowerIndex(ctx, validator)
//	validator = validator.RemoveTokens(tokensToRemove)
//	k.SetValidator(ctx, validator)
//	k.SetValidatorByPowerIndex(ctx, validator)
//	return validator
//}

// remove the validator record and associated indexes
// except for the bonded validator index which is only handled in ApplyAndReturnTendermintUpdates
// TODO, this function panics, and it's not good.
func (k Keeper) RemoveValidator(ctx sdk.Context, address sdk.ValAddress) {
	// first retrieve the old validator record
	validator, found := k.GetValidator(ctx, address)
	if !found {
		return
	}

	if !validator.IsUnbonded() {
		panic("cannot call RemoveValidator on bonded or unbonding validators")
	}

	if validator.TotalRewards.IsPositive() {
		panic("attempting to remove a validator which still contains tokens")
	}

	valConsAddr, err := validator.GetConsAddr()
	if err != nil {
		panic(err)
	}

	// delete the old validator record
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorKey(address))
	store.Delete(types.GetValidatorByConsAddrIndexKey(valConsAddr))
	store.Delete(types.GetValidatorRewards(address))
	//store.Delete(k.GetValidatorsByPowerIndexKey(ctx, validator))

	// call hooks
	k.AfterValidatorRemoved(ctx, valConsAddr, validator.GetOperator())
}

// get groups of validators

// get the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetValidatorsKey())
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		k.MustGetValidatorRewards(ctx, &validator)

		validators = append(validators, validator)
	}

	return validators
}

// return a given amount of all the validators
func (k Keeper) GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validators = make([]types.Validator, maxRetrieve)

	iterator := sdk.KVStorePrefixIterator(store, types.GetValidatorsKey())
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		k.MustGetValidatorRewards(ctx, &validator)

		validators[i] = validator
		i++
	}

	return validators[:i] // trim if the array length < maxRetrieve
}

// get the current group of bonded validators sorted by power-rank
func (k Keeper) GetBondedValidatorsByPower(ctx sdk.Context) []types.Validator {
	maxValidators := k.MaxValidators(ctx)
	validators := make([]types.Validator, maxValidators)

	iterator := k.ValidatorsPowerStoreIterator(ctx)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxValidators); iterator.Next() {
		address := iterator.Value()
		validator := k.mustGetValidator(ctx, address)
		k.MustGetValidatorRewards(ctx, &validator)

		if validator.IsBonded() {
			validators[i] = validator
			i++
		}
	}

	return validators[:i] // trim
}

// returns an iterator for the current validator power store
func (k Keeper) ValidatorsPowerStoreIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStoreReversePrefixIterator(store, types.GetValidatorsByPowerIndexKey())
}

// Last Validator Index

// Load the last validator power.
// Returns zero if the operator was not a validator last block.
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetLastValidatorPowerKey(operator))
	if bz == nil {
		return 0
	}

	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(bz, &intV)

	return intV.GetValue()
}

// Set the last validator power.
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetLastValidatorPowerKey(operator), bz)
}

// Delete the last validator power.
func (k Keeper) DeleteLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLastValidatorPowerKey(operator))
}

// returns an iterator for the consensus validators in the last block
func (k Keeper) LastValidatorsIterator(ctx sdk.Context) (iterator sdk.Iterator) {
	store := ctx.KVStore(k.storeKey)
	iterator = sdk.KVStorePrefixIterator(store, types.GetLastValidatorPowersKey())

	return iterator
}

// Iterate over last validator powers.
func (k Keeper) IterateLastValidatorPowers(ctx sdk.Context, handler func(operator sdk.ValAddress, power int64) (stop bool)) {
	iter := k.LastValidatorsIterator(ctx)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(types.AddressFromLastValidatorPowerKey(iter.Key()))
		intV := &gogotypes.Int64Value{}

		k.cdc.MustUnmarshal(iter.Value(), intV)

		if handler(addr, intV.GetValue()) {
			break
		}
	}
}

// get the group of the bonded validators
func (k Keeper) GetLastValidators(ctx sdk.Context) (validators []types.Validator) {
	// add the actual validator power sorted store
	maxValidators := k.MaxValidators(ctx)
	validators = make([]types.Validator, maxValidators)

	iterator := k.LastValidatorsIterator(ctx)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid(); iterator.Next() {
		// sanity check
		if i >= int(maxValidators) {
			panic("more validators than maxValidators found")
		}

		address := types.AddressFromLastValidatorPowerKey(iterator.Key())
		validator := k.mustGetValidator(ctx, address)

		validators[i] = validator
		i++
	}

	return validators[:i] // trim
}

// GetUnbondingValidators returns a slice of mature validator addresses that
// complete their unbonding at a given time and height.
func (k Keeper) GetUnbondingValidators(ctx sdk.Context, endTime time.Time, endHeight int64) []string {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetValidatorQueueKey(endTime, endHeight))
	if bz == nil {
		return []string{}
	}

	addrs := types.ValAddresses{}
	k.cdc.MustUnmarshal(bz, &addrs)

	return addrs.Addresses
}

// SetUnbondingValidatorsQueue sets a given slice of validator addresses into
// the unbonding validator queue by a given height and time.
func (k Keeper) SetUnbondingValidatorsQueue(ctx sdk.Context, endTime time.Time, endHeight int64, addrs []string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.ValAddresses{Addresses: addrs})
	store.Set(types.GetValidatorQueueKey(endTime, endHeight), bz)
}

// InsertUnbondingValidatorQueue inserts a given unbonding validator address into
// the unbonding validator queue for a given height and time.
func (k Keeper) InsertUnbondingValidatorQueue(ctx sdk.Context, val types.Validator) {
	addrs := k.GetUnbondingValidators(ctx, val.UnbondingTime, val.UnbondingHeight)
	addrs = append(addrs, val.OperatorAddress)
	k.SetUnbondingValidatorsQueue(ctx, val.UnbondingTime, val.UnbondingHeight, addrs)
}

// DeleteValidatorQueueTimeSlice deletes all entries in the queue indexed by a
// given height and time.
func (k Keeper) DeleteValidatorQueueTimeSlice(ctx sdk.Context, endTime time.Time, endHeight int64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorQueueKey(endTime, endHeight))
}

// DeleteValidatorQueue removes a validator by address from the unbonding queue
// indexed by a given height and time.
func (k Keeper) DeleteValidatorQueue(ctx sdk.Context, val types.Validator) {
	addrs := k.GetUnbondingValidators(ctx, val.UnbondingTime, val.UnbondingHeight)
	newAddrs := []string{}

	for _, addr := range addrs {
		if addr != val.OperatorAddress {
			newAddrs = append(newAddrs, addr)
		}
	}

	if len(newAddrs) == 0 {
		k.DeleteValidatorQueueTimeSlice(ctx, val.UnbondingTime, val.UnbondingHeight)
	} else {
		k.SetUnbondingValidatorsQueue(ctx, val.UnbondingTime, val.UnbondingHeight, newAddrs)
	}
}

// ValidatorQueueIterator returns an interator ranging over validators that are
// unbonding whose unbonding completion occurs at the given height and time.
func (k Keeper) ValidatorQueueIterator(ctx sdk.Context, endTime time.Time, endHeight int64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.GetAllValidatorQueueKey(), sdk.InclusiveEndBytes(types.GetValidatorQueueKey(endTime, endHeight)))
}

// UnbondAllMatureValidators unbonds all the mature unbonding validators that
// have finished their unbonding period.
func (k Keeper) UnbondAllMatureValidators(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	blockTime := ctx.BlockTime()
	blockHeight := ctx.BlockHeight()

	// unbondingValIterator will contains all validator addresses indexed under
	// the ValidatorQueueKey prefix. Note, the entire index key is composed as
	// ValidatorQueueKey | timeBzLen (8-byte big endian) | timeBz | heightBz (8-byte big endian),
	// so it may be possible that certain validator addresses that are iterated
	// over are not ready to unbond, so an explicit check is required.
	unbondingValIterator := k.ValidatorQueueIterator(ctx, blockTime, blockHeight)
	defer unbondingValIterator.Close()

	for ; unbondingValIterator.Valid(); unbondingValIterator.Next() {
		key := unbondingValIterator.Key()
		keyTime, keyHeight, err := types.ParseValidatorQueueKey(key)
		if err != nil {
			panic(fmt.Errorf("failed to parse unbonding key: %w", err))
		}

		// All addresses for the given key have the same unbonding height and time.
		// We only unbond if the height and time are less than the current height
		// and time.
		if keyHeight <= blockHeight && (keyTime.Before(blockTime) || keyTime.Equal(blockTime)) {
			addrs := types.ValAddresses{}
			k.cdc.MustUnmarshal(unbondingValIterator.Value(), &addrs)

			for _, valAddr := range addrs.Addresses {
				addr, err := sdk.ValAddressFromBech32(valAddr)
				if err != nil {
					panic(err)
				}
				val, found := k.GetValidator(ctx, addr)
				if !found {
					panic("validator in the unbonding queue was not found")
				}

				if !val.IsUnbonding() {
					panic("unexpected validator in unbonding queue; status was not unbonding")
				}

				//val = k.UnbondingToUnbonded(ctx, val)
				//if val.GetDelegatorShares().IsZero() {
				//	k.RemoveValidator(ctx, val.GetOperator())
				//}
			}

			store.Delete(key)
		}
	}
}

func (k Keeper) mustGetValidator(ctx sdk.Context, addr sdk.ValAddress) types.Validator {
	validator, found := k.GetValidator(ctx, addr)
	if !found {
		panic(fmt.Sprintf("validator record not found for address: %s\n", addr.String()))
	}
	return validator
}

func (k Keeper) mustGetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) types.Validator {
	validator, found := k.GetValidatorByConsAddrDecimal(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}
	return validator
}
