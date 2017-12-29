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
	"github.com/nebulaim/telegramd/zproto"
	"github.com/golang/glog"
	"fmt"
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/auth_key_server/cache"
)

type AuthKeyStorageServiceImpl struct {
	cache cache.AuthKeyStorager
}

func NewAuthKeyStorageService(cache cache.AuthKeyStorager) *AuthKeyStorageServiceImpl {
	s := &AuthKeyStorageServiceImpl{
		cache: cache,
	}
	return s
}

// TODO(@benqi): cache
func (s *AuthKeyStorageServiceImpl) GetAuthKey(ctx context.Context, request *zproto.GetAuthKeyReq) (*zproto.AuthKeyData, error) {
	glog.Infof("GetAuthKey - request: %s", logger.JsonDebugData(request))

	authKey := s.cache.GetAuthKey(request.GetAuthKeyId())
	if authKey == nil {
		err := fmt.Errorf("process GetAuthKey - Wrong GetAuthKey")
		glog.Error(err)
		return nil, err
	}

	reply := &zproto.AuthKeyData{
		AuthKeyId: request.GetAuthKeyId(),
		AuthKey:   authKey,
	}
	glog.Infof("GetAuthKey - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

func (s *AuthKeyStorageServiceImpl) PutAuthKey(ctx context.Context, request *zproto.AuthKeyData) (*zproto.VoidRsp, error) {
	glog.Infof("PutAuthKey - request: %s", logger.JsonDebugData(request))

	s.cache.PutAuthKey(request.GetAuthKeyId(), request.GetAuthKey())

	glog.Infof("PutAuthKey - reply: {VoidRsp}")
	return &zproto.VoidRsp{}, nil
}

