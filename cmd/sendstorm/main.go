package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	dscApi "bitbucket.org/decimalteam/go-smart-node/sdk/api"
	dscTx "bitbucket.org/decimalteam/go-smart-node/sdk/tx"
	"bitbucket.org/decimalteam/go-smart-node/sdk/wallet"
	dscWallet "bitbucket.org/decimalteam/go-smart-node/sdk/wallet"
	helpers "bitbucket.org/decimalteam/go-smart-node/utils/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	mnemonicsFlag  = "mnemonics_file"
	nodeFlag       = "node"
	tendermintPort = "tport"
	restPort       = "rport"
	turnOnDebug    = "debug"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "sendstorm",
		Short: "SendStorm is testing application for decimal blockchain",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	rootCmd.PersistentFlags().String(mnemonicsFlag, "mnemonics.cfg", "path to mnemonics file")
	rootCmd.PersistentFlags().String(nodeFlag, "http://localhost", "hostname of decimal node as http://... without port")
	rootCmd.PersistentFlags().Int(tendermintPort, 26657, "tendermint RPC port of decimal node")
	rootCmd.PersistentFlags().Int(restPort, 1317, "REST port of decimal node")
	rootCmd.PersistentFlags().Bool(turnOnDebug, false, "write api requests/responses to sendstorm.log")

	rootCmd.AddCommand(
		cmdGenerate(),
		cmdFaucet(),
		cmdRun(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmdGenerate() *cobra.Command {
	var mnemonicsCount int

	var cmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate mnemonics and save them to file",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Flags().Parse(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			mnemonicsFile, err := cmd.Flags().GetString(mnemonicsFlag)
			if err != nil {
				fmt.Println(err)
				return
			}
			f, err := os.Create(mnemonicsFile)
			if err != nil {
				fmt.Printf("create mnemonic file error: %s\n", err.Error())
				return
			}
			defer f.Close()
			fmt.Printf("create %d mnemonics\n", mnemonicsCount)
			for i := 0; i < mnemonicsCount; i++ {
				mn, err := wallet.NewMnemonic("")
				if err != nil {
					fmt.Printf("create mnemonic error: %s\n", err.Error())
					continue
				}
				_, err = fmt.Fprintf(f, "%s\n", mn.Words())
				if err != nil {
					fmt.Printf("write to mnemonic file error: %s\n", err.Error())
				}
			}
		},
	}

	cmd.PersistentFlags().IntVar(&mnemonicsCount, "count", 10, "count of mnemonics")

	return cmd
}

func cmdFaucet() *cobra.Command {
	var onlyEmpty bool
	var amountToSend int64

	var cmd = &cobra.Command{
		Use:   "faucet",
		Short: "Send some base coins to accounts from mnemonics",
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
			err = reactor.initAccounts(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			err = reactor.initFaucet(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			// do action
			for i, acc := range reactor.accounts {
				err := acc.Update()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(acc.currentBalance)
				fmt.Printf("account: (%d) %s, balance: %s\n", i, acc.Address(), acc.BalanceForCoin(reactor.api.BaseCoin()))
				if onlyEmpty && acc.BalanceForCoin(reactor.api.BaseCoin()).GT(sdk.ZeroInt()) {
					continue
				}
				msg := dscTx.NewMsgSendCoin(
					reactor.faucetAccount.SdkAddress(),
					sdk.NewCoin(reactor.api.BaseCoin(), helpers.EtherToWei(sdk.NewInt(amountToSend))),
					acc.account.SdkAddress(),
				)
				tx, err := dscTx.BuildTransaction([]sdk.Msg{msg}, "", reactor.api.BaseCoin(), 0) /// TODO: 0 gas is temporary
				if err != nil {
					fmt.Println(err)
					continue
				}
				tx, err = tx.SignTransaction(reactor.faucetAccount)
				if err != nil {
					fmt.Println(err)
					continue
				}
				data, err := tx.BytesToSend()
				if err != nil {
					fmt.Println(err)
					continue
				}
				res, err := reactor.api.BroadcastTxSync(data)
				if err != nil {
					fmt.Println(err)
					continue
				}
				if res.Code != 0 {
					fmt.Println(res)
					time.Sleep(time.Second * 6) // wait ~block
					// reset faucet
					err = reactor.initFaucet(cmd.Flags())
					if err != nil {
						fmt.Println(err)
						return
					}
					continue
				}
				reactor.faucetAccount.IncrementSequence()
			}
		},
	}

	cmd.PersistentFlags().String("faucet_mnemonic", "", "faucet mnemonic")
	cmd.PersistentFlags().BoolVar(&onlyEmpty, "only_empty", true, "send coins to account with zero balance")
	cmd.PersistentFlags().Int64Var(&amountToSend, "amount", 10000, "amount of base coins to send")

	return cmd
}

func cmdRun() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "run",
		Short: "Run actions",
		Run: func(cmd *cobra.Command, args []string) {
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
			err = reactor.initAccounts(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			err = reactor.initActionReactor(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			err = reactor.initLimiter(cmd.Flags())
			if err != nil {
				fmt.Println(err)
				return
			}
			reactor.updateGeneratorsInfo()
			// infinite loop
			n := 0
			for {
				n++
				if n >= 100 {
					reactor.updateGeneratorsInfo()
					n = 0
				}
				action := reactor.actionReactor.Generate()
				acc := reactor.accounts[rand.Intn(len(reactor.accounts))]
				if !action.CanPerform(acc) {
					/*
						amt := action.(*ActionSend).coin.Amount
						blnc := acc.BalanceForCoin("del")
						adr := action.(*ActionSend).address
						fmt.Printf(" dirty: %v, less: %v, addr eq: %v\n",
							acc.IsDirty(),
							blnc.LT(amt),
							acc.Address() == adr,
						)
					*/
					//fmt.Printf("cannot %#v for %s\n", action, acc.Address())
					continue
				}
				if !reactor.limiter.CanMake() {
					continue
				}
				bytesToSend, err := action.GenerateTx(acc)
				if err != nil {
					fmt.Println(err)
					return
				}
				res, err := reactor.api.BroadcastTxSync(bytesToSend)
				if err != nil {
					fmt.Println(err)
					acc.MarkDirty()
					// TODO: Update() returns error
					go acc.Update()
					continue
				}
				fmt.Printf("%v\n", res)
				if res.Code != 0 {
					fmt.Printf("%s: (%d) %s\n", acc.Address(), res.Code, res.Log)
					acc.MarkDirty()
					// TODO: Update() returns error
					go acc.Update()
					continue
				}
				acc.IncrementSequence()
				go acc.UpdateBalance()
			}
		},
	}

	cmd.PersistentFlags().StringSlice("actions", []string{},
		"actions list in format: action1=weight1,action2=weight2,... weight must be integer")
	cmd.PersistentFlags().Int64("tps", 1, "transactions per second")

	return cmd
}

// helpers
func loadMnemonics(fname string) ([]string, error) {
	var result []string
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		mn := fileScanner.Text()
		result = append(result, mn)
	}
	return result, nil
}

type stormReactor struct {
	accounts      []*StormAccount
	faucetAccount *dscWallet.Account
	api           *dscApi.API
	actionReactor *ActionReactor
	limiter       *TPSLimiter
}

func (reactor *stormReactor) initApi(flags *pflag.FlagSet) error {
	nodeHost, err := flags.GetString(nodeFlag)
	if err != nil {
		return err
	}
	tPort, err := flags.GetInt(tendermintPort)
	if err != nil {
		return err
	}
	rPort, err := flags.GetInt(restPort)
	if err != nil {
		return err
	}
	debug, err := flags.GetBool(turnOnDebug)
	if err != nil {
		return err
	}
	// TODO: make this dirty debug more nice
	if debug {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		logfile, err := os.Create("sendstorm.log")
		if err != nil {
			log.Fatalln("Cannot create log file")
		}
		//defer logfile.Close()
		log.SetOutput(logfile)
	}

	reactor.api = dscApi.NewAPI(dscApi.ConnectionOptions{
		EndpointHost:   nodeHost,
		TendermintPort: tPort,
		RestPort:       rPort,
		Timeout:        10,
		Debug:          debug,
	})
	err = reactor.api.GetParameters()
	if err != nil {
		return err
	}
	return nil
}

func (reactor *stormReactor) initAccounts(flags *pflag.FlagSet) error {
	mnemonicsFile, err := flags.GetString(mnemonicsFlag)
	if err != nil {
		return err
	}
	mnemonics, err := loadMnemonics(mnemonicsFile)
	if err != nil {
		return err
	}
	reactor.accounts = make([]*StormAccount, 0)
	for _, mn := range mnemonics {
		acc, err := NewStormAccount(mn, reactor.api)
		if err != nil {
			return err
		}
		reactor.accounts = append(reactor.accounts, acc)
	}
	return nil
}

func (reactor *stormReactor) initFaucet(flags *pflag.FlagSet) error {
	faucetMnemonic, err := flags.GetString("faucet_mnemonic")
	if err != nil {
		return err
	}
	reactor.faucetAccount, err = dscWallet.NewAccountFromMnemonicWords(faucetMnemonic, "")
	if err != nil {
		return err
	}
	fmt.Printf("faucet address: %s\n", reactor.faucetAccount.Address())
	an, as, err := reactor.api.AccountNumberAndSequence(reactor.faucetAccount.Address())
	if err != nil {
		return err
	}
	reactor.faucetAccount = reactor.faucetAccount.WithAccountNumber(an).WithSequence(as).WithChainID(reactor.api.ChainID())
	return nil
}

func (reactor *stormReactor) initActionReactor(flags *pflag.FlagSet) error {
	// simple action parser
	actions, err := flags.GetStringSlice("actions")
	if err != nil {
		return err
	}
	reactor.actionReactor = &ActionReactor{}
	for _, act := range actions {
		ss := strings.Split(act, "=")
		if len(ss) != 2 {
			return fmt.Errorf("'%s' must be actionName=weight", act)
		}
		generatorName := ss[0]
		weight, err := strconv.ParseInt(ss[1], 10, 64)
		if err != nil {
			return fmt.Errorf("'%s' weight must be integer, go error '%s'", act, err.Error())
		}
		err = reactor.actionReactor.Add(generatorName, weight)
		if err != nil {
			return fmt.Errorf("'%s': %s", act, err.Error())
		}
	}
	return nil
}

func (reactor *stormReactor) initLimiter(flags *pflag.FlagSet) error {
	limit, err := flags.GetInt64("tps")
	if err != nil {
		return err
	}
	reactor.limiter = NewTPSLimiter(limit)
	return nil
}

func (reactor *stormReactor) updateGeneratorsInfo() {
	// update info
	ui := UpdateInfo{}
	coins, err := reactor.api.Coins()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range coins {
		ui.Coins = append(ui.Coins, c.Symbol)
		ui.FullCoins = append(ui.FullCoins, c)
	}
	for _, acc := range reactor.accounts {
		err = acc.Update()
		if err != nil {
			fmt.Println(err)
		}
		ui.Addresses = append(ui.Addresses, acc.Address())
	}
	reactor.actionReactor.Update(ui)
}
