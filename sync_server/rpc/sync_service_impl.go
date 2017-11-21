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

package rpc

import (
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/biz_model/model"
	"sync"
	"errors"
	"github.com/golang/glog"
)

type SyncServiceImpl struct {
	// status *model.OnlineStatusModel

	mu sync.RWMutex

	// TODO(@benqi): 多个连接
	updates map[int32]chan *zproto.PushUpdatesData
}

func NewSyncService() *SyncServiceImpl {
	s := &SyncServiceImpl{}
	// s.status = status
	s.updates = make(map[int32]chan *zproto.PushUpdatesData)
	return s
}

func (s *SyncServiceImpl) withReadLock(f func()) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f()
}

func (s *SyncServiceImpl) withWriteLock(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f()
}

func (s *SyncServiceImpl) unsafeExpire(sid int32) {
	if buf, ok := s.updates[sid]; ok {
		close(buf)
	}
	delete(s.updates, sid)
}

func (s *SyncServiceImpl) PushUpdatesStream(auth *zproto.ServerAuthReq, stream zproto.RPCSync_PushUpdatesStreamServer) error {
	// TODO(@benqi): chan数量
	var update  chan *zproto.PushUpdatesData = make(chan *zproto.PushUpdatesData, 1000)

	var err error
	s.withWriteLock(func() {
		if _, ok := s.updates[auth.ServerId]; ok {
			err = errors.New("already connected")
			glog.Errorf("PushUpdatesStream - %s\n", err)
			return
		}
		s.updates[auth.ServerId] = update
	})

	if err != nil {
		return err
	}

	defer s.withWriteLock(func() { s.unsafeExpire(auth.ServerId) })

	for {
		select {
		case <-stream.Context().Done():
			err = stream.Context().Err()
			glog.Errorf("PushUpdatesStream - %s\n", err)
			return stream.Context().Err()
		case data := <-update:
			if err = stream.Send(data); err != nil {
				return err
			}

			glog.Infof("PushUpdatesStream: update: {%v}", data)
		}
	}
	return nil
}

func (s *SyncServiceImpl) DeliveryUpdates(ctx context.Context, deliver *zproto.DeliveryUpdatesToUsers) (reply *zproto.VoidRsp, err error) {
	glog.Infof("DeliveryUpdates: {%v}", deliver)

	statusList, err := model.GetOnlineStatusModel().GetOnlineByUserIdList(deliver.SendtoUserIdList)
	ss := make(map[int32][]*model.SessionStatus)
	for _, status := range statusList {
		if _, ok := ss[status.ServerId]; !ok {
			ss[status.ServerId] = []*model.SessionStatus{}
		}
		// 会不会出问题？？
		ss[status.ServerId] = append(ss[status.ServerId], status)
	}
	for k, ss3 := range ss {
		// glog.Infof("DeliveryUpdates: k: {%v}, v: {%v}", k, ss3)

		go s.withReadLock(func() {
			for _, ss4 := range ss3 {
				if ss4.NetlibSessionId != deliver.MyNetlibSessionId {
					update := &zproto.PushUpdatesData{}
					update.AuthKeyId = ss4.AuthKeyId
					update.SessionId = ss4.SessionId
					update.NetlibSessionId = ss4.NetlibSessionId
					// update.RawDataHeader = deliver.RawDataHeader
					update.RawData = deliver.RawData

					glog.Infof("DeliveryUpdates: update: {%v}", update)
					s.updates[k] <- update
				} else {
					glog.Infof("Not delivery me: {%v}", ss4)
				}
			}
		})
	}

	reply = &zproto.VoidRsp{}
	return
}
