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

package model

// TODO(@benqi): 配置中心管理配置
// dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
type DcOption struct {
	Ipv6      bool
	MediaOnly bool
	TcpoOnly  bool
	Cdn       bool
	Static    bool
	Id        int32
	IpAddress string
	Port      int32
}

type DisabledFeature struct {
	Feature     string
	Description string
}

type Config struct {
	PhonecallsEnabled     bool
	Date                  int32
	Expires               int32
	TestMode              bool
	ThisDc                int32
	DcOptions             []DcOption
	ChatSizeMax           int32
	MegagroupSizeMax      int32
	ForwardedCountMax     int32
	OnlineUpdatePeriodMs  int32
	OfflineBlurTimeoutMs  int32
	OfflineIdleTimeoutMs  int32
	OnlineCloudTimeoutMs  int32
	NotifyCloudDelayMs    int32
	NotifyDefaultDelayMs  int32
	ChatBigSize           int32
	PushChatPeriodMs      int32
	PushChatLimit         int32
	SavedGifsLimit        int32
	EditTimeLimit         int32
	RatingEDecay          int32
	StickersRecentLimit   int32
	StickersFavedLimit    int32
	TmpSessions           int32
	PinnedDialogsCountMax int32
	CallReceiveTimeoutMs  int32
	CallRingTimeoutMs     int32
	CallConnectTimeoutMs  int32
	CallPacketTimeoutMs   int32
	MeUrlPrefix           string
	SuggestedLangCode     string
	LangPackVersion       int32
	DisabledFeatures      []DisabledFeature
}
