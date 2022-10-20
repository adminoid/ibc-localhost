package errors

import (
	"cosmossdk.io/errors"
)

var codespace = "validator"

var (
	Internal                        = errors.New(codespace, 101, "internal error")
	ValidatorAlreadyExists          = errors.New(codespace, 102, "validator already exists")
	InvalidConsensusPubKey          = errors.New(codespace, 103, "invalid consensus public key")
	ValidatorPublicKeyAlreadyExists = errors.New(codespace, 104, "validator public key already exists")
	UnsupportedPubKeyType           = errors.New(codespace, 105, "unsupported public key type")
	ValidatorNotFound               = errors.New(codespace, 106, "validator not found")
	ValidatorAlreadyOnline          = errors.New(codespace, 107, "validator already online")
	ValidatorAlreadyOffline         = errors.New(codespace, 108, "validator already offline")
	NFTSubTokenNotFound             = errors.New(codespace, 109, "NFT subtoken does not exists")
	DelegatorIsNotOwnerOfSubtoken   = errors.New(codespace, 110, "delegator is not owner of NFT subtoken")
	NFTTokenNotFound                = errors.New(codespace, 112, "NFT does not exists")
	DelegationNotFound              = errors.New(codespace, 113, "delegation not found")
	DelegationWrongType             = errors.New(codespace, 114, "delegation has wrong type")
	StakeTooSmall                   = errors.New(codespace, 116, "stake too small")
	SubTokenIDsDublicates           = errors.New(codespace, 117, "subtokes ID set has dublicates")
	StakeDoesNotHaveSubTokenID      = errors.New(codespace, 118, "stake does not have subtoken id")
	DelegateSubTokenTwice           = errors.New(codespace, 119, "trying to delegate subtoken id twice")
	BadRedelegationDst              = errors.New(codespace, 120, "redelegation destination validator not found")
	BadRedelegationSrc              = errors.New(codespace, 121, "redelegation source validator not found")
	SelfRedelegation                = errors.New(codespace, 122, "cannot redelegate to the same validator")
	TransitiveRedelegation          = errors.New(codespace, 123, "redelegation to this validator already in progress; first redelegation to this validator must complete before next redelegation")
	MaxRedelegationEntries          = errors.New(codespace, 124, "too many redelegation entries for (delegator, src-validator, dst-validator) tuple")
	IncompatibleBondStatuses        = errors.New(codespace, 125, "incompatible bond statuses")
	ValidatorStatusUnknown          = errors.New(codespace, 126, "validator status unknown")
	WrongStakeType                  = errors.New(codespace, 127, "wrong stake type")
	SubTokenExistsInStake           = errors.New(codespace, 128, "subtoken exists in stake")
	MaxUndelegationEntries          = errors.New(codespace, 129, "too many unbonding delegation entries for (delegator, validator) tuple")
	ValidatorJailed                 = errors.New(codespace, 130, "validator for this address is currently jailed")
	UBDNotFound                     = errors.New(codespace, 131, "undelegation not found for pair (validator, delegator)")
	UBDEntryNotFound                = errors.New(codespace, 132, "undelegation entry not found for height")
	UBDAlreadyProcessed             = errors.New(codespace, 133, "undelegation is already processed")
	WrongStakeID                    = errors.New(codespace, 134, "wrong stake id")
	REDNotFound                     = errors.New(codespace, 135, "redelegation not found for pair (validatorSrc, validatorDst, delegator)")
	REDEntryNotFound                = errors.New(codespace, 136, "redelegation entry not found for height")
	REDAlreadyProcessed             = errors.New(codespace, 137, "redelegation is already processed")
	CoinDoesNotExist                = errors.New(codespace, 138, "coin does not exist")
	RewardsNotFound                 = errors.New(codespace, 139, "not found rewards for validator")
)
