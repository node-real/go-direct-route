package client

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

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
	err := ec.CallContext(context.Background(), &hash, "eth_sendBundle", bundle)
	return hash, err
}

func (ec *Client) BundlePrice(ctx context.Context) (*big.Int, error) {
	var bundlePrice big.Int

	err := ec.CallContext(context.Background(), &bundlePrice, "eth_bundlePrice")
	return &bundlePrice, err
}

func (ec *Client) GetBundleByHash(ctx context.Context, hash common.Hash) (*Bundle, error) {
	var bundle Bundle

	err := ec.CallContext(context.Background(), &bundle, "txpool_getBundleByHash", hash)
	return &bundle, err
}
