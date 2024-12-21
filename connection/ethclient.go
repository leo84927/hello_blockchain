package connection

import (
	"fmt"
	"hello_blockchain/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rotisserie/eris"
)

func NewEthClient() *ethclient.Client {

	cl, err := ethclient.Dial(config.EthNodeHttp)
	if err != nil {
		fmt.Println("node:", config.EthNodeHttp)
		panic(eris.Wrap(err, "ethclient init error"))
	}

	return cl
}

func GetEthClient() *ethclient.Client {
	return _ethClient
}
