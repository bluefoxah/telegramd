/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_encode_decode.py'
 *
 *  Copyright (c) 2017, https://github.com/nebula-im/nebula
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

package mtproto

import (
	"encoding/binary"
	"fmt"
)

type newTLObjectFunc func() TLObject

var registers2 = map[int32]newTLObjectFunc{
	int32(TLConstructor_CRC32_message2):                                         func() TLObject { return new(TLMessage2) },
	int32(TLConstructor_CRC32_msg_container):                                    func() TLObject { return new(TLMsgContainer) },
	int32(TLConstructor_CRC32_msg_copy):                                         func() TLObject { return new(TLMsgCopy) },
	int32(TLConstructor_CRC32_gzip_packed):                                      func() TLObject { return new(TLGzipPacked) },
	int32(TLConstructor_CRC32_resPQ):                                            func() TLObject { return new(TLResPQ) },
	int32(TLConstructor_CRC32_p_q_inner_data):                                   func() TLObject { return new(TLPQInnerData) },
	int32(TLConstructor_CRC32_server_DH_params_fail):                            func() TLObject { return new(TLServer_DHParamsFail) },
	int32(TLConstructor_CRC32_server_DH_params_ok):                              func() TLObject { return new(TLServer_DHParamsOk) },
	int32(TLConstructor_CRC32_server_DH_inner_data):                             func() TLObject { return new(TLServer_DHInnerData) },
	int32(TLConstructor_CRC32_client_DH_inner_data):                             func() TLObject { return new(TLClient_DHInnerData) },
	int32(TLConstructor_CRC32_dh_gen_ok):                                        func() TLObject { return new(TLDhGenOk) },
	int32(TLConstructor_CRC32_dh_gen_retry):                                     func() TLObject { return new(TLDhGenRetry) },
	int32(TLConstructor_CRC32_dh_gen_fail):                                      func() TLObject { return new(TLDhGenFail) },
	int32(TLConstructor_CRC32_destroy_auth_key_ok):                              func() TLObject { return new(TLDestroyAuthKeyOk) },
	int32(TLConstructor_CRC32_destroy_auth_key_none):                            func() TLObject { return new(TLDestroyAuthKeyNone) },
	int32(TLConstructor_CRC32_destroy_auth_key_fail):                            func() TLObject { return new(TLDestroyAuthKeyFail) },
	int32(TLConstructor_CRC32_req_pq):                                           func() TLObject { return new(TLReqPq) },
	int32(TLConstructor_CRC32_req_DH_params):                                    func() TLObject { return new(TLReq_DHParams) },
	int32(TLConstructor_CRC32_set_client_DH_params):                             func() TLObject { return new(TLSetClient_DHParams) },
	int32(TLConstructor_CRC32_destroy_auth_key):                                 func() TLObject { return new(TLDestroyAuthKey) },
	int32(TLConstructor_CRC32_msgs_ack):                                         func() TLObject { return new(TLMsgsAck) },
	int32(TLConstructor_CRC32_bad_msg_notification):                             func() TLObject { return new(TLBadMsgNotification) },
	int32(TLConstructor_CRC32_bad_server_salt):                                  func() TLObject { return new(TLBadServerSalt) },
	int32(TLConstructor_CRC32_msgs_state_req):                                   func() TLObject { return new(TLMsgsStateReq) },
	int32(TLConstructor_CRC32_msgs_state_info):                                  func() TLObject { return new(TLMsgsStateInfo) },
	int32(TLConstructor_CRC32_msgs_all_info):                                    func() TLObject { return new(TLMsgsAllInfo) },
	int32(TLConstructor_CRC32_msg_detailed_info):                                func() TLObject { return new(TLMsgDetailedInfo) },
	int32(TLConstructor_CRC32_msg_new_detailed_info):                            func() TLObject { return new(TLMsgNewDetailedInfo) },
	int32(TLConstructor_CRC32_msg_resend_req):                                   func() TLObject { return new(TLMsgResendReq) },
	int32(TLConstructor_CRC32_rpc_error):                                        func() TLObject { return new(TLRpcError) },
	int32(TLConstructor_CRC32_rpc_answer_unknown):                               func() TLObject { return new(TLRpcAnswerUnknown) },
	int32(TLConstructor_CRC32_rpc_answer_dropped_running):                       func() TLObject { return new(TLRpcAnswerDroppedRunning) },
	int32(TLConstructor_CRC32_rpc_answer_dropped):                               func() TLObject { return new(TLRpcAnswerDropped) },
	int32(TLConstructor_CRC32_future_salt):                                      func() TLObject { return new(TLFutureSalt) },
	int32(TLConstructor_CRC32_future_salts):                                     func() TLObject { return new(TLFutureSalts) },
	int32(TLConstructor_CRC32_pong):                                             func() TLObject { return new(TLPong) },
	int32(TLConstructor_CRC32_destroy_session_ok):                               func() TLObject { return new(TLDestroySessionOk) },
	int32(TLConstructor_CRC32_destroy_session_none):                             func() TLObject { return new(TLDestroySessionNone) },
	int32(TLConstructor_CRC32_new_session_created):                              func() TLObject { return new(TLNewSessionCreated) },
	int32(TLConstructor_CRC32_http_wait):                                        func() TLObject { return new(TLHttpWait) },
	int32(TLConstructor_CRC32_ipPort):                                           func() TLObject { return new(TLIpPort) },
	int32(TLConstructor_CRC32_help_configSimple):                                func() TLObject { return new(TLHelpConfigSimple) },
	int32(TLConstructor_CRC32_rpc_drop_answer):                                  func() TLObject { return new(TLRpcDropAnswer) },
	int32(TLConstructor_CRC32_get_future_salts):                                 func() TLObject { return new(TLGetFutureSalts) },
	int32(TLConstructor_CRC32_ping):                                             func() TLObject { return new(TLPing) },
	int32(TLConstructor_CRC32_ping_delay_disconnect):                            func() TLObject { return new(TLPingDelayDisconnect) },
	int32(TLConstructor_CRC32_destroy_session):                                  func() TLObject { return new(TLDestroySession) },
	int32(TLConstructor_CRC32_contest_saveDeveloperInfo):                        func() TLObject { return new(TLContestSaveDeveloperInfo) },
	int32(TLConstructor_CRC32_boolFalse):                                        func() TLObject { return new(TLBoolFalse) },
	int32(TLConstructor_CRC32_boolTrue):                                         func() TLObject { return new(TLBoolTrue) },
	int32(TLConstructor_CRC32_true):                                             func() TLObject { return new(TLTrue) },
	int32(TLConstructor_CRC32_error):                                            func() TLObject { return new(TLError) },
	int32(TLConstructor_CRC32_null):                                             func() TLObject { return new(TLNull) },
	int32(TLConstructor_CRC32_inputPeerEmpty):                                   func() TLObject { return new(TLInputPeerEmpty) },
	int32(TLConstructor_CRC32_inputPeerSelf):                                    func() TLObject { return new(TLInputPeerSelf) },
	int32(TLConstructor_CRC32_inputPeerChat):                                    func() TLObject { return new(TLInputPeerChat) },
	int32(TLConstructor_CRC32_inputPeerUser):                                    func() TLObject { return new(TLInputPeerUser) },
	int32(TLConstructor_CRC32_inputPeerChannel):                                 func() TLObject { return new(TLInputPeerChannel) },
	int32(TLConstructor_CRC32_inputUserEmpty):                                   func() TLObject { return new(TLInputUserEmpty) },
	int32(TLConstructor_CRC32_inputUserSelf):                                    func() TLObject { return new(TLInputUserSelf) },
	int32(TLConstructor_CRC32_inputUser):                                        func() TLObject { return new(TLInputUser) },
	int32(TLConstructor_CRC32_inputPhoneContact):                                func() TLObject { return new(TLInputPhoneContact) },
	int32(TLConstructor_CRC32_inputFile):                                        func() TLObject { return new(TLInputFile) },
	int32(TLConstructor_CRC32_inputFileBig):                                     func() TLObject { return new(TLInputFileBig) },
	int32(TLConstructor_CRC32_inputMediaEmpty):                                  func() TLObject { return new(TLInputMediaEmpty) },
	int32(TLConstructor_CRC32_inputMediaUploadedPhoto):                          func() TLObject { return new(TLInputMediaUploadedPhoto) },
	int32(TLConstructor_CRC32_inputMediaPhoto):                                  func() TLObject { return new(TLInputMediaPhoto) },
	int32(TLConstructor_CRC32_inputMediaGeoPoint):                               func() TLObject { return new(TLInputMediaGeoPoint) },
	int32(TLConstructor_CRC32_inputMediaContact):                                func() TLObject { return new(TLInputMediaContact) },
	int32(TLConstructor_CRC32_inputMediaUploadedDocument):                       func() TLObject { return new(TLInputMediaUploadedDocument) },
	int32(TLConstructor_CRC32_inputMediaDocument):                               func() TLObject { return new(TLInputMediaDocument) },
	int32(TLConstructor_CRC32_inputMediaVenue):                                  func() TLObject { return new(TLInputMediaVenue) },
	int32(TLConstructor_CRC32_inputMediaGifExternal):                            func() TLObject { return new(TLInputMediaGifExternal) },
	int32(TLConstructor_CRC32_inputMediaPhotoExternal):                          func() TLObject { return new(TLInputMediaPhotoExternal) },
	int32(TLConstructor_CRC32_inputMediaDocumentExternal):                       func() TLObject { return new(TLInputMediaDocumentExternal) },
	int32(TLConstructor_CRC32_inputMediaGame):                                   func() TLObject { return new(TLInputMediaGame) },
	int32(TLConstructor_CRC32_inputMediaInvoice):                                func() TLObject { return new(TLInputMediaInvoice) },
	int32(TLConstructor_CRC32_inputChatPhotoEmpty):                              func() TLObject { return new(TLInputChatPhotoEmpty) },
	int32(TLConstructor_CRC32_inputChatUploadedPhoto):                           func() TLObject { return new(TLInputChatUploadedPhoto) },
	int32(TLConstructor_CRC32_inputChatPhoto):                                   func() TLObject { return new(TLInputChatPhoto) },
	int32(TLConstructor_CRC32_inputGeoPointEmpty):                               func() TLObject { return new(TLInputGeoPointEmpty) },
	int32(TLConstructor_CRC32_inputGeoPoint):                                    func() TLObject { return new(TLInputGeoPoint) },
	int32(TLConstructor_CRC32_inputPhotoEmpty):                                  func() TLObject { return new(TLInputPhotoEmpty) },
	int32(TLConstructor_CRC32_inputPhoto):                                       func() TLObject { return new(TLInputPhoto) },
	int32(TLConstructor_CRC32_inputFileLocation):                                func() TLObject { return new(TLInputFileLocation) },
	int32(TLConstructor_CRC32_inputEncryptedFileLocation):                       func() TLObject { return new(TLInputEncryptedFileLocation) },
	int32(TLConstructor_CRC32_inputDocumentFileLocation):                        func() TLObject { return new(TLInputDocumentFileLocation) },
	int32(TLConstructor_CRC32_inputAppEvent):                                    func() TLObject { return new(TLInputAppEvent) },
	int32(TLConstructor_CRC32_peerUser):                                         func() TLObject { return new(TLPeerUser) },
	int32(TLConstructor_CRC32_peerChat):                                         func() TLObject { return new(TLPeerChat) },
	int32(TLConstructor_CRC32_peerChannel):                                      func() TLObject { return new(TLPeerChannel) },
	int32(TLConstructor_CRC32_storage_fileUnknown):                              func() TLObject { return new(TLStorageFileUnknown) },
	int32(TLConstructor_CRC32_storage_filePartial):                              func() TLObject { return new(TLStorageFilePartial) },
	int32(TLConstructor_CRC32_storage_fileJpeg):                                 func() TLObject { return new(TLStorageFileJpeg) },
	int32(TLConstructor_CRC32_storage_fileGif):                                  func() TLObject { return new(TLStorageFileGif) },
	int32(TLConstructor_CRC32_storage_filePng):                                  func() TLObject { return new(TLStorageFilePng) },
	int32(TLConstructor_CRC32_storage_filePdf):                                  func() TLObject { return new(TLStorageFilePdf) },
	int32(TLConstructor_CRC32_storage_fileMp3):                                  func() TLObject { return new(TLStorageFileMp3) },
	int32(TLConstructor_CRC32_storage_fileMov):                                  func() TLObject { return new(TLStorageFileMov) },
	int32(TLConstructor_CRC32_storage_fileMp4):                                  func() TLObject { return new(TLStorageFileMp4) },
	int32(TLConstructor_CRC32_storage_fileWebp):                                 func() TLObject { return new(TLStorageFileWebp) },
	int32(TLConstructor_CRC32_fileLocationUnavailable):                          func() TLObject { return new(TLFileLocationUnavailable) },
	int32(TLConstructor_CRC32_fileLocation):                                     func() TLObject { return new(TLFileLocation) },
	int32(TLConstructor_CRC32_userEmpty):                                        func() TLObject { return new(TLUserEmpty) },
	int32(TLConstructor_CRC32_user):                                             func() TLObject { return new(TLUser) },
	int32(TLConstructor_CRC32_userProfilePhotoEmpty):                            func() TLObject { return new(TLUserProfilePhotoEmpty) },
	int32(TLConstructor_CRC32_userProfilePhoto):                                 func() TLObject { return new(TLUserProfilePhoto) },
	int32(TLConstructor_CRC32_userStatusEmpty):                                  func() TLObject { return new(TLUserStatusEmpty) },
	int32(TLConstructor_CRC32_userStatusOnline):                                 func() TLObject { return new(TLUserStatusOnline) },
	int32(TLConstructor_CRC32_userStatusOffline):                                func() TLObject { return new(TLUserStatusOffline) },
	int32(TLConstructor_CRC32_userStatusRecently):                               func() TLObject { return new(TLUserStatusRecently) },
	int32(TLConstructor_CRC32_userStatusLastWeek):                               func() TLObject { return new(TLUserStatusLastWeek) },
	int32(TLConstructor_CRC32_userStatusLastMonth):                              func() TLObject { return new(TLUserStatusLastMonth) },
	int32(TLConstructor_CRC32_chatEmpty):                                        func() TLObject { return new(TLChatEmpty) },
	int32(TLConstructor_CRC32_chat):                                             func() TLObject { return new(TLChat) },
	int32(TLConstructor_CRC32_chatForbidden):                                    func() TLObject { return new(TLChatForbidden) },
	int32(TLConstructor_CRC32_channel):                                          func() TLObject { return new(TLChannel) },
	int32(TLConstructor_CRC32_channelForbidden):                                 func() TLObject { return new(TLChannelForbidden) },
	int32(TLConstructor_CRC32_chatFull):                                         func() TLObject { return new(TLChatFull) },
	int32(TLConstructor_CRC32_channelFull):                                      func() TLObject { return new(TLChannelFull) },
	int32(TLConstructor_CRC32_chatParticipant):                                  func() TLObject { return new(TLChatParticipant) },
	int32(TLConstructor_CRC32_chatParticipantCreator):                           func() TLObject { return new(TLChatParticipantCreator) },
	int32(TLConstructor_CRC32_chatParticipantAdmin):                             func() TLObject { return new(TLChatParticipantAdmin) },
	int32(TLConstructor_CRC32_chatParticipantsForbidden):                        func() TLObject { return new(TLChatParticipantsForbidden) },
	int32(TLConstructor_CRC32_chatParticipants):                                 func() TLObject { return new(TLChatParticipants) },
	int32(TLConstructor_CRC32_chatPhotoEmpty):                                   func() TLObject { return new(TLChatPhotoEmpty) },
	int32(TLConstructor_CRC32_chatPhoto):                                        func() TLObject { return new(TLChatPhoto) },
	int32(TLConstructor_CRC32_messageEmpty):                                     func() TLObject { return new(TLMessageEmpty) },
	int32(TLConstructor_CRC32_message):                                          func() TLObject { return new(TLMessage) },
	int32(TLConstructor_CRC32_messageService):                                   func() TLObject { return new(TLMessageService) },
	int32(TLConstructor_CRC32_messageMediaEmpty):                                func() TLObject { return new(TLMessageMediaEmpty) },
	int32(TLConstructor_CRC32_messageMediaPhoto):                                func() TLObject { return new(TLMessageMediaPhoto) },
	int32(TLConstructor_CRC32_messageMediaGeo):                                  func() TLObject { return new(TLMessageMediaGeo) },
	int32(TLConstructor_CRC32_messageMediaContact):                              func() TLObject { return new(TLMessageMediaContact) },
	int32(TLConstructor_CRC32_messageMediaUnsupported):                          func() TLObject { return new(TLMessageMediaUnsupported) },
	int32(TLConstructor_CRC32_messageMediaDocument):                             func() TLObject { return new(TLMessageMediaDocument) },
	int32(TLConstructor_CRC32_messageMediaWebPage):                              func() TLObject { return new(TLMessageMediaWebPage) },
	int32(TLConstructor_CRC32_messageMediaVenue):                                func() TLObject { return new(TLMessageMediaVenue) },
	int32(TLConstructor_CRC32_messageMediaGame):                                 func() TLObject { return new(TLMessageMediaGame) },
	int32(TLConstructor_CRC32_messageMediaInvoice):                              func() TLObject { return new(TLMessageMediaInvoice) },
	int32(TLConstructor_CRC32_messageActionEmpty):                               func() TLObject { return new(TLMessageActionEmpty) },
	int32(TLConstructor_CRC32_messageActionChatCreate):                          func() TLObject { return new(TLMessageActionChatCreate) },
	int32(TLConstructor_CRC32_messageActionChatEditTitle):                       func() TLObject { return new(TLMessageActionChatEditTitle) },
	int32(TLConstructor_CRC32_messageActionChatEditPhoto):                       func() TLObject { return new(TLMessageActionChatEditPhoto) },
	int32(TLConstructor_CRC32_messageActionChatDeletePhoto):                     func() TLObject { return new(TLMessageActionChatDeletePhoto) },
	int32(TLConstructor_CRC32_messageActionChatAddUser):                         func() TLObject { return new(TLMessageActionChatAddUser) },
	int32(TLConstructor_CRC32_messageActionChatDeleteUser):                      func() TLObject { return new(TLMessageActionChatDeleteUser) },
	int32(TLConstructor_CRC32_messageActionChatJoinedByLink):                    func() TLObject { return new(TLMessageActionChatJoinedByLink) },
	int32(TLConstructor_CRC32_messageActionChannelCreate):                       func() TLObject { return new(TLMessageActionChannelCreate) },
	int32(TLConstructor_CRC32_messageActionChatMigrateTo):                       func() TLObject { return new(TLMessageActionChatMigrateTo) },
	int32(TLConstructor_CRC32_messageActionChannelMigrateFrom):                  func() TLObject { return new(TLMessageActionChannelMigrateFrom) },
	int32(TLConstructor_CRC32_messageActionPinMessage):                          func() TLObject { return new(TLMessageActionPinMessage) },
	int32(TLConstructor_CRC32_messageActionHistoryClear):                        func() TLObject { return new(TLMessageActionHistoryClear) },
	int32(TLConstructor_CRC32_messageActionGameScore):                           func() TLObject { return new(TLMessageActionGameScore) },
	int32(TLConstructor_CRC32_messageActionPaymentSentMe):                       func() TLObject { return new(TLMessageActionPaymentSentMe) },
	int32(TLConstructor_CRC32_messageActionPaymentSent):                         func() TLObject { return new(TLMessageActionPaymentSent) },
	int32(TLConstructor_CRC32_messageActionPhoneCall):                           func() TLObject { return new(TLMessageActionPhoneCall) },
	int32(TLConstructor_CRC32_messageActionScreenshotTaken):                     func() TLObject { return new(TLMessageActionScreenshotTaken) },
	int32(TLConstructor_CRC32_dialog):                                           func() TLObject { return new(TLDialog) },
	int32(TLConstructor_CRC32_photoEmpty):                                       func() TLObject { return new(TLPhotoEmpty) },
	int32(TLConstructor_CRC32_photo):                                            func() TLObject { return new(TLPhoto) },
	int32(TLConstructor_CRC32_photoSizeEmpty):                                   func() TLObject { return new(TLPhotoSizeEmpty) },
	int32(TLConstructor_CRC32_photoSize):                                        func() TLObject { return new(TLPhotoSize) },
	int32(TLConstructor_CRC32_photoCachedSize):                                  func() TLObject { return new(TLPhotoCachedSize) },
	int32(TLConstructor_CRC32_geoPointEmpty):                                    func() TLObject { return new(TLGeoPointEmpty) },
	int32(TLConstructor_CRC32_geoPoint):                                         func() TLObject { return new(TLGeoPoint) },
	int32(TLConstructor_CRC32_auth_checkedPhone):                                func() TLObject { return new(TLAuthCheckedPhone) },
	int32(TLConstructor_CRC32_auth_sentCode):                                    func() TLObject { return new(TLAuthSentCode) },
	int32(TLConstructor_CRC32_auth_authorization):                               func() TLObject { return new(TLAuthAuthorization) },
	int32(TLConstructor_CRC32_auth_exportedAuthorization):                       func() TLObject { return new(TLAuthExportedAuthorization) },
	int32(TLConstructor_CRC32_inputNotifyPeer):                                  func() TLObject { return new(TLInputNotifyPeer) },
	int32(TLConstructor_CRC32_inputNotifyUsers):                                 func() TLObject { return new(TLInputNotifyUsers) },
	int32(TLConstructor_CRC32_inputNotifyChats):                                 func() TLObject { return new(TLInputNotifyChats) },
	int32(TLConstructor_CRC32_inputNotifyAll):                                   func() TLObject { return new(TLInputNotifyAll) },
	int32(TLConstructor_CRC32_inputPeerNotifyEventsEmpty):                       func() TLObject { return new(TLInputPeerNotifyEventsEmpty) },
	int32(TLConstructor_CRC32_inputPeerNotifyEventsAll):                         func() TLObject { return new(TLInputPeerNotifyEventsAll) },
	int32(TLConstructor_CRC32_inputPeerNotifySettings):                          func() TLObject { return new(TLInputPeerNotifySettings) },
	int32(TLConstructor_CRC32_peerNotifyEventsEmpty):                            func() TLObject { return new(TLPeerNotifyEventsEmpty) },
	int32(TLConstructor_CRC32_peerNotifyEventsAll):                              func() TLObject { return new(TLPeerNotifyEventsAll) },
	int32(TLConstructor_CRC32_peerNotifySettingsEmpty):                          func() TLObject { return new(TLPeerNotifySettingsEmpty) },
	int32(TLConstructor_CRC32_peerNotifySettings):                               func() TLObject { return new(TLPeerNotifySettings) },
	int32(TLConstructor_CRC32_peerSettings):                                     func() TLObject { return new(TLPeerSettings) },
	int32(TLConstructor_CRC32_wallPaper):                                        func() TLObject { return new(TLWallPaper) },
	int32(TLConstructor_CRC32_wallPaperSolid):                                   func() TLObject { return new(TLWallPaperSolid) },
	int32(TLConstructor_CRC32_inputReportReasonSpam):                            func() TLObject { return new(TLInputReportReasonSpam) },
	int32(TLConstructor_CRC32_inputReportReasonViolence):                        func() TLObject { return new(TLInputReportReasonViolence) },
	int32(TLConstructor_CRC32_inputReportReasonPornography):                     func() TLObject { return new(TLInputReportReasonPornography) },
	int32(TLConstructor_CRC32_inputReportReasonOther):                           func() TLObject { return new(TLInputReportReasonOther) },
	int32(TLConstructor_CRC32_userFull):                                         func() TLObject { return new(TLUserFull) },
	int32(TLConstructor_CRC32_contact):                                          func() TLObject { return new(TLContact) },
	int32(TLConstructor_CRC32_importedContact):                                  func() TLObject { return new(TLImportedContact) },
	int32(TLConstructor_CRC32_contactBlocked):                                   func() TLObject { return new(TLContactBlocked) },
	int32(TLConstructor_CRC32_contactStatus):                                    func() TLObject { return new(TLContactStatus) },
	int32(TLConstructor_CRC32_contacts_link):                                    func() TLObject { return new(TLContactsLink) },
	int32(TLConstructor_CRC32_contacts_contactsNotModified):                     func() TLObject { return new(TLContactsContactsNotModified) },
	int32(TLConstructor_CRC32_contacts_contacts):                                func() TLObject { return new(TLContactsContacts) },
	int32(TLConstructor_CRC32_contacts_importedContacts):                        func() TLObject { return new(TLContactsImportedContacts) },
	int32(TLConstructor_CRC32_contacts_blocked):                                 func() TLObject { return new(TLContactsBlocked) },
	int32(TLConstructor_CRC32_contacts_blockedSlice):                            func() TLObject { return new(TLContactsBlockedSlice) },
	int32(TLConstructor_CRC32_messages_dialogs):                                 func() TLObject { return new(TLMessagesDialogs) },
	int32(TLConstructor_CRC32_messages_dialogsSlice):                            func() TLObject { return new(TLMessagesDialogsSlice) },
	int32(TLConstructor_CRC32_messages_messages):                                func() TLObject { return new(TLMessagesMessages) },
	int32(TLConstructor_CRC32_messages_messagesSlice):                           func() TLObject { return new(TLMessagesMessagesSlice) },
	int32(TLConstructor_CRC32_messages_channelMessages):                         func() TLObject { return new(TLMessagesChannelMessages) },
	int32(TLConstructor_CRC32_messages_chats):                                   func() TLObject { return new(TLMessagesChats) },
	int32(TLConstructor_CRC32_messages_chatsSlice):                              func() TLObject { return new(TLMessagesChatsSlice) },
	int32(TLConstructor_CRC32_messages_chatFull):                                func() TLObject { return new(TLMessagesChatFull) },
	int32(TLConstructor_CRC32_messages_affectedHistory):                         func() TLObject { return new(TLMessagesAffectedHistory) },
	int32(TLConstructor_CRC32_inputMessagesFilterEmpty):                         func() TLObject { return new(TLInputMessagesFilterEmpty) },
	int32(TLConstructor_CRC32_inputMessagesFilterPhotos):                        func() TLObject { return new(TLInputMessagesFilterPhotos) },
	int32(TLConstructor_CRC32_inputMessagesFilterVideo):                         func() TLObject { return new(TLInputMessagesFilterVideo) },
	int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideo):                    func() TLObject { return new(TLInputMessagesFilterPhotoVideo) },
	int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideoDocuments):           func() TLObject { return new(TLInputMessagesFilterPhotoVideoDocuments) },
	int32(TLConstructor_CRC32_inputMessagesFilterDocument):                      func() TLObject { return new(TLInputMessagesFilterDocument) },
	int32(TLConstructor_CRC32_inputMessagesFilterUrl):                           func() TLObject { return new(TLInputMessagesFilterUrl) },
	int32(TLConstructor_CRC32_inputMessagesFilterGif):                           func() TLObject { return new(TLInputMessagesFilterGif) },
	int32(TLConstructor_CRC32_inputMessagesFilterVoice):                         func() TLObject { return new(TLInputMessagesFilterVoice) },
	int32(TLConstructor_CRC32_inputMessagesFilterMusic):                         func() TLObject { return new(TLInputMessagesFilterMusic) },
	int32(TLConstructor_CRC32_inputMessagesFilterChatPhotos):                    func() TLObject { return new(TLInputMessagesFilterChatPhotos) },
	int32(TLConstructor_CRC32_inputMessagesFilterPhoneCalls):                    func() TLObject { return new(TLInputMessagesFilterPhoneCalls) },
	int32(TLConstructor_CRC32_inputMessagesFilterRoundVoice):                    func() TLObject { return new(TLInputMessagesFilterRoundVoice) },
	int32(TLConstructor_CRC32_inputMessagesFilterRoundVideo):                    func() TLObject { return new(TLInputMessagesFilterRoundVideo) },
	int32(TLConstructor_CRC32_inputMessagesFilterMyMentions):                    func() TLObject { return new(TLInputMessagesFilterMyMentions) },
	int32(TLConstructor_CRC32_updateNewMessage):                                 func() TLObject { return new(TLUpdateNewMessage) },
	int32(TLConstructor_CRC32_updateMessageID):                                  func() TLObject { return new(TLUpdateMessageID) },
	int32(TLConstructor_CRC32_updateDeleteMessages):                             func() TLObject { return new(TLUpdateDeleteMessages) },
	int32(TLConstructor_CRC32_updateUserTyping):                                 func() TLObject { return new(TLUpdateUserTyping) },
	int32(TLConstructor_CRC32_updateChatUserTyping):                             func() TLObject { return new(TLUpdateChatUserTyping) },
	int32(TLConstructor_CRC32_updateChatParticipants):                           func() TLObject { return new(TLUpdateChatParticipants) },
	int32(TLConstructor_CRC32_updateUserStatus):                                 func() TLObject { return new(TLUpdateUserStatus) },
	int32(TLConstructor_CRC32_updateUserName):                                   func() TLObject { return new(TLUpdateUserName) },
	int32(TLConstructor_CRC32_updateUserPhoto):                                  func() TLObject { return new(TLUpdateUserPhoto) },
	int32(TLConstructor_CRC32_updateContactRegistered):                          func() TLObject { return new(TLUpdateContactRegistered) },
	int32(TLConstructor_CRC32_updateContactLink):                                func() TLObject { return new(TLUpdateContactLink) },
	int32(TLConstructor_CRC32_updateNewEncryptedMessage):                        func() TLObject { return new(TLUpdateNewEncryptedMessage) },
	int32(TLConstructor_CRC32_updateEncryptedChatTyping):                        func() TLObject { return new(TLUpdateEncryptedChatTyping) },
	int32(TLConstructor_CRC32_updateEncryption):                                 func() TLObject { return new(TLUpdateEncryption) },
	int32(TLConstructor_CRC32_updateEncryptedMessagesRead):                      func() TLObject { return new(TLUpdateEncryptedMessagesRead) },
	int32(TLConstructor_CRC32_updateChatParticipantAdd):                         func() TLObject { return new(TLUpdateChatParticipantAdd) },
	int32(TLConstructor_CRC32_updateChatParticipantDelete):                      func() TLObject { return new(TLUpdateChatParticipantDelete) },
	int32(TLConstructor_CRC32_updateDcOptions):                                  func() TLObject { return new(TLUpdateDcOptions) },
	int32(TLConstructor_CRC32_updateUserBlocked):                                func() TLObject { return new(TLUpdateUserBlocked) },
	int32(TLConstructor_CRC32_updateNotifySettings):                             func() TLObject { return new(TLUpdateNotifySettings) },
	int32(TLConstructor_CRC32_updateServiceNotification):                        func() TLObject { return new(TLUpdateServiceNotification) },
	int32(TLConstructor_CRC32_updatePrivacy):                                    func() TLObject { return new(TLUpdatePrivacy) },
	int32(TLConstructor_CRC32_updateUserPhone):                                  func() TLObject { return new(TLUpdateUserPhone) },
	int32(TLConstructor_CRC32_updateReadHistoryInbox):                           func() TLObject { return new(TLUpdateReadHistoryInbox) },
	int32(TLConstructor_CRC32_updateReadHistoryOutbox):                          func() TLObject { return new(TLUpdateReadHistoryOutbox) },
	int32(TLConstructor_CRC32_updateWebPage):                                    func() TLObject { return new(TLUpdateWebPage) },
	int32(TLConstructor_CRC32_updateReadMessagesContents):                       func() TLObject { return new(TLUpdateReadMessagesContents) },
	int32(TLConstructor_CRC32_updateChannelTooLong):                             func() TLObject { return new(TLUpdateChannelTooLong) },
	int32(TLConstructor_CRC32_updateChannel):                                    func() TLObject { return new(TLUpdateChannel) },
	int32(TLConstructor_CRC32_updateNewChannelMessage):                          func() TLObject { return new(TLUpdateNewChannelMessage) },
	int32(TLConstructor_CRC32_updateReadChannelInbox):                           func() TLObject { return new(TLUpdateReadChannelInbox) },
	int32(TLConstructor_CRC32_updateDeleteChannelMessages):                      func() TLObject { return new(TLUpdateDeleteChannelMessages) },
	int32(TLConstructor_CRC32_updateChannelMessageViews):                        func() TLObject { return new(TLUpdateChannelMessageViews) },
	int32(TLConstructor_CRC32_updateChatAdmins):                                 func() TLObject { return new(TLUpdateChatAdmins) },
	int32(TLConstructor_CRC32_updateChatParticipantAdmin):                       func() TLObject { return new(TLUpdateChatParticipantAdmin) },
	int32(TLConstructor_CRC32_updateNewStickerSet):                              func() TLObject { return new(TLUpdateNewStickerSet) },
	int32(TLConstructor_CRC32_updateStickerSetsOrder):                           func() TLObject { return new(TLUpdateStickerSetsOrder) },
	int32(TLConstructor_CRC32_updateStickerSets):                                func() TLObject { return new(TLUpdateStickerSets) },
	int32(TLConstructor_CRC32_updateSavedGifs):                                  func() TLObject { return new(TLUpdateSavedGifs) },
	int32(TLConstructor_CRC32_updateBotInlineQuery):                             func() TLObject { return new(TLUpdateBotInlineQuery) },
	int32(TLConstructor_CRC32_updateBotInlineSend):                              func() TLObject { return new(TLUpdateBotInlineSend) },
	int32(TLConstructor_CRC32_updateEditChannelMessage):                         func() TLObject { return new(TLUpdateEditChannelMessage) },
	int32(TLConstructor_CRC32_updateChannelPinnedMessage):                       func() TLObject { return new(TLUpdateChannelPinnedMessage) },
	int32(TLConstructor_CRC32_updateBotCallbackQuery):                           func() TLObject { return new(TLUpdateBotCallbackQuery) },
	int32(TLConstructor_CRC32_updateEditMessage):                                func() TLObject { return new(TLUpdateEditMessage) },
	int32(TLConstructor_CRC32_updateInlineBotCallbackQuery):                     func() TLObject { return new(TLUpdateInlineBotCallbackQuery) },
	int32(TLConstructor_CRC32_updateReadChannelOutbox):                          func() TLObject { return new(TLUpdateReadChannelOutbox) },
	int32(TLConstructor_CRC32_updateDraftMessage):                               func() TLObject { return new(TLUpdateDraftMessage) },
	int32(TLConstructor_CRC32_updateReadFeaturedStickers):                       func() TLObject { return new(TLUpdateReadFeaturedStickers) },
	int32(TLConstructor_CRC32_updateRecentStickers):                             func() TLObject { return new(TLUpdateRecentStickers) },
	int32(TLConstructor_CRC32_updateConfig):                                     func() TLObject { return new(TLUpdateConfig) },
	int32(TLConstructor_CRC32_updatePtsChanged):                                 func() TLObject { return new(TLUpdatePtsChanged) },
	int32(TLConstructor_CRC32_updateChannelWebPage):                             func() TLObject { return new(TLUpdateChannelWebPage) },
	int32(TLConstructor_CRC32_updateDialogPinned):                               func() TLObject { return new(TLUpdateDialogPinned) },
	int32(TLConstructor_CRC32_updatePinnedDialogs):                              func() TLObject { return new(TLUpdatePinnedDialogs) },
	int32(TLConstructor_CRC32_updateBotWebhookJSON):                             func() TLObject { return new(TLUpdateBotWebhookJSON) },
	int32(TLConstructor_CRC32_updateBotWebhookJSONQuery):                        func() TLObject { return new(TLUpdateBotWebhookJSONQuery) },
	int32(TLConstructor_CRC32_updateBotShippingQuery):                           func() TLObject { return new(TLUpdateBotShippingQuery) },
	int32(TLConstructor_CRC32_updateBotPrecheckoutQuery):                        func() TLObject { return new(TLUpdateBotPrecheckoutQuery) },
	int32(TLConstructor_CRC32_updatePhoneCall):                                  func() TLObject { return new(TLUpdatePhoneCall) },
	int32(TLConstructor_CRC32_updateLangPackTooLong):                            func() TLObject { return new(TLUpdateLangPackTooLong) },
	int32(TLConstructor_CRC32_updateLangPack):                                   func() TLObject { return new(TLUpdateLangPack) },
	int32(TLConstructor_CRC32_updateFavedStickers):                              func() TLObject { return new(TLUpdateFavedStickers) },
	int32(TLConstructor_CRC32_updateChannelReadMessagesContents):                func() TLObject { return new(TLUpdateChannelReadMessagesContents) },
	int32(TLConstructor_CRC32_updateContactsReset):                              func() TLObject { return new(TLUpdateContactsReset) },
	int32(TLConstructor_CRC32_updates_state):                                    func() TLObject { return new(TLUpdatesState) },
	int32(TLConstructor_CRC32_updates_differenceEmpty):                          func() TLObject { return new(TLUpdatesDifferenceEmpty) },
	int32(TLConstructor_CRC32_updates_difference):                               func() TLObject { return new(TLUpdatesDifference) },
	int32(TLConstructor_CRC32_updates_differenceSlice):                          func() TLObject { return new(TLUpdatesDifferenceSlice) },
	int32(TLConstructor_CRC32_updates_differenceTooLong):                        func() TLObject { return new(TLUpdatesDifferenceTooLong) },
	int32(TLConstructor_CRC32_updatesTooLong):                                   func() TLObject { return new(TLUpdatesTooLong) },
	int32(TLConstructor_CRC32_updateShortMessage):                               func() TLObject { return new(TLUpdateShortMessage) },
	int32(TLConstructor_CRC32_updateShortChatMessage):                           func() TLObject { return new(TLUpdateShortChatMessage) },
	int32(TLConstructor_CRC32_updateShort):                                      func() TLObject { return new(TLUpdateShort) },
	int32(TLConstructor_CRC32_updatesCombined):                                  func() TLObject { return new(TLUpdatesCombined) },
	int32(TLConstructor_CRC32_updates):                                          func() TLObject { return new(TLUpdates) },
	int32(TLConstructor_CRC32_updateShortSentMessage):                           func() TLObject { return new(TLUpdateShortSentMessage) },
	int32(TLConstructor_CRC32_photos_photos):                                    func() TLObject { return new(TLPhotosPhotos) },
	int32(TLConstructor_CRC32_photos_photosSlice):                               func() TLObject { return new(TLPhotosPhotosSlice) },
	int32(TLConstructor_CRC32_photos_photo):                                     func() TLObject { return new(TLPhotosPhoto) },
	int32(TLConstructor_CRC32_upload_file):                                      func() TLObject { return new(TLUploadFile) },
	int32(TLConstructor_CRC32_upload_fileCdnRedirect):                           func() TLObject { return new(TLUploadFileCdnRedirect) },
	int32(TLConstructor_CRC32_dcOption):                                         func() TLObject { return new(TLDcOption) },
	int32(TLConstructor_CRC32_config):                                           func() TLObject { return new(TLConfig) },
	int32(TLConstructor_CRC32_nearestDc):                                        func() TLObject { return new(TLNearestDc) },
	int32(TLConstructor_CRC32_help_appUpdate):                                   func() TLObject { return new(TLHelpAppUpdate) },
	int32(TLConstructor_CRC32_help_noAppUpdate):                                 func() TLObject { return new(TLHelpNoAppUpdate) },
	int32(TLConstructor_CRC32_help_inviteText):                                  func() TLObject { return new(TLHelpInviteText) },
	int32(TLConstructor_CRC32_encryptedChatEmpty):                               func() TLObject { return new(TLEncryptedChatEmpty) },
	int32(TLConstructor_CRC32_encryptedChatWaiting):                             func() TLObject { return new(TLEncryptedChatWaiting) },
	int32(TLConstructor_CRC32_encryptedChatRequested):                           func() TLObject { return new(TLEncryptedChatRequested) },
	int32(TLConstructor_CRC32_encryptedChat):                                    func() TLObject { return new(TLEncryptedChat) },
	int32(TLConstructor_CRC32_encryptedChatDiscarded):                           func() TLObject { return new(TLEncryptedChatDiscarded) },
	int32(TLConstructor_CRC32_inputEncryptedChat):                               func() TLObject { return new(TLInputEncryptedChat) },
	int32(TLConstructor_CRC32_encryptedFileEmpty):                               func() TLObject { return new(TLEncryptedFileEmpty) },
	int32(TLConstructor_CRC32_encryptedFile):                                    func() TLObject { return new(TLEncryptedFile) },
	int32(TLConstructor_CRC32_inputEncryptedFileEmpty):                          func() TLObject { return new(TLInputEncryptedFileEmpty) },
	int32(TLConstructor_CRC32_inputEncryptedFileUploaded):                       func() TLObject { return new(TLInputEncryptedFileUploaded) },
	int32(TLConstructor_CRC32_inputEncryptedFile):                               func() TLObject { return new(TLInputEncryptedFile) },
	int32(TLConstructor_CRC32_inputEncryptedFileBigUploaded):                    func() TLObject { return new(TLInputEncryptedFileBigUploaded) },
	int32(TLConstructor_CRC32_encryptedMessage):                                 func() TLObject { return new(TLEncryptedMessage) },
	int32(TLConstructor_CRC32_encryptedMessageService):                          func() TLObject { return new(TLEncryptedMessageService) },
	int32(TLConstructor_CRC32_messages_dhConfigNotModified):                     func() TLObject { return new(TLMessagesDhConfigNotModified) },
	int32(TLConstructor_CRC32_messages_dhConfig):                                func() TLObject { return new(TLMessagesDhConfig) },
	int32(TLConstructor_CRC32_messages_sentEncryptedMessage):                    func() TLObject { return new(TLMessagesSentEncryptedMessage) },
	int32(TLConstructor_CRC32_messages_sentEncryptedFile):                       func() TLObject { return new(TLMessagesSentEncryptedFile) },
	int32(TLConstructor_CRC32_inputDocumentEmpty):                               func() TLObject { return new(TLInputDocumentEmpty) },
	int32(TLConstructor_CRC32_inputDocument):                                    func() TLObject { return new(TLInputDocument) },
	int32(TLConstructor_CRC32_documentEmpty):                                    func() TLObject { return new(TLDocumentEmpty) },
	int32(TLConstructor_CRC32_document):                                         func() TLObject { return new(TLDocument) },
	int32(TLConstructor_CRC32_help_support):                                     func() TLObject { return new(TLHelpSupport) },
	int32(TLConstructor_CRC32_notifyPeer):                                       func() TLObject { return new(TLNotifyPeer) },
	int32(TLConstructor_CRC32_notifyUsers):                                      func() TLObject { return new(TLNotifyUsers) },
	int32(TLConstructor_CRC32_notifyChats):                                      func() TLObject { return new(TLNotifyChats) },
	int32(TLConstructor_CRC32_notifyAll):                                        func() TLObject { return new(TLNotifyAll) },
	int32(TLConstructor_CRC32_sendMessageTypingAction):                          func() TLObject { return new(TLSendMessageTypingAction) },
	int32(TLConstructor_CRC32_sendMessageCancelAction):                          func() TLObject { return new(TLSendMessageCancelAction) },
	int32(TLConstructor_CRC32_sendMessageRecordVideoAction):                     func() TLObject { return new(TLSendMessageRecordVideoAction) },
	int32(TLConstructor_CRC32_sendMessageUploadVideoAction):                     func() TLObject { return new(TLSendMessageUploadVideoAction) },
	int32(TLConstructor_CRC32_sendMessageRecordAudioAction):                     func() TLObject { return new(TLSendMessageRecordAudioAction) },
	int32(TLConstructor_CRC32_sendMessageUploadAudioAction):                     func() TLObject { return new(TLSendMessageUploadAudioAction) },
	int32(TLConstructor_CRC32_sendMessageUploadPhotoAction):                     func() TLObject { return new(TLSendMessageUploadPhotoAction) },
	int32(TLConstructor_CRC32_sendMessageUploadDocumentAction):                  func() TLObject { return new(TLSendMessageUploadDocumentAction) },
	int32(TLConstructor_CRC32_sendMessageGeoLocationAction):                     func() TLObject { return new(TLSendMessageGeoLocationAction) },
	int32(TLConstructor_CRC32_sendMessageChooseContactAction):                   func() TLObject { return new(TLSendMessageChooseContactAction) },
	int32(TLConstructor_CRC32_sendMessageGamePlayAction):                        func() TLObject { return new(TLSendMessageGamePlayAction) },
	int32(TLConstructor_CRC32_sendMessageRecordRoundAction):                     func() TLObject { return new(TLSendMessageRecordRoundAction) },
	int32(TLConstructor_CRC32_sendMessageUploadRoundAction):                     func() TLObject { return new(TLSendMessageUploadRoundAction) },
	int32(TLConstructor_CRC32_contacts_found):                                   func() TLObject { return new(TLContactsFound) },
	int32(TLConstructor_CRC32_inputPrivacyKeyStatusTimestamp):                   func() TLObject { return new(TLInputPrivacyKeyStatusTimestamp) },
	int32(TLConstructor_CRC32_inputPrivacyKeyChatInvite):                        func() TLObject { return new(TLInputPrivacyKeyChatInvite) },
	int32(TLConstructor_CRC32_inputPrivacyKeyPhoneCall):                         func() TLObject { return new(TLInputPrivacyKeyPhoneCall) },
	int32(TLConstructor_CRC32_privacyKeyStatusTimestamp):                        func() TLObject { return new(TLPrivacyKeyStatusTimestamp) },
	int32(TLConstructor_CRC32_privacyKeyChatInvite):                             func() TLObject { return new(TLPrivacyKeyChatInvite) },
	int32(TLConstructor_CRC32_privacyKeyPhoneCall):                              func() TLObject { return new(TLPrivacyKeyPhoneCall) },
	int32(TLConstructor_CRC32_inputPrivacyValueAllowContacts):                   func() TLObject { return new(TLInputPrivacyValueAllowContacts) },
	int32(TLConstructor_CRC32_inputPrivacyValueAllowAll):                        func() TLObject { return new(TLInputPrivacyValueAllowAll) },
	int32(TLConstructor_CRC32_inputPrivacyValueAllowUsers):                      func() TLObject { return new(TLInputPrivacyValueAllowUsers) },
	int32(TLConstructor_CRC32_inputPrivacyValueDisallowContacts):                func() TLObject { return new(TLInputPrivacyValueDisallowContacts) },
	int32(TLConstructor_CRC32_inputPrivacyValueDisallowAll):                     func() TLObject { return new(TLInputPrivacyValueDisallowAll) },
	int32(TLConstructor_CRC32_inputPrivacyValueDisallowUsers):                   func() TLObject { return new(TLInputPrivacyValueDisallowUsers) },
	int32(TLConstructor_CRC32_privacyValueAllowContacts):                        func() TLObject { return new(TLPrivacyValueAllowContacts) },
	int32(TLConstructor_CRC32_privacyValueAllowAll):                             func() TLObject { return new(TLPrivacyValueAllowAll) },
	int32(TLConstructor_CRC32_privacyValueAllowUsers):                           func() TLObject { return new(TLPrivacyValueAllowUsers) },
	int32(TLConstructor_CRC32_privacyValueDisallowContacts):                     func() TLObject { return new(TLPrivacyValueDisallowContacts) },
	int32(TLConstructor_CRC32_privacyValueDisallowAll):                          func() TLObject { return new(TLPrivacyValueDisallowAll) },
	int32(TLConstructor_CRC32_privacyValueDisallowUsers):                        func() TLObject { return new(TLPrivacyValueDisallowUsers) },
	int32(TLConstructor_CRC32_account_privacyRules):                             func() TLObject { return new(TLAccountPrivacyRules) },
	int32(TLConstructor_CRC32_accountDaysTTL):                                   func() TLObject { return new(TLAccountDaysTTL) },
	int32(TLConstructor_CRC32_documentAttributeImageSize):                       func() TLObject { return new(TLDocumentAttributeImageSize) },
	int32(TLConstructor_CRC32_documentAttributeAnimated):                        func() TLObject { return new(TLDocumentAttributeAnimated) },
	int32(TLConstructor_CRC32_documentAttributeSticker):                         func() TLObject { return new(TLDocumentAttributeSticker) },
	int32(TLConstructor_CRC32_documentAttributeVideo):                           func() TLObject { return new(TLDocumentAttributeVideo) },
	int32(TLConstructor_CRC32_documentAttributeAudio):                           func() TLObject { return new(TLDocumentAttributeAudio) },
	int32(TLConstructor_CRC32_documentAttributeFilename):                        func() TLObject { return new(TLDocumentAttributeFilename) },
	int32(TLConstructor_CRC32_documentAttributeHasStickers):                     func() TLObject { return new(TLDocumentAttributeHasStickers) },
	int32(TLConstructor_CRC32_messages_stickersNotModified):                     func() TLObject { return new(TLMessagesStickersNotModified) },
	int32(TLConstructor_CRC32_messages_stickers):                                func() TLObject { return new(TLMessagesStickers) },
	int32(TLConstructor_CRC32_stickerPack):                                      func() TLObject { return new(TLStickerPack) },
	int32(TLConstructor_CRC32_messages_allStickersNotModified):                  func() TLObject { return new(TLMessagesAllStickersNotModified) },
	int32(TLConstructor_CRC32_messages_allStickers):                             func() TLObject { return new(TLMessagesAllStickers) },
	int32(TLConstructor_CRC32_disabledFeature):                                  func() TLObject { return new(TLDisabledFeature) },
	int32(TLConstructor_CRC32_messages_affectedMessages):                        func() TLObject { return new(TLMessagesAffectedMessages) },
	int32(TLConstructor_CRC32_contactLinkUnknown):                               func() TLObject { return new(TLContactLinkUnknown) },
	int32(TLConstructor_CRC32_contactLinkNone):                                  func() TLObject { return new(TLContactLinkNone) },
	int32(TLConstructor_CRC32_contactLinkHasPhone):                              func() TLObject { return new(TLContactLinkHasPhone) },
	int32(TLConstructor_CRC32_contactLinkContact):                               func() TLObject { return new(TLContactLinkContact) },
	int32(TLConstructor_CRC32_webPageEmpty):                                     func() TLObject { return new(TLWebPageEmpty) },
	int32(TLConstructor_CRC32_webPagePending):                                   func() TLObject { return new(TLWebPagePending) },
	int32(TLConstructor_CRC32_webPage):                                          func() TLObject { return new(TLWebPage) },
	int32(TLConstructor_CRC32_webPageNotModified):                               func() TLObject { return new(TLWebPageNotModified) },
	int32(TLConstructor_CRC32_authorization):                                    func() TLObject { return new(TLAuthorization) },
	int32(TLConstructor_CRC32_account_authorizations):                           func() TLObject { return new(TLAccountAuthorizations) },
	int32(TLConstructor_CRC32_account_noPassword):                               func() TLObject { return new(TLAccountNoPassword) },
	int32(TLConstructor_CRC32_account_password):                                 func() TLObject { return new(TLAccountPassword) },
	int32(TLConstructor_CRC32_account_passwordSettings):                         func() TLObject { return new(TLAccountPasswordSettings) },
	int32(TLConstructor_CRC32_account_passwordInputSettings):                    func() TLObject { return new(TLAccountPasswordInputSettings) },
	int32(TLConstructor_CRC32_auth_passwordRecovery):                            func() TLObject { return new(TLAuthPasswordRecovery) },
	int32(TLConstructor_CRC32_receivedNotifyMessage):                            func() TLObject { return new(TLReceivedNotifyMessage) },
	int32(TLConstructor_CRC32_chatInviteEmpty):                                  func() TLObject { return new(TLChatInviteEmpty) },
	int32(TLConstructor_CRC32_chatInviteExported):                               func() TLObject { return new(TLChatInviteExported) },
	int32(TLConstructor_CRC32_chatInviteAlready):                                func() TLObject { return new(TLChatInviteAlready) },
	int32(TLConstructor_CRC32_chatInvite):                                       func() TLObject { return new(TLChatInvite) },
	int32(TLConstructor_CRC32_inputStickerSetEmpty):                             func() TLObject { return new(TLInputStickerSetEmpty) },
	int32(TLConstructor_CRC32_inputStickerSetID):                                func() TLObject { return new(TLInputStickerSetID) },
	int32(TLConstructor_CRC32_inputStickerSetShortName):                         func() TLObject { return new(TLInputStickerSetShortName) },
	int32(TLConstructor_CRC32_stickerSet):                                       func() TLObject { return new(TLStickerSet) },
	int32(TLConstructor_CRC32_messages_stickerSet):                              func() TLObject { return new(TLMessagesStickerSet) },
	int32(TLConstructor_CRC32_botCommand):                                       func() TLObject { return new(TLBotCommand) },
	int32(TLConstructor_CRC32_botInfo):                                          func() TLObject { return new(TLBotInfo) },
	int32(TLConstructor_CRC32_keyboardButton):                                   func() TLObject { return new(TLKeyboardButton) },
	int32(TLConstructor_CRC32_keyboardButtonUrl):                                func() TLObject { return new(TLKeyboardButtonUrl) },
	int32(TLConstructor_CRC32_keyboardButtonCallback):                           func() TLObject { return new(TLKeyboardButtonCallback) },
	int32(TLConstructor_CRC32_keyboardButtonRequestPhone):                       func() TLObject { return new(TLKeyboardButtonRequestPhone) },
	int32(TLConstructor_CRC32_keyboardButtonRequestGeoLocation):                 func() TLObject { return new(TLKeyboardButtonRequestGeoLocation) },
	int32(TLConstructor_CRC32_keyboardButtonSwitchInline):                       func() TLObject { return new(TLKeyboardButtonSwitchInline) },
	int32(TLConstructor_CRC32_keyboardButtonGame):                               func() TLObject { return new(TLKeyboardButtonGame) },
	int32(TLConstructor_CRC32_keyboardButtonBuy):                                func() TLObject { return new(TLKeyboardButtonBuy) },
	int32(TLConstructor_CRC32_keyboardButtonRow):                                func() TLObject { return new(TLKeyboardButtonRow) },
	int32(TLConstructor_CRC32_replyKeyboardHide):                                func() TLObject { return new(TLReplyKeyboardHide) },
	int32(TLConstructor_CRC32_replyKeyboardForceReply):                          func() TLObject { return new(TLReplyKeyboardForceReply) },
	int32(TLConstructor_CRC32_replyKeyboardMarkup):                              func() TLObject { return new(TLReplyKeyboardMarkup) },
	int32(TLConstructor_CRC32_replyInlineMarkup):                                func() TLObject { return new(TLReplyInlineMarkup) },
	int32(TLConstructor_CRC32_messageEntityUnknown):                             func() TLObject { return new(TLMessageEntityUnknown) },
	int32(TLConstructor_CRC32_messageEntityMention):                             func() TLObject { return new(TLMessageEntityMention) },
	int32(TLConstructor_CRC32_messageEntityHashtag):                             func() TLObject { return new(TLMessageEntityHashtag) },
	int32(TLConstructor_CRC32_messageEntityBotCommand):                          func() TLObject { return new(TLMessageEntityBotCommand) },
	int32(TLConstructor_CRC32_messageEntityUrl):                                 func() TLObject { return new(TLMessageEntityUrl) },
	int32(TLConstructor_CRC32_messageEntityEmail):                               func() TLObject { return new(TLMessageEntityEmail) },
	int32(TLConstructor_CRC32_messageEntityBold):                                func() TLObject { return new(TLMessageEntityBold) },
	int32(TLConstructor_CRC32_messageEntityItalic):                              func() TLObject { return new(TLMessageEntityItalic) },
	int32(TLConstructor_CRC32_messageEntityCode):                                func() TLObject { return new(TLMessageEntityCode) },
	int32(TLConstructor_CRC32_messageEntityPre):                                 func() TLObject { return new(TLMessageEntityPre) },
	int32(TLConstructor_CRC32_messageEntityTextUrl):                             func() TLObject { return new(TLMessageEntityTextUrl) },
	int32(TLConstructor_CRC32_messageEntityMentionName):                         func() TLObject { return new(TLMessageEntityMentionName) },
	int32(TLConstructor_CRC32_inputMessageEntityMentionName):                    func() TLObject { return new(TLInputMessageEntityMentionName) },
	int32(TLConstructor_CRC32_inputChannelEmpty):                                func() TLObject { return new(TLInputChannelEmpty) },
	int32(TLConstructor_CRC32_inputChannel):                                     func() TLObject { return new(TLInputChannel) },
	int32(TLConstructor_CRC32_contacts_resolvedPeer):                            func() TLObject { return new(TLContactsResolvedPeer) },
	int32(TLConstructor_CRC32_messageRange):                                     func() TLObject { return new(TLMessageRange) },
	int32(TLConstructor_CRC32_updates_channelDifferenceEmpty):                   func() TLObject { return new(TLUpdatesChannelDifferenceEmpty) },
	int32(TLConstructor_CRC32_updates_channelDifferenceTooLong):                 func() TLObject { return new(TLUpdatesChannelDifferenceTooLong) },
	int32(TLConstructor_CRC32_updates_channelDifference):                        func() TLObject { return new(TLUpdatesChannelDifference) },
	int32(TLConstructor_CRC32_channelMessagesFilterEmpty):                       func() TLObject { return new(TLChannelMessagesFilterEmpty) },
	int32(TLConstructor_CRC32_channelMessagesFilter):                            func() TLObject { return new(TLChannelMessagesFilter) },
	int32(TLConstructor_CRC32_channelParticipant):                               func() TLObject { return new(TLChannelParticipant) },
	int32(TLConstructor_CRC32_channelParticipantSelf):                           func() TLObject { return new(TLChannelParticipantSelf) },
	int32(TLConstructor_CRC32_channelParticipantCreator):                        func() TLObject { return new(TLChannelParticipantCreator) },
	int32(TLConstructor_CRC32_channelParticipantAdmin):                          func() TLObject { return new(TLChannelParticipantAdmin) },
	int32(TLConstructor_CRC32_channelParticipantBanned):                         func() TLObject { return new(TLChannelParticipantBanned) },
	int32(TLConstructor_CRC32_channelParticipantsRecent):                        func() TLObject { return new(TLChannelParticipantsRecent) },
	int32(TLConstructor_CRC32_channelParticipantsAdmins):                        func() TLObject { return new(TLChannelParticipantsAdmins) },
	int32(TLConstructor_CRC32_channelParticipantsKicked):                        func() TLObject { return new(TLChannelParticipantsKicked) },
	int32(TLConstructor_CRC32_channelParticipantsBots):                          func() TLObject { return new(TLChannelParticipantsBots) },
	int32(TLConstructor_CRC32_channelParticipantsBanned):                        func() TLObject { return new(TLChannelParticipantsBanned) },
	int32(TLConstructor_CRC32_channelParticipantsSearch):                        func() TLObject { return new(TLChannelParticipantsSearch) },
	int32(TLConstructor_CRC32_channels_channelParticipants):                     func() TLObject { return new(TLChannelsChannelParticipants) },
	int32(TLConstructor_CRC32_channels_channelParticipant):                      func() TLObject { return new(TLChannelsChannelParticipant) },
	int32(TLConstructor_CRC32_help_termsOfService):                              func() TLObject { return new(TLHelpTermsOfService) },
	int32(TLConstructor_CRC32_foundGif):                                         func() TLObject { return new(TLFoundGif) },
	int32(TLConstructor_CRC32_foundGifCached):                                   func() TLObject { return new(TLFoundGifCached) },
	int32(TLConstructor_CRC32_messages_foundGifs):                               func() TLObject { return new(TLMessagesFoundGifs) },
	int32(TLConstructor_CRC32_messages_savedGifsNotModified):                    func() TLObject { return new(TLMessagesSavedGifsNotModified) },
	int32(TLConstructor_CRC32_messages_savedGifs):                               func() TLObject { return new(TLMessagesSavedGifs) },
	int32(TLConstructor_CRC32_inputBotInlineMessageMediaAuto):                   func() TLObject { return new(TLInputBotInlineMessageMediaAuto) },
	int32(TLConstructor_CRC32_inputBotInlineMessageText):                        func() TLObject { return new(TLInputBotInlineMessageText) },
	int32(TLConstructor_CRC32_inputBotInlineMessageMediaGeo):                    func() TLObject { return new(TLInputBotInlineMessageMediaGeo) },
	int32(TLConstructor_CRC32_inputBotInlineMessageMediaVenue):                  func() TLObject { return new(TLInputBotInlineMessageMediaVenue) },
	int32(TLConstructor_CRC32_inputBotInlineMessageMediaContact):                func() TLObject { return new(TLInputBotInlineMessageMediaContact) },
	int32(TLConstructor_CRC32_inputBotInlineMessageGame):                        func() TLObject { return new(TLInputBotInlineMessageGame) },
	int32(TLConstructor_CRC32_inputBotInlineResult):                             func() TLObject { return new(TLInputBotInlineResult) },
	int32(TLConstructor_CRC32_inputBotInlineResultPhoto):                        func() TLObject { return new(TLInputBotInlineResultPhoto) },
	int32(TLConstructor_CRC32_inputBotInlineResultDocument):                     func() TLObject { return new(TLInputBotInlineResultDocument) },
	int32(TLConstructor_CRC32_inputBotInlineResultGame):                         func() TLObject { return new(TLInputBotInlineResultGame) },
	int32(TLConstructor_CRC32_botInlineMessageMediaAuto):                        func() TLObject { return new(TLBotInlineMessageMediaAuto) },
	int32(TLConstructor_CRC32_botInlineMessageText):                             func() TLObject { return new(TLBotInlineMessageText) },
	int32(TLConstructor_CRC32_botInlineMessageMediaGeo):                         func() TLObject { return new(TLBotInlineMessageMediaGeo) },
	int32(TLConstructor_CRC32_botInlineMessageMediaVenue):                       func() TLObject { return new(TLBotInlineMessageMediaVenue) },
	int32(TLConstructor_CRC32_botInlineMessageMediaContact):                     func() TLObject { return new(TLBotInlineMessageMediaContact) },
	int32(TLConstructor_CRC32_botInlineResult):                                  func() TLObject { return new(TLBotInlineResult) },
	int32(TLConstructor_CRC32_botInlineMediaResult):                             func() TLObject { return new(TLBotInlineMediaResult) },
	int32(TLConstructor_CRC32_messages_botResults):                              func() TLObject { return new(TLMessagesBotResults) },
	int32(TLConstructor_CRC32_exportedMessageLink):                              func() TLObject { return new(TLExportedMessageLink) },
	int32(TLConstructor_CRC32_messageFwdHeader):                                 func() TLObject { return new(TLMessageFwdHeader) },
	int32(TLConstructor_CRC32_auth_codeTypeSms):                                 func() TLObject { return new(TLAuthCodeTypeSms) },
	int32(TLConstructor_CRC32_auth_codeTypeCall):                                func() TLObject { return new(TLAuthCodeTypeCall) },
	int32(TLConstructor_CRC32_auth_codeTypeFlashCall):                           func() TLObject { return new(TLAuthCodeTypeFlashCall) },
	int32(TLConstructor_CRC32_auth_sentCodeTypeApp):                             func() TLObject { return new(TLAuthSentCodeTypeApp) },
	int32(TLConstructor_CRC32_auth_sentCodeTypeSms):                             func() TLObject { return new(TLAuthSentCodeTypeSms) },
	int32(TLConstructor_CRC32_auth_sentCodeTypeCall):                            func() TLObject { return new(TLAuthSentCodeTypeCall) },
	int32(TLConstructor_CRC32_auth_sentCodeTypeFlashCall):                       func() TLObject { return new(TLAuthSentCodeTypeFlashCall) },
	int32(TLConstructor_CRC32_messages_botCallbackAnswer):                       func() TLObject { return new(TLMessagesBotCallbackAnswer) },
	int32(TLConstructor_CRC32_messages_messageEditData):                         func() TLObject { return new(TLMessagesMessageEditData) },
	int32(TLConstructor_CRC32_inputBotInlineMessageID):                          func() TLObject { return new(TLInputBotInlineMessageID) },
	int32(TLConstructor_CRC32_inlineBotSwitchPM):                                func() TLObject { return new(TLInlineBotSwitchPM) },
	int32(TLConstructor_CRC32_messages_peerDialogs):                             func() TLObject { return new(TLMessagesPeerDialogs) },
	int32(TLConstructor_CRC32_topPeer):                                          func() TLObject { return new(TLTopPeer) },
	int32(TLConstructor_CRC32_topPeerCategoryBotsPM):                            func() TLObject { return new(TLTopPeerCategoryBotsPM) },
	int32(TLConstructor_CRC32_topPeerCategoryBotsInline):                        func() TLObject { return new(TLTopPeerCategoryBotsInline) },
	int32(TLConstructor_CRC32_topPeerCategoryCorrespondents):                    func() TLObject { return new(TLTopPeerCategoryCorrespondents) },
	int32(TLConstructor_CRC32_topPeerCategoryGroups):                            func() TLObject { return new(TLTopPeerCategoryGroups) },
	int32(TLConstructor_CRC32_topPeerCategoryChannels):                          func() TLObject { return new(TLTopPeerCategoryChannels) },
	int32(TLConstructor_CRC32_topPeerCategoryPhoneCalls):                        func() TLObject { return new(TLTopPeerCategoryPhoneCalls) },
	int32(TLConstructor_CRC32_topPeerCategoryPeers):                             func() TLObject { return new(TLTopPeerCategoryPeers) },
	int32(TLConstructor_CRC32_contacts_topPeersNotModified):                     func() TLObject { return new(TLContactsTopPeersNotModified) },
	int32(TLConstructor_CRC32_contacts_topPeers):                                func() TLObject { return new(TLContactsTopPeers) },
	int32(TLConstructor_CRC32_draftMessageEmpty):                                func() TLObject { return new(TLDraftMessageEmpty) },
	int32(TLConstructor_CRC32_draftMessage):                                     func() TLObject { return new(TLDraftMessage) },
	int32(TLConstructor_CRC32_messages_featuredStickersNotModified):             func() TLObject { return new(TLMessagesFeaturedStickersNotModified) },
	int32(TLConstructor_CRC32_messages_featuredStickers):                        func() TLObject { return new(TLMessagesFeaturedStickers) },
	int32(TLConstructor_CRC32_messages_recentStickersNotModified):               func() TLObject { return new(TLMessagesRecentStickersNotModified) },
	int32(TLConstructor_CRC32_messages_recentStickers):                          func() TLObject { return new(TLMessagesRecentStickers) },
	int32(TLConstructor_CRC32_messages_archivedStickers):                        func() TLObject { return new(TLMessagesArchivedStickers) },
	int32(TLConstructor_CRC32_messages_stickerSetInstallResultSuccess):          func() TLObject { return new(TLMessagesStickerSetInstallResultSuccess) },
	int32(TLConstructor_CRC32_messages_stickerSetInstallResultArchive):          func() TLObject { return new(TLMessagesStickerSetInstallResultArchive) },
	int32(TLConstructor_CRC32_stickerSetCovered):                                func() TLObject { return new(TLStickerSetCovered) },
	int32(TLConstructor_CRC32_stickerSetMultiCovered):                           func() TLObject { return new(TLStickerSetMultiCovered) },
	int32(TLConstructor_CRC32_maskCoords):                                       func() TLObject { return new(TLMaskCoords) },
	int32(TLConstructor_CRC32_inputStickeredMediaPhoto):                         func() TLObject { return new(TLInputStickeredMediaPhoto) },
	int32(TLConstructor_CRC32_inputStickeredMediaDocument):                      func() TLObject { return new(TLInputStickeredMediaDocument) },
	int32(TLConstructor_CRC32_game):                                             func() TLObject { return new(TLGame) },
	int32(TLConstructor_CRC32_inputGameID):                                      func() TLObject { return new(TLInputGameID) },
	int32(TLConstructor_CRC32_inputGameShortName):                               func() TLObject { return new(TLInputGameShortName) },
	int32(TLConstructor_CRC32_highScore):                                        func() TLObject { return new(TLHighScore) },
	int32(TLConstructor_CRC32_messages_highScores):                              func() TLObject { return new(TLMessagesHighScores) },
	int32(TLConstructor_CRC32_textEmpty):                                        func() TLObject { return new(TLTextEmpty) },
	int32(TLConstructor_CRC32_textPlain):                                        func() TLObject { return new(TLTextPlain) },
	int32(TLConstructor_CRC32_textBold):                                         func() TLObject { return new(TLTextBold) },
	int32(TLConstructor_CRC32_textItalic):                                       func() TLObject { return new(TLTextItalic) },
	int32(TLConstructor_CRC32_textUnderline):                                    func() TLObject { return new(TLTextUnderline) },
	int32(TLConstructor_CRC32_textStrike):                                       func() TLObject { return new(TLTextStrike) },
	int32(TLConstructor_CRC32_textFixed):                                        func() TLObject { return new(TLTextFixed) },
	int32(TLConstructor_CRC32_textUrl):                                          func() TLObject { return new(TLTextUrl) },
	int32(TLConstructor_CRC32_textEmail):                                        func() TLObject { return new(TLTextEmail) },
	int32(TLConstructor_CRC32_textConcat):                                       func() TLObject { return new(TLTextConcat) },
	int32(TLConstructor_CRC32_pageBlockUnsupported):                             func() TLObject { return new(TLPageBlockUnsupported) },
	int32(TLConstructor_CRC32_pageBlockTitle):                                   func() TLObject { return new(TLPageBlockTitle) },
	int32(TLConstructor_CRC32_pageBlockSubtitle):                                func() TLObject { return new(TLPageBlockSubtitle) },
	int32(TLConstructor_CRC32_pageBlockAuthorDate):                              func() TLObject { return new(TLPageBlockAuthorDate) },
	int32(TLConstructor_CRC32_pageBlockHeader):                                  func() TLObject { return new(TLPageBlockHeader) },
	int32(TLConstructor_CRC32_pageBlockSubheader):                               func() TLObject { return new(TLPageBlockSubheader) },
	int32(TLConstructor_CRC32_pageBlockParagraph):                               func() TLObject { return new(TLPageBlockParagraph) },
	int32(TLConstructor_CRC32_pageBlockPreformatted):                            func() TLObject { return new(TLPageBlockPreformatted) },
	int32(TLConstructor_CRC32_pageBlockFooter):                                  func() TLObject { return new(TLPageBlockFooter) },
	int32(TLConstructor_CRC32_pageBlockDivider):                                 func() TLObject { return new(TLPageBlockDivider) },
	int32(TLConstructor_CRC32_pageBlockAnchor):                                  func() TLObject { return new(TLPageBlockAnchor) },
	int32(TLConstructor_CRC32_pageBlockList):                                    func() TLObject { return new(TLPageBlockList) },
	int32(TLConstructor_CRC32_pageBlockBlockquote):                              func() TLObject { return new(TLPageBlockBlockquote) },
	int32(TLConstructor_CRC32_pageBlockPullquote):                               func() TLObject { return new(TLPageBlockPullquote) },
	int32(TLConstructor_CRC32_pageBlockPhoto):                                   func() TLObject { return new(TLPageBlockPhoto) },
	int32(TLConstructor_CRC32_pageBlockVideo):                                   func() TLObject { return new(TLPageBlockVideo) },
	int32(TLConstructor_CRC32_pageBlockCover):                                   func() TLObject { return new(TLPageBlockCover) },
	int32(TLConstructor_CRC32_pageBlockEmbed):                                   func() TLObject { return new(TLPageBlockEmbed) },
	int32(TLConstructor_CRC32_pageBlockEmbedPost):                               func() TLObject { return new(TLPageBlockEmbedPost) },
	int32(TLConstructor_CRC32_pageBlockCollage):                                 func() TLObject { return new(TLPageBlockCollage) },
	int32(TLConstructor_CRC32_pageBlockSlideshow):                               func() TLObject { return new(TLPageBlockSlideshow) },
	int32(TLConstructor_CRC32_pageBlockChannel):                                 func() TLObject { return new(TLPageBlockChannel) },
	int32(TLConstructor_CRC32_pageBlockAudio):                                   func() TLObject { return new(TLPageBlockAudio) },
	int32(TLConstructor_CRC32_pagePart):                                         func() TLObject { return new(TLPagePart) },
	int32(TLConstructor_CRC32_pageFull):                                         func() TLObject { return new(TLPageFull) },
	int32(TLConstructor_CRC32_phoneCallDiscardReasonMissed):                     func() TLObject { return new(TLPhoneCallDiscardReasonMissed) },
	int32(TLConstructor_CRC32_phoneCallDiscardReasonDisconnect):                 func() TLObject { return new(TLPhoneCallDiscardReasonDisconnect) },
	int32(TLConstructor_CRC32_phoneCallDiscardReasonHangup):                     func() TLObject { return new(TLPhoneCallDiscardReasonHangup) },
	int32(TLConstructor_CRC32_phoneCallDiscardReasonBusy):                       func() TLObject { return new(TLPhoneCallDiscardReasonBusy) },
	int32(TLConstructor_CRC32_dataJSON):                                         func() TLObject { return new(TLDataJSON) },
	int32(TLConstructor_CRC32_labeledPrice):                                     func() TLObject { return new(TLLabeledPrice) },
	int32(TLConstructor_CRC32_invoice):                                          func() TLObject { return new(TLInvoice) },
	int32(TLConstructor_CRC32_paymentCharge):                                    func() TLObject { return new(TLPaymentCharge) },
	int32(TLConstructor_CRC32_postAddress):                                      func() TLObject { return new(TLPostAddress) },
	int32(TLConstructor_CRC32_paymentRequestedInfo):                             func() TLObject { return new(TLPaymentRequestedInfo) },
	int32(TLConstructor_CRC32_paymentSavedCredentialsCard):                      func() TLObject { return new(TLPaymentSavedCredentialsCard) },
	int32(TLConstructor_CRC32_webDocument):                                      func() TLObject { return new(TLWebDocument) },
	int32(TLConstructor_CRC32_inputWebDocument):                                 func() TLObject { return new(TLInputWebDocument) },
	int32(TLConstructor_CRC32_inputWebFileLocation):                             func() TLObject { return new(TLInputWebFileLocation) },
	int32(TLConstructor_CRC32_upload_webFile):                                   func() TLObject { return new(TLUploadWebFile) },
	int32(TLConstructor_CRC32_payments_paymentForm):                             func() TLObject { return new(TLPaymentsPaymentForm) },
	int32(TLConstructor_CRC32_payments_validatedRequestedInfo):                  func() TLObject { return new(TLPaymentsValidatedRequestedInfo) },
	int32(TLConstructor_CRC32_payments_paymentResult):                           func() TLObject { return new(TLPaymentsPaymentResult) },
	int32(TLConstructor_CRC32_payments_paymentVerficationNeeded):                func() TLObject { return new(TLPaymentsPaymentVerficationNeeded) },
	int32(TLConstructor_CRC32_payments_paymentReceipt):                          func() TLObject { return new(TLPaymentsPaymentReceipt) },
	int32(TLConstructor_CRC32_payments_savedInfo):                               func() TLObject { return new(TLPaymentsSavedInfo) },
	int32(TLConstructor_CRC32_inputPaymentCredentialsSaved):                     func() TLObject { return new(TLInputPaymentCredentialsSaved) },
	int32(TLConstructor_CRC32_inputPaymentCredentials):                          func() TLObject { return new(TLInputPaymentCredentials) },
	int32(TLConstructor_CRC32_account_tmpPassword):                              func() TLObject { return new(TLAccountTmpPassword) },
	int32(TLConstructor_CRC32_shippingOption):                                   func() TLObject { return new(TLShippingOption) },
	int32(TLConstructor_CRC32_inputStickerSetItem):                              func() TLObject { return new(TLInputStickerSetItem) },
	int32(TLConstructor_CRC32_inputPhoneCall):                                   func() TLObject { return new(TLInputPhoneCall) },
	int32(TLConstructor_CRC32_phoneCallEmpty):                                   func() TLObject { return new(TLPhoneCallEmpty) },
	int32(TLConstructor_CRC32_phoneCallWaiting):                                 func() TLObject { return new(TLPhoneCallWaiting) },
	int32(TLConstructor_CRC32_phoneCallRequested):                               func() TLObject { return new(TLPhoneCallRequested) },
	int32(TLConstructor_CRC32_phoneCallAccepted):                                func() TLObject { return new(TLPhoneCallAccepted) },
	int32(TLConstructor_CRC32_phoneCall):                                        func() TLObject { return new(TLPhoneCall) },
	int32(TLConstructor_CRC32_phoneCallDiscarded):                               func() TLObject { return new(TLPhoneCallDiscarded) },
	int32(TLConstructor_CRC32_phoneConnection):                                  func() TLObject { return new(TLPhoneConnection) },
	int32(TLConstructor_CRC32_phoneCallProtocol):                                func() TLObject { return new(TLPhoneCallProtocol) },
	int32(TLConstructor_CRC32_phone_phoneCall):                                  func() TLObject { return new(TLPhonePhoneCall) },
	int32(TLConstructor_CRC32_upload_cdnFileReuploadNeeded):                     func() TLObject { return new(TLUploadCdnFileReuploadNeeded) },
	int32(TLConstructor_CRC32_upload_cdnFile):                                   func() TLObject { return new(TLUploadCdnFile) },
	int32(TLConstructor_CRC32_cdnPublicKey):                                     func() TLObject { return new(TLCdnPublicKey) },
	int32(TLConstructor_CRC32_cdnConfig):                                        func() TLObject { return new(TLCdnConfig) },
	int32(TLConstructor_CRC32_langPackString):                                   func() TLObject { return new(TLLangPackString) },
	int32(TLConstructor_CRC32_langPackStringPluralized):                         func() TLObject { return new(TLLangPackStringPluralized) },
	int32(TLConstructor_CRC32_langPackStringDeleted):                            func() TLObject { return new(TLLangPackStringDeleted) },
	int32(TLConstructor_CRC32_langPackDifference):                               func() TLObject { return new(TLLangPackDifference) },
	int32(TLConstructor_CRC32_langPackLanguage):                                 func() TLObject { return new(TLLangPackLanguage) },
	int32(TLConstructor_CRC32_channelAdminRights):                               func() TLObject { return new(TLChannelAdminRights) },
	int32(TLConstructor_CRC32_channelBannedRights):                              func() TLObject { return new(TLChannelBannedRights) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionChangeTitle):            func() TLObject { return new(TLChannelAdminLogEventActionChangeTitle) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionChangeAbout):            func() TLObject { return new(TLChannelAdminLogEventActionChangeAbout) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionChangeUsername):         func() TLObject { return new(TLChannelAdminLogEventActionChangeUsername) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionChangePhoto):            func() TLObject { return new(TLChannelAdminLogEventActionChangePhoto) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionToggleInvites):          func() TLObject { return new(TLChannelAdminLogEventActionToggleInvites) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionToggleSignatures):       func() TLObject { return new(TLChannelAdminLogEventActionToggleSignatures) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionUpdatePinned):           func() TLObject { return new(TLChannelAdminLogEventActionUpdatePinned) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionEditMessage):            func() TLObject { return new(TLChannelAdminLogEventActionEditMessage) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionDeleteMessage):          func() TLObject { return new(TLChannelAdminLogEventActionDeleteMessage) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantJoin):        func() TLObject { return new(TLChannelAdminLogEventActionParticipantJoin) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantLeave):       func() TLObject { return new(TLChannelAdminLogEventActionParticipantLeave) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantInvite):      func() TLObject { return new(TLChannelAdminLogEventActionParticipantInvite) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleBan):   func() TLObject { return new(TLChannelAdminLogEventActionParticipantToggleBan) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleAdmin): func() TLObject { return new(TLChannelAdminLogEventActionParticipantToggleAdmin) },
	int32(TLConstructor_CRC32_channelAdminLogEventActionChangeStickerSet):       func() TLObject { return new(TLChannelAdminLogEventActionChangeStickerSet) },
	int32(TLConstructor_CRC32_channelAdminLogEvent):                             func() TLObject { return new(TLChannelAdminLogEvent) },
	int32(TLConstructor_CRC32_channels_adminLogResults):                         func() TLObject { return new(TLChannelsAdminLogResults) },
	int32(TLConstructor_CRC32_channelAdminLogEventsFilter):                      func() TLObject { return new(TLChannelAdminLogEventsFilter) },
	int32(TLConstructor_CRC32_popularContact):                                   func() TLObject { return new(TLPopularContact) },
	int32(TLConstructor_CRC32_cdnFileHash):                                      func() TLObject { return new(TLCdnFileHash) },
	int32(TLConstructor_CRC32_messages_favedStickersNotModified):                func() TLObject { return new(TLMessagesFavedStickersNotModified) },
	int32(TLConstructor_CRC32_messages_favedStickers):                           func() TLObject { return new(TLMessagesFavedStickers) },
	int32(TLConstructor_CRC32_invokeAfterMsg):                                   func() TLObject { return new(TLInvokeAfterMsg) },
	int32(TLConstructor_CRC32_invokeAfterMsgs):                                  func() TLObject { return new(TLInvokeAfterMsgs) },
	int32(TLConstructor_CRC32_initConnection):                                   func() TLObject { return new(TLInitConnection) },
	int32(TLConstructor_CRC32_invokeWithLayer):                                  func() TLObject { return new(TLInvokeWithLayer) },
	int32(TLConstructor_CRC32_invokeWithoutUpdates):                             func() TLObject { return new(TLInvokeWithoutUpdates) },
	int32(TLConstructor_CRC32_auth_checkPhone):                                  func() TLObject { return new(TLAuthCheckPhone) },
	int32(TLConstructor_CRC32_auth_sendCode):                                    func() TLObject { return new(TLAuthSendCode) },
	int32(TLConstructor_CRC32_auth_signUp):                                      func() TLObject { return new(TLAuthSignUp) },
	int32(TLConstructor_CRC32_auth_signIn):                                      func() TLObject { return new(TLAuthSignIn) },
	int32(TLConstructor_CRC32_auth_logOut):                                      func() TLObject { return new(TLAuthLogOut) },
	int32(TLConstructor_CRC32_auth_resetAuthorizations):                         func() TLObject { return new(TLAuthResetAuthorizations) },
	int32(TLConstructor_CRC32_auth_sendInvites):                                 func() TLObject { return new(TLAuthSendInvites) },
	int32(TLConstructor_CRC32_auth_exportAuthorization):                         func() TLObject { return new(TLAuthExportAuthorization) },
	int32(TLConstructor_CRC32_auth_importAuthorization):                         func() TLObject { return new(TLAuthImportAuthorization) },
	int32(TLConstructor_CRC32_auth_bindTempAuthKey):                             func() TLObject { return new(TLAuthBindTempAuthKey) },
	int32(TLConstructor_CRC32_auth_importBotAuthorization):                      func() TLObject { return new(TLAuthImportBotAuthorization) },
	int32(TLConstructor_CRC32_auth_checkPassword):                               func() TLObject { return new(TLAuthCheckPassword) },
	int32(TLConstructor_CRC32_auth_requestPasswordRecovery):                     func() TLObject { return new(TLAuthRequestPasswordRecovery) },
	int32(TLConstructor_CRC32_auth_recoverPassword):                             func() TLObject { return new(TLAuthRecoverPassword) },
	int32(TLConstructor_CRC32_auth_resendCode):                                  func() TLObject { return new(TLAuthResendCode) },
	int32(TLConstructor_CRC32_auth_cancelCode):                                  func() TLObject { return new(TLAuthCancelCode) },
	int32(TLConstructor_CRC32_auth_dropTempAuthKeys):                            func() TLObject { return new(TLAuthDropTempAuthKeys) },
	int32(TLConstructor_CRC32_account_registerDevice):                           func() TLObject { return new(TLAccountRegisterDevice) },
	int32(TLConstructor_CRC32_account_unregisterDevice):                         func() TLObject { return new(TLAccountUnregisterDevice) },
	int32(TLConstructor_CRC32_account_updateNotifySettings):                     func() TLObject { return new(TLAccountUpdateNotifySettings) },
	int32(TLConstructor_CRC32_account_getNotifySettings):                        func() TLObject { return new(TLAccountGetNotifySettings) },
	int32(TLConstructor_CRC32_account_resetNotifySettings):                      func() TLObject { return new(TLAccountResetNotifySettings) },
	int32(TLConstructor_CRC32_account_updateProfile):                            func() TLObject { return new(TLAccountUpdateProfile) },
	int32(TLConstructor_CRC32_account_updateStatus):                             func() TLObject { return new(TLAccountUpdateStatus) },
	int32(TLConstructor_CRC32_account_getWallPapers):                            func() TLObject { return new(TLAccountGetWallPapers) },
	int32(TLConstructor_CRC32_account_reportPeer):                               func() TLObject { return new(TLAccountReportPeer) },
	int32(TLConstructor_CRC32_account_checkUsername):                            func() TLObject { return new(TLAccountCheckUsername) },
	int32(TLConstructor_CRC32_account_updateUsername):                           func() TLObject { return new(TLAccountUpdateUsername) },
	int32(TLConstructor_CRC32_account_getPrivacy):                               func() TLObject { return new(TLAccountGetPrivacy) },
	int32(TLConstructor_CRC32_account_setPrivacy):                               func() TLObject { return new(TLAccountSetPrivacy) },
	int32(TLConstructor_CRC32_account_deleteAccount):                            func() TLObject { return new(TLAccountDeleteAccount) },
	int32(TLConstructor_CRC32_account_getAccountTTL):                            func() TLObject { return new(TLAccountGetAccountTTL) },
	int32(TLConstructor_CRC32_account_setAccountTTL):                            func() TLObject { return new(TLAccountSetAccountTTL) },
	int32(TLConstructor_CRC32_account_sendChangePhoneCode):                      func() TLObject { return new(TLAccountSendChangePhoneCode) },
	int32(TLConstructor_CRC32_account_changePhone):                              func() TLObject { return new(TLAccountChangePhone) },
	int32(TLConstructor_CRC32_account_updateDeviceLocked):                       func() TLObject { return new(TLAccountUpdateDeviceLocked) },
	int32(TLConstructor_CRC32_account_getAuthorizations):                        func() TLObject { return new(TLAccountGetAuthorizations) },
	int32(TLConstructor_CRC32_account_resetAuthorization):                       func() TLObject { return new(TLAccountResetAuthorization) },
	int32(TLConstructor_CRC32_account_getPassword):                              func() TLObject { return new(TLAccountGetPassword) },
	int32(TLConstructor_CRC32_account_getPasswordSettings):                      func() TLObject { return new(TLAccountGetPasswordSettings) },
	int32(TLConstructor_CRC32_account_updatePasswordSettings):                   func() TLObject { return new(TLAccountUpdatePasswordSettings) },
	int32(TLConstructor_CRC32_account_sendConfirmPhoneCode):                     func() TLObject { return new(TLAccountSendConfirmPhoneCode) },
	int32(TLConstructor_CRC32_account_confirmPhone):                             func() TLObject { return new(TLAccountConfirmPhone) },
	int32(TLConstructor_CRC32_account_getTmpPassword):                           func() TLObject { return new(TLAccountGetTmpPassword) },
	int32(TLConstructor_CRC32_users_getUsers):                                   func() TLObject { return new(TLUsersGetUsers) },
	int32(TLConstructor_CRC32_users_getFullUser):                                func() TLObject { return new(TLUsersGetFullUser) },
	int32(TLConstructor_CRC32_contacts_getStatuses):                             func() TLObject { return new(TLContactsGetStatuses) },
	int32(TLConstructor_CRC32_contacts_getContacts):                             func() TLObject { return new(TLContactsGetContacts) },
	int32(TLConstructor_CRC32_contacts_importContacts):                          func() TLObject { return new(TLContactsImportContacts) },
	int32(TLConstructor_CRC32_contacts_deleteContact):                           func() TLObject { return new(TLContactsDeleteContact) },
	int32(TLConstructor_CRC32_contacts_deleteContacts):                          func() TLObject { return new(TLContactsDeleteContacts) },
	int32(TLConstructor_CRC32_contacts_block):                                   func() TLObject { return new(TLContactsBlock) },
	int32(TLConstructor_CRC32_contacts_unblock):                                 func() TLObject { return new(TLContactsUnblock) },
	int32(TLConstructor_CRC32_contacts_getBlocked):                              func() TLObject { return new(TLContactsGetBlocked) },
	int32(TLConstructor_CRC32_contacts_exportCard):                              func() TLObject { return new(TLContactsExportCard) },
	int32(TLConstructor_CRC32_contacts_importCard):                              func() TLObject { return new(TLContactsImportCard) },
	int32(TLConstructor_CRC32_contacts_search):                                  func() TLObject { return new(TLContactsSearch) },
	int32(TLConstructor_CRC32_contacts_resolveUsername):                         func() TLObject { return new(TLContactsResolveUsername) },
	int32(TLConstructor_CRC32_contacts_getTopPeers):                             func() TLObject { return new(TLContactsGetTopPeers) },
	int32(TLConstructor_CRC32_contacts_resetTopPeerRating):                      func() TLObject { return new(TLContactsResetTopPeerRating) },
	int32(TLConstructor_CRC32_contacts_resetSaved):                              func() TLObject { return new(TLContactsResetSaved) },
	int32(TLConstructor_CRC32_messages_getMessages):                             func() TLObject { return new(TLMessagesGetMessages) },
	int32(TLConstructor_CRC32_messages_getDialogs):                              func() TLObject { return new(TLMessagesGetDialogs) },
	int32(TLConstructor_CRC32_messages_getHistory):                              func() TLObject { return new(TLMessagesGetHistory) },
	int32(TLConstructor_CRC32_messages_search):                                  func() TLObject { return new(TLMessagesSearch) },
	int32(TLConstructor_CRC32_messages_readHistory):                             func() TLObject { return new(TLMessagesReadHistory) },
	int32(TLConstructor_CRC32_messages_deleteHistory):                           func() TLObject { return new(TLMessagesDeleteHistory) },
	int32(TLConstructor_CRC32_messages_deleteMessages):                          func() TLObject { return new(TLMessagesDeleteMessages) },
	int32(TLConstructor_CRC32_messages_receivedMessages):                        func() TLObject { return new(TLMessagesReceivedMessages) },
	int32(TLConstructor_CRC32_messages_setTyping):                               func() TLObject { return new(TLMessagesSetTyping) },
	int32(TLConstructor_CRC32_messages_sendMessage):                             func() TLObject { return new(TLMessagesSendMessage) },
	int32(TLConstructor_CRC32_messages_sendMedia):                               func() TLObject { return new(TLMessagesSendMedia) },
	int32(TLConstructor_CRC32_messages_forwardMessages):                         func() TLObject { return new(TLMessagesForwardMessages) },
	int32(TLConstructor_CRC32_messages_reportSpam):                              func() TLObject { return new(TLMessagesReportSpam) },
	int32(TLConstructor_CRC32_messages_hideReportSpam):                          func() TLObject { return new(TLMessagesHideReportSpam) },
	int32(TLConstructor_CRC32_messages_getPeerSettings):                         func() TLObject { return new(TLMessagesGetPeerSettings) },
	int32(TLConstructor_CRC32_messages_getChats):                                func() TLObject { return new(TLMessagesGetChats) },
	int32(TLConstructor_CRC32_messages_getFullChat):                             func() TLObject { return new(TLMessagesGetFullChat) },
	int32(TLConstructor_CRC32_messages_editChatTitle):                           func() TLObject { return new(TLMessagesEditChatTitle) },
	int32(TLConstructor_CRC32_messages_editChatPhoto):                           func() TLObject { return new(TLMessagesEditChatPhoto) },
	int32(TLConstructor_CRC32_messages_addChatUser):                             func() TLObject { return new(TLMessagesAddChatUser) },
	int32(TLConstructor_CRC32_messages_deleteChatUser):                          func() TLObject { return new(TLMessagesDeleteChatUser) },
	int32(TLConstructor_CRC32_messages_createChat):                              func() TLObject { return new(TLMessagesCreateChat) },
	int32(TLConstructor_CRC32_messages_forwardMessage):                          func() TLObject { return new(TLMessagesForwardMessage) },
	int32(TLConstructor_CRC32_messages_getDhConfig):                             func() TLObject { return new(TLMessagesGetDhConfig) },
	int32(TLConstructor_CRC32_messages_requestEncryption):                       func() TLObject { return new(TLMessagesRequestEncryption) },
	int32(TLConstructor_CRC32_messages_acceptEncryption):                        func() TLObject { return new(TLMessagesAcceptEncryption) },
	int32(TLConstructor_CRC32_messages_discardEncryption):                       func() TLObject { return new(TLMessagesDiscardEncryption) },
	int32(TLConstructor_CRC32_messages_setEncryptedTyping):                      func() TLObject { return new(TLMessagesSetEncryptedTyping) },
	int32(TLConstructor_CRC32_messages_readEncryptedHistory):                    func() TLObject { return new(TLMessagesReadEncryptedHistory) },
	int32(TLConstructor_CRC32_messages_sendEncrypted):                           func() TLObject { return new(TLMessagesSendEncrypted) },
	int32(TLConstructor_CRC32_messages_sendEncryptedFile):                       func() TLObject { return new(TLMessagesSendEncryptedFile) },
	int32(TLConstructor_CRC32_messages_sendEncryptedService):                    func() TLObject { return new(TLMessagesSendEncryptedService) },
	int32(TLConstructor_CRC32_messages_receivedQueue):                           func() TLObject { return new(TLMessagesReceivedQueue) },
	int32(TLConstructor_CRC32_messages_reportEncryptedSpam):                     func() TLObject { return new(TLMessagesReportEncryptedSpam) },
	int32(TLConstructor_CRC32_messages_readMessageContents):                     func() TLObject { return new(TLMessagesReadMessageContents) },
	int32(TLConstructor_CRC32_messages_getAllStickers):                          func() TLObject { return new(TLMessagesGetAllStickers) },
	int32(TLConstructor_CRC32_messages_getWebPagePreview):                       func() TLObject { return new(TLMessagesGetWebPagePreview) },
	int32(TLConstructor_CRC32_messages_exportChatInvite):                        func() TLObject { return new(TLMessagesExportChatInvite) },
	int32(TLConstructor_CRC32_messages_checkChatInvite):                         func() TLObject { return new(TLMessagesCheckChatInvite) },
	int32(TLConstructor_CRC32_messages_importChatInvite):                        func() TLObject { return new(TLMessagesImportChatInvite) },
	int32(TLConstructor_CRC32_messages_getStickerSet):                           func() TLObject { return new(TLMessagesGetStickerSet) },
	int32(TLConstructor_CRC32_messages_installStickerSet):                       func() TLObject { return new(TLMessagesInstallStickerSet) },
	int32(TLConstructor_CRC32_messages_uninstallStickerSet):                     func() TLObject { return new(TLMessagesUninstallStickerSet) },
	int32(TLConstructor_CRC32_messages_startBot):                                func() TLObject { return new(TLMessagesStartBot) },
	int32(TLConstructor_CRC32_messages_getMessagesViews):                        func() TLObject { return new(TLMessagesGetMessagesViews) },
	int32(TLConstructor_CRC32_messages_toggleChatAdmins):                        func() TLObject { return new(TLMessagesToggleChatAdmins) },
	int32(TLConstructor_CRC32_messages_editChatAdmin):                           func() TLObject { return new(TLMessagesEditChatAdmin) },
	int32(TLConstructor_CRC32_messages_migrateChat):                             func() TLObject { return new(TLMessagesMigrateChat) },
	int32(TLConstructor_CRC32_messages_searchGlobal):                            func() TLObject { return new(TLMessagesSearchGlobal) },
	int32(TLConstructor_CRC32_messages_reorderStickerSets):                      func() TLObject { return new(TLMessagesReorderStickerSets) },
	int32(TLConstructor_CRC32_messages_getDocumentByHash):                       func() TLObject { return new(TLMessagesGetDocumentByHash) },
	int32(TLConstructor_CRC32_messages_searchGifs):                              func() TLObject { return new(TLMessagesSearchGifs) },
	int32(TLConstructor_CRC32_messages_getSavedGifs):                            func() TLObject { return new(TLMessagesGetSavedGifs) },
	int32(TLConstructor_CRC32_messages_saveGif):                                 func() TLObject { return new(TLMessagesSaveGif) },
	int32(TLConstructor_CRC32_messages_getInlineBotResults):                     func() TLObject { return new(TLMessagesGetInlineBotResults) },
	int32(TLConstructor_CRC32_messages_setInlineBotResults):                     func() TLObject { return new(TLMessagesSetInlineBotResults) },
	int32(TLConstructor_CRC32_messages_sendInlineBotResult):                     func() TLObject { return new(TLMessagesSendInlineBotResult) },
	int32(TLConstructor_CRC32_messages_getMessageEditData):                      func() TLObject { return new(TLMessagesGetMessageEditData) },
	int32(TLConstructor_CRC32_messages_editMessage):                             func() TLObject { return new(TLMessagesEditMessage) },
	int32(TLConstructor_CRC32_messages_editInlineBotMessage):                    func() TLObject { return new(TLMessagesEditInlineBotMessage) },
	int32(TLConstructor_CRC32_messages_getBotCallbackAnswer):                    func() TLObject { return new(TLMessagesGetBotCallbackAnswer) },
	int32(TLConstructor_CRC32_messages_setBotCallbackAnswer):                    func() TLObject { return new(TLMessagesSetBotCallbackAnswer) },
	int32(TLConstructor_CRC32_messages_getPeerDialogs):                          func() TLObject { return new(TLMessagesGetPeerDialogs) },
	int32(TLConstructor_CRC32_messages_saveDraft):                               func() TLObject { return new(TLMessagesSaveDraft) },
	int32(TLConstructor_CRC32_messages_getAllDrafts):                            func() TLObject { return new(TLMessagesGetAllDrafts) },
	int32(TLConstructor_CRC32_messages_getFeaturedStickers):                     func() TLObject { return new(TLMessagesGetFeaturedStickers) },
	int32(TLConstructor_CRC32_messages_readFeaturedStickers):                    func() TLObject { return new(TLMessagesReadFeaturedStickers) },
	int32(TLConstructor_CRC32_messages_getRecentStickers):                       func() TLObject { return new(TLMessagesGetRecentStickers) },
	int32(TLConstructor_CRC32_messages_saveRecentSticker):                       func() TLObject { return new(TLMessagesSaveRecentSticker) },
	int32(TLConstructor_CRC32_messages_clearRecentStickers):                     func() TLObject { return new(TLMessagesClearRecentStickers) },
	int32(TLConstructor_CRC32_messages_getArchivedStickers):                     func() TLObject { return new(TLMessagesGetArchivedStickers) },
	int32(TLConstructor_CRC32_messages_getMaskStickers):                         func() TLObject { return new(TLMessagesGetMaskStickers) },
	int32(TLConstructor_CRC32_messages_getAttachedStickers):                     func() TLObject { return new(TLMessagesGetAttachedStickers) },
	int32(TLConstructor_CRC32_messages_setGameScore):                            func() TLObject { return new(TLMessagesSetGameScore) },
	int32(TLConstructor_CRC32_messages_setInlineGameScore):                      func() TLObject { return new(TLMessagesSetInlineGameScore) },
	int32(TLConstructor_CRC32_messages_getGameHighScores):                       func() TLObject { return new(TLMessagesGetGameHighScores) },
	int32(TLConstructor_CRC32_messages_getInlineGameHighScores):                 func() TLObject { return new(TLMessagesGetInlineGameHighScores) },
	int32(TLConstructor_CRC32_messages_getCommonChats):                          func() TLObject { return new(TLMessagesGetCommonChats) },
	int32(TLConstructor_CRC32_messages_getAllChats):                             func() TLObject { return new(TLMessagesGetAllChats) },
	int32(TLConstructor_CRC32_messages_getWebPage):                              func() TLObject { return new(TLMessagesGetWebPage) },
	int32(TLConstructor_CRC32_messages_toggleDialogPin):                         func() TLObject { return new(TLMessagesToggleDialogPin) },
	int32(TLConstructor_CRC32_messages_reorderPinnedDialogs):                    func() TLObject { return new(TLMessagesReorderPinnedDialogs) },
	int32(TLConstructor_CRC32_messages_getPinnedDialogs):                        func() TLObject { return new(TLMessagesGetPinnedDialogs) },
	int32(TLConstructor_CRC32_messages_setBotShippingResults):                   func() TLObject { return new(TLMessagesSetBotShippingResults) },
	int32(TLConstructor_CRC32_messages_setBotPrecheckoutResults):                func() TLObject { return new(TLMessagesSetBotPrecheckoutResults) },
	int32(TLConstructor_CRC32_messages_uploadMedia):                             func() TLObject { return new(TLMessagesUploadMedia) },
	int32(TLConstructor_CRC32_messages_sendScreenshotNotification):              func() TLObject { return new(TLMessagesSendScreenshotNotification) },
	int32(TLConstructor_CRC32_messages_getFavedStickers):                        func() TLObject { return new(TLMessagesGetFavedStickers) },
	int32(TLConstructor_CRC32_messages_faveSticker):                             func() TLObject { return new(TLMessagesFaveSticker) },
	int32(TLConstructor_CRC32_messages_getUnreadMentions):                       func() TLObject { return new(TLMessagesGetUnreadMentions) },
	int32(TLConstructor_CRC32_updates_getState):                                 func() TLObject { return new(TLUpdatesGetState) },
	int32(TLConstructor_CRC32_updates_getDifference):                            func() TLObject { return new(TLUpdatesGetDifference) },
	int32(TLConstructor_CRC32_updates_getChannelDifference):                     func() TLObject { return new(TLUpdatesGetChannelDifference) },
	int32(TLConstructor_CRC32_photos_updateProfilePhoto):                        func() TLObject { return new(TLPhotosUpdateProfilePhoto) },
	int32(TLConstructor_CRC32_photos_uploadProfilePhoto):                        func() TLObject { return new(TLPhotosUploadProfilePhoto) },
	int32(TLConstructor_CRC32_photos_deletePhotos):                              func() TLObject { return new(TLPhotosDeletePhotos) },
	int32(TLConstructor_CRC32_photos_getUserPhotos):                             func() TLObject { return new(TLPhotosGetUserPhotos) },
	int32(TLConstructor_CRC32_upload_saveFilePart):                              func() TLObject { return new(TLUploadSaveFilePart) },
	int32(TLConstructor_CRC32_upload_getFile):                                   func() TLObject { return new(TLUploadGetFile) },
	int32(TLConstructor_CRC32_upload_saveBigFilePart):                           func() TLObject { return new(TLUploadSaveBigFilePart) },
	int32(TLConstructor_CRC32_upload_getWebFile):                                func() TLObject { return new(TLUploadGetWebFile) },
	int32(TLConstructor_CRC32_upload_getCdnFile):                                func() TLObject { return new(TLUploadGetCdnFile) },
	int32(TLConstructor_CRC32_upload_reuploadCdnFile):                           func() TLObject { return new(TLUploadReuploadCdnFile) },
	int32(TLConstructor_CRC32_upload_getCdnFileHashes):                          func() TLObject { return new(TLUploadGetCdnFileHashes) },
	int32(TLConstructor_CRC32_help_getConfig):                                   func() TLObject { return new(TLHelpGetConfig) },
	int32(TLConstructor_CRC32_help_getNearestDc):                                func() TLObject { return new(TLHelpGetNearestDc) },
	int32(TLConstructor_CRC32_help_getAppUpdate):                                func() TLObject { return new(TLHelpGetAppUpdate) },
	int32(TLConstructor_CRC32_help_saveAppLog):                                  func() TLObject { return new(TLHelpSaveAppLog) },
	int32(TLConstructor_CRC32_help_getInviteText):                               func() TLObject { return new(TLHelpGetInviteText) },
	int32(TLConstructor_CRC32_help_getSupport):                                  func() TLObject { return new(TLHelpGetSupport) },
	int32(TLConstructor_CRC32_help_getAppChangelog):                             func() TLObject { return new(TLHelpGetAppChangelog) },
	int32(TLConstructor_CRC32_help_getTermsOfService):                           func() TLObject { return new(TLHelpGetTermsOfService) },
	int32(TLConstructor_CRC32_help_setBotUpdatesStatus):                         func() TLObject { return new(TLHelpSetBotUpdatesStatus) },
	int32(TLConstructor_CRC32_help_getCdnConfig):                                func() TLObject { return new(TLHelpGetCdnConfig) },
	int32(TLConstructor_CRC32_channels_readHistory):                             func() TLObject { return new(TLChannelsReadHistory) },
	int32(TLConstructor_CRC32_channels_deleteMessages):                          func() TLObject { return new(TLChannelsDeleteMessages) },
	int32(TLConstructor_CRC32_channels_deleteUserHistory):                       func() TLObject { return new(TLChannelsDeleteUserHistory) },
	int32(TLConstructor_CRC32_channels_reportSpam):                              func() TLObject { return new(TLChannelsReportSpam) },
	int32(TLConstructor_CRC32_channels_getMessages):                             func() TLObject { return new(TLChannelsGetMessages) },
	int32(TLConstructor_CRC32_channels_getParticipants):                         func() TLObject { return new(TLChannelsGetParticipants) },
	int32(TLConstructor_CRC32_channels_getParticipant):                          func() TLObject { return new(TLChannelsGetParticipant) },
	int32(TLConstructor_CRC32_channels_getChannels):                             func() TLObject { return new(TLChannelsGetChannels) },
	int32(TLConstructor_CRC32_channels_getFullChannel):                          func() TLObject { return new(TLChannelsGetFullChannel) },
	int32(TLConstructor_CRC32_channels_createChannel):                           func() TLObject { return new(TLChannelsCreateChannel) },
	int32(TLConstructor_CRC32_channels_editAbout):                               func() TLObject { return new(TLChannelsEditAbout) },
	int32(TLConstructor_CRC32_channels_editAdmin):                               func() TLObject { return new(TLChannelsEditAdmin) },
	int32(TLConstructor_CRC32_channels_editTitle):                               func() TLObject { return new(TLChannelsEditTitle) },
	int32(TLConstructor_CRC32_channels_editPhoto):                               func() TLObject { return new(TLChannelsEditPhoto) },
	int32(TLConstructor_CRC32_channels_checkUsername):                           func() TLObject { return new(TLChannelsCheckUsername) },
	int32(TLConstructor_CRC32_channels_updateUsername):                          func() TLObject { return new(TLChannelsUpdateUsername) },
	int32(TLConstructor_CRC32_channels_joinChannel):                             func() TLObject { return new(TLChannelsJoinChannel) },
	int32(TLConstructor_CRC32_channels_leaveChannel):                            func() TLObject { return new(TLChannelsLeaveChannel) },
	int32(TLConstructor_CRC32_channels_inviteToChannel):                         func() TLObject { return new(TLChannelsInviteToChannel) },
	int32(TLConstructor_CRC32_channels_exportInvite):                            func() TLObject { return new(TLChannelsExportInvite) },
	int32(TLConstructor_CRC32_channels_deleteChannel):                           func() TLObject { return new(TLChannelsDeleteChannel) },
	int32(TLConstructor_CRC32_channels_toggleInvites):                           func() TLObject { return new(TLChannelsToggleInvites) },
	int32(TLConstructor_CRC32_channels_exportMessageLink):                       func() TLObject { return new(TLChannelsExportMessageLink) },
	int32(TLConstructor_CRC32_channels_toggleSignatures):                        func() TLObject { return new(TLChannelsToggleSignatures) },
	int32(TLConstructor_CRC32_channels_updatePinnedMessage):                     func() TLObject { return new(TLChannelsUpdatePinnedMessage) },
	int32(TLConstructor_CRC32_channels_getAdminedPublicChannels):                func() TLObject { return new(TLChannelsGetAdminedPublicChannels) },
	int32(TLConstructor_CRC32_channels_editBanned):                              func() TLObject { return new(TLChannelsEditBanned) },
	int32(TLConstructor_CRC32_channels_getAdminLog):                             func() TLObject { return new(TLChannelsGetAdminLog) },
	int32(TLConstructor_CRC32_channels_setStickers):                             func() TLObject { return new(TLChannelsSetStickers) },
	int32(TLConstructor_CRC32_channels_readMessageContents):                     func() TLObject { return new(TLChannelsReadMessageContents) },
	int32(TLConstructor_CRC32_bots_sendCustomRequest):                           func() TLObject { return new(TLBotsSendCustomRequest) },
	int32(TLConstructor_CRC32_bots_answerWebhookJSONQuery):                      func() TLObject { return new(TLBotsAnswerWebhookJSONQuery) },
	int32(TLConstructor_CRC32_payments_getPaymentForm):                          func() TLObject { return new(TLPaymentsGetPaymentForm) },
	int32(TLConstructor_CRC32_payments_getPaymentReceipt):                       func() TLObject { return new(TLPaymentsGetPaymentReceipt) },
	int32(TLConstructor_CRC32_payments_validateRequestedInfo):                   func() TLObject { return new(TLPaymentsValidateRequestedInfo) },
	int32(TLConstructor_CRC32_payments_sendPaymentForm):                         func() TLObject { return new(TLPaymentsSendPaymentForm) },
	int32(TLConstructor_CRC32_payments_getSavedInfo):                            func() TLObject { return new(TLPaymentsGetSavedInfo) },
	int32(TLConstructor_CRC32_payments_clearSavedInfo):                          func() TLObject { return new(TLPaymentsClearSavedInfo) },
	int32(TLConstructor_CRC32_stickers_createStickerSet):                        func() TLObject { return new(TLStickersCreateStickerSet) },
	int32(TLConstructor_CRC32_stickers_removeStickerFromSet):                    func() TLObject { return new(TLStickersRemoveStickerFromSet) },
	int32(TLConstructor_CRC32_stickers_changeStickerPosition):                   func() TLObject { return new(TLStickersChangeStickerPosition) },
	int32(TLConstructor_CRC32_stickers_addStickerToSet):                         func() TLObject { return new(TLStickersAddStickerToSet) },
	int32(TLConstructor_CRC32_phone_getCallConfig):                              func() TLObject { return new(TLPhoneGetCallConfig) },
	int32(TLConstructor_CRC32_phone_requestCall):                                func() TLObject { return new(TLPhoneRequestCall) },
	int32(TLConstructor_CRC32_phone_acceptCall):                                 func() TLObject { return new(TLPhoneAcceptCall) },
	int32(TLConstructor_CRC32_phone_confirmCall):                                func() TLObject { return new(TLPhoneConfirmCall) },
	int32(TLConstructor_CRC32_phone_receivedCall):                               func() TLObject { return new(TLPhoneReceivedCall) },
	int32(TLConstructor_CRC32_phone_discardCall):                                func() TLObject { return new(TLPhoneDiscardCall) },
	int32(TLConstructor_CRC32_phone_setCallRating):                              func() TLObject { return new(TLPhoneSetCallRating) },
	int32(TLConstructor_CRC32_phone_saveCallDebug):                              func() TLObject { return new(TLPhoneSaveCallDebug) },
	int32(TLConstructor_CRC32_langpack_getLangPack):                             func() TLObject { return new(TLLangpackGetLangPack) },
	int32(TLConstructor_CRC32_langpack_getStrings):                              func() TLObject { return new(TLLangpackGetStrings) },
	int32(TLConstructor_CRC32_langpack_getDifference):                           func() TLObject { return new(TLLangpackGetDifference) },
	int32(TLConstructor_CRC32_langpack_getLanguages):                            func() TLObject { return new(TLLangpackGetLanguages) },
}

func NewTLObjectByClassID(classId int32) TLObject {
	m, ok := registers2[classId]
	if !ok {
		return nil
	}
	return m()
}

//////////////////////////////////////////////////////////////////////////////////////////
func (m *ResPQ) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ResPQ_ResPQ:
		m2, _ := m.Payload.(*ResPQ_ResPQ)
		b = m2.ResPQ.Encode()
	}
	return
}

func (m *ResPQ) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_resPQ):
		m2 := ResPQ_ResPQ{}
		m2.ResPQ.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *P_QInnerData) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *P_QInnerData_PQInnerData:
		m2, _ := m.Payload.(*P_QInnerData_PQInnerData)
		b = m2.PQInnerData.Encode()
	}
	return
}

func (m *P_QInnerData) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_p_q_inner_data):
		m2 := P_QInnerData_PQInnerData{}
		m2.PQInnerData.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Server_DH_Params) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Server_DH_Params_Server_DHParamsFail:
		m2, _ := m.Payload.(*Server_DH_Params_Server_DHParamsFail)
		b = m2.Server_DHParamsFail.Encode()
	case *Server_DH_Params_Server_DHParamsOk:
		m2, _ := m.Payload.(*Server_DH_Params_Server_DHParamsOk)
		b = m2.Server_DHParamsOk.Encode()
	}
	return
}

func (m *Server_DH_Params) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_server_DH_params_fail):
		m2 := Server_DH_Params_Server_DHParamsFail{}
		m2.Server_DHParamsFail.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_server_DH_params_ok):
		m2 := Server_DH_Params_Server_DHParamsOk{}
		m2.Server_DHParamsOk.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Server_DHInnerData) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Server_DHInnerData_Server_DHInnerData:
		m2, _ := m.Payload.(*Server_DHInnerData_Server_DHInnerData)
		b = m2.Server_DHInnerData.Encode()
	}
	return
}

func (m *Server_DHInnerData) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_server_DH_inner_data):
		m2 := Server_DHInnerData_Server_DHInnerData{}
		m2.Server_DHInnerData.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Client_DH_Inner_Data) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Client_DH_Inner_Data_Client_DHInnerData:
		m2, _ := m.Payload.(*Client_DH_Inner_Data_Client_DHInnerData)
		b = m2.Client_DHInnerData.Encode()
	}
	return
}

func (m *Client_DH_Inner_Data) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_client_DH_inner_data):
		m2 := Client_DH_Inner_Data_Client_DHInnerData{}
		m2.Client_DHInnerData.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *SetClient_DHParamsAnswer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *SetClient_DHParamsAnswer_DhGenOk:
		m2, _ := m.Payload.(*SetClient_DHParamsAnswer_DhGenOk)
		b = m2.DhGenOk.Encode()
	case *SetClient_DHParamsAnswer_DhGenRetry:
		m2, _ := m.Payload.(*SetClient_DHParamsAnswer_DhGenRetry)
		b = m2.DhGenRetry.Encode()
	case *SetClient_DHParamsAnswer_DhGenFail:
		m2, _ := m.Payload.(*SetClient_DHParamsAnswer_DhGenFail)
		b = m2.DhGenFail.Encode()
	}
	return
}

func (m *SetClient_DHParamsAnswer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_dh_gen_ok):
		m2 := SetClient_DHParamsAnswer_DhGenOk{}
		m2.DhGenOk.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_dh_gen_retry):
		m2 := SetClient_DHParamsAnswer_DhGenRetry{}
		m2.DhGenRetry.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_dh_gen_fail):
		m2 := SetClient_DHParamsAnswer_DhGenFail{}
		m2.DhGenFail.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DestroyAuthKeyRes) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DestroyAuthKeyRes_DestroyAuthKeyOk:
		m2, _ := m.Payload.(*DestroyAuthKeyRes_DestroyAuthKeyOk)
		b = m2.DestroyAuthKeyOk.Encode()
	case *DestroyAuthKeyRes_DestroyAuthKeyNone:
		m2, _ := m.Payload.(*DestroyAuthKeyRes_DestroyAuthKeyNone)
		b = m2.DestroyAuthKeyNone.Encode()
	case *DestroyAuthKeyRes_DestroyAuthKeyFail:
		m2, _ := m.Payload.(*DestroyAuthKeyRes_DestroyAuthKeyFail)
		b = m2.DestroyAuthKeyFail.Encode()
	}
	return
}

func (m *DestroyAuthKeyRes) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_destroy_auth_key_ok):
		m2 := DestroyAuthKeyRes_DestroyAuthKeyOk{}
		m2.DestroyAuthKeyOk.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_destroy_auth_key_none):
		m2 := DestroyAuthKeyRes_DestroyAuthKeyNone{}
		m2.DestroyAuthKeyNone.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_destroy_auth_key_fail):
		m2 := DestroyAuthKeyRes_DestroyAuthKeyFail{}
		m2.DestroyAuthKeyFail.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MsgsAck) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MsgsAck_MsgsAck:
		m2, _ := m.Payload.(*MsgsAck_MsgsAck)
		b = m2.MsgsAck.Encode()
	}
	return
}

func (m *MsgsAck) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_msgs_ack):
		m2 := MsgsAck_MsgsAck{}
		m2.MsgsAck.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *BadMsgNotification) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *BadMsgNotification_BadMsgNotification:
		m2, _ := m.Payload.(*BadMsgNotification_BadMsgNotification)
		b = m2.BadMsgNotification.Encode()
	case *BadMsgNotification_BadServerSalt:
		m2, _ := m.Payload.(*BadMsgNotification_BadServerSalt)
		b = m2.BadServerSalt.Encode()
	}
	return
}

func (m *BadMsgNotification) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_bad_msg_notification):
		m2 := BadMsgNotification_BadMsgNotification{}
		m2.BadMsgNotification.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_bad_server_salt):
		m2 := BadMsgNotification_BadServerSalt{}
		m2.BadServerSalt.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MsgsStateReq) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MsgsStateReq_MsgsStateReq:
		m2, _ := m.Payload.(*MsgsStateReq_MsgsStateReq)
		b = m2.MsgsStateReq.Encode()
	}
	return
}

func (m *MsgsStateReq) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_msgs_state_req):
		m2 := MsgsStateReq_MsgsStateReq{}
		m2.MsgsStateReq.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MsgsStateInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MsgsStateInfo_MsgsStateInfo:
		m2, _ := m.Payload.(*MsgsStateInfo_MsgsStateInfo)
		b = m2.MsgsStateInfo.Encode()
	}
	return
}

func (m *MsgsStateInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_msgs_state_info):
		m2 := MsgsStateInfo_MsgsStateInfo{}
		m2.MsgsStateInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MsgsAllInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MsgsAllInfo_MsgsAllInfo:
		m2, _ := m.Payload.(*MsgsAllInfo_MsgsAllInfo)
		b = m2.MsgsAllInfo.Encode()
	}
	return
}

func (m *MsgsAllInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_msgs_all_info):
		m2 := MsgsAllInfo_MsgsAllInfo{}
		m2.MsgsAllInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MsgDetailedInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MsgDetailedInfo_MsgDetailedInfo:
		m2, _ := m.Payload.(*MsgDetailedInfo_MsgDetailedInfo)
		b = m2.MsgDetailedInfo.Encode()
	case *MsgDetailedInfo_MsgNewDetailedInfo:
		m2, _ := m.Payload.(*MsgDetailedInfo_MsgNewDetailedInfo)
		b = m2.MsgNewDetailedInfo.Encode()
	}
	return
}

func (m *MsgDetailedInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_msg_detailed_info):
		m2 := MsgDetailedInfo_MsgDetailedInfo{}
		m2.MsgDetailedInfo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_msg_new_detailed_info):
		m2 := MsgDetailedInfo_MsgNewDetailedInfo{}
		m2.MsgNewDetailedInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MsgResendReq) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MsgResendReq_MsgResendReq:
		m2, _ := m.Payload.(*MsgResendReq_MsgResendReq)
		b = m2.MsgResendReq.Encode()
	}
	return
}

func (m *MsgResendReq) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_msg_resend_req):
		m2 := MsgResendReq_MsgResendReq{}
		m2.MsgResendReq.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *RpcError) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *RpcError_RpcError:
		m2, _ := m.Payload.(*RpcError_RpcError)
		b = m2.RpcError.Encode()
	}
	return
}

func (m *RpcError) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_rpc_error):
		m2 := RpcError_RpcError{}
		m2.RpcError.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *RpcDropAnswer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *RpcDropAnswer_RpcAnswerUnknown:
		m2, _ := m.Payload.(*RpcDropAnswer_RpcAnswerUnknown)
		b = m2.RpcAnswerUnknown.Encode()
	case *RpcDropAnswer_RpcAnswerDroppedRunning:
		m2, _ := m.Payload.(*RpcDropAnswer_RpcAnswerDroppedRunning)
		b = m2.RpcAnswerDroppedRunning.Encode()
	case *RpcDropAnswer_RpcAnswerDropped:
		m2, _ := m.Payload.(*RpcDropAnswer_RpcAnswerDropped)
		b = m2.RpcAnswerDropped.Encode()
	}
	return
}

func (m *RpcDropAnswer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_rpc_answer_unknown):
		m2 := RpcDropAnswer_RpcAnswerUnknown{}
		m2.RpcAnswerUnknown.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_rpc_answer_dropped_running):
		m2 := RpcDropAnswer_RpcAnswerDroppedRunning{}
		m2.RpcAnswerDroppedRunning.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_rpc_answer_dropped):
		m2 := RpcDropAnswer_RpcAnswerDropped{}
		m2.RpcAnswerDropped.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *FutureSalt) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *FutureSalt_FutureSalt:
		m2, _ := m.Payload.(*FutureSalt_FutureSalt)
		b = m2.FutureSalt.Encode()
	}
	return
}

func (m *FutureSalt) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_future_salt):
		m2 := FutureSalt_FutureSalt{}
		m2.FutureSalt.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *FutureSalts) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *FutureSalts_FutureSalts:
		m2, _ := m.Payload.(*FutureSalts_FutureSalts)
		b = m2.FutureSalts.Encode()
	}
	return
}

func (m *FutureSalts) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_future_salts):
		m2 := FutureSalts_FutureSalts{}
		m2.FutureSalts.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Pong) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Pong_Pong:
		m2, _ := m.Payload.(*Pong_Pong)
		b = m2.Pong.Encode()
	}
	return
}

func (m *Pong) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_pong):
		m2 := Pong_Pong{}
		m2.Pong.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DestroySessionRes) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DestroySessionRes_DestroySessionOk:
		m2, _ := m.Payload.(*DestroySessionRes_DestroySessionOk)
		b = m2.DestroySessionOk.Encode()
	case *DestroySessionRes_DestroySessionNone:
		m2, _ := m.Payload.(*DestroySessionRes_DestroySessionNone)
		b = m2.DestroySessionNone.Encode()
	}
	return
}

func (m *DestroySessionRes) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_destroy_session_ok):
		m2 := DestroySessionRes_DestroySessionOk{}
		m2.DestroySessionOk.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_destroy_session_none):
		m2 := DestroySessionRes_DestroySessionNone{}
		m2.DestroySessionNone.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *NewSession) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *NewSession_NewSessionCreated:
		m2, _ := m.Payload.(*NewSession_NewSessionCreated)
		b = m2.NewSessionCreated.Encode()
	}
	return
}

func (m *NewSession) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_new_session_created):
		m2 := NewSession_NewSessionCreated{}
		m2.NewSessionCreated.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *HttpWait) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *HttpWait_HttpWait:
		m2, _ := m.Payload.(*HttpWait_HttpWait)
		b = m2.HttpWait.Encode()
	}
	return
}

func (m *HttpWait) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_http_wait):
		m2 := HttpWait_HttpWait{}
		m2.HttpWait.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *IpPort) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *IpPort_IpPort:
		m2, _ := m.Payload.(*IpPort_IpPort)
		b = m2.IpPort.Encode()
	}
	return
}

func (m *IpPort) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_ipPort):
		m2 := IpPort_IpPort{}
		m2.IpPort.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Help_ConfigSimple) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Help_ConfigSimple_HelpConfigSimple:
		m2, _ := m.Payload.(*Help_ConfigSimple_HelpConfigSimple)
		b = m2.HelpConfigSimple.Encode()
	}
	return
}

func (m *Help_ConfigSimple) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_help_configSimple):
		m2 := Help_ConfigSimple_HelpConfigSimple{}
		m2.HelpConfigSimple.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Bool) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Bool_BoolFalse:
		m2, _ := m.Payload.(*Bool_BoolFalse)
		b = m2.BoolFalse.Encode()
	case *Bool_BoolTrue:
		m2, _ := m.Payload.(*Bool_BoolTrue)
		b = m2.BoolTrue.Encode()
	}
	return
}

func (m *Bool) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_boolFalse):
		m2 := Bool_BoolFalse{}
		m2.BoolFalse.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_boolTrue):
		m2 := Bool_BoolTrue{}
		m2.BoolTrue.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *True) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *True_True:
		m2, _ := m.Payload.(*True_True)
		b = m2.True.Encode()
	}
	return
}

func (m *True) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_true):
		m2 := True_True{}
		m2.True.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Error) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Error_Error:
		m2, _ := m.Payload.(*Error_Error)
		b = m2.Error.Encode()
	}
	return
}

func (m *Error) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_error):
		m2 := Error_Error{}
		m2.Error.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Null) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Null_Null:
		m2, _ := m.Payload.(*Null_Null)
		b = m2.Null.Encode()
	}
	return
}

func (m *Null) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_null):
		m2 := Null_Null{}
		m2.Null.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPeer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPeer_InputPeerEmpty:
		m2, _ := m.Payload.(*InputPeer_InputPeerEmpty)
		b = m2.InputPeerEmpty.Encode()
	case *InputPeer_InputPeerSelf:
		m2, _ := m.Payload.(*InputPeer_InputPeerSelf)
		b = m2.InputPeerSelf.Encode()
	case *InputPeer_InputPeerChat:
		m2, _ := m.Payload.(*InputPeer_InputPeerChat)
		b = m2.InputPeerChat.Encode()
	case *InputPeer_InputPeerUser:
		m2, _ := m.Payload.(*InputPeer_InputPeerUser)
		b = m2.InputPeerUser.Encode()
	case *InputPeer_InputPeerChannel:
		m2, _ := m.Payload.(*InputPeer_InputPeerChannel)
		b = m2.InputPeerChannel.Encode()
	}
	return
}

func (m *InputPeer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPeerEmpty):
		m2 := InputPeer_InputPeerEmpty{}
		m2.InputPeerEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPeerSelf):
		m2 := InputPeer_InputPeerSelf{}
		m2.InputPeerSelf.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPeerChat):
		m2 := InputPeer_InputPeerChat{}
		m2.InputPeerChat.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPeerUser):
		m2 := InputPeer_InputPeerUser{}
		m2.InputPeerUser.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPeerChannel):
		m2 := InputPeer_InputPeerChannel{}
		m2.InputPeerChannel.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputUser) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputUser_InputUserEmpty:
		m2, _ := m.Payload.(*InputUser_InputUserEmpty)
		b = m2.InputUserEmpty.Encode()
	case *InputUser_InputUserSelf:
		m2, _ := m.Payload.(*InputUser_InputUserSelf)
		b = m2.InputUserSelf.Encode()
	case *InputUser_InputUser:
		m2, _ := m.Payload.(*InputUser_InputUser)
		b = m2.InputUser.Encode()
	}
	return
}

func (m *InputUser) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputUserEmpty):
		m2 := InputUser_InputUserEmpty{}
		m2.InputUserEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputUserSelf):
		m2 := InputUser_InputUserSelf{}
		m2.InputUserSelf.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputUser):
		m2 := InputUser_InputUser{}
		m2.InputUser.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputContact) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputContact_InputPhoneContact:
		m2, _ := m.Payload.(*InputContact_InputPhoneContact)
		b = m2.InputPhoneContact.Encode()
	}
	return
}

func (m *InputContact) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPhoneContact):
		m2 := InputContact_InputPhoneContact{}
		m2.InputPhoneContact.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputFile) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputFile_InputFile:
		m2, _ := m.Payload.(*InputFile_InputFile)
		b = m2.InputFile.Encode()
	case *InputFile_InputFileBig:
		m2, _ := m.Payload.(*InputFile_InputFileBig)
		b = m2.InputFileBig.Encode()
	}
	return
}

func (m *InputFile) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputFile):
		m2 := InputFile_InputFile{}
		m2.InputFile.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputFileBig):
		m2 := InputFile_InputFileBig{}
		m2.InputFileBig.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputMedia) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputMedia_InputMediaEmpty:
		m2, _ := m.Payload.(*InputMedia_InputMediaEmpty)
		b = m2.InputMediaEmpty.Encode()
	case *InputMedia_InputMediaUploadedPhoto:
		m2, _ := m.Payload.(*InputMedia_InputMediaUploadedPhoto)
		b = m2.InputMediaUploadedPhoto.Encode()
	case *InputMedia_InputMediaPhoto:
		m2, _ := m.Payload.(*InputMedia_InputMediaPhoto)
		b = m2.InputMediaPhoto.Encode()
	case *InputMedia_InputMediaGeoPoint:
		m2, _ := m.Payload.(*InputMedia_InputMediaGeoPoint)
		b = m2.InputMediaGeoPoint.Encode()
	case *InputMedia_InputMediaContact:
		m2, _ := m.Payload.(*InputMedia_InputMediaContact)
		b = m2.InputMediaContact.Encode()
	case *InputMedia_InputMediaUploadedDocument:
		m2, _ := m.Payload.(*InputMedia_InputMediaUploadedDocument)
		b = m2.InputMediaUploadedDocument.Encode()
	case *InputMedia_InputMediaDocument:
		m2, _ := m.Payload.(*InputMedia_InputMediaDocument)
		b = m2.InputMediaDocument.Encode()
	case *InputMedia_InputMediaVenue:
		m2, _ := m.Payload.(*InputMedia_InputMediaVenue)
		b = m2.InputMediaVenue.Encode()
	case *InputMedia_InputMediaGifExternal:
		m2, _ := m.Payload.(*InputMedia_InputMediaGifExternal)
		b = m2.InputMediaGifExternal.Encode()
	case *InputMedia_InputMediaPhotoExternal:
		m2, _ := m.Payload.(*InputMedia_InputMediaPhotoExternal)
		b = m2.InputMediaPhotoExternal.Encode()
	case *InputMedia_InputMediaDocumentExternal:
		m2, _ := m.Payload.(*InputMedia_InputMediaDocumentExternal)
		b = m2.InputMediaDocumentExternal.Encode()
	case *InputMedia_InputMediaGame:
		m2, _ := m.Payload.(*InputMedia_InputMediaGame)
		b = m2.InputMediaGame.Encode()
	case *InputMedia_InputMediaInvoice:
		m2, _ := m.Payload.(*InputMedia_InputMediaInvoice)
		b = m2.InputMediaInvoice.Encode()
	}
	return
}

func (m *InputMedia) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputMediaEmpty):
		m2 := InputMedia_InputMediaEmpty{}
		m2.InputMediaEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaUploadedPhoto):
		m2 := InputMedia_InputMediaUploadedPhoto{}
		m2.InputMediaUploadedPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaPhoto):
		m2 := InputMedia_InputMediaPhoto{}
		m2.InputMediaPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaGeoPoint):
		m2 := InputMedia_InputMediaGeoPoint{}
		m2.InputMediaGeoPoint.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaContact):
		m2 := InputMedia_InputMediaContact{}
		m2.InputMediaContact.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaUploadedDocument):
		m2 := InputMedia_InputMediaUploadedDocument{}
		m2.InputMediaUploadedDocument.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaDocument):
		m2 := InputMedia_InputMediaDocument{}
		m2.InputMediaDocument.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaVenue):
		m2 := InputMedia_InputMediaVenue{}
		m2.InputMediaVenue.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaGifExternal):
		m2 := InputMedia_InputMediaGifExternal{}
		m2.InputMediaGifExternal.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaPhotoExternal):
		m2 := InputMedia_InputMediaPhotoExternal{}
		m2.InputMediaPhotoExternal.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaDocumentExternal):
		m2 := InputMedia_InputMediaDocumentExternal{}
		m2.InputMediaDocumentExternal.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaGame):
		m2 := InputMedia_InputMediaGame{}
		m2.InputMediaGame.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMediaInvoice):
		m2 := InputMedia_InputMediaInvoice{}
		m2.InputMediaInvoice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputChatPhoto) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputChatPhoto_InputChatPhotoEmpty:
		m2, _ := m.Payload.(*InputChatPhoto_InputChatPhotoEmpty)
		b = m2.InputChatPhotoEmpty.Encode()
	case *InputChatPhoto_InputChatUploadedPhoto:
		m2, _ := m.Payload.(*InputChatPhoto_InputChatUploadedPhoto)
		b = m2.InputChatUploadedPhoto.Encode()
	case *InputChatPhoto_InputChatPhoto:
		m2, _ := m.Payload.(*InputChatPhoto_InputChatPhoto)
		b = m2.InputChatPhoto.Encode()
	}
	return
}

func (m *InputChatPhoto) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputChatPhotoEmpty):
		m2 := InputChatPhoto_InputChatPhotoEmpty{}
		m2.InputChatPhotoEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputChatUploadedPhoto):
		m2 := InputChatPhoto_InputChatUploadedPhoto{}
		m2.InputChatUploadedPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputChatPhoto):
		m2 := InputChatPhoto_InputChatPhoto{}
		m2.InputChatPhoto.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputGeoPoint) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputGeoPoint_InputGeoPointEmpty:
		m2, _ := m.Payload.(*InputGeoPoint_InputGeoPointEmpty)
		b = m2.InputGeoPointEmpty.Encode()
	case *InputGeoPoint_InputGeoPoint:
		m2, _ := m.Payload.(*InputGeoPoint_InputGeoPoint)
		b = m2.InputGeoPoint.Encode()
	}
	return
}

func (m *InputGeoPoint) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputGeoPointEmpty):
		m2 := InputGeoPoint_InputGeoPointEmpty{}
		m2.InputGeoPointEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputGeoPoint):
		m2 := InputGeoPoint_InputGeoPoint{}
		m2.InputGeoPoint.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPhoto) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPhoto_InputPhotoEmpty:
		m2, _ := m.Payload.(*InputPhoto_InputPhotoEmpty)
		b = m2.InputPhotoEmpty.Encode()
	case *InputPhoto_InputPhoto:
		m2, _ := m.Payload.(*InputPhoto_InputPhoto)
		b = m2.InputPhoto.Encode()
	}
	return
}

func (m *InputPhoto) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPhotoEmpty):
		m2 := InputPhoto_InputPhotoEmpty{}
		m2.InputPhotoEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPhoto):
		m2 := InputPhoto_InputPhoto{}
		m2.InputPhoto.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputFileLocation) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputFileLocation_InputFileLocation:
		m2, _ := m.Payload.(*InputFileLocation_InputFileLocation)
		b = m2.InputFileLocation.Encode()
	case *InputFileLocation_InputEncryptedFileLocation:
		m2, _ := m.Payload.(*InputFileLocation_InputEncryptedFileLocation)
		b = m2.InputEncryptedFileLocation.Encode()
	case *InputFileLocation_InputDocumentFileLocation:
		m2, _ := m.Payload.(*InputFileLocation_InputDocumentFileLocation)
		b = m2.InputDocumentFileLocation.Encode()
	}
	return
}

func (m *InputFileLocation) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputFileLocation):
		m2 := InputFileLocation_InputFileLocation{}
		m2.InputFileLocation.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputEncryptedFileLocation):
		m2 := InputFileLocation_InputEncryptedFileLocation{}
		m2.InputEncryptedFileLocation.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputDocumentFileLocation):
		m2 := InputFileLocation_InputDocumentFileLocation{}
		m2.InputDocumentFileLocation.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputAppEvent) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputAppEvent_InputAppEvent:
		m2, _ := m.Payload.(*InputAppEvent_InputAppEvent)
		b = m2.InputAppEvent.Encode()
	}
	return
}

func (m *InputAppEvent) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputAppEvent):
		m2 := InputAppEvent_InputAppEvent{}
		m2.InputAppEvent.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Peer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Peer_PeerUser:
		m2, _ := m.Payload.(*Peer_PeerUser)
		b = m2.PeerUser.Encode()
	case *Peer_PeerChat:
		m2, _ := m.Payload.(*Peer_PeerChat)
		b = m2.PeerChat.Encode()
	case *Peer_PeerChannel:
		m2, _ := m.Payload.(*Peer_PeerChannel)
		b = m2.PeerChannel.Encode()
	}
	return
}

func (m *Peer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_peerUser):
		m2 := Peer_PeerUser{}
		m2.PeerUser.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_peerChat):
		m2 := Peer_PeerChat{}
		m2.PeerChat.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_peerChannel):
		m2 := Peer_PeerChannel{}
		m2.PeerChannel.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Storage_FileType) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Storage_FileType_StorageFileUnknown:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileUnknown)
		b = m2.StorageFileUnknown.Encode()
	case *Storage_FileType_StorageFilePartial:
		m2, _ := m.Payload.(*Storage_FileType_StorageFilePartial)
		b = m2.StorageFilePartial.Encode()
	case *Storage_FileType_StorageFileJpeg:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileJpeg)
		b = m2.StorageFileJpeg.Encode()
	case *Storage_FileType_StorageFileGif:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileGif)
		b = m2.StorageFileGif.Encode()
	case *Storage_FileType_StorageFilePng:
		m2, _ := m.Payload.(*Storage_FileType_StorageFilePng)
		b = m2.StorageFilePng.Encode()
	case *Storage_FileType_StorageFilePdf:
		m2, _ := m.Payload.(*Storage_FileType_StorageFilePdf)
		b = m2.StorageFilePdf.Encode()
	case *Storage_FileType_StorageFileMp3:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileMp3)
		b = m2.StorageFileMp3.Encode()
	case *Storage_FileType_StorageFileMov:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileMov)
		b = m2.StorageFileMov.Encode()
	case *Storage_FileType_StorageFileMp4:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileMp4)
		b = m2.StorageFileMp4.Encode()
	case *Storage_FileType_StorageFileWebp:
		m2, _ := m.Payload.(*Storage_FileType_StorageFileWebp)
		b = m2.StorageFileWebp.Encode()
	}
	return
}

func (m *Storage_FileType) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_storage_fileUnknown):
		m2 := Storage_FileType_StorageFileUnknown{}
		m2.StorageFileUnknown.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_filePartial):
		m2 := Storage_FileType_StorageFilePartial{}
		m2.StorageFilePartial.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_fileJpeg):
		m2 := Storage_FileType_StorageFileJpeg{}
		m2.StorageFileJpeg.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_fileGif):
		m2 := Storage_FileType_StorageFileGif{}
		m2.StorageFileGif.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_filePng):
		m2 := Storage_FileType_StorageFilePng{}
		m2.StorageFilePng.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_filePdf):
		m2 := Storage_FileType_StorageFilePdf{}
		m2.StorageFilePdf.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_fileMp3):
		m2 := Storage_FileType_StorageFileMp3{}
		m2.StorageFileMp3.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_fileMov):
		m2 := Storage_FileType_StorageFileMov{}
		m2.StorageFileMov.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_fileMp4):
		m2 := Storage_FileType_StorageFileMp4{}
		m2.StorageFileMp4.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_storage_fileWebp):
		m2 := Storage_FileType_StorageFileWebp{}
		m2.StorageFileWebp.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *FileLocation) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *FileLocation_FileLocationUnavailable:
		m2, _ := m.Payload.(*FileLocation_FileLocationUnavailable)
		b = m2.FileLocationUnavailable.Encode()
	case *FileLocation_FileLocation:
		m2, _ := m.Payload.(*FileLocation_FileLocation)
		b = m2.FileLocation.Encode()
	}
	return
}

func (m *FileLocation) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_fileLocationUnavailable):
		m2 := FileLocation_FileLocationUnavailable{}
		m2.FileLocationUnavailable.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_fileLocation):
		m2 := FileLocation_FileLocation{}
		m2.FileLocation.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *User) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *User_UserEmpty:
		m2, _ := m.Payload.(*User_UserEmpty)
		b = m2.UserEmpty.Encode()
	case *User_User:
		m2, _ := m.Payload.(*User_User)
		b = m2.User.Encode()
	}
	return
}

func (m *User) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_userEmpty):
		m2 := User_UserEmpty{}
		m2.UserEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_user):
		m2 := User_User{}
		m2.User.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *UserProfilePhoto) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *UserProfilePhoto_UserProfilePhotoEmpty:
		m2, _ := m.Payload.(*UserProfilePhoto_UserProfilePhotoEmpty)
		b = m2.UserProfilePhotoEmpty.Encode()
	case *UserProfilePhoto_UserProfilePhoto:
		m2, _ := m.Payload.(*UserProfilePhoto_UserProfilePhoto)
		b = m2.UserProfilePhoto.Encode()
	}
	return
}

func (m *UserProfilePhoto) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_userProfilePhotoEmpty):
		m2 := UserProfilePhoto_UserProfilePhotoEmpty{}
		m2.UserProfilePhotoEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_userProfilePhoto):
		m2 := UserProfilePhoto_UserProfilePhoto{}
		m2.UserProfilePhoto.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *UserStatus) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *UserStatus_UserStatusEmpty:
		m2, _ := m.Payload.(*UserStatus_UserStatusEmpty)
		b = m2.UserStatusEmpty.Encode()
	case *UserStatus_UserStatusOnline:
		m2, _ := m.Payload.(*UserStatus_UserStatusOnline)
		b = m2.UserStatusOnline.Encode()
	case *UserStatus_UserStatusOffline:
		m2, _ := m.Payload.(*UserStatus_UserStatusOffline)
		b = m2.UserStatusOffline.Encode()
	case *UserStatus_UserStatusRecently:
		m2, _ := m.Payload.(*UserStatus_UserStatusRecently)
		b = m2.UserStatusRecently.Encode()
	case *UserStatus_UserStatusLastWeek:
		m2, _ := m.Payload.(*UserStatus_UserStatusLastWeek)
		b = m2.UserStatusLastWeek.Encode()
	case *UserStatus_UserStatusLastMonth:
		m2, _ := m.Payload.(*UserStatus_UserStatusLastMonth)
		b = m2.UserStatusLastMonth.Encode()
	}
	return
}

func (m *UserStatus) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_userStatusEmpty):
		m2 := UserStatus_UserStatusEmpty{}
		m2.UserStatusEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_userStatusOnline):
		m2 := UserStatus_UserStatusOnline{}
		m2.UserStatusOnline.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_userStatusOffline):
		m2 := UserStatus_UserStatusOffline{}
		m2.UserStatusOffline.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_userStatusRecently):
		m2 := UserStatus_UserStatusRecently{}
		m2.UserStatusRecently.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_userStatusLastWeek):
		m2 := UserStatus_UserStatusLastWeek{}
		m2.UserStatusLastWeek.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_userStatusLastMonth):
		m2 := UserStatus_UserStatusLastMonth{}
		m2.UserStatusLastMonth.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Chat) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Chat_ChatEmpty:
		m2, _ := m.Payload.(*Chat_ChatEmpty)
		b = m2.ChatEmpty.Encode()
	case *Chat_Chat:
		m2, _ := m.Payload.(*Chat_Chat)
		b = m2.Chat.Encode()
	case *Chat_ChatForbidden:
		m2, _ := m.Payload.(*Chat_ChatForbidden)
		b = m2.ChatForbidden.Encode()
	case *Chat_Channel:
		m2, _ := m.Payload.(*Chat_Channel)
		b = m2.Channel.Encode()
	case *Chat_ChannelForbidden:
		m2, _ := m.Payload.(*Chat_ChannelForbidden)
		b = m2.ChannelForbidden.Encode()
	}
	return
}

func (m *Chat) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatEmpty):
		m2 := Chat_ChatEmpty{}
		m2.ChatEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chat):
		m2 := Chat_Chat{}
		m2.Chat.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatForbidden):
		m2 := Chat_ChatForbidden{}
		m2.ChatForbidden.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channel):
		m2 := Chat_Channel{}
		m2.Channel.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelForbidden):
		m2 := Chat_ChannelForbidden{}
		m2.ChannelForbidden.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChatFull) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChatFull_ChatFull:
		m2, _ := m.Payload.(*ChatFull_ChatFull)
		b = m2.ChatFull.Encode()
	case *ChatFull_ChannelFull:
		m2, _ := m.Payload.(*ChatFull_ChannelFull)
		b = m2.ChannelFull.Encode()
	}
	return
}

func (m *ChatFull) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatFull):
		m2 := ChatFull_ChatFull{}
		m2.ChatFull.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelFull):
		m2 := ChatFull_ChannelFull{}
		m2.ChannelFull.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChatParticipant) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChatParticipant_ChatParticipant:
		m2, _ := m.Payload.(*ChatParticipant_ChatParticipant)
		b = m2.ChatParticipant.Encode()
	case *ChatParticipant_ChatParticipantCreator:
		m2, _ := m.Payload.(*ChatParticipant_ChatParticipantCreator)
		b = m2.ChatParticipantCreator.Encode()
	case *ChatParticipant_ChatParticipantAdmin:
		m2, _ := m.Payload.(*ChatParticipant_ChatParticipantAdmin)
		b = m2.ChatParticipantAdmin.Encode()
	}
	return
}

func (m *ChatParticipant) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatParticipant):
		m2 := ChatParticipant_ChatParticipant{}
		m2.ChatParticipant.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatParticipantCreator):
		m2 := ChatParticipant_ChatParticipantCreator{}
		m2.ChatParticipantCreator.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatParticipantAdmin):
		m2 := ChatParticipant_ChatParticipantAdmin{}
		m2.ChatParticipantAdmin.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChatParticipants) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChatParticipants_ChatParticipantsForbidden:
		m2, _ := m.Payload.(*ChatParticipants_ChatParticipantsForbidden)
		b = m2.ChatParticipantsForbidden.Encode()
	case *ChatParticipants_ChatParticipants:
		m2, _ := m.Payload.(*ChatParticipants_ChatParticipants)
		b = m2.ChatParticipants.Encode()
	}
	return
}

func (m *ChatParticipants) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatParticipantsForbidden):
		m2 := ChatParticipants_ChatParticipantsForbidden{}
		m2.ChatParticipantsForbidden.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatParticipants):
		m2 := ChatParticipants_ChatParticipants{}
		m2.ChatParticipants.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChatPhoto) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChatPhoto_ChatPhotoEmpty:
		m2, _ := m.Payload.(*ChatPhoto_ChatPhotoEmpty)
		b = m2.ChatPhotoEmpty.Encode()
	case *ChatPhoto_ChatPhoto:
		m2, _ := m.Payload.(*ChatPhoto_ChatPhoto)
		b = m2.ChatPhoto.Encode()
	}
	return
}

func (m *ChatPhoto) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatPhotoEmpty):
		m2 := ChatPhoto_ChatPhotoEmpty{}
		m2.ChatPhotoEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatPhoto):
		m2 := ChatPhoto_ChatPhoto{}
		m2.ChatPhoto.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Message) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Message_MessageEmpty:
		m2, _ := m.Payload.(*Message_MessageEmpty)
		b = m2.MessageEmpty.Encode()
	case *Message_Message:
		m2, _ := m.Payload.(*Message_Message)
		b = m2.Message.Encode()
	case *Message_MessageService:
		m2, _ := m.Payload.(*Message_MessageService)
		b = m2.MessageService.Encode()
	}
	return
}

func (m *Message) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messageEmpty):
		m2 := Message_MessageEmpty{}
		m2.MessageEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_message):
		m2 := Message_Message{}
		m2.Message.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageService):
		m2 := Message_MessageService{}
		m2.MessageService.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MessageMedia) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MessageMedia_MessageMediaEmpty:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaEmpty)
		b = m2.MessageMediaEmpty.Encode()
	case *MessageMedia_MessageMediaPhoto:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaPhoto)
		b = m2.MessageMediaPhoto.Encode()
	case *MessageMedia_MessageMediaGeo:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaGeo)
		b = m2.MessageMediaGeo.Encode()
	case *MessageMedia_MessageMediaContact:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaContact)
		b = m2.MessageMediaContact.Encode()
	case *MessageMedia_MessageMediaUnsupported:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaUnsupported)
		b = m2.MessageMediaUnsupported.Encode()
	case *MessageMedia_MessageMediaDocument:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaDocument)
		b = m2.MessageMediaDocument.Encode()
	case *MessageMedia_MessageMediaWebPage:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaWebPage)
		b = m2.MessageMediaWebPage.Encode()
	case *MessageMedia_MessageMediaVenue:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaVenue)
		b = m2.MessageMediaVenue.Encode()
	case *MessageMedia_MessageMediaGame:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaGame)
		b = m2.MessageMediaGame.Encode()
	case *MessageMedia_MessageMediaInvoice:
		m2, _ := m.Payload.(*MessageMedia_MessageMediaInvoice)
		b = m2.MessageMediaInvoice.Encode()
	}
	return
}

func (m *MessageMedia) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messageMediaEmpty):
		m2 := MessageMedia_MessageMediaEmpty{}
		m2.MessageMediaEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaPhoto):
		m2 := MessageMedia_MessageMediaPhoto{}
		m2.MessageMediaPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaGeo):
		m2 := MessageMedia_MessageMediaGeo{}
		m2.MessageMediaGeo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaContact):
		m2 := MessageMedia_MessageMediaContact{}
		m2.MessageMediaContact.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaUnsupported):
		m2 := MessageMedia_MessageMediaUnsupported{}
		m2.MessageMediaUnsupported.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaDocument):
		m2 := MessageMedia_MessageMediaDocument{}
		m2.MessageMediaDocument.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaWebPage):
		m2 := MessageMedia_MessageMediaWebPage{}
		m2.MessageMediaWebPage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaVenue):
		m2 := MessageMedia_MessageMediaVenue{}
		m2.MessageMediaVenue.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaGame):
		m2 := MessageMedia_MessageMediaGame{}
		m2.MessageMediaGame.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageMediaInvoice):
		m2 := MessageMedia_MessageMediaInvoice{}
		m2.MessageMediaInvoice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MessageAction) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MessageAction_MessageActionEmpty:
		m2, _ := m.Payload.(*MessageAction_MessageActionEmpty)
		b = m2.MessageActionEmpty.Encode()
	case *MessageAction_MessageActionChatCreate:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatCreate)
		b = m2.MessageActionChatCreate.Encode()
	case *MessageAction_MessageActionChatEditTitle:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatEditTitle)
		b = m2.MessageActionChatEditTitle.Encode()
	case *MessageAction_MessageActionChatEditPhoto:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatEditPhoto)
		b = m2.MessageActionChatEditPhoto.Encode()
	case *MessageAction_MessageActionChatDeletePhoto:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatDeletePhoto)
		b = m2.MessageActionChatDeletePhoto.Encode()
	case *MessageAction_MessageActionChatAddUser:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatAddUser)
		b = m2.MessageActionChatAddUser.Encode()
	case *MessageAction_MessageActionChatDeleteUser:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatDeleteUser)
		b = m2.MessageActionChatDeleteUser.Encode()
	case *MessageAction_MessageActionChatJoinedByLink:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatJoinedByLink)
		b = m2.MessageActionChatJoinedByLink.Encode()
	case *MessageAction_MessageActionChannelCreate:
		m2, _ := m.Payload.(*MessageAction_MessageActionChannelCreate)
		b = m2.MessageActionChannelCreate.Encode()
	case *MessageAction_MessageActionChatMigrateTo:
		m2, _ := m.Payload.(*MessageAction_MessageActionChatMigrateTo)
		b = m2.MessageActionChatMigrateTo.Encode()
	case *MessageAction_MessageActionChannelMigrateFrom:
		m2, _ := m.Payload.(*MessageAction_MessageActionChannelMigrateFrom)
		b = m2.MessageActionChannelMigrateFrom.Encode()
	case *MessageAction_MessageActionPinMessage:
		m2, _ := m.Payload.(*MessageAction_MessageActionPinMessage)
		b = m2.MessageActionPinMessage.Encode()
	case *MessageAction_MessageActionHistoryClear:
		m2, _ := m.Payload.(*MessageAction_MessageActionHistoryClear)
		b = m2.MessageActionHistoryClear.Encode()
	case *MessageAction_MessageActionGameScore:
		m2, _ := m.Payload.(*MessageAction_MessageActionGameScore)
		b = m2.MessageActionGameScore.Encode()
	case *MessageAction_MessageActionPaymentSentMe:
		m2, _ := m.Payload.(*MessageAction_MessageActionPaymentSentMe)
		b = m2.MessageActionPaymentSentMe.Encode()
	case *MessageAction_MessageActionPaymentSent:
		m2, _ := m.Payload.(*MessageAction_MessageActionPaymentSent)
		b = m2.MessageActionPaymentSent.Encode()
	case *MessageAction_MessageActionPhoneCall:
		m2, _ := m.Payload.(*MessageAction_MessageActionPhoneCall)
		b = m2.MessageActionPhoneCall.Encode()
	case *MessageAction_MessageActionScreenshotTaken:
		m2, _ := m.Payload.(*MessageAction_MessageActionScreenshotTaken)
		b = m2.MessageActionScreenshotTaken.Encode()
	}
	return
}

func (m *MessageAction) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messageActionEmpty):
		m2 := MessageAction_MessageActionEmpty{}
		m2.MessageActionEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatCreate):
		m2 := MessageAction_MessageActionChatCreate{}
		m2.MessageActionChatCreate.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatEditTitle):
		m2 := MessageAction_MessageActionChatEditTitle{}
		m2.MessageActionChatEditTitle.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatEditPhoto):
		m2 := MessageAction_MessageActionChatEditPhoto{}
		m2.MessageActionChatEditPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatDeletePhoto):
		m2 := MessageAction_MessageActionChatDeletePhoto{}
		m2.MessageActionChatDeletePhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatAddUser):
		m2 := MessageAction_MessageActionChatAddUser{}
		m2.MessageActionChatAddUser.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatDeleteUser):
		m2 := MessageAction_MessageActionChatDeleteUser{}
		m2.MessageActionChatDeleteUser.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatJoinedByLink):
		m2 := MessageAction_MessageActionChatJoinedByLink{}
		m2.MessageActionChatJoinedByLink.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChannelCreate):
		m2 := MessageAction_MessageActionChannelCreate{}
		m2.MessageActionChannelCreate.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChatMigrateTo):
		m2 := MessageAction_MessageActionChatMigrateTo{}
		m2.MessageActionChatMigrateTo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionChannelMigrateFrom):
		m2 := MessageAction_MessageActionChannelMigrateFrom{}
		m2.MessageActionChannelMigrateFrom.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionPinMessage):
		m2 := MessageAction_MessageActionPinMessage{}
		m2.MessageActionPinMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionHistoryClear):
		m2 := MessageAction_MessageActionHistoryClear{}
		m2.MessageActionHistoryClear.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionGameScore):
		m2 := MessageAction_MessageActionGameScore{}
		m2.MessageActionGameScore.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionPaymentSentMe):
		m2 := MessageAction_MessageActionPaymentSentMe{}
		m2.MessageActionPaymentSentMe.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionPaymentSent):
		m2 := MessageAction_MessageActionPaymentSent{}
		m2.MessageActionPaymentSent.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionPhoneCall):
		m2 := MessageAction_MessageActionPhoneCall{}
		m2.MessageActionPhoneCall.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageActionScreenshotTaken):
		m2 := MessageAction_MessageActionScreenshotTaken{}
		m2.MessageActionScreenshotTaken.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Dialog) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Dialog_Dialog:
		m2, _ := m.Payload.(*Dialog_Dialog)
		b = m2.Dialog.Encode()
	}
	return
}

func (m *Dialog) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_dialog):
		m2 := Dialog_Dialog{}
		m2.Dialog.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Photo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Photo_PhotoEmpty:
		m2, _ := m.Payload.(*Photo_PhotoEmpty)
		b = m2.PhotoEmpty.Encode()
	case *Photo_Photo:
		m2, _ := m.Payload.(*Photo_Photo)
		b = m2.Photo.Encode()
	}
	return
}

func (m *Photo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_photoEmpty):
		m2 := Photo_PhotoEmpty{}
		m2.PhotoEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_photo):
		m2 := Photo_Photo{}
		m2.Photo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PhotoSize) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PhotoSize_PhotoSizeEmpty:
		m2, _ := m.Payload.(*PhotoSize_PhotoSizeEmpty)
		b = m2.PhotoSizeEmpty.Encode()
	case *PhotoSize_PhotoSize:
		m2, _ := m.Payload.(*PhotoSize_PhotoSize)
		b = m2.PhotoSize.Encode()
	case *PhotoSize_PhotoCachedSize:
		m2, _ := m.Payload.(*PhotoSize_PhotoCachedSize)
		b = m2.PhotoCachedSize.Encode()
	}
	return
}

func (m *PhotoSize) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_photoSizeEmpty):
		m2 := PhotoSize_PhotoSizeEmpty{}
		m2.PhotoSizeEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_photoSize):
		m2 := PhotoSize_PhotoSize{}
		m2.PhotoSize.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_photoCachedSize):
		m2 := PhotoSize_PhotoCachedSize{}
		m2.PhotoCachedSize.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *GeoPoint) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *GeoPoint_GeoPointEmpty:
		m2, _ := m.Payload.(*GeoPoint_GeoPointEmpty)
		b = m2.GeoPointEmpty.Encode()
	case *GeoPoint_GeoPoint:
		m2, _ := m.Payload.(*GeoPoint_GeoPoint)
		b = m2.GeoPoint.Encode()
	}
	return
}

func (m *GeoPoint) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_geoPointEmpty):
		m2 := GeoPoint_GeoPointEmpty{}
		m2.GeoPointEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_geoPoint):
		m2 := GeoPoint_GeoPoint{}
		m2.GeoPoint.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_CheckedPhone) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_CheckedPhone_AuthCheckedPhone:
		m2, _ := m.Payload.(*Auth_CheckedPhone_AuthCheckedPhone)
		b = m2.AuthCheckedPhone.Encode()
	}
	return
}

func (m *Auth_CheckedPhone) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_checkedPhone):
		m2 := Auth_CheckedPhone_AuthCheckedPhone{}
		m2.AuthCheckedPhone.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_SentCode) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_SentCode_AuthSentCode:
		m2, _ := m.Payload.(*Auth_SentCode_AuthSentCode)
		b = m2.AuthSentCode.Encode()
	}
	return
}

func (m *Auth_SentCode) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_sentCode):
		m2 := Auth_SentCode_AuthSentCode{}
		m2.AuthSentCode.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_Authorization) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_Authorization_AuthAuthorization:
		m2, _ := m.Payload.(*Auth_Authorization_AuthAuthorization)
		b = m2.AuthAuthorization.Encode()
	}
	return
}

func (m *Auth_Authorization) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_authorization):
		m2 := Auth_Authorization_AuthAuthorization{}
		m2.AuthAuthorization.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_ExportedAuthorization) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_ExportedAuthorization_AuthExportedAuthorization:
		m2, _ := m.Payload.(*Auth_ExportedAuthorization_AuthExportedAuthorization)
		b = m2.AuthExportedAuthorization.Encode()
	}
	return
}

func (m *Auth_ExportedAuthorization) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_exportedAuthorization):
		m2 := Auth_ExportedAuthorization_AuthExportedAuthorization{}
		m2.AuthExportedAuthorization.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputNotifyPeer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputNotifyPeer_InputNotifyPeer:
		m2, _ := m.Payload.(*InputNotifyPeer_InputNotifyPeer)
		b = m2.InputNotifyPeer.Encode()
	case *InputNotifyPeer_InputNotifyUsers:
		m2, _ := m.Payload.(*InputNotifyPeer_InputNotifyUsers)
		b = m2.InputNotifyUsers.Encode()
	case *InputNotifyPeer_InputNotifyChats:
		m2, _ := m.Payload.(*InputNotifyPeer_InputNotifyChats)
		b = m2.InputNotifyChats.Encode()
	case *InputNotifyPeer_InputNotifyAll:
		m2, _ := m.Payload.(*InputNotifyPeer_InputNotifyAll)
		b = m2.InputNotifyAll.Encode()
	}
	return
}

func (m *InputNotifyPeer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputNotifyPeer):
		m2 := InputNotifyPeer_InputNotifyPeer{}
		m2.InputNotifyPeer.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputNotifyUsers):
		m2 := InputNotifyPeer_InputNotifyUsers{}
		m2.InputNotifyUsers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputNotifyChats):
		m2 := InputNotifyPeer_InputNotifyChats{}
		m2.InputNotifyChats.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputNotifyAll):
		m2 := InputNotifyPeer_InputNotifyAll{}
		m2.InputNotifyAll.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPeerNotifyEvents) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPeerNotifyEvents_InputPeerNotifyEventsEmpty:
		m2, _ := m.Payload.(*InputPeerNotifyEvents_InputPeerNotifyEventsEmpty)
		b = m2.InputPeerNotifyEventsEmpty.Encode()
	case *InputPeerNotifyEvents_InputPeerNotifyEventsAll:
		m2, _ := m.Payload.(*InputPeerNotifyEvents_InputPeerNotifyEventsAll)
		b = m2.InputPeerNotifyEventsAll.Encode()
	}
	return
}

func (m *InputPeerNotifyEvents) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPeerNotifyEventsEmpty):
		m2 := InputPeerNotifyEvents_InputPeerNotifyEventsEmpty{}
		m2.InputPeerNotifyEventsEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPeerNotifyEventsAll):
		m2 := InputPeerNotifyEvents_InputPeerNotifyEventsAll{}
		m2.InputPeerNotifyEventsAll.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPeerNotifySettings) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPeerNotifySettings_InputPeerNotifySettings:
		m2, _ := m.Payload.(*InputPeerNotifySettings_InputPeerNotifySettings)
		b = m2.InputPeerNotifySettings.Encode()
	}
	return
}

func (m *InputPeerNotifySettings) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPeerNotifySettings):
		m2 := InputPeerNotifySettings_InputPeerNotifySettings{}
		m2.InputPeerNotifySettings.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PeerNotifyEvents) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PeerNotifyEvents_PeerNotifyEventsEmpty:
		m2, _ := m.Payload.(*PeerNotifyEvents_PeerNotifyEventsEmpty)
		b = m2.PeerNotifyEventsEmpty.Encode()
	case *PeerNotifyEvents_PeerNotifyEventsAll:
		m2, _ := m.Payload.(*PeerNotifyEvents_PeerNotifyEventsAll)
		b = m2.PeerNotifyEventsAll.Encode()
	}
	return
}

func (m *PeerNotifyEvents) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_peerNotifyEventsEmpty):
		m2 := PeerNotifyEvents_PeerNotifyEventsEmpty{}
		m2.PeerNotifyEventsEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_peerNotifyEventsAll):
		m2 := PeerNotifyEvents_PeerNotifyEventsAll{}
		m2.PeerNotifyEventsAll.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PeerNotifySettings) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PeerNotifySettings_PeerNotifySettingsEmpty:
		m2, _ := m.Payload.(*PeerNotifySettings_PeerNotifySettingsEmpty)
		b = m2.PeerNotifySettingsEmpty.Encode()
	case *PeerNotifySettings_PeerNotifySettings:
		m2, _ := m.Payload.(*PeerNotifySettings_PeerNotifySettings)
		b = m2.PeerNotifySettings.Encode()
	}
	return
}

func (m *PeerNotifySettings) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_peerNotifySettingsEmpty):
		m2 := PeerNotifySettings_PeerNotifySettingsEmpty{}
		m2.PeerNotifySettingsEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_peerNotifySettings):
		m2 := PeerNotifySettings_PeerNotifySettings{}
		m2.PeerNotifySettings.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PeerSettings) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PeerSettings_PeerSettings:
		m2, _ := m.Payload.(*PeerSettings_PeerSettings)
		b = m2.PeerSettings.Encode()
	}
	return
}

func (m *PeerSettings) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_peerSettings):
		m2 := PeerSettings_PeerSettings{}
		m2.PeerSettings.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *WallPaper) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *WallPaper_WallPaper:
		m2, _ := m.Payload.(*WallPaper_WallPaper)
		b = m2.WallPaper.Encode()
	case *WallPaper_WallPaperSolid:
		m2, _ := m.Payload.(*WallPaper_WallPaperSolid)
		b = m2.WallPaperSolid.Encode()
	}
	return
}

func (m *WallPaper) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_wallPaper):
		m2 := WallPaper_WallPaper{}
		m2.WallPaper.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_wallPaperSolid):
		m2 := WallPaper_WallPaperSolid{}
		m2.WallPaperSolid.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ReportReason) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ReportReason_InputReportReasonSpam:
		m2, _ := m.Payload.(*ReportReason_InputReportReasonSpam)
		b = m2.InputReportReasonSpam.Encode()
	case *ReportReason_InputReportReasonViolence:
		m2, _ := m.Payload.(*ReportReason_InputReportReasonViolence)
		b = m2.InputReportReasonViolence.Encode()
	case *ReportReason_InputReportReasonPornography:
		m2, _ := m.Payload.(*ReportReason_InputReportReasonPornography)
		b = m2.InputReportReasonPornography.Encode()
	case *ReportReason_InputReportReasonOther:
		m2, _ := m.Payload.(*ReportReason_InputReportReasonOther)
		b = m2.InputReportReasonOther.Encode()
	}
	return
}

func (m *ReportReason) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputReportReasonSpam):
		m2 := ReportReason_InputReportReasonSpam{}
		m2.InputReportReasonSpam.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputReportReasonViolence):
		m2 := ReportReason_InputReportReasonViolence{}
		m2.InputReportReasonViolence.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputReportReasonPornography):
		m2 := ReportReason_InputReportReasonPornography{}
		m2.InputReportReasonPornography.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputReportReasonOther):
		m2 := ReportReason_InputReportReasonOther{}
		m2.InputReportReasonOther.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *UserFull) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *UserFull_UserFull:
		m2, _ := m.Payload.(*UserFull_UserFull)
		b = m2.UserFull.Encode()
	}
	return
}

func (m *UserFull) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_userFull):
		m2 := UserFull_UserFull{}
		m2.UserFull.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contact) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contact_Contact:
		m2, _ := m.Payload.(*Contact_Contact)
		b = m2.Contact.Encode()
	}
	return
}

func (m *Contact) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contact):
		m2 := Contact_Contact{}
		m2.Contact.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ImportedContact) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ImportedContact_ImportedContact:
		m2, _ := m.Payload.(*ImportedContact_ImportedContact)
		b = m2.ImportedContact.Encode()
	}
	return
}

func (m *ImportedContact) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_importedContact):
		m2 := ImportedContact_ImportedContact{}
		m2.ImportedContact.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ContactBlocked) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ContactBlocked_ContactBlocked:
		m2, _ := m.Payload.(*ContactBlocked_ContactBlocked)
		b = m2.ContactBlocked.Encode()
	}
	return
}

func (m *ContactBlocked) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contactBlocked):
		m2 := ContactBlocked_ContactBlocked{}
		m2.ContactBlocked.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ContactStatus) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ContactStatus_ContactStatus:
		m2, _ := m.Payload.(*ContactStatus_ContactStatus)
		b = m2.ContactStatus.Encode()
	}
	return
}

func (m *ContactStatus) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contactStatus):
		m2 := ContactStatus_ContactStatus{}
		m2.ContactStatus.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_Link) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_Link_ContactsLink:
		m2, _ := m.Payload.(*Contacts_Link_ContactsLink)
		b = m2.ContactsLink.Encode()
	}
	return
}

func (m *Contacts_Link) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_link):
		m2 := Contacts_Link_ContactsLink{}
		m2.ContactsLink.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_Contacts) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_Contacts_ContactsContactsNotModified:
		m2, _ := m.Payload.(*Contacts_Contacts_ContactsContactsNotModified)
		b = m2.ContactsContactsNotModified.Encode()
	case *Contacts_Contacts_ContactsContacts:
		m2, _ := m.Payload.(*Contacts_Contacts_ContactsContacts)
		b = m2.ContactsContacts.Encode()
	}
	return
}

func (m *Contacts_Contacts) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_contactsNotModified):
		m2 := Contacts_Contacts_ContactsContactsNotModified{}
		m2.ContactsContactsNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_contacts_contacts):
		m2 := Contacts_Contacts_ContactsContacts{}
		m2.ContactsContacts.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_ImportedContacts) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_ImportedContacts_ContactsImportedContacts:
		m2, _ := m.Payload.(*Contacts_ImportedContacts_ContactsImportedContacts)
		b = m2.ContactsImportedContacts.Encode()
	}
	return
}

func (m *Contacts_ImportedContacts) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_importedContacts):
		m2 := Contacts_ImportedContacts_ContactsImportedContacts{}
		m2.ContactsImportedContacts.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_Blocked) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_Blocked_ContactsBlocked:
		m2, _ := m.Payload.(*Contacts_Blocked_ContactsBlocked)
		b = m2.ContactsBlocked.Encode()
	case *Contacts_Blocked_ContactsBlockedSlice:
		m2, _ := m.Payload.(*Contacts_Blocked_ContactsBlockedSlice)
		b = m2.ContactsBlockedSlice.Encode()
	}
	return
}

func (m *Contacts_Blocked) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_blocked):
		m2 := Contacts_Blocked_ContactsBlocked{}
		m2.ContactsBlocked.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_contacts_blockedSlice):
		m2 := Contacts_Blocked_ContactsBlockedSlice{}
		m2.ContactsBlockedSlice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_Dialogs) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_Dialogs_MessagesDialogs:
		m2, _ := m.Payload.(*Messages_Dialogs_MessagesDialogs)
		b = m2.MessagesDialogs.Encode()
	case *Messages_Dialogs_MessagesDialogsSlice:
		m2, _ := m.Payload.(*Messages_Dialogs_MessagesDialogsSlice)
		b = m2.MessagesDialogsSlice.Encode()
	}
	return
}

func (m *Messages_Dialogs) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_dialogs):
		m2 := Messages_Dialogs_MessagesDialogs{}
		m2.MessagesDialogs.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_dialogsSlice):
		m2 := Messages_Dialogs_MessagesDialogsSlice{}
		m2.MessagesDialogsSlice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_Messages) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_Messages_MessagesMessages:
		m2, _ := m.Payload.(*Messages_Messages_MessagesMessages)
		b = m2.MessagesMessages.Encode()
	case *Messages_Messages_MessagesMessagesSlice:
		m2, _ := m.Payload.(*Messages_Messages_MessagesMessagesSlice)
		b = m2.MessagesMessagesSlice.Encode()
	case *Messages_Messages_MessagesChannelMessages:
		m2, _ := m.Payload.(*Messages_Messages_MessagesChannelMessages)
		b = m2.MessagesChannelMessages.Encode()
	}
	return
}

func (m *Messages_Messages) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_messages):
		m2 := Messages_Messages_MessagesMessages{}
		m2.MessagesMessages.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_messagesSlice):
		m2 := Messages_Messages_MessagesMessagesSlice{}
		m2.MessagesMessagesSlice.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_channelMessages):
		m2 := Messages_Messages_MessagesChannelMessages{}
		m2.MessagesChannelMessages.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_Chats) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_Chats_MessagesChats:
		m2, _ := m.Payload.(*Messages_Chats_MessagesChats)
		b = m2.MessagesChats.Encode()
	case *Messages_Chats_MessagesChatsSlice:
		m2, _ := m.Payload.(*Messages_Chats_MessagesChatsSlice)
		b = m2.MessagesChatsSlice.Encode()
	}
	return
}

func (m *Messages_Chats) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_chats):
		m2 := Messages_Chats_MessagesChats{}
		m2.MessagesChats.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_chatsSlice):
		m2 := Messages_Chats_MessagesChatsSlice{}
		m2.MessagesChatsSlice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_ChatFull) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_ChatFull_MessagesChatFull:
		m2, _ := m.Payload.(*Messages_ChatFull_MessagesChatFull)
		b = m2.MessagesChatFull.Encode()
	}
	return
}

func (m *Messages_ChatFull) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_chatFull):
		m2 := Messages_ChatFull_MessagesChatFull{}
		m2.MessagesChatFull.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_AffectedHistory) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_AffectedHistory_MessagesAffectedHistory:
		m2, _ := m.Payload.(*Messages_AffectedHistory_MessagesAffectedHistory)
		b = m2.MessagesAffectedHistory.Encode()
	}
	return
}

func (m *Messages_AffectedHistory) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_affectedHistory):
		m2 := Messages_AffectedHistory_MessagesAffectedHistory{}
		m2.MessagesAffectedHistory.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MessagesFilter) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MessagesFilter_InputMessagesFilterEmpty:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterEmpty)
		b = m2.InputMessagesFilterEmpty.Encode()
	case *MessagesFilter_InputMessagesFilterPhotos:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterPhotos)
		b = m2.InputMessagesFilterPhotos.Encode()
	case *MessagesFilter_InputMessagesFilterVideo:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterVideo)
		b = m2.InputMessagesFilterVideo.Encode()
	case *MessagesFilter_InputMessagesFilterPhotoVideo:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterPhotoVideo)
		b = m2.InputMessagesFilterPhotoVideo.Encode()
	case *MessagesFilter_InputMessagesFilterPhotoVideoDocuments:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterPhotoVideoDocuments)
		b = m2.InputMessagesFilterPhotoVideoDocuments.Encode()
	case *MessagesFilter_InputMessagesFilterDocument:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterDocument)
		b = m2.InputMessagesFilterDocument.Encode()
	case *MessagesFilter_InputMessagesFilterUrl:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterUrl)
		b = m2.InputMessagesFilterUrl.Encode()
	case *MessagesFilter_InputMessagesFilterGif:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterGif)
		b = m2.InputMessagesFilterGif.Encode()
	case *MessagesFilter_InputMessagesFilterVoice:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterVoice)
		b = m2.InputMessagesFilterVoice.Encode()
	case *MessagesFilter_InputMessagesFilterMusic:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterMusic)
		b = m2.InputMessagesFilterMusic.Encode()
	case *MessagesFilter_InputMessagesFilterChatPhotos:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterChatPhotos)
		b = m2.InputMessagesFilterChatPhotos.Encode()
	case *MessagesFilter_InputMessagesFilterPhoneCalls:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterPhoneCalls)
		b = m2.InputMessagesFilterPhoneCalls.Encode()
	case *MessagesFilter_InputMessagesFilterRoundVoice:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterRoundVoice)
		b = m2.InputMessagesFilterRoundVoice.Encode()
	case *MessagesFilter_InputMessagesFilterRoundVideo:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterRoundVideo)
		b = m2.InputMessagesFilterRoundVideo.Encode()
	case *MessagesFilter_InputMessagesFilterMyMentions:
		m2, _ := m.Payload.(*MessagesFilter_InputMessagesFilterMyMentions)
		b = m2.InputMessagesFilterMyMentions.Encode()
	}
	return
}

func (m *MessagesFilter) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputMessagesFilterEmpty):
		m2 := MessagesFilter_InputMessagesFilterEmpty{}
		m2.InputMessagesFilterEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterPhotos):
		m2 := MessagesFilter_InputMessagesFilterPhotos{}
		m2.InputMessagesFilterPhotos.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterVideo):
		m2 := MessagesFilter_InputMessagesFilterVideo{}
		m2.InputMessagesFilterVideo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideo):
		m2 := MessagesFilter_InputMessagesFilterPhotoVideo{}
		m2.InputMessagesFilterPhotoVideo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideoDocuments):
		m2 := MessagesFilter_InputMessagesFilterPhotoVideoDocuments{}
		m2.InputMessagesFilterPhotoVideoDocuments.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterDocument):
		m2 := MessagesFilter_InputMessagesFilterDocument{}
		m2.InputMessagesFilterDocument.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterUrl):
		m2 := MessagesFilter_InputMessagesFilterUrl{}
		m2.InputMessagesFilterUrl.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterGif):
		m2 := MessagesFilter_InputMessagesFilterGif{}
		m2.InputMessagesFilterGif.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterVoice):
		m2 := MessagesFilter_InputMessagesFilterVoice{}
		m2.InputMessagesFilterVoice.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterMusic):
		m2 := MessagesFilter_InputMessagesFilterMusic{}
		m2.InputMessagesFilterMusic.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterChatPhotos):
		m2 := MessagesFilter_InputMessagesFilterChatPhotos{}
		m2.InputMessagesFilterChatPhotos.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterPhoneCalls):
		m2 := MessagesFilter_InputMessagesFilterPhoneCalls{}
		m2.InputMessagesFilterPhoneCalls.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterRoundVoice):
		m2 := MessagesFilter_InputMessagesFilterRoundVoice{}
		m2.InputMessagesFilterRoundVoice.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterRoundVideo):
		m2 := MessagesFilter_InputMessagesFilterRoundVideo{}
		m2.InputMessagesFilterRoundVideo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessagesFilterMyMentions):
		m2 := MessagesFilter_InputMessagesFilterMyMentions{}
		m2.InputMessagesFilterMyMentions.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Update) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Update_UpdateNewMessage:
		m2, _ := m.Payload.(*Update_UpdateNewMessage)
		b = m2.UpdateNewMessage.Encode()
	case *Update_UpdateMessageID:
		m2, _ := m.Payload.(*Update_UpdateMessageID)
		b = m2.UpdateMessageID.Encode()
	case *Update_UpdateDeleteMessages:
		m2, _ := m.Payload.(*Update_UpdateDeleteMessages)
		b = m2.UpdateDeleteMessages.Encode()
	case *Update_UpdateUserTyping:
		m2, _ := m.Payload.(*Update_UpdateUserTyping)
		b = m2.UpdateUserTyping.Encode()
	case *Update_UpdateChatUserTyping:
		m2, _ := m.Payload.(*Update_UpdateChatUserTyping)
		b = m2.UpdateChatUserTyping.Encode()
	case *Update_UpdateChatParticipants:
		m2, _ := m.Payload.(*Update_UpdateChatParticipants)
		b = m2.UpdateChatParticipants.Encode()
	case *Update_UpdateUserStatus:
		m2, _ := m.Payload.(*Update_UpdateUserStatus)
		b = m2.UpdateUserStatus.Encode()
	case *Update_UpdateUserName:
		m2, _ := m.Payload.(*Update_UpdateUserName)
		b = m2.UpdateUserName.Encode()
	case *Update_UpdateUserPhoto:
		m2, _ := m.Payload.(*Update_UpdateUserPhoto)
		b = m2.UpdateUserPhoto.Encode()
	case *Update_UpdateContactRegistered:
		m2, _ := m.Payload.(*Update_UpdateContactRegistered)
		b = m2.UpdateContactRegistered.Encode()
	case *Update_UpdateContactLink:
		m2, _ := m.Payload.(*Update_UpdateContactLink)
		b = m2.UpdateContactLink.Encode()
	case *Update_UpdateNewEncryptedMessage:
		m2, _ := m.Payload.(*Update_UpdateNewEncryptedMessage)
		b = m2.UpdateNewEncryptedMessage.Encode()
	case *Update_UpdateEncryptedChatTyping:
		m2, _ := m.Payload.(*Update_UpdateEncryptedChatTyping)
		b = m2.UpdateEncryptedChatTyping.Encode()
	case *Update_UpdateEncryption:
		m2, _ := m.Payload.(*Update_UpdateEncryption)
		b = m2.UpdateEncryption.Encode()
	case *Update_UpdateEncryptedMessagesRead:
		m2, _ := m.Payload.(*Update_UpdateEncryptedMessagesRead)
		b = m2.UpdateEncryptedMessagesRead.Encode()
	case *Update_UpdateChatParticipantAdd:
		m2, _ := m.Payload.(*Update_UpdateChatParticipantAdd)
		b = m2.UpdateChatParticipantAdd.Encode()
	case *Update_UpdateChatParticipantDelete:
		m2, _ := m.Payload.(*Update_UpdateChatParticipantDelete)
		b = m2.UpdateChatParticipantDelete.Encode()
	case *Update_UpdateDcOptions:
		m2, _ := m.Payload.(*Update_UpdateDcOptions)
		b = m2.UpdateDcOptions.Encode()
	case *Update_UpdateUserBlocked:
		m2, _ := m.Payload.(*Update_UpdateUserBlocked)
		b = m2.UpdateUserBlocked.Encode()
	case *Update_UpdateNotifySettings:
		m2, _ := m.Payload.(*Update_UpdateNotifySettings)
		b = m2.UpdateNotifySettings.Encode()
	case *Update_UpdateServiceNotification:
		m2, _ := m.Payload.(*Update_UpdateServiceNotification)
		b = m2.UpdateServiceNotification.Encode()
	case *Update_UpdatePrivacy:
		m2, _ := m.Payload.(*Update_UpdatePrivacy)
		b = m2.UpdatePrivacy.Encode()
	case *Update_UpdateUserPhone:
		m2, _ := m.Payload.(*Update_UpdateUserPhone)
		b = m2.UpdateUserPhone.Encode()
	case *Update_UpdateReadHistoryInbox:
		m2, _ := m.Payload.(*Update_UpdateReadHistoryInbox)
		b = m2.UpdateReadHistoryInbox.Encode()
	case *Update_UpdateReadHistoryOutbox:
		m2, _ := m.Payload.(*Update_UpdateReadHistoryOutbox)
		b = m2.UpdateReadHistoryOutbox.Encode()
	case *Update_UpdateWebPage:
		m2, _ := m.Payload.(*Update_UpdateWebPage)
		b = m2.UpdateWebPage.Encode()
	case *Update_UpdateReadMessagesContents:
		m2, _ := m.Payload.(*Update_UpdateReadMessagesContents)
		b = m2.UpdateReadMessagesContents.Encode()
	case *Update_UpdateChannelTooLong:
		m2, _ := m.Payload.(*Update_UpdateChannelTooLong)
		b = m2.UpdateChannelTooLong.Encode()
	case *Update_UpdateChannel:
		m2, _ := m.Payload.(*Update_UpdateChannel)
		b = m2.UpdateChannel.Encode()
	case *Update_UpdateNewChannelMessage:
		m2, _ := m.Payload.(*Update_UpdateNewChannelMessage)
		b = m2.UpdateNewChannelMessage.Encode()
	case *Update_UpdateReadChannelInbox:
		m2, _ := m.Payload.(*Update_UpdateReadChannelInbox)
		b = m2.UpdateReadChannelInbox.Encode()
	case *Update_UpdateDeleteChannelMessages:
		m2, _ := m.Payload.(*Update_UpdateDeleteChannelMessages)
		b = m2.UpdateDeleteChannelMessages.Encode()
	case *Update_UpdateChannelMessageViews:
		m2, _ := m.Payload.(*Update_UpdateChannelMessageViews)
		b = m2.UpdateChannelMessageViews.Encode()
	case *Update_UpdateChatAdmins:
		m2, _ := m.Payload.(*Update_UpdateChatAdmins)
		b = m2.UpdateChatAdmins.Encode()
	case *Update_UpdateChatParticipantAdmin:
		m2, _ := m.Payload.(*Update_UpdateChatParticipantAdmin)
		b = m2.UpdateChatParticipantAdmin.Encode()
	case *Update_UpdateNewStickerSet:
		m2, _ := m.Payload.(*Update_UpdateNewStickerSet)
		b = m2.UpdateNewStickerSet.Encode()
	case *Update_UpdateStickerSetsOrder:
		m2, _ := m.Payload.(*Update_UpdateStickerSetsOrder)
		b = m2.UpdateStickerSetsOrder.Encode()
	case *Update_UpdateStickerSets:
		m2, _ := m.Payload.(*Update_UpdateStickerSets)
		b = m2.UpdateStickerSets.Encode()
	case *Update_UpdateSavedGifs:
		m2, _ := m.Payload.(*Update_UpdateSavedGifs)
		b = m2.UpdateSavedGifs.Encode()
	case *Update_UpdateBotInlineQuery:
		m2, _ := m.Payload.(*Update_UpdateBotInlineQuery)
		b = m2.UpdateBotInlineQuery.Encode()
	case *Update_UpdateBotInlineSend:
		m2, _ := m.Payload.(*Update_UpdateBotInlineSend)
		b = m2.UpdateBotInlineSend.Encode()
	case *Update_UpdateEditChannelMessage:
		m2, _ := m.Payload.(*Update_UpdateEditChannelMessage)
		b = m2.UpdateEditChannelMessage.Encode()
	case *Update_UpdateChannelPinnedMessage:
		m2, _ := m.Payload.(*Update_UpdateChannelPinnedMessage)
		b = m2.UpdateChannelPinnedMessage.Encode()
	case *Update_UpdateBotCallbackQuery:
		m2, _ := m.Payload.(*Update_UpdateBotCallbackQuery)
		b = m2.UpdateBotCallbackQuery.Encode()
	case *Update_UpdateEditMessage:
		m2, _ := m.Payload.(*Update_UpdateEditMessage)
		b = m2.UpdateEditMessage.Encode()
	case *Update_UpdateInlineBotCallbackQuery:
		m2, _ := m.Payload.(*Update_UpdateInlineBotCallbackQuery)
		b = m2.UpdateInlineBotCallbackQuery.Encode()
	case *Update_UpdateReadChannelOutbox:
		m2, _ := m.Payload.(*Update_UpdateReadChannelOutbox)
		b = m2.UpdateReadChannelOutbox.Encode()
	case *Update_UpdateDraftMessage:
		m2, _ := m.Payload.(*Update_UpdateDraftMessage)
		b = m2.UpdateDraftMessage.Encode()
	case *Update_UpdateReadFeaturedStickers:
		m2, _ := m.Payload.(*Update_UpdateReadFeaturedStickers)
		b = m2.UpdateReadFeaturedStickers.Encode()
	case *Update_UpdateRecentStickers:
		m2, _ := m.Payload.(*Update_UpdateRecentStickers)
		b = m2.UpdateRecentStickers.Encode()
	case *Update_UpdateConfig:
		m2, _ := m.Payload.(*Update_UpdateConfig)
		b = m2.UpdateConfig.Encode()
	case *Update_UpdatePtsChanged:
		m2, _ := m.Payload.(*Update_UpdatePtsChanged)
		b = m2.UpdatePtsChanged.Encode()
	case *Update_UpdateChannelWebPage:
		m2, _ := m.Payload.(*Update_UpdateChannelWebPage)
		b = m2.UpdateChannelWebPage.Encode()
	case *Update_UpdateDialogPinned:
		m2, _ := m.Payload.(*Update_UpdateDialogPinned)
		b = m2.UpdateDialogPinned.Encode()
	case *Update_UpdatePinnedDialogs:
		m2, _ := m.Payload.(*Update_UpdatePinnedDialogs)
		b = m2.UpdatePinnedDialogs.Encode()
	case *Update_UpdateBotWebhookJSON:
		m2, _ := m.Payload.(*Update_UpdateBotWebhookJSON)
		b = m2.UpdateBotWebhookJSON.Encode()
	case *Update_UpdateBotWebhookJSONQuery:
		m2, _ := m.Payload.(*Update_UpdateBotWebhookJSONQuery)
		b = m2.UpdateBotWebhookJSONQuery.Encode()
	case *Update_UpdateBotShippingQuery:
		m2, _ := m.Payload.(*Update_UpdateBotShippingQuery)
		b = m2.UpdateBotShippingQuery.Encode()
	case *Update_UpdateBotPrecheckoutQuery:
		m2, _ := m.Payload.(*Update_UpdateBotPrecheckoutQuery)
		b = m2.UpdateBotPrecheckoutQuery.Encode()
	case *Update_UpdatePhoneCall:
		m2, _ := m.Payload.(*Update_UpdatePhoneCall)
		b = m2.UpdatePhoneCall.Encode()
	case *Update_UpdateLangPackTooLong:
		m2, _ := m.Payload.(*Update_UpdateLangPackTooLong)
		b = m2.UpdateLangPackTooLong.Encode()
	case *Update_UpdateLangPack:
		m2, _ := m.Payload.(*Update_UpdateLangPack)
		b = m2.UpdateLangPack.Encode()
	case *Update_UpdateFavedStickers:
		m2, _ := m.Payload.(*Update_UpdateFavedStickers)
		b = m2.UpdateFavedStickers.Encode()
	case *Update_UpdateChannelReadMessagesContents:
		m2, _ := m.Payload.(*Update_UpdateChannelReadMessagesContents)
		b = m2.UpdateChannelReadMessagesContents.Encode()
	case *Update_UpdateContactsReset:
		m2, _ := m.Payload.(*Update_UpdateContactsReset)
		b = m2.UpdateContactsReset.Encode()
	}
	return
}

func (m *Update) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_updateNewMessage):
		m2 := Update_UpdateNewMessage{}
		m2.UpdateNewMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateMessageID):
		m2 := Update_UpdateMessageID{}
		m2.UpdateMessageID.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateDeleteMessages):
		m2 := Update_UpdateDeleteMessages{}
		m2.UpdateDeleteMessages.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateUserTyping):
		m2 := Update_UpdateUserTyping{}
		m2.UpdateUserTyping.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChatUserTyping):
		m2 := Update_UpdateChatUserTyping{}
		m2.UpdateChatUserTyping.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChatParticipants):
		m2 := Update_UpdateChatParticipants{}
		m2.UpdateChatParticipants.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateUserStatus):
		m2 := Update_UpdateUserStatus{}
		m2.UpdateUserStatus.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateUserName):
		m2 := Update_UpdateUserName{}
		m2.UpdateUserName.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateUserPhoto):
		m2 := Update_UpdateUserPhoto{}
		m2.UpdateUserPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateContactRegistered):
		m2 := Update_UpdateContactRegistered{}
		m2.UpdateContactRegistered.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateContactLink):
		m2 := Update_UpdateContactLink{}
		m2.UpdateContactLink.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateNewEncryptedMessage):
		m2 := Update_UpdateNewEncryptedMessage{}
		m2.UpdateNewEncryptedMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateEncryptedChatTyping):
		m2 := Update_UpdateEncryptedChatTyping{}
		m2.UpdateEncryptedChatTyping.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateEncryption):
		m2 := Update_UpdateEncryption{}
		m2.UpdateEncryption.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateEncryptedMessagesRead):
		m2 := Update_UpdateEncryptedMessagesRead{}
		m2.UpdateEncryptedMessagesRead.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChatParticipantAdd):
		m2 := Update_UpdateChatParticipantAdd{}
		m2.UpdateChatParticipantAdd.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChatParticipantDelete):
		m2 := Update_UpdateChatParticipantDelete{}
		m2.UpdateChatParticipantDelete.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateDcOptions):
		m2 := Update_UpdateDcOptions{}
		m2.UpdateDcOptions.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateUserBlocked):
		m2 := Update_UpdateUserBlocked{}
		m2.UpdateUserBlocked.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateNotifySettings):
		m2 := Update_UpdateNotifySettings{}
		m2.UpdateNotifySettings.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateServiceNotification):
		m2 := Update_UpdateServiceNotification{}
		m2.UpdateServiceNotification.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updatePrivacy):
		m2 := Update_UpdatePrivacy{}
		m2.UpdatePrivacy.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateUserPhone):
		m2 := Update_UpdateUserPhone{}
		m2.UpdateUserPhone.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateReadHistoryInbox):
		m2 := Update_UpdateReadHistoryInbox{}
		m2.UpdateReadHistoryInbox.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateReadHistoryOutbox):
		m2 := Update_UpdateReadHistoryOutbox{}
		m2.UpdateReadHistoryOutbox.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateWebPage):
		m2 := Update_UpdateWebPage{}
		m2.UpdateWebPage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateReadMessagesContents):
		m2 := Update_UpdateReadMessagesContents{}
		m2.UpdateReadMessagesContents.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChannelTooLong):
		m2 := Update_UpdateChannelTooLong{}
		m2.UpdateChannelTooLong.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChannel):
		m2 := Update_UpdateChannel{}
		m2.UpdateChannel.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateNewChannelMessage):
		m2 := Update_UpdateNewChannelMessage{}
		m2.UpdateNewChannelMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateReadChannelInbox):
		m2 := Update_UpdateReadChannelInbox{}
		m2.UpdateReadChannelInbox.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateDeleteChannelMessages):
		m2 := Update_UpdateDeleteChannelMessages{}
		m2.UpdateDeleteChannelMessages.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChannelMessageViews):
		m2 := Update_UpdateChannelMessageViews{}
		m2.UpdateChannelMessageViews.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChatAdmins):
		m2 := Update_UpdateChatAdmins{}
		m2.UpdateChatAdmins.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChatParticipantAdmin):
		m2 := Update_UpdateChatParticipantAdmin{}
		m2.UpdateChatParticipantAdmin.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateNewStickerSet):
		m2 := Update_UpdateNewStickerSet{}
		m2.UpdateNewStickerSet.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateStickerSetsOrder):
		m2 := Update_UpdateStickerSetsOrder{}
		m2.UpdateStickerSetsOrder.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateStickerSets):
		m2 := Update_UpdateStickerSets{}
		m2.UpdateStickerSets.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateSavedGifs):
		m2 := Update_UpdateSavedGifs{}
		m2.UpdateSavedGifs.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotInlineQuery):
		m2 := Update_UpdateBotInlineQuery{}
		m2.UpdateBotInlineQuery.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotInlineSend):
		m2 := Update_UpdateBotInlineSend{}
		m2.UpdateBotInlineSend.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateEditChannelMessage):
		m2 := Update_UpdateEditChannelMessage{}
		m2.UpdateEditChannelMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChannelPinnedMessage):
		m2 := Update_UpdateChannelPinnedMessage{}
		m2.UpdateChannelPinnedMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotCallbackQuery):
		m2 := Update_UpdateBotCallbackQuery{}
		m2.UpdateBotCallbackQuery.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateEditMessage):
		m2 := Update_UpdateEditMessage{}
		m2.UpdateEditMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateInlineBotCallbackQuery):
		m2 := Update_UpdateInlineBotCallbackQuery{}
		m2.UpdateInlineBotCallbackQuery.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateReadChannelOutbox):
		m2 := Update_UpdateReadChannelOutbox{}
		m2.UpdateReadChannelOutbox.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateDraftMessage):
		m2 := Update_UpdateDraftMessage{}
		m2.UpdateDraftMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateReadFeaturedStickers):
		m2 := Update_UpdateReadFeaturedStickers{}
		m2.UpdateReadFeaturedStickers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateRecentStickers):
		m2 := Update_UpdateRecentStickers{}
		m2.UpdateRecentStickers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateConfig):
		m2 := Update_UpdateConfig{}
		m2.UpdateConfig.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updatePtsChanged):
		m2 := Update_UpdatePtsChanged{}
		m2.UpdatePtsChanged.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChannelWebPage):
		m2 := Update_UpdateChannelWebPage{}
		m2.UpdateChannelWebPage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateDialogPinned):
		m2 := Update_UpdateDialogPinned{}
		m2.UpdateDialogPinned.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updatePinnedDialogs):
		m2 := Update_UpdatePinnedDialogs{}
		m2.UpdatePinnedDialogs.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotWebhookJSON):
		m2 := Update_UpdateBotWebhookJSON{}
		m2.UpdateBotWebhookJSON.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotWebhookJSONQuery):
		m2 := Update_UpdateBotWebhookJSONQuery{}
		m2.UpdateBotWebhookJSONQuery.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotShippingQuery):
		m2 := Update_UpdateBotShippingQuery{}
		m2.UpdateBotShippingQuery.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateBotPrecheckoutQuery):
		m2 := Update_UpdateBotPrecheckoutQuery{}
		m2.UpdateBotPrecheckoutQuery.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updatePhoneCall):
		m2 := Update_UpdatePhoneCall{}
		m2.UpdatePhoneCall.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateLangPackTooLong):
		m2 := Update_UpdateLangPackTooLong{}
		m2.UpdateLangPackTooLong.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateLangPack):
		m2 := Update_UpdateLangPack{}
		m2.UpdateLangPack.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateFavedStickers):
		m2 := Update_UpdateFavedStickers{}
		m2.UpdateFavedStickers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateChannelReadMessagesContents):
		m2 := Update_UpdateChannelReadMessagesContents{}
		m2.UpdateChannelReadMessagesContents.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateContactsReset):
		m2 := Update_UpdateContactsReset{}
		m2.UpdateContactsReset.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Updates_State) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Updates_State_UpdatesState:
		m2, _ := m.Payload.(*Updates_State_UpdatesState)
		b = m2.UpdatesState.Encode()
	}
	return
}

func (m *Updates_State) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_updates_state):
		m2 := Updates_State_UpdatesState{}
		m2.UpdatesState.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Updates_Difference) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Updates_Difference_UpdatesDifferenceEmpty:
		m2, _ := m.Payload.(*Updates_Difference_UpdatesDifferenceEmpty)
		b = m2.UpdatesDifferenceEmpty.Encode()
	case *Updates_Difference_UpdatesDifference:
		m2, _ := m.Payload.(*Updates_Difference_UpdatesDifference)
		b = m2.UpdatesDifference.Encode()
	case *Updates_Difference_UpdatesDifferenceSlice:
		m2, _ := m.Payload.(*Updates_Difference_UpdatesDifferenceSlice)
		b = m2.UpdatesDifferenceSlice.Encode()
	case *Updates_Difference_UpdatesDifferenceTooLong:
		m2, _ := m.Payload.(*Updates_Difference_UpdatesDifferenceTooLong)
		b = m2.UpdatesDifferenceTooLong.Encode()
	}
	return
}

func (m *Updates_Difference) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_updates_differenceEmpty):
		m2 := Updates_Difference_UpdatesDifferenceEmpty{}
		m2.UpdatesDifferenceEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updates_difference):
		m2 := Updates_Difference_UpdatesDifference{}
		m2.UpdatesDifference.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updates_differenceSlice):
		m2 := Updates_Difference_UpdatesDifferenceSlice{}
		m2.UpdatesDifferenceSlice.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updates_differenceTooLong):
		m2 := Updates_Difference_UpdatesDifferenceTooLong{}
		m2.UpdatesDifferenceTooLong.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Updates) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Updates_UpdatesTooLong:
		m2, _ := m.Payload.(*Updates_UpdatesTooLong)
		b = m2.UpdatesTooLong.Encode()
	case *Updates_UpdateShortMessage:
		m2, _ := m.Payload.(*Updates_UpdateShortMessage)
		b = m2.UpdateShortMessage.Encode()
	case *Updates_UpdateShortChatMessage:
		m2, _ := m.Payload.(*Updates_UpdateShortChatMessage)
		b = m2.UpdateShortChatMessage.Encode()
	case *Updates_UpdateShort:
		m2, _ := m.Payload.(*Updates_UpdateShort)
		b = m2.UpdateShort.Encode()
	case *Updates_UpdatesCombined:
		m2, _ := m.Payload.(*Updates_UpdatesCombined)
		b = m2.UpdatesCombined.Encode()
	case *Updates_Updates:
		m2, _ := m.Payload.(*Updates_Updates)
		b = m2.Updates.Encode()
	case *Updates_UpdateShortSentMessage:
		m2, _ := m.Payload.(*Updates_UpdateShortSentMessage)
		b = m2.UpdateShortSentMessage.Encode()
	}
	return
}

func (m *Updates) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_updatesTooLong):
		m2 := Updates_UpdatesTooLong{}
		m2.UpdatesTooLong.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateShortMessage):
		m2 := Updates_UpdateShortMessage{}
		m2.UpdateShortMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateShortChatMessage):
		m2 := Updates_UpdateShortChatMessage{}
		m2.UpdateShortChatMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateShort):
		m2 := Updates_UpdateShort{}
		m2.UpdateShort.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updatesCombined):
		m2 := Updates_UpdatesCombined{}
		m2.UpdatesCombined.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updates):
		m2 := Updates_Updates{}
		m2.Updates.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updateShortSentMessage):
		m2 := Updates_UpdateShortSentMessage{}
		m2.UpdateShortSentMessage.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Photos_Photos) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Photos_Photos_PhotosPhotos:
		m2, _ := m.Payload.(*Photos_Photos_PhotosPhotos)
		b = m2.PhotosPhotos.Encode()
	case *Photos_Photos_PhotosPhotosSlice:
		m2, _ := m.Payload.(*Photos_Photos_PhotosPhotosSlice)
		b = m2.PhotosPhotosSlice.Encode()
	}
	return
}

func (m *Photos_Photos) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_photos_photos):
		m2 := Photos_Photos_PhotosPhotos{}
		m2.PhotosPhotos.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_photos_photosSlice):
		m2 := Photos_Photos_PhotosPhotosSlice{}
		m2.PhotosPhotosSlice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Photos_Photo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Photos_Photo_PhotosPhoto:
		m2, _ := m.Payload.(*Photos_Photo_PhotosPhoto)
		b = m2.PhotosPhoto.Encode()
	}
	return
}

func (m *Photos_Photo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_photos_photo):
		m2 := Photos_Photo_PhotosPhoto{}
		m2.PhotosPhoto.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Upload_File) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Upload_File_UploadFile:
		m2, _ := m.Payload.(*Upload_File_UploadFile)
		b = m2.UploadFile.Encode()
	case *Upload_File_UploadFileCdnRedirect:
		m2, _ := m.Payload.(*Upload_File_UploadFileCdnRedirect)
		b = m2.UploadFileCdnRedirect.Encode()
	}
	return
}

func (m *Upload_File) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_upload_file):
		m2 := Upload_File_UploadFile{}
		m2.UploadFile.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_upload_fileCdnRedirect):
		m2 := Upload_File_UploadFileCdnRedirect{}
		m2.UploadFileCdnRedirect.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DcOption) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DcOption_DcOption:
		m2, _ := m.Payload.(*DcOption_DcOption)
		b = m2.DcOption.Encode()
	}
	return
}

func (m *DcOption) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_dcOption):
		m2 := DcOption_DcOption{}
		m2.DcOption.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Config) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Config_Config:
		m2, _ := m.Payload.(*Config_Config)
		b = m2.Config.Encode()
	}
	return
}

func (m *Config) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_config):
		m2 := Config_Config{}
		m2.Config.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *NearestDc) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *NearestDc_NearestDc:
		m2, _ := m.Payload.(*NearestDc_NearestDc)
		b = m2.NearestDc.Encode()
	}
	return
}

func (m *NearestDc) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_nearestDc):
		m2 := NearestDc_NearestDc{}
		m2.NearestDc.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Help_AppUpdate) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Help_AppUpdate_HelpAppUpdate:
		m2, _ := m.Payload.(*Help_AppUpdate_HelpAppUpdate)
		b = m2.HelpAppUpdate.Encode()
	case *Help_AppUpdate_HelpNoAppUpdate:
		m2, _ := m.Payload.(*Help_AppUpdate_HelpNoAppUpdate)
		b = m2.HelpNoAppUpdate.Encode()
	}
	return
}

func (m *Help_AppUpdate) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_help_appUpdate):
		m2 := Help_AppUpdate_HelpAppUpdate{}
		m2.HelpAppUpdate.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_help_noAppUpdate):
		m2 := Help_AppUpdate_HelpNoAppUpdate{}
		m2.HelpNoAppUpdate.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Help_InviteText) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Help_InviteText_HelpInviteText:
		m2, _ := m.Payload.(*Help_InviteText_HelpInviteText)
		b = m2.HelpInviteText.Encode()
	}
	return
}

func (m *Help_InviteText) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_help_inviteText):
		m2 := Help_InviteText_HelpInviteText{}
		m2.HelpInviteText.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *EncryptedChat) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *EncryptedChat_EncryptedChatEmpty:
		m2, _ := m.Payload.(*EncryptedChat_EncryptedChatEmpty)
		b = m2.EncryptedChatEmpty.Encode()
	case *EncryptedChat_EncryptedChatWaiting:
		m2, _ := m.Payload.(*EncryptedChat_EncryptedChatWaiting)
		b = m2.EncryptedChatWaiting.Encode()
	case *EncryptedChat_EncryptedChatRequested:
		m2, _ := m.Payload.(*EncryptedChat_EncryptedChatRequested)
		b = m2.EncryptedChatRequested.Encode()
	case *EncryptedChat_EncryptedChat:
		m2, _ := m.Payload.(*EncryptedChat_EncryptedChat)
		b = m2.EncryptedChat.Encode()
	case *EncryptedChat_EncryptedChatDiscarded:
		m2, _ := m.Payload.(*EncryptedChat_EncryptedChatDiscarded)
		b = m2.EncryptedChatDiscarded.Encode()
	}
	return
}

func (m *EncryptedChat) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_encryptedChatEmpty):
		m2 := EncryptedChat_EncryptedChatEmpty{}
		m2.EncryptedChatEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_encryptedChatWaiting):
		m2 := EncryptedChat_EncryptedChatWaiting{}
		m2.EncryptedChatWaiting.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_encryptedChatRequested):
		m2 := EncryptedChat_EncryptedChatRequested{}
		m2.EncryptedChatRequested.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_encryptedChat):
		m2 := EncryptedChat_EncryptedChat{}
		m2.EncryptedChat.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_encryptedChatDiscarded):
		m2 := EncryptedChat_EncryptedChatDiscarded{}
		m2.EncryptedChatDiscarded.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputEncryptedChat) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputEncryptedChat_InputEncryptedChat:
		m2, _ := m.Payload.(*InputEncryptedChat_InputEncryptedChat)
		b = m2.InputEncryptedChat.Encode()
	}
	return
}

func (m *InputEncryptedChat) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputEncryptedChat):
		m2 := InputEncryptedChat_InputEncryptedChat{}
		m2.InputEncryptedChat.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *EncryptedFile) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *EncryptedFile_EncryptedFileEmpty:
		m2, _ := m.Payload.(*EncryptedFile_EncryptedFileEmpty)
		b = m2.EncryptedFileEmpty.Encode()
	case *EncryptedFile_EncryptedFile:
		m2, _ := m.Payload.(*EncryptedFile_EncryptedFile)
		b = m2.EncryptedFile.Encode()
	}
	return
}

func (m *EncryptedFile) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_encryptedFileEmpty):
		m2 := EncryptedFile_EncryptedFileEmpty{}
		m2.EncryptedFileEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_encryptedFile):
		m2 := EncryptedFile_EncryptedFile{}
		m2.EncryptedFile.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputEncryptedFile) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputEncryptedFile_InputEncryptedFileEmpty:
		m2, _ := m.Payload.(*InputEncryptedFile_InputEncryptedFileEmpty)
		b = m2.InputEncryptedFileEmpty.Encode()
	case *InputEncryptedFile_InputEncryptedFileUploaded:
		m2, _ := m.Payload.(*InputEncryptedFile_InputEncryptedFileUploaded)
		b = m2.InputEncryptedFileUploaded.Encode()
	case *InputEncryptedFile_InputEncryptedFile:
		m2, _ := m.Payload.(*InputEncryptedFile_InputEncryptedFile)
		b = m2.InputEncryptedFile.Encode()
	case *InputEncryptedFile_InputEncryptedFileBigUploaded:
		m2, _ := m.Payload.(*InputEncryptedFile_InputEncryptedFileBigUploaded)
		b = m2.InputEncryptedFileBigUploaded.Encode()
	}
	return
}

func (m *InputEncryptedFile) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputEncryptedFileEmpty):
		m2 := InputEncryptedFile_InputEncryptedFileEmpty{}
		m2.InputEncryptedFileEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputEncryptedFileUploaded):
		m2 := InputEncryptedFile_InputEncryptedFileUploaded{}
		m2.InputEncryptedFileUploaded.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputEncryptedFile):
		m2 := InputEncryptedFile_InputEncryptedFile{}
		m2.InputEncryptedFile.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputEncryptedFileBigUploaded):
		m2 := InputEncryptedFile_InputEncryptedFileBigUploaded{}
		m2.InputEncryptedFileBigUploaded.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *EncryptedMessage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *EncryptedMessage_EncryptedMessage:
		m2, _ := m.Payload.(*EncryptedMessage_EncryptedMessage)
		b = m2.EncryptedMessage.Encode()
	case *EncryptedMessage_EncryptedMessageService:
		m2, _ := m.Payload.(*EncryptedMessage_EncryptedMessageService)
		b = m2.EncryptedMessageService.Encode()
	}
	return
}

func (m *EncryptedMessage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_encryptedMessage):
		m2 := EncryptedMessage_EncryptedMessage{}
		m2.EncryptedMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_encryptedMessageService):
		m2 := EncryptedMessage_EncryptedMessageService{}
		m2.EncryptedMessageService.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_DhConfig) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_DhConfig_MessagesDhConfigNotModified:
		m2, _ := m.Payload.(*Messages_DhConfig_MessagesDhConfigNotModified)
		b = m2.MessagesDhConfigNotModified.Encode()
	case *Messages_DhConfig_MessagesDhConfig:
		m2, _ := m.Payload.(*Messages_DhConfig_MessagesDhConfig)
		b = m2.MessagesDhConfig.Encode()
	}
	return
}

func (m *Messages_DhConfig) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_dhConfigNotModified):
		m2 := Messages_DhConfig_MessagesDhConfigNotModified{}
		m2.MessagesDhConfigNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_dhConfig):
		m2 := Messages_DhConfig_MessagesDhConfig{}
		m2.MessagesDhConfig.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_SentEncryptedMessage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_SentEncryptedMessage_MessagesSentEncryptedMessage:
		m2, _ := m.Payload.(*Messages_SentEncryptedMessage_MessagesSentEncryptedMessage)
		b = m2.MessagesSentEncryptedMessage.Encode()
	case *Messages_SentEncryptedMessage_MessagesSentEncryptedFile:
		m2, _ := m.Payload.(*Messages_SentEncryptedMessage_MessagesSentEncryptedFile)
		b = m2.MessagesSentEncryptedFile.Encode()
	}
	return
}

func (m *Messages_SentEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_sentEncryptedMessage):
		m2 := Messages_SentEncryptedMessage_MessagesSentEncryptedMessage{}
		m2.MessagesSentEncryptedMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_sentEncryptedFile):
		m2 := Messages_SentEncryptedMessage_MessagesSentEncryptedFile{}
		m2.MessagesSentEncryptedFile.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputDocument) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputDocument_InputDocumentEmpty:
		m2, _ := m.Payload.(*InputDocument_InputDocumentEmpty)
		b = m2.InputDocumentEmpty.Encode()
	case *InputDocument_InputDocument:
		m2, _ := m.Payload.(*InputDocument_InputDocument)
		b = m2.InputDocument.Encode()
	}
	return
}

func (m *InputDocument) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputDocumentEmpty):
		m2 := InputDocument_InputDocumentEmpty{}
		m2.InputDocumentEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputDocument):
		m2 := InputDocument_InputDocument{}
		m2.InputDocument.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Document) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Document_DocumentEmpty:
		m2, _ := m.Payload.(*Document_DocumentEmpty)
		b = m2.DocumentEmpty.Encode()
	case *Document_Document:
		m2, _ := m.Payload.(*Document_Document)
		b = m2.Document.Encode()
	}
	return
}

func (m *Document) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_documentEmpty):
		m2 := Document_DocumentEmpty{}
		m2.DocumentEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_document):
		m2 := Document_Document{}
		m2.Document.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Help_Support) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Help_Support_HelpSupport:
		m2, _ := m.Payload.(*Help_Support_HelpSupport)
		b = m2.HelpSupport.Encode()
	}
	return
}

func (m *Help_Support) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_help_support):
		m2 := Help_Support_HelpSupport{}
		m2.HelpSupport.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *NotifyPeer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *NotifyPeer_NotifyPeer:
		m2, _ := m.Payload.(*NotifyPeer_NotifyPeer)
		b = m2.NotifyPeer.Encode()
	case *NotifyPeer_NotifyUsers:
		m2, _ := m.Payload.(*NotifyPeer_NotifyUsers)
		b = m2.NotifyUsers.Encode()
	case *NotifyPeer_NotifyChats:
		m2, _ := m.Payload.(*NotifyPeer_NotifyChats)
		b = m2.NotifyChats.Encode()
	case *NotifyPeer_NotifyAll:
		m2, _ := m.Payload.(*NotifyPeer_NotifyAll)
		b = m2.NotifyAll.Encode()
	}
	return
}

func (m *NotifyPeer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_notifyPeer):
		m2 := NotifyPeer_NotifyPeer{}
		m2.NotifyPeer.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_notifyUsers):
		m2 := NotifyPeer_NotifyUsers{}
		m2.NotifyUsers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_notifyChats):
		m2 := NotifyPeer_NotifyChats{}
		m2.NotifyChats.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_notifyAll):
		m2 := NotifyPeer_NotifyAll{}
		m2.NotifyAll.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *SendMessageAction) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *SendMessageAction_SendMessageTypingAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageTypingAction)
		b = m2.SendMessageTypingAction.Encode()
	case *SendMessageAction_SendMessageCancelAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageCancelAction)
		b = m2.SendMessageCancelAction.Encode()
	case *SendMessageAction_SendMessageRecordVideoAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageRecordVideoAction)
		b = m2.SendMessageRecordVideoAction.Encode()
	case *SendMessageAction_SendMessageUploadVideoAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageUploadVideoAction)
		b = m2.SendMessageUploadVideoAction.Encode()
	case *SendMessageAction_SendMessageRecordAudioAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageRecordAudioAction)
		b = m2.SendMessageRecordAudioAction.Encode()
	case *SendMessageAction_SendMessageUploadAudioAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageUploadAudioAction)
		b = m2.SendMessageUploadAudioAction.Encode()
	case *SendMessageAction_SendMessageUploadPhotoAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageUploadPhotoAction)
		b = m2.SendMessageUploadPhotoAction.Encode()
	case *SendMessageAction_SendMessageUploadDocumentAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageUploadDocumentAction)
		b = m2.SendMessageUploadDocumentAction.Encode()
	case *SendMessageAction_SendMessageGeoLocationAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageGeoLocationAction)
		b = m2.SendMessageGeoLocationAction.Encode()
	case *SendMessageAction_SendMessageChooseContactAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageChooseContactAction)
		b = m2.SendMessageChooseContactAction.Encode()
	case *SendMessageAction_SendMessageGamePlayAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageGamePlayAction)
		b = m2.SendMessageGamePlayAction.Encode()
	case *SendMessageAction_SendMessageRecordRoundAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageRecordRoundAction)
		b = m2.SendMessageRecordRoundAction.Encode()
	case *SendMessageAction_SendMessageUploadRoundAction:
		m2, _ := m.Payload.(*SendMessageAction_SendMessageUploadRoundAction)
		b = m2.SendMessageUploadRoundAction.Encode()
	}
	return
}

func (m *SendMessageAction) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_sendMessageTypingAction):
		m2 := SendMessageAction_SendMessageTypingAction{}
		m2.SendMessageTypingAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageCancelAction):
		m2 := SendMessageAction_SendMessageCancelAction{}
		m2.SendMessageCancelAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageRecordVideoAction):
		m2 := SendMessageAction_SendMessageRecordVideoAction{}
		m2.SendMessageRecordVideoAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageUploadVideoAction):
		m2 := SendMessageAction_SendMessageUploadVideoAction{}
		m2.SendMessageUploadVideoAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageRecordAudioAction):
		m2 := SendMessageAction_SendMessageRecordAudioAction{}
		m2.SendMessageRecordAudioAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageUploadAudioAction):
		m2 := SendMessageAction_SendMessageUploadAudioAction{}
		m2.SendMessageUploadAudioAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageUploadPhotoAction):
		m2 := SendMessageAction_SendMessageUploadPhotoAction{}
		m2.SendMessageUploadPhotoAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageUploadDocumentAction):
		m2 := SendMessageAction_SendMessageUploadDocumentAction{}
		m2.SendMessageUploadDocumentAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageGeoLocationAction):
		m2 := SendMessageAction_SendMessageGeoLocationAction{}
		m2.SendMessageGeoLocationAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageChooseContactAction):
		m2 := SendMessageAction_SendMessageChooseContactAction{}
		m2.SendMessageChooseContactAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageGamePlayAction):
		m2 := SendMessageAction_SendMessageGamePlayAction{}
		m2.SendMessageGamePlayAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageRecordRoundAction):
		m2 := SendMessageAction_SendMessageRecordRoundAction{}
		m2.SendMessageRecordRoundAction.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_sendMessageUploadRoundAction):
		m2 := SendMessageAction_SendMessageUploadRoundAction{}
		m2.SendMessageUploadRoundAction.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_Found) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_Found_ContactsFound:
		m2, _ := m.Payload.(*Contacts_Found_ContactsFound)
		b = m2.ContactsFound.Encode()
	}
	return
}

func (m *Contacts_Found) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_found):
		m2 := Contacts_Found_ContactsFound{}
		m2.ContactsFound.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPrivacyKey) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPrivacyKey_InputPrivacyKeyStatusTimestamp:
		m2, _ := m.Payload.(*InputPrivacyKey_InputPrivacyKeyStatusTimestamp)
		b = m2.InputPrivacyKeyStatusTimestamp.Encode()
	case *InputPrivacyKey_InputPrivacyKeyChatInvite:
		m2, _ := m.Payload.(*InputPrivacyKey_InputPrivacyKeyChatInvite)
		b = m2.InputPrivacyKeyChatInvite.Encode()
	case *InputPrivacyKey_InputPrivacyKeyPhoneCall:
		m2, _ := m.Payload.(*InputPrivacyKey_InputPrivacyKeyPhoneCall)
		b = m2.InputPrivacyKeyPhoneCall.Encode()
	}
	return
}

func (m *InputPrivacyKey) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPrivacyKeyStatusTimestamp):
		m2 := InputPrivacyKey_InputPrivacyKeyStatusTimestamp{}
		m2.InputPrivacyKeyStatusTimestamp.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyKeyChatInvite):
		m2 := InputPrivacyKey_InputPrivacyKeyChatInvite{}
		m2.InputPrivacyKeyChatInvite.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyKeyPhoneCall):
		m2 := InputPrivacyKey_InputPrivacyKeyPhoneCall{}
		m2.InputPrivacyKeyPhoneCall.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PrivacyKey) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PrivacyKey_PrivacyKeyStatusTimestamp:
		m2, _ := m.Payload.(*PrivacyKey_PrivacyKeyStatusTimestamp)
		b = m2.PrivacyKeyStatusTimestamp.Encode()
	case *PrivacyKey_PrivacyKeyChatInvite:
		m2, _ := m.Payload.(*PrivacyKey_PrivacyKeyChatInvite)
		b = m2.PrivacyKeyChatInvite.Encode()
	case *PrivacyKey_PrivacyKeyPhoneCall:
		m2, _ := m.Payload.(*PrivacyKey_PrivacyKeyPhoneCall)
		b = m2.PrivacyKeyPhoneCall.Encode()
	}
	return
}

func (m *PrivacyKey) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_privacyKeyStatusTimestamp):
		m2 := PrivacyKey_PrivacyKeyStatusTimestamp{}
		m2.PrivacyKeyStatusTimestamp.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyKeyChatInvite):
		m2 := PrivacyKey_PrivacyKeyChatInvite{}
		m2.PrivacyKeyChatInvite.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyKeyPhoneCall):
		m2 := PrivacyKey_PrivacyKeyPhoneCall{}
		m2.PrivacyKeyPhoneCall.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPrivacyRule) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPrivacyRule_InputPrivacyValueAllowContacts:
		m2, _ := m.Payload.(*InputPrivacyRule_InputPrivacyValueAllowContacts)
		b = m2.InputPrivacyValueAllowContacts.Encode()
	case *InputPrivacyRule_InputPrivacyValueAllowAll:
		m2, _ := m.Payload.(*InputPrivacyRule_InputPrivacyValueAllowAll)
		b = m2.InputPrivacyValueAllowAll.Encode()
	case *InputPrivacyRule_InputPrivacyValueAllowUsers:
		m2, _ := m.Payload.(*InputPrivacyRule_InputPrivacyValueAllowUsers)
		b = m2.InputPrivacyValueAllowUsers.Encode()
	case *InputPrivacyRule_InputPrivacyValueDisallowContacts:
		m2, _ := m.Payload.(*InputPrivacyRule_InputPrivacyValueDisallowContacts)
		b = m2.InputPrivacyValueDisallowContacts.Encode()
	case *InputPrivacyRule_InputPrivacyValueDisallowAll:
		m2, _ := m.Payload.(*InputPrivacyRule_InputPrivacyValueDisallowAll)
		b = m2.InputPrivacyValueDisallowAll.Encode()
	case *InputPrivacyRule_InputPrivacyValueDisallowUsers:
		m2, _ := m.Payload.(*InputPrivacyRule_InputPrivacyValueDisallowUsers)
		b = m2.InputPrivacyValueDisallowUsers.Encode()
	}
	return
}

func (m *InputPrivacyRule) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPrivacyValueAllowContacts):
		m2 := InputPrivacyRule_InputPrivacyValueAllowContacts{}
		m2.InputPrivacyValueAllowContacts.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyValueAllowAll):
		m2 := InputPrivacyRule_InputPrivacyValueAllowAll{}
		m2.InputPrivacyValueAllowAll.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyValueAllowUsers):
		m2 := InputPrivacyRule_InputPrivacyValueAllowUsers{}
		m2.InputPrivacyValueAllowUsers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyValueDisallowContacts):
		m2 := InputPrivacyRule_InputPrivacyValueDisallowContacts{}
		m2.InputPrivacyValueDisallowContacts.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyValueDisallowAll):
		m2 := InputPrivacyRule_InputPrivacyValueDisallowAll{}
		m2.InputPrivacyValueDisallowAll.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPrivacyValueDisallowUsers):
		m2 := InputPrivacyRule_InputPrivacyValueDisallowUsers{}
		m2.InputPrivacyValueDisallowUsers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PrivacyRule) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PrivacyRule_PrivacyValueAllowContacts:
		m2, _ := m.Payload.(*PrivacyRule_PrivacyValueAllowContacts)
		b = m2.PrivacyValueAllowContacts.Encode()
	case *PrivacyRule_PrivacyValueAllowAll:
		m2, _ := m.Payload.(*PrivacyRule_PrivacyValueAllowAll)
		b = m2.PrivacyValueAllowAll.Encode()
	case *PrivacyRule_PrivacyValueAllowUsers:
		m2, _ := m.Payload.(*PrivacyRule_PrivacyValueAllowUsers)
		b = m2.PrivacyValueAllowUsers.Encode()
	case *PrivacyRule_PrivacyValueDisallowContacts:
		m2, _ := m.Payload.(*PrivacyRule_PrivacyValueDisallowContacts)
		b = m2.PrivacyValueDisallowContacts.Encode()
	case *PrivacyRule_PrivacyValueDisallowAll:
		m2, _ := m.Payload.(*PrivacyRule_PrivacyValueDisallowAll)
		b = m2.PrivacyValueDisallowAll.Encode()
	case *PrivacyRule_PrivacyValueDisallowUsers:
		m2, _ := m.Payload.(*PrivacyRule_PrivacyValueDisallowUsers)
		b = m2.PrivacyValueDisallowUsers.Encode()
	}
	return
}

func (m *PrivacyRule) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_privacyValueAllowContacts):
		m2 := PrivacyRule_PrivacyValueAllowContacts{}
		m2.PrivacyValueAllowContacts.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyValueAllowAll):
		m2 := PrivacyRule_PrivacyValueAllowAll{}
		m2.PrivacyValueAllowAll.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyValueAllowUsers):
		m2 := PrivacyRule_PrivacyValueAllowUsers{}
		m2.PrivacyValueAllowUsers.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyValueDisallowContacts):
		m2 := PrivacyRule_PrivacyValueDisallowContacts{}
		m2.PrivacyValueDisallowContacts.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyValueDisallowAll):
		m2 := PrivacyRule_PrivacyValueDisallowAll{}
		m2.PrivacyValueDisallowAll.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_privacyValueDisallowUsers):
		m2 := PrivacyRule_PrivacyValueDisallowUsers{}
		m2.PrivacyValueDisallowUsers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Account_PrivacyRules) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Account_PrivacyRules_AccountPrivacyRules:
		m2, _ := m.Payload.(*Account_PrivacyRules_AccountPrivacyRules)
		b = m2.AccountPrivacyRules.Encode()
	}
	return
}

func (m *Account_PrivacyRules) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_account_privacyRules):
		m2 := Account_PrivacyRules_AccountPrivacyRules{}
		m2.AccountPrivacyRules.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *AccountDaysTTL) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *AccountDaysTTL_AccountDaysTTL:
		m2, _ := m.Payload.(*AccountDaysTTL_AccountDaysTTL)
		b = m2.AccountDaysTTL.Encode()
	}
	return
}

func (m *AccountDaysTTL) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_accountDaysTTL):
		m2 := AccountDaysTTL_AccountDaysTTL{}
		m2.AccountDaysTTL.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DocumentAttribute) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DocumentAttribute_DocumentAttributeImageSize:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeImageSize)
		b = m2.DocumentAttributeImageSize.Encode()
	case *DocumentAttribute_DocumentAttributeAnimated:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeAnimated)
		b = m2.DocumentAttributeAnimated.Encode()
	case *DocumentAttribute_DocumentAttributeSticker:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeSticker)
		b = m2.DocumentAttributeSticker.Encode()
	case *DocumentAttribute_DocumentAttributeVideo:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeVideo)
		b = m2.DocumentAttributeVideo.Encode()
	case *DocumentAttribute_DocumentAttributeAudio:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeAudio)
		b = m2.DocumentAttributeAudio.Encode()
	case *DocumentAttribute_DocumentAttributeFilename:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeFilename)
		b = m2.DocumentAttributeFilename.Encode()
	case *DocumentAttribute_DocumentAttributeHasStickers:
		m2, _ := m.Payload.(*DocumentAttribute_DocumentAttributeHasStickers)
		b = m2.DocumentAttributeHasStickers.Encode()
	}
	return
}

func (m *DocumentAttribute) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_documentAttributeImageSize):
		m2 := DocumentAttribute_DocumentAttributeImageSize{}
		m2.DocumentAttributeImageSize.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_documentAttributeAnimated):
		m2 := DocumentAttribute_DocumentAttributeAnimated{}
		m2.DocumentAttributeAnimated.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_documentAttributeSticker):
		m2 := DocumentAttribute_DocumentAttributeSticker{}
		m2.DocumentAttributeSticker.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_documentAttributeVideo):
		m2 := DocumentAttribute_DocumentAttributeVideo{}
		m2.DocumentAttributeVideo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_documentAttributeAudio):
		m2 := DocumentAttribute_DocumentAttributeAudio{}
		m2.DocumentAttributeAudio.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_documentAttributeFilename):
		m2 := DocumentAttribute_DocumentAttributeFilename{}
		m2.DocumentAttributeFilename.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_documentAttributeHasStickers):
		m2 := DocumentAttribute_DocumentAttributeHasStickers{}
		m2.DocumentAttributeHasStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_Stickers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_Stickers_MessagesStickersNotModified:
		m2, _ := m.Payload.(*Messages_Stickers_MessagesStickersNotModified)
		b = m2.MessagesStickersNotModified.Encode()
	case *Messages_Stickers_MessagesStickers:
		m2, _ := m.Payload.(*Messages_Stickers_MessagesStickers)
		b = m2.MessagesStickers.Encode()
	}
	return
}

func (m *Messages_Stickers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_stickersNotModified):
		m2 := Messages_Stickers_MessagesStickersNotModified{}
		m2.MessagesStickersNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_stickers):
		m2 := Messages_Stickers_MessagesStickers{}
		m2.MessagesStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *StickerPack) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *StickerPack_StickerPack:
		m2, _ := m.Payload.(*StickerPack_StickerPack)
		b = m2.StickerPack.Encode()
	}
	return
}

func (m *StickerPack) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_stickerPack):
		m2 := StickerPack_StickerPack{}
		m2.StickerPack.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_AllStickers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_AllStickers_MessagesAllStickersNotModified:
		m2, _ := m.Payload.(*Messages_AllStickers_MessagesAllStickersNotModified)
		b = m2.MessagesAllStickersNotModified.Encode()
	case *Messages_AllStickers_MessagesAllStickers:
		m2, _ := m.Payload.(*Messages_AllStickers_MessagesAllStickers)
		b = m2.MessagesAllStickers.Encode()
	}
	return
}

func (m *Messages_AllStickers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_allStickersNotModified):
		m2 := Messages_AllStickers_MessagesAllStickersNotModified{}
		m2.MessagesAllStickersNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_allStickers):
		m2 := Messages_AllStickers_MessagesAllStickers{}
		m2.MessagesAllStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DisabledFeature) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DisabledFeature_DisabledFeature:
		m2, _ := m.Payload.(*DisabledFeature_DisabledFeature)
		b = m2.DisabledFeature.Encode()
	}
	return
}

func (m *DisabledFeature) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_disabledFeature):
		m2 := DisabledFeature_DisabledFeature{}
		m2.DisabledFeature.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_AffectedMessages) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_AffectedMessages_MessagesAffectedMessages:
		m2, _ := m.Payload.(*Messages_AffectedMessages_MessagesAffectedMessages)
		b = m2.MessagesAffectedMessages.Encode()
	}
	return
}

func (m *Messages_AffectedMessages) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_affectedMessages):
		m2 := Messages_AffectedMessages_MessagesAffectedMessages{}
		m2.MessagesAffectedMessages.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ContactLink) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ContactLink_ContactLinkUnknown:
		m2, _ := m.Payload.(*ContactLink_ContactLinkUnknown)
		b = m2.ContactLinkUnknown.Encode()
	case *ContactLink_ContactLinkNone:
		m2, _ := m.Payload.(*ContactLink_ContactLinkNone)
		b = m2.ContactLinkNone.Encode()
	case *ContactLink_ContactLinkHasPhone:
		m2, _ := m.Payload.(*ContactLink_ContactLinkHasPhone)
		b = m2.ContactLinkHasPhone.Encode()
	case *ContactLink_ContactLinkContact:
		m2, _ := m.Payload.(*ContactLink_ContactLinkContact)
		b = m2.ContactLinkContact.Encode()
	}
	return
}

func (m *ContactLink) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contactLinkUnknown):
		m2 := ContactLink_ContactLinkUnknown{}
		m2.ContactLinkUnknown.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_contactLinkNone):
		m2 := ContactLink_ContactLinkNone{}
		m2.ContactLinkNone.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_contactLinkHasPhone):
		m2 := ContactLink_ContactLinkHasPhone{}
		m2.ContactLinkHasPhone.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_contactLinkContact):
		m2 := ContactLink_ContactLinkContact{}
		m2.ContactLinkContact.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *WebPage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *WebPage_WebPageEmpty:
		m2, _ := m.Payload.(*WebPage_WebPageEmpty)
		b = m2.WebPageEmpty.Encode()
	case *WebPage_WebPagePending:
		m2, _ := m.Payload.(*WebPage_WebPagePending)
		b = m2.WebPagePending.Encode()
	case *WebPage_WebPage:
		m2, _ := m.Payload.(*WebPage_WebPage)
		b = m2.WebPage.Encode()
	case *WebPage_WebPageNotModified:
		m2, _ := m.Payload.(*WebPage_WebPageNotModified)
		b = m2.WebPageNotModified.Encode()
	}
	return
}

func (m *WebPage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_webPageEmpty):
		m2 := WebPage_WebPageEmpty{}
		m2.WebPageEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_webPagePending):
		m2 := WebPage_WebPagePending{}
		m2.WebPagePending.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_webPage):
		m2 := WebPage_WebPage{}
		m2.WebPage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_webPageNotModified):
		m2 := WebPage_WebPageNotModified{}
		m2.WebPageNotModified.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Authorization) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Authorization_Authorization:
		m2, _ := m.Payload.(*Authorization_Authorization)
		b = m2.Authorization.Encode()
	}
	return
}

func (m *Authorization) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_authorization):
		m2 := Authorization_Authorization{}
		m2.Authorization.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Account_Authorizations) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Account_Authorizations_AccountAuthorizations:
		m2, _ := m.Payload.(*Account_Authorizations_AccountAuthorizations)
		b = m2.AccountAuthorizations.Encode()
	}
	return
}

func (m *Account_Authorizations) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_account_authorizations):
		m2 := Account_Authorizations_AccountAuthorizations{}
		m2.AccountAuthorizations.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Account_Password) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Account_Password_AccountNoPassword:
		m2, _ := m.Payload.(*Account_Password_AccountNoPassword)
		b = m2.AccountNoPassword.Encode()
	case *Account_Password_AccountPassword:
		m2, _ := m.Payload.(*Account_Password_AccountPassword)
		b = m2.AccountPassword.Encode()
	}
	return
}

func (m *Account_Password) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_account_noPassword):
		m2 := Account_Password_AccountNoPassword{}
		m2.AccountNoPassword.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_account_password):
		m2 := Account_Password_AccountPassword{}
		m2.AccountPassword.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Account_PasswordSettings) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Account_PasswordSettings_AccountPasswordSettings:
		m2, _ := m.Payload.(*Account_PasswordSettings_AccountPasswordSettings)
		b = m2.AccountPasswordSettings.Encode()
	}
	return
}

func (m *Account_PasswordSettings) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_account_passwordSettings):
		m2 := Account_PasswordSettings_AccountPasswordSettings{}
		m2.AccountPasswordSettings.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Account_PasswordInputSettings) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Account_PasswordInputSettings_AccountPasswordInputSettings:
		m2, _ := m.Payload.(*Account_PasswordInputSettings_AccountPasswordInputSettings)
		b = m2.AccountPasswordInputSettings.Encode()
	}
	return
}

func (m *Account_PasswordInputSettings) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_account_passwordInputSettings):
		m2 := Account_PasswordInputSettings_AccountPasswordInputSettings{}
		m2.AccountPasswordInputSettings.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_PasswordRecovery) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_PasswordRecovery_AuthPasswordRecovery:
		m2, _ := m.Payload.(*Auth_PasswordRecovery_AuthPasswordRecovery)
		b = m2.AuthPasswordRecovery.Encode()
	}
	return
}

func (m *Auth_PasswordRecovery) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_passwordRecovery):
		m2 := Auth_PasswordRecovery_AuthPasswordRecovery{}
		m2.AuthPasswordRecovery.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ReceivedNotifyMessage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ReceivedNotifyMessage_ReceivedNotifyMessage:
		m2, _ := m.Payload.(*ReceivedNotifyMessage_ReceivedNotifyMessage)
		b = m2.ReceivedNotifyMessage.Encode()
	}
	return
}

func (m *ReceivedNotifyMessage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_receivedNotifyMessage):
		m2 := ReceivedNotifyMessage_ReceivedNotifyMessage{}
		m2.ReceivedNotifyMessage.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ExportedChatInvite) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ExportedChatInvite_ChatInviteEmpty:
		m2, _ := m.Payload.(*ExportedChatInvite_ChatInviteEmpty)
		b = m2.ChatInviteEmpty.Encode()
	case *ExportedChatInvite_ChatInviteExported:
		m2, _ := m.Payload.(*ExportedChatInvite_ChatInviteExported)
		b = m2.ChatInviteExported.Encode()
	}
	return
}

func (m *ExportedChatInvite) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatInviteEmpty):
		m2 := ExportedChatInvite_ChatInviteEmpty{}
		m2.ChatInviteEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatInviteExported):
		m2 := ExportedChatInvite_ChatInviteExported{}
		m2.ChatInviteExported.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChatInvite) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChatInvite_ChatInviteAlready:
		m2, _ := m.Payload.(*ChatInvite_ChatInviteAlready)
		b = m2.ChatInviteAlready.Encode()
	case *ChatInvite_ChatInvite:
		m2, _ := m.Payload.(*ChatInvite_ChatInvite)
		b = m2.ChatInvite.Encode()
	}
	return
}

func (m *ChatInvite) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_chatInviteAlready):
		m2 := ChatInvite_ChatInviteAlready{}
		m2.ChatInviteAlready.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_chatInvite):
		m2 := ChatInvite_ChatInvite{}
		m2.ChatInvite.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputStickerSet) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputStickerSet_InputStickerSetEmpty:
		m2, _ := m.Payload.(*InputStickerSet_InputStickerSetEmpty)
		b = m2.InputStickerSetEmpty.Encode()
	case *InputStickerSet_InputStickerSetID:
		m2, _ := m.Payload.(*InputStickerSet_InputStickerSetID)
		b = m2.InputStickerSetID.Encode()
	case *InputStickerSet_InputStickerSetShortName:
		m2, _ := m.Payload.(*InputStickerSet_InputStickerSetShortName)
		b = m2.InputStickerSetShortName.Encode()
	}
	return
}

func (m *InputStickerSet) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputStickerSetEmpty):
		m2 := InputStickerSet_InputStickerSetEmpty{}
		m2.InputStickerSetEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputStickerSetID):
		m2 := InputStickerSet_InputStickerSetID{}
		m2.InputStickerSetID.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputStickerSetShortName):
		m2 := InputStickerSet_InputStickerSetShortName{}
		m2.InputStickerSetShortName.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *StickerSet) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *StickerSet_StickerSet:
		m2, _ := m.Payload.(*StickerSet_StickerSet)
		b = m2.StickerSet.Encode()
	}
	return
}

func (m *StickerSet) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_stickerSet):
		m2 := StickerSet_StickerSet{}
		m2.StickerSet.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_StickerSet) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_StickerSet_MessagesStickerSet:
		m2, _ := m.Payload.(*Messages_StickerSet_MessagesStickerSet)
		b = m2.MessagesStickerSet.Encode()
	}
	return
}

func (m *Messages_StickerSet) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_stickerSet):
		m2 := Messages_StickerSet_MessagesStickerSet{}
		m2.MessagesStickerSet.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *BotCommand) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *BotCommand_BotCommand:
		m2, _ := m.Payload.(*BotCommand_BotCommand)
		b = m2.BotCommand.Encode()
	}
	return
}

func (m *BotCommand) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_botCommand):
		m2 := BotCommand_BotCommand{}
		m2.BotCommand.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *BotInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *BotInfo_BotInfo:
		m2, _ := m.Payload.(*BotInfo_BotInfo)
		b = m2.BotInfo.Encode()
	}
	return
}

func (m *BotInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_botInfo):
		m2 := BotInfo_BotInfo{}
		m2.BotInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *KeyboardButton) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *KeyboardButton_KeyboardButton:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButton)
		b = m2.KeyboardButton.Encode()
	case *KeyboardButton_KeyboardButtonUrl:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonUrl)
		b = m2.KeyboardButtonUrl.Encode()
	case *KeyboardButton_KeyboardButtonCallback:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonCallback)
		b = m2.KeyboardButtonCallback.Encode()
	case *KeyboardButton_KeyboardButtonRequestPhone:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonRequestPhone)
		b = m2.KeyboardButtonRequestPhone.Encode()
	case *KeyboardButton_KeyboardButtonRequestGeoLocation:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonRequestGeoLocation)
		b = m2.KeyboardButtonRequestGeoLocation.Encode()
	case *KeyboardButton_KeyboardButtonSwitchInline:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonSwitchInline)
		b = m2.KeyboardButtonSwitchInline.Encode()
	case *KeyboardButton_KeyboardButtonGame:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonGame)
		b = m2.KeyboardButtonGame.Encode()
	case *KeyboardButton_KeyboardButtonBuy:
		m2, _ := m.Payload.(*KeyboardButton_KeyboardButtonBuy)
		b = m2.KeyboardButtonBuy.Encode()
	}
	return
}

func (m *KeyboardButton) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_keyboardButton):
		m2 := KeyboardButton_KeyboardButton{}
		m2.KeyboardButton.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonUrl):
		m2 := KeyboardButton_KeyboardButtonUrl{}
		m2.KeyboardButtonUrl.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonCallback):
		m2 := KeyboardButton_KeyboardButtonCallback{}
		m2.KeyboardButtonCallback.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonRequestPhone):
		m2 := KeyboardButton_KeyboardButtonRequestPhone{}
		m2.KeyboardButtonRequestPhone.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonRequestGeoLocation):
		m2 := KeyboardButton_KeyboardButtonRequestGeoLocation{}
		m2.KeyboardButtonRequestGeoLocation.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonSwitchInline):
		m2 := KeyboardButton_KeyboardButtonSwitchInline{}
		m2.KeyboardButtonSwitchInline.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonGame):
		m2 := KeyboardButton_KeyboardButtonGame{}
		m2.KeyboardButtonGame.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_keyboardButtonBuy):
		m2 := KeyboardButton_KeyboardButtonBuy{}
		m2.KeyboardButtonBuy.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *KeyboardButtonRow) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *KeyboardButtonRow_KeyboardButtonRow:
		m2, _ := m.Payload.(*KeyboardButtonRow_KeyboardButtonRow)
		b = m2.KeyboardButtonRow.Encode()
	}
	return
}

func (m *KeyboardButtonRow) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_keyboardButtonRow):
		m2 := KeyboardButtonRow_KeyboardButtonRow{}
		m2.KeyboardButtonRow.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ReplyMarkup) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ReplyMarkup_ReplyKeyboardHide:
		m2, _ := m.Payload.(*ReplyMarkup_ReplyKeyboardHide)
		b = m2.ReplyKeyboardHide.Encode()
	case *ReplyMarkup_ReplyKeyboardForceReply:
		m2, _ := m.Payload.(*ReplyMarkup_ReplyKeyboardForceReply)
		b = m2.ReplyKeyboardForceReply.Encode()
	case *ReplyMarkup_ReplyKeyboardMarkup:
		m2, _ := m.Payload.(*ReplyMarkup_ReplyKeyboardMarkup)
		b = m2.ReplyKeyboardMarkup.Encode()
	case *ReplyMarkup_ReplyInlineMarkup:
		m2, _ := m.Payload.(*ReplyMarkup_ReplyInlineMarkup)
		b = m2.ReplyInlineMarkup.Encode()
	}
	return
}

func (m *ReplyMarkup) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_replyKeyboardHide):
		m2 := ReplyMarkup_ReplyKeyboardHide{}
		m2.ReplyKeyboardHide.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_replyKeyboardForceReply):
		m2 := ReplyMarkup_ReplyKeyboardForceReply{}
		m2.ReplyKeyboardForceReply.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_replyKeyboardMarkup):
		m2 := ReplyMarkup_ReplyKeyboardMarkup{}
		m2.ReplyKeyboardMarkup.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_replyInlineMarkup):
		m2 := ReplyMarkup_ReplyInlineMarkup{}
		m2.ReplyInlineMarkup.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MessageEntity) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MessageEntity_MessageEntityUnknown:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityUnknown)
		b = m2.MessageEntityUnknown.Encode()
	case *MessageEntity_MessageEntityMention:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityMention)
		b = m2.MessageEntityMention.Encode()
	case *MessageEntity_MessageEntityHashtag:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityHashtag)
		b = m2.MessageEntityHashtag.Encode()
	case *MessageEntity_MessageEntityBotCommand:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityBotCommand)
		b = m2.MessageEntityBotCommand.Encode()
	case *MessageEntity_MessageEntityUrl:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityUrl)
		b = m2.MessageEntityUrl.Encode()
	case *MessageEntity_MessageEntityEmail:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityEmail)
		b = m2.MessageEntityEmail.Encode()
	case *MessageEntity_MessageEntityBold:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityBold)
		b = m2.MessageEntityBold.Encode()
	case *MessageEntity_MessageEntityItalic:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityItalic)
		b = m2.MessageEntityItalic.Encode()
	case *MessageEntity_MessageEntityCode:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityCode)
		b = m2.MessageEntityCode.Encode()
	case *MessageEntity_MessageEntityPre:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityPre)
		b = m2.MessageEntityPre.Encode()
	case *MessageEntity_MessageEntityTextUrl:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityTextUrl)
		b = m2.MessageEntityTextUrl.Encode()
	case *MessageEntity_MessageEntityMentionName:
		m2, _ := m.Payload.(*MessageEntity_MessageEntityMentionName)
		b = m2.MessageEntityMentionName.Encode()
	case *MessageEntity_InputMessageEntityMentionName:
		m2, _ := m.Payload.(*MessageEntity_InputMessageEntityMentionName)
		b = m2.InputMessageEntityMentionName.Encode()
	}
	return
}

func (m *MessageEntity) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messageEntityUnknown):
		m2 := MessageEntity_MessageEntityUnknown{}
		m2.MessageEntityUnknown.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityMention):
		m2 := MessageEntity_MessageEntityMention{}
		m2.MessageEntityMention.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityHashtag):
		m2 := MessageEntity_MessageEntityHashtag{}
		m2.MessageEntityHashtag.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityBotCommand):
		m2 := MessageEntity_MessageEntityBotCommand{}
		m2.MessageEntityBotCommand.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityUrl):
		m2 := MessageEntity_MessageEntityUrl{}
		m2.MessageEntityUrl.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityEmail):
		m2 := MessageEntity_MessageEntityEmail{}
		m2.MessageEntityEmail.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityBold):
		m2 := MessageEntity_MessageEntityBold{}
		m2.MessageEntityBold.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityItalic):
		m2 := MessageEntity_MessageEntityItalic{}
		m2.MessageEntityItalic.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityCode):
		m2 := MessageEntity_MessageEntityCode{}
		m2.MessageEntityCode.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityPre):
		m2 := MessageEntity_MessageEntityPre{}
		m2.MessageEntityPre.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityTextUrl):
		m2 := MessageEntity_MessageEntityTextUrl{}
		m2.MessageEntityTextUrl.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messageEntityMentionName):
		m2 := MessageEntity_MessageEntityMentionName{}
		m2.MessageEntityMentionName.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputMessageEntityMentionName):
		m2 := MessageEntity_InputMessageEntityMentionName{}
		m2.InputMessageEntityMentionName.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputChannel) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputChannel_InputChannelEmpty:
		m2, _ := m.Payload.(*InputChannel_InputChannelEmpty)
		b = m2.InputChannelEmpty.Encode()
	case *InputChannel_InputChannel:
		m2, _ := m.Payload.(*InputChannel_InputChannel)
		b = m2.InputChannel.Encode()
	}
	return
}

func (m *InputChannel) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputChannelEmpty):
		m2 := InputChannel_InputChannelEmpty{}
		m2.InputChannelEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputChannel):
		m2 := InputChannel_InputChannel{}
		m2.InputChannel.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_ResolvedPeer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_ResolvedPeer_ContactsResolvedPeer:
		m2, _ := m.Payload.(*Contacts_ResolvedPeer_ContactsResolvedPeer)
		b = m2.ContactsResolvedPeer.Encode()
	}
	return
}

func (m *Contacts_ResolvedPeer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_resolvedPeer):
		m2 := Contacts_ResolvedPeer_ContactsResolvedPeer{}
		m2.ContactsResolvedPeer.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MessageRange) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MessageRange_MessageRange:
		m2, _ := m.Payload.(*MessageRange_MessageRange)
		b = m2.MessageRange.Encode()
	}
	return
}

func (m *MessageRange) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messageRange):
		m2 := MessageRange_MessageRange{}
		m2.MessageRange.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Updates_ChannelDifference) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Updates_ChannelDifference_UpdatesChannelDifferenceEmpty:
		m2, _ := m.Payload.(*Updates_ChannelDifference_UpdatesChannelDifferenceEmpty)
		b = m2.UpdatesChannelDifferenceEmpty.Encode()
	case *Updates_ChannelDifference_UpdatesChannelDifferenceTooLong:
		m2, _ := m.Payload.(*Updates_ChannelDifference_UpdatesChannelDifferenceTooLong)
		b = m2.UpdatesChannelDifferenceTooLong.Encode()
	case *Updates_ChannelDifference_UpdatesChannelDifference:
		m2, _ := m.Payload.(*Updates_ChannelDifference_UpdatesChannelDifference)
		b = m2.UpdatesChannelDifference.Encode()
	}
	return
}

func (m *Updates_ChannelDifference) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_updates_channelDifferenceEmpty):
		m2 := Updates_ChannelDifference_UpdatesChannelDifferenceEmpty{}
		m2.UpdatesChannelDifferenceEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updates_channelDifferenceTooLong):
		m2 := Updates_ChannelDifference_UpdatesChannelDifferenceTooLong{}
		m2.UpdatesChannelDifferenceTooLong.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_updates_channelDifference):
		m2 := Updates_ChannelDifference_UpdatesChannelDifference{}
		m2.UpdatesChannelDifference.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelMessagesFilter) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelMessagesFilter_ChannelMessagesFilterEmpty:
		m2, _ := m.Payload.(*ChannelMessagesFilter_ChannelMessagesFilterEmpty)
		b = m2.ChannelMessagesFilterEmpty.Encode()
	case *ChannelMessagesFilter_ChannelMessagesFilter:
		m2, _ := m.Payload.(*ChannelMessagesFilter_ChannelMessagesFilter)
		b = m2.ChannelMessagesFilter.Encode()
	}
	return
}

func (m *ChannelMessagesFilter) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelMessagesFilterEmpty):
		m2 := ChannelMessagesFilter_ChannelMessagesFilterEmpty{}
		m2.ChannelMessagesFilterEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelMessagesFilter):
		m2 := ChannelMessagesFilter_ChannelMessagesFilter{}
		m2.ChannelMessagesFilter.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelParticipant) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelParticipant_ChannelParticipant:
		m2, _ := m.Payload.(*ChannelParticipant_ChannelParticipant)
		b = m2.ChannelParticipant.Encode()
	case *ChannelParticipant_ChannelParticipantSelf:
		m2, _ := m.Payload.(*ChannelParticipant_ChannelParticipantSelf)
		b = m2.ChannelParticipantSelf.Encode()
	case *ChannelParticipant_ChannelParticipantCreator:
		m2, _ := m.Payload.(*ChannelParticipant_ChannelParticipantCreator)
		b = m2.ChannelParticipantCreator.Encode()
	case *ChannelParticipant_ChannelParticipantAdmin:
		m2, _ := m.Payload.(*ChannelParticipant_ChannelParticipantAdmin)
		b = m2.ChannelParticipantAdmin.Encode()
	case *ChannelParticipant_ChannelParticipantBanned:
		m2, _ := m.Payload.(*ChannelParticipant_ChannelParticipantBanned)
		b = m2.ChannelParticipantBanned.Encode()
	}
	return
}

func (m *ChannelParticipant) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelParticipant):
		m2 := ChannelParticipant_ChannelParticipant{}
		m2.ChannelParticipant.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantSelf):
		m2 := ChannelParticipant_ChannelParticipantSelf{}
		m2.ChannelParticipantSelf.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantCreator):
		m2 := ChannelParticipant_ChannelParticipantCreator{}
		m2.ChannelParticipantCreator.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantAdmin):
		m2 := ChannelParticipant_ChannelParticipantAdmin{}
		m2.ChannelParticipantAdmin.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantBanned):
		m2 := ChannelParticipant_ChannelParticipantBanned{}
		m2.ChannelParticipantBanned.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelParticipantsFilter) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelParticipantsFilter_ChannelParticipantsRecent:
		m2, _ := m.Payload.(*ChannelParticipantsFilter_ChannelParticipantsRecent)
		b = m2.ChannelParticipantsRecent.Encode()
	case *ChannelParticipantsFilter_ChannelParticipantsAdmins:
		m2, _ := m.Payload.(*ChannelParticipantsFilter_ChannelParticipantsAdmins)
		b = m2.ChannelParticipantsAdmins.Encode()
	case *ChannelParticipantsFilter_ChannelParticipantsKicked:
		m2, _ := m.Payload.(*ChannelParticipantsFilter_ChannelParticipantsKicked)
		b = m2.ChannelParticipantsKicked.Encode()
	case *ChannelParticipantsFilter_ChannelParticipantsBots:
		m2, _ := m.Payload.(*ChannelParticipantsFilter_ChannelParticipantsBots)
		b = m2.ChannelParticipantsBots.Encode()
	case *ChannelParticipantsFilter_ChannelParticipantsBanned:
		m2, _ := m.Payload.(*ChannelParticipantsFilter_ChannelParticipantsBanned)
		b = m2.ChannelParticipantsBanned.Encode()
	case *ChannelParticipantsFilter_ChannelParticipantsSearch:
		m2, _ := m.Payload.(*ChannelParticipantsFilter_ChannelParticipantsSearch)
		b = m2.ChannelParticipantsSearch.Encode()
	}
	return
}

func (m *ChannelParticipantsFilter) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelParticipantsRecent):
		m2 := ChannelParticipantsFilter_ChannelParticipantsRecent{}
		m2.ChannelParticipantsRecent.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantsAdmins):
		m2 := ChannelParticipantsFilter_ChannelParticipantsAdmins{}
		m2.ChannelParticipantsAdmins.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantsKicked):
		m2 := ChannelParticipantsFilter_ChannelParticipantsKicked{}
		m2.ChannelParticipantsKicked.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantsBots):
		m2 := ChannelParticipantsFilter_ChannelParticipantsBots{}
		m2.ChannelParticipantsBots.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantsBanned):
		m2 := ChannelParticipantsFilter_ChannelParticipantsBanned{}
		m2.ChannelParticipantsBanned.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelParticipantsSearch):
		m2 := ChannelParticipantsFilter_ChannelParticipantsSearch{}
		m2.ChannelParticipantsSearch.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Channels_ChannelParticipants) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Channels_ChannelParticipants_ChannelsChannelParticipants:
		m2, _ := m.Payload.(*Channels_ChannelParticipants_ChannelsChannelParticipants)
		b = m2.ChannelsChannelParticipants.Encode()
	}
	return
}

func (m *Channels_ChannelParticipants) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channels_channelParticipants):
		m2 := Channels_ChannelParticipants_ChannelsChannelParticipants{}
		m2.ChannelsChannelParticipants.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Channels_ChannelParticipant) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Channels_ChannelParticipant_ChannelsChannelParticipant:
		m2, _ := m.Payload.(*Channels_ChannelParticipant_ChannelsChannelParticipant)
		b = m2.ChannelsChannelParticipant.Encode()
	}
	return
}

func (m *Channels_ChannelParticipant) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channels_channelParticipant):
		m2 := Channels_ChannelParticipant_ChannelsChannelParticipant{}
		m2.ChannelsChannelParticipant.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Help_TermsOfService) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Help_TermsOfService_HelpTermsOfService:
		m2, _ := m.Payload.(*Help_TermsOfService_HelpTermsOfService)
		b = m2.HelpTermsOfService.Encode()
	}
	return
}

func (m *Help_TermsOfService) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_help_termsOfService):
		m2 := Help_TermsOfService_HelpTermsOfService{}
		m2.HelpTermsOfService.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *FoundGif) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *FoundGif_FoundGif:
		m2, _ := m.Payload.(*FoundGif_FoundGif)
		b = m2.FoundGif.Encode()
	case *FoundGif_FoundGifCached:
		m2, _ := m.Payload.(*FoundGif_FoundGifCached)
		b = m2.FoundGifCached.Encode()
	}
	return
}

func (m *FoundGif) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_foundGif):
		m2 := FoundGif_FoundGif{}
		m2.FoundGif.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_foundGifCached):
		m2 := FoundGif_FoundGifCached{}
		m2.FoundGifCached.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_FoundGifs) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_FoundGifs_MessagesFoundGifs:
		m2, _ := m.Payload.(*Messages_FoundGifs_MessagesFoundGifs)
		b = m2.MessagesFoundGifs.Encode()
	}
	return
}

func (m *Messages_FoundGifs) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_foundGifs):
		m2 := Messages_FoundGifs_MessagesFoundGifs{}
		m2.MessagesFoundGifs.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_SavedGifs) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_SavedGifs_MessagesSavedGifsNotModified:
		m2, _ := m.Payload.(*Messages_SavedGifs_MessagesSavedGifsNotModified)
		b = m2.MessagesSavedGifsNotModified.Encode()
	case *Messages_SavedGifs_MessagesSavedGifs:
		m2, _ := m.Payload.(*Messages_SavedGifs_MessagesSavedGifs)
		b = m2.MessagesSavedGifs.Encode()
	}
	return
}

func (m *Messages_SavedGifs) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_savedGifsNotModified):
		m2 := Messages_SavedGifs_MessagesSavedGifsNotModified{}
		m2.MessagesSavedGifsNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_savedGifs):
		m2 := Messages_SavedGifs_MessagesSavedGifs{}
		m2.MessagesSavedGifs.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputBotInlineMessage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputBotInlineMessage_InputBotInlineMessageMediaAuto:
		m2, _ := m.Payload.(*InputBotInlineMessage_InputBotInlineMessageMediaAuto)
		b = m2.InputBotInlineMessageMediaAuto.Encode()
	case *InputBotInlineMessage_InputBotInlineMessageText:
		m2, _ := m.Payload.(*InputBotInlineMessage_InputBotInlineMessageText)
		b = m2.InputBotInlineMessageText.Encode()
	case *InputBotInlineMessage_InputBotInlineMessageMediaGeo:
		m2, _ := m.Payload.(*InputBotInlineMessage_InputBotInlineMessageMediaGeo)
		b = m2.InputBotInlineMessageMediaGeo.Encode()
	case *InputBotInlineMessage_InputBotInlineMessageMediaVenue:
		m2, _ := m.Payload.(*InputBotInlineMessage_InputBotInlineMessageMediaVenue)
		b = m2.InputBotInlineMessageMediaVenue.Encode()
	case *InputBotInlineMessage_InputBotInlineMessageMediaContact:
		m2, _ := m.Payload.(*InputBotInlineMessage_InputBotInlineMessageMediaContact)
		b = m2.InputBotInlineMessageMediaContact.Encode()
	case *InputBotInlineMessage_InputBotInlineMessageGame:
		m2, _ := m.Payload.(*InputBotInlineMessage_InputBotInlineMessageGame)
		b = m2.InputBotInlineMessageGame.Encode()
	}
	return
}

func (m *InputBotInlineMessage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputBotInlineMessageMediaAuto):
		m2 := InputBotInlineMessage_InputBotInlineMessageMediaAuto{}
		m2.InputBotInlineMessageMediaAuto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineMessageText):
		m2 := InputBotInlineMessage_InputBotInlineMessageText{}
		m2.InputBotInlineMessageText.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineMessageMediaGeo):
		m2 := InputBotInlineMessage_InputBotInlineMessageMediaGeo{}
		m2.InputBotInlineMessageMediaGeo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineMessageMediaVenue):
		m2 := InputBotInlineMessage_InputBotInlineMessageMediaVenue{}
		m2.InputBotInlineMessageMediaVenue.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineMessageMediaContact):
		m2 := InputBotInlineMessage_InputBotInlineMessageMediaContact{}
		m2.InputBotInlineMessageMediaContact.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineMessageGame):
		m2 := InputBotInlineMessage_InputBotInlineMessageGame{}
		m2.InputBotInlineMessageGame.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputBotInlineResult) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputBotInlineResult_InputBotInlineResult:
		m2, _ := m.Payload.(*InputBotInlineResult_InputBotInlineResult)
		b = m2.InputBotInlineResult.Encode()
	case *InputBotInlineResult_InputBotInlineResultPhoto:
		m2, _ := m.Payload.(*InputBotInlineResult_InputBotInlineResultPhoto)
		b = m2.InputBotInlineResultPhoto.Encode()
	case *InputBotInlineResult_InputBotInlineResultDocument:
		m2, _ := m.Payload.(*InputBotInlineResult_InputBotInlineResultDocument)
		b = m2.InputBotInlineResultDocument.Encode()
	case *InputBotInlineResult_InputBotInlineResultGame:
		m2, _ := m.Payload.(*InputBotInlineResult_InputBotInlineResultGame)
		b = m2.InputBotInlineResultGame.Encode()
	}
	return
}

func (m *InputBotInlineResult) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputBotInlineResult):
		m2 := InputBotInlineResult_InputBotInlineResult{}
		m2.InputBotInlineResult.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineResultPhoto):
		m2 := InputBotInlineResult_InputBotInlineResultPhoto{}
		m2.InputBotInlineResultPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineResultDocument):
		m2 := InputBotInlineResult_InputBotInlineResultDocument{}
		m2.InputBotInlineResultDocument.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputBotInlineResultGame):
		m2 := InputBotInlineResult_InputBotInlineResultGame{}
		m2.InputBotInlineResultGame.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *BotInlineMessage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *BotInlineMessage_BotInlineMessageMediaAuto:
		m2, _ := m.Payload.(*BotInlineMessage_BotInlineMessageMediaAuto)
		b = m2.BotInlineMessageMediaAuto.Encode()
	case *BotInlineMessage_BotInlineMessageText:
		m2, _ := m.Payload.(*BotInlineMessage_BotInlineMessageText)
		b = m2.BotInlineMessageText.Encode()
	case *BotInlineMessage_BotInlineMessageMediaGeo:
		m2, _ := m.Payload.(*BotInlineMessage_BotInlineMessageMediaGeo)
		b = m2.BotInlineMessageMediaGeo.Encode()
	case *BotInlineMessage_BotInlineMessageMediaVenue:
		m2, _ := m.Payload.(*BotInlineMessage_BotInlineMessageMediaVenue)
		b = m2.BotInlineMessageMediaVenue.Encode()
	case *BotInlineMessage_BotInlineMessageMediaContact:
		m2, _ := m.Payload.(*BotInlineMessage_BotInlineMessageMediaContact)
		b = m2.BotInlineMessageMediaContact.Encode()
	}
	return
}

func (m *BotInlineMessage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_botInlineMessageMediaAuto):
		m2 := BotInlineMessage_BotInlineMessageMediaAuto{}
		m2.BotInlineMessageMediaAuto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_botInlineMessageText):
		m2 := BotInlineMessage_BotInlineMessageText{}
		m2.BotInlineMessageText.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_botInlineMessageMediaGeo):
		m2 := BotInlineMessage_BotInlineMessageMediaGeo{}
		m2.BotInlineMessageMediaGeo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_botInlineMessageMediaVenue):
		m2 := BotInlineMessage_BotInlineMessageMediaVenue{}
		m2.BotInlineMessageMediaVenue.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_botInlineMessageMediaContact):
		m2 := BotInlineMessage_BotInlineMessageMediaContact{}
		m2.BotInlineMessageMediaContact.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *BotInlineResult) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *BotInlineResult_BotInlineResult:
		m2, _ := m.Payload.(*BotInlineResult_BotInlineResult)
		b = m2.BotInlineResult.Encode()
	case *BotInlineResult_BotInlineMediaResult:
		m2, _ := m.Payload.(*BotInlineResult_BotInlineMediaResult)
		b = m2.BotInlineMediaResult.Encode()
	}
	return
}

func (m *BotInlineResult) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_botInlineResult):
		m2 := BotInlineResult_BotInlineResult{}
		m2.BotInlineResult.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_botInlineMediaResult):
		m2 := BotInlineResult_BotInlineMediaResult{}
		m2.BotInlineMediaResult.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_BotResults) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_BotResults_MessagesBotResults:
		m2, _ := m.Payload.(*Messages_BotResults_MessagesBotResults)
		b = m2.MessagesBotResults.Encode()
	}
	return
}

func (m *Messages_BotResults) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_botResults):
		m2 := Messages_BotResults_MessagesBotResults{}
		m2.MessagesBotResults.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ExportedMessageLink) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ExportedMessageLink_ExportedMessageLink:
		m2, _ := m.Payload.(*ExportedMessageLink_ExportedMessageLink)
		b = m2.ExportedMessageLink.Encode()
	}
	return
}

func (m *ExportedMessageLink) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_exportedMessageLink):
		m2 := ExportedMessageLink_ExportedMessageLink{}
		m2.ExportedMessageLink.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MessageFwdHeader) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MessageFwdHeader_MessageFwdHeader:
		m2, _ := m.Payload.(*MessageFwdHeader_MessageFwdHeader)
		b = m2.MessageFwdHeader.Encode()
	}
	return
}

func (m *MessageFwdHeader) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messageFwdHeader):
		m2 := MessageFwdHeader_MessageFwdHeader{}
		m2.MessageFwdHeader.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_CodeType) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_CodeType_AuthCodeTypeSms:
		m2, _ := m.Payload.(*Auth_CodeType_AuthCodeTypeSms)
		b = m2.AuthCodeTypeSms.Encode()
	case *Auth_CodeType_AuthCodeTypeCall:
		m2, _ := m.Payload.(*Auth_CodeType_AuthCodeTypeCall)
		b = m2.AuthCodeTypeCall.Encode()
	case *Auth_CodeType_AuthCodeTypeFlashCall:
		m2, _ := m.Payload.(*Auth_CodeType_AuthCodeTypeFlashCall)
		b = m2.AuthCodeTypeFlashCall.Encode()
	}
	return
}

func (m *Auth_CodeType) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_codeTypeSms):
		m2 := Auth_CodeType_AuthCodeTypeSms{}
		m2.AuthCodeTypeSms.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_auth_codeTypeCall):
		m2 := Auth_CodeType_AuthCodeTypeCall{}
		m2.AuthCodeTypeCall.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_auth_codeTypeFlashCall):
		m2 := Auth_CodeType_AuthCodeTypeFlashCall{}
		m2.AuthCodeTypeFlashCall.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Auth_SentCodeType) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Auth_SentCodeType_AuthSentCodeTypeApp:
		m2, _ := m.Payload.(*Auth_SentCodeType_AuthSentCodeTypeApp)
		b = m2.AuthSentCodeTypeApp.Encode()
	case *Auth_SentCodeType_AuthSentCodeTypeSms:
		m2, _ := m.Payload.(*Auth_SentCodeType_AuthSentCodeTypeSms)
		b = m2.AuthSentCodeTypeSms.Encode()
	case *Auth_SentCodeType_AuthSentCodeTypeCall:
		m2, _ := m.Payload.(*Auth_SentCodeType_AuthSentCodeTypeCall)
		b = m2.AuthSentCodeTypeCall.Encode()
	case *Auth_SentCodeType_AuthSentCodeTypeFlashCall:
		m2, _ := m.Payload.(*Auth_SentCodeType_AuthSentCodeTypeFlashCall)
		b = m2.AuthSentCodeTypeFlashCall.Encode()
	}
	return
}

func (m *Auth_SentCodeType) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_auth_sentCodeTypeApp):
		m2 := Auth_SentCodeType_AuthSentCodeTypeApp{}
		m2.AuthSentCodeTypeApp.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_auth_sentCodeTypeSms):
		m2 := Auth_SentCodeType_AuthSentCodeTypeSms{}
		m2.AuthSentCodeTypeSms.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_auth_sentCodeTypeCall):
		m2 := Auth_SentCodeType_AuthSentCodeTypeCall{}
		m2.AuthSentCodeTypeCall.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_auth_sentCodeTypeFlashCall):
		m2 := Auth_SentCodeType_AuthSentCodeTypeFlashCall{}
		m2.AuthSentCodeTypeFlashCall.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_BotCallbackAnswer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_BotCallbackAnswer_MessagesBotCallbackAnswer:
		m2, _ := m.Payload.(*Messages_BotCallbackAnswer_MessagesBotCallbackAnswer)
		b = m2.MessagesBotCallbackAnswer.Encode()
	}
	return
}

func (m *Messages_BotCallbackAnswer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_botCallbackAnswer):
		m2 := Messages_BotCallbackAnswer_MessagesBotCallbackAnswer{}
		m2.MessagesBotCallbackAnswer.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_MessageEditData) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_MessageEditData_MessagesMessageEditData:
		m2, _ := m.Payload.(*Messages_MessageEditData_MessagesMessageEditData)
		b = m2.MessagesMessageEditData.Encode()
	}
	return
}

func (m *Messages_MessageEditData) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_messageEditData):
		m2 := Messages_MessageEditData_MessagesMessageEditData{}
		m2.MessagesMessageEditData.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputBotInlineMessageID) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputBotInlineMessageID_InputBotInlineMessageID:
		m2, _ := m.Payload.(*InputBotInlineMessageID_InputBotInlineMessageID)
		b = m2.InputBotInlineMessageID.Encode()
	}
	return
}

func (m *InputBotInlineMessageID) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputBotInlineMessageID):
		m2 := InputBotInlineMessageID_InputBotInlineMessageID{}
		m2.InputBotInlineMessageID.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InlineBotSwitchPM) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InlineBotSwitchPM_InlineBotSwitchPM:
		m2, _ := m.Payload.(*InlineBotSwitchPM_InlineBotSwitchPM)
		b = m2.InlineBotSwitchPM.Encode()
	}
	return
}

func (m *InlineBotSwitchPM) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inlineBotSwitchPM):
		m2 := InlineBotSwitchPM_InlineBotSwitchPM{}
		m2.InlineBotSwitchPM.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_PeerDialogs) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_PeerDialogs_MessagesPeerDialogs:
		m2, _ := m.Payload.(*Messages_PeerDialogs_MessagesPeerDialogs)
		b = m2.MessagesPeerDialogs.Encode()
	}
	return
}

func (m *Messages_PeerDialogs) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_peerDialogs):
		m2 := Messages_PeerDialogs_MessagesPeerDialogs{}
		m2.MessagesPeerDialogs.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *TopPeer) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *TopPeer_TopPeer:
		m2, _ := m.Payload.(*TopPeer_TopPeer)
		b = m2.TopPeer.Encode()
	}
	return
}

func (m *TopPeer) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_topPeer):
		m2 := TopPeer_TopPeer{}
		m2.TopPeer.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *TopPeerCategory) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *TopPeerCategory_TopPeerCategoryBotsPM:
		m2, _ := m.Payload.(*TopPeerCategory_TopPeerCategoryBotsPM)
		b = m2.TopPeerCategoryBotsPM.Encode()
	case *TopPeerCategory_TopPeerCategoryBotsInline:
		m2, _ := m.Payload.(*TopPeerCategory_TopPeerCategoryBotsInline)
		b = m2.TopPeerCategoryBotsInline.Encode()
	case *TopPeerCategory_TopPeerCategoryCorrespondents:
		m2, _ := m.Payload.(*TopPeerCategory_TopPeerCategoryCorrespondents)
		b = m2.TopPeerCategoryCorrespondents.Encode()
	case *TopPeerCategory_TopPeerCategoryGroups:
		m2, _ := m.Payload.(*TopPeerCategory_TopPeerCategoryGroups)
		b = m2.TopPeerCategoryGroups.Encode()
	case *TopPeerCategory_TopPeerCategoryChannels:
		m2, _ := m.Payload.(*TopPeerCategory_TopPeerCategoryChannels)
		b = m2.TopPeerCategoryChannels.Encode()
	case *TopPeerCategory_TopPeerCategoryPhoneCalls:
		m2, _ := m.Payload.(*TopPeerCategory_TopPeerCategoryPhoneCalls)
		b = m2.TopPeerCategoryPhoneCalls.Encode()
	}
	return
}

func (m *TopPeerCategory) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_topPeerCategoryBotsPM):
		m2 := TopPeerCategory_TopPeerCategoryBotsPM{}
		m2.TopPeerCategoryBotsPM.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_topPeerCategoryBotsInline):
		m2 := TopPeerCategory_TopPeerCategoryBotsInline{}
		m2.TopPeerCategoryBotsInline.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_topPeerCategoryCorrespondents):
		m2 := TopPeerCategory_TopPeerCategoryCorrespondents{}
		m2.TopPeerCategoryCorrespondents.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_topPeerCategoryGroups):
		m2 := TopPeerCategory_TopPeerCategoryGroups{}
		m2.TopPeerCategoryGroups.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_topPeerCategoryChannels):
		m2 := TopPeerCategory_TopPeerCategoryChannels{}
		m2.TopPeerCategoryChannels.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_topPeerCategoryPhoneCalls):
		m2 := TopPeerCategory_TopPeerCategoryPhoneCalls{}
		m2.TopPeerCategoryPhoneCalls.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *TopPeerCategoryPeers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *TopPeerCategoryPeers_TopPeerCategoryPeers:
		m2, _ := m.Payload.(*TopPeerCategoryPeers_TopPeerCategoryPeers)
		b = m2.TopPeerCategoryPeers.Encode()
	}
	return
}

func (m *TopPeerCategoryPeers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_topPeerCategoryPeers):
		m2 := TopPeerCategoryPeers_TopPeerCategoryPeers{}
		m2.TopPeerCategoryPeers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Contacts_TopPeers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Contacts_TopPeers_ContactsTopPeersNotModified:
		m2, _ := m.Payload.(*Contacts_TopPeers_ContactsTopPeersNotModified)
		b = m2.ContactsTopPeersNotModified.Encode()
	case *Contacts_TopPeers_ContactsTopPeers:
		m2, _ := m.Payload.(*Contacts_TopPeers_ContactsTopPeers)
		b = m2.ContactsTopPeers.Encode()
	}
	return
}

func (m *Contacts_TopPeers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_contacts_topPeersNotModified):
		m2 := Contacts_TopPeers_ContactsTopPeersNotModified{}
		m2.ContactsTopPeersNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_contacts_topPeers):
		m2 := Contacts_TopPeers_ContactsTopPeers{}
		m2.ContactsTopPeers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DraftMessage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DraftMessage_DraftMessageEmpty:
		m2, _ := m.Payload.(*DraftMessage_DraftMessageEmpty)
		b = m2.DraftMessageEmpty.Encode()
	case *DraftMessage_DraftMessage:
		m2, _ := m.Payload.(*DraftMessage_DraftMessage)
		b = m2.DraftMessage.Encode()
	}
	return
}

func (m *DraftMessage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_draftMessageEmpty):
		m2 := DraftMessage_DraftMessageEmpty{}
		m2.DraftMessageEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_draftMessage):
		m2 := DraftMessage_DraftMessage{}
		m2.DraftMessage.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_FeaturedStickers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_FeaturedStickers_MessagesFeaturedStickersNotModified:
		m2, _ := m.Payload.(*Messages_FeaturedStickers_MessagesFeaturedStickersNotModified)
		b = m2.MessagesFeaturedStickersNotModified.Encode()
	case *Messages_FeaturedStickers_MessagesFeaturedStickers:
		m2, _ := m.Payload.(*Messages_FeaturedStickers_MessagesFeaturedStickers)
		b = m2.MessagesFeaturedStickers.Encode()
	}
	return
}

func (m *Messages_FeaturedStickers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_featuredStickersNotModified):
		m2 := Messages_FeaturedStickers_MessagesFeaturedStickersNotModified{}
		m2.MessagesFeaturedStickersNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_featuredStickers):
		m2 := Messages_FeaturedStickers_MessagesFeaturedStickers{}
		m2.MessagesFeaturedStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_RecentStickers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_RecentStickers_MessagesRecentStickersNotModified:
		m2, _ := m.Payload.(*Messages_RecentStickers_MessagesRecentStickersNotModified)
		b = m2.MessagesRecentStickersNotModified.Encode()
	case *Messages_RecentStickers_MessagesRecentStickers:
		m2, _ := m.Payload.(*Messages_RecentStickers_MessagesRecentStickers)
		b = m2.MessagesRecentStickers.Encode()
	}
	return
}

func (m *Messages_RecentStickers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_recentStickersNotModified):
		m2 := Messages_RecentStickers_MessagesRecentStickersNotModified{}
		m2.MessagesRecentStickersNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_recentStickers):
		m2 := Messages_RecentStickers_MessagesRecentStickers{}
		m2.MessagesRecentStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_ArchivedStickers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_ArchivedStickers_MessagesArchivedStickers:
		m2, _ := m.Payload.(*Messages_ArchivedStickers_MessagesArchivedStickers)
		b = m2.MessagesArchivedStickers.Encode()
	}
	return
}

func (m *Messages_ArchivedStickers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_archivedStickers):
		m2 := Messages_ArchivedStickers_MessagesArchivedStickers{}
		m2.MessagesArchivedStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_StickerSetInstallResult) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_StickerSetInstallResult_MessagesStickerSetInstallResultSuccess:
		m2, _ := m.Payload.(*Messages_StickerSetInstallResult_MessagesStickerSetInstallResultSuccess)
		b = m2.MessagesStickerSetInstallResultSuccess.Encode()
	case *Messages_StickerSetInstallResult_MessagesStickerSetInstallResultArchive:
		m2, _ := m.Payload.(*Messages_StickerSetInstallResult_MessagesStickerSetInstallResultArchive)
		b = m2.MessagesStickerSetInstallResultArchive.Encode()
	}
	return
}

func (m *Messages_StickerSetInstallResult) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_stickerSetInstallResultSuccess):
		m2 := Messages_StickerSetInstallResult_MessagesStickerSetInstallResultSuccess{}
		m2.MessagesStickerSetInstallResultSuccess.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_stickerSetInstallResultArchive):
		m2 := Messages_StickerSetInstallResult_MessagesStickerSetInstallResultArchive{}
		m2.MessagesStickerSetInstallResultArchive.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *StickerSetCovered) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *StickerSetCovered_StickerSetCovered:
		m2, _ := m.Payload.(*StickerSetCovered_StickerSetCovered)
		b = m2.StickerSetCovered.Encode()
	case *StickerSetCovered_StickerSetMultiCovered:
		m2, _ := m.Payload.(*StickerSetCovered_StickerSetMultiCovered)
		b = m2.StickerSetMultiCovered.Encode()
	}
	return
}

func (m *StickerSetCovered) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_stickerSetCovered):
		m2 := StickerSetCovered_StickerSetCovered{}
		m2.StickerSetCovered.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_stickerSetMultiCovered):
		m2 := StickerSetCovered_StickerSetMultiCovered{}
		m2.StickerSetMultiCovered.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *MaskCoords) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *MaskCoords_MaskCoords:
		m2, _ := m.Payload.(*MaskCoords_MaskCoords)
		b = m2.MaskCoords.Encode()
	}
	return
}

func (m *MaskCoords) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_maskCoords):
		m2 := MaskCoords_MaskCoords{}
		m2.MaskCoords.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputStickeredMedia) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputStickeredMedia_InputStickeredMediaPhoto:
		m2, _ := m.Payload.(*InputStickeredMedia_InputStickeredMediaPhoto)
		b = m2.InputStickeredMediaPhoto.Encode()
	case *InputStickeredMedia_InputStickeredMediaDocument:
		m2, _ := m.Payload.(*InputStickeredMedia_InputStickeredMediaDocument)
		b = m2.InputStickeredMediaDocument.Encode()
	}
	return
}

func (m *InputStickeredMedia) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputStickeredMediaPhoto):
		m2 := InputStickeredMedia_InputStickeredMediaPhoto{}
		m2.InputStickeredMediaPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputStickeredMediaDocument):
		m2 := InputStickeredMedia_InputStickeredMediaDocument{}
		m2.InputStickeredMediaDocument.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Game) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Game_Game:
		m2, _ := m.Payload.(*Game_Game)
		b = m2.Game.Encode()
	}
	return
}

func (m *Game) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_game):
		m2 := Game_Game{}
		m2.Game.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputGame) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputGame_InputGameID:
		m2, _ := m.Payload.(*InputGame_InputGameID)
		b = m2.InputGameID.Encode()
	case *InputGame_InputGameShortName:
		m2, _ := m.Payload.(*InputGame_InputGameShortName)
		b = m2.InputGameShortName.Encode()
	}
	return
}

func (m *InputGame) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputGameID):
		m2 := InputGame_InputGameID{}
		m2.InputGameID.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputGameShortName):
		m2 := InputGame_InputGameShortName{}
		m2.InputGameShortName.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *HighScore) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *HighScore_HighScore:
		m2, _ := m.Payload.(*HighScore_HighScore)
		b = m2.HighScore.Encode()
	}
	return
}

func (m *HighScore) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_highScore):
		m2 := HighScore_HighScore{}
		m2.HighScore.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_HighScores) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_HighScores_MessagesHighScores:
		m2, _ := m.Payload.(*Messages_HighScores_MessagesHighScores)
		b = m2.MessagesHighScores.Encode()
	}
	return
}

func (m *Messages_HighScores) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_highScores):
		m2 := Messages_HighScores_MessagesHighScores{}
		m2.MessagesHighScores.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *RichText) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *RichText_TextEmpty:
		m2, _ := m.Payload.(*RichText_TextEmpty)
		b = m2.TextEmpty.Encode()
	case *RichText_TextPlain:
		m2, _ := m.Payload.(*RichText_TextPlain)
		b = m2.TextPlain.Encode()
	case *RichText_TextBold:
		m2, _ := m.Payload.(*RichText_TextBold)
		b = m2.TextBold.Encode()
	case *RichText_TextItalic:
		m2, _ := m.Payload.(*RichText_TextItalic)
		b = m2.TextItalic.Encode()
	case *RichText_TextUnderline:
		m2, _ := m.Payload.(*RichText_TextUnderline)
		b = m2.TextUnderline.Encode()
	case *RichText_TextStrike:
		m2, _ := m.Payload.(*RichText_TextStrike)
		b = m2.TextStrike.Encode()
	case *RichText_TextFixed:
		m2, _ := m.Payload.(*RichText_TextFixed)
		b = m2.TextFixed.Encode()
	case *RichText_TextUrl:
		m2, _ := m.Payload.(*RichText_TextUrl)
		b = m2.TextUrl.Encode()
	case *RichText_TextEmail:
		m2, _ := m.Payload.(*RichText_TextEmail)
		b = m2.TextEmail.Encode()
	case *RichText_TextConcat:
		m2, _ := m.Payload.(*RichText_TextConcat)
		b = m2.TextConcat.Encode()
	}
	return
}

func (m *RichText) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_textEmpty):
		m2 := RichText_TextEmpty{}
		m2.TextEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textPlain):
		m2 := RichText_TextPlain{}
		m2.TextPlain.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textBold):
		m2 := RichText_TextBold{}
		m2.TextBold.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textItalic):
		m2 := RichText_TextItalic{}
		m2.TextItalic.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textUnderline):
		m2 := RichText_TextUnderline{}
		m2.TextUnderline.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textStrike):
		m2 := RichText_TextStrike{}
		m2.TextStrike.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textFixed):
		m2 := RichText_TextFixed{}
		m2.TextFixed.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textUrl):
		m2 := RichText_TextUrl{}
		m2.TextUrl.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textEmail):
		m2 := RichText_TextEmail{}
		m2.TextEmail.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_textConcat):
		m2 := RichText_TextConcat{}
		m2.TextConcat.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PageBlock) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PageBlock_PageBlockUnsupported:
		m2, _ := m.Payload.(*PageBlock_PageBlockUnsupported)
		b = m2.PageBlockUnsupported.Encode()
	case *PageBlock_PageBlockTitle:
		m2, _ := m.Payload.(*PageBlock_PageBlockTitle)
		b = m2.PageBlockTitle.Encode()
	case *PageBlock_PageBlockSubtitle:
		m2, _ := m.Payload.(*PageBlock_PageBlockSubtitle)
		b = m2.PageBlockSubtitle.Encode()
	case *PageBlock_PageBlockAuthorDate:
		m2, _ := m.Payload.(*PageBlock_PageBlockAuthorDate)
		b = m2.PageBlockAuthorDate.Encode()
	case *PageBlock_PageBlockHeader:
		m2, _ := m.Payload.(*PageBlock_PageBlockHeader)
		b = m2.PageBlockHeader.Encode()
	case *PageBlock_PageBlockSubheader:
		m2, _ := m.Payload.(*PageBlock_PageBlockSubheader)
		b = m2.PageBlockSubheader.Encode()
	case *PageBlock_PageBlockParagraph:
		m2, _ := m.Payload.(*PageBlock_PageBlockParagraph)
		b = m2.PageBlockParagraph.Encode()
	case *PageBlock_PageBlockPreformatted:
		m2, _ := m.Payload.(*PageBlock_PageBlockPreformatted)
		b = m2.PageBlockPreformatted.Encode()
	case *PageBlock_PageBlockFooter:
		m2, _ := m.Payload.(*PageBlock_PageBlockFooter)
		b = m2.PageBlockFooter.Encode()
	case *PageBlock_PageBlockDivider:
		m2, _ := m.Payload.(*PageBlock_PageBlockDivider)
		b = m2.PageBlockDivider.Encode()
	case *PageBlock_PageBlockAnchor:
		m2, _ := m.Payload.(*PageBlock_PageBlockAnchor)
		b = m2.PageBlockAnchor.Encode()
	case *PageBlock_PageBlockList:
		m2, _ := m.Payload.(*PageBlock_PageBlockList)
		b = m2.PageBlockList.Encode()
	case *PageBlock_PageBlockBlockquote:
		m2, _ := m.Payload.(*PageBlock_PageBlockBlockquote)
		b = m2.PageBlockBlockquote.Encode()
	case *PageBlock_PageBlockPullquote:
		m2, _ := m.Payload.(*PageBlock_PageBlockPullquote)
		b = m2.PageBlockPullquote.Encode()
	case *PageBlock_PageBlockPhoto:
		m2, _ := m.Payload.(*PageBlock_PageBlockPhoto)
		b = m2.PageBlockPhoto.Encode()
	case *PageBlock_PageBlockVideo:
		m2, _ := m.Payload.(*PageBlock_PageBlockVideo)
		b = m2.PageBlockVideo.Encode()
	case *PageBlock_PageBlockCover:
		m2, _ := m.Payload.(*PageBlock_PageBlockCover)
		b = m2.PageBlockCover.Encode()
	case *PageBlock_PageBlockEmbed:
		m2, _ := m.Payload.(*PageBlock_PageBlockEmbed)
		b = m2.PageBlockEmbed.Encode()
	case *PageBlock_PageBlockEmbedPost:
		m2, _ := m.Payload.(*PageBlock_PageBlockEmbedPost)
		b = m2.PageBlockEmbedPost.Encode()
	case *PageBlock_PageBlockCollage:
		m2, _ := m.Payload.(*PageBlock_PageBlockCollage)
		b = m2.PageBlockCollage.Encode()
	case *PageBlock_PageBlockSlideshow:
		m2, _ := m.Payload.(*PageBlock_PageBlockSlideshow)
		b = m2.PageBlockSlideshow.Encode()
	case *PageBlock_PageBlockChannel:
		m2, _ := m.Payload.(*PageBlock_PageBlockChannel)
		b = m2.PageBlockChannel.Encode()
	case *PageBlock_PageBlockAudio:
		m2, _ := m.Payload.(*PageBlock_PageBlockAudio)
		b = m2.PageBlockAudio.Encode()
	}
	return
}

func (m *PageBlock) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_pageBlockUnsupported):
		m2 := PageBlock_PageBlockUnsupported{}
		m2.PageBlockUnsupported.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockTitle):
		m2 := PageBlock_PageBlockTitle{}
		m2.PageBlockTitle.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockSubtitle):
		m2 := PageBlock_PageBlockSubtitle{}
		m2.PageBlockSubtitle.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockAuthorDate):
		m2 := PageBlock_PageBlockAuthorDate{}
		m2.PageBlockAuthorDate.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockHeader):
		m2 := PageBlock_PageBlockHeader{}
		m2.PageBlockHeader.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockSubheader):
		m2 := PageBlock_PageBlockSubheader{}
		m2.PageBlockSubheader.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockParagraph):
		m2 := PageBlock_PageBlockParagraph{}
		m2.PageBlockParagraph.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockPreformatted):
		m2 := PageBlock_PageBlockPreformatted{}
		m2.PageBlockPreformatted.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockFooter):
		m2 := PageBlock_PageBlockFooter{}
		m2.PageBlockFooter.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockDivider):
		m2 := PageBlock_PageBlockDivider{}
		m2.PageBlockDivider.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockAnchor):
		m2 := PageBlock_PageBlockAnchor{}
		m2.PageBlockAnchor.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockList):
		m2 := PageBlock_PageBlockList{}
		m2.PageBlockList.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockBlockquote):
		m2 := PageBlock_PageBlockBlockquote{}
		m2.PageBlockBlockquote.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockPullquote):
		m2 := PageBlock_PageBlockPullquote{}
		m2.PageBlockPullquote.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockPhoto):
		m2 := PageBlock_PageBlockPhoto{}
		m2.PageBlockPhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockVideo):
		m2 := PageBlock_PageBlockVideo{}
		m2.PageBlockVideo.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockCover):
		m2 := PageBlock_PageBlockCover{}
		m2.PageBlockCover.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockEmbed):
		m2 := PageBlock_PageBlockEmbed{}
		m2.PageBlockEmbed.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockEmbedPost):
		m2 := PageBlock_PageBlockEmbedPost{}
		m2.PageBlockEmbedPost.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockCollage):
		m2 := PageBlock_PageBlockCollage{}
		m2.PageBlockCollage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockSlideshow):
		m2 := PageBlock_PageBlockSlideshow{}
		m2.PageBlockSlideshow.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockChannel):
		m2 := PageBlock_PageBlockChannel{}
		m2.PageBlockChannel.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageBlockAudio):
		m2 := PageBlock_PageBlockAudio{}
		m2.PageBlockAudio.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Page) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Page_PagePart:
		m2, _ := m.Payload.(*Page_PagePart)
		b = m2.PagePart.Encode()
	case *Page_PageFull:
		m2, _ := m.Payload.(*Page_PageFull)
		b = m2.PageFull.Encode()
	}
	return
}

func (m *Page) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_pagePart):
		m2 := Page_PagePart{}
		m2.PagePart.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_pageFull):
		m2 := Page_PageFull{}
		m2.PageFull.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PhoneCallDiscardReason) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PhoneCallDiscardReason_PhoneCallDiscardReasonMissed:
		m2, _ := m.Payload.(*PhoneCallDiscardReason_PhoneCallDiscardReasonMissed)
		b = m2.PhoneCallDiscardReasonMissed.Encode()
	case *PhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect:
		m2, _ := m.Payload.(*PhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect)
		b = m2.PhoneCallDiscardReasonDisconnect.Encode()
	case *PhoneCallDiscardReason_PhoneCallDiscardReasonHangup:
		m2, _ := m.Payload.(*PhoneCallDiscardReason_PhoneCallDiscardReasonHangup)
		b = m2.PhoneCallDiscardReasonHangup.Encode()
	case *PhoneCallDiscardReason_PhoneCallDiscardReasonBusy:
		m2, _ := m.Payload.(*PhoneCallDiscardReason_PhoneCallDiscardReasonBusy)
		b = m2.PhoneCallDiscardReasonBusy.Encode()
	}
	return
}

func (m *PhoneCallDiscardReason) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_phoneCallDiscardReasonMissed):
		m2 := PhoneCallDiscardReason_PhoneCallDiscardReasonMissed{}
		m2.PhoneCallDiscardReasonMissed.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallDiscardReasonDisconnect):
		m2 := PhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect{}
		m2.PhoneCallDiscardReasonDisconnect.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallDiscardReasonHangup):
		m2 := PhoneCallDiscardReason_PhoneCallDiscardReasonHangup{}
		m2.PhoneCallDiscardReasonHangup.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallDiscardReasonBusy):
		m2 := PhoneCallDiscardReason_PhoneCallDiscardReasonBusy{}
		m2.PhoneCallDiscardReasonBusy.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *DataJSON) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *DataJSON_DataJSON:
		m2, _ := m.Payload.(*DataJSON_DataJSON)
		b = m2.DataJSON.Encode()
	}
	return
}

func (m *DataJSON) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_dataJSON):
		m2 := DataJSON_DataJSON{}
		m2.DataJSON.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *LabeledPrice) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *LabeledPrice_LabeledPrice:
		m2, _ := m.Payload.(*LabeledPrice_LabeledPrice)
		b = m2.LabeledPrice.Encode()
	}
	return
}

func (m *LabeledPrice) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_labeledPrice):
		m2 := LabeledPrice_LabeledPrice{}
		m2.LabeledPrice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Invoice) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Invoice_Invoice:
		m2, _ := m.Payload.(*Invoice_Invoice)
		b = m2.Invoice.Encode()
	}
	return
}

func (m *Invoice) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_invoice):
		m2 := Invoice_Invoice{}
		m2.Invoice.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PaymentCharge) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PaymentCharge_PaymentCharge:
		m2, _ := m.Payload.(*PaymentCharge_PaymentCharge)
		b = m2.PaymentCharge.Encode()
	}
	return
}

func (m *PaymentCharge) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_paymentCharge):
		m2 := PaymentCharge_PaymentCharge{}
		m2.PaymentCharge.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PostAddress) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PostAddress_PostAddress:
		m2, _ := m.Payload.(*PostAddress_PostAddress)
		b = m2.PostAddress.Encode()
	}
	return
}

func (m *PostAddress) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_postAddress):
		m2 := PostAddress_PostAddress{}
		m2.PostAddress.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PaymentRequestedInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PaymentRequestedInfo_PaymentRequestedInfo:
		m2, _ := m.Payload.(*PaymentRequestedInfo_PaymentRequestedInfo)
		b = m2.PaymentRequestedInfo.Encode()
	}
	return
}

func (m *PaymentRequestedInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_paymentRequestedInfo):
		m2 := PaymentRequestedInfo_PaymentRequestedInfo{}
		m2.PaymentRequestedInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PaymentSavedCredentials) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PaymentSavedCredentials_PaymentSavedCredentialsCard:
		m2, _ := m.Payload.(*PaymentSavedCredentials_PaymentSavedCredentialsCard)
		b = m2.PaymentSavedCredentialsCard.Encode()
	}
	return
}

func (m *PaymentSavedCredentials) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_paymentSavedCredentialsCard):
		m2 := PaymentSavedCredentials_PaymentSavedCredentialsCard{}
		m2.PaymentSavedCredentialsCard.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *WebDocument) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *WebDocument_WebDocument:
		m2, _ := m.Payload.(*WebDocument_WebDocument)
		b = m2.WebDocument.Encode()
	}
	return
}

func (m *WebDocument) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_webDocument):
		m2 := WebDocument_WebDocument{}
		m2.WebDocument.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputWebDocument) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputWebDocument_InputWebDocument:
		m2, _ := m.Payload.(*InputWebDocument_InputWebDocument)
		b = m2.InputWebDocument.Encode()
	}
	return
}

func (m *InputWebDocument) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputWebDocument):
		m2 := InputWebDocument_InputWebDocument{}
		m2.InputWebDocument.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputWebFileLocation) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputWebFileLocation_InputWebFileLocation:
		m2, _ := m.Payload.(*InputWebFileLocation_InputWebFileLocation)
		b = m2.InputWebFileLocation.Encode()
	}
	return
}

func (m *InputWebFileLocation) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputWebFileLocation):
		m2 := InputWebFileLocation_InputWebFileLocation{}
		m2.InputWebFileLocation.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Upload_WebFile) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Upload_WebFile_UploadWebFile:
		m2, _ := m.Payload.(*Upload_WebFile_UploadWebFile)
		b = m2.UploadWebFile.Encode()
	}
	return
}

func (m *Upload_WebFile) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_upload_webFile):
		m2 := Upload_WebFile_UploadWebFile{}
		m2.UploadWebFile.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Payments_PaymentForm) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Payments_PaymentForm_PaymentsPaymentForm:
		m2, _ := m.Payload.(*Payments_PaymentForm_PaymentsPaymentForm)
		b = m2.PaymentsPaymentForm.Encode()
	}
	return
}

func (m *Payments_PaymentForm) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_payments_paymentForm):
		m2 := Payments_PaymentForm_PaymentsPaymentForm{}
		m2.PaymentsPaymentForm.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Payments_ValidatedRequestedInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Payments_ValidatedRequestedInfo_PaymentsValidatedRequestedInfo:
		m2, _ := m.Payload.(*Payments_ValidatedRequestedInfo_PaymentsValidatedRequestedInfo)
		b = m2.PaymentsValidatedRequestedInfo.Encode()
	}
	return
}

func (m *Payments_ValidatedRequestedInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_payments_validatedRequestedInfo):
		m2 := Payments_ValidatedRequestedInfo_PaymentsValidatedRequestedInfo{}
		m2.PaymentsValidatedRequestedInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Payments_PaymentResult) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Payments_PaymentResult_PaymentsPaymentResult:
		m2, _ := m.Payload.(*Payments_PaymentResult_PaymentsPaymentResult)
		b = m2.PaymentsPaymentResult.Encode()
	case *Payments_PaymentResult_PaymentsPaymentVerficationNeeded:
		m2, _ := m.Payload.(*Payments_PaymentResult_PaymentsPaymentVerficationNeeded)
		b = m2.PaymentsPaymentVerficationNeeded.Encode()
	}
	return
}

func (m *Payments_PaymentResult) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_payments_paymentResult):
		m2 := Payments_PaymentResult_PaymentsPaymentResult{}
		m2.PaymentsPaymentResult.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_payments_paymentVerficationNeeded):
		m2 := Payments_PaymentResult_PaymentsPaymentVerficationNeeded{}
		m2.PaymentsPaymentVerficationNeeded.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Payments_PaymentReceipt) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Payments_PaymentReceipt_PaymentsPaymentReceipt:
		m2, _ := m.Payload.(*Payments_PaymentReceipt_PaymentsPaymentReceipt)
		b = m2.PaymentsPaymentReceipt.Encode()
	}
	return
}

func (m *Payments_PaymentReceipt) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_payments_paymentReceipt):
		m2 := Payments_PaymentReceipt_PaymentsPaymentReceipt{}
		m2.PaymentsPaymentReceipt.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Payments_SavedInfo) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Payments_SavedInfo_PaymentsSavedInfo:
		m2, _ := m.Payload.(*Payments_SavedInfo_PaymentsSavedInfo)
		b = m2.PaymentsSavedInfo.Encode()
	}
	return
}

func (m *Payments_SavedInfo) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_payments_savedInfo):
		m2 := Payments_SavedInfo_PaymentsSavedInfo{}
		m2.PaymentsSavedInfo.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPaymentCredentials) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPaymentCredentials_InputPaymentCredentialsSaved:
		m2, _ := m.Payload.(*InputPaymentCredentials_InputPaymentCredentialsSaved)
		b = m2.InputPaymentCredentialsSaved.Encode()
	case *InputPaymentCredentials_InputPaymentCredentials:
		m2, _ := m.Payload.(*InputPaymentCredentials_InputPaymentCredentials)
		b = m2.InputPaymentCredentials.Encode()
	}
	return
}

func (m *InputPaymentCredentials) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPaymentCredentialsSaved):
		m2 := InputPaymentCredentials_InputPaymentCredentialsSaved{}
		m2.InputPaymentCredentialsSaved.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_inputPaymentCredentials):
		m2 := InputPaymentCredentials_InputPaymentCredentials{}
		m2.InputPaymentCredentials.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Account_TmpPassword) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Account_TmpPassword_AccountTmpPassword:
		m2, _ := m.Payload.(*Account_TmpPassword_AccountTmpPassword)
		b = m2.AccountTmpPassword.Encode()
	}
	return
}

func (m *Account_TmpPassword) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_account_tmpPassword):
		m2 := Account_TmpPassword_AccountTmpPassword{}
		m2.AccountTmpPassword.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ShippingOption) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ShippingOption_ShippingOption:
		m2, _ := m.Payload.(*ShippingOption_ShippingOption)
		b = m2.ShippingOption.Encode()
	}
	return
}

func (m *ShippingOption) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_shippingOption):
		m2 := ShippingOption_ShippingOption{}
		m2.ShippingOption.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputStickerSetItem) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputStickerSetItem_InputStickerSetItem:
		m2, _ := m.Payload.(*InputStickerSetItem_InputStickerSetItem)
		b = m2.InputStickerSetItem.Encode()
	}
	return
}

func (m *InputStickerSetItem) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputStickerSetItem):
		m2 := InputStickerSetItem_InputStickerSetItem{}
		m2.InputStickerSetItem.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *InputPhoneCall) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *InputPhoneCall_InputPhoneCall:
		m2, _ := m.Payload.(*InputPhoneCall_InputPhoneCall)
		b = m2.InputPhoneCall.Encode()
	}
	return
}

func (m *InputPhoneCall) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_inputPhoneCall):
		m2 := InputPhoneCall_InputPhoneCall{}
		m2.InputPhoneCall.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PhoneCall) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PhoneCall_PhoneCallEmpty:
		m2, _ := m.Payload.(*PhoneCall_PhoneCallEmpty)
		b = m2.PhoneCallEmpty.Encode()
	case *PhoneCall_PhoneCallWaiting:
		m2, _ := m.Payload.(*PhoneCall_PhoneCallWaiting)
		b = m2.PhoneCallWaiting.Encode()
	case *PhoneCall_PhoneCallRequested:
		m2, _ := m.Payload.(*PhoneCall_PhoneCallRequested)
		b = m2.PhoneCallRequested.Encode()
	case *PhoneCall_PhoneCallAccepted:
		m2, _ := m.Payload.(*PhoneCall_PhoneCallAccepted)
		b = m2.PhoneCallAccepted.Encode()
	case *PhoneCall_PhoneCall:
		m2, _ := m.Payload.(*PhoneCall_PhoneCall)
		b = m2.PhoneCall.Encode()
	case *PhoneCall_PhoneCallDiscarded:
		m2, _ := m.Payload.(*PhoneCall_PhoneCallDiscarded)
		b = m2.PhoneCallDiscarded.Encode()
	}
	return
}

func (m *PhoneCall) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_phoneCallEmpty):
		m2 := PhoneCall_PhoneCallEmpty{}
		m2.PhoneCallEmpty.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallWaiting):
		m2 := PhoneCall_PhoneCallWaiting{}
		m2.PhoneCallWaiting.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallRequested):
		m2 := PhoneCall_PhoneCallRequested{}
		m2.PhoneCallRequested.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallAccepted):
		m2 := PhoneCall_PhoneCallAccepted{}
		m2.PhoneCallAccepted.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCall):
		m2 := PhoneCall_PhoneCall{}
		m2.PhoneCall.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_phoneCallDiscarded):
		m2 := PhoneCall_PhoneCallDiscarded{}
		m2.PhoneCallDiscarded.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PhoneConnection) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PhoneConnection_PhoneConnection:
		m2, _ := m.Payload.(*PhoneConnection_PhoneConnection)
		b = m2.PhoneConnection.Encode()
	}
	return
}

func (m *PhoneConnection) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_phoneConnection):
		m2 := PhoneConnection_PhoneConnection{}
		m2.PhoneConnection.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PhoneCallProtocol) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PhoneCallProtocol_PhoneCallProtocol:
		m2, _ := m.Payload.(*PhoneCallProtocol_PhoneCallProtocol)
		b = m2.PhoneCallProtocol.Encode()
	}
	return
}

func (m *PhoneCallProtocol) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_phoneCallProtocol):
		m2 := PhoneCallProtocol_PhoneCallProtocol{}
		m2.PhoneCallProtocol.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Phone_PhoneCall) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Phone_PhoneCall_PhonePhoneCall:
		m2, _ := m.Payload.(*Phone_PhoneCall_PhonePhoneCall)
		b = m2.PhonePhoneCall.Encode()
	}
	return
}

func (m *Phone_PhoneCall) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_phone_phoneCall):
		m2 := Phone_PhoneCall_PhonePhoneCall{}
		m2.PhonePhoneCall.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Upload_CdnFile) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Upload_CdnFile_UploadCdnFileReuploadNeeded:
		m2, _ := m.Payload.(*Upload_CdnFile_UploadCdnFileReuploadNeeded)
		b = m2.UploadCdnFileReuploadNeeded.Encode()
	case *Upload_CdnFile_UploadCdnFile:
		m2, _ := m.Payload.(*Upload_CdnFile_UploadCdnFile)
		b = m2.UploadCdnFile.Encode()
	}
	return
}

func (m *Upload_CdnFile) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_upload_cdnFileReuploadNeeded):
		m2 := Upload_CdnFile_UploadCdnFileReuploadNeeded{}
		m2.UploadCdnFileReuploadNeeded.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_upload_cdnFile):
		m2 := Upload_CdnFile_UploadCdnFile{}
		m2.UploadCdnFile.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *CdnPublicKey) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *CdnPublicKey_CdnPublicKey:
		m2, _ := m.Payload.(*CdnPublicKey_CdnPublicKey)
		b = m2.CdnPublicKey.Encode()
	}
	return
}

func (m *CdnPublicKey) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_cdnPublicKey):
		m2 := CdnPublicKey_CdnPublicKey{}
		m2.CdnPublicKey.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *CdnConfig) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *CdnConfig_CdnConfig:
		m2, _ := m.Payload.(*CdnConfig_CdnConfig)
		b = m2.CdnConfig.Encode()
	}
	return
}

func (m *CdnConfig) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_cdnConfig):
		m2 := CdnConfig_CdnConfig{}
		m2.CdnConfig.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *LangPackString) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *LangPackString_LangPackString:
		m2, _ := m.Payload.(*LangPackString_LangPackString)
		b = m2.LangPackString.Encode()
	case *LangPackString_LangPackStringPluralized:
		m2, _ := m.Payload.(*LangPackString_LangPackStringPluralized)
		b = m2.LangPackStringPluralized.Encode()
	case *LangPackString_LangPackStringDeleted:
		m2, _ := m.Payload.(*LangPackString_LangPackStringDeleted)
		b = m2.LangPackStringDeleted.Encode()
	}
	return
}

func (m *LangPackString) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_langPackString):
		m2 := LangPackString_LangPackString{}
		m2.LangPackString.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_langPackStringPluralized):
		m2 := LangPackString_LangPackStringPluralized{}
		m2.LangPackStringPluralized.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_langPackStringDeleted):
		m2 := LangPackString_LangPackStringDeleted{}
		m2.LangPackStringDeleted.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *LangPackDifference) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *LangPackDifference_LangPackDifference:
		m2, _ := m.Payload.(*LangPackDifference_LangPackDifference)
		b = m2.LangPackDifference.Encode()
	}
	return
}

func (m *LangPackDifference) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_langPackDifference):
		m2 := LangPackDifference_LangPackDifference{}
		m2.LangPackDifference.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *LangPackLanguage) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *LangPackLanguage_LangPackLanguage:
		m2, _ := m.Payload.(*LangPackLanguage_LangPackLanguage)
		b = m2.LangPackLanguage.Encode()
	}
	return
}

func (m *LangPackLanguage) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_langPackLanguage):
		m2 := LangPackLanguage_LangPackLanguage{}
		m2.LangPackLanguage.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelAdminRights) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelAdminRights_ChannelAdminRights:
		m2, _ := m.Payload.(*ChannelAdminRights_ChannelAdminRights)
		b = m2.ChannelAdminRights.Encode()
	}
	return
}

func (m *ChannelAdminRights) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelAdminRights):
		m2 := ChannelAdminRights_ChannelAdminRights{}
		m2.ChannelAdminRights.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelBannedRights) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelBannedRights_ChannelBannedRights:
		m2, _ := m.Payload.(*ChannelBannedRights_ChannelBannedRights)
		b = m2.ChannelBannedRights.Encode()
	}
	return
}

func (m *ChannelBannedRights) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelBannedRights):
		m2 := ChannelBannedRights_ChannelBannedRights{}
		m2.ChannelBannedRights.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelAdminLogEventAction) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle)
		b = m2.ChannelAdminLogEventActionChangeTitle.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout)
		b = m2.ChannelAdminLogEventActionChangeAbout.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername)
		b = m2.ChannelAdminLogEventActionChangeUsername.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto)
		b = m2.ChannelAdminLogEventActionChangePhoto.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites)
		b = m2.ChannelAdminLogEventActionToggleInvites.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures)
		b = m2.ChannelAdminLogEventActionToggleSignatures.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned)
		b = m2.ChannelAdminLogEventActionUpdatePinned.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage)
		b = m2.ChannelAdminLogEventActionEditMessage.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage)
		b = m2.ChannelAdminLogEventActionDeleteMessage.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin)
		b = m2.ChannelAdminLogEventActionParticipantJoin.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave)
		b = m2.ChannelAdminLogEventActionParticipantLeave.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite)
		b = m2.ChannelAdminLogEventActionParticipantInvite.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan)
		b = m2.ChannelAdminLogEventActionParticipantToggleBan.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin)
		b = m2.ChannelAdminLogEventActionParticipantToggleAdmin.Encode()
	case *ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet:
		m2, _ := m.Payload.(*ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet)
		b = m2.ChannelAdminLogEventActionChangeStickerSet.Encode()
	}
	return
}

func (m *ChannelAdminLogEventAction) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelAdminLogEventActionChangeTitle):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle{}
		m2.ChannelAdminLogEventActionChangeTitle.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionChangeAbout):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout{}
		m2.ChannelAdminLogEventActionChangeAbout.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionChangeUsername):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername{}
		m2.ChannelAdminLogEventActionChangeUsername.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionChangePhoto):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto{}
		m2.ChannelAdminLogEventActionChangePhoto.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionToggleInvites):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites{}
		m2.ChannelAdminLogEventActionToggleInvites.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionToggleSignatures):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures{}
		m2.ChannelAdminLogEventActionToggleSignatures.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionUpdatePinned):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned{}
		m2.ChannelAdminLogEventActionUpdatePinned.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionEditMessage):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage{}
		m2.ChannelAdminLogEventActionEditMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionDeleteMessage):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage{}
		m2.ChannelAdminLogEventActionDeleteMessage.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantJoin):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin{}
		m2.ChannelAdminLogEventActionParticipantJoin.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantLeave):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave{}
		m2.ChannelAdminLogEventActionParticipantLeave.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantInvite):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite{}
		m2.ChannelAdminLogEventActionParticipantInvite.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleBan):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan{}
		m2.ChannelAdminLogEventActionParticipantToggleBan.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleAdmin):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin{}
		m2.ChannelAdminLogEventActionParticipantToggleAdmin.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_channelAdminLogEventActionChangeStickerSet):
		m2 := ChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet{}
		m2.ChannelAdminLogEventActionChangeStickerSet.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelAdminLogEvent) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelAdminLogEvent_ChannelAdminLogEvent:
		m2, _ := m.Payload.(*ChannelAdminLogEvent_ChannelAdminLogEvent)
		b = m2.ChannelAdminLogEvent.Encode()
	}
	return
}

func (m *ChannelAdminLogEvent) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelAdminLogEvent):
		m2 := ChannelAdminLogEvent_ChannelAdminLogEvent{}
		m2.ChannelAdminLogEvent.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Channels_AdminLogResults) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Channels_AdminLogResults_ChannelsAdminLogResults:
		m2, _ := m.Payload.(*Channels_AdminLogResults_ChannelsAdminLogResults)
		b = m2.ChannelsAdminLogResults.Encode()
	}
	return
}

func (m *Channels_AdminLogResults) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channels_adminLogResults):
		m2 := Channels_AdminLogResults_ChannelsAdminLogResults{}
		m2.ChannelsAdminLogResults.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *ChannelAdminLogEventsFilter) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *ChannelAdminLogEventsFilter_ChannelAdminLogEventsFilter:
		m2, _ := m.Payload.(*ChannelAdminLogEventsFilter_ChannelAdminLogEventsFilter)
		b = m2.ChannelAdminLogEventsFilter.Encode()
	}
	return
}

func (m *ChannelAdminLogEventsFilter) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_channelAdminLogEventsFilter):
		m2 := ChannelAdminLogEventsFilter_ChannelAdminLogEventsFilter{}
		m2.ChannelAdminLogEventsFilter.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *PopularContact) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *PopularContact_PopularContact:
		m2, _ := m.Payload.(*PopularContact_PopularContact)
		b = m2.PopularContact.Encode()
	}
	return
}

func (m *PopularContact) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_popularContact):
		m2 := PopularContact_PopularContact{}
		m2.PopularContact.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *CdnFileHash) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *CdnFileHash_CdnFileHash:
		m2, _ := m.Payload.(*CdnFileHash_CdnFileHash)
		b = m2.CdnFileHash.Encode()
	}
	return
}

func (m *CdnFileHash) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_cdnFileHash):
		m2 := CdnFileHash_CdnFileHash{}
		m2.CdnFileHash.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

func (m *Messages_FavedStickers) Encode() (b []byte) {
	b = nil
	switch m.Payload.(type) {
	case *Messages_FavedStickers_MessagesFavedStickersNotModified:
		m2, _ := m.Payload.(*Messages_FavedStickers_MessagesFavedStickersNotModified)
		b = m2.MessagesFavedStickersNotModified.Encode()
	case *Messages_FavedStickers_MessagesFavedStickers:
		m2, _ := m.Payload.(*Messages_FavedStickers_MessagesFavedStickers)
		b = m2.MessagesFavedStickers.Encode()
	}
	return
}

func (m *Messages_FavedStickers) Decode(dbuf *DecodeBuf) error {
	classId := dbuf.Int()
	switch classId {
	case int32(TLConstructor_CRC32_messages_favedStickersNotModified):
		m2 := Messages_FavedStickers_MessagesFavedStickersNotModified{}
		m2.MessagesFavedStickersNotModified.Decode(dbuf)
		m.Payload = &m2
	case int32(TLConstructor_CRC32_messages_favedStickers):
		m2 := Messages_FavedStickers_MessagesFavedStickers{}
		m2.MessagesFavedStickers.Decode(dbuf)
		m.Payload = &m2
	}
	return dbuf.err
}

// resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
func (m *TLResPQ) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_resPQ))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.String(m.Pq)
	x.VectorLong(m.ServerPublicKeyFingerprints)
	return x.buf
}

func (m *TLResPQ) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.Pq = dbuf.String()
	m.ServerPublicKeyFingerprints = dbuf.VectorLong()
	return dbuf.err
}

// p_q_inner_data#83c95aec pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 = P_Q_inner_data;
func (m *TLPQInnerData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_p_q_inner_data))
	x.String(m.Pq)
	x.String(m.P)
	x.String(m.Q)
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Bytes(m.NewNonce)
	return x.buf
}

func (m *TLPQInnerData) Decode(dbuf *DecodeBuf) error {
	m.Pq = dbuf.String()
	m.P = dbuf.String()
	m.Q = dbuf.String()
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.NewNonce = dbuf.Bytes(32)
	return dbuf.err
}

// server_DH_params_fail#79cb045d nonce:int128 server_nonce:int128 new_nonce_hash:int128 = Server_DH_Params;
func (m *TLServer_DHParamsFail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_server_DH_params_fail))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Bytes(m.NewNonceHash)
	return x.buf
}

func (m *TLServer_DHParamsFail) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.NewNonceHash = dbuf.Bytes(16)
	return dbuf.err
}

// server_DH_params_ok#d0e8075c nonce:int128 server_nonce:int128 encrypted_answer:string = Server_DH_Params;
func (m *TLServer_DHParamsOk) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_server_DH_params_ok))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.String(m.EncryptedAnswer)
	return x.buf
}

func (m *TLServer_DHParamsOk) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.EncryptedAnswer = dbuf.String()
	return dbuf.err
}

// server_DH_inner_data#b5890dba nonce:int128 server_nonce:int128 g:int dh_prime:string g_a:string server_time:int = Server_DH_inner_data;
func (m *TLServer_DHInnerData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_server_DH_inner_data))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Int(m.G)
	x.String(m.DhPrime)
	x.String(m.GA)
	x.Int(m.ServerTime)
	return x.buf
}

func (m *TLServer_DHInnerData) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.G = dbuf.Int()
	m.DhPrime = dbuf.String()
	m.GA = dbuf.String()
	m.ServerTime = dbuf.Int()
	return dbuf.err
}

// client_DH_inner_data#6643b654 nonce:int128 server_nonce:int128 retry_id:long g_b:string = Client_DH_Inner_Data;
func (m *TLClient_DHInnerData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_client_DH_inner_data))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Long(m.RetryId)
	x.String(m.GB)
	return x.buf
}

func (m *TLClient_DHInnerData) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.RetryId = dbuf.Long()
	m.GB = dbuf.String()
	return dbuf.err
}

// dh_gen_ok#3bcbf734 nonce:int128 server_nonce:int128 new_nonce_hash1:int128 = Set_client_DH_params_answer;
func (m *TLDhGenOk) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dh_gen_ok))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Bytes(m.NewNonceHash1)
	return x.buf
}

func (m *TLDhGenOk) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.NewNonceHash1 = dbuf.Bytes(16)
	return dbuf.err
}

// dh_gen_retry#46dc1fb9 nonce:int128 server_nonce:int128 new_nonce_hash2:int128 = Set_client_DH_params_answer;
func (m *TLDhGenRetry) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dh_gen_retry))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Bytes(m.NewNonceHash2)
	return x.buf
}

func (m *TLDhGenRetry) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.NewNonceHash2 = dbuf.Bytes(16)
	return dbuf.err
}

// dh_gen_fail#a69dae02 nonce:int128 server_nonce:int128 new_nonce_hash3:int128 = Set_client_DH_params_answer;
func (m *TLDhGenFail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dh_gen_fail))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.Bytes(m.NewNonceHash3)
	return x.buf
}

func (m *TLDhGenFail) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.NewNonceHash3 = dbuf.Bytes(16)
	return dbuf.err
}

// destroy_auth_key_ok#f660e1d4 = DestroyAuthKeyRes;
func (m *TLDestroyAuthKeyOk) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_auth_key_ok))
	return x.buf
}

func (m *TLDestroyAuthKeyOk) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// destroy_auth_key_none#0a9f2259 = DestroyAuthKeyRes;
func (m *TLDestroyAuthKeyNone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_auth_key_none))
	return x.buf
}

func (m *TLDestroyAuthKeyNone) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// destroy_auth_key_fail#ea109b13 = DestroyAuthKeyRes;
func (m *TLDestroyAuthKeyFail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_auth_key_fail))
	return x.buf
}

func (m *TLDestroyAuthKeyFail) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// msgs_ack#62d6b459 msg_ids:Vector<long> = MsgsAck;
func (m *TLMsgsAck) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msgs_ack))
	x.VectorLong(m.MsgIds)
	return x.buf
}

func (m *TLMsgsAck) Decode(dbuf *DecodeBuf) error {
	m.MsgIds = dbuf.VectorLong()
	return dbuf.err
}

// bad_msg_notification#a7eff811 bad_msg_id:long bad_msg_seqno:int error_code:int = BadMsgNotification;
func (m *TLBadMsgNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_bad_msg_notification))
	x.Long(m.BadMsgId)
	x.Int(m.BadMsgSeqno)
	x.Int(m.ErrorCode)
	return x.buf
}

func (m *TLBadMsgNotification) Decode(dbuf *DecodeBuf) error {
	m.BadMsgId = dbuf.Long()
	m.BadMsgSeqno = dbuf.Int()
	m.ErrorCode = dbuf.Int()
	return dbuf.err
}

// bad_server_salt#edab447b bad_msg_id:long bad_msg_seqno:int error_code:int new_server_salt:long = BadMsgNotification;
func (m *TLBadServerSalt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_bad_server_salt))
	x.Long(m.BadMsgId)
	x.Int(m.BadMsgSeqno)
	x.Int(m.ErrorCode)
	x.Long(m.NewServerSalt)
	return x.buf
}

func (m *TLBadServerSalt) Decode(dbuf *DecodeBuf) error {
	m.BadMsgId = dbuf.Long()
	m.BadMsgSeqno = dbuf.Int()
	m.ErrorCode = dbuf.Int()
	m.NewServerSalt = dbuf.Long()
	return dbuf.err
}

// msgs_state_req#da69fb52 msg_ids:Vector<long> = MsgsStateReq;
func (m *TLMsgsStateReq) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msgs_state_req))
	x.VectorLong(m.MsgIds)
	return x.buf
}

func (m *TLMsgsStateReq) Decode(dbuf *DecodeBuf) error {
	m.MsgIds = dbuf.VectorLong()
	return dbuf.err
}

// msgs_state_info#04deb57d req_msg_id:long info:string = MsgsStateInfo;
func (m *TLMsgsStateInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msgs_state_info))
	x.Long(m.ReqMsgId)
	x.String(m.Info)
	return x.buf
}

func (m *TLMsgsStateInfo) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	m.Info = dbuf.String()
	return dbuf.err
}

// msgs_all_info#8cc0d131 msg_ids:Vector<long> info:string = MsgsAllInfo;
func (m *TLMsgsAllInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msgs_all_info))
	x.VectorLong(m.MsgIds)
	x.String(m.Info)
	return x.buf
}

func (m *TLMsgsAllInfo) Decode(dbuf *DecodeBuf) error {
	m.MsgIds = dbuf.VectorLong()
	m.Info = dbuf.String()
	return dbuf.err
}

// msg_detailed_info#276d3ec6 msg_id:long answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
func (m *TLMsgDetailedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_detailed_info))
	x.Long(m.MsgId)
	x.Long(m.AnswerMsgId)
	x.Int(m.Bytes)
	x.Int(m.Status)
	return x.buf
}

func (m *TLMsgDetailedInfo) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.AnswerMsgId = dbuf.Long()
	m.Bytes = dbuf.Int()
	m.Status = dbuf.Int()
	return dbuf.err
}

// msg_new_detailed_info#809db6df answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
func (m *TLMsgNewDetailedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_new_detailed_info))
	x.Long(m.AnswerMsgId)
	x.Int(m.Bytes)
	x.Int(m.Status)
	return x.buf
}

func (m *TLMsgNewDetailedInfo) Decode(dbuf *DecodeBuf) error {
	m.AnswerMsgId = dbuf.Long()
	m.Bytes = dbuf.Int()
	m.Status = dbuf.Int()
	return dbuf.err
}

// msg_resend_req#7d861a08 msg_ids:Vector<long> = MsgResendReq;
func (m *TLMsgResendReq) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_resend_req))
	x.VectorLong(m.MsgIds)
	return x.buf
}

func (m *TLMsgResendReq) Decode(dbuf *DecodeBuf) error {
	m.MsgIds = dbuf.VectorLong()
	return dbuf.err
}

// rpc_error#2144ca19 error_code:int error_message:string = RpcError;
func (m *TLRpcError) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_error))
	x.Int(m.ErrorCode)
	x.String(m.ErrorMessage)
	return x.buf
}

func (m *TLRpcError) Decode(dbuf *DecodeBuf) error {
	m.ErrorCode = dbuf.Int()
	m.ErrorMessage = dbuf.String()
	return dbuf.err
}

// rpc_answer_unknown#5e2ad36e = RpcDropAnswer;
func (m *TLRpcAnswerUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_answer_unknown))
	return x.buf
}

func (m *TLRpcAnswerUnknown) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// rpc_answer_dropped_running#cd78e586 = RpcDropAnswer;
func (m *TLRpcAnswerDroppedRunning) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_answer_dropped_running))
	return x.buf
}

func (m *TLRpcAnswerDroppedRunning) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// rpc_answer_dropped#a43ad8b7 msg_id:long seq_no:int bytes:int = RpcDropAnswer;
func (m *TLRpcAnswerDropped) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_answer_dropped))
	x.Long(m.MsgId)
	x.Int(m.SeqNo)
	x.Int(m.Bytes)
	return x.buf
}

func (m *TLRpcAnswerDropped) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.SeqNo = dbuf.Int()
	m.Bytes = dbuf.Int()
	return dbuf.err
}

// future_salt#0949d9dc valid_since:int valid_until:int salt:long = FutureSalt;
func (m *TLFutureSalt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_future_salt))
	x.Int(m.ValidSince)
	x.Int(m.ValidUntil)
	x.Long(m.Salt)
	return x.buf
}

func (m *TLFutureSalt) Decode(dbuf *DecodeBuf) error {
	m.ValidSince = dbuf.Int()
	m.ValidUntil = dbuf.Int()
	m.Salt = dbuf.Long()
	return dbuf.err
}

// future_salts#ae500895 req_msg_id:long now:int salts:vector<future_salt> = FutureSalts;
func (m *TLFutureSalts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_future_salts))
	x.Long(m.ReqMsgId)
	x.Int(m.Now)
	// x.VectorMessage(m.Salts);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Salts)))
	for _, v := range m.Salts {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLFutureSalts) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	m.Now = dbuf.Int()
	// x.VectorMessage(m.Salts);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Salts = make([]*TLFutureSalt, l3)
	for i := 0; i < int(l3); i++ {
		m.Salts[i] = &TLFutureSalt{}
		(*m.Salts[i]).Decode(dbuf)
		// TODO(@benqi): Check classID valid!!!
		dbuf.Int()
	}
	return dbuf.err
}

// pong#347773c5 msg_id:long ping_id:long = Pong;
func (m *TLPong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pong))
	x.Long(m.MsgId)
	x.Long(m.PingId)
	return x.buf
}

func (m *TLPong) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.PingId = dbuf.Long()
	return dbuf.err
}

// destroy_session_ok#e22045fc session_id:long = DestroySessionRes;
func (m *TLDestroySessionOk) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_session_ok))
	x.Long(m.SessionId)
	return x.buf
}

func (m *TLDestroySessionOk) Decode(dbuf *DecodeBuf) error {
	m.SessionId = dbuf.Long()
	return dbuf.err
}

// destroy_session_none#62d350c9 session_id:long = DestroySessionRes;
func (m *TLDestroySessionNone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_session_none))
	x.Long(m.SessionId)
	return x.buf
}

func (m *TLDestroySessionNone) Decode(dbuf *DecodeBuf) error {
	m.SessionId = dbuf.Long()
	return dbuf.err
}

// new_session_created#9ec20908 first_msg_id:long unique_id:long server_salt:long = NewSession;
func (m *TLNewSessionCreated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_new_session_created))
	x.Long(m.FirstMsgId)
	x.Long(m.UniqueId)
	x.Long(m.ServerSalt)
	return x.buf
}

func (m *TLNewSessionCreated) Decode(dbuf *DecodeBuf) error {
	m.FirstMsgId = dbuf.Long()
	m.UniqueId = dbuf.Long()
	m.ServerSalt = dbuf.Long()
	return dbuf.err
}

// http_wait#9299359f max_delay:int wait_after:int max_wait:int = HttpWait;
func (m *TLHttpWait) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_http_wait))
	x.Int(m.MaxDelay)
	x.Int(m.WaitAfter)
	x.Int(m.MaxWait)
	return x.buf
}

func (m *TLHttpWait) Decode(dbuf *DecodeBuf) error {
	m.MaxDelay = dbuf.Int()
	m.WaitAfter = dbuf.Int()
	m.MaxWait = dbuf.Int()
	return dbuf.err
}

// ipPort#d433ad73 ipv4:int port:int = IpPort;
func (m *TLIpPort) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_ipPort))
	x.Int(m.Ipv4)
	x.Int(m.Port)
	return x.buf
}

func (m *TLIpPort) Decode(dbuf *DecodeBuf) error {
	m.Ipv4 = dbuf.Int()
	m.Port = dbuf.Int()
	return dbuf.err
}

// help.configSimple#d997c3c5 date:int expires:int dc_id:int ip_port_list:Vector<ipPort> = help.ConfigSimple;
func (m *TLHelpConfigSimple) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_configSimple))
	x.Int(m.Date)
	x.Int(m.Expires)
	x.Int(m.DcId)
	// x.VectorMessage(m.IpPortList);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.IpPortList)))
	for _, v := range m.IpPortList {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLHelpConfigSimple) Decode(dbuf *DecodeBuf) error {
	m.Date = dbuf.Int()
	m.Expires = dbuf.Int()
	m.DcId = dbuf.Int()
	// x.VectorMessage(m.IpPortList);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.IpPortList = make([]*TLIpPort, l4)
	for i := 0; i < int(l4); i++ {
		m.IpPortList[i] = &TLIpPort{}
		(*m.IpPortList[i]).Decode(dbuf)
		// TODO(@benqi): Check classID valid!!!
		dbuf.Int()
	}
	return dbuf.err
}

// boolFalse#bc799737 = Bool;
func (m *TLBoolFalse) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_boolFalse))
	return x.buf
}

func (m *TLBoolFalse) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// boolTrue#997275b5 = Bool;
func (m *TLBoolTrue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_boolTrue))
	return x.buf
}

func (m *TLBoolTrue) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// true#3fedd339 = True;
func (m *TLTrue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_true))
	return x.buf
}

func (m *TLTrue) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// error#c4b9f9bb code:int text:string = Error;
func (m *TLError) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_error))
	x.Int(m.Code)
	x.String(m.Text)
	return x.buf
}

func (m *TLError) Decode(dbuf *DecodeBuf) error {
	m.Code = dbuf.Int()
	m.Text = dbuf.String()
	return dbuf.err
}

// null#56730bcc = Null;
func (m *TLNull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_null))
	return x.buf
}

func (m *TLNull) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPeerEmpty#7f3b18ea = InputPeer;
func (m *TLInputPeerEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerEmpty))
	return x.buf
}

func (m *TLInputPeerEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPeerSelf#7da07ec9 = InputPeer;
func (m *TLInputPeerSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerSelf))
	return x.buf
}

func (m *TLInputPeerSelf) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPeerChat#179be863 chat_id:int = InputPeer;
func (m *TLInputPeerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerChat))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLInputPeerChat) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer;
func (m *TLInputPeerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerUser))
	x.Int(m.UserId)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputPeerUser) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer;
func (m *TLInputPeerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerChannel))
	x.Int(m.ChannelId)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputPeerChannel) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputUserEmpty#b98886cf = InputUser;
func (m *TLInputUserEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputUserEmpty))
	return x.buf
}

func (m *TLInputUserEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputUserSelf#f7c1b13f = InputUser;
func (m *TLInputUserSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputUserSelf))
	return x.buf
}

func (m *TLInputUserSelf) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputUser#d8292816 user_id:int access_hash:long = InputUser;
func (m *TLInputUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputUser))
	x.Int(m.UserId)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputUser) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;
func (m *TLInputPhoneContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhoneContact))
	x.Long(m.ClientId)
	x.String(m.Phone)
	x.String(m.FirstName)
	x.String(m.LastName)
	return x.buf
}

func (m *TLInputPhoneContact) Decode(dbuf *DecodeBuf) error {
	m.ClientId = dbuf.Long()
	m.Phone = dbuf.String()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	return dbuf.err
}

// inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
func (m *TLInputFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputFile))
	x.Long(m.Id)
	x.Int(m.Parts)
	x.String(m.Name)
	x.String(m.Md5Checksum)
	return x.buf
}

func (m *TLInputFile) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Parts = dbuf.Int()
	m.Name = dbuf.String()
	m.Md5Checksum = dbuf.String()
	return dbuf.err
}

// inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;
func (m *TLInputFileBig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputFileBig))
	x.Long(m.Id)
	x.Int(m.Parts)
	x.String(m.Name)
	return x.buf
}

func (m *TLInputFileBig) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Parts = dbuf.Int()
	m.Name = dbuf.String()
	return dbuf.err
}

// inputMediaEmpty#9664f57f = InputMedia;
func (m *TLInputMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaEmpty))
	return x.buf
}

func (m *TLInputMediaEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (m *TLInputMediaUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaUploadedPhoto))
	x.Int(m.Flags)
	x.Bytes(m.File.Encode())
	x.String(m.Caption)
	// x.VectorMessage(m.Stickers);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Stickers)))
	for _, v := range m.Stickers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLInputMediaUploadedPhoto) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.File = &InputFile{}
	m.Decode(dbuf)
	m.Caption = dbuf.String()
	// x.VectorMessage(m.Stickers);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Stickers = make([]*InputDocument, l4)
	for i := 0; i < int(l4); i++ {
		m.Stickers[i] = &InputDocument{}
		(*m.Stickers[i]).Decode(dbuf)
	}
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaPhoto))
	x.Int(m.Flags)
	x.Bytes(m.Id.Encode())
	x.String(m.Caption)
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLInputMediaPhoto) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = &InputPhoto{}
	m.Decode(dbuf)
	m.Caption = dbuf.String()
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
func (m *TLInputMediaGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaGeoPoint))
	x.Bytes(m.GeoPoint.Encode())
	return x.buf
}

func (m *TLInputMediaGeoPoint) Decode(dbuf *DecodeBuf) error {
	m.GeoPoint = &InputGeoPoint{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
func (m *TLInputMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaContact))
	x.String(m.PhoneNumber)
	x.String(m.FirstName)
	x.String(m.LastName)
	return x.buf
}

func (m *TLInputMediaContact) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	return dbuf.err
}

// inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (m *TLInputMediaUploadedDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaUploadedDocument))
	x.Int(m.Flags)
	x.Bytes(m.File.Encode())
	x.Bytes(m.Thumb.Encode())
	x.String(m.MimeType)
	// x.VectorMessage(m.Attributes);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Attributes)))
	for _, v := range m.Attributes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.String(m.Caption)
	// x.VectorMessage(m.Stickers);
	x7 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x7, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x7[4:], uint32(len(m.Stickers)))
	for _, v := range m.Stickers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLInputMediaUploadedDocument) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.File = &InputFile{}
	m.Decode(dbuf)
	m.Thumb = &InputFile{}
	m.Decode(dbuf)
	m.MimeType = dbuf.String()
	// x.VectorMessage(m.Attributes);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Attributes = make([]*DocumentAttribute, l5)
	for i := 0; i < int(l5); i++ {
		m.Attributes[i] = &DocumentAttribute{}
		(*m.Attributes[i]).Decode(dbuf)
	}
	m.Caption = dbuf.String()
	// x.VectorMessage(m.Stickers);
	c7 := dbuf.Int()
	if c7 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c7)
	}
	l7 := dbuf.Int()
	m.Stickers = make([]*InputDocument, l7)
	for i := 0; i < int(l7); i++ {
		m.Stickers[i] = &InputDocument{}
		(*m.Stickers[i]).Decode(dbuf)
	}
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaDocument))
	x.Int(m.Flags)
	x.Bytes(m.Id.Encode())
	x.String(m.Caption)
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLInputMediaDocument) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = &InputDocument{}
	m.Decode(dbuf)
	m.Caption = dbuf.String()
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;
func (m *TLInputMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaVenue))
	x.Bytes(m.GeoPoint.Encode())
	x.String(m.Title)
	x.String(m.Address)
	x.String(m.Provider)
	x.String(m.VenueId)
	return x.buf
}

func (m *TLInputMediaVenue) Decode(dbuf *DecodeBuf) error {
	m.GeoPoint = &InputGeoPoint{}
	m.Decode(dbuf)
	m.Title = dbuf.String()
	m.Address = dbuf.String()
	m.Provider = dbuf.String()
	m.VenueId = dbuf.String()
	return dbuf.err
}

// inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
func (m *TLInputMediaGifExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaGifExternal))
	x.String(m.Url)
	x.String(m.Q)
	return x.buf
}

func (m *TLInputMediaGifExternal) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.Q = dbuf.String()
	return dbuf.err
}

// inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaPhotoExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaPhotoExternal))
	x.Int(m.Flags)
	x.String(m.Url)
	x.String(m.Caption)
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLInputMediaPhotoExternal) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Url = dbuf.String()
	m.Caption = dbuf.String()
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaDocumentExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaDocumentExternal))
	x.Int(m.Flags)
	x.String(m.Url)
	x.String(m.Caption)
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLInputMediaDocumentExternal) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Url = dbuf.String()
	m.Caption = dbuf.String()
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// inputMediaGame#d33f43f3 id:InputGame = InputMedia;
func (m *TLInputMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaGame))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLInputMediaGame) Decode(dbuf *DecodeBuf) error {
	m.Id = &InputGame{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;
func (m *TLInputMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaInvoice))
	x.Int(m.Flags)
	x.String(m.Title)
	x.String(m.Description)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.Invoice.Encode())
	x.StringBytes(m.Payload)
	x.String(m.Provider)
	x.String(m.StartParam)
	return x.buf
}

func (m *TLInputMediaInvoice) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Photo = &InputWebDocument{}
	m.Decode(dbuf)
	m.Invoice = &Invoice{}
	m.Decode(dbuf)
	m.Payload = dbuf.StringBytes()
	m.Provider = dbuf.String()
	m.StartParam = dbuf.String()
	return dbuf.err
}

// inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;
func (m *TLInputChatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChatPhotoEmpty))
	return x.buf
}

func (m *TLInputChatPhotoEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputChatUploadedPhoto#927c55b4 file:InputFile = InputChatPhoto;
func (m *TLInputChatUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChatUploadedPhoto))
	x.Bytes(m.File.Encode())
	return x.buf
}

func (m *TLInputChatUploadedPhoto) Decode(dbuf *DecodeBuf) error {
	m.File = &InputFile{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputChatPhoto#8953ad37 id:InputPhoto = InputChatPhoto;
func (m *TLInputChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChatPhoto))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLInputChatPhoto) Decode(dbuf *DecodeBuf) error {
	m.Id = &InputPhoto{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
func (m *TLInputGeoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGeoPointEmpty))
	return x.buf
}

func (m *TLInputGeoPointEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
func (m *TLInputGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGeoPoint))
	x.Double(m.Lat)
	x.Double(m.Long)
	return x.buf
}

func (m *TLInputGeoPoint) Decode(dbuf *DecodeBuf) error {
	m.Lat = dbuf.Double()
	m.Long = dbuf.Double()
	return dbuf.err
}

// inputPhotoEmpty#1cd7bf0d = InputPhoto;
func (m *TLInputPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhotoEmpty))
	return x.buf
}

func (m *TLInputPhotoEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;
func (m *TLInputPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhoto))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputPhoto) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
func (m *TLInputFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputFileLocation))
	x.Long(m.VolumeId)
	x.Int(m.LocalId)
	x.Long(m.Secret)
	return x.buf
}

func (m *TLInputFileLocation) Decode(dbuf *DecodeBuf) error {
	m.VolumeId = dbuf.Long()
	m.LocalId = dbuf.Int()
	m.Secret = dbuf.Long()
	return dbuf.err
}

// inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
func (m *TLInputEncryptedFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileLocation))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputEncryptedFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
func (m *TLInputDocumentFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocumentFileLocation))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Version)
	return x.buf
}

func (m *TLInputDocumentFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Version = dbuf.Int()
	return dbuf.err
}

// inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;
func (m *TLInputAppEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputAppEvent))
	x.Double(m.Time)
	x.String(m.Type)
	x.Long(m.Peer)
	x.String(m.Data)
	return x.buf
}

func (m *TLInputAppEvent) Decode(dbuf *DecodeBuf) error {
	m.Time = dbuf.Double()
	m.Type = dbuf.String()
	m.Peer = dbuf.Long()
	m.Data = dbuf.String()
	return dbuf.err
}

// peerUser#9db1bc6d user_id:int = Peer;
func (m *TLPeerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerUser))
	x.Int(m.UserId)
	return x.buf
}

func (m *TLPeerUser) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	return dbuf.err
}

// peerChat#bad0e5bb chat_id:int = Peer;
func (m *TLPeerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerChat))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLPeerChat) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// peerChannel#bddde532 channel_id:int = Peer;
func (m *TLPeerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerChannel))
	x.Int(m.ChannelId)
	return x.buf
}

func (m *TLPeerChannel) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	return dbuf.err
}

// storage.fileUnknown#aa963b05 = storage.FileType;
func (m *TLStorageFileUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileUnknown))
	return x.buf
}

func (m *TLStorageFileUnknown) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.filePartial#40bc6f52 = storage.FileType;
func (m *TLStorageFilePartial) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_filePartial))
	return x.buf
}

func (m *TLStorageFilePartial) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.fileJpeg#7efe0e = storage.FileType;
func (m *TLStorageFileJpeg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileJpeg))
	return x.buf
}

func (m *TLStorageFileJpeg) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.fileGif#cae1aadf = storage.FileType;
func (m *TLStorageFileGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileGif))
	return x.buf
}

func (m *TLStorageFileGif) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.filePng#a4f63c0 = storage.FileType;
func (m *TLStorageFilePng) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_filePng))
	return x.buf
}

func (m *TLStorageFilePng) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.filePdf#ae1e508d = storage.FileType;
func (m *TLStorageFilePdf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_filePdf))
	return x.buf
}

func (m *TLStorageFilePdf) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.fileMp3#528a0677 = storage.FileType;
func (m *TLStorageFileMp3) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileMp3))
	return x.buf
}

func (m *TLStorageFileMp3) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.fileMov#4b09ebbc = storage.FileType;
func (m *TLStorageFileMov) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileMov))
	return x.buf
}

func (m *TLStorageFileMov) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.fileMp4#b3cea0e4 = storage.FileType;
func (m *TLStorageFileMp4) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileMp4))
	return x.buf
}

func (m *TLStorageFileMp4) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// storage.fileWebp#1081464c = storage.FileType;
func (m *TLStorageFileWebp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileWebp))
	return x.buf
}

func (m *TLStorageFileWebp) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;
func (m *TLFileLocationUnavailable) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_fileLocationUnavailable))
	x.Long(m.VolumeId)
	x.Int(m.LocalId)
	x.Long(m.Secret)
	return x.buf
}

func (m *TLFileLocationUnavailable) Decode(dbuf *DecodeBuf) error {
	m.VolumeId = dbuf.Long()
	m.LocalId = dbuf.Int()
	m.Secret = dbuf.Long()
	return dbuf.err
}

// fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
func (m *TLFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_fileLocation))
	x.Int(m.DcId)
	x.Long(m.VolumeId)
	x.Int(m.LocalId)
	x.Long(m.Secret)
	return x.buf
}

func (m *TLFileLocation) Decode(dbuf *DecodeBuf) error {
	m.DcId = dbuf.Int()
	m.VolumeId = dbuf.Long()
	m.LocalId = dbuf.Int()
	m.Secret = dbuf.Long()
	return dbuf.err
}

// userEmpty#200250ba id:int = User;
func (m *TLUserEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userEmpty))
	x.Int(m.Id)
	return x.buf
}

func (m *TLUserEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	return dbuf.err
}

// user#2e13f4c3 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
func (m *TLUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_user))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Long(m.AccessHash)
	x.String(m.FirstName)
	x.String(m.LastName)
	x.String(m.Username)
	x.String(m.Phone)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.Status.Encode())
	x.Int(m.BotInfoVersion)
	x.String(m.RestrictionReason)
	x.String(m.BotInlinePlaceholder)
	x.String(m.LangCode)
	return x.buf
}

func (m *TLUser) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.AccessHash = dbuf.Long()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	m.Username = dbuf.String()
	m.Phone = dbuf.String()
	m.Photo = &UserProfilePhoto{}
	m.Decode(dbuf)
	m.Status = &UserStatus{}
	m.Decode(dbuf)
	m.BotInfoVersion = dbuf.Int()
	m.RestrictionReason = dbuf.String()
	m.BotInlinePlaceholder = dbuf.String()
	m.LangCode = dbuf.String()
	return dbuf.err
}

// userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;
func (m *TLUserProfilePhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userProfilePhotoEmpty))
	return x.buf
}

func (m *TLUserProfilePhotoEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;
func (m *TLUserProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userProfilePhoto))
	x.Long(m.PhotoId)
	x.Bytes(m.PhotoSmall.Encode())
	x.Bytes(m.PhotoBig.Encode())
	return x.buf
}

func (m *TLUserProfilePhoto) Decode(dbuf *DecodeBuf) error {
	m.PhotoId = dbuf.Long()
	m.PhotoSmall = &FileLocation{}
	m.Decode(dbuf)
	m.PhotoBig = &FileLocation{}
	m.Decode(dbuf)
	return dbuf.err
}

// userStatusEmpty#9d05049 = UserStatus;
func (m *TLUserStatusEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusEmpty))
	return x.buf
}

func (m *TLUserStatusEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// userStatusOnline#edb93949 expires:int = UserStatus;
func (m *TLUserStatusOnline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusOnline))
	x.Int(m.Expires)
	return x.buf
}

func (m *TLUserStatusOnline) Decode(dbuf *DecodeBuf) error {
	m.Expires = dbuf.Int()
	return dbuf.err
}

// userStatusOffline#8c703f was_online:int = UserStatus;
func (m *TLUserStatusOffline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusOffline))
	x.Int(m.WasOnline)
	return x.buf
}

func (m *TLUserStatusOffline) Decode(dbuf *DecodeBuf) error {
	m.WasOnline = dbuf.Int()
	return dbuf.err
}

// userStatusRecently#e26f42f1 = UserStatus;
func (m *TLUserStatusRecently) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusRecently))
	return x.buf
}

func (m *TLUserStatusRecently) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// userStatusLastWeek#7bf09fc = UserStatus;
func (m *TLUserStatusLastWeek) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusLastWeek))
	return x.buf
}

func (m *TLUserStatusLastWeek) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// userStatusLastMonth#77ebc742 = UserStatus;
func (m *TLUserStatusLastMonth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusLastMonth))
	return x.buf
}

func (m *TLUserStatusLastMonth) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// chatEmpty#9ba2d800 id:int = Chat;
func (m *TLChatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatEmpty))
	x.Int(m.Id)
	return x.buf
}

func (m *TLChatEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	return dbuf.err
}

// chat#d91cdd54 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel = Chat;
func (m *TLChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chat))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.String(m.Title)
	x.Bytes(m.Photo.Encode())
	x.Int(m.ParticipantsCount)
	x.Int(m.Date)
	x.Int(m.Version)
	x.Bytes(m.MigratedTo.Encode())
	return x.buf
}

func (m *TLChat) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.Title = dbuf.String()
	m.Photo = &ChatPhoto{}
	m.Decode(dbuf)
	m.ParticipantsCount = dbuf.Int()
	m.Date = dbuf.Int()
	m.Version = dbuf.Int()
	m.MigratedTo = &InputChannel{}
	m.Decode(dbuf)
	return dbuf.err
}

// chatForbidden#7328bdb id:int title:string = Chat;
func (m *TLChatForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatForbidden))
	x.Int(m.Id)
	x.String(m.Title)
	return x.buf
}

func (m *TLChatForbidden) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Title = dbuf.String()
	return dbuf.err
}

// channel#cb44b1c flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights = Chat;
func (m *TLChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channel))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Long(m.AccessHash)
	x.String(m.Title)
	x.String(m.Username)
	x.Bytes(m.Photo.Encode())
	x.Int(m.Date)
	x.Int(m.Version)
	x.String(m.RestrictionReason)
	x.Bytes(m.AdminRights.Encode())
	x.Bytes(m.BannedRights.Encode())
	return x.buf
}

func (m *TLChannel) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.AccessHash = dbuf.Long()
	m.Title = dbuf.String()
	m.Username = dbuf.String()
	m.Photo = &ChatPhoto{}
	m.Decode(dbuf)
	m.Date = dbuf.Int()
	m.Version = dbuf.Int()
	m.RestrictionReason = dbuf.String()
	m.AdminRights = &ChannelAdminRights{}
	m.Decode(dbuf)
	m.BannedRights = &ChannelBannedRights{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelForbidden#289da732 flags:# broadcast:flags.5?true megagroup:flags.8?true id:int access_hash:long title:string until_date:flags.16?int = Chat;
func (m *TLChannelForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelForbidden))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Long(m.AccessHash)
	x.String(m.Title)
	x.Int(m.UntilDate)
	return x.buf
}

func (m *TLChannelForbidden) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.AccessHash = dbuf.Long()
	m.Title = dbuf.String()
	m.UntilDate = dbuf.Int()
	return dbuf.err
}

// chatFull#2e02a614 id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> = ChatFull;
func (m *TLChatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatFull))
	x.Int(m.Id)
	x.Bytes(m.Participants.Encode())
	x.Bytes(m.ChatPhoto.Encode())
	x.Bytes(m.NotifySettings.Encode())
	x.Bytes(m.ExportedInvite.Encode())
	// x.VectorMessage(m.BotInfo);
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.BotInfo)))
	for _, v := range m.BotInfo {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChatFull) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Participants = &ChatParticipants{}
	m.Decode(dbuf)
	m.ChatPhoto = &Photo{}
	m.Decode(dbuf)
	m.NotifySettings = &PeerNotifySettings{}
	m.Decode(dbuf)
	m.ExportedInvite = &ExportedChatInvite{}
	m.Decode(dbuf)
	// x.VectorMessage(m.BotInfo);
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.BotInfo = make([]*BotInfo, l6)
	for i := 0; i < int(l6); i++ {
		m.BotInfo[i] = &BotInfo{}
		(*m.BotInfo[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channelFull#17f45fcf flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet = ChatFull;
func (m *TLChannelFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelFull))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.String(m.About)
	x.Int(m.ParticipantsCount)
	x.Int(m.AdminsCount)
	x.Int(m.KickedCount)
	x.Int(m.BannedCount)
	x.Int(m.ReadInboxMaxId)
	x.Int(m.ReadOutboxMaxId)
	x.Int(m.UnreadCount)
	x.Bytes(m.ChatPhoto.Encode())
	x.Bytes(m.NotifySettings.Encode())
	x.Bytes(m.ExportedInvite.Encode())
	// x.VectorMessage(m.BotInfo);
	x17 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x17, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x17[4:], uint32(len(m.BotInfo)))
	for _, v := range m.BotInfo {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.MigratedFromChatId)
	x.Int(m.MigratedFromMaxId)
	x.Int(m.PinnedMsgId)
	x.Bytes(m.Stickerset.Encode())
	return x.buf
}

func (m *TLChannelFull) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.About = dbuf.String()
	m.ParticipantsCount = dbuf.Int()
	m.AdminsCount = dbuf.Int()
	m.KickedCount = dbuf.Int()
	m.BannedCount = dbuf.Int()
	m.ReadInboxMaxId = dbuf.Int()
	m.ReadOutboxMaxId = dbuf.Int()
	m.UnreadCount = dbuf.Int()
	m.ChatPhoto = &Photo{}
	m.Decode(dbuf)
	m.NotifySettings = &PeerNotifySettings{}
	m.Decode(dbuf)
	m.ExportedInvite = &ExportedChatInvite{}
	m.Decode(dbuf)
	// x.VectorMessage(m.BotInfo);
	c17 := dbuf.Int()
	if c17 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c17)
	}
	l17 := dbuf.Int()
	m.BotInfo = make([]*BotInfo, l17)
	for i := 0; i < int(l17); i++ {
		m.BotInfo[i] = &BotInfo{}
		(*m.BotInfo[i]).Decode(dbuf)
	}
	m.MigratedFromChatId = dbuf.Int()
	m.MigratedFromMaxId = dbuf.Int()
	m.PinnedMsgId = dbuf.Int()
	m.Stickerset = &StickerSet{}
	m.Decode(dbuf)
	return dbuf.err
}

// chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
func (m *TLChatParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipant))
	x.Int(m.UserId)
	x.Int(m.InviterId)
	x.Int(m.Date)
	return x.buf
}

func (m *TLChatParticipant) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.InviterId = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// chatParticipantCreator#da13538a user_id:int = ChatParticipant;
func (m *TLChatParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipantCreator))
	x.Int(m.UserId)
	return x.buf
}

func (m *TLChatParticipantCreator) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	return dbuf.err
}

// chatParticipantAdmin#e2d6e436 user_id:int inviter_id:int date:int = ChatParticipant;
func (m *TLChatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipantAdmin))
	x.Int(m.UserId)
	x.Int(m.InviterId)
	x.Int(m.Date)
	return x.buf
}

func (m *TLChatParticipantAdmin) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.InviterId = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// chatParticipantsForbidden#fc900c2b flags:# chat_id:int self_participant:flags.0?ChatParticipant = ChatParticipants;
func (m *TLChatParticipantsForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipantsForbidden))
	x.Int(m.Flags)
	x.Int(m.ChatId)
	x.Bytes(m.SelfParticipant.Encode())
	return x.buf
}

func (m *TLChatParticipantsForbidden) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.ChatId = dbuf.Int()
	m.SelfParticipant = &ChatParticipant{}
	m.Decode(dbuf)
	return dbuf.err
}

// chatParticipants#3f460fed chat_id:int participants:Vector<ChatParticipant> version:int = ChatParticipants;
func (m *TLChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipants))
	x.Int(m.ChatId)
	// x.VectorMessage(m.Participants);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Participants)))
	for _, v := range m.Participants {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.Version)
	return x.buf
}

func (m *TLChatParticipants) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	// x.VectorMessage(m.Participants);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Participants = make([]*ChatParticipant, l2)
	for i := 0; i < int(l2); i++ {
		m.Participants[i] = &ChatParticipant{}
		(*m.Participants[i]).Decode(dbuf)
	}
	m.Version = dbuf.Int()
	return dbuf.err
}

// chatPhotoEmpty#37c1011c = ChatPhoto;
func (m *TLChatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatPhotoEmpty))
	return x.buf
}

func (m *TLChatPhotoEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;
func (m *TLChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatPhoto))
	x.Bytes(m.PhotoSmall.Encode())
	x.Bytes(m.PhotoBig.Encode())
	return x.buf
}

func (m *TLChatPhoto) Decode(dbuf *DecodeBuf) error {
	m.PhotoSmall = &FileLocation{}
	m.Decode(dbuf)
	m.PhotoBig = &FileLocation{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageEmpty#83e5de54 id:int = Message;
func (m *TLMessageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEmpty))
	x.Int(m.Id)
	return x.buf
}

func (m *TLMessageEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	return dbuf.err
}

// message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
func (m *TLMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_message))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Int(m.FromId)
	x.Bytes(m.ToId.Encode())
	x.Bytes(m.FwdFrom.Encode())
	x.Int(m.ViaBotId)
	x.Int(m.ReplyToMsgId)
	x.Int(m.Date)
	x.String(m.Message)
	x.Bytes(m.Media.Encode())
	x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities);
	x17 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x17, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x17[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.Views)
	x.Int(m.EditDate)
	x.String(m.PostAuthor)
	return x.buf
}

func (m *TLMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.FromId = dbuf.Int()
	m.ToId = &Peer{}
	m.Decode(dbuf)
	m.FwdFrom = &MessageFwdHeader{}
	m.Decode(dbuf)
	m.ViaBotId = dbuf.Int()
	m.ReplyToMsgId = dbuf.Int()
	m.Date = dbuf.Int()
	m.Message = dbuf.String()
	m.Media = &MessageMedia{}
	m.Decode(dbuf)
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Entities);
	c17 := dbuf.Int()
	if c17 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c17)
	}
	l17 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l17)
	for i := 0; i < int(l17); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	m.Views = dbuf.Int()
	m.EditDate = dbuf.Int()
	m.PostAuthor = dbuf.String()
	return dbuf.err
}

// messageService#9e19a1f6 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;
func (m *TLMessageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageService))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Int(m.FromId)
	x.Bytes(m.ToId.Encode())
	x.Int(m.ReplyToMsgId)
	x.Int(m.Date)
	x.Bytes(m.Action.Encode())
	return x.buf
}

func (m *TLMessageService) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.FromId = dbuf.Int()
	m.ToId = &Peer{}
	m.Decode(dbuf)
	m.ReplyToMsgId = dbuf.Int()
	m.Date = dbuf.Int()
	m.Action = &MessageAction{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageMediaEmpty#3ded6320 = MessageMedia;
func (m *TLMessageMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaEmpty))
	return x.buf
}

func (m *TLMessageMediaEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
func (m *TLMessageMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaPhoto))
	x.Int(m.Flags)
	x.Bytes(m.Photo.Encode())
	x.String(m.Caption)
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLMessageMediaPhoto) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Photo = &Photo{}
	m.Decode(dbuf)
	m.Caption = dbuf.String()
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
func (m *TLMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaGeo))
	x.Bytes(m.Geo.Encode())
	return x.buf
}

func (m *TLMessageMediaGeo) Decode(dbuf *DecodeBuf) error {
	m.Geo = &GeoPoint{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
func (m *TLMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaContact))
	x.String(m.PhoneNumber)
	x.String(m.FirstName)
	x.String(m.LastName)
	x.Int(m.UserId)
	return x.buf
}

func (m *TLMessageMediaContact) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	m.UserId = dbuf.Int()
	return dbuf.err
}

// messageMediaUnsupported#9f84f49e = MessageMedia;
func (m *TLMessageMediaUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaUnsupported))
	return x.buf
}

func (m *TLMessageMediaUnsupported) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
func (m *TLMessageMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaDocument))
	x.Int(m.Flags)
	x.Bytes(m.Document.Encode())
	x.String(m.Caption)
	x.Int(m.TtlSeconds)
	return x.buf
}

func (m *TLMessageMediaDocument) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Document = &Document{}
	m.Decode(dbuf)
	m.Caption = dbuf.String()
	m.TtlSeconds = dbuf.Int()
	return dbuf.err
}

// messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;
func (m *TLMessageMediaWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaWebPage))
	x.Bytes(m.Webpage.Encode())
	return x.buf
}

func (m *TLMessageMediaWebPage) Decode(dbuf *DecodeBuf) error {
	m.Webpage = &WebPage{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
func (m *TLMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaVenue))
	x.Bytes(m.Geo.Encode())
	x.String(m.Title)
	x.String(m.Address)
	x.String(m.Provider)
	x.String(m.VenueId)
	return x.buf
}

func (m *TLMessageMediaVenue) Decode(dbuf *DecodeBuf) error {
	m.Geo = &GeoPoint{}
	m.Decode(dbuf)
	m.Title = dbuf.String()
	m.Address = dbuf.String()
	m.Provider = dbuf.String()
	m.VenueId = dbuf.String()
	return dbuf.err
}

// messageMediaGame#fdb19008 game:Game = MessageMedia;
func (m *TLMessageMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaGame))
	x.Bytes(m.Game.Encode())
	return x.buf
}

func (m *TLMessageMediaGame) Decode(dbuf *DecodeBuf) error {
	m.Game = &Game{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;
func (m *TLMessageMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaInvoice))
	x.Int(m.Flags)
	x.String(m.Title)
	x.String(m.Description)
	x.Bytes(m.Photo.Encode())
	x.Int(m.ReceiptMsgId)
	x.String(m.Currency)
	x.Long(m.TotalAmount)
	x.String(m.StartParam)
	return x.buf
}

func (m *TLMessageMediaInvoice) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Photo = &WebDocument{}
	m.Decode(dbuf)
	m.ReceiptMsgId = dbuf.Int()
	m.Currency = dbuf.String()
	m.TotalAmount = dbuf.Long()
	m.StartParam = dbuf.String()
	return dbuf.err
}

// messageActionEmpty#b6aef7b0 = MessageAction;
func (m *TLMessageActionEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionEmpty))
	return x.buf
}

func (m *TLMessageActionEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
func (m *TLMessageActionChatCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatCreate))
	x.String(m.Title)
	x.VectorInt(m.Users)
	return x.buf
}

func (m *TLMessageActionChatCreate) Decode(dbuf *DecodeBuf) error {
	m.Title = dbuf.String()
	m.Users = dbuf.VectorInt()
	return dbuf.err
}

// messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
func (m *TLMessageActionChatEditTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatEditTitle))
	x.String(m.Title)
	return x.buf
}

func (m *TLMessageActionChatEditTitle) Decode(dbuf *DecodeBuf) error {
	m.Title = dbuf.String()
	return dbuf.err
}

// messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
func (m *TLMessageActionChatEditPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatEditPhoto))
	x.Bytes(m.Photo.Encode())
	return x.buf
}

func (m *TLMessageActionChatEditPhoto) Decode(dbuf *DecodeBuf) error {
	m.Photo = &Photo{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageActionChatDeletePhoto#95e3fbef = MessageAction;
func (m *TLMessageActionChatDeletePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatDeletePhoto))
	return x.buf
}

func (m *TLMessageActionChatDeletePhoto) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
func (m *TLMessageActionChatAddUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatAddUser))
	x.VectorInt(m.Users)
	return x.buf
}

func (m *TLMessageActionChatAddUser) Decode(dbuf *DecodeBuf) error {
	m.Users = dbuf.VectorInt()
	return dbuf.err
}

// messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
func (m *TLMessageActionChatDeleteUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatDeleteUser))
	x.Int(m.UserId)
	return x.buf
}

func (m *TLMessageActionChatDeleteUser) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	return dbuf.err
}

// messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
func (m *TLMessageActionChatJoinedByLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatJoinedByLink))
	x.Int(m.InviterId)
	return x.buf
}

func (m *TLMessageActionChatJoinedByLink) Decode(dbuf *DecodeBuf) error {
	m.InviterId = dbuf.Int()
	return dbuf.err
}

// messageActionChannelCreate#95d2ac92 title:string = MessageAction;
func (m *TLMessageActionChannelCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChannelCreate))
	x.String(m.Title)
	return x.buf
}

func (m *TLMessageActionChannelCreate) Decode(dbuf *DecodeBuf) error {
	m.Title = dbuf.String()
	return dbuf.err
}

// messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
func (m *TLMessageActionChatMigrateTo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatMigrateTo))
	x.Int(m.ChannelId)
	return x.buf
}

func (m *TLMessageActionChatMigrateTo) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	return dbuf.err
}

// messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
func (m *TLMessageActionChannelMigrateFrom) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChannelMigrateFrom))
	x.String(m.Title)
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLMessageActionChannelMigrateFrom) Decode(dbuf *DecodeBuf) error {
	m.Title = dbuf.String()
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// messageActionPinMessage#94bd38ed = MessageAction;
func (m *TLMessageActionPinMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPinMessage))
	return x.buf
}

func (m *TLMessageActionPinMessage) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messageActionHistoryClear#9fbab604 = MessageAction;
func (m *TLMessageActionHistoryClear) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionHistoryClear))
	return x.buf
}

func (m *TLMessageActionHistoryClear) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messageActionGameScore#92a72876 game_id:long score:int = MessageAction;
func (m *TLMessageActionGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionGameScore))
	x.Long(m.GameId)
	x.Int(m.Score)
	return x.buf
}

func (m *TLMessageActionGameScore) Decode(dbuf *DecodeBuf) error {
	m.GameId = dbuf.Long()
	m.Score = dbuf.Int()
	return dbuf.err
}

// messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;
func (m *TLMessageActionPaymentSentMe) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPaymentSentMe))
	x.Int(m.Flags)
	x.String(m.Currency)
	x.Long(m.TotalAmount)
	x.StringBytes(m.Payload)
	x.Bytes(m.Info.Encode())
	x.String(m.ShippingOptionId)
	x.Bytes(m.Charge.Encode())
	return x.buf
}

func (m *TLMessageActionPaymentSentMe) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Currency = dbuf.String()
	m.TotalAmount = dbuf.Long()
	m.Payload = dbuf.StringBytes()
	m.Info = &PaymentRequestedInfo{}
	m.Decode(dbuf)
	m.ShippingOptionId = dbuf.String()
	m.Charge = &PaymentCharge{}
	m.Decode(dbuf)
	return dbuf.err
}

// messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;
func (m *TLMessageActionPaymentSent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPaymentSent))
	x.String(m.Currency)
	x.Long(m.TotalAmount)
	return x.buf
}

func (m *TLMessageActionPaymentSent) Decode(dbuf *DecodeBuf) error {
	m.Currency = dbuf.String()
	m.TotalAmount = dbuf.Long()
	return dbuf.err
}

// messageActionPhoneCall#80e11a7f flags:# call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
func (m *TLMessageActionPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPhoneCall))
	x.Int(m.Flags)
	x.Long(m.CallId)
	x.Bytes(m.Reason.Encode())
	x.Int(m.Duration)
	return x.buf
}

func (m *TLMessageActionPhoneCall) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.CallId = dbuf.Long()
	m.Reason = &PhoneCallDiscardReason{}
	m.Decode(dbuf)
	m.Duration = dbuf.Int()
	return dbuf.err
}

// messageActionScreenshotTaken#4792929b = MessageAction;
func (m *TLMessageActionScreenshotTaken) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionScreenshotTaken))
	return x.buf
}

func (m *TLMessageActionScreenshotTaken) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
func (m *TLDialog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dialog))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.TopMessage)
	x.Int(m.ReadInboxMaxId)
	x.Int(m.ReadOutboxMaxId)
	x.Int(m.UnreadCount)
	x.Int(m.UnreadMentionsCount)
	x.Bytes(m.NotifySettings.Encode())
	x.Int(m.Pts)
	x.Bytes(m.Draft.Encode())
	return x.buf
}

func (m *TLDialog) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Peer = &Peer{}
	m.Decode(dbuf)
	m.TopMessage = dbuf.Int()
	m.ReadInboxMaxId = dbuf.Int()
	m.ReadOutboxMaxId = dbuf.Int()
	m.UnreadCount = dbuf.Int()
	m.UnreadMentionsCount = dbuf.Int()
	m.NotifySettings = &PeerNotifySettings{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.Draft = &DraftMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// photoEmpty#2331b22d id:long = Photo;
func (m *TLPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoEmpty))
	x.Long(m.Id)
	return x.buf
}

func (m *TLPhotoEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	return dbuf.err
}

// photo#9288dd29 flags:# has_stickers:flags.0?true id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
func (m *TLPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photo))
	x.Int(m.Flags)
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	// x.VectorMessage(m.Sizes);
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Sizes)))
	for _, v := range m.Sizes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPhoto) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	// x.VectorMessage(m.Sizes);
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Sizes = make([]*PhotoSize, l6)
	for i := 0; i < int(l6); i++ {
		m.Sizes[i] = &PhotoSize{}
		(*m.Sizes[i]).Decode(dbuf)
	}
	return dbuf.err
}

// photoSizeEmpty#e17e23c type:string = PhotoSize;
func (m *TLPhotoSizeEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoSizeEmpty))
	x.String(m.Type)
	return x.buf
}

func (m *TLPhotoSizeEmpty) Decode(dbuf *DecodeBuf) error {
	m.Type = dbuf.String()
	return dbuf.err
}

// photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;
func (m *TLPhotoSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoSize))
	x.String(m.Type)
	x.Bytes(m.Location.Encode())
	x.Int(m.W)
	x.Int(m.H)
	x.Int(m.Size)
	return x.buf
}

func (m *TLPhotoSize) Decode(dbuf *DecodeBuf) error {
	m.Type = dbuf.String()
	m.Location = &FileLocation{}
	m.Decode(dbuf)
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	m.Size = dbuf.Int()
	return dbuf.err
}

// photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;
func (m *TLPhotoCachedSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoCachedSize))
	x.String(m.Type)
	x.Bytes(m.Location.Encode())
	x.Int(m.W)
	x.Int(m.H)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLPhotoCachedSize) Decode(dbuf *DecodeBuf) error {
	m.Type = dbuf.String()
	m.Location = &FileLocation{}
	m.Decode(dbuf)
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// geoPointEmpty#1117dd5f = GeoPoint;
func (m *TLGeoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_geoPointEmpty))
	return x.buf
}

func (m *TLGeoPointEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// geoPoint#2049d70c long:double lat:double = GeoPoint;
func (m *TLGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_geoPoint))
	x.Double(m.Long)
	x.Double(m.Lat)
	return x.buf
}

func (m *TLGeoPoint) Decode(dbuf *DecodeBuf) error {
	m.Long = dbuf.Double()
	m.Lat = dbuf.Double()
	return dbuf.err
}

// auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;
func (m *TLAuthCheckedPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_checkedPhone))
	x.Bytes(m.PhoneRegistered.Encode())
	return x.buf
}

func (m *TLAuthCheckedPhone) Decode(dbuf *DecodeBuf) error {
	m.PhoneRegistered = &Bool{}
	m.Decode(dbuf)
	return dbuf.err
}

// auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;
func (m *TLAuthSentCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCode))
	x.Int(m.Flags)
	x.Bytes(m.Type.Encode())
	x.String(m.PhoneCodeHash)
	x.Bytes(m.NextType.Encode())
	x.Int(m.Timeout)
	return x.buf
}

func (m *TLAuthSentCode) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Type = &Auth_SentCodeType{}
	m.Decode(dbuf)
	m.PhoneCodeHash = dbuf.String()
	m.NextType = &Auth_CodeType{}
	m.Decode(dbuf)
	m.Timeout = dbuf.Int()
	return dbuf.err
}

// auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;
func (m *TLAuthAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_authorization))
	x.Int(m.Flags)
	x.Int(m.TmpSessions)
	x.Bytes(m.User.Encode())
	return x.buf
}

func (m *TLAuthAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.TmpSessions = dbuf.Int()
	m.User = &User{}
	m.Decode(dbuf)
	return dbuf.err
}

// auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
func (m *TLAuthExportedAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_exportedAuthorization))
	x.Int(m.Id)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLAuthExportedAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;
func (m *TLInputNotifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyPeer))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLInputNotifyPeer) Decode(dbuf *DecodeBuf) error {
	m.Peer = &InputPeer{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputNotifyUsers#193b4417 = InputNotifyPeer;
func (m *TLInputNotifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyUsers))
	return x.buf
}

func (m *TLInputNotifyUsers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputNotifyChats#4a95e84e = InputNotifyPeer;
func (m *TLInputNotifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyChats))
	return x.buf
}

func (m *TLInputNotifyChats) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputNotifyAll#a429b886 = InputNotifyPeer;
func (m *TLInputNotifyAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyAll))
	return x.buf
}

func (m *TLInputNotifyAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
func (m *TLInputPeerNotifyEventsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerNotifyEventsEmpty))
	return x.buf
}

func (m *TLInputPeerNotifyEventsEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;
func (m *TLInputPeerNotifyEventsAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerNotifyEventsAll))
	return x.buf
}

func (m *TLInputPeerNotifyEventsAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPeerNotifySettings#38935eb2 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = InputPeerNotifySettings;
func (m *TLInputPeerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerNotifySettings))
	x.Int(m.Flags)
	x.Int(m.MuteUntil)
	x.String(m.Sound)
	return x.buf
}

func (m *TLInputPeerNotifySettings) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.MuteUntil = dbuf.Int()
	m.Sound = dbuf.String()
	return dbuf.err
}

// peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;
func (m *TLPeerNotifyEventsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifyEventsEmpty))
	return x.buf
}

func (m *TLPeerNotifyEventsEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;
func (m *TLPeerNotifyEventsAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifyEventsAll))
	return x.buf
}

func (m *TLPeerNotifyEventsAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;
func (m *TLPeerNotifySettingsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifySettingsEmpty))
	return x.buf
}

func (m *TLPeerNotifySettingsEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// peerNotifySettings#9acda4c0 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = PeerNotifySettings;
func (m *TLPeerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifySettings))
	x.Int(m.Flags)
	x.Int(m.MuteUntil)
	x.String(m.Sound)
	return x.buf
}

func (m *TLPeerNotifySettings) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.MuteUntil = dbuf.Int()
	m.Sound = dbuf.String()
	return dbuf.err
}

// peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;
func (m *TLPeerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerSettings))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLPeerSettings) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
func (m *TLWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_wallPaper))
	x.Int(m.Id)
	x.String(m.Title)
	// x.VectorMessage(m.Sizes);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Sizes)))
	for _, v := range m.Sizes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.Color)
	return x.buf
}

func (m *TLWallPaper) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Title = dbuf.String()
	// x.VectorMessage(m.Sizes);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Sizes = make([]*PhotoSize, l3)
	for i := 0; i < int(l3); i++ {
		m.Sizes[i] = &PhotoSize{}
		(*m.Sizes[i]).Decode(dbuf)
	}
	m.Color = dbuf.Int()
	return dbuf.err
}

// wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
func (m *TLWallPaperSolid) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_wallPaperSolid))
	x.Int(m.Id)
	x.String(m.Title)
	x.Int(m.BgColor)
	x.Int(m.Color)
	return x.buf
}

func (m *TLWallPaperSolid) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Title = dbuf.String()
	m.BgColor = dbuf.Int()
	m.Color = dbuf.Int()
	return dbuf.err
}

// inputReportReasonSpam#58dbcab8 = ReportReason;
func (m *TLInputReportReasonSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonSpam))
	return x.buf
}

func (m *TLInputReportReasonSpam) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputReportReasonViolence#1e22c78d = ReportReason;
func (m *TLInputReportReasonViolence) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonViolence))
	return x.buf
}

func (m *TLInputReportReasonViolence) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputReportReasonPornography#2e59d922 = ReportReason;
func (m *TLInputReportReasonPornography) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonPornography))
	return x.buf
}

func (m *TLInputReportReasonPornography) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputReportReasonOther#e1746d0a text:string = ReportReason;
func (m *TLInputReportReasonOther) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonOther))
	x.String(m.Text)
	return x.buf
}

func (m *TLInputReportReasonOther) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// userFull#f220f3f flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo common_chats_count:int = UserFull;
func (m *TLUserFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userFull))
	x.Int(m.Flags)
	x.Bytes(m.User.Encode())
	x.String(m.About)
	x.Bytes(m.Link.Encode())
	x.Bytes(m.ProfilePhoto.Encode())
	x.Bytes(m.NotifySettings.Encode())
	x.Bytes(m.BotInfo.Encode())
	x.Int(m.CommonChatsCount)
	return x.buf
}

func (m *TLUserFull) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.User = &User{}
	m.Decode(dbuf)
	m.About = dbuf.String()
	m.Link = &Contacts_Link{}
	m.Decode(dbuf)
	m.ProfilePhoto = &Photo{}
	m.Decode(dbuf)
	m.NotifySettings = &PeerNotifySettings{}
	m.Decode(dbuf)
	m.BotInfo = &BotInfo{}
	m.Decode(dbuf)
	m.CommonChatsCount = dbuf.Int()
	return dbuf.err
}

// contact#f911c994 user_id:int mutual:Bool = Contact;
func (m *TLContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contact))
	x.Int(m.UserId)
	x.Bytes(m.Mutual.Encode())
	return x.buf
}

func (m *TLContact) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Mutual = &Bool{}
	m.Decode(dbuf)
	return dbuf.err
}

// importedContact#d0028438 user_id:int client_id:long = ImportedContact;
func (m *TLImportedContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_importedContact))
	x.Int(m.UserId)
	x.Long(m.ClientId)
	return x.buf
}

func (m *TLImportedContact) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.ClientId = dbuf.Long()
	return dbuf.err
}

// contactBlocked#561bc879 user_id:int date:int = ContactBlocked;
func (m *TLContactBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactBlocked))
	x.Int(m.UserId)
	x.Int(m.Date)
	return x.buf
}

func (m *TLContactBlocked) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;
func (m *TLContactStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactStatus))
	x.Int(m.UserId)
	x.Bytes(m.Status.Encode())
	return x.buf
}

func (m *TLContactStatus) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Status = &UserStatus{}
	m.Decode(dbuf)
	return dbuf.err
}

// contacts.link#3ace484c my_link:ContactLink foreign_link:ContactLink user:User = contacts.Link;
func (m *TLContactsLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_link))
	x.Bytes(m.MyLink.Encode())
	x.Bytes(m.ForeignLink.Encode())
	x.Bytes(m.User.Encode())
	return x.buf
}

func (m *TLContactsLink) Decode(dbuf *DecodeBuf) error {
	m.MyLink = &ContactLink{}
	m.Decode(dbuf)
	m.ForeignLink = &ContactLink{}
	m.Decode(dbuf)
	m.User = &User{}
	m.Decode(dbuf)
	return dbuf.err
}

// contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;
func (m *TLContactsContactsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_contactsNotModified))
	return x.buf
}

func (m *TLContactsContactsNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// contacts.contacts#eae87e42 contacts:Vector<Contact> saved_count:int users:Vector<User> = contacts.Contacts;
func (m *TLContactsContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_contacts))
	// x.VectorMessage(m.Contacts);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Contacts)))
	for _, v := range m.Contacts {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.SavedCount)
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsContacts) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Contacts);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Contacts = make([]*Contact, l1)
	for i := 0; i < int(l1); i++ {
		m.Contacts[i] = &Contact{}
		(*m.Contacts[i]).Decode(dbuf)
	}
	m.SavedCount = dbuf.Int()
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// contacts.importedContacts#77d01c3b imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
func (m *TLContactsImportedContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_importedContacts))
	// x.VectorMessage(m.Imported);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Imported)))
	for _, v := range m.Imported {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.PopularInvites);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.PopularInvites)))
	for _, v := range m.PopularInvites {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.VectorLong(m.RetryContacts)
	// x.VectorMessage(m.Users);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsImportedContacts) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Imported);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Imported = make([]*ImportedContact, l1)
	for i := 0; i < int(l1); i++ {
		m.Imported[i] = &ImportedContact{}
		(*m.Imported[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.PopularInvites);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.PopularInvites = make([]*PopularContact, l2)
	for i := 0; i < int(l2); i++ {
		m.PopularInvites[i] = &PopularContact{}
		(*m.PopularInvites[i]).Decode(dbuf)
	}
	m.RetryContacts = dbuf.VectorLong()
	// x.VectorMessage(m.Users);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Users = make([]*User, l4)
	for i := 0; i < int(l4); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
func (m *TLContactsBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_blocked))
	// x.VectorMessage(m.Blocked);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Blocked)))
	for _, v := range m.Blocked {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsBlocked) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Blocked);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Blocked = make([]*ContactBlocked, l1)
	for i := 0; i < int(l1); i++ {
		m.Blocked[i] = &ContactBlocked{}
		(*m.Blocked[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
func (m *TLContactsBlockedSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_blockedSlice))
	x.Int(m.Count)
	// x.VectorMessage(m.Blocked);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Blocked)))
	for _, v := range m.Blocked {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsBlockedSlice) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Blocked);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Blocked = make([]*ContactBlocked, l2)
	for i := 0; i < int(l2); i++ {
		m.Blocked[i] = &ContactBlocked{}
		(*m.Blocked[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
func (m *TLMessagesDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dialogs))
	// x.VectorMessage(m.Dialogs);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Dialogs)))
	for _, v := range m.Dialogs {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Messages);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesDialogs) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Dialogs);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Dialogs = make([]*Dialog, l1)
	for i := 0; i < int(l1); i++ {
		m.Dialogs[i] = &Dialog{}
		(*m.Dialogs[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Messages);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Messages = make([]*Message, l2)
	for i := 0; i < int(l2); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Chats = make([]*Chat, l3)
	for i := 0; i < int(l3); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Users = make([]*User, l4)
	for i := 0; i < int(l4); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
func (m *TLMessagesDialogsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dialogsSlice))
	x.Int(m.Count)
	// x.VectorMessage(m.Dialogs);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Dialogs)))
	for _, v := range m.Dialogs {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Messages);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesDialogsSlice) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Dialogs);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Dialogs = make([]*Dialog, l2)
	for i := 0; i < int(l2); i++ {
		m.Dialogs[i] = &Dialog{}
		(*m.Dialogs[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Messages);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Messages = make([]*Message, l3)
	for i := 0; i < int(l3); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Chats = make([]*Chat, l4)
	for i := 0; i < int(l4); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Users = make([]*User, l5)
	for i := 0; i < int(l5); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *TLMessagesMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_messages))
	// x.VectorMessage(m.Messages);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesMessages) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Messages);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Messages = make([]*Message, l1)
	for i := 0; i < int(l1); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *TLMessagesMessagesSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_messagesSlice))
	x.Int(m.Count)
	// x.VectorMessage(m.Messages);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesMessagesSlice) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Messages);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Messages = make([]*Message, l2)
	for i := 0; i < int(l2); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Chats = make([]*Chat, l3)
	for i := 0; i < int(l3); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Users = make([]*User, l4)
	for i := 0; i < int(l4); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *TLMessagesChannelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_channelMessages))
	x.Int(m.Flags)
	x.Int(m.Pts)
	x.Int(m.Count)
	// x.VectorMessage(m.Messages);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesChannelMessages) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Pts = dbuf.Int()
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Messages);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Messages = make([]*Message, l4)
	for i := 0; i < int(l4); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Chats = make([]*Chat, l5)
	for i := 0; i < int(l5); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Users = make([]*User, l6)
	for i := 0; i < int(l6); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.chats#64ff9fd5 chats:Vector<Chat> = messages.Chats;
func (m *TLMessagesChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_chats))
	// x.VectorMessage(m.Chats);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesChats) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Chats);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Chats = make([]*Chat, l1)
	for i := 0; i < int(l1); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.chatsSlice#9cd81144 count:int chats:Vector<Chat> = messages.Chats;
func (m *TLMessagesChatsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_chatsSlice))
	x.Int(m.Count)
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesChatsSlice) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;
func (m *TLMessagesChatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_chatFull))
	x.Bytes(m.FullChat.Encode())
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesChatFull) Decode(dbuf *DecodeBuf) error {
	m.FullChat = &ChatFull{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.affectedHistory#b45c69d1 pts:int pts_count:int offset:int = messages.AffectedHistory;
func (m *TLMessagesAffectedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_affectedHistory))
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	x.Int(m.Offset)
	return x.buf
}

func (m *TLMessagesAffectedHistory) Decode(dbuf *DecodeBuf) error {
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	m.Offset = dbuf.Int()
	return dbuf.err
}

// inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
func (m *TLInputMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterEmpty))
	return x.buf
}

func (m *TLInputMessagesFilterEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterPhotos#9609a51c = MessagesFilter;
func (m *TLInputMessagesFilterPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhotos))
	return x.buf
}

func (m *TLInputMessagesFilterPhotos) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
func (m *TLInputMessagesFilterVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterVideo))
	return x.buf
}

func (m *TLInputMessagesFilterVideo) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
func (m *TLInputMessagesFilterPhotoVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideo))
	return x.buf
}

func (m *TLInputMessagesFilterPhotoVideo) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;
func (m *TLInputMessagesFilterPhotoVideoDocuments) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideoDocuments))
	return x.buf
}

func (m *TLInputMessagesFilterPhotoVideoDocuments) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterDocument#9eddf188 = MessagesFilter;
func (m *TLInputMessagesFilterDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterDocument))
	return x.buf
}

func (m *TLInputMessagesFilterDocument) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
func (m *TLInputMessagesFilterUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterUrl))
	return x.buf
}

func (m *TLInputMessagesFilterUrl) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterGif#ffc86587 = MessagesFilter;
func (m *TLInputMessagesFilterGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterGif))
	return x.buf
}

func (m *TLInputMessagesFilterGif) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterVoice#50f5c392 = MessagesFilter;
func (m *TLInputMessagesFilterVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterVoice))
	return x.buf
}

func (m *TLInputMessagesFilterVoice) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterMusic#3751b49e = MessagesFilter;
func (m *TLInputMessagesFilterMusic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterMusic))
	return x.buf
}

func (m *TLInputMessagesFilterMusic) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
func (m *TLInputMessagesFilterChatPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterChatPhotos))
	return x.buf
}

func (m *TLInputMessagesFilterChatPhotos) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
func (m *TLInputMessagesFilterPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhoneCalls))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLInputMessagesFilterPhoneCalls) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
func (m *TLInputMessagesFilterRoundVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterRoundVoice))
	return x.buf
}

func (m *TLInputMessagesFilterRoundVoice) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
func (m *TLInputMessagesFilterRoundVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterRoundVideo))
	return x.buf
}

func (m *TLInputMessagesFilterRoundVideo) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
func (m *TLInputMessagesFilterMyMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterMyMentions))
	return x.buf
}

func (m *TLInputMessagesFilterMyMentions) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
func (m *TLUpdateNewMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewMessage))
	x.Bytes(m.Message.Encode())
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateNewMessage) Decode(dbuf *DecodeBuf) error {
	m.Message = &Message{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateMessageID#4e90bfd6 id:int random_id:long = Update;
func (m *TLUpdateMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateMessageID))
	x.Int(m.Id)
	x.Long(m.RandomId)
	return x.buf
}

func (m *TLUpdateMessageID) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.RandomId = dbuf.Long()
	return dbuf.err
}

// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
func (m *TLUpdateDeleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDeleteMessages))
	x.VectorInt(m.Messages)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateDeleteMessages) Decode(dbuf *DecodeBuf) error {
	m.Messages = dbuf.VectorInt()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
func (m *TLUpdateUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserTyping))
	x.Int(m.UserId)
	x.Bytes(m.Action.Encode())
	return x.buf
}

func (m *TLUpdateUserTyping) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Action = &SendMessageAction{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
func (m *TLUpdateChatUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatUserTyping))
	x.Int(m.ChatId)
	x.Int(m.UserId)
	x.Bytes(m.Action.Encode())
	return x.buf
}

func (m *TLUpdateChatUserTyping) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.UserId = dbuf.Int()
	m.Action = &SendMessageAction{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateChatParticipants#7761198 participants:ChatParticipants = Update;
func (m *TLUpdateChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipants))
	x.Bytes(m.Participants.Encode())
	return x.buf
}

func (m *TLUpdateChatParticipants) Decode(dbuf *DecodeBuf) error {
	m.Participants = &ChatParticipants{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;
func (m *TLUpdateUserStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserStatus))
	x.Int(m.UserId)
	x.Bytes(m.Status.Encode())
	return x.buf
}

func (m *TLUpdateUserStatus) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Status = &UserStatus{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
func (m *TLUpdateUserName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserName))
	x.Int(m.UserId)
	x.String(m.FirstName)
	x.String(m.LastName)
	x.String(m.Username)
	return x.buf
}

func (m *TLUpdateUserName) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	m.Username = dbuf.String()
	return dbuf.err
}

// updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
func (m *TLUpdateUserPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserPhoto))
	x.Int(m.UserId)
	x.Int(m.Date)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.Previous.Encode())
	return x.buf
}

func (m *TLUpdateUserPhoto) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Date = dbuf.Int()
	m.Photo = &UserProfilePhoto{}
	m.Decode(dbuf)
	m.Previous = &Bool{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateContactRegistered#2575bbb9 user_id:int date:int = Update;
func (m *TLUpdateContactRegistered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateContactRegistered))
	x.Int(m.UserId)
	x.Int(m.Date)
	return x.buf
}

func (m *TLUpdateContactRegistered) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// updateContactLink#9d2e67c5 user_id:int my_link:ContactLink foreign_link:ContactLink = Update;
func (m *TLUpdateContactLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateContactLink))
	x.Int(m.UserId)
	x.Bytes(m.MyLink.Encode())
	x.Bytes(m.ForeignLink.Encode())
	return x.buf
}

func (m *TLUpdateContactLink) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.MyLink = &ContactLink{}
	m.Decode(dbuf)
	m.ForeignLink = &ContactLink{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;
func (m *TLUpdateNewEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewEncryptedMessage))
	x.Bytes(m.Message.Encode())
	x.Int(m.Qts)
	return x.buf
}

func (m *TLUpdateNewEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	m.Message = &EncryptedMessage{}
	m.Decode(dbuf)
	m.Qts = dbuf.Int()
	return dbuf.err
}

// updateEncryptedChatTyping#1710f156 chat_id:int = Update;
func (m *TLUpdateEncryptedChatTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEncryptedChatTyping))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLUpdateEncryptedChatTyping) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;
func (m *TLUpdateEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEncryption))
	x.Bytes(m.Chat.Encode())
	x.Int(m.Date)
	return x.buf
}

func (m *TLUpdateEncryption) Decode(dbuf *DecodeBuf) error {
	m.Chat = &EncryptedChat{}
	m.Decode(dbuf)
	m.Date = dbuf.Int()
	return dbuf.err
}

// updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;
func (m *TLUpdateEncryptedMessagesRead) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEncryptedMessagesRead))
	x.Int(m.ChatId)
	x.Int(m.MaxDate)
	x.Int(m.Date)
	return x.buf
}

func (m *TLUpdateEncryptedMessagesRead) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.MaxDate = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
func (m *TLUpdateChatParticipantAdd) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipantAdd))
	x.Int(m.ChatId)
	x.Int(m.UserId)
	x.Int(m.InviterId)
	x.Int(m.Date)
	x.Int(m.Version)
	return x.buf
}

func (m *TLUpdateChatParticipantAdd) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.UserId = dbuf.Int()
	m.InviterId = dbuf.Int()
	m.Date = dbuf.Int()
	m.Version = dbuf.Int()
	return dbuf.err
}

// updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int version:int = Update;
func (m *TLUpdateChatParticipantDelete) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipantDelete))
	x.Int(m.ChatId)
	x.Int(m.UserId)
	x.Int(m.Version)
	return x.buf
}

func (m *TLUpdateChatParticipantDelete) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.UserId = dbuf.Int()
	m.Version = dbuf.Int()
	return dbuf.err
}

// updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;
func (m *TLUpdateDcOptions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDcOptions))
	// x.VectorMessage(m.DcOptions);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.DcOptions)))
	for _, v := range m.DcOptions {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdateDcOptions) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.DcOptions);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.DcOptions = make([]*DcOption, l1)
	for i := 0; i < int(l1); i++ {
		m.DcOptions[i] = &DcOption{}
		(*m.DcOptions[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;
func (m *TLUpdateUserBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserBlocked))
	x.Int(m.UserId)
	x.Bytes(m.Blocked.Encode())
	return x.buf
}

func (m *TLUpdateUserBlocked) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Blocked = &Bool{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;
func (m *TLUpdateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNotifySettings))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.NotifySettings.Encode())
	return x.buf
}

func (m *TLUpdateNotifySettings) Decode(dbuf *DecodeBuf) error {
	m.Peer = &NotifyPeer{}
	m.Decode(dbuf)
	m.NotifySettings = &PeerNotifySettings{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;
func (m *TLUpdateServiceNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateServiceNotification))
	x.Int(m.Flags)
	x.Int(m.InboxDate)
	x.String(m.Type)
	x.String(m.Message)
	x.Bytes(m.Media.Encode())
	// x.VectorMessage(m.Entities);
	x7 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x7, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x7[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdateServiceNotification) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.InboxDate = dbuf.Int()
	m.Type = dbuf.String()
	m.Message = dbuf.String()
	m.Media = &MessageMedia{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Entities);
	c7 := dbuf.Int()
	if c7 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c7)
	}
	l7 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l7)
	for i := 0; i < int(l7); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;
func (m *TLUpdatePrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePrivacy))
	x.Bytes(m.Key.Encode())
	// x.VectorMessage(m.Rules);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Rules)))
	for _, v := range m.Rules {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdatePrivacy) Decode(dbuf *DecodeBuf) error {
	m.Key = &PrivacyKey{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Rules);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Rules = make([]*PrivacyRule, l2)
	for i := 0; i < int(l2); i++ {
		m.Rules[i] = &PrivacyRule{}
		(*m.Rules[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updateUserPhone#12b9417b user_id:int phone:string = Update;
func (m *TLUpdateUserPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserPhone))
	x.Int(m.UserId)
	x.String(m.Phone)
	return x.buf
}

func (m *TLUpdateUserPhone) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Phone = dbuf.String()
	return dbuf.err
}

// updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
func (m *TLUpdateReadHistoryInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadHistoryInbox))
	x.Bytes(m.Peer.Encode())
	x.Int(m.MaxId)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateReadHistoryInbox) Decode(dbuf *DecodeBuf) error {
	m.Peer = &Peer{}
	m.Decode(dbuf)
	m.MaxId = dbuf.Int()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
func (m *TLUpdateReadHistoryOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadHistoryOutbox))
	x.Bytes(m.Peer.Encode())
	x.Int(m.MaxId)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateReadHistoryOutbox) Decode(dbuf *DecodeBuf) error {
	m.Peer = &Peer{}
	m.Decode(dbuf)
	m.MaxId = dbuf.Int()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;
func (m *TLUpdateWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateWebPage))
	x.Bytes(m.Webpage.Encode())
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateWebPage) Decode(dbuf *DecodeBuf) error {
	m.Webpage = &WebPage{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
func (m *TLUpdateReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadMessagesContents))
	x.VectorInt(m.Messages)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateReadMessagesContents) Decode(dbuf *DecodeBuf) error {
	m.Messages = dbuf.VectorInt()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;
func (m *TLUpdateChannelTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelTooLong))
	x.Int(m.Flags)
	x.Int(m.ChannelId)
	x.Int(m.Pts)
	return x.buf
}

func (m *TLUpdateChannelTooLong) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.ChannelId = dbuf.Int()
	m.Pts = dbuf.Int()
	return dbuf.err
}

// updateChannel#b6d45656 channel_id:int = Update;
func (m *TLUpdateChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannel))
	x.Int(m.ChannelId)
	return x.buf
}

func (m *TLUpdateChannel) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	return dbuf.err
}

// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
func (m *TLUpdateNewChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewChannelMessage))
	x.Bytes(m.Message.Encode())
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateNewChannelMessage) Decode(dbuf *DecodeBuf) error {
	m.Message = &Message{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
func (m *TLUpdateReadChannelInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadChannelInbox))
	x.Int(m.ChannelId)
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLUpdateReadChannelInbox) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
func (m *TLUpdateDeleteChannelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDeleteChannelMessages))
	x.Int(m.ChannelId)
	x.VectorInt(m.Messages)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateDeleteChannelMessages) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.Messages = dbuf.VectorInt()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateChannelMessageViews#98a12b4b channel_id:int id:int views:int = Update;
func (m *TLUpdateChannelMessageViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelMessageViews))
	x.Int(m.ChannelId)
	x.Int(m.Id)
	x.Int(m.Views)
	return x.buf
}

func (m *TLUpdateChannelMessageViews) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.Id = dbuf.Int()
	m.Views = dbuf.Int()
	return dbuf.err
}

// updateChatAdmins#6e947941 chat_id:int enabled:Bool version:int = Update;
func (m *TLUpdateChatAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatAdmins))
	x.Int(m.ChatId)
	x.Bytes(m.Enabled.Encode())
	x.Int(m.Version)
	return x.buf
}

func (m *TLUpdateChatAdmins) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.Enabled = &Bool{}
	m.Decode(dbuf)
	m.Version = dbuf.Int()
	return dbuf.err
}

// updateChatParticipantAdmin#b6901959 chat_id:int user_id:int is_admin:Bool version:int = Update;
func (m *TLUpdateChatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipantAdmin))
	x.Int(m.ChatId)
	x.Int(m.UserId)
	x.Bytes(m.IsAdmin.Encode())
	x.Int(m.Version)
	return x.buf
}

func (m *TLUpdateChatParticipantAdmin) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.UserId = dbuf.Int()
	m.IsAdmin = &Bool{}
	m.Decode(dbuf)
	m.Version = dbuf.Int()
	return dbuf.err
}

// updateNewStickerSet#688a30aa stickerset:messages.StickerSet = Update;
func (m *TLUpdateNewStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewStickerSet))
	x.Bytes(m.Stickerset.Encode())
	return x.buf
}

func (m *TLUpdateNewStickerSet) Decode(dbuf *DecodeBuf) error {
	m.Stickerset = &Messages_StickerSet{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateStickerSetsOrder#bb2d201 flags:# masks:flags.0?true order:Vector<long> = Update;
func (m *TLUpdateStickerSetsOrder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateStickerSetsOrder))
	x.Int(m.Flags)
	x.VectorLong(m.Order)
	return x.buf
}

func (m *TLUpdateStickerSetsOrder) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Order = dbuf.VectorLong()
	return dbuf.err
}

// updateStickerSets#43ae3dec = Update;
func (m *TLUpdateStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateStickerSets))
	return x.buf
}

func (m *TLUpdateStickerSets) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateSavedGifs#9375341e = Update;
func (m *TLUpdateSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateSavedGifs))
	return x.buf
}

func (m *TLUpdateSavedGifs) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateBotInlineQuery#54826690 flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint offset:string = Update;
func (m *TLUpdateBotInlineQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotInlineQuery))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.Int(m.UserId)
	x.String(m.Query)
	x.Bytes(m.Geo.Encode())
	x.String(m.Offset)
	return x.buf
}

func (m *TLUpdateBotInlineQuery) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.UserId = dbuf.Int()
	m.Query = dbuf.String()
	m.Geo = &GeoPoint{}
	m.Decode(dbuf)
	m.Offset = dbuf.String()
	return dbuf.err
}

// updateBotInlineSend#e48f964 flags:# user_id:int query:string geo:flags.0?GeoPoint id:string msg_id:flags.1?InputBotInlineMessageID = Update;
func (m *TLUpdateBotInlineSend) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotInlineSend))
	x.Int(m.Flags)
	x.Int(m.UserId)
	x.String(m.Query)
	x.Bytes(m.Geo.Encode())
	x.String(m.Id)
	x.Bytes(m.MsgId.Encode())
	return x.buf
}

func (m *TLUpdateBotInlineSend) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.UserId = dbuf.Int()
	m.Query = dbuf.String()
	m.Geo = &GeoPoint{}
	m.Decode(dbuf)
	m.Id = dbuf.String()
	m.MsgId = &InputBotInlineMessageID{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
func (m *TLUpdateEditChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEditChannelMessage))
	x.Bytes(m.Message.Encode())
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateEditChannelMessage) Decode(dbuf *DecodeBuf) error {
	m.Message = &Message{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
func (m *TLUpdateChannelPinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelPinnedMessage))
	x.Int(m.ChannelId)
	x.Int(m.Id)
	return x.buf
}

func (m *TLUpdateChannelPinnedMessage) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.Id = dbuf.Int()
	return dbuf.err
}

// updateBotCallbackQuery#e73547e1 flags:# query_id:long user_id:int peer:Peer msg_id:int chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
func (m *TLUpdateBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotCallbackQuery))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.Int(m.UserId)
	x.Bytes(m.Peer.Encode())
	x.Int(m.MsgId)
	x.Long(m.ChatInstance)
	x.StringBytes(m.Data)
	x.String(m.GameShortName)
	return x.buf
}

func (m *TLUpdateBotCallbackQuery) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.UserId = dbuf.Int()
	m.Peer = &Peer{}
	m.Decode(dbuf)
	m.MsgId = dbuf.Int()
	m.ChatInstance = dbuf.Long()
	m.Data = dbuf.StringBytes()
	m.GameShortName = dbuf.String()
	return dbuf.err
}

// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
func (m *TLUpdateEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEditMessage))
	x.Bytes(m.Message.Encode())
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateEditMessage) Decode(dbuf *DecodeBuf) error {
	m.Message = &Message{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateInlineBotCallbackQuery#f9d27a5a flags:# query_id:long user_id:int msg_id:InputBotInlineMessageID chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
func (m *TLUpdateInlineBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateInlineBotCallbackQuery))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.Int(m.UserId)
	x.Bytes(m.MsgId.Encode())
	x.Long(m.ChatInstance)
	x.StringBytes(m.Data)
	x.String(m.GameShortName)
	return x.buf
}

func (m *TLUpdateInlineBotCallbackQuery) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.UserId = dbuf.Int()
	m.MsgId = &InputBotInlineMessageID{}
	m.Decode(dbuf)
	m.ChatInstance = dbuf.Long()
	m.Data = dbuf.StringBytes()
	m.GameShortName = dbuf.String()
	return dbuf.err
}

// updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
func (m *TLUpdateReadChannelOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadChannelOutbox))
	x.Int(m.ChannelId)
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLUpdateReadChannelOutbox) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// updateDraftMessage#ee2bb969 peer:Peer draft:DraftMessage = Update;
func (m *TLUpdateDraftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDraftMessage))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Draft.Encode())
	return x.buf
}

func (m *TLUpdateDraftMessage) Decode(dbuf *DecodeBuf) error {
	m.Peer = &Peer{}
	m.Decode(dbuf)
	m.Draft = &DraftMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateReadFeaturedStickers#571d2742 = Update;
func (m *TLUpdateReadFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadFeaturedStickers))
	return x.buf
}

func (m *TLUpdateReadFeaturedStickers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateRecentStickers#9a422c20 = Update;
func (m *TLUpdateRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateRecentStickers))
	return x.buf
}

func (m *TLUpdateRecentStickers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateConfig#a229dd06 = Update;
func (m *TLUpdateConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateConfig))
	return x.buf
}

func (m *TLUpdateConfig) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updatePtsChanged#3354678f = Update;
func (m *TLUpdatePtsChanged) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePtsChanged))
	return x.buf
}

func (m *TLUpdatePtsChanged) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
func (m *TLUpdateChannelWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelWebPage))
	x.Int(m.ChannelId)
	x.Bytes(m.Webpage.Encode())
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLUpdateChannelWebPage) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.Webpage = &WebPage{}
	m.Decode(dbuf)
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// updateDialogPinned#d711a2cc flags:# pinned:flags.0?true peer:Peer = Update;
func (m *TLUpdateDialogPinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDialogPinned))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLUpdateDialogPinned) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Peer = &Peer{}
	m.Decode(dbuf)
	return dbuf.err
}

// updatePinnedDialogs#d8caf68d flags:# order:flags.0?Vector<Peer> = Update;
func (m *TLUpdatePinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePinnedDialogs))
	x.Int(m.Flags)
	// x.VectorMessage(m.Order);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Order)))
	for _, v := range m.Order {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdatePinnedDialogs) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.VectorMessage(m.Order);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Order = make([]*Peer, l2)
	for i := 0; i < int(l2); i++ {
		m.Order[i] = &Peer{}
		(*m.Order[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updateBotWebhookJSON#8317c0c3 data:DataJSON = Update;
func (m *TLUpdateBotWebhookJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotWebhookJSON))
	x.Bytes(m.Data.Encode())
	return x.buf
}

func (m *TLUpdateBotWebhookJSON) Decode(dbuf *DecodeBuf) error {
	m.Data = &DataJSON{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateBotWebhookJSONQuery#9b9240a6 query_id:long data:DataJSON timeout:int = Update;
func (m *TLUpdateBotWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotWebhookJSONQuery))
	x.Long(m.QueryId)
	x.Bytes(m.Data.Encode())
	x.Int(m.Timeout)
	return x.buf
}

func (m *TLUpdateBotWebhookJSONQuery) Decode(dbuf *DecodeBuf) error {
	m.QueryId = dbuf.Long()
	m.Data = &DataJSON{}
	m.Decode(dbuf)
	m.Timeout = dbuf.Int()
	return dbuf.err
}

// updateBotShippingQuery#e0cdc940 query_id:long user_id:int payload:bytes shipping_address:PostAddress = Update;
func (m *TLUpdateBotShippingQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotShippingQuery))
	x.Long(m.QueryId)
	x.Int(m.UserId)
	x.StringBytes(m.Payload)
	x.Bytes(m.ShippingAddress.Encode())
	return x.buf
}

func (m *TLUpdateBotShippingQuery) Decode(dbuf *DecodeBuf) error {
	m.QueryId = dbuf.Long()
	m.UserId = dbuf.Int()
	m.Payload = dbuf.StringBytes()
	m.ShippingAddress = &PostAddress{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateBotPrecheckoutQuery#5d2f3aa9 flags:# query_id:long user_id:int payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string currency:string total_amount:long = Update;
func (m *TLUpdateBotPrecheckoutQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotPrecheckoutQuery))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.Int(m.UserId)
	x.StringBytes(m.Payload)
	x.Bytes(m.Info.Encode())
	x.String(m.ShippingOptionId)
	x.String(m.Currency)
	x.Long(m.TotalAmount)
	return x.buf
}

func (m *TLUpdateBotPrecheckoutQuery) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.UserId = dbuf.Int()
	m.Payload = dbuf.StringBytes()
	m.Info = &PaymentRequestedInfo{}
	m.Decode(dbuf)
	m.ShippingOptionId = dbuf.String()
	m.Currency = dbuf.String()
	m.TotalAmount = dbuf.Long()
	return dbuf.err
}

// updatePhoneCall#ab0f6b1e phone_call:PhoneCall = Update;
func (m *TLUpdatePhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePhoneCall))
	x.Bytes(m.PhoneCall.Encode())
	return x.buf
}

func (m *TLUpdatePhoneCall) Decode(dbuf *DecodeBuf) error {
	m.PhoneCall = &PhoneCall{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateLangPackTooLong#10c2404b = Update;
func (m *TLUpdateLangPackTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateLangPackTooLong))
	return x.buf
}

func (m *TLUpdateLangPackTooLong) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateLangPack#56022f4d difference:LangPackDifference = Update;
func (m *TLUpdateLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateLangPack))
	x.Bytes(m.Difference.Encode())
	return x.buf
}

func (m *TLUpdateLangPack) Decode(dbuf *DecodeBuf) error {
	m.Difference = &LangPackDifference{}
	m.Decode(dbuf)
	return dbuf.err
}

// updateFavedStickers#e511996d = Update;
func (m *TLUpdateFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateFavedStickers))
	return x.buf
}

func (m *TLUpdateFavedStickers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
func (m *TLUpdateChannelReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelReadMessagesContents))
	x.Int(m.ChannelId)
	x.VectorInt(m.Messages)
	return x.buf
}

func (m *TLUpdateChannelReadMessagesContents) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.Messages = dbuf.VectorInt()
	return dbuf.err
}

// updateContactsReset#7084a7be = Update;
func (m *TLUpdateContactsReset) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateContactsReset))
	return x.buf
}

func (m *TLUpdateContactsReset) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;
func (m *TLUpdatesState) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_state))
	x.Int(m.Pts)
	x.Int(m.Qts)
	x.Int(m.Date)
	x.Int(m.Seq)
	x.Int(m.UnreadCount)
	return x.buf
}

func (m *TLUpdatesState) Decode(dbuf *DecodeBuf) error {
	m.Pts = dbuf.Int()
	m.Qts = dbuf.Int()
	m.Date = dbuf.Int()
	m.Seq = dbuf.Int()
	m.UnreadCount = dbuf.Int()
	return dbuf.err
}

// updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
func (m *TLUpdatesDifferenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_differenceEmpty))
	x.Int(m.Date)
	x.Int(m.Seq)
	return x.buf
}

func (m *TLUpdatesDifferenceEmpty) Decode(dbuf *DecodeBuf) error {
	m.Date = dbuf.Int()
	m.Seq = dbuf.Int()
	return dbuf.err
}

// updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
func (m *TLUpdatesDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_difference))
	// x.VectorMessage(m.NewMessages);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.NewMessages)))
	for _, v := range m.NewMessages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.NewEncryptedMessages);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.NewEncryptedMessages)))
	for _, v := range m.NewEncryptedMessages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.OtherUpdates);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.OtherUpdates)))
	for _, v := range m.OtherUpdates {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.State.Encode())
	return x.buf
}

func (m *TLUpdatesDifference) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.NewMessages);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.NewMessages = make([]*Message, l1)
	for i := 0; i < int(l1); i++ {
		m.NewMessages[i] = &Message{}
		(*m.NewMessages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.NewEncryptedMessages);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.NewEncryptedMessages = make([]*EncryptedMessage, l2)
	for i := 0; i < int(l2); i++ {
		m.NewEncryptedMessages[i] = &EncryptedMessage{}
		(*m.NewEncryptedMessages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.OtherUpdates);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.OtherUpdates = make([]*Update, l3)
	for i := 0; i < int(l3); i++ {
		m.OtherUpdates[i] = &Update{}
		(*m.OtherUpdates[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Chats = make([]*Chat, l4)
	for i := 0; i < int(l4); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Users = make([]*User, l5)
	for i := 0; i < int(l5); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	m.State = &Updates_State{}
	m.Decode(dbuf)
	return dbuf.err
}

// updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;
func (m *TLUpdatesDifferenceSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_differenceSlice))
	// x.VectorMessage(m.NewMessages);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.NewMessages)))
	for _, v := range m.NewMessages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.NewEncryptedMessages);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.NewEncryptedMessages)))
	for _, v := range m.NewEncryptedMessages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.OtherUpdates);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.OtherUpdates)))
	for _, v := range m.OtherUpdates {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.IntermediateState.Encode())
	return x.buf
}

func (m *TLUpdatesDifferenceSlice) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.NewMessages);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.NewMessages = make([]*Message, l1)
	for i := 0; i < int(l1); i++ {
		m.NewMessages[i] = &Message{}
		(*m.NewMessages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.NewEncryptedMessages);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.NewEncryptedMessages = make([]*EncryptedMessage, l2)
	for i := 0; i < int(l2); i++ {
		m.NewEncryptedMessages[i] = &EncryptedMessage{}
		(*m.NewEncryptedMessages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.OtherUpdates);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.OtherUpdates = make([]*Update, l3)
	for i := 0; i < int(l3); i++ {
		m.OtherUpdates[i] = &Update{}
		(*m.OtherUpdates[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Chats = make([]*Chat, l4)
	for i := 0; i < int(l4); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Users = make([]*User, l5)
	for i := 0; i < int(l5); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	m.IntermediateState = &Updates_State{}
	m.Decode(dbuf)
	return dbuf.err
}

// updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;
func (m *TLUpdatesDifferenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_differenceTooLong))
	x.Int(m.Pts)
	return x.buf
}

func (m *TLUpdatesDifferenceTooLong) Decode(dbuf *DecodeBuf) error {
	m.Pts = dbuf.Int()
	return dbuf.err
}

// updatesTooLong#e317af7e = Updates;
func (m *TLUpdatesTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatesTooLong))
	return x.buf
}

func (m *TLUpdatesTooLong) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func (m *TLUpdateShortMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShortMessage))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Int(m.UserId)
	x.String(m.Message)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	x.Int(m.Date)
	x.Bytes(m.FwdFrom.Encode())
	x.Int(m.ViaBotId)
	x.Int(m.ReplyToMsgId)
	// x.VectorMessage(m.Entities);
	x15 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x15, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x15[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdateShortMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.UserId = dbuf.Int()
	m.Message = dbuf.String()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	m.Date = dbuf.Int()
	m.FwdFrom = &MessageFwdHeader{}
	m.Decode(dbuf)
	m.ViaBotId = dbuf.Int()
	m.ReplyToMsgId = dbuf.Int()
	// x.VectorMessage(m.Entities);
	c15 := dbuf.Int()
	if c15 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c15)
	}
	l15 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l15)
	for i := 0; i < int(l15); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func (m *TLUpdateShortChatMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShortChatMessage))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Int(m.FromId)
	x.Int(m.ChatId)
	x.String(m.Message)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	x.Int(m.Date)
	x.Bytes(m.FwdFrom.Encode())
	x.Int(m.ViaBotId)
	x.Int(m.ReplyToMsgId)
	// x.VectorMessage(m.Entities);
	x16 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x16, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x16[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdateShortChatMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.FromId = dbuf.Int()
	m.ChatId = dbuf.Int()
	m.Message = dbuf.String()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	m.Date = dbuf.Int()
	m.FwdFrom = &MessageFwdHeader{}
	m.Decode(dbuf)
	m.ViaBotId = dbuf.Int()
	m.ReplyToMsgId = dbuf.Int()
	// x.VectorMessage(m.Entities);
	c16 := dbuf.Int()
	if c16 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c16)
	}
	l16 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l16)
	for i := 0; i < int(l16); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updateShort#78d4dec1 update:Update date:int = Updates;
func (m *TLUpdateShort) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShort))
	x.Bytes(m.Update.Encode())
	x.Int(m.Date)
	return x.buf
}

func (m *TLUpdateShort) Decode(dbuf *DecodeBuf) error {
	m.Update = &Update{}
	m.Decode(dbuf)
	m.Date = dbuf.Int()
	return dbuf.err
}

// updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
func (m *TLUpdatesCombined) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatesCombined))
	// x.VectorMessage(m.Updates);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Updates)))
	for _, v := range m.Updates {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.Date)
	x.Int(m.SeqStart)
	x.Int(m.Seq)
	return x.buf
}

func (m *TLUpdatesCombined) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Updates);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Updates = make([]*Update, l1)
	for i := 0; i < int(l1); i++ {
		m.Updates[i] = &Update{}
		(*m.Updates[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Chats = make([]*Chat, l3)
	for i := 0; i < int(l3); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	m.Date = dbuf.Int()
	m.SeqStart = dbuf.Int()
	m.Seq = dbuf.Int()
	return dbuf.err
}

// updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
func (m *TLUpdates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates))
	// x.VectorMessage(m.Updates);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Updates)))
	for _, v := range m.Updates {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.Date)
	x.Int(m.Seq)
	return x.buf
}

func (m *TLUpdates) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Updates);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Updates = make([]*Update, l1)
	for i := 0; i < int(l1); i++ {
		m.Updates[i] = &Update{}
		(*m.Updates[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Chats = make([]*Chat, l3)
	for i := 0; i < int(l3); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	m.Date = dbuf.Int()
	m.Seq = dbuf.Int()
	return dbuf.err
}

// updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
func (m *TLUpdateShortSentMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShortSentMessage))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	x.Int(m.Date)
	x.Bytes(m.Media.Encode())
	// x.VectorMessage(m.Entities);
	x8 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x8, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x8[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdateShortSentMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	m.Date = dbuf.Int()
	m.Media = &MessageMedia{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Entities);
	c8 := dbuf.Int()
	if c8 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c8)
	}
	l8 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l8)
	for i := 0; i < int(l8); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;
func (m *TLPhotosPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_photos))
	// x.VectorMessage(m.Photos);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Photos)))
	for _, v := range m.Photos {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPhotosPhotos) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Photos);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Photos = make([]*Photo, l1)
	for i := 0; i < int(l1); i++ {
		m.Photos[i] = &Photo{}
		(*m.Photos[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;
func (m *TLPhotosPhotosSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_photosSlice))
	x.Int(m.Count)
	// x.VectorMessage(m.Photos);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Photos)))
	for _, v := range m.Photos {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPhotosPhotosSlice) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Photos);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Photos = make([]*Photo, l2)
	for i := 0; i < int(l2); i++ {
		m.Photos[i] = &Photo{}
		(*m.Photos[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;
func (m *TLPhotosPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_photo))
	x.Bytes(m.Photo.Encode())
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPhotosPhoto) Decode(dbuf *DecodeBuf) error {
	m.Photo = &Photo{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;
func (m *TLUploadFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_file))
	x.Bytes(m.Type.Encode())
	x.Int(m.Mtime)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLUploadFile) Decode(dbuf *DecodeBuf) error {
	m.Type = &Storage_FileType{}
	m.Decode(dbuf)
	m.Mtime = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// upload.fileCdnRedirect#ea52fe5a dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes cdn_file_hashes:Vector<CdnFileHash> = upload.File;
func (m *TLUploadFileCdnRedirect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_fileCdnRedirect))
	x.Int(m.DcId)
	x.StringBytes(m.FileToken)
	x.StringBytes(m.EncryptionKey)
	x.StringBytes(m.EncryptionIv)
	// x.VectorMessage(m.CdnFileHashes);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.CdnFileHashes)))
	for _, v := range m.CdnFileHashes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUploadFileCdnRedirect) Decode(dbuf *DecodeBuf) error {
	m.DcId = dbuf.Int()
	m.FileToken = dbuf.StringBytes()
	m.EncryptionKey = dbuf.StringBytes()
	m.EncryptionIv = dbuf.StringBytes()
	// x.VectorMessage(m.CdnFileHashes);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.CdnFileHashes = make([]*CdnFileHash, l5)
	for i := 0; i < int(l5); i++ {
		m.CdnFileHashes[i] = &CdnFileHash{}
		(*m.CdnFileHashes[i]).Decode(dbuf)
	}
	return dbuf.err
}

// dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
func (m *TLDcOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dcOption))
	x.Int(m.Flags)
	x.Int(m.Id)
	x.String(m.IpAddress)
	x.Int(m.Port)
	return x.buf
}

func (m *TLDcOption) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Int()
	m.IpAddress = dbuf.String()
	m.Port = dbuf.Int()
	return dbuf.err
}

// config#8df376a4 flags:# phonecalls_enabled:flags.1?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string suggested_lang_code:flags.2?string lang_pack_version:flags.2?int disabled_features:Vector<DisabledFeature> = Config;
func (m *TLConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_config))
	x.Int(m.Flags)
	x.Int(m.Date)
	x.Int(m.Expires)
	x.Bytes(m.TestMode.Encode())
	x.Int(m.ThisDc)
	// x.VectorMessage(m.DcOptions);
	x7 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x7, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x7[4:], uint32(len(m.DcOptions)))
	for _, v := range m.DcOptions {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.ChatSizeMax)
	x.Int(m.MegagroupSizeMax)
	x.Int(m.ForwardedCountMax)
	x.Int(m.OnlineUpdatePeriodMs)
	x.Int(m.OfflineBlurTimeoutMs)
	x.Int(m.OfflineIdleTimeoutMs)
	x.Int(m.OnlineCloudTimeoutMs)
	x.Int(m.NotifyCloudDelayMs)
	x.Int(m.NotifyDefaultDelayMs)
	x.Int(m.ChatBigSize)
	x.Int(m.PushChatPeriodMs)
	x.Int(m.PushChatLimit)
	x.Int(m.SavedGifsLimit)
	x.Int(m.EditTimeLimit)
	x.Int(m.RatingEDecay)
	x.Int(m.StickersRecentLimit)
	x.Int(m.StickersFavedLimit)
	x.Int(m.TmpSessions)
	x.Int(m.PinnedDialogsCountMax)
	x.Int(m.CallReceiveTimeoutMs)
	x.Int(m.CallRingTimeoutMs)
	x.Int(m.CallConnectTimeoutMs)
	x.Int(m.CallPacketTimeoutMs)
	x.String(m.MeUrlPrefix)
	x.String(m.SuggestedLangCode)
	x.Int(m.LangPackVersion)
	// x.VectorMessage(m.DisabledFeatures);
	x34 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x34, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x34[4:], uint32(len(m.DisabledFeatures)))
	for _, v := range m.DisabledFeatures {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLConfig) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Date = dbuf.Int()
	m.Expires = dbuf.Int()
	m.TestMode = &Bool{}
	m.Decode(dbuf)
	m.ThisDc = dbuf.Int()
	// x.VectorMessage(m.DcOptions);
	c7 := dbuf.Int()
	if c7 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c7)
	}
	l7 := dbuf.Int()
	m.DcOptions = make([]*DcOption, l7)
	for i := 0; i < int(l7); i++ {
		m.DcOptions[i] = &DcOption{}
		(*m.DcOptions[i]).Decode(dbuf)
	}
	m.ChatSizeMax = dbuf.Int()
	m.MegagroupSizeMax = dbuf.Int()
	m.ForwardedCountMax = dbuf.Int()
	m.OnlineUpdatePeriodMs = dbuf.Int()
	m.OfflineBlurTimeoutMs = dbuf.Int()
	m.OfflineIdleTimeoutMs = dbuf.Int()
	m.OnlineCloudTimeoutMs = dbuf.Int()
	m.NotifyCloudDelayMs = dbuf.Int()
	m.NotifyDefaultDelayMs = dbuf.Int()
	m.ChatBigSize = dbuf.Int()
	m.PushChatPeriodMs = dbuf.Int()
	m.PushChatLimit = dbuf.Int()
	m.SavedGifsLimit = dbuf.Int()
	m.EditTimeLimit = dbuf.Int()
	m.RatingEDecay = dbuf.Int()
	m.StickersRecentLimit = dbuf.Int()
	m.StickersFavedLimit = dbuf.Int()
	m.TmpSessions = dbuf.Int()
	m.PinnedDialogsCountMax = dbuf.Int()
	m.CallReceiveTimeoutMs = dbuf.Int()
	m.CallRingTimeoutMs = dbuf.Int()
	m.CallConnectTimeoutMs = dbuf.Int()
	m.CallPacketTimeoutMs = dbuf.Int()
	m.MeUrlPrefix = dbuf.String()
	m.SuggestedLangCode = dbuf.String()
	m.LangPackVersion = dbuf.Int()
	// x.VectorMessage(m.DisabledFeatures);
	c34 := dbuf.Int()
	if c34 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c34)
	}
	l34 := dbuf.Int()
	m.DisabledFeatures = make([]*DisabledFeature, l34)
	for i := 0; i < int(l34); i++ {
		m.DisabledFeatures[i] = &DisabledFeature{}
		(*m.DisabledFeatures[i]).Decode(dbuf)
	}
	return dbuf.err
}

// nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;
func (m *TLNearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_nearestDc))
	x.String(m.Country)
	x.Int(m.ThisDc)
	x.Int(m.NearestDc)
	return x.buf
}

func (m *TLNearestDc) Decode(dbuf *DecodeBuf) error {
	m.Country = dbuf.String()
	m.ThisDc = dbuf.Int()
	m.NearestDc = dbuf.Int()
	return dbuf.err
}

// help.appUpdate#8987f311 id:int critical:Bool url:string text:string = help.AppUpdate;
func (m *TLHelpAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_appUpdate))
	x.Int(m.Id)
	x.Bytes(m.Critical.Encode())
	x.String(m.Url)
	x.String(m.Text)
	return x.buf
}

func (m *TLHelpAppUpdate) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Critical = &Bool{}
	m.Decode(dbuf)
	m.Url = dbuf.String()
	m.Text = dbuf.String()
	return dbuf.err
}

// help.noAppUpdate#c45a6536 = help.AppUpdate;
func (m *TLHelpNoAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_noAppUpdate))
	return x.buf
}

func (m *TLHelpNoAppUpdate) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.inviteText#18cb9f78 message:string = help.InviteText;
func (m *TLHelpInviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_inviteText))
	x.String(m.Message)
	return x.buf
}

func (m *TLHelpInviteText) Decode(dbuf *DecodeBuf) error {
	m.Message = dbuf.String()
	return dbuf.err
}

// encryptedChatEmpty#ab7ec0a0 id:int = EncryptedChat;
func (m *TLEncryptedChatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatEmpty))
	x.Int(m.Id)
	return x.buf
}

func (m *TLEncryptedChatEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	return dbuf.err
}

// encryptedChatWaiting#3bf703dc id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;
func (m *TLEncryptedChatWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatWaiting))
	x.Int(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	return x.buf
}

func (m *TLEncryptedChatWaiting) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	return dbuf.err
}

// encryptedChatRequested#c878527e id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
func (m *TLEncryptedChatRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatRequested))
	x.Int(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	x.StringBytes(m.GA)
	return x.buf
}

func (m *TLEncryptedChatRequested) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	m.GA = dbuf.StringBytes()
	return dbuf.err
}

// encryptedChat#fa56ce36 id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;
func (m *TLEncryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChat))
	x.Int(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	x.StringBytes(m.GAOrB)
	x.Long(m.KeyFingerprint)
	return x.buf
}

func (m *TLEncryptedChat) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	m.GAOrB = dbuf.StringBytes()
	m.KeyFingerprint = dbuf.Long()
	return dbuf.err
}

// encryptedChatDiscarded#13d6dd27 id:int = EncryptedChat;
func (m *TLEncryptedChatDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatDiscarded))
	x.Int(m.Id)
	return x.buf
}

func (m *TLEncryptedChatDiscarded) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	return dbuf.err
}

// inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;
func (m *TLInputEncryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedChat))
	x.Int(m.ChatId)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputEncryptedChat) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// encryptedFileEmpty#c21f497e = EncryptedFile;
func (m *TLEncryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedFileEmpty))
	return x.buf
}

func (m *TLEncryptedFileEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// encryptedFile#4a70994c id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;
func (m *TLEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedFile))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Size)
	x.Int(m.DcId)
	x.Int(m.KeyFingerprint)
	return x.buf
}

func (m *TLEncryptedFile) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Size = dbuf.Int()
	m.DcId = dbuf.Int()
	m.KeyFingerprint = dbuf.Int()
	return dbuf.err
}

// inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;
func (m *TLInputEncryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileEmpty))
	return x.buf
}

func (m *TLInputEncryptedFileEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputEncryptedFileUploaded#64bd0306 id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
func (m *TLInputEncryptedFileUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileUploaded))
	x.Long(m.Id)
	x.Int(m.Parts)
	x.String(m.Md5Checksum)
	x.Int(m.KeyFingerprint)
	return x.buf
}

func (m *TLInputEncryptedFileUploaded) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Parts = dbuf.Int()
	m.Md5Checksum = dbuf.String()
	m.KeyFingerprint = dbuf.Int()
	return dbuf.err
}

// inputEncryptedFile#5a17b5e5 id:long access_hash:long = InputEncryptedFile;
func (m *TLInputEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFile))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputEncryptedFile) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputEncryptedFileBigUploaded#2dc173c8 id:long parts:int key_fingerprint:int = InputEncryptedFile;
func (m *TLInputEncryptedFileBigUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileBigUploaded))
	x.Long(m.Id)
	x.Int(m.Parts)
	x.Int(m.KeyFingerprint)
	return x.buf
}

func (m *TLInputEncryptedFileBigUploaded) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Parts = dbuf.Int()
	m.KeyFingerprint = dbuf.Int()
	return dbuf.err
}

// encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;
func (m *TLEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedMessage))
	x.Long(m.RandomId)
	x.Int(m.ChatId)
	x.Int(m.Date)
	x.StringBytes(m.Bytes)
	x.Bytes(m.File.Encode())
	return x.buf
}

func (m *TLEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	m.RandomId = dbuf.Long()
	m.ChatId = dbuf.Int()
	m.Date = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	m.File = &EncryptedFile{}
	m.Decode(dbuf)
	return dbuf.err
}

// encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;
func (m *TLEncryptedMessageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedMessageService))
	x.Long(m.RandomId)
	x.Int(m.ChatId)
	x.Int(m.Date)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLEncryptedMessageService) Decode(dbuf *DecodeBuf) error {
	m.RandomId = dbuf.Long()
	m.ChatId = dbuf.Int()
	m.Date = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;
func (m *TLMessagesDhConfigNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dhConfigNotModified))
	x.StringBytes(m.Random)
	return x.buf
}

func (m *TLMessagesDhConfigNotModified) Decode(dbuf *DecodeBuf) error {
	m.Random = dbuf.StringBytes()
	return dbuf.err
}

// messages.dhConfig#2c221edd g:int p:bytes version:int random:bytes = messages.DhConfig;
func (m *TLMessagesDhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dhConfig))
	x.Int(m.G)
	x.StringBytes(m.P)
	x.Int(m.Version)
	x.StringBytes(m.Random)
	return x.buf
}

func (m *TLMessagesDhConfig) Decode(dbuf *DecodeBuf) error {
	m.G = dbuf.Int()
	m.P = dbuf.StringBytes()
	m.Version = dbuf.Int()
	m.Random = dbuf.StringBytes()
	return dbuf.err
}

// messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;
func (m *TLMessagesSentEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sentEncryptedMessage))
	x.Int(m.Date)
	return x.buf
}

func (m *TLMessagesSentEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	m.Date = dbuf.Int()
	return dbuf.err
}

// messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;
func (m *TLMessagesSentEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sentEncryptedFile))
	x.Int(m.Date)
	x.Bytes(m.File.Encode())
	return x.buf
}

func (m *TLMessagesSentEncryptedFile) Decode(dbuf *DecodeBuf) error {
	m.Date = dbuf.Int()
	m.File = &EncryptedFile{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputDocumentEmpty#72f0eaae = InputDocument;
func (m *TLInputDocumentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocumentEmpty))
	return x.buf
}

func (m *TLInputDocumentEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputDocument#18798952 id:long access_hash:long = InputDocument;
func (m *TLInputDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocument))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputDocument) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// documentEmpty#36f8c871 id:long = Document;
func (m *TLDocumentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentEmpty))
	x.Long(m.Id)
	return x.buf
}

func (m *TLDocumentEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	return dbuf.err
}

// document#87232bc7 id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int version:int attributes:Vector<DocumentAttribute> = Document;
func (m *TLDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_document))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.String(m.MimeType)
	x.Int(m.Size)
	x.Bytes(m.Thumb.Encode())
	x.Int(m.DcId)
	x.Int(m.Version)
	// x.VectorMessage(m.Attributes);
	x9 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x9, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x9[4:], uint32(len(m.Attributes)))
	for _, v := range m.Attributes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLDocument) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.MimeType = dbuf.String()
	m.Size = dbuf.Int()
	m.Thumb = &PhotoSize{}
	m.Decode(dbuf)
	m.DcId = dbuf.Int()
	m.Version = dbuf.Int()
	// x.VectorMessage(m.Attributes);
	c9 := dbuf.Int()
	if c9 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c9)
	}
	l9 := dbuf.Int()
	m.Attributes = make([]*DocumentAttribute, l9)
	for i := 0; i < int(l9); i++ {
		m.Attributes[i] = &DocumentAttribute{}
		(*m.Attributes[i]).Decode(dbuf)
	}
	return dbuf.err
}

// help.support#17c6b5f6 phone_number:string user:User = help.Support;
func (m *TLHelpSupport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_support))
	x.String(m.PhoneNumber)
	x.Bytes(m.User.Encode())
	return x.buf
}

func (m *TLHelpSupport) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.User = &User{}
	m.Decode(dbuf)
	return dbuf.err
}

// notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;
func (m *TLNotifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyPeer))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLNotifyPeer) Decode(dbuf *DecodeBuf) error {
	m.Peer = &Peer{}
	m.Decode(dbuf)
	return dbuf.err
}

// notifyUsers#b4c83b4c = NotifyPeer;
func (m *TLNotifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyUsers))
	return x.buf
}

func (m *TLNotifyUsers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// notifyChats#c007cec3 = NotifyPeer;
func (m *TLNotifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyChats))
	return x.buf
}

func (m *TLNotifyChats) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// notifyAll#74d07c60 = NotifyPeer;
func (m *TLNotifyAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyAll))
	return x.buf
}

func (m *TLNotifyAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageTypingAction#16bf744e = SendMessageAction;
func (m *TLSendMessageTypingAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageTypingAction))
	return x.buf
}

func (m *TLSendMessageTypingAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageCancelAction#fd5ec8f5 = SendMessageAction;
func (m *TLSendMessageCancelAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageCancelAction))
	return x.buf
}

func (m *TLSendMessageCancelAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageRecordVideoAction#a187d66f = SendMessageAction;
func (m *TLSendMessageRecordVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageRecordVideoAction))
	return x.buf
}

func (m *TLSendMessageRecordVideoAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageUploadVideoAction#e9763aec progress:int = SendMessageAction;
func (m *TLSendMessageUploadVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadVideoAction))
	x.Int(m.Progress)
	return x.buf
}

func (m *TLSendMessageUploadVideoAction) Decode(dbuf *DecodeBuf) error {
	m.Progress = dbuf.Int()
	return dbuf.err
}

// sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;
func (m *TLSendMessageRecordAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageRecordAudioAction))
	return x.buf
}

func (m *TLSendMessageRecordAudioAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageUploadAudioAction#f351d7ab progress:int = SendMessageAction;
func (m *TLSendMessageUploadAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadAudioAction))
	x.Int(m.Progress)
	return x.buf
}

func (m *TLSendMessageUploadAudioAction) Decode(dbuf *DecodeBuf) error {
	m.Progress = dbuf.Int()
	return dbuf.err
}

// sendMessageUploadPhotoAction#d1d34a26 progress:int = SendMessageAction;
func (m *TLSendMessageUploadPhotoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadPhotoAction))
	x.Int(m.Progress)
	return x.buf
}

func (m *TLSendMessageUploadPhotoAction) Decode(dbuf *DecodeBuf) error {
	m.Progress = dbuf.Int()
	return dbuf.err
}

// sendMessageUploadDocumentAction#aa0cd9e4 progress:int = SendMessageAction;
func (m *TLSendMessageUploadDocumentAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadDocumentAction))
	x.Int(m.Progress)
	return x.buf
}

func (m *TLSendMessageUploadDocumentAction) Decode(dbuf *DecodeBuf) error {
	m.Progress = dbuf.Int()
	return dbuf.err
}

// sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;
func (m *TLSendMessageGeoLocationAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageGeoLocationAction))
	return x.buf
}

func (m *TLSendMessageGeoLocationAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageChooseContactAction#628cbc6f = SendMessageAction;
func (m *TLSendMessageChooseContactAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageChooseContactAction))
	return x.buf
}

func (m *TLSendMessageChooseContactAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageGamePlayAction#dd6a8f48 = SendMessageAction;
func (m *TLSendMessageGamePlayAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageGamePlayAction))
	return x.buf
}

func (m *TLSendMessageGamePlayAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageRecordRoundAction#88f27fbc = SendMessageAction;
func (m *TLSendMessageRecordRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageRecordRoundAction))
	return x.buf
}

func (m *TLSendMessageRecordRoundAction) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// sendMessageUploadRoundAction#243e1c66 progress:int = SendMessageAction;
func (m *TLSendMessageUploadRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadRoundAction))
	x.Int(m.Progress)
	return x.buf
}

func (m *TLSendMessageUploadRoundAction) Decode(dbuf *DecodeBuf) error {
	m.Progress = dbuf.Int()
	return dbuf.err
}

// contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
func (m *TLContactsFound) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_found))
	// x.VectorMessage(m.Results);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Results)))
	for _, v := range m.Results {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsFound) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Results);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Results = make([]*Peer, l1)
	for i := 0; i < int(l1); i++ {
		m.Results[i] = &Peer{}
		(*m.Results[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
func (m *TLInputPrivacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyKeyStatusTimestamp))
	return x.buf
}

func (m *TLInputPrivacyKeyStatusTimestamp) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
func (m *TLInputPrivacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyKeyChatInvite))
	return x.buf
}

func (m *TLInputPrivacyKeyChatInvite) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
func (m *TLInputPrivacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyKeyPhoneCall))
	return x.buf
}

func (m *TLInputPrivacyKeyPhoneCall) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
func (m *TLPrivacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyKeyStatusTimestamp))
	return x.buf
}

func (m *TLPrivacyKeyStatusTimestamp) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyKeyChatInvite#500e6dfa = PrivacyKey;
func (m *TLPrivacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyKeyChatInvite))
	return x.buf
}

func (m *TLPrivacyKeyChatInvite) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyKeyPhoneCall#3d662b7b = PrivacyKey;
func (m *TLPrivacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyKeyPhoneCall))
	return x.buf
}

func (m *TLPrivacyKeyPhoneCall) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
func (m *TLInputPrivacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueAllowContacts))
	return x.buf
}

func (m *TLInputPrivacyValueAllowContacts) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
func (m *TLInputPrivacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueAllowAll))
	return x.buf
}

func (m *TLInputPrivacyValueAllowAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
func (m *TLInputPrivacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueAllowUsers))
	// x.VectorMessage(m.Users);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLInputPrivacyValueAllowUsers) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Users);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Users = make([]*InputUser, l1)
	for i := 0; i < int(l1); i++ {
		m.Users[i] = &InputUser{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
func (m *TLInputPrivacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueDisallowContacts))
	return x.buf
}

func (m *TLInputPrivacyValueDisallowContacts) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
func (m *TLInputPrivacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueDisallowAll))
	return x.buf
}

func (m *TLInputPrivacyValueDisallowAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
func (m *TLInputPrivacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueDisallowUsers))
	// x.VectorMessage(m.Users);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLInputPrivacyValueDisallowUsers) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Users);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Users = make([]*InputUser, l1)
	for i := 0; i < int(l1); i++ {
		m.Users[i] = &InputUser{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// privacyValueAllowContacts#fffe1bac = PrivacyRule;
func (m *TLPrivacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueAllowContacts))
	return x.buf
}

func (m *TLPrivacyValueAllowContacts) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyValueAllowAll#65427b82 = PrivacyRule;
func (m *TLPrivacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueAllowAll))
	return x.buf
}

func (m *TLPrivacyValueAllowAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
func (m *TLPrivacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueAllowUsers))
	x.VectorInt(m.Users)
	return x.buf
}

func (m *TLPrivacyValueAllowUsers) Decode(dbuf *DecodeBuf) error {
	m.Users = dbuf.VectorInt()
	return dbuf.err
}

// privacyValueDisallowContacts#f888fa1a = PrivacyRule;
func (m *TLPrivacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueDisallowContacts))
	return x.buf
}

func (m *TLPrivacyValueDisallowContacts) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyValueDisallowAll#8b73e763 = PrivacyRule;
func (m *TLPrivacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueDisallowAll))
	return x.buf
}

func (m *TLPrivacyValueDisallowAll) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
func (m *TLPrivacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueDisallowUsers))
	x.VectorInt(m.Users)
	return x.buf
}

func (m *TLPrivacyValueDisallowUsers) Decode(dbuf *DecodeBuf) error {
	m.Users = dbuf.VectorInt()
	return dbuf.err
}

// account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;
func (m *TLAccountPrivacyRules) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_privacyRules))
	// x.VectorMessage(m.Rules);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Rules)))
	for _, v := range m.Rules {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLAccountPrivacyRules) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Rules);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Rules = make([]*PrivacyRule, l1)
	for i := 0; i < int(l1); i++ {
		m.Rules[i] = &PrivacyRule{}
		(*m.Rules[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;
func (m *TLAccountDaysTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_accountDaysTTL))
	x.Int(m.Days)
	return x.buf
}

func (m *TLAccountDaysTTL) Decode(dbuf *DecodeBuf) error {
	m.Days = dbuf.Int()
	return dbuf.err
}

// documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
func (m *TLDocumentAttributeImageSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeImageSize))
	x.Int(m.W)
	x.Int(m.H)
	return x.buf
}

func (m *TLDocumentAttributeImageSize) Decode(dbuf *DecodeBuf) error {
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	return dbuf.err
}

// documentAttributeAnimated#11b58939 = DocumentAttribute;
func (m *TLDocumentAttributeAnimated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeAnimated))
	return x.buf
}

func (m *TLDocumentAttributeAnimated) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;
func (m *TLDocumentAttributeSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeSticker))
	x.Int(m.Flags)
	x.String(m.Alt)
	x.Bytes(m.Stickerset.Encode())
	x.Bytes(m.MaskCoords.Encode())
	return x.buf
}

func (m *TLDocumentAttributeSticker) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Alt = dbuf.String()
	m.Stickerset = &InputStickerSet{}
	m.Decode(dbuf)
	m.MaskCoords = &MaskCoords{}
	m.Decode(dbuf)
	return dbuf.err
}

// documentAttributeVideo#ef02ce6 flags:# round_message:flags.0?true duration:int w:int h:int = DocumentAttribute;
func (m *TLDocumentAttributeVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeVideo))
	x.Int(m.Flags)
	x.Int(m.Duration)
	x.Int(m.W)
	x.Int(m.H)
	return x.buf
}

func (m *TLDocumentAttributeVideo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Duration = dbuf.Int()
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	return dbuf.err
}

// documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;
func (m *TLDocumentAttributeAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeAudio))
	x.Int(m.Flags)
	x.Int(m.Duration)
	x.String(m.Title)
	x.String(m.Performer)
	x.StringBytes(m.Waveform)
	return x.buf
}

func (m *TLDocumentAttributeAudio) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Duration = dbuf.Int()
	m.Title = dbuf.String()
	m.Performer = dbuf.String()
	m.Waveform = dbuf.StringBytes()
	return dbuf.err
}

// documentAttributeFilename#15590068 file_name:string = DocumentAttribute;
func (m *TLDocumentAttributeFilename) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeFilename))
	x.String(m.FileName)
	return x.buf
}

func (m *TLDocumentAttributeFilename) Decode(dbuf *DecodeBuf) error {
	m.FileName = dbuf.String()
	return dbuf.err
}

// documentAttributeHasStickers#9801d2f7 = DocumentAttribute;
func (m *TLDocumentAttributeHasStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeHasStickers))
	return x.buf
}

func (m *TLDocumentAttributeHasStickers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.stickersNotModified#f1749a22 = messages.Stickers;
func (m *TLMessagesStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickersNotModified))
	return x.buf
}

func (m *TLMessagesStickersNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.stickers#8a8ecd32 hash:string stickers:Vector<Document> = messages.Stickers;
func (m *TLMessagesStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickers))
	x.String(m.Hash)
	// x.VectorMessage(m.Stickers);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Stickers)))
	for _, v := range m.Stickers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.String()
	// x.VectorMessage(m.Stickers);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Stickers = make([]*Document, l2)
	for i := 0; i < int(l2); i++ {
		m.Stickers[i] = &Document{}
		(*m.Stickers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;
func (m *TLStickerPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerPack))
	x.String(m.Emoticon)
	x.VectorLong(m.Documents)
	return x.buf
}

func (m *TLStickerPack) Decode(dbuf *DecodeBuf) error {
	m.Emoticon = dbuf.String()
	m.Documents = dbuf.VectorLong()
	return dbuf.err
}

// messages.allStickersNotModified#e86602c3 = messages.AllStickers;
func (m *TLMessagesAllStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_allStickersNotModified))
	return x.buf
}

func (m *TLMessagesAllStickersNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.allStickers#edfd405f hash:int sets:Vector<StickerSet> = messages.AllStickers;
func (m *TLMessagesAllStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_allStickers))
	x.Int(m.Hash)
	// x.VectorMessage(m.Sets);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Sets)))
	for _, v := range m.Sets {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesAllStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	// x.VectorMessage(m.Sets);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Sets = make([]*StickerSet, l2)
	for i := 0; i < int(l2); i++ {
		m.Sets[i] = &StickerSet{}
		(*m.Sets[i]).Decode(dbuf)
	}
	return dbuf.err
}

// disabledFeature#ae636f24 feature:string description:string = DisabledFeature;
func (m *TLDisabledFeature) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_disabledFeature))
	x.String(m.Feature)
	x.String(m.Description)
	return x.buf
}

func (m *TLDisabledFeature) Decode(dbuf *DecodeBuf) error {
	m.Feature = dbuf.String()
	m.Description = dbuf.String()
	return dbuf.err
}

// messages.affectedMessages#84d19185 pts:int pts_count:int = messages.AffectedMessages;
func (m *TLMessagesAffectedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_affectedMessages))
	x.Int(m.Pts)
	x.Int(m.PtsCount)
	return x.buf
}

func (m *TLMessagesAffectedMessages) Decode(dbuf *DecodeBuf) error {
	m.Pts = dbuf.Int()
	m.PtsCount = dbuf.Int()
	return dbuf.err
}

// contactLinkUnknown#5f4f9247 = ContactLink;
func (m *TLContactLinkUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkUnknown))
	return x.buf
}

func (m *TLContactLinkUnknown) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// contactLinkNone#feedd3ad = ContactLink;
func (m *TLContactLinkNone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkNone))
	return x.buf
}

func (m *TLContactLinkNone) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// contactLinkHasPhone#268f3f59 = ContactLink;
func (m *TLContactLinkHasPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkHasPhone))
	return x.buf
}

func (m *TLContactLinkHasPhone) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// contactLinkContact#d502c2d0 = ContactLink;
func (m *TLContactLinkContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkContact))
	return x.buf
}

func (m *TLContactLinkContact) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// webPageEmpty#eb1477e8 id:long = WebPage;
func (m *TLWebPageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPageEmpty))
	x.Long(m.Id)
	return x.buf
}

func (m *TLWebPageEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	return dbuf.err
}

// webPagePending#c586da1c id:long date:int = WebPage;
func (m *TLWebPagePending) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPagePending))
	x.Long(m.Id)
	x.Int(m.Date)
	return x.buf
}

func (m *TLWebPagePending) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Date = dbuf.Int()
	return dbuf.err
}

// webPage#5f07b4bc flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;
func (m *TLWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPage))
	x.Int(m.Flags)
	x.Long(m.Id)
	x.String(m.Url)
	x.String(m.DisplayUrl)
	x.Int(m.Hash)
	x.String(m.Type)
	x.String(m.SiteName)
	x.String(m.Title)
	x.String(m.Description)
	x.Bytes(m.Photo.Encode())
	x.String(m.EmbedUrl)
	x.String(m.EmbedType)
	x.Int(m.EmbedWidth)
	x.Int(m.EmbedHeight)
	x.Int(m.Duration)
	x.String(m.Author)
	x.Bytes(m.Document.Encode())
	x.Bytes(m.CachedPage.Encode())
	return x.buf
}

func (m *TLWebPage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Long()
	m.Url = dbuf.String()
	m.DisplayUrl = dbuf.String()
	m.Hash = dbuf.Int()
	m.Type = dbuf.String()
	m.SiteName = dbuf.String()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Photo = &Photo{}
	m.Decode(dbuf)
	m.EmbedUrl = dbuf.String()
	m.EmbedType = dbuf.String()
	m.EmbedWidth = dbuf.Int()
	m.EmbedHeight = dbuf.Int()
	m.Duration = dbuf.Int()
	m.Author = dbuf.String()
	m.Document = &Document{}
	m.Decode(dbuf)
	m.CachedPage = &Page{}
	m.Decode(dbuf)
	return dbuf.err
}

// webPageNotModified#85849473 = WebPage;
func (m *TLWebPageNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPageNotModified))
	return x.buf
}

func (m *TLWebPageNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// authorization#7bf2e6f6 hash:long flags:int device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
func (m *TLAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_authorization))
	x.Long(m.Hash)
	x.Int(m.Flags)
	x.String(m.DeviceModel)
	x.String(m.Platform)
	x.String(m.SystemVersion)
	x.Int(m.ApiId)
	x.String(m.AppName)
	x.String(m.AppVersion)
	x.Int(m.DateCreated)
	x.Int(m.DateActive)
	x.String(m.Ip)
	x.String(m.Country)
	x.String(m.Region)
	return x.buf
}

func (m *TLAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Long()
	m.Flags = dbuf.Int()
	m.DeviceModel = dbuf.String()
	m.Platform = dbuf.String()
	m.SystemVersion = dbuf.String()
	m.ApiId = dbuf.Int()
	m.AppName = dbuf.String()
	m.AppVersion = dbuf.String()
	m.DateCreated = dbuf.Int()
	m.DateActive = dbuf.Int()
	m.Ip = dbuf.String()
	m.Country = dbuf.String()
	m.Region = dbuf.String()
	return dbuf.err
}

// account.authorizations#1250abde authorizations:Vector<Authorization> = account.Authorizations;
func (m *TLAccountAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_authorizations))
	// x.VectorMessage(m.Authorizations);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Authorizations)))
	for _, v := range m.Authorizations {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLAccountAuthorizations) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Authorizations);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Authorizations = make([]*Authorization, l1)
	for i := 0; i < int(l1); i++ {
		m.Authorizations[i] = &Authorization{}
		(*m.Authorizations[i]).Decode(dbuf)
	}
	return dbuf.err
}

// account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
func (m *TLAccountNoPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_noPassword))
	x.StringBytes(m.NewSalt)
	x.String(m.EmailUnconfirmedPattern)
	return x.buf
}

func (m *TLAccountNoPassword) Decode(dbuf *DecodeBuf) error {
	m.NewSalt = dbuf.StringBytes()
	m.EmailUnconfirmedPattern = dbuf.String()
	return dbuf.err
}

// account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
func (m *TLAccountPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_password))
	x.StringBytes(m.CurrentSalt)
	x.StringBytes(m.NewSalt)
	x.String(m.Hint)
	x.Bytes(m.HasRecovery.Encode())
	x.String(m.EmailUnconfirmedPattern)
	return x.buf
}

func (m *TLAccountPassword) Decode(dbuf *DecodeBuf) error {
	m.CurrentSalt = dbuf.StringBytes()
	m.NewSalt = dbuf.StringBytes()
	m.Hint = dbuf.String()
	m.HasRecovery = &Bool{}
	m.Decode(dbuf)
	m.EmailUnconfirmedPattern = dbuf.String()
	return dbuf.err
}

// account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;
func (m *TLAccountPasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_passwordSettings))
	x.String(m.Email)
	return x.buf
}

func (m *TLAccountPasswordSettings) Decode(dbuf *DecodeBuf) error {
	m.Email = dbuf.String()
	return dbuf.err
}

// account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;
func (m *TLAccountPasswordInputSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_passwordInputSettings))
	x.Int(m.Flags)
	x.StringBytes(m.NewSalt)
	x.StringBytes(m.NewPasswordHash)
	x.String(m.Hint)
	x.String(m.Email)
	return x.buf
}

func (m *TLAccountPasswordInputSettings) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.NewSalt = dbuf.StringBytes()
	m.NewPasswordHash = dbuf.StringBytes()
	m.Hint = dbuf.String()
	m.Email = dbuf.String()
	return dbuf.err
}

// auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;
func (m *TLAuthPasswordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_passwordRecovery))
	x.String(m.EmailPattern)
	return x.buf
}

func (m *TLAuthPasswordRecovery) Decode(dbuf *DecodeBuf) error {
	m.EmailPattern = dbuf.String()
	return dbuf.err
}

// receivedNotifyMessage#a384b779 id:int flags:int = ReceivedNotifyMessage;
func (m *TLReceivedNotifyMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_receivedNotifyMessage))
	x.Int(m.Id)
	x.Int(m.Flags)
	return x.buf
}

func (m *TLReceivedNotifyMessage) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Flags = dbuf.Int()
	return dbuf.err
}

// chatInviteEmpty#69df3769 = ExportedChatInvite;
func (m *TLChatInviteEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInviteEmpty))
	return x.buf
}

func (m *TLChatInviteEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// chatInviteExported#fc2e05bc link:string = ExportedChatInvite;
func (m *TLChatInviteExported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInviteExported))
	x.String(m.Link)
	return x.buf
}

func (m *TLChatInviteExported) Decode(dbuf *DecodeBuf) error {
	m.Link = dbuf.String()
	return dbuf.err
}

// chatInviteAlready#5a686d7c chat:Chat = ChatInvite;
func (m *TLChatInviteAlready) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInviteAlready))
	x.Bytes(m.Chat.Encode())
	return x.buf
}

func (m *TLChatInviteAlready) Decode(dbuf *DecodeBuf) error {
	m.Chat = &Chat{}
	m.Decode(dbuf)
	return dbuf.err
}

// chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;
func (m *TLChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInvite))
	x.Int(m.Flags)
	x.String(m.Title)
	x.Bytes(m.Photo.Encode())
	x.Int(m.ParticipantsCount)
	// x.VectorMessage(m.Participants);
	x9 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x9, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x9[4:], uint32(len(m.Participants)))
	for _, v := range m.Participants {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChatInvite) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Title = dbuf.String()
	m.Photo = &ChatPhoto{}
	m.Decode(dbuf)
	m.ParticipantsCount = dbuf.Int()
	// x.VectorMessage(m.Participants);
	c9 := dbuf.Int()
	if c9 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c9)
	}
	l9 := dbuf.Int()
	m.Participants = make([]*User, l9)
	for i := 0; i < int(l9); i++ {
		m.Participants[i] = &User{}
		(*m.Participants[i]).Decode(dbuf)
	}
	return dbuf.err
}

// inputStickerSetEmpty#ffb62b95 = InputStickerSet;
func (m *TLInputStickerSetEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetEmpty))
	return x.buf
}

func (m *TLInputStickerSetEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputStickerSetID#9de7a269 id:long access_hash:long = InputStickerSet;
func (m *TLInputStickerSetID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetID))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputStickerSetID) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputStickerSetShortName#861cc8a0 short_name:string = InputStickerSet;
func (m *TLInputStickerSetShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetShortName))
	x.String(m.ShortName)
	return x.buf
}

func (m *TLInputStickerSetShortName) Decode(dbuf *DecodeBuf) error {
	m.ShortName = dbuf.String()
	return dbuf.err
}

// stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
func (m *TLStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerSet))
	x.Int(m.Flags)
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.String(m.Title)
	x.String(m.ShortName)
	x.Int(m.Count)
	x.Int(m.Hash)
	return x.buf
}

func (m *TLStickerSet) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Title = dbuf.String()
	m.ShortName = dbuf.String()
	m.Count = dbuf.Int()
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.stickerSet#b60a24a6 set:StickerSet packs:Vector<StickerPack> documents:Vector<Document> = messages.StickerSet;
func (m *TLMessagesStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickerSet))
	x.Bytes(m.Set.Encode())
	// x.VectorMessage(m.Packs);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Packs)))
	for _, v := range m.Packs {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Documents);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Documents)))
	for _, v := range m.Documents {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesStickerSet) Decode(dbuf *DecodeBuf) error {
	m.Set = &StickerSet{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Packs);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Packs = make([]*StickerPack, l2)
	for i := 0; i < int(l2); i++ {
		m.Packs[i] = &StickerPack{}
		(*m.Packs[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Documents);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Documents = make([]*Document, l3)
	for i := 0; i < int(l3); i++ {
		m.Documents[i] = &Document{}
		(*m.Documents[i]).Decode(dbuf)
	}
	return dbuf.err
}

// botCommand#c27ac8c7 command:string description:string = BotCommand;
func (m *TLBotCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botCommand))
	x.String(m.Command)
	x.String(m.Description)
	return x.buf
}

func (m *TLBotCommand) Decode(dbuf *DecodeBuf) error {
	m.Command = dbuf.String()
	m.Description = dbuf.String()
	return dbuf.err
}

// botInfo#98e81d3a user_id:int description:string commands:Vector<BotCommand> = BotInfo;
func (m *TLBotInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInfo))
	x.Int(m.UserId)
	x.String(m.Description)
	// x.VectorMessage(m.Commands);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Commands)))
	for _, v := range m.Commands {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLBotInfo) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Description = dbuf.String()
	// x.VectorMessage(m.Commands);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Commands = make([]*BotCommand, l3)
	for i := 0; i < int(l3); i++ {
		m.Commands[i] = &BotCommand{}
		(*m.Commands[i]).Decode(dbuf)
	}
	return dbuf.err
}

// keyboardButton#a2fa4880 text:string = KeyboardButton;
func (m *TLKeyboardButton) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButton))
	x.String(m.Text)
	return x.buf
}

func (m *TLKeyboardButton) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// keyboardButtonUrl#258aff05 text:string url:string = KeyboardButton;
func (m *TLKeyboardButtonUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonUrl))
	x.String(m.Text)
	x.String(m.Url)
	return x.buf
}

func (m *TLKeyboardButtonUrl) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	m.Url = dbuf.String()
	return dbuf.err
}

// keyboardButtonCallback#683a5e46 text:string data:bytes = KeyboardButton;
func (m *TLKeyboardButtonCallback) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonCallback))
	x.String(m.Text)
	x.StringBytes(m.Data)
	return x.buf
}

func (m *TLKeyboardButtonCallback) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	m.Data = dbuf.StringBytes()
	return dbuf.err
}

// keyboardButtonRequestPhone#b16a6c29 text:string = KeyboardButton;
func (m *TLKeyboardButtonRequestPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonRequestPhone))
	x.String(m.Text)
	return x.buf
}

func (m *TLKeyboardButtonRequestPhone) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// keyboardButtonRequestGeoLocation#fc796b3f text:string = KeyboardButton;
func (m *TLKeyboardButtonRequestGeoLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonRequestGeoLocation))
	x.String(m.Text)
	return x.buf
}

func (m *TLKeyboardButtonRequestGeoLocation) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// keyboardButtonSwitchInline#568a748 flags:# same_peer:flags.0?true text:string query:string = KeyboardButton;
func (m *TLKeyboardButtonSwitchInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonSwitchInline))
	x.Int(m.Flags)
	x.String(m.Text)
	x.String(m.Query)
	return x.buf
}

func (m *TLKeyboardButtonSwitchInline) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Text = dbuf.String()
	m.Query = dbuf.String()
	return dbuf.err
}

// keyboardButtonGame#50f41ccf text:string = KeyboardButton;
func (m *TLKeyboardButtonGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonGame))
	x.String(m.Text)
	return x.buf
}

func (m *TLKeyboardButtonGame) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// keyboardButtonBuy#afd93fbb text:string = KeyboardButton;
func (m *TLKeyboardButtonBuy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonBuy))
	x.String(m.Text)
	return x.buf
}

func (m *TLKeyboardButtonBuy) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// keyboardButtonRow#77608b83 buttons:Vector<KeyboardButton> = KeyboardButtonRow;
func (m *TLKeyboardButtonRow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonRow))
	// x.VectorMessage(m.Buttons);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Buttons)))
	for _, v := range m.Buttons {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLKeyboardButtonRow) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Buttons);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Buttons = make([]*KeyboardButton, l1)
	for i := 0; i < int(l1); i++ {
		m.Buttons[i] = &KeyboardButton{}
		(*m.Buttons[i]).Decode(dbuf)
	}
	return dbuf.err
}

// replyKeyboardHide#a03e5b85 flags:# selective:flags.2?true = ReplyMarkup;
func (m *TLReplyKeyboardHide) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyKeyboardHide))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLReplyKeyboardHide) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// replyKeyboardForceReply#f4108aa0 flags:# single_use:flags.1?true selective:flags.2?true = ReplyMarkup;
func (m *TLReplyKeyboardForceReply) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyKeyboardForceReply))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLReplyKeyboardForceReply) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// replyKeyboardMarkup#3502758c flags:# resize:flags.0?true single_use:flags.1?true selective:flags.2?true rows:Vector<KeyboardButtonRow> = ReplyMarkup;
func (m *TLReplyKeyboardMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyKeyboardMarkup))
	x.Int(m.Flags)
	// x.VectorMessage(m.Rows);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Rows)))
	for _, v := range m.Rows {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLReplyKeyboardMarkup) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.VectorMessage(m.Rows);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Rows = make([]*KeyboardButtonRow, l5)
	for i := 0; i < int(l5); i++ {
		m.Rows[i] = &KeyboardButtonRow{}
		(*m.Rows[i]).Decode(dbuf)
	}
	return dbuf.err
}

// replyInlineMarkup#48a30254 rows:Vector<KeyboardButtonRow> = ReplyMarkup;
func (m *TLReplyInlineMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyInlineMarkup))
	// x.VectorMessage(m.Rows);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Rows)))
	for _, v := range m.Rows {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLReplyInlineMarkup) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Rows);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Rows = make([]*KeyboardButtonRow, l1)
	for i := 0; i < int(l1); i++ {
		m.Rows[i] = &KeyboardButtonRow{}
		(*m.Rows[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;
func (m *TLMessageEntityUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityUnknown))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityUnknown) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityMention#fa04579d offset:int length:int = MessageEntity;
func (m *TLMessageEntityMention) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityMention))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityMention) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;
func (m *TLMessageEntityHashtag) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityHashtag))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityHashtag) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;
func (m *TLMessageEntityBotCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityBotCommand))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityBotCommand) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;
func (m *TLMessageEntityUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityUrl))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityUrl) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;
func (m *TLMessageEntityEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityEmail))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityEmail) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;
func (m *TLMessageEntityBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityBold))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityBold) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;
func (m *TLMessageEntityItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityItalic))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityItalic) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityCode#28a20571 offset:int length:int = MessageEntity;
func (m *TLMessageEntityCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityCode))
	x.Int(m.Offset)
	x.Int(m.Length)
	return x.buf
}

func (m *TLMessageEntityCode) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	return dbuf.err
}

// messageEntityPre#73924be0 offset:int length:int language:string = MessageEntity;
func (m *TLMessageEntityPre) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityPre))
	x.Int(m.Offset)
	x.Int(m.Length)
	x.String(m.Language)
	return x.buf
}

func (m *TLMessageEntityPre) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	m.Language = dbuf.String()
	return dbuf.err
}

// messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;
func (m *TLMessageEntityTextUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityTextUrl))
	x.Int(m.Offset)
	x.Int(m.Length)
	x.String(m.Url)
	return x.buf
}

func (m *TLMessageEntityTextUrl) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	m.Url = dbuf.String()
	return dbuf.err
}

// messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;
func (m *TLMessageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityMentionName))
	x.Int(m.Offset)
	x.Int(m.Length)
	x.Int(m.UserId)
	return x.buf
}

func (m *TLMessageEntityMentionName) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	m.UserId = dbuf.Int()
	return dbuf.err
}

// inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;
func (m *TLInputMessageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessageEntityMentionName))
	x.Int(m.Offset)
	x.Int(m.Length)
	x.Bytes(m.UserId.Encode())
	return x.buf
}

func (m *TLInputMessageEntityMentionName) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Length = dbuf.Int()
	m.UserId = &InputUser{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputChannelEmpty#ee8c1e86 = InputChannel;
func (m *TLInputChannelEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChannelEmpty))
	return x.buf
}

func (m *TLInputChannelEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// inputChannel#afeb712e channel_id:int access_hash:long = InputChannel;
func (m *TLInputChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChannel))
	x.Int(m.ChannelId)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputChannel) Decode(dbuf *DecodeBuf) error {
	m.ChannelId = dbuf.Int()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// contacts.resolvedPeer#7f077ad9 peer:Peer chats:Vector<Chat> users:Vector<User> = contacts.ResolvedPeer;
func (m *TLContactsResolvedPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_resolvedPeer))
	x.Bytes(m.Peer.Encode())
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsResolvedPeer) Decode(dbuf *DecodeBuf) error {
	m.Peer = &Peer{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messageRange#ae30253 min_id:int max_id:int = MessageRange;
func (m *TLMessageRange) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageRange))
	x.Int(m.MinId)
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLMessageRange) Decode(dbuf *DecodeBuf) error {
	m.MinId = dbuf.Int()
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// updates.channelDifferenceEmpty#3e11affb flags:# final:flags.0?true pts:int timeout:flags.1?int = updates.ChannelDifference;
func (m *TLUpdatesChannelDifferenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_channelDifferenceEmpty))
	x.Int(m.Flags)
	x.Int(m.Pts)
	x.Int(m.Timeout)
	return x.buf
}

func (m *TLUpdatesChannelDifferenceEmpty) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Pts = dbuf.Int()
	m.Timeout = dbuf.Int()
	return dbuf.err
}

// updates.channelDifferenceTooLong#6a9d7b35 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
func (m *TLUpdatesChannelDifferenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_channelDifferenceTooLong))
	x.Int(m.Flags)
	x.Int(m.Pts)
	x.Int(m.Timeout)
	x.Int(m.TopMessage)
	x.Int(m.ReadInboxMaxId)
	x.Int(m.ReadOutboxMaxId)
	x.Int(m.UnreadCount)
	x.Int(m.UnreadMentionsCount)
	// x.VectorMessage(m.Messages);
	x10 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x10, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x10[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x11 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x11, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x11[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x12 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x12, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x12[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdatesChannelDifferenceTooLong) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Pts = dbuf.Int()
	m.Timeout = dbuf.Int()
	m.TopMessage = dbuf.Int()
	m.ReadInboxMaxId = dbuf.Int()
	m.ReadOutboxMaxId = dbuf.Int()
	m.UnreadCount = dbuf.Int()
	m.UnreadMentionsCount = dbuf.Int()
	// x.VectorMessage(m.Messages);
	c10 := dbuf.Int()
	if c10 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c10)
	}
	l10 := dbuf.Int()
	m.Messages = make([]*Message, l10)
	for i := 0; i < int(l10); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c11 := dbuf.Int()
	if c11 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c11)
	}
	l11 := dbuf.Int()
	m.Chats = make([]*Chat, l11)
	for i := 0; i < int(l11); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c12 := dbuf.Int()
	if c12 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c12)
	}
	l12 := dbuf.Int()
	m.Users = make([]*User, l12)
	for i := 0; i < int(l12); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
func (m *TLUpdatesChannelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_channelDifference))
	x.Int(m.Flags)
	x.Int(m.Pts)
	x.Int(m.Timeout)
	// x.VectorMessage(m.NewMessages);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.NewMessages)))
	for _, v := range m.NewMessages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.OtherUpdates);
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.OtherUpdates)))
	for _, v := range m.OtherUpdates {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x7 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x7, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x7[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x8 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x8, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x8[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUpdatesChannelDifference) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Pts = dbuf.Int()
	m.Timeout = dbuf.Int()
	// x.VectorMessage(m.NewMessages);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.NewMessages = make([]*Message, l5)
	for i := 0; i < int(l5); i++ {
		m.NewMessages[i] = &Message{}
		(*m.NewMessages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.OtherUpdates);
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.OtherUpdates = make([]*Update, l6)
	for i := 0; i < int(l6); i++ {
		m.OtherUpdates[i] = &Update{}
		(*m.OtherUpdates[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c7 := dbuf.Int()
	if c7 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c7)
	}
	l7 := dbuf.Int()
	m.Chats = make([]*Chat, l7)
	for i := 0; i < int(l7); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c8 := dbuf.Int()
	if c8 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c8)
	}
	l8 := dbuf.Int()
	m.Users = make([]*User, l8)
	for i := 0; i < int(l8); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channelMessagesFilterEmpty#94d42ee7 = ChannelMessagesFilter;
func (m *TLChannelMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelMessagesFilterEmpty))
	return x.buf
}

func (m *TLChannelMessagesFilterEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channelMessagesFilter#cd77d957 flags:# exclude_new_messages:flags.1?true ranges:Vector<MessageRange> = ChannelMessagesFilter;
func (m *TLChannelMessagesFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelMessagesFilter))
	x.Int(m.Flags)
	// x.VectorMessage(m.Ranges);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Ranges)))
	for _, v := range m.Ranges {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChannelMessagesFilter) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.VectorMessage(m.Ranges);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Ranges = make([]*MessageRange, l3)
	for i := 0; i < int(l3); i++ {
		m.Ranges[i] = &MessageRange{}
		(*m.Ranges[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channelParticipant#15ebac1d user_id:int date:int = ChannelParticipant;
func (m *TLChannelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipant))
	x.Int(m.UserId)
	x.Int(m.Date)
	return x.buf
}

func (m *TLChannelParticipant) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// channelParticipantSelf#a3289a6d user_id:int inviter_id:int date:int = ChannelParticipant;
func (m *TLChannelParticipantSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantSelf))
	x.Int(m.UserId)
	x.Int(m.InviterId)
	x.Int(m.Date)
	return x.buf
}

func (m *TLChannelParticipantSelf) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	m.InviterId = dbuf.Int()
	m.Date = dbuf.Int()
	return dbuf.err
}

// channelParticipantCreator#e3e2e1f9 user_id:int = ChannelParticipant;
func (m *TLChannelParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantCreator))
	x.Int(m.UserId)
	return x.buf
}

func (m *TLChannelParticipantCreator) Decode(dbuf *DecodeBuf) error {
	m.UserId = dbuf.Int()
	return dbuf.err
}

// channelParticipantAdmin#a82fa898 flags:# can_edit:flags.0?true user_id:int inviter_id:int promoted_by:int date:int admin_rights:ChannelAdminRights = ChannelParticipant;
func (m *TLChannelParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantAdmin))
	x.Int(m.Flags)
	x.Int(m.UserId)
	x.Int(m.InviterId)
	x.Int(m.PromotedBy)
	x.Int(m.Date)
	x.Bytes(m.AdminRights.Encode())
	return x.buf
}

func (m *TLChannelParticipantAdmin) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.UserId = dbuf.Int()
	m.InviterId = dbuf.Int()
	m.PromotedBy = dbuf.Int()
	m.Date = dbuf.Int()
	m.AdminRights = &ChannelAdminRights{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelParticipantBanned#222c1886 flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChannelBannedRights = ChannelParticipant;
func (m *TLChannelParticipantBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantBanned))
	x.Int(m.Flags)
	x.Int(m.UserId)
	x.Int(m.KickedBy)
	x.Int(m.Date)
	x.Bytes(m.BannedRights.Encode())
	return x.buf
}

func (m *TLChannelParticipantBanned) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.UserId = dbuf.Int()
	m.KickedBy = dbuf.Int()
	m.Date = dbuf.Int()
	m.BannedRights = &ChannelBannedRights{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelParticipantsRecent#de3f3c79 = ChannelParticipantsFilter;
func (m *TLChannelParticipantsRecent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsRecent))
	return x.buf
}

func (m *TLChannelParticipantsRecent) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channelParticipantsAdmins#b4608969 = ChannelParticipantsFilter;
func (m *TLChannelParticipantsAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsAdmins))
	return x.buf
}

func (m *TLChannelParticipantsAdmins) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channelParticipantsKicked#a3b54985 q:string = ChannelParticipantsFilter;
func (m *TLChannelParticipantsKicked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsKicked))
	x.String(m.Q)
	return x.buf
}

func (m *TLChannelParticipantsKicked) Decode(dbuf *DecodeBuf) error {
	m.Q = dbuf.String()
	return dbuf.err
}

// channelParticipantsBots#b0d1865b = ChannelParticipantsFilter;
func (m *TLChannelParticipantsBots) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsBots))
	return x.buf
}

func (m *TLChannelParticipantsBots) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channelParticipantsBanned#1427a5e1 q:string = ChannelParticipantsFilter;
func (m *TLChannelParticipantsBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsBanned))
	x.String(m.Q)
	return x.buf
}

func (m *TLChannelParticipantsBanned) Decode(dbuf *DecodeBuf) error {
	m.Q = dbuf.String()
	return dbuf.err
}

// channelParticipantsSearch#656ac4b q:string = ChannelParticipantsFilter;
func (m *TLChannelParticipantsSearch) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsSearch))
	x.String(m.Q)
	return x.buf
}

func (m *TLChannelParticipantsSearch) Decode(dbuf *DecodeBuf) error {
	m.Q = dbuf.String()
	return dbuf.err
}

// channels.channelParticipants#f56ee2a8 count:int participants:Vector<ChannelParticipant> users:Vector<User> = channels.ChannelParticipants;
func (m *TLChannelsChannelParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_channelParticipants))
	x.Int(m.Count)
	// x.VectorMessage(m.Participants);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Participants)))
	for _, v := range m.Participants {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChannelsChannelParticipants) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Participants);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Participants = make([]*ChannelParticipant, l2)
	for i := 0; i < int(l2); i++ {
		m.Participants[i] = &ChannelParticipant{}
		(*m.Participants[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channels.channelParticipant#d0d9b163 participant:ChannelParticipant users:Vector<User> = channels.ChannelParticipant;
func (m *TLChannelsChannelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_channelParticipant))
	x.Bytes(m.Participant.Encode())
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChannelsChannelParticipant) Decode(dbuf *DecodeBuf) error {
	m.Participant = &ChannelParticipant{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// help.termsOfService#f1ee3e90 text:string = help.TermsOfService;
func (m *TLHelpTermsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_termsOfService))
	x.String(m.Text)
	return x.buf
}

func (m *TLHelpTermsOfService) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// foundGif#162ecc1f url:string thumb_url:string content_url:string content_type:string w:int h:int = FoundGif;
func (m *TLFoundGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_foundGif))
	x.String(m.Url)
	x.String(m.ThumbUrl)
	x.String(m.ContentUrl)
	x.String(m.ContentType)
	x.Int(m.W)
	x.Int(m.H)
	return x.buf
}

func (m *TLFoundGif) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.ThumbUrl = dbuf.String()
	m.ContentUrl = dbuf.String()
	m.ContentType = dbuf.String()
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	return dbuf.err
}

// foundGifCached#9c750409 url:string photo:Photo document:Document = FoundGif;
func (m *TLFoundGifCached) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_foundGifCached))
	x.String(m.Url)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.Document.Encode())
	return x.buf
}

func (m *TLFoundGifCached) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.Photo = &Photo{}
	m.Decode(dbuf)
	m.Document = &Document{}
	m.Decode(dbuf)
	return dbuf.err
}

// messages.foundGifs#450a1c0a next_offset:int results:Vector<FoundGif> = messages.FoundGifs;
func (m *TLMessagesFoundGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_foundGifs))
	x.Int(m.NextOffset)
	// x.VectorMessage(m.Results);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Results)))
	for _, v := range m.Results {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesFoundGifs) Decode(dbuf *DecodeBuf) error {
	m.NextOffset = dbuf.Int()
	// x.VectorMessage(m.Results);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Results = make([]*FoundGif, l2)
	for i := 0; i < int(l2); i++ {
		m.Results[i] = &FoundGif{}
		(*m.Results[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.savedGifsNotModified#e8025ca2 = messages.SavedGifs;
func (m *TLMessagesSavedGifsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_savedGifsNotModified))
	return x.buf
}

func (m *TLMessagesSavedGifsNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.savedGifs#2e0709a5 hash:int gifs:Vector<Document> = messages.SavedGifs;
func (m *TLMessagesSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_savedGifs))
	x.Int(m.Hash)
	// x.VectorMessage(m.Gifs);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Gifs)))
	for _, v := range m.Gifs {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesSavedGifs) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	// x.VectorMessage(m.Gifs);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Gifs = make([]*Document, l2)
	for i := 0; i < int(l2); i++ {
		m.Gifs[i] = &Document{}
		(*m.Gifs[i]).Decode(dbuf)
	}
	return dbuf.err
}

// inputBotInlineMessageMediaAuto#292fed13 flags:# caption:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaAuto))
	x.Int(m.Flags)
	x.String(m.Caption)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLInputBotInlineMessageMediaAuto) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Caption = dbuf.String()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineMessageText#3dcd7a87 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageText))
	x.Int(m.Flags)
	x.String(m.Message)
	// x.VectorMessage(m.Entities);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLInputBotInlineMessageText) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Message = dbuf.String()
	// x.VectorMessage(m.Entities);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l4)
	for i := 0; i < int(l4); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineMessageMediaGeo#f4a59de1 flags:# geo_point:InputGeoPoint reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaGeo))
	x.Int(m.Flags)
	x.Bytes(m.GeoPoint.Encode())
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLInputBotInlineMessageMediaGeo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.GeoPoint = &InputGeoPoint{}
	m.Decode(dbuf)
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineMessageMediaVenue#aaafadc8 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaVenue))
	x.Int(m.Flags)
	x.Bytes(m.GeoPoint.Encode())
	x.String(m.Title)
	x.String(m.Address)
	x.String(m.Provider)
	x.String(m.VenueId)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLInputBotInlineMessageMediaVenue) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.GeoPoint = &InputGeoPoint{}
	m.Decode(dbuf)
	m.Title = dbuf.String()
	m.Address = dbuf.String()
	m.Provider = dbuf.String()
	m.VenueId = dbuf.String()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineMessageMediaContact#2daf01a7 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaContact))
	x.Int(m.Flags)
	x.String(m.PhoneNumber)
	x.String(m.FirstName)
	x.String(m.LastName)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLInputBotInlineMessageMediaContact) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.PhoneNumber = dbuf.String()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineMessageGame#4b425864 flags:# reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageGame))
	x.Int(m.Flags)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLInputBotInlineMessageGame) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineResult#2cbbe15a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResult))
	x.Int(m.Flags)
	x.String(m.Id)
	x.String(m.Type)
	x.String(m.Title)
	x.String(m.Description)
	x.String(m.Url)
	x.String(m.ThumbUrl)
	x.String(m.ContentUrl)
	x.String(m.ContentType)
	x.Int(m.W)
	x.Int(m.H)
	x.Int(m.Duration)
	x.Bytes(m.SendMessage.Encode())
	return x.buf
}

func (m *TLInputBotInlineResult) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.String()
	m.Type = dbuf.String()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Url = dbuf.String()
	m.ThumbUrl = dbuf.String()
	m.ContentUrl = dbuf.String()
	m.ContentType = dbuf.String()
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	m.Duration = dbuf.Int()
	m.SendMessage = &InputBotInlineMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineResultPhoto#a8d864a7 id:string type:string photo:InputPhoto send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResultPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResultPhoto))
	x.String(m.Id)
	x.String(m.Type)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.SendMessage.Encode())
	return x.buf
}

func (m *TLInputBotInlineResultPhoto) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.String()
	m.Type = dbuf.String()
	m.Photo = &InputPhoto{}
	m.Decode(dbuf)
	m.SendMessage = &InputBotInlineMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineResultDocument#fff8fdc4 flags:# id:string type:string title:flags.1?string description:flags.2?string document:InputDocument send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResultDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResultDocument))
	x.Int(m.Flags)
	x.String(m.Id)
	x.String(m.Type)
	x.String(m.Title)
	x.String(m.Description)
	x.Bytes(m.Document.Encode())
	x.Bytes(m.SendMessage.Encode())
	return x.buf
}

func (m *TLInputBotInlineResultDocument) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.String()
	m.Type = dbuf.String()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Document = &InputDocument{}
	m.Decode(dbuf)
	m.SendMessage = &InputBotInlineMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputBotInlineResultGame#4fa417f2 id:string short_name:string send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResultGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResultGame))
	x.String(m.Id)
	x.String(m.ShortName)
	x.Bytes(m.SendMessage.Encode())
	return x.buf
}

func (m *TLInputBotInlineResultGame) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.String()
	m.ShortName = dbuf.String()
	m.SendMessage = &InputBotInlineMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineMessageMediaAuto#a74b15b flags:# caption:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaAuto))
	x.Int(m.Flags)
	x.String(m.Caption)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLBotInlineMessageMediaAuto) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Caption = dbuf.String()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineMessageText#8c7f65e2 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageText))
	x.Int(m.Flags)
	x.String(m.Message)
	// x.VectorMessage(m.Entities);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLBotInlineMessageText) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Message = dbuf.String()
	// x.VectorMessage(m.Entities);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l4)
	for i := 0; i < int(l4); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineMessageMediaGeo#3a8fd8b8 flags:# geo:GeoPoint reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaGeo))
	x.Int(m.Flags)
	x.Bytes(m.Geo.Encode())
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLBotInlineMessageMediaGeo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Geo = &GeoPoint{}
	m.Decode(dbuf)
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineMessageMediaVenue#4366232e flags:# geo:GeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaVenue))
	x.Int(m.Flags)
	x.Bytes(m.Geo.Encode())
	x.String(m.Title)
	x.String(m.Address)
	x.String(m.Provider)
	x.String(m.VenueId)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLBotInlineMessageMediaVenue) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Geo = &GeoPoint{}
	m.Decode(dbuf)
	m.Title = dbuf.String()
	m.Address = dbuf.String()
	m.Provider = dbuf.String()
	m.VenueId = dbuf.String()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineMessageMediaContact#35edb4d4 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaContact))
	x.Int(m.Flags)
	x.String(m.PhoneNumber)
	x.String(m.FirstName)
	x.String(m.LastName)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLBotInlineMessageMediaContact) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.PhoneNumber = dbuf.String()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	m.ReplyMarkup = &ReplyMarkup{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineResult#9bebaeb9 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:BotInlineMessage = BotInlineResult;
func (m *TLBotInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineResult))
	x.Int(m.Flags)
	x.String(m.Id)
	x.String(m.Type)
	x.String(m.Title)
	x.String(m.Description)
	x.String(m.Url)
	x.String(m.ThumbUrl)
	x.String(m.ContentUrl)
	x.String(m.ContentType)
	x.Int(m.W)
	x.Int(m.H)
	x.Int(m.Duration)
	x.Bytes(m.SendMessage.Encode())
	return x.buf
}

func (m *TLBotInlineResult) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.String()
	m.Type = dbuf.String()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Url = dbuf.String()
	m.ThumbUrl = dbuf.String()
	m.ContentUrl = dbuf.String()
	m.ContentType = dbuf.String()
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	m.Duration = dbuf.Int()
	m.SendMessage = &BotInlineMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// botInlineMediaResult#17db940b flags:# id:string type:string photo:flags.0?Photo document:flags.1?Document title:flags.2?string description:flags.3?string send_message:BotInlineMessage = BotInlineResult;
func (m *TLBotInlineMediaResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMediaResult))
	x.Int(m.Flags)
	x.String(m.Id)
	x.String(m.Type)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.Document.Encode())
	x.String(m.Title)
	x.String(m.Description)
	x.Bytes(m.SendMessage.Encode())
	return x.buf
}

func (m *TLBotInlineMediaResult) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.String()
	m.Type = dbuf.String()
	m.Photo = &Photo{}
	m.Decode(dbuf)
	m.Document = &Document{}
	m.Decode(dbuf)
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.SendMessage = &BotInlineMessage{}
	m.Decode(dbuf)
	return dbuf.err
}

// messages.botResults#ccd3563d flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int = messages.BotResults;
func (m *TLMessagesBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_botResults))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.String(m.NextOffset)
	x.Bytes(m.SwitchPm.Encode())
	// x.VectorMessage(m.Results);
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Results)))
	for _, v := range m.Results {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.CacheTime)
	return x.buf
}

func (m *TLMessagesBotResults) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.NextOffset = dbuf.String()
	m.SwitchPm = &InlineBotSwitchPM{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Results);
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Results = make([]*BotInlineResult, l6)
	for i := 0; i < int(l6); i++ {
		m.Results[i] = &BotInlineResult{}
		(*m.Results[i]).Decode(dbuf)
	}
	m.CacheTime = dbuf.Int()
	return dbuf.err
}

// exportedMessageLink#1f486803 link:string = ExportedMessageLink;
func (m *TLExportedMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_exportedMessageLink))
	x.String(m.Link)
	return x.buf
}

func (m *TLExportedMessageLink) Decode(dbuf *DecodeBuf) error {
	m.Link = dbuf.String()
	return dbuf.err
}

// messageFwdHeader#fadff4ac flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string = MessageFwdHeader;
func (m *TLMessageFwdHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageFwdHeader))
	x.Int(m.Flags)
	x.Int(m.FromId)
	x.Int(m.Date)
	x.Int(m.ChannelId)
	x.Int(m.ChannelPost)
	x.String(m.PostAuthor)
	return x.buf
}

func (m *TLMessageFwdHeader) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.FromId = dbuf.Int()
	m.Date = dbuf.Int()
	m.ChannelId = dbuf.Int()
	m.ChannelPost = dbuf.Int()
	m.PostAuthor = dbuf.String()
	return dbuf.err
}

// auth.codeTypeSms#72a3158c = auth.CodeType;
func (m *TLAuthCodeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_codeTypeSms))
	return x.buf
}

func (m *TLAuthCodeTypeSms) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// auth.codeTypeCall#741cd3e3 = auth.CodeType;
func (m *TLAuthCodeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_codeTypeCall))
	return x.buf
}

func (m *TLAuthCodeTypeCall) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// auth.codeTypeFlashCall#226ccefb = auth.CodeType;
func (m *TLAuthCodeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_codeTypeFlashCall))
	return x.buf
}

func (m *TLAuthCodeTypeFlashCall) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
func (m *TLAuthSentCodeTypeApp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeApp))
	x.Int(m.Length)
	return x.buf
}

func (m *TLAuthSentCodeTypeApp) Decode(dbuf *DecodeBuf) error {
	m.Length = dbuf.Int()
	return dbuf.err
}

// auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
func (m *TLAuthSentCodeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeSms))
	x.Int(m.Length)
	return x.buf
}

func (m *TLAuthSentCodeTypeSms) Decode(dbuf *DecodeBuf) error {
	m.Length = dbuf.Int()
	return dbuf.err
}

// auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
func (m *TLAuthSentCodeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeCall))
	x.Int(m.Length)
	return x.buf
}

func (m *TLAuthSentCodeTypeCall) Decode(dbuf *DecodeBuf) error {
	m.Length = dbuf.Int()
	return dbuf.err
}

// auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
func (m *TLAuthSentCodeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeFlashCall))
	x.String(m.Pattern)
	return x.buf
}

func (m *TLAuthSentCodeTypeFlashCall) Decode(dbuf *DecodeBuf) error {
	m.Pattern = dbuf.String()
	return dbuf.err
}

// messages.botCallbackAnswer#36585ea4 flags:# alert:flags.1?true has_url:flags.3?true message:flags.0?string url:flags.2?string cache_time:int = messages.BotCallbackAnswer;
func (m *TLMessagesBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_botCallbackAnswer))
	x.Int(m.Flags)
	x.String(m.Message)
	x.String(m.Url)
	x.Int(m.CacheTime)
	return x.buf
}

func (m *TLMessagesBotCallbackAnswer) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Message = dbuf.String()
	m.Url = dbuf.String()
	m.CacheTime = dbuf.Int()
	return dbuf.err
}

// messages.messageEditData#26b5dde6 flags:# caption:flags.0?true = messages.MessageEditData;
func (m *TLMessagesMessageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_messageEditData))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLMessagesMessageEditData) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// inputBotInlineMessageID#890c3d89 dc_id:int id:long access_hash:long = InputBotInlineMessageID;
func (m *TLInputBotInlineMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageID))
	x.Int(m.DcId)
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputBotInlineMessageID) Decode(dbuf *DecodeBuf) error {
	m.DcId = dbuf.Int()
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inlineBotSwitchPM#3c20629f text:string start_param:string = InlineBotSwitchPM;
func (m *TLInlineBotSwitchPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inlineBotSwitchPM))
	x.String(m.Text)
	x.String(m.StartParam)
	return x.buf
}

func (m *TLInlineBotSwitchPM) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	m.StartParam = dbuf.String()
	return dbuf.err
}

// messages.peerDialogs#3371c354 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> state:updates.State = messages.PeerDialogs;
func (m *TLMessagesPeerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_peerDialogs))
	// x.VectorMessage(m.Dialogs);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Dialogs)))
	for _, v := range m.Dialogs {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Messages);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Messages)))
	for _, v := range m.Messages {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.State.Encode())
	return x.buf
}

func (m *TLMessagesPeerDialogs) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Dialogs);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Dialogs = make([]*Dialog, l1)
	for i := 0; i < int(l1); i++ {
		m.Dialogs[i] = &Dialog{}
		(*m.Dialogs[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Messages);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Messages = make([]*Message, l2)
	for i := 0; i < int(l2); i++ {
		m.Messages[i] = &Message{}
		(*m.Messages[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Chats = make([]*Chat, l3)
	for i := 0; i < int(l3); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Users = make([]*User, l4)
	for i := 0; i < int(l4); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	m.State = &Updates_State{}
	m.Decode(dbuf)
	return dbuf.err
}

// topPeer#edcdc05b peer:Peer rating:double = TopPeer;
func (m *TLTopPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeer))
	x.Bytes(m.Peer.Encode())
	x.Double(m.Rating)
	return x.buf
}

func (m *TLTopPeer) Decode(dbuf *DecodeBuf) error {
	m.Peer = &Peer{}
	m.Decode(dbuf)
	m.Rating = dbuf.Double()
	return dbuf.err
}

// topPeerCategoryBotsPM#ab661b5b = TopPeerCategory;
func (m *TLTopPeerCategoryBotsPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryBotsPM))
	return x.buf
}

func (m *TLTopPeerCategoryBotsPM) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// topPeerCategoryBotsInline#148677e2 = TopPeerCategory;
func (m *TLTopPeerCategoryBotsInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryBotsInline))
	return x.buf
}

func (m *TLTopPeerCategoryBotsInline) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// topPeerCategoryCorrespondents#637b7ed = TopPeerCategory;
func (m *TLTopPeerCategoryCorrespondents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryCorrespondents))
	return x.buf
}

func (m *TLTopPeerCategoryCorrespondents) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// topPeerCategoryGroups#bd17a14a = TopPeerCategory;
func (m *TLTopPeerCategoryGroups) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryGroups))
	return x.buf
}

func (m *TLTopPeerCategoryGroups) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// topPeerCategoryChannels#161d9628 = TopPeerCategory;
func (m *TLTopPeerCategoryChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryChannels))
	return x.buf
}

func (m *TLTopPeerCategoryChannels) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// topPeerCategoryPhoneCalls#1e76a78c = TopPeerCategory;
func (m *TLTopPeerCategoryPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryPhoneCalls))
	return x.buf
}

func (m *TLTopPeerCategoryPhoneCalls) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// topPeerCategoryPeers#fb834291 category:TopPeerCategory count:int peers:Vector<TopPeer> = TopPeerCategoryPeers;
func (m *TLTopPeerCategoryPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryPeers))
	x.Bytes(m.Category.Encode())
	x.Int(m.Count)
	// x.VectorMessage(m.Peers);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Peers)))
	for _, v := range m.Peers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLTopPeerCategoryPeers) Decode(dbuf *DecodeBuf) error {
	m.Category = &TopPeerCategory{}
	m.Decode(dbuf)
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Peers);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Peers = make([]*TopPeer, l3)
	for i := 0; i < int(l3); i++ {
		m.Peers[i] = &TopPeer{}
		(*m.Peers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// contacts.topPeersNotModified#de266ef5 = contacts.TopPeers;
func (m *TLContactsTopPeersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_topPeersNotModified))
	return x.buf
}

func (m *TLContactsTopPeersNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// contacts.topPeers#70b772a8 categories:Vector<TopPeerCategoryPeers> chats:Vector<Chat> users:Vector<User> = contacts.TopPeers;
func (m *TLContactsTopPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_topPeers))
	// x.VectorMessage(m.Categories);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Categories)))
	for _, v := range m.Categories {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsTopPeers) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Categories);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Categories = make([]*TopPeerCategoryPeers, l1)
	for i := 0; i < int(l1); i++ {
		m.Categories[i] = &TopPeerCategoryPeers{}
		(*m.Categories[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// draftMessageEmpty#ba4baec5 = DraftMessage;
func (m *TLDraftMessageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_draftMessageEmpty))
	return x.buf
}

func (m *TLDraftMessageEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;
func (m *TLDraftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_draftMessage))
	x.Int(m.Flags)
	x.Int(m.ReplyToMsgId)
	x.String(m.Message)
	// x.VectorMessage(m.Entities);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.Date)
	return x.buf
}

func (m *TLDraftMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.ReplyToMsgId = dbuf.Int()
	m.Message = dbuf.String()
	// x.VectorMessage(m.Entities);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l5)
	for i := 0; i < int(l5); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	m.Date = dbuf.Int()
	return dbuf.err
}

// messages.featuredStickersNotModified#4ede3cf = messages.FeaturedStickers;
func (m *TLMessagesFeaturedStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_featuredStickersNotModified))
	return x.buf
}

func (m *TLMessagesFeaturedStickersNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.featuredStickers#f89d88e5 hash:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
func (m *TLMessagesFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_featuredStickers))
	x.Int(m.Hash)
	// x.VectorMessage(m.Sets);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Sets)))
	for _, v := range m.Sets {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.VectorLong(m.Unread)
	return x.buf
}

func (m *TLMessagesFeaturedStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	// x.VectorMessage(m.Sets);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Sets = make([]*StickerSetCovered, l2)
	for i := 0; i < int(l2); i++ {
		m.Sets[i] = &StickerSetCovered{}
		(*m.Sets[i]).Decode(dbuf)
	}
	m.Unread = dbuf.VectorLong()
	return dbuf.err
}

// messages.recentStickersNotModified#b17f890 = messages.RecentStickers;
func (m *TLMessagesRecentStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_recentStickersNotModified))
	return x.buf
}

func (m *TLMessagesRecentStickersNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.recentStickers#5ce20970 hash:int stickers:Vector<Document> = messages.RecentStickers;
func (m *TLMessagesRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_recentStickers))
	x.Int(m.Hash)
	// x.VectorMessage(m.Stickers);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Stickers)))
	for _, v := range m.Stickers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesRecentStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	// x.VectorMessage(m.Stickers);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Stickers = make([]*Document, l2)
	for i := 0; i < int(l2); i++ {
		m.Stickers[i] = &Document{}
		(*m.Stickers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.archivedStickers#4fcba9c8 count:int sets:Vector<StickerSetCovered> = messages.ArchivedStickers;
func (m *TLMessagesArchivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_archivedStickers))
	x.Int(m.Count)
	// x.VectorMessage(m.Sets);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Sets)))
	for _, v := range m.Sets {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesArchivedStickers) Decode(dbuf *DecodeBuf) error {
	m.Count = dbuf.Int()
	// x.VectorMessage(m.Sets);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Sets = make([]*StickerSetCovered, l2)
	for i := 0; i < int(l2); i++ {
		m.Sets[i] = &StickerSetCovered{}
		(*m.Sets[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.stickerSetInstallResultSuccess#38641628 = messages.StickerSetInstallResult;
func (m *TLMessagesStickerSetInstallResultSuccess) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickerSetInstallResultSuccess))
	return x.buf
}

func (m *TLMessagesStickerSetInstallResultSuccess) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.stickerSetInstallResultArchive#35e410a8 sets:Vector<StickerSetCovered> = messages.StickerSetInstallResult;
func (m *TLMessagesStickerSetInstallResultArchive) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickerSetInstallResultArchive))
	// x.VectorMessage(m.Sets);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Sets)))
	for _, v := range m.Sets {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesStickerSetInstallResultArchive) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Sets);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Sets = make([]*StickerSetCovered, l1)
	for i := 0; i < int(l1); i++ {
		m.Sets[i] = &StickerSetCovered{}
		(*m.Sets[i]).Decode(dbuf)
	}
	return dbuf.err
}

// stickerSetCovered#6410a5d2 set:StickerSet cover:Document = StickerSetCovered;
func (m *TLStickerSetCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerSetCovered))
	x.Bytes(m.Set.Encode())
	x.Bytes(m.Cover.Encode())
	return x.buf
}

func (m *TLStickerSetCovered) Decode(dbuf *DecodeBuf) error {
	m.Set = &StickerSet{}
	m.Decode(dbuf)
	m.Cover = &Document{}
	m.Decode(dbuf)
	return dbuf.err
}

// stickerSetMultiCovered#3407e51b set:StickerSet covers:Vector<Document> = StickerSetCovered;
func (m *TLStickerSetMultiCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerSetMultiCovered))
	x.Bytes(m.Set.Encode())
	// x.VectorMessage(m.Covers);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Covers)))
	for _, v := range m.Covers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLStickerSetMultiCovered) Decode(dbuf *DecodeBuf) error {
	m.Set = &StickerSet{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Covers);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Covers = make([]*Document, l2)
	for i := 0; i < int(l2); i++ {
		m.Covers[i] = &Document{}
		(*m.Covers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// maskCoords#aed6dbb2 n:int x:double y:double zoom:double = MaskCoords;
func (m *TLMaskCoords) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_maskCoords))
	x.Int(m.N)
	x.Double(m.X)
	x.Double(m.Y)
	x.Double(m.Zoom)
	return x.buf
}

func (m *TLMaskCoords) Decode(dbuf *DecodeBuf) error {
	m.N = dbuf.Int()
	m.X = dbuf.Double()
	m.Y = dbuf.Double()
	m.Zoom = dbuf.Double()
	return dbuf.err
}

// inputStickeredMediaPhoto#4a992157 id:InputPhoto = InputStickeredMedia;
func (m *TLInputStickeredMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickeredMediaPhoto))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLInputStickeredMediaPhoto) Decode(dbuf *DecodeBuf) error {
	m.Id = &InputPhoto{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputStickeredMediaDocument#438865b id:InputDocument = InputStickeredMedia;
func (m *TLInputStickeredMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickeredMediaDocument))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLInputStickeredMediaDocument) Decode(dbuf *DecodeBuf) error {
	m.Id = &InputDocument{}
	m.Decode(dbuf)
	return dbuf.err
}

// game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
func (m *TLGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_game))
	x.Int(m.Flags)
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.String(m.ShortName)
	x.String(m.Title)
	x.String(m.Description)
	x.Bytes(m.Photo.Encode())
	x.Bytes(m.Document.Encode())
	return x.buf
}

func (m *TLGame) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.ShortName = dbuf.String()
	m.Title = dbuf.String()
	m.Description = dbuf.String()
	m.Photo = &Photo{}
	m.Decode(dbuf)
	m.Document = &Document{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputGameID#32c3e77 id:long access_hash:long = InputGame;
func (m *TLInputGameID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGameID))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputGameID) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;
func (m *TLInputGameShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGameShortName))
	x.Bytes(m.BotId.Encode())
	x.String(m.ShortName)
	return x.buf
}

func (m *TLInputGameShortName) Decode(dbuf *DecodeBuf) error {
	m.BotId = &InputUser{}
	m.Decode(dbuf)
	m.ShortName = dbuf.String()
	return dbuf.err
}

// highScore#58fffcd0 pos:int user_id:int score:int = HighScore;
func (m *TLHighScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_highScore))
	x.Int(m.Pos)
	x.Int(m.UserId)
	x.Int(m.Score)
	return x.buf
}

func (m *TLHighScore) Decode(dbuf *DecodeBuf) error {
	m.Pos = dbuf.Int()
	m.UserId = dbuf.Int()
	m.Score = dbuf.Int()
	return dbuf.err
}

// messages.highScores#9a3bfd99 scores:Vector<HighScore> users:Vector<User> = messages.HighScores;
func (m *TLMessagesHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_highScores))
	// x.VectorMessage(m.Scores);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Scores)))
	for _, v := range m.Scores {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesHighScores) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Scores);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Scores = make([]*HighScore, l1)
	for i := 0; i < int(l1); i++ {
		m.Scores[i] = &HighScore{}
		(*m.Scores[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// textEmpty#dc3d824f = RichText;
func (m *TLTextEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textEmpty))
	return x.buf
}

func (m *TLTextEmpty) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// textPlain#744694e0 text:string = RichText;
func (m *TLTextPlain) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textPlain))
	x.String(m.Text)
	return x.buf
}

func (m *TLTextPlain) Decode(dbuf *DecodeBuf) error {
	m.Text = dbuf.String()
	return dbuf.err
}

// textBold#6724abc4 text:RichText = RichText;
func (m *TLTextBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textBold))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLTextBold) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// textItalic#d912a59c text:RichText = RichText;
func (m *TLTextItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textItalic))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLTextItalic) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// textUnderline#c12622c4 text:RichText = RichText;
func (m *TLTextUnderline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textUnderline))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLTextUnderline) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// textStrike#9bf8bb95 text:RichText = RichText;
func (m *TLTextStrike) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textStrike))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLTextStrike) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// textFixed#6c3f19b9 text:RichText = RichText;
func (m *TLTextFixed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textFixed))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLTextFixed) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// textUrl#3c2884c1 text:RichText url:string webpage_id:long = RichText;
func (m *TLTextUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textUrl))
	x.Bytes(m.Text.Encode())
	x.String(m.Url)
	x.Long(m.WebpageId)
	return x.buf
}

func (m *TLTextUrl) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	m.Url = dbuf.String()
	m.WebpageId = dbuf.Long()
	return dbuf.err
}

// textEmail#de5a0dd6 text:RichText email:string = RichText;
func (m *TLTextEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textEmail))
	x.Bytes(m.Text.Encode())
	x.String(m.Email)
	return x.buf
}

func (m *TLTextEmail) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	m.Email = dbuf.String()
	return dbuf.err
}

// textConcat#7e6260d7 texts:Vector<RichText> = RichText;
func (m *TLTextConcat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textConcat))
	// x.VectorMessage(m.Texts);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Texts)))
	for _, v := range m.Texts {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLTextConcat) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Texts);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Texts = make([]*RichText, l1)
	for i := 0; i < int(l1); i++ {
		m.Texts[i] = &RichText{}
		(*m.Texts[i]).Decode(dbuf)
	}
	return dbuf.err
}

// pageBlockUnsupported#13567e8a = PageBlock;
func (m *TLPageBlockUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockUnsupported))
	return x.buf
}

func (m *TLPageBlockUnsupported) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// pageBlockTitle#70abc3fd text:RichText = PageBlock;
func (m *TLPageBlockTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockTitle))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLPageBlockTitle) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockSubtitle#8ffa9a1f text:RichText = PageBlock;
func (m *TLPageBlockSubtitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockSubtitle))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLPageBlockSubtitle) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockAuthorDate#baafe5e0 author:RichText published_date:int = PageBlock;
func (m *TLPageBlockAuthorDate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockAuthorDate))
	x.Bytes(m.Author.Encode())
	x.Int(m.PublishedDate)
	return x.buf
}

func (m *TLPageBlockAuthorDate) Decode(dbuf *DecodeBuf) error {
	m.Author = &RichText{}
	m.Decode(dbuf)
	m.PublishedDate = dbuf.Int()
	return dbuf.err
}

// pageBlockHeader#bfd064ec text:RichText = PageBlock;
func (m *TLPageBlockHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockHeader))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLPageBlockHeader) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockSubheader#f12bb6e1 text:RichText = PageBlock;
func (m *TLPageBlockSubheader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockSubheader))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLPageBlockSubheader) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockParagraph#467a0766 text:RichText = PageBlock;
func (m *TLPageBlockParagraph) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockParagraph))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLPageBlockParagraph) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockPreformatted#c070d93e text:RichText language:string = PageBlock;
func (m *TLPageBlockPreformatted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockPreformatted))
	x.Bytes(m.Text.Encode())
	x.String(m.Language)
	return x.buf
}

func (m *TLPageBlockPreformatted) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	m.Language = dbuf.String()
	return dbuf.err
}

// pageBlockFooter#48870999 text:RichText = PageBlock;
func (m *TLPageBlockFooter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockFooter))
	x.Bytes(m.Text.Encode())
	return x.buf
}

func (m *TLPageBlockFooter) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockDivider#db20b188 = PageBlock;
func (m *TLPageBlockDivider) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockDivider))
	return x.buf
}

func (m *TLPageBlockDivider) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// pageBlockAnchor#ce0d37b0 name:string = PageBlock;
func (m *TLPageBlockAnchor) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockAnchor))
	x.String(m.Name)
	return x.buf
}

func (m *TLPageBlockAnchor) Decode(dbuf *DecodeBuf) error {
	m.Name = dbuf.String()
	return dbuf.err
}

// pageBlockList#3a58c7f4 ordered:Bool items:Vector<RichText> = PageBlock;
func (m *TLPageBlockList) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockList))
	x.Bytes(m.Ordered.Encode())
	// x.VectorMessage(m.Items);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Items)))
	for _, v := range m.Items {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPageBlockList) Decode(dbuf *DecodeBuf) error {
	m.Ordered = &Bool{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Items);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Items = make([]*RichText, l2)
	for i := 0; i < int(l2); i++ {
		m.Items[i] = &RichText{}
		(*m.Items[i]).Decode(dbuf)
	}
	return dbuf.err
}

// pageBlockBlockquote#263d7c26 text:RichText caption:RichText = PageBlock;
func (m *TLPageBlockBlockquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockBlockquote))
	x.Bytes(m.Text.Encode())
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockBlockquote) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockPullquote#4f4456d3 text:RichText caption:RichText = PageBlock;
func (m *TLPageBlockPullquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockPullquote))
	x.Bytes(m.Text.Encode())
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockPullquote) Decode(dbuf *DecodeBuf) error {
	m.Text = &RichText{}
	m.Decode(dbuf)
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockPhoto#e9c69982 photo_id:long caption:RichText = PageBlock;
func (m *TLPageBlockPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockPhoto))
	x.Long(m.PhotoId)
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockPhoto) Decode(dbuf *DecodeBuf) error {
	m.PhotoId = dbuf.Long()
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockVideo#d9d71866 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:RichText = PageBlock;
func (m *TLPageBlockVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockVideo))
	x.Int(m.Flags)
	x.Long(m.VideoId)
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockVideo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.VideoId = dbuf.Long()
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockCover#39f23300 cover:PageBlock = PageBlock;
func (m *TLPageBlockCover) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockCover))
	x.Bytes(m.Cover.Encode())
	return x.buf
}

func (m *TLPageBlockCover) Decode(dbuf *DecodeBuf) error {
	m.Cover = &PageBlock{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockEmbed#cde200d1 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:int h:int caption:RichText = PageBlock;
func (m *TLPageBlockEmbed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockEmbed))
	x.Int(m.Flags)
	x.String(m.Url)
	x.String(m.Html)
	x.Long(m.PosterPhotoId)
	x.Int(m.W)
	x.Int(m.H)
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockEmbed) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Url = dbuf.String()
	m.Html = dbuf.String()
	m.PosterPhotoId = dbuf.Long()
	m.W = dbuf.Int()
	m.H = dbuf.Int()
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockEmbedPost#292c7be9 url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:RichText = PageBlock;
func (m *TLPageBlockEmbedPost) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockEmbedPost))
	x.String(m.Url)
	x.Long(m.WebpageId)
	x.Long(m.AuthorPhotoId)
	x.String(m.Author)
	x.Int(m.Date)
	// x.VectorMessage(m.Blocks);
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Blocks)))
	for _, v := range m.Blocks {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockEmbedPost) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.WebpageId = dbuf.Long()
	m.AuthorPhotoId = dbuf.Long()
	m.Author = dbuf.String()
	m.Date = dbuf.Int()
	// x.VectorMessage(m.Blocks);
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Blocks = make([]*PageBlock, l6)
	for i := 0; i < int(l6); i++ {
		m.Blocks[i] = &PageBlock{}
		(*m.Blocks[i]).Decode(dbuf)
	}
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockCollage#8b31c4f items:Vector<PageBlock> caption:RichText = PageBlock;
func (m *TLPageBlockCollage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockCollage))
	// x.VectorMessage(m.Items);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Items)))
	for _, v := range m.Items {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockCollage) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Items);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Items = make([]*PageBlock, l1)
	for i := 0; i < int(l1); i++ {
		m.Items[i] = &PageBlock{}
		(*m.Items[i]).Decode(dbuf)
	}
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockSlideshow#130c8963 items:Vector<PageBlock> caption:RichText = PageBlock;
func (m *TLPageBlockSlideshow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockSlideshow))
	// x.VectorMessage(m.Items);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Items)))
	for _, v := range m.Items {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockSlideshow) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Items);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Items = make([]*PageBlock, l1)
	for i := 0; i < int(l1); i++ {
		m.Items[i] = &PageBlock{}
		(*m.Items[i]).Decode(dbuf)
	}
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockChannel#ef1751b5 channel:Chat = PageBlock;
func (m *TLPageBlockChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockChannel))
	x.Bytes(m.Channel.Encode())
	return x.buf
}

func (m *TLPageBlockChannel) Decode(dbuf *DecodeBuf) error {
	m.Channel = &Chat{}
	m.Decode(dbuf)
	return dbuf.err
}

// pageBlockAudio#31b81a7f audio_id:long caption:RichText = PageBlock;
func (m *TLPageBlockAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockAudio))
	x.Long(m.AudioId)
	x.Bytes(m.Caption.Encode())
	return x.buf
}

func (m *TLPageBlockAudio) Decode(dbuf *DecodeBuf) error {
	m.AudioId = dbuf.Long()
	m.Caption = &RichText{}
	m.Decode(dbuf)
	return dbuf.err
}

// pagePart#8e3f9ebe blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
func (m *TLPagePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pagePart))
	// x.VectorMessage(m.Blocks);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Blocks)))
	for _, v := range m.Blocks {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Photos);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Photos)))
	for _, v := range m.Photos {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Documents);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Documents)))
	for _, v := range m.Documents {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPagePart) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Blocks);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Blocks = make([]*PageBlock, l1)
	for i := 0; i < int(l1); i++ {
		m.Blocks[i] = &PageBlock{}
		(*m.Blocks[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Photos);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Photos = make([]*Photo, l2)
	for i := 0; i < int(l2); i++ {
		m.Photos[i] = &Photo{}
		(*m.Photos[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Documents);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Documents = make([]*Document, l3)
	for i := 0; i < int(l3); i++ {
		m.Documents[i] = &Document{}
		(*m.Documents[i]).Decode(dbuf)
	}
	return dbuf.err
}

// pageFull#556ec7aa blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
func (m *TLPageFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageFull))
	// x.VectorMessage(m.Blocks);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Blocks)))
	for _, v := range m.Blocks {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Photos);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Photos)))
	for _, v := range m.Photos {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Documents);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Documents)))
	for _, v := range m.Documents {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPageFull) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Blocks);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Blocks = make([]*PageBlock, l1)
	for i := 0; i < int(l1); i++ {
		m.Blocks[i] = &PageBlock{}
		(*m.Blocks[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Photos);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Photos = make([]*Photo, l2)
	for i := 0; i < int(l2); i++ {
		m.Photos[i] = &Photo{}
		(*m.Photos[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Documents);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Documents = make([]*Document, l3)
	for i := 0; i < int(l3); i++ {
		m.Documents[i] = &Document{}
		(*m.Documents[i]).Decode(dbuf)
	}
	return dbuf.err
}

// phoneCallDiscardReasonMissed#85e42301 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonMissed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonMissed))
	return x.buf
}

func (m *TLPhoneCallDiscardReasonMissed) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// phoneCallDiscardReasonDisconnect#e095c1a0 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonDisconnect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonDisconnect))
	return x.buf
}

func (m *TLPhoneCallDiscardReasonDisconnect) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// phoneCallDiscardReasonHangup#57adc690 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonHangup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonHangup))
	return x.buf
}

func (m *TLPhoneCallDiscardReasonHangup) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// phoneCallDiscardReasonBusy#faf7e8c9 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonBusy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonBusy))
	return x.buf
}

func (m *TLPhoneCallDiscardReasonBusy) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// dataJSON#7d748d04 data:string = DataJSON;
func (m *TLDataJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dataJSON))
	x.String(m.Data)
	return x.buf
}

func (m *TLDataJSON) Decode(dbuf *DecodeBuf) error {
	m.Data = dbuf.String()
	return dbuf.err
}

// labeledPrice#cb296bf8 label:string amount:long = LabeledPrice;
func (m *TLLabeledPrice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_labeledPrice))
	x.String(m.Label)
	x.Long(m.Amount)
	return x.buf
}

func (m *TLLabeledPrice) Decode(dbuf *DecodeBuf) error {
	m.Label = dbuf.String()
	m.Amount = dbuf.Long()
	return dbuf.err
}

// invoice#c30aa358 flags:# test:flags.0?true name_requested:flags.1?true phone_requested:flags.2?true email_requested:flags.3?true shipping_address_requested:flags.4?true flexible:flags.5?true currency:string prices:Vector<LabeledPrice> = Invoice;
func (m *TLInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_invoice))
	x.Int(m.Flags)
	x.String(m.Currency)
	// x.VectorMessage(m.Prices);
	x9 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x9, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x9[4:], uint32(len(m.Prices)))
	for _, v := range m.Prices {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLInvoice) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Currency = dbuf.String()
	// x.VectorMessage(m.Prices);
	c9 := dbuf.Int()
	if c9 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c9)
	}
	l9 := dbuf.Int()
	m.Prices = make([]*LabeledPrice, l9)
	for i := 0; i < int(l9); i++ {
		m.Prices[i] = &LabeledPrice{}
		(*m.Prices[i]).Decode(dbuf)
	}
	return dbuf.err
}

// paymentCharge#ea02c27e id:string provider_charge_id:string = PaymentCharge;
func (m *TLPaymentCharge) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_paymentCharge))
	x.String(m.Id)
	x.String(m.ProviderChargeId)
	return x.buf
}

func (m *TLPaymentCharge) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.String()
	m.ProviderChargeId = dbuf.String()
	return dbuf.err
}

// postAddress#1e8caaeb street_line1:string street_line2:string city:string state:string country_iso2:string post_code:string = PostAddress;
func (m *TLPostAddress) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_postAddress))
	x.String(m.StreetLine1)
	x.String(m.StreetLine2)
	x.String(m.City)
	x.String(m.State)
	x.String(m.CountryIso2)
	x.String(m.PostCode)
	return x.buf
}

func (m *TLPostAddress) Decode(dbuf *DecodeBuf) error {
	m.StreetLine1 = dbuf.String()
	m.StreetLine2 = dbuf.String()
	m.City = dbuf.String()
	m.State = dbuf.String()
	m.CountryIso2 = dbuf.String()
	m.PostCode = dbuf.String()
	return dbuf.err
}

// paymentRequestedInfo#909c3f94 flags:# name:flags.0?string phone:flags.1?string email:flags.2?string shipping_address:flags.3?PostAddress = PaymentRequestedInfo;
func (m *TLPaymentRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_paymentRequestedInfo))
	x.Int(m.Flags)
	x.String(m.Name)
	x.String(m.Phone)
	x.String(m.Email)
	x.Bytes(m.ShippingAddress.Encode())
	return x.buf
}

func (m *TLPaymentRequestedInfo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Name = dbuf.String()
	m.Phone = dbuf.String()
	m.Email = dbuf.String()
	m.ShippingAddress = &PostAddress{}
	m.Decode(dbuf)
	return dbuf.err
}

// paymentSavedCredentialsCard#cdc27a1f id:string title:string = PaymentSavedCredentials;
func (m *TLPaymentSavedCredentialsCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_paymentSavedCredentialsCard))
	x.String(m.Id)
	x.String(m.Title)
	return x.buf
}

func (m *TLPaymentSavedCredentialsCard) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.String()
	m.Title = dbuf.String()
	return dbuf.err
}

// webDocument#c61acbd8 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> dc_id:int = WebDocument;
func (m *TLWebDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webDocument))
	x.String(m.Url)
	x.Long(m.AccessHash)
	x.Int(m.Size)
	x.String(m.MimeType)
	// x.VectorMessage(m.Attributes);
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Attributes)))
	for _, v := range m.Attributes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.DcId)
	return x.buf
}

func (m *TLWebDocument) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.AccessHash = dbuf.Long()
	m.Size = dbuf.Int()
	m.MimeType = dbuf.String()
	// x.VectorMessage(m.Attributes);
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Attributes = make([]*DocumentAttribute, l5)
	for i := 0; i < int(l5); i++ {
		m.Attributes[i] = &DocumentAttribute{}
		(*m.Attributes[i]).Decode(dbuf)
	}
	m.DcId = dbuf.Int()
	return dbuf.err
}

// inputWebDocument#9bed434d url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = InputWebDocument;
func (m *TLInputWebDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputWebDocument))
	x.String(m.Url)
	x.Int(m.Size)
	x.String(m.MimeType)
	// x.VectorMessage(m.Attributes);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Attributes)))
	for _, v := range m.Attributes {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLInputWebDocument) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.Size = dbuf.Int()
	m.MimeType = dbuf.String()
	// x.VectorMessage(m.Attributes);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Attributes = make([]*DocumentAttribute, l4)
	for i := 0; i < int(l4); i++ {
		m.Attributes[i] = &DocumentAttribute{}
		(*m.Attributes[i]).Decode(dbuf)
	}
	return dbuf.err
}

// inputWebFileLocation#c239d686 url:string access_hash:long = InputWebFileLocation;
func (m *TLInputWebFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputWebFileLocation))
	x.String(m.Url)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputWebFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// upload.webFile#21e753bc size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes = upload.WebFile;
func (m *TLUploadWebFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_webFile))
	x.Int(m.Size)
	x.String(m.MimeType)
	x.Bytes(m.FileType.Encode())
	x.Int(m.Mtime)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLUploadWebFile) Decode(dbuf *DecodeBuf) error {
	m.Size = dbuf.Int()
	m.MimeType = dbuf.String()
	m.FileType = &Storage_FileType{}
	m.Decode(dbuf)
	m.Mtime = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// payments.paymentForm#3f56aea3 flags:# can_save_credentials:flags.2?true password_missing:flags.3?true bot_id:int invoice:Invoice provider_id:int url:string native_provider:flags.4?string native_params:flags.4?DataJSON saved_info:flags.0?PaymentRequestedInfo saved_credentials:flags.1?PaymentSavedCredentials users:Vector<User> = payments.PaymentForm;
func (m *TLPaymentsPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentForm))
	x.Int(m.Flags)
	x.Int(m.BotId)
	x.Bytes(m.Invoice.Encode())
	x.Int(m.ProviderId)
	x.String(m.Url)
	x.String(m.NativeProvider)
	x.Bytes(m.NativeParams.Encode())
	x.Bytes(m.SavedInfo.Encode())
	x.Bytes(m.SavedCredentials.Encode())
	// x.VectorMessage(m.Users);
	x12 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x12, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x12[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPaymentsPaymentForm) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.BotId = dbuf.Int()
	m.Invoice = &Invoice{}
	m.Decode(dbuf)
	m.ProviderId = dbuf.Int()
	m.Url = dbuf.String()
	m.NativeProvider = dbuf.String()
	m.NativeParams = &DataJSON{}
	m.Decode(dbuf)
	m.SavedInfo = &PaymentRequestedInfo{}
	m.Decode(dbuf)
	m.SavedCredentials = &PaymentSavedCredentials{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Users);
	c12 := dbuf.Int()
	if c12 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c12)
	}
	l12 := dbuf.Int()
	m.Users = make([]*User, l12)
	for i := 0; i < int(l12); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// payments.validatedRequestedInfo#d1451883 flags:# id:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = payments.ValidatedRequestedInfo;
func (m *TLPaymentsValidatedRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_validatedRequestedInfo))
	x.Int(m.Flags)
	x.String(m.Id)
	// x.VectorMessage(m.ShippingOptions);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.ShippingOptions)))
	for _, v := range m.ShippingOptions {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPaymentsValidatedRequestedInfo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.String()
	// x.VectorMessage(m.ShippingOptions);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.ShippingOptions = make([]*ShippingOption, l3)
	for i := 0; i < int(l3); i++ {
		m.ShippingOptions[i] = &ShippingOption{}
		(*m.ShippingOptions[i]).Decode(dbuf)
	}
	return dbuf.err
}

// payments.paymentResult#4e5f810d updates:Updates = payments.PaymentResult;
func (m *TLPaymentsPaymentResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentResult))
	x.Bytes(m.Updates.Encode())
	return x.buf
}

func (m *TLPaymentsPaymentResult) Decode(dbuf *DecodeBuf) error {
	m.Updates = &Updates{}
	m.Decode(dbuf)
	return dbuf.err
}

// payments.paymentVerficationNeeded#6b56b921 url:string = payments.PaymentResult;
func (m *TLPaymentsPaymentVerficationNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentVerficationNeeded))
	x.String(m.Url)
	return x.buf
}

func (m *TLPaymentsPaymentVerficationNeeded) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	return dbuf.err
}

// payments.paymentReceipt#500911e1 flags:# date:int bot_id:int invoice:Invoice provider_id:int info:flags.0?PaymentRequestedInfo shipping:flags.1?ShippingOption currency:string total_amount:long credentials_title:string users:Vector<User> = payments.PaymentReceipt;
func (m *TLPaymentsPaymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentReceipt))
	x.Int(m.Flags)
	x.Int(m.Date)
	x.Int(m.BotId)
	x.Bytes(m.Invoice.Encode())
	x.Int(m.ProviderId)
	x.Bytes(m.Info.Encode())
	x.Bytes(m.Shipping.Encode())
	x.String(m.Currency)
	x.Long(m.TotalAmount)
	x.String(m.CredentialsTitle)
	// x.VectorMessage(m.Users);
	x11 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x11, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x11[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPaymentsPaymentReceipt) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Date = dbuf.Int()
	m.BotId = dbuf.Int()
	m.Invoice = &Invoice{}
	m.Decode(dbuf)
	m.ProviderId = dbuf.Int()
	m.Info = &PaymentRequestedInfo{}
	m.Decode(dbuf)
	m.Shipping = &ShippingOption{}
	m.Decode(dbuf)
	m.Currency = dbuf.String()
	m.TotalAmount = dbuf.Long()
	m.CredentialsTitle = dbuf.String()
	// x.VectorMessage(m.Users);
	c11 := dbuf.Int()
	if c11 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c11)
	}
	l11 := dbuf.Int()
	m.Users = make([]*User, l11)
	for i := 0; i < int(l11); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// payments.savedInfo#fb8fe43c flags:# has_saved_credentials:flags.1?true saved_info:flags.0?PaymentRequestedInfo = payments.SavedInfo;
func (m *TLPaymentsSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_savedInfo))
	x.Int(m.Flags)
	x.Bytes(m.SavedInfo.Encode())
	return x.buf
}

func (m *TLPaymentsSavedInfo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.SavedInfo = &PaymentRequestedInfo{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputPaymentCredentialsSaved#c10eb2cf id:string tmp_password:bytes = InputPaymentCredentials;
func (m *TLInputPaymentCredentialsSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPaymentCredentialsSaved))
	x.String(m.Id)
	x.StringBytes(m.TmpPassword)
	return x.buf
}

func (m *TLInputPaymentCredentialsSaved) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.String()
	m.TmpPassword = dbuf.StringBytes()
	return dbuf.err
}

// inputPaymentCredentials#3417d728 flags:# save:flags.0?true data:DataJSON = InputPaymentCredentials;
func (m *TLInputPaymentCredentials) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPaymentCredentials))
	x.Int(m.Flags)
	x.Bytes(m.Data.Encode())
	return x.buf
}

func (m *TLInputPaymentCredentials) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Data = &DataJSON{}
	m.Decode(dbuf)
	return dbuf.err
}

// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
func (m *TLAccountTmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_tmpPassword))
	x.StringBytes(m.TmpPassword)
	x.Int(m.ValidUntil)
	return x.buf
}

func (m *TLAccountTmpPassword) Decode(dbuf *DecodeBuf) error {
	m.TmpPassword = dbuf.StringBytes()
	m.ValidUntil = dbuf.Int()
	return dbuf.err
}

// shippingOption#b6213cdf id:string title:string prices:Vector<LabeledPrice> = ShippingOption;
func (m *TLShippingOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_shippingOption))
	x.String(m.Id)
	x.String(m.Title)
	// x.VectorMessage(m.Prices);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Prices)))
	for _, v := range m.Prices {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLShippingOption) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.String()
	m.Title = dbuf.String()
	// x.VectorMessage(m.Prices);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Prices = make([]*LabeledPrice, l3)
	for i := 0; i < int(l3); i++ {
		m.Prices[i] = &LabeledPrice{}
		(*m.Prices[i]).Decode(dbuf)
	}
	return dbuf.err
}

// inputStickerSetItem#ffa0a496 flags:# document:InputDocument emoji:string mask_coords:flags.0?MaskCoords = InputStickerSetItem;
func (m *TLInputStickerSetItem) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetItem))
	x.Int(m.Flags)
	x.Bytes(m.Document.Encode())
	x.String(m.Emoji)
	x.Bytes(m.MaskCoords.Encode())
	return x.buf
}

func (m *TLInputStickerSetItem) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Document = &InputDocument{}
	m.Decode(dbuf)
	m.Emoji = dbuf.String()
	m.MaskCoords = &MaskCoords{}
	m.Decode(dbuf)
	return dbuf.err
}

// inputPhoneCall#1e36fded id:long access_hash:long = InputPhoneCall;
func (m *TLInputPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhoneCall))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	return x.buf
}

func (m *TLInputPhoneCall) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	return dbuf.err
}

// phoneCallEmpty#5366c915 id:long = PhoneCall;
func (m *TLPhoneCallEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallEmpty))
	x.Long(m.Id)
	return x.buf
}

func (m *TLPhoneCallEmpty) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	return dbuf.err
}

// phoneCallWaiting#1b8f4ad1 flags:# id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
func (m *TLPhoneCallWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallWaiting))
	x.Int(m.Flags)
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	x.Bytes(m.Protocol.Encode())
	x.Int(m.ReceiveDate)
	return x.buf
}

func (m *TLPhoneCallWaiting) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	m.Protocol = &PhoneCallProtocol{}
	m.Decode(dbuf)
	m.ReceiveDate = dbuf.Int()
	return dbuf.err
}

// phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
func (m *TLPhoneCallRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallRequested))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	x.StringBytes(m.GAHash)
	x.Bytes(m.Protocol.Encode())
	return x.buf
}

func (m *TLPhoneCallRequested) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	m.GAHash = dbuf.StringBytes()
	m.Protocol = &PhoneCallProtocol{}
	m.Decode(dbuf)
	return dbuf.err
}

// phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
func (m *TLPhoneCallAccepted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallAccepted))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	x.StringBytes(m.GB)
	x.Bytes(m.Protocol.Encode())
	return x.buf
}

func (m *TLPhoneCallAccepted) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	m.GB = dbuf.StringBytes()
	m.Protocol = &PhoneCallProtocol{}
	m.Decode(dbuf)
	return dbuf.err
}

// phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
func (m *TLPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCall))
	x.Long(m.Id)
	x.Long(m.AccessHash)
	x.Int(m.Date)
	x.Int(m.AdminId)
	x.Int(m.ParticipantId)
	x.StringBytes(m.GAOrB)
	x.Long(m.KeyFingerprint)
	x.Bytes(m.Protocol.Encode())
	x.Bytes(m.Connection.Encode())
	// x.VectorMessage(m.AlternativeConnections);
	x10 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x10, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x10[4:], uint32(len(m.AlternativeConnections)))
	for _, v := range m.AlternativeConnections {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.StartDate)
	return x.buf
}

func (m *TLPhoneCall) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.AccessHash = dbuf.Long()
	m.Date = dbuf.Int()
	m.AdminId = dbuf.Int()
	m.ParticipantId = dbuf.Int()
	m.GAOrB = dbuf.StringBytes()
	m.KeyFingerprint = dbuf.Long()
	m.Protocol = &PhoneCallProtocol{}
	m.Decode(dbuf)
	m.Connection = &PhoneConnection{}
	m.Decode(dbuf)
	// x.VectorMessage(m.AlternativeConnections);
	c10 := dbuf.Int()
	if c10 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c10)
	}
	l10 := dbuf.Int()
	m.AlternativeConnections = make([]*PhoneConnection, l10)
	for i := 0; i < int(l10); i++ {
		m.AlternativeConnections[i] = &PhoneConnection{}
		(*m.AlternativeConnections[i]).Decode(dbuf)
	}
	m.StartDate = dbuf.Int()
	return dbuf.err
}

// phoneCallDiscarded#50ca4de1 flags:# need_rating:flags.2?true need_debug:flags.3?true id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = PhoneCall;
func (m *TLPhoneCallDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscarded))
	x.Int(m.Flags)
	x.Long(m.Id)
	x.Bytes(m.Reason.Encode())
	x.Int(m.Duration)
	return x.buf
}

func (m *TLPhoneCallDiscarded) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.Long()
	m.Reason = &PhoneCallDiscardReason{}
	m.Decode(dbuf)
	m.Duration = dbuf.Int()
	return dbuf.err
}

// phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
func (m *TLPhoneConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneConnection))
	x.Long(m.Id)
	x.String(m.Ip)
	x.String(m.Ipv6)
	x.Int(m.Port)
	x.StringBytes(m.PeerTag)
	return x.buf
}

func (m *TLPhoneConnection) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Ip = dbuf.String()
	m.Ipv6 = dbuf.String()
	m.Port = dbuf.Int()
	m.PeerTag = dbuf.StringBytes()
	return dbuf.err
}

// phoneCallProtocol#a2bb35cb flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int = PhoneCallProtocol;
func (m *TLPhoneCallProtocol) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallProtocol))
	x.Int(m.Flags)
	x.Int(m.MinLayer)
	x.Int(m.MaxLayer)
	return x.buf
}

func (m *TLPhoneCallProtocol) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.MinLayer = dbuf.Int()
	m.MaxLayer = dbuf.Int()
	return dbuf.err
}

// phone.phoneCall#ec82e140 phone_call:PhoneCall users:Vector<User> = phone.PhoneCall;
func (m *TLPhonePhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_phoneCall))
	x.Bytes(m.PhoneCall.Encode())
	// x.VectorMessage(m.Users);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPhonePhoneCall) Decode(dbuf *DecodeBuf) error {
	m.PhoneCall = &PhoneCall{}
	m.Decode(dbuf)
	// x.VectorMessage(m.Users);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*User, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// upload.cdnFileReuploadNeeded#eea8e46e request_token:bytes = upload.CdnFile;
func (m *TLUploadCdnFileReuploadNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_cdnFileReuploadNeeded))
	x.StringBytes(m.RequestToken)
	return x.buf
}

func (m *TLUploadCdnFileReuploadNeeded) Decode(dbuf *DecodeBuf) error {
	m.RequestToken = dbuf.StringBytes()
	return dbuf.err
}

// upload.cdnFile#a99fca4f bytes:bytes = upload.CdnFile;
func (m *TLUploadCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_cdnFile))
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLUploadCdnFile) Decode(dbuf *DecodeBuf) error {
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// cdnPublicKey#c982eaba dc_id:int public_key:string = CdnPublicKey;
func (m *TLCdnPublicKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_cdnPublicKey))
	x.Int(m.DcId)
	x.String(m.PublicKey)
	return x.buf
}

func (m *TLCdnPublicKey) Decode(dbuf *DecodeBuf) error {
	m.DcId = dbuf.Int()
	m.PublicKey = dbuf.String()
	return dbuf.err
}

// cdnConfig#5725e40a public_keys:Vector<CdnPublicKey> = CdnConfig;
func (m *TLCdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_cdnConfig))
	// x.VectorMessage(m.PublicKeys);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.PublicKeys)))
	for _, v := range m.PublicKeys {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLCdnConfig) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.PublicKeys);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.PublicKeys = make([]*CdnPublicKey, l1)
	for i := 0; i < int(l1); i++ {
		m.PublicKeys[i] = &CdnPublicKey{}
		(*m.PublicKeys[i]).Decode(dbuf)
	}
	return dbuf.err
}

// langPackString#cad181f6 key:string value:string = LangPackString;
func (m *TLLangPackString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackString))
	x.String(m.Key)
	x.String(m.Value)
	return x.buf
}

func (m *TLLangPackString) Decode(dbuf *DecodeBuf) error {
	m.Key = dbuf.String()
	m.Value = dbuf.String()
	return dbuf.err
}

// langPackStringPluralized#6c47ac9f flags:# key:string zero_value:flags.0?string one_value:flags.1?string two_value:flags.2?string few_value:flags.3?string many_value:flags.4?string other_value:string = LangPackString;
func (m *TLLangPackStringPluralized) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackStringPluralized))
	x.Int(m.Flags)
	x.String(m.Key)
	x.String(m.ZeroValue)
	x.String(m.OneValue)
	x.String(m.TwoValue)
	x.String(m.FewValue)
	x.String(m.ManyValue)
	x.String(m.OtherValue)
	return x.buf
}

func (m *TLLangPackStringPluralized) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Key = dbuf.String()
	m.ZeroValue = dbuf.String()
	m.OneValue = dbuf.String()
	m.TwoValue = dbuf.String()
	m.FewValue = dbuf.String()
	m.ManyValue = dbuf.String()
	m.OtherValue = dbuf.String()
	return dbuf.err
}

// langPackStringDeleted#2979eeb2 key:string = LangPackString;
func (m *TLLangPackStringDeleted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackStringDeleted))
	x.String(m.Key)
	return x.buf
}

func (m *TLLangPackStringDeleted) Decode(dbuf *DecodeBuf) error {
	m.Key = dbuf.String()
	return dbuf.err
}

// langPackDifference#f385c1f6 lang_code:string from_version:int version:int strings:Vector<LangPackString> = LangPackDifference;
func (m *TLLangPackDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackDifference))
	x.String(m.LangCode)
	x.Int(m.FromVersion)
	x.Int(m.Version)
	// x.VectorMessage(m.Strings);
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.Strings)))
	for _, v := range m.Strings {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLLangPackDifference) Decode(dbuf *DecodeBuf) error {
	m.LangCode = dbuf.String()
	m.FromVersion = dbuf.Int()
	m.Version = dbuf.Int()
	// x.VectorMessage(m.Strings);
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.Strings = make([]*LangPackString, l4)
	for i := 0; i < int(l4); i++ {
		m.Strings[i] = &LangPackString{}
		(*m.Strings[i]).Decode(dbuf)
	}
	return dbuf.err
}

// langPackLanguage#117698f1 name:string native_name:string lang_code:string = LangPackLanguage;
func (m *TLLangPackLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackLanguage))
	x.String(m.Name)
	x.String(m.NativeName)
	x.String(m.LangCode)
	return x.buf
}

func (m *TLLangPackLanguage) Decode(dbuf *DecodeBuf) error {
	m.Name = dbuf.String()
	m.NativeName = dbuf.String()
	m.LangCode = dbuf.String()
	return dbuf.err
}

// channelAdminRights#5d7ceba5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true invite_link:flags.6?true pin_messages:flags.7?true add_admins:flags.9?true = ChannelAdminRights;
func (m *TLChannelAdminRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminRights))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLChannelAdminRights) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// channelBannedRights#58cf4249 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true until_date:int = ChannelBannedRights;
func (m *TLChannelBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelBannedRights))
	x.Int(m.Flags)
	x.Int(m.UntilDate)
	return x.buf
}

func (m *TLChannelBannedRights) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.UntilDate = dbuf.Int()
	return dbuf.err
}

// channelAdminLogEventActionChangeTitle#e6dfb825 prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeTitle))
	x.String(m.PrevValue)
	x.String(m.NewValue)
	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeTitle) Decode(dbuf *DecodeBuf) error {
	m.PrevValue = dbuf.String()
	m.NewValue = dbuf.String()
	return dbuf.err
}

// channelAdminLogEventActionChangeAbout#55188a2e prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeAbout))
	x.String(m.PrevValue)
	x.String(m.NewValue)
	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeAbout) Decode(dbuf *DecodeBuf) error {
	m.PrevValue = dbuf.String()
	m.NewValue = dbuf.String()
	return dbuf.err
}

// channelAdminLogEventActionChangeUsername#6a4afc38 prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeUsername))
	x.String(m.PrevValue)
	x.String(m.NewValue)
	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeUsername) Decode(dbuf *DecodeBuf) error {
	m.PrevValue = dbuf.String()
	m.NewValue = dbuf.String()
	return dbuf.err
}

// channelAdminLogEventActionChangePhoto#b82f55c3 prev_photo:ChatPhoto new_photo:ChatPhoto = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangePhoto))
	x.Bytes(m.PrevPhoto.Encode())
	x.Bytes(m.NewPhoto.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionChangePhoto) Decode(dbuf *DecodeBuf) error {
	m.PrevPhoto = &ChatPhoto{}
	m.Decode(dbuf)
	m.NewPhoto = &ChatPhoto{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionToggleInvites#1b7907ae new_value:Bool = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionToggleInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionToggleInvites))
	x.Bytes(m.NewValue.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionToggleInvites) Decode(dbuf *DecodeBuf) error {
	m.NewValue = &Bool{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionToggleSignatures#26ae0971 new_value:Bool = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionToggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionToggleSignatures))
	x.Bytes(m.NewValue.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionToggleSignatures) Decode(dbuf *DecodeBuf) error {
	m.NewValue = &Bool{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionUpdatePinned#e9e82c18 message:Message = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionUpdatePinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionUpdatePinned))
	x.Bytes(m.Message.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionUpdatePinned) Decode(dbuf *DecodeBuf) error {
	m.Message = &Message{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionEditMessage#709b2405 prev_message:Message new_message:Message = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionEditMessage))
	x.Bytes(m.PrevMessage.Encode())
	x.Bytes(m.NewMessage.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionEditMessage) Decode(dbuf *DecodeBuf) error {
	m.PrevMessage = &Message{}
	m.Decode(dbuf)
	m.NewMessage = &Message{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionDeleteMessage#42e047bb message:Message = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionDeleteMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionDeleteMessage))
	x.Bytes(m.Message.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionDeleteMessage) Decode(dbuf *DecodeBuf) error {
	m.Message = &Message{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionParticipantJoin#183040d3 = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantJoin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantJoin))
	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantJoin) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channelAdminLogEventActionParticipantLeave#f89777f2 = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantLeave) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantLeave))
	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantLeave) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channelAdminLogEventActionParticipantInvite#e31c34d8 participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantInvite))
	x.Bytes(m.Participant.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantInvite) Decode(dbuf *DecodeBuf) error {
	m.Participant = &ChannelParticipant{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionParticipantToggleBan#e6d83d7e prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantToggleBan) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleBan))
	x.Bytes(m.PrevParticipant.Encode())
	x.Bytes(m.NewParticipant.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) Decode(dbuf *DecodeBuf) error {
	m.PrevParticipant = &ChannelParticipant{}
	m.Decode(dbuf)
	m.NewParticipant = &ChannelParticipant{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionParticipantToggleAdmin#d5676710 prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleAdmin))
	x.Bytes(m.PrevParticipant.Encode())
	x.Bytes(m.NewParticipant.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) Decode(dbuf *DecodeBuf) error {
	m.PrevParticipant = &ChannelParticipant{}
	m.Decode(dbuf)
	m.NewParticipant = &ChannelParticipant{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEventActionChangeStickerSet#b1c3caa7 prev_stickerset:InputStickerSet new_stickerset:InputStickerSet = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeStickerSet))
	x.Bytes(m.PrevStickerset.Encode())
	x.Bytes(m.NewStickerset.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) Decode(dbuf *DecodeBuf) error {
	m.PrevStickerset = &InputStickerSet{}
	m.Decode(dbuf)
	m.NewStickerset = &InputStickerSet{}
	m.Decode(dbuf)
	return dbuf.err
}

// channelAdminLogEvent#3b5a3e40 id:long date:int user_id:int action:ChannelAdminLogEventAction = ChannelAdminLogEvent;
func (m *TLChannelAdminLogEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEvent))
	x.Long(m.Id)
	x.Int(m.Date)
	x.Int(m.UserId)
	x.Bytes(m.Action.Encode())
	return x.buf
}

func (m *TLChannelAdminLogEvent) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Long()
	m.Date = dbuf.Int()
	m.UserId = dbuf.Int()
	m.Action = &ChannelAdminLogEventAction{}
	m.Decode(dbuf)
	return dbuf.err
}

// channels.adminLogResults#ed8af74d events:Vector<ChannelAdminLogEvent> chats:Vector<Chat> users:Vector<User> = channels.AdminLogResults;
func (m *TLChannelsAdminLogResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_adminLogResults))
	// x.VectorMessage(m.Events);
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Events)))
	for _, v := range m.Events {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Chats);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	for _, v := range m.Chats {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Users);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChannelsAdminLogResults) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Events);
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Events = make([]*ChannelAdminLogEvent, l1)
	for i := 0; i < int(l1); i++ {
		m.Events[i] = &ChannelAdminLogEvent{}
		(*m.Events[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Chats);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Chats = make([]*Chat, l2)
	for i := 0; i < int(l2); i++ {
		m.Chats[i] = &Chat{}
		(*m.Chats[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Users);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Users = make([]*User, l3)
	for i := 0; i < int(l3); i++ {
		m.Users[i] = &User{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channelAdminLogEventsFilter#ea107ae4 flags:# join:flags.0?true leave:flags.1?true invite:flags.2?true ban:flags.3?true unban:flags.4?true kick:flags.5?true unkick:flags.6?true promote:flags.7?true demote:flags.8?true info:flags.9?true settings:flags.10?true pinned:flags.11?true edit:flags.12?true delete:flags.13?true = ChannelAdminLogEventsFilter;
func (m *TLChannelAdminLogEventsFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventsFilter))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLChannelAdminLogEventsFilter) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// popularContact#5ce14175 client_id:long importers:int = PopularContact;
func (m *TLPopularContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_popularContact))
	x.Long(m.ClientId)
	x.Int(m.Importers)
	return x.buf
}

func (m *TLPopularContact) Decode(dbuf *DecodeBuf) error {
	m.ClientId = dbuf.Long()
	m.Importers = dbuf.Int()
	return dbuf.err
}

// cdnFileHash#77eec38f offset:int limit:int hash:bytes = CdnFileHash;
func (m *TLCdnFileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_cdnFileHash))
	x.Int(m.Offset)
	x.Int(m.Limit)
	x.StringBytes(m.Hash)
	return x.buf
}

func (m *TLCdnFileHash) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	m.Hash = dbuf.StringBytes()
	return dbuf.err
}

// messages.favedStickersNotModified#9e8fa6d3 = messages.FavedStickers;
func (m *TLMessagesFavedStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_favedStickersNotModified))
	return x.buf
}

func (m *TLMessagesFavedStickersNotModified) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.favedStickers#f37f2f16 hash:int packs:Vector<StickerPack> stickers:Vector<Document> = messages.FavedStickers;
func (m *TLMessagesFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_favedStickers))
	x.Int(m.Hash)
	// x.VectorMessage(m.Packs);
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Packs)))
	for _, v := range m.Packs {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	// x.VectorMessage(m.Stickers);
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Stickers)))
	for _, v := range m.Stickers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesFavedStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	// x.VectorMessage(m.Packs);
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Packs = make([]*StickerPack, l2)
	for i := 0; i < int(l2); i++ {
		m.Packs[i] = &StickerPack{}
		(*m.Packs[i]).Decode(dbuf)
	}
	// x.VectorMessage(m.Stickers);
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Stickers = make([]*Document, l3)
	for i := 0; i < int(l3); i++ {
		m.Stickers[i] = &Document{}
		(*m.Stickers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// RPC
// req_pq#60469778 nonce:int128 = ResPQ;
func (m *TLReqPq) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_req_pq))
	x.Bytes(m.Nonce)
	return x.buf
}

func (m *TLReqPq) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	return dbuf.err
}

// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
func (m *TLReq_DHParams) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_req_DH_params))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.String(m.P)
	x.String(m.Q)
	x.Long(m.PublicKeyFingerprint)
	x.String(m.EncryptedData)
	return x.buf
}

func (m *TLReq_DHParams) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.P = dbuf.String()
	m.Q = dbuf.String()
	m.PublicKeyFingerprint = dbuf.Long()
	m.EncryptedData = dbuf.String()
	return dbuf.err
}

// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
func (m *TLSetClient_DHParams) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_set_client_DH_params))
	x.Bytes(m.Nonce)
	x.Bytes(m.ServerNonce)
	x.String(m.EncryptedData)
	return x.buf
}

func (m *TLSetClient_DHParams) Decode(dbuf *DecodeBuf) error {
	m.Nonce = dbuf.Bytes(16)
	m.ServerNonce = dbuf.Bytes(16)
	m.EncryptedData = dbuf.String()
	return dbuf.err
}

// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
func (m *TLDestroyAuthKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_auth_key))
	return x.buf
}

func (m *TLDestroyAuthKey) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// rpc_drop_answer#58e4a740 req_msg_id:long = RpcDropAnswer;
func (m *TLRpcDropAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_rpc_drop_answer))
	x.Long(m.ReqMsgId)
	return x.buf
}

func (m *TLRpcDropAnswer) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	return dbuf.err
}

// get_future_salts#b921bd04 num:int = FutureSalts;
func (m *TLGetFutureSalts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_get_future_salts))
	x.Int(m.Num)
	return x.buf
}

func (m *TLGetFutureSalts) Decode(dbuf *DecodeBuf) error {
	m.Num = dbuf.Int()
	return dbuf.err
}

// ping#7abe77ec ping_id:long = Pong;
func (m *TLPing) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_ping))
	x.Long(m.PingId)
	return x.buf
}

func (m *TLPing) Decode(dbuf *DecodeBuf) error {
	m.PingId = dbuf.Long()
	return dbuf.err
}

// ping_delay_disconnect#f3427b8c ping_id:long disconnect_delay:int = Pong;
func (m *TLPingDelayDisconnect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_ping_delay_disconnect))
	x.Long(m.PingId)
	x.Int(m.DisconnectDelay)
	return x.buf
}

func (m *TLPingDelayDisconnect) Decode(dbuf *DecodeBuf) error {
	m.PingId = dbuf.Long()
	m.DisconnectDelay = dbuf.Int()
	return dbuf.err
}

// destroy_session#e7512126 session_id:long = DestroySessionRes;
func (m *TLDestroySession) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_destroy_session))
	x.Long(m.SessionId)
	return x.buf
}

func (m *TLDestroySession) Decode(dbuf *DecodeBuf) error {
	m.SessionId = dbuf.Long()
	return dbuf.err
}

// contest.saveDeveloperInfo#9a5f6e95 vk_id:int name:string phone_number:string age:int city:string = Bool;
func (m *TLContestSaveDeveloperInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contest_saveDeveloperInfo))
	x.Int(m.VkId)
	x.String(m.Name)
	x.String(m.PhoneNumber)
	x.Int(m.Age)
	x.String(m.City)
	return x.buf
}

func (m *TLContestSaveDeveloperInfo) Decode(dbuf *DecodeBuf) error {
	m.VkId = dbuf.Int()
	m.Name = dbuf.String()
	m.PhoneNumber = dbuf.String()
	m.Age = dbuf.Int()
	m.City = dbuf.String()
	return dbuf.err
}

// auth.logOut#5717da40 = Bool;
func (m *TLAuthLogOut) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_logOut))
	return x.buf
}

func (m *TLAuthLogOut) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// auth.resetAuthorizations#9fab0d1a = Bool;
func (m *TLAuthResetAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_resetAuthorizations))
	return x.buf
}

func (m *TLAuthResetAuthorizations) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// auth.sendInvites#771c1d97 phone_numbers:Vector<string> message:string = Bool;
func (m *TLAuthSendInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sendInvites))
	x.VectorString(m.PhoneNumbers)
	x.String(m.Message)
	return x.buf
}

func (m *TLAuthSendInvites) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumbers = dbuf.VectorString()
	m.Message = dbuf.String()
	return dbuf.err
}

// auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (m *TLAuthBindTempAuthKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_bindTempAuthKey))
	x.Long(m.PermAuthKeyId)
	x.Long(m.Nonce)
	x.Int(m.ExpiresAt)
	x.StringBytes(m.EncryptedMessage)
	return x.buf
}

func (m *TLAuthBindTempAuthKey) Decode(dbuf *DecodeBuf) error {
	m.PermAuthKeyId = dbuf.Long()
	m.Nonce = dbuf.Long()
	m.ExpiresAt = dbuf.Int()
	m.EncryptedMessage = dbuf.StringBytes()
	return dbuf.err
}

// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (m *TLAuthCancelCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_cancelCode))
	x.String(m.PhoneNumber)
	x.String(m.PhoneCodeHash)
	return x.buf
}

func (m *TLAuthCancelCode) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.PhoneCodeHash = dbuf.String()
	return dbuf.err
}

// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
func (m *TLAuthDropTempAuthKeys) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_dropTempAuthKeys))
	x.VectorLong(m.ExceptAuthKeys)
	return x.buf
}

func (m *TLAuthDropTempAuthKeys) Decode(dbuf *DecodeBuf) error {
	m.ExceptAuthKeys = dbuf.VectorLong()
	return dbuf.err
}

// account.registerDevice#637ea878 token_type:int token:string = Bool;
func (m *TLAccountRegisterDevice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_registerDevice))
	x.Int(m.TokenType)
	x.String(m.Token)
	return x.buf
}

func (m *TLAccountRegisterDevice) Decode(dbuf *DecodeBuf) error {
	m.TokenType = dbuf.Int()
	m.Token = dbuf.String()
	return dbuf.err
}

// account.unregisterDevice#65c55b40 token_type:int token:string = Bool;
func (m *TLAccountUnregisterDevice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_unregisterDevice))
	x.Int(m.TokenType)
	x.String(m.Token)
	return x.buf
}

func (m *TLAccountUnregisterDevice) Decode(dbuf *DecodeBuf) error {
	m.TokenType = dbuf.Int()
	m.Token = dbuf.String()
	return dbuf.err
}

// account.updateNotifySettings#84be5b93 peer:InputNotifyPeer settings:InputPeerNotifySettings = Bool;
func (m *TLAccountUpdateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_updateNotifySettings))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Settings.Encode())
	return x.buf
}

func (m *TLAccountUpdateNotifySettings) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.Settings.Encode())
	return dbuf.err
}

// account.resetNotifySettings#db7e1747 = Bool;
func (m *TLAccountResetNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_resetNotifySettings))
	return x.buf
}

func (m *TLAccountResetNotifySettings) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// account.updateStatus#6628562c offline:Bool = Bool;
func (m *TLAccountUpdateStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_updateStatus))
	x.Bytes(m.Offline.Encode())
	return x.buf
}

func (m *TLAccountUpdateStatus) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Offline.Encode())
	return dbuf.err
}

// account.reportPeer#ae189d5f peer:InputPeer reason:ReportReason = Bool;
func (m *TLAccountReportPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_reportPeer))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Reason.Encode())
	return x.buf
}

func (m *TLAccountReportPeer) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.Reason.Encode())
	return dbuf.err
}

// account.checkUsername#2714d86c username:string = Bool;
func (m *TLAccountCheckUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_checkUsername))
	x.String(m.Username)
	return x.buf
}

func (m *TLAccountCheckUsername) Decode(dbuf *DecodeBuf) error {
	m.Username = dbuf.String()
	return dbuf.err
}

// account.deleteAccount#418d4e0b reason:string = Bool;
func (m *TLAccountDeleteAccount) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_deleteAccount))
	x.String(m.Reason)
	return x.buf
}

func (m *TLAccountDeleteAccount) Decode(dbuf *DecodeBuf) error {
	m.Reason = dbuf.String()
	return dbuf.err
}

// account.setAccountTTL#2442485e ttl:AccountDaysTTL = Bool;
func (m *TLAccountSetAccountTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_setAccountTTL))
	x.Bytes(m.Ttl.Encode())
	return x.buf
}

func (m *TLAccountSetAccountTTL) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Ttl.Encode())
	return dbuf.err
}

// account.updateDeviceLocked#38df3532 period:int = Bool;
func (m *TLAccountUpdateDeviceLocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_updateDeviceLocked))
	x.Int(m.Period)
	return x.buf
}

func (m *TLAccountUpdateDeviceLocked) Decode(dbuf *DecodeBuf) error {
	m.Period = dbuf.Int()
	return dbuf.err
}

// account.resetAuthorization#df77f3bc hash:long = Bool;
func (m *TLAccountResetAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_resetAuthorization))
	x.Long(m.Hash)
	return x.buf
}

func (m *TLAccountResetAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Long()
	return dbuf.err
}

// account.updatePasswordSettings#fa7c4b86 current_password_hash:bytes new_settings:account.PasswordInputSettings = Bool;
func (m *TLAccountUpdatePasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_updatePasswordSettings))
	x.StringBytes(m.CurrentPasswordHash)
	x.Bytes(m.NewSettings.Encode())
	return x.buf
}

func (m *TLAccountUpdatePasswordSettings) Decode(dbuf *DecodeBuf) error {
	m.CurrentPasswordHash = dbuf.StringBytes()
	// x.Bytes(m.NewSettings.Encode())
	return dbuf.err
}

// account.confirmPhone#5f2178c3 phone_code_hash:string phone_code:string = Bool;
func (m *TLAccountConfirmPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_confirmPhone))
	x.String(m.PhoneCodeHash)
	x.String(m.PhoneCode)
	return x.buf
}

func (m *TLAccountConfirmPhone) Decode(dbuf *DecodeBuf) error {
	m.PhoneCodeHash = dbuf.String()
	m.PhoneCode = dbuf.String()
	return dbuf.err
}

// contacts.deleteContacts#59ab389e id:Vector<InputUser> = Bool;
func (m *TLContactsDeleteContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_deleteContacts))
	// x.VectorMessage(m.Id)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Id)))
	for _, v := range m.Id {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsDeleteContacts) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Id)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Id = make([]*InputUser, l1)
	for i := 0; i < int(l1); i++ {
		m.Id[i] = &InputUser{}
		(*m.Id[i]).Decode(dbuf)
	}
	return dbuf.err
}

// contacts.block#332b49fc id:InputUser = Bool;
func (m *TLContactsBlock) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_block))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLContactsBlock) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	return dbuf.err
}

// contacts.unblock#e54100bd id:InputUser = Bool;
func (m *TLContactsUnblock) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_unblock))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLContactsUnblock) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	return dbuf.err
}

// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (m *TLContactsResetTopPeerRating) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_resetTopPeerRating))
	x.Bytes(m.Category.Encode())
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLContactsResetTopPeerRating) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Category.Encode())
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// contacts.resetSaved#879537f1 = Bool;
func (m *TLContactsResetSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_resetSaved))
	return x.buf
}

func (m *TLContactsResetSaved) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.setTyping#a3825e50 peer:InputPeer action:SendMessageAction = Bool;
func (m *TLMessagesSetTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setTyping))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Action.Encode())
	return x.buf
}

func (m *TLMessagesSetTyping) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.Action.Encode())
	return dbuf.err
}

// messages.reportSpam#cf1592db peer:InputPeer = Bool;
func (m *TLMessagesReportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_reportSpam))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLMessagesReportSpam) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// messages.hideReportSpam#a8f1709b peer:InputPeer = Bool;
func (m *TLMessagesHideReportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_hideReportSpam))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLMessagesHideReportSpam) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// messages.discardEncryption#edd923c5 chat_id:int = Bool;
func (m *TLMessagesDiscardEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_discardEncryption))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLMessagesDiscardEncryption) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// messages.setEncryptedTyping#791451ed peer:InputEncryptedChat typing:Bool = Bool;
func (m *TLMessagesSetEncryptedTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setEncryptedTyping))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Typing.Encode())
	return x.buf
}

func (m *TLMessagesSetEncryptedTyping) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.Typing.Encode())
	return dbuf.err
}

// messages.readEncryptedHistory#7f4b690a peer:InputEncryptedChat max_date:int = Bool;
func (m *TLMessagesReadEncryptedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_readEncryptedHistory))
	x.Bytes(m.Peer.Encode())
	x.Int(m.MaxDate)
	return x.buf
}

func (m *TLMessagesReadEncryptedHistory) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.MaxDate = dbuf.Int()
	return dbuf.err
}

// messages.reportEncryptedSpam#4b0c8c0f peer:InputEncryptedChat = Bool;
func (m *TLMessagesReportEncryptedSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_reportEncryptedSpam))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLMessagesReportEncryptedSpam) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// messages.uninstallStickerSet#f96e55de stickerset:InputStickerSet = Bool;
func (m *TLMessagesUninstallStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_uninstallStickerSet))
	x.Bytes(m.Stickerset.Encode())
	return x.buf
}

func (m *TLMessagesUninstallStickerSet) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Stickerset.Encode())
	return dbuf.err
}

// messages.editChatAdmin#a9e69f2e chat_id:int user_id:InputUser is_admin:Bool = Bool;
func (m *TLMessagesEditChatAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_editChatAdmin))
	x.Int(m.ChatId)
	x.Bytes(m.UserId.Encode())
	x.Bytes(m.IsAdmin.Encode())
	return x.buf
}

func (m *TLMessagesEditChatAdmin) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	// x.Bytes(m.UserId.Encode())
	// x.Bytes(m.IsAdmin.Encode())
	return dbuf.err
}

// messages.reorderStickerSets#78337739 flags:# masks:flags.0?true order:Vector<long> = Bool;
func (m *TLMessagesReorderStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_reorderStickerSets))
	x.Int(m.Flags)
	x.VectorLong(m.Order)
	return x.buf
}

func (m *TLMessagesReorderStickerSets) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Order = dbuf.VectorLong()
	return dbuf.err
}

// messages.saveGif#327a30cb id:InputDocument unsave:Bool = Bool;
func (m *TLMessagesSaveGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_saveGif))
	x.Bytes(m.Id.Encode())
	x.Bytes(m.Unsave.Encode())
	return x.buf
}

func (m *TLMessagesSaveGif) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	// x.Bytes(m.Unsave.Encode())
	return dbuf.err
}

// messages.setInlineBotResults#eb5ea206 flags:# gallery:flags.0?true private:flags.1?true query_id:long results:Vector<InputBotInlineResult> cache_time:int next_offset:flags.2?string switch_pm:flags.3?InlineBotSwitchPM = Bool;
func (m *TLMessagesSetInlineBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setInlineBotResults))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	// x.VectorMessage(m.Results)
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Results)))
	for _, v := range m.Results {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Int(m.CacheTime)
	x.String(m.NextOffset)
	x.Bytes(m.SwitchPm.Encode())
	return x.buf
}

func (m *TLMessagesSetInlineBotResults) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	// x.VectorMessage(m.Results)
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Results = make([]*InputBotInlineResult, l5)
	for i := 0; i < int(l5); i++ {
		m.Results[i] = &InputBotInlineResult{}
		(*m.Results[i]).Decode(dbuf)
	}
	m.CacheTime = dbuf.Int()
	m.NextOffset = dbuf.String()
	// x.Bytes(m.SwitchPm.Encode())
	return dbuf.err
}

// messages.editInlineBotMessage#130c2c85 flags:# no_webpage:flags.1?true id:InputBotInlineMessageID message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Bool;
func (m *TLMessagesEditInlineBotMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_editInlineBotMessage))
	x.Int(m.Flags)
	x.Bytes(m.Id.Encode())
	x.String(m.Message)
	x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities)
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesEditInlineBotMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Id.Encode())
	m.Message = dbuf.String()
	// x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities)
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l6)
	for i := 0; i < int(l6); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.setBotCallbackAnswer#d58f130a flags:# alert:flags.1?true query_id:long message:flags.0?string url:flags.2?string cache_time:int = Bool;
func (m *TLMessagesSetBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setBotCallbackAnswer))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.String(m.Message)
	x.String(m.Url)
	x.Int(m.CacheTime)
	return x.buf
}

func (m *TLMessagesSetBotCallbackAnswer) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.Message = dbuf.String()
	m.Url = dbuf.String()
	m.CacheTime = dbuf.Int()
	return dbuf.err
}

// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (m *TLMessagesSaveDraft) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_saveDraft))
	x.Int(m.Flags)
	x.Int(m.ReplyToMsgId)
	x.Bytes(m.Peer.Encode())
	x.String(m.Message)
	// x.VectorMessage(m.Entities)
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesSaveDraft) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.ReplyToMsgId = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.Message = dbuf.String()
	// x.VectorMessage(m.Entities)
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l6)
	for i := 0; i < int(l6); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.readFeaturedStickers#5b118126 id:Vector<long> = Bool;
func (m *TLMessagesReadFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_readFeaturedStickers))
	x.VectorLong(m.Id)
	return x.buf
}

func (m *TLMessagesReadFeaturedStickers) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.VectorLong()
	return dbuf.err
}

// messages.saveRecentSticker#392718f8 flags:# attached:flags.0?true id:InputDocument unsave:Bool = Bool;
func (m *TLMessagesSaveRecentSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_saveRecentSticker))
	x.Int(m.Flags)
	x.Bytes(m.Id.Encode())
	x.Bytes(m.Unsave.Encode())
	return x.buf
}

func (m *TLMessagesSaveRecentSticker) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Id.Encode())
	// x.Bytes(m.Unsave.Encode())
	return dbuf.err
}

// messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool;
func (m *TLMessagesClearRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_clearRecentStickers))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLMessagesClearRecentStickers) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// messages.setInlineGameScore#15ad9f64 flags:# edit_message:flags.0?true force:flags.1?true id:InputBotInlineMessageID user_id:InputUser score:int = Bool;
func (m *TLMessagesSetInlineGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setInlineGameScore))
	x.Int(m.Flags)
	x.Bytes(m.Id.Encode())
	x.Bytes(m.UserId.Encode())
	x.Int(m.Score)
	return x.buf
}

func (m *TLMessagesSetInlineGameScore) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Id.Encode())
	// x.Bytes(m.UserId.Encode())
	m.Score = dbuf.Int()
	return dbuf.err
}

// messages.toggleDialogPin#3289be6a flags:# pinned:flags.0?true peer:InputPeer = Bool;
func (m *TLMessagesToggleDialogPin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_toggleDialogPin))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLMessagesToggleDialogPin) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// messages.reorderPinnedDialogs#959ff644 flags:# force:flags.0?true order:Vector<InputPeer> = Bool;
func (m *TLMessagesReorderPinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_reorderPinnedDialogs))
	x.Int(m.Flags)
	// x.VectorMessage(m.Order)
	x3 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Order)))
	for _, v := range m.Order {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesReorderPinnedDialogs) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.VectorMessage(m.Order)
	c3 := dbuf.Int()
	if c3 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c3)
	}
	l3 := dbuf.Int()
	m.Order = make([]*InputPeer, l3)
	for i := 0; i < int(l3); i++ {
		m.Order[i] = &InputPeer{}
		(*m.Order[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.setBotShippingResults#e5f672fa flags:# query_id:long error:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = Bool;
func (m *TLMessagesSetBotShippingResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setBotShippingResults))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.String(m.Error)
	// x.VectorMessage(m.ShippingOptions)
	x4 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x4, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x4[4:], uint32(len(m.ShippingOptions)))
	for _, v := range m.ShippingOptions {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesSetBotShippingResults) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.Error = dbuf.String()
	// x.VectorMessage(m.ShippingOptions)
	c4 := dbuf.Int()
	if c4 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c4)
	}
	l4 := dbuf.Int()
	m.ShippingOptions = make([]*ShippingOption, l4)
	for i := 0; i < int(l4); i++ {
		m.ShippingOptions[i] = &ShippingOption{}
		(*m.ShippingOptions[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.setBotPrecheckoutResults#9c2dd95 flags:# success:flags.1?true query_id:long error:flags.0?string = Bool;
func (m *TLMessagesSetBotPrecheckoutResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setBotPrecheckoutResults))
	x.Int(m.Flags)
	x.Long(m.QueryId)
	x.String(m.Error)
	return x.buf
}

func (m *TLMessagesSetBotPrecheckoutResults) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.QueryId = dbuf.Long()
	m.Error = dbuf.String()
	return dbuf.err
}

// messages.faveSticker#b9ffc55b id:InputDocument unfave:Bool = Bool;
func (m *TLMessagesFaveSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_faveSticker))
	x.Bytes(m.Id.Encode())
	x.Bytes(m.Unfave.Encode())
	return x.buf
}

func (m *TLMessagesFaveSticker) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	// x.Bytes(m.Unfave.Encode())
	return dbuf.err
}

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (m *TLUploadSaveFilePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_saveFilePart))
	x.Long(m.FileId)
	x.Int(m.FilePart)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLUploadSaveFilePart) Decode(dbuf *DecodeBuf) error {
	m.FileId = dbuf.Long()
	m.FilePart = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (m *TLUploadSaveBigFilePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_saveBigFilePart))
	x.Long(m.FileId)
	x.Int(m.FilePart)
	x.Int(m.FileTotalParts)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLUploadSaveBigFilePart) Decode(dbuf *DecodeBuf) error {
	m.FileId = dbuf.Long()
	m.FilePart = dbuf.Int()
	m.FileTotalParts = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// help.saveAppLog#6f02f748 events:Vector<InputAppEvent> = Bool;
func (m *TLHelpSaveAppLog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_saveAppLog))
	// x.VectorMessage(m.Events)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Events)))
	for _, v := range m.Events {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLHelpSaveAppLog) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Events)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Events = make([]*InputAppEvent, l1)
	for i := 0; i < int(l1); i++ {
		m.Events[i] = &InputAppEvent{}
		(*m.Events[i]).Decode(dbuf)
	}
	return dbuf.err
}

// help.setBotUpdatesStatus#ec22cfcd pending_updates_count:int message:string = Bool;
func (m *TLHelpSetBotUpdatesStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_setBotUpdatesStatus))
	x.Int(m.PendingUpdatesCount)
	x.String(m.Message)
	return x.buf
}

func (m *TLHelpSetBotUpdatesStatus) Decode(dbuf *DecodeBuf) error {
	m.PendingUpdatesCount = dbuf.Int()
	m.Message = dbuf.String()
	return dbuf.err
}

// channels.readHistory#cc104937 channel:InputChannel max_id:int = Bool;
func (m *TLChannelsReadHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_readHistory))
	x.Bytes(m.Channel.Encode())
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLChannelsReadHistory) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// channels.reportSpam#fe087810 channel:InputChannel user_id:InputUser id:Vector<int> = Bool;
func (m *TLChannelsReportSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_reportSpam))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.UserId.Encode())
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLChannelsReportSpam) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.UserId.Encode())
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// channels.editAbout#13e27f1e channel:InputChannel about:string = Bool;
func (m *TLChannelsEditAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_editAbout))
	x.Bytes(m.Channel.Encode())
	x.String(m.About)
	return x.buf
}

func (m *TLChannelsEditAbout) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.About = dbuf.String()
	return dbuf.err
}

// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (m *TLChannelsCheckUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_checkUsername))
	x.Bytes(m.Channel.Encode())
	x.String(m.Username)
	return x.buf
}

func (m *TLChannelsCheckUsername) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Username = dbuf.String()
	return dbuf.err
}

// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (m *TLChannelsUpdateUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_updateUsername))
	x.Bytes(m.Channel.Encode())
	x.String(m.Username)
	return x.buf
}

func (m *TLChannelsUpdateUsername) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Username = dbuf.String()
	return dbuf.err
}

// channels.setStickers#ea8ca4f9 channel:InputChannel stickerset:InputStickerSet = Bool;
func (m *TLChannelsSetStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_setStickers))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.Stickerset.Encode())
	return x.buf
}

func (m *TLChannelsSetStickers) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.Stickerset.Encode())
	return dbuf.err
}

// channels.readMessageContents#eab5dc38 channel:InputChannel id:Vector<int> = Bool;
func (m *TLChannelsReadMessageContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_readMessageContents))
	x.Bytes(m.Channel.Encode())
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLChannelsReadMessageContents) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// bots.answerWebhookJSONQuery#e6213f4d query_id:long data:DataJSON = Bool;
func (m *TLBotsAnswerWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_bots_answerWebhookJSONQuery))
	x.Long(m.QueryId)
	x.Bytes(m.Data.Encode())
	return x.buf
}

func (m *TLBotsAnswerWebhookJSONQuery) Decode(dbuf *DecodeBuf) error {
	m.QueryId = dbuf.Long()
	// x.Bytes(m.Data.Encode())
	return dbuf.err
}

// payments.clearSavedInfo#d83d70c1 flags:# credentials:flags.0?true info:flags.1?true = Bool;
func (m *TLPaymentsClearSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_clearSavedInfo))
	x.Int(m.Flags)
	return x.buf
}

func (m *TLPaymentsClearSavedInfo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	return dbuf.err
}

// phone.receivedCall#17d54f61 peer:InputPhoneCall = Bool;
func (m *TLPhoneReceivedCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_receivedCall))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLPhoneReceivedCall) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// phone.saveCallDebug#277add7e peer:InputPhoneCall debug:DataJSON = Bool;
func (m *TLPhoneSaveCallDebug) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_saveCallDebug))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Debug.Encode())
	return x.buf
}

func (m *TLPhoneSaveCallDebug) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.Debug.Encode())
	return dbuf.err
}

// invokeAfterMsg#cb9f372d {X:Type} msg_id:long query:!X = X;
func (m *TLInvokeAfterMsg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_invokeAfterMsg))
	x.Long(m.MsgId)
	x.Bytes(m.Query)
	return x.buf
}

func (m *TLInvokeAfterMsg) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好
	o2 := dbuf.Object()
	m.Query = o2.Encode()
	return dbuf.err
}

// invokeAfterMsgs#3dc4b4f0 {X:Type} msg_ids:Vector<long> query:!X = X;
func (m *TLInvokeAfterMsgs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_invokeAfterMsgs))
	x.VectorLong(m.MsgIds)
	x.Bytes(m.Query)
	return x.buf
}

func (m *TLInvokeAfterMsgs) Decode(dbuf *DecodeBuf) error {
	m.MsgIds = dbuf.VectorLong()
	// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好
	o2 := dbuf.Object()
	m.Query = o2.Encode()
	return dbuf.err
}

// initConnection#c7481da6 {X:Type} api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string query:!X = X;
func (m *TLInitConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_initConnection))
	x.Int(m.ApiId)
	x.String(m.DeviceModel)
	x.String(m.SystemVersion)
	x.String(m.AppVersion)
	x.String(m.SystemLangCode)
	x.String(m.LangPack)
	x.String(m.LangCode)
	x.Bytes(m.Query)
	return x.buf
}

func (m *TLInitConnection) Decode(dbuf *DecodeBuf) error {
	m.ApiId = dbuf.Int()
	m.DeviceModel = dbuf.String()
	m.SystemVersion = dbuf.String()
	m.AppVersion = dbuf.String()
	m.SystemLangCode = dbuf.String()
	m.LangPack = dbuf.String()
	m.LangCode = dbuf.String()
	// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好
	o8 := dbuf.Object()
	m.Query = o8.Encode()
	return dbuf.err
}

// invokeWithLayer#da9b0d0d {X:Type} layer:int query:!X = X;
func (m *TLInvokeWithLayer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_invokeWithLayer))
	x.Int(m.Layer)
	x.Bytes(m.Query)
	return x.buf
}

func (m *TLInvokeWithLayer) Decode(dbuf *DecodeBuf) error {
	m.Layer = dbuf.Int()
	// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好
	o2 := dbuf.Object()
	m.Query = o2.Encode()
	return dbuf.err
}

// invokeWithoutUpdates#bf9459b7 {X:Type} query:!X = X;
func (m *TLInvokeWithoutUpdates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_invokeWithoutUpdates))
	x.Bytes(m.Query)
	return x.buf
}

func (m *TLInvokeWithoutUpdates) Decode(dbuf *DecodeBuf) error {
	// TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好
	o1 := dbuf.Object()
	m.Query = o1.Encode()
	return dbuf.err
}

// auth.checkPhone#6fe51dfb phone_number:string = auth.CheckedPhone;
func (m *TLAuthCheckPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_checkPhone))
	x.String(m.PhoneNumber)
	return x.buf
}

func (m *TLAuthCheckPhone) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	return dbuf.err
}

// auth.sendCode#86aef0ec flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool api_id:int api_hash:string = auth.SentCode;
func (m *TLAuthSendCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sendCode))
	x.Int(m.Flags)
	x.String(m.PhoneNumber)
	x.Bytes(m.CurrentNumber.Encode())
	x.Int(m.ApiId)
	x.String(m.ApiHash)
	return x.buf
}

func (m *TLAuthSendCode) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.PhoneNumber = dbuf.String()
	// x.Bytes(m.CurrentNumber.Encode())
	m.ApiId = dbuf.Int()
	m.ApiHash = dbuf.String()
	return dbuf.err
}

// auth.resendCode#3ef1a9bf phone_number:string phone_code_hash:string = auth.SentCode;
func (m *TLAuthResendCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_resendCode))
	x.String(m.PhoneNumber)
	x.String(m.PhoneCodeHash)
	return x.buf
}

func (m *TLAuthResendCode) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.PhoneCodeHash = dbuf.String()
	return dbuf.err
}

// account.sendChangePhoneCode#8e57deb flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool = auth.SentCode;
func (m *TLAccountSendChangePhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_sendChangePhoneCode))
	x.Int(m.Flags)
	x.String(m.PhoneNumber)
	x.Bytes(m.CurrentNumber.Encode())
	return x.buf
}

func (m *TLAccountSendChangePhoneCode) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.PhoneNumber = dbuf.String()
	// x.Bytes(m.CurrentNumber.Encode())
	return dbuf.err
}

// account.sendConfirmPhoneCode#1516d7bd flags:# allow_flashcall:flags.0?true hash:string current_number:flags.0?Bool = auth.SentCode;
func (m *TLAccountSendConfirmPhoneCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_sendConfirmPhoneCode))
	x.Int(m.Flags)
	x.String(m.Hash)
	x.Bytes(m.CurrentNumber.Encode())
	return x.buf
}

func (m *TLAccountSendConfirmPhoneCode) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Hash = dbuf.String()
	// x.Bytes(m.CurrentNumber.Encode())
	return dbuf.err
}

// auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;
func (m *TLAuthSignUp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_signUp))
	x.String(m.PhoneNumber)
	x.String(m.PhoneCodeHash)
	x.String(m.PhoneCode)
	x.String(m.FirstName)
	x.String(m.LastName)
	return x.buf
}

func (m *TLAuthSignUp) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.PhoneCodeHash = dbuf.String()
	m.PhoneCode = dbuf.String()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	return dbuf.err
}

// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
func (m *TLAuthSignIn) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_signIn))
	x.String(m.PhoneNumber)
	x.String(m.PhoneCodeHash)
	x.String(m.PhoneCode)
	return x.buf
}

func (m *TLAuthSignIn) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.PhoneCodeHash = dbuf.String()
	m.PhoneCode = dbuf.String()
	return dbuf.err
}

// auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
func (m *TLAuthImportAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_importAuthorization))
	x.Int(m.Id)
	x.StringBytes(m.Bytes)
	return x.buf
}

func (m *TLAuthImportAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.Int()
	m.Bytes = dbuf.StringBytes()
	return dbuf.err
}

// auth.importBotAuthorization#67a3ff2c flags:int api_id:int api_hash:string bot_auth_token:string = auth.Authorization;
func (m *TLAuthImportBotAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_importBotAuthorization))
	x.Int(m.Flags)
	x.Int(m.ApiId)
	x.String(m.ApiHash)
	x.String(m.BotAuthToken)
	return x.buf
}

func (m *TLAuthImportBotAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.ApiId = dbuf.Int()
	m.ApiHash = dbuf.String()
	m.BotAuthToken = dbuf.String()
	return dbuf.err
}

// auth.checkPassword#a63011e password_hash:bytes = auth.Authorization;
func (m *TLAuthCheckPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_checkPassword))
	x.StringBytes(m.PasswordHash)
	return x.buf
}

func (m *TLAuthCheckPassword) Decode(dbuf *DecodeBuf) error {
	m.PasswordHash = dbuf.StringBytes()
	return dbuf.err
}

// auth.recoverPassword#4ea56e92 code:string = auth.Authorization;
func (m *TLAuthRecoverPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_recoverPassword))
	x.String(m.Code)
	return x.buf
}

func (m *TLAuthRecoverPassword) Decode(dbuf *DecodeBuf) error {
	m.Code = dbuf.String()
	return dbuf.err
}

// auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
func (m *TLAuthExportAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_exportAuthorization))
	x.Int(m.DcId)
	return x.buf
}

func (m *TLAuthExportAuthorization) Decode(dbuf *DecodeBuf) error {
	m.DcId = dbuf.Int()
	return dbuf.err
}

// auth.requestPasswordRecovery#d897bc66 = auth.PasswordRecovery;
func (m *TLAuthRequestPasswordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_requestPasswordRecovery))
	return x.buf
}

func (m *TLAuthRequestPasswordRecovery) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (m *TLAccountGetNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getNotifySettings))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLAccountGetNotifySettings) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (m *TLAccountUpdateProfile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_updateProfile))
	x.Int(m.Flags)
	x.String(m.FirstName)
	x.String(m.LastName)
	x.String(m.About)
	return x.buf
}

func (m *TLAccountUpdateProfile) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.FirstName = dbuf.String()
	m.LastName = dbuf.String()
	m.About = dbuf.String()
	return dbuf.err
}

// account.updateUsername#3e0bdd7c username:string = User;
func (m *TLAccountUpdateUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_updateUsername))
	x.String(m.Username)
	return x.buf
}

func (m *TLAccountUpdateUsername) Decode(dbuf *DecodeBuf) error {
	m.Username = dbuf.String()
	return dbuf.err
}

// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (m *TLAccountChangePhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_changePhone))
	x.String(m.PhoneNumber)
	x.String(m.PhoneCodeHash)
	x.String(m.PhoneCode)
	return x.buf
}

func (m *TLAccountChangePhone) Decode(dbuf *DecodeBuf) error {
	m.PhoneNumber = dbuf.String()
	m.PhoneCodeHash = dbuf.String()
	m.PhoneCode = dbuf.String()
	return dbuf.err
}

// contacts.importCard#4fe196fe export_card:Vector<int> = User;
func (m *TLContactsImportCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_importCard))
	x.VectorInt(m.ExportCard)
	return x.buf
}

func (m *TLContactsImportCard) Decode(dbuf *DecodeBuf) error {
	m.ExportCard = dbuf.VectorInt()
	return dbuf.err
}

// account.getWallPapers#c04cfac2 = Vector<WallPaper>;
func (m *TLAccountGetWallPapers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getWallPapers))
	return x.buf
}

func (m *TLAccountGetWallPapers) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (m *TLAccountGetPrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getPrivacy))
	x.Bytes(m.Key.Encode())
	return x.buf
}

func (m *TLAccountGetPrivacy) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Key.Encode())
	return dbuf.err
}

// account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
func (m *TLAccountSetPrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_setPrivacy))
	x.Bytes(m.Key.Encode())
	// x.VectorMessage(m.Rules)
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Rules)))
	for _, v := range m.Rules {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLAccountSetPrivacy) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Key.Encode())
	// x.VectorMessage(m.Rules)
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Rules = make([]*InputPrivacyRule, l2)
	for i := 0; i < int(l2); i++ {
		m.Rules[i] = &InputPrivacyRule{}
		(*m.Rules[i]).Decode(dbuf)
	}
	return dbuf.err
}

// account.getAccountTTL#8fc711d = AccountDaysTTL;
func (m *TLAccountGetAccountTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getAccountTTL))
	return x.buf
}

func (m *TLAccountGetAccountTTL) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// account.getAuthorizations#e320c158 = account.Authorizations;
func (m *TLAccountGetAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getAuthorizations))
	return x.buf
}

func (m *TLAccountGetAuthorizations) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// account.getPassword#548a30f5 = account.Password;
func (m *TLAccountGetPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getPassword))
	return x.buf
}

func (m *TLAccountGetPassword) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// account.getPasswordSettings#bc8d11bb current_password_hash:bytes = account.PasswordSettings;
func (m *TLAccountGetPasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getPasswordSettings))
	x.StringBytes(m.CurrentPasswordHash)
	return x.buf
}

func (m *TLAccountGetPasswordSettings) Decode(dbuf *DecodeBuf) error {
	m.CurrentPasswordHash = dbuf.StringBytes()
	return dbuf.err
}

// account.getTmpPassword#4a82327e password_hash:bytes period:int = account.TmpPassword;
func (m *TLAccountGetTmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_getTmpPassword))
	x.StringBytes(m.PasswordHash)
	x.Int(m.Period)
	return x.buf
}

func (m *TLAccountGetTmpPassword) Decode(dbuf *DecodeBuf) error {
	m.PasswordHash = dbuf.StringBytes()
	m.Period = dbuf.Int()
	return dbuf.err
}

// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (m *TLUsersGetUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_users_getUsers))
	// x.VectorMessage(m.Id)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Id)))
	for _, v := range m.Id {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLUsersGetUsers) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Id)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Id = make([]*InputUser, l1)
	for i := 0; i < int(l1); i++ {
		m.Id[i] = &InputUser{}
		(*m.Id[i]).Decode(dbuf)
	}
	return dbuf.err
}

// users.getFullUser#ca30a5b1 id:InputUser = UserFull;
func (m *TLUsersGetFullUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_users_getFullUser))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLUsersGetFullUser) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	return dbuf.err
}

// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (m *TLContactsGetStatuses) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_getStatuses))
	return x.buf
}

func (m *TLContactsGetStatuses) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// contacts.getContacts#c023849f hash:int = contacts.Contacts;
func (m *TLContactsGetContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_getContacts))
	x.Int(m.Hash)
	return x.buf
}

func (m *TLContactsGetContacts) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	return dbuf.err
}

// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (m *TLContactsImportContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_importContacts))
	// x.VectorMessage(m.Contacts)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Contacts)))
	for _, v := range m.Contacts {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLContactsImportContacts) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Contacts)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Contacts = make([]*InputContact, l1)
	for i := 0; i < int(l1); i++ {
		m.Contacts[i] = &InputContact{}
		(*m.Contacts[i]).Decode(dbuf)
	}
	return dbuf.err
}

// contacts.deleteContact#8e953744 id:InputUser = contacts.Link;
func (m *TLContactsDeleteContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_deleteContact))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLContactsDeleteContact) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	return dbuf.err
}

// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
func (m *TLContactsGetBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_getBlocked))
	x.Int(m.Offset)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLContactsGetBlocked) Decode(dbuf *DecodeBuf) error {
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// contacts.exportCard#84e53737 = Vector<int>;
func (m *TLContactsExportCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_exportCard))
	return x.buf
}

func (m *TLContactsExportCard) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.getMessagesViews#c4c8a55d peer:InputPeer id:Vector<int> increment:Bool = Vector<int>;
func (m *TLMessagesGetMessagesViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getMessagesViews))
	x.Bytes(m.Peer.Encode())
	x.VectorInt(m.Id)
	x.Bytes(m.Increment.Encode())
	return x.buf
}

func (m *TLMessagesGetMessagesViews) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.Id = dbuf.VectorInt()
	// x.Bytes(m.Increment.Encode())
	return dbuf.err
}

// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (m *TLContactsSearch) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_search))
	x.String(m.Q)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLContactsSearch) Decode(dbuf *DecodeBuf) error {
	m.Q = dbuf.String()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// contacts.resolveUsername#f93ccba3 username:string = contacts.ResolvedPeer;
func (m *TLContactsResolveUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_resolveUsername))
	x.String(m.Username)
	return x.buf
}

func (m *TLContactsResolveUsername) Decode(dbuf *DecodeBuf) error {
	m.Username = dbuf.String()
	return dbuf.err
}

// contacts.getTopPeers#d4982db5 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:int = contacts.TopPeers;
func (m *TLContactsGetTopPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_getTopPeers))
	x.Int(m.Flags)
	x.Int(m.Offset)
	x.Int(m.Limit)
	x.Int(m.Hash)
	return x.buf
}

func (m *TLContactsGetTopPeers) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getMessages#4222fa74 id:Vector<int> = messages.Messages;
func (m *TLMessagesGetMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getMessages))
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLMessagesGetMessages) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// messages.getHistory#afa92846 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (m *TLMessagesGetHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getHistory))
	x.Bytes(m.Peer.Encode())
	x.Int(m.OffsetId)
	x.Int(m.OffsetDate)
	x.Int(m.AddOffset)
	x.Int(m.Limit)
	x.Int(m.MaxId)
	x.Int(m.MinId)
	return x.buf
}

func (m *TLMessagesGetHistory) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.OffsetId = dbuf.Int()
	m.OffsetDate = dbuf.Int()
	m.AddOffset = dbuf.Int()
	m.Limit = dbuf.Int()
	m.MaxId = dbuf.Int()
	m.MinId = dbuf.Int()
	return dbuf.err
}

// messages.search#39e9ea0 flags:# peer:InputPeer q:string from_id:flags.0?InputUser filter:MessagesFilter min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (m *TLMessagesSearch) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_search))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.String(m.Q)
	x.Bytes(m.FromId.Encode())
	x.Bytes(m.Filter.Encode())
	x.Int(m.MinDate)
	x.Int(m.MaxDate)
	x.Int(m.OffsetId)
	x.Int(m.AddOffset)
	x.Int(m.Limit)
	x.Int(m.MaxId)
	x.Int(m.MinId)
	return x.buf
}

func (m *TLMessagesSearch) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.Q = dbuf.String()
	// x.Bytes(m.FromId.Encode())
	// x.Bytes(m.Filter.Encode())
	m.MinDate = dbuf.Int()
	m.MaxDate = dbuf.Int()
	m.OffsetId = dbuf.Int()
	m.AddOffset = dbuf.Int()
	m.Limit = dbuf.Int()
	m.MaxId = dbuf.Int()
	m.MinId = dbuf.Int()
	return dbuf.err
}

// messages.searchGlobal#9e3cacb0 q:string offset_date:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (m *TLMessagesSearchGlobal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_searchGlobal))
	x.String(m.Q)
	x.Int(m.OffsetDate)
	x.Bytes(m.OffsetPeer.Encode())
	x.Int(m.OffsetId)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLMessagesSearchGlobal) Decode(dbuf *DecodeBuf) error {
	m.Q = dbuf.String()
	m.OffsetDate = dbuf.Int()
	// x.Bytes(m.OffsetPeer.Encode())
	m.OffsetId = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// messages.getUnreadMentions#46578472 peer:InputPeer offset_id:int add_offset:int limit:int max_id:int min_id:int = messages.Messages;
func (m *TLMessagesGetUnreadMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getUnreadMentions))
	x.Bytes(m.Peer.Encode())
	x.Int(m.OffsetId)
	x.Int(m.AddOffset)
	x.Int(m.Limit)
	x.Int(m.MaxId)
	x.Int(m.MinId)
	return x.buf
}

func (m *TLMessagesGetUnreadMentions) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.OffsetId = dbuf.Int()
	m.AddOffset = dbuf.Int()
	m.Limit = dbuf.Int()
	m.MaxId = dbuf.Int()
	m.MinId = dbuf.Int()
	return dbuf.err
}

// channels.getMessages#93d7b347 channel:InputChannel id:Vector<int> = messages.Messages;
func (m *TLChannelsGetMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getMessages))
	x.Bytes(m.Channel.Encode())
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLChannelsGetMessages) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// messages.getDialogs#191ba9c5 flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int = messages.Dialogs;
func (m *TLMessagesGetDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getDialogs))
	x.Int(m.Flags)
	x.Int(m.OffsetDate)
	x.Int(m.OffsetId)
	x.Bytes(m.OffsetPeer.Encode())
	x.Int(m.Limit)
	return x.buf
}

func (m *TLMessagesGetDialogs) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.OffsetDate = dbuf.Int()
	m.OffsetId = dbuf.Int()
	// x.Bytes(m.OffsetPeer.Encode())
	m.Limit = dbuf.Int()
	return dbuf.err
}

// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (m *TLMessagesReadHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_readHistory))
	x.Bytes(m.Peer.Encode())
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLMessagesReadHistory) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (m *TLMessagesDeleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_deleteMessages))
	x.Int(m.Flags)
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLMessagesDeleteMessages) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (m *TLMessagesReadMessageContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_readMessageContents))
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLMessagesReadMessageContents) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// channels.deleteMessages#84c1fd4e channel:InputChannel id:Vector<int> = messages.AffectedMessages;
func (m *TLChannelsDeleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_deleteMessages))
	x.Bytes(m.Channel.Encode())
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLChannelsDeleteMessages) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// messages.deleteHistory#1c015b09 flags:# just_clear:flags.0?true peer:InputPeer max_id:int = messages.AffectedHistory;
func (m *TLMessagesDeleteHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_deleteHistory))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLMessagesDeleteHistory) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// channels.deleteUserHistory#d10dd71b channel:InputChannel user_id:InputUser = messages.AffectedHistory;
func (m *TLChannelsDeleteUserHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_deleteUserHistory))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.UserId.Encode())
	return x.buf
}

func (m *TLChannelsDeleteUserHistory) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.UserId.Encode())
	return dbuf.err
}

// messages.receivedMessages#5a954c0 max_id:int = Vector<ReceivedNotifyMessage>;
func (m *TLMessagesReceivedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_receivedMessages))
	x.Int(m.MaxId)
	return x.buf
}

func (m *TLMessagesReceivedMessages) Decode(dbuf *DecodeBuf) error {
	m.MaxId = dbuf.Int()
	return dbuf.err
}

// messages.sendMessage#fa88427a flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;
func (m *TLMessagesSendMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendMessage))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.ReplyToMsgId)
	x.String(m.Message)
	x.Long(m.RandomId)
	x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities)
	x11 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x11, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x11[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesSendMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.ReplyToMsgId = dbuf.Int()
	m.Message = dbuf.String()
	m.RandomId = dbuf.Long()
	// x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities)
	c11 := dbuf.Int()
	if c11 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c11)
	}
	l11 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l11)
	for i := 0; i < int(l11); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.sendMedia#c8f16791 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia random_id:long reply_markup:flags.2?ReplyMarkup = Updates;
func (m *TLMessagesSendMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendMedia))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.ReplyToMsgId)
	x.Bytes(m.Media.Encode())
	x.Long(m.RandomId)
	x.Bytes(m.ReplyMarkup.Encode())
	return x.buf
}

func (m *TLMessagesSendMedia) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.ReplyToMsgId = dbuf.Int()
	// x.Bytes(m.Media.Encode())
	m.RandomId = dbuf.Long()
	// x.Bytes(m.ReplyMarkup.Encode())
	return dbuf.err
}

// messages.forwardMessages#708e0195 flags:# silent:flags.5?true background:flags.6?true with_my_score:flags.8?true from_peer:InputPeer id:Vector<int> random_id:Vector<long> to_peer:InputPeer = Updates;
func (m *TLMessagesForwardMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_forwardMessages))
	x.Int(m.Flags)
	x.Bytes(m.FromPeer.Encode())
	x.VectorInt(m.Id)
	x.VectorLong(m.RandomId)
	x.Bytes(m.ToPeer.Encode())
	return x.buf
}

func (m *TLMessagesForwardMessages) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.FromPeer.Encode())
	m.Id = dbuf.VectorInt()
	m.RandomId = dbuf.VectorLong()
	// x.Bytes(m.ToPeer.Encode())
	return dbuf.err
}

// messages.editChatTitle#dc452855 chat_id:int title:string = Updates;
func (m *TLMessagesEditChatTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_editChatTitle))
	x.Int(m.ChatId)
	x.String(m.Title)
	return x.buf
}

func (m *TLMessagesEditChatTitle) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	m.Title = dbuf.String()
	return dbuf.err
}

// messages.editChatPhoto#ca4c79d8 chat_id:int photo:InputChatPhoto = Updates;
func (m *TLMessagesEditChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_editChatPhoto))
	x.Int(m.ChatId)
	x.Bytes(m.Photo.Encode())
	return x.buf
}

func (m *TLMessagesEditChatPhoto) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	// x.Bytes(m.Photo.Encode())
	return dbuf.err
}

// messages.addChatUser#f9a0aa09 chat_id:int user_id:InputUser fwd_limit:int = Updates;
func (m *TLMessagesAddChatUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_addChatUser))
	x.Int(m.ChatId)
	x.Bytes(m.UserId.Encode())
	x.Int(m.FwdLimit)
	return x.buf
}

func (m *TLMessagesAddChatUser) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	// x.Bytes(m.UserId.Encode())
	m.FwdLimit = dbuf.Int()
	return dbuf.err
}

// messages.deleteChatUser#e0611f16 chat_id:int user_id:InputUser = Updates;
func (m *TLMessagesDeleteChatUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_deleteChatUser))
	x.Int(m.ChatId)
	x.Bytes(m.UserId.Encode())
	return x.buf
}

func (m *TLMessagesDeleteChatUser) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	// x.Bytes(m.UserId.Encode())
	return dbuf.err
}

// messages.createChat#9cb126e users:Vector<InputUser> title:string = Updates;
func (m *TLMessagesCreateChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_createChat))
	// x.VectorMessage(m.Users)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.String(m.Title)
	return x.buf
}

func (m *TLMessagesCreateChat) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Users)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Users = make([]*InputUser, l1)
	for i := 0; i < int(l1); i++ {
		m.Users[i] = &InputUser{}
		(*m.Users[i]).Decode(dbuf)
	}
	m.Title = dbuf.String()
	return dbuf.err
}

// messages.forwardMessage#33963bf9 peer:InputPeer id:int random_id:long = Updates;
func (m *TLMessagesForwardMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_forwardMessage))
	x.Bytes(m.Peer.Encode())
	x.Int(m.Id)
	x.Long(m.RandomId)
	return x.buf
}

func (m *TLMessagesForwardMessage) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.Id = dbuf.Int()
	m.RandomId = dbuf.Long()
	return dbuf.err
}

// messages.importChatInvite#6c50051c hash:string = Updates;
func (m *TLMessagesImportChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_importChatInvite))
	x.String(m.Hash)
	return x.buf
}

func (m *TLMessagesImportChatInvite) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.String()
	return dbuf.err
}

// messages.startBot#e6df7378 bot:InputUser peer:InputPeer random_id:long start_param:string = Updates;
func (m *TLMessagesStartBot) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_startBot))
	x.Bytes(m.Bot.Encode())
	x.Bytes(m.Peer.Encode())
	x.Long(m.RandomId)
	x.String(m.StartParam)
	return x.buf
}

func (m *TLMessagesStartBot) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Bot.Encode())
	// x.Bytes(m.Peer.Encode())
	m.RandomId = dbuf.Long()
	m.StartParam = dbuf.String()
	return dbuf.err
}

// messages.toggleChatAdmins#ec8bd9e1 chat_id:int enabled:Bool = Updates;
func (m *TLMessagesToggleChatAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_toggleChatAdmins))
	x.Int(m.ChatId)
	x.Bytes(m.Enabled.Encode())
	return x.buf
}

func (m *TLMessagesToggleChatAdmins) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	// x.Bytes(m.Enabled.Encode())
	return dbuf.err
}

// messages.migrateChat#15a3b8e3 chat_id:int = Updates;
func (m *TLMessagesMigrateChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_migrateChat))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLMessagesMigrateChat) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// messages.sendInlineBotResult#b16e06fe flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int random_id:long query_id:long id:string = Updates;
func (m *TLMessagesSendInlineBotResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendInlineBotResult))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.ReplyToMsgId)
	x.Long(m.RandomId)
	x.Long(m.QueryId)
	x.String(m.Id)
	return x.buf
}

func (m *TLMessagesSendInlineBotResult) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.ReplyToMsgId = dbuf.Int()
	m.RandomId = dbuf.Long()
	m.QueryId = dbuf.Long()
	m.Id = dbuf.String()
	return dbuf.err
}

// messages.editMessage#ce91e4ca flags:# no_webpage:flags.1?true peer:InputPeer id:int message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;
func (m *TLMessagesEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_editMessage))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.Id)
	x.String(m.Message)
	x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities)
	x7 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x7, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x7[4:], uint32(len(m.Entities)))
	for _, v := range m.Entities {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesEditMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.Id = dbuf.Int()
	m.Message = dbuf.String()
	// x.Bytes(m.ReplyMarkup.Encode())
	// x.VectorMessage(m.Entities)
	c7 := dbuf.Int()
	if c7 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c7)
	}
	l7 := dbuf.Int()
	m.Entities = make([]*MessageEntity, l7)
	for i := 0; i < int(l7); i++ {
		m.Entities[i] = &MessageEntity{}
		(*m.Entities[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.getAllDrafts#6a3f8d65 = Updates;
func (m *TLMessagesGetAllDrafts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getAllDrafts))
	return x.buf
}

func (m *TLMessagesGetAllDrafts) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.setGameScore#8ef8ecc0 flags:# edit_message:flags.0?true force:flags.1?true peer:InputPeer id:int user_id:InputUser score:int = Updates;
func (m *TLMessagesSetGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_setGameScore))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.Id)
	x.Bytes(m.UserId.Encode())
	x.Int(m.Score)
	return x.buf
}

func (m *TLMessagesSetGameScore) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.Id = dbuf.Int()
	// x.Bytes(m.UserId.Encode())
	m.Score = dbuf.Int()
	return dbuf.err
}

// messages.sendScreenshotNotification#c97df020 peer:InputPeer reply_to_msg_id:int random_id:long = Updates;
func (m *TLMessagesSendScreenshotNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendScreenshotNotification))
	x.Bytes(m.Peer.Encode())
	x.Int(m.ReplyToMsgId)
	x.Long(m.RandomId)
	return x.buf
}

func (m *TLMessagesSendScreenshotNotification) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.ReplyToMsgId = dbuf.Int()
	m.RandomId = dbuf.Long()
	return dbuf.err
}

// help.getAppChangelog#9010ef6f prev_app_version:string = Updates;
func (m *TLHelpGetAppChangelog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getAppChangelog))
	x.String(m.PrevAppVersion)
	return x.buf
}

func (m *TLHelpGetAppChangelog) Decode(dbuf *DecodeBuf) error {
	m.PrevAppVersion = dbuf.String()
	return dbuf.err
}

// channels.createChannel#f4893d7f flags:# broadcast:flags.0?true megagroup:flags.1?true title:string about:string = Updates;
func (m *TLChannelsCreateChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_createChannel))
	x.Int(m.Flags)
	x.String(m.Title)
	x.String(m.About)
	return x.buf
}

func (m *TLChannelsCreateChannel) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Title = dbuf.String()
	m.About = dbuf.String()
	return dbuf.err
}

// channels.editAdmin#20b88214 channel:InputChannel user_id:InputUser admin_rights:ChannelAdminRights = Updates;
func (m *TLChannelsEditAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_editAdmin))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.UserId.Encode())
	x.Bytes(m.AdminRights.Encode())
	return x.buf
}

func (m *TLChannelsEditAdmin) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.UserId.Encode())
	// x.Bytes(m.AdminRights.Encode())
	return dbuf.err
}

// channels.editTitle#566decd0 channel:InputChannel title:string = Updates;
func (m *TLChannelsEditTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_editTitle))
	x.Bytes(m.Channel.Encode())
	x.String(m.Title)
	return x.buf
}

func (m *TLChannelsEditTitle) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Title = dbuf.String()
	return dbuf.err
}

// channels.editPhoto#f12e57c9 channel:InputChannel photo:InputChatPhoto = Updates;
func (m *TLChannelsEditPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_editPhoto))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.Photo.Encode())
	return x.buf
}

func (m *TLChannelsEditPhoto) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.Photo.Encode())
	return dbuf.err
}

// channels.joinChannel#24b524c5 channel:InputChannel = Updates;
func (m *TLChannelsJoinChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_joinChannel))
	x.Bytes(m.Channel.Encode())
	return x.buf
}

func (m *TLChannelsJoinChannel) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	return dbuf.err
}

// channels.leaveChannel#f836aa95 channel:InputChannel = Updates;
func (m *TLChannelsLeaveChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_leaveChannel))
	x.Bytes(m.Channel.Encode())
	return x.buf
}

func (m *TLChannelsLeaveChannel) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	return dbuf.err
}

// channels.inviteToChannel#199f3a6c channel:InputChannel users:Vector<InputUser> = Updates;
func (m *TLChannelsInviteToChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_inviteToChannel))
	x.Bytes(m.Channel.Encode())
	// x.VectorMessage(m.Users)
	x2 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Users)))
	for _, v := range m.Users {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChannelsInviteToChannel) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.VectorMessage(m.Users)
	c2 := dbuf.Int()
	if c2 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c2)
	}
	l2 := dbuf.Int()
	m.Users = make([]*InputUser, l2)
	for i := 0; i < int(l2); i++ {
		m.Users[i] = &InputUser{}
		(*m.Users[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channels.deleteChannel#c0111fe3 channel:InputChannel = Updates;
func (m *TLChannelsDeleteChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_deleteChannel))
	x.Bytes(m.Channel.Encode())
	return x.buf
}

func (m *TLChannelsDeleteChannel) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	return dbuf.err
}

// channels.toggleInvites#49609307 channel:InputChannel enabled:Bool = Updates;
func (m *TLChannelsToggleInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_toggleInvites))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.Enabled.Encode())
	return x.buf
}

func (m *TLChannelsToggleInvites) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.Enabled.Encode())
	return dbuf.err
}

// channels.toggleSignatures#1f69b606 channel:InputChannel enabled:Bool = Updates;
func (m *TLChannelsToggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_toggleSignatures))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.Enabled.Encode())
	return x.buf
}

func (m *TLChannelsToggleSignatures) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.Enabled.Encode())
	return dbuf.err
}

// channels.updatePinnedMessage#a72ded52 flags:# silent:flags.0?true channel:InputChannel id:int = Updates;
func (m *TLChannelsUpdatePinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_updatePinnedMessage))
	x.Int(m.Flags)
	x.Bytes(m.Channel.Encode())
	x.Int(m.Id)
	return x.buf
}

func (m *TLChannelsUpdatePinnedMessage) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Channel.Encode())
	m.Id = dbuf.Int()
	return dbuf.err
}

// channels.editBanned#bfd915cd channel:InputChannel user_id:InputUser banned_rights:ChannelBannedRights = Updates;
func (m *TLChannelsEditBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_editBanned))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.UserId.Encode())
	x.Bytes(m.BannedRights.Encode())
	return x.buf
}

func (m *TLChannelsEditBanned) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.UserId.Encode())
	// x.Bytes(m.BannedRights.Encode())
	return dbuf.err
}

// phone.discardCall#78d413a6 peer:InputPhoneCall duration:int reason:PhoneCallDiscardReason connection_id:long = Updates;
func (m *TLPhoneDiscardCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_discardCall))
	x.Bytes(m.Peer.Encode())
	x.Int(m.Duration)
	x.Bytes(m.Reason.Encode())
	x.Long(m.ConnectionId)
	return x.buf
}

func (m *TLPhoneDiscardCall) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.Duration = dbuf.Int()
	// x.Bytes(m.Reason.Encode())
	m.ConnectionId = dbuf.Long()
	return dbuf.err
}

// phone.setCallRating#1c536a34 peer:InputPhoneCall rating:int comment:string = Updates;
func (m *TLPhoneSetCallRating) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_setCallRating))
	x.Bytes(m.Peer.Encode())
	x.Int(m.Rating)
	x.String(m.Comment)
	return x.buf
}

func (m *TLPhoneSetCallRating) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.Rating = dbuf.Int()
	m.Comment = dbuf.String()
	return dbuf.err
}

// messages.getPeerSettings#3672e09c peer:InputPeer = PeerSettings;
func (m *TLMessagesGetPeerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getPeerSettings))
	x.Bytes(m.Peer.Encode())
	return x.buf
}

func (m *TLMessagesGetPeerSettings) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	return dbuf.err
}

// messages.getChats#3c6aa187 id:Vector<int> = messages.Chats;
func (m *TLMessagesGetChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getChats))
	x.VectorInt(m.Id)
	return x.buf
}

func (m *TLMessagesGetChats) Decode(dbuf *DecodeBuf) error {
	m.Id = dbuf.VectorInt()
	return dbuf.err
}

// messages.getCommonChats#d0a48c4 user_id:InputUser max_id:int limit:int = messages.Chats;
func (m *TLMessagesGetCommonChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getCommonChats))
	x.Bytes(m.UserId.Encode())
	x.Int(m.MaxId)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLMessagesGetCommonChats) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.UserId.Encode())
	m.MaxId = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// messages.getAllChats#eba80ff0 except_ids:Vector<int> = messages.Chats;
func (m *TLMessagesGetAllChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getAllChats))
	x.VectorInt(m.ExceptIds)
	return x.buf
}

func (m *TLMessagesGetAllChats) Decode(dbuf *DecodeBuf) error {
	m.ExceptIds = dbuf.VectorInt()
	return dbuf.err
}

// channels.getChannels#a7f6bbb id:Vector<InputChannel> = messages.Chats;
func (m *TLChannelsGetChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getChannels))
	// x.VectorMessage(m.Id)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Id)))
	for _, v := range m.Id {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLChannelsGetChannels) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Id)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Id = make([]*InputChannel, l1)
	for i := 0; i < int(l1); i++ {
		m.Id[i] = &InputChannel{}
		(*m.Id[i]).Decode(dbuf)
	}
	return dbuf.err
}

// channels.getAdminedPublicChannels#8d8d82d7 = messages.Chats;
func (m *TLChannelsGetAdminedPublicChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getAdminedPublicChannels))
	return x.buf
}

func (m *TLChannelsGetAdminedPublicChannels) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.getFullChat#3b831c66 chat_id:int = messages.ChatFull;
func (m *TLMessagesGetFullChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getFullChat))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLMessagesGetFullChat) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// channels.getFullChannel#8736a09 channel:InputChannel = messages.ChatFull;
func (m *TLChannelsGetFullChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getFullChannel))
	x.Bytes(m.Channel.Encode())
	return x.buf
}

func (m *TLChannelsGetFullChannel) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	return dbuf.err
}

// messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
func (m *TLMessagesGetDhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getDhConfig))
	x.Int(m.Version)
	x.Int(m.RandomLength)
	return x.buf
}

func (m *TLMessagesGetDhConfig) Decode(dbuf *DecodeBuf) error {
	m.Version = dbuf.Int()
	m.RandomLength = dbuf.Int()
	return dbuf.err
}

// messages.requestEncryption#f64daf43 user_id:InputUser random_id:int g_a:bytes = EncryptedChat;
func (m *TLMessagesRequestEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_requestEncryption))
	x.Bytes(m.UserId.Encode())
	x.Int(m.RandomId)
	x.StringBytes(m.GA)
	return x.buf
}

func (m *TLMessagesRequestEncryption) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.UserId.Encode())
	m.RandomId = dbuf.Int()
	m.GA = dbuf.StringBytes()
	return dbuf.err
}

// messages.acceptEncryption#3dbc0415 peer:InputEncryptedChat g_b:bytes key_fingerprint:long = EncryptedChat;
func (m *TLMessagesAcceptEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_acceptEncryption))
	x.Bytes(m.Peer.Encode())
	x.StringBytes(m.GB)
	x.Long(m.KeyFingerprint)
	return x.buf
}

func (m *TLMessagesAcceptEncryption) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.GB = dbuf.StringBytes()
	m.KeyFingerprint = dbuf.Long()
	return dbuf.err
}

// messages.sendEncrypted#a9776773 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
func (m *TLMessagesSendEncrypted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendEncrypted))
	x.Bytes(m.Peer.Encode())
	x.Long(m.RandomId)
	x.StringBytes(m.Data)
	return x.buf
}

func (m *TLMessagesSendEncrypted) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.RandomId = dbuf.Long()
	m.Data = dbuf.StringBytes()
	return dbuf.err
}

// messages.sendEncryptedFile#9a901b66 peer:InputEncryptedChat random_id:long data:bytes file:InputEncryptedFile = messages.SentEncryptedMessage;
func (m *TLMessagesSendEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendEncryptedFile))
	x.Bytes(m.Peer.Encode())
	x.Long(m.RandomId)
	x.StringBytes(m.Data)
	x.Bytes(m.File.Encode())
	return x.buf
}

func (m *TLMessagesSendEncryptedFile) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.RandomId = dbuf.Long()
	m.Data = dbuf.StringBytes()
	// x.Bytes(m.File.Encode())
	return dbuf.err
}

// messages.sendEncryptedService#32d439a4 peer:InputEncryptedChat random_id:long data:bytes = messages.SentEncryptedMessage;
func (m *TLMessagesSendEncryptedService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sendEncryptedService))
	x.Bytes(m.Peer.Encode())
	x.Long(m.RandomId)
	x.StringBytes(m.Data)
	return x.buf
}

func (m *TLMessagesSendEncryptedService) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.RandomId = dbuf.Long()
	m.Data = dbuf.StringBytes()
	return dbuf.err
}

// messages.receivedQueue#55a5bb66 max_qts:int = Vector<long>;
func (m *TLMessagesReceivedQueue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_receivedQueue))
	x.Int(m.MaxQts)
	return x.buf
}

func (m *TLMessagesReceivedQueue) Decode(dbuf *DecodeBuf) error {
	m.MaxQts = dbuf.Int()
	return dbuf.err
}

// photos.deletePhotos#87cf7f2f id:Vector<InputPhoto> = Vector<long>;
func (m *TLPhotosDeletePhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_deletePhotos))
	// x.VectorMessage(m.Id)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Id)))
	for _, v := range m.Id {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLPhotosDeletePhotos) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Id)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Id = make([]*InputPhoto, l1)
	for i := 0; i < int(l1); i++ {
		m.Id[i] = &InputPhoto{}
		(*m.Id[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.getAllStickers#1c9618b1 hash:int = messages.AllStickers;
func (m *TLMessagesGetAllStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getAllStickers))
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetAllStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getMaskStickers#65b8c79f hash:int = messages.AllStickers;
func (m *TLMessagesGetMaskStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getMaskStickers))
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetMaskStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getWebPagePreview#25223e24 message:string = MessageMedia;
func (m *TLMessagesGetWebPagePreview) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getWebPagePreview))
	x.String(m.Message)
	return x.buf
}

func (m *TLMessagesGetWebPagePreview) Decode(dbuf *DecodeBuf) error {
	m.Message = dbuf.String()
	return dbuf.err
}

// messages.uploadMedia#519bc2b1 peer:InputPeer media:InputMedia = MessageMedia;
func (m *TLMessagesUploadMedia) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_uploadMedia))
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.Media.Encode())
	return x.buf
}

func (m *TLMessagesUploadMedia) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.Media.Encode())
	return dbuf.err
}

// messages.exportChatInvite#7d885289 chat_id:int = ExportedChatInvite;
func (m *TLMessagesExportChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_exportChatInvite))
	x.Int(m.ChatId)
	return x.buf
}

func (m *TLMessagesExportChatInvite) Decode(dbuf *DecodeBuf) error {
	m.ChatId = dbuf.Int()
	return dbuf.err
}

// channels.exportInvite#c7560885 channel:InputChannel = ExportedChatInvite;
func (m *TLChannelsExportInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_exportInvite))
	x.Bytes(m.Channel.Encode())
	return x.buf
}

func (m *TLChannelsExportInvite) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	return dbuf.err
}

// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (m *TLMessagesCheckChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_checkChatInvite))
	x.String(m.Hash)
	return x.buf
}

func (m *TLMessagesCheckChatInvite) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.String()
	return dbuf.err
}

// messages.getStickerSet#2619a90e stickerset:InputStickerSet = messages.StickerSet;
func (m *TLMessagesGetStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getStickerSet))
	x.Bytes(m.Stickerset.Encode())
	return x.buf
}

func (m *TLMessagesGetStickerSet) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Stickerset.Encode())
	return dbuf.err
}

// stickers.createStickerSet#9bd86e6a flags:# masks:flags.0?true user_id:InputUser title:string short_name:string stickers:Vector<InputStickerSetItem> = messages.StickerSet;
func (m *TLStickersCreateStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickers_createStickerSet))
	x.Int(m.Flags)
	x.Bytes(m.UserId.Encode())
	x.String(m.Title)
	x.String(m.ShortName)
	// x.VectorMessage(m.Stickers)
	x6 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x6, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x6[4:], uint32(len(m.Stickers)))
	for _, v := range m.Stickers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLStickersCreateStickerSet) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.UserId.Encode())
	m.Title = dbuf.String()
	m.ShortName = dbuf.String()
	// x.VectorMessage(m.Stickers)
	c6 := dbuf.Int()
	if c6 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c6)
	}
	l6 := dbuf.Int()
	m.Stickers = make([]*InputStickerSetItem, l6)
	for i := 0; i < int(l6); i++ {
		m.Stickers[i] = &InputStickerSetItem{}
		(*m.Stickers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// stickers.removeStickerFromSet#f7760f51 sticker:InputDocument = messages.StickerSet;
func (m *TLStickersRemoveStickerFromSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickers_removeStickerFromSet))
	x.Bytes(m.Sticker.Encode())
	return x.buf
}

func (m *TLStickersRemoveStickerFromSet) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Sticker.Encode())
	return dbuf.err
}

// stickers.changeStickerPosition#ffb6d4ca sticker:InputDocument position:int = messages.StickerSet;
func (m *TLStickersChangeStickerPosition) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickers_changeStickerPosition))
	x.Bytes(m.Sticker.Encode())
	x.Int(m.Position)
	return x.buf
}

func (m *TLStickersChangeStickerPosition) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Sticker.Encode())
	m.Position = dbuf.Int()
	return dbuf.err
}

// stickers.addStickerToSet#8653febe stickerset:InputStickerSet sticker:InputStickerSetItem = messages.StickerSet;
func (m *TLStickersAddStickerToSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickers_addStickerToSet))
	x.Bytes(m.Stickerset.Encode())
	x.Bytes(m.Sticker.Encode())
	return x.buf
}

func (m *TLStickersAddStickerToSet) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Stickerset.Encode())
	// x.Bytes(m.Sticker.Encode())
	return dbuf.err
}

// messages.installStickerSet#c78fe460 stickerset:InputStickerSet archived:Bool = messages.StickerSetInstallResult;
func (m *TLMessagesInstallStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_installStickerSet))
	x.Bytes(m.Stickerset.Encode())
	x.Bytes(m.Archived.Encode())
	return x.buf
}

func (m *TLMessagesInstallStickerSet) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Stickerset.Encode())
	// x.Bytes(m.Archived.Encode())
	return dbuf.err
}

// messages.getDocumentByHash#338e2464 sha256:bytes size:int mime_type:string = Document;
func (m *TLMessagesGetDocumentByHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getDocumentByHash))
	x.StringBytes(m.Sha256)
	x.Int(m.Size)
	x.String(m.MimeType)
	return x.buf
}

func (m *TLMessagesGetDocumentByHash) Decode(dbuf *DecodeBuf) error {
	m.Sha256 = dbuf.StringBytes()
	m.Size = dbuf.Int()
	m.MimeType = dbuf.String()
	return dbuf.err
}

// messages.searchGifs#bf9a776b q:string offset:int = messages.FoundGifs;
func (m *TLMessagesSearchGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_searchGifs))
	x.String(m.Q)
	x.Int(m.Offset)
	return x.buf
}

func (m *TLMessagesSearchGifs) Decode(dbuf *DecodeBuf) error {
	m.Q = dbuf.String()
	m.Offset = dbuf.Int()
	return dbuf.err
}

// messages.getSavedGifs#83bf3d52 hash:int = messages.SavedGifs;
func (m *TLMessagesGetSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getSavedGifs))
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetSavedGifs) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getInlineBotResults#514e999d flags:# bot:InputUser peer:InputPeer geo_point:flags.0?InputGeoPoint query:string offset:string = messages.BotResults;
func (m *TLMessagesGetInlineBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getInlineBotResults))
	x.Int(m.Flags)
	x.Bytes(m.Bot.Encode())
	x.Bytes(m.Peer.Encode())
	x.Bytes(m.GeoPoint.Encode())
	x.String(m.Query)
	x.String(m.Offset)
	return x.buf
}

func (m *TLMessagesGetInlineBotResults) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Bot.Encode())
	// x.Bytes(m.Peer.Encode())
	// x.Bytes(m.GeoPoint.Encode())
	m.Query = dbuf.String()
	m.Offset = dbuf.String()
	return dbuf.err
}

// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (m *TLMessagesGetMessageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getMessageEditData))
	x.Bytes(m.Peer.Encode())
	x.Int(m.Id)
	return x.buf
}

func (m *TLMessagesGetMessageEditData) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.Id = dbuf.Int()
	return dbuf.err
}

// messages.getBotCallbackAnswer#810a9fec flags:# game:flags.1?true peer:InputPeer msg_id:int data:flags.0?bytes = messages.BotCallbackAnswer;
func (m *TLMessagesGetBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getBotCallbackAnswer))
	x.Int(m.Flags)
	x.Bytes(m.Peer.Encode())
	x.Int(m.MsgId)
	x.StringBytes(m.Data)
	return x.buf
}

func (m *TLMessagesGetBotCallbackAnswer) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Peer.Encode())
	m.MsgId = dbuf.Int()
	m.Data = dbuf.StringBytes()
	return dbuf.err
}

// messages.getPeerDialogs#2d9776b9 peers:Vector<InputPeer> = messages.PeerDialogs;
func (m *TLMessagesGetPeerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getPeerDialogs))
	// x.VectorMessage(m.Peers)
	x1 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Peers)))
	for _, v := range m.Peers {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	return x.buf
}

func (m *TLMessagesGetPeerDialogs) Decode(dbuf *DecodeBuf) error {
	// x.VectorMessage(m.Peers)
	c1 := dbuf.Int()
	if c1 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c1)
	}
	l1 := dbuf.Int()
	m.Peers = make([]*InputPeer, l1)
	for i := 0; i < int(l1); i++ {
		m.Peers[i] = &InputPeer{}
		(*m.Peers[i]).Decode(dbuf)
	}
	return dbuf.err
}

// messages.getPinnedDialogs#e254d64e = messages.PeerDialogs;
func (m *TLMessagesGetPinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getPinnedDialogs))
	return x.buf
}

func (m *TLMessagesGetPinnedDialogs) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// messages.getFeaturedStickers#2dacca4f hash:int = messages.FeaturedStickers;
func (m *TLMessagesGetFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getFeaturedStickers))
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetFeaturedStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getRecentStickers#5ea192c9 flags:# attached:flags.0?true hash:int = messages.RecentStickers;
func (m *TLMessagesGetRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getRecentStickers))
	x.Int(m.Flags)
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetRecentStickers) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getArchivedStickers#57f17692 flags:# masks:flags.0?true offset_id:long limit:int = messages.ArchivedStickers;
func (m *TLMessagesGetArchivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getArchivedStickers))
	x.Int(m.Flags)
	x.Long(m.OffsetId)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLMessagesGetArchivedStickers) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.OffsetId = dbuf.Long()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// messages.getAttachedStickers#cc5b67cc media:InputStickeredMedia = Vector<StickerSetCovered>;
func (m *TLMessagesGetAttachedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getAttachedStickers))
	x.Bytes(m.Media.Encode())
	return x.buf
}

func (m *TLMessagesGetAttachedStickers) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Media.Encode())
	return dbuf.err
}

// messages.getGameHighScores#e822649d peer:InputPeer id:int user_id:InputUser = messages.HighScores;
func (m *TLMessagesGetGameHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getGameHighScores))
	x.Bytes(m.Peer.Encode())
	x.Int(m.Id)
	x.Bytes(m.UserId.Encode())
	return x.buf
}

func (m *TLMessagesGetGameHighScores) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.Id = dbuf.Int()
	// x.Bytes(m.UserId.Encode())
	return dbuf.err
}

// messages.getInlineGameHighScores#f635e1b id:InputBotInlineMessageID user_id:InputUser = messages.HighScores;
func (m *TLMessagesGetInlineGameHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getInlineGameHighScores))
	x.Bytes(m.Id.Encode())
	x.Bytes(m.UserId.Encode())
	return x.buf
}

func (m *TLMessagesGetInlineGameHighScores) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	// x.Bytes(m.UserId.Encode())
	return dbuf.err
}

// messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
func (m *TLMessagesGetWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getWebPage))
	x.String(m.Url)
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetWebPage) Decode(dbuf *DecodeBuf) error {
	m.Url = dbuf.String()
	m.Hash = dbuf.Int()
	return dbuf.err
}

// messages.getFavedStickers#21ce0b0e hash:int = messages.FavedStickers;
func (m *TLMessagesGetFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_getFavedStickers))
	x.Int(m.Hash)
	return x.buf
}

func (m *TLMessagesGetFavedStickers) Decode(dbuf *DecodeBuf) error {
	m.Hash = dbuf.Int()
	return dbuf.err
}

// updates.getState#edd4882a = updates.State;
func (m *TLUpdatesGetState) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_getState))
	return x.buf
}

func (m *TLUpdatesGetState) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (m *TLUpdatesGetDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_getDifference))
	x.Int(m.Flags)
	x.Int(m.Pts)
	x.Int(m.PtsTotalLimit)
	x.Int(m.Date)
	x.Int(m.Qts)
	return x.buf
}

func (m *TLUpdatesGetDifference) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.Pts = dbuf.Int()
	m.PtsTotalLimit = dbuf.Int()
	m.Date = dbuf.Int()
	m.Qts = dbuf.Int()
	return dbuf.err
}

// updates.getChannelDifference#3173d78 flags:# force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (m *TLUpdatesGetChannelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_getChannelDifference))
	x.Int(m.Flags)
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.Filter.Encode())
	x.Int(m.Pts)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLUpdatesGetChannelDifference) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.Filter.Encode())
	m.Pts = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// photos.updateProfilePhoto#f0bb5152 id:InputPhoto = UserProfilePhoto;
func (m *TLPhotosUpdateProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_updateProfilePhoto))
	x.Bytes(m.Id.Encode())
	return x.buf
}

func (m *TLPhotosUpdateProfilePhoto) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Id.Encode())
	return dbuf.err
}

// photos.uploadProfilePhoto#4f32c098 file:InputFile = photos.Photo;
func (m *TLPhotosUploadProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_uploadProfilePhoto))
	x.Bytes(m.File.Encode())
	return x.buf
}

func (m *TLPhotosUploadProfilePhoto) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.File.Encode())
	return dbuf.err
}

// photos.getUserPhotos#91cd32a8 user_id:InputUser offset:int max_id:long limit:int = photos.Photos;
func (m *TLPhotosGetUserPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_getUserPhotos))
	x.Bytes(m.UserId.Encode())
	x.Int(m.Offset)
	x.Long(m.MaxId)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLPhotosGetUserPhotos) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.UserId.Encode())
	m.Offset = dbuf.Int()
	m.MaxId = dbuf.Long()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
func (m *TLUploadGetFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_getFile))
	x.Bytes(m.Location.Encode())
	x.Int(m.Offset)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLUploadGetFile) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Location.Encode())
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (m *TLUploadGetWebFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_getWebFile))
	x.Bytes(m.Location.Encode())
	x.Int(m.Offset)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLUploadGetWebFile) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Location.Encode())
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
func (m *TLUploadGetCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_getCdnFile))
	x.StringBytes(m.FileToken)
	x.Int(m.Offset)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLUploadGetCdnFile) Decode(dbuf *DecodeBuf) error {
	m.FileToken = dbuf.StringBytes()
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// upload.reuploadCdnFile#1af91c09 file_token:bytes request_token:bytes = Vector<CdnFileHash>;
func (m *TLUploadReuploadCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_reuploadCdnFile))
	x.StringBytes(m.FileToken)
	x.StringBytes(m.RequestToken)
	return x.buf
}

func (m *TLUploadReuploadCdnFile) Decode(dbuf *DecodeBuf) error {
	m.FileToken = dbuf.StringBytes()
	m.RequestToken = dbuf.StringBytes()
	return dbuf.err
}

// upload.getCdnFileHashes#f715c87b file_token:bytes offset:int = Vector<CdnFileHash>;
func (m *TLUploadGetCdnFileHashes) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_getCdnFileHashes))
	x.StringBytes(m.FileToken)
	x.Int(m.Offset)
	return x.buf
}

func (m *TLUploadGetCdnFileHashes) Decode(dbuf *DecodeBuf) error {
	m.FileToken = dbuf.StringBytes()
	m.Offset = dbuf.Int()
	return dbuf.err
}

// help.getConfig#c4f9186b = Config;
func (m *TLHelpGetConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getConfig))
	return x.buf
}

func (m *TLHelpGetConfig) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.getNearestDc#1fb33026 = NearestDc;
func (m *TLHelpGetNearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getNearestDc))
	return x.buf
}

func (m *TLHelpGetNearestDc) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.getAppUpdate#ae2de196 = help.AppUpdate;
func (m *TLHelpGetAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getAppUpdate))
	return x.buf
}

func (m *TLHelpGetAppUpdate) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.getInviteText#4d392343 = help.InviteText;
func (m *TLHelpGetInviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getInviteText))
	return x.buf
}

func (m *TLHelpGetInviteText) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.getSupport#9cdf08cd = help.Support;
func (m *TLHelpGetSupport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getSupport))
	return x.buf
}

func (m *TLHelpGetSupport) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.getTermsOfService#350170f3 = help.TermsOfService;
func (m *TLHelpGetTermsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getTermsOfService))
	return x.buf
}

func (m *TLHelpGetTermsOfService) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// help.getCdnConfig#52029342 = CdnConfig;
func (m *TLHelpGetCdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_getCdnConfig))
	return x.buf
}

func (m *TLHelpGetCdnConfig) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// channels.getParticipants#24d98f92 channel:InputChannel filter:ChannelParticipantsFilter offset:int limit:int = channels.ChannelParticipants;
func (m *TLChannelsGetParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getParticipants))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.Filter.Encode())
	x.Int(m.Offset)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLChannelsGetParticipants) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.Filter.Encode())
	m.Offset = dbuf.Int()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// channels.getParticipant#546dd7a6 channel:InputChannel user_id:InputUser = channels.ChannelParticipant;
func (m *TLChannelsGetParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getParticipant))
	x.Bytes(m.Channel.Encode())
	x.Bytes(m.UserId.Encode())
	return x.buf
}

func (m *TLChannelsGetParticipant) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	// x.Bytes(m.UserId.Encode())
	return dbuf.err
}

// channels.exportMessageLink#c846d22d channel:InputChannel id:int = ExportedMessageLink;
func (m *TLChannelsExportMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_exportMessageLink))
	x.Bytes(m.Channel.Encode())
	x.Int(m.Id)
	return x.buf
}

func (m *TLChannelsExportMessageLink) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Channel.Encode())
	m.Id = dbuf.Int()
	return dbuf.err
}

// channels.getAdminLog#33ddf480 flags:# channel:InputChannel q:string events_filter:flags.0?ChannelAdminLogEventsFilter admins:flags.1?Vector<InputUser> max_id:long min_id:long limit:int = channels.AdminLogResults;
func (m *TLChannelsGetAdminLog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_getAdminLog))
	x.Int(m.Flags)
	x.Bytes(m.Channel.Encode())
	x.String(m.Q)
	x.Bytes(m.EventsFilter.Encode())
	// x.VectorMessage(m.Admins)
	x5 := make([]byte, 8)
	binary.LittleEndian.PutUint32(x5, uint32(TLConstructor_CRC32_vector))
	binary.LittleEndian.PutUint32(x5[4:], uint32(len(m.Admins)))
	for _, v := range m.Admins {
		x.buf = append(x.buf, (*v).Encode()...)
	}
	x.Long(m.MaxId)
	x.Long(m.MinId)
	x.Int(m.Limit)
	return x.buf
}

func (m *TLChannelsGetAdminLog) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	// x.Bytes(m.Channel.Encode())
	m.Q = dbuf.String()
	// x.Bytes(m.EventsFilter.Encode())
	// x.VectorMessage(m.Admins)
	c5 := dbuf.Int()
	if c5 != int32(TLConstructor_CRC32_vector) {
		return fmt.Errorf("Not vector, classID: ", c5)
	}
	l5 := dbuf.Int()
	m.Admins = make([]*InputUser, l5)
	for i := 0; i < int(l5); i++ {
		m.Admins[i] = &InputUser{}
		(*m.Admins[i]).Decode(dbuf)
	}
	m.MaxId = dbuf.Long()
	m.MinId = dbuf.Long()
	m.Limit = dbuf.Int()
	return dbuf.err
}

// bots.sendCustomRequest#aa2769ed custom_method:string params:DataJSON = DataJSON;
func (m *TLBotsSendCustomRequest) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_bots_sendCustomRequest))
	x.String(m.CustomMethod)
	x.Bytes(m.Params.Encode())
	return x.buf
}

func (m *TLBotsSendCustomRequest) Decode(dbuf *DecodeBuf) error {
	m.CustomMethod = dbuf.String()
	// x.Bytes(m.Params.Encode())
	return dbuf.err
}

// phone.getCallConfig#55451fa9 = DataJSON;
func (m *TLPhoneGetCallConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_getCallConfig))
	return x.buf
}

func (m *TLPhoneGetCallConfig) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// payments.getPaymentForm#99f09745 msg_id:int = payments.PaymentForm;
func (m *TLPaymentsGetPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_getPaymentForm))
	x.Int(m.MsgId)
	return x.buf
}

func (m *TLPaymentsGetPaymentForm) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Int()
	return dbuf.err
}

// payments.getPaymentReceipt#a092a980 msg_id:int = payments.PaymentReceipt;
func (m *TLPaymentsGetPaymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_getPaymentReceipt))
	x.Int(m.MsgId)
	return x.buf
}

func (m *TLPaymentsGetPaymentReceipt) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Int()
	return dbuf.err
}

// payments.validateRequestedInfo#770a8e74 flags:# save:flags.0?true msg_id:int info:PaymentRequestedInfo = payments.ValidatedRequestedInfo;
func (m *TLPaymentsValidateRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_validateRequestedInfo))
	x.Int(m.Flags)
	x.Int(m.MsgId)
	x.Bytes(m.Info.Encode())
	return x.buf
}

func (m *TLPaymentsValidateRequestedInfo) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.MsgId = dbuf.Int()
	// x.Bytes(m.Info.Encode())
	return dbuf.err
}

// payments.sendPaymentForm#2b8879b3 flags:# msg_id:int requested_info_id:flags.0?string shipping_option_id:flags.1?string credentials:InputPaymentCredentials = payments.PaymentResult;
func (m *TLPaymentsSendPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_sendPaymentForm))
	x.Int(m.Flags)
	x.Int(m.MsgId)
	x.String(m.RequestedInfoId)
	x.String(m.ShippingOptionId)
	x.Bytes(m.Credentials.Encode())
	return x.buf
}

func (m *TLPaymentsSendPaymentForm) Decode(dbuf *DecodeBuf) error {
	m.Flags = dbuf.Int()
	m.MsgId = dbuf.Int()
	m.RequestedInfoId = dbuf.String()
	m.ShippingOptionId = dbuf.String()
	// x.Bytes(m.Credentials.Encode())
	return dbuf.err
}

// payments.getSavedInfo#227d824b = payments.SavedInfo;
func (m *TLPaymentsGetSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_getSavedInfo))
	return x.buf
}

func (m *TLPaymentsGetSavedInfo) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}

// phone.requestCall#5b95b3d4 user_id:InputUser random_id:int g_a_hash:bytes protocol:PhoneCallProtocol = phone.PhoneCall;
func (m *TLPhoneRequestCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_requestCall))
	x.Bytes(m.UserId.Encode())
	x.Int(m.RandomId)
	x.StringBytes(m.GAHash)
	x.Bytes(m.Protocol.Encode())
	return x.buf
}

func (m *TLPhoneRequestCall) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.UserId.Encode())
	m.RandomId = dbuf.Int()
	m.GAHash = dbuf.StringBytes()
	// x.Bytes(m.Protocol.Encode())
	return dbuf.err
}

// phone.acceptCall#3bd2b4a0 peer:InputPhoneCall g_b:bytes protocol:PhoneCallProtocol = phone.PhoneCall;
func (m *TLPhoneAcceptCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_acceptCall))
	x.Bytes(m.Peer.Encode())
	x.StringBytes(m.GB)
	x.Bytes(m.Protocol.Encode())
	return x.buf
}

func (m *TLPhoneAcceptCall) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.GB = dbuf.StringBytes()
	// x.Bytes(m.Protocol.Encode())
	return dbuf.err
}

// phone.confirmCall#2efe1722 peer:InputPhoneCall g_a:bytes key_fingerprint:long protocol:PhoneCallProtocol = phone.PhoneCall;
func (m *TLPhoneConfirmCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_confirmCall))
	x.Bytes(m.Peer.Encode())
	x.StringBytes(m.GA)
	x.Long(m.KeyFingerprint)
	x.Bytes(m.Protocol.Encode())
	return x.buf
}

func (m *TLPhoneConfirmCall) Decode(dbuf *DecodeBuf) error {
	// x.Bytes(m.Peer.Encode())
	m.GA = dbuf.StringBytes()
	m.KeyFingerprint = dbuf.Long()
	// x.Bytes(m.Protocol.Encode())
	return dbuf.err
}

// langpack.getLangPack#9ab5c58e lang_code:string = LangPackDifference;
func (m *TLLangpackGetLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langpack_getLangPack))
	x.String(m.LangCode)
	return x.buf
}

func (m *TLLangpackGetLangPack) Decode(dbuf *DecodeBuf) error {
	m.LangCode = dbuf.String()
	return dbuf.err
}

// langpack.getDifference#b2e4d7d from_version:int = LangPackDifference;
func (m *TLLangpackGetDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langpack_getDifference))
	x.Int(m.FromVersion)
	return x.buf
}

func (m *TLLangpackGetDifference) Decode(dbuf *DecodeBuf) error {
	m.FromVersion = dbuf.Int()
	return dbuf.err
}

// langpack.getStrings#2e1ee318 lang_code:string keys:Vector<string> = Vector<LangPackString>;
func (m *TLLangpackGetStrings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langpack_getStrings))
	x.String(m.LangCode)
	x.VectorString(m.Keys)
	return x.buf
}

func (m *TLLangpackGetStrings) Decode(dbuf *DecodeBuf) error {
	m.LangCode = dbuf.String()
	m.Keys = dbuf.VectorString()
	return dbuf.err
}

// langpack.getLanguages#800fd57d = Vector<LangPackLanguage>;
func (m *TLLangpackGetLanguages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langpack_getLanguages))
	return x.buf
}

func (m *TLLangpackGetLanguages) Decode(dbuf *DecodeBuf) error {
	return dbuf.err
}
