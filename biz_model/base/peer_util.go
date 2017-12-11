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
	// "github.com/nebulaim/telegramd/mtproto"
	"fmt"
)

const (
	PEER_EMPTY 		= 0
	PEER_SELF 		= 1
	PEER_USER 		= 2
	PEER_CHAT 		= 3
	PEER_CHANNEL 	= 4
	PEER_USERS		= 5
	PEER_CHATS 		= 6
	PEER_ALL 		= 7
	PEER_UNKNOWN 	= -1
)

type PeerUtil struct {
	PeerType 		int32
	PeerId 			int32
	AccessHash    	int64
}


func (p PeerUtil) String() (s string) {
	switch p.PeerType {
	case PEER_EMPTY:
		return fmt.Sprintf("PEER_EMPTY: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_SELF:
		return fmt.Sprintf("PEER_SELF: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USER:
		return fmt.Sprintf("PEER_USER: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHAT:
		return fmt.Sprintf("PEER_CHAT: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHANNEL:
		return fmt.Sprintf("PEER_CHANNEL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USERS:
		return fmt.Sprintf("PEER_USERS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHATS:
		return fmt.Sprintf("PEER_CHATS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_ALL:
		return fmt.Sprintf("PEER_ALL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	default:
		return fmt.Sprintf("PEER_UNKNOWN: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	}
	return
}

/*
func FromInputPeer(peer *mtproto.InputPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.Payload.(type) {
	case *mtproto.InputPeer_InputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case *mtproto.InputPeer_InputPeerSelf:
		p.PeerType = PEER_SELF
	case *mtproto.InputPeer_InputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.GetInputPeerUser().UserId
		p.AccessHash = peer.GetInputPeerUser().AccessHash
	case *mtproto.InputPeer_InputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.GetInputPeerChat().ChatId
	case *mtproto.InputPeer_InputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.GetInputPeerChannel().ChannelId
		p.AccessHash = peer.GetInputPeerChannel().AccessHash
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return
}

func (p *PeerUtil) ToInputPeer() (peer *mtproto.InputPeer) {
	switch p.PeerType {
	case PEER_EMPTY:
		peer = mtproto.MakeInputPeer(&mtproto.TLInputPeerEmpty{})
	case PEER_SELF:
		peer = mtproto.MakeInputPeer(&mtproto.TLInputPeerSelf{})
	case PEER_USER:
		p2 := &mtproto.TLInputPeerUser{}
		p2.UserId = p.PeerId
		p2.AccessHash = p.AccessHash
		peer = p2.ToInputPeer()
	case PEER_CHAT:
		p2 := &mtproto.TLInputPeerChat{}
		p2.ChatId = p.PeerId
		peer = p2.ToInputPeer()
	case PEER_CHANNEL:
		p2 := &mtproto.TLInputPeerChannel{}
		p2.ChannelId = p.PeerId
		p2.AccessHash = p.AccessHash
		peer = p2.ToInputPeer()
	default:
		panic(fmt.Sprintf("ToInputPeer(%v) error!", p))
	}
	return
}

func FromPeer(peer *mtproto.Peer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.Payload.(type) {
	case *mtproto.Peer_PeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.GetPeerUser().UserId
	case *mtproto.Peer_PeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.GetPeerChat().ChatId
	case *mtproto.Peer_PeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.GetPeerChannel().ChannelId
	default:
		panic(fmt.Sprintf("FromPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToPeer() (peer *mtproto.Peer) {
	switch p.PeerType {
	case PEER_USER:
		p2 := &mtproto.TLPeerUser{}
		p2.UserId = p.PeerId
		peer = p2.ToPeer()
	case PEER_CHAT:
		p2 := &mtproto.TLPeerChat{}
		p2.ChatId = p.PeerId
		peer = p2.ToPeer()
	case PEER_CHANNEL:
		p2 := &mtproto.TLPeerChannel{}
		p2.ChannelId = p.PeerId
		peer = p2.ToPeer()
	default:
		panic(fmt.Sprintf("ToPeer(%v) error!", p))
	}
	return
}

func FromInputNotifyPeer(peer *mtproto.InputNotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.Payload.(type) {
	case *mtproto.InputNotifyPeer_InputNotifyPeer:
		p = FromInputPeer(peer.GetInputNotifyPeer().GetPeer())
	case *mtproto.InputNotifyPeer_InputNotifyUsers:
		p.PeerType = PEER_USERS
	case *mtproto.InputNotifyPeer_InputNotifyChats:
		p.PeerType = PEER_CHATS
	case *mtproto.InputNotifyPeer_InputNotifyAll:
		p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromInputNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToInputNotifyPeer(peer *mtproto.InputNotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		p2 := &mtproto.TLInputNotifyPeer{}
		p2.Peer = p.ToInputPeer()
		peer = p2.ToInputNotifyPeer()
	case PEER_USERS:
		p2 := &mtproto.TLInputNotifyUsers{}
		peer = p2.ToInputNotifyPeer()
	case PEER_CHATS:
		p2 := &mtproto.TLInputNotifyChats{}
		peer = p2.ToInputNotifyPeer()
	case PEER_ALL:
		p2 := &mtproto.TLInputNotifyAll{}
		peer = p2.ToInputNotifyPeer()
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", p))
	}
	return
}

func FromNotifyPeer(peer *mtproto.NotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.Payload.(type) {
	case *mtproto.NotifyPeer_NotifyPeer:
		p = FromPeer(peer.GetNotifyPeer().GetPeer())
	case *mtproto.NotifyPeer_NotifyUsers:
		p.PeerType = PEER_USERS
	case *mtproto.NotifyPeer_NotifyChats:
		p.PeerType = PEER_CHATS
	case *mtproto.NotifyPeer_NotifyAll:
		p.PeerType = PEER_ALL
	default:
		panic(fmt.Sprintf("FromNotifyPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToNotifyPeer() (peer *mtproto.NotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		p2 := &mtproto.TLNotifyPeer{}
		p2.Peer = p.ToPeer()
		peer = p2.ToNotifyPeer()
	case PEER_USERS:
		p2 := &mtproto.TLNotifyUsers{}
		peer = p2.ToNotifyPeer()
	case PEER_CHATS:
		p2 := &mtproto.TLNotifyChats{}
		peer = p2.ToNotifyPeer()
	case PEER_ALL:
		p2 := &mtproto.TLNotifyAll{}
		peer = p2.ToNotifyPeer()
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", p))
	}
	return
}
*/
