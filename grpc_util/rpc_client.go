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

package grpc_util

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/nebulaim/telegramd/mtproto"
	"fmt"
	"context"
	"google.golang.org/grpc/metadata"
	"time"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type RPCClient struct {
	conn *grpc.ClientConn
}

func NewRPCClient(target string) (c *RPCClient, err error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	c = &RPCClient{
		conn: conn,
	}
	return
}

// 通用grpc转发器
func (c* RPCClient) Invoke(rpcMetaData *mtproto.RpcMetaData, object mtproto.TLObject) (mtproto.TLObject, error) {
	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err := fmt.Errorf("Invoke error: %v not regist!\n", object)
		return nil, err
	}

	glog.Infof("Invoke - method: {%s}, req: {%v}", t.Method, object)
	r := t.NewReplyFunc()
	// glog.Infof("Invoke - NewReplyFunc: {%v}\n", r)

	var header, trailer metadata.MD
	ctx := metadata.NewOutgoingContext(context.Background(), rpcMetaData.Encode())
	// ctx := context.Background()
	err := c.conn.Invoke(ctx, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))

	// TODO(@benqi): 哪些情况需要断开客户端连接
	if s, ok := status.FromError(err); !ok {
		glog.Errorf("RPC method: %s,  >> %v.Invoke(_) = _, %v: \n", t.Method, c.conn, err)
		return mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR"), nil
	} else {
		switch s.Code() {
		case codes.OK:
			glog.Infof("Invoke - Invoke reply: {%v}\n", r)
			reply, ok := r.(mtproto.TLObject)

			glog.Infof("Invoke %s time: %d", t.Method, (time.Now().Unix() - rpcMetaData.ReceiveTime))

			if !ok {
				err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
				glog.Error(err)
				return nil, err
			}
			return reply, nil
		case codes.Unknown:
			rpcErr, err := mtproto.NewRpcErrorFromMetadata(trailer)
			if err != nil {
				return mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR"), nil
			} else {
				return rpcErr, nil
			}
		default:
			return mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR"), nil
		}
	}
}
