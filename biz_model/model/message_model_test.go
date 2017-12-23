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
	"testing"
	"fmt"
	"github.com/nebulaim/telegramd/base/mysql_client"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/base/logger"
)

func init()  {
	mysqlConfig1 := mysql_client.MySQLConfig{
		Name:   "immaster",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}

	mysqlConfig2 := mysql_client.MySQLConfig{
		Name:   "imslave",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}

	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig1, mysqlConfig2})
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
}

func TestGetMessagesByUserIdPeerOffsetLimit(t *testing.T) {
	// 1. 先从message_boxes取出message_id
	boxesList := dao.GetMessageBoxesDAO(dao.DB_SLAVE).SelectBackwardByPeerOffsetLimit(2, base.PEER_USER, 3, 429, 50)
	fmt.Printf("GetMessagesByUserIdPeerOffsetLimit - boxesList: %s", logger.JsonDebugData(boxesList))

	// messags := GetMessageModel().GetMessagesByUserIdPeerOffsetLimit(2, base.PEER_USER, 3, 429, 50)
	// fmt.Printf("messages - %s\n", logger.JsonDebugData(messags))
}
