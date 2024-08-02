package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() {
	// 初始化 Viper
	viper.SetConfigName("config.yaml")   // 配置文件的名称（不包括后缀）
	viper.SetConfigType("yaml")          // 配置文件的类型（可根据实际情况选择）
	viper.AddConfigPath("./config.yaml") // 配置文件所在的路径（可以添加多个路径）

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 处理读取配置文件时的错误
		fmt.Printf("Failed to read config.yaml file: %v\n", err)
		return
	}
}
