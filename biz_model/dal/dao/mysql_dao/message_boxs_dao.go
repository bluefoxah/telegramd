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

type MessageBoxsDAO struct {
	db *sqlx.DB
}

func NewMessageBoxsDAO(db *sqlx.DB) *MessageBoxsDAO {
	return &MessageBoxsDAO{db}
}

// insert into message_boxs(user_id, message_box_type, message_id, pts, created_at) values (:user_id, :message_box_type, :message_id, :pts, :created_at)
// TODO(@benqi): sqlmap
func (dao *MessageBoxsDAO) Insert(do *dataobject.MessageBoxsDO) int64 {
	var query = "insert into message_boxs(user_id, message_box_type, message_id, pts, created_at) values (:user_id, :message_box_type, :message_id, :pts, :created_at)"
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

// select pts from message_boxs where user_id = :user_id order by pts desc limit 1
// TODO(@benqi): sqlmap
func (dao *MessageBoxsDAO) SelectLastPts(user_id int32) *dataobject.MessageBoxsDO {
	var query = "select pts from message_boxs where user_id = ? order by pts desc limit 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastPts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select pts from message_boxs where user_id = :user_id and message_id > :message_id order by pts asc
// TODO(@benqi): sqlmap
func (dao *MessageBoxsDAO) SelectPtsByGTMessageID(user_id int32, message_id int32) []dataobject.MessageBoxsDO {
	var query = "select pts from message_boxs where user_id = ? and message_id > ? order by pts asc"
	rows, err := dao.db.Queryx(query, user_id, message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPtsByGTMessageID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxsDO
	for rows.Next() {
		v := dataobject.MessageBoxsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPtsByGTMessageID(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}
