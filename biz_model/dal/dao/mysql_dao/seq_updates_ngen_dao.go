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

type SeqUpdatesNgenDAO struct {
	db *sqlx.DB
}

func NewSeqUpdatesNgenDAO(db *sqlx.DB) *SeqUpdatesNgenDAO {
	return &SeqUpdatesNgenDAO{db}
}

// insert into seq_updates_ngen(seq_name, seq, created_at) values (:seq_name, :seq, :created_at)
// TODO(@benqi): sqlmap
func (dao *SeqUpdatesNgenDAO) Insert(do *do.SeqUpdatesNgenDO) (id int64, err error) {
	var query = "insert into seq_updates_ngen(seq_name, seq, created_at) values (:seq_name, :seq, :created_at)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("SeqUpdatesNgenDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("SeqUpdatesNgenDAO/LastInsertId error: ", err)
	}
	return
}

// update seq_updates_ngen set seq = :seq where seq_name = :seq_name
// TODO(@benqi): sqlmap
func (dao *SeqUpdatesNgenDAO) UpdateSeqBySeqName(seq int64, seq_name string) (rows int64, err error) {
	var query = "update seq_updates_ngen set seq = ? where seq_name = ?"
	r, err := dao.db.Exec(query, seq, seq_name)

	if err != nil {
		glog.Error("SeqUpdatesNgenDAO/UpdateSeqBySeqName error: ", err)
		return
	}

	rows, err = r.RowsAffected()
	if err != nil {
		glog.Error("SeqUpdatesNgenDAO/RowsAffected error: ", err)
	}
	return
}

// select seq_name, seq from seq_updates_ngen where seq_name = :seq_name
// TODO(@benqi): sqlmap
func (dao *SeqUpdatesNgenDAO) SelectBySeqName(seq_name string) (*do.SeqUpdatesNgenDO, error) {
	var query = "select seq_name, seq from seq_updates_ngen where seq_name = ?"
	rows, err := dao.db.Queryx(query, seq_name)

	if err != nil {
		glog.Error("SeqUpdatesNgenDAO/SelectBySeqName error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.SeqUpdatesNgenDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("SeqUpdatesNgenDAO/SelectBySeqName error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
