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

package rpc

/*
import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/nebulaim/telegramd/biz_server/langpack/model"
)

const (
	LANG_PACK_EN_FILE = "./lang_pack_en.toml"
)

var langs model.LangPacks

func init()  {
	if _, err := toml.DecodeFile(LANG_PACK_EN_FILE, &langs); err != nil {
		panic(err)
	}
}

type LangpackServiceImpl struct {
}

func (s *LangpackServiceImpl) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	glog.Infof("LangpackGetLangPack - Process: %v", request)


	if _, err := toml.DecodeFile(LANG_PACK_EN_FILE, &langs); err != nil {
		fmt.Errorf("%s\n", err)
		return nil, err
	}

	diff := &mtproto.TLLangPackDifference{}
	diff.LangCode = request.LangCode
	diff.Version = langs.Version
	diff.FromVersion = 0

	for _, strings := range langs.Strings {
		diff.Strings = append(diff.Strings, mtproto.MakeLangPackString(strings))
	}

	for _, stringPluralizeds := range langs.StringPluralizeds {
		diff.Strings = append(diff.Strings, mtproto.MakeLangPackString(stringPluralizeds))
	}

	reply := mtproto.MakeLangPackDifference(diff)
	glog.Infof("LangpackGetLangPack - reply: {%v}\n", reply)
	return reply, nil
}

func (s *LangpackServiceImpl) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	glog.Infof("LangpackGetDifference - Process: %v", request)

	diff := &mtproto.TLLangPackDifference{}
	diff.LangCode = "en"
	diff.Version = langs.Version
	diff.FromVersion = request.FromVersion

	if request.FromVersion < langs.Version {
		// TODO(@benqi): 找出不同版本的增量更新数据
	}

	reply := mtproto.MakeLangPackDifference(diff)
	glog.Infof("LangpackGetDifference - reply: {%v}\n", reply)
	return reply, nil
}

// func (s *LangpackServiceImpl)LangpackGetStrings(ctx context.Context,  request *mtproto.TLLangpackGetStrings) (*mtproto.Vector<LangPackString>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

// func (s *LangpackServiceImpl)LangpackGetLanguages(ctx context.Context,  request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector<LangPackLanguage>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }
*/