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

type UpdatesServiceImpl struct {
}

func (s *UpdatesServiceImpl) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *UpdatesServiceImpl) UpdatesGetDifference(ctx context.Context, request *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *UpdatesServiceImpl) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
