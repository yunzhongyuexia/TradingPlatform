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

func Casting() {
	privateKey, err := crypto.HexToECDSA(Account1)
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 构造交易
	toAddress := common.HexToAddress(Account1Address) // 替换为接收方的地址
	amount := big.NewInt(1000)                        // 替换为要铸造的代币数量

	// 通过方法签名获取方法元数据
	//Method: mintTokens,Inputs: [{to address false} {amount uint256 false}], Outputs: []
	mintTokensSignature := "mintTokens(address, uint256)"
	mintTokensMethod, exist := ContractABI.Methods[mintTokensSignature]
	if !exist {
		log.Fatalf("Method %s not found in the ABI", mintTokensSignature)
	}

	// 编码mintTokens函数的参数
	inputArgs, err := mintTokensMethod.Inputs.Pack(toAddress, amount)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := NFTClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := NFTClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(nftContractAddress),
		big.NewInt(0),
		500000, // 适当设置Gas限制，根据实际情况调整
		gasPrice,
		inputArgs,
	)

	chainID, err := NFTClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = NFTClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
