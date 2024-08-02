package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Mysql *gorm.DB

func NewMysql() {
	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		})
	// 从配置文件中读取值
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	database := viper.GetString("mysql.database")

	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, database)
	ret, err := gorm.Open(mysql.Open(my), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	Mysql = ret

}
func MClose() error {
	db, err := Mysql.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
