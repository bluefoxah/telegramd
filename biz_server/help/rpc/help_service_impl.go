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
	"errors"
	"time"
)

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
	glog.Infof("Process: %v", request)

	config := &mtproto.TLConfig{}
	config.PhonecallsEnabled = true
	config.Date = int32(time.Now().Unix())
	return nil, errors.New("Not impl")
}

func (s *HelpServiceImpl) HelpGetNearestDc(ctx context.Context, request *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	glog.Infof("Process: %v", request)
	return nil, errors.New("Not impl")
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
