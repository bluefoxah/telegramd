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
	"errors"
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

// handshake
func (c *Client) OnHandshake(request *UnencryptedMessage) error {
	var rspObject TLObject
	switch request.Object.(type) {
	case *TLMsgsAck:
		msg_acks, _ := request.Object.(*TLMsgsAck)
		c.onHandshakeMsgsAck(msg_acks)
		return nil
	case *TLReqPq:
		rspObject = c.onReqPq(request)
	case *TLReq_DHParams:
		rspObject = c.onReq_DHParams(request)
	case *TLSetClient_DHParams:
		rspObject = c.onSetClient_DHParams(request)
	default:
		glog.Errorf("Invalid request!!!!")
		rspObject = nil
	}

	if rspObject == nil {
		return errors.New("handshake: process error!")
	}

	m := &UnencryptedMessage{
		NeedAck : false,
		Object:		rspObject,
	}

	return c.Session.Send(m)
}

// MsgsAck
func (c *Client) OnUnencryptedMessage(request *UnencryptedMessage) error {
	// var rspObject TLObject
	switch request.Object.(type) {
	case *TLMsgsAck:
		msg_acks, _ := request.Object.(*TLMsgsAck)
		c.onMsgsAck(msg_acks)
	default:
		glog.Info("processUnencryptedMessage - Recv authKey created message: ", *request)
	}
	return nil
}

func (c *Client) OnEncryptedMessage(request *EncryptedMessage2) error {
	var rspObject TLObject

	switch request.Object.(type) {
	case *TLPing:
		rspObject = c.onPing(request)
	case *TLPingDelayDisconnect:
		rspObject = c.onPingDelayDisconnect(request)
	case *TLDestroySession:
		rspObject = c.onDestroySession(request)
	case *TLGetFutureSalts:
		rspObject = c.onGetFutureSalts(request)
	case *TLRpcDropAnswer:
		rspObject = c.onRpcDropAnswer(request)
	case *TLContestSaveDeveloperInfo:
		rspObject = c.onContestSaveDeveloperInfo(request)
	case *TLInvokeWithLayer:
		return c.onInvokeWithLayer(request)
	case *TLInvokeAfterMsg:
		return c.onInvokeAfterMsg(request)
	case *TLMsgContainer:
		return c.onMsgContainer(request)
	case *TLGzipPacked:
		return c.onGzipPacked(request)
	default:
		glog.Error("processEncryptedMessage - Not impl processor")
		rspObject = nil
	}

	if rspObject == nil {
		return errors.New("processEncryptedMessage - process error!")
	}

	m := &EncryptedMessage2{
		NeedAck : false,
		SeqNo:	  request.SeqNo,
		Object:   rspObject,
	}

	return c.Session.Send(m)
}
