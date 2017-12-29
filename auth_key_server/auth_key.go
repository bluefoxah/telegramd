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

package main

import (
	"github.com/nebulaim/telegramd/base/redis_client"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/golang/glog"
	"flag"
	"github.com/nebulaim/telegramd/grpc_util/service_discovery"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/grpc_util"
	"google.golang.org/grpc"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/auth_key_server/rpc"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/base/mysql_client"
	cache2 "github.com/nebulaim/telegramd/auth_key_server/cache"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type RpcServerConfig struct {
	Addr string
}

type AuthKeyServerConfig struct {
	Server    RpcServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig
	Mysql     []mysql_client.MySQLConfig
	Redis     []redis_client.RedisConfig
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	authKeyServerConfig := &AuthKeyServerConfig{}
	if _, err := toml.DecodeFile("./auth_key.toml", authKeyServerConfig); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info("%v", authKeyServerConfig)

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(authKeyServerConfig.Redis)
	mysql_client.InstallMysqlClientManager(authKeyServerConfig.Mysql)

	// 初始化redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	authKeyServer := grpc_util.NewRpcServer(authKeyServerConfig.Server.Addr, &authKeyServerConfig.Discovery)
	authKeyServer.Serve(func(s *grpc.Server) {
		cache := cache2.NewAuthKeyCacheManager()
		mtproto.RegisterRPCAuthKeyServer(s, rpc.NewAuthKeyService(cache))
		zproto.RegisterRPCAuthKeyStorageServer(s, rpc.NewAuthKeyStorageService(cache))
	})
}
