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
	"github.com/nebulaim/telegramd/biz_model/model"
)

// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountUpdateProfile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	affected := dao.GetUsersDAO(dao.DB_SLAVE).UpdateProfile(
		request.GetFirstName(),
		request.GetLastName(),
		request.GetAbout(),
		md.UserId)

	ok := affected == 1

	if !ok {
		// TODO(@benqi): return rpc error!
		// panic()
	}

	user := model.GetUserModel().GetUser(md.UserId)
	// TODO(@benqi): Delivery updateUserName updates

	glog.Infof("AccountUpdateProfile - reply: {%v}\n", user)
	return user.To_User(), nil
}
