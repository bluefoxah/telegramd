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

func (m *messageModel) getMessagesByIDList(idList []int32, order bool) (messages []*mtproto.Message) {
	// TODO(@benqi): Check messageDAO
	messageDAO := dao.GetMessagesDAO(dao.DB_SLAVE)

	var messagesDOList []dataobject.MessagesDO
	if order {
		messagesDOList = messageDAO.SelectByIdList(idList)

	} else {
		messagesDOList = messageDAO.SelectByIdList(idList)
	}
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
	return m.getMessagesByMessageBoxes(boxesList, true)
}

func (m *messageModel) GetMessagesByUserIdPeerOffsetLimit(userId int32, peerType , peerId int32, offset int32, limit int32) []*mtproto.Message {
	// 1. 先从message_boxes取出message_id
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByPeerOffsetLimit(userId, int8(peerType), peerId, offset, limit)
	glog.Infof("GetMessagesByUserIdPeerOffsetLimit - boxesList: %v", boxesList)
	if len(boxesList) == 0 {
		return make([]*mtproto.Message, 0)
	}
	return m.getMessagesByMessageBoxes(boxesList, false)
}

func (m *messageModel) getMessagesByMessageBoxes(boxes []dataobject.MessageBoxesDO, order bool) []*mtproto.Message {
	glog.Infof("getMessagesByMessageBoxes - boxes: {%v}", boxes)
	messageIdList := make([]int32, 0, len(boxes))
	for _, do := range boxes {
		messageIdList = append(messageIdList, do.MessageId)
	}
	messageList := m.getMessagesByIDList(messageIdList, order)
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
		return m.getMessagesByMessageBoxes(boxDOList, true)
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
