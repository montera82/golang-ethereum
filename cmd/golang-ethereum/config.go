package main

type Config struct {
	TransfererPrivateKey string `env:"TRANSFERER_PRIVATE_KEY" required:"true"`
	Web3Provider string `env:"WEB3_PROVIDER" required:"true"`
}
