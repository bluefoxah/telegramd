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
{{else if eq .QueryType "SELECT_STRUCT_SINGLE"}}
{{template "SELECT_STRUCT_SINGLE" $v}}
{{else if eq .QueryType "SELECT_STRUCT_LIST"}}
{{template "SELECT_STRUCT_LIST" $v}}
{{else if eq .QueryType "UPDATE"}}
{{template "UPDATE" $v}}
{{else if eq .QueryType "DELETE"}}
{{template "DELETE" $v}}
{{end}}
{{end}}

{{define "INSERT"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}(do *do.{{.TableName}}DO) (id int64, err error) {
	var query = "{{.Sql}}"
	r, err := dao.db.NamedExec(query, do)
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

{{define "SELECT_STRUCT_SINGLE"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (*do.{{.TableName}}DO, error) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    rows, err := dao.db.Queryx(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    rows, err := dao.db.Queryx(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
		glog.Error("{{.TableName}}DAO/{{.FuncName}} error: ", err)
		return nil, err
	}

	defer rows.Close()

	do := &do.{{.TableName}}DO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			glog.Error("{{.TableName}}DAO/{{.FuncName}} error: ", err)
			return nil, err
		}
	} else {
		return nil, nil
	}

	return do, nil
}
{{end}}

{{define "SELECT_STRUCT_LIST"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) ([]do.{{.TableName}}DO, error) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    rows, err := dao.db.Queryx(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    rows, err := dao.db.Queryx(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
	if err != nil {
		glog.Errorf("{{.TableName}}DAO/{{.FuncName}} error: ", err)
		return nil, err
	}

	defer rows.Close()

	var values []do.{{.TableName}}DO
	for rows.Next() {
        v := do.{{.TableName}}DO{}

        // TODO(@benqi): 不使用反射
        err := rows.StructScan(&v)
        if err != nil {
            glog.Errorf("{{.TableName}}DAO/{{.FuncName}} error: %s", err)
            return nil, err
        }
		values = append(values, v)
    }

    return values, nil
}
{{end}}


{{define "UPDATE"}}
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (rows int64, err error) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    r, err := dao.db.Exec(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    r, err := dao.db.Exec(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
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
// {{.Sql}}
// TODO(@benqi): sqlmap
func (dao *{{.TableName}}DAO) {{.FuncName}}({{ range $i, $v := .Params }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{.Type}} {{end}}) (rows int64, err error) {
{{if eq .ParamHasList "true"}}  var q = "{{.CompiledByNamedSql}}"
    query, a, err := sqlx.In(q, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
    r, err := dao.db.Exec(query, a...)
{{else}} var query = "{{.CompiledByNamedSql}}"
    r, err := dao.db.Exec(query, {{range $i, $v := .QueryParams }} {{if ne $i 0 }} , {{end}} {{.FieldName}} {{end}})
{{end}}
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
