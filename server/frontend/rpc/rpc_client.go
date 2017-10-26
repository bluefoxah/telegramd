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
	"github.com/nebulaim/telegramd/mtproto"
	"fmt"
	"context"
	"errors"
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
func (c* RPCClient) Invoke(object mtproto.TLObject) (mtproto.TLObject, error) {
	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err := fmt.Errorf("Invoke error: %v not regist!\n", object)
		return nil, err
	}

	r := t.NewReplyFunc()
	err := c.conn.Invoke(context.Background(), t.Method, r, r)
	if err != nil {
		fmt.Errorf("%v.Invoke(_) = _, %v: \n", c.conn, err)
	}

	reply, ok := r.(mtproto.TLObject)
	if !ok {
		return reply, nil
	}
	err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
	return nil, err
}
