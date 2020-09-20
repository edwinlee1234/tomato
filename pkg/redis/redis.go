package redis

import (
	"context"
	"tomato/config"

	redis "github.com/go-redis/redis/v8"
)

// Conn Conn
var Conn *redis.Client

// Init Init
func Init() {
	Conn = redis.NewClient(&redis.Options{
		Addr:     config.Val.RedisHost + ":" + config.Val.RedisPort,
		Password: config.Val.RedisPass,
		DB:       0,
	})

	_, err := Conn.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
