package types

import (
	sdk "github.com/adminoid/cosmos-sdk/types"
)

// query endpoints supported by the staking Querier
const (
	QueryValidators             = "validators"
	QueryValidator              = "validator"
	QueryDelegatorDelegations   = "delegatorDelegations"
	QueryDelegatorUndelegations = "delegatorUndelegations"
	QueryRedelegations          = "redelegations"
	QueryValidatorDelegations   = "validatorDelegations"
	QueryValidatorRedelegations = "validatorRedelegations"
	QueryValidatorUndelegations = "validatorUndelegations"
	QueryDelegation             = "delegation"
	QueryUndelegation           = "undelegation"
	QueryDelegatorValidators    = "delegatorValidators"
	QueryDelegatorValidator     = "delegatorValidator"
	QueryPool                   = "pool"
	QueryParameters             = "parameters"
	QueryHistoricalInfo         = "historicalInfo"
)

// defines the params for the following queries:
// - 'custom/staking/delegatorDelegations'
// - 'custom/staking/delegatorUndelegations'
// - 'custom/staking/delegatorValidators'
type QueryDelegatorParams struct {
	DelegatorAddr sdk.AccAddress
}

func NewQueryDelegatorParams(delegatorAddr sdk.AccAddress) QueryDelegatorParams {
	return QueryDelegatorParams{
		DelegatorAddr: delegatorAddr,
	}
}

// defines the params for the following queries:
// - 'custom/staking/validator'
// - 'custom/staking/validatorDelegations'
// - 'custom/staking/validatorUndelegations'
type QueryValidatorParams struct {
	ValidatorAddr sdk.ValAddress
	Page, Limit   int
}

func NewQueryValidatorParams(validatorAddr sdk.ValAddress, page, limit int) QueryValidatorParams {
	return QueryValidatorParams{
		ValidatorAddr: validatorAddr,
		Page:          page,
		Limit:         limit,
	}
}

// defines the params for the following queries:
// - 'custom/staking/redelegation'
type QueryRedelegationParams struct {
	DelegatorAddr    sdk.AccAddress
	SrcValidatorAddr sdk.ValAddress
	DstValidatorAddr sdk.ValAddress
}

func NewQueryRedelegationParams(delegatorAddr sdk.AccAddress, srcValidatorAddr, dstValidatorAddr sdk.ValAddress) QueryRedelegationParams {
	return QueryRedelegationParams{
		DelegatorAddr:    delegatorAddr,
		SrcValidatorAddr: srcValidatorAddr,
		DstValidatorAddr: dstValidatorAddr,
	}
}

// QueryValidatorsParams defines the params for the following queries:
// - 'custom/staking/validators'
type QueryValidatorsParams struct {
	Page, Limit int
	Status      string
}

func NewQueryValidatorsParams(page, limit int, status string) QueryValidatorsParams {
	return QueryValidatorsParams{page, limit, status}
}
