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
	"github.com/nebulaim/telegramd/base/redis_client"
	"github.com/golang/glog"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"github.com/nebulaim/telegramd/base/base"
	"time"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"sync"
)

// - 简单设计思路(@benqi)
// 	- IM里，在线状态是一种Weak状态，不需要非常的精确，在线状态只需要达到最终一致性要求即可
//   故可以将在线状态存储在redis里
//  - 以一种租约机制设置用户状态的在线状态
//    > 一旦收到account_status，如果不存在，则设置为在线，并为这个key设置一个过期时间
//   如果key存在，则延长过期时间
//  - 回到我们的应用场景，一个客户端对应于一个AuthKeyId，以如下结构存储
//    > key(online_user_id) --> hash { [auth_key_id]_[session_id] : {time, server_id, net_lib_session_id} }
//    > - key为用户Id
//    > - value为hash结构，实际存储值通过连接符'_'拼装
//    >   - hash里的field为<auth_key_id, session_id>二元组
//    >   - field的value为<time, server_id, net_lib_session_id>三元组
//  - why？？
//    > telegram最大的一个特色是允许多端漫游，由auth_key_id标识一个客户端设备，连接到服务端以后，由auth_key_id+session_id对应于
//    > telegram实例（https://core.telegram.org/mtproto/description），一个用户可能会有多个由auth_key_id，而且每个由auth_key_id
//    > 也可能会有多个session_id，所以以<auth_key_id, session_id>二元组做为字段唯一标识客户端实例已经上线，但这样会存在如下问题：
//    > - hash没有指定某一个field设置过期时间，会导致已经下线的客户端还存在field里，故field的value里要缓存一个time时间，由业务层通过
//    >   过期时间来判断是否已经失效
//  - 运维需要注意的地方
//    - 各服务器要对时，服务器时间误差不能太大
//    - ......
// auth_key_id ->
const (
	ONLINE_TIMEOUT = 15  			// 15秒
	CHECK_ONLINE_TIMEOUT = 17  		// 17秒, 15+2秒的误差
	onlineKeyPrefix = "online"		//
)

//var p1, p2 struct {
//	Title  string `redis:"title"`
//	Author string `redis:"author"`
//	Body   string `redis:"body"`
//}

type SessionStatus struct {
	UserId			int32		//

	AuthKeyId 		int64		//
	SessionId		int64		//

	ServerId		int32		// ServerId
	NetlibSessionId	int64		// 网络库SessionID，不是
	Now				int64		// 上报时间
}

func (status *SessionStatus) ToKey() string {
	return fmt.Sprintf("%s_%d", onlineKeyPrefix, status.UserId)
}

func (status *SessionStatus) ToField() string {
	return fmt.Sprintf("%d@%d", status.AuthKeyId, status.SessionId)
}

func (status *SessionStatus) ToValue() string {
	return fmt.Sprintf("%d@%d@%d", status.ServerId, status.NetlibSessionId, status.Now)
}

func (status *SessionStatus) FromKeyValue(k, v string) (err error) {
	// TODO(@benqi): 检查数据合法性
	ks := strings.Split(k, "@")
	if len(ks) != 2 {
		err = fmt.Errorf("FromKeyValue - Invalid field: %s", k)
		return
	}

	if status.AuthKeyId, err = base.StringToInt64(ks[0]); err != nil {
		return
	}
	if status.SessionId, err = base.StringToInt64(ks[1]); err != nil {
		return
	}

	vs := strings.Split(v, "@")
	if len(vs) != 3 {
		err = fmt.Errorf("FromKeyValue - Invalid value: %s", v)
		return
	}

	if status.ServerId, err = base.StringToInt32(vs[0]); err != nil {
		return
	}
	if status.NetlibSessionId, err = base.StringToInt64(vs[1]); err != nil {
		return
	}
	if status.Now, err = base.StringToInt64(vs[2]); err != nil {
		return
	}

	return
}

type onlineStatusModel struct {
	// redis* redis_client.RedisPool
}

var (
	statusInstance *onlineStatusModel
	statusInstanceOnce sync.Once
)

func GetOnlineStatusModel() *onlineStatusModel {
	statusInstanceOnce.Do(func() {
		statusInstance = &onlineStatusModel{}
	})
	return statusInstance
}

func (s *onlineStatusModel) SetOnline(status *SessionStatus) (err error) {
	redis_client := redis_client.GetRedisClient(dao.CACHE)

	conn := redis_client.Get()
	defer conn.Close()

	// 设置键盘
	if _, err = conn.Do("HSET", status.ToKey(), status.ToField(), status.ToValue()); err != nil {
		glog.Errorf("SetOnline - HSET {%v}, error: %s", status, err)
		return
	}

	if _, err = conn.Do("EXPIRE", status.ToKey(), ONLINE_TIMEOUT); err != nil {
		glog.Errorf("SetOnline - EXPIRE {%v}, error: %s", status, err)
		return
	}
	return
}

func (s *onlineStatusModel) SetOffline(status *SessionStatus) (err error) {
	// 设置离线将Now减少CHECK_ONLINE_TIMEOUT
	now := status.Now
	status.Now -= CHECK_ONLINE_TIMEOUT
	defer func() {
		status.Now = now
	}()

	err = s.SetOnline(status)
	return
}

func (s *onlineStatusModel) getOnline(conn redis.Conn, userId int32) (statusList []*SessionStatus, err error) {
	// 设置键盘
	fmt.Printf("%s_%d\n", onlineKeyPrefix, userId)
	m, err := redis.StringMap(conn.Do("HGETALL", fmt.Sprintf("%s_%d", onlineKeyPrefix, userId)));
	if err != nil {
		glog.Errorf("GetOnlineByUserId - HGETALL {online_%d}, error: %s", userId, err)
		return
	}

	fmt.Println(m)
	for k, v := range m {
		status := &SessionStatus{}
		status.UserId = userId
		if err2 := status.FromKeyValue(k, v); err2 != nil {
			glog.Errorf("GetOnlineByUserId - FromKeyValue {online_%d} error: {k: %s, v: %s}, error: %s", userId, k, v)
			continue
		}

		if time.Now().Unix() < status.Now + CHECK_ONLINE_TIMEOUT {
			statusList = append(statusList, status)
			fmt.Println(status)
		}
	}
	return
}

func (s *onlineStatusModel) GetOnlineByUserId(userId int32) ([]*SessionStatus, error) {
	redis_client := redis_client.GetRedisClient(dao.CACHE)

	conn := redis_client.Get()
	defer conn.Close()

	return s.getOnline(conn, userId)
}

// TODO(@benqi): 优化
// 取多个用户的状态信息，可以使用lua脚本，避免多次请求
// eval "local rst={}; for i,v in pairs(KEYS) do rst[i]=redis.call('hgetall', v) end; return rst" 2 user:1 user:2
func (s *onlineStatusModel) GetOnlineByUserIdList(userIdList []int32) (statusList []*SessionStatus, err error) {
	redis_client := redis_client.GetRedisClient(dao.CACHE)

	conn := redis_client.Get()
	defer conn.Close()

	for _, userId := range userIdList {
		ss, err := s.getOnline(conn, userId)
		if err != nil {
			glog.Errorf("GetOnlineByUserIdList - HGETALL {online_%d}, error: %s", userId, err)
			return nil, err
		}

		statusList = append(statusList, ss...)
	}
	return
}
