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

type UploadServiceImpl struct {
}

func (s *UploadServiceImpl) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *UploadServiceImpl) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *UploadServiceImpl) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *UploadServiceImpl) UploadGetWebFile(ctx context.Context, request *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

func (s *UploadServiceImpl) UploadGetCdnFile(ctx context.Context, request *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// func (s *UploadServiceImpl)UploadReuploadCdnFile(ctx context.Context,  request *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector<CdnFileHash>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

// func (s *UploadServiceImpl)UploadGetCdnFileHashes(ctx context.Context,  request *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Vector<CdnFileHash>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }
