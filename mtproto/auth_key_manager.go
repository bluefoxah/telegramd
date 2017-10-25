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
	"sync"
	"io/ioutil"
	"encoding/binary"
	"github.com/golang/glog"
)

const (
	AUTH_KEY_DB = "/Users/benqi/Github/nebula-im/out/imengine/bin/Debug/auth_key.dat"
	AUTH_KEY_SIZE = 264
)

type cacheAuthKeys struct {
	keys sync.Map
	storeMutex sync.Mutex
	once sync.Once
}

var keysManager = &cacheAuthKeys{}

func initializeOnce() {
	db, err := ioutil.ReadFile(AUTH_KEY_DB);
	if err != nil {
		glog.Errorf("Read file error: ", AUTH_KEY_DB)
		return
	}

	size := len(db) / AUTH_KEY_SIZE

	for i := 0; i < size; i++ {
		kid := binary.BigEndian.Uint64(db[i*264:])
		// key := db[8:]
		keysManager.keys.Store(kid, db[i*264+8:i*264+8+256])
		glog.Info("Load authKey: kid=", kid)
	}
}

func FindAuthKey(keyID uint64) []byte {
	// 初始化
	keysManager.once.Do(initializeOnce);

	k, b := keysManager.keys.Load(keyID)
	if b == false {
		return nil
	}

	return k.([]byte)
}

func StoreAuthKey(keyID uint64, key []byte) {
	// 初始化
	keysManager.once.Do(initializeOnce);

	// 存储key
	keysManager.keys.Store(keyID, key)

	var keys = []byte{}
	keysManager.keys.Range(func(key, value interface{}) bool {
		data := make([]byte, AUTH_KEY_SIZE)
		binary.BigEndian.PutUint64(data, key.(uint64))
		copy(data[8:], value.([]byte))
		keys = append(keys, data...)
		return true
	})

	keysManager.storeMutex.Lock()
	defer keysManager.storeMutex.Unlock()

	ioutil.WriteFile(AUTH_KEY_DB, keys, 0666)
}
