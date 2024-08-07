package service_eth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

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
	// 跳过前缀'0x04'，因为公钥通常以'0x04'开头，后面跟随64字节的X和Y坐标
	hash.Write(publicKeyBytes[1:])

	// 截取哈希值的最后20个字节
	address := hexutil.Encode(hash.Sum(nil)[12:])
	return address
}
