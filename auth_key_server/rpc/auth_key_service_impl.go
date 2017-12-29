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

package rpc

import (
	"context"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/base/logger"
	"fmt"
	"github.com/nebulaim/telegramd/base/crypto"
	"crypto/sha1"
	"bytes"
	"math/big"
	"time"
	"encoding/binary"
	"github.com/golang/protobuf/ptypes"
	"github.com/nebulaim/telegramd/zproto"
	"google.golang.org/grpc/metadata"
	"github.com/gogo/protobuf/proto"
	"encoding/base64"
	"google.golang.org/grpc"
	"encoding/hex"
	"github.com/nebulaim/telegramd/auth_key_server/cache"
)

const (
	SHA_DIGEST_LENGTH = 20
)

var (
	rsa = crypto.NewRSACryptor()
	headerRpcMetadata = "auth_key_metadata"


// TODO(@benqi): 预先计算出fingerprint
	// 这里直接使用了0xc3b42b026ce86b21
	fingerprint uint64 = 12240908862933197005

	// TODO(@benqi): 使用算法生成PQ
	// 这里直接指定了PQ值: {0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81}
	pq = string([]byte{0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81})

	// TODO(@benqi): 直接指定了p和q
	p = []byte{0x49, 0x4C, 0x55, 0x3B}
	q = []byte{0x53, 0x91, 0x10, 0x73}

	// TODO(@benqi): 直接指定了dh2048_p和dh2048_g!!!
	dh2048_p =[]byte{
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

	dh2048_g = []byte{ 0x02,}
)


type AuthKeyServiceImpl struct {
	cache cache.AuthKeyStorager
}

func NewAuthKeyService(cache cache.AuthKeyStorager) *AuthKeyServiceImpl {
	s := &AuthKeyServiceImpl{
		cache: cache,
	}
	return s
}

// req_pq#60469778 nonce:int128 = ResPQ;
func (s *AuthKeyServiceImpl) ReqPq(ctx context.Context, request *mtproto.TLReqPq) (*mtproto.ResPQ, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	authKeyMD := &zproto.AuthKeyMetadata{}
	if err := ptypes.UnmarshalAny(md.Extend, authKeyMD); err != nil {
		glog.Errorf("ReqPq - Unmarshal auth_key_metadata error")
		return nil, fmt.Errorf("ReqPq - Unmarshal auth_key_metadata error")
	}
	glog.Infof("ReqPq - metadata: %s, auth_key_metadata, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(authKeyMD), logger.JsonDebugData(request))

	// 检查数据是否合法
	if request.GetNonce() == nil || len(request.GetNonce()) != 16 {
		glog.Errorf("ReqPq - nonce not int128 type")
		return nil, fmt.Errorf("ReqPq - nonce not int128 type")
	}

	resPQ := mtproto.NewTLResPQ()
	resPQ.SetNonce(make([]byte, 16))
	copy(resPQ.Data2.Nonce, request.GetNonce())
	resPQ.SetServerNonce(crypto.GenerateNonce(16))
	resPQ.SetPq(pq)
	resPQ.SetServerPublicKeyFingerprints([]int64{int64(fingerprint)})

	// 缓存客户端Nonce
	authKeyMD.Nonce = request.GetNonce()
	authKeyMD.ServerNonce = resPQ.GetServerNonce()
	s.authKeyMetadataToTrailer(ctx, authKeyMD)

	glog.Infof("ReqPq - reply: %v", resPQ)
	return resPQ.To_ResPQ(), nil
}

// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
func (s *AuthKeyServiceImpl) Req_DHParams(ctx context.Context, request *mtproto.TLReq_DHParams) (*mtproto.Server_DH_Params, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	authKeyMD := &zproto.AuthKeyMetadata{}
	if err := ptypes.UnmarshalAny(md.Extend, authKeyMD); err != nil {
		glog.Errorf("Req_DHParams - Unmarshal auth_key_metadata error")
		return nil, fmt.Errorf("Req_DHParams - Unmarshal auth_key_metadata error")
	}
	glog.Infof("Req_DHParams - metadata: {%v}, auth_key_metadata: {%v}, request: {%v}", md, authKeyMD, request)

	// 客户端传输数据解析
	// Nonce
	if !bytes.Equal(request.Nonce, authKeyMD.Nonce) {
		glog.Error("process Req_DHParams - Invalid Nonce")
		return nil, fmt.Errorf("process Req_DHParams - InvalidNonce")
	}

	// ServerNonce
	if !bytes.Equal(request.ServerNonce, authKeyMD.ServerNonce) {
		glog.Error("process Req_DHParams - Wrong ServerNonce")
		return nil, fmt.Errorf("process Req_DHParams - Wrong ServerNonce")
	}

	// P
	if !bytes.Equal([]byte(request.P), p) {
		glog.Error("process Req_DHParams - Invalid p value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p value")
	}

	// Q
	if !bytes.Equal([]byte(request.Q), q) {
		glog.Error("process Req_DHParams - Invalid q value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid q value")
	}

	if request.PublicKeyFingerprint != int64(fingerprint) {
		glog.Error("process Req_DHParams - Invalid PublicKeyFingerprint value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid PublicKeyFingerprint value")
	}

	// encryptedData := []byte(request.EncryptedData)
	// glog.Info("EncryptedData: len = ", len(encryptedData), ", data: ", hex.EncodeToString(encryptedData))

	// 1. 解密
	encryptedPQInnerData := rsa.Decrypt([]byte(request.EncryptedData))

	// TODO(@benqi): sha1_check
	sha1Check := sha1.Sum(encryptedPQInnerData[20:])
	glog.Info(hex.EncodeToString(sha1Check[:]))
	glog.Info(hex.EncodeToString(encryptedPQInnerData[:20]))
	//if !bytes.Equal(sha1Check[:], encryptedPQInnerData[0:20]) {
	//	glog.Error("process Req_DHParams - sha1Check error")
	//	return nil, fmt.Errorf("process Req_DHParams - sha1Check error")
	//}

	// 2. 反序列化出pqInnerData
	pqInnerData := mtproto.NewTLPQInnerData()
	// &TLPQInnerData{}
	dbuf := mtproto.NewDecodeBuf(encryptedPQInnerData[SHA_DIGEST_LENGTH+4:])
	err := pqInnerData.Decode(dbuf)
	if err != nil {
		glog.Errorf("process Req_DHParams - TLPQInnerData decode error: %v", err)
		return nil, fmt.Errorf("process Req_DHParams - TLPQInnerData decode error: %v", err)
	}

	// glog.Info("processReq_DHParams - pqInnerData Decode sucess: ", pqInnerData.String())

	// 缓存NewNonce
	authKeyMD.NewNonce = pqInnerData.GetNewNonce()
	authKeyMD.A = crypto.GenerateNonce(256)
	authKeyMD.P = dh2048_p

	bigIntA := new(big.Int).SetBytes(authKeyMD.A)
	bigIntP := new(big.Int).SetBytes(authKeyMD.P)

	//c.A = new(big.Int).SetBytes()
	//c.P = new(big.Int).SetBytes(dh2048_p)

	g_a := new(big.Int)
	g_a.Exp(new(big.Int).SetBytes(dh2048_g), bigIntA, bigIntP)
	// ServerNonce
	server_DHInnerData := &mtproto.TLServer_DHInnerData{ Data2: &mtproto.Server_DHInnerData_Data{
		Nonce: authKeyMD.Nonce,
		ServerNonce: authKeyMD.ServerNonce,
		G: int32(dh2048_g[0]),
		GA: string(g_a.Bytes()),
		DhPrime: string(dh2048_p),
		ServerTime: int32(time.Now().Unix()),
	}}

	server_DHInnerData_buf := server_DHInnerData.Encode()
	// server_DHInnerData_buf_sha1 := sha1.Sum(server_DHInnerData_buf)

	// 创建aes和iv key
	tmp_aes_key_and_iv := make([]byte, 64)
	sha1_a := sha1.Sum(append(authKeyMD.NewNonce, authKeyMD.ServerNonce...))
	sha1_b := sha1.Sum(append(authKeyMD.ServerNonce, authKeyMD.NewNonce...))
	sha1_c := sha1.Sum(append(authKeyMD.NewNonce, authKeyMD.NewNonce...))
	copy(tmp_aes_key_and_iv, sha1_a[:])
	copy(tmp_aes_key_and_iv[20:], sha1_b[:])
	copy(tmp_aes_key_and_iv[40:], sha1_c[:])
	copy(tmp_aes_key_and_iv[60:], authKeyMD.NewNonce[:4])

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

	server_DHParamsOk := &mtproto.TLServer_DHParamsOk{ Data2: &mtproto.Server_DH_Params_Data{
		Nonce:  authKeyMD.Nonce,
		ServerNonce: authKeyMD.ServerNonce,
		EncryptedAnswer: string(tmp_encrypted_answer),
	}}

	glog.Infof("process Req_DHParams - reply: %s", logger.JsonDebugData(server_DHParamsOk))
	s.authKeyMetadataToTrailer(ctx, authKeyMD)
	return server_DHParamsOk.To_Server_DH_Params(), nil
}

// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
func (s *AuthKeyServiceImpl) SetClient_DHParams(ctx context.Context, request *mtproto.TLSetClient_DHParams) (*mtproto.SetClient_DHParamsAnswer, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	authKeyMD := &zproto.AuthKeyMetadata{}
	if err := ptypes.UnmarshalAny(md.Extend, authKeyMD); err != nil {
		glog.Errorf("SetClient_DHParams - Unmarshal auth_key_metadata error")
		return nil, fmt.Errorf("SetClient_DHParams - Unmarshal auth_key_metadata error")
	}
	glog.Infof("SetClient_DHParams - metadata: %s, auth_key_metadata, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(authKeyMD), logger.JsonDebugData(request))

	// TODO(@benqi): Impl SetClient_DHParams logic
	// 客户端传输数据解析
	// Nonce
	if !bytes.Equal(request.Nonce, authKeyMD.Nonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong Nonce")
		glog.Error(err)
		return nil, err
	}

	// ServerNonce
	if !bytes.Equal(request.ServerNonce, authKeyMD.ServerNonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong ServerNonce")
		glog.Error(err)
		return nil, err
	}

	bEncryptedData := []byte(request.EncryptedData)

	// 创建aes和iv key
	tmp_aes_key_and_iv := make([]byte, 64)
	sha1_a := sha1.Sum(append(authKeyMD.NewNonce, authKeyMD.ServerNonce...))
	sha1_b := sha1.Sum(append(authKeyMD.ServerNonce, authKeyMD.NewNonce...))
	sha1_c := sha1.Sum(append(authKeyMD.NewNonce, authKeyMD.NewNonce...))
	copy(tmp_aes_key_and_iv, sha1_a[:])
	copy(tmp_aes_key_and_iv[20:], sha1_b[:])
	copy(tmp_aes_key_and_iv[40:], sha1_c[:])
	copy(tmp_aes_key_and_iv[60:], authKeyMD.NewNonce[:4])

	d := crypto.NewAES256IGECryptor(tmp_aes_key_and_iv[:32], tmp_aes_key_and_iv[32:64])
	decryptedData, err := d.Decrypt(bEncryptedData)
	if err != nil {
		err := fmt.Errorf("process SetClient_DHParams - AES256IGECryptor descrypt error")
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): 检查签名是否合法
	dbuf := mtproto.NewDecodeBuf(decryptedData[24:])
	client_DHInnerData := mtproto.NewTLClient_DHInnerData()
	// &TLClient_DHInnerData{}
	err = client_DHInnerData.Decode(dbuf)
	if err != nil {
		glog.Errorf("processSetClient_DHParams - TLClient_DHInnerData decode error: %s", err)
		return nil, err
	}

	glog.Info("processSetClient_DHParams - client_DHInnerData: ", client_DHInnerData.String())

	//
	if !bytes.Equal(client_DHInnerData.GetNonce(), authKeyMD.Nonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong client_DHInnerData's Nonce")
		glog.Error(err)
		return nil, err
	}

	// ServerNonce
	if !bytes.Equal(client_DHInnerData.GetServerNonce(), authKeyMD.ServerNonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong client_DHInnerData's ServerNonce")
		glog.Error(err)
		return nil, err
	}

	bigIntA := new(big.Int).SetBytes(authKeyMD.A)
	bigIntP := new(big.Int).SetBytes(authKeyMD.P)

	// hash_key
	authKeyNum := new(big.Int)
	authKeyNum.Exp(new(big.Int).SetBytes([]byte(client_DHInnerData.GetGB())), bigIntA, bigIntP)
	authKey := make([]byte, 256)
	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())


	authKeyAuxHash := make([]byte, len(authKeyMD.NewNonce))
	copy(authKeyAuxHash, authKeyMD.NewNonce)
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	sha1_d := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1_d[:]...)
	sha1_e := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1_e[:]...)

	dhGenOk := &mtproto.TLDhGenOk{ Data2: &mtproto.SetClient_DHParamsAnswer_Data{
		Nonce: authKeyMD.Nonce,
		ServerNonce: authKeyMD.ServerNonce,
		NewNonceHash1: authKeyAuxHash[len(authKeyAuxHash)-16:len(authKeyAuxHash)],
	}}

	authKeyMD.AuthKeyId = int64(binary.LittleEndian.Uint64(authKeyAuxHash[len(authKeyMD.NewNonce)+1+12:len(authKeyMD.NewNonce)+1+12+8]))
	authKeyMD.AuthKey = authKey

	// TODO(@benqi): error 处理
	s.cache.PutAuthKey(authKeyMD.AuthKeyId, authKeyMD.AuthKey)

	glog.Infof("process Req_DHParams - reply: %s", logger.JsonDebugData(dhGenOk))
	s.authKeyMetadataToTrailer(ctx, authKeyMD)
	return dhGenOk.To_SetClient_DHParamsAnswer(), nil
}

// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
func (s *AuthKeyServiceImpl) DestroyAuthKey(ctx context.Context, request *mtproto.TLDestroyAuthKey) (*mtproto.DestroyAuthKeyRes, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("DestroyAuthKey - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl DestroyAuthKey logic

	return nil, fmt.Errorf("Not impl DestroyAuthKey")
}


func (s *AuthKeyServiceImpl) authKeyMetadataToTrailer(ctx context.Context, authKeyMD *zproto.AuthKeyMetadata) error {
	buf, err := proto.Marshal(authKeyMD)
	if err != nil {
		glog.Errorf("Marshal rpc_metadata error: %v", err)
		return err
	}

	md := metadata.Pairs(headerRpcMetadata, base64.StdEncoding.EncodeToString(buf))
	grpc.SetTrailer(ctx, md)
	return nil
}
