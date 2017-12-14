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

type FilesDAO struct {
	db *sqlx.DB
}

func NewFilesDAO(db *sqlx.DB) *FilesDAO {
	return &FilesDAO{db}
}

// insert into files(creator_user_id, file_id, access_hash, file_parts, file_size, md5_checksum) values (:creator_user_id, :file_id, :access_hash, :file_parts, :file_size, :md5_checksum)
// TODO(@benqi): sqlmap
func (dao *FilesDAO) Insert(do *dataobject.FilesDO) int64 {
	var query = "insert into files(creator_user_id, file_id, access_hash, file_parts, file_size, md5_checksum) values (:creator_user_id, :file_id, :access_hash, :file_parts, :file_size, :md5_checksum)"
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

// select id, creator_user_id, file_id, access_hash, file_parts, file_size, md5_checksum from files where file_id = :file_id and file_parts = :file_parts limit 1
// TODO(@benqi): sqlmap
func (dao *FilesDAO) SelectByIDAndParts(file_id int64, file_parts int32) *dataobject.FilesDO {
	var query = "select id, creator_user_id, file_id, access_hash, file_parts, file_size, md5_checksum from files where file_id = ? and file_parts = ? limit 1"
	rows, err := dao.db.Queryx(query, file_id, file_parts)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByIDAndParts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.FilesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByIDAndParts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
