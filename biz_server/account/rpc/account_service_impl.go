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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"google.golang.org/grpc/metadata"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/cosiner/gohper/errors"
)

const (
	TOKEN_TYPE_APNS = 1
	TOKEN_TYPE_GCM = 2
	TOKEN_TYPE_MPNS = 3
	TOKEN_TYPE_SIMPLE_PUSH = 4
	TOKEN_TYPE_UBUNTU_PHONE = 5
	TOKEN_TYPE_BLACKBERRY = 6
)

type AccountServiceImpl struct {
	UsersDAO *dao.UsersDAO
	DeviceDAO *dao.DevicesDAO
}

func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	glog.Info("AccountRegisterDevice - Process: %v", request)

	// 查出来
	md, _ := metadata.FromIncomingContext(ctx)
	rpcMetaData := mtproto.RpcMetaData{}
	rpcMetaData.Decode(md)

	// TODO(@benqi): check token_type

	do, err := s.DeviceDAO.SelectIdByAuthId(rpcMetaData.AuthId, int8(request.TokenType), request.Token)
	if err != nil {
		glog.Errorf("AccountRegisterDevice - s.DeviceDAO.SelectIdByAuthId error: %s", err)
		return nil, err
	}

	if do == nil {
		//
		do = &dataobject.DevicesDO{
			AuthId: rpcMetaData.AuthId,
			UserId: rpcMetaData.UserId,
			TokenType: int8(request.TokenType),
			Token: request.Token,
		}

		_, err := s.DeviceDAO.Insert(do)
		if err != nil {
			glog.Errorf("AccountRegisterDevice - s.DeviceDAO.Insert error: %s", err)
			return nil, err
		}
	} else {
		_, err := s.DeviceDAO.UpdateStateById(0, do.Id)
		if err != nil {
			glog.Errorf("AccountRegisterDevice - s.DeviceDAO.UpdateStateById error: %s", err)
			return nil, err
		}
	}

	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})
	glog.Infof("AccountRegisterDevice - reply: {%v}\n", reply)

	return reply, nil
}

func (s *AccountServiceImpl) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	glog.Info("AccountUnregisterDevice - Process: %v", request)

	// 查出来
	md, _ := metadata.FromIncomingContext(ctx)
	rpcMetaData := mtproto.RpcMetaData{}
	rpcMetaData.Decode(md)

	// TODO(@benqi): check token_type

	do, err := s.DeviceDAO.SelectIdByAuthId(rpcMetaData.AuthId, int8(request.TokenType), request.Token)
	if err != nil {
		glog.Errorf("AccountUnregisterDevice - s.DeviceDAO.SelectIdByAuthId error: %s", err)
		return nil, err
	}

	if do == nil {
		// glog.Errorf("AccountUnregisterDevice - s.DeviceDAO.Insert error: %s", err)
	} else {
		_, err := s.DeviceDAO.UpdateStateById(1, do.Id)
		if err != nil {
			glog.Errorf("AccountUnregisterDevice - s.DeviceDAO.UpdateStateById error: %s", err)
			return nil, err
		}
	}

	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})
	glog.Infof("AccountUnregisterDevice - reply: {%v}\n", reply)

	return reply, nil
}

func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	glog.Info("AccountUpdateNotifySettings - Process: %v", request)

	// TODO(@benqi): 实现逻辑
	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	glog.Infof("AccountUpdateNotifySettings - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AccountServiceImpl) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	glog.Info("AccountResetNotifySettings - Process: %v", request)

	// TODO(@benqi): 实现逻辑
	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	glog.Infof("AccountResetNotifySettings - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AccountServiceImpl) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	glog.Infof("AccountUpdateStatus - Process: {%v}", request)

	// TODO(@benqi): 实现逻辑
	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	glog.Infof("AccountUpdateStatus - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	glog.Info("AccountReportPeer - Process: %v", request)

	// TODO(@benqi): 实现逻辑
	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	glog.Infof("AccountReportPeer - reply: {%v}\n", reply)
	return reply, nil
}

func (s *AccountServiceImpl) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	glog.Infof("AccountCheckUsername - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	glog.Infof("AccountDeleteAccount - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	glog.Infof("AccountSetAccountTTL - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	glog.Infof("AccountUpdateDeviceLocked - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	glog.Infof("AccountResetAuthorization - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	glog.Infof("AccountUpdatePasswordSettings - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	glog.Infof("AccountConfirmPhone - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("AccountSendChangePhoneCode - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("AccountSendConfirmPhoneCode - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	glog.Infof("AccountGetNotifySettings - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	glog.Infof("AccountUpdateProfile - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	glog.Infof("AccountUpdateUsername - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	glog.Infof("AccountChangePhone - Process: %v", request)
	return nil, errors.New("Not impl")
}

// func (s *AccountServiceImpl)AccountGetWallPapers(ctx context.Context,  request *mtproto.TLAccountGetWallPapers) (*mtproto.Vector<WallPaper>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *AccountServiceImpl) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	glog.Infof("AccountGetPrivacy - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	glog.Infof("AccountSetPrivacy - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	glog.Infof("AccountGetAccountTTL - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	glog.Infof("AccountGetAuthorizations - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
// account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
func (s *AccountServiceImpl) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	glog.Infof("AccountGetPassword - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	glog.Infof("AccountGetPasswordSettings - Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *AccountServiceImpl) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	glog.Infof("AccountGetTmpPassword - Process: %v", request)
	return nil, errors.New("Not impl")
}
