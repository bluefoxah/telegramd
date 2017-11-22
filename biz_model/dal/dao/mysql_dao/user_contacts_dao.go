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

type UserContactsDAO struct {
	db *sqlx.DB
}

func NewUserContactsDAO(db *sqlx.DB) *UserContactsDAO {
	return &UserContactsDAO{db}
}

// insert into user_contacts(owner_user_id, contact_user_id, created_at) values (:owner_user_id, :contact_user_id, :created_at)
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) Insert(do *dataobject.UserContactsDO) int64 {
	var query = "insert into user_contacts(owner_user_id, contact_user_id, created_at) values (:owner_user_id, :contact_user_id, :created_at)"
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

// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = 0
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) SelectUserContacts(owner_user_id int32) []dataobject.UserContactsDO {
	var query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = 0"
	rows, err := dao.db.Queryx(query, owner_user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectUserContacts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UserContactsDO
	for rows.Next() {
		v := dataobject.UserContactsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectUserContacts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// update user_contacts set is_blocked = :is_blocked where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) UpdateBlock(is_blocked int8, owner_user_id int32, contact_user_id int32) int64 {
	var query = "update user_contacts set is_blocked = ? where owner_user_id = ? and contact_user_id = ?"
	r, err := dao.db.Exec(query, is_blocked, owner_user_id, contact_user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateBlock(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateBlock(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update user_contacts set is_deleted = 1 where owner_user_id = :owner_user_id and contact_user_id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *UserContactsDAO) DeleteContacts(owner_user_id int32, id_list []int32) int64 {
	var q = "update user_contacts set is_deleted = 1 where owner_user_id = ? and contact_user_id in (?)"
	query, a, err := sqlx.In(q, owner_user_id, id_list)
	r, err := dao.db.Exec(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in DeleteContacts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in DeleteContacts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
