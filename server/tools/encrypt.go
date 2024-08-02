package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func Encrypt(pwd string) string {
	newPwd := pwd + "云中月下"
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hashBytes := hash.Sum(nil) // 获取md5 hash值
	hashString := hex.EncodeToString(hashBytes)
	fmt.Println("加密后的密码：", hashString)
	return hashString
}

// ECDSAKey 生成并返回一个新的ECDSA私钥的十六进制字符串
func ECDSAKey() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "生成密钥失败"
	}
	// 将私钥的字节表示形式转换为十六进制字符串
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	return privateKeyHex
}
