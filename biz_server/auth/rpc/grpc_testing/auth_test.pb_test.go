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

package grpc_testing

import (
	"github.com/golang/protobuf/proto"
	"testing"
	"fmt"
	. "github.com/nebulaim/telegramd/mtproto"
)

// auth.sentCode#5e002502
// flags:# phone_registered:flags.0?true
// type:auth.SentCodeType
// phone_code_hash:string
// next_type:flags.1?auth.CodeType
// timeout:flags.2?int
// = auth.SentCode;
func (m *TLAuthSentCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCode))

	flags := uint32(0)
	if m.PhoneRegistered == true {
		flags |= 1<<0
	}
	if m.NextType != nil {
		flags |= 1<<1
	}
	if m.Timeout != 0 {
		flags |= 1<<2
	}
	x.UInt(flags)
	x.Bytes(m.Type.Encode())
	x.String(m.PhoneCodeHash)
	if (flags & (1 << 1)) != 0 {
		x.Bytes(m.NextType.Encode())
	}
	if (flags & (1 << 2)) != 0 {
		x.Int(m.Timeout)
	}
	return x.GetBuf()
}

func (m *TLAuthSentCode) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()

	if (flags & (1 << 0)) != 0 {
		m.PhoneRegistered = true
	}

	m.Type = &Auth_SentCodeType{}
	m.Decode(dbuf)
	m.PhoneCodeHash = dbuf.String()

	if (flags & (1 << 1)) != 0 {
		m.NextType = &Auth_CodeType{}
		m.Decode(dbuf)
	}

	if (flags & (1 << 2)) != 0 {
		m.Timeout = dbuf.Int()
	}

	return dbuf.GetError()
}

func MakeAuth_SentCode(message proto.Message) (m *Auth_SentCode) {
	switch message.(type) {
	case *TLAuthSentCode:
		m2, _ := message.(*TLAuthSentCode)
		m = &Auth_SentCode{
			Payload: &Auth_SentCode_AuthSentCode{
				AuthSentCode: m2,
			},
		}
	}
	return
}

func TestMakeAuth_SentCode(t *testing.T) {
	m2 := &TLAuthSentCode{
		Flags: 1,
	}

	// m2 := &TLAuthSendCode{}

	fmt.Printf("{%v}\n", MakeAuth_SentCode(m2))
}
