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

func (dao *MessagesDAO) SelectByIdList(idList []int32) ([]do.MessagesDO, error) {
	// TODO(@benqi): sqlmap
	params := make(map[string]interface{})
	// params["idList"] = base.JoinInt32List(idList, ",")

	var sql = fmt.Sprintf("select id, user_id, peer_type, peer_id, random_id, message, date from messages where id in (%s)", base.JoinInt32List(idList, ","))
	rows, err := dao.db.NamedQuery(sql, params)
	if err != nil {
		glog.Errorf("MessagesDAO/SelectByIdList error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.MessagesDO
	for rows.Next() {
		v := do.MessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("MessagesDAO/SelectByIdList error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
