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
	"sync"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/golang/glog"
	base2 "github.com/nebulaim/telegramd/base/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"time"
	"github.com/nebulaim/telegramd/base/logger"
	"fmt"
	"encoding/json"
)

const (
	MESSAGE_TYPE_UNKNOWN = 0
	MESSAGE_TYPE_MESSAGE_EMPTY = 1
	MESSAGE_TYPE_MESSAGE = 2
	MESSAGE_TYPE_MESSAGE_SERVICE = 3
)
const (
	MESSAGE_BOX_TYPE_INCOMING = 0
	MESSAGE_BOX_TYPE_OUTGOING = 1
)

const (
	PTS_UNKNOWN = 0
	PTS_MESSAGE_OUTBOX = 1
	PTS_MESSAGE_INBOX = 2
	PTS_READ_HISTORY_OUTBOX = 3
	PTS_READ_HISTORY_INBOX = 4
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
		// TODO(@benqi): 不推给DB，程序内排序
		messagesDOList = messageDAO.SelectOrderByIdList(idList)
	} else {
		messagesDOList = messageDAO.SelectByIdList(idList)
	}

	messages = []*mtproto.Message{}
	for _, messageDO := range messagesDOList {
		message := &mtproto.Message{
			Data2: &mtproto.Message_Data{},
		}
		switch messageDO.MessageType {
		case MESSAGE_TYPE_MESSAGE_EMPTY:
			message.Constructor = mtproto.TLConstructor_CRC32_messageEmpty
		case MESSAGE_TYPE_MESSAGE:
			message.Constructor = mtproto.TLConstructor_CRC32_message
			// err := proto.Unmarshal(messageDO.MessageData, message)
			err := json.Unmarshal([]byte(messageDO.MessageData), message)
			if err != nil {
				glog.Errorf("GetMessagesByIDList - Unmarshal message(%d)error: %v", messageDO.Id, err)
				continue
			}
			message.Data2.Id = messageDO.Id
			//messages = append(messages, message.To_Message())
		case MESSAGE_TYPE_MESSAGE_SERVICE:
			message.Constructor = mtproto.TLConstructor_CRC32_messageService
			// err := proto.Unmarshal(messageDO.MessageData, message)
			err := json.Unmarshal([]byte(messageDO.MessageData), message)
			if err != nil {
				glog.Errorf("GetMessagesByIDList - Unmarshal message(%d)error: %v", messageDO.Id, err)
				continue
			}
			message.Data2.Id = messageDO.Id
		default:
			glog.Error("Invalid messageType, db's data error: %s", logger.JsonDebugData(messageDO))
			continue
		}

		messages = append(messages, message)
	}
	glog.Infof("GetMessagesByIDList(%s) - %s", base2.JoinInt32List(idList, ","), logger.JsonDebugData(messages))
	return
}

func (m *messageModel) GetMessagesByPeerAndMessageIdList(userId int32, idList []int32) (messages []*mtproto.Message) {
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByMessageIdList(userId, idList)
	return m.getMessagesByMessageBoxes(boxesList, true)
}

func (m *messageModel) GetMessagesByPeerAndMessageIdList2(userId int32, idList []int32) (messages []*mtproto.Message) {
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByMessageIdList(userId, idList)
	return m.getMessagesByMessageBoxes(boxesList, false)
}

func (m *messageModel) LoadBackwardHistoryMessages(userId int32, peerType , peerId int32, offset int32, limit int32) []*mtproto.Message {
	// 1. 先从message_boxes取出message_id
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectBackwardByPeerOffsetLimit(userId, int8(peerType), peerId, offset, limit)
	glog.Infof("GetMessagesByUserIdPeerOffsetLimit - boxesList: %v", boxesList)
	if len(boxesList) == 0 {
		return make([]*mtproto.Message, 0)
	}
	return m.getMessagesByMessageBoxes(boxesList, true)
}

func (m *messageModel) LoadForwardHistoryMessages(userId int32, peerType , peerId int32, offset int32, limit int32) []*mtproto.Message {
	// 1. 先从message_boxes取出message_id
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectForwardByPeerOffsetLimit(userId, int8(peerType), peerId, offset, limit)
	glog.Infof("GetMessagesByUserIdPeerOffsetLimit - boxesList: %v", boxesList)
	if len(boxesList) == 0 {
		return make([]*mtproto.Message, 0)
	}
	return m.getMessagesByMessageBoxes(boxesList, true)
}

// TODO(@benqi): 出问题了！！！
func (m *messageModel) getMessagesByMessageBoxes(boxes []dataobject.MessageBoxesDO, order bool) []*mtproto.Message {
	glog.Infof("getMessagesByMessageBoxes - boxes: %s", logger.JsonDebugData(boxes))
	messageIdList := make([]int32, 0, len(boxes))
	for _, do := range boxes {
		messageIdList = append(messageIdList, do.MessageId)
	}
	messageList := m.getMessagesByIDList(messageIdList, order)
	// TODO(@benqi): 假设数据一致，后续还是要考虑数据不一致情况
	for i, message := range messageList {
		// TODO(@benqi): 数据不一致会有问题
		boxDO := boxes[i]
		if boxDO.MessageBoxType == MESSAGE_BOX_TYPE_OUTGOING {
			message.Data2.Out = true
		} else {
			message.Data2.Out = false
		}
		// message.Data2.Silent = true
		message.Data2.MediaUnread = boxDO.MediaUnread != 0
		message.Data2.Mentioned = false

		// 使用UserMessageBoxId作为messageBoxId
		message.Data2.Id = boxDO.UserMessageBoxId
		// glog.Infof("message(%d): %s", i, logger.JsonDebugData(message))
	}

	return messageList
}

func (m *messageModel) GetMessagesByGtPts(userId int32, pts int32) []*mtproto.Message {
	boxDOList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectByGtPts(userId, pts)
	if len(boxDOList) == 0 {
		return make([]*mtproto.Message, 0)
	} else {
		return m.getMessagesByMessageBoxes(boxDOList, false)
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
	messageBox := &dataobject.MessageBoxesDO{
		UserId:       userId,
		SenderUserId: fromId,
		PeerType:     int8(peerType),
		PeerId:       peerId,
		MessageId:    messageId,
		Date2:        int32(time.Now().Unix()),
		CreatedAt:    base2.NowFormatYMDHMS(),
	}

	if incoming {
		messageBox.MessageBoxType = MESSAGE_BOX_TYPE_INCOMING
		outPts := GetSequenceModel().NextPtsId(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(outPts)
	} else {
		messageBox.MessageBoxType = MESSAGE_BOX_TYPE_OUTGOING
		inPts := GetSequenceModel().NextPtsId(base2.Int32ToString(messageBox.UserId))
		messageBox.Pts = int32(inPts)
	}

	dao.GetMessageBoxesDAO(dao.DB_MASTER).Insert(messageBox)
	return messageBox.Pts
}

// CreateHistoryMessage2
func (m *messageModel) CreateHistoryMessage2(fromId int32, peer *base.PeerUtil, randomId int64, date int32, message *mtproto.Message) (messageId int32) {
	// TODO(@benqi): 重复插入出错处理
	messageDO := &dataobject.MessagesDO{
		SenderUserId: fromId,
		PeerType:     peer.PeerType,
		PeerId:       peer.PeerId,
		RandomId:     randomId,
		Date2:        date,
	}

	switch message.GetConstructor() {
	case mtproto.TLConstructor_CRC32_messageEmpty:
		messageDO.MessageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.TLConstructor_CRC32_message:
		messageDO.MessageType = MESSAGE_TYPE_MESSAGE
	case mtproto.TLConstructor_CRC32_messageService:
		messageDO.MessageType = MESSAGE_TYPE_MESSAGE_SERVICE
	default:
		panic(fmt.Errorf("Invalid message_type: {%v}", message))
	}

	// TODO(@benqi): 测试阶段使用Json!!!
	// messageDO.MessageData, _ = proto.Marshal(message)
	messageData, _ := json.Marshal(message)
	messageDO.MessageData = string(messageData)
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

type IDMessage struct {
	UserId int32
	MessageBoxId int32
}

// SendMessage
func (m *messageModel) SendMessage(fromId int32, peerType int32, peerId int32, clientRandomId int64, message *mtproto.Message) (ids []*IDMessage) {
	switch peerType {
	case base.PEER_USER:
		ids = m.sendUserMessage(fromId, peerId, clientRandomId, message)
	case base.PEER_CHAT:
		ids = m.sendChatMessage(fromId, peerId, clientRandomId, message)
	case base.PEER_CHANNEL:
		ids = m.sendChannelMessage(fromId, peerId, clientRandomId, message)
	default:
		glog.Errorf("SendMessage - invalid peerType: %d", peerType)
	}

	return
/*
	// TODO(@benqi): 重复插入出错处理
	messageDO := &dataobject.MessagesDO{
		SenderUserId: fromId,
		PeerType:     peer.PeerType,
		PeerId:       peer.PeerId,
		RandomId:     randomId,
		Date2:        date,
	}

	// TODO(@benqi): 测试阶段使用Json!!!
	// messageDO.MessageData, _ = proto.Marshal(message)
	messageData, _ := json.Marshal(message)
	messageDO.MessageData = string(messageData)
	messageId = int32(dao.GetMessagesDAO(dao.DB_MASTER).Insert(messageDO))
 */
	// Insert
}

func (m *messageModel) insertMessage(fromId, peerType, peerId int32, clientRandomId int64, message *mtproto.Message) int32 {
	messageDO := &dataobject.MessagesDO{
		// UserId:       userId,
		SenderUserId: fromId,
		PeerType:     base.PEER_USER,
		PeerId:       peerId,
		RandomId:     clientRandomId,
		Date2:        message.GetData2().GetDate(),
	}

	switch message.GetConstructor() {
	case mtproto.TLConstructor_CRC32_messageEmpty:
		messageDO.MessageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.TLConstructor_CRC32_message:
		messageDO.MessageType = MESSAGE_TYPE_MESSAGE
	case mtproto.TLConstructor_CRC32_messageService:
		messageDO.MessageType = MESSAGE_TYPE_MESSAGE_SERVICE
	default:
		panic(fmt.Errorf("Invalid message_type: {%v}", message))
	}

	// TODO(@benqi): 测试阶段使用Json!!!
	// messageDO.MessageData, _ = proto.Marshal(message)
	messageData, _ := json.Marshal(message)
	messageDO.MessageData = string(messageData)

	return int32(dao.GetMessagesDAO(dao.DB_MASTER).Insert(messageDO))
}

// CreateMessage
func (m *messageModel) insertMessageBox(userId, fromId, peerType, peerId int32, messageBoxType int8, messageId int32) (int32) {
	messageBox := &dataobject.MessageBoxesDO{
		UserId:       userId,
		UserMessageBoxId: int32(GetSequenceModel().NextMessageBoxId(base2.Int32ToString(userId))),
		SenderUserId: fromId,
		MessageBoxType: messageBoxType,
		PeerType:     int8(peerType),
		PeerId:       peerId,
		MessageId:    messageId,
		Date2:        int32(time.Now().Unix()),
		CreatedAt:    base2.NowFormatYMDHMS(),
	}

	// TODO(@benqi): check UserMessageBoxId
	if messageBox.UserMessageBoxId == 0 {
		glog.Errorf("insertMessageBox - error: NextMessageBoxId is 0")
	} else {
		dao.GetMessageBoxesDAO(dao.DB_MASTER).Insert(messageBox)
	}

	return messageBox.UserMessageBoxId
}

func (m *messageModel) sendUserMessage(fromId, peerId int32, clientRandomId int64, message *mtproto.Message) (ids []*IDMessage) {
	var boxId int32

	// 存历史消息
	messageId := m.insertMessage(fromId, base.PEER_USER, peerId, clientRandomId, message)
	if fromId == peerId {
		// PeerSelf

		// 存message_box
		boxId = m.insertMessageBox(fromId, fromId, base.PEER_USER, peerId, MESSAGE_BOX_TYPE_INCOMING, messageId)
		// dialog
		_  = GetDialogModel().CreateOrUpdateByLastMessage(fromId, base.PEER_USER, peerId, boxId, message.GetData2().GetMentioned(), true)
		ids = append(ids, &IDMessage{UserId: fromId, MessageBoxId: boxId})
	} else {
		// PeerUser

		// outbox
		// 存历史消息
		// messageId = m.insertMessage(fromId, base.PEER_USER, peerId, clientRandomId, message)
		// 存message_box
		boxId = m.insertMessageBox(fromId, fromId, base.PEER_USER, peerId, MESSAGE_BOX_TYPE_OUTGOING, messageId)
		// dialog
		_  = GetDialogModel().CreateOrUpdateByLastMessage(fromId, base.PEER_USER, peerId, boxId, message.GetData2().GetMentioned(), false)
		ids = append(ids, &IDMessage{UserId: fromId, MessageBoxId: boxId})

		// inbox
		// 存历史消息
		// messageId = m.insertMessage(peerId, fromId, base.PEER_USER, peerId, clientRandomId, message)
		// 存message_box
		boxId = m.insertMessageBox(peerId, fromId, base.PEER_USER, peerId, MESSAGE_BOX_TYPE_INCOMING, messageId)
		// dialog
		_  = GetDialogModel().CreateOrUpdateByLastMessage(peerId, base.PEER_USER, fromId, boxId, message.GetData2().GetMentioned(), true)
		ids = append(ids, &IDMessage{UserId: peerId, MessageBoxId: boxId})
	}
	return
}

func (m *messageModel) sendChatMessage(fromId, peerId int32, clientRandomId int64, message *mtproto.Message) (ids []*IDMessage) {
	return
}

func (m *messageModel) sendChannelMessage(fromId, peerId int32, clientRandomId int64, message *mtproto.Message) (ids []*IDMessage) {
	return
}