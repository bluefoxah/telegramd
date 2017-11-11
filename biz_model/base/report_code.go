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

package base

import (
	"github.com/nebulaim/telegramd/mtproto"
)

type ReportReasonType int8

const (
	//inputReportReasonOther#e1746d0a text:string = ReportReason;
	REASON_OTHER 		ReportReasonType 	= 0		// 其它

	//inputReportReasonSpam#58dbcab8 = ReportReason;
	REASON_SPAM 		ReportReasonType	= 1		// 垃圾

	//inputReportReasonViolence#1e22c78d = ReportReason;
	REASON_VIOLENCE 	ReportReasonType	= 2		// 暴力

	//inputReportReasonPornography#2e59d922 = ReportReason;
	REASON_PORNOGRAPHY	ReportReasonType	= 3		// 色情
)

func (i ReportReasonType) String() (s string) {
	switch i {
	case REASON_OTHER:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	case REASON_SPAM:
		s = "nputReportReasonSpam#58dbcab8 = ReportReason"
	case REASON_VIOLENCE:
		s = "inputReportReasonPornography#2e59d922 = ReportReason"
	case REASON_PORNOGRAPHY:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	}
	return
}

func (i *ReportReasonType) FromReportReason(reason *mtproto.ReportReason) {
	switch reason.Payload.(type) {
	case *mtproto.ReportReason_InputReportReasonSpam:
		*i = REASON_OTHER
	case *mtproto.ReportReason_InputReportReasonViolence:
		*i = REASON_SPAM
	case *mtproto.ReportReason_InputReportReasonPornography:
		*i = REASON_VIOLENCE
	case *mtproto.ReportReason_InputReportReasonOther:
		*i = REASON_PORNOGRAPHY
	}
}

func (i ReportReasonType) ToReportReason(reason *mtproto.ReportReason) {
	switch i {
	case REASON_OTHER:
		reason = mtproto.MakeReportReason(&mtproto.TLInputReportReasonOther{})
	case REASON_SPAM:
		reason = mtproto.MakeReportReason(&mtproto.TLInputReportReasonSpam{})
	case REASON_VIOLENCE:
		reason = mtproto.MakeReportReason(&mtproto.TLInputReportReasonViolence{})
	case REASON_PORNOGRAPHY:
		reason = mtproto.MakeReportReason(&mtproto.TLInputReportReasonPornography{})
	}
	return
}
