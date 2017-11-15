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
	"github.com/nebulaim/telegramd/base/base"
	do "github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"fmt"
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
	params := make(map[string]interface{})
	params["phone"] = phone

	var sql = "select id, access_hash, first_name, last_name, username from users where phone = :phone limit 1"
	rows, err := dao.db.NamedQuery(sql, params)
	if err != nil {
		glog.Error("UsersDAO/SelectByPhoneNumber error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.UsersDO{}
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
	params := make(map[string]interface{})
	params["id"] = id

	var sql = "select id, access_hash, first_name, last_name, username from users where id = :id limit 1"
	rows, err := dao.db.NamedQuery(sql, params)
	if err != nil {
		glog.Error("UsersDAO/SelectById error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.UsersDO{}
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

func (dao *UsersDAO) SelectUsersByIdList(id_list []int32) ([]do.UsersDO, error) {
	// TODO(@benqi): sqlmap
	params := make(map[string]interface{})
	//params["id_list"] = base.JoinInt32List(id_list, ",")

	var sql = fmt.Sprintf("select id, access_hash, first_name, last_name, username from users where id in (%s)", base.JoinInt32List(id_list, ","))
	rows, err := dao.db.NamedQuery(sql, params)
	if err != nil {
		glog.Errorf("UsersDAO/SelectUsersByIdList error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UsersDO
	for rows.Next() {
		v := do.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UsersDAO/SelectUsersByIdList error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}

func (dao *UsersDAO) SelectByQueryString(first_name string, last_name string, phone string, username string) ([]do.UsersDO, error) {
	// TODO(@benqi): sqlmap
	params := make(map[string]interface{})
	params["first_name"] = first_name
	params["last_name"] = last_name
	params["phone"] = phone
	params["username"] = username

	var sql = "select id, access_hash, first_name, last_name, username, phone from users where username = :username or first_name = :first_name or last_name = :last_name or phone = :phone limit 20"
	rows, err := dao.db.NamedQuery(sql, params)
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
