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

package client

import (
	net2 "github.com/nebulaim/telegramd/net"
	. "github.com/nebulaim/telegramd/mtproto"
	"net"
	"github.com/golang/glog"
	"math/big"
)
//CODEC_UNKNOWN = iota
//CODEC_req_pq
//CODEC_resPQ
//CODEC_req_DH_params
//CODEC_Server_DH_Params_OK
//CODEC_Server_DH_Params_FAILED
//CODEC_set_client_DH_params
//CODEC_dh_gen_ok
//CODEC_dh_gen_retry
//CODEC_dh_gen_fail
//CODEC_AUTH_KEY_OK
//CODEC_ERROR

type Client struct {
	Session *net2.Session
	Codec   *MTProtoCodec
	RemoteAddr net.Addr
	LocalAddr  net.Addr

	Nonce []byte			// 每连接缓存客户端生成的Nonce
	ServerNonce []byte		// 每连接缓存服务生成的ServerNonce
	NewNonce []byte
	A *big.Int
	P *big.Int

	// AuthKeyID uint64
	// AuthKey   []byte
	// MessageSessionID int64
}

func NewClient(session *net2.Session) (c *Client) {
	c = &Client{
		Session: 	session,
		Codec:		session.Codec().(*MTProtoCodec),
		// session.Codec().()
	}

	c.RemoteAddr = c.Codec.RemoteAddr()
	c.LocalAddr = c.Codec.LocalAddr()

	return c
}

func (c *Client) ProcessMessage(message *MTProtoMessage) error {
	glog.Info("ProcessMessage")
	//panic(message)
	//// 状态
	//switch message () {
	//case 1:
	//
	//}
	//if c.Codec.State == mtproto.CODEC_req_pq {
	//}

	return nil
}
