# Direct Route Go SDK

The Direct Route GO SDK provides a thin wrapper around the Direct Route API for sending private transactions.
It includes the following core components:

* **client** - implementations of Direct Route client, such as for querying bundle price, sending bundles.
* **example** - provide examples about how to using Direct Route in many different scenarios.

## What is Direct Route

Direct Route achieves following goals:
1. **Transaction privacy**. Transactions submitted through MEV can never be detected by others before they have been included in a block. 
2. **First-price sealed-bid auction**. It allows users to privately communicate their bid and granular transaction order preference.
3. **No paying for failed transactions**. Losing bids are never included in a block, thus never exposed to the public and no need to pay any transaction fees.
4. **Bundle transactions**. Multiple transactions are submitted as a bundle, the bundle transactions are all successfully validated on chain in the same block or never included on chain at all.
5. **Efficiency**. MEV extraction is performed without causing unnecessary network or chain congestion.


## Install

### Requirement

Go version above 1.16

### Use go mod(recommend)

Add `github.com/node-real/go-direct-route` dependency into your go.mod file. Example:

```go
require (
	github.com/node-real/go-direct-route latest
)
```

### Init Client

```
var directRouteEndPoint = "https://api.nodereal.io/direct-route"
client, _ := client.Dial(directRouteEndPoint)
```

### Quick Start with APIs

1. Query suggested bundle price
```
price, _ := client.BundlePrice(context.Background())
```

2. Send bundle
```
	data1, _ := bep20ABI.Pack("transfer", account2.Addr, big.NewInt(1))
	data2, _ := bep20ABI.Pack("transfer", account2.Addr, big.NewInt(1))

	tx1, hash1, _ := utils.SignTransaction(account1, common.HexToAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56"), valueToTransfer, data1, n1, gasLimit, price, chainId)
	tx2, hash2, _ := utils.SignTransaction(account1, common.HexToAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56"), valueToTransfer, data2, n1+1, gasLimit, price, chainId)
	maxTime := uint64(time.Now().Unix() + 80)
	minTime := uint64(time.Now().Unix() + 20)

	bundle := &client.SendBundleArgs{
		Txs:               []string{hexutil.Encode(tx1), hexutil.Encode(tx2)},
		MaxBlockNumber:    "",
		MinTimestamp:      &minTime,
		MaxTimestamp:      &maxTime,
		RevertingTxHashes: []common.Hash{hash2},
	}
	bundleHash, err := directClient.SendBundle(context.Background(), bundle)
```

After the bundle is successfully submitted, you may need wait at lest 3-60 seconds before the transaction been verified on chain.

So please use `MaxBlockNumber` and `MaxTimestamp` a relative lager one, better 60 seconds later, otherwise DR may 
get no chance to include the bundle.

Note that only one tx sender is allowed with one bundle.

3. Query bundle

```
bundle, _ := directClient.GetBundleByHash(context.Background(), bundleHash)
```


### SDK Example

We provide three demos in `example.go`:
1. `getBundlePriceDemo`. The bundle price is volatile according to the 
network congestion, the demo shows you how to get proper bundle price.
2. `sendBNBByBundleDemo`. In this case, we use two different 
accounts to send BNB to each other, the two transaction should be all 
successful or all failed.
3. `sendBUSDByBundleDemo`. In this case, we use two accounts to send BUSD 
to each other, the second transaction is allowed to be failed,
and the bundle should be verified on chain during [now+20 second, now+80 second].
This case shows you how to interact with smart contract through direct-route,
and how to control the timing to be verified.

If you want to try with above examples, what you need to do is just to 
replace the private keys of `account1` and `account2` in `example.go`


