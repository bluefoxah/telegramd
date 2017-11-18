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

package model

import (
	"time"
	"sync"
	dao2 "github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

var (
	sequenceInstance *sequnceModel
	sequenceInstanceOnce sync.Once
)

type sequnceModel struct {
	// chatDAO *dao.UserDialogsDAO
}

func GetSequenceModel() *sequnceModel {
	sequenceInstanceOnce.Do(func() {
		sequenceInstance = &sequnceModel{}
	})
	return sequenceInstance
}

// TODO(@benqi):
//  使用数据库和REDIS获取sequence
//  redis: sequence
//  暂时不考虑DB等异常处理
func (dao *sequnceModel) NextID(key string) (seq int64) {
	sequenceDAO := dao2.GetSequenceDAO(dao2.CACHE)

	seq, _ = sequenceDAO.Incr(key)
	var do *dataobject.SeqUpdatesNgenDO = nil

	// 使用seq==1做为哨兵减少DB和REDIS操作
	if seq == 1 {
		// seq为1，有两种情况:
		// 1. 没有指定key的seq，第一次生成seq，DB也无纪录
		// 2. redis重新启动，DB里可能已经有值

		SeqUpdatesNgenDAO := dao2.GetSeqUpdatesNgenDAO(dao2.DB_SLAVE)
		do = SeqUpdatesNgenDAO.SelectBySeqName(key)
		if do == nil {
			// DB无值，插入数据
			do = &dataobject.SeqUpdatesNgenDO{
				SeqName:   key,
				Seq:       seq,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			}
		} else {
			// DB有seq
			do.Seq += 1
			sequenceDAO.Set(key, do.Seq)
		}
	} else {
		do = &dataobject.SeqUpdatesNgenDO{
			SeqName: key,
			Seq:     seq,
		}
	}

	// TODO(@benqi): 使用一些策略减少存盘次数
	SeqUpdatesNgenDAO := dao2.GetSeqUpdatesNgenDAO(dao2.DB_MASTER)

	if do.Seq == 1 {
		SeqUpdatesNgenDAO.Insert(do)
	} else {
		SeqUpdatesNgenDAO.UpdateSeqBySeqName(do.Seq, key)
	}

	return
}

