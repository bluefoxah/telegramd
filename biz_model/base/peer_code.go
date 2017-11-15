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

import "github.com/nebulaim/telegramd/mtproto"

type PeerType int32

const (
	PEER_EMPTY = 0
	PEER_SELF = 1
	PEER_USER = 2
	PEER_CHAT = 3
	PEER_CHANNEL = 4
	PEER_INVALID = 5
)

func (i PeerType) String() (s string) {
	switch i {
	case PEER_EMPTY:
		s = "inputPeerEmpty#7f3b18ea = InputPeer"
	case PEER_SELF:
		s = "inputPeerSelf#7da07ec9 = InputPeer"
	case PEER_USER:
		s = "inputPeerChat#179be863 chat_id:int = InputPeer"
	case PEER_CHAT:
		s = "inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer"
	case PEER_CHANNEL:
		s = "inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer"
	}
	return
}

func FromInputPeer(reason *mtproto.InputPeer) (i PeerType) {
	switch reason.Payload.(type) {
	case *mtproto.InputPeer_InputPeerEmpty:
		i = PEER_EMPTY
	case *mtproto.InputPeer_InputPeerSelf:
		i = PEER_SELF
	case *mtproto.InputPeer_InputPeerUser:
		i = PEER_USER
	case *mtproto.InputPeer_InputPeerChat:
		i = PEER_CHAT
	case *mtproto.InputPeer_InputPeerChannel:
		i = PEER_CHANNEL
	default:
		i = PEER_INVALID
	}
	return
}

func (i PeerType) ToInputPeer(reason *mtproto.InputPeer) {
	switch i {
	case PEER_EMPTY:
		reason = mtproto.MakeInputPeer(&mtproto.TLInputPeerEmpty{})
	case PEER_SELF:
		reason = mtproto.MakeInputPeer(&mtproto.TLInputPeerSelf{})
	case PEER_USER:
		reason = mtproto.MakeInputPeer(&mtproto.TLInputPeerUser{})
	case PEER_CHAT:
		reason = mtproto.MakeInputPeer(&mtproto.TLInputPeerChat{})
	case PEER_CHANNEL:
		reason = mtproto.MakeInputPeer(&mtproto.TLInputPeerChannel{})
	}
	return
}
