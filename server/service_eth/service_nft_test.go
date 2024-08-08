package service_eth

import (
	"fmt"
	"testing"
)

func TestTransferAccounts(t *testing.T) {
	ConnectNFTClient()
	//0.7819
	fromAccount := Account1
	//2.21
	toAccount := Account2Address
	//转0.01
	TransferAccounts(fromAccount, toAccount)
	defer NFTClient.Close()
}

func TestCheckETHBalance(t *testing.T) {
	ConnectNFTClient()
	address := Account2Address
	balance := CheckETHBalance(address)
	fmt.Println(balance)
	defer NFTClient.Close()
}

// tx sent: 0x5f60c05a6439756e3716f595307fb3b6db4c9f2bf01751823e29667652a5c35e
func TestQueryTransaction(t *testing.T) {
	ConnectNFTClient()
	QueryTransaction("0x5f60c05a6439756e3716f595307fb3b6db4c9f2bf01751823e29667652a5c35e")
	defer NFTClient.Close()
}

func TestCasting(t *testing.T) {
	ConnectNFTClient()
	found()
	defer NFTClient.Close()
}
