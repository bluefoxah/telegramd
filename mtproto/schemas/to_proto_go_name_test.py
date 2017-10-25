#!/usr/bin/python
#-*- coding: utf-8 -*-
#encoding=utf-8

'''
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
'''

def to_proto_go_name(name):
    ss = name.split("_")
    for i in range(len(ss)):
        s = ss[i]
        if i!=0 and s[0:1].isupper():
            ss[i] = '_' + s
        else:
            ss[i] = s[0:1].upper() + s[1:]
    return ''.join(ss)

#print to_proto_go_name('resPQ')
#print to_proto_go_name('CRC32_p_q_inner_data')
#print to_proto_go_name('CRC32_server_DH_params_fail')
#print to_proto_go_name('CRC32_server_DH_params_ok')
#print to_proto_go_name('CRC32_server_DH_inner_data')
print to_proto_go_name('client_DH_inner_data')

