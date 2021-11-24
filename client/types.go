package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SendBundleArgs represents the arguments for a call.
type SendBundleArgs struct {
	Txs               []string      `json:"txs,omitempty"`
	MaxBlockNumber    string        `json:"maxBlockNumber,omitempty"`
	MinTimestamp      *uint64       `json:"minTimestamp,omitempty"`
	MaxTimestamp      *uint64       `json:"maxTimestamp,omitempty"`
	RevertingTxHashes []common.Hash `json:"revertingTxHashes,omitempty"`
}

type Bundle struct {
	Txs               types.Transactions
	MaxBlockNumber    *big.Int
	MinTimestamp      uint64
	MaxTimestamp      uint64
	RevertingTxHashes []common.Hash
	Hash              common.Hash
	Price             *big.Int
}

type Status struct {
	Status     int64            `json:"status"`
	Validators map[string]int64 `json:"validators"`
}
