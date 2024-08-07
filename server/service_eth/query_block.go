package service_eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
)

// QueryBlock 查询区块
func QueryBlock() {
	//调用客户端的HeaderByNumber来返回有关一个区块的头信息。若传入nil，它将返回最新的区块头。
	header, err := NFTClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("最新区块头:", header.Number.Int64())

	//查询最新区块头的完整区块
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := NFTClient.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 打印其他区块信息
	fmt.Println("区块时间戳:", block.Time())
	fmt.Println("区块难度:", block.Difficulty().Uint64())
	fmt.Println("区块哈希值:", block.Hash().Hex())
	fmt.Println("区块中交易的数量:", len(block.Transactions()))

	//调用Transaction只返回一个区块的交易数目。
	//count, err := Client.TransactionCount(context.Background(), block.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(count) // 144
}
