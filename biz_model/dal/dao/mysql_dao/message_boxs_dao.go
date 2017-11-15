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

type MessageBoxsDAO struct {
	db *sqlx.DB
}

func NewMessageBoxsDAO(db *sqlx.DB) *MessageBoxsDAO {
	return &MessageBoxsDAO{db}
}

// insert into message_boxs(user_id, message_box_type, message_id, pts, created_at) values (:user_id, :message_box_type, :message_id, :pts, :created_at)
// TODO(@benqi): sqlmap
func (dao *MessageBoxsDAO) Insert(do *do.MessageBoxsDO) (id int64, err error) {
	var query = "insert into message_boxs(user_id, message_box_type, message_id, pts, created_at) values (:user_id, :message_box_type, :message_id, :pts, :created_at)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("MessageBoxsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("MessageBoxsDAO/LastInsertId error: ", err)
	}
	return
}
