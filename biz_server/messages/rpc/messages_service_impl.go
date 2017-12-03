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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"time"
	"github.com/nebulaim/telegramd/biz_model/base"
	base2 "github.com/nebulaim/telegramd/base/base"
	"github.com/nebulaim/telegramd/biz_model/model"
	"github.com/nebulaim/telegramd/grpc_util"
	delivery2 "github.com/nebulaim/telegramd/biz_server/delivery"
	"github.com/nebulaim/telegramd/frontend/id"
)

type MessagesServiceImpl struct {
}

func (s *MessagesServiceImpl) MessagesSetTyping(ctx context.Context, request *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	glog.Infof("MessagesSetTyping - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputPeer(request.GetPeer())

	if peer.PeerType == base.PEER_SELF || peer.PeerType == base.PEER_USER {
		typing := &mtproto.TLUpdateUserTyping{}
		typing.UserId = md.UserId
		typing.Action = request.Action

		updates := &mtproto.TLUpdateShort{}
		updates.Update = mtproto.MakeUpdate(&mtproto.TLUpdateUserTyping{md.UserId, request.GetAction()})
		updates.Date = int32(time.Now().Unix())
		delivery2.GetDeliveryInstance().DeliveryUpdates(
			md.AuthId,
			md.SessionId,
			md.NetlibSessionId,
			[]int32{peer.PeerId},
			updates.ToUpdates().Encode())
	} else {
		// 其他的不需要推送
	}

	return mtproto.ToBool(true), nil
}

func (s *MessagesServiceImpl) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	glog.Infof("MessagesReportSpam - Process: {%v}", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputPeer(request.GetPeer())

	if peer.PeerType == base.PEER_USER || peer.PeerType == base.PEER_CHAT {
		// TODO(@benqi): 入库
	}

	return mtproto.ToBool(true), nil
}

func (s *MessagesServiceImpl) MessagesHideReportSpam(ctx context.Context, request *mtproto.TLMessagesHideReportSpam) (*mtproto.Bool, error) {
	glog.Infof("MessagesHideReportSpam - Process: {%v}", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputPeer(request.GetPeer())

	if peer.PeerType == base.PEER_USER || peer.PeerType == base.PEER_CHAT {
		// TODO(@benqi): 入库
	}

	return mtproto.ToBool(true), nil
}

func (s *MessagesServiceImpl) MessagesDiscardEncryption(ctx context.Context, request *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	glog.Infof("MessagesReportEncryptedSpam - Process: {%v}", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)
	// peer := base.FromInputPeer(request.GetPeer())
	//
	// if peer.PeerType == base.PEER_USER || peer.PeerType == base.PEER_CHAT {
	//	// TODO(@benqi): 入库
	// }

	return mtproto.ToBool(true), nil
}

func (s *MessagesServiceImpl) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSaveGif(ctx context.Context, request *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesSetInlineBotResults) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditInlineBotMessage(ctx context.Context, request *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesSetBotCallbackAnswer) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSaveRecentSticker(ctx context.Context, request *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetInlineGameScore(ctx context.Context, request *mtproto.TLMessagesSetInlineGameScore) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesToggleDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReorderPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetBotShippingResults(ctx context.Context, request *mtproto.TLMessagesSetBotShippingResults) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetBotPrecheckoutResults(ctx context.Context, request *mtproto.TLMessagesSetBotPrecheckoutResults) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesFaveSticker(ctx context.Context, request *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesGetMessagesViews(ctx context.Context,  request *mtproto.TLMessagesGetMessagesViews) (*mtproto.Vector<int32T>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

/*
message TL_messages_getHistory {
  InputPeer peer = 1;
  int32 offset_id = 2;
  int32 offset_date = 3;
  int32 add_offset = 4;
  int32 limit = 5;
  int32 max_id = 6;
  int32 min_id = 7;
};
 */
// messages.getHistory#afa92846 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
// messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	glog.Infof("MessagesGetHistory - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputPeer(request.Peer)

	chatIdList := []int32{}
	userIdList := []int32{}

	offsetId := request.OffsetId + request.AddOffset
	messages := model.GetMessageModel().GetMessagesByUserIdPeerOffsetLimit(md.UserId, peer.PeerType, peer.PeerId, offsetId, request.Limit)
	for _, message := range messages {
		switch message.Payload.(type) {
		case *mtproto.Message_Message:
			m := message.GetMessage()
			userIdList = append(userIdList, m.FromId)
			p := base.FromPeer(m.GetToId())
			switch p.PeerType {
			case base.PEER_SELF, base.PEER_USER:
				userIdList = append(userIdList, p.PeerId)
			case base.PEER_CHAT:
				chatIdList = append(chatIdList, p.PeerId)
			case base.PEER_CHANNEL:
				// TODO(@benqi): add channel
			}
		case *mtproto.Message_MessageService:
			m := message.GetMessageService()
			userIdList = append(userIdList, m.FromId)
			chatIdList = append(chatIdList, m.GetToId().GetPeerChat().GetChatId())
		}
	}

	messagesMessages := &mtproto.TLMessagesMessages{}
	messagesMessages.Messages = messages
	if len(userIdList) > 0 {
		users := model.GetUserModel().GetUserList(userIdList)
		for _, u := range users {
			if u.Id == md.UserId {
				u.Self = true
			}
			u.Contact = true
			messagesMessages.Users = append(messagesMessages.Users, u.ToUser())
		}
	}
	if len(chatIdList) > 0 {
		messagesMessages.Chats = model.GetChatModel().GetChatListByIDList(chatIdList)
	}

	glog.Infof("MessagesGetHistory - reply: {%v}", messagesMessages)
	return messagesMessages.ToMessages_Messages(), nil
}


// messages.search#39e9ea0 flags:# peer:InputPeer q:string from_id:flags.0?InputUser filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
// {peer:<inputPeerUser:<user_id:2 access_hash:5537087501845505974 > > filter:<inputMessagesFilterUrl:<> > }
func (s *MessagesServiceImpl) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	glog.Infof("MessagesSearch - Process: {%v}", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)
	_ = base.FromInputPeer(request.GetPeer())

	// TODO(@benqi): Not impl
	reply := &mtproto.TLMessagesMessages{}
	glog.Infof("MessagesGetHistory - reply: {%v}", reply)
	return reply.ToMessages_Messages(), nil
}

func (s *MessagesServiceImpl) MessagesSearchGlobal(ctx context.Context, request *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetUnreadMentions(ctx context.Context, request *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

/*
	query: { messages_getDialogs
	  flags: 1 [INT],
	  exclude_pinned: YES [ BY BIT 0 IN FIELD flags ],
	  offset_date: 0 [INT],
	  offset_id: 0 [INT],
	  offset_peer: { inputPeerEmpty },
	  limit: 20 [INT],
	},
 */
func (s *MessagesServiceImpl) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	glog.Infof("MessagesGetDialogs - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
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
	messageDialogs := &mtproto.TLMessagesDialogs{}

	messageIdList := []int32{}
	userIdList := []int32{}
	chatIdList := []int32{}

	for _, dialog := range dialogs {
		// dialog.Peer
		messageIdList = append(messageIdList, dialog.TopMessage)
		p := dialog.GetPeer()
		// TODO(@benqi): 先假设只有PEER_USER
		switch p.Payload.(type) {
		case *mtproto.Peer_PeerUser:
			userIdList = append(userIdList, p.GetPeerUser().GetUserId())
		case *mtproto.Peer_PeerChat:
			chatIdList = append(chatIdList, p.GetPeerChat().GetChatId())
		case *mtproto.Peer_PeerChannel:
		}
		messageDialogs.Dialogs = append(messageDialogs.Dialogs, dialog.ToDialog())
	}

	glog.Infof("messageIdList - %v", messageIdList)
	if len(messageIdList) > 0 {
		messageDialogs.Messages = model.GetMessageModel().GetMessagesByPeerAndMessageIdList(md.UserId, messageIdList)
	}

	// userIdList = append(userIdList, md.UserId)
	users := model.GetUserModel().GetUserList(userIdList)
	for _, user := range users {
		if user.Id == md.UserId {
			user.Self = true
		} else {
			user.Self = false
		}
		user.Contact = true
		user.MutualContact = true
		messageDialogs.Users = append(messageDialogs.Users, user.ToUser())
	}

	if len(chatIdList) > 0 {
		messageDialogs.Chats = model.GetChatModel().GetChatListByIDList(chatIdList)
	}

	// d, _ := json.Marshal(messageDialogs)
	glog.Infof("MessagesGetDialogs - reply: {%v}", messageDialogs)
	return messageDialogs.ToMessages_Dialogs(), nil
}

func (s *MessagesServiceImpl) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	glog.Infof("MessagesReadHistory - Process: {%v}", request)

/*
	md := grpc_util.RpcMetadataFromIncoming(ctx)

	var affected *mtproto.TLMessagesAffectedMessages = nil

	switch base.FromInputPeer(request.Peer) {
	case base.PEER_SELF:
	case base.PEER_USER:
		affected = model.GetUpdatesModel().GetAffectedMessage(md.UserId, request.MaxId)
	case base.PEER_CHAT:
	case base.PEER_CHANNEL:
	default:
	}
 */

 	// TODO(@benqi): 实现逻辑
	affected := &mtproto.TLMessagesAffectedMessages{}
	affected.Pts = -1
	affected.PtsCount = 0

	glog.Infof("MessagesReadHistory - reply: %v", affected)
	return affected.ToMessages_AffectedMessages(), nil
}

func (s *MessagesServiceImpl) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesReceivedMessages(ctx context.Context,  request *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector<ReceivedNotifyMessage>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }


/*
	// messages.sendMessage#fa88427a flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;
	message TL_messages_sendMessage {
	  bool no_webpage = 1;
	  bool silent = 2;
	  bool background = 3;
	  bool clear_draft = 4;
	  InputPeer peer = 5;
	  int32 reply_to_msg_id = 6;
	  string message = 7;
	  int64 random_id = 8;
	  ReplyMarkup reply_markup = 9;
	  repeated MessageEntity entities = 10;
	};

	// updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
	message TL_updateShortSentMessage {
	  bool out = 1;
	  int32 id = 2;
	  int32 pts = 3;
	  int32 pts_count = 4;
	  int32 date = 5;
	  MessageMedia media = 6;
	  repeated MessageEntity entities = 7;
	}

	// updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
	message TL_updateShortMessage {
	  bool out = 1;
	  bool mentioned = 2;
	  bool media_unread = 3;
	  bool silent = 4;
	  int32 id = 5;
	  int32 user_id = 6;
	  string message = 7;
	  int32 pts = 8;
	  int32 pts_count = 9;
	  int32 date = 10;
	  MessageFwdHeader fwd_from = 11;
	  int32 via_bot_id = 12;
	  int32 reply_to_msg_id = 13;
	  repeated MessageEntity entities = 14;
	}
 */
func (s *MessagesServiceImpl) MessagesSendMessage(ctx context.Context, request *mtproto.TLMessagesSendMessage) (reply *mtproto.Updates, err error) {
	glog.Infof("MessagesSendMessage - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	peer := base.FromInputPeer(request.GetPeer())

	message := &mtproto.TLMessage{}

	message.Silent = request.Silent
	// TODO(@benqi): ???
	// request.Background
	// request.NoWebpage
	// request.ClearDraft
	message.FromId = md.UserId
	if peer.PeerType == base.PEER_SELF {
		to := &mtproto.TLPeerUser{md.UserId}
		message.ToId = to.ToPeer()
	} else {
		message.ToId = peer.ToPeer()
	}
	message.Message = request.Message
	message.ReplyToMsgId = request.ReplyToMsgId
	message.ReplyMarkup = request.ReplyMarkup
	message.Entities = request.Entities
	message.Date = int32(time.Now().Unix())
	// glog.Infof("metadata: {%v}, rpcMetaData: {%v}", md, rpcMetaData)

	sentMessage := &mtproto.TLUpdateShortSentMessage{}
	switch peer.PeerType {
	case base.PEER_SELF:
		// 1. SaveMessage
		messageId := model.GetMessageModel().CreateHistoryMessage(md.UserId, peer, request.RandomId, message)
		// 2. MessageBoxes
		pts := model.GetMessageModel().CreateMessageBoxes(md.UserId, message.FromId, peer.PeerType, md.UserId, false, messageId)
		// 3. dialog
		model.GetDialogModel().CreateOrUpdateByLastMessage(md.UserId, peer.PeerType, md.UserId, messageId, message.Mentioned)

		// 推送给sync
		// 推给客户端的updates
		updates := mtproto.TLUpdateShortMessage{}
		updates.Id = int32(messageId)
		updates.UserId = md.UserId
		updates.Pts = pts
		updates.PtsCount = 1
		updates.Message = request.Message
		updates.Date = int32(time.Now().Unix())

		delivery2.GetDeliveryInstance().DeliveryUpdatesNotMe(
			md.AuthId,
			md.SessionId,
			md.NetlibSessionId,
			[]int32{md.UserId},
			updates.ToUpdates().Encode())

		// 返回给客户端
		// sentMessage := &mtproto.TLUpdateShortSentMessage{}
		sentMessage.Out = true
		sentMessage.Id = int32(messageId)
		sentMessage.Pts = pts
		sentMessage.PtsCount = 1
		sentMessage.Date = int32(time.Now().Unix())
		sentMessage.Media = mtproto.MakeMessageMedia(&mtproto.TLMessageMediaEmpty{})
		glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		// reply = sentMessage.ToUpdates()

	case base.PEER_CHAT:
		// 返回给客户端
		// sentMessage := &mtproto.TLUpdateShortSentMessage{}
		sentMessage.Out = true
		// sentMessage.Id = int32(messageId)
		// sentMessage.Pts = outPts
		sentMessage.PtsCount = 1
		sentMessage.Date = message.Date

		// 1. SaveMessage
		messageId := model.GetMessageModel().CreateHistoryMessage(md.UserId, peer, request.RandomId, message)

		participants := model.GetChatModel().GetChatParticipants(peer.PeerId)
		var userId int32 = 0
		for _, participan := range participants.GetParticipants() {
			switch participan.Payload.(type) {
			case *mtproto.ChatParticipant_ChatParticipantCreator:
				userId = participan.GetChatParticipantCreator().GetUserId()
			case *mtproto.ChatParticipant_ChatParticipantAdmin:
				userId = participan.GetChatParticipantAdmin().GetUserId()
			case *mtproto.ChatParticipant_ChatParticipant:
				userId = participan.GetChatParticipant().GetUserId()
			}

			// 2. MessageBoxes
			pts := model.GetMessageModel().CreateMessageBoxes(userId, md.UserId, peer.PeerType, peer.PeerId, false, messageId)
			model.GetDialogModel().CreateOrUpdateByLastMessage(userId, peer.PeerType, peer.PeerId, messageId, message.Mentioned)

			// inPts := model.GetMessageModel().CreateMessageBoxes(peer.PeerId, message.FromId, peer, true, messageId)
			// 3. dialog
			// model.GetDialogModel().CreateOrUpdateByLastMessage(peer.PeerId, peer, messageId, message.Mentioned)

			// 推送给sync
			// 推给客户端的updates
			updates := mtproto.TLUpdateShortChatMessage{}
			updates.Id = int32(messageId)
			updates.FromId = md.UserId
			updates.ChatId = peer.PeerId
			updates.Pts = pts
			updates.PtsCount = 1
			updates.Message = request.Message
			updates.Date = message.Date

			if md.UserId == userId {
				sentMessage.Id = int32(messageId)
				sentMessage.Pts = pts

				delivery2.GetDeliveryInstance().DeliveryUpdatesNotMe(
					md.AuthId,
					md.SessionId,
					md.NetlibSessionId,
					[]int32{userId},
					updates.ToUpdates().Encode())
			} else {
				delivery2.GetDeliveryInstance().DeliveryUpdates(
					md.AuthId,
					md.SessionId,
					md.NetlibSessionId,
					[]int32{userId},
					updates.ToUpdates().Encode())
			}
		}
		glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		// reply = sentMessage.ToUpdates()
	case base.PEER_USER:
		// 1. SaveMessage
		messageId := model.GetMessageModel().CreateHistoryMessage(md.UserId, peer, request.RandomId, message)
		// 2. MessageBoxes
		outPts := model.GetMessageModel().CreateMessageBoxes(md.UserId, message.FromId, peer.PeerType, peer.PeerId, false, messageId)
		inPts := model.GetMessageModel().CreateMessageBoxes(peer.PeerId, message.FromId, peer.PeerType, md.UserId, true, messageId)
		// 3. dialog
		model.GetDialogModel().CreateOrUpdateByLastMessage(md.UserId, peer.PeerType, peer.PeerId, messageId, message.Mentioned)
		model.GetDialogModel().CreateOrUpdateByLastMessage(peer.PeerId, peer.PeerType, md.UserId, messageId, message.Mentioned)

		// 推送给sync
		// 推给客户端的updates
		updates := mtproto.TLUpdateShortMessage{}
		updates.Id = int32(messageId)
		updates.UserId = md.UserId
		updates.Pts = inPts
		updates.PtsCount = 1
		updates.Message = request.Message
		updates.Date = message.Date

		delivery2.GetDeliveryInstance().DeliveryUpdatesNotMe(
			md.AuthId,
			md.SessionId,
			md.NetlibSessionId,
			[]int32{md.UserId, peer.PeerId},
			updates.ToUpdates().Encode())

		// 返回给客户端
		// sentMessage := &mtproto.TLUpdateShortSentMessage{}
		sentMessage.Out = true
		sentMessage.Id = int32(messageId)
		sentMessage.Pts = outPts
		sentMessage.PtsCount = 1
		sentMessage.Date = message.Date

		// glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		// reply = sentMessage.ToUpdates()
	case base.PEER_CHANNEL:

	default:
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "InputPeer invalid"))
	}

	glog.Infof("MessagesSendMessage - reply: {%v}", sentMessage)
	return sentMessage.ToUpdates(), nil
}

func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// messages.editChatTitle#dc452855 chat_id:int title:string = Updates;
func (s *MessagesServiceImpl) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	_ = dao.GetChatsDAO(dao.DB_MASTER).UpdateTitle(request.Title, md.UserId, id.NextId(), base2.NowFormatYMDHMS(), md.UserId)

	chat := &mtproto.TLChat{}
	chat.Id = request.ChatId
	chat.Title = request.Title
	chat.Photo = mtproto.MakeChatPhoto(&mtproto.TLChatPhotoEmpty{})
	chat.Date = int32(time.Now().Unix())
	chat.Version = 1

	updates := &mtproto.TLUpdates{}

	updates.Chats = append(updates.Chats, chat.ToChat())

	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	glog.Infof("MessagesEditChatPhoto - Process: {%v}", request)
	return nil, nil
}

// messages.addChatUser#f9a0aa09 chat_id:int user_id:InputUser fwd_limit:int = Updates;
func (s *MessagesServiceImpl) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// messages.deleteChatUser#e0611f16 chat_id:int user_id:InputUser = Updates;
func (s *MessagesServiceImpl) MessagesDeleteChatUser(ctx context.Context, request *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (s *MessagesServiceImpl) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (reply *mtproto.Updates, err error) {
	glog.Infof("MessagesCreateChat - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	randomId := md.ClientMsgId

	chatUserIdList := make([]int32, 0, len(request.GetUsers())+1)
	// chatUserIdList = append(chatUserIdList, md.UserId)

	for _, u := range request.Users {
		switch u.Payload.(type) {
		case *mtproto.InputUser_InputUserEmpty:
			// TODO(@benqi): ignore??
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "InputPeer invalid"))
		case *mtproto.InputUser_InputUserSelf:
			chatUserIdList = append(chatUserIdList, md.UserId)
		case *mtproto.InputUser_InputUser:
			chatUserIdList = append(chatUserIdList, u.GetInputUser().UserId)
		}
 	}

 	chat, participants := model.GetChatModel().CreateChat(md.UserId, request.Title, chatUserIdList, md.ClientMsgId)
	chatUserIdList = append(chatUserIdList, md.UserId)
	peer := &base.PeerUtil{}
	peer.PeerType = base.PEER_CHAT
	peer.PeerId = chat.Id

	messageService := &mtproto.TLMessageService{}
	messageService.Out = true
	messageService.Date = chat.Date
	messageService.FromId = md.UserId
	messageService.ToId = peer.ToPeer()
	// mtproto.MakePeer(&mtproto.TLPeerChat{chat.Id})
	action := &mtproto.TLMessageActionChatCreate{}
	action.Title = request.Title
	action.Users = chatUserIdList
	messageService.Action = action.ToMessageAction()

	messageServiceId := model.GetMessageModel().CreateHistoryMessageService(md.UserId, peer, md.ClientMsgId, messageService)
	messageService.Id = int32(messageServiceId)

	users := model.GetUserModel().GetUserList(chatUserIdList)
	updateUsers := make([]*mtproto.User, 0, len(users))
	for _, u := range users {
		u.Self = true
		updates := &mtproto.TLUpdates{}

		// 2. MessageBoxes
		pts := model.GetMessageModel().CreateMessageBoxes(u.Id, md.UserId, peer.PeerType, peer.PeerId, false, messageServiceId)
		// 3. dialog
		model.GetDialogModel().CreateOrUpdateByLastMessage(u.Id, peer.PeerType, peer.PeerId, messageServiceId, false)

		if u.GetId() == md.UserId {
			updateMessageID := &mtproto.TLUpdateMessageID{}
			updateMessageID.Id = int32(messageServiceId)
			updateMessageID.RandomId = randomId
			updates.Updates = append(updates.Updates, updateMessageID.ToUpdate())
			updates.Seq = 0
		} else {
			// TODO(@benqi): seq++
			updates.Seq = 0
		}

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{}
		updateChatParticipants.Participants = participants.ToChatParticipants()
		updates.Updates = append(updates.Updates, updateChatParticipants.ToUpdate())
		updateNewMessage := &mtproto.TLUpdateNewMessage{}
		updateNewMessage.Pts = pts
		updateNewMessage.PtsCount = 1
		updateNewMessage.Message = messageService.ToMessage()
		updates.Updates = append(updates.Updates, updateNewMessage.ToUpdate())

		updates.Users = updateUsers
		updates.Chats = append(updates.Chats, chat.ToChat())

		updates.Date = chat.Date

		if u.Id == md.UserId {
			reply = updates.ToUpdates()
			delivery2.GetDeliveryInstance().DeliveryUpdatesNotMe(
				md.AuthId,
				md.SessionId,
				md.NetlibSessionId,
				[]int32{u.Id},
				updates.ToUpdates().Encode())
		} else {
			delivery2.GetDeliveryInstance().DeliveryUpdates(
				md.AuthId,
				md.SessionId,
				md.NetlibSessionId,
				[]int32{u.Id},
				updates.ToUpdates().Encode())
		}
		u.Self = false
	}

	for _, u := range users {
		// updates := &mtproto.TLUpdates{}
		if u.Id == md.UserId {
			u.Self = true
		}
		updateUsers = append(updateUsers, u.ToUser())
	}

	glog.Infof("MessagesCreateChat - reply: {%v}", reply)
	return
}

func (s *MessagesServiceImpl) MessagesForwardMessage(ctx context.Context, request *mtproto.TLMessagesForwardMessage) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesStartBot(ctx context.Context, request *mtproto.TLMessagesStartBot) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesToggleChatAdmins(ctx context.Context, request *mtproto.TLMessagesToggleChatAdmins) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesMigrateChat(ctx context.Context, request *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendInlineBotResult(ctx context.Context, request *mtproto.TLMessagesSendInlineBotResult) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetAllDrafts(ctx context.Context, request *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetGameScore(ctx context.Context, request *mtproto.TLMessagesSetGameScore) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendScreenshotNotification(ctx context.Context, request *mtproto.TLMessagesSendScreenshotNotification) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetPeerSettings(ctx context.Context, request *mtproto.TLMessagesGetPeerSettings) (*mtproto.PeerSettings, error) {
	glog.Infof("MessagesGetPeerSettings - Process: {%v}", request)

	settings := &mtproto.TLPeerSettings{}
	settings.ReportSpam = false

	glog.Infof("MessagesGetPeerSettings - reply: {%v}\n", settings)
	return settings.ToPeerSettings(), nil

}

func (s *MessagesServiceImpl) MessagesGetChats(ctx context.Context, request *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	glog.Infof("MessagesGetFullChat - Process: {%v}", request)
	md := grpc_util.RpcMetadataFromIncoming(ctx)

	messagesChatFull := &mtproto.TLMessagesChatFull{}

	chatFull := model.GetChatModel().GetChatFull(request.ChatId)
	peer := &base.PeerUtil{}
	peer.PeerType = base.PEER_CHAT
	peer.PeerId = request.ChatId
	chatFull.NotifySettings = model.GetAccountModel().GetNotifySettings(md.UserId, peer)

	chat := model.GetChatModel().GetChat(request.ChatId)
	// chat.ParticipantsCount = len(chatFull.GetParticipants().GetChatParticipants().GetParticipants())

	chatUserIdList := make([]int32, 0)
	participants := chatFull.GetParticipants().GetChatParticipants().GetParticipants()
	for _, participant := range participants {
		switch participant.Payload.(type) {
		case *mtproto.ChatParticipant_ChatParticipantCreator:
			chatUserIdList = append(chatUserIdList, md.UserId)
		case *mtproto.ChatParticipant_ChatParticipant:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipant().GetUserId())
		case *mtproto.ChatParticipant_ChatParticipantAdmin:
			chatUserIdList = append(chatUserIdList, participant.GetChatParticipantAdmin().GetUserId())
		}
	}
	chat.ParticipantsCount = int32(len(participants))

	users := model.GetUserModel().GetUserList(chatUserIdList)
	for _, u := range users {
		if u.Id == md.UserId {
			u.Self = true
		}
		messagesChatFull.Users = append(messagesChatFull.Users, u.ToUser())
	}
	messagesChatFull.Chats = append(messagesChatFull.Chats, chat.ToChat())
	messagesChatFull.FullChat = chatFull.ToChatFull()

	glog.Infof("MessagesGetFullChat - Process: {%v}", chatFull)
	return messagesChatFull.ToMessages_ChatFull(), nil
}

func (s *MessagesServiceImpl) MessagesGetDhConfig(ctx context.Context, request *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesRequestEncryption(ctx context.Context, request *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesAcceptEncryption(ctx context.Context, request *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendEncrypted(ctx context.Context, request *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendEncryptedFile(ctx context.Context, request *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendEncryptedService(ctx context.Context, request *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesReceivedQueue(ctx context.Context,  request *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector<int64T>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	glog.Infof("MessagesGetAllStickers - Process: {%v}", request)

	stickers := &mtproto.TLMessagesAllStickersNotModified{}

	glog.Infof("MessagesGetAllStickers - reply: {%v}\n", stickers)
	return stickers.ToMessages_AllStickers(), nil
}

func (s *MessagesServiceImpl) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	glog.Infof("MessagesGetMaskStickers - Process: {%v}", request)

	stickers := &mtproto.TLMessagesAllStickersNotModified{}

	glog.Infof("MessagesGetMaskStickers - reply: {%v}\n", stickers)
	return stickers.ToMessages_AllStickers(), nil
}

func (s *MessagesServiceImpl) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesUploadMedia(ctx context.Context, request *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesInstallStickerSet(ctx context.Context, request *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSearchGifs(ctx context.Context, request *mtproto.TLMessagesSearchGifs) (*mtproto.Messages_FoundGifs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	glog.Infof("MessagesGetSavedGifs - Process: {%v}", request)

	stickers := &mtproto.TLMessagesSavedGifsNotModified{}

	glog.Infof("MessagesGetSavedGifs - reply: {%v}\n", stickers)
	return stickers.ToMessages_SavedGifs(), nil
}

func (s *MessagesServiceImpl) MessagesGetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesGetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesGetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	glog.Infof("MessagesGetPinnedDialogs - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	dialogs := model.GetDialogModel().GetPinnedDialogs(md.UserId)
	peerDialogs := &mtproto.TLMessagesPeerDialogs{}

	messageIdList := []int32{}
	userIdList := []int32{md.UserId}
	chatIdList := []int32{}

	for _, dialog := range dialogs {
		// dialog.Peer
		messageIdList = append(messageIdList, dialog.TopMessage)
		peer := base.FromPeer(dialog.GetPeer())
		// TODO(@benqi): 先假设只有PEER_USER
		if peer.PeerType == base.PEER_USER {
			userIdList = append(userIdList, dialog.Peer.GetPeerUser().UserId)
		} else if peer.PeerType == base.PEER_SELF {
			userIdList = append(userIdList, md.UserId)
		} else if peer.PeerType == base.PEER_CHAT {
			chatIdList = append(chatIdList, dialog.Peer.GetPeerChat().ChatId)
		}
		peerDialogs.Dialogs = append(peerDialogs.Dialogs, dialog.ToDialog())
	}

	glog.Infof("messageIdList - %v", messageIdList)
	if len(messageIdList) > 0 {
		peerDialogs.Messages = model.GetMessageModel().GetMessagesByPeerAndMessageIdList(md.UserId, messageIdList)
	}

	users := model.GetUserModel().GetUserList(userIdList)
	for _, user := range users {
		if user.Id == md.UserId {
			user.Self = true
		} else {
			user.Self = false
		}
		user.Contact = true
		user.MutualContact = true
		peerDialogs.Users = append(peerDialogs.Users, user.ToUser())
	}

	if len(chatIdList) > 0 {
		peerDialogs.Chats = model.GetChatModel().GetChatListByIDList(chatIdList)
	}

	state := &mtproto.TLUpdatesState{}
	state.Date = int32(time.Now().Unix())
	state.Pts = model.GetMessageModel().GetLastPtsByUserId(md.UserId)
	state.Qts = 0
	state.Seq = 1
	state.UnreadCount = 0
	peerDialogs.State = state.ToUpdates_State()

	glog.Infof("MessagesGetPinnedDialogs - reply: {%v}\n", peerDialogs)
	return peerDialogs.ToMessages_PeerDialogs(), nil
}

func (s *MessagesServiceImpl) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	glog.Infof("MessagesGetFeaturedStickers - Process: {%v}", request)

	stickers := &mtproto.TLMessagesFeaturedStickersNotModified{}

	glog.Infof("MessagesGetFeaturedStickers - reply: {%v}\n", stickers)
	return stickers.ToMessages_FeaturedStickers(), nil
}

func (s *MessagesServiceImpl) MessagesGetRecentStickers(ctx context.Context, request *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	glog.Infof("MessagesGetRecentStickers - Process: {%v}", request)

	stickers := &mtproto.TLMessagesRecentStickersNotModified{}

	glog.Infof("MessagesGetPinnedDialogs - reply: {%v}\n", stickers)
	return stickers.ToMessages_RecentStickers(), nil
}

func (s *MessagesServiceImpl) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	glog.Infof("MessagesGetArchivedStickers - Process: {%v}", request)

	stickers := &mtproto.TLMessagesArchivedStickers{}
	stickers.Count = 0

	glog.Infof("MessagesGetArchivedStickers - reply: {%v}\n", stickers)
	return stickers.ToMessages_ArchivedStickers(), nil
}

// func (s *MessagesServiceImpl)MessagesGetAttachedStickers(ctx context.Context,  request *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector<StickerSetCovered>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesGetGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetInlineGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetInlineGameHighScores) (*mtproto.Messages_HighScores, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	glog.Infof("MessagesGetFavedStickers - Process: {%v}", request)

	stickers := &mtproto.TLMessagesFavedStickersNotModified{}

	glog.Infof("MessagesGetFavedStickers - reply: {%v}\n", stickers)
	return stickers.ToMessages_FavedStickers(), nil
}
