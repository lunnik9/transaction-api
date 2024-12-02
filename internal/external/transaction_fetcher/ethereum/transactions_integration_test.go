package ethereum

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"local/transaction/internal/rpc/client"
)

const ethereumHost = "https://ethereum-rpc.publicnode.com"

func TestGetBlockTransactionsSuccess(t *testing.T) {
	ctx := context.Background()

	clnt := New(client.NewClientWithOpts(ethereumHost, &client.RPCClientOpts{}))

	transactions, err := clnt.GetBlockTransactions(ctx, "latest")
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
}
