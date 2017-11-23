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
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/biz_model/model"
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/cosiner/gohper/errors"
)

type ContactsServiceImpl struct {
}

// contacts.deleteContacts#59ab389e id:Vector<InputUser> = Bool;
func (s *ContactsServiceImpl) ContactsDeleteContacts(ctx context.Context, request *mtproto.TLContactsDeleteContacts) (*mtproto.Bool, error) {
	glog.Infof("ContactsDeleteContacts - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	idList := make([]int32, 0, len(request.GetId()))
	for _, inputPeer := range request.GetId() {
		switch inputPeer.Payload.(type) {
		case *mtproto.InputUser_InputUserEmpty:
		case *mtproto.InputUser_InputUserSelf:
			idList = append(idList, md.UserId)
		case *mtproto.InputUser_InputUser:
			// TODO(@benqi): Check InputUser's userId and access_hash
			idList = append(idList, inputPeer.GetInputUser().GetUserId())
		}
	}

	dao.GetUserContactsDAO(dao.DB_MASTER).DeleteContacts(md.UserId, idList)
	glog.Infof("ContactsDeleteContacts - reply: {true}")
	return mtproto.ToBool(true), nil
}

// contacts.block#332b49fc id:InputUser = Bool;
func (s *ContactsServiceImpl) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	glog.Infof("ContactsBlock - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	switch request.GetId().Payload.(type) {
	case *mtproto.InputUser_InputUserEmpty:
	case *mtproto.InputUser_InputUserSelf:
		dao.GetUserContactsDAO(dao.DB_MASTER).UpdateBlock(1, md.UserId, md.UserId)
	case *mtproto.InputUser_InputUser:
		// TODO(@benqi): Check InputUser's userId and access_hash
		dao.GetUserContactsDAO(dao.DB_MASTER).UpdateBlock(1, md.UserId, request.GetId().GetInputUser().GetUserId())
	}

	glog.Infof("ContactsBlock - reply: {true}")
	return mtproto.ToBool(true), nil
}

// contacts.unblock#e54100bd id:InputUser = Bool;
func (s *ContactsServiceImpl) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	glog.Infof("ContactsUnblock - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	switch request.GetId().Payload.(type) {
	case *mtproto.InputUser_InputUserEmpty:
	case *mtproto.InputUser_InputUserSelf:
		dao.GetUserContactsDAO(dao.DB_MASTER).UpdateBlock(0, md.UserId, md.UserId)
	case *mtproto.InputUser_InputUser:
		// TODO(@benqi): Check InputUser's userId and access_hash
		dao.GetUserContactsDAO(dao.DB_MASTER).UpdateBlock(0, md.UserId, request.GetId().GetInputUser().GetUserId())
	}

	glog.Infof("ContactsUnblock - reply: {true}")
	return mtproto.ToBool(true), nil
}

// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
// Android client request:
// 	TLRPC.TL_contacts_getBlocked req = new TLRPC.TL_contacts_getBlocked();
// 	req.offset = 0;
// 	req.limit = 200;
func (s *ContactsServiceImpl) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	glog.Infof("ContactsGetBlocked - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	blockedList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectBlockedList(md.UserId, request.Offset, request.Limit)

	blocks := &mtproto.TLContactsBlocked{}

	if len(blockedList) > 0 {
		blockedIdList := make([]int32, 0, len(blockedList))
		for _, c := range blockedList {
			blocked := &mtproto.TLContactBlocked{}
			blocked.UserId = c.ContactUserId
			blocked.Date = c.Date2
			blocks.Blocked = append(blocks.Blocked, blocked.ToContactBlocked())
			blockedIdList = append(blockedIdList, c.ContactUserId)
		}

		users := model.GetUserModel().GetUserList(blockedIdList)
		for _, u := range users {
			blocks.Users = append(blocks.Users, u.ToUser())
		}
	}

	glog.Infof("ContactsSearch - reply: {%v}\n", blocks)
	return blocks.ToContacts_Blocked(), nil
}

// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (s *ContactsServiceImpl) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	glog.Infof("ContactsResetTopPeerRating - Process: {%v}", request)

	_ = grpc_util.RpcMetadataFromIncoming(ctx)
	_ = base.FromInputPeer(request.Peer)

	// TODO(@benqi): 看看客户端代码，什么情况会调用
	switch request.GetCategory().GetPayload().(type) {
	case *mtproto.TopPeerCategory_TopPeerCategoryBotsPM:
	case *mtproto.TopPeerCategory_TopPeerCategoryBotsInline:
	case *mtproto.TopPeerCategory_TopPeerCategoryCorrespondents:
	case *mtproto.TopPeerCategory_TopPeerCategoryGroups:
	case *mtproto.TopPeerCategory_TopPeerCategoryChannels:
	case *mtproto.TopPeerCategory_TopPeerCategoryPhoneCalls:
	}

	glog.Infof("ContactsResetTopPeerRating - reply: {true}")
	return mtproto.ToBool(true), nil
}

// contacts.resetSaved#879537f1 = Bool;
func (s *ContactsServiceImpl) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	glog.Infof("ContactsResetSaved - Process: %v", request)

	// TODO(@benqi): 客户端未调用此请求
	glog.Infof("ContactsResetSaved - reply: {true}")
	return mtproto.ToBool(true), nil
}

// contacts.importCard#4fe196fe export_card:Vector<int> = User;
func (s *ContactsServiceImpl) ContactsImportCard(ctx context.Context, request *mtproto.TLContactsImportCard) (*mtproto.User, error) {
	glog.Infof("ContactsImportCard - Process: %v", request)
	return nil, errors.New("Not impl")
}

// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
// func (s *ContactsServiceImpl)ContactsGetStatuses(ctx context.Context,  request *mtproto.TLContactsGetStatuses) (*mtproto.Vector<ContactStatus>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

// contacts.getContacts#c023849f hash:int = contacts.Contacts;
func (s *ContactsServiceImpl) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	glog.Infof("ContactsGetContacts: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): Logout逻辑处理，失效AuthKey
	// reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	contacts := &mtproto.TLContactsContacts{}

	contactsDOList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectUserContacts(md.UserId)
	contacts.SavedCount = int32(len(contactsDOList))

	for _, do := range contactsDOList {
		contact := &mtproto.TLContact{}
		contact.UserId = do.ContactUserId
		contact.Mutual = mtproto.MakeBool(&mtproto.TLBoolFalse{})

		contacts.Contacts = append(contacts.Contacts, mtproto.MakeContact(contact))

		userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(do.ContactUserId)
		user := &mtproto.TLUser{}
		user.Id = userDO.Id
		if user.Id == md.UserId {
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

// @benqi: Android client
// contacts.getContacts#22c6aa08 hash:string = contacts.Contacts;
func (s *ContactsServiceImpl) ContactsGetContacts2(ctx context.Context, request *mtproto.TLContactsGetContacts2) (*mtproto.Contacts_Contacts, error) {
	glog.Infof("ContactsGetContacts2: %v", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	// TODO(@benqi): Logout逻辑处理，失效AuthKey
	// reply := mtproto.MakeBool(&mtproto.TLBoolTrue{})

	contacts := &mtproto.TLContactsContacts{}

	contactsDOList := dao.GetUserContactsDAO(dao.DB_SLAVE).SelectUserContacts(md.UserId)
	contacts.SavedCount = int32(len(contactsDOList))

	for _, do := range contactsDOList {
		contact := &mtproto.TLContact{}
		contact.UserId = do.ContactUserId
		contact.Mutual = mtproto.MakeBool(&mtproto.TLBoolFalse{})

		contacts.Contacts = append(contacts.Contacts, mtproto.MakeContact(contact))

		userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(do.ContactUserId)
		user := &mtproto.TLUser{}
		user.Id = userDO.Id
		if user.Id == md.UserId {
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

	glog.Infof("ContactsGetContacts2 - reply: {%v}\n", reply)
	return reply, nil
}

// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (s *ContactsServiceImpl) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	glog.Infof("ContactsImportContacts - Process: {%v}", request)
	return nil, errors.New("Not impl")
}

// contacts.deleteContact#8e953744 id:InputUser = contacts.Link;
func (s *ContactsServiceImpl) ContactsDeleteContact(ctx context.Context, request *mtproto.TLContactsDeleteContact) (*mtproto.Contacts_Link, error) {
	glog.Infof("ContactsDeleteContact - Process: %v", request)
	return nil, errors.New("Not impl")
}

// contacts.exportCard#84e53737 = Vector<int>;
// func (s *ContactsServiceImpl)ContactsExportCard(ctx context.Context,  request *mtproto.TLContactsExportCard) (*mtproto.Vector<int32T>, error) {
//   glog.Info("Process: %v", request)
//	 return nil, errors.New("Not impl")
// }

// contacts.search#11f812d8 q:string limit:int = contacts.Found;
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

// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (s *ContactsServiceImpl) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	glog.Infof("ContactsResolveUsername - Process: {%v}", request)
	return nil, errors.New("Not impl")
}

// contacts.getTopPeers#d4982db5 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:int = contacts.TopPeers;
func (s *ContactsServiceImpl) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	glog.Info("ContactsGetTopPeers - Process: {%v}", request)
	return nil, errors.New("Not impl")
}
