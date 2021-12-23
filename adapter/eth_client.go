package adapter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	EthClient interface {
		PendingNonceAt(ctx context.Context, fromAddress common.Address) (uint64, error)
		SuggestGasPrice(ctx context.Context) (*big.Int, error)
		NetworkID(ctx context.Context) (*big.Int, error)
		SendTransaction(ctx context.Context, signedTx *types.Transaction) error
	}
)

