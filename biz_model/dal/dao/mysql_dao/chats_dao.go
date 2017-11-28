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

type ChatsDAO struct {
	db *sqlx.DB
}

func NewChatsDAO(db *sqlx.DB) *ChatsDAO {
	return &ChatsDAO{db}
}

// insert into chats(creator_user_id, access_hash, participant_count, create_random_id, title, title_changer_user_id, title_change_random_id, title_changed_at, avatar_changer_user_id, avatar_change_random_id, avatar_changed_at, created_at) values (:creator_user_id, :access_hash, :participant_count, :create_random_id, :title, :title_changer_user_id, :title_change_random_id, :title_changed_at, :avatar_changer_user_id, :avatar_change_random_id, :avatar_changed_at, :created_at)
// TODO(@benqi): sqlmap
func (dao *ChatsDAO) Insert(do *dataobject.ChatsDO) int64 {
	var query = "insert into chats(creator_user_id, access_hash, participant_count, create_random_id, title, title_changer_user_id, title_change_random_id, title_changed_at, avatar_changer_user_id, avatar_change_random_id, avatar_changed_at, created_at) values (:creator_user_id, :access_hash, :participant_count, :create_random_id, :title, :title_changer_user_id, :title_change_random_id, :title_changed_at, :avatar_changer_user_id, :avatar_change_random_id, :avatar_changed_at, :created_at)"
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

// select id, participant_count, title, version from chats where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatsDAO) Select(id int32) *dataobject.ChatsDO {
	var query = "select id, participant_count, title, version from chats where id = ?"
	rows, err := dao.db.Queryx(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in Select(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ChatsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in Select(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// update chats set title = :title, title_changer_user_id = :title_changer_user_id, title_change_random_id = :title_change_random_id, title_changed_at = :title_changed_at where id = :id
// TODO(@benqi): sqlmap
func (dao *ChatsDAO) UpdateTitle(title string, title_changer_user_id int32, title_change_random_id int64, title_changed_at string, id int32) int64 {
	var query = "update chats set title = ?, title_changer_user_id = ?, title_change_random_id = ?, title_changed_at = ? where id = ?"
	r, err := dao.db.Exec(query, title, title_changer_user_id, title_change_random_id, title_changed_at, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTitle(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTitle(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select id, participant_count, title, version from chats where id in (:idList)
// TODO(@benqi): sqlmap
func (dao *ChatsDAO) SelectByIdList(idList []int32) []dataobject.ChatsDO {
	var q = "select id, participant_count, title, version from chats where id in (?)"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.ChatsDO
	for rows.Next() {
		v := dataobject.ChatsDO{}

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
