package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

const goSDKSource = "2"

type Client struct {
	rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{*c}
}

func (ec *Client) Close() {
	ec.Client.Close()
}

func (ec *Client) SendBundle(ctx context.Context, bundle *SendBundleArgs) (common.Hash, error) {
	var hash common.Hash
	ec.SetHeader("Request-Source", goSDKSource)
	err := ec.CallContext(ctx, &hash, "eth_sendBundle", bundle)
	return hash, err
}

func (ec *Client) BundlePrice(ctx context.Context) (*big.Int, error) {
	var bundlePrice big.Int

	err := ec.CallContext(ctx, &bundlePrice, "eth_bundlePrice")
	return &bundlePrice, err
}

func (ec *Client) GetBundleByHash(ctx context.Context, hash common.Hash) (*Bundle, error) {
	var bundle Bundle

	err := ec.CallContext(ctx, &bundle, "txpool_getBundleByHash", hash)
	return &bundle, err
}

func (ec *Client) GetStatus(ctx context.Context) (*Status, error) {
	var status Status

	err := ec.CallContext(ctx, &status, "eth_validatorStatus")
	return &status, err
}