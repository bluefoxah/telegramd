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
	"testing"
	"github.com/nebulaim/telegramd/frontend/id"
	"time"
	"fmt"
)

func TestRpcMetaData(t *testing.T) {
	rpcMetadata := &RpcMetaData{}
	rpcMetadata.ServerId = 1
	rpcMetadata.UserId = 2
	rpcMetadata.AuthId = 100000
	rpcMetadata.SessionId = 100001
	rpcMetadata.ClientAddr = "127.0.0.1"
	rpcMetadata.TraceId = id.NextId()
	rpcMetadata.SpanId = id.NextId()
	rpcMetadata.ReceiveTime = time.Now().Unix()


	md := rpcMetadata.Encode()

	rpcMetadata2 := &RpcMetaData{}
	rpcMetadata2.Decode(md)

	fmt.Printf("rpcMetadata: {%v}, md: {%v}, rpcMetadata2: {%v}", rpcMetadata, md, rpcMetadata2)
}
