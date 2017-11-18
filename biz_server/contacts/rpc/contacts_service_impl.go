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
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"google.golang.org/grpc/metadata"
)

type ContactsServiceImpl struct {
}

func (s *ContactsServiceImpl) ContactsDeleteContacts(ctx context.Context, request *mtproto.TLContactsDeleteContacts) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsImportCard(ctx context.Context, request *mtproto.TLContactsImportCard) (*mtproto.User, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *ContactsServiceImpl)ContactsGetStatuses(ctx context.Context,  request *mtproto.TLContactsGetStatuses) (*mtproto.Vector<ContactStatus>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *ContactsServiceImpl) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	glog.Infof("ContactsGetContacts: %v", request)

	md, _ := metadata.FromIncomingContext(ctx)
	rpcMetaData := mtproto.RpcMetaData{}
	rpcMetaData.Decode(md)

	// TODO(@benqi): Logout逻辑处理，失效AuthKey
	// reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	contacts := &mtproto.TLContactsContacts{}

	contactsDOList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectUserContacts(rpcMetaData.UserId)
	contacts.SavedCount = int32(len(contactsDOList))

	for _, do := range contactsDOList {
		contact := &mtproto.TLContact{}
		contact.UserId = do.ContactUserId
		contact.Mutual = mtproto.MakeBool(&mtproto.TLBoolFalse{})

		contacts.Contacts = append(contacts.Contacts, mtproto.MakeContact(contact))

		userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(do.ContactUserId)
		user := &mtproto.TLUser{}
		user.Id = userDO.Id
		if user.Id == rpcMetaData.UserId {
			user.Self = true
		} else {
			user.Self = false
		}
		user.Contact = true
		user.AccessHash = userDO.AccessHash
		user.FirstName = userDO.FirstName
		user.LastName = userDO.LastName
		user.Username = userDO.Username
		user.Phone = userDO.Phone

		contacts.Users = append(contacts.Users, mtproto.MakeUser(user))
	}

	reply := mtproto.MakeContacts_Contacts(contacts)

	glog.Infof("ContactsGetContacts - reply: {%v}\n", reply)
	return reply, nil
}

func (s *ContactsServiceImpl) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsDeleteContact(ctx context.Context, request *mtproto.TLContactsDeleteContact) (*mtproto.Contacts_Link, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *ContactsServiceImpl)ContactsExportCard(ctx context.Context,  request *mtproto.TLContactsExportCard) (*mtproto.Vector<int32T>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *ContactsServiceImpl) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	glog.Infof("ContactsSearch - Process: {%v}", request)

	// TODO(@benqi) 使用ES查询
	usersDOList := dao.GetUsersDAO(dao.DB_SLAVE).SelectByQueryString(request.Q, request.Q, request.Q, request.Q)

	found := &mtproto.TLContactsFound{}
	// Peer/Chat/User
	for _, usersDO := range usersDOList {
		found.Results = append(found.Results, mtproto.MakePeer(&mtproto.TLPeerUser{UserId: usersDO.Id}))

		user := &mtproto.TLUser{}
		user.Id = usersDO.Id
		user.Self = false
		user.Contact = true
		user.AccessHash = usersDO.AccessHash
		user.FirstName = usersDO.FirstName
		user.LastName = usersDO.LastName
		user.Username = usersDO.Username
		user.Phone = usersDO.Phone

		found.Users = append(found.Users, mtproto.MakeUser(user))
	}

	reply := mtproto.MakeContacts_Found(found)

	glog.Infof("ContactsSearch - reply: {%v}\n", reply)
	return reply, nil
}

func (s *ContactsServiceImpl) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
