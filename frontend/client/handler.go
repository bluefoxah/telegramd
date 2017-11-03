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
	"time"
	"github.com/nebulaim/telegramd/frontend/id"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

func (c *Client) onMsgsAck(msgId int64, seqNo int32, request TLObject) {
	glog.Info("processMsgsAck - request: %s", request.String())
}

func (c *Client) onNewSessionCreated(sessionId, msgId int64, seqNo int32) (*TLNewSessionCreated) {
	// glog.Info("processMsgsAck - request: %s", request.String())

	// TODO(@benqi): 客户端保存的initConnection信息推到后台服务存储
	authSaltsDO, err := c.authSaltsDAO.SelectByAuthId(c.Codec.AuthKeyId)
	if err == nil {
		// TODO(@benqi): 处理数据库出错
		glog.Error("c.authSaltsDAO.SelectByAuthId - query error: ", err)
		return nil
	}

	if authSaltsDO == nil {
		// salts不存在，插入一条记录
		authSaltsDO = &dataobject.AuthSaltsDO{ AuthId: c.Codec.AuthKeyId, Salt: id.NextId() }
		_, err := c.authSaltsDAO.Insert(authSaltsDO)
		if err != nil {
			glog.Error("c.authSaltsDAO.Insert - insert error: ", err)
			return nil
		}
	} else {
		// TODO(@benqi): salts是否已经过有效期
	}

	// c.Codec.SessionId =
	notify := &TLNewSessionCreated{
		FirstMsgId: msgId,
		UniqueId: id.NextId(),
		ServerSalt: authSaltsDO.Salt,
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

func (c *Client) onDestroyAuthKey(msgId int64, seqNo int32, request TLObject) (TLObject) {
	destroyAuthKey, _ := request.(*TLDestroyAuthKey)
	glog.Info("onDestroyAuthKey - request data: ", destroyAuthKey.String())

	destroyAuthKeyRes := &TLDestroyAuthKeyOk{}
	return destroyAuthKeyRes
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
	do, err := c.authsDAO.SelectConnectionHashByAuthId(c.Codec.AuthKeyId)
	if err != nil {
		glog.Errorf("c.authsDAO.SelectConnectionHashByAuthId: query db error: %s", err)
		return err
	}

	if do == nil {
		do = &dataobject.AuthsDO{
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
		_, err := c.authsDAO.Insert(do)
		if err != nil {
			glog.Errorf("c.authsDAO.SelectConnectionHashByAuthId: query db error: %s", err)
			return err
		}
	} else {
		// TODO(@benqi): 更新initConnection信息
	}

	dbuf = NewDecodeBuf(initConnection.Query)
	query := dbuf.Object()
	if query == nil {
		return fmt.Errorf("Decode query error: %s", hex.EncodeToString(invokeWithLayer.Query))
	}

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
