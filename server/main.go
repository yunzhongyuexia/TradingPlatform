package main

import (
	"server/config"
	"server/db"
	"server/router"
)

func main() {
	config.NewViper()
	db.NewMysql()
	db.NewRedis()
	defer func() {
		_ = db.MClose()
		_ = db.RClose()
	}()
	router.NewRouter()
}

//先部署nft
//truffle compile                     //编译生成build目录和json文件
//truffle migrate --network holesky   //  编译并部署到区块链网络上
