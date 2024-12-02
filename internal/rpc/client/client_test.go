package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"local/transaction/internal/rpc"
	"net/http"
	"reflect"
	"testing"
)

func TestNewClientWithOpts(t *testing.T) {
	t.Run("Without options", func(t *testing.T) {
		url := "https://example.com/jsonrpc"
		client := NewClientWithOpts(url, nil)

		if client.url != url {
			t.Errorf("Expected URL to be %s, got %s", url, client.url)
		}
		if client.httpClient == nil {
			t.Error("Expected httpClient to be initialized")
		}
		if len(client.customHeaders) != 0 {
			t.Errorf("Expected customHeaders to be empty, got %v", client.customHeaders)
		}
		if client.defaultRequestID != 0 {
			t.Errorf("Expected defaultRequestID to be 0, got %d", client.defaultRequestID)
		}
	})

	t.Run("With options", func(t *testing.T) {
		url := "https://example.com/jsonrpc"
		customClient := &http.Client{}
		customHeaders := map[string]string{"Authorization": "Bearer token"}
		defaultRequestID := 42

		opts := &RPCClientOpts{
			HTTPClient:       customClient,
			Headers:          customHeaders,
			DefaultRequestID: defaultRequestID,
		}
		client := NewClientWithOpts(url, opts)

		if client.url != url {
			t.Errorf("Expected URL to be %s, got %s", url, client.url)
		}
		if !reflect.DeepEqual(client.httpClient, customClient) {
			t.Errorf("Expected httpClient to be customClient, got %v", client.httpClient)
		}
		if !reflect.DeepEqual(client.customHeaders, customHeaders) {
			t.Errorf("Expected customHeaders to be %v, got %v", customHeaders, client.customHeaders)
		}
		if client.defaultRequestID != defaultRequestID {
			t.Errorf("Expected defaultRequestID to be %d, got %d", defaultRequestID, client.defaultRequestID)
		}
	})
}

func TestRPCClientCall(t *testing.T) {
	rpcClient := NewClientWithOpts("https://ethereum-rpc.publicnode.com", &RPCClientOpts{})

	params := []interface{}{"latest"}

	request := &rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  &params,
		ID:      1,
	}

	ctx := context.Background()
	response, err := rpcClient.Call(ctx, request)

	require.NoError(t, err, "Expected no error from RPCClient.Call")
	assert.NotNil(t, response, "Response should not be nil")
}
