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
	"fmt"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
)

// import (
// )

// req_pq#60469778 nonce:int128 = ResPQ;
//type TLReqPq struct {
//	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
//}

func (m []*TLFutureSalt) Encode() []byte {
	x := NewEncodeBuf(512)
	fmt.Println(len(m))
	// x.UInt(CRC32_req_pq)
	return x.buf
}

func (m* TLReqPq) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC32_req_pq)
	x.Bytes(m.GetNonce())
	return x.buf
}

func (m *DecodeBuf) Object2() (r proto.Message) {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}

	switch constructor {

	case CRC32_req_pq:
		r = &TLReqPq{
			Nonce: m.Bytes(16),
		}

	default:
		r = nil
	}
	return
}

func TestTLReqPqEncode(t *testing.T) {
	m := TLReqPq{
		Nonce: GenerateNonce(16),
	}

	b := m.encode()
	fmt.Println(hex.EncodeToString(b))
	fmt.Println(m.String())

	dbuf := NewDecodeBuf(b)
	m2 := dbuf.Object2()

	reqPq, _ := m2.(*TLReqPq)
	fmt.Println(reqPq.String())
}
