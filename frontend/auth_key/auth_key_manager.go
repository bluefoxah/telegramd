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

package auth_key

import (
	"github.com/golang/glog"
	"encoding/base64"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/cosiner/gohper/errors"
)

// "root:@/nebulaim?charset=utf8"
// 30
func NewAuthKeyCacheManager() *AuthKeyCacheManager {
	return &AuthKeyCacheManager{}
}

type AuthKeyCacheManager struct {
}

// TODO(@benqi): 暂时在这里操作数据库，需要改善的地方:
// 1. 如果数据库连接有问题，尝试存储到本地缓存
// 2. 所有需要读写数据库和缓存的地方，全部推给后端服务
// GetAuthKey(uint64) []byte
// PutAuthKey(uint64, []byte) error
func (s *AuthKeyCacheManager) GetAuthKey(keyID int64) (authKey []byte) {
	defer func() {
		if r := recover(); r != nil {
			authKey = nil
		}
	}()

	do := dao.GetAuthKeysDAO(dao.DB_SLAVE).SelectByAuthId(keyID)
	if do == nil {
		glog.Errorf("Read keyData error: not find keyId\n")
		return nil
	}
	authKey, _ = base64.RawStdEncoding.DecodeString(do.Body)
	return
}

func (s *AuthKeyCacheManager) PutAuthKey(keyID int64, key []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r)
		}
	}()

	do := &dataobject.AuthKeysDO{ AuthId: keyID}
	do.Body = base64.RawStdEncoding.EncodeToString(key)
	dao.GetAuthKeysDAO(dao.DB_MASTER).Insert(do)
	return
}

