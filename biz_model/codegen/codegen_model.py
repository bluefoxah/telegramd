#!/usr/bin/python
#-*- coding: utf-8 -*-
#encoding=utf-8

'''
/*
 *  Copyright (c) 2017, https://github.com/nebula-im/nebula
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
 // 最后一步，flags检查
'''

import MySQLdb


header = '''/*
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
)

'''

init = '''
func init()  {
    orm.RegisterModel(new(%s))
}
'''

def ToCamelName(name):
    ss = name.split("_")
    for i in range(len(ss)):
        s = ss[i]
        ss[i] = s[0:1].upper() + s[1:]

    return ''.join(ss)

def ToType(name):
    v = name
    idx = name.find('(')
    if (idx > 0):
        v = name[0:idx]

    if (v == 'int'):
        v = 'int32'
    elif (v == 'bigint'):
        v = 'int64'
    elif (v == 'varchar'):
        v = 'string'
    elif (v == 'timestamp'):
        #v = 'time.Time'
        v = 'string'
    elif (v == 'tinyint'):
        v = 'int32'
    else:
        v = name
    return v


conn= MySQLdb.connect(
    host='127.0.0.1',
    port = 3306,
    user='root',
    passwd='',
    db ='nebulaim',
)
cur = conn.cursor()

sz = cur.execute("show tables")
tables = cur.fetchmany(sz)

'''
already_header = ''
if os.path.isfile(proto_file):
    with open(output_proto, 'r') as already:
        already_header = already.read()
if already_header != proto_file:
    with open(output_proto, 'w') as out:
        out.write(proto_file)
'''

for t in tables:
    t_sz = cur.execute('desc ' + t[0])
    t_flds = cur.fetchmany(t_sz)

    content = header
    content += 'type %s struct {\n' % (ToCamelName(t[0]))
    for flds in t_flds:
        content += '  ' + ToCamelName(flds[0]) + ' ' + ToType(flds[1]) + '\n'

    content += '}\n'

    content += init % (ToCamelName(t[0]))

    with open('../model'+t[0]+'.go', 'w') as out:
        out.write(content)

cur.close()
conn.commit()
conn.close()
