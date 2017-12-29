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
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/grpc_util/service_discovery"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/zproto"
	"context"
)


type AuthKeyRPCClient struct {
	Client mtproto.RPCAuthKeyClient
	CacheClient zproto.RPCAuthKeyStorageClient
}

func NewAuthKeyRPCClient(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *AuthKeyRPCClient, err error) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	c = &AuthKeyRPCClient{}
	c.Client = mtproto.NewRPCAuthKeyClient(conn)
	c.CacheClient = zproto.NewRPCAuthKeyStorageClient(conn)
	return
}

func (c *AuthKeyRPCClient) GetAuthKey(keyID int64) (authKey []byte) {
	var cacheKey CacheAuthKeyItem
	if k, ok := cacheAuthKey.Load(keyID); ok {
		// 本地缓存命中
		cacheKey = k.(CacheAuthKeyItem)
		if cacheKey.AuthKey != nil {
			authKey = cacheKey.AuthKey
			return
		}
	}

	// 本地缓存未命中
	authKeyData, err := c.CacheClient.GetAuthKey(context.Background(), &zproto.GetAuthKeyReq{AuthKeyId: keyID})
	if err != nil {
		glog.Errorf("GetAuthKey error: %v", err)
		// return nil
	} else {
		// 存入缓存
		cacheKey.AuthKey = authKeyData.AuthKey
		cacheAuthKey.Store(keyID, cacheKey)

		// 返回key
		authKey = authKeyData.AuthKey
	}

	return
}
