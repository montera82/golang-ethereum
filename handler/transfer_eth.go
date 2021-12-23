package handler

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/montera82/golang-ethereum/adapter"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transferer struct {
	client adapter.EthClient
	privateKey string
}

func NewTransferer(client adapter.EthClient, privateKey string) *Transferer {
	return &Transferer{client, privateKey}
}

func (t *Transferer) TransferEth(amount int64, to string) error {
	privateKey, err := crypto.HexToECDSA(t.privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("assert error: publicKey is not of type *ecdsa.PublicKey")
		return err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	value := big.NewInt(amount)
	nonce, err := t.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}
	gasLimit := uint64(21000)
	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	toAddress := common.HexToAddress(to)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainId, err := t.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = t.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return nil

}
