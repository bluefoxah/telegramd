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
	base2 "github.com/nebulaim/telegramd/base/base"
	base "github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_server/delivery"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

type AccountServiceImpl struct {
}

// account.registerDevice#637ea878 token_type:int token:string = Bool;
func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	glog.Info("AccountRegisterDevice - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): check toke_type invalid
	model.GetAccountModel().RegisterDevice(md.AuthId, md.UserId, int8(request.TokenType), request.Token)

	glog.Infof("AccountRegisterDevice - reply: {true}")
	return mtproto.ToBool(true), nil
}

// account.unregisterDevice#65c55b40 token_type:int token:string = Bool;
func (s *AccountServiceImpl) AccountUnregisterDevice(ctx context.Context, request *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	glog.Info("AccountUnregisterDevice - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): check toke_type invalid
	ok := model.GetAccountModel().UnRegisterDevice(md.AuthId, md.UserId, int8(request.TokenType), request.Token)

	glog.Infof("AccountUnregisterDevice - reply: {%v}\n", ok)
	return mtproto.ToBool(ok), nil
}

// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (s *AccountServiceImpl) AccountUpdateNotifySettings(ctx context.Context, request *mtproto.TLAccountUpdateNotifySettings) (reply *mtproto.Bool, err error) {
	glog.Infof("AccountUpdateNotifySettings - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputNotifyPeer(request.GetPeer())
	settings := request.GetSettings().GetInputPeerNotifySettings()

	model.GetAccountModel().SetNotifySettings(md.UserId, peer, settings)

	update := &mtproto.TLUpdateNotifySettings{}
	update.Peer = peer.ToNotifyPeer()
	updateSettings := &mtproto.TLPeerNotifySettings{}
	updateSettings.ShowPreviews = settings.ShowPreviews
	updateSettings.Silent = settings.Silent
	updateSettings.MuteUntil = settings.MuteUntil
	updateSettings.Sound = settings.Sound
	update.NotifySettings = updateSettings.ToPeerNotifySettings()

	updates := &mtproto.TLUpdateShort{}
	updates.Date = int32(time.Now().Unix())
	updates.Update = update.ToUpdate()

	delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
		md.AuthId,
		md.SessionId,
		md.NetlibSessionId,
		[]int32{md.UserId},
		updates.ToUpdates().Encode())
	//glog.Infof("AccountUpdateNotifySettings - delivery: %v", delivery)
	//_, _ = s.SyncRPCClient.Client.DeliveryUpdates(context.Background(), delivery)

	reply = mtproto.MakeBool(&mtproto.TLBoolTrue{})
	glog.Infof("AccountUpdateNotifySettings - reply: {%v}\n", reply)
	return
}

// account.resetNotifySettings#db7e1747 = Bool;
func (s *AccountServiceImpl) AccountResetNotifySettings(ctx context.Context, request *mtproto.TLAccountResetNotifySettings) (*mtproto.Bool, error) {
	glog.Info("AccountResetNotifySettings - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	model.GetAccountModel().ResetNotifySettings(md.UserId)
	peer := &base.PeerUtil{}
	peer.PeerType = base.PEER_ALL
	update := &mtproto.TLUpdateNotifySettings{}
	update.Peer = peer.ToNotifyPeer()
	updateSettings := &mtproto.TLPeerNotifySettings{}
	updateSettings.ShowPreviews = true
	updateSettings.Silent = false
	updateSettings.MuteUntil = 0
	updateSettings.Sound = "default"
	update.NotifySettings = updateSettings.ToPeerNotifySettings()

	updates := &mtproto.TLUpdateShort{}
	updates.Date = int32(time.Now().Unix())
	updates.Update = update.ToUpdate()

	delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
		md.AuthId,
		md.SessionId,
		md.NetlibSessionId,
		[]int32{md.UserId},
		updates.ToUpdates().Encode())

	reply := &mtproto.TLBoolTrue{}
	glog.Infof("AccountResetNotifySettings - reply: {%v}\n", reply)
	return reply.ToBool(), nil
}

// account.updateStatus#6628562c offline:Bool = Bool;
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

	// TODO(@benqi): broadcast online status???

	reply := &mtproto.TLBoolTrue{}
	glog.Infof("AccountUpdateStatus - reply: {%v}\n", reply)
	return reply.ToBool(), nil
}

//account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	glog.Infof("AccountReportPeer - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	peer := base.FromInputPeer(request.GetPeer())
	reason := base.FromReportReason(request.GetReason())

	// Insert to db
	do := &dataobject.ReportsDO{}
	do.AuthId = md.AuthId
	do.UserId = md.UserId
	do.PeerType = peer.PeerType
	do.PeerId = peer.PeerId
	do.Reason = int8(reason)
	if reason == base.REASON_OTHER {
		do.Content = request.Reason.GetInputReportReasonOther().GetText()
	}

	dao.GetReportsDAO(dao.DB_MASTER).Insert(do)

	reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})
	glog.Infof("AccountReportPeer - reply: {%v}\n", reply)
	return reply, nil
}

//account.checkUsername#2714d86c username:string = Bool;
func (s *AccountServiceImpl) AccountCheckUsername(ctx context.Context, request *mtproto.TLAccountCheckUsername) (*mtproto.Bool, error) {
	glog.Infof("AccountCheckUsername - Process: %v", request)

	params := make(map[string]interface{})
	params["username"] = request.GetUsername()

	r := dao.GetCommonDAO(dao.DB_SLAVE).CheckExists("users", params)
	glog.Infof("AccountReportPeer - reply: {%v}\n", r)
	return mtproto.ToBool(r), nil
}

// account.deleteAccount#418d4e0b reason:string = Bool;
func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	glog.Infof("AccountDeleteAccount - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	affected := dao.GetUsersDAO(dao.DB_MASTER).Delete(
		request.GetReason(),
		base2.NowFormatYMDHMS(),
		md.UserId)

	deletedOk := affected == 1
	// TODO(@benqi): 1. Clear account data 2. Kickoff other client

	glog.Infof("AccountDeleteAccount - reply: {%v}\n", deletedOk)
	return mtproto.ToBool(deletedOk), nil
}

// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (s *AccountServiceImpl) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	glog.Infof("AccountSetAccountTTL - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	affected := dao.GetUserPrivacysDAO(dao.DB_MASTER).UpdateTTL(
		request.GetTtl().GetAccountDaysTTL().GetDays(),
		int32(time.Now().Unix()),
		md.UserId)

	updatedOk := affected == 1

	glog.Infof("AccountSetAccountTTL - reply: {%v}\n", updatedOk)
	return mtproto.ToBool(updatedOk), nil
}

// account.updateDeviceLocked#38df3532 period:int = Bool;
// Android's client and tdesktop client not impl this functions
func (s *AccountServiceImpl) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	glog.Infof("AccountUpdateDeviceLocked - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.resetAuthorization#df77f3bc hash:long = Bool;
func (s *AccountServiceImpl) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	glog.Infof("AccountResetAuthorization - Process: %v", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)
	// TODO(@benqi): Terminal session



	return nil, errors.New("Not impl")
}

//account.updatePasswordSettings#fa7c4b86 current_password_hash:bytes new_settings:account.PasswordInputSettings = Bool;
func (s *AccountServiceImpl) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	glog.Infof("AccountUpdatePasswordSettings - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (s *AccountServiceImpl) AccountConfirmPhone(ctx context.Context, request *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	glog.Infof("AccountConfirmPhone - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.sendChangePhoneCode#8e57deb flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool = auth.SentCode;
func (s *AccountServiceImpl) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("AccountSendChangePhoneCode - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.sendConfirmPhoneCode#1516d7bd flags:# allow_flashcall:flags.0?true hash:string current_number:flags.0?Bool = auth.SentCode;
func (s *AccountServiceImpl) AccountSendConfirmPhoneCode(ctx context.Context, request *mtproto.TLAccountSendConfirmPhoneCode) (*mtproto.Auth_SentCode, error) {
	glog.Infof("AccountSendConfirmPhoneCode - Process: %v", request)
	return nil, errors.New("Not impl")
}

//account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	glog.Infof("AccountGetNotifySettings - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputNotifyPeer(request.GetPeer())

	reply := model.GetAccountModel().GetNotifySettings(md.UserId, peer)
	glog.Infof("AccountReportPeer - reply: {%v}\n", reply)
	return reply, nil
}

// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	glog.Infof("AccountUpdateProfile - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

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
	return user.ToUser(), nil
}

// account.updateUsername#3e0bdd7c username:string = User;
func (s *AccountServiceImpl) AccountUpdateUsername(ctx context.Context, request *mtproto.TLAccountUpdateUsername) (*mtproto.User, error) {
	glog.Infof("AccountUpdateUsername - Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	affected := dao.GetUsersDAO(dao.DB_MASTER).UpdateUsername(request.GetUsername(), md.UserId)
	ok := affected == 1

	if !ok {
		// panic()
	}

	user := model.GetUserModel().GetUser(md.UserId)
	// TODO(@benqi): Delivery updateUserName updates

	glog.Infof("AccountReportPeer - reply: {%v}\n", user)
	return user.ToUser(), nil
}

// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (s *AccountServiceImpl) AccountChangePhone(ctx context.Context, request *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	glog.Infof("AccountChangePhone - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.getWallPapers#c04cfac2 = Vector<WallPaper>;
// func (s *AccountServiceImpl)AccountGetWallPapers(ctx context.Context,  request *mtproto.TLAccountGetWallPapers) (*mtproto.Vector<WallPaper>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (s *AccountServiceImpl) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	glog.Infof("AccountGetPrivacy - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (s *AccountServiceImpl) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	glog.Infof("AccountSetPrivacy - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (s *AccountServiceImpl) AccountGetAccountTTL(ctx context.Context, request *mtproto.TLAccountGetAccountTTL) (*mtproto.AccountDaysTTL, error) {
	glog.Infof("AccountGetAccountTTL - Process: {%v}", request)

	// TODO(@benqi): 估计不是这个规则
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	do := dao.GetUserPrivacysDAO(dao.DB_SLAVE).SelectTTL(md.UserId)
	ttl := &mtproto.TLAccountDaysTTL{}
	if do == nil {
		ttl.Days = 180
	} else {
		ttl.Days = do.Ttl
	}

	glog.Infof("AccountReportPeer - reply: {%v}\n", ttl)
	return ttl.ToAccountDaysTTL(), nil
}

// account.getAuthorizations#e320c158 = account.Authorizations;
func (s *AccountServiceImpl) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	glog.Infof("AccountGetAuthorizations - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.getPassword#548a30f5 = account.Password;
// account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
// account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
func (s *AccountServiceImpl) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	glog.Infof("AccountGetPassword - Process: {%v}", request)

	// md := grpc_util.RpcMetadataFromIncoming(ctx)
	password := &mtproto.TLAccountNoPassword{}
	password.NewSalt = []byte("111")
	password.EmailUnconfirmedPattern = "EmailUnconfirmedPattern"

	return password.ToAccount_Password(), nil
}

// account.getPasswordSettings#bc8d11bb current_password_hash:bytes = account.PasswordSettings;
func (s *AccountServiceImpl) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	glog.Infof("AccountGetPasswordSettings - Process: %v", request)
	return nil, errors.New("Not impl")
}

// account.getTmpPassword#4a82327e password_hash:bytes period:int = account.TmpPassword;
func (s *AccountServiceImpl) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	glog.Infof("AccountGetTmpPassword - Process: %v", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): Check password_hash invalid, android source code
	// byte[] hash = new byte[currentPassword.current_salt.length * 2 + passwordBytes.length];
	// System.arraycopy(currentPassword.current_salt, 0, hash, 0, currentPassword.current_salt.length);
	// System.arraycopy(passwordBytes, 0, hash, currentPassword.current_salt.length, passwordBytes.length);
	// System.arraycopy(currentPassword.current_salt, 0, hash, hash.length - currentPassword.current_salt.length, currentPassword.current_salt.length);

	// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
	tmpPassword := &mtproto.TLAccountTmpPassword{}
	tmpPassword.TmpPassword = []byte("01234567899876543210")
	tmpPassword.ValidUntil = int32(time.Now().Unix()) + request.Period

	return tmpPassword.ToAccount_TmpPassword(), nil
}
