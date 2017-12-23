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
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/model"
)

// android client request source code
/*
	TLRPC.TL_messages_getDialogs req = new TLRPC.TL_messages_getDialogs();
	req.limit = count;
	req.exclude_pinned = true;
	if (UserConfig.dialogsLoadOffsetId != -1) {
		if (UserConfig.dialogsLoadOffsetId == Integer.MAX_VALUE) {
			dialogsEndReached = true;
			serverDialogsEndReached = true;
			loadingDialogs = false;
			NotificationCenter.getInstance().postNotificationName(NotificationCenter.dialogsNeedReload);
			return;
		}
		req.offset_id = UserConfig.dialogsLoadOffsetId;
		req.offset_date = UserConfig.dialogsLoadOffsetDate;
		if (req.offset_id == 0) {
			req.offset_peer = new TLRPC.TL_inputPeerEmpty();
		} else {
			if (UserConfig.dialogsLoadOffsetChannelId != 0) {
				req.offset_peer = new TLRPC.TL_inputPeerChannel();
				req.offset_peer.channel_id = UserConfig.dialogsLoadOffsetChannelId;
			} else if (UserConfig.dialogsLoadOffsetUserId != 0) {
				req.offset_peer = new TLRPC.TL_inputPeerUser();
				req.offset_peer.user_id = UserConfig.dialogsLoadOffsetUserId;
			} else {
				req.offset_peer = new TLRPC.TL_inputPeerChat();
				req.offset_peer.chat_id = UserConfig.dialogsLoadOffsetChatId;
			}
			req.offset_peer.access_hash = UserConfig.dialogsLoadOffsetAccess;
		}
	} else {
		boolean found = false;
		for (int a = dialogs.size() - 1; a >= 0; a--) {
			TLRPC.TL_dialog dialog = dialogs.get(a);
			if (dialog.pinned) {
				continue;
			}
			int lower_id = (int) dialog.id;
			int high_id = (int) (dialog.id >> 32);
			if (lower_id != 0 && high_id != 1 && dialog.top_message > 0) {
				MessageObject message = dialogMessage.get(dialog.id);
				if (message != null && message.getId() > 0) {
					req.offset_date = message.messageOwner.date;
					req.offset_id = message.messageOwner.id;
					int id;
					if (message.messageOwner.to_id.channel_id != 0) {
						id = -message.messageOwner.to_id.channel_id;
					} else if (message.messageOwner.to_id.chat_id != 0) {
						id = -message.messageOwner.to_id.chat_id;
					} else {
						id = message.messageOwner.to_id.user_id;
					}
					req.offset_peer = getInputPeer(id);
					found = true;
					break;
				}
			}
		}
		if (!found) {
			req.offset_peer = new TLRPC.TL_inputPeerEmpty();
		}
	}
 */
// 由客户端代码: offset_id为当前用户最后一条消息ID，offset_peer为最后一条消息的接收者peer
// offset_date
// messages.getDialogs#191ba9c5 flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int = messages.Dialogs;
func (s *MessagesServiceImpl) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetDialogs - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.OffsetPeer)
	var dialogs []*mtproto.TLDialog
	if peer.PeerType == base.PEER_EMPTY {
		// 取出全部
	} else {
		// 通过message_boxs表检查offset_peer
		offsetMessageDO := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByUserIdAndMessageBoxId(md.UserId, request.OffsetId)
		// TODO(@benqi): date, access_hash check
		if offsetMessageDO == nil || ( peer.PeerType != int32(offsetMessageDO.PeerType)  && peer.PeerId != offsetMessageDO.PeerId) {
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "InputPeer invalid"))
		}
	}

	dialogs = model.GetDialogModel().GetDialogsByOffsetId(md.UserId, request.GetExcludePinned(), request.GetOffsetId(), request.GetLimit())
	messageDialogs := mtproto.NewTLMessagesDialogs()
	messageIdList := []int32{}
	userIdList := []int32{md.UserId}
	chatIdList := []int32{}
	for _, dialog := range dialogs {
		// dialog.Peer
		messageIdList = append(messageIdList, dialog.GetTopMessage())
		p := dialog.GetPeer()
		// TODO(@benqi): 先假设只有PEER_USER
		switch p.GetConstructor() {
		case mtproto.TLConstructor_CRC32_peerUser:
			userIdList = append(userIdList, p.GetData2().GetUserId())
		case mtproto.TLConstructor_CRC32_peerChat:
			chatIdList = append(chatIdList, p.GetData2().GetChatId())
		case mtproto.TLConstructor_CRC32_peerChannel:
		}
		messageDialogs.Data2.Dialogs = append(messageDialogs.Data2.Dialogs, dialog.To_Dialog())
	}
	glog.Infof("messageIdList - %v", messageIdList)
	if len(messageIdList) > 0 {
		messageDialogs.SetMessages(model.GetMessageModel().GetMessagesByPeerAndMessageIdList(md.UserId, messageIdList))
	}
	// userIdList = append(userIdList, md.UserId)
	users := model.GetUserModel().GetUserList(userIdList)
	for _, user := range users {
		if user.GetId() == md.UserId {
			user.SetSelf(true)
		} else {
			user.SetSelf(false)
		}
		user.SetContact(true)
		user.SetMutualContact(true)
		messageDialogs.Data2.Users = append(messageDialogs.Data2.Users, user.To_User())
	}
	if len(chatIdList) > 0 {
		messageDialogs.SetChats(model.GetChatModel().GetChatListByIDList(chatIdList))
	}
	// d, _ := json.Marshal(messageDialogs)
	glog.Infof("MessagesGetDialogs - reply: %s", logger.JsonDebugData(messageDialogs))
	return messageDialogs.To_Messages_Dialogs(), nil
}
