package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	web3hexutil "github.com/ethereum/go-ethereum/common/hexutil"
	web3types "github.com/ethereum/go-ethereum/core/types"
	web3 "github.com/ethereum/go-ethereum/ethclient"

	"github.com/tendermint/tendermint/libs/log"
	rpc "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	"bitbucket.org/decimalteam/go-smart-node/utils/helpers"
)

type Worker struct {
	ctx         context.Context
	httpClient  *http.Client
	cdc         params.EncodingConfig
	logger      log.Logger
	config      *Config
	hostname    string
	rpcClient   *rpc.HTTP
	web3Client  *web3.Client
	web3ChainId *big.Int
	query       chan *ParseTask
}

type Config struct {
	IndexerEndpoint string
	RpcEndpoint     string
	Web3Endpoint    string
	WorkersCount    int
}

func NewWorker(cdc params.EncodingConfig, logger log.Logger, config *Config) (*Worker, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	rpcClient, err := rpc.NewWithClient(config.RpcEndpoint, config.RpcEndpoint, httpClient)
	if err != nil {
		return nil, err
	}
	web3Client, err := web3.Dial(config.Web3Endpoint)
	if err != nil {
		return nil, err
	}
	web3ChainId, err := web3Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	worker := &Worker{
		ctx:         context.Background(),
		httpClient:  httpClient,
		cdc:         cdc,
		logger:      logger,
		config:      config,
		hostname:    hostname,
		rpcClient:   rpcClient,
		web3Client:  web3Client,
		web3ChainId: web3ChainId,
		query:       make(chan *ParseTask, 1000),
	}
	return worker, nil
}

func (w *Worker) Start() {
	wg := &sync.WaitGroup{}
	wg.Add(w.config.WorkersCount)
	for i := 0; i < w.config.WorkersCount; i++ {
		go w.executeFromQuery(wg)
	}
	wg.Wait()
}

func (w *Worker) executeFromQuery(wg *sync.WaitGroup) {
	defer wg.Done()
	w.getWork()
	for {
		task := <-w.query
		b := w.GetBlockResult(task.height, task.txNum)
		// Send
		data, err := json.Marshal(*b)
		w.panicError(err)
		w.sendBlock(task.height, data)

		w.getWork()
	}
}

func (w *Worker) GetBlockResult(height int64, txNum int) *Block {
	accum := NewEventAccumulator()

	// Fetch requested block from Tendermint RPC
	block := w.fetchBlock(height)

	// Fetch everything needed from Tendermint RPC aand EVM
	start := time.Now()
	txsChan := make(chan []Tx)
	resultsChan := make(chan *ctypes.ResultBlockResults)
	sizeChan := make(chan int)
	web3BlockChan := make(chan *web3types.Block)
	web3ReceiptsChan := make(chan web3types.Receipts)
	var parseTxNum int
	if txNum == -1 {
		parseTxNum = len(block.Block.Data.Txs)
	} else {
		parseTxNum = txNum
	}
	go w.fetchBlockTxs(height, parseTxNum, accum, txsChan)
	go w.fetchBlockTxResults(height, resultsChan)
	go w.fetchBlockSize(height, sizeChan)
	go w.fetchBlockWeb3(height, web3BlockChan)

	web3Block := <-web3BlockChan
	go w.fetchBlockTxReceiptsWeb3(web3Block, web3ReceiptsChan)
	txs := <-txsChan
	results := <-resultsChan
	size := <-sizeChan
	web3Body := web3Block.Body()
	web3Transactions := make([]*TransactionEVM, len(web3Body.Transactions))
	for i, tx := range web3Body.Transactions {
		msg, err := tx.AsMessage(web3types.NewLondonSigner(w.web3ChainId), nil)
		w.panicError(err)
		web3Transactions[i] = &TransactionEVM{
			Type:             web3hexutil.Uint64(tx.Type()),
			Hash:             tx.Hash(),
			Nonce:            web3hexutil.Uint64(tx.Nonce()),
			BlockHash:        web3Block.Hash(),
			BlockNumber:      web3hexutil.Uint64(web3Block.NumberU64()),
			TransactionIndex: web3hexutil.Uint64(uint64(i)),
			From:             msg.From(),
			To:               msg.To(),
			Value:            (*web3hexutil.Big)(msg.Value()),
			Data:             web3hexutil.Bytes(msg.Data()),
			Gas:              web3hexutil.Uint64(msg.Gas()),
			GasPrice:         (*web3hexutil.Big)(msg.GasPrice()),
			ChainId:          (*web3hexutil.Big)(tx.ChainId()),
			AccessList:       msg.AccessList(),
			GasTipCap:        (*web3hexutil.Big)(msg.GasTipCap()),
			GasFeeCap:        (*web3hexutil.Big)(msg.GasFeeCap()),
		}
	}
	web3Receipts := <-web3ReceiptsChan

	for _, event := range results.BeginBlockEvents {
		err := accum.AddEvent(event, "", results.Height)
		if err != nil {
			w.panicError(err)
		}
	}
	for _, event := range results.EndBlockEvents {
		err := accum.AddEvent(event, "", results.Height)
		if err != nil {
			w.panicError(err)
		}
	}

	// TODO: move to event accumulator
	// Retrieve emission and rewards
	var emission string
	var rewards []ProposerReward
	var commissionRewards []CommissionReward
	for _, event := range results.EndBlockEvents {
		switch event.Type {
		case "emission":
			// Parse emission
			emission = string(event.Attributes[0].Value)

		case "proposer_reward":
			// Parse proposer rewards
			var reward ProposerReward
			for _, attr := range event.Attributes {
				switch string(attr.Key) {
				case "amount", "accum_rewards":
					reward.Reward = string(attr.Value)
				case "validator", "accum_rewards_validator":
					reward.Address = string(attr.Value)
				case "delegator":
					reward.Delegator = string(attr.Value)
				}
			}
			rewards = append(rewards, reward)

		case "commission_reward":
			// Parser commission reward
			var reward CommissionReward
			for _, attr := range event.Attributes {
				switch string(attr.Key) {
				case "amount":
					reward.Amount = string(attr.Value)
				case "validator":
					reward.Validator = string(attr.Value)
				case "reward_address":
					reward.RewardAddress = string(attr.Value)
				}
			}
			commissionRewards = append(commissionRewards, reward)
		}
	}

	w.logger.Info(
		fmt.Sprintf("Compiled block (%s)", helpers.DurationToString(time.Since(start))),
		"block", height,
		"txs", len(txs),
		"begin-block-events", len(results.BeginBlockEvents),
		"end-block-events", len(results.EndBlockEvents),
	)

	// Create and fill Block object and then marshal to JSON
	return &Block{
		ID:                block.BlockID,
		Evidence:          block.Block.Evidence,
		Header:            block.Block.Header,
		LastCommit:        block.Block.LastCommit,
		Data:              BlockData{Txs: txs},
		Emission:          emission,
		Rewards:           rewards,
		CommissionRewards: commissionRewards,
		EndBlockEvents:    w.parseEvents(results.EndBlockEvents),
		BeginBlockEvents:  w.parseEvents(results.BeginBlockEvents),
		Size:              size,
		StateChanges:      *accum,
		EVM: BlockEVM{
			Header:       web3Block.Header(),
			Transactions: web3Transactions,
			Uncles:       web3Body.Uncles,
			Receipts:     web3Receipts,
		},
	}
}

func (w *Worker) panicError(err error) {
	if err != nil {
		w.logger.Error(fmt.Sprintf("Error: %v", err))
		panic(err)
	}
}
