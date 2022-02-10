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
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

// Check function searches for a key in the database, creates the key
func Check(rClient *redis.Client, key string) {
	key = strings.ToUpper(key)
	_, err := rClient.Get(ctx, key).Result()
	if err != nil {
		rClient.Set(ctx, key, 0, 0)
	}
}

// Increment function increases the value for a given key.
func Increment(rClient *redis.Client, key string) {
	key = strings.ToUpper(key)
	Check(rClient, key)
	rClient.Incr(ctx, key)
}

// Decrement function decrements the value for a given key.
func Decrement(rClient *redis.Client, key string) {
	key = strings.ToUpper(key)
	Check(rClient, key)
	val, _ := rClient.Get(ctx, key).Result()
	v, _ := strconv.Atoi(val)
	if int(v) > 0 {
		rClient.Decr(ctx, key)
	}
}
