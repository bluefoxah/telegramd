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

package model

import (
	"sync"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/nebulaim/telegramd/biz_model/base"
	base2 "github.com/nebulaim/telegramd/base/base"
	"github.com/nebulaim/telegramd/mtproto"
)

const (
	TOKEN_TYPE_APNS = 1
	TOKEN_TYPE_GCM = 2
	TOKEN_TYPE_MPNS = 3
	TOKEN_TYPE_SIMPLE_PUSH = 4
	TOKEN_TYPE_UBUNTU_PHONE = 5
	TOKEN_TYPE_BLACKBERRY = 6
	// Android里使用
	TOKEN_TYPE_INTERNAL_PUSH = 7
)

type accountModel struct {
	//
}

var (
	accountInstance *accountModel
	accountInstanceOnce sync.Once
)

func GetAccountModel() *accountModel {
	accountInstanceOnce.Do(func() {
		accountInstance = &accountModel{}
	})
	return accountInstance
}

func (m *accountModel) RegisterDevice(authKeyId int64, userId int32, tokenType int8, token string) {
	slave := dao.GetDevicesDAO(dao.DB_SLAVE)
	master := dao.GetDevicesDAO(dao.DB_MASTER)
	do := slave.SelectId(authKeyId, tokenType, token)
	if do == nil {
		do = &dataobject.DevicesDO{
			AuthId: authKeyId,
			UserId: userId,
			TokenType: tokenType,
			State: 0,
			CreatedAt: base2.NowFormatYMDHMS(),
		}
		master.Insert(do)
	} else {
		master.UpdateStateById(1, do.Id)
	}
}

func (m *accountModel) UnRegisterDevice(authKeyId int64, userId int32, tokenType int8, token string) bool {
	master := dao.GetDevicesDAO(dao.DB_MASTER)
	rows := master.UpdateState(tokenType, authKeyId, tokenType, token)
	return rows == 1

}

func (m *accountModel) GetNotifySettings(userId int32, peer *base.PeerUtil) *mtproto.PeerNotifySettings {
	do := dao.GetUserNotifySettingsDAO(dao.DB_SLAVE).SelectByPeer(userId, int8(peer.PeerType), peer.PeerId)

	if do == nil {
		settings := mtproto.NewTLPeerNotifySettingsEmpty()
		return settings.To_PeerNotifySettings()
	} else {
		settings := mtproto.NewTLPeerNotifySettings()
		settings.SetShowPreviews(do.ShowPreviews == 1)
		settings.SetSilent(do.Silent == 1)
		settings.SetMuteUntil(do.MuteUntil)
		settings.SetSound(do.Sound)
		return settings.To_PeerNotifySettings()
	}
}

func (m *accountModel) SetNotifySettings(userId int32, peer *base.PeerUtil, settings *mtproto.TLInputPeerNotifySettings) {
	slave := dao.GetUserNotifySettingsDAO(dao.DB_SLAVE)
	master := dao.GetUserNotifySettingsDAO(dao.DB_MASTER)

	var showPreviews int8 = 0
	var slient int8 = 0
	if settings.GetShowPreviews() {
		showPreviews = 1
	}
	if settings.GetSilent() {
		slient = 1
	}

	do := slave.SelectByPeer(userId, int8(peer.PeerType), peer.PeerId)
	if do == nil {
		do = &dataobject.UserNotifySettingsDO{}
		do.UserId = userId
		do.PeerType = int8(peer.PeerType)
		do.PeerId = peer.PeerId
		do.ShowPreviews = showPreviews
		do.Silent = slient
		do.MuteUntil = settings.GetMuteUntil()
		do.Sound = settings.GetSound()

		master.Insert(do)
	} else {
		master.UpdateByPeer(showPreviews, slient, settings.GetMuteUntil(), settings.GetSound(), 0, userId, int8(peer.PeerType), peer.PeerId)
	}
}

func (m *accountModel) ResetNotifySettings(userId int32) {
	slave := dao.GetUserNotifySettingsDAO(dao.DB_SLAVE)
	master := dao.GetUserNotifySettingsDAO(dao.DB_MASTER)

	master.DeleteNotAll(userId)
	do := slave.SelectByAll(userId)
	if do == nil {
		do = &dataobject.UserNotifySettingsDO{}
		do.UserId = userId
		do.PeerType = base.PEER_ALL
		do.PeerId = 0
		do.ShowPreviews = 1
		do.Silent = 0
		do.MuteUntil = 0
		master.Insert(do)
	} else {
		master.UpdateByPeer(1, 0, 0, "default", 0, userId, base.PEER_ALL, 0)
	}
}
