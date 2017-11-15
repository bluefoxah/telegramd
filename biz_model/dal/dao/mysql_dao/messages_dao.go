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

type MessagesDAO struct {
	db *sqlx.DB
}

func NewMessagesDAO(db *sqlx.DB) *MessagesDAO {
	return &MessagesDAO{db}
}

// insert into messages(user_id, peer_type, peer_id, random_id, message, `date`, created_at) values (:user_id, :peer_type, :peer_id, :random_id, :message, :date, :created_at)
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) Insert(do *do.MessagesDO) (id int64, err error) {
	var query = "insert into messages(user_id, peer_type, peer_id, random_id, message, `date`, created_at) values (:user_id, :peer_type, :peer_id, :random_id, :message, :date, :created_at)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		glog.Error("MessagesDAO/Insert error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("MessagesDAO/LastInsertId error: ", err)
	}
	return
}

// select id, user_id, peer_type, peer_id, random_id, message, `date` from messages where id in (:idList)
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) SelectByIdList(idList []int32) ([]do.MessagesDO, error) {
	var q = "select id, user_id, peer_type, peer_id, random_id, message, `date` from messages where id in (?)"
	query, a, err := sqlx.In(q, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		glog.Errorf("MessagesDAO/SelectByIdList error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.MessagesDO
	for rows.Next() {
		v := do.MessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("MessagesDAO/SelectByIdList error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}

// select id, user_id, peer_type, peer_id, random_id, message, `date` from messages where peer_type = :peer_type and (user_id = :user_id and peer_id = :peer_id) or (user_id = :peer_id and peer_id = :user_id)
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) SelectByUserIdAndPeer(peer_type int32, user_id int32, peer_id int32) ([]do.MessagesDO, error) {
	var query = "select id, user_id, peer_type, peer_id, random_id, message, `date` from messages where peer_type = ? and (user_id = ? and peer_id = ?) or (user_id = ? and peer_id = ?)"
	rows, err := dao.db.Queryx(query, peer_type, user_id, peer_id, peer_id, user_id)

	if err != nil {
		glog.Errorf("MessagesDAO/SelectByUserIdAndPeer error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.MessagesDO
	for rows.Next() {
		v := do.MessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("MessagesDAO/SelectByUserIdAndPeer error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}

// select id, user_id, peer_type, peer_id, random_id, message, `date` from messages where id > :offset_id and peer_type = :peer_type and ((user_id = :user_id and peer_id = :peer_id) or (user_id = :peer_id and peer_id = :user_id)) limit :limit
// TODO(@benqi): sqlmap
func (dao *MessagesDAO) SelectByUserIdAndPeerOffsetLimit(offset_id int32, peer_type int32, user_id int32, peer_id int32, limit int32) ([]do.MessagesDO, error) {
	var query = "select id, user_id, peer_type, peer_id, random_id, message, `date` from messages where id > ? and peer_type = ? and ((user_id = ? and peer_id = ?) or (user_id = ? and peer_id = ?)) limit ?"
	rows, err := dao.db.Queryx(query, offset_id, peer_type, user_id, peer_id, peer_id, user_id, limit)

	if err != nil {
		glog.Errorf("MessagesDAO/SelectByUserIdAndPeerOffsetLimit error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.MessagesDO
	for rows.Next() {
		v := do.MessagesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			glog.Errorf("MessagesDAO/SelectByUserIdAndPeerOffsetLimit error: %s", err)
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
