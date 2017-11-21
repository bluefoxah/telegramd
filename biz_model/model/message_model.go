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
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/biz_model/base"
	"sync"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/golang/glog"
	base2 "github.com/nebulaim/telegramd/base/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

type messageModel struct {
	// messageDAO *dao.MessagesDAO
}

var (
	messageInstance *messageModel
	messageInstanceOnce sync.Once
)

func GetMessageModel() *messageModel {
	messageInstanceOnce.Do(func() {
		messageInstance = &messageModel{}
	})
	return messageInstance
}

/*
// message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
message TL_message {
  bool out = 1;
  bool mentioned = 2;
  bool media_unread = 3;
  bool silent = 4;
  bool post = 5;
  int32 id = 6;
  int32 from_id = 7;
  Peer to_id = 8;
  MessageFwdHeader fwd_from = 9;
  int32 via_bot_id = 10;
  int32 reply_to_msg_id = 11;
  int32 date = 12;
  string message = 13;
  MessageMedia media = 14;
  ReplyMarkup reply_markup = 15;
  repeated MessageEntity entities = 16;
  int32 views = 17;
  int32 edit_date = 18;
  string post_author = 19;
}
 */

func messagesDOToMessage(do *dataobject.MessagesDO) (message *mtproto.TLMessage) {
	message = &mtproto.TLMessage{}
	message.Out = true
	message.Id = do.Id
	message.FromId = do.UserId
	switch do.PeerType {
	//case base.PEER_EMPTY:
	//	peer := &mtproto.TLInputPeerEmpty{}
	//	message.ToId = peer.ToPeer()
	case base.PEER_SELF, base.PEER_USER:
		peer := &mtproto.TLPeerUser{do.PeerId}
		message.ToId = peer.ToPeer()
	case base.PEER_CHAT:
		peer := &mtproto.TLPeerChat{do.PeerId}
		message.ToId = peer.ToPeer()
	case base.PEER_CHANNEL:
		peer := &mtproto.TLPeerChannel{do.PeerId}
		message.ToId = peer.ToPeer()
	}
	message.Date = do.Date2
	message.Message = do.Message
	return message
}

func (m *messageModel) GetMessagesByIDList(idList []int32) (messages []*mtproto.TLMessage) {
	// TODO(@benqi): Check messageDAO
	messageDAO := dao.GetMessagesDAO(dao.DB_SLAVE)

	messagesDOList := messageDAO.SelectByIdList(idList)
	messages = []*mtproto.TLMessage{}

	for _, messageDO := range messagesDOList {
		message := &mtproto.TLMessage{}
		message.Out = true
		message.Id = messageDO.Id
		message.FromId = messageDO.UserId
		switch messageDO.PeerType {
		case base.PEER_EMPTY:
			continue
		case base.PEER_SELF, base.PEER_USER:
			peer := &mtproto.TLPeerUser{messageDO.PeerId}
			message.ToId = peer.ToPeer()
		case base.PEER_CHAT:
			peer := &mtproto.TLPeerChat{messageDO.PeerId}
			message.ToId = peer.ToPeer()
		case base.PEER_CHANNEL:
			peer := &mtproto.TLPeerChannel{messageDO.PeerId}
			message.ToId = peer.ToPeer()
		}
		message.Date = messageDO.Date2
		message.Message = messageDO.Message

		messages = append(messages, message)
	}

	glog.Infof("SelectByIdList(%s) - {%v}", base2.JoinInt32List(idList, ","), messages)
	return
}

func (m *messageModel) GetMessagesByUserIdPeerOffsetLimit(userId int32, peerType , peerId int32, offset int32, limit int32) (messages []*mtproto.TLMessage) {
	// TODO(@benqi): Check messageDAO
	messageDAO := dao.GetMessagesDAO(dao.DB_SLAVE)

	var maxId = offset
	//if maxId < limit {
	//	maxId = 0
	//} else {
	//	maxId = maxId - limit
	//}

	messagesDOList := messageDAO.SelectByUserIdAndPeerOffsetLimit(maxId, peerType, userId, peerId, limit)
	messages = []*mtproto.TLMessage{}

	for _, messageDO := range messagesDOList {
		message := &mtproto.TLMessage{}
		message.Id = messageDO.Id
		message.FromId = messageDO.UserId
		switch messageDO.PeerType {
		case base.PEER_EMPTY:
			continue
		case base.PEER_SELF, base.PEER_USER:
			if messageDO.UserId == userId {
				message.Out = true
			} else {
				message.Out = false
			}
			peer := &mtproto.TLPeerUser{messageDO.PeerId}
			message.ToId = peer.ToPeer()
		case base.PEER_CHAT:
			peer := &mtproto.TLPeerChat{messageDO.PeerId}
			message.ToId = peer.ToPeer()
		case base.PEER_CHANNEL:
			peer := &mtproto.TLPeerChannel{messageDO.PeerId}
			message.ToId = peer.ToPeer()
		}
		message.Date = messageDO.Date2
		message.Message = messageDO.Message

		messages = append(messages, message)
	}
	return
}

func (m *messageModel) GetMessagesByPts(userId int32, pts int32) (messages []*mtproto.TLMessage) {
	messageDAO := dao.GetMessagesDAO(dao.DB_SLAVE)

	messagesDOList := messageDAO.SelectByPts(userId, pts)
	messages = []*mtproto.TLMessage{}
	for _, messageDO := range messagesDOList {
		message := messagesDOToMessage(&messageDO)
		messages = append(messages, message)
	}

	return messages
}
