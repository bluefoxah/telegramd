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

output_proto = output_path # + '/codec_schema.tl.pb.go'

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
resClassTypesTexts2 = '';
resClassTypesTexts3 = '';
classFuncsTexts = '';
registers = [];

registers.append('  int32(TLConstructor_CRC32_message2) : func() (TLObject) { return new(TLMessage2) },\n');
registers.append('  int32(TLConstructor_CRC32_msg_container) : func() (TLObject) { return new(TLMsgContainer) },\n');
registers.append('  int32(TLConstructor_CRC32_msg_copy) : func() (TLObject) { return new(TLMsgCopy) },\n');
registers.append('  int32(TLConstructor_CRC32_gzip_packed) : func() (TLObject) { return new(TLGzipPacked) },\n');
registers.append('  int32(TLConstructor_CRC32_rpc_result) : func() (TLObject) { return new(TLRpcResult) },\n');

with open(input_file) as f:
  for line in f:
    line=line.strip('\n')
    layerline = re.match(r'// LAYER (\d+)', line)
    if (layerline):
      # 当前层
      layer = 'const CURRENT_LAYER = ' + layerline.group(1);

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
    registers.append('  int32(TLConstructor_CRC32_' + name + ') : func() (TLObject) { return new(TL' + to_proto_go_name(name) + ') },\n');

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
        ptype2 = 'int128';
      if (ptype2 == 'int256'):
        ptype2 = 'int256';
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

for restype in typesList:
  v = typesDict[restype];
  resType = TypesDict[restype];

  # print('restype:' + restype + ', resType: ' + resType);

  withData = 0;
  creatorsText = '';
  constructsText = '';
  constructsInline = '';

  withType = (len(v) > 1);
  switchLines = '';
  friendDecl = '';
  getters = '';
  reader = '';
  writer = '';
  sizeList = [];
  sizeFast = '';
  newFast = '';
  sizeCases = '';
  #  print(v)

  resClassTypesTexts2 = ''
  resClassTypesTexts3 = ''

  resClassTypesTexts += 'func (m *' + to_proto_go_name(resType) + ') Encode() (b []byte) {\n'
  resClassTypesTexts += '  b = nil\n'
  resClassTypesTexts += '  switch m.Payload.(type) {\n';

  resClassTypesTexts2 += 'func (m *' + to_proto_go_name(resType) + ') Decode(dbuf *DecodeBuf) error {\n'
  resClassTypesTexts2 += '  classId := dbuf.Int()\n'
  resClassTypesTexts2 += '  switch classId {\n'

  resClassTypesTexts3 += 'func Make' + to_proto_go_name(resType) + '(message proto.Message) (m *' + to_proto_go_name(resType) + ') {\n'
  resClassTypesTexts3 += '  switch message.(type) {\n'

  #### resClassTypesTexts += '  oneof payload {\n';

  ij = 1;
  for data in v:
    name = data[0];
    typeid = data[1];
    prmsList = data[2];
    prms = data[3];
    hasFlags = data[4];
    conditionsList = data[5];
    conditions = data[6];
    trivialConditions = data[7];
    line = data[8]

    classTypesTexts += '// ' + line + '\n';

    # classTypesTexts += 'message Z' + name + ' : public ' + resType + ' {\n'; # type class declaration
    classTypesTexts += 'func (m* TL' + to_proto_go_name(name) + ') Encode() []byte {\n'; # type class declaration
    classTypesTexts += '  x := NewEncodeBuf(512)\n'; # type class declaration
    classTypesTexts += '  x.Int(int32(TLConstructor_CRC32_' + name + '))\n'

    #### resClassTypesTexts += '    TL_' + name + ' ' + name + ' = ' + str(ij) + ';\n'
    resClassTypesTexts += '    case *' + to_proto_go_name(resType) + '_' + to_proto_go_name(name) +':\n'
    resClassTypesTexts += '      m2, _ := m.Payload.(*' + to_proto_go_name(resType) + '_' + to_proto_go_name(name) + ')\n'
    resClassTypesTexts += '      b = m2.' + to_proto_go_name(name) + '.Encode()\n'

    resClassTypesTexts2 += '    case int32(TLConstructor_CRC32_' + name + '):\n'
    resClassTypesTexts2 += '      m2 := ' + to_proto_go_name(resType) + '_' + to_proto_go_name(name) + '{}\n'
    resClassTypesTexts2 += '      m2.' + to_proto_go_name(name) + ' = &TL' + to_proto_go_name(name) + '{}\n'
    resClassTypesTexts2 += '      m2.' + to_proto_go_name(name) + '.Decode(dbuf)\n'
    resClassTypesTexts2 += '      m.Payload = &m2\n'

    resClassTypesTexts3 += '    case *TL' + to_proto_go_name(name) +':\n'
    resClassTypesTexts3 += '      m2, _ := message.(*TL' + to_proto_go_name(name) + ')\n'
    resClassTypesTexts3 += '      m = &' + to_proto_go_name(resType) + '{\n'
    resClassTypesTexts3 += '         Payload: &' + to_proto_go_name(resType) + '_' + to_proto_go_name(name) + '{\n'
    resClassTypesTexts3 += '           ' + to_proto_go_name(name) + ': m2,\n'
    resClassTypesTexts3 += '         },\n'
    resClassTypesTexts3 += '      }\n'

    ij += 1;

    ii = 1;

    if (hasFlags):
      classTypesTexts += '\n  var flags uint32 = 0\n';

  ## Encode()

    for paramName in prmsList:
      paramType = prms[paramName];

      if (paramName in conditionsList):
        if (paramType in ['bool']):
          # print '  if m.' + to_proto_go_name(paramName) + ' == true {'
          classTypesTexts += '  if m.' + to_proto_go_name(paramName) + ' == true {\n';
        elif (paramType in ['int32', 'int64']):
          #print '  if  m.' + to_proto_go_name(paramName) + ' != 0 {'
          classTypesTexts += '  if  m.' + to_proto_go_name(paramName) + ' != 0 {\n';
        elif (paramType in ['string']):
          #print '  if  m.' + to_proto_go_name(paramName) + ' != "" {'
          classTypesTexts += '  if  m.' + to_proto_go_name(paramName) + ' != "" {\n';
        else:
          #print '  if m.' + to_proto_go_name(paramName) + ' != nil {'
          classTypesTexts += '  if m.' + to_proto_go_name(paramName) + ' != nil {\n';

        classTypesTexts += '    flags |= 1<<' + conditions[paramName] + '\n';
        classTypesTexts += '  }\n';

    if (hasFlags):
      classTypesTexts += '  x.UInt(flags)\n\n'

    for paramName in prmsList:
      if (paramName == 'flags'):
        continue;

      paramType = prms[paramName];

      if (paramName in conditionsList):
        classTypesTexts += '  if (flags & (1 << ' + conditions[paramName] + ')) != 0 {\n';

      if (paramType == 'bool'):
        classTypesTexts += '  // ignore\n';
        # classTypesTexts += '  if (flags & (1 << ' + conditions[paramName] + ')) != 0 {\n';

      elif (paramType =='int32'):
        classTypesTexts += '  x.Int(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'int64'):
        classTypesTexts += '  x.Long(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'double'):
        classTypesTexts += '  x.Double(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'string'):
        classTypesTexts += '  x.String(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'int128'):
        classTypesTexts += '  x.Bytes(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'int256'):
        classTypesTexts += '  x.Bytes(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'bytes'):
        classTypesTexts += '  x.StringBytes(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'repeated int32'):
        classTypesTexts += '  x.VectorInt(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'repeated int64'):
        classTypesTexts += '  x.VectorLong(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'repeated string'):
        classTypesTexts += '  x.VectorString(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType in TypesList):
        classTypesTexts += '  x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';
      elif (paramType.find('std::vector') >= 0):
        eptype = txt_wrap_by('<', '*', paramType);
        classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ');\n';

        classTypesTexts += '  // x%d := make([]byte, 8)\n' % (ii)
        classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d, uint32(TLConstructor_CRC32_vector))\n' % (ii)
        classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d[4:], uint32(len(m.%s)))\n' % (ii, to_proto_go_name(paramName))
        classTypesTexts += '  x.Int(int32(len(m.%s)))\n' % (to_proto_go_name(paramName))
        classTypesTexts += '  for _, v := range m.' + to_proto_go_name(paramName) + ' {\n'
        classTypesTexts += '     x.buf = append(x.buf, (*v).Encode()...)\n'
        classTypesTexts += '  }\n'

      elif (paramType.find('TLObjectVector') >= 0):
        eptype = txt_wrap_by('<', '>', paramType);
        classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ');\n';

        classTypesTexts += '  // x%d := make([]byte, 8)\n' % (ii)
        classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d, uint32(TLConstructor_CRC32_vector))\n' % (ii)
        classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d[4:], uint32(len(m.%s)))\n' % (ii, to_proto_go_name(paramName))
        classTypesTexts += '  x.Int(int32(TLConstructor_CRC32_vector))\n'
        classTypesTexts += '  x.Int(int32(len(m.%s)))\n' % (to_proto_go_name(paramName))
        classTypesTexts += '  for _, v := range m.' + to_proto_go_name(paramName) + ' {\n'
        classTypesTexts += '     x.buf = append(x.buf, (*v).Encode()...)\n'
        classTypesTexts += '  }\n'

      else:
        # classTypesTexts += '  // 2. ' + paramType + ' ' + paramName + ' = ' + str(ii) + ';\n';
        classTypesTexts += '  x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';

      if (paramName in conditionsList):
        classTypesTexts += '  }\n';

      ii += 1;

    classTypesTexts += '  return x.buf\n'
    classTypesTexts += '}\n\n';


    ## Decode()
    # classTypesTexts += 'message Z' + name + ' : public ' + resType + ' {\n'; # type class declaration
    classTypesTexts += 'func (m* TL' + to_proto_go_name(name) + ') Decode(dbuf *DecodeBuf) error {\n'; # type class declaration
    # classTypesTexts += '  x.Int(int32(TLConstructor_CRC32_' + name + '))\n'

    ii = 1;
    if (hasFlags):
      classTypesTexts += '  flags := dbuf.UInt()\n'
      if (name=='messages_channelMessages'):
        classTypesTexts += '  if flags != 0 {}\n'

    for paramName in prmsList:
      if (paramName == 'flags'):
        continue;

      paramType = prms[paramName];

      if (paramName in conditionsList):
        classTypesTexts += '  if (flags & (1 << ' + conditions[paramName] + ')) != 0 {\n';

      if (paramType == 'bool'):
        if (paramName in conditionsList):
          # classTypesTexts += '';
          classTypesTexts += '    m.' + to_proto_go_name(paramName) + ' = true\n';

          # classTypesTexts += '  bool ' + paramName + ' = ' + str(ii) + ';\n';
      elif (paramType =='int32'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Int()\n';
      elif (paramType == 'int64'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Long()\n';
      elif (paramType == 'double'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Double()\n';
      elif (paramType == 'int128'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Bytes(16)\n';
      elif (paramType == 'int256'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Bytes(32)\n';
      elif (paramType == 'string'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.String()\n';
      elif (paramType == 'bytes'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.StringBytes()\n';
      elif (paramType == 'repeated int32'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.VectorInt()\n';
      elif (paramType == 'repeated int64'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.VectorLong()\n';
      elif (paramType == 'repeated string'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.VectorString()\n';
      elif (paramType in TypesList):
        classTypesTexts += '  // dbuf.Int()\n'
        classTypesTexts += '  m.' + to_proto_go_name(paramName) + ' = &' + to_proto_go_name(paramType) + '{}\n';
        classTypesTexts += '  (*m.' + to_proto_go_name(paramName) + ').Decode(dbuf)\n';

      elif (paramType.find('std::vector') >= 0):
        eptype = txt_wrap_by('<', '*', paramType);
        classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ');\n';

        classTypesTexts += '  // c%d := dbuf.Int()\n' % (ii)
        classTypesTexts += '  // if c%d != int32(TLConstructor_CRC32_vector) {\n' % (ii)
        classTypesTexts += '  //   return fmt.Errorf("Not vector, classID: ", c%d)\n' % (ii)
        classTypesTexts += '  // }\n'
        classTypesTexts += '  l%d := dbuf.Int()\n' % (ii)
        classTypesTexts += '  m.%s = make([]*%s, l%d)\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype), ii)


        classTypesTexts += '  for i := 0; i < int(l%d); i++ {\n' % (ii)
        classTypesTexts += '    m.%s[i] = &%s{}\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype))
        if (eptype in TypesList):
          classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
        else:
          classTypesTexts += '    dbuf.Int()\n'
          classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
          classTypesTexts += '    // TODO(@benqi): Check classID valid!!!\n'
          classTypesTexts += '    // dbuf.Int()\n'
        classTypesTexts += '  }\n'

      elif (paramType.find('TLObjectVector') >= 0):
        eptype = txt_wrap_by('<', '>', paramType);
        classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ');\n';
        classTypesTexts += '  //  dbuf.Int()\n'
        classTypesTexts += '  c%d := dbuf.Int()\n' % (ii)
        classTypesTexts += '  if c%d != int32(TLConstructor_CRC32_vector) {\n' % (ii)
        classTypesTexts += '    return fmt.Errorf("Not vector, classID: ", c%d)\n' % (ii)
        classTypesTexts += '  }\n'
        classTypesTexts += '  l%d := dbuf.Int()\n' % (ii)
        classTypesTexts += '  m.%s = make([]*%s, l%d)\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype), ii)

        classTypesTexts += '  for i := 0; i < int(l%d); i++ {\n' % (ii)
        classTypesTexts += '    m.%s[i] = &%s{}\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype))
        if (eptype in TypesList):
          classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
        else:
          classTypesTexts += '    dbuf.Int()\n'
          classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
          classTypesTexts += '    // TODO(@benqi): Check classID valid!!!\n'
          classTypesTexts += '    // dbuf.Int()\n'
        classTypesTexts += '  }\n'

      else:
        # classTypesTexts += '  // 2. ' + paramType + ' ' + paramName + ' = ' + str(ii) + ';\n';
        classTypesTexts += '  // other!!!! x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';

      if (paramName in conditionsList):
        classTypesTexts += '  }\n';

      ii += 1;

    classTypesTexts += '  return dbuf.err\n'
    classTypesTexts += '}\n\n';

  resClassTypesTexts += '  }\n'
  resClassTypesTexts += '  return\n'
  resClassTypesTexts += '}\n\n';

  resClassTypesTexts2 += '  }\n'
  resClassTypesTexts2 += '  return dbuf.err\n'
  resClassTypesTexts2 += '}\n\n'

  resClassTypesTexts3 += '  }\n'
  resClassTypesTexts3 += '  return\n'
  resClassTypesTexts3 += '}\n\n'

  resClassTypesTexts = resClassTypesTexts + resClassTypesTexts2 + resClassTypesTexts3

classTypesTexts += '\n// RPC\n';
for restype in funcsList:
  v = funcsDict[restype];
  for data in v:
    name = data[0];
    typeid = data[1];
    prmsList = data[2];
    prms = data[3];
    hasFlags = data[4];
    conditionsList = data[5];
    conditions = data[6];
    trivialConditions = data[7];
    line = data[8]

    classTypesTexts += '// ' + line + '\n'; # type class declaration

    classTypesTexts += 'func (m* TL' + to_proto_go_name(name) + ') Encode() []byte {\n'; # type class declaration
    classTypesTexts += '  x := NewEncodeBuf(512)\n'; # type class declaration
    classTypesTexts += '  x.Int(int32(TLConstructor_CRC32_' + name + '))\n'

    ii = 1;

    if (hasFlags):
      classTypesTexts += '\n  var flags uint32 = 0\n';

    for paramName in prmsList:
      paramType = prms[paramName];
      if (paramName in conditionsList):
        if (paramType in ['bool']):
          #print '  if m.' + to_proto_go_name(paramName) + ' == true {'
          classTypesTexts += '  if m.' + to_proto_go_name(paramName) + ' == true {\n';
        elif (paramType in ['int32', 'int64']):
          #print '  if  m.' + to_proto_go_name(paramName) + ' != 0 {'
          classTypesTexts += '  if  m.' + to_proto_go_name(paramName) + ' != 0 {\n';
        elif (paramType in ['string']):
          #print '  if  m.' + to_proto_go_name(paramName) + ' != "" {'
          classTypesTexts += '  if  m.' + to_proto_go_name(paramName) + ' != "" {\n';
        else:
          #print '  if m.' + to_proto_go_name(paramName) + ' != nil {'
          classTypesTexts += '  if m.' + to_proto_go_name(paramName) + ' != nil {\n';

        classTypesTexts += '    flags |= 1<<' + conditions[paramName] + '\n';
        classTypesTexts += '  }\n';

    if (hasFlags):
      classTypesTexts += '  x.UInt(flags)\n\n'

    for paramName in prmsList:
      if (paramName == 'flags'):
        continue;

      paramType = prms[paramName];

      if (paramName in conditionsList):
        classTypesTexts += '  if (flags & (1 << ' + conditions[paramName] + ')) != 0 {\n';

      if (paramType == 'bool'):
        if (paramName in conditionsList):
          # classTypesTexts += '';
          classTypesTexts += '   //  m.' + to_proto_go_name(paramName) + ' = true\n';
      elif (paramType =='int32'):
        classTypesTexts += '  x.Int(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'int64'):
        classTypesTexts += '  x.Long(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'double'):
        classTypesTexts += '  x.Double(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'int128'):
        classTypesTexts += '  x.Bytes(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'int256'):
        classTypesTexts += '  x.Bytes(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'string'):
        classTypesTexts += '  x.String(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'bytes'):
        classTypesTexts += '  x.StringBytes(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'repeated int32'):
        classTypesTexts += '  x.VectorInt(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'repeated int64'):
        classTypesTexts += '  x.VectorLong(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType == 'repeated string'):
        classTypesTexts += '  x.VectorString(m.' +  to_proto_go_name(paramName) + ')\n';
      elif (paramType in TypesList):
        classTypesTexts += '  x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';
      else:
        if (paramType == 'TQueryType'):
          # classTypesTexts += '  bytes ' + paramName + ' = ' + str(ii) + ';\n';
          classTypesTexts += '  x.Bytes(m.' +  to_proto_go_name(paramName) + ')\n';
        elif (paramType.find('std::vector') >= 0):
          eptype = txt_wrap_by('<', '*', paramType);
          classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ')\n';

          classTypesTexts += '  // x%d := make([]byte, 8)\n' % (ii)
          classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d, uint32(TLConstructor_CRC32_vector))\n' % (ii)
          classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d[4:], uint32(len(m.%s)))\n' % (ii, to_proto_go_name(paramName))
          classTypesTexts += '  x.Int(int32(len(m.%s)))\n' % (to_proto_go_name(paramName))
          classTypesTexts += '  for _, v := range m.' + to_proto_go_name(paramName) + ' {\n'
          classTypesTexts += '     x.buf = append(x.buf, (*v).Encode()...)\n'
          classTypesTexts += '  }\n'

        elif (paramType.find('TLObjectVector') >= 0):
          eptype = txt_wrap_by('<', '>', paramType);

          classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ')\n';

          classTypesTexts += '  // x%d := make([]byte, 8)\n' % (ii)
          classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d, uint32(TLConstructor_CRC32_vector))\n' % (ii)
          classTypesTexts += '  // binary.LittleEndian.PutUint32(x%d[4:], uint32(len(m.%s)))\n' % (ii, to_proto_go_name(paramName))
          classTypesTexts += '  x.Int(int32(TLConstructor_CRC32_vector))\n'
          classTypesTexts += '  x.Int(int32(len(m.%s)))\n' % (to_proto_go_name(paramName))
          classTypesTexts += '  for _, v := range m.' + to_proto_go_name(paramName) + ' {\n'
          classTypesTexts += '     x.buf = append(x.buf, (*v).Encode()...)\n'
          classTypesTexts += '  }\n'
        else:
          # classTypesTexts += '  // 2. ' + paramType + ' ' + paramName + ' = ' + str(ii) + ';\n';
          classTypesTexts += '  x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';

      if (paramName in conditionsList):
        classTypesTexts += '  }\n';

      ii += 1;

    classTypesTexts += '  return x.buf\n'
    classTypesTexts += '}\n\n';

    classTypesTexts += 'func (m* TL' + to_proto_go_name(name) + ') Decode(dbuf *DecodeBuf) error {\n'; # type class declaration

    if (hasFlags):
      classTypesTexts += '  flags := dbuf.UInt()\n'
      if (name=='messages_channelMessages'):
        classTypesTexts += '  if flags != 0 {}\n'

    ii = 1;
    for paramName in prmsList:
      paramType = prms[paramName];
      if (paramName == 'flags'):
        continue;

      if (paramName in conditionsList):
        classTypesTexts += '  if (flags & (1 << ' + conditions[paramName] + ')) != 0 {\n';

      if (paramType == 'bool'):
        if (paramName in conditionsList):
          # classTypesTexts += '';
          classTypesTexts += '    m.' + to_proto_go_name(paramName) + ' = true\n';
      elif (paramType =='int32'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Int()\n';
      elif (paramType == 'int64'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Long() \n';
      elif (paramType == 'double'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Double()\n';
      elif (paramType == 'int128'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Bytes(16)\n';
      elif (paramType == 'int256'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.Bytes(32)\n';
      elif (paramType == 'string'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.String()\n';
      elif (paramType == 'bytes'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.StringBytes()\n';
      elif (paramType == 'repeated int32'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.VectorInt()\n';
      elif (paramType == 'repeated int64'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.VectorLong()\n';
      elif (paramType == 'repeated string'):
        classTypesTexts += '  m.' +  to_proto_go_name(paramName) + ' = dbuf.VectorString()\n';
      elif (paramType in TypesList):
        # classTypesTexts += '  // x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';
        classTypesTexts += '  // dbuf.Int()\n'
        classTypesTexts += '  m.' + to_proto_go_name(paramName) + ' = &' + to_proto_go_name(paramType) + '{}\n';
        classTypesTexts += '  (*m.' + to_proto_go_name(paramName) + ').Decode(dbuf)\n';

      else:
        if (paramType == 'TQueryType'):
          # classTypesTexts += '  bytes ' + paramName + ' = ' + str(ii) + ';\n';
          classTypesTexts += '  // TODO(@benqi): 暂时这么做，估计还是使用Any类型比较好\n'
          classTypesTexts += '  o%d := dbuf.Object()\n' % (ii)
          classTypesTexts += '  m.' + to_proto_go_name(paramName) + ' = o%d.Encode()\n' % (ii)
        elif (paramType.find('std::vector') >= 0):
          eptype = txt_wrap_by('<', '*', paramType);
          classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ')\n';

          classTypesTexts += '  c%d := dbuf.Int()\n' % (ii)
          classTypesTexts += '  // if c%d != int32(TLConstructor_CRC32_vector) {\n' % (ii)
          classTypesTexts += '  //   return fmt.Errorf("Not vector, classID: ", c%d)\n' % (ii)
          classTypesTexts += '  // }\n'
          classTypesTexts += '  l%d := dbuf.Int()\n' % (ii)
          classTypesTexts += '  m.%s = make([]*%s, l%d)\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype), ii)

          classTypesTexts += '  for i := 0; i < int(l%d); i++ {\n' % (ii)
          classTypesTexts += '    m.%s[i] = &%s{}\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype))
          if (eptype in TypesList):
            classTypesTexts += '   (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
          else:
            classTypesTexts += '    dbuf.Int()\n'
            classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
            classTypesTexts += '    // TODO(@benqi): Check classID valid!!!\n'
            classTypesTexts += '    // dbuf.Int()\n'
          classTypesTexts += '  }\n'

        elif (paramType.find('TLObjectVector') >= 0):
          eptype = txt_wrap_by('<', '>', paramType);

          classTypesTexts += '  // x.VectorMessage(m.' + to_proto_go_name(paramName) + ')\n';

          classTypesTexts += '  c%d := dbuf.Int()\n' % (ii)
          classTypesTexts += '  if c%d != int32(TLConstructor_CRC32_vector) {\n' % (ii)
          classTypesTexts += '    return fmt.Errorf("Not vector, classID: ", c%d)\n' % (ii)
          classTypesTexts += '  }\n'
          classTypesTexts += '  l%d := dbuf.Int()\n' % (ii)
          classTypesTexts += '  m.%s = make([]*%s, l%d)\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype), ii)

          classTypesTexts += '  for i := 0; i < int(l%d); i++ {\n' % (ii)
          classTypesTexts += '    m.%s[i] = &%s{}\n' % (to_proto_go_name(paramName), to_proto_go_name(eptype))
          if (eptype in TypesList):
            classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
          else:
            classTypesTexts += '    dbuf.Int()\n'
            classTypesTexts += '    (*m.%s[i]).Decode(dbuf)\n' % (to_proto_go_name(paramName))
            classTypesTexts += '    // TODO(@benqi): Check classID valid!!!\n'
            classTypesTexts += '    // dbuf.Int()\n'
          classTypesTexts += '  }\n'

        else:
          # classTypesTexts += '  // 2. ' + paramType + ' ' + paramName + ' = ' + str(ii) + ';\n';
          classTypesTexts += '  // other!!!! x.Bytes(m.' + to_proto_go_name(paramName) + '.Encode())\n';

      if (paramName in conditionsList):
        classTypesTexts += '  }\n';

      ii += 1;

    classTypesTexts += '  return dbuf.err\n'
    classTypesTexts += '}\n\n';

proto_file = '\
/*\n\
 * WARNING! All changes made in this file will be lost!\n\
 * Created from \'scheme.tl\' by \'codegen_encode_decode.py\'\n\
 *\n\
 *  Copyright (c) 2017, https://github.com/nebulaim\n\
 *  All rights reserved.\n\
 *\n\
 * Licensed under the Apache License, Version 2.0 (the "License");\n\
 * you may not use this file except in compliance with the License.\n\
 * You may obtain a copy of the License at\n\
 *\n\
 *   http://www.apache.org/licenses/LICENSE-2.0\n\
 *\n\
 * Unless required by applicable law or agreed to in writing, software\n\
 * distributed under the License is distributed on an "AS IS" BASIS,\n\
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n\
 * See the License for the specific language governing permissions and\n\
 * limitations under the License.\n\
 */\n\n\
package mtproto\n\n\
import ( \n\
  // "encoding/binary" \n\
  "fmt" \n\
  "github.com/golang/protobuf/proto"\n\
)\n\n\
type newTLObjectFunc func() TLObject\n\n\
var registers2 = map[int32]newTLObjectFunc {\n\
' + ''.join(registers) + '}\n\n\
func NewTLObjectByClassID(classId int32) TLObject { \n\
  m, ok := registers2[classId]\n\
  if !ok {\n\
    return nil\n\
  }\n\
  return m()\n\
}\n\n\
//////////////////////////////////////////////////////////////////////////////////////////\n\
' + resClassTypesTexts + '\n\n\
' + classTypesTexts + '\n'

already_header = ''
if os.path.isfile(proto_file):
  with open(output_proto, 'r') as already:
    already_header = already.read()
if already_header != proto_file:
  with open(output_proto, 'w') as out:
    out.write(proto_file)
