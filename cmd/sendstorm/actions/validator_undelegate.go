package actions

import (
	"math/rand"
	"time"

	stormTypes "bitbucket.org/decimalteam/go-smart-node/cmd/sendstorm/types"
	dscTx "bitbucket.org/decimalteam/go-smart-node/sdk/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type UndelegateGenerator struct {
	knownStakes []GenericStake
	rnd         *rand.Rand
}

type UndelegateAction struct {
	coin             sdk.Coin
	delegatorAddress string
	validatorAddress string
}

func NewUndelegateGenerator() *UndelegateGenerator {
	return &UndelegateGenerator{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (gg *UndelegateGenerator) Update(ui UpdateInfo) {
	gg.knownStakes = ui.Stakes
}

func (gg *UndelegateGenerator) Generate() Action {
	if len(gg.knownStakes) == 0 {
		return &EmptyAction{}
	}
	stake := RandomChoice(gg.rnd, gg.knownStakes)
	return &UndelegateAction{
		coin:             stake.Coin,
		delegatorAddress: stake.Delegator,
		validatorAddress: stake.Validator,
	}
}

func (ac *UndelegateAction) ChooseAccounts(saList []*stormTypes.StormAccount) []*stormTypes.StormAccount {
	var res []*stormTypes.StormAccount
	for i := range saList {
		if saList[i].IsDirty() {
			continue
		}
		if saList[i].Address() != ac.delegatorAddress {
			continue
		}
		res = append(res, saList[i])
	}
	return res
}

func (ac *UndelegateAction) GenerateTx(sa *stormTypes.StormAccount, feeConfig *stormTypes.FeeConfiguration) ([]byte, error) {
	_, err := sdk.AccAddressFromBech32(sa.Address())
	if err != nil {
		return nil, err
	}

	// TODO

	tx, err := dscTx.BuildTransaction(sa.Account(), []sdk.Msg{}, "", sa.FeeDenom(), feeConfig.DelPrice, feeConfig.Params)
	if err != nil {
		return nil, err
	}
	err = tx.SignTransaction(sa.Account())
	if err != nil {
		return nil, err
	}
	return tx.BytesToSend()
}

func (ac *UndelegateAction) String() string {
	return "UndelegateAction"
}
