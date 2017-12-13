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
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"time"
	"github.com/nebulaim/telegramd/biz_model/model"
)

// account.updateStatus#6628562c offline:Bool = Bool;
func (s *AccountServiceImpl) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountUpdateStatus - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	status := &model.SessionStatus{}
	status.UserId = md.UserId
	status.AuthKeyId = md.AuthId
	status.SessionId = md.SessionId
	status.ServerId = md.ServerId
	status.Now = time.Now().Unix()

	// Offline可能为nil，由grpc中间件保证Offline必须设置值
	// offline := request.GetOffline()
	if mtproto.FromBool(request.GetOffline()) {
		model.GetOnlineStatusModel().SetOnline(status)
	} else {
		model.GetOnlineStatusModel().SetOffline(status)
	}

	// TODO(@benqi): broadcast online status???

	glog.Infof("AccountUpdateStatus - reply: {true}")
	return mtproto.ToBool(true), nil
}
