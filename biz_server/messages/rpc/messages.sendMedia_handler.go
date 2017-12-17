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
	"time"
	"github.com/nebulaim/telegramd/biz_model/model"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_server/delivery"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/frontend/id"
)

// messages.sendMedia#c8f16791 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia random_id:long reply_markup:flags.2?ReplyMarkup = Updates;
func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesSendMedia - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	now := int32(time.Now().Unix())
	peer := base.FromInputPeer(request.GetPeer())
	message := mtproto.NewTLMessage()
	message.SetSilent(request.GetSilent())

	// TODO(@benqi): ???
	// request.Background
	// request.NoWebpage
	// request.ClearDraft

	message.SetFromId(md.UserId)
	switch request.GetPeer().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputPeerSelf:
		to := mtproto.NewTLPeerUser()
		to.SetUserId(md.UserId)
		message.SetToId(to.To_Peer())
	case mtproto.TLConstructor_CRC32_inputPeerUser:
		to := mtproto.NewTLPeerUser()
		to.SetUserId(request.GetPeer().GetData2().GetUserId())
		message.SetToId(to.To_Peer())
	case mtproto.TLConstructor_CRC32_inputPeerChat:
		to := mtproto.NewTLPeerChat()
		to.SetChatId(request.GetPeer().GetData2().GetChatId())
		message.SetToId(to.To_Peer())
	case mtproto.TLConstructor_CRC32_inputPeerChannel:
		to := mtproto.NewTLPeerChannel()
		to.SetChannelId(request.GetPeer().GetData2().GetChannelId())
		message.SetToId(to.To_Peer())
	default:
		// mtproto.TLConstructor_CRC32_inputPeerEmpty
		// TODO(@benqi): Bad request
	}

	message.SetReplyToMsgId(request.GetReplyToMsgId())
	message.SetReplyMarkup(request.GetReplyMarkup())
	message.SetDate(now)

	mediaData := request.GetMedia().GetData2()
	switch request.Media.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputMediaUploadedPhoto:
		fileData := mediaData.GetFile().GetData2()

		media := mtproto.NewTLMessageMediaPhoto()
		photo := mtproto.NewTLPhoto()

		// get access_hash
		filesDO := dao.GetFilesDAO(dao.DB_MASTER).SelectByIDAndParts(fileData.GetId(), fileData.GetParts())
		photo.SetAccessHash(filesDO.AccessHash)

		// @benqi
		//   重要，客户端发送图片文件后，要求返回的id不能和上传文件相同，如果相同认为文件上传未结束
		photoId := id.NextId()
		photo.SetDate(now)
		photo.SetId(photoId)
		// photo.SetId(fileData.GetId())
		sizes, err := model.GetPhotoModel().UploadPhoto(md.UserId, photoId, fileData.GetId(), fileData.GetParts(), fileData.GetName(), fileData.GetMd5Checksum())
		if err != nil {
			glog.Errorf("UploadPhoto error: %v, by %s", err, logger.JsonDebugData(media))
		} else {
			photo.SetSizes(sizes)
		}

		media.SetCaption(mediaData.GetCaption())
		media.SetTtlSeconds(mediaData.GetTtlSeconds())
		media.SetPhoto(photo.To_Photo())

		// glog.Infof("inputMediaUploadedPhoto: %s", logger.JsonDebugData(media))

		message.SetMedia(media.To_MessageMedia())
	}

	sentMessage := mtproto.NewTLUpdateShortSentMessage()
	switch request.GetPeer().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputPeerSelf:
		// 1. SaveMessage
		messageId := model.GetMessageModel().CreateHistoryMessage2(md.UserId, peer, request.RandomId, now, message.To_Message())
		// 2. MessageBoxes
		pts := model.GetMessageModel().CreateMessageBoxes(md.UserId, message.GetFromId(), base.PEER_SELF, md.UserId, false, messageId)
		// 3. dialog
		model.GetDialogModel().CreateOrUpdateByLastMessage(md.UserId, base.PEER_SELF, md.UserId, messageId, message.GetMentioned())



		// 推送给sync
		// 推给客户端的updates
		updates := mtproto.NewTLUpdateShortMessage()
		updates.SetId(int32(messageId))
		updates.SetUserId(md.UserId)
		updates.SetPts(pts)
		updates.SetPtsCount(1)
		// updates.Message = request.Message

		updates.SetDate(now)
		delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
			md.AuthId,
			md.SessionId,
			md.NetlibSessionId,
			[]int32{md.UserId},
			updates.To_Updates().Encode())
		// 返回给客户端
		// sentMessage := &mtproto.TLUpdateShortSentMessage{}

		sentMessage.SetOut(true)
		sentMessage.SetId(int32(messageId))
		sentMessage.SetPts(pts)
		sentMessage.SetPtsCount(1)
		sentMessage.SetDate(int32(time.Now().Unix()))
		sentMessage.SetMedia(mtproto.NewTLMessageMediaEmpty().To_MessageMedia())

		// glog.Infof("MessagesSendMessage - reply: %s", logger.JsonDebugData(sentMessage))
		// reply = sentMessage.ToUpdates()
	case mtproto.TLConstructor_CRC32_inputPeerUser:
		// peer := request.GetPeer().To_InputPeerUser()
		// 1. SaveMessage
		messageId := model.GetMessageModel().CreateHistoryMessage2(md.UserId, peer, request.RandomId, now, message.To_Message())
		message.SetId(messageId)

		// 2. MessageBoxes
		outPts := model.GetMessageModel().CreateMessageBoxes(md.UserId, message.GetFromId(), base.PEER_USER, peer.PeerId, false, messageId)
		inPts := model.GetMessageModel().CreateMessageBoxes(peer.PeerId, message.GetFromId(), base.PEER_USER, md.UserId, true, messageId)
		// 3. dialog
		model.GetDialogModel().CreateOrUpdateByLastMessage(md.UserId, base.PEER_USER, peer.PeerId, messageId, message.GetMentioned())
		model.GetDialogModel().CreateOrUpdateByLastMessage(peer.PeerId, base.PEER_USER, md.UserId, messageId, message.GetMentioned())

		users := model.GetUserModel().GetUserList([]int32{md.UserId, peer.PeerId})

		// sender
		// 推送给sync
		// 推给客户端的updates
		updates := mtproto.NewTLUpdates()

		updates.SetSeq(0)
		updates.SetDate(now)

		updateNewMessage := mtproto.NewTLUpdateNewMessage()
		updateNewMessage.SetMessage(message.To_Message())
		updateNewMessage.SetPts(outPts)
		updateNewMessage.SetPtsCount(1)

		updates.Data2.Updates = append(updates.Data2.Updates, updateNewMessage.To_Update())

		message.SetOut(false)

		for _, u := range  users {
			if u.GetId() != md.UserId {
				u.SetSelf(true)
			} else {
				u.SetSelf(false)
			}
			u.SetContact(true)
			updates.Data2.Users = append(updates.Data2.Users, u.To_User())
			// .SetUsers(users)
		}

		delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
			md.AuthId,
			md.SessionId,
			md.NetlibSessionId,
			// []int32{md.UserId, peer.PeerId},
			[]int32{peer.PeerId},
			updates.To_Updates().Encode())

		//	updates := &mtproto.TLUpdates{}

		updateMessageID := mtproto.NewTLUpdateMessageID()
		updateMessageID.SetId(int32(messageId))
		updateMessageID.SetRandomId(request.GetRandomId())
		// updates.Data2.Updates = append(updates.Data2.Updates, updateMessageID.To_Update())

		updatesList := []*mtproto.Update{updateMessageID.To_Update()}
		updatesList = append(updatesList, updates.GetUpdates()...)
		updates.SetUpdates(updatesList)

		// fix message data
		// message.SetOut(false)
		updateNewMessage.SetPts(inPts)

		updates.SetUsers([]*mtproto.User{})
		for _, u := range  users {
			if u.GetId() == md.UserId {
				u.SetSelf(true)
				message.SetOut(true)
			} else {
				u.SetSelf(false)
				message.SetOut(false)
			}
			u.SetContact(true)
			updates.Data2.Users = append(updates.Data2.Users, u.To_User())
		}

		message.SetOut(true)

		glog.Infof("MessagesSendMedia - reply: %s", logger.JsonDebugData(updates))
		return updates.To_Updates(), nil

		// updates.SetId(int32(messageId))
		// updates.SetUserId(md.UserId)
		// updates.SetPts(inPts)
		// updates.SetPtsCount(1)
		// updates.Message = request.Message
		// updates.SetDate(now)
		// 返回给客户端
		// sentMessage := &mtproto.TLUpdateShortSentMessage{}
		// sentMessage.SetOut(true)
		// sentMessage.SetId(int32(messageId))
		// sentMessage.SetPts(outPts)
		// sentMessage.SetPtsCount(1)
		// sentMessage.SetDate(now)
		// glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		// reply = sentMessage.ToUpdates()
	case mtproto.TLConstructor_CRC32_inputPeerChat:
		// 返回给客户端
		// sentMessage := &mtproto.TLUpdateShortSentMessage{}
		sentMessage.SetOut(true)
		// sentMessage.Id = int32(messageId)
		// sentMessage.Pts = outPts
		sentMessage.SetPtsCount(1)
		sentMessage.SetDate(now)
		// 1. SaveMessage
		messageId := model.GetMessageModel().CreateHistoryMessage2(md.UserId, peer, request.RandomId, now, message.To_Message())
		participants := model.GetChatModel().GetChatParticipants(peer.PeerId)
		var userId int32 = 0
		for _, participan := range participants.GetParticipants() {
			switch participan.GetConstructor() {
			case mtproto.TLConstructor_CRC32_chatParticipantCreator:
				userId = participan.GetData2().GetUserId()
			case mtproto.TLConstructor_CRC32_chatParticipantAdmin:
				userId = participan.GetData2().GetUserId()
			case mtproto.TLConstructor_CRC32_chatParticipant:
				userId = participan.GetData2().GetUserId()
			}
			// 2. MessageBoxes
			outgoing := userId == md.UserId
			pts := model.GetMessageModel().CreateMessageBoxes(userId, md.UserId, peer.PeerType, peer.PeerId, outgoing, messageId)
			model.GetDialogModel().CreateOrUpdateByLastMessage(userId, peer.PeerType, peer.PeerId, messageId, message.GetMentioned())
			// inPts := model.GetMessageModel().CreateMessageBoxes(peer.PeerId, message.FromId, peer, true, messageId)
			// 3. dialog
			// model.GetDialogModel().CreateOrUpdateByLastMessage(peer.PeerId, peer, messageId, message.Mentioned)
			// 推送给sync
			// 推给客户端的updates
			updates := mtproto.NewTLUpdateShortChatMessage()
			updates.SetId(int32(messageId))
			updates.SetFromId(md.UserId)
			updates.SetChatId(peer.PeerId)
			updates.SetPts(pts)
			updates.SetPtsCount(1)
			// updates.Message = request.Message
			updates.SetDate(now)
			if md.UserId == userId {
				sentMessage.SetId(int32(messageId))
				sentMessage.SetPts(pts)
				delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
					md.AuthId,
					md.SessionId,
					md.NetlibSessionId,
					[]int32{userId},
					updates.To_Updates().Encode())
			} else {
				delivery.GetDeliveryInstance().DeliveryUpdates(
					md.AuthId,
					md.SessionId,
					md.NetlibSessionId,
					[]int32{userId},
					updates.To_Updates().Encode())
			}
		}
		glog.Infof("MessagesSendMessage - reply: %v", sentMessage)
		// reply = sentMessage.ToUpdates()
	case mtproto.TLConstructor_CRC32_inputPeerChannel:
	default:
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "InputPeer invalid"))
	}
	glog.Infof("MessagesSendMessage - reply: %s", logger.JsonDebugData(sentMessage))
	return sentMessage.To_Updates(), nil
}
