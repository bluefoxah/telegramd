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
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/biz_server/help/model"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"time"
)

const (
	CONFIG_FILE = "./config.toml"

	// date = 1509066502,    2017/10/27 09:08:22
	// expires = 1509070295, 2017/10/27 10:11:35
	EXPIRES_TIMEOUT = 3600 // 超时时间设置为3600秒
)

var config model.Config

func init()  {
	if _, err := toml.DecodeFile(CONFIG_FILE, &config); err != nil {
		panic(err)
	}
}

type HelpServiceImpl struct {
}

func (s *HelpServiceImpl) HelpSaveAppLog(ctx context.Context, request *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpSetBotUpdatesStatus(ctx context.Context, request *mtproto.TLHelpSetBotUpdatesStatus) (*mtproto.Bool, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetAppChangelog(ctx context.Context, request *mtproto.TLHelpGetAppChangelog) (*mtproto.Updates, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	glog.Infof("HelpGetConfig - Process: {%v}", request)

	// TODO(@benqi): 设置Reply对象累死人了, 得想个办法实现model和mtproto的自动转换
	helpConfig := &mtproto.TLConfig{}
	helpConfig.PhonecallsEnabled = config.PhonecallsEnabled
	helpConfig.Date = int32(time.Now().Unix())
	helpConfig.Expires = helpConfig.Date + EXPIRES_TIMEOUT
	if config.TestMode == true {
		helpConfig.TestMode = mtproto.MakeBool(new(mtproto.TLBoolTrue))
	} else {
		helpConfig.TestMode = mtproto.MakeBool(new(mtproto.TLBoolFalse))
	}
	helpConfig.ThisDc = config.ThisDc
	for _, opt := range config.DcOptions {
		dcOption := &mtproto.TLDcOption{}
		dcOption.Ipv6 = opt.Ipv6
		dcOption.MediaOnly = opt.MediaOnly
		dcOption.TcpoOnly = opt.TcpoOnly
		dcOption.Cdn = opt.Cdn
		dcOption.Static = opt.Static
		dcOption.Id = opt.Id
		dcOption.IpAddress = opt.IpAddress
		dcOption.Port = opt.Port
		helpConfig.DcOptions = append(helpConfig.DcOptions, mtproto.MakeDcOption(dcOption))
	}
	helpConfig.ChatSizeMax = config.ChatSizeMax
	helpConfig.MegagroupSizeMax = config.MegagroupSizeMax
	helpConfig.ForwardedCountMax = config.ForwardedCountMax
	helpConfig.OnlineUpdatePeriodMs = config.OnlineUpdatePeriodMs
	helpConfig.OfflineBlurTimeoutMs = config.OfflineBlurTimeoutMs
	helpConfig.OnlineCloudTimeoutMs = config.OnlineCloudTimeoutMs
	helpConfig.NotifyCloudDelayMs = config.NotifyCloudDelayMs
	helpConfig.NotifyDefaultDelayMs = config.NotifyDefaultDelayMs
	helpConfig.ChatBigSize = config.ChatBigSize
	helpConfig.PushChatPeriodMs = config.PushChatPeriodMs
	helpConfig.PushChatLimit = config.PushChatLimit
	helpConfig.SavedGifsLimit = config.SavedGifsLimit
	helpConfig.EditTimeLimit = config.EditTimeLimit
	helpConfig.RatingEDecay = config.RatingEDecay
	helpConfig.StickersRecentLimit = config.StickersRecentLimit
	helpConfig.StickersFavedLimit = config.StickersFavedLimit
	helpConfig.TmpSessions = config.TmpSessions
	helpConfig.PinnedDialogsCountMax = config.PinnedDialogsCountMax
	helpConfig.CallReceiveTimeoutMs = config.CallReceiveTimeoutMs
	helpConfig.CallRingTimeoutMs = config.CallRingTimeoutMs
	helpConfig.CallConnectTimeoutMs = config.CallConnectTimeoutMs
	helpConfig.CallPacketTimeoutMs = config.CallPacketTimeoutMs
	helpConfig.MeUrlPrefix = config.MeUrlPrefix
	helpConfig.SuggestedLangCode = config.SuggestedLangCode
	helpConfig.LangPackVersion = config.LangPackVersion

	for _, disabled := range config.DisabledFeatures {
		disabledFeature := &mtproto.TLDisabledFeature{
			Feature:     disabled.Feature,
			Description: disabled.Description,
		}
		helpConfig.DisabledFeatures = append(helpConfig.DisabledFeatures, mtproto.MakeDisabledFeature(disabledFeature))
	}

	reply := mtproto.MakeConfig(helpConfig)
	// glog.Infof("HelpGetConfig - reply: {%v}\n", reply)
	return reply, nil
}

func (s *HelpServiceImpl) HelpGetNearestDc(ctx context.Context, request *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	glog.Infof("HelpGetNearestDc - Process: {%v}", request)

	dc := &mtproto.TLNearestDc{}
	dc.Country = "US"
	dc.ThisDc = 2
	dc.NearestDc = 2

	reply := mtproto.MakeNearestDc(dc)
	glog.Infof("HelpGetNearestDc - reply: {%v}\n", reply)
	return reply, nil
}

func (s *HelpServiceImpl) HelpGetAppUpdate(ctx context.Context, request *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetInviteText(ctx context.Context, request *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetSupport(ctx context.Context, request *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetTermsOfService(ctx context.Context, request *mtproto.TLHelpGetTermsOfService) (*mtproto.Help_TermsOfService, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetCdnConfig(ctx context.Context, request *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
}
