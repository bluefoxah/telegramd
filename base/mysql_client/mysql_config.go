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

package mysql_client

import (
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
)

type MySQLConfig struct {
	Name   string // for trace
	DSN    string // data source name
	Active int    // pool
	Idle   int    // pool
}

func NewSqlxDB(c* MySQLConfig) (db *sqlx.DB) {
	if db, err := sqlx.Connect("mysql", c.DSN); db != nil {
		glog.Errorf("Connect db error: %s", err)
	}

	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	return
}
