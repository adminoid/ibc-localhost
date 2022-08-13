package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"bitbucket.org/decimalteam/go-smart-node/x/swap/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for the module.
func GetTxCmd() *cobra.Command {
	coinCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	coinCmd.AddCommand(
		NewSwapInitializeCmd(),
		NewSwapRedeemCmd(),
		NewChainActivateCmd(),
		NewChainDeactivateCmd(),
	)

	return coinCmd
}

func NewSwapInitializeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [recipient] [amount] [token_symbol] [tx_number] [from_chain] [dest_chain]",
		Short: "Swap initialize",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			recipient := args[0]
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount")
			}
			tokenSymbol := args[2]
			txNumber := args[3]
			fromChain, err := strconv.ParseUint(args[4], 10, 32)
			if err != nil {
				return err
			}
			destChain, err := strconv.ParseUint(args[5], 10, 32)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwapInitialize(from, recipient, amount, tokenSymbol, txNumber, uint32(fromChain), uint32(destChain))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	// workaround for cosmos
	cmd.Flags().String(flags.FlagChainID, "", "network chain id")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func NewSwapRedeemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem [from] [recipient] [amount] [token_symbol] [tx_number] [from_chain] [dest_chain] [v] [r] [s]",
		Short: "Swap redeem",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			from := args[0]
			recipient := args[1]
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid amount")
			}
			tokenSymbol := args[3]
			txNumber := args[4]
			fromChain, err := strconv.ParseUint(args[5], 10, 32)
			if err != nil {
				return err
			}
			destChain, err := strconv.ParseUint(args[6], 10, 32)
			if err != nil {
				return err
			}

			v, err := strconv.Atoi(args[7])
			if err != nil {
				return err
			}

			r, err := hex.DecodeString(args[8])
			if err != nil {
				return err
			}
			s, err := hex.DecodeString(args[9])
			if err != nil {
				return err
			}
			var _r types.Hash
			copy(_r[:], r)

			var _s types.Hash
			copy(_s[:], s)

			msg := types.NewMsgSwapRedeem(
				sender, from, recipient, amount, tokenSymbol, txNumber, uint32(fromChain), uint32(destChain), uint32(v), &_r, &_s)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	// workaround for cosmos
	cmd.Flags().String(flags.FlagChainID, "", "network chain id")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func NewChainActivateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-activate [number] [name]",
		Short: "Activate chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			number, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			name := args[1]

			msg := types.NewMsgChainActivate(sender, uint32(number), name)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	// workaround for cosmos
	cmd.Flags().String(flags.FlagChainID, "", "network chain id")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func NewChainDeactivateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-deactivate [number]",
		Short: "Deactivate chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			number, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			msg := types.NewMsgChainDeactivate(sender, uint32(number))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	// workaround for cosmos
	cmd.Flags().String(flags.FlagChainID, "", "network chain id")

	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}