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

package chat_test

import (
	"net"
	"google.golang.org/grpc"
	"context"
	"sync"
	"errors"
	"github.com/golang/glog"
)

type RpcChatTestServer struct {
	mu   sync.RWMutex
	buf  map[string]chan *ChatMessage
}

func NewChatTestServer() *RpcChatTestServer {
	return &RpcChatTestServer{
		buf:  make(map[string]chan *ChatMessage),
	}
}

func (s *RpcChatTestServer) withReadLock(f func()) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f()
}

func (s *RpcChatTestServer) withWriteLock(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f()
}

func (s *RpcChatTestServer) unsafeExpire(sid string) {
	if buf, ok := s.buf[sid]; ok {
		close(buf)
	}
	delete(s.buf, sid)
}

func (s *RpcChatTestServer) Connect(request *Session, stream ChatTest_ConnectServer) (err error) {
	var buf  chan *ChatMessage = make(chan *ChatMessage, 1000)

	s.withWriteLock(func() {
		if _, ok := s.buf[request.SessionId]; ok {
			err = errors.New("already connected")
			return
		}
		s.buf[request.SessionId] = buf
	})

	if err != nil {
		return err
	}

	defer s.withWriteLock(func() { s.unsafeExpire(request.SessionId) })

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-buf:
			if err := stream.Send(event); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *RpcChatTestServer) SendChat(ctx context.Context, request *ChatMessage) (reply *VoidRsp, err error) {
	if len(request.MessageData) == 0 {
		return nil, errors.New("message must be not empty")
	}
	if len(request.MessageData) > 140 {
		return nil, errors.New("message must be less than or equal 140 characters")
	}

	go s.withReadLock(func() {
		glog.Infof("Log message={%v}", request)
		for _, buf := range s.buf {
			buf <- request
		}
	})
	return &VoidRsp{}, nil
}

func DoChatServe() {
	lis, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		panic(err)
		// glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterChatTestServer(grpcServer, NewChatTestServer())
	grpcServer.Serve(lis)
}
