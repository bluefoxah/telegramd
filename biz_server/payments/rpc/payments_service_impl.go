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
	"github.com/nebulaim/telegramd/base/orm"
)

type PaymentsServiceImpl struct {
	zorm orm.Ormer
}

func (s *PaymentsServiceImpl) PaymentsClearSavedInfo(ctx context.Context, request *mtproto.TLPaymentsClearSavedInfo) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PaymentsServiceImpl) PaymentsGetPaymentForm(ctx context.Context, request *mtproto.TLPaymentsGetPaymentForm) (*mtproto.Payments_PaymentForm, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PaymentsServiceImpl) PaymentsGetPaymentReceipt(ctx context.Context, request *mtproto.TLPaymentsGetPaymentReceipt) (*mtproto.Payments_PaymentReceipt, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PaymentsServiceImpl) PaymentsValidateRequestedInfo(ctx context.Context, request *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PaymentsServiceImpl) PaymentsSendPaymentForm(ctx context.Context, request *mtproto.TLPaymentsSendPaymentForm) (*mtproto.Payments_PaymentResult, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *PaymentsServiceImpl) PaymentsGetSavedInfo(ctx context.Context, request *mtproto.TLPaymentsGetSavedInfo) (*mtproto.Payments_SavedInfo, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}
