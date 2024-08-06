package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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
