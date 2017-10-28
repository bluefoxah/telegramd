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
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"errors"
)

type AuthServiceImpl struct {
}

func (s *AuthServiceImpl) AuthLogOut(ctx context.Context, request *mtproto.TLAuthLogOut) (*mtproto.Bool, error) {
	glog.Infof("AuthLogOut - Process: {%v}", request)

	// TODO(@benqi): Logout逻辑
	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	glog.Infof("AuthLogOut - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AuthServiceImpl) AuthResetAuthorizations(ctx context.Context, request *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthSendInvites(ctx context.Context, request *mtproto.TLAuthSendInvites) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthBindTempAuthKey(ctx context.Context, request *mtproto.TLAuthBindTempAuthKey) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthDropTempAuthKeys(ctx context.Context, request *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthCheckPhone(ctx context.Context, request *mtproto.TLAuthCheckPhone) (*mtproto.Auth_CheckedPhone, error) {
	glog.Infof("AuthCheckPhone - Process: {%v}", request)

	// TODO(@benqi): 实现PhoneNumber检查逻辑
	reply := mtproto.MakeAuth_CheckedPhone(&mtproto.TLAuthCheckedPhone{
			PhoneRegistered: mtproto.MakeBool(&mtproto.TLBoolTrue{}),
		})

	glog.Infof("AuthCheckPhone - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AuthServiceImpl) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("AuthSendCode - Process: {%v}", request)

	authSentCode := &mtproto.TLAuthSentCode{}
	authSentCode.Type = mtproto.MakeAuth_SentCodeType(&mtproto.TLAuthSentCodeTypeApp{
		Length: 6,
	})
	authSentCode.PhoneCodeHash = "123456";

	reply := mtproto.MakeAuth_SentCode(authSentCode)
	glog.Infof("AuthSendCode - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AuthServiceImpl) AuthResendCode(ctx context.Context, request *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthSignUp(ctx context.Context, request *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	glog.Infof("AuthSignIn - Process: {%v}", request)

	// TODO(@benqi): 从数据库加载
	authAuthorization := &mtproto.TLAuthAuthorization{}
	user := &mtproto.TLUser{}
	user.Self = true
	user.Id = 1
	user.AccessHash = 1
	user.FirstName = "test1"
	user.LastName = "test1"
	user.Username = "test1"
	user.Phone = "+86 111 1111 1111"
	authAuthorization.User = mtproto.MakeUser(user)

	reply := mtproto.MakeAuth_Authorization(authAuthorization)
	glog.Infof("AuthSignIn - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AuthServiceImpl) AuthImportAuthorization(ctx context.Context, request *mtproto.TLAuthImportAuthorization) (*mtproto.Auth_Authorization, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthImportBotAuthorization(ctx context.Context, request *mtproto.TLAuthImportBotAuthorization) (*mtproto.Auth_Authorization, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthCheckPassword(ctx context.Context, request *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthRecoverPassword(ctx context.Context, request *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthExportAuthorization(ctx context.Context, request *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthRequestPasswordRecovery(ctx context.Context, request *mtproto.TLAuthRequestPasswordRecovery) (*mtproto.Auth_PasswordRecovery, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}
