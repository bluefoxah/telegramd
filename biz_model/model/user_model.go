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
	"sync"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/base/base"
)

type userModel struct {
	// usersDAO *dao.UsersDAO
}

var (
	userInstance *userModel
	userInstanceOnce sync.Once
)

func GetUserModel() *userModel {
	userInstanceOnce.Do(func() {
		userInstance = &userModel{}
	})
	return userInstance
}

func (m *userModel) GetUser(userId int32) (user* mtproto.TLUser) {
	usersDAO := dao.GetUsersDAO(dao.DB_SLAVE)

	userDO := usersDAO.SelectById(userId)
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

func (m *userModel) GetUserList(userIdList []int32) (users []*mtproto.TLUser) {
	usersDAO := dao.GetUsersDAO(dao.DB_SLAVE)

	userDOList := usersDAO.SelectUsersByIdList(userIdList)
	users = []*mtproto.TLUser{}
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

	glog.Infof("SelectUsersByIdList(%s) - {%v}", base.JoinInt32List(userIdList, ","), users)

	return
}

func (m *userModel) GetUserFull(userId int32) (userFull *mtproto.TLUserFull) {
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
