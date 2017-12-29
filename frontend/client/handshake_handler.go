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
	"github.com/nebulaim/telegramd/frontend/rpc"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/zproto"
	"google.golang.org/grpc/metadata"
	"github.com/nebulaim/telegramd/grpc_util"
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"encoding/base64"
	"github.com/gogo/protobuf/proto"
	"github.com/nebulaim/telegramd/base/logger"
)

var (
	headerRpcMetadata = "auth_key_metadata"
)

type HandshakeHandler struct {
	AuthKeyClient   *rpc.AuthKeyRPCClient
	RpcMetaData     *zproto.RpcMetadata
	AuthKeyMetadata *zproto.AuthKeyMetadata
}

func NewHandshakeHandler(authKeyClient *rpc.AuthKeyRPCClient) *HandshakeHandler {
	return &HandshakeHandler{
		AuthKeyClient: authKeyClient,
		AuthKeyMetadata: &zproto.AuthKeyMetadata{},
		RpcMetaData: &zproto.RpcMetadata{},
	}
}

// Server To Client
func (c* HandshakeHandler) authKeyFromMD(md metadata.MD) *zproto.AuthKeyMetadata {
	val := metautils.NiceMD(md).Get(headerRpcMetadata)
	if val == "" {
		glog.Errorf("Unknown error!")
		return nil
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		glog.Errorf("Base64 decode error, rpc_error: %s, error: %v", val, err)
		return nil
	}

	err = proto.Unmarshal(buf, c.AuthKeyMetadata)
	if err != nil {
		glog.Errorf("RpcError unmarshal error, rpc_error: %s, error: %v", val, err)
		return nil
	}

	return c.AuthKeyMetadata
}

func (c* HandshakeHandler) onHandshakeMsgsAck(client *Client, request *mtproto.TLMsgsAck) {
	glog.Info("onHandshakeMsgsAck - request: %s", logger.JsonDebugData(request))
}

func (c* HandshakeHandler) onReqPq(client *Client, request *mtproto.UnencryptedMessage) (mtproto.TLObject) {
	reqPq, _ := request.Object.(*mtproto.TLReqPq)
	glog.Infof("onReqPq - request data: %v", reqPq)

	var err error
	var header, trailer metadata.MD

	c.RpcMetaData.Extend, err = ptypes.MarshalAny(c.AuthKeyMetadata)
	ctx, _ := grpc_util.RpcMetadatToOutgoing(context.Background(), c.RpcMetaData, )
	res, err := c.AuthKeyClient.Client.ReqPq(ctx, reqPq, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		glog.Error("onReqPq - error: %v", err)
		return nil
	}

	authKeyMD := c.authKeyFromMD(trailer)

	client.Codec.State = mtproto.CODEC_resPQ
	glog.Infof("onReqPq - reply: %v, c: %v", res, authKeyMD)
	return res
}

func (c* HandshakeHandler) onReq_DHParams(client *Client, request *mtproto.UnencryptedMessage) (mtproto.TLObject) {
	req_DH_params, _ := request.Object.(*mtproto.TLReq_DHParams)
	glog.Info("processReq_DHParams - request: ", req_DH_params.String())

	var err error
	var header, trailer metadata.MD

	c.RpcMetaData.Extend, err = ptypes.MarshalAny(c.AuthKeyMetadata)
	ctx, _ := grpc_util.RpcMetadatToOutgoing(context.Background(), c.RpcMetaData)
	res, err := c.AuthKeyClient.Client.Req_DHParams(ctx, req_DH_params, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		glog.Error("onReq_DHParams - error: %v", err)
		return nil
	}

	authKeyMD := c.authKeyFromMD(trailer)

	client.Codec.State = mtproto.CODEC_server_DH_params_ok
	glog.Infof("onReq_DHParams - reply: %v, c: %v", res, authKeyMD)
	return res
}

func (c* HandshakeHandler) onSetClient_DHParams(client *Client, request *mtproto.UnencryptedMessage) (mtproto.TLObject) {
	setClient_DHParams, _ := request.Object.(*mtproto.TLSetClient_DHParams)
	glog.Info("onSetClient_DHParams - request: ", setClient_DHParams.String())

	var err error
	var header, trailer metadata.MD

	c.RpcMetaData.Extend, err = ptypes.MarshalAny(c.AuthKeyMetadata)
	ctx, _ := grpc_util.RpcMetadatToOutgoing(context.Background(), c.RpcMetaData)
	res, err := c.AuthKeyClient.Client.SetClient_DHParams(ctx, setClient_DHParams, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		glog.Error("onSetClient_DHParams - error: %v", err)
		return nil
	}

	authKeyMD := c.authKeyFromMD(trailer)

	// TODO(@benqi): remove...
	client.Codec.AuthKeyId = authKeyMD.AuthKeyId
	client.Codec.AuthKey = authKeyMD.AuthKey
	client.Codec.State = mtproto.CODEC_dh_gen_ok

	glog.Infof("onSetClient_DHParams - reply: %v, c: %v", res, authKeyMD)
	return res
}
