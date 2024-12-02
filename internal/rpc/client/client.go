package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"local/transaction/internal/rpc"
)

const (
	jsonrpcVersion = "2.0"
)

type RPCClient struct {
	url              string
	httpClient       *http.Client
	customHeaders    map[string]string
	defaultRequestID int
}

type RPCClientOpts struct {
	HTTPClient       *http.Client
	Headers          map[string]string
	DefaultRequestID int
}

func NewClientWithOpts(url string, opts *RPCClientOpts) *RPCClient {
	rpcClient := &RPCClient{
		url:           url,
		httpClient:    &http.Client{},
		customHeaders: make(map[string]string),
	}

	if opts == nil {
		return rpcClient
	}

	if opts.HTTPClient != nil {
		rpcClient.httpClient = opts.HTTPClient
	}

	if opts.Headers != nil {
		rpcClient.customHeaders = opts.Headers
	}

	rpcClient.defaultRequestID = opts.DefaultRequestID

	return rpcClient
}

func (c *RPCClient) Call(ctx context.Context, request *rpc.RPCRequest) (*rpc.RPCResponse, error) {
	if request.JSONRPC == "" {
		request.JSONRPC = jsonrpcVersion
	}

	httpRequest, err := c.newRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %w", request.Method, c.url, err)
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %w", request.Method, httpRequest.URL.Redacted(), err)
	}
	defer httpResponse.Body.Close()

	var rpcResponse *rpc.RPCResponse
	decoder := json.NewDecoder(httpResponse.Body)
	decoder.UseNumber()
	err = decoder.Decode(&rpcResponse)

	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. could't decode body to rpc response: %w", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode, err)
	}

	if rpcResponse == nil {
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode)
	}

	if httpResponse.StatusCode >= 400 {
		return rpcResponse, fmt.Errorf("rpc call %v() on %v status code: %v. rpc response error: %v", request.Method, httpRequest.URL.Redacted(), httpResponse.StatusCode, rpcResponse.Error)
	}

	return rpcResponse, nil
}

func (c *RPCClient) newRequest(ctx context.Context, req interface{}) (*http.Request, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	for k, v := range c.customHeaders {
		request.Header.Set(k, v)
	}

	return request, nil
}
