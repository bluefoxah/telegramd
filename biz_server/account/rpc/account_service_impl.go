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
	"github.com/cosiner/gohper/errors"
	"github.com/nebulaim/telegramd/biz_model/model"
	"time"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/biz_model/base"
)

type AccountServiceImpl struct {
}

func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	glog.Info("AccountRegisterDevice - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): check toke_type invalid
	model.GetAccountModel().RegisterDevice(md.AuthId, md.UserId, int8(request.TokenType), request.Token)

	reply := &mtproto.TLBoolTrue{}
	glog.Infof("AccountRegisterDevice - reply: {%v}\n", reply)

	return reply.ToBool(), nil
}

func (s *AccountServiceImpl) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (reply *mtproto.Bool, err error) {
	glog.Info("AccountUnregisterDevice - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): check toke_type invalid
	if ok := model.GetAccountModel().UnRegisterDevice(md.AuthId, md.UserId, int8(request.TokenType), request.Token); ok {
		reply = mtproto.MakeBool(&mtproto.TLBoolTrue{})
	} else {
		reply = mtproto.MakeBool(&mtproto.TLBoolFalse{})
	}

	err = nil
	glog.Infof("AccountUnregisterDevice - reply: {%v}\n", reply)

	return
}

//inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
//inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;
//inputPeerNotifySettings#38935eb2
// flags:#
// 	show_previews:flags.0?true
// 	silent:flags.1?true
//  mute_until:int
// 	sound:string = InputPeerNotifySettings;
func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (reply *mtproto.Bool, err error) {
	glog.Info("AccountUpdateNotifySettings - Process: %v", request)

	peer := base.FromInputNotifyPeer(request.GetPeer())
	settings := request.GetSettings().GetInputPeerNotifySettings()
	_ = settings
	switch peer.PeerType {
	case base.PEER_EMPTY:
	case base.PEER_SELF:
	case base.PEER_USER:
	case base.PEER_CHAT:
	case base.PEER_CHANNEL:
	case base.PEER_USERS:
	case base.PEER_CHATS:
	case base.PEER_ALL:
		reply = mtproto.MakeBool(&mtproto.TLBoolFalse{})
	default:
	}
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

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	status := &model.SessionStatus{}
	status.UserId = md.UserId
	status.AuthKeyId = md.AuthId
	status.SessionId = md.SessionId
	status.ServerId = md.ServerId
	status.Now = time.Now().Unix()

	// Offline可能为nil，由grpc中间件保证Offline必须设置值
	if request.Offline.GetBoolTrue() != nil {
		model.GetOnlineStatusModel().SetOnline(status)
	} else {
		model.GetOnlineStatusModel().SetOffline(status)
	}

	reply := &mtproto.TLBoolTrue{}

	glog.Infof("AccountUpdateStatus - reply: {%v}\n", reply)
	return reply.ToBool(), nil
}

func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	glog.Info("AccountReportPeer - Process: %v", request)

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
