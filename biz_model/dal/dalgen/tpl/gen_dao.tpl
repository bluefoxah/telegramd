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

package dao

import(
	do "github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/jmoiron/sqlx"
	"github.com/golang/glog"
)

type {{.Name}}DAO struct {
	db *sqlx.DB
}

func New{{.Name}}DAO(db* sqlx.DB) *{{.Name}}DAO {
	return &{{.Name}}DAO{db}
}

{{range $i, $v := .Funcs }}
{{if eq .QueryType "INSERT"}}
{{template "INSERT" $v}}
{{else if eq .QueryType "SELECT"}}
{{template "SELECT" $v}}
{{else if eq .QueryType "UPDATE"}}
{{template "UPDATE" $v}}
{{else if eq .QueryType "DELETE"}}
{{template "DELETE" $v}}
{{end}}
{{end}}

{{define "INSERT"}}
func (dao *{{.TableName}}DAO) {{.FuncName}}(do *do.{{.TableName}}DO) (id int64, err error) {
	// TODO(@benqi): sqlmap
	id = 0

	var sql = "{{.Sql}}"
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("{{.TableName}}DAO/{{.FuncName}} error: ", err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		glog.Error("{{.TableName}}DAO/LastInsertId error: ", err)
	}
	return
}
{{end}}

{{define "SELECT"}}
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (*do.{{.TableName}}DO, error) {
	// TODO(@benqi): sqlmap
	var sql = "{{.Sql}}"
	do := &do.{{.TableName}}DO{ {{ range $i, $v := .Params }} {{.Name}} : {{.FieldName}}, {{end}} }
	r, err := dao.db.NamedQuery(sql, do)
	if err != nil {
		glog.Error("{{.TableName}}DAO/SelectById error: ", err)
		return nil, err
	}

	if r.Next() {
		err = r.StructScan(do)
		if err != nil {
			glog.Error("{{.TableName}}DAO/SelectById error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
{{end}}

{{define "UPDATE"}}
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (rows int64, err error) {
	// TODO(@benqi): sqlmap
	rows = 0

	var sql = "{{.Sql}}"
	do := &do.{{.TableName}}DO{ {{ range $i, $v := .Params }} {{.Name}} : {{.FieldName}}, {{end}} }
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("{{.TableName}}DAO/{{.FuncName}} error: ", err)
		return
	}

	rows, err = r.RowsAffected()
	if err != nil {
		glog.Error("{{.TableName}}DAO/RowsAffected error: ", err)
	}
	return
}
{{end}}

{{define "DELETE"}}
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (rows int64, err error) {
	// TODO(@benqi): sqlmap
	rows = 0

	var sql = "{{.Sql}}"
	do := &do.{{.TableName}}DO{ {{ range $i, $v := .Params }} {{.Name}} : {{.FieldName}}, {{end}} }
	r, err := dao.db.NamedExec(sql, do)
	if err != nil {
		glog.Error("{{.TableName}}DAO/{{.FuncName}} error: ", err)
		return
	}

	rows, err = r.RowsAffected()
	if err != nil {
		glog.Error("{{.TableName}}DAO/RowsAffected error: ", err)
	}
	return
}
{{end}}
