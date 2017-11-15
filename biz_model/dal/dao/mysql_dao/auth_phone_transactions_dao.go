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

// insert into auth_phone_transactions(transaction_hash, api_id, api_hash, phone_number, code, created_at) values (:transaction_hash, :api_id, :api_hash, :phone_number, :code, :created_at)
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) Insert(do *do.AuthPhoneTransactionsDO) (id int64, err error) {
	var query = "insert into auth_phone_transactions(transaction_hash, api_id, api_hash, phone_number, code, created_at) values (:transaction_hash, :api_id, :api_hash, :phone_number, :code, :created_at)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/LastInsertId error: ", err)
	}
	return
}

// select transaction_hash from auth_phone_transactions where phone_number = :phone_number and api_id = :api_id and api_hash = :api_hash
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) SelectByPhoneAndApiIdAndHash(phone_number string, api_id int32, api_hash string) (*do.AuthPhoneTransactionsDO, error) {
	var query = "select transaction_hash from auth_phone_transactions where phone_number = ? and api_id = ? and api_hash = ?"
	rows, err := dao.db.Queryx(query, phone_number, api_id, api_hash)

	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/SelectByPhoneAndApiIdAndHash error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.AuthPhoneTransactionsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("AuthPhoneTransactionsDAO/SelectByPhoneAndApiIdAndHash error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}

// select id from auth_phone_transactions where transaction_hash = :transaction_hash and code = :code and phone_number = :phone_number
// TODO(@benqi): sqlmap
func (dao *AuthPhoneTransactionsDAO) SelectByPhoneCode(transaction_hash string, code string, phone_number string) (*do.AuthPhoneTransactionsDO, error) {
	var query = "select id from auth_phone_transactions where transaction_hash = ? and code = ? and phone_number = ?"
	rows, err := dao.db.Queryx(query, transaction_hash, code, phone_number)

	if err != nil {
		glog.Error("AuthPhoneTransactionsDAO/SelectByPhoneCode error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.AuthPhoneTransactionsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("AuthPhoneTransactionsDAO/SelectByPhoneCode error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
