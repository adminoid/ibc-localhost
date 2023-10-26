package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"bitbucket.org/decimalteam/go-smart-node/x/validator/types"
)

func RegisterInvariants(registry sdk.InvariantRegistry, k Keeper) {
	registry.RegisterRoute(types.ModuleName, "staked-custom-coins", StakedCustomCoinInvariant(k))
	registry.RegisterRoute(types.ModuleName, "delegations-counter", DelegationsCounterInvariant(k))
}

// StakedCustomCoinInvariant checks that helper amounts of custom coins is equal
// to sum of delegation/undelegations/redelegations
func StakedCustomCoinInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allCoins := sdk.NewCoins()
		k.IterateAllDelegations(ctx, func(delegation types.Delegation) bool {
			if delegation.Stake.Stake.Denom == k.BaseDenom(ctx) {
				return false
			}
			allCoins = allCoins.Add(delegation.Stake.GetStake())
			return false
		})
		k.IterateUndelegations(ctx, func(index int64, ubd types.Undelegation) bool {
			for _, entry := range ubd.Entries {
				if entry.Stake.Stake.Denom == k.BaseDenom(ctx) {
					continue
				}
				allCoins = allCoins.Add(entry.Stake.GetStake())
			}
			return false
		})
		k.IterateRedelegations(ctx, func(index int64, red types.Redelegation) bool {
			for _, entry := range red.Entries {
				if entry.Stake.Stake.Denom == k.BaseDenom(ctx) {
					continue
				}
				allCoins = allCoins.Add(entry.Stake.GetStake())
			}
			return false
		})
		broken := false
		diff := []string{}
		for denom, amount := range k.GetAllCustomCoinsStaked(ctx) {
			if !allCoins.AmountOf(denom).Equal(amount) {
				diff = append(diff, fmt.Sprintf("coin: %s, (re/un)delegations: %s, staked: %s",
					denom, allCoins.AmountOf(denom), amount))
				broken = true
			}
		}
		ccs := k.GetAllCustomCoinsStaked(ctx)
		for _, coin := range allCoins {
			_, ok := ccs[coin.Denom]
			if !ok {
				diff = append(diff, fmt.Sprintf("coin '%s' is not found in staked custom coin", coin.Denom))
				broken = true
			}
		}

		description := strings.Join(diff, "\n")

		// Bonded tokens should equal sum of tokens with bonded validators
		// Not-bonded tokens should equal unbonding delegations	plus tokens on unbonded validators
		return sdk.FormatInvariant(types.ModuleName, "staked custom coins", "\tDifferences:\n"+description), broken
	}
}

func DelegationsCounterInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		counter1 := k.GetAllDelegationsCount(ctx)
		counter2 := make(map[string]uint32)
		k.IterateAllDelegations(ctx, func(delegation types.Delegation) bool {
			counter2[delegation.Validator] = counter2[delegation.Validator] + 1
			return false
		})
		vals := k.GetAllValidators(ctx)
		diff := []string{}
		broken := false
		for _, val := range vals {
			if counter2[val.OperatorAddress] != counter1[val.OperatorAddress] {
				broken = true
				diff = append(diff, fmt.Sprintf("validator: %s, stored count: %d, calculated count: %d",
					val.OperatorAddress, counter1[val.OperatorAddress], counter2[val.OperatorAddress],
				))
			}
		}
		description := strings.Join(diff, "\n")
		return sdk.FormatInvariant(types.ModuleName, "delegations counter", "\tDifferences:\n"+description), broken
	}
}

//
//// RegisterInvariants registers all the module's invariants.
//func RegisterInvariants(registry sdk.InvariantRegistry, k Keeper) {
//	registry.RegisterRoute(types.ModuleName, "module-accounts", ModuleAccountInvariants(k))
//	registry.RegisterRoute(types.ModuleName, "nonnegative-power", NonNegativePowerInvariant(k))
//	registry.RegisterRoute(types.ModuleName, "positive-delegation", PositiveDelegationInvariant(k))
//	registry.RegisterRoute(types.ModuleName, "delegator-shares", DelegatorSharesInvariant(k))
//}
//
//// AllInvariants runs all invariants of the module.
//func AllInvariants(k Keeper) sdk.Invariant {
//	return func(ctx sdk.Context) (string, bool) {
//		res, stop := ModuleAccountInvariants(k)(ctx)
//		if stop {
//			return res, stop
//		}
//
//		res, stop = NonNegativePowerInvariant(k)(ctx)
//		if stop {
//			return res, stop
//		}
//
//		res, stop = PositiveDelegationInvariant(k)(ctx)
//		if stop {
//			return res, stop
//		}
//
//		return DelegatorSharesInvariant(k)(ctx)
//	}
//}
//
//// ModuleAccountInvariants checks that the bonded and notBonded ModuleAccounts pools
//// reflects the tokens actively bonded and not bonded
//func ModuleAccountInvariants(k Keeper) sdk.Invariant {
//	return func(ctx sdk.Context) (string, bool) {
//		bonded := sdk.ZeroInt()
//		notBonded := sdk.ZeroInt()
//		bondedPool := k.GetBondedPool(ctx)
//		notBondedPool := k.GetNotBondedPool(ctx)
//		bondDenom := k.BondDenom(ctx)
//
//		k.IterateValidators(ctx, func(_ int64, validator types.ValidatorI) bool {
//			switch validator.GetStatus() {
//			case types.Bonded:
//				bonded = bonded.Add(validator.GetTokens())
//			case types.Unbonding, types.Unbonded:
//				notBonded = notBonded.Add(validator.GetTokens())
//			default:
//				panic("invalid validator status")
//			}
//			return false
//		})
//
//		k.IterateUndelegations(ctx, func(_ int64, ubd types.Undelegation) bool {
//			for _, entry := range ubd.Entries {
//				notBonded = notBonded.Add(entry.Balance)
//			}
//			return false
//		})
//
//		poolBonded := k.bankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom)
//		poolNotBonded := k.bankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom)
//		broken := !poolBonded.Amount.Equal(bonded) || !poolNotBonded.Amount.Equal(notBonded)
//
//		// Bonded tokens should equal sum of tokens with bonded validators
//		// Not-bonded tokens should equal unbonding delegations	plus tokens on unbonded validators
//		return sdk.FormatInvariant(types.ModuleName, "bonded and not bonded module account coins", fmt.Sprintf(
//			"\tPool's bonded tokens: %v\n"+
//				"\tsum of bonded tokens: %v\n"+
//				"not bonded token invariance:\n"+
//				"\tPool's not bonded tokens: %v\n"+
//				"\tsum of not bonded tokens: %v\n"+
//				"module accounts total (bonded + not bonded):\n"+
//				"\tModule Accounts' tokens: %v\n"+
//				"\tsum tokens:              %v\n",
//			poolBonded, bonded, poolNotBonded, notBonded, poolBonded.Add(poolNotBonded), bonded.Add(notBonded))), broken
//	}
//}
//
//// NonNegativePowerInvariant checks that all stored validators have >= 0 power.
//func NonNegativePowerInvariant(k Keeper) sdk.Invariant {
//	return func(ctx sdk.Context) (string, bool) {
//		var (
//			msg    string
//			broken bool
//		)
//
//		iterator := k.ValidatorsPowerStoreIterator(ctx)
//		for ; iterator.Valid(); iterator.Next() {
//			validator, found := k.GetValidator(ctx, iterator.Value())
//			if !found {
//				panic(fmt.Sprintf("validator record not found for address: %X\n", iterator.Value()))
//			}
//
//			powerKey := k.GetValidatorsByPowerIndexKey(ctx, validator)
//
//			if !bytes.Equal(iterator.Key(), powerKey) {
//				broken = true
//				msg += fmt.Sprintf("power store invariance:\n\tvalidator.Power: %v"+
//					"\n\tkey should be: %v\n\tkey in store: %v\n",
//					validator.GetConsensusPower(ethtypes.PowerReduction), powerKey, iterator.Key())
//			}
//
//			if validator.Tokens.IsNegative() {
//				broken = true
//				msg += fmt.Sprintf("\tnegative tokens for validator: %v\n", validator)
//			}
//		}
//		iterator.Close()
//
//		return sdk.FormatInvariant(types.ModuleName, "nonnegative power", fmt.Sprintf("found invalid validator powers\n%s", msg)), broken
//	}
//}
//
//// PositiveDelegationInvariant checks that all stored delegations have > 0 shares.
//func PositiveDelegationInvariant(k Keeper) sdk.Invariant {
//	return func(ctx sdk.Context) (string, bool) {
//		var (
//			msg   string
//			count int
//		)
//
//		delegations := k.GetAllDelegations(ctx)
//		for _, delegation := range delegations {
//			if delegation.Shares.IsNegative() {
//				count++
//				msg += fmt.Sprintf("\tdelegation with negative shares: %+v\n", delegation)
//			}
//
//			if delegation.Shares.IsZero() {
//				count++
//				msg += fmt.Sprintf("\tdelegation with zero shares: %+v\n", delegation)
//			}
//		}
//
//		broken := count != 0
//
//		return sdk.FormatInvariant(types.ModuleName, "positive delegations", fmt.Sprintf(
//			"%d invalid delegations found\n%s", count, msg)), broken
//	}
//}
//
//// DelegatorSharesInvariant checks whether all the delegator shares which persist
//// in the delegator object add up to the correct total delegator shares
//// amount stored in each validator.
//func DelegatorSharesInvariant(k Keeper) sdk.Invariant {
//	return func(ctx sdk.Context) (string, bool) {
//		var (
//			msg    string
//			broken bool
//		)
//
//		validators := k.GetAllValidators(ctx)
//		validatorsDelegationShares := map[string]sdk.Dec{}
//
//		// initialize a map: validator -> its delegation shares
//		for _, validator := range validators {
//			validatorsDelegationShares[validator.GetOperator().String()] = sdk.ZeroDec()
//		}
//
//		// iterate through all the delegations to calculate the total delegation shares for each validator
//		delegations := k.GetAllDelegations(ctx)
//		for _, delegation := range delegations {
//			delegationValidatorAddr := delegation.GetValidatorAddr().String()
//			validatorDelegationShares := validatorsDelegationShares[delegationValidatorAddr]
//			validatorsDelegationShares[delegationValidatorAddr] = validatorDelegationShares.Add(delegation.Shares)
//		}
//
//		// for each validator, check if its total delegation shares calculated from the step above equals to its expected delegation shares
//		for _, validator := range validators {
//			expValTotalDelShares := validator.GetDelegatorShares()
//			calculatedValTotalDelShares := validatorsDelegationShares[validator.GetOperator().String()]
//			if !calculatedValTotalDelShares.Equal(expValTotalDelShares) {
//				broken = true
//				msg += fmt.Sprintf("broken delegator shares invariance:\n"+
//					"\tvalidator.DelegatorShares: %v\n"+
//					"\tsum of Delegator.Shares: %v\n", expValTotalDelShares, calculatedValTotalDelShares)
//			}
//		}
//
//		return sdk.FormatInvariant(types.ModuleName, "delegator shares", msg), broken
//	}
//}
