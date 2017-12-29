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
	"github.com/nebulaim/telegramd/grpc_util/service_discovery"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/zproto"
	"context"
)


type AuthSessionRPCClient struct {
	Client zproto.RPCAuthSessionClient
}

func NewAuthSessionRPCClient(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *AuthSessionRPCClient, err error) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	c = &AuthSessionRPCClient{}
	c.Client = zproto.NewRPCAuthSessionClient(conn)
	return
}

func (c *AuthSessionRPCClient) GetUserIDByAuthKey(keyID int64) int32 {
	var cacheKey CacheAuthKeyItem
	if k, ok := cacheAuthKey.Load(keyID); ok {
		// 本地缓存命中
		cacheKey = k.(CacheAuthKeyItem)
		if cacheKey.UserId != 0 {
			return cacheKey.UserId
		}
	}

	// 本地缓存未命中
	userRes, err := c.Client.GetAuthUser(context.Background(), &zproto.AuthUserReq{AuthKeyId: keyID})
	if err != nil {
		glog.Errorf("GetAuthUser error: %v", err)
		return 0
	} else {
		// 存入缓存
		cacheKey.UserId = userRes.UserId
		cacheAuthKey.Store(keyID, cacheKey)

		return cacheKey.UserId
	}
}

func (c *AuthSessionRPCClient) UpdateOnline(status *zproto.OnlineStatus) {
	_, err := c.Client.UpdateOnline(context.Background(), status)
	if err != nil {
		glog.Errorf("UpdateOnline error: {%v}", err)
	}
}
