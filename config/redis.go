package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Rdb *redis.Client

func init() {

	client := redis.NewClient(&redis.Options{
		Addr:     K.Redis.Addr,
		Password: K.Redis.Password, // no password set
		DB:       K.Redis.DB,       // uses default DB
		PoolSize: K.Redis.PoolSize,
	})
	ctx := context.Background()

	ping := client.Ping(ctx)
	if ping.String() == "ping: PONG" {
		log.Println("连接redis 成功!")
		Rdb = client
	}
	log.Println("连接redis失败")

	//panic("连接失败")

}
