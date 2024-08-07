package service_eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

// QueryTransaction 查询交易
func QueryTransaction(txHash string) {
	// 将十六进制的交易哈希字符串转换为 common.Hash 类型
	hash := common.HexToHash(txHash)

	// 获取交易
	tx, isPending, err := NFTClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Fatalf("Failed to retrieve the transaction: %v", err)
	}

	// 打印交易的基本信息
	fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
	if isPending {
		fmt.Println("This transaction is still pending.")
	} else {
		//fmt.Printf("From: %s\n", tx.From().Hex())
		fmt.Printf("To: %s\n", tx.To().Hex())
		fmt.Printf("Value: %s\n", tx.Value())
		fmt.Printf("Gas: %d\n", tx.Gas())
		fmt.Printf("Gas Price: %s\n", tx.GasPrice())
		fmt.Println("Transaction Nonce:", tx.Nonce())

		// 获取交易收据以检查交易状态
		receipt, err := NFTClient.TransactionReceipt(context.Background(), hash)
		if err != nil {
			log.Fatalf("Failed to retrieve the transaction receipt: %v", err)
		}
		if receipt.Status == types.ReceiptStatusFailed {
			fmt.Println("Transaction failed.")
		} else {
			fmt.Println("Transaction succeeded.")
		}
	}
}
