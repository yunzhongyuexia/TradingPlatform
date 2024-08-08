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

func Airdrop() {
	privateKey, err := crypto.HexToECDSA(Account1)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := NFTClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := NFTClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 构造空投交易的参数
	recipients := []common.Address{
		common.HexToAddress(Account1Address),
		common.HexToAddress(Account2Address),
		// 更多接收者地址...
	}
	values := []*big.Int{
		big.NewInt(1000), // 每个NFT的价值
		big.NewInt(2000), // 使用*big.Int表示2000
		// 更多NFT的价值...
	}
	encoded, err := JSONABI.Pack("airdropNFT", recipients, values)
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(0)    // 空投NFT通常不需要以太币
	gasLimit := uint64(300000) // 适当设置Gas限制，根据实际情况调整
	tx := types.NewTransaction(nonce, common.HexToAddress(nftContractAddress), amount, gasLimit, gasPrice, encoded)

	// 使用私钥签名交易
	chainIDBig := big.NewInt(int64(chainID))
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainIDBig), privateKey)
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

	fmt.Printf("Airdrop transaction mined in block: %d\n", receipt.BlockNumber)
}
