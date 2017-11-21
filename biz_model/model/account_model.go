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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"time"
)

const (
	TOKEN_TYPE_APNS = 1
	TOKEN_TYPE_GCM = 2
	TOKEN_TYPE_MPNS = 3
	TOKEN_TYPE_SIMPLE_PUSH = 4
	TOKEN_TYPE_UBUNTU_PHONE = 5
	TOKEN_TYPE_BLACKBERRY = 6
	// Android里使用
	TOKEN_TYPE_INTERNAL_PUSH = 7
)

type accountModel struct {
	//
}

var (
	accountInstance *accountModel
	accountInstanceOnce sync.Once
)

func GetAccountModel() *accountModel {
	accountInstanceOnce.Do(func() {
		accountInstance = &accountModel{}
	})
	return accountInstance
}

func (m *accountModel) RegisterDevice(authKeyId int64, userId int32, tokenType int8, token string) {
	slave := dao.GetDevicesDAO(dao.DB_SLAVE)
	master := dao.GetDevicesDAO(dao.DB_MASTER)
	do := slave.SelectId(authKeyId, tokenType, token)
	if do == nil {
		do = &dataobject.DevicesDO{}
		do.AuthId = authKeyId
		do.UserId = userId
		do.TokenType = tokenType
		do.Token = token
		do.State = 0
		do.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		master.Insert(do)
	} else {
		master.UpdateStateById(1, do.Id)
	}
}

func (m *accountModel) UnRegisterDevice(authKeyId int64, userId int32, tokenType int8, token string) bool {
	master := dao.GetDevicesDAO(dao.DB_MASTER)
	rows := master.UpdateState(tokenType, authKeyId, tokenType, token)
	return rows == 1
}
