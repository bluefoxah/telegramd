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

type MessagesDAO struct {
	db *sqlx.DB
}

func NewMessagesDAO(db *sqlx.DB) *MessagesDAO {
	return &MessagesDAO{db}
}

func (dao *MessagesDAO) Insert(do *do.MessagesDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into messages(user_id, peer_type, peer_id, random_id, message, `date`, created_at) values (:user_id, :peer_type, :peer_id, :random_id, :message, :date, :created_at)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("MessagesDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("MessagesDAO/LastInsertId error: ", err)
	}
	return
}
