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
)

// help.getConfig#c4f9186b = Config;
func (s *HelpServiceImpl) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("HelpGetConfig - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 设置Reply对象累死人了, 得想个办法实现model和mtproto的自动转换
	helpConfig := mtproto.NewTLConfig()
	// &mtproto.TLConfig{}
	helpConfig.SetPhonecallsEnabled(config.PhonecallsEnabled)
	now := int32(time.Now().Unix())
	helpConfig.SetDate(now)
	helpConfig.SetExpires(now + EXPIRES_TIMEOUT)
	if config.TestMode == true {
		// mtproto.NewTLBoolTrue().To_Bool()
		helpConfig.SetTestMode(mtproto.ToBool(true))
		// MakeBool(new(mtproto.TLBoolTrue))
	} else {
		helpConfig.SetTestMode(mtproto.ToBool(false))
		// MakeBool(new(mtproto.TLBoolFalse))
	}
	helpConfig.SetThisDc(config.ThisDc)
	// = config.ThisDc
	dcOptions := make([]*mtproto.DcOption, 0, len(config.DcOptions))
	for _, opt := range config.DcOptions {
		dcOption := mtproto.NewTLDcOption()
		dcOption.SetIpv6(opt.Ipv6)
		dcOption.SetMediaOnly(opt.MediaOnly)
		dcOption.SetTcpoOnly(opt.TcpoOnly)
		dcOption.SetCdn(opt.Cdn)
		dcOption.SetStatic(opt.Static)
		dcOption.SetId(opt.Id)
		dcOption.SetIpAddress(opt.IpAddress)
		dcOption.SetPort(opt.Port)
		dcOptions = append(dcOptions, dcOption.To_DcOption())
		// helpConfig.SetDcOptions = append(helpConfig.DcOptions, mtproto.MakeDcOption(dcOption))
	}
	helpConfig.SetDcOptions(dcOptions)

	helpConfig.SetChatSizeMax(config.ChatSizeMax)
	helpConfig.SetMegagroupSizeMax(config.MegagroupSizeMax)
	helpConfig.SetForwardedCountMax(config.ForwardedCountMax)
	helpConfig.SetOnlineUpdatePeriodMs(config.OnlineUpdatePeriodMs)
	helpConfig.SetOfflineBlurTimeoutMs(config.OfflineBlurTimeoutMs)
	helpConfig.SetOnlineCloudTimeoutMs(config.OnlineCloudTimeoutMs)
	helpConfig.SetNotifyCloudDelayMs(config.NotifyCloudDelayMs)
	helpConfig.SetNotifyDefaultDelayMs(config.NotifyDefaultDelayMs)
	helpConfig.SetChatBigSize(config.ChatBigSize)
	helpConfig.SetPushChatPeriodMs(config.PushChatPeriodMs)
	helpConfig.SetPushChatLimit(config.PushChatLimit)
	helpConfig.SetSavedGifsLimit(config.SavedGifsLimit)
	helpConfig.SetEditTimeLimit(config.EditTimeLimit)
	helpConfig.SetRatingEDecay(config.RatingEDecay)
	helpConfig.SetStickersRecentLimit(config.StickersRecentLimit)
	helpConfig.SetStickersFavedLimit(config.StickersFavedLimit)
	helpConfig.SetTmpSessions(config.TmpSessions)
	helpConfig.SetPinnedDialogsCountMax(config.PinnedDialogsCountMax)
	helpConfig.SetCallReceiveTimeoutMs(config.CallReceiveTimeoutMs)
	helpConfig.SetCallRingTimeoutMs(config.CallRingTimeoutMs)
	helpConfig.SetCallConnectTimeoutMs(config.CallConnectTimeoutMs)
	helpConfig.SetCallPacketTimeoutMs(config.CallPacketTimeoutMs)
	helpConfig.SetMeUrlPrefix(config.MeUrlPrefix)
	helpConfig.SetSuggestedLangCode(config.SuggestedLangCode)
	helpConfig.SetLangPackVersion(config.LangPackVersion)

	disabledFeatures := make([]*mtproto.DisabledFeature, 0, len(config.DisabledFeatures))
	for _, disabled := range config.DisabledFeatures {
		disabledFeature := mtproto.NewTLDisabledFeature()
		disabledFeature.SetFeature(disabled.Feature)
		disabledFeature.SetDescription(disabled.Description)
		disabledFeatures = append(disabledFeatures, disabledFeature.To_DisabledFeature())
		// helpConfig.DisabledFeatures = append(helpConfig.DisabledFeatures, mtproto.MakeDisabledFeature(disabledFeature))
	}
	helpConfig.SetDisabledFeatures(disabledFeatures)

	reply := helpConfig.To_Config()
	// mtproto.MakeConfig(helpConfig)
	glog.Infof("HelpGetConfig - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}
