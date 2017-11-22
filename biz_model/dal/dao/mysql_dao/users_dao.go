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

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

// insert into users(first_name, last_name, access_hash, username, phone, country_code, bio, about, is_bot) values (:first_name, :last_name, :access_hash, :username, :phone, :country_code, :bio, :about, :is_bot)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) Insert(do *dataobject.UsersDO) int64 {
	var query = "insert into users(first_name, last_name, access_hash, username, phone, country_code, bio, about, is_bot) values (:first_name, :last_name, :access_hash, :username, :phone, :country_code, :bio, :about, :is_bot)"
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

// select id, access_hash, first_name, last_name, username from users where phone = :phone limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByPhoneNumber(phone string) *dataobject.UsersDO {
	var query = "select id, access_hash, first_name, last_name, username from users where phone = ? limit 1"
	rows, err := dao.db.Queryx(query, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPhoneNumber(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPhoneNumber(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select id, access_hash, first_name, last_name, username from users where id = :id limit 1
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectById(id int32) *dataobject.UsersDO {
	var query = "select id, access_hash, first_name, last_name, username from users where id = ? limit 1"
	rows, err := dao.db.Queryx(query, id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectById(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectById(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select id, access_hash, first_name, last_name, username from users where id in (:id_list)
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectUsersByIdList(id_list []int32) []dataobject.UsersDO {
	var q = "select id, access_hash, first_name, last_name, username from users where id in (?)"
	query, a, err := sqlx.In(q, id_list)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectUsersByIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectUsersByIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select id, access_hash, first_name, last_name, username, phone from users where username = :username or first_name = :first_name or last_name = :last_name or phone = :phone limit 20
// TODO(@benqi): sqlmap
func (dao *UsersDAO) SelectByQueryString(username string, first_name string, last_name string, phone string) []dataobject.UsersDO {
	var query = "select id, access_hash, first_name, last_name, username, phone from users where username = ? or first_name = ? or last_name = ? or phone = ? limit 20"
	rows, err := dao.db.Queryx(query, username, first_name, last_name, phone)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByQueryString(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByQueryString(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// update users set deleted = 1, deleted_reason = :deleted_reason, deleted_at = :deleted_at where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) Delete(deleted_reason string, deleted_at string, id int32) int64 {
	var query = "update users set deleted = 1, deleted_reason = ?, deleted_at = ? where id = ?"
	r, err := dao.db.Exec(query, deleted_reason, deleted_at, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in Delete(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in Delete(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update users set username = :username where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateUsername(username string, id int32) int64 {
	var query = "update users set username = ? where id = ?"
	r, err := dao.db.Exec(query, username, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateUsername(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateUsername(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}

// update users set first_name = :first_name, last_name = :last_name, about = :about where id = :id
// TODO(@benqi): sqlmap
func (dao *UsersDAO) UpdateProfile(first_name string, last_name string, about string, id int32) int64 {
	var query = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
	r, err := dao.db.Exec(query, first_name, last_name, about, id)

	if err != nil {
		errDesc := fmt.Sprintf("Exec in UpdateProfile(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	rows, err := r.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("RowsAffected in UpdateProfile(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	return rows
}
