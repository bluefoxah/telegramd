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


/*
import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/grpc_util"
)

type UsersServiceImpl struct {
	// UsersDAO *dao.UsersDAO
}

// func (s *UsersServiceImpl)UsersGetUsers(ctx context.Context,  request *mtproto.TLUsersGetUsers) (*mtproto.Vector<User>, error) {
//	 glog.Infof("UsersGetFullUser - Process: {%v}", request)
//   return nil, errors.New("UsersGetFullUser - Not impl")
// }

// userFull#f220f3f
// 	flags:#
//  	blocked:flags.0?true
//  	phone_calls_available:flags.4?true
//  	phone_calls_private:flags.5?true
//  user:User
// 		about:flags.1?string
//  link:contacts.Link
//  	profile_photo:flags.2?Photo
//  notify_settings:PeerNotifySettings
//  	bot_info:flags.3?BotInfo
//  common_chats_count:int
// = UserFull;
func (s *UsersServiceImpl) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.UserFull, error) {
	glog.Infof("UsersGetFullUser - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)

	fullUser := &mtproto.TLUserFull{}
	fullUser.PhoneCallsAvailable = true
	fullUser.PhoneCallsPrivate = true
	fullUser.About = "@Benqi"
	fullUser.CommonChatsCount = 0

	switch request.Id.Payload.(type) {
	case *mtproto.InputUser_InputUserSelf:
		// User
		userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(md.UserId)
		user := &mtproto.TLUser{}
		user.Self = true
		user.Contact = false
		user.Id = userDO.Id
		user.FirstName = userDO.FirstName
		user.LastName = userDO.LastName
		user.Username = userDO.Username
		user.AccessHash = userDO.AccessHash
		user.Phone = userDO.Phone
		fullUser.User = mtproto.MakeUser(user)

		// Link
		link := &mtproto.TLContactsLink{}
		link.MyLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
		link.ForeignLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
		link.User = mtproto.MakeUser(user)
		fullUser.Link = mtproto.MakeContacts_Link(link)
	case *mtproto.InputUser_InputUser:
		inputUser := request.Id.Payload.(*mtproto.InputUser_InputUser).InputUser
		// User
		userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(inputUser.UserId)
		user := &mtproto.TLUser{}
		user.Self = md.UserId == inputUser.UserId
		user.Contact = true
		user.Id = userDO.Id
		user.FirstName = userDO.FirstName
		user.LastName = userDO.LastName
		user.Username = userDO.Username
		user.AccessHash = userDO.AccessHash
		user.Phone = userDO.Phone
		fullUser.User = mtproto.MakeUser(user)

		// Link
		link := &mtproto.TLContactsLink{}
		link.MyLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
		link.ForeignLink = mtproto.MakeContactLink(&mtproto.TLContactLinkContact{})
		link.User = mtproto.MakeUser(user)
		fullUser.Link = mtproto.MakeContacts_Link(link)
	case *mtproto.InputUser_InputUserEmpty:
		// TODO(@benqi): BAD_REQUEST: 400
	}

	// NotifySettings
	peerNotifySettings := &mtproto.TLPeerNotifySettings{}
	peerNotifySettings.ShowPreviews = true
	peerNotifySettings.MuteUntil = 0
	peerNotifySettings.Sound = "default"
	fullUser.NotifySettings = mtproto.MakePeerNotifySettings(peerNotifySettings)

	reply := mtproto.MakeUserFull(fullUser)
	glog.Infof("UsersGetFullUser - reply: {%v}\n", reply)
	return reply, nil
}
*/
