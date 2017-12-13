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

package rpc

import (
	"flag"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"google.golang.org/grpc"
	"net"
)

func DoMainServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	mtproto.RegisterRPCBotsServer(grpcServer, &BotsServiceImpl{})
	grpcServer.Serve(lis)
}
