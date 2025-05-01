package utils

import (
	"context"
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient() *ethclient.Client {
	rpc := os.Getenv("RPC_URL")
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	return client
}

func GetPrivateKey() *ecdsa.PrivateKey {
	privHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privHex[2:])
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	return privateKey
}

func GetTransactOpts(client *ethclient.Client) *bind.TransactOpts {
	privateKey := GetPrivateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.From = fromAddress
	auth.Context = context.Background()
	auth.GasLimit = uint64(8000000)
	return auth
}
