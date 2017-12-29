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
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/base/mysql_client"
	"github.com/nebulaim/telegramd/auth_session_server/rpc"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type RpcServerConfig struct {
	Addr string
}

type AuthSessionConfig struct {
	Server    RpcServerConfig
	Discovery service_discovery.ServiceDiscoveryServerConfig
	Mysql     []mysql_client.MySQLConfig
	Redis     []redis_client.RedisConfig
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	authSessionConfig := &AuthSessionConfig{}
	if _, err := toml.DecodeFile("./auth_session.toml", authSessionConfig); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info("%v", authSessionConfig)

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(authSessionConfig.Redis)
	mysql_client.InstallMysqlClientManager(authSessionConfig.Mysql)

	// 初始化redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	server := grpc_util.NewRpcServer(authSessionConfig.Server.Addr, &authSessionConfig.Discovery)
	server.Serve(func(s *grpc.Server) {
		zproto.RegisterRPCAuthSessionServer(s, rpc.NewAuthSessionServiceImpl())
	})
}
