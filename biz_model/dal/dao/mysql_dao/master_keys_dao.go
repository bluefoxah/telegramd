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

type MasterKeysDAO struct {
	db *sqlx.DB
}

func NewMasterKeysDAO(db *sqlx.DB) *MasterKeysDAO {
	return &MasterKeysDAO{db}
}

func (dao *MasterKeysDAO) Insert(do *do.MasterKeysDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into master_keys(auth_id, body) values (:auth_id, :body)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("MasterKeysDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("MasterKeysDAO/LastInsertId error: ", err)
	}
	return
}

func (dao *MasterKeysDAO) SelectByAuthId(auth_id int64) (*do.MasterKeysDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select body from master_keys where auth_id = :auth_id"
	do := &do.MasterKeysDO{AuthId: auth_id}
	rows, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("MasterKeysDAO/SelectById error: ", err)
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("MasterKeysDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
