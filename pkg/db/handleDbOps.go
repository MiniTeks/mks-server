package db

import (
	"strconv"

	"github.com/go-redis/redis"
)

func Check(rClient *redis.Client, key string) error {
	_, err := rClient.Get(key).Result()
	return err
}

func Increment(rClient *redis.Client, key string) {
	rClient.Incr(key)
}
func Decrement(rClient *redis.Client, key string) {
	val, _ := rClient.Get(key).Result()
	v, _ := strconv.Atoi(val)
	if int(v) > 0 {
		rClient.Decr(key)
	}
}
