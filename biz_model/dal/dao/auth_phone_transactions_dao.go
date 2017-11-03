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

package dao

import (
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	do "github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

type AuthPhoneTransactionsDAO struct {
	db *sqlx.DB
}

func NewAuthPhoneTransactionsDAO(db *sqlx.DB) *AuthPhoneTransactionsDAO {
	return &AuthPhoneTransactionsDAO{db}
}

func (dao *AuthPhoneTransactionsDAO) Insert(do *do.AuthPhoneTransactionsDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	var sql = "insert into auth_phone_transactions(transaction_hash, api_id, api_hash, phone_number, code, created_at) values (:transaction_hash, :api_id, :api_hash, :phone_number, :code, :created_at)"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/Insert error: ", err)
		return 0, nil
	}

	return r.LastInsertId()
}

func (dao *AuthPhoneTransactionsDAO) SelectByPhoneAndApiIdAndHash(phone_number string, api_id int32, api_hash string) (*do.AuthPhoneTransactionsDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select transaction_hash from auth_phone_transactions where phone_number = :phone_number and api_id = :api_id and api_hash = :api_hash"
	do := &do.AuthPhoneTransactionsDO{PhoneNumber: phone_number, ApiId: api_id, ApiHash: api_hash}
	r, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/SelectById error: ", err)
		return nil, err
	}

	if r.Next() {
		err = r.StructScan(do)
		if err != nil {
			glog.Error("AuthPhoneTransactionsDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

func (dao *AuthPhoneTransactionsDAO) SelectByPhoneCode(transaction_hash string, code string, phone_number string) (*do.AuthPhoneTransactionsDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select id from auth_phone_transactions where transaction_hash = :transaction_hash and code = :code and phone_number = :phone_number"
	do := &do.AuthPhoneTransactionsDO{TransactionHash: transaction_hash, Code: code, PhoneNumber: phone_number}
	r, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/SelectById error: ", err)
		return nil, err
	}

	if r.Next() {
		err = r.StructScan(do)
		if err != nil {
			glog.Error("AuthPhoneTransactionsDAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
