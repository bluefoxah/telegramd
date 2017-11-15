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
	server2 "github.com/nebulaim/telegramd/frontend/server"
	"github.com/nebulaim/telegramd/frontend/rpc"
	"flag"
	"github.com/nebulaim/telegramd/base/mysql_client"
	"github.com/nebulaim/telegramd/base/redis_client"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type ServerConfig struct {
	Addr string
}

type RpcClientConfig struct {
	ServiceName string
	Addr string
}

type FrontendConfig struct{
	ServerId 	int32			// 服务器ID
	Server 		*ServerConfig
	BizRpcClient	*RpcClientConfig
	SyncRpcClient	*RpcClientConfig
	Mysql		[]mysql_client.MySQLConfig
	Redis 		[]redis_client.RedisConfig
}

func main() {
	flag.Parse()

	frontendConfig := &FrontendConfig{}
	if _, err := toml.DecodeFile("./frontend.toml", frontendConfig); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info(frontendConfig)

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(frontendConfig.Redis)
	mysql_client.InstallMysqlClientManager(frontendConfig.Mysql)

	// 初始化redis_dao、mysql_dao
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	server := server2.NewServer(frontendConfig.Server.Addr)
	rpc_client, _ := rpc.NewRPCClient(frontendConfig.BizRpcClient.Addr)
	sync_rpc_client, _ := rpc.NewSyncRPCClient(frontendConfig.SyncRpcClient.Addr)
	server.Serve(rpc_client, sync_rpc_client)
}
