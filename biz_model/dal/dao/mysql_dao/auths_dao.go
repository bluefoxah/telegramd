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

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}

// insert into auths(auth_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, connection_hash) values (:auth_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :connection_hash)
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) Insert(do *do.AuthsDO) (id int64, err error) {
	var query = "insert into auths(auth_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, connection_hash) values (:auth_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :connection_hash)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("AuthsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("AuthsDAO/LastInsertId error: ", err)
	}
	return
}

// select connection_hash from auths where auth_id = :auth_id
// TODO(@benqi): sqlmap
func (dao *AuthsDAO) SelectConnectionHashByAuthId(auth_id int64) (*do.AuthsDO, error) {
	var query = "select connection_hash from auths where auth_id = ?"
	rows, err := dao.db.Queryx(query, auth_id)

	if err != nil {
		glog.Error("AuthsDAO/SelectConnectionHashByAuthId error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.AuthsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("AuthsDAO/SelectConnectionHashByAuthId error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
