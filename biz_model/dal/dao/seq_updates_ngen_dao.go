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

type SeqUpdatesNgenDAO struct {
	db *sqlx.DB
}

func NewSeqUpdatesNgenDAO(db *sqlx.DB) *SeqUpdatesNgenDAO {
	return &SeqUpdatesNgenDAO{db}
}

func (dao *SeqUpdatesNgenDAO) Insert(do *do.SeqUpdatesNgenDO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "insert into seq_updates_ngen(seq_name, seq, created_at) values (:seq_name, :seq, :created_at)"
	r, err := dao.db.NamedExec(sql, do)
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

func (dao *SeqUpdatesNgenDAO) UpdateSeqBySeqName(seq int64, seq_name string) (rows int64, err error) {
	// TODO(@benqi): sqlmap
	rows = 0

	var sql = "update seq_updates_ngen set seq = :seq where seq_name = :seq_name"
	do := &do.SeqUpdatesNgenDO{Seq: seq, SeqName: seq_name}
	r, err := dao.db.NamedExec(sql, do)
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

func (dao *SeqUpdatesNgenDAO) SelectBySeqName(seq_name string) (*do.SeqUpdatesNgenDO, error) {
	// TODO(@benqi): sqlmap
	var sql = "select seq_name, seq from seq_updates_ngen where seq_name = :seq_name"
	do := &do.SeqUpdatesNgenDO{SeqName: seq_name}
	rows, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("SeqUpdatesNgenDAO/SelectBySeqName error: ", err)
		return nil, err
	}

	defer rows.Close()
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
