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

package client_test

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/nebulaim/telegramd/grpc_util/service_discovery/etcd3"
	"github.com/nebulaim/telegramd/grpc_util/load_balancer"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/base/crypto"
	"github.com/golang/protobuf/ptypes"
	"github.com/nebulaim/telegramd/grpc_util"
	"google.golang.org/grpc/metadata"
	"github.com/nebulaim/telegramd/zproto"
	"fmt"
)

func main() {
	etcdConfg := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	r := etcd3.NewResolver("/nebulaim", "auth_key", etcdConfg)
	b := load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	c, err := grpc.Dial("", grpc.WithInsecure(),  grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()

	client := mtproto.NewRPCAuthKeyClient(c)

	for i := 0; i < 1000; i++ {
		var err error
		var header, trailer metadata.MD
		md := &zproto.RpcMetadata{}
		auth := &zproto.AuthKeyMetadata{}
		md.Extend, err = ptypes.MarshalAny(auth)
		ctx, _ := grpc_util.RpcMetadatToOutgoing(context.Background(), md)
		resp, err := client.ReqPq(ctx, &mtproto.TLReqPq{Nonce: crypto.GenerateNonce(16)}, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(resp)
	}
}
