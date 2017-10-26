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
	"github.com/nebulaim/telegramd/server/frontend/client"
	"net"
)

type Server struct {
	Server        *net2.Server
}

func NewServer(addr string) (s *Server) {
	mtproto := NewMTProto()
	lsn := listen("server", "0.0.0.0:12345")
	server := net2.NewServer(lsn, mtproto, 1024, net2.HandlerFunc(mtprotoSessionLoop))

	s = &Server{
		Server: server,
	}
	return
}

func (s* Server) Serve() {
	glog.Info("loop")
	s.Server.Serve()
}

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

func mtprotoSessionLoop(session *net2.Session) {
	client := client.NewClient(session)
	glog.Info("NewClient, sessionId: ", session.ID(), ", addr: ", client.RemoteAddr)

	for {
		// 接收数据包
		msg, err := client.Session.Receive()
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
 		if client.Codec.State == CODEC_CONNECTED {
			switch msg.(type) {
			case *EncryptedMessage2:
				// 第一个包
				// Encrypted
				// 第一个包为加密包，则可推断出key已经创建
				client.Codec.State = CODEC_AUTH_KEY_OK
			case *UnencryptedMessage:
				m, _ := msg.(*UnencryptedMessage)
				switch m.Object.(type) {
				case *TLReqPq:
					client.Codec.State = CODEC_req_pq
				default:
					// 未加密第一个包不是TL_req_pq，那么能推断出是RPC消息，key也已经创建
					// Encrypted
					client.Codec.State = CODEC_AUTH_KEY_OK
				}
			default:
				// 不可能发生
				glog.Errorf("Unknown error");
				return
			}
		}

		switch client.Codec.State {
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
			err = client.OnHandshake(m)
			if err != nil {
				return
			}

		case CODEC_AUTH_KEY_OK:
			switch msg.(type) {
			case *EncryptedMessage2:
				m, _ := msg.(*EncryptedMessage2)
				err = client.OnEncryptedMessage(m)
			case *UnencryptedMessage:
				m, _ := msg.(*UnencryptedMessage)
				err = client.OnUnencryptedMessage(m)
			}

			if err!= nil {
				return
			}

		default:
			glog.Errorf("Invalid state: ", client.Codec.State)
			return
		}
	}
}

/*
// MsgsAck
func processUnencryptedMessage(client *client.Client, request *UnencryptedMessage) error {
	// var rspObject TLObject
	switch request.Object.(type) {
	case *TLMsgsAck:
		msg_acks, _ := request.Object.(*TLMsgsAck)
		processMsgsAck(client, msg_acks)
	default:
		glog.Info("processUnencryptedMessage - Recv authKey created message: ", *request)
	}
	return nil
}

func processEncryptedMessage(client *client.Client, request *EncryptedMessage2) error {
	var rspObject TLObject

	switch request.Object.(type) {
	case *TLPing:
		rspObject = processPing(client, request)
	case *TLPingDelayDisconnect:
		rspObject = processPingDelayDisconnect(client, request)
	case *TLDestroySession:
		rspObject = processDestroySession(client, request)
	case *TLGetFutureSalts:
		rspObject = processGetFutureSalts(client, request)
	case *TLRpcDropAnswer:
		rspObject = processRpcDropAnswer(client, request)
	case *TLContestSaveDeveloperInfo:
		rspObject = processContestSaveDeveloperInfo(client, request)
	case *TLInvokeWithLayer:
		return processInvokeWithLayer(client, request)
	case *TLInvokeAfterMsg:
		return processInvokeAfterMsg(client, request)
	case *TLMsgContainer:
		return processMsgContainer(client, request)
	case *TLGzipPacked:
		return processGzipPacked(client, request)
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

	return client.Session.Send(m)
}

// handshake
func processHandshake(client *client.Client, request *UnencryptedMessage) error {
	var rspObject TLObject
	switch request.Object.(type) {
	case *TLMsgsAck:
		msg_acks, _ := request.Object.(*TLMsgsAck)
		processHandshakeMsgsAck(client, msg_acks)
		return nil
	case *TLReqPq:
		rspObject = processReqPq(client, request)
	case *TLReq_DHParams:
		rspObject = processReq_DHParams(client, request)
	case *TLSetClient_DHParams:
		rspObject = processSetClient_DHParams(client, request)
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

	return client.Session.Send(m)
}

func processHandshakeMsgsAck(client *client.Client, request *TLMsgsAck) {
	glog.Info("processHandshakeMsgsAck - request: %s", request.String())
}

func processReqPq(client *client.Client, request *UnencryptedMessage) (TLObject) {
	reqPq, _ := request.Object.(*TLReqPq)
	glog.Info("processReqPq - request data: ", reqPq.String())

	// 检查数据是否合法
	if reqPq.GetNonce() == nil || len(reqPq.GetNonce()) != 16 {
		glog.Errorf("processReqPq - nonce not int128 type")
		return nil
	}

	resPQ := &TLResPQ{}

	// Nonce
	resPQ.Nonce = make([]byte, 16)
	copy(resPQ.Nonce, reqPq.Nonce)

	resPQ.ServerNonce = crypto.GenerateNonce(16)

	// TODO(@benqi): 使用算法生成PQ
	// 这里直接指定了PQ值: {0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81}
	resPQ.Pq = string([]byte{0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81})

	// TODO(@benqi): 预先计算出fingerprint
	// 这里直接使用了0xc3b42b026ce86b21
	// var a = 0xc3b42b026ce86b21
	// var u uint64 = 0xc3b42b026ce86b21
	// var u uint64 = 14404393623918429762
	var u uint64 = 12240908862933197005

	resPQ.ServerPublicKeyFingerprints = []int64{int64(u)}

	// 缓存客户端Nonce
	client.Nonce = reqPq.Nonce
	client.ServerNonce = resPQ.ServerNonce
	client.Codec.State = CODEC_resPQ
	return resPQ
}

func processReq_DHParams(client *client.Client, request *UnencryptedMessage) (TLObject) {
	req_DH_params, _ := request.Object.(*TLReq_DHParams)
	glog.Info("processReq_DHParams - request: ", req_DH_params.String())

	// 客户端传输数据解析
	// Nonce
	if !bytes.Equal(req_DH_params.Nonce, client.Nonce) {
		glog.Info("processReq_DHParams - Wrong Nonce")
		return nil
	}

	// ServerNonce
	if !bytes.Equal(req_DH_params.ServerNonce, client.ServerNonce) {
		glog.Info("processReq_DHParams - Wrong ServerNonce")
		return nil
	}

	var defaultP = []byte{0x49, 0x4C, 0x55, 0x3B}
	var defaultQ = []byte{0x53, 0x91, 0x10, 0x73}
	var u uint64 = 12240908862933197005

	// P
	if !bytes.Equal([]byte(req_DH_params.P), defaultP) {
		glog.Info("processReq_DHParams - Wrong p value")
		return nil
	}

	// Q
	if !bytes.Equal([]byte(req_DH_params.Q), defaultQ) {
		glog.Info("processReq_DHParams - Wrong q value")
		return nil
	}

	if req_DH_params.PublicKeyFingerprint != int64(u) {
		glog.Info("processReq_DHParams - Wrong PublicKeyFingerprint value")
		return nil
	}

	encryptedData := []byte(req_DH_params.EncryptedData)
	glog.Info("EncryptedData: len = ", len(encryptedData), ", data: ", hex.EncodeToString(encryptedData))

	// 1. 解密
	encryptedPQInnerData := rsa.Decrypt([]byte(req_DH_params.EncryptedData))
	sha1_check := sha1.Sum(encryptedPQInnerData[20:])

	glog.Info(hex.EncodeToString(sha1_check[:]))
	glog.Info(hex.EncodeToString(encryptedPQInnerData[:20]))

	// 2. 反序列化出pqInnerData
	pqInnerData := &TLPQInnerData{}
	dbuf := NewDecodeBuf(encryptedPQInnerData[SHA_DIGEST_LENGTH+4:])
	err := pqInnerData.Decode(dbuf)
	if err != nil {
		glog.Errorf("processReq_DHParams - TLPQInnerData decode error: ", err)
		return nil
	}

	glog.Info("processReq_DHParams - pqInnerData Decode sucess: ", pqInnerData.String())

	// 缓存NewNonce
	client.NewNonce = pqInnerData.NewNonce

	// TODO(@benqi): 直接指定了dh2048_p和dh2048_g!!!
	var dh2048_p =[]byte{
		0xF8,0x87,0xA5,0x15,0x98,0x35,0x20,0x1E,0xF5,0x81,0xE5,0x95,
		0x1B,0xE4,0x54,0xEA,0x53,0xF5,0xE7,0x26,0x30,0x03,0x06,0x79,
		0x3C,0xC1,0x0B,0xAD,0x3B,0x59,0x3C,0x61,0x13,0x03,0x7B,0x02,
		0x70,0xDE,0xC1,0x20,0x11,0x9E,0x94,0x13,0x50,0xF7,0x62,0xFC,
		0x99,0x0D,0xC1,0x12,0x6E,0x03,0x95,0xA3,0x57,0xC7,0x3C,0xB8,
		0x6B,0x40,0x56,0x65,0x70,0xFB,0x7A,0xE9,0x02,0xEC,0xD2,0xB6,
		0x54,0xD7,0x34,0xAD,0x3D,0x9E,0x11,0x61,0x53,0xBE,0xEA,0xB8,
		0x17,0x48,0xA8,0xDC,0x70,0xAE,0x65,0x99,0x3F,0x82,0x4C,0xFF,
		0x6A,0xC9,0xFA,0xB1,0xFA,0xE4,0x4F,0x5D,0xA4,0x05,0xC2,0x8E,
		0x55,0xC0,0xB1,0x1D,0xCC,0x17,0xF3,0xFA,0x65,0xD8,0x6B,0x09,
		0x13,0x01,0x2A,0x39,0xF1,0x86,0x73,0xE3,0x7A,0xC8,0xDB,0x7D,
		0xDA,0x1C,0xA1,0x2D,0xBA,0x2C,0x00,0x6B,0x2C,0x55,0x28,0x2B,
		0xD5,0xF5,0x3C,0x9F,0x50,0xA7,0xB7,0x28,0x9F,0x22,0xD5,0x3A,
		0xC4,0x53,0x01,0xC9,0xF3,0x69,0xB1,0x8D,0x01,0x36,0xF8,0xA8,
		0x89,0xCA,0x2E,0x72,0xBC,0x36,0x3A,0x42,0xC1,0x06,0xD6,0x0E,
		0xCB,0x4D,0x5C,0x1F,0xE4,0xA1,0x17,0xBF,0x55,0x64,0x1B,0xB4,
		0x52,0xEC,0x15,0xED,0x32,0xB1,0x81,0x07,0xC9,0x71,0x25,0xF9,
		0x4D,0x48,0x3D,0x18,0xF4,0x12,0x09,0x32,0xC4,0x0B,0x7A,0x4E,
		0x83,0xC3,0x10,0x90,0x51,0x2E,0xBE,0x87,0xF9,0xDE,0xB4,0xE6,
		0x3C,0x29,0xB5,0x32,0x01,0x9D,0x95,0x04,0xBD,0x42,0x89,0xFD,
		0x21,0xEB,0xE9,0x88,0x5A,0x27,0xBB,0x31,0xC4,0x26,0x99,0xAB,
		0x8C,0xA1,0x76,0xDB,
	}

	var dh2048_g = []byte{ 0x02,}

	client.A = new(big.Int).SetBytes(crypto.GenerateNonce(256))
	client.P = new(big.Int).SetBytes(dh2048_p)

	g_a := new(big.Int)
	g_a.Exp(new(big.Int).SetBytes(dh2048_g), client.A, client.P)
	// ServerNonce
	server_DHInnerData := &TLServer_DHInnerData{}

	server_DHInnerData.Nonce = client.Nonce
	server_DHInnerData.ServerNonce = client.ServerNonce
	server_DHInnerData.G = int32(dh2048_g[0])
	server_DHInnerData.GA = string(g_a.Bytes())
	server_DHInnerData.DhPrime = string(dh2048_p)
	server_DHInnerData.ServerTime = int32(time.Now().Unix())

	server_DHInnerData_buf := server_DHInnerData.Encode()
	// server_DHInnerData_buf_sha1 := sha1.Sum(server_DHInnerData_buf)

	// 创建aes和iv key
	tmp_aes_key_and_iv := make([]byte, 64)
	sha1_a := sha1.Sum(append(client.NewNonce, client.ServerNonce...))
	sha1_b := sha1.Sum(append(client.ServerNonce, client.NewNonce...))
	sha1_c := sha1.Sum(append(client.NewNonce, client.NewNonce...))
	copy(tmp_aes_key_and_iv, sha1_a[:])
	copy(tmp_aes_key_and_iv[20:], sha1_b[:])
	copy(tmp_aes_key_and_iv[40:], sha1_c[:])
	copy(tmp_aes_key_and_iv[60:], client.NewNonce[:4])

	tmpLen := 20+len(server_DHInnerData_buf)
	if tmpLen % 16 > 0 {
		tmpLen = (tmpLen / 16 + 1) * 16
	} else {
		tmpLen = 20 + len(server_DHInnerData_buf)
	}

	tmp_encrypted_answer := make([]byte, tmpLen)
	sha1_tmp := sha1.Sum(server_DHInnerData_buf)
	copy(tmp_encrypted_answer, sha1_tmp[:])
	copy(tmp_encrypted_answer[20:], server_DHInnerData_buf)

	e := crypto.NewAES256IGECryptor(tmp_aes_key_and_iv[:32], tmp_aes_key_and_iv[32:64])
	tmp_encrypted_answer, _ = e.Encrypt(tmp_encrypted_answer)

	server_DHParamsOk := &TLServer_DHParamsOk{}
	server_DHParamsOk.Nonce = client.Nonce
	server_DHParamsOk.ServerNonce = client.ServerNonce
	server_DHParamsOk.EncryptedAnswer = string(tmp_encrypted_answer)

	client.Codec.State = CODEC_server_DH_params_ok

	return server_DHParamsOk
}

func processSetClient_DHParams(client *client.Client, request *UnencryptedMessage) (TLObject) {
	setClient_DHParams, _ := request.Object.(*TLSetClient_DHParams)
	glog.Info("processSetClient_DHParams - request: ", setClient_DHParams.String())

	// 客户端传输数据解析
	// Nonce
	if !bytes.Equal(setClient_DHParams.Nonce, client.Nonce) {
		glog.Error("processSetClient_DHParams - Wrong Nonce")
		return nil
	}

	// ServerNonce
	if !bytes.Equal(setClient_DHParams.ServerNonce, client.ServerNonce) {
		glog.Error("processSetClient_DHParams - Wrong ServerNonce")
		return nil
	}

	bEncryptedData := []byte(setClient_DHParams.EncryptedData)

	// 创建aes和iv key
	tmp_aes_key_and_iv := make([]byte, 64)
	sha1_a := sha1.Sum(append(client.NewNonce, client.ServerNonce...))
	sha1_b := sha1.Sum(append(client.ServerNonce, client.NewNonce...))
	sha1_c := sha1.Sum(append(client.NewNonce, client.NewNonce...))
	copy(tmp_aes_key_and_iv, sha1_a[:])
	copy(tmp_aes_key_and_iv[20:], sha1_b[:])
	copy(tmp_aes_key_and_iv[40:], sha1_c[:])
	copy(tmp_aes_key_and_iv[60:], client.NewNonce[:4])

	d := crypto.NewAES256IGECryptor(tmp_aes_key_and_iv[:32], tmp_aes_key_and_iv[32:64])
	decryptedData, err := d.Decrypt(bEncryptedData)
	if err != nil {
		glog.Error("processSetClient_DHParams - AES256IGECryptor descrypt error")
		return nil
	}

	// TODO(@benqi): 检查签名是否合法
	dbuf := NewDecodeBuf(decryptedData[24:])
	client_DHInnerData := &TLClient_DHInnerData{}
	err = client_DHInnerData.Decode(dbuf)
	if err != nil {
		glog.Errorf("processSetClient_DHParams - TLClient_DHInnerData decode error: %s", err)
		return nil
	}

	glog.Info("processSetClient_DHParams - client_DHInnerData: ", client_DHInnerData.String())

	//
	if !bytes.Equal(client_DHInnerData.Nonce, client.Nonce) {
		glog.Error("processSetClient_DHParams - Wrong client_DHInnerData's Nonce")
		return nil
	}

	// ServerNonce
	if !bytes.Equal(client_DHInnerData.ServerNonce, client.ServerNonce) {
		glog.Error("processSetClient_DHParams - Wrong client_DHInnerData's ServerNonce")
		return nil
	}

	// hash_key
	authKeyNum := new(big.Int)
	authKeyNum.Exp(new(big.Int).SetBytes([]byte(client_DHInnerData.GB)), client.A, client.P)
	authKey := make([]byte, 256)
	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())


	authKeyAuxHash := make([]byte, len(client.NewNonce))
	copy(authKeyAuxHash, client.NewNonce)
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	sha1_d := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1_d[:]...)
	sha1_e := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1_e[:]...)

	dhGenOk := &TLDhGenOk{}
	dhGenOk.Nonce = client.Nonce
	dhGenOk.ServerNonce = client.ServerNonce
	dhGenOk.NewNonceHash1 = authKeyAuxHash[len(authKeyAuxHash)-16:len(authKeyAuxHash)]

	client.Codec.AuthKeyId = binary.LittleEndian.Uint64(authKeyAuxHash[len(client.NewNonce)+1+12:len(client.NewNonce)+1+12+8])
	client.Codec.AuthKey = authKey

	StoreAuthKey(client.Codec.AuthKeyId, client.Codec.AuthKey)

	client.Session.State = CODEC_dh_gen_ok

	return dhGenOk
}


func processMsgsAck(client *client.Client, request *TLMsgsAck) {
	glog.Info("processMsgsAck - request: %s", request.String())
}

func processPing(client *client.Client, request *EncryptedMessage2) (TLObject) {
	ping, _ := request.Object.(*TLPing)
	glog.Info("processPing - request data: ", ping.String())

	pong := &TLPong{
		PingId: ping.PingId,
	}

	return pong
}

func processPingDelayDisconnect(client *client.Client, request *EncryptedMessage2) (TLObject) {
	pingDelayDissconnect, _ := request.Object.(*TLPingDelayDisconnect)
	glog.Info("processPingDelayDisconnect - request data: ", pingDelayDissconnect.String())

	pong := &TLPong{
		PingId: pingDelayDissconnect.PingId,
	}

	return pong
}

func processDestroySession(client *client.Client, request *EncryptedMessage2) (TLObject) {
	destroySession, _ := request.Object.(*TLDestroySession)
	glog.Info("processDestroySession - request data: ", destroySession.String())

	// TODO(@benqi): 实现destroySession处理逻辑
	destroy_session_ok := &TLDestroySessionOk{
		SessionId: destroySession.SessionId,
	}
	return destroy_session_ok
}

func processGetFutureSalts(client *client.Client, request *EncryptedMessage2) (TLObject) {
	getFutureSalts, _ := request.Object.(*TLGetFutureSalts)
	glog.Info("processGetFutureSalts - request data: ", getFutureSalts.String())

	// TODO(@benqi): 实现getFutureSalts处理逻辑
	futureSalts := &TLFutureSalts{
	}

	return futureSalts
}

func processRpcDropAnswer(client *client.Client, request *EncryptedMessage2) (TLObject) {
	rpcDropAnswer, _ := request.Object.(*TLRpcDropAnswer)
	glog.Info("processRpcDropAnswer - request data: ", rpcDropAnswer.String())

	// TODO(@benqi): 实现rpcDropAnswer处理逻辑

	return nil
}

func processContestSaveDeveloperInfo(client *client.Client, request *EncryptedMessage2) (TLObject) {
	contestSaveDeveloperInfo, _ := request.Object.(*TLContestSaveDeveloperInfo)
	glog.Info("processGetFutureSalts - request data: ", contestSaveDeveloperInfo.String())

	// TODO(@benqi): 实现scontestSaveDeveloperInfo处理逻辑
	r := &TLTrue{}

	return r
}

func processInvokeWithLayer(client *client.Client, request *EncryptedMessage2) error {
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

func processInvokeAfterMsg(client *client.Client, request *EncryptedMessage2) error {
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

func processMsgContainer(client *client.Client, request *EncryptedMessage2) error {
	msgContainer, _ := request.Object.(*TLMsgContainer)
	glog.Info("processMsgContainer - request data: ", msgContainer.String())

	return nil
}

func processGzipPacked(client *client.Client, request *EncryptedMessage2) error {
	gzipPacked, _ := request.Object.(*TLGzipPacked)
	glog.Info("processGzipPacked - request data: ", gzipPacked.String())

	return nil
}
*/