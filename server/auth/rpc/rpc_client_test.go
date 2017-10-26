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
	"context"
	"testing"
	"fmt"
)

func TestRPCClient(t *testing.T)  {
	fmt.Println("TestRPCClient...")
	conn, err := grpc.Dial("127.0.0.1:10001", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := NewAuthClient(conn)
	authSendCode := &TLAuthSendCode{}
	// glog.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	auth_SentCode, err := client.AuthSentCode(context.Background(), authSendCode)
	if err != nil {
		fmt.Errorf("%v.AuthSentCode(_) = _, %v: \n", client, err)
	}
	fmt.Printf("%v\n", auth_SentCode)
}
