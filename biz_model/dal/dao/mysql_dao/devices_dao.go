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

type DevicesDAO struct {
	db *sqlx.DB
}

func NewDevicesDAO(db *sqlx.DB) *DevicesDAO {
	return &DevicesDAO{db}
}

// insert into devices(user_id, token_type, token, state) values (:user_id, :token_type, :token, :state)
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) Insert(do *do.DevicesDO) (id int64, err error) {
	var query = "insert into devices(user_id, token_type, token, state) values (:user_id, :token_type, :token, :state)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("DevicesDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("DevicesDAO/LastInsertId error: ", err)
	}
	return
}

// select id from devices where user_id in (select id from auth_users where auth_id = :auth_id) and token_type = :token_type and token = :token limit 1
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) SelectIdByAuthId(auth_id int64, token_type int8, token string) (*do.DevicesDO, error) {
	var query = "select id from devices where user_id in (select id from auth_users where auth_id = ?) and token_type = ? and token = ? limit 1"
	rows, err := dao.db.Queryx(query, auth_id, token_type, token)

	if err != nil {
		glog.Error("DevicesDAO/SelectIdByAuthId error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.DevicesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("DevicesDAO/SelectIdByAuthId error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

// update devices set state = :state where id = :id
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateById(state int8, id int32) (rows int64, err error) {
	var query = "update devices set state = ? where id = ?"
	r, err := dao.db.Exec(query, state, id)

	if err != nil {
		glog.Error("DevicesDAO/UpdateStateById error: ", err)
		return
	}

	rows, err = r.RowsAffected()
	if err != nil {
		glog.Error("DevicesDAO/RowsAffected error: ", err)
	}
	return
}

// update devices set state = 1 where id = :id
// TODO(@benqi): sqlmap
func (dao *DevicesDAO) UpdateStateByAuthId(id int32) (rows int64, err error) {
	var query = "update devices set state = 1 where id = ?"
	r, err := dao.db.Exec(query, id)

	if err != nil {
		glog.Error("DevicesDAO/UpdateStateByAuthId error: ", err)
		return
	}

	rows, err = r.RowsAffected()
	if err != nil {
		glog.Error("DevicesDAO/RowsAffected error: ", err)
	}
	return
}
