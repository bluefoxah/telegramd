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

func (m *photoModel) UploadPhoto(userId int32, id int64, parts int32, name, md5Checksum string) ([]*mtproto.PhotoSize, error) {
	sizes := make([]*mtproto.PhotoSize, 0, 4)

	// 图片压缩和处理
	// ext := filepath.Ext(name)

	filesDO := dao.GetFilesDAO(dao.DB_MASTER).SelectByIDAndParts(id, parts)
	if filesDO == nil {
		return nil, fmt.Errorf("File exists: id = %d, parts = %d", id, parts)
	}

	// check md5Checksum, big file's md5 is empty
	if md5Checksum != "" && md5Checksum != filesDO.Md5Checksum {
		return nil, fmt.Errorf("Invalid md5Checksum: md5Checksum = %s, but filesDO = {%v}", md5Checksum, filesDO)
	}

	// select file data
	filePartsDOList := dao.GetFilePartsDAO(dao.DB_MASTER).SelectFileParts(id)
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
			FileId: id,
			DcId: 2,
			VolumeId: vId,
			// LocalId: 12345,
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
		localId := int32(dao.GetPhotoDatasDAO(dao.DB_MASTER).Insert(photoDatasDO))

		photoSizeData := &mtproto.PhotoSize_Data{
			Type: getSizeType(i),
			W:    photoDatasDO.Width,
			H:    photoDatasDO.Height,
			Size: int32(len(photoDatasDO.Bytes)),
			Location: &mtproto.FileLocation{
				Constructor: mtproto.TLConstructor_CRC32_fileLocation,
				Data2: &mtproto.FileLocation_Data{
					VolumeId: photoDatasDO.VolumeId,
					LocalId:  localId,
					Secret:   id2.NextId(),
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