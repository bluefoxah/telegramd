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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountSetAccountTTL - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	ttl := request.GetTtl().To_AccountDaysTTL()
	affected := dao.GetUserPrivacysDAO(dao.DB_MASTER).UpdateTTL(
		ttl.GetDays(),
		int32(time.Now().Unix()),
		md.UserId)

	updatedOk := affected == 1

	glog.Infof("AccountSetAccountTTL - reply: {%v}", updatedOk)
	return mtproto.ToBool(updatedOk), nil
}
