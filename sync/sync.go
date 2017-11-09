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
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	sync "github.com/nebulaim/telegramd/sync/rpc"
	"google.golang.org/grpc"
	"net"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/base/redis_client"
	"github.com/nebulaim/telegramd/biz_model/model"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")

}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	// TODO(@benqi): 配置驱动
	redisConfig := &redis_client.RedisConfig{
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

	redisPool := redis_client.NewRedisPool(redisConfig)
	onlineModel := model.NewOnlineStatusModel(redisPool)

	lis, err := net.Listen("tcp", "0.0.0.0:10002")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	zproto.RegisterRPCSyncServer(grpcServer, sync.NewSyncService(onlineModel))

	glog.Info("NewRPCServer in 0.0.0.0:10002.")
	grpcServer.Serve(lis)
}
