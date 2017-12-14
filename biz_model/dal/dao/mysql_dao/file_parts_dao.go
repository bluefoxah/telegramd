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

type FilePartsDAO struct {
	db *sqlx.DB
}

func NewFilePartsDAO(db *sqlx.DB) *FilePartsDAO {
	return &FilePartsDAO{db}
}

// insert into file_parts(creator_user_id, file_id, file_part, is_big_file, file_total_parts, bytes) values (:creator_user_id, :file_id, :file_part, :is_big_file, :file_total_parts, :bytes)
// TODO(@benqi): sqlmap
func (dao *FilePartsDAO) Insert(do *dataobject.FilePartsDO) int64 {
	var query = "insert into file_parts(creator_user_id, file_id, file_part, is_big_file, file_total_parts, bytes) values (:creator_user_id, :file_id, :file_part, :is_big_file, :file_total_parts, :bytes)"
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

// select creator_user_id, file_id, file_part, is_big_file, file_total_parts, bytes from file_parts where file_id = :file_id order by file_part asc
// TODO(@benqi): sqlmap
func (dao *FilePartsDAO) SelectFileParts(file_id int64) []dataobject.FilePartsDO {
	var query = "select creator_user_id, file_id, file_part, is_big_file, file_total_parts, bytes from file_parts where file_id = ? order by file_part asc"
	rows, err := dao.db.Queryx(query, file_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectFileParts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.FilePartsDO
	for rows.Next() {
		v := dataobject.FilePartsDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectFileParts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}
