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
	"errors"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/biz_model/model"
	"github.com/nebulaim/telegramd/biz_model/base"
	"time"
)

type UpdatesServiceImpl struct {
}

func (s *UpdatesServiceImpl) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UpdatesGetState(%v) - Process: {%v}", md, request)

	state := model.GetUpdatesModel().GetState(md.AuthId, md.UserId)
	glog.Infof("UpdatesGetState - reply: {%v}", state)
	return state.ToUpdates_State(), nil
}

// updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
func (s *UpdatesServiceImpl) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	glog.Infof("UpdatesGetDifference - Process: {%v}", request)

	md := grpc_util.RpcMetadataFromIncoming(ctx)


	difference := &mtproto.TLUpdatesDifference{}

	messages := model.GetMessageModel().GetMessagesByPts(md.UserId, request.Pts)
	userIdList := []int32{md.UserId}
	for _, m := range messages {
		peer := base.FromPeer(m.ToId)
		switch peer.PeerType {
		case base.PEER_USER:
			userIdList = append(userIdList, peer.PeerId)
		case base.PEER_CHAT:
		case base.PEER_CHANNEL:
		}

		difference.NewMessages = append(difference.NewMessages, m.ToMessage())
	}

	usersList := model.GetUserModel().GetUserList(userIdList)
	for _, u := range usersList {
		difference.Users = append(difference.Users, u.ToUser())
	}

	state := &mtproto.TLUpdatesState{}

	// TODO(@benqi): Pts通过规则计算出来
	state.Pts = request.Pts + int32(len(messages))
	state.Date = int32(time.Now().Unix())
	state.UnreadCount = 0
	difference.State = state.ToUpdates_State()

	glog.Infof("UpdatesGetDifference - reply: {%v}", difference)
	return difference.ToUpdates_Difference(), nil
}

func (s *UpdatesServiceImpl) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	glog.Infof("UpdatesGetChannelDifference - Process: {%v}", request)
	return nil, errors.New("UpdatesGetChannelDifference - Not impl")
}
