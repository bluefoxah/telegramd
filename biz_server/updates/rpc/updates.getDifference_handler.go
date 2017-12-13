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
	"time"
	"github.com/nebulaim/telegramd/biz_model/model"
	"github.com/nebulaim/telegramd/biz_model/base"
)

// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (s *UpdatesServiceImpl) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UpdatesGetDifference - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	difference := mtproto.NewTLUpdatesDifference()
	messages := model.GetMessageModel().GetMessagesByGtPts(md.UserId, request.Pts)
	userIdList := []int32{}
	chatIdList := []int32{}

	for _, m := range messages {
	    switch m.GetConstructor()  {
	    case mtproto.TLConstructor_CRC32_message:
	        m2 := m.To_Message()
	        userIdList = append(userIdList, m2.GetFromId())
	        peer := base.FromPeer(m2.GetToId())
	        switch peer.PeerType {
	        case base.PEER_USER:
	            userIdList = append(userIdList, peer.PeerId)
	        case base.PEER_CHAT:
	            chatIdList = append(chatIdList, peer.PeerId)
	        case base.PEER_CHANNEL:
	            // TODO(@benqi): add channel
	        }
	    case mtproto.TLConstructor_CRC32_messageService:
	        m2 := m.To_MessageService()
	        userIdList = append(userIdList, m2.GetFromId())
	        chatIdList = append(chatIdList, m2.GetToId().GetData2().GetChatId())
	    case mtproto.TLConstructor_CRC32_messageEmpty:
	    }
	    difference.Data2.NewMessages = append(difference.Data2.NewMessages, m)
	}

	if len(userIdList) > 0 {
	    usersList := model.GetUserModel().GetUserList(userIdList)
	    for _, u := range usersList {
	        if u.GetId() == md.UserId {
	            u.SetSelf(true)
	        }
	        u.SetContact(true)
	        u.SetMutualContact(true)
	        difference.Data2.Users = append(difference.Data2.Users, u.To_User())
	    }
	}

	state := mtproto.NewTLUpdatesState()

	// TODO(@benqi): Pts通过规则计算出来
	state.SetPts(request.Pts + int32(len(messages)))
	state.SetDate(int32(time.Now().Unix()))
	state.SetUnreadCount(0)
	difference.SetState(state.To_Updates_State())

	glog.Infof("UpdatesGetDifference - reply: %s", difference)
	return difference.To_Updates_Difference(), nil
}
