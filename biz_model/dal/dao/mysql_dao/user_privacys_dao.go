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

type UserPrivacysDAO struct {
	db *sqlx.DB
}

func NewUserPrivacysDAO(db *sqlx.DB) *UserPrivacysDAO {
	return &UserPrivacysDAO{db}
}

// insert into user_privacys(user_id, password, recovery_mail, status_timestamp, chat_invite, phone_call, ttl, ttl_created_at, created_at) values (:user_id, :password, :recovery_mail, :status_timestamp, :chat_invite, :phone_call, :ttl, :ttl_created_at, :created_at)
// TODO(@benqi): sqlmap
func (dao *UserPrivacysDAO) Insert(do *dataobject.UserPrivacysDO) int64 {
	var query = "insert into user_privacys(user_id, password, recovery_mail, status_timestamp, chat_invite, phone_call, ttl, ttl_created_at, created_at) values (:user_id, :password, :recovery_mail, :status_timestamp, :chat_invite, :phone_call, :ttl, :ttl_created_at, :created_at)"
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

// update user_privacys set ttl = :ttl, ttl_created_at = :ttl_created_at where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPrivacysDAO) UpdateTTL(ttl int32, ttl_created_at int32, user_id int32) int64 {
	var query = "update user_privacys set ttl = ?, ttl_created_at = ? where user_id = ?"
	r, err := dao.db.Exec(query, ttl, ttl_created_at, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// select ttl, ttl_created_at from user_privacys where user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *UserPrivacysDAO) SelectTTL(user_id int32) *dataobject.UserPrivacysDO {
	var query = "select ttl, ttl_created_at from user_privacys where user_id = ?"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectTTL(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UserPrivacysDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectTTL(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
