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
//"github.com/nebulaim/telegramd/grpc_util"
//"github.com/golang/glog"
//"github.com/nebulaim/telegramd/base/logger"
//"fmt"
)

/*
import (
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
)

// @benqi: Android client
// contacts.getContacts#22c6aa08 hash:string = contacts.Contacts;
func (s *ContactsServiceImpl) ContactsGetContacts2(ctx context.Context, request *mtproto.TLContactsGetContacts2) (*mtproto.Contacts_Contacts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ContactsGetContacts - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// TODO(@benqi): Logout逻辑处理，失效AuthKey
	//// reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})
	//
	//contacts := &mtproto.TLContactsContacts{}
	//
	//contactsDOList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectUserContacts(md.UserId)
	//contacts.SavedCount = int32(len(contactsDOList))
	//
	//for _, do := range contactsDOList {
	//	contact := &mtproto.TLContact{}
	//	contact.UserId = do.ContactUserId
	//	contact.Mutual = mtproto.MakeBool(&mtproto.TLBoolFalse{})
	//
	//	contacts.Contacts = append(contacts.Contacts, mtproto.MakeContact(contact))
	//
	//	userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(do.ContactUserId)
	//	user := &mtproto.TLUser{}
	//	user.Id = userDO.Id
	//	if user.Id == md.UserId {
	//		user.Self = true
	//	} else {
	//		user.Self = false
	//	}
	//	user.Contact = true
	//	user.AccessHash = userDO.AccessHash
	//	user.FirstName = userDO.FirstName
	//	user.LastName = userDO.LastName
	//	user.Username = userDO.Username
	//	user.Phone = userDO.Phone
	//
	//	contacts.Users = append(contacts.Users, mtproto.MakeUser(user))
	//}
	//
	//reply := mtproto.MakeContacts_Contacts(contacts)
	//
	//glog.Infof("ContactsGetContacts2 - reply: {%v}\n", reply)
	//return reply, nil

	return nil, fmt.Errorf("Not impl ContactsGetContacts")

}
*/
