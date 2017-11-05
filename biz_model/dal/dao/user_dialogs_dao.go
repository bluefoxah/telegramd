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

type UserDialogsDAO struct {
	db *sqlx.DB
}

func NewUserDialogsDAO(db *sqlx.DB) *UserDialogsDAO {
	return &UserDialogsDAO{db}
}

func (dao *UserDialogsDAO) Insert(do *do.UserDialogsDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into user_dialogs(user_id, peer_type, peer_id, created_at) values (:user_id, peer_type, peer_id, created_at)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("UserDialogsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("UserDialogsDAO/LastInsertId error: ", err)
	}
	return
}

func (dao *UserDialogsDAO) SelectPinnedDialogs(user_id int32) ([]do.UserDialogsDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select peer_type, peer_id from user_dialogs where user_id = :user_id and is_pinned = 1"
	do2 := &do.UserDialogsDO{UserId: user_id}
	glog.Info("do2: ", do2)
	rows, err := dao.db.NamedQuery(sql, do2)

	if err != nil {
		glog.Errorf("UserDialogsDAO/SelectPinnedDialogs error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.UserDialogsDO
	for rows.Next() {
		v := do.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("UserDialogsDAO/SelectPinnedDialogs error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
