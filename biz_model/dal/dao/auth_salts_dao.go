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

package dao

import (
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	do "github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

type AuthSaltsDAO struct {
	db *sqlx.DB
}

func NewAuthSaltsDAO(db *sqlx.DB) *AuthSaltsDAO {
	return &AuthSaltsDAO{db}
}

func (dao *AuthSaltsDAO) Insert(do *do.AuthSaltsDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	var sql = "insert into auth_salts(auth_id, salt) values (:auth_id, :salt)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("AuthSaltsDAO/Insert error: ", err)
		return 0, nil
	}

	return r.LastInsertId()
}

func (dao *AuthSaltsDAO) SelectByAuthId(auth_id int64) (*do.AuthSaltsDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select auth_id, salt from auth_salts where auth_id = :auth_id"
	do := &do.AuthSaltsDO{AuthId: auth_id}
	r, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("AppsDAO/SelectById error: ", err)
		return nil, err
	}

	if r.Next() {
		err = r.StructScan(do)
		if err != nil {
			glog.Error("AppsDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
