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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/base/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz_model/model"
)

// upload.getFile#e3a6cfb5 location:InputFileLocation offset:int limit:int = upload.File;
func (s *UploadServiceImpl) UploadGetFile(ctx context.Context, request *mtproto.TLUploadGetFile) (*mtproto.Upload_File, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadGetFile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UploadGetFile logic
	// upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;

	// inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
	// inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
	// inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
	//uploadFile := mtproto.Upload_File{ Data2 :
	//}
	switch request.GetLocation().GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputFileLocation:
		inputFileLocation := request.GetLocation().To_InputFileLocation()
		uploadfile := model.GetPhotoModel().GetPhotoFileData(inputFileLocation.GetVolumeId(),
			inputFileLocation.GetLocalId(), inputFileLocation.GetSecret(), request.GetOffset(), request.GetLimit())

		glog.Infof("UploadGetFile - reply: %s", logger.JsonDebugData(uploadfile))
		return uploadfile, nil
	case mtproto.TLConstructor_CRC32_inputEncryptedFileLocation:
	case mtproto.TLConstructor_CRC32_inputDocumentFileLocation:
	default:
		glog.Errorf("Invalid InputFileLocation type!")
	}

	return nil, fmt.Errorf("Not impl UploadGetFile")
}
