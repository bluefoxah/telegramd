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
	"github.com/nebulaim/telegramd/base/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"testing"
	"github.com/nebulaim/telegramd/frontend/id"
	"fmt"
)


func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	// register model
	// orm.RegisterModel(new(AuthSessions))
	// set default database
	err = orm.RegisterDataBase("default", "mysql", "root:@/nebulaim?charset=utf8", 30)
	if err != nil {
		panic(err)
	}
}

func TestAuthSessions(t *testing.T) {
	orm := orm.NewOrm()

	tb := &AuthSessions{
		AuthId: 	1,
		SessionId:	1,
		UniqueId:   id.NextId(),
	}

	fmt.Printf("auth_sessions: %v\n", tb)
	_, id, _ := orm.ReadOrCreate(tb, "AuthId", "SessionId")
	tb.Id = uint32(id)
	tb.UniqueId = 1
	orm.Read(tb)
	fmt.Printf("auth_sessions: %v\n", tb)
}

