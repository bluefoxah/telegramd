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
	"github.com/golang/glog"
	"fmt"
	"encoding/hex"
)

//const (
//	TLConstructor_CRC32_message2  		= 0x5bb8e511
//	TLConstructor_CRC32_msg_container  	= 0x73f1f8dc
//	TLConstructor_CRC32_msg_copy  		= 0xe06046b2
//	TLConstructor_CRC32_gzip_packed 	= 0x3072cfa1
//)

//message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessage2 struct {
	MsgId 	int64
	Seqno 	int32
	Bytes 	int32
	Object 	TLObject
}

func (m *TLMessage2) String() string {
	return "{message2#5bb8e511}"
}

func (m *TLMessage2) Encode() []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(TLConstructor_CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.StringBytes(m.Object.Encode())
	return x.buf
}

func (m *TLMessage2) Decode(dbuf *DecodeBuf) error {
	m.MsgId = dbuf.Long()
	m.Seqno = dbuf.Int()
	m.Bytes = dbuf.Int()
	b := dbuf.Bytes(int(m.Bytes))

	dbuf2 := NewDecodeBuf(b)
	m.Object = dbuf2.Object()
	if m.Object == nil {
		err := fmt.Errorf("Decode core_message error: %s", hex.EncodeToString(b))
		glog.Error(err)
		return err
	}

	glog.Info("Sucess decoded core_message: ", m.Object.String())
	return dbuf2.err
}

//msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgContainer struct {
	Messages []TLMessage2
}

func (m *TLMsgContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgContainer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_container))
	x.Int(int32(len(m.Messages)))
	for _, m := range m.Messages {
		x.Bytes(m.Encode())
	}
	return x.buf
}

func (m *TLMsgContainer) Decode(dbuf *DecodeBuf) error {
	len := dbuf.Int()
	glog.Info("TLMsgContainer: messages len: ", len)
	for i := 0; i < int(len); i++ {
		glog.Infof("TLMsgContainer: messages[%d]: ", i)
		// classID := dbuf.Int()
		// if classID != (int32)(TLConstructor_CRC32_message2) {
		// 	err := fmt.Errorf("Decode TL_message2 error, invalid TL_message2 classID, classID: 0x%x", uint32(classID))
		// 	glog.Error(err)
		// 	return err
		// }
		message2 := &TLMessage2{}
		err := message2.Decode(dbuf)
		if err != nil {
			glog.Error("Decode message2 error: ", err)
			return err
		}

		m.Messages = append(m.Messages, *message2)
	}
	return dbuf.err
}

//msg_copy#e06046b2 orig_message:Message2 = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	OrigMessage TLMessage2
}

func (m *TLMsgCopy) String() string {
	return "{msg_copy#e06046b2}"
}

func (m *TLMsgCopy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_msg_copy))
	x.Bytes(m.OrigMessage.Encode())
	return x.buf
}

func (m *TLMsgCopy) Decode(dbuf *DecodeBuf) error {
	o := dbuf.Object()
	message2, _ := o.(*TLMessage2)
	m.OrigMessage = *message2
	return dbuf.err
}

//gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TLGzipPacked struct {
	PackedData []byte
}

func (m *TLGzipPacked) String() string {
	return "{gzip_packed#3072cfa1}"
}

func (m *TLGzipPacked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_gzip_packed))
	x.Bytes(m.PackedData)
	return x.buf
}

func (m *TLGzipPacked) Decode(dbuf *DecodeBuf) error {
	m.PackedData = dbuf.buf
	return dbuf.err
}
