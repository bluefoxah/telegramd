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

type PhoneServiceImpl struct {
}

func (s *PhoneServiceImpl) PhoneReceivedCall(ctx context.Context, request *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneSaveCallDebug(ctx context.Context, request *mtproto.TLPhoneSaveCallDebug) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneDiscardCall(ctx context.Context, request *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneSetCallRating(ctx context.Context, request *mtproto.TLPhoneSetCallRating) (*mtproto.Updates, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneGetCallConfig(ctx context.Context, request *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneRequestCall(ctx context.Context, request *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneAcceptCall(ctx context.Context, request *mtproto.TLPhoneAcceptCall) (*mtproto.Phone_PhoneCall, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PhoneServiceImpl) PhoneConfirmCall(ctx context.Context, request *mtproto.TLPhoneConfirmCall) (*mtproto.Phone_PhoneCall, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
