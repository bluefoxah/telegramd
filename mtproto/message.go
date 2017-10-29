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
	"fmt"
	"encoding/binary"
	"bytes"
	"github.com/golang/glog"
	"encoding/hex"
	"github.com/nebulaim/telegramd/base/crypto"
)

const (
	QUICK_ACKID = iota
	UNENCRYPTED_MESSAGE
	ENCRYPTED_MESSAGE
)

type MTProtoMessage interface {
	// encode([]byte) ([]byte, error)
	// decode([]byte) error
	// MessageType() int
}

type QuickAckMessage struct {
	ackId int32
}

func (m *QuickAckMessage) MessageType() int {
	return QUICK_ACKID
}

func (m *QuickAckMessage) encode() ([]byte, error) {
	return nil, nil
}

func (m *QuickAckMessage) decode(b []byte) error {
	if len(b) != 4 {
		return fmt.Errorf("Message len: %d (need 4)", len(b))
	}
	m.ackId = int32(binary.LittleEndian.Uint32(b))
	return nil
}

type UnencryptedMessage struct {
	NeedAck bool

	// authKeyId int64
	MessageId int64
	// messageDataLength int32
	// messageData []byte

	// classID int32
	Object TLObject
}

func (m *UnencryptedMessage) MessageType() int {
	return UNENCRYPTED_MESSAGE
}

func (m *UnencryptedMessage) encode() ([]byte, error) {
	x := NewEncodeBuf(512)
	x.Long(0)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)
	b := m.Object.Encode()
	x.Int(int32(len(b)))
	x.Bytes(b)

	// glog.Info("Encode object: ", m.Object)
	return x.buf, nil
}

func (m *UnencryptedMessage) decode(b []byte) error {
	dbuf := NewDecodeBuf(b)
	// m.authKeyId = dbuf.Long()
	m.MessageId = dbuf.Long()

	// glog.Info("messageId:", m.messageId)
	// mod := m.messageId & 3
	// if mod != 1 && mod != 3 {
	// 	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	messageLen := dbuf.Int()
	// glog.Info("messageLen:", m.messageId)

	if int(messageLen) != dbuf.size-12 {
		return fmt.Errorf("Message len: %d (need %d)", messageLen, dbuf.size-12)
	}

	m.Object = dbuf.Object()
	if m.Object == nil {
		return fmt.Errorf("Decode object is nil")
	}

	// proto.Message()
	// glog.Info("Recved object: ", m.Object.(proto.Message).String())
	return dbuf.err
}

// TODO(@benqi): 将Encrypt和Descrypt移到底层
type EncryptedMessage2 struct {
	// AuthKeyId int64
	NeedAck bool

	msgKey []byte
	salt int64
	SessionId int64
	MessageId int64
	SeqNo int32
	Object TLObject
}

func (m *EncryptedMessage2) MessageType() int {
	return ENCRYPTED_MESSAGE
}

func (m *EncryptedMessage2) encode(authKeyId int64, authKey []byte) ([]byte, error) {
	objData := m.Object.Encode()
	var additional_size = (32 + len(objData)) % 16
	if (additional_size != 0) {
		additional_size = 16 - additional_size
	}
	if MTPROTO_VERSION == 2 && additional_size < 12 {
		additional_size += 16
	}

	x := NewEncodeBuf(32+len(objData)+additional_size)
	// x.Long(authKeyId)
	// msgKey := make([]byte, 16)
	// x.Bytes(msgKey)
	x.Long(m.salt)
	x.Long(m.SessionId)
	m.MessageId = GenerateMessageId()
	x.Long(m.MessageId)
	x.Int(m.SeqNo)
	x.Int(int32(len(objData)))
	x.Bytes(objData)
	x.Bytes(crypto.GenerateNonce(additional_size))

	// glog.Info("Encode object: ", m.Object)

	encryptedData, _ := m.encrypt(authKey, x.buf)
	x2 := NewEncodeBuf(56+len(objData)+additional_size)
	x2.Long(authKeyId)
	x2.Bytes(m.msgKey)
	x2.Bytes(encryptedData)

	return x2.buf, nil
}

func (m *EncryptedMessage2) decode(authKey []byte, b []byte) error {
	msgKey := b[:16]
	// aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	// x, err := doAES256IGEdecrypt(b[16:], aesKey, aesIV)

	x, err := m.descrypt(msgKey, authKey, b[16:])
	if err != nil {
		return err
	}

	dbuf := NewDecodeBuf(x)

	m.salt = dbuf.Long() // salt
	m.SessionId = dbuf.Long() // session_id
	m.MessageId = dbuf.Long()

	// mod := m.messageId & 3
	// if mod != 1 && mod != 3 {
	//	return fmt.Errorf("Wrong bits of message_id: %d", mod)
	// }

	m.SeqNo = dbuf.Int()
	messageLen := dbuf.Int()
	if int(messageLen) > dbuf.size-32 {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
	}

	// glog.Infof("salt: %d, sessionId: %d, messageId: %d, seqNo: %d, messageLen: %d", m.salt, m.SessionId, m.MessageId, m.SeqNo, messageLen)
	m.Object = dbuf.Object()
	if m.Object == nil {
		return fmt.Errorf("Decode object is nil")
	}

	// glog.Info("Recved object: ", m.Object.String())

	return nil
}

func (m *EncryptedMessage2) descrypt(msgKey, authKey, data []byte) ([]byte, error) {
	// dbuf := NewDecodeBuf(data)
	// m.authKeyId = dbuf.Long()
	// msgKey := dbuf.Bytes(16)

	var dataLen = uint32(len(data))
	// 创建aesKey, aesIV
	aesKey, aesIV := generateMessageKey(msgKey, authKey, false)
	d := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := d.Decrypt(data)
	if err != nil {
		glog.Error("descrypted data error: ", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	messageLen := binary.LittleEndian.Uint32(x[28:])
	// glog.Info("descrypt - messageLen = ", messageLen)
	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("descrypted data error: Wrong message length %d", messageLen)
		glog.Error(err)
		return nil, err
	}

	sha256MsgKey := make([]byte, 32)
	switch MTPROTO_VERSION {
	case 2:
		t_d := make([]byte, 0, 32 + dataLen)
		t_d = append(t_d, authKey[88:88+32]...)
		t_d = append(t_d, x[:dataLen]...)
		copy(sha256MsgKey, crypto.Sha256Digest(t_d))
	default:
		copy(sha256MsgKey[4:], crypto.Sha1Digest(x))
	}

	if !bytes.Equal(sha256MsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("descrypted data error: msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(sha256MsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		glog.Error(err)
		return nil, err
	}

	return x, nil
}

func (m *EncryptedMessage2) encrypt(authKey []byte, data []byte) ([]byte, error) {
	message_key := make([]byte, 32)
	switch MTPROTO_VERSION {
	case 2:
		t_d := make([]byte, 0, 32+len(data))
		t_d = append(t_d, authKey[88+8:88+8+32]...)
		t_d = append(t_d, data...)
		copy(message_key, crypto.Sha256Digest(t_d))
	default:
		copy(message_key[4:], crypto.Sha1Digest(data))
	}


	// copy(message_key[8:], )
	// memcpy(p_data + 8, message_key + 8, 16);

	aesKey, aesIV := generateMessageKey(message_key[8:8+16], authKey, true)
	e := crypto.NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(data)
	if err != nil {
		glog.Error("Encrypt data error: ", err)
		return nil, err
	}

	m.msgKey = message_key[8:8+16]
	return x, nil
}
