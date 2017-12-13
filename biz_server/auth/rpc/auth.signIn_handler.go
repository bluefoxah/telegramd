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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/ttacon/libphonenumber"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
func (s *AuthServiceImpl) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthSignIn - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumber := libphonenumber.NormalizeDigitsOnly(request.PhoneNumber)

	// Check code
	authPhoneTransactionsDAO := dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE)
	do1 := authPhoneTransactionsDAO.SelectByPhoneCode(request.PhoneCodeHash, request.PhoneCode, phoneNumber)
	if do1 == nil {
	    err := fmt.Errorf("SelectByPhoneCode(_) return empty in request: {}%v", request)
	    glog.Error(err)
	    return nil, err
	}

	usersDAO := dao.GetUsersDAO(dao.DB_SLAVE)
	do2 := usersDAO.SelectByPhoneNumber(phoneNumber)
	if do2 == nil {
	    err := fmt.Errorf("SelectByPhoneNumber(_) return empty in request{}%v", request)
	    glog.Error(err)
	    return nil, err
	}

	do3 := dao.GetAuthUsersDAO(dao.DB_SLAVE).SelectByAuthId(md.AuthId)
	if do3 == nil {
	    do3 := &dataobject.AuthUsersDO{}
	    do3.AuthId = md.AuthId
	    do3.UserId = do2.Id
	    dao.GetAuthUsersDAO(dao.DB_MASTER).Insert(do3)
	}

	// TODO(@benqi): 从数据库加载
	authAuthorization := mtproto.NewTLAuthAuthorization()
	user := mtproto.NewTLUser()
	user.SetSelf(true)
	user.SetId(do2.Id)
	user.SetAccessHash(do2.AccessHash)
	user.SetFirstName(do2.FirstName)
	user.SetLastName(do2.LastName)
	user.SetUsername(do2.Username)
	user.SetPhone(phoneNumber)
	authAuthorization.SetUser(user.To_User())

	glog.Infof("AuthSignIn - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
}
