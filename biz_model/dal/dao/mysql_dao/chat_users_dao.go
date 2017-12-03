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

type ChatUsersDAO struct {
	db *sqlx.DB
}

func NewChatUsersDAO(db *sqlx.DB) *ChatUsersDAO {
	return &ChatUsersDAO{db}
}

// insert into chat_users(chat_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state, created_at) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, :joined_at, :state, :created_at)
// TODO(@benqi): sqlmap
func (dao *ChatUsersDAO) Insert(do *dataobject.ChatUsersDO) int64 {
	var query = "insert into chat_users(chat_id, user_id, participant_type, inviter_user_id, invited_at, joined_at, state, created_at) values (:chat_id, :user_id, :participant_type, :inviter_user_id, :invited_at, :joined_at, :state, :created_at)"
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

// select id, chat_id, user_id, participant_type, inviter_user_id, invited_at, joined_at from chat_users where chat_id = :chat_id
// TODO(@benqi): sqlmap
func (dao *ChatUsersDAO) SelectByChatId(chat_id int32) []dataobject.ChatUsersDO {
	var query = "select id, chat_id, user_id, participant_type, inviter_user_id, invited_at, joined_at from chat_users where chat_id = ?"
	rows, err := dao.db.Queryx(query, chat_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByChatId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChatUsersDO
	for rows.Next() {
		v := dataobject.ChatUsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByChatId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// delete from chat_users where chat_id = :chat_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ChatUsersDAO) DeleteChatUser(chat_id int32, user_id int32) int64 {
	var query = "delete from chat_users where chat_id = ? and user_id = ?"
	r, err := dao.db.Exec(query, chat_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in DeleteChatUser(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in DeleteChatUser(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return rows
}
