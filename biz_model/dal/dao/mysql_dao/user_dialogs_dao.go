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

type UserDialogsDAO struct {
	db *sqlx.DB
}

func NewUserDialogsDAO(db *sqlx.DB) *UserDialogsDAO {
	return &UserDialogsDAO{db}
}

// insert into user_dialogs(user_id, peer_type, peer_id, created_at) values (:user_id, :peer_type, :peer_id, :created_at)
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) Insert(do *dataobject.UserDialogsDO) int64 {
	var query = "insert into user_dialogs(user_id, peer_type, peer_id, created_at) values (:user_id, :peer_type, :peer_id, :created_at)"
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

// select peer_type, peer_id from user_dialogs where user_id = :user_id and is_pinned = 1
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectPinnedDialogs(user_id int32) []dataobject.UserDialogsDO {
	var query = "select peer_type, peer_id from user_dialogs where user_id = ? and is_pinned = 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPinnedDialogs(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPinnedDialogs(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select id from user_dialogs where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) CheckExists(user_id int32, peer_type int32, peer_id int32) *dataobject.UserDialogsDO {
	var query = "select id from user_dialogs where user_id = ? and peer_type = ? and peer_id = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in CheckExists(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserDialogsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in CheckExists(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByUserID(user_id int32) []dataobject.UserDialogsDO {
	var query = "select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = ?"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogsByUserID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogsByUserID(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = :user_id and peer_type = :peer_type
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) SelectDialogsByPeerType(user_id int32, peer_type int32) []dataobject.UserDialogsDO {
	var query = "select peer_type, peer_id, is_pinned, top_message, read_inbox_max_id, read_outbox_max_id, unread_count, unread_mentions_count from user_dialogs where user_id = ? and peer_type = ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectDialogsByPeerType(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserDialogsDO
	for rows.Next() {
		v := dataobject.UserDialogsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectDialogsByPeerType(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// update user_dialogs set top_message = :top_message where id = :id
// TODO(@benqi): sqlmap
func (dao *UserDialogsDAO) UpdateTopMessage(top_message int32, id int32) int64 {
	var query = "update user_dialogs set top_message = ? where id = ?"
	r, err := dao.db.Exec(query, top_message, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTopMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTopMessage(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
