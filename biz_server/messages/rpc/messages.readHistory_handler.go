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

/*
	// updateReadHistoryOutbox
	// updateReadHistoryInbox
	// messages_affectedMessages
 */
// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("MessagesReadHistory - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.GetPeer())

	// 消息已读逻辑
	// 1. inbox，设置unread_count为0以及read_inbox_max_id
	// inBoxDO := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPeer(md.UserId, int8(peer.PeerType), peer.PeerId)
	dao.GetUserDialogsDAO(dao.DB_MASTER).UpdateUnreadByPeer(request.GetMaxId(), md.UserId, int8(peer.PeerType), peer.PeerId)
	// return me
	pts := int32(model.GetSequenceModel().NextPtsId(base2.Int32ToString(md.UserId)))
	model.GetUpdatesModel().AddPtsToUpdatesQueue(md.UserId, pts, base.PEER_USER, peer.PeerId, model.PTS_READ_HISTORY_INBOX, 0, request.GetMaxId())

	affected := mtproto.NewTLMessagesAffectedMessages()
	// pts = model.GetSequenceModel().NextPtsId(base2.Int32ToString(peer.PeerId))
	affected.SetPts(int32(pts))
	affected.SetPtsCount(1)

	// outboxPeer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
	// 	UserId: md.UserId,
	// }}
	// 消息漫游
	updateReadHistoryInbox := mtproto.NewTLUpdateReadHistoryInbox()
	updateReadHistoryInbox.SetPeer(peer.ToPeer())
	updateReadHistoryInbox.SetPts(pts)
	updateReadHistoryInbox.SetPtsCount(1)
	updateReadHistoryInbox.SetMaxId(request.MaxId)

	updates := mtproto.NewTLUpdates()
	updates.SetSeq(0)
	updates.SetDate(int32(time.Now().Unix()))
	updates.SetUpdates([]*mtproto.Update{updateReadHistoryInbox.To_Update()})

	delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
		md.AuthId,
		md.SessionId,
		md.NetlibSessionId,
		[]int32{md.UserId},
		updates.To_Updates().Encode())

	// 2. outbox, 设置read_outbox_max_id
	outboxDO := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPeer(peer.PeerId, int8(peer.PeerType), md.UserId)
	dao.GetUserDialogsDAO(dao.DB_MASTER).UpdateReadOutboxMaxIdByPeer(outboxDO.TopMessage, peer.PeerId, int8(peer.PeerType), md.UserId)
	pts = int32(model.GetSequenceModel().NextPtsId(base2.Int32ToString(peer.PeerId)))
	model.GetUpdatesModel().AddPtsToUpdatesQueue(peer.PeerId, pts, base.PEER_USER, md.UserId, model.PTS_READ_HISTORY_OUTBOX, 0, outboxDO.TopMessage)

	updateReadHistoryOutbox := mtproto.NewTLUpdateReadHistoryOutbox()
	// oudboxDO := dao.GetUserDialogsDAO(dao.DB_SLAVE).SelectByPeer(peer.PeerId, int8(peer.PeerType), md.UserId)
	outboxPeer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
		UserId: md.UserId,
	}}
	updateReadHistoryOutbox.SetPeer(outboxPeer.To_Peer())
	updateReadHistoryOutbox.SetPts(pts)
	updateReadHistoryOutbox.SetPtsCount(1)
	updateReadHistoryOutbox.SetMaxId(outboxDO.TopMessage)

	updates = mtproto.NewTLUpdates()
	updates.SetSeq(0)
	updates.SetDate(int32(time.Now().Unix()))
	updates.SetUpdates([]*mtproto.Update{updateReadHistoryOutbox.To_Update()})
	delivery.GetDeliveryInstance().DeliveryUpdatesNotMe(
		md.AuthId,
		md.SessionId,
		md.NetlibSessionId,
		[]int32{peer.PeerId},
		updates.To_Updates().Encode())

	glog.Infof("MessagesReadHistory - reply: %s", logger.JsonDebugData(affected))
	return affected.To_Messages_AffectedMessages(), nil
}
