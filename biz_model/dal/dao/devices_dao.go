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

type DevicesDAO struct {
	db *sqlx.DB
}

func NewDevicesDAO(db *sqlx.DB) *DevicesDAO {
	return &DevicesDAO{db}
}

func (dao *DevicesDAO) Insert(do *do.DevicesDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into devices(user_id, token_type, token, state) values (:user_id, :token_type, :token, :state)"
	r, err := dao.db.NamedExec(sql, do)
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

func (dao *DevicesDAO) SelectIdByAuthId(auth_id int64, token_type int8, token string) (*do.DevicesDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select id from devices where user_id in (select id from auth_users where auth_id = :auth_id limit 1) and token_type = :token_type and token = :token limit 1"
	do := &do.DevicesDO{AuthId: auth_id, TokenType: token_type, Token: token}
	r, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("DevicesDAO/SelectById error: ", err)
		return nil, err
	}

	if r.Next() {
		err = r.StructScan(do)
		if err != nil {
			glog.Error("DevicesDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

func (dao *DevicesDAO) UpdateStateById(state int8, id int32) (rows int64, err error) {
	// TODO(@benqi): sqlmap
	rows = 0

	var sql = "update devices set state = :state where id = :id"
	do := &do.DevicesDO{State: state, Id: id}
	r, err := dao.db.NamedExec(sql, do)
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

func (dao *DevicesDAO) UpdateStateByAuthId(id int32) (rows int64, err error) {
	// TODO(@benqi): sqlmap
	rows = 0

	var sql = "update devices set state = 1 where id = :id"
	do := &do.DevicesDO{Id: id}
	r, err := dao.db.NamedExec(sql, do)
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
