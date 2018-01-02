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
	//"github.com/nebulaim/telegramd/mtproto"
	//"github.com/nebulaim/telegramd/biz_model/dal/dao"
	//"time"
	"time"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

type updatesModel struct {
}

var (
	updatesInstance *updatesModel
	updatesInstanceOnce sync.Once
)

func GetUpdatesModel() *updatesModel {
	updatesInstanceOnce.Do(func() {
		updatesInstance = &updatesModel{}
	})
	return updatesInstance
}

func (m *updatesModel) GetState(authKeyId int64, userId int32) *mtproto.TLUpdatesState {
	state := mtproto.NewTLUpdatesState()
	stateData := state.GetData2()

	// TODO(@benqi): 从数据库取出date
	stateData.Date = int32(time.Now().Unix())

	do := dao.GetAuthUpdatesStateDAO(dao.DB_SLAVE).SelectById(authKeyId, userId)
	if do == nil || do.Pts == 0 {
		// TODO(@benqi):
		stateData.Date = int32(time.Now().Unix())
		stateData.Pts = 1
		stateData.Qts = 0
		stateData.Seq = 1
		stateData.UnreadCount = 0
	} else {
		stateData.Pts = do.Pts
		stateData.Qts = do.Qts
		stateData.Seq = do.Seq
		stateData.UnreadCount = 0
	}
	return state
}

func (m *updatesModel) AddPtsToUpdatesQueue(userId, pts, peerType, peerId, updateType, messageBoxId, maxMessageBoxId int32, ) int32 {
	do := &dataobject.UserPtsUpdatesDO{
		UserId:          userId,
		PeerType:		 int8(peerType),
		PeerId:			 peerId,
		Pts:             pts,
		UpdateType:      updateType,
		MessageBoxId:    messageBoxId,
		MaxMessageBoxId: maxMessageBoxId,
		Date2:           int32(time.Now().Unix()),
	}

	return int32(dao.GetUserPtsUpdatesDAO(dao.DB_MASTER).Insert(do))
}

func (m *updatesModel) AddQtsToUpdatesQueue(userId, qts, updateType int32, updateData []byte) int32 {
	do := &dataobject.UserQtsUpdatesDO{
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: updateData,
		Date2:      int32(time.Now().Unix()),
		Qts:        qts,
	}

	return int32(dao.GetUserQtsUpdatesDAO(dao.DB_MASTER).Insert(do))
}

func (m *updatesModel) AddSeqToUpdatesQueue(userId, seq, updateType int32, updateData []byte) int32 {
	do := &dataobject.UserSeqUpdatesDO{
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: updateData,
		Date2:      int32(time.Now().Unix()),
		Seq:        seq,
	}

	return int32(dao.GetUserSeqUpdatesDAO(dao.DB_MASTER).Insert(do))
}

//func (m *updatesModel) GetAffectedMessage(userId, messageId int32) *mtproto.TLMessagesAffectedMessages {
//	doList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectPtsByGTMessageID(userId, messageId)
//	if len(doList) == 0 {
//		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2), fmt.Sprintf("GetAffectedMessage(%d, %d) empty", userId, messageId)))
//	}
//
//	affected := &mtproto.TLMessagesAffectedMessages{}
//	affected.Pts = doList[0].Pts
//	affected.PtsCount = int32(len(doList))
//	return affected
//}

