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
	"time"
	"github.com/golang/protobuf/proto"
)

const (
	MESSAGE_TYPE_UNKNOWN = 0
	MESSAGE_TYPE_MESSAGE = 1
	MESSAGE_TYPE_MESSAGE_SERVICE = 2
)
const (
	MESSAGE_BOX_TYPE_INCOMING = 0
	MESSAGE_BOX_TYPE_OUTGOING = 1
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

/*
func toMessage(do *dataobject.MessagesDO) (message *mtproto.TLMessage) {
	message = &mtproto.TLMessage{}
	err := proto.Unmarshal(do.MessageData, message)
	if err != nil {
		return nil
	}

	//message.Out = true
	//message.Id = do.Id
	//message.FromId = do.UserId
	//switch do.PeerType {
	////case base.PEER_EMPTY:
	////	peer := &mtproto.TLInputPeerEmpty{}
	////	message.ToId = peer.ToPeer()
	//case base.PEER_SELF, base.PEER_USER:
	//	peer := &mtproto.TLPeerUser{do.PeerId}
	//	message.ToId = peer.ToPeer()
	//case base.PEER_CHAT:
	//	peer := &mtproto.TLPeerChat{do.PeerId}
	//	message.ToId = peer.ToPeer()
	//case base.PEER_CHANNEL:
	//	peer := &mtproto.TLPeerChannel{do.PeerId}
	//	message.ToId = peer.ToPeer()
	//}
	// message.Date = do.Date2
	// message.Message = do.Message
	return message
}

func toMessageService(do *dataobject.MessagesDO) (message *mtproto.TLMessageService) {
	message = &mtproto.TLMessageService{}
	err := proto.Unmarshal(do.MessageData, message)
	if err != nil {
		return nil
	}

	//message.Out = true
	//message.Id = do.Id
	//message.FromId = do.UserId
	//switch do.PeerType {
	////case base.PEER_EMPTY:
	////	peer := &mtproto.TLInputPeerEmpty{}
	////	message.ToId = peer.ToPeer()
	//case base.PEER_SELF, base.PEER_USER:
	//	peer := &mtproto.TLPeerUser{do.PeerId}
	//	message.ToId = peer.ToPeer()
	//case base.PEER_CHAT:
	//	peer := &mtproto.TLPeerChat{do.PeerId}
	//	message.ToId = peer.ToPeer()
	//case base.PEER_CHANNEL:
	//	peer := &mtproto.TLPeerChannel{do.PeerId}
	//	message.ToId = peer.ToPeer()
	//}
	//message.Date = do.Date2
	// message.Message = do.Message
	return message
}
*/

func (m *messageModel) getMessagesByIDList(idList []int32) (messages []*mtproto.Message) {
	// TODO(@benqi): Check messageDAO
	messageDAO := dao.GetMessagesDAO(dao.DB_SLAVE)

	messagesDOList := messageDAO.SelectByIdList(idList)
	// messages = []*mtproto.TLMessage{}

	for _, messageDO := range messagesDOList {
		if messageDO.MessageType == MESSAGE_TYPE_MESSAGE {
			message := &mtproto.TLMessage{}
			err := proto.Unmarshal(messageDO.MessageData, message)
			if err != nil {
				glog.Errorf("GetMessagesByIDList - Unmarshal message(%d)error: %v", messageDO.Id, err)
				continue
			}
			message.Id = messageDO.Id
			messages = append(messages, message.ToMessage())
		} else if messageDO.MessageType == MESSAGE_TYPE_MESSAGE_SERVICE {
			message := &mtproto.TLMessageService{}
			err := proto.Unmarshal(messageDO.MessageData, message)
			if err != nil {
				glog.Errorf("GetMessagesByIDList - Unmarshal message(%d)error: %v", messageDO.Id, err)
				continue
			}
			message.Id = messageDO.Id
			messages = append(messages, message.ToMessage())
		}
	}
	glog.Infof("GetMessagesByIDList(%s) - {%v}", base2.JoinInt32List(idList, ","), messages)
	return
}

func (m *messageModel) GetMessagesByPeerAndMessageIdList(userId int32, idList []int32) (messages []*mtproto.Message) {
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByMessageIdList(userId, idList)
	return m.getMessagesByMessageBoxes(boxesList)
}

func (m *messageModel) GetMessagesByUserIdPeerOffsetLimit(userId int32, peerType , peerId int32, offset int32, limit int32) (messages []*mtproto.Message) {
	// 1. 先从message_boxes取出message_id
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByPeerOffsetLimit(userId, int8(peerType), peerId, offset, limit)
	glog.Infof("GetMessagesByUserIdPeerOffsetLimit - boxesList: %v", boxesList)
	if len(boxesList) == 0 {
		return make([]*mtproto.Message, 0)
	}
	return m.getMessagesByMessageBoxes(boxesList)
/*
	// TODO(@benqi): Check messageDAO
	messageDAO := dao.GetMessagesDAO(dao.DB_SLAVE)

	var maxId = offset
	//if maxId < limit {
	//	maxId = 0
	//} else {
	//	maxId = maxId - limit
	//}

	var messagesDOList []dataobject.MessagesDO
	switch peerType {
	case base.PEER_SELF, base.PEER_USER:
		messagesDOList = messageDAO.SelectByUserIdAndPeerOffsetLimit(maxId, peerType, userId, peerId, limit)
	case base.PEER_CHAT:
		messagesDOList = messageDAO.SelectChatMessageByOffsetLimit(maxId, peerType, peerId, limit)
		glog.Infof("GetMessagesByUserIdPeerOffsetLimit - {%v}", messagesDOList)
	case base.PEER_CHANNEL:
	}

	messages = []*mtproto.TLMessage{}

	for _, messageDO := range messagesDOList {
		message := &mtproto.TLMessage{}
		message.Id = messageDO.Id
		message.FromId = messageDO.UserId
		if messageDO.UserId == userId {
			message.Out = true
		} else {
			message.Out = false
		}
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
 */
	return
}

func (m *messageModel) getMessagesByMessageBoxes(boxes []dataobject.MessageBoxesDO) []*mtproto.Message {
	glog.Infof("getMessagesByMessageBoxes - boxes: {%v}", boxes)
	messageIdList := make([]int32, 0, len(boxes))
	for _, do := range boxes {
		messageIdList = append(messageIdList, do.MessageId)
	}
	messageList := m.getMessagesByIDList(messageIdList)
	// TODO(@benqi): 假设数据一致
	for i, message := range messageList {
		boxDO := boxes[i]
		switch message.Payload.(type) {
		case *mtproto.Message_Message:
			m2 := message.GetMessage()
			if boxDO.MessageBoxType == MESSAGE_BOX_TYPE_OUTGOING {
				m2.Out = true
			} else {
				m2.Out = false
			}
			m2.Silent = true
			m2.MediaUnread = boxDO.MediaUnread != 0
			m2.Mentioned = false
			glog.Infof("message(%d): %v", i, m2)
		case *mtproto.Message_MessageService:
			m2 := message.GetMessageService()
			if boxDO.MessageBoxType == MESSAGE_BOX_TYPE_OUTGOING {
				m2.Out = true
			} else {
				m2.Out = false
			}
			m2.Silent = true
			m2.MediaUnread = boxDO.MediaUnread != 0
			m2.Mentioned = false
			glog.Infof("message2(%d): %v", i, m2)
		}
		glog.Infof("message(%d): %v", i, messageList[i])
	}

	return messageList
}

func (m *messageModel) GetMessagesByGtPts(userId int32, pts int32) []*mtproto.Message {
	boxDOList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByGtPts(userId, pts)
	if len(boxDOList) == 0 {
		return make([]*mtproto.Message, 0)
	} else {
		return m.getMessagesByMessageBoxes(boxDOList)
	}
}

func (m *messageModel) GetLastPtsByUserId(userId int32) int32 {
	do := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectLastPts(userId)
	if do == nil {
		return 0
	} else {
		return  do.Pts
	}
}

// CreateMessage
func (m *messageModel) CreateMessageBoxes(userId, fromId int32, peerType int32, peerId int32, incoming bool, messageId int32) (int32) {
	messageBox := &dataobject.MessageBoxesDO{}
	if incoming {
		messageBox.UserId = userId
		messageBox.SenderUserId = fromId
		messageBox.MessageBoxType = MESSAGE_BOX_TYPE_INCOMING
		messageBox.PeerType = int8(peerType)
		messageBox.PeerId = peerId
		messageBox.MessageId = messageId
		outPts := GetSequenceModel().NextID(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(outPts)
		messageBox.Date2 = int32(time.Now().Unix())
		messageBox.CreatedAt = base2.NowFormatYMDHMS()
	} else {
		messageBox.UserId = userId
		messageBox.SenderUserId = fromId
		messageBox.MessageBoxType = MESSAGE_BOX_TYPE_OUTGOING
		messageBox.PeerType = int8(peerType)
		messageBox.PeerId = peerId
		messageBox.MessageId = messageId
		inPts := GetSequenceModel().NextID(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(inPts)
		messageBox.Date2 = int32(time.Now().Unix())
		messageBox.CreatedAt = base2.NowFormatYMDHMS()
	}
	dao.GetMessageBoxesDAO(dao.DB_MASTER).Insert(messageBox)
	return messageBox.Pts
}

// CreateMessage
func (m *messageModel) CreateHistoryMessage(fromId int32, peer *base.PeerUtil, randomId int64, message *mtproto.TLMessage) (messageId int32) {
	// TODO(@benqi): 重复插入出错处理
	messageDO := &dataobject.MessagesDO{}
	messageDO.SenderUserId = fromId
	messageDO.PeerType = peer.PeerType
	messageDO.PeerId = peer.PeerId
	// Message
	messageDO.MessageType = MESSAGE_TYPE_MESSAGE
	messageDO.RandomId = randomId
	messageDO.Date2 = message.Date
	messageDO.MessageData, _ = proto.Marshal(message)
	messageId = int32(dao.GetMessagesDAO(dao.DB_MASTER).Insert(messageDO))
	return
}

// CreateMessage
func (m *messageModel) CreateHistoryMessageService(fromId int32, peer *base.PeerUtil, randomId int64, message *mtproto.TLMessageService) (messageId int32) {
	// TODO(@benqi): 重复插入出错处理
	messageDO := &dataobject.MessagesDO{}
	messageDO.SenderUserId = fromId
	messageDO.PeerType = peer.PeerType
	messageDO.PeerId = peer.PeerId
	messageDO.MessageType = MESSAGE_TYPE_MESSAGE_SERVICE
	messageDO.RandomId = randomId
	messageDO.Date2 = int32(time.Now().Unix())
	messageDO.MessageData, _ = proto.Marshal(message)
	messageId = int32(dao.GetMessagesDAO(dao.DB_MASTER).Insert(messageDO))
	return
}

//func (m *messageModel) MakeUpdatesByMessage(randomId int64, message *mtproto.TLMessageService) (updates *mtproto.Updates) {
//	//// 插入消息
//	peer := base.FromPeer(message.ToId)
//	messageDO := &dataobject.MessagesDO{}
//
//	messageDO.UserId = message.FromId
//	messageDO.PeerType = peer.PeerType
//	messageDO.PeerId = peer.PeerId
//	messageDO.RandomId = randomId
//	messageDO.Date2 = message.Date
//
//	messageId := dao.GetMessagesDAO(dao.DB_MASTER).Insert(messageDO)
//	message.Id = int32(messageId)
//	return
//}
