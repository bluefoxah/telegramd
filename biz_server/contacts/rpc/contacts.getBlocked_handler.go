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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/model"
)

// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
func (s *ContactsServiceImpl) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsGetBlocked - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	blockedList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectBlockedList(md.UserId, request.Offset, request.Limit)

	blocks := &mtproto.TLContactsBlocked{}

	if len(blockedList) > 0 {
		blockedIdList := make([]int32, 0, len(blockedList))
		for _, c := range blockedList {
			blocked := mtproto.NewTLContactBlocked()
			blocked.SetUserId(c.ContactUserId)
			blocked.SetDate(c.Date2)
			blocks.Data2.Blocked = append(blocks.Data2.Blocked, blocked.To_ContactBlocked())
			blockedIdList = append(blockedIdList, c.ContactUserId)
		}

		users := model.GetUserModel().GetUserList(blockedIdList)
		for _, u := range users {
			blocks.Data2.Users = append(blocks.Data2.Users, u.To_User())
		}
	}

	glog.Infof("ContactsSearch - reply: %s\n", logger.JsonDebugData(blocks))
	return blocks.To_Contacts_Blocked(), nil
}
