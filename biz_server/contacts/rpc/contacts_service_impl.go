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
	"github.com/nebulaim/telegramd/base/orm"
)

type ContactsServiceImpl struct {
	zorm orm.Ormer
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
	glog.Info("Process: %v", request)
	return nil, nil
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
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsResolveUsername(ctx context.Context, request *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ContactsServiceImpl) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
