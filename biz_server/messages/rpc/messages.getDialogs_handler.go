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

// messages.getDialogs#191ba9c5 flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int = messages.Dialogs;
func (s *MessagesServiceImpl) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetDialogs - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.OffsetPeer)
	var dialogs []*mtproto.TLDialog
	if peer.PeerType == base.PEER_EMPTY {
		// 取出全部
	} else {
		dialogDO := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPeer(md.UserId, int8(peer.PeerType), peer.PeerId)
		// TODO(@benqi): date, access_hash check
		if dialogDO == nil || request.OffsetId > dialogDO.TopMessage {
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "InputPeer invalid"))
		}
	}
	dialogs = model.GetDialogModel().GetDialogsByOffsetDate(md.UserId, request.ExcludePinned, request.OffsetDate, request.Limit)
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
