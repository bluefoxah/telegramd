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
	//"github.com/nebulaim/telegramd/mtproto"
	//"github.com/nebulaim/telegramd/biz_model/base"
	//"github.com/nebulaim/telegramd/biz_model/dal/dao"
	//"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	//"github.com/golang/glog"
	//base2 "github.com/nebulaim/telegramd/base/base"
	//"time"
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

/*
func dialogDOToDialog(dialogDO* dataobject.UserDialogsDO) *mtproto.TLDialog {
	dialog := &mtproto.TLDialog{}
	dialog.Pinned = dialogDO.IsPinned == 1
	switch dialogDO.PeerType {
	case base.PEER_EMPTY:
	case base.PEER_SELF, base.PEER_USER:
		peer := &mtproto.TLPeerUser{dialogDO.PeerId}
		dialog.Peer = peer.ToPeer()
	case base.PEER_CHAT:
		peer := &mtproto.TLPeerChat{dialogDO.PeerId}
		dialog.Peer = peer.ToPeer()
	case base.PEER_CHANNEL:
		peer := &mtproto.TLPeerChannel{dialogDO.PeerId}
		dialog.Peer = peer.ToPeer()
	}
	dialog.TopMessage = dialogDO.TopMessage
	dialog.ReadInboxMaxId = dialogDO.ReadInboxMaxId
	dialog.ReadOutboxMaxId = dialogDO.ReadOutboxMaxId
	dialog.UnreadCount = dialogDO.UnreadCount
	dialog.UnreadMentionsCount = dialogDO.UnreadMentionsCount

	// TODO(@benqi): pts/draft
	// NotifySettings
	peerNotifySettings := &mtproto.TLPeerNotifySettings{}
	peerNotifySettings.ShowPreviews = true
	peerNotifySettings.MuteUntil = 0
	peerNotifySettings.Sound = "default"
	dialog.NotifySettings = peerNotifySettings.ToPeerNotifySettings()

	return dialog
}

// dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
//message TL_dialog {
//	bool pinned = 1;
//	Peer peer = 2;
//	int32 top_message = 3;
//	int32 read_inbox_max_id = 4;
//	int32 read_outbox_max_id = 5;
//	int32 unread_count = 6;
//	int32 unread_mentions_count = 7;
//	PeerNotifySettings notify_settings = 8;
//	int32 pts = 9;
//	DraftMessage draft = 10;
//}

//func (m *dialogModel) GetDialog(userId int32, offsetPeer *base.PeerUtil) (dialog *mtproto.TLDialog) {
//	slave := dao.GetUserDialogsDAO(dao.DB_SLAVE)
//
//	dialogDO := slave.SelectByPeer(userId, int8(offsetPeer.PeerType), offsetPeer.PeerId)
//	// _ = dialog
//	// dialog := slave.SelectByPeer(userId, int8(offsetPeer.PeerType), offsetPeer.PeerId)
//
//}
//
//
/// *
//  exclude_pinned: YES [ BY BIT 0 IN FIELD flags ],
//  offset_date: 0 [INT],
//  offset_id: 0 [INT],
//  offset_peer: { inputPeerEmpty },
// *  ///
func (m *dialogModel) GetDialogsByOffsetDate(userId int32, excludePinned bool, offsetData int32, limit int32) (dialogs []*mtproto.TLDialog) {
	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectDialogsByPinnedAndOffsetDate(
		userId, base2.BoolToInt8(!excludePinned), offsetData, limit)
	for _, dialogDO := range dialogDOList {
		dialog := &mtproto.TLDialog{}
		dialog.Pinned = dialogDO.IsPinned == 1
		switch dialogDO.PeerType {
		case base.PEER_EMPTY:
			continue
		case base.PEER_SELF, base.PEER_USER:
			peer := &mtproto.TLPeerUser{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		case base.PEER_CHAT:
			peer := &mtproto.TLPeerChat{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		case base.PEER_CHANNEL:
			peer := &mtproto.TLPeerChannel{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		}
		dialog.TopMessage = dialogDO.TopMessage
		dialog.ReadInboxMaxId = dialogDO.ReadInboxMaxId
		dialog.ReadOutboxMaxId = dialogDO.ReadOutboxMaxId
		dialog.UnreadCount = dialogDO.UnreadCount
		dialog.UnreadMentionsCount = dialogDO.UnreadMentionsCount

		// TODO(@benqi): pts/draft
		// NotifySettings
		peerNotifySettings := &mtproto.TLPeerNotifySettings{}
		peerNotifySettings.ShowPreviews = true
		peerNotifySettings.MuteUntil = 0
		peerNotifySettings.Sound = "default"
		dialog.NotifySettings = peerNotifySettings.ToPeerNotifySettings()

		dialogs = append(dialogs, dialog)
	}
	return
}

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

	// []do.UserDialogsDO
	// dialogDOList, _ := dialogsDAO.SelectDialogsByUserID(userId)
	dialogs = []*mtproto.TLDialog{}
	for _, dialogDO := range dialogDOList {
		dialog := &mtproto.TLDialog{}
		dialog.Pinned = dialogDO.IsPinned == 1
		switch dialogDO.PeerType {
		case base.PEER_EMPTY:
			continue
		case base.PEER_SELF, base.PEER_USER:
			peer := &mtproto.TLPeerUser{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		case base.PEER_CHAT:
			peer := &mtproto.TLPeerChat{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		case base.PEER_CHANNEL:
			peer := &mtproto.TLPeerChannel{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		}
		dialog.TopMessage = dialogDO.TopMessage
		dialog.ReadInboxMaxId = dialogDO.ReadInboxMaxId
		dialog.ReadOutboxMaxId = dialogDO.ReadOutboxMaxId
		dialog.UnreadCount = dialogDO.UnreadCount
		dialog.UnreadMentionsCount = dialogDO.UnreadMentionsCount

		// TODO(@benqi): pts/draft
		// NotifySettings
		peerNotifySettings := &mtproto.TLPeerNotifySettings{}
		peerNotifySettings.ShowPreviews = true
		peerNotifySettings.MuteUntil = 0
		peerNotifySettings.Sound = "default"
		dialog.NotifySettings = peerNotifySettings.ToPeerNotifySettings()

		dialogs = append(dialogs, dialog)
	}

	glog.Infof("SelectDialogsByPeerType(%d, %d) - {%v}", userId, int32(peerType), dialogs)
	return
}

func (m *dialogModel) GetPinnedDialogs(userId int32) (dialogs []*mtproto.TLDialog) {
	dialogDOList := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectPinnedDialogs(userId)
	for _, dialogDO := range dialogDOList {
		dialog := &mtproto.TLDialog{}
		dialog.Pinned = dialogDO.IsPinned == 1
		switch dialogDO.PeerType {
		case base.PEER_EMPTY:
			continue
		case base.PEER_SELF, base.PEER_USER:
			peer := &mtproto.TLPeerUser{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		case base.PEER_CHAT:
			peer := &mtproto.TLPeerChat{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		case base.PEER_CHANNEL:
			peer := &mtproto.TLPeerChannel{dialogDO.PeerId}
			dialog.Peer = peer.ToPeer()
		}
		dialog.TopMessage = dialogDO.TopMessage
		dialog.ReadInboxMaxId = dialogDO.ReadInboxMaxId
		dialog.ReadOutboxMaxId = dialogDO.ReadOutboxMaxId
		dialog.UnreadCount = dialogDO.UnreadCount
		dialog.UnreadMentionsCount = dialogDO.UnreadMentionsCount

		// TODO(@benqi): pts/draft
		// NotifySettings
		peerNotifySettings := &mtproto.TLPeerNotifySettings{}
		peerNotifySettings.ShowPreviews = true
		peerNotifySettings.MuteUntil = 0
		peerNotifySettings.Sound = "default"
		dialog.NotifySettings = peerNotifySettings.ToPeerNotifySettings()

		dialogs = append(dialogs, dialog)
	}
	return
}

//func (m *dialogModel) UpdateTopMessage(dialogId, topMessage int32) {
//	dialogsDAO := dao.GetUserDialogsDAO(dao.DB_MASTER)
//	dialogsDAO.UpdateTopMessage(topMessage, dialogId)
//}

func (m *dialogModel) CreateOrUpdateByLastMessage(userId int32, peerType int32, peerId int32, topMessage int32, unreadMentions bool) (dialogId int32) {
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
		dialog.UnreadCount += 1
		dialog.TopMessage = topMessage
		dialog.Date2 = date
		dialogId = dialog.Id
		master.UpdateTopMessage(topMessage, dialog.UnreadCount, dialog.UnreadMentionsCount, date, dialog.Id)
	}
	return
}

*/
