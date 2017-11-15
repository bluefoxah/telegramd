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

package mysql_dao

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

// insert into auth_salts(auth_id, salt) values (:auth_id, :salt)
// TODO(@benqi): sqlmap
func (dao *AuthSaltsDAO) Insert(do *do.AuthSaltsDO) (id int64, err error) {
	var query = "insert into auth_salts(auth_id, salt) values (:auth_id, :salt)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("AuthSaltsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("AuthSaltsDAO/LastInsertId error: ", err)
	}
	return
}

// select auth_id, salt from auth_salts where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthSaltsDAO) SelectByAuthId(auth_id int64) (*do.AuthSaltsDO, error) {
	var query = "select auth_id, salt from auth_salts where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		glog.Error("AuthSaltsDAO/SelectByAuthId error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.AuthSaltsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("AuthSaltsDAO/SelectByAuthId error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
