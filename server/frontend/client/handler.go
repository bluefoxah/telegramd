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
)

func (c *Client) onMsgsAck(request *TLMsgsAck) {
	glog.Info("processMsgsAck - request: %s", request.String())
}

func (c *Client) onPing(request *EncryptedMessage2) (TLObject) {
	ping, _ := request.Object.(*TLPing)
	glog.Info("processPing - request data: ", ping.String())

	pong := &TLPong{
		PingId: ping.PingId,
	}

	return pong
}

func (c *Client) onPingDelayDisconnect(request *EncryptedMessage2) (TLObject) {
	pingDelayDissconnect, _ := request.Object.(*TLPingDelayDisconnect)
	glog.Info("processPingDelayDisconnect - request data: ", pingDelayDissconnect.String())

	pong := &TLPong{
		PingId: pingDelayDissconnect.PingId,
	}

	return pong
}

func (c *Client) onDestroySession(request *EncryptedMessage2) (TLObject) {
	destroySession, _ := request.Object.(*TLDestroySession)
	glog.Info("processDestroySession - request data: ", destroySession.String())

	// TODO(@benqi): 实现destroySession处理逻辑
	destroy_session_ok := &TLDestroySessionOk{
		SessionId: destroySession.SessionId,
	}
	return destroy_session_ok
}

func (c *Client) onGetFutureSalts(request *EncryptedMessage2) (TLObject) {
	getFutureSalts, _ := request.Object.(*TLGetFutureSalts)
	glog.Info("processGetFutureSalts - request data: ", getFutureSalts.String())

	// TODO(@benqi): 实现getFutureSalts处理逻辑
	futureSalts := &TLFutureSalts{
	}

	return futureSalts
}

func (c *Client) onRpcDropAnswer(request *EncryptedMessage2) (TLObject) {
	rpcDropAnswer, _ := request.Object.(*TLRpcDropAnswer)
	glog.Info("processRpcDropAnswer - request data: ", rpcDropAnswer.String())

	// TODO(@benqi): 实现rpcDropAnswer处理逻辑

	return nil
}

func (c *Client) onContestSaveDeveloperInfo(request *EncryptedMessage2) (TLObject) {
	contestSaveDeveloperInfo, _ := request.Object.(*TLContestSaveDeveloperInfo)
	glog.Info("processGetFutureSalts - request data: ", contestSaveDeveloperInfo.String())

	// TODO(@benqi): 实现scontestSaveDeveloperInfo处理逻辑
	r := &TLTrue{}

	return r
}

func (c *Client) onInvokeWithLayer(request *EncryptedMessage2) error {
	invokeWithLayer, _ := request.Object.(*TLInvokeWithLayer)
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

	dbuf = NewDecodeBuf(initConnection.Query)
	query := dbuf.Object()
	if query == nil {
		return fmt.Errorf("Decode query error: %s", hex.EncodeToString(invokeWithLayer.Query))
	}

	return nil
}

func (c *Client) onInvokeAfterMsg(request *EncryptedMessage2) error {
	invokeAfterMsg, _ := request.Object.(*TLInvokeAfterMsg)
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

func (c *Client) onMsgContainer(request *EncryptedMessage2) error {
	msgContainer, _ := request.Object.(*TLMsgContainer)
	glog.Info("processMsgContainer - request data: ", msgContainer.String())

	return nil
}

func (c *Client) onGzipPacked(request *EncryptedMessage2) error {
	gzipPacked, _ := request.Object.(*TLGzipPacked)
	glog.Info("processGzipPacked - request data: ", gzipPacked.String())

	return nil
}
