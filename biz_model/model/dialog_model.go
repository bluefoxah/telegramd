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
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/golang/glog"
	base2 "github.com/nebulaim/telegramd/base/base"
	"time"
)

var (
	dialogInstance *dialogModel
	dialogInstanceOnce sync.Once
)

type dialogModel struct {
}

func GetDialogModel() *dialogModel {
	dialogInstanceOnce.Do(func() {
		dialogInstance = &dialogModel{}
	})
	return dialogInstance
}

func dialogDOToDialog(dialogDO* dataobject.UserDialogsDO) *mtproto.TLDialog {
	dialog := mtproto.NewTLDialog()
	// dialogData := dialog.GetData2()
	// draftIdList := make([]int32, 0)

	dialog.SetPinned(dialogDO.IsPinned == 1)
	dialog.SetPeer(base.ToPeerByTypeAndID(dialogDO.PeerType, dialogDO.PeerId))
	if dialogDO.PeerType == base.PEER_CHANNEL {
		// TODO(@benqi): only channel has pts
		// dialog.SetPts(messageBoxsDO.Pts)
		// peerChannlIdList = append(peerChannlIdList, dialogDO.PeerId)
		dialog.SetPts(dialogDO.Pts)
	}

	dialog.SetTopMessage(dialogDO.TopMessage)
	dialog.SetReadInboxMaxId(dialogDO.ReadInboxMaxId)
	dialog.SetReadOutboxMaxId(dialogDO.ReadOutboxMaxId)
	dialog.SetUnreadCount(dialogDO.UnreadCount)
	dialog.SetUnreadMentionsCount(dialogDO.UnreadMentionsCount)

	// NotifySettings
	peerNotifySettings := mtproto.NewTLPeerNotifySettings()
	peerNotifySettings.SetShowPreviews(dialogDO.ShowPreviews == 1)
	peerNotifySettings.SetSilent(dialogDO.Silent == 1)
	peerNotifySettings.SetMuteUntil(dialogDO.MuteUntil)
	peerNotifySettings.SetSound(dialogDO.Sound)
	dialog.SetNotifySettings(peerNotifySettings.To_PeerNotifySettings())
	return dialog
}

func dialogDOListToDialogList(dialogDOList []dataobject.UserDialogsDO) (dialogs []*mtproto.TLDialog) {
	draftIdList := make([]int32, 0)
	for _, dialogDO := range dialogDOList {
		if dialogDO.DraftId > 0 {
			draftIdList = append(draftIdList, dialogDO.DraftId)
		}
		dialogs = append(dialogs, dialogDOToDialog(&dialogDO))
	}

	// TODO(@benqi): fetch draft message list
	return
}

func (m *dialogModel) GetDialogsByOffsetId(userId int32, isPinned bool, offsetId int32, limit int32) (dialogs []*mtproto.TLDialog) {
	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPinnedAndOffset(
		userId, base2.BoolToInt8(isPinned), offsetId, limit)

	dialogs = dialogDOListToDialogList(dialogDOList)
	return
}

//func (m *dialogModel) GetDialogsByOffsetDate(userId int32, excludePinned bool, offsetData int32, limit int32) (dialogs []*mtproto.TLDialog) {
//	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectDialogsByPinnedAndOffsetDate(
//		userId, base2.BoolToInt8(!excludePinned), offsetData, limit)
//	for _, dialogDO := range dialogDOList {
//		dialogs = append(dialogs, dialogDOToDialog(&dialogDO))
//	}
//	return
//}

func (m *dialogModel) GetDialogsByUserIDAndType(userId, peerType int32) (dialogs []*mtproto.TLDialog) {
	dialogsDAO := dao.GetUserDialogsDAO(dao.DB_SLAVE)

	var dialogDOList []dataobject.UserDialogsDO
	if peerType == base.PEER_UNKNOWN || peerType == base.PEER_EMPTY {
		dialogDOList = dialogsDAO.SelectDialogsByUserID(userId)
		glog.Infof("SelectDialogsByUserID(%d) - {%v}", userId, dialogDOList)
	} else {
		dialogDOList = dialogsDAO.SelectDialogsByPeerType(userId, int8(peerType))
		glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogDOList)
	}

	dialogs = dialogDOListToDialogList(dialogDOList)
	// glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogs)
	return
}

func (m *dialogModel) GetPinnedDialogs(userId int32) (dialogs []*mtproto.TLDialog) {
	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectPinnedDialogs(userId)
	dialogs = dialogDOListToDialogList(dialogDOList)
	return
}

func (m *dialogModel) CreateOrUpdateByLastMessage(userId int32, peerType int32, peerId int32, topMessage int32, unreadMentions, inbox bool) (dialogId int32) {
	// TODO(@benqi): 事务
	// 创建会话
	slave := dao.GetUserDialogsDAO(dao.DB_SLAVE)
	master := dao.GetUserDialogsDAO(dao.DB_MASTER)

	date := int32(time.Now().Unix())

	dialog :=slave.SelectByPeer(userId, int8(peerType), peerId)
	if dialog == nil {
		dialog = &dataobject.UserDialogsDO{}
		dialog.UserId = userId
		dialog.PeerType = int8(peerType)
		dialog.PeerId = peerId
		if unreadMentions {
			dialog.UnreadMentionsCount = 1
		} else {
			dialog.UnreadMentionsCount = 0
		}
		dialog.UnreadCount = 1
		dialog.TopMessage = topMessage
		dialog.CreatedAt = base2.NowFormatYMDHMS()
		dialog.Date2 = date
		dialogId = int32(master.Insert(dialog))
	} else {
		if unreadMentions {
			dialog.UnreadMentionsCount += 1
		}
		if inbox {
			dialog.UnreadCount += 1
		}
		dialog.TopMessage = topMessage
		dialog.Date2 = date
		dialogId = dialog.Id
		master.UpdateTopMessage(topMessage, dialog.UnreadCount, dialog.UnreadMentionsCount, date, dialog.Id)
	}
	return
}
