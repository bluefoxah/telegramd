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
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/base/redis_client"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/sync_server/rpc"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")

}

type RpcServerConfig struct {
	Addr string
}

type SyncServerConfig struct{
	Server 		*RpcServerConfig
	Redis 		[]redis_client.RedisConfig
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	syncServerConfig := &SyncServerConfig{}
	if _, err := toml.DecodeFile("./sync_server.toml", syncServerConfig); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}

	glog.Info(syncServerConfig)

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(syncServerConfig.Redis)

	// 初始化redis_dao、mysql_dao
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	//// TODO(@benqi): 配置驱动
	//redisConfig := &redis_client.RedisConfig{
	//	Name: "test",
	//	Addr: "127.0.0.1:6379",
	//	Idle: 100,
	//	Active: 100,
	//	DialTimeout: 1000000,
	//	ReadTimeout: 1000000,
	//	WriteTimeout: 1000000,
	//	IdleTimeout: 15000000,
	//	DBNum: "0",
	//	Password: "",
	//}
	//
	//redisPool := redis_client.NewRedisPool(redisConfig)
	//onlineModel := model.NewOnlineStatusModel(redisPool)

	lis, err := net.Listen("tcp", syncServerConfig.Server.Addr)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	zproto.RegisterRPCSyncServer(grpcServer, rpc.NewSyncService())

	glog.Info("NewRPCServer in 0.0.0.0:10002.")
	grpcServer.Serve(lis)
}
