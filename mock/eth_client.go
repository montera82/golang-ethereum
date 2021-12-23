package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	EthClientMock struct {
		SendTransactionInvoked bool
		SendTransactionFunc    func(ctx context.Context, signedTx *types.Transaction) error

		PendingNonceAtInvoked bool
		PendingNonceAtFunc    func(ctx context.Context, fromAddress common.Address) (uint64, error)

		SuggestGasPriceInvoked bool
		SuggestGasPriceFunc    func(ctx context.Context) (*big.Int, error)

		NetworkIDInvoked bool
		NetworkIDFunc    func(ctx context.Context) (*big.Int, error)
	}
)

func (e *EthClientMock) PendingNonceAt(ctx context.Context, fromAddress common.Address) (uint64, error) {

	e.PendingNonceAtInvoked = true
	return e.PendingNonceAtFunc(ctx, fromAddress)
}

func (e *EthClientMock) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	e.SuggestGasPriceInvoked = true
	return e.SuggestGasPriceFunc(ctx)
}

func (e *EthClientMock) NetworkID(ctx context.Context) (*big.Int, error) {
	e.NetworkIDInvoked = true
	return e.NetworkIDFunc(ctx)
}

func (e *EthClientMock) SendTransaction(ctx context.Context, signedTx *types.Transaction) error {
	e.SendTransactionInvoked = true
	return e.SendTransactionFunc(ctx, signedTx)
}
