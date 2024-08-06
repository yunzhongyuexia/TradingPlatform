package setETH

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"
	"strings"
)

var Client *ethclient.Client

// ClientURL holesky测试网
var clientURL = "https://ethereum-holesky.core.chainstack.com/5cb26e90d4925a7f24a1c7b51c2d8263"

// Account1 私钥
var Account1 = "0x3b35e382c432bfdfbcafcdf5e428fff8998ccce50204ce0336025e5a01bd63e6"

// Account1Address Account1地址
var Account1Address = "0x44cc3545d27fc25a4fe827e76ade6f74d4ac7a02"

// Account2 私钥
var Account2 = "0x5e1d6f1a64cc1d3ad1c155fcdb3165a0a844d797a6e99f770910fb5bc81134b3"

// Account2Address Account2地址
var Account2Address = "0x8ff0bf277ea5ae4e3047c178923d973e6eac27f1"

func ConnectETH() {
	// 连接到以太坊节点
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	ReadABIFromFile()
}

// Balance 查询ETH余额
func Balance() {
	account := common.HexToAddress(Account1)
	blockNumber := big.NewInt(5532993)
	balanceAt, err := Client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
}

// GenerateWallet 生成钱包
func GenerateWallet() string {
	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//返回公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 获取公钥的序列化形式
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// 使用Keccak-256哈希算法计算哈希值
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:]) // 跳过前缀'0x04'，因为公钥通常以'0x04'开头，后面跟随64字节的X和Y坐标

	// 截取哈希值的最后20个字节
	address := hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Println("wallet:", address)
	return address
}

// QueryBlock 查询区块
func QueryBlock() {
	//调用客户端的HeaderByNumber来返回有关一个区块的头信息。若传入nil，它将返回最新的区块头。
	header, err := Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.Int64()) // 最新区块头

	blockNumber := big.NewInt(header.Number.Int64()) //查询最新区块头的完整区块
	block, err := Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 打印其他区块信息
	fmt.Println("Block Time:", block.Time())                          // 区块时间戳
	fmt.Println("Block Difficulty:", block.Difficulty().Uint64())     // 区块难度
	fmt.Println("Block Hash:", block.Hash().Hex())                    // 区块哈希值
	fmt.Println("Number of Transactions:", len(block.Transactions())) // 区块中交易的数量

	//调用Transaction只返回一个区块的交易数目。
	//count, err := Client.TransactionCount(context.Background(), block.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(count) // 144
}

// QueryTransaction 查询交易
func QueryTransaction(txHash string) {
	// 将十六进制的交易哈希字符串转换为 common.Hash 类型
	hash := common.HexToHash(txHash)

	// 获取交易
	tx, isPending, err := Client.TransactionByHash(context.Background(), hash)
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
		receipt, err := Client.TransactionReceipt(context.Background(), hash)
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

// TransferAccounts ETH转帐
func TransferAccounts() {
	// 假设 Account1 是一个不含 '0x' 前缀的十六进制私钥字符串
	// 如果 Account1 包含 '0x'，需要去掉它，例如：
	fromAccount1 := strings.TrimPrefix(Account1, "0x")
	privateKey, err := crypto.HexToECDSA(fromAccount1)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(50000000000000000) // in wei (0.05 eth)
	gasLimit := uint64(21000)              // in units
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(Account2Address)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
