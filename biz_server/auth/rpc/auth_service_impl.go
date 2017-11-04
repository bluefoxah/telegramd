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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"time"
	"github.com/nebulaim/telegramd/frontend/id"
	"fmt"
)

type AuthServiceImpl struct {
	UsersDAO *dao.UsersDAO
	AuthPhoneTransactionsDAO *dao.AuthPhoneTransactionsDAO
}

func (s *AuthServiceImpl) AuthLogOut(ctx context.Context, request *mtproto.TLAuthLogOut) (*mtproto.Bool, error) {
	glog.Infof("AuthLogOut - Process: {%v}", request)

	// TODO(@benqi): Logout逻辑处理，失效AuthKey
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

// 检查手机号码是否已经注册
func (s *AuthServiceImpl) AuthCheckPhone(ctx context.Context, request *mtproto.TLAuthCheckPhone) (*mtproto.Auth_CheckedPhone, error) {
	glog.Infof("AuthCheckPhone - Process: {%v}", request)

	// TODO(@benqi): panic/recovery
	usersDO, err := s.UsersDAO.SelectByPhoneNumber(request.PhoneNumber)
	if err != nil {
		glog.Errorf("AuthCheckPhone - s.UsersDAO.SelectUserIdByPhoneNumber: %s", err)
		return nil, err
	}

	var reply *mtproto.Auth_CheckedPhone

	if usersDO == nil {
		// 未注册
		reply = mtproto.MakeAuth_CheckedPhone(&mtproto.TLAuthCheckedPhone{
			PhoneRegistered: mtproto.MakeBool(&mtproto.TLBoolFalse{}),
		})
	} else {
		// 已经注册
		reply = mtproto.MakeAuth_CheckedPhone(&mtproto.TLAuthCheckedPhone{
			PhoneRegistered: mtproto.MakeBool(&mtproto.TLBoolTrue{}),
		})
	}

	glog.Infof("AuthCheckPhone - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AuthServiceImpl) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("AuthSendCode - Process: {%v}", request)

	// Check TLAuthSendCode
	// CurrentNumber: 是否为本机电话号码
	// 检查数据是否合法
	//switch request.CurrentNumber.(type) {
	//case *mtproto.Bool_BoolFalse:
	//	// 本机电话号码，AllowFlashcall为false
	//	if request.AllowFlashcall == false {
	//		// TODO(@benqi): 数据包非法
	//	}
	//}

	// TODO(@benqi): 独立出统一消息推送系统
	// 检查phpne是否存在，若存在是否在线决定是否通过短信发送或通过其他客户端发送
	// 透传AuthId，UserId，终端类型等
	// 检查满足条件的TransactionHash是否存在，可能的条件：
	//  1. is_deleted !=0 and now - created_at < 15 分钟
	do, err := s.AuthPhoneTransactionsDAO.SelectByPhoneAndApiIdAndHash(request.ApiId, request.ApiHash, request.PhoneNumber)
	if err != nil {
		glog.Errorf("AuthSendCode - s.AuthPhoneTransactionsDAO.SelectByPhoneAndApiIdAndHash: %s", err)
		return nil, err
	}

	if do == nil {
		do = &dataobject.AuthPhoneTransactionsDO{}
		do.ApiId = request.ApiId
		do.ApiHash = request.ApiHash
		do.PhoneNumber = request.PhoneNumber
		do.Code = "123456"
		do.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		// TODO(@benqi): 生成一个32字节的随机字串
		do.TransactionHash = fmt.Sprintf("%20d", id.NextId())

		_, err := s.AuthPhoneTransactionsDAO.Insert(do)
		if err != nil {
			glog.Errorf("AuthSendCode - s.AuthPhoneTransactionsDAO.Insert: %s", err)
			return nil, err
		}
	} else {
		// TODO(@benqi): 检查是否已经过了失效期
	}

	authSentCode := &mtproto.TLAuthSentCode{}
	authSentCode.Type = mtproto.MakeAuth_SentCodeType(&mtproto.TLAuthSentCodeTypeApp{
		Length: 6,
	})
	authSentCode.PhoneCodeHash = do.TransactionHash

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

	//// auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;
	//type TLAuthSignUp struct {
	//	PhoneNumber   string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	//	PhoneCodeHash string `protobuf:"bytes,2,opt,name=phone_code_hash,json=phoneCodeHash" json:"phone_code_hash,omitempty"`
	//	PhoneCode     string `protobuf:"bytes,3,opt,name=phone_code,json=phoneCode" json:"phone_code,omitempty"`
	//	FirstName     string `protobuf:"bytes,4,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	//	LastName      string `protobuf:"bytes,5,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	//}

	return nil, errors.New("Not impl")
}

func (s *AuthServiceImpl) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	glog.Infof("AuthSignIn - Process: {%v}", request)

	// Check code
	do1, err := s.AuthPhoneTransactionsDAO.SelectByPhoneCode(request.PhoneCodeHash, request.PhoneCode, request.PhoneNumber)
	if do1 == nil {
		glog.Errorf("AuthSignIn - s.AuthPhoneTransactionsDAO.SelectByPhoneCode: %s", err)
		return nil, err
	}

	do2, err := s.UsersDAO.SelectByPhoneNumber(request.PhoneNumber)
	if do2 == nil {
		glog.Errorf("AuthSignIn - s.UsersDAO.SelectByPhoneNumber: %s", err)
		return nil, err
	}

	// TODO(@benqi): 从数据库加载
	authAuthorization := &mtproto.TLAuthAuthorization{}
	user := &mtproto.TLUser{}
	user.Self = true
	user.Id = do2.Id
	user.AccessHash = do2.AccessHash
	user.FirstName = do2.FirstName
	user.LastName = do2.LastName
	user.Username = do2.Username
	user.Phone = request.PhoneNumber
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
