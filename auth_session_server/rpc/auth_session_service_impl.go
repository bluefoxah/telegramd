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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/model"
	"time"
)

type AuthSessionServiceImpl struct {
}

func NewAuthSessionServiceImpl() *AuthSessionServiceImpl {
	return &AuthSessionServiceImpl{}
}

func (s *AuthSessionServiceImpl) GetAuthUser(ctx context.Context, request *zproto.AuthUserReq) (*zproto.AuthUserRsp, error) {
	glog.Infof("GetAuthUser - request: {%v}", request)

	do := dao.GetAuthUsersDAO(dao.DB_SLAVE).SelectByAuthId(request.AuthKeyId)

	reply := &zproto.AuthUserRsp{
		AuthKeyId: request.AuthKeyId,
		UserId:    0,
	}

	// glog.Info("SelectByAuthId : ", do)
	if do == nil {
		glog.Errorf("GetAuthUser - Not find AuthKeyID: %v", request)
	} else {
		// TODO(@benqi): set online!!!
		reply.UserId = do.UserId
	}

	glog.Infof("GetAuthUser - reply: {%v}", reply)
	return reply, nil
}

func (s *AuthSessionServiceImpl) UpdateOnline(ctx context.Context, request *zproto.OnlineStatus) (*zproto.VoidRsp, error) {
	glog.Infof("UpdateOnline - request: {%v}", request)

	status := &model.SessionStatus{
		ServerId:        request.ServerId,
		UserId:          request.UserId,
		AuthKeyId:       request.AuthKeyId,
		SessionId:       request.SessionId,
		NetlibSessionId: request.NetlibSessionId,
		Now:             time.Now().Unix(),
	}
	model.GetOnlineStatusModel().SetOnline(status)

	glog.Infof("UpdateOnline - reply: {VoidRsp}")
	return &zproto.VoidRsp{}, nil
}
