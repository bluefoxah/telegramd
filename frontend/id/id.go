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

package id

import "github.com/nebulaim/telegramd/base/snowflake"

var id *snowflake.IdWorker

// = &snowflake.IdWorker{
//
//}

const (
	workerId	   	= int64(1)
	dataCenterId	= int64(1)
	twepoch        	= int64(1288834974657)
)

func init()  {
	id, _ = snowflake.NewIdWorker(workerId, dataCenterId, twepoch)
}

func NextId() (int64) {
	r, _ := id.NextId()
	return r
}
