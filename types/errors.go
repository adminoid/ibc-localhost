package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/errors"
)

// RootCodespace is the codespace for all errors defined in the project.
const RootCodespace = "decimal"

// root error codes for Decimal
const (
	codeKeyTypeNotSupported = iota + 2
)

// errors
var (
	ErrKeyTypeNotSupported = sdkerrors.New(RootCodespace, codeKeyTypeNotSupported, "key type 'secp256k1' not supported")
)
