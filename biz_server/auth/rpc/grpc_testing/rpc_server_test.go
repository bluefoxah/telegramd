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

package grpc_testing

import (
	"flag"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"testing"
	"google.golang.org/grpc/metadata"
	"fmt"
)

type TestRPCServer struct {
}

func (s *TestRPCServer)AuthSentCode(ctx context.Context,  sendCode *TLAuthSendCode) (*Auth_SentCode, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("Recive AuthSentCode: md: {%v}, query: {%v}\n", md, sendCode)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	return &Auth_SentCode{
		Payload: &Auth_SentCode_AuthSentCode{
			AuthSentCode: & TLAuthSentCode{
				Flags: 1,
				PhoneRegistered: false,
				Type: nil,
				PhoneCodeHash: "12345",
				NextType: nil,
				Timeout: 1,
			},
		},
	}, nil
}

func TestRPCServer2(t *testing.T) {
	flag.Parse()
	lis, err := net.Listen("tcp", ("localhost:10001"))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterAuthServer(grpcServer, &TestRPCServer{})
	grpcServer.Serve(lis)

}
