package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/montera82/golang-ethereum/contracts"
	"github.com/montera82/golang-ethereum/cmd"
)

func main() {
	var c cmd.Config

	if err := cleanenv.ReadConfig("../../.env", &c); err != nil {
		log.Fatal("unable to load config")
	}
	client, err := ethclient.Dial(c.Web3Provider)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(c.TransfererPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("assert error: publicKey is not of type *ecdsa.PublicKey")

	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nounce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nounce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := contracts.DeployContracts(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance

}
