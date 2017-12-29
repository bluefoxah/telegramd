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

// TODO(@benqi): 考虑专门独立出session_server

package client

import (
	"sync"
	"github.com/golang/glog"
)

var cacheAuthSessionGroup sync.Map

// import "sync"

//const (
//	AUTH_SESSION_STATE_UNKNOWN = 0
//	// AUTH_SESSION_STATE_UNKNOWN = 0
//	// AUTH_SESSION_STATE_UNKNOWN = 0
//	// AUTH_SESSION_STATE_UNKNOWN = 0
//)
//

const (
	GENERIC = 0
	DOWNLOAD = 1
	UPLOAD = 3
	PUSH = 7
	TEMP = 8

	// INVALID = math.MaxInt32
)

const (
	DOWNLOAD_CONNECTIONS_COUNT = 2 	// Download conn count
	UPLOAD_CONNECTIONS_COUNT = 4	//
	MAX_CONNECTIONS_COUNT = 9		// 最大连接数为9
)

//type AuthSessionType int
//type AuthSessionState int

//type AuthSessionCache struct {
//	authSessions map[string]*AuthSession
//	mu           sync.Mutex
//}

type AuthSession struct {
	Id              int64
	Type            int   	// 连接类型
	NetlibSessionId int64 	// 网络连接SessionId，为0认为离线

	used 			bool	//
	//SessionIds []int64
	//Salts      []int64
	MsgIds []int64
	//
}

type AuthSessionGroup struct {
	AuthSessionList []*AuthSession
	lock *sync.RWMutex
}

func NewAuthSessionGroup() *AuthSessionGroup {
	return &AuthSessionGroup{
		AuthSessionList: make([]*AuthSession, MAX_CONNECTIONS_COUNT),
		lock:            new(sync.RWMutex),
	}
}

func (m *AuthSessionGroup) GetOrCreate(sessId int64) (sess *AuthSession) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, s := range m.AuthSessionList {
		if s != nil || (s.Id == sessId && !s.used) {
			s.used = true
			sess = s
			break
		} else {
			glog.Errorf("AuthSessionGroup - get error: found but used sess {%v}", s)
		}
	}

	if sess == nil {
		sess = &AuthSession{
			Id:   0,
			Type: -1,
		}
	}
	return
}

func (m *AuthSessionGroup) Update(sess *AuthSession) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	switch sess.Type {
	case GENERIC, PUSH, TEMP:
		if m.AuthSessionList[sess.Type] != nil {
			glog.Infof("AuthSessionGroup - Set: existed session {%v}, set session {%v}", m.AuthSessionList[sess.Type], sess)
		}

		sess.used = sess.NetlibSessionId != 0
		m.AuthSessionList[sess.Type] = sess
	case DOWNLOAD, UPLOAD:
		var sesses []*AuthSession
		if sess.Type == DOWNLOAD {
			sesses = m.AuthSessionList[DOWNLOAD:DOWNLOAD+DOWNLOAD_CONNECTIONS_COUNT]
		} else {
			sesses = m.AuthSessionList[UPLOAD:UPLOAD+UPLOAD_CONNECTIONS_COUNT]
		}
		var found = false
		for i, s := range sesses {
			if s == nil || s.Id == sess.Id {
				sess.used = sess.NetlibSessionId != 0
				if sess.Type == DOWNLOAD {
					m.AuthSessionList[DOWNLOAD+i] = sess
				} else {
					m.AuthSessionList[UPLOAD+i] = sess
				}
				found = true
				break
			}
		}
		if !found {
			glog.Infof("AuthSessionGroup - Set download: {%v}", sesses)
		}
	}
}

func GetOrCreateSession(keyId, sessId int64) *AuthSession {
	var sessGroup *AuthSessionGroup
	if k, ok := cacheAuthSessionGroup.Load(keyId); ok {
		// 本地缓存命中
		sessGroup = k.(*AuthSessionGroup)
	} else {
		sessGroup = NewAuthSessionGroup()

		// TODO(@benqi): 可能会被其它线程覆盖，但关系不大
		cacheAuthSessionGroup.Store(keyId, sessGroup)
	}

	return sessGroup.GetOrCreate(sessId)
}

func UpdateAuthSession(keyId int64, sess *AuthSession) {
	var sessGroup *AuthSessionGroup
	if k, ok := cacheAuthSessionGroup.Load(keyId); ok {
		// 本地缓存命中
		sessGroup = k.(*AuthSessionGroup)
		sessGroup.Update(sess)
	} else {
		glog.Errorf("UpdateAuthSession - Not execute GetOrCreateSession!!!!")
	}
}

