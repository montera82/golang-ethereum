package main

import (
	// "context"
	"fmt"
	"log"

	// "math"
	// "math/big"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/montera82/golang-ethereum/transfer"
)

func main() {
	var c Config

	if err := cleanenv.ReadConfig(".env", &c); err!= nil {
		log.Fatal("unable to load config")
	}

	client, err := ethclient.Dial(c.Web3Provider)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection!")

	t := transfer.NewTransferer(
		client,
		c.TransfererPrivateKey,
	)
	err = t.TransferEth(
		1000000000000000000, // 1 eth
		"0xf17f52151EbEF6C7334FAD080c5704D77216b732",
	)
	if err != nil {
		log.Fatal(err)
	}

	// address := common.HexToAddress("0x627306090abaB3A6e1400e9345bC60c78a8BEf57")
	// fmt.Println(address.Hex())
	// fmt.Println(address.Hash().Hex())
	// fmt.Println(address.Bytes())

	// // get account balance
	// balance, err := client.BalanceAt(context.Background(), address, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fbalance := new(big.Float)
	// fbalance.SetString(balance.String())
	// ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	// fmt.Println(ethValue)
	// fmt.Println(balance)
}
