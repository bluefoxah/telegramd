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
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/nebulaim/telegramd/mtproto"
)

type MessagesDAO struct {
	db *sqlx.DB
}

func NewMessagesDAO(db *sqlx.DB) *MessagesDAO {
	return &MessagesDAO{db}
}

// insert into messages(sender_user_id, peer_type, peer_id, random_id, message_type, message_data, date2) values (:sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :date2)
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) Insert(do *dataobject.MessagesDO) int64 {
	var query = "insert into messages(sender_user_id, peer_type, peer_id, random_id, message_type, message_data, date2) values (:sender_user_id, :peer_type, :peer_id, :random_id, :message_type, :message_data, :date2)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in Insert(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in Insert(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}

// select id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, date2 from messages where id in (:idList) order by id desc
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) SelectByIdList(idList []int32) []dataobject.MessagesDO {
	var q = "select id, sender_user_id, peer_type, peer_id, random_id, message_type, message_data, date2 from messages where id in (?) order by id desc"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessagesDO
	for rows.Next() {
		v := dataobject.MessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}
