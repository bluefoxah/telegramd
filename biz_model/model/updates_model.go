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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"time"
	"fmt"
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
	state := &mtproto.TLUpdatesState{}

	// TODO(@benqi): 从数据库取出date
	state.Date = int32(time.Now().Unix())

	do := dao.GetMessageBoxsDAO(dao.DB_SLAVE).SelectLastPts(userId)
	if do == nil || do.Pts == 0 {
		// TODO(@benqi):
		state.Date = int32(time.Now().Unix())
		state.Pts = 1
		state.Qts = 0
		state.Seq = 1
		state.UnreadCount = 0
	} else {
		state.Pts = do.Pts
		state.Qts = 0
		state.Seq = 1
		state.UnreadCount = 0
	}
	return state
}

func (m *updatesModel) GetAffectedMessage(userId, messageId int32) *mtproto.TLMessagesAffectedMessages {
	doList := dao.GetMessageBoxsDAO(dao.DB_SLAVE).SelectPtsByGTMessageID(userId, messageId)
	if len(doList) == 0 {
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2), fmt.Sprintf("GetAffectedMessage(%d, %d) empty", userId, messageId)))
	}

	affected := &mtproto.TLMessagesAffectedMessages{}
	affected.Pts = doList[0].Pts
	affected.PtsCount = int32(len(doList))
	return affected
}
