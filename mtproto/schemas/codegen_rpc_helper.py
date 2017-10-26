#!/usr/bin/python
#-*- coding: utf-8 -*-
#encoding=utf-8

'''
/*
 *  Copyright (c) 2016, https://github.com/nebulaim
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

import glob, re, binascii, os, sys


def txt_wrap_by(start_str, end, line):
  start = line.find(start_str)
  if start >= 0:
    start += len(start_str)
  end = line.find(end, start)
  if end >= 0:
    return line[start:end].strip()

input_file = ''
output_path = ''
next_output_path = False
for arg in sys.argv[1:]:
  if next_output_path:
    next_output_path = False
    output_path = arg
  elif arg == '-o':
    next_output_path = True
  elif re.match(r'^-o(.+)', arg):
    output_path = arg[2:]
  else:
    input_file = arg

if input_file == '':
  print('Input file required.')
  sys.exit(1)
if output_path == '':
  print('Output path required.')
  sys.exit(1)

output_proto = output_path # + '/schema.tl.proto'

# this is a map (key flags -> map (flag name -> flag bit))
# each key flag of parentFlags should be a subset of the value flag here
parentFlagsCheck = {};

layer = '';
funcs = 0
types = 0;
consts = 0
funcsNow = 0
enums = [];

funcsDict = {};
FuncsDict = {};
funcsList = [];

typesDict = {};
TypesDict = {};
typesList = [];
TypesList = [];

typesText = '';
creatorProxyText = '';
inlineMethods = '';
textSerializeInit = '';
textSerializeMethods = '';

classTypesTexts = '';
resClassTypesTexts = '';
classFuncsTexts = '';
registers = [];

with open(input_file) as f:
  for line in f:
    line=line.strip('\n')
    layerline = re.match(r'// LAYER (\d+)', line)
    if (layerline):
      # 当前层
      layer = 'constexpr int CURRENT_LAYER = ' + layerline.group(1) + ';';

    nocomment = re.match(r'^(.*?)//', line)
    if (nocomment):
      line = nocomment.group(1);
    if (re.match(r'\-\-\-functions\-\-\-', line)):
      funcsNow = 1;
      continue;
    if (re.match(r'\-\-\-types\-\-\-', line)):
      funcsNow = 0;
      continue;
    if (re.match(r'^\s*$', line)):
      continue;

    nametype = re.match(r'([a-zA-Z\.0-9_]+)#([0-9a-f]+)([^=]*)=\s*([a-zA-Z\.<>0-9_]+);', line);
    if (not nametype):
      # 特殊处理 vector#1cb5c415 {t:Type} # [ t ] = Vector t;
      if (not re.match(r'vector#1cb5c415 \{t:Type\} # \[ t \] = Vector t;', line)):
        print('Bad line found: ' + line);
      print(line)
      continue;

    # resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
    # name为: resPQ
    # contest.saveDeveloperInfo#9a5f6e95 vk_id:int name:string phone_number:string age:int city:string = Bool;
    # name为: contest_saveDeveloperInfo
    name = nametype.group(1);
    nameInd = name.find('.');
    if (nameInd >= 0):
      Name = name[0:nameInd] + '_' + name[nameInd + 1:nameInd + 2].upper() + name[nameInd + 2:];
      name = name.replace('.', '_');
    else:
      Name = name[0:1].upper() + name[1:];

    # typeid为: 05162463
    typeid = nametype.group(2);
    # 去掉前面的0
    while (len(typeid) > 0 and typeid[0] == '0'):
      typeid = typeid[1:];
    if (len(typeid) == 0):
      typeid = '0';
    typeid = '0x' + typeid;

    cleanline = nametype.group(1) + nametype.group(3) + '= ' + nametype.group(4);
    cleanline = re.sub(r' [a-zA-Z0-9_]+\:flags\.[0-9]+\?true', '', cleanline);
    cleanline = cleanline.replace('<', ' ').replace('>', ' ').replace('  ', ' ');
    cleanline = re.sub(r'^ ', '', cleanline);
    cleanline = re.sub(r' $', '', cleanline);
    cleanline = cleanline.replace(':bytes ', ':string ');
    cleanline = cleanline.replace('?bytes ', '?string ');
    cleanline = cleanline.replace('{', '');
    cleanline = cleanline.replace('}', '');

    # 通过cleanline计算出typeid并进行验证
    countTypeId = binascii.crc32(binascii.a2b_qp(cleanline));
    if (countTypeId < 0):
      countTypeId += 2 ** 32;
    countTypeId = '0x' + re.sub(r'^0x|L$', '', hex(countTypeId));
    if (typeid != countTypeId):
      print('Warning: counted ' + countTypeId + ' mismatch with provided ' + typeid + ' (' + cleanline + ')');
      continue;

    typeid = binascii.crc32(binascii.a2b_qp(cleanline));

    # params为: nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long>
    params = nametype.group(3);
    # restype为: ResPQ
    restype = nametype.group(4);
    if (restype.find('<') >= 0):
      # vector
      templ = re.match(r'^([vV]ector<)([A-Za-z0-9\._]+)>$', restype);
      if (templ):
        vectemplate = templ.group(2);
        if (re.match(r'^[A-Z]', vectemplate) or re.match(r'^[a-zA-Z0-9]+_[A-Z]', vectemplate)):
          # restype = templ.group(1) + 'MTP' + vectemplate.replace('.', '_') + '>';
          # restype = 'std::vector<std::shared_ptr<' + vectemplate.replace('.', '_') + '>>';
          restype = templ.group(1) + vectemplate.replace('.', '_') + '>';
          # print('name: ' + name + ', object: ' + restype);
        elif (vectemplate == 'int' or vectemplate == 'long' or vectemplate == 'string'):
          if (vectemplate == 'int'):
            vectemplate = 'int32_t';
          elif (vectemplate == 'long'):
            vectemplate = 'int64_t';
          else:
            vectemplate = 'std::string';
          # restype = templ.group(1) + 'MTP' + vectemplate.replace('.', '_') + '>';
          # restype = 'std::vector<' + vectemplate.replace('.', '_') + '>';
          restype = templ.group(1) + vectemplate.replace('.', '_') + '>';
          # print('name: ' + name + ', int/long/string: ' + restype);
        else:
          foundmeta = '';
          for metatype in typesDict:
            for typedata in typesDict[metatype]:
              if (typedata[0] == vectemplate):
                foundmeta = metatype;
                break;
            if (len(foundmeta) > 0):
              break;
          if (len(foundmeta) > 0):
            # ptype = templ.group(1) + 'MTP' + foundmeta.replace('.', '_') + '>';
            ptype = templ.group(1) + foundmeta.replace('.', '_') + '>';
            # print('name: ' + name + ', foundmeta: ' + ptype);
          else:
            print('Bad vector param: ' + vectemplate);
            continue;
      else:
        print('Bad template type: ' + restype);
        continue;

    resType = restype.replace('.', '_');
    # print('restype: ' + restype + ', resType: ' + resType);

    if (restype.find('.') >= 0):
      parts = re.match(r'([a-z]+)\.([A-Z][A-Za-z0-9<>\._]+)', restype)
      if (parts):
        restype = parts.group(1) + '_' + parts.group(2)[0:1].lower() + parts.group(2)[1:];
      else:
        print('Bad result type name with dot: ' + restype);
        continue;
    else:
      if (re.match(r'^[A-Z]', restype)):
        restype = restype[:1].lower() + restype[1:];
      else:
        print('Bad result type name: ' + restype);
        continue;

    # print('name: %s, typeid: %s, params: %s, resType: %s, restype: %s' %(name, typeid, params, resType, restype));

    # 生成: REGISTER_TLOBJECT(TL_resPQ);
    registers.append('  CRC32_' + name + ' = ' + str(typeid) + ';\n');

    # 参数处理
    paramsList = params.strip().split(' ');
    prms = {};
    conditions = {}; # 条件: flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string
    trivialConditions = {}; # true type, allow_flashcall:flags.0?true
    prmsList = [];
    conditionsList = [];
    isTemplate = hasFlags = hasTemplate = '';
    for param in paramsList:
      if (re.match(r'^\s*$', param)):
        continue;
      templ = re.match(r'^{([A-Za-z]+):Type}$', param); # vector#1cb5c415 {t:Type} # [ t ] = Vector t;
      if (templ):
        hasTemplate = templ.group(1);
        # print('hasTemplate: ' + hasTemplate + ', in: ' + cleanline);
        continue;

      pnametype = re.match(r'([a-z_][a-z0-9_]*):([A-Za-z0-9<>\._]+|![a-zA-Z]+|\#|[a-z_][a-z0-9_]*\.[0-9]+\?[A-Za-z0-9<>\._]+)$', param);
      if (not pnametype):
        print('Bad param found: "' + param + '" in line: ' + line);
        continue;

      pname = pnametype.group(1); # 参数名
      ptypewide = pnametype.group(2); # 参数类型

      if (re.match(r'^!([a-zA-Z]+)$', ptypewide)):
        if ('!' + hasTemplate == ptypewide):
          # 模板类型
          isTemplate = pname;
          ptype = 'TQueryType';
          # print('template param name: ' + pname + ', type: TQueryType');
        else:
          print('Bad template param name: "' + param + '" in line: ' + line);
          continue;
      elif (ptypewide == '#'):
        # flags, 类似protobuf的optional字段
        hasFlags = pname;
        ptype = 'int32';

      else:
        ptype = ptypewide;
        if (ptype.find('?') >= 0):
          # flags.0?string
          pmasktype = re.match(r'([a-z_][a-z0-9_]*)\.([0-9]+)\?([A-Za-z0-9<>\._]+)', ptype);
          if (not pmasktype or pmasktype.group(1) != hasFlags):
            print('Bad param found: "' + param + '" in line: ' + line);
            continue;
          ptype = pmasktype.group(3);
          if (ptype.find('<') >= 0):
            # inputMediaUploadedPhoto#630c9af1 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> = InputMedia;
            # print('flags\'s template type: ' + ptype);
            templ = re.match(r'^([vV]ector<)([A-Za-z0-9\._]+)>$', ptype);
            if (templ):
              vectemplate = templ.group(2);
              if (re.match(r'^[A-Z]', vectemplate) or re.match(r'^[a-zA-Z0-9]+_[A-Z]', vectemplate)):
                # ptype = templ.group(1) + 'MTP' + vectemplate.replace('.', '_') + '>';
                # ptype = 'std::vector<std::shared_ptr<' + vectemplate.replace('.', '_') + '>>';
                ptype = 'TLObjectVector<' + vectemplate.replace('.', '_') + '>';

              elif (vectemplate == 'int' or vectemplate == 'long' or vectemplate == 'string'):
                # ptype = templ.group(1) + 'MTP' + vectemplate.replace('.', '_') + '>';

                if (vectemplate == 'int'):
                  ptype = 'repeated int32';

                  # vectemplate = 'int32_t';
                elif (vectemplate == 'long'):
                  ptype = 'repeated int64';

                  # vectemplate = 'int64_t';
                else:
                  ptype = 'repeated int32 string';

                  # vectemplate = 'std::string';

                # ptype = 'std::vector<' + vectemplate.replace('.', '_') + '>';

              else:
                foundmeta = '';
                for metatype in typesDict:
                  for typedata in typesDict[metatype]:
                    if (typedata[0] == vectemplate):
                      foundmeta = metatype;
                      break;
                  if (len(foundmeta) > 0):
                    break;
                if (len(foundmeta) > 0):
                  # ptype = templ.group(1) + 'MTP' + foundmeta.replace('.', '_') + '>';
                  ptype = 'std::vector<' + foundmeta.replace('.', '_') + '>';
                  print('foundmeta: ' + ptype);
                else:
                  print('Bad vector param: ' + vectemplate);
                  continue;
            else:
              print('Bad template type: ' + ptype);
              continue;
          if (not pname in conditions):
            conditionsList.append(pname);
            conditions[pname] = pmasktype.group(2);
            # print('condition: ' + pname + ' --> ' + pmasktype.group(2) + ', ' + ptype);
            if (ptype == 'true'):
              trivialConditions[pname] = 1;
        elif (ptype.find('<') >= 0):
          templ = re.match(r'^([vV]ector<)([A-Za-z0-9\._]+)>$', ptype);
          if (templ):
            vectemplate = templ.group(2);
            if (re.match(r'^[A-Z]', vectemplate) or re.match(r'^[a-zA-Z0-9]+_[A-Z]', vectemplate)):
              # ptype = templ.group(1) + 'MTP' + vectemplate.replace('.', '_') + '>';
              # ptype = 'std::vector<std::shared_ptr<' + vectemplate.replace('.', '_') + '>>';
              ptype = 'TLObjectVector<' + vectemplate.replace('.', '_') + '>';

            elif (vectemplate == 'int' or vectemplate == 'long' or vectemplate == 'string'):
              # ptype = templ.group(1) + 'MTP' + vectemplate.replace('.', '_') + '>';
              if (vectemplate == 'int'):
                ptype = 'repeated int32';

                # vectemplate = 'int32_t';
              elif (vectemplate == 'long'):
                ptype = 'repeated int64';
                # vectemplate = 'int64_t';
              else:
                ptype = 'repeated string';
                # vectemplate = 'std::string';
              # ptype = 'std::vector<' + vectemplate.replace('.', '_') + '>';

            else:
              foundmeta = '';
              for metatype in typesDict:
                for typedata in typesDict[metatype]:
                  if (typedata[0] == vectemplate):
                    foundmeta = metatype;
                    break;
                if (len(foundmeta) > 0):
                  break;
              if (len(foundmeta) > 0):
                # ptype = templ.group(1) + 'MTP' + foundmeta.replace('.', '_') + '>';
                ptype = 'std::vector<TL_' + vectemplate + '*>';
                print('ptype: ' + ptype + ', metatype: ' + metatype + ', vectemplate: ' + vectemplate);
              else:
                print('Bad vector param: ' + vectemplate);
                continue;
          else:
            print('Bad template type: ' + ptype);
            continue;
      prmsList.append(pname);
      # prms[pname] = ptype.replace('.', '_');

      ptype2 = ptype.replace('.', '_');
      if (ptype2 == 'true'):
        ptype2 = 'bool';
      if (ptype2 == 'int'):
        ptype2 = 'int32';
      if (ptype2 == 'long'):
        ptype2 = 'int64';
      if (ptype2 == 'int128'):
        ptype2 = 'bytes';
      if (ptype2 == 'int256'):
        ptype2 = 'bytes';
      if (ptype2 == 'string'):
        ptype2 = 'string';
      if (ptype2 == 'bytes'):
        ptype2 = 'bytes';

      prms[pname] = ptype2;

    # print(prms);

    if (isTemplate == '' and resType == 'X'):
      print('Bad response type "X" in "' + name +'" in line: ' + line);
      continue;

    if funcsNow:
      if (not restype in funcsDict):
        funcsList.append(restype);
        funcsDict[restype] = [];
        # TypesDict[restype] = resType;
      FuncsDict[restype] = resType;
      funcsDict[restype].append([name, typeid, prmsList, prms, hasFlags, conditionsList, conditions, trivialConditions, line]);

      # print(funcsDict[restype])

    else:
      if (isTemplate != ''):
        print('Template types not allowed: "' + resType + '" in line: ' + line);
        continue;
      if (not restype in typesDict):
        typesList.append(restype);
        TypesList.append(resType);

        # print('typeList added: ' + restype);
        typesDict[restype] = [];
      TypesDict[restype] = resType;
      typesDict[restype].append([name, typeid, prmsList, prms, hasFlags, conditionsList, conditions, trivialConditions, line]);

      consts = consts + 1;

      # print(TypesDict[restype])

innerRPC=[]
innerRPC.append('req_pq')
innerRPC.append('req_DH_params')
innerRPC.append('set_client_DH_params')
innerRPC.append('destroy_auth_key')
innerRPC.append('rpc_drop_answer')
innerRPC.append('get_future_salts')
innerRPC.append('ping')
innerRPC.append('ping_delay_disconnect')
innerRPC.append('destroy_session')
innerRPC.append('contest.saveDeveloperInfo')
innerRPC.append('contest_saveDeveloperInfo')
innerRPC.append('invokeAfterMsg')
innerRPC.append('invokeAfterMsgs')
innerRPC.append('initConnection')
innerRPC.append('invokeWithLayer')
innerRPC.append('invokeWithoutUpdates')

def to_proto_go_name(name):
  name2 = name
  if (name == 'udp_p2p'):
    name2 = 'UdpP2P'
  ss = name2.split("_")
  for i in range(len(ss)):
    s = ss[i]
    if i!=0 and s[0:1].isupper():
      ss[i] = '_' + s
    else:
      ss[i] = s[0:1].upper() + s[1:]
  return ''.join(ss)

def to_proto_first_name(name):
  ss = name.split("_")
  s = ss[0]
  return s[0:1].upper() + s[1:]

classTypesTexts2 = ''
#classTypesTexts += 'service RPCQuery {\n'
for restype in funcsList:
  v = funcsDict[restype];
  for data in v:
    name = data[0]

    if (name in innerRPC):
      continue

    #resType2 = TypesDict[''.restype]
    resType = FuncsDict[restype]
    if (resType.find('Vector') >= 0):
      classTypesTexts2 += '  // "TL' + to_proto_go_name(name) + '": RPCContextTuple{"/mtproto.RPC' + to_proto_first_name(name) + '/' + name + '", func() interface{} { return new(' + to_proto_go_name(resType) + ') }},\n'
    else:
      classTypesTexts += '  "TL' + to_proto_go_name(name) + '": RPCContextTuple{"/mtproto.RPC' + to_proto_first_name(name) + '/' + name + '", func() interface{} { return new(' + to_proto_go_name(resType) + ') }},\n'

proto_file = '''/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_rpc_helper.py'
 *
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

package mtproto

import (
	"reflect"
	"github.com/golang/glog"
)

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct  {
	Method string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple {
'''

proto_file += classTypesTexts

proto_file += '\n  // TODO(@benqi): Not support Vector Reply\n'

proto_file += classTypesTexts2

proto_file += '''}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		glog.Error("Can't find name: ", rt.Name())
		return nil
	}
	return &m
}
'''

already_header = ''
if os.path.isfile(proto_file):
  with open(output_proto, 'r') as already:
    already_header = already.read()
if already_header != proto_file:
  with open(output_proto, 'w') as out:
    out.write(proto_file)

