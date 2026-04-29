package queue

import (
	"strconv"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func ConnectQueue(RedisAddr, RedisPassword string, RedisDB string) {
	db, _ := strconv.Atoi(RedisDB)
	rdb = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       db,
	})
}
