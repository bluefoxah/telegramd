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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
)

// messages.editChatTitle#dc452855 chat_id:int title:string = Updates;
func (s *MessagesServiceImpl) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesEditChatTitle - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): Impl MessagesEditChatTitle logic
	// now := int32(time.Now().Unix())
	//_ = dao.GetChatsDAO(dao.DB_MASTER).UpdateTitle(request.Title, md.UserId, id.NextId(), base2.NowFormatYMDHMS(), request.ChatId)
	//chat := model.GetChatModel().GetChat(request.ChatId)
	//chat.Title = request.Title
	//participants := model.GetChatModel().GetChatParticipants(chat.Id)
	//peer := &base.PeerUtil{}
	//peer.PeerType = base.PEER_CHAT
	//peer.PeerId = chat.Id
	//messageService := &mtproto.TLMessageService{}
	//messageService.Out = true
	//messageService.Date = chat.Date
	//messageService.FromId = md.UserId
	//messageService.ToId = peer.ToPeer()
	//// mtproto.MakePeer(&mtproto.TLPeerChat{chat.Id})
	//action := &mtproto.TLMessageActionChatEditTitle{}
	//action.Title = request.Title
	//messageService.Action = action.ToMessageAction()
	//messageServiceId := model.GetMessageModel().CreateHistoryMessage2(md.UserId, peer, md.ClientMsgId, now, messageService.To_Message())
	//messageService.Id = int32(messageServiceId)
	//chatUserIdList := []int32{}
	//for _, participant := range participants.GetParticipants() {
	//	switch participant.Payload.(type) {
	//	case *mtproto.ChatParticipant_ChatParticipant:
	//		chatUserIdList = append(chatUserIdList, participant.GetChatParticipant().GetUserId())
	//	case *mtproto.ChatParticipant_ChatParticipantAdmin:
	//		chatUserIdList = append(chatUserIdList, participant.GetChatParticipantAdmin().GetUserId())
	//	case *mtproto.ChatParticipant_ChatParticipantCreator:
	//		chatUserIdList = append(chatUserIdList, participant.GetChatParticipantCreator().GetUserId())
	//	}
	//}
	//users := model.GetUserModel().GetUserList(chatUserIdList)
	//updateUsers := make([]*mtproto.User, 0, len(users))
	//for _, u := range users {
	//	u.Self = true
	//	updates := &mtproto.TLUpdates{}
	//	// 2. MessageBoxes
	//	pts := model.GetMessageModel().CreateMessageBoxes(u.Id, md.UserId, peer.PeerType, peer.PeerId, false, messageServiceId)
	//	// 3. dialog
	//	model.GetDialogModel().CreateOrUpdateByLastMessage(u.Id, peer.PeerType, peer.PeerId, messageServiceId, false)
	//	if u.GetId() == md.UserId {
	//		updateMessageID := &mtproto.TLUpdateMessageID{}
	//		updateMessageID.Id = int32(messageServiceId)
	//		updateMessageID.RandomId = md.ClientMsgId
	//		updates.Updates = append(updates.Updates, updateMessageID.ToUpdate())
	//		updates.Seq = 0
	//	} else {
	//		// TODO(@benqi): seq++
	//		updates.Seq = 0
	//	}
	//	updateChatParticipants := &mtproto.TLUpdateChatParticipants{}
	//	updateChatParticipants.Participants = participants.ToChatParticipants()
	//	updates.Updates = append(updates.Updates, updateChatParticipants.ToUpdate())
	//	updateNewMessage := &mtproto.TLUpdateNewMessage{}
	//	updateNewMessage.Pts = pts
	//	updateNewMessage.PtsCount = 1
	//	updateNewMessage.Message = messageService.ToMessage()
	//	updates.Updates = append(updates.Updates, updateNewMessage.ToUpdate())
	//	updates.Users = updateUsers
	//	updates.Chats = append(updates.Chats, chat.ToChat())
	//	updates.Date = chat.Date
	//	if u.Id == md.UserId {
	//		reply = updates.ToUpdates()
	//		delivery2.GetDeliveryInstance().DeliveryUpdatesNotMe(
	//			md.AuthId,
	//			md.SessionId,
	//			md.NetlibSessionId,
	//			[]int32{u.Id},
	//			updates.ToUpdates().Encode())
	//	} else {
	//		delivery2.GetDeliveryInstance().DeliveryUpdates(
	//			md.AuthId,
	//			md.SessionId,
	//			md.NetlibSessionId,
	//			[]int32{u.Id},
	//			updates.ToUpdates().Encode())
	//	}
	//	u.Self = false
	//}
	//for _, u := range users {
	//	// updates := &mtproto.TLUpdates{}
	//	if u.Id == md.UserId {
	//		u.Self = true
	//	}
	//	updateUsers = append(updateUsers, u.ToUser())
	//}
	//glog.Infof("MessagesEditChatTitle - reply: {%v}", reply)
	//return
	return nil, fmt.Errorf("Not impl MessagesEditChatTitle")
}
