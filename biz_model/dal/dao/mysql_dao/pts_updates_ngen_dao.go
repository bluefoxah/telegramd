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

type PtsUpdatesNgenDAO struct {
	db *sqlx.DB
}

func NewPtsUpdatesNgenDAO(db *sqlx.DB) *PtsUpdatesNgenDAO {
	return &PtsUpdatesNgenDAO{db}
}

// insert into pts_updates_ngen(user_id, pts, update_type, update_data, date2) values (:user_id, :pts, :update_type, :update_data, :date2)
// TODO(@benqi): sqlmap
func (dao *PtsUpdatesNgenDAO) Insert(do *dataobject.PtsUpdatesNgenDO) int64 {
	var query = "insert into pts_updates_ngen(user_id, pts, update_type, update_data, date2) values (:user_id, :pts, :update_type, :update_data, :date2)"
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

// select pts from pts_updates_ngen where user_id = :user_id order by pts desc limit 1
// TODO(@benqi): sqlmap
func (dao *PtsUpdatesNgenDAO) SelectLastPtsByUserId(user_id int32) *dataobject.PtsUpdatesNgenDO {
	var query = "select pts from pts_updates_ngen where user_id = ? order by pts desc limit 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastPtsByUserId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.PtsUpdatesNgenDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastPtsByUserId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
