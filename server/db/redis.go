package db

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func NewRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := Redis.Ping(context.TODO()).Result(); err != nil {
		panic(err)
	}
}

func NewRedisStoreCode(ctx context.Context, Redis *redis.Client) *redisstore.RedisStore {
	storeCode, _ := redisstore.NewRedisStore(ctx, Redis)
	storeCode.Options(sessions.Options{
		MaxAge: 300, // 设置过期时间为5分钟
	})
	return storeCode
}

func RClose() error {
	err := Redis.Close()
	if err != nil {
		return err
	}
	return nil
}
