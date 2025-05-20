package eth

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	ClientUrl string
	Client    *ethclient.Client
	ChainID   big.Int
}

func Init(clientUrl string, chainID string) *EthClient {
	parsedID, err := strconv.Atoi(chainID)
	if err != nil {
		log.Fatal("Invalid chain id:", err)
		return nil
	}
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatal("Connection to Ethereum client FAILED", err)
		return nil
	}
	fmt.Println("Connected to Ethereum client")
	return &EthClient{
		ClientUrl: clientUrl,
		Client:    client,
		ChainID:   *big.NewInt(int64(parsedID)),
	}
}
