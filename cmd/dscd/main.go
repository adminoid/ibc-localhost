package main

import (
	"os"

	"github.com/adminoid/cosmos-sdk/server"
	svrcmd "github.com/adminoid/cosmos-sdk/server/cmd"
	sdk "github.com/adminoid/cosmos-sdk/types"

	"bitbucket.org/decimalteam/go-smart-node/app"
	cmdcfg "bitbucket.org/decimalteam/go-smart-node/cmd/config"
)

func main() {

	// Setup config
	cfg := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(cfg)
	cmdcfg.SetBip44CoinType(cfg)
	cmdcfg.RegisterBaseDenom()
	cfg.Seal()

	// Create root command
	rootCmd, _ := NewRootCmd()

	// Execute root command
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
