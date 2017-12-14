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
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"crypto/md5"
	"fmt"
)

const (
	maxFilePartSize = 32768
)

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveFilePart(ctx context.Context, request *mtproto.TLUploadSaveFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadSaveFilePart - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): 最简单的实现保证跑通流程
	filePartsDO := &dataobject.FilePartsDO{
	    CreatorUserId: md.UserId,
	    FileId: request.FileId,
	    FilePart: request.FilePart,
	    IsBigFile: 0,
	    Bytes: request.Bytes,
	}
	dao.GetFilePartsDAO(dao.DB_MASTER).Insert(filePartsDO)

	if len(request.Bytes) < maxFilePartSize {
		// 文件上传结束, 计算出文件大小和md5，存盘
		filePartsDOList := dao.GetFilePartsDAO(dao.DB_MASTER).SelectFileParts(request.FileId)
		// datas := make([]byte, 0, len(filePartsDOList)*maxFilePartSize)
		md5Hash := md5.New()
		fileSize := 0
		for _, v := range filePartsDOList {
			fileSize += len(v.Bytes)
			md5Hash.Write(v.Bytes)
		}

		filesDO := &dataobject.FilesDO{
			CreatorUserId: md.UserId,
			FileId: request.FileId,
			FileParts: int32(len(filePartsDOList)),
			FileSize: int64(fileSize),
			Md5Checksum: fmt.Sprintf("%x", md5Hash.Sum(nil)),
		}
		dao.GetFilesDAO(dao.DB_MASTER).Insert(filesDO)
	}

	glog.Infof("UploadSaveFilePart - reply: {true}")
	return mtproto.ToBool(true), nil
}
