package service_eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func Synthesis() {
	privateKeyBytes, err := crypto.HexToECDSA(Account1)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKeyBytes.PublicKey)
	nonce, err := NFTClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := NFTClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 构造合成交易的参数
	nftIds := []uint64{1, 2, 3} // 假设我们要合成的NFT ID列表
	encoded, err := JSONABI.Pack("synthesizeNFTs", nftIds)
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(0)    // 合成NFT通常不需要以太币
	gasLimit := uint64(300000) // 适当设置Gas限制，根据实际情况调整
	tx := types.NewTransaction(nonce, common.HexToAddress(nftContractAddress), amount, gasLimit, gasPrice, encoded)

	// 使用私钥签名交易
	chainIDBig := big.NewInt(int64(chainID))
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainIDBig), privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = NFTClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// 等待交易确认
	receipt, err := NFTClient.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatalf("Transaction failed: %v", receipt)
	}

	fmt.Printf("Transaction mined in block: %d\n", receipt.BlockNumber)
}
