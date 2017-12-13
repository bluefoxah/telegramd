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

// users.getFullUser#ca30a5b1 id:InputUser = UserFull;
func (s *UsersServiceImpl) UsersGetFullUser(ctx context.Context, request *mtproto.TLUsersGetFullUser) (*mtproto.UserFull, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UsersGetFullUser - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	fullUser := mtproto.NewTLUserFull()
	fullUser.SetPhoneCallsAvailable(true)
	fullUser.SetPhoneCallsPrivate(true)
	fullUser.SetAbout("@Benqi")
	fullUser.SetCommonChatsCount(0)

	switch request.GetId().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
	    // User
	    userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(md.UserId)
	    user := &mtproto.TLUser{ Data2: &mtproto.User_Data{
			Self:       true,
			Contact:    true,
			Id:         userDO.Id,
			FirstName:  userDO.FirstName,
			LastName:   userDO.LastName,
			Username:   userDO.Username,
			AccessHash: userDO.AccessHash,
			Phone:      userDO.Phone,
		}}
	    fullUser.SetUser(user.To_User())

	    // Link
	    link := &mtproto.TLContactsLink{ Data2: &mtproto.Contacts_Link_Data{
	    	MyLink:	mtproto.NewTLContactLinkContact().To_ContactLink(),
	    	ForeignLink: mtproto.NewTLContactLinkContact().To_ContactLink(),
	    	User: user.To_User(),
		}}
	    fullUser.SetLink(link.To_Contacts_Link())
	case mtproto.TLConstructor_CRC32_inputUser:
	    inputUser := request.GetId().To_InputUser()
	    // request.Id.Payload.(*mtproto.InputUser_InputUser).InputUser
	    // User
	    userDO := dao.GetUsersDAO(dao.DB_SLAVE).SelectById(inputUser.GetUserId())
		user := &mtproto.TLUser{ Data2: &mtproto.User_Data{
			Self:       md.UserId == inputUser.GetUserId(),
			Contact:    true,
			Id:         userDO.Id,
			FirstName:  userDO.FirstName,
			LastName:   userDO.LastName,
			Username:   userDO.Username,
			AccessHash: userDO.AccessHash,
			Phone:      userDO.Phone,
		}}
		fullUser.SetUser(user.To_User())

	    // Link
		link := &mtproto.TLContactsLink{ Data2: &mtproto.Contacts_Link_Data{
			MyLink:	mtproto.NewTLContactLinkContact().To_ContactLink(),
			ForeignLink: mtproto.NewTLContactLinkContact().To_ContactLink(),
			User: user.To_User(),
		}}
		fullUser.SetLink(link.To_Contacts_Link())
	case mtproto.TLConstructor_CRC32_inputUserEmpty:
	    // TODO(@benqi): BAD_REQUEST: 400
	}

	// NotifySettings
	peerNotifySettings := &mtproto.TLPeerNotifySettings{ Data2: &mtproto.PeerNotifySettings_Data{
		ShowPreviews: true,
		MuteUntil:    0,
		Sound:        "default",
	}}

	fullUser.SetNotifySettings(peerNotifySettings.To_PeerNotifySettings())

	glog.Infof("UsersGetFullUser - reply: %s", logger.JsonDebugData(fullUser))
	return fullUser.To_UserFull(), nil
}
