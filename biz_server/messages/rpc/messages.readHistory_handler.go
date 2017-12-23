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
	"github.com/nebulaim/telegramd/biz_model/base"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"time"
	"github.com/nebulaim/telegramd/biz_server/delivery"
	"github.com/nebulaim/telegramd/biz_model/model"
	base2 "github.com/nebulaim/telegramd/base/base"
)

// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReadHistory - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	/*
        body: { rpc_result
          req_msg_id: 6498830637519766076 [LONG],
          result: { messages_affectedMessages
            pts: 567 [INT],
            pts_count: 1 [INT],
          },
        },
	 */

	peer := base.FromInputPeer(request.GetPeer())

	// Delivery updateReadHistoryOutbox to peer
	// peer := base.FromInputPeer(request.GetPeer())
	updateReadHistoryOutbox := mtproto.NewTLUpdateReadHistoryOutbox()
	oudboxDO := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPeer(peer.PeerId, int8(peer.PeerType), md.UserId)
	outboxPeer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
		UserId: md.UserId,
	}}

	// lastPts := model.GetMessageModel().GetLastPtsByUserId(peer.PeerId)
	// _ = lastPts
	updateReadHistoryOutbox.SetPeer(outboxPeer.To_Peer())

	// TODO(@benqi): db里存放什么？
	pts := model.GetSequenceModel().NextID(base2.Int32ToString(peer.PeerId))
	updateReadHistoryOutbox.SetPts(int32(pts))
	updateReadHistoryOutbox.SetPtsCount(1)
	updateReadHistoryOutbox.SetMaxId(oudboxDO.TopMessage)

	updates := mtproto.NewTLUpdates()
	updates.SetSeq(0)
	updates.SetDate(int32(time.Now().Unix()))
	updates.SetUpdates([]*mtproto.Update{updateReadHistoryOutbox.To_Update()})

	delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
		md.AuthId,
		md.SessionId,
		md.NetlibSessionId,
		[]int32{peer.PeerId},
		updates.To_Updates().Encode())

	// TODO(@benqi): 如何处理多端登录？

	// TODO(@benqi):
	dao.GetUserDialogsDAO(dao.DB_MASTER).UpdateUnreadByPeer(md.UserId, int8(peer.PeerType), peer.PeerId)
	// model.GetDialogModel().get
	// 1. maxId找到
	// lastPts = model.GetMessageModel().GetLastPtsByUserId(peer.PeerId)
	affected := mtproto.NewTLMessagesAffectedMessages()

	// TODO(@benqi): db里存放什么？
	pts = model.GetSequenceModel().NextID(base2.Int32ToString(peer.PeerId))
	affected.SetPts(int32(pts))
	affected.SetPtsCount(1)

	glog.Infof("MessagesReadHistory - reply: %s", logger.JsonDebugData(affected))
	return affected.To_Messages_AffectedMessages(), nil
}
