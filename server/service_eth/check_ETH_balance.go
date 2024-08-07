package service_eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
	"strings"
)

// CheckETHBalance 查询ETH余额
func CheckETHBalance(address string) *big.Float {
	// 假设 Account1Address 是一个不含 '0x' 前缀的十六进制私钥字符串
	// 如果 Account1Address 包含 '0x'，需要去掉它，例如：
	fromAccount := strings.TrimPrefix(address, "0x")
	account := common.HexToAddress(fromAccount)
	//指定最新区块
	balanceAt, err := NFTClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}
