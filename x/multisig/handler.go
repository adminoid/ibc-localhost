package multisig

import (
	"fmt"
	"runtime/debug"

	"bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler defines the module handler instance.
func NewHandler(server types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		// Defer hook to catch panic and log it
		defer func() {
			if r := recover(); r != nil {
				ctx.Logger().Error(fmt.Sprintf("stacktrace from panic: %s\n%s\n", r, string(debug.Stack())))
			}
		}()
		// Handle the message
		switch msg := msg.(type) {
		case *types.MsgCreateWallet:
			res, err := server.CreateWallet(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateTransaction:
			res, err := server.CreateTransaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSignTransaction:
			res, err := server.SignTransaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
