package types

import (
	"bitbucket.org/decimalteam/go-smart-node/x/coin/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"regexp"
)

var (
	_ sdk.Msg = &MsgCreateCoin{}
	_ sdk.Msg = &MsgUpdateCoin{}
	_ sdk.Msg = &MsgSendCoin{}
	_ sdk.Msg = &MsgMultiSendCoin{}
	_ sdk.Msg = &MsgBuyCoin{}
	_ sdk.Msg = &MsgSellCoin{}
	_ sdk.Msg = &MsgSellAllCoin{}
	_ sdk.Msg = &MsgBurnCoin{}
	_ sdk.Msg = &MsgRedeemCheck{}
)

const (
	TypeMsgCreateCoin    = "create_coin"
	TypeMsgUpdateCoin    = "update_coin"
	TypeMsgSendCoin      = "send_coin"
	TypeMsgMultiSendCoin = "multi_send_coin"
	TypeMsgBuyCoin       = "buy_coin"
	TypeMsgSellCoin      = "sell_coin"
	TypeMsgSellAllCoin   = "sell_all_coin"
	TypeMsgBurnCoin      = "burn_coin"
	TypeMsgRedeemCheck   = "redeem_check"
)

////////////////////////////////////////////////////////////////
// MsgCreateCoin
////////////////////////////////////////////////////////////////

// NewMsgCreateCoin creates a new instance of MsgCreateCoin.
func NewMsgCreateCoin(
	sender sdk.AccAddress,
	title string,
	symbol string,
	crr uint64,
	initVolume sdk.Int,
	initReserve sdk.Int,
	limitVolume sdk.Int,
	identity string,
) *MsgCreateCoin {
	return &MsgCreateCoin{
		Sender:         sender.String(),
		Title:          title,
		Symbol:         symbol,
		CRR:            crr,
		InitialVolume:  initVolume,
		InitialReserve: initReserve,
		LimitVolume:    limitVolume,
		Identity:       identity,
	}
}

// Route should return the name of the module.
func (msg MsgCreateCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgCreateCoin) Type() string { return TypeMsgCreateCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgCreateCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgCreateCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgCreateCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Validate coin title
	if len(msg.Title) > maxCoinNameBytes {
		return errors.InvalidCoinTitle
	}
	// Validate coin symbol
	if match, _ := regexp.MatchString(allowedCoinSymbols, msg.Symbol); !match {
		return errors.InvalidCoinSymbol
	}
	// TODO: Looks like no need since should be no more possible to create such coins anyway
	// // Forbid creating coin with symbol DEL in testnet
	// if strings.HasPrefix(config.ChainID, "decimal-testnet") {
	// 	if strings.ToLower(msg.Symbol) == config.SymbolBaseCoin {
	// 		return ErrForbiddenCoinSymbol(msg.Symbol)
	// 	}
	// }
	// Validate coin CRR
	if msg.CRR < 10 || msg.CRR > 100 {
		return errors.InvalidCRR
	}
	// Check coin initial volume to be correct
	if msg.InitialVolume.LT(MinCoinSupply) || msg.InitialVolume.GT(MaxCoinSupply) {
		return errors.InvalidCoinInitialVolume
	}
	if msg.InitialVolume.GT(msg.LimitVolume) {
		return errors.InvalidLimitVolume
	}
	// Check coin initial reserve to be correct
	if msg.InitialReserve.LT(MinCoinReserve) {
		return errors.InvalidCoinInitialReserve
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgUpdateCoin
////////////////////////////////////////////////////////////////

// NewMsgUpdateCoin creates a new instance of MsgUpdateCoin.
func NewMsgUpdateCoin(
	sender sdk.AccAddress,
	symbol string,
	limitVolume sdk.Int,
	identity string,
) *MsgUpdateCoin {
	return &MsgUpdateCoin{
		Sender:      sender.String(),
		Symbol:      symbol,
		LimitVolume: limitVolume,
		Identity:    identity,
	}
}

// Route should return the name of the module.
func (msg MsgUpdateCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgUpdateCoin) Type() string { return TypeMsgUpdateCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgUpdateCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgUpdateCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgUpdateCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Validate coin symbol
	if match, _ := regexp.MatchString(allowedCoinSymbols, msg.Symbol); !match {
		return errors.InvalidCoinSymbol
	}
	// Check coin limit volume to be less than max coin supply
	if msg.LimitVolume.GT(MaxCoinSupply) {
		return errors.InvalidLimitVolume
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgSendCoin
////////////////////////////////////////////////////////////////

// NewMsgSendCoin creates a new instance of MsgSendCoin.
func NewMsgSendCoin(
	sender sdk.AccAddress,
	coin sdk.Coin,
	receiver sdk.AccAddress,
) *MsgSendCoin {
	return &MsgSendCoin{
		Sender:   sender.String(),
		Coin:     coin,
		Receiver: receiver.String(),
	}
}

// Route should return the name of the module.
func (msg MsgSendCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgSendCoin) Type() string { return TypeMsgSendCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgSendCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgSendCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgSendCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Validate receiver
	if _, err := sdk.AccAddressFromBech32(msg.Receiver); err != nil {
		return errors.InvalidReceiverAddress
	}
	// Amount should be positive
	if msg.Coin.Amount.LT(sdk.NewInt(0)) {
		return errors.InvalidAmount
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgMultiSendCoin
////////////////////////////////////////////////////////////////

// NewMsgMultiSendCoin creates a new instance of MsgMultiSendCoin.
func NewMsgMultiSendCoin(
	sender sdk.AccAddress,
	sends []Send,
) *MsgMultiSendCoin {
	return &MsgMultiSendCoin{
		Sender: sender.String(),
		Sends:  sends,
	}
}

// Route should return the name of the module.
func (msg MsgMultiSendCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgMultiSendCoin) Type() string { return TypeMsgMultiSendCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgMultiSendCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgMultiSendCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgMultiSendCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	for i := range msg.Sends {
		// Validate receiver
		if _, err := sdk.AccAddressFromBech32(msg.Sends[i].Receiver); err != nil {
			return errors.InvalidReceiverAddress
		}
		// Amount should be positive
		if msg.Sends[i].Coin.Amount.LT(sdk.NewInt(0)) {
			return errors.InvalidAmount
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgBuyCoin
////////////////////////////////////////////////////////////////

// NewMsgBuyCoin creates a new instance of MsgBuyCoin.
func NewMsgBuyCoin(
	sender sdk.AccAddress,
	coinToBuy sdk.Coin,
	maxCoinToSell sdk.Coin,
) *MsgBuyCoin {
	return &MsgBuyCoin{
		Sender:        sender.String(),
		CoinToBuy:     coinToBuy,
		MaxCoinToSell: maxCoinToSell,
	}
}

// Route should return the name of the module.
func (msg MsgBuyCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgBuyCoin) Type() string { return TypeMsgBuyCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgBuyCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgBuyCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgBuyCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Denoms of specified coins cannot be the same
	if msg.CoinToBuy.Denom == msg.MaxCoinToSell.Denom {
		return errors.SameCoin
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgSellCoin
////////////////////////////////////////////////////////////////

// NewMsgSellCoin creates a new instance of MsgSellCoin.
func NewMsgSellCoin(
	sender sdk.AccAddress,
	coinToSell sdk.Coin,
	minCoinToBuy sdk.Coin,
) *MsgSellCoin {
	return &MsgSellCoin{
		Sender:       sender.String(),
		CoinToSell:   coinToSell,
		MinCoinToBuy: minCoinToBuy,
	}
}

// Route should return the name of the module.
func (msg MsgSellCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgSellCoin) Type() string { return TypeMsgSellCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgSellCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgSellCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgSellCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Denoms of specified coins cannot be the same
	if msg.CoinToSell.Denom == msg.MinCoinToBuy.Denom {
		return errors.SameCoin
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgSellAllCoin
////////////////////////////////////////////////////////////////

// NewMsgSellAllCoin creates a new instance of MsgSellAllCoin.
func NewMsgSellAllCoin(
	sender sdk.AccAddress,
	coinToSell sdk.Coin,
	minCoinToBuy sdk.Coin,
) *MsgSellAllCoin {
	return &MsgSellAllCoin{
		Sender:       sender.String(),
		CoinToSell:   coinToSell,
		MinCoinToBuy: minCoinToBuy,
	}
}

// Route should return the name of the module.
func (msg MsgSellAllCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgSellAllCoin) Type() string { return TypeMsgSellAllCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgSellAllCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgSellAllCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgSellAllCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Denoms of specified coins cannot be the same
	if msg.CoinToSell.Denom == msg.MinCoinToBuy.Denom {
		return errors.SameCoin
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgBurnCoin
////////////////////////////////////////////////////////////////

// NewMsgSendCoin creates a new instance of MsgSendCoin.
func NewMsgBurnCoin(
	sender sdk.AccAddress,
	coin sdk.Coin,
) *MsgBurnCoin {
	return &MsgBurnCoin{
		Sender: sender.String(),
		Coin:   coin,
	}
}

// Route should return the name of the module.
func (msg MsgBurnCoin) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgBurnCoin) Type() string { return TypeMsgSendCoin }

// GetSignBytes encodes the message for signing.
func (msg *MsgBurnCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgBurnCoin) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgBurnCoin) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// Amount should be positive
	if !msg.Coin.Amount.IsPositive() {
		return errors.InvalidAmount
	}
	return nil
}

////////////////////////////////////////////////////////////////
// MsgRedeemCheck
////////////////////////////////////////////////////////////////

// NewMsgRedeemCheck creates a new instance of MsgRedeemCheck.
func NewMsgRedeemCheck(
	sender sdk.AccAddress,
	check string,
	proof string,
) *MsgRedeemCheck {
	return &MsgRedeemCheck{
		Sender: sender.String(),
		Check:  check,
		Proof:  proof,
	}
}

// Route should return the name of the module.
func (msg MsgRedeemCheck) Route() string { return RouterKey }

// Type should return the action.
func (msg MsgRedeemCheck) Type() string { return TypeMsgRedeemCheck }

// GetSignBytes encodes the message for signing.
func (msg *MsgRedeemCheck) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required.
func (msg MsgRedeemCheck) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs stateless checks on the message.
func (msg MsgRedeemCheck) ValidateBasic() error {
	// Validate sender
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.InvalidSenderAddress
	}
	// TODO
	return nil
}
