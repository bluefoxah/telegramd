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

func (m *updatesModel) GetState(userId int32) *mtproto.TLUpdatesState {
	state := &mtproto.TLUpdatesState{}

	// updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;

	return state
}