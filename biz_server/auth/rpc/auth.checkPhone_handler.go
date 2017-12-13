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
	"github.com/ttacon/libphonenumber"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

// auth.checkPhone#6fe51dfb phone_number:string = auth.CheckedPhone;
func (s *AuthServiceImpl) AuthCheckPhone(ctx context.Context, request *mtproto.TLAuthCheckPhone) (*mtproto.Auth_CheckedPhone, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthCheckPhone - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): panic/recovery
	usersDAO := dao.GetUsersDAO(dao.DB_SLAVE)

	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumer := libphonenumber.NormalizeDigitsOnly(request.PhoneNumber)

	usersDO := usersDAO.SelectByPhoneNumber(phoneNumer)

	var reply *mtproto.Auth_CheckedPhone
	if usersDO == nil {
	    // 未注册
	    checkedPhone := mtproto.NewTLAuthCheckedPhone()
	    checkedPhone.SetPhoneRegistered(mtproto.ToBool(false))
	    reply = checkedPhone.To_Auth_CheckedPhone()
	} else {
	    // 已经注册
		checkedPhone := mtproto.NewTLAuthCheckedPhone()
		checkedPhone.SetPhoneRegistered(mtproto.ToBool(true))
		reply = checkedPhone.To_Auth_CheckedPhone()
	}

	glog.Infof("AuthCheckPhone - reply: %s\n", reply)
	return reply, nil
}
