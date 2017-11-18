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

type AuthPhoneTransactionsDAO struct {
	db *sqlx.DB
}

func NewAuthPhoneTransactionsDAO(db *sqlx.DB) *AuthPhoneTransactionsDAO {
	return &AuthPhoneTransactionsDAO{db}
}

// insert into auth_phone_transactions(transaction_hash, api_id, api_hash, phone_number, code, created_at) values (:transaction_hash, :api_id, :api_hash, :phone_number, :code, :created_at)
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) Insert(do *dataobject.AuthPhoneTransactionsDO) int64 {
	var query = "insert into auth_phone_transactions(transaction_hash, api_id, api_hash, phone_number, code, created_at) values (:transaction_hash, :api_id, :api_hash, :phone_number, :code, :created_at)"
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

// select transaction_hash from auth_phone_transactions where phone_number = :phone_number and api_id = :api_id and api_hash = :api_hash
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) SelectByPhoneAndApiIdAndHash(phone_number string, api_id int32, api_hash string) *dataobject.AuthPhoneTransactionsDO {
	var query = "select transaction_hash from auth_phone_transactions where phone_number = ? and api_id = ? and api_hash = ?"
	rows, err := dao.db.Queryx(query, phone_number, api_id, api_hash)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPhoneAndApiIdAndHash(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthPhoneTransactionsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPhoneAndApiIdAndHash(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select id from auth_phone_transactions where transaction_hash = :transaction_hash and code = :code and phone_number = :phone_number
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) SelectByPhoneCode(transaction_hash string, code string, phone_number string) *dataobject.AuthPhoneTransactionsDO {
	var query = "select id from auth_phone_transactions where transaction_hash = ? and code = ? and phone_number = ?"
	rows, err := dao.db.Queryx(query, transaction_hash, code, phone_number)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByPhoneCode(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.AuthPhoneTransactionsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByPhoneCode(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
