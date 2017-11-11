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
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/nebulaim/telegramd/zproto"
	"io"
	"context"
	"time"
	"github.com/nebulaim/telegramd/mtproto"
	net2 "github.com/nebulaim/telegramd/net"
)

type SyncRPCClient struct {
	client zproto.RPCSyncClient
}

func NewSyncRPCClient(target string) (c *SyncRPCClient, err error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	c = &SyncRPCClient{}
	c.client = zproto.NewRPCSyncClient(conn)
	return
}

// TODO(@benqi): 可能有问题
func (c* SyncRPCClient) RunUpdatesStreamLoop(server *net2.Server) {
	auth := &zproto.ServerAuthReq{}
	auth.ServerId = 1
	auth.ServerName = "frontend"

	// TODO(@benqi): 简单等待10s
	for {
		stream, err := c.client.PushUpdatesStream(context.Background(), auth)
		if err != nil {
			glog.Errorf("PushUpdatesStream:", err)
			time.Sleep(10 * time.Second)
			continue
		}


		for {
			update, err := stream.Recv()
			if err == io.EOF {
				time.Sleep(10 * time.Second)
				break
			}
			if err != nil {
				glog.Fatalf("%v.PushUpdatesStream(_) = _, %v", update, err)
				time.Sleep(10 * time.Second)
				break
			}

			// TODO(@benqi): 这是一种简单粗暴的实现方式
			dbuf := mtproto.NewDecodeBuf(update.RawData)
			o := dbuf.Object()
			sendBySessionID(server, update.NetlibSessionId, o)
		}
	}
}

// TODO(@benqi): 使用chan
func sendBySessionID(server *net2.Server, sessionId int64, message mtproto.TLObject)  {
	session := server.GetSession(uint64(sessionId))
	if session != nil {
		m := &mtproto.EncryptedMessage2{
			NeedAck : false,
			Object:   message,
		}
		session.Send(m)
	} else {
		glog.Errorf("SendBySessionID - can't found sessionId: %d", sessionId)
	}
}
