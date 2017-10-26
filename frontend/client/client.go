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
	"github.com/nebulaim/telegramd/frontend/rpc"
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
	RPCClient *rpc.RPCClient
	Codec   *MTProtoCodec

	RemoteAddr net.Addr
	LocalAddr  net.Addr

	// TODO(@benqi): 移到handshake处理器里
	Nonce []byte			// 每连接缓存客户端生成的Nonce
	ServerNonce []byte		// 每连接缓存服务生成的ServerNonce
	NewNonce []byte
	A *big.Int
	P *big.Int
}

func NewClient(session *net2.Session, rpcClient *rpc.RPCClient) (c *Client) {
	c = &Client{
		Session: 	session,
		RPCClient:	rpcClient,
		Codec:		session.Codec().(*MTProtoCodec),
	}

	c.RemoteAddr = c.Codec.RemoteAddr()
	c.LocalAddr = c.Codec.LocalAddr()

	return c
}

// handshake
func (c *Client) OnHandshake(request *UnencryptedMessage) error {
	var reply TLObject
	switch request.Object.(type) {
	case *TLMsgsAck:
		msg_acks, _ := request.Object.(*TLMsgsAck)
		c.onHandshakeMsgsAck(msg_acks)
		return nil
	case *TLReqPq:
		reply = c.onReqPq(request)
	case *TLReq_DHParams:
		reply = c.onReq_DHParams(request)
	case *TLSetClient_DHParams:
		reply = c.onSetClient_DHParams(request)
	default:
		glog.Errorf("Invalid request!!!!")
		reply = nil
	}

	if reply == nil {
		return errors.New("handshake: process error!")
	}

	m := &UnencryptedMessage{
		NeedAck : false,
		Object:	  reply,
	}

	return c.Session.Send(m)
}

// MsgsAck
func (c *Client) OnUnencryptedMessage(request *UnencryptedMessage) error {
	// var rspObject TLObject
	switch request.Object.(type) {
	case *TLMsgsAck:
		msg_acks, _ := request.Object.(*TLMsgsAck)
		c.onMsgsAck(request.MessageId, 0, msg_acks)
	default:
		glog.Info("processUnencryptedMessage - Recv authKey created message: ", *request)
	}
	return nil
}

func (c *Client) OnEncryptedMessage(request *EncryptedMessage2) error {
	return c.OnMessage(request.MessageId, request.SeqNo, request.Object)
}

func (c *Client) OnMessage(msgId int64, seqNo int32, request TLObject) error {
	var reply TLObject
	var err error

	switch request.(type) {
	case *TLPing:
		reply = c.onPing(msgId, seqNo, request)
	case *TLPingDelayDisconnect:
		reply = c.onPingDelayDisconnect(msgId, seqNo, request)
	case *TLDestroySession:
		reply = c.onDestroySession(msgId, seqNo, request)
	case *TLGetFutureSalts:
		reply = c.onGetFutureSalts(msgId, seqNo, request)
	case *TLRpcDropAnswer:
		reply = c.onRpcDropAnswer(msgId, seqNo, request)
	case *TLContestSaveDeveloperInfo:
		reply = c.onContestSaveDeveloperInfo(msgId, seqNo, request)
	case *TLInvokeWithLayer:
		return c.onInvokeWithLayer(msgId, seqNo, request)
	case *TLInvokeAfterMsg:
		return c.onInvokeAfterMsg(msgId, seqNo, request)
	case *TLMsgContainer:
		return c.onMsgContainer(msgId, seqNo, request)
	case *TLGzipPacked:
		return c.onGzipPacked(msgId, seqNo, request)
	default:
		// glog.Error("processEncryptedMessage - Not impl processor")
		// rspObject = nil
		reply, err = c.RPCClient.Invoke(request)
		if err != nil {
			return err
		}
	}

	if reply == nil {
		return errors.New("processEncryptedMessage - process error!")
	}

	m := &EncryptedMessage2{
		NeedAck : false,
		SeqNo:	  seqNo,
		Object:   reply,
	}

	return c.Session.Send(m)
}