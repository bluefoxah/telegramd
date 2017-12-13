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

// messages.addChatUser#f9a0aa09 chat_id:int user_id:InputUser fwd_limit:int = Updates;
func (s *MessagesServiceImpl) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesAddChatUser - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): Impl MessagesAddChatUser logic
	//chat := model.GetChatModel().GetChat(request.ChatId)
	//participants := model.GetChatModel().GetChatParticipants(chat.Id)
	//var addChatUserId int32
	//// peer := base.fr(request.UserId)
	//switch request.UserId.Payload.(type) {
	//case *mtproto.InputUser_InputUser:
	//	addChatUserId = request.GetUserId().GetInputUser().GetUserId()
	//case *mtproto.InputUser_InputUserSelf:
	//	addChatUserId = md.UserId
	//case *mtproto.InputUser_InputUserEmpty:
	//	panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "InputPeer invalid"))
	//}
	//addChatParticipant := &mtproto.TLChatParticipant{}
	//addChatParticipant.UserId = addChatUserId
	//addChatParticipant.InviterId = md.UserId
	//addChatParticipant.Date = int32(time.Now().Unix())
	//model.GetChatModel().AddChatParticipant(chat.Id, addChatParticipant.UserId, addChatParticipant.UserId, 0)
	//dao.GetChatsDAO(dao.DB_MASTER).UpdateParticipantCount(chat.ParticipantsCount+1, chat.Version+1, chat.Id)
	//chat.ParticipantsCount += 1
	//chat.Version += 1
	//participants.Version = chat.Version
	//// participantUsers := participants.GetParticipants()
	//participants.Participants = append(participants.Participants, addChatParticipant.ToChatParticipant())
	//chatUserIdList := mtproto.GetUserIdListByChatParticipants(participants)
	//peer := &base.PeerUtil{}
	//peer.PeerType = base.PEER_CHAT
	//peer.PeerId = chat.Id
	//messageService := &mtproto.TLMessageService{}
	//messageService.Out = true
	//messageService.Date = chat.Date
	//messageService.FromId = md.UserId
	//messageService.ToId = peer.ToPeer()
	//// mtproto.MakePeer(&mtproto.TLPeerChat{chat.Id})
	//action := &mtproto.TLMessageActionChatAddUser {}
	//action.Users = append(action.Users, addChatUserId)
	//messageService.Action = action.ToMessageAction()
	//messageServiceId := model.GetMessageModel().CreateHistoryMessageService(md.UserId, peer, md.ClientMsgId, messageService)
	//messageService.Id = int32(messageServiceId)
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
	//	updates.Date = chat.Date
	//	if u.Id == md.UserId {
	//		// TODO(@benqi): Delete me
	//		updates.Chats = append(updates.Chats, chat.ToChat())
	//		reply = updates.ToUpdates()
	//		delivery2.GetDeliveryInstance().DeliveryUpdatesNotMe(
	//			md.AuthId,
	//			md.SessionId,
	//			md.NetlibSessionId,
	//			[]int32{u.Id},
	//			updates.ToUpdates().Encode())
	//	} else {
	//		updates.Chats = append(updates.Chats, chat.ToChat())
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
	//glog.Infof("MessagesAddChatUser - reply: {%v}", reply)
	//return
	return nil, fmt.Errorf("Not impl MessagesAddChatUser")
}
