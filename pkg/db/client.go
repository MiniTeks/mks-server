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
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// GetResiClient returns a new redis client given the credential.
func GetRedisClient(cred *RClient) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cred.Addr,
		Password: cred.Pass, // no password set
		DB:       cred.Db,   // use default DB 0
	})
	// keys := []string{"mksTaskcreated", "mksTaskdeleted", "mksTaskactive", "mksTaskfailed", "mksTaskRuncreated", "mksTaskRundeleted", "mksTaskRunactive", "mksTaskRunfailed", "mksPipelineRuncreated", "mksPipelineRundeleted", "mksPipelineRunactive", "mksPipelineRunfailed"}
	// for _, item := range keys {
	// 	rdb.Set(item, 0, 0)
	// }
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Couldn't connect to redis-database")
	}
	fmt.Print(ping, err)
	return rdb
}
