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

type AppsDAO struct {
	db *sqlx.DB
}

func NewAppsDAO(db *sqlx.DB) *AppsDAO {
	return &AppsDAO{db}
}

func (dao *AppsDAO) Insert(do *do.AppsDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into apps(api_id, api_hash, title, short_name) values (:api_id, :api_hash, :title, :short_name)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("AppsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("AppsDAO/LastInsertId error: ", err)
	}
	return
}

func (dao *AppsDAO) SelectById(id int32) (*do.AppsDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select id, api_id, api_hash, title, short_name from apps where id = :id"
	do := &do.AppsDO{Id: id}
	rows, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("AppsDAO/SelectById error: ", err)
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("AppsDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
