/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"testing"
	"github.com/nebulaim/telegramd/base/redis_client"
	"time"
)

// var onlineModel *OnlineStatusModel

func init()  {
	_ = &redis_client.RedisConfig{
		Name: "test",
		Addr: "127.0.0.1:6379",
		Idle: 100,
		Active: 100,
		DialTimeout: 1000000,
		ReadTimeout: 1000000,
		WriteTimeout: 1000000,
		IdleTimeout: 15000000,
		DBNum: "0",
		Password: "",
	}

	// redisPool := redis_client.NewRedisPool(redisConfig)
	// onlineModel = NewOnlineStatusModel(redisPool)
}

func TestSetOnline(t *testing.T) {
	_ = &SessionStatus{
		UserId: 1,
		AuthKeyId: 1,
		SessionId: 1,
		ServerId: 1,
		NetlibSessionId: 1,
		Now: time.Now().Unix(),
	}
/*
	onlineModel.SetOnline(s)

	s.SessionId = 2
	s.ServerId = 1
	s.NetlibSessionId = 2
	s.Now = time.Now().Unix()
	onlineModel.SetOnline(s)

	s.SessionId = 2
	s.ServerId = 1
	s.NetlibSessionId = 2
	s.Now = time.Now().Unix()
	onlineModel.SetOnline(s)

	s.AuthKeyId = 2
	s.SessionId = 2
	s.ServerId = 1
	s.NetlibSessionId = 3
	s.Now = time.Now().Unix()
	onlineModel.SetOnline(s)
 */
}

func TestGetOnline(t *testing.T) {
	// statusList, _ := onlineModel.GetOnlineByUserId(1)
	// fmt.Println(statusList)
}
