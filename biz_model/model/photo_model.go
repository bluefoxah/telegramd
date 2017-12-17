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

package model

import (
	"sync"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	// "github.com/cosiner/gohper/errors"
	"github.com/disintegration/imaging"
	"bytes"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
	"image"
	"github.com/nebulaim/telegramd/base/base"
	"fmt"
	"github.com/golang/glog"
	id2 "github.com/nebulaim/telegramd/frontend/id"
	"time"
)

const (
	PHOTO_SIZE_SMALL_TYPE	 	= "s"
	PHOTO_SIZE_MEDIUMN_TYPE 	= "m"
	PHOTO_SIZE_XLARGE_TYPE 		= "x"
	PHOTO_SIZE_YLARGE_TYPE 		= "y"

	PHOTO_SIZE_SMALL_SIZE	 	= 90
	PHOTO_SIZE_MEDIUMN_SIZE 	= 320
	PHOTO_SIZE_XLARGE_SIZE 		= 800
	PHOTO_SIZE_YLARGE_SIZE 		= 1280
)

/*
	storage.fileUnknown#aa963b05 = storage.FileType;
	storage.filePartial#40bc6f52 = storage.FileType;
	storage.fileJpeg#7efe0e = storage.FileType;
	storage.fileGif#cae1aadf = storage.FileType;
	storage.filePng#a4f63c0 = storage.FileType;
	storage.filePdf#ae1e508d = storage.FileType;
	storage.fileMp3#528a0677 = storage.FileType;
	storage.fileMov#4b09ebbc = storage.FileType;
	storage.fileMp4#b3cea0e4 = storage.FileType;
	storage.fileWebp#1081464c = storage.FileType;
 */

type photoModel struct {
	// usersDAO *dao.UsersDAO
}

var sizeList = []int{PHOTO_SIZE_SMALL_SIZE, PHOTO_SIZE_MEDIUMN_SIZE, PHOTO_SIZE_XLARGE_SIZE, PHOTO_SIZE_YLARGE_SIZE}
func getSizeType(idx int) string {
	switch idx {
	case 0:
		return PHOTO_SIZE_SMALL_TYPE
	case 1:
		return PHOTO_SIZE_MEDIUMN_TYPE
	case 2:
		return PHOTO_SIZE_XLARGE_TYPE
	default:
		return PHOTO_SIZE_YLARGE_TYPE
	}
}

var (
	photoInstance *photoModel
	photoInstanceOnce sync.Once
)

func GetPhotoModel() *photoModel {
	photoInstanceOnce.Do(func() {
		photoInstance = &photoModel{}
	})
	return photoInstance
}

type resizeInfo struct {
	isWidth bool
	size int
}

func MakeResizeInfo(img image.Image) resizeInfo {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	if w >= h {
		return resizeInfo{
			isWidth: true,
			size: w,
		}
	} else {
		return resizeInfo{
			isWidth: false,
			size : h,
		}
	}
}

// TODO: @benqi
// 	我们未来的图片存储系统可能会按facebook的Haystack论文来实现
// 	mtproto协议也定义了一套自己的文件存储方案，fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
// 	在这里，我们重新定义mtproto的volume_id和local_id，对应Haystack的key和alternate_key，secret对应cookie
//  在当前简单实现里，volume_id由sonwflake生成，local_id对应于图片类型，secret为access_hash
func (m *photoModel) UploadPhoto(userId int32, photoId, fileId int64, parts int32, name, md5Checksum string) ([]*mtproto.PhotoSize, error) {
	sizes := make([]*mtproto.PhotoSize, 0, 4)

	// 图片压缩和处理
	// ext := filepath.Ext(name)

	filesDO := dao.GetFilesDAO(dao.DB_MASTER).SelectByIDAndParts(fileId, parts)
	if filesDO == nil {
		return nil, fmt.Errorf("File exists: id = %d, parts = %d", fileId, parts)
	}

	// check md5Checksum, big file's md5 is empty
	if md5Checksum != "" && md5Checksum != filesDO.Md5Checksum {
		return nil, fmt.Errorf("Invalid md5Checksum: md5Checksum = %s, but filesDO = {%v}", md5Checksum, filesDO)
	}

	// select file data
	filePartsDOList := dao.GetFilePartsDAO(dao.DB_MASTER).SelectFileParts(fileId)
	fileDatas := []byte{}

	for _, p := range filePartsDOList {
		fileDatas = append(fileDatas, p.Bytes...)
	}

	// bufio.Reader{}
	img, err := imaging.Decode(bytes.NewReader(fileDatas))
	if err != nil {
		glog.Errorf("Decode error: {%v}", err)
		return nil, err
	}

	imgSz := MakeResizeInfo(img)

	vId := id2.NextId()
	for i, sz := range sizeList {
		photoDatasDO := &dataobject.PhotoDatasDO{
			PhotoId: photoId,
			DcId: 2,
			VolumeId: vId,
			LocalId: int32(i),
			AccessHash: id2.NextId(),
		}

		var dst *image.NRGBA
		if imgSz.isWidth {
			dst = imaging.Resize(img, sz, 0, imaging.Lanczos)
		} else {
			dst = imaging.Resize(img, 0, sz, imaging.Lanczos)
		}

		photoDatasDO.Width = int32(dst.Bounds().Dx())
		photoDatasDO.Height = int32(dst.Bounds().Dy())
		imgBuf := base.MakeBuffer(0, len(fileDatas))
		err = imaging.Encode(imgBuf, dst, imaging.JPEG)
		if err != nil {
			glog.Errorf("Encode error: {%v}", err)
			return nil, err
		}

		photoDatasDO.Bytes = imgBuf.Bytes()
		dao.GetPhotoDatasDAO(dao.DB_MASTER).Insert(photoDatasDO)

		photoSizeData := &mtproto.PhotoSize_Data{
			Type: getSizeType(i),
			W:    photoDatasDO.Width,
			H:    photoDatasDO.Height,
			Size: int32(len(photoDatasDO.Bytes)),
			Location: &mtproto.FileLocation{
				Constructor: mtproto.TLConstructor_CRC32_fileLocation,
				Data2: &mtproto.FileLocation_Data{
					VolumeId: photoDatasDO.VolumeId,
					LocalId:  int32(i),
					Secret:   photoDatasDO.AccessHash,
					DcId: 	photoDatasDO.DcId}}}

		if i== 0 {
			sizes = append(sizes, &mtproto.PhotoSize{
				Constructor: mtproto.TLConstructor_CRC32_photoCachedSize,
				Data2:       photoSizeData,})
			photoSizeData.Bytes = photoDatasDO.Bytes
		} else {
			sizes = append(sizes, &mtproto.PhotoSize{
				Constructor: mtproto.TLConstructor_CRC32_photoSize,
				Data2:       photoSizeData,})
		}
	}

	return sizes, nil
}

func (m *photoModel) GetPhotoFileData(volumeId int64, localId int32, secret int64, offset int32, limit int32) *mtproto.Upload_File {
	// inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
	photoDatasDO := dao.GetPhotoDatasDAO(dao.DB_MASTER).SelectByFileLocation(volumeId, localId, secret)
	if photoDatasDO != nil {
		if offset <= int32(len(photoDatasDO.Bytes)) {
			uploadFile := mtproto.NewTLUploadFile()
			uploadFile.SetType(mtproto.NewTLStorageFileJpeg().To_Storage_FileType())
			if offset + limit > int32(len(photoDatasDO.Bytes)) {
				limit = int32(len(photoDatasDO.Bytes)) - offset
			}
			uploadFile.SetBytes(photoDatasDO.Bytes[offset:offset+limit])
			uploadFile.SetMtime(int32(time.Now().Unix()))
			return uploadFile.To_Upload_File()
		}
	} else {
		glog.Errorf("GetPhotoDatasDAO nil")
	}
	return nil
}
