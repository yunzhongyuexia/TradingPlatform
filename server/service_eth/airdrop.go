package service_eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Airdrop() {
	// 智能合约地址
	contractAddress := common.HexToAddress(nftContractAddress)

	// 接收空投的地址列表
	recipients := []string{
		Account1Address,
		Account2Address,
		// ...更多地址
	}

	// 发送者地址和私钥（需要有足够余额和权限发送交易）
	fromAddress := common.HexToAddress(Account1Address)
	privateKey, _ := crypto.HexToECDSA(Account1)
	ctx := context.Background()
	// 执行空投
	if err := AirdropNFTs(ctx, NFTClient, contractAddress, recipients, fromAddress, privateKey); err != nil {
		log.Fatal(err)
	}
}
func AirdropNFTs(ctx context.Context, client *ethclient.Client, contractAddress common.Address, recipients []string, fromAddress common.Address, privateKey *ecdsa.PrivateKey) error {
	chainID, _ := client.ChainID(ctx)        // 获取链ID
	signer := types.NewEIP155Signer(chainID) // 创建签名者

	// 构造调用数据，这里需要根据智能合约的具体情况来构造
	data, err := ContractABI.Pack("airdropNFT", recipients)
	if err != nil {
		return err
	}

	// 创建交易
	tx := types.NewTransaction(
		0, // nonce
		contractAddress,
		big.NewInt(0), // 空投函数可能不需要支付 Ether
		500000,        // gas limit
		big.NewInt(0), // gas price
		data,
	)

	// 签署交易
	var chainIDBig *big.Int
	fmt.Println(chainIDBig)
	if chainID != nil {
		chainIDBig = new(big.Int).Set(chainID)
	}
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return err
	}

	// 发送交易
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}

	fmt.Println("Transaction has been sent.")
	return nil
}
