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
	"fmt"
)

type MysqlClientManager struct{
	// TODO(@benqi): 使用sync.Map，动态添加和卸载数据库
	mysqlClients map[string]*sqlx.DB
}

var mysqlClients = &MysqlClientManager{make(map[string]*sqlx.DB)}

func  InstallMysqlClientManager(configs []MySQLConfig) {
	for _, config := range configs {
		client := NewSqlxDB(&config)
		if client == nil {
			err := fmt.Errorf("InstallModelManager - NewSqlxDB {%v} error!", config)
			panic(err)
			// continue
		}

		// TODO(@benqi): 检查config数据合法性
		mysqlClients.mysqlClients[config.Name] = client
	}
}

func GetMysqlClient(dbName string) (client *sqlx.DB) {
	client, ok := mysqlClients.mysqlClients[dbName]
	if !ok {
		glog.Errorf("GetMysqlClient - Not found client: %s", dbName)
	}
	return
}

func GetMysqlClientManager() map[string]*sqlx.DB {
	return mysqlClients.mysqlClients
}
