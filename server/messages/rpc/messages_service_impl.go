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
)

type MessagesServiceImpl struct {
}

func (s *MessagesServiceImpl) MessagesSetTyping(ctx context.Context, request *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesHideReportSpam(ctx context.Context, request *mtproto.TLMessagesHideReportSpam) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesDiscardEncryption(ctx context.Context, request *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSaveGif(ctx context.Context, request *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesSetInlineBotResults) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditInlineBotMessage(ctx context.Context, request *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesSetBotCallbackAnswer) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSaveRecentSticker(ctx context.Context, request *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetInlineGameScore(ctx context.Context, request *mtproto.TLMessagesSetInlineGameScore) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesToggleDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReorderPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetBotShippingResults(ctx context.Context, request *mtproto.TLMessagesSetBotShippingResults) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetBotPrecheckoutResults(ctx context.Context, request *mtproto.TLMessagesSetBotPrecheckoutResults) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesFaveSticker(ctx context.Context, request *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesGetMessagesViews(ctx context.Context,  request *mtproto.TLMessagesGetMessagesViews) (*mtproto.Vector<int32T>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetHistory(ctx context.Context, request *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSearchGlobal(ctx context.Context, request *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetUnreadMentions(ctx context.Context, request *mtproto.TLMessagesGetUnreadMentions) (*mtproto.Messages_Messages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesReceivedMessages(ctx context.Context,  request *mtproto.TLMessagesReceivedMessages) (*mtproto.Vector<ReceivedNotifyMessage>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesSendMessage(ctx context.Context, request *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatTitle(ctx context.Context, request *mtproto.TLMessagesEditChatTitle) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesDeleteChatUser(ctx context.Context, request *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesCreateChat(ctx context.Context, request *mtproto.TLMessagesCreateChat) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesForwardMessage(ctx context.Context, request *mtproto.TLMessagesForwardMessage) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesImportChatInvite(ctx context.Context, request *mtproto.TLMessagesImportChatInvite) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesStartBot(ctx context.Context, request *mtproto.TLMessagesStartBot) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesToggleChatAdmins(ctx context.Context, request *mtproto.TLMessagesToggleChatAdmins) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesMigrateChat(ctx context.Context, request *mtproto.TLMessagesMigrateChat) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendInlineBotResult(ctx context.Context, request *mtproto.TLMessagesSendInlineBotResult) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetAllDrafts(ctx context.Context, request *mtproto.TLMessagesGetAllDrafts) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSetGameScore(ctx context.Context, request *mtproto.TLMessagesSetGameScore) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendScreenshotNotification(ctx context.Context, request *mtproto.TLMessagesSendScreenshotNotification) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetPeerSettings(ctx context.Context, request *mtproto.TLMessagesGetPeerSettings) (*mtproto.PeerSettings, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetChats(ctx context.Context, request *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetAllChats(ctx context.Context, request *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetFullChat(ctx context.Context, request *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetDhConfig(ctx context.Context, request *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesRequestEncryption(ctx context.Context, request *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesAcceptEncryption(ctx context.Context, request *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendEncrypted(ctx context.Context, request *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendEncryptedFile(ctx context.Context, request *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSendEncryptedService(ctx context.Context, request *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesReceivedQueue(ctx context.Context,  request *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector<int64T>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesUploadMedia(ctx context.Context, request *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesInstallStickerSet(ctx context.Context, request *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetDocumentByHash(ctx context.Context, request *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesSearchGifs(ctx context.Context, request *mtproto.TLMessagesSearchGifs) (*mtproto.Messages_FoundGifs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesGetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesGetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetRecentStickers(ctx context.Context, request *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *MessagesServiceImpl)MessagesGetAttachedStickers(ctx context.Context,  request *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector<StickerSetCovered>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

func (s *MessagesServiceImpl) MessagesGetGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetInlineGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetInlineGameHighScores) (*mtproto.Messages_HighScores, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *MessagesServiceImpl) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
