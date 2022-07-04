package main

import (
	"encoding/json"
	"fmt"

	dscApi "bitbucket.org/decimalteam/go-smart-node/sdk/api"
	dscTx "bitbucket.org/decimalteam/go-smart-node/sdk/tx"
	dscWallet "bitbucket.org/decimalteam/go-smart-node/sdk/wallet"
	helpers "bitbucket.org/decimalteam/go-smart-node/utils/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//helper function
func formatAsJSON(obj interface{}) string {
	objStr, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s\n", objStr)
}

var mult = sdk.NewInt(1000000000000000000)

func main() {
	api := dscApi.NewAPI(dscApi.ConnectionOptions{EndpointHost: "http://127.0.0.1", Timeout: 40})
	err := api.GetParameters()
	if err != nil {
		fmt.Printf("get ChainID error: %v\n", err)
		return
	}
	chainID := api.ChainID()
	fmt.Printf("ChainID=%s\n", chainID)
	for _, addr := range []string{"dx10n5429jflppqv9grprd3ywglzlheelarcvzsk6", "dx1r75hkp8q6ay535tskjdskynz3y3373pujqxyfr"} {
		res, err := api.Address(addr)
		if err != nil {
			fmt.Printf("get Address(%s) error: %v\n", addr, err)
			continue
		}
		fmt.Printf("Address=%s\n", formatAsJSON(res))
	}
	coins, err := api.Coins()
	if err != nil {
		fmt.Printf("api.Coins() error: %v\n", err)
	}
	fmt.Printf("api.Coins() result: %v\n", coins)
	///////////////
	faucetMnemonic := "bench own lock april loyal wheel nest slogan practice patch lamp change work scout slide city limb short festival card tornado enrich submit food"
	faucet, err := dscWallet.NewAccountFromMnemonicWords(faucetMnemonic, "")
	if err != nil {
		fmt.Printf("create wallet error: %v\n", err)
		return
	}
	an, as, err := api.AccountNumberAndSequence(faucet.Address())
	if err != nil {
		fmt.Printf("AccountNumberAndSequence error: %v\n", err)
		return
	}
	faucet = faucet.WithAccountNumber(an).WithSequence(as).WithChainID(chainID)
	mnemonics := []string{
		"possible hedgehog buddy desk smart camera frost vacant ridge robust seminar riot boost gauge jar aunt frozen morning system ordinary volcano rescue bind trust",
		"drip charge ridge between primary comic core fatigue evidence member fault tank tennis venue young lawsuit shock skull hybrid enlist shield opera please panther",
		"asthma science hawk hip piano enrich avoid myself divide seek number satoshi matter bunker question disease foster toward rare depth fame catch artefact woman",
	}
	for i, mn := range mnemonics {
		fmt.Printf("menmonic #%d\n", i)
		w, err := dscWallet.NewAccountFromMnemonicWords(mn, "")
		if err != nil {
			fmt.Printf("create wallet error: %v\n", err)
			return
		}
		tx, err := dscTx.BuildTransaction([]sdk.Msg{dscTx.NewMsgSendCoin(
			faucet.SdkAddress(),
			sdk.NewCoin(api.BaseCoin(), helpers.EtherToWei(sdk.NewInt(10))),
			w.SdkAddress(),
		)}, "some send", api.BaseCoin(), api.MaxGas())
		if err != nil {
			fmt.Printf("BuildTransaction error: %v\n", err)
			return
		}
		tx, err = tx.SignTransaction(faucet)
		if err != nil {
			fmt.Printf("SignTransaction error: %v\n", err)
			return
		}
		bytes, err := tx.BytesToSend()
		if err != nil {
			fmt.Printf("BytesToSend error: %v\n", err)
			return
		}
		r, err := api.BroadcastTxSync(bytes)
		if err != nil {
			fmt.Printf("err = %v\n", err)
			continue
		}
		fmt.Printf("result = %+v\n", r)
		faucet.IncrementSequence()
	}
}

/*
	tx, err := w.BuildTransaction([]sdk.Msg{dscTx.NewMsgCreateCoin(
		w.SdkAddress(),
		"aaa",
		"aaa",
		10,
		sdk.NewInt(1000).Mul(mult),
		sdk.NewInt(1000).Mul(mult),
		sdk.NewInt(100000).Mul(mult),
		"aaa",
	)}, "", "del")
	if err != nil {
		fmt.Printf("BuildTransaction error: %v\n", err)
		return
	}
	tx, err = w.SignTransaction(tx)
	if err != nil {
		fmt.Printf("SignTransaction error: %v\n", err)
		return
	}
	bytes, err := tx.BytesToSend()
	if err != nil {
		fmt.Printf("BytesToSend error: %v\n", err)
		return
	}
	fmt.Printf("%s\n", string(bytes))
	r, err := api.BroadcastTxSync(bytes)
	fmt.Println(string(r.Body()))
*/
