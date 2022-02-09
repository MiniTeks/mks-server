package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var (
	client *redis.Client
	mr     *miniredis.Miniredis
)
var (
	key = "KEY"
	val = "1"
)

func TestMain(m *testing.M) {
	var err error
	mr, err = miniredis.Run()
	if err != nil {
		log.Fatalf("An error '%s' occured while running redis-server", err)
	}
	defer mr.Close()
	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	code := m.Run()
	os.Exit(code)
}

// test redis is down
func TestRedisCanBePinged(t *testing.T) {
	t.Log("Ping redis-db")
	assert.True(t, RedisIsAvailable(client))
}

// This would be your production code.
func RedisIsAvailable(client *redis.Client) bool {
	return client.Ping(context.Background()).Err() == nil
}

func TestIncrement(t *testing.T) {
	mr.Set(key, val)
	t.Log("Calling increment on key")
	Increment(client, key)
	t.Log("Fetching value of key after increment")
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("Error occured fething value of '%s'", key)
	} else if val != "2" {
		t.Fatalf("Expected \"2\" but found %s", val)
	}
}

func TestDecrement(t *testing.T) {
	mr.Set(key, val)
	t.Log("Calling decreament on key")
	Decrement(client, key)
	t.Log("Fetching value of key after decreament")
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("Error occured fething value of '%s'", key)
	} else if val != "0" {
		t.Fatalf("Expected \"2\" but found %s", val)
	}
}

func TestCheck(t *testing.T) {
	// a key which doesn't yet exist  in the database
	key = "NEWK"
	Check(client, key)
	_, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("Error occured fething value of '%s'", key)
	}
}
