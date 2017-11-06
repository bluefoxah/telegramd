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

type UserContactsDAO struct {
	db *sqlx.DB
}

func NewUserContactsDAO(db *sqlx.DB) *UserContactsDAO {
	return &UserContactsDAO{db}
}

func (dao *UserContactsDAO) Insert(do *do.UserContactsDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into user_contacts(owner_user_id, contact_user_id) values (:owner_user_id, :contact_user_id)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("UserContactsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("UserContactsDAO/LastInsertId error: ", err)
	}
	return
}

func (dao *UserContactsDAO) SelectUserContacts(owner_user_id int32) ([]do.UserContactsDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0"
	do2 := &do.UserContactsDO{OwnerUserId: owner_user_id}
	rows, err := dao.db.NamedQuery(sql, do2)
	if err != nil {
		glog.Errorf("UserContactsDAO/SelectUserContacts error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UserContactsDO
	for rows.Next() {
		v := do.UserContactsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UserContactsDAO/SelectUserContacts error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
