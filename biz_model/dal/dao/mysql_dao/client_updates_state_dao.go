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

type ClientUpdatesStateDAO struct {
	db *sqlx.DB
}

func NewClientUpdatesStateDAO(db *sqlx.DB) *ClientUpdatesStateDAO {
	return &ClientUpdatesStateDAO{db}
}

// insert into client_updates_state(auth_key_id, user_id, pts, qts, seq, date2) values (:auth_key_id, :user_id, :pts, :qts, :seq, :date2)
// TODO(@benqi): sqlmap
func (dao *ClientUpdatesStateDAO) Insert(do *dataobject.ClientUpdatesStateDO) int64 {
	var query = "insert into client_updates_state(auth_key_id, user_id, pts, qts, seq, date2) values (:auth_key_id, :user_id, :pts, :qts, :seq, :date2)"
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

// select auth_key_id, user_id, pts, qts, seq, date2 from client_updates_state where auth_key_id = :auth_key_id and user_id = :user_id
// TODO(@benqi): sqlmap
func (dao *ClientUpdatesStateDAO) SelectByAuthKeyAndUserId(auth_key_id int64, user_id int32) *dataobject.ClientUpdatesStateDO {
	var query = "select auth_key_id, user_id, pts, qts, seq, date2 from client_updates_state where auth_key_id = ? and user_id = ?"
	rows, err := dao.db.Queryx(query, auth_key_id, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByAuthKeyAndUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.ClientUpdatesStateDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByAuthKeyAndUserId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
