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

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

func (dao *UsersDAO) Insert(do *do.UsersDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into users(phone) values (:phone)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("UsersDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("UsersDAO/LastInsertId error: ", err)
	}
	return
}

func (dao *UsersDAO) SelectByPhoneNumber(phone string) (*do.UsersDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select id, access_hash, first_name, last_name, username from users where phone = :phone limit 1"
	do := &do.UsersDO{Phone: phone}
	rows, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("UsersDAO/SelectByPhoneNumber error: ", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("UsersDAO/SelectByPhoneNumber error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

func (dao *UsersDAO) SelectById(id int32) (*do.UsersDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select id, access_hash, first_name, last_name, username from users where id = :id limit 1"
	do := &do.UsersDO{Id: id}
	rows, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("UsersDAO/SelectById error: ", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("UsersDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

func (dao *UsersDAO) SelectByQueryString(username string, first_name string, last_name string, phone string) ([]do.UsersDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select id, access_hash, first_name, last_name, username, phone from users where username = :username or first_name = :first_name or last_name = :last_name or phone = :phone limit 20"
	do2 := &do.UsersDO{Username: username, FirstName: first_name, LastName: last_name, Phone: phone}
	rows, err := dao.db.NamedQuery(sql, do2)
	if err != nil {
		glog.Errorf("UsersDAO/SelectByQueryString error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UsersDO
	for rows.Next() {
		v := do.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UsersDAO/SelectByQueryString error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
