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
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/mtproto"
	"os/user"
)

type UserModel struct {
	usersDAO *dao.UsersDAO
}

func NewUserModel(user *dao.UsersDAO) *UserModel {
	return &UserModel{user}
}

func (m *UserModel) GetUser(userId int32) (user* mtproto.TLUser) {
	userDO, _ := m.usersDAO.SelectById(userId)
	if userDO != nil {
		// TODO(@benqi): fill bot, photo, about...
		user = &mtproto.TLUser{}
		// user.Self由业务层进行判断
		// user.Self = true
		user.Id = userDO.Id
		user.AccessHash = userDO.AccessHash
		user.FirstName = userDO.FirstName
		user.LastName = userDO.LastName
		user.Username = userDO.Username
		user.Phone = userDO.Phone
	}
	return
}

func (m *UserModel) GetUserList(userIdList []int32) (users []*mtproto.TLUser) {
	userDOList, _ := m.usersDAO.SelectUsersByIdList(userIdList)
	users = make([]*mtproto.TLUser, len(userDOList))
	for _, userDO := range userDOList {
		// TODO(@benqi): fill bot, photo, about...
		user := &mtproto.TLUser{}
		// user.Self由业务层进行判断
		// user.Self = true
		user.Id = userDO.Id
		user.AccessHash = userDO.AccessHash
		user.FirstName = userDO.FirstName
		user.LastName = userDO.LastName
		user.Username = userDO.Username
		user.Phone = userDO.Phone

		users = append(users, user)
	}
	return
}

func (m *UserModel) GetUserFull(userId int32) (userFull *mtproto.TLUserFull) {
	//TODO(@benqi): 等Link和NotifySettings实现后再来完善
	//fullUser := &mtproto.TLUserFull{}
	//fullUser.PhoneCallsAvailable = true
	//fullUser.PhoneCallsPrivate = true
	//fullUser.About = "@Benqi"
	//fullUser.CommonChatsCount = 0
	//
	//switch request.Id.Payload.(type) {
	//case *mtproto.InputUser_InputUserSelf:
	//	// User
	//	userDO, _ := s.UsersDAO.SelectById(2)
	//	user := &mtproto.TLUser{}
	//	user.Self = true
	//	user.Contact = false
	//	user.Id = userDO.Id
	//	user.FirstName = userDO.FirstName
	//	user.LastName = userDO.LastName
	//	user.Username = userDO.Username
	//	user.AccessHash = userDO.AccessHash
	//	user.Phone = userDO.Phone
	//	fullUser.User = mtproto.MakeUser(user)
	//
	//	// Link
	//	link := &mtproto.TLContactsLink{}
	//	link.MyLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
	//	link.ForeignLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
	//	link.User = mtproto.MakeUser(user)
	//	fullUser.Link = mtproto.MakeContacts_Link(link)
	//case *mtproto.InputUser_InputUser:
	//	inputUser := request.Id.Payload.(*mtproto.InputUser_InputUser).InputUser
	//	// User
	//	userDO, _ := s.UsersDAO.SelectById(inputUser.UserId)
	//	user := &mtproto.TLUser{}
	//	user.Self = false
	//	user.Contact = true
	//	user.Id = userDO.Id
	//	user.FirstName = userDO.FirstName
	//	user.LastName = userDO.LastName
	//	user.Username = userDO.Username
	//	user.AccessHash = userDO.AccessHash
	//	user.Phone = userDO.Phone
	//	fullUser.User = mtproto.MakeUser(user)
	//
	//	// Link
	//	link := &mtproto.TLContactsLink{}
	//	link.MyLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
	//	link.ForeignLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
	//	link.User = mtproto.MakeUser(user)
	//	fullUser.Link = mtproto.MakeContacts_Link(link)
	//case *mtproto.InputUser_InputUserEmpty:
	//	// TODO(@benqi): BAD_REQUEST: 400
	//}
	//
	//// NotifySettings
	//peerNotifySettings := &mtproto.TLPeerNotifySettings{}
	//peerNotifySettings.ShowPreviews = true
	//peerNotifySettings.MuteUntil = 0
	//peerNotifySettings.Sound = "default"
	//fullUser.NotifySettings = mtproto.MakePeerNotifySettings(peerNotifySettings)
	return nil
}

//func (m *UserModel) GetUserFullList(userId []int32) (user []*mtproto.TLUserFull) {
//	return nil
//}
