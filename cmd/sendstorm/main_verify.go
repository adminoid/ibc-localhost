package main

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	cosmosAuthTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	dscApi "bitbucket.org/decimalteam/go-smart-node/sdk/api"
)

func cmdVerify() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify-coins",
		Short: "Verify custom coins volume",
		Run: func(cmd *cobra.Command, args []string) {
			//
			err := cmd.Flags().Parse(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			reactor := stormReactor{}
			// init
			err = reactor.initApi(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			addresses, err := reactor.api.AllAccounts()
			if err != nil {
				fmt.Println(err)
				return
			}
			// coins info from coin keeper
			coinsInfo, err := reactor.coins()
			if err != nil {
				fmt.Println(err)
				return
			}
			balances := sdk.NewCoins()
			for _, adr := range addresses {
				coins, err := reactor.api.AddressBalance(adr)
				if err != nil {
					fmt.Println(err)
					return
				}
				balances = balances.Add(coins...)
			}
			for _, coinInfo := range coinsInfo {
				//if coinInfo.Denom == reactor.api.BaseCoin() {
				//	continue
				//}
				diff := "none"
				bal := balances.AmountOf(coinInfo.Denom)
				if !bal.Equal(coinInfo.Volume) {
					diff = fmt.Sprintf("volume=%s, balances=%s",
						coinInfo.Volume.String(), bal.String())
				}
				fmt.Printf("coin: (symbol: %s, volume: %s, reserve: %s), difference: %s\n",
					coinInfo.Denom, coinInfo.Volume, coinInfo.Reserve, diff)
			}
		},
	}

	return cmd
}

func cmdValidators() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "validators",
		Short: "Show validators info",
		Run: func(cmd *cobra.Command, args []string) {
			//
			err := cmd.Flags().Parse(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			reactor := stormReactor{}
			// init
			err = reactor.initApi(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			// validators info
			vals, err := reactor.api.Validators()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, val := range vals {
				dels, err := reactor.api.ValidatorDelegations(val.OperatorAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("moniker: %s, address: %s, status: %d, online: %v, jailed: %v, stake: %d, rewards: %s, delegation: %d\n",
					val.Description.Moniker, val.OperatorAddress, val.Status, val.Online, val.Jailed, val.Stake, val.Rewards, len(dels))
			}
		},
	}

	return cmd
}

func cmdVerifyPools() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify-pools",
		Short: "Verify (un/re)delegations and validator pools",
		Run: func(cmd *cobra.Command, args []string) {
			//
			err := cmd.Flags().Parse(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			reactor := stormReactor{}
			// init
			err = reactor.initApi(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			// validators info
			vals, err := reactor.api.Validators()
			if err != nil {
				fmt.Println(err)
				return
			}
			expectBondedPool := sdk.NewCoins()
			expectNotBondedPool := sdk.NewCoins()
			for _, val := range vals {
				// delegations
				dels, err := reactor.api.ValidatorDelegations(val.OperatorAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				for _, del := range dels {
					if del.Stake.Type == dscApi.StakeType_Coin {
						switch val.Status {
						case dscApi.BondStatus_Bonded:
							expectBondedPool = expectBondedPool.Add(del.Stake.Stake)
						case dscApi.BondStatus_Unbonded, dscApi.BondStatus_Unbonding:
							expectNotBondedPool = expectNotBondedPool.Add(del.Stake.Stake)
						}
					}
				}
				// redelegations
				reds, err := reactor.api.ValidatorRedelegations(val.OperatorAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				for _, red := range reds {
					for _, ent := range red.Entries {
						if ent.Stake.Type == dscApi.StakeType_Coin {
							expectNotBondedPool = expectNotBondedPool.Add(ent.Stake.Stake)
						}
					}
				}
				// undelegations
				ubds, err := reactor.api.ValidatorUndelegations(val.OperatorAddress)
				if err != nil {
					fmt.Println(err)
					continue
				}
				for _, ubd := range ubds {
					for _, ent := range ubd.Entries {
						if ent.Stake.Type == dscApi.StakeType_Coin {
							expectNotBondedPool = expectNotBondedPool.Add(ent.Stake.Stake)
						}
					}
				}
			}

			// check pool
			balanceBondedPool, err := reactor.api.AddressBalance(moduleNameToAddress("bonded_tokens_pool"))
			if err != nil {
				fmt.Println(err)
				return
			}
			balanceNotBondedPool, err := reactor.api.AddressBalance(moduleNameToAddress("not_bonded_tokens_pool"))
			if err != nil {
				fmt.Println(err)
				return
			}
			compareCoins("bonded_tokens_pool", balanceBondedPool, expectBondedPool)
			compareCoins("not_bonded_tokens_pool", balanceNotBondedPool, expectNotBondedPool)
		},
	}

	return cmd
}

func moduleNameToAddress(name string) string {
	address, err := bech32.ConvertAndEncode("dx", cosmosAuthTypes.NewModuleAddress(name))
	if err != nil {
		panic(fmt.Sprintf("moduleNameToAddress(%s) = %s", name, err.Error()))
	}
	return address
}

func compareCoins(name string, coins1, coins2 sdk.Coins) {
	if coins1.IsEqual(coins2) {
		fmt.Printf("pool '%s' is correct\n", name)
	} else {
		fmt.Printf("pool '%s' differs:\n", name)
		denoms := make([]string, 0)
		for _, coin := range coins1.Add(coins2...) {
			denoms = append(denoms, coin.Denom)
		}
		sort.Strings(denoms)
		for _, denom := range denoms {
			if !coins1.AmountOf(denom).Equal(coins2.AmountOf(denom)) {
				fmt.Printf("denom '%s' %s != %s\n", denom, coins1.AmountOf(denom), coins2.AmountOf(denom))
			}
		}
	}
}
