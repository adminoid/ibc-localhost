package api

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ethermintTypes "github.com/evmos/ethermint/types"
	"github.com/gogo/protobuf/proto"
)

// AddressResult contains API response fields.
type AddressResult struct {
	ID      uint64
	Address string
	Nonce   uint64
}

// Address requests full information about specified address
func (api *API) Address(address string) (*AddressResult, error) {
	client := authTypes.NewQueryClient(api.grpcClient)
	res, err := client.Account(
		context.Background(),
		&authTypes.QueryAccountRequest{address},
	)
	if err != nil {
		return nil, err
	}
	var acc ethermintTypes.EthAccount
	err = proto.Unmarshal(res.GetAccount().Value, &acc)
	if err != nil {
		return nil, err
	}

	return &AddressResult{
		ID:      acc.AccountNumber,
		Address: acc.Address,
		Nonce:   acc.Sequence,
	}, nil
}

// AccountNumberAndSequence requests account number and current sequence (nonce) of specified address.
func (api *API) AccountNumberAndSequence(address string) (uint64, uint64, error) {
	adr, err := api.Address(address)
	if err != nil {
		return 0, 0, err
	}
	return adr.ID, adr.Nonce, nil
}

func (api *API) AddressBalance(address string) (sdk.Coins, error) {
	bankClient := bankTypes.NewQueryClient(api.grpcClient)
	balance := sdk.NewCoins()
	req := &bankTypes.QueryAllBalancesRequest{
		Address:    address,
		Pagination: &query.PageRequest{Limit: queryLimit},
	}
	for {
		res, err := bankClient.AllBalances(
			context.Background(),
			req,
		)
		if err != nil {
			return sdk.NewCoins(), err
		}
		if res.Balances.Empty() {
			break
		}
		balance = balance.Add(res.Balances...)
		if len(res.Pagination.NextKey) == 0 {
			break
		}
		req.Pagination.Key = res.Pagination.NextKey
	}
	return balance, nil
}
