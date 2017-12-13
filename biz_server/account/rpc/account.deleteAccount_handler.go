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
	base2 "github.com/nebulaim/telegramd/base/base"
)

// account.deleteAccount#418d4e0b reason:string = Bool;
func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountDeleteAccount - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl AccountDeleteAccount logic
	affected := dao.GetUsersDAO(dao.DB_MASTER).Delete(
		request.GetReason(),
		base2.NowFormatYMDHMS(),
		md.UserId)

	deletedOk := affected == 1
	// TODO(@benqi): 1. Clear account data 2. Kickoff other client

	glog.Infof("AccountDeleteAccount - reply: {%v}", deletedOk)
	return mtproto.ToBool(deletedOk), nil
}
