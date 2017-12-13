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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (s *AccountServiceImpl) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountGetAccountTTL - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 估计不是这个规则
	do := dao.GetUserPrivacysDAO(dao.DB_SLAVE).SelectTTL(md.UserId)
	ttl := mtproto.NewTLAccountDaysTTL()
	if do == nil {
		ttl.SetDays(180)
	} else {
		ttl.SetDays(do.Ttl)
	}

	glog.Infof("AccountReportPeer - reply: %s\n", logger.JsonDebugData(ttl))
	return ttl.To_AccountDaysTTL(), nil
}
