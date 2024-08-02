package setETH

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 以太坊节点
var clientURL = "https://ethereum-holesky.core.chainstack.com/5cb26e90d4925a7f24a1c7b51c2d8263"
var recipientStr = "0x8ff0bf277eA5aE4E3047c178923d973e6eac27f1"
var contractAddr = "0xAC55603B967c46184D5112093dF502fd0a9289C8"

func ConnectETH() {
	// 连接到以太坊节点
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ReadABIFromFile()

}
