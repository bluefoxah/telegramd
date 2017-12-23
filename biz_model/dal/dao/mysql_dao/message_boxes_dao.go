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

type MessageBoxesDAO struct {
	db *sqlx.DB
}

func NewMessageBoxesDAO(db *sqlx.DB) *MessageBoxesDAO {
	return &MessageBoxesDAO{db}
}

// insert into message_boxes(user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, date2, created_at) values (:user_id, :sender_user_id, :message_box_type, :peer_type, :peer_id, :pts, :message_id, :media_unread, :date2, :created_at)
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) Insert(do *dataobject.MessageBoxesDO) int64 {
	var query = "insert into message_boxes(user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, date2, created_at) values (:user_id, :sender_user_id, :message_box_type, :peer_type, :peer_id, :pts, :message_id, :media_unread, :date2, :created_at)"
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

// select pts from message_boxes where user_id = :user_id order by pts desc limit 1
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectLastPts(user_id int32) *dataobject.MessageBoxesDO {
	var query = "select pts from message_boxes where user_id = ? order by pts desc limit 1"
	rows, err := dao.db.Queryx(query, user_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastPts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select pts from message_boxes where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id order by pts desc limit 1
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectLastPtsByPeer(user_id int32, peer_type int8, peer_id int32) *dataobject.MessageBoxesDO {
	var query = "select pts from message_boxes where user_id = ? and peer_type = ? and peer_id = ? order by pts desc limit 1"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectLastPtsByPeer(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectLastPtsByPeer(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}

// select pts from message_boxes where user_id = :user_id and message_id > :message_id order by pts asc
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectPtsByGtMessageID(user_id int32, message_id int32) []dataobject.MessageBoxesDO {
	var query = "select pts from message_boxes where user_id = ? and message_id > ? order by pts asc"
	rows, err := dao.db.Queryx(query, user_id, message_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectPtsByGtMessageID(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectPtsByGtMessageID(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and message_id < :message_id order by message_id desc limit :limit
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectBackwardByPeerOffsetLimit(user_id int32, peer_type int8, peer_id int32, message_id int32, limit int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = ? and peer_type = ? and peer_id = ? and message_id < ? order by message_id desc limit ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id, message_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectBackwardByPeerOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectBackwardByPeerOffsetLimit(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and message_id >= :message_id order by message_id asc limit :limit
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectForwardByPeerOffsetLimit(user_id int32, peer_type int8, peer_id int32, message_id int32, limit int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = ? and peer_type = ? and peer_id = ? and message_id >= ? order by message_id asc limit ?"
	rows, err := dao.db.Queryx(query, user_id, peer_type, peer_id, message_id, limit)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectForwardByPeerOffsetLimit(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectForwardByPeerOffsetLimit(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = :user_id and pts > :pts order by pts asc
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByGtPts(user_id int32, pts int32) []dataobject.MessageBoxesDO {
	var query = "select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = ? and pts > ? order by pts asc"
	rows, err := dao.db.Queryx(query, user_id, pts)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByGtPts(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByGtPts(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = :user_id and message_id in (:idList)
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByMessageIdList(user_id int32, idList []int32) []dataobject.MessageBoxesDO {
	var q = "select user_id, sender_user_id, message_box_type, peer_type, peer_id, pts, message_id, media_unread, state, date2 from message_boxes where user_id = ? and message_id in (?)"
	query, a, err := sqlx.In(q, user_id, idList)
	rows, err := dao.db.Queryx(query, a...)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByMessageIdList(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	var values []dataobject.MessageBoxesDO
	for rows.Next() {
		v := dataobject.MessageBoxesDO{}

		// TODO(@benqi): 不使用反射
		err := rows.StructScan(&v)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByMessageIdList(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
		values = append(values, v)
	}

	return values
}

// select user_id, user_message_box_id, message_id, sender_user_id, message_box_type, peer_type, peer_id, media_unread, date2 from message_boxed where user_id = :user_id and user_message_box_id = :user_message_box_id limit 1
// TODO(@benqi): sqlmap
func (dao *MessageBoxesDAO) SelectByUserIdAndMessageBoxId(user_id int32, user_message_box_id int32) *dataobject.MessageBoxesDO {
	var query = "select user_id, user_message_box_id, message_id, sender_user_id, message_box_type, peer_type, peer_id, media_unread, date2 from message_boxed where user_id = ? and user_message_box_id = ? limit 1"
	rows, err := dao.db.Queryx(query, user_id, user_message_box_id)

	if err != nil {
		errDesc := fmt.Sprintf("Queryx in SelectByUserIdAndMessageBoxId(_), error: %v", err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	defer rows.Close()

	do := &dataobject.MessageBoxesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			errDesc := fmt.Sprintf("StructScan in SelectByUserIdAndMessageBoxId(_), error: %v", err)
			glog.Error(errDesc)
			panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
		}
	} else {
		return nil
	}

	return do
}
