package tx

import (
	coinTypes "bitbucket.org/decimalteam/go-smart-node/x/coin/types"
	nftTypes "bitbucket.org/decimalteam/go-smart-node/x/nft/types"
)

type (
	MsgCreateCoin          = coinTypes.MsgCreateCoin
	MsgUpdateCoin          = coinTypes.MsgUpdateCoin
	MsgMultiSendCoin       = coinTypes.MsgMultiSendCoin
	MsgBuyCoin             = coinTypes.MsgBuyCoin
	MsgSellCoin            = coinTypes.MsgSellCoin
	MsgSellAllCoin         = coinTypes.MsgSellAllCoin
	MsgSendCoin            = coinTypes.MsgSendCoin
	MsgRedeemCheck         = coinTypes.MsgRedeemCheck
	MsgReturnLegacyBalance = coinTypes.MsgReturnLegacyBalance

	OneSend = coinTypes.Send

	MsgMintNFT          = nftTypes.MsgMintNFT
	MsgBurnNFT          = nftTypes.MsgBurnNFT
	MsgUpdateReserveNFT = nftTypes.MsgUpdateReserveNFT
	MsgTransferNFT      = nftTypes.MsgTransferNFT
	MsgEditNFTMetadata  = nftTypes.MsgEditNFTMetadata
)

var (
	NewMsgCreateCoin          = coinTypes.NewMsgCreateCoin
	NewMsgUpdateCoin          = coinTypes.NewMsgUpdateCoin
	NewMsgMultiSendCoin       = coinTypes.NewMsgMultiSendCoin
	NewMsgBuyCoin             = coinTypes.NewMsgBuyCoin
	NewMsgSellCoin            = coinTypes.NewMsgSellCoin
	NewMsgSellAllCoin         = coinTypes.NewMsgSellAllCoin
	NewMsgSendCoin            = coinTypes.NewMsgSendCoin
	NewMsgRedeemCheck         = coinTypes.NewMsgRedeemCheck
	NewMsgReturnLegacyBalance = coinTypes.NewMsgReturnLegacyBalance

	NewMsgMintNFT          = nftTypes.NewMsgMintNFT
	NewMsgBurnNFT          = nftTypes.NewMsgBurnNFT
	NewMsgUpdateReserveNFT = nftTypes.NewMsgUpdateReserveNFT
	NewMsgTransferNFT      = nftTypes.NewMsgTransferNFT
	NewMsgEditNFTMetadata  = nftTypes.NewMsgEditNFTMetadata
)
