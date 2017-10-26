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

import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
)

type LangpackServiceImpl struct {
}

func (s *LangpackServiceImpl) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *LangpackServiceImpl) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *LangpackServiceImpl)LangpackGetStrings(ctx context.Context,  request *mtproto.TLLangpackGetStrings) (*mtproto.Vector<LangPackString>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

// func (s *LangpackServiceImpl)LangpackGetLanguages(ctx context.Context,  request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector<LangPackLanguage>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }
