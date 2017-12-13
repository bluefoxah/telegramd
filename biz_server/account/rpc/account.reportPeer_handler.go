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
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

// account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
func (s *AccountServiceImpl) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountReportPeer - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer(request.GetPeer())
	reason := base.FromReportReason(request.GetReason())

	// Insert to db
	do := &dataobject.ReportsDO{}
	do.AuthId = md.AuthId
	do.UserId = md.UserId
	do.PeerType = peer.PeerType
	do.PeerId = peer.PeerId
	do.Reason = int8(reason)
	if reason == base.REASON_OTHER {
		reason := request.GetReason().To_InputReportReasonOther()
		do.Content = reason.GetText()
	}
	dao.GetReportsDAO(dao.DB_MASTER).Insert(do)

	glog.Infof("AccountReportPeer - reply: {true}",)
	return mtproto.ToBool(true), nil
}
