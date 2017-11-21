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
	"github.com/nebulaim/telegramd/biz_server/sync_client"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	base2 "github.com/nebulaim/telegramd/base/base"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/biz_model/model"
	"github.com/nebulaim/telegramd/grpc_util"
)

type MessagesServiceImpl struct {
	SyncRPCClient *sync_client.SyncRPCClient
}

func (s *MessagesServiceImpl) MessagesSetTyping(ctx context.Context, request *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	glog.Infof("MessagesSetTyping - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	var peerUserId = int32(0)

	var reply *mtproto.Bool = nil
	switch request.Peer.Payload.(type) {
	case *mtproto.InputPeer_InputPeerUser:
		b := &mtproto.TLBoolTrue{}
		reply = b.ToBool()
		peerUserId = request.Peer.GetInputPeerUser().UserId
		// mtproto.MakeBool(&mtproto.TLBoolTrue{})
		// typing.UserId = request.Peer.Payload.(*mtproto.InputPeer_InputPeerUser).InputPeerUser.UserId
	case *mtproto.InputPeer_InputPeerChat:
		b := &mtproto.TLBoolTrue{}
		reply = b.ToBool()
		// reply = mtproto.MakeBool(&mtproto.TLBoolTrue{})
		// typing.UserId = request.Peer.Payload.(*mtproto.InputPeer_InputPeerChat).InputPeerChat.ChatId
	default:
		glog.Errorf("MessagesSetTyping - BadRequest!")
		// reply = mtproto.MakeBool(&mtproto.TLBoolFalse{})
		b := &mtproto.TLBoolTrue{}
		reply = b.ToBool()
		return reply, nil
	}

	reply = mtproto.MakeBool(&mtproto.TLBoolTrue{})
	glog.Infof("MessagesSetTyping - reply: {%v}\n", reply)

	if peerUserId != 0 {
		// TODO(@benqi): Dispatch to updates
		// var update *mtproto.Update
		// updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
		// updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
		// 转发
		typing := &mtproto.TLUpdateUserTyping{}
		typing.UserId = md.UserId
		typing.Action = request.Action

		updates := &mtproto.TLUpdateShort{}
		updates.Update = typing.ToUpdate()
		updates.Date = int32(time.Now().Unix())

		delivery := &zproto.DeliveryUpdatesToUsers{}
		delivery.MyAuthKeyId = md.AuthId
		delivery.MySessionId = md.SessionId
		delivery.SendtoUserIdList = append(delivery.SendtoUserIdList, peerUserId)
		updateRawData := updates.ToUpdates().Encode()
		delivery.RawData = updateRawData
		glog.Infof("MessagesSendMessage - delivery: %v", delivery)
		s.SyncRPCClient.Client.DeliveryUpdates(context.Background(), delivery)
	}

	return reply, nil
}

func (s *MessagesServiceImpl) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesHideReportSpam(ctx context.Context, request *mtproto.TLMessagesHideReportSpam) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
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
	glog.Info("Process: %v", request)
	return nil, nil
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

	messagesMessages := &mtproto.TLMessagesMessages{}
	peer := base.FromInputPeer(request.Peer)

	chatIdList := []int32{}
	_ = chatIdList
	userIdList := []int32{md.UserId}

	switch peer.PeerType {
	case base.PEER_EMPTY:
	case base.PEER_CHAT:
		// TODO(@benqi): chats
	case base.PEER_USER:
		ms := model.GetMessageModel().GetMessagesByUserIdPeerOffsetLimit(md.UserId, peer.PeerType, request.Peer.GetInputPeerUser().UserId, request.OffsetId, request.Limit)
		for _, message := range ms {
			messagesMessages.Messages = append(messagesMessages.Messages, message.ToMessage())
		}
		userIdList = append(userIdList, request.Peer.GetInputPeerUser().UserId)
	case base.PEER_SELF:
	case base.PEER_CHANNEL:
	default:
	}

	users := model.GetUserModel().GetUserList(userIdList)
	for _, u := range users {
		messagesMessages.Users = append(messagesMessages.Users, u.ToUser())
	}

	glog.Infof("MessagesGetHistory - reply: {%v}", messagesMessages)
	return messagesMessages.ToMessages_Messages(), nil
}

func (s *MessagesServiceImpl) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	glog.Infof("MessagesSearch - Process: {%v}", request)



	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSearchGlobal(ctx context.Context, request *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetUnreadMentions(ctx context.Context, request *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// messages.getDialogs#191ba9c5
// 	flags:#
// 		exclude_pinned:flags.0?true
// 	offset_date:int
//  offset_id:int
//  offset_peer:InputPeer
//  limit:int = messages.Dialogs;
// messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
// dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
// Message message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
// Chat
// User
func (s *MessagesServiceImpl) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	glog.Infof("MessagesGetDialogs - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	peer := base.FromInputPeer(request.OffsetPeer)
	messageDialogs := &mtproto.TLMessagesDialogs{}

	dialogs := model.GetDialogModel().GetDialogsByUserIDAndType(md.UserId, peer.PeerType)
	messageIdList := []int32{}
	userIdList := []int32{md.UserId}
	chatIdList := []int32{}
	for _, dialog := range dialogs {
		// dialog.Peer
		messageIdList = append(messageIdList, dialog.TopMessage)
		// TODO(@benqi): 先假设只有PEER_USER
		//if ptype == base.PEER_USER {
		//	userIdList = append(userIdList, dialog.Peer.GetPeerUser().UserId)
		//} else if ptype == base.PEER_CHAT {
		//	chatIdList = append(chatIdList, dialog.Peer.GetPeerChat().ChatId)
		//}
		userIdList = append(userIdList, dialog.Peer.GetPeerUser().UserId)
		messageDialogs.Dialogs = append(messageDialogs.Dialogs, dialog.ToDialog())

	}

	if len(messageIdList) > 0 {
		messages := model.GetMessageModel().GetMessagesByIDList(messageIdList)
		for _, message := range messages {
			messageDialogs.Messages = append(messageDialogs.Messages, message.ToMessage())
		}
	}

	if len(userIdList) > 0 {
		users := model.GetUserModel().GetUserList(userIdList)
		for _, user := range users {
			messageDialogs.Users = append(messageDialogs.Users, user.ToUser())
		}
	}

	if len(chatIdList) > 0 {
		// TODO(@benqi): 还未实现chat
		//chats := model.GetChatModel().GetChatList(chatIdList)
		//for _, chat := range chats {
		//	messageDialogs.Users = append(messageDialogs.Users, chat.ToChat())
		//}
	}

	glog.Infof("MessagesGetDialogs - reply: %v", messageDialogs)
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
	// glog.Infof("metadata: {%v}, rpcMetaData: {%v}", md, rpcMetaData)

	// TODO(@benqi): 仅仅验证逻辑，未考虑出错处理等，
	sentMessage := &mtproto.TLUpdateShortSentMessage{}
	_ = sentMessage
	switch request.Peer.Payload.(type) {
	case *mtproto.InputPeer_InputPeerEmpty:
	case *mtproto.InputPeer_InputPeerSelf:
		//TODO(@benqi): 消除重复代码

		// TODO(@benqi): 事务
		// 创建会话
		dialog := dao.GetUserDialogsDAO(dao.DB_SLAVE).CheckExists(md.UserId, base.PEER_USER, md.UserId)
		if dialog == nil {
			dialog = &dataobject.UserDialogsDO{}
			dialog.UserId = md.UserId
			dialog.PeerType = base.PEER_USER
			dialog.PeerId = md.UserId
			dao.GetUserDialogsDAO(dao.DB_MASTER).Insert(dialog)
		}

		// 插入消息
		message := &dataobject.MessagesDO{}
		message.UserId = md.UserId
		message.PeerType = base.PEER_USER
		message.PeerId = md.UserId
		message.RandomId = request.RandomId
		message.Message = request.Message
		messageId := dao.GetMessagesDAO(dao.DB_MASTER).Insert(message)

		// inbox和outbox
		// 发件箱
		messageBox := &dataobject.MessageBoxsDO{}
		messageBox.UserId = md.UserId
		messageBox.MessageBoxType = 0
		messageBox.MessageId = int32(messageId)
		outPts := model.GetSequenceModel().NextID(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(outPts)
		dao.GetMessageBoxsDAO(dao.DB_MASTER).Insert(messageBox)

		// 推送给sync
		// 推给客户端的updates
		update := mtproto.TLUpdateShortMessage{}
		update.Id = int32(messageId)
		update.UserId = md.UserId
		update.Pts = int32(outPts)
		update.PtsCount = 1
		update.Message = request.Message
		update.Date = int32(time.Now().Unix())
		updateRawData := update.ToUpdates().Encode()

		delivery := &zproto.DeliveryUpdatesToUsers{}
		delivery.MyAuthKeyId = md.AuthId
		delivery.MySessionId = md.SessionId
		delivery.SendtoUserIdList = append(delivery.SendtoUserIdList, md.UserId)
		delivery.RawData = updateRawData

		glog.Infof("MessagesSendMessage - delivery: %v", delivery)
		_, _ = s.SyncRPCClient.Client.DeliveryUpdates(context.Background(), delivery)
		// update.Encode()

		// 返回给客户端
		sentMessage := &mtproto.TLUpdateShortSentMessage{}
		sentMessage.Out = true
		sentMessage.Id = int32(messageId)
		sentMessage.Pts = int32(outPts)
		sentMessage.PtsCount = 1
		sentMessage.Date = int32(time.Now().Unix())

		glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		reply = sentMessage.ToUpdates()

	case *mtproto.InputPeer_InputPeerChat:
	case *mtproto.InputPeer_InputPeerUser:
		inputPeerUser := request.Peer.GetInputPeerUser()

		// TODO(@benqi): 事务
		// 创建会话
		dialog1 := dao.GetUserDialogsDAO(dao.DB_SLAVE).CheckExists(md.UserId, base.PEER_USER, inputPeerUser.UserId)
		if dialog1 == nil {
			dialog1 = &dataobject.UserDialogsDO{}
			dialog1.UserId = md.UserId
			dialog1.PeerType = base.PEER_USER
			dialog1.PeerId = inputPeerUser.UserId
			dao.GetUserDialogsDAO(dao.DB_MASTER).Insert(dialog1)
		}
		dialog2 := dao.GetUserDialogsDAO(dao.DB_SLAVE).CheckExists(inputPeerUser.UserId, base.PEER_USER, md.UserId)
		if dialog2 == nil {
			dialog2 = &dataobject.UserDialogsDO{}
			dialog2.UserId = inputPeerUser.UserId
			dialog2.PeerType = base.PEER_USER
			dialog2.PeerId = md.UserId
			dao.GetUserDialogsDAO(dao.DB_MASTER).Insert(dialog2)
		}

		// 插入消息
		message := &dataobject.MessagesDO{}
		message.UserId = md.UserId
		message.PeerType = base.PEER_USER
		message.PeerId = inputPeerUser.UserId
		message.RandomId = request.RandomId
		message.Message = request.Message
		message.Date2 = int32(time.Now().Unix())
		messageId := dao.GetMessagesDAO(dao.DB_MASTER).Insert(message)

		// 更新最后一条消息
		model.GetDialogModel().UpdateTopMessage(dialog1.Id, int32(messageId))
		model.GetDialogModel().UpdateTopMessage(dialog2.Id, int32(messageId))

		// inbox和outbox
		// 发件箱
		messageBox := &dataobject.MessageBoxsDO{}
		messageBox.UserId = md.UserId
		messageBox.MessageBoxType = 0
		messageBox.MessageId = int32(messageId)
		outPts := model.GetSequenceModel().NextID(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(outPts)
		dao.GetMessageBoxsDAO(dao.DB_MASTER).Insert(messageBox)

		// 收件箱
		messageBox.UserId = inputPeerUser.UserId
		messageBox.MessageBoxType = 1
		inPts := model.GetSequenceModel().NextID(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(inPts)
		dao.GetMessageBoxsDAO(dao.DB_MASTER).Insert(messageBox)

		// 推送给sync
		// 推给客户端的updates
		update := mtproto.TLUpdateShortMessage{}
		update.Id = int32(messageId)
		update.UserId = md.UserId
		update.Pts = int32(inPts)
		update.PtsCount = 1
		update.Message = request.Message
		update.Date = message.Date2
		updateRawData := update.ToUpdates().Encode()

		delivery := &zproto.DeliveryUpdatesToUsers{}
		delivery.MyAuthKeyId = md.AuthId
		delivery.MySessionId = md.SessionId
		delivery.MyNetlibSessionId = md.NetlibSessionId
		delivery.SendtoUserIdList = append(delivery.SendtoUserIdList, md.UserId)
		delivery.SendtoUserIdList = append(delivery.SendtoUserIdList, inputPeerUser.UserId)
		delivery.RawData = updateRawData

		glog.Infof("MessagesSendMessage - delivery: %v", delivery)
		_, _ = s.SyncRPCClient.Client.DeliveryUpdates(context.Background(), delivery)
		// update.Encode()

		// 返回给客户端
		sentMessage := &mtproto.TLUpdateShortSentMessage{}
		sentMessage.Out = true
		sentMessage.Id = int32(messageId)
		sentMessage.Pts = int32(outPts)
		sentMessage.PtsCount = 1
		sentMessage.Date = int32(time.Now().Unix())

		glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		reply = sentMessage.ToUpdates()
	case *mtproto.InputPeer_InputPeerChannel:
	}

	return
}

func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesDeleteChatUser(ctx context.Context, request *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
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
	glog.Info("Process: %v", request)
	return nil, nil
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

	// 查出来
	// md := grpc_util.RpcMetadataFromIncoming(ctx)

/*
	// TODO(@benqi): check error!
	// authUsersDO, _ := s.AuthUsersDAO.SelectByAuthId(rpcMetaData.AuthId)
	// glog.Info("user_id: ", authUsersDO)
	// userDialogsDO, _ := s.UserDialogsDAO.SelectPinnedDialogs(authUsersDO.UserId)
	userDialogsDO, _ := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectPinnedDialogs(rpcMetaData.UserId)
	_ = userDialogsDO

	peerDialogs := &mtproto.TLMessagesPeerDialogs{}
	state := &mtproto.TLUpdatesState{}
	state.Date = int32(time.Now().Unix())

	peerDialogs.State = mtproto.MakeUpdates_State(state)

	reply := mtproto.MakeMessages_PeerDialogs(peerDialogs)
	glog.Infof("MessagesGetPinnedDialogs - reply: {%v}\n", reply)

	return reply, nil
 */
 	peerDialogs := &mtproto.TLMessagesPeerDialogs{}
	state := &mtproto.TLUpdatesState{}
	state.Date = int32(time.Now().Unix())
	state.Pts = 1
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
