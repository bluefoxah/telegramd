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
)

type ChannelsServiceImpl struct {
}

func (s *ChannelsServiceImpl) ChannelsReadHistory(ctx context.Context, request *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsReportSpam(ctx context.Context, request *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsEditAbout(ctx context.Context, request *mtproto.TLChannelsEditAbout) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsCheckUsername(ctx context.Context, request *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsSetStickers(ctx context.Context, request *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsReadMessageContents(ctx context.Context, request *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetMessages(ctx context.Context, request *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsDeleteMessages(ctx context.Context, request *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsDeleteUserHistory(ctx context.Context, request *mtproto.TLChannelsDeleteUserHistory) (*mtproto.Messages_AffectedHistory, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsCreateChannel(ctx context.Context, request *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsEditAdmin(ctx context.Context, request *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsEditTitle(ctx context.Context, request *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsEditPhoto(ctx context.Context, request *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsJoinChannel(ctx context.Context, request *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsLeaveChannel(ctx context.Context, request *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsInviteToChannel(ctx context.Context, request *mtproto.TLChannelsInviteToChannel) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsDeleteChannel(ctx context.Context, request *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsToggleInvites(ctx context.Context, request *mtproto.TLChannelsToggleInvites) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsToggleSignatures(ctx context.Context, request *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsUpdatePinnedMessage(ctx context.Context, request *mtproto.TLChannelsUpdatePinnedMessage) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsEditBanned(ctx context.Context, request *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetChannels(ctx context.Context, request *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetAdminedPublicChannels(ctx context.Context, request *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetFullChannel(ctx context.Context, request *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsExportInvite(ctx context.Context, request *mtproto.TLChannelsExportInvite) (*mtproto.ExportedChatInvite, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetParticipants(ctx context.Context, request *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetParticipant(ctx context.Context, request *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsExportMessageLink(ctx context.Context, request *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *ChannelsServiceImpl) ChannelsGetAdminLog(ctx context.Context, request *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
*/
