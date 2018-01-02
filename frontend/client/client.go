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
	"errors"
	"fmt"
	"github.com/nebulaim/telegramd/frontend/id"
	"time"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/frontend/rpc"
)

type Client struct {
	Session    *net2.Session
	RPCClient  *grpc_util.RPCClient
	Codec      *MTProtoCodec
	RemoteAddr net.Addr
	LocalAddr  net.Addr

	AuthKeyClient     *rpc.AuthKeyRPCClient
	AuthSessionClient *rpc.AuthSessionRPCClient
	Handshake         *HandshakeHandler
	AuthSession		  *AuthSession

	ConnectionType 	  int
}

func NewClient(session *net2.Session, rpcClient *grpc_util.RPCClient, authKeyClient *rpc.AuthKeyRPCClient, authSessionClient *rpc.AuthSessionRPCClient) (c *Client) {
	c = &Client{
		Session:           session,
		RPCClient:         rpcClient,
		AuthKeyClient:     authKeyClient,
		AuthSessionClient: authSessionClient,
		Codec:             session.Codec().(*MTProtoCodec),
		ConnectionType:	   -1,
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
		c.Handshake.onHandshakeMsgsAck(c, msg_acks)
		// c.onHandshakeMsgsAck(msg_acks)
		return nil
	case *TLReqPq:
		reply = c.Handshake.onReqPq(c, request)
		// reply = c.onReqPq(request)
	case *TLReq_DHParams:
		reply = c.Handshake.onReq_DHParams(c, request)
		// reply = c.onReq_DHParams(request)
	case *TLSetClient_DHParams:
		reply = c.Handshake.onSetClient_DHParams(c, request)
		// reply = c.onSetClient_DHParams(request)
	default:
		glog.Errorf("OnHandshake: Invalid request!!!!")
		reply = nil
	}

	if reply == nil {
		return errors.New("OnHandshake: process error!")
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
		glog.Info("OnUnencryptedMessage - Recv authKey created message: ", *request)
	}
	return nil
}

func (c *Client) OnEncryptedMessage(request *EncryptedMessage2) error {
	// NewSessionCreated
	if c.AuthSession == nil {
		c.AuthSession = GetOrCreateSession(c.Codec.AuthKeyId, request.SessionId)
		//if c.AuthSession.Id == 0 {
		//	c.AuthSession.NetlibSessionId = int64(c.Session.ID())
		//	c.AuthSession.Id = request.SessionId
		//	// AuthSession
		//	if c.ConnectionType != -1 {
		//		c.AuthSession.Type = c.ConnectionType
		//		UpdateAuthSession(c.Codec.AuthKeyId, c.AuthSession)
		//	}
		//}
	}

	if c.AuthSession.Id != request.SessionId {
		// Telegram 以session为核心，我们可以将session理解为虚连接，由客户端创建，
		// 一般情况下一个session对应与客户端的运行实例的生命周期，
		// 客户端每次启动时创建一个session，退出时销毁session。
		// 服务端也应该存储session，但如果服务端内存不足需要回收session或者服务端异常丢失session后，
		// 会主动要求客户端重新生成session。
		// session一般情况下只对应一条tcp连接或整个生命周期内多次http请求
		//
		// 首先检查session是否已经创建

		newSessionCreated := c.onNewSessionCreated(request.SessionId, request.MessageId, request.SeqNo)
		if newSessionCreated == nil {
			return fmt.Errorf("onNewSessionCreated error!")
		}

		c.AuthSession.Id = request.SessionId
		c.AuthSession.NetlibSessionId = int64(c.Session.ID())
		if c.ConnectionType != -1 {
			c.AuthSession.Type = c.ConnectionType
			UpdateAuthSession(c.Codec.AuthKeyId, c.AuthSession)
		}

		// TODO(@benqi): remove Codec.SessionId
		c.Codec.SessionId = request.SessionId
		c.Codec.Salt = newSessionCreated.GetServerSalt()

		m := &EncryptedMessage2{
			// NeedAck : false,
			NeedAck : false,
			Object:   newSessionCreated,
		}

		c.Session.Send(m)
	}

	glog.Info("OnEncryptedMessage - sessionId: ", request.SessionId, ", seqNo: ", request.SeqNo, ", messageId: ", request.MessageId)

	// TODO(@benqi): 检查sessionId
	return c.OnMessage(request.MessageId, request.SeqNo, request.Object)
}

func (c *Client) OnMessage(msgId int64, seqNo int32, request TLObject) error {
	if c.Codec.UserId != 0 {
		defer func() {
			if r := recover(); r != nil {
				glog.Error(r)
			}
		}()

		// TODO(@benqi): Genernal或Push才会后设置online，如果Genernal则必须已经注册
		c.setOnline()
	}

	var reply TLObject
	// var err error

	switch request.(type) {
	case *TLDestroyAuthKey:
		reply = c.onDestroyAuthKey(msgId, seqNo, request)
	case *TLPing:
		reply = c.onPing(msgId, seqNo, request)
	case *TLPingDelayDisconnect:
		reply = c.onPingDelayDisconnect(msgId, seqNo, request)
	case *TLMsgsAck:
		// msg_acks, _ := request.Object.(*TLMsgsAck)
		c.onMsgsAck(msgId, seqNo, request)
		return nil
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
		if c.ConnectionType != -1 {
			switch request.(type) {
			case *TLUploadSaveFilePart, *TLUploadSaveBigFilePart, *TLUploadReuploadCdnFile:
				c.AuthSession.Type = UPLOAD
				c.ConnectionType = UPLOAD
				UpdateAuthSession(c.Codec.AuthKeyId, c.AuthSession)
			case *TLUploadGetFile, *TLUploadGetWebFile, *TLUploadGetCdnFile, *TLUploadCdnFileReuploadNeeded:
				c.AuthSession.Type = DOWNLOAD
				c.ConnectionType = DOWNLOAD
				UpdateAuthSession(c.Codec.AuthKeyId, c.AuthSession)
			case *TLHelpGetConfig:
				// TODO(@benqi): 暂时还不能确定是否TEMP
			default:
				c.AuthSession.Type = GENERIC
				c.ConnectionType = GENERIC
				UpdateAuthSession(c.Codec.AuthKeyId, c.AuthSession)
			}
		}

		// glog.Error("processEncryptedMessage - Not impl processor")
		// rspObject = nil

		// TODO(@benqi): [权限判断](https://core.telegram.org/api/auth)
		/*
		 *	Only a small portion of the API methods are available to unauthorized users:
	     *
		 *	- auth.sendCode
		 *	- auth.sendCall
		 *	- auth.checkPhone
		 *	- auth.signUp
		 *	- auth.signIn
		 *	- auth.importAuthorization
		 *	- help.getConfig
		 *	- help.getNearestDc
		 *
		 *	Other methods will result in an error: 401 UNAUTHORIZED.
		 */

		glog.Info("rpc request authId: ", c.Codec.AuthKeyId)
		// TODO(@benqi): 透传UserID

		if c.Codec.UserId == 0 {
			defer func() {
				if r := recover(); r != nil {
					glog.Error(r)
				}
			}()

			userId := c.AuthSessionClient.GetUserIDByAuthKey(c.Codec.AuthKeyId)
			if userId != 0 {
				c.Codec.UserId = userId
			}
			//
			//do := dao.GetAuthUsersDAO(dao.DB_SLAVE).SelectByAuthId(c.Codec.AuthKeyId)
			//glog.Info("SelectByAuthId : ", do)
			//if do != nil {
			//	c.setOnline()
			//}
		}

		// 初始化metadata
		rpcMetadata := &zproto.RpcMetadata{}
		rpcMetadata.ServerId = 1
		rpcMetadata.NetlibSessionId = int64(c.Session.ID())
		rpcMetadata.UserId = c.Codec.UserId
		rpcMetadata.AuthId = c.Codec.AuthKeyId
		rpcMetadata.SessionId = c.Codec.SessionId
		rpcMetadata.ClientAddr = c.Codec.RemoteAddr().String()
		rpcMetadata.TraceId = id.NextId()
		rpcMetadata.SpanId = id.NextId()
		rpcMetadata.ReceiveTime = time.Now().Unix()

		rpcResult, err := c.RPCClient.Invoke(rpcMetadata, request)

		if err != nil {
			return nil
			// return err
		}

		glog.Infof("OnMessage - rpc_result: {%v}\n", rpcResult)
		// 构造rpc_result
		reply = &TLRpcResult{
			ReqMsgId: msgId,
			Result: rpcResult,
		}

		// TODO(@benqi): 协议底层处理
		if _, ok := request.(*TLMessagesSendMedia); ok {
			if _, ok := rpcResult.(*TLRpcError); !ok {
				// TODO(@benqi): 由底层处理，通过多种策略（gzip, msg_container等）来打包并发送给客户端
				m := &MsgDetailedInfoContainer{Message: &EncryptedMessage2{
					NeedAck: false,
					SeqNo:   seqNo,
					Object:  reply,
				}}
				return c.Session.Send(m)
			}
		}
	}

	if reply == nil {
		return errors.New("OnMessage - process error!")
	}

	// TODO(@benqi): 由底层处理，通过多种策略（gzip, msg_container等）来打包并发送给客户端
	m := &EncryptedMessage2{
		NeedAck : false,
		SeqNo:	  seqNo,
		Object:   reply,
	}

	return c.Session.Send(m)
}

func (c *Client) OnClose() {
	if c.AuthSession != nil {
		c.AuthSession.NetlibSessionId = 0
		UpdateAuthSession(c.Codec.AuthKeyId, c.AuthSession)
	}
}
