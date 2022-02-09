package db

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

func Check(rClient *redis.Client, key string) {
	key = strings.ToUpper(key)
	_, err := rClient.Get(ctx, key).Result()
	if err != nil {
		rClient.Set(ctx, key, 0, 0)
	}
}

func Increment(rClient *redis.Client, key string) {
	key = strings.ToUpper(key)
	Check(rClient, key)
	rClient.Incr(ctx, key)
}

func Decrement(rClient *redis.Client, key string) {
	key = strings.ToUpper(key)
	Check(rClient, key)
	val, _ := rClient.Get(ctx, key).Result()
	v, _ := strconv.Atoi(val)
	if int(v) > 0 {
		rClient.Decr(ctx, key)
	}
}
