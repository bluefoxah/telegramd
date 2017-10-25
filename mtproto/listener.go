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
	"net"
	"os"
	"sync"
	"io"
	"github.com/golang/glog"
	"sync/atomic"
	"encoding/hex"
	"encoding/binary"
	"github.com/nebulaim/telegramd/base/crypto"
)

const (
	FIRST1_INT32 = 0x44414548
	FIRST2_INT32 = 0x54534f50
	FIRST3_INT32 = 0x20544547
	FIRST4_INT32 = 0x20544547
	FIRST5_INT32 = 0xeeeeeeee
	SECOND_INT32 = 0x00000000
	FIRST_BYTE 	 = 0xef
)

var _ net.Listener = &Listener{}

type Listener struct {
	base         net.Listener
	acceptChan   chan net.Conn
	closed       bool
	closeOnce    sync.Once
	closeChan    chan struct{}
	atomicConnID uint64
	connsMutex   sync.Mutex
	conns        map[uint64]*MTProtoConn
}

func Listen(listenFunc func() (net.Listener, error)) (*Listener, error) {
	listener, err := listenFunc()
	if err != nil {
		return nil, err
	}
	l := &Listener{
		base:       listener,
		closeChan:  make(chan struct{}),
		acceptChan: make(chan net.Conn, 1000),
		conns:      make(map[uint64]*MTProtoConn),
	}
	go l.acceptLoop()
	return l, nil
}

func (l *Listener) Addr() net.Addr {
	return l.base.Addr()
}

func (l *Listener) Close() error {
	l.closeOnce.Do(func() {
		l.closed = true
		close(l.closeChan)
	})
	return l.base.Close()
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.acceptChan:
		return conn, nil
	case <-l.closeChan:
	}
	return nil, os.ErrInvalid
}

func (l *Listener) acceptLoop() {
	for {
		conn, err := l.base.Accept()
		if err != nil {
			if !l.closed {
				// l.trace("accept failed: %v", err)
			}
			break
		}
		go l.handshake(conn)
	}
}

func (l *Listener) handshake(conn net.Conn) {
	glog.Info("New connection by ", conn.RemoteAddr())
	var buf [64]byte
	if n, err := io.ReadFull(conn, buf[:]); err != nil {
		glog.Error("Receive faild: ", err, "len: ", n)
		conn.Close()
		return
	}

	glog.Info("Read first 64 bytes: ", hex.EncodeToString(buf[:]))

	// 检查val和val2
	first := binary.BigEndian.Uint32(buf[:4])
	second := binary.BigEndian.Uint32(buf[4:8])
	if buf[0]  == FIRST_BYTE   ||
		first  == FIRST1_INT32 ||
		first  == FIRST2_INT32 ||
		first  == FIRST3_INT32 ||
		first  == FIRST4_INT32 ||
		first  == FIRST5_INT32 ||
		second == SECOND_INT32 {

		glog.Errorf("Invalid key: ", hex.EncodeToString(buf[:8]))
		conn.Close()
		return
	}

	var tmp [64]byte
	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = buf[55 - i]
	}

	var connID = atomic.AddUint64(&l.atomicConnID, 1)

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	if err != nil {
		glog.Error("NewAesCTR128Encrypt error: %s", err)
		conn.Close()
		return
	}

	d, err := crypto.NewAesCTR128Encrypt(buf[8:40], buf[40:56])
	if err != nil {
		glog.Error("NewAesCTR128Encrypt error: %s", err)
		conn.Close()
		return
	}

	d.Encrypt(buf[:])
	if buf[56] != 0xef && buf[57] != 0xef && buf[58] != 0xef && buf[59] != 0xef {
		glog.Error("Invalid buf[56:60] error: ", hex.EncodeToString(buf[56:64]))
		conn.Close()
		return
	}

	mtprotoConn := newMTProtoConn(conn, connID, e, d)
	mtprotoConn.listener = l
	l.putConn(connID, mtprotoConn)

	glog.Info("Create AesCTR128 key sucessful in connID: ", connID, ", by ", conn.RemoteAddr())
	select {
		case l.acceptChan <- mtprotoConn:
		case <-l.closeChan:
	}
}

func (l *Listener) getConn(id uint64) (*MTProtoConn, bool) {
	l.connsMutex.Lock()
	defer l.connsMutex.Unlock()
	conn, exists := l.conns[id]
	return conn, exists
}

func (l *Listener) putConn(id uint64, conn *MTProtoConn) {
	l.connsMutex.Lock()
	defer l.connsMutex.Unlock()
	l.conns[id] = conn
}

func (l *Listener) delConn(id uint64) {
	l.connsMutex.Lock()
	defer l.connsMutex.Unlock()
	if _, exists := l.conns[id]; exists {
		delete(l.conns, id)
	}
}
