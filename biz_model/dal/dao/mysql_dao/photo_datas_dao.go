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

type PhotoDatasDAO struct {
	db *sqlx.DB
}

func NewPhotoDatasDAO(db *sqlx.DB) *PhotoDatasDAO {
	return &PhotoDatasDAO{db}
}

// insert into photo_datas(photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, bytes) values (:photo_id, :photo_type, :dc_id, :volume_id, :local_id, :access_hash, :width, :height, :bytes)
// TODO(@benqi): sqlmap
func (dao *PhotoDatasDAO) Insert(do *dataobject.PhotoDatasDO) int64 {
	var query = "insert into photo_datas(photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, bytes) values (:photo_id, :photo_type, :dc_id, :volume_id, :local_id, :access_hash, :width, :height, :bytes)"
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

// select id, photo_id, photo_type, bytes from photo_datas where dc_id = 2 and volume_id = :volume_id and local_id = :local_id and access_hash = :access_hash
// TODO(@benqi): sqlmap
func (dao *PhotoDatasDAO) SelectByFileLocation(volume_id int64, local_id int32, access_hash int64) *dataobject.PhotoDatasDO {
	var query = "select id, photo_id, photo_type, bytes from photo_datas where dc_id = 2 and volume_id = ? and local_id = ? and access_hash = ?"
	rows, err := dao.db.Queryx(query, volume_id, local_id, access_hash)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByFileLocation(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.PhotoDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByFileLocation(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
