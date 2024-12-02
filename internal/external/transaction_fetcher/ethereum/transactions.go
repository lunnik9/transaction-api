package ethereum

import (
	"context"
	"encoding/json"
	"fmt"

	"local/transaction/internal/domain"
	"local/transaction/internal/rpc"
)

const ethGetBlockByNumberMethod = "eth_getBlockByNumber"

type EthereumClient struct {
	rpcClient rpc.RPCClient
}

func New(rpcClient rpc.RPCClient) *EthereumClient {
	return &EthereumClient{rpcClient: rpcClient}
}

func (c *EthereumClient) GetBlockTransactions(ctx context.Context, blockID string) ([]domain.Transaction, error) {
	params := []interface{}{blockID, true}

	rpcResp, err := c.rpcClient.Call(ctx, &rpc.RPCRequest{
		Method: ethGetBlockByNumberMethod,
		Params: &params,
	})
	if err != nil {
		return nil, err
	}

	if rpcResp.Result == nil {
		return nil, domain.BlockNotFoundError
	}

	var transactions domain.Transactions
	err = json.Unmarshal(rpcResp.Result, &transactions)
	if err != nil {
		return nil, fmt.Errorf("can't umarshal GetBlockTransactions responce %w", err)
	}

	return transactions.Transactions, nil
}
