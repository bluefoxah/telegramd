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
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
)

type UploadServiceImpl struct {
}

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	glog.Infof("UploadSaveFilePart - Process: file_id = %d, file_part = %d", request.FileId, request.FilePart)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	filePartsDO := &dataobject.FilePartsDO{
		CreatorUserId: md.UserId,
		FileId: request.FileId,
		FilePart: request.FilePart,
		IsBigFile: 0,
		Bytes: request.Bytes,
	}
	dao.GetFilePartsDAO(dao.DB_MASTER).Insert(filePartsDO)

	glog.Infof("UploadSaveFilePart - reply: {true}")
	return mtproto.ToBool(true), nil
}

// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	glog.Infof("UploadSaveBigFilePart - Process: file_id = %d, file_part = %d", request.FileId, request.FilePart)

	md := grpc_util.RpcMetadataFromIncoming(ctx)
	filePartsDO := &dataobject.FilePartsDO{
		CreatorUserId: md.UserId,
		FileId: request.FileId,
		FilePart: request.FilePart,
		IsBigFile: 1,
		FileTotalParts: request.FileTotalParts,
		Bytes: request.Bytes,
	}
	dao.GetFilePartsDAO(dao.DB_MASTER).Insert(filePartsDO)

	glog.Infof("UploadSaveBigFilePart - reply: {true}")
	return mtproto.ToBool(true), nil
}


// upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
func (s *UploadServiceImpl) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// upload.getWebFile#24e6818d location:InputWebFileLocation offset:int limit:int = upload.WebFile;
func (s *UploadServiceImpl) UploadGetWebFile(ctx context.Context, request *mtproto.TLUploadGetWebFile) (*mtproto.Upload_WebFile, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// upload.getCdnFile#2000bcc3 file_token:bytes offset:int limit:int = upload.CdnFile;
func (s *UploadServiceImpl) UploadGetCdnFile(ctx context.Context, request *mtproto.TLUploadGetCdnFile) (*mtproto.Upload_CdnFile, error) {
	glog.Info("Process: %v", request)
	return nil, nil
}

// upload.reuploadCdnFile#1af91c09 file_token:bytes request_token:bytes = Vector<CdnFileHash>;
// func (s *UploadServiceImpl)UploadReuploadCdnFile(ctx context.Context,  request *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector<CdnFileHash>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }

// upload.getCdnFileHashes#f715c87b file_token:bytes offset:int = Vector<CdnFileHash>;
// func (s *UploadServiceImpl)UploadGetCdnFileHashes(ctx context.Context,  request *mtproto.TLUploadGetCdnFileHashes) (*mtproto.Vector<CdnFileHash>, error) {
//   glog.Info("Process: %v", request)
//   return nil, nil
// }
*/
