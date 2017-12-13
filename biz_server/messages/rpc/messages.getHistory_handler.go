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
	"github.com/nebulaim/telegramd/biz_model/model"
	"math"
)

// messages.getHistory#afa92846 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesGetHistory - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.GetPeer())
	chatIdList := []int32{}
	userIdList := []int32{md.UserId}
	offsetId := request.OffsetId + request.AddOffset
	if offsetId <= 0 {
		offsetId = math.MaxInt32
	}
	messages := model.GetMessageModel().GetMessagesByUserIdPeerOffsetLimit(md.UserId, peer.PeerType, peer.PeerId, offsetId, request.Limit)
	for _, message := range messages {
		switch message.GetConstructor() {
		case mtproto.TLConstructor_CRC32_message:
			m := message.To_Message()
			userIdList = append(userIdList, m.GetFromId())
			p := base.FromPeer(m.GetToId())
			switch p.PeerType {
			case base.PEER_SELF, base.PEER_USER:
				userIdList = append(userIdList, p.PeerId)
			case base.PEER_CHAT:
				chatIdList = append(chatIdList, p.PeerId)
			case base.PEER_CHANNEL:
				// TODO(@benqi): add channel
			}
		case mtproto.TLConstructor_CRC32_messageService:
			m := message.To_MessageService()
			userIdList = append(userIdList, m.GetFromId())
			chatIdList = append(chatIdList, m.GetToId().GetData2().GetChatId())
		}
	}
	messagesMessages := mtproto.NewTLMessagesMessages()
	messagesMessages.SetMessages(messages)
	if len(userIdList) > 0 {
		users := model.GetUserModel().GetUserList(userIdList)
		for _, u := range users {
			if u.GetId() == md.UserId {
				u.SetSelf(true)
			}
			u.SetContact(true)
			messagesMessages.Data2.Users = append(messagesMessages.Data2.Users, u.To_User())
		}
	}
	if len(chatIdList) > 0 {
		messagesMessages.SetChats(model.GetChatModel().GetChatListByIDList(chatIdList))
	}
	glog.Infof("MessagesGetHistory - reply: %s", messagesMessages)
	return messagesMessages.To_Messages_Messages(), nil
}
