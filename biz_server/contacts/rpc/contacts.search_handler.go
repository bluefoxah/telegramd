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
)

// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (s *ContactsServiceImpl) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsSearch - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi) 使用ES查询
	usersDOList := dao.GetUsersDAO(dao.DB_SLAVE).SelectByQueryString(request.Q, request.Q, request.Q, request.Q)

	found := &mtproto.TLContactsFound{}
	// Peer/Chat/User
	for _, usersDO := range usersDOList {
		p := mtproto.NewTLPeerUser()
		p.SetUserId(usersDO.Id)
		found.Data2.Results = append(found.Data2.Results, p.To_Peer())

		user := mtproto.NewTLUser()
		user.SetId(usersDO.Id)
		if md.UserId == usersDO.Id {
			user.SetSelf(true)
		} else {
			user.SetSelf(false)
		}
		user.SetContact(true)
		user.SetAccessHash(usersDO.AccessHash)
		user.SetFirstName(usersDO.FirstName)
		user.SetLastName(usersDO.LastName)
		user.SetUsername(usersDO.Username)
		user.SetPhone(usersDO.Phone)

		found.Data2.Users = append(found.Data2.Users, user.To_User())
	}

	glog.Infof("ContactsSearch - reply: s%s\n", found)
	return found.To_Contacts_Found(), nil
}
