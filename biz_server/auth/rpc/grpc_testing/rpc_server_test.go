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
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"testing"
)

type TestRPCServer struct {
}

func (s *TestRPCServer)AuthSentCode(ctx context.Context,  sendCode *TLAuthSendCode) (*Auth_SentCode, error) {
	glog.Infof("Recive AuthSentCode query: {%v}", sendCode)

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
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterAuthServer(grpcServer, &TestRPCServer{})
	grpcServer.Serve(lis)

}
