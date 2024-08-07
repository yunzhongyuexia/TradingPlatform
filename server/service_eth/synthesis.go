package service_eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

type NFTSynthesizedEvent struct {
	TokenId   *big.Int
	Recipient common.Address
}

func Synthesis(tokenId1, tokenId2, tokenId3 *big.Int, newTokenURI string) *int {
	// Convert recipient address from string to common.Address
	recipient := common.HexToAddress(Account1Address)

	contractAddress := common.HexToAddress(nftContractAddress)

	// Convert private key from hex string
	privateKey, err := crypto.HexToECDSA(Account1)
	if err != nil {
		log.Fatal("failed to convert private key:", err)
	}

	// Create an authorized transactor
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasPrice = big.NewInt(20000000000) // Set gas price (20 Gwei)
	auth.GasLimit = uint64(300000)          // Set gas limit
	// Get the nonce for the sender's account
	nonce, err := NFTClient.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal("failed to get nonce:", err)
	}

	// Pack the transaction data, including the recipient address
	data, err := ContractABI.Pack("synthesizeNFT", tokenId1, tokenId2, tokenId3, newTokenURI, recipient)
	if err != nil {
		log.Fatal("failed to pack data:", err)
	}

	// Create the transaction
	tx := types.NewTransaction(
		nonce,           // Nonce
		contractAddress, // Contract address
		big.NewInt(0),   // Value to send (zero for contract interaction)
		auth.GasLimit,   // Gas limit
		auth.GasPrice,   // Gas price
		data,            // Data
	)

	// Sign the transaction
	chainID, err := NFTClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal("failed to get chain ID: ", err)
	}

	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatal("failed to sign transaction:", err)
	}

	// Send the signed transaction
	err = NFTClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("failed to send transaction:", err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// Wait for transaction receipt
	receipt, err := bind.WaitMined(context.Background(), NFTClient, signedTx)
	if err != nil {
		log.Fatal("failed to get transaction receipt:", err)
	}

	if receipt.Status != 1 {
		log.Fatal("transaction failed with status: ", receipt.Status)
	}

	// Decode the transaction receipt logs
	var newTokenId *big.Int
	for _, log := range receipt.Logs {
		event, err := ContractABI.EventByID(log.Topics[0])
		if err != nil {
			continue
		}
		if event.Name == "NFTSynthesized" {
			parsedEvent := NFTSynthesizedEvent{}
			err := ContractABI.UnpackIntoInterface(&parsedEvent, "NFTSynthesized", log.Data)
			if err != nil {
				return nil
			}
			newTokenId = parsedEvent.TokenId
			break
		}
	}

	if newTokenId == nil {
		log.Fatal("获取tokenId失败")
	}
	intValue := int(newTokenId.Int64())

	return &intValue
}

//func SynthesizeDb(tokenId1, tokenId2, tokenId3 string, newToken nft.Token) error {
//	// 获取数据库实例
//	db := db.Mysql("nft")
//
//	// 开始一个事务
//	tx := db.Begin()
//
//	// 确保在事务结束时进行回滚或提交
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//			panic(r) // 重新抛出 panic
//		}
//	}()
//
//	// 删除操作
//	if err := tx.Delete(&nft.Token{}, "id = ?", tokenId1).Error; err != nil {
//		tx.Rollback()
//		return fmt.Errorf("failed to delete token with id %s: %w", tokenId1, err)
//	}
//	if err := tx.Delete(&nft.Token{}, "id = ?", tokenId2).Error; err != nil {
//		tx.Rollback()
//		return fmt.Errorf("failed to delete token with id %s: %w", tokenId2, err)
//	}
//	if err := tx.Delete(&nft.Token{}, "id = ?", tokenId3).Error; err != nil {
//		tx.Rollback()
//		return fmt.Errorf("failed to delete token with id %s: %w", tokenId3, err)
//	}
//
//	// 创建新记录
//	if err := tx.Updates(&newToken).Error; err != nil {
//		tx.Rollback()
//		return fmt.Errorf("failed to create new token: %w", err)
//	}
//
//	// 提交事务
//	return tx.Commit().Error
//}
