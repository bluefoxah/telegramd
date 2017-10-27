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

package server

import (
	"github.com/golang/glog"
	net2 "github.com/nebulaim/telegramd/net"
	. "github.com/nebulaim/telegramd/mtproto"
	"net"
	"github.com/nebulaim/telegramd/frontend/rpc"
	"github.com/nebulaim/telegramd/frontend/client"
	"github.com/nebulaim/telegramd/frontend/auth_key"
)

type Server struct {
	cacheKeys	*auth_key.AuthKeyCacheManager
	Server      *net2.Server
}

func NewServer(addr, dsn string) (s *Server) {
	mtproto := NewMTProto()
	lsn := listen("server", "0.0.0.0:12345")
	server := net2.NewServer(lsn, mtproto, 1024, net2.HandlerFunc(emptySessionLoop))

	s = &Server{
		Server: 	server,
		cacheKeys:  auth_key.NewAuthKeyCacheManager(dsn),
	}
	return
}

func (s* Server) Serve(rpcClient *rpc.RPCClient) {
	glog.Info("Serve...")

	for {
		session, err := s.Server.Accept2()
		if err != nil {
			glog.Error(err)
		}
		glog.Info("a new client ", session.ID())

		c := client.NewClient(session, rpcClient)

		// 使用很土的办法，注入cacheKeys
		c.Codec.AuthKeyStorager = s.cacheKeys
		go s.sessionLoop(c)
	}

	// s.Server.Serve()
}

func (s* Server) sessionLoop(c *client.Client) {
	// client := client.NewClient(c)
	// .Info("NewClient, sessionId: ", session.ID(), ", addr: ", client.RemoteAddr)

	for {
		// 接收数据包
		msg, err := c.Session.Receive()
		if err != nil {
			glog.Error(err)
			return
		}

		if msg == nil {
			glog.Errorf("Recv nil msg, err: ", err)
			return
		}

		// glog.Info("Recved mtproto message!! ", msg)
		// mtprotoMessage2, ok := msg.(*MTProtoMessage)
		// if !ok {
		// 	glog.Info("mtprotoMessage error!")
		// 	return
		// }
		// mtprotoMessage := &mtprotoMessage2
		//m1, ok1 := msg.(EncryptedMessage2)
		//m2, _ := msg.(UnencryptedMessage)
 		if c.Codec.State == CODEC_CONNECTED {
			switch msg.(type) {
			case *EncryptedMessage2:
				// 第一个包
				// Encrypted
				// 第一个包为加密包，则可推断出key已经创建
				c.Codec.State = CODEC_AUTH_KEY_OK
			case *UnencryptedMessage:
				m, _ := msg.(*UnencryptedMessage)
				switch m.Object.(type) {
				case *TLReqPq:
					c.Codec.State = CODEC_req_pq
				default:
					// 未加密第一个包不是TL_req_pq，那么能推断出是RPC消息，key也已经创建
					// Encrypted
					c.Codec.State = CODEC_AUTH_KEY_OK
				}
			default:
				// 不可能发生
				glog.Errorf("Unknown error");
				return
			}
		}

		switch c.Codec.State {
		case CODEC_req_pq,
			 CODEC_resPQ,
			 CODEC_req_DH_params,
			 CODEC_server_DH_params_ok,
			 CODEC_server_DH_params_fail,
			 CODEC_set_client_DH_params,
			 CODEC_dh_gen_ok,
			 CODEC_dh_gen_retry,
			 CODEC_dh_gen_fail:

			m, _ := msg.(*UnencryptedMessage)
			err = c.OnHandshake(m)
			if err != nil {
				return
			}

		case CODEC_AUTH_KEY_OK:
			switch msg.(type) {
			case *EncryptedMessage2:
				m, _ := msg.(*EncryptedMessage2)
				err = c.OnEncryptedMessage(m)
			case *UnencryptedMessage:
				m, _ := msg.(*UnencryptedMessage)
				err = c.OnUnencryptedMessage(m)
			}

			if err!= nil {
				return
			}

		default:
			glog.Errorf("Invalid state: ", c.Codec.State)
			return
		}
	}
}

func emptySessionLoop(session *net2.Session) {
}

// TODO(@benqi): 移植到API层
func listen(who, addr string) net.Listener {
	var lsn net.Listener
	var err error

	lsn, err = net.Listen("tcp", addr)

	if err != nil {
		glog.Fatal("setup ", who, " listener at ", addr, " failed - ", err)
	}

	lsn, _ = Listen(func() (net.Listener, error) {
		return lsn, nil
	})

	glog.Info("setup ", who, " listener at - ", lsn.Addr())
	return lsn
}
