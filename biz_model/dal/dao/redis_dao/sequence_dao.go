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

// TODO(@benqi): 可以使用如下方法来生成seq
// - 如果运维能保证redis数据可靠性，可移除数据库seq_updates_ngen的存储
// - 可使用[seqsvr](https://github.com/nebula-in/seqsvr)服务来生成seq
// - 可调研艺龙的序列号生成器
// - 直接使用etcd或zk
package redis_dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/nebulaim/telegramd/base/base"
	"github.com/nebulaim/telegramd/base/redis_client"
)

const (
	seqUpdatesNgenId = "seq_updates_ngen"
)

type SequenceDAO struct {
	redis *redis_client.RedisPool
	// ngen  *SeqUpdatesNgenDAO
}

func NewSequenceDAO(redis *redis_client.RedisPool) *SequenceDAO {
	return &SequenceDAO{
		redis: redis,
		// ngen:  ngen,
	}
}

// 独立出incr和set的原因
// 在NextID直接获取redis的连接，incr执行完后可能存在操作数据库的大事物
// 有可能会导致redis在一段时间内未释放
// 独立出来后，一旦执行incr或set则立即释放redis连接
func (dao *SequenceDAO) Incr(key string) (seq int64, err error) {
	conn := dao.redis.Get()
	defer conn.Close()

	// 设置键
	seq, err = redis.Int64(conn.Do("INCR", fmt.Sprintf("%s_%s", seqUpdatesNgenId, key)))
	if err != nil {
		// glog.Errorf("NextID - INCR {%d}, error: %s", k, err)
		return
	}

	return
}

func (dao *SequenceDAO) Set(key string, seq int64) (err error) {
	conn := dao.redis.Get()
	defer conn.Close()

	_, err = redis.Bool(conn.Do("SET", fmt.Sprintf("%s_%s", seqUpdatesNgenId, key), base.Int64ToString(seq)))
	if err != nil {
		// glog.Errorf("NextID - SET {%s}, error: %s", k, err)
		return
	}

	return
}

