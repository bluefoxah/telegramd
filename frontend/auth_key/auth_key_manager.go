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
	"github.com/nebulaim/telegramd/base/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/golang/glog"
	"encoding/base64"
)

// Model Struct
type MasterKeys struct {
	AuthId   int64    	`orm:"pk"`
	Body     string
}

// "root:@/nebulaim?charset=utf8"
// 30
func NewAuthKeyCacheManager(dsn string) *AuthKeyCacheManager {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	// register model
	orm.RegisterModel(new(MasterKeys))
	// set default database
	err = orm.RegisterDataBase("default", "mysql", dsn, 30)
	if err != nil {
		panic(err)
	}

	return &AuthKeyCacheManager{ orm.NewOrm(), }
}

type AuthKeyCacheManager struct {
	ZOrm orm.Ormer
}

// TODO(@benqi): 暂时在这里操作数据库，需要改善的地方:
// 1. 如果数据库连接有问题，尝试存储到本地缓存
// 2. 所有需要读写数据库和缓存的地方，全部推给后端服务              2
// GetAuthKey(uint64) []byte
// PutAuthKey(uint64, []byte) error
func (s *AuthKeyCacheManager) GetAuthKey(keyID int64) (authKey []byte) {
	k := &MasterKeys{ AuthId: keyID, }
	err := s.ZOrm.Read(k)
	if err != nil {
		glog.Errorf("Read keyData error: %s\n", err)
		return nil
	}

	authKey, err = base64.RawStdEncoding.DecodeString(k.Body)
	return
}

func (s *AuthKeyCacheManager) PutAuthKey(keyID int64, key []byte) (err error) {
	k := &MasterKeys{ AuthId: keyID, }
	k.Body = base64.RawStdEncoding.EncodeToString(key)
	_, err = s.ZOrm.Insert(k)
	if err != nil {
		glog.Errorf("Write keyData error: %s\n", err)
	}

	return err
}

