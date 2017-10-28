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
	"github.com/golang/glog"
	"fmt"
	"encoding/hex"
	. "github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/frontend/auth_key"
	"github.com/nebulaim/telegramd/frontend/model"
	"time"
	"github.com/nebulaim/telegramd/frontend/id"
)

func (c *Client) onMsgsAck(msgId int64, seqNo int32, request TLObject) {
	glog.Info("processMsgsAck - request: %s", request.String())
}

func (c *Client) onNewSessionCreated(sessionId, msgId int64, seqNo int32) (*TLNewSessionCreated) {
	// glog.Info("processMsgsAck - request: %s", request.String())

	// TODO(@benqi): 客户端保存的initConnection信息推到后台服务存储
	// 先用最老土的办法实现
	authSessions := &model.AuthSessions{
			AuthId: 	c.Codec.AuthKeyId,
			SessionId:	sessionId,
			UniqueId:   id.NextId(),
		}

	authSalts := &model.AuthSalts{AuthId: 	c.Codec.AuthKeyId,}
	if c.Codec.Salt == 0 {
		authSalts.Salt = id.NextId()
	}

	// 先这样吧
	cacheKey, _ := c.Codec.AuthKeyStorager.(*auth_key.AuthKeyCacheManager)

	// TODO(@benqi): 检查数据库操作是否成功
	cacheKey.ZOrm.ReadOrCreate(authSessions, "AuthId", "SessionId")
	cacheKey.ZOrm.ReadOrCreate(authSalts, "AuthId", "Salt")

	// c.Codec.SessionId =
	notify := &TLNewSessionCreated{
		FirstMsgId: msgId,
		UniqueId:   authSessions.UniqueId,
		ServerSalt: authSalts.Salt,
	}
	return notify
}

func (c *Client) onPing(msgId int64, seqNo int32, request TLObject) (TLObject) {
	ping, _ := request.(*TLPing)
	glog.Info("processPing - request data: ", ping.String())

	pong := &TLPong{
		MsgId: msgId,
		PingId: ping.PingId,
	}

	return pong
}

func (c *Client) onPingDelayDisconnect(msgId int64, seqNo int32, request TLObject) (TLObject) {
	pingDelayDissconnect, _ := request.(*TLPingDelayDisconnect)
	glog.Info("processPingDelayDisconnect - request data: ", pingDelayDissconnect.String())

	pong := &TLPong{
		MsgId: msgId,
		PingId: pingDelayDissconnect.PingId,
	}

	return pong
}

func (c *Client) onDestroySession(msgId int64, seqNo int32, request TLObject) (TLObject) {
	destroySession, _ := request.(*TLDestroySession)
	glog.Info("processDestroySession - request data: ", destroySession.String())

	// TODO(@benqi): 实现destroySession处理逻辑
	destroy_session_ok := &TLDestroySessionOk{
		SessionId: destroySession.SessionId,
	}
	return destroy_session_ok
}

func (c *Client) onGetFutureSalts(msgId int64, seqNo int32, request TLObject) (TLObject) {
	getFutureSalts, _ := request.(*TLGetFutureSalts)
	glog.Info("processGetFutureSalts - request data: ", getFutureSalts.String())

	// TODO(@benqi): 实现getFutureSalts处理逻辑
	futureSalts := &TLFutureSalts{
	}

	return futureSalts
}

func (c *Client) onRpcDropAnswer(msgId int64, seqNo int32, request TLObject) (TLObject) {
	rpcDropAnswer, _ := request.(*TLRpcDropAnswer)
	glog.Info("processRpcDropAnswer - request data: ", rpcDropAnswer.String())

	// TODO(@benqi): 实现rpcDropAnswer处理逻辑

	return nil
}

func (c *Client) onContestSaveDeveloperInfo(msgId int64, seqNo int32, request TLObject) (TLObject) {
	contestSaveDeveloperInfo, _ := request.(*TLContestSaveDeveloperInfo)
	glog.Info("processGetFutureSalts - request data: ", contestSaveDeveloperInfo.String())

	// TODO(@benqi): 实现scontestSaveDeveloperInfo处理逻辑
	r := &TLTrue{}

	return r
}

func (c *Client) onInvokeWithLayer(msgId int64, seqNo int32, request TLObject) error {
	invokeWithLayer, _ := request.(*TLInvokeWithLayer)
	glog.Info("processInvokeWithLayer - request data: ", invokeWithLayer.String())

	// Check api layer
	// if invokeWithLayer.Layer > API_LAYER {
	// 	return fmt.Errorf("Not suppoer api layer: %d", invokeWithLayer.layer)
	// }

	if invokeWithLayer.GetQuery() == nil {
		return fmt.Errorf("invokeWithLayer Query is nil")
	}

	dbuf := NewDecodeBuf(invokeWithLayer.Query)
	classID := dbuf.Int();
	if classID != int32(TLConstructor_CRC32_initConnection) {
		return fmt.Errorf("Not initConnection classID: %d", classID)
	}

	initConnection := &TLInitConnection{}
	err := initConnection.Decode(dbuf)
	if err != nil {
		glog.Error("Decode initConnection error: ", err)
		return err
	}

	// TODO(@benqi): 客户端保存的initConnection信息推到后台服务存储
	// 先用最老土的办法实现
	connectionsModel := &model.AuthConnections{
		AuthId: c.Codec.AuthKeyId,
		ApiId:  initConnection.ApiId,
		DeviceModel: initConnection.DeviceModel,
		SystemVersion: initConnection.SystemVersion,
		AppVersion: initConnection.AppVersion,
		SystemLangCode: initConnection.SystemLangCode,
		LangPack: initConnection.LangPack,
		LangCode: initConnection.LangCode,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 先这样吧
	cacheKey, _ := c.Codec.AuthKeyStorager.(*auth_key.AuthKeyCacheManager)
	cacheKey.ZOrm.InsertOrUpdate(connectionsModel, "auth_id")

	dbuf = NewDecodeBuf(initConnection.Query)
	query := dbuf.Object()
	if query == nil {
		return fmt.Errorf("Decode query error: %s", hex.EncodeToString(invokeWithLayer.Query))
	}

	// 
	c.OnMessage(msgId, seqNo, query)

	return nil
}

func (c *Client) onInvokeAfterMsg(msgId int64, seqNo int32, request TLObject) error {
	invokeAfterMsg, _ := request.(*TLInvokeAfterMsg)
	glog.Info("processInvokeAfterMsg - request data: ", invokeAfterMsg.String())

	if invokeAfterMsg.GetQuery() == nil {
		return fmt.Errorf("invokeAfterMsg Query is nil")
	}

	dbuf := NewDecodeBuf(invokeAfterMsg.Query)
	query := dbuf.Object()
	if query == nil {
		return fmt.Errorf("Decode query error: %s", hex.EncodeToString(invokeAfterMsg.Query))
	}

	return nil
}

func (c *Client) onMsgContainer(msgId int64, seqNo int32, request TLObject) error {
	msgContainer, _ := request.(*TLMsgContainer)
	glog.Info("processMsgContainer - request data: ", msgContainer.String())


	for _, m := range msgContainer.Messages {
		c.OnMessage(m.MsgId, m.Seqno, m.Object)
	}

	return nil
}

func (c *Client) onGzipPacked(msgId int64, seqNo int32, request TLObject) error {
	gzipPacked, _ := request.(*TLGzipPacked)
	glog.Info("processGzipPacked - request data: ", gzipPacked.String())

	return nil
}
