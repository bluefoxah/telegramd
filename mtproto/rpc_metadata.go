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

package mtproto

import (
	"google.golang.org/grpc/metadata"
	"strconv"
	"github.com/nebulaim/telegramd/base/base"
)

type RpcMetaData struct {
	ServerId	int32
	NetlibSessionId int64
	ClientAddr	string
	AuthId		int64
	SessionId	int64

	TraceId 	int64
	SpanId		int64
	ReceiveTime int64

	// 用户ID
	UserId		int32
}

func getFirstKeyVal(md metadata.MD, k string) (string, bool) {
	if v, ok := md[k]; ok {
		if len(v) > 0 {
			return v[0], ok
		}
	}

	return "", false
}

func (m *RpcMetaData) Decode(md metadata.MD) {
	if v, ok := getFirstKeyVal(md, "serverid"); ok {
		m.ServerId, _ = base.StringToInt32(v)
	}

	if v, ok := getFirstKeyVal(md, "clientaddr"); ok {
		m.ClientAddr = v
	}

	if v, ok := getFirstKeyVal(md, "authid"); ok {
		m.AuthId, _ = base.StringToInt64(v)
	}

	if v, ok := getFirstKeyVal(md, "sessionid"); ok {
		m.SessionId, _ = base.StringToInt64(v)
	}

	if v, ok := getFirstKeyVal(md, "traceid"); ok {
		m.TraceId, _ = base.StringToInt64(v)
	}

	if v, ok := getFirstKeyVal(md, "spanid"); ok {
		m.SpanId, _ = base.StringToInt64(v)
	}

	if v, ok := getFirstKeyVal(md, "receivetime"); ok {
		m.ReceiveTime, _ = base.StringToInt64(v)
	}

	if v, ok := getFirstKeyVal(md, "userid"); ok {
		m.UserId, _ = base.StringToInt32(v)
	}
}

func (m *RpcMetaData) Encode() (metadata.MD) {
	return metadata.Pairs(
		"serverid", strconv.FormatInt(int64(m.ServerId), 10),
		"clientaddr", m.ClientAddr,
		"authid", strconv.FormatInt(m.AuthId, 10),
		"sessionid", strconv.FormatInt(m.SessionId, 10),
		"traceid", strconv.FormatInt(m.TraceId, 10),
		"spanid", strconv.FormatInt(m.SpanId, 10),
		"receivetime", strconv.FormatInt(m.ReceiveTime, 10),
		"userid", strconv.FormatInt(int64(m.UserId), 10))
}
