package types

import (
	"sort"

	"github.com/gogo/protobuf/proto"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/adminoid/cosmos-sdk/codec"
	codectypes "github.com/adminoid/cosmos-sdk/codec/types"
	sdkerrors "github.com/adminoid/cosmos-sdk/types/errors"
)

// NewHistoricalInfo will create a historical information struct from header and validator set.
// It will first sort validator set before inclusion into historical info.
func NewHistoricalInfo(header tmproto.Header, valSet Validators) HistoricalInfo {
	// Must sort in the same way that tendermint does
	sort.SliceStable(valSet, func(i, j int) bool {
		return ValidatorsByVotingPower(valSet).Less(i, j)
	})

	// pruning unused data (for staking.Validator), see keeper.GetHistoricalInfo
	var prunedValSet Validators
	for _, v := range valSet {
		prunedValSet = append(prunedValSet, Validator{
			OperatorAddress: v.OperatorAddress,
			ConsensusPubkey: v.ConsensusPubkey,
			Status:          v.Status,
			Jailed:          v.Jailed,
			Stake:           v.Stake,
		})
	}

	return HistoricalInfo{
		Header: header,
		Valset: prunedValSet,
	}
}

// MustUnmarshalHistoricalInfo will unmarshal historical info and panic on error.
func MustUnmarshalHistoricalInfo(cdc codec.BinaryCodec, value []byte) HistoricalInfo {
	hi, err := UnmarshalHistoricalInfo(cdc, value)
	if err != nil {
		panic(err)
	}

	return hi
}

// UnmarshalHistoricalInfo will unmarshal historical info and return any error.
func UnmarshalHistoricalInfo(cdc codec.BinaryCodec, value []byte) (hi HistoricalInfo, err error) {
	err = cdc.Unmarshal(value, &hi)
	return hi, err
}

// ValidateBasic will ensure HistoricalInfo is not nil and sorted.
func ValidateBasic(hi HistoricalInfo) error {
	if len(hi.Valset) == 0 {
		return sdkerrors.Wrap(ErrInvalidHistoricalInfo, "validator set is empty")
	}

	if !sort.IsSorted(Validators(hi.Valset)) {
		return sdkerrors.Wrap(ErrInvalidHistoricalInfo, "validator set is not sorted by address")
	}

	return nil
}

// Equal checks if receiver is equal to the parameter.
func (hi *HistoricalInfo) Equal(hi2 *HistoricalInfo) bool {
	if !proto.Equal(&hi.Header, &hi2.Header) {
		return false
	}
	if len(hi.Valset) != len(hi2.Valset) {
		return false
	}
	for i := range hi.Valset {
		if !hi.Valset[i].Equal(&hi2.Valset[i]) {
			return false
		}
	}
	return true
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces.
func (hi HistoricalInfo) UnpackInterfaces(c codectypes.AnyUnpacker) error {
	for i := range hi.Valset {
		if err := hi.Valset[i].UnpackInterfaces(c); err != nil {
			return err
		}
	}
	return nil
}
