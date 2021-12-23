package handler_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/montera82/golang-ethereum/handler"
	"github.com/montera82/golang-ethereum/mock"
	"github.com/montera82/golang-ethereum/testhelper"
)

func TestTransferer(t *testing.T) {

	tests := []struct {
		scenario string
		function func(*testing.T, *mock.EthClientMock)
	}{
		{
			"Should be able to successfully transfer eth",
			shouldBeAbleToSuccessfullyTransferEth,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			e := new(mock.EthClientMock)
			test.function(t, e)
		})
	}

}

func shouldBeAbleToSuccessfullyTransferEth(t *testing.T, ethclient *mock.EthClientMock) {
	ethclient.PendingNonceAtFunc = func(ctx context.Context, fromAddress common.Address) (uint64, error) {
		return 0, nil
	}
	ethclient.SuggestGasPriceFunc = func(ctx context.Context) (*big.Int, error) {
		return big.NewInt(153000000000000), nil
	}
	ethclient.NetworkIDFunc = func(ctx context.Context) (*big.Int, error) {
		return big.NewInt(1), nil
	}
	ethclient.SendTransactionFunc = func(ctx context.Context, signedTx *types.Transaction) error {
		return nil
	}
	fakePrivateKey := "24469d3bc6d70c3f7cd8a0ed15869a03395d2edb0327520c6eb25293fec48003"
	transferer := handler.NewTransferer(ethclient, fakePrivateKey)
	err := transferer.TransferEth(1, "0xD72716A5A55dB1eE226b6fEC93e860affA8c12f7")

	testhelper.Assert(t, err == nil, "Expected to successfully send eth but got: %v", err)
	testhelper.Assert(t, ethclient.SendTransactionInvoked == true, "Expected sendtransaction to be invoked")
}
