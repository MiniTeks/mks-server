// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
