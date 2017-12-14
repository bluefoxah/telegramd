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
	"testing"
	"fmt"
	"github.com/disintegration/imaging"
	"bytes"
	"github.com/nebulaim/telegramd/base/mysql_client"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/golang/glog"
	"image"
)

func init()  {
	mysqlConfig := mysql_client.MySQLConfig{
		Name:   "immaster",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig})
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
}

func TestResize(t *testing.T) {
	id := int64(-8540733062663239681)
	filePartsDOList := dao.GetFilePartsDAO(dao.DB_MASTER).SelectFileParts(id)
	fileDatas := []byte{}
	for _, p := range filePartsDOList {
		fileDatas = append(fileDatas, p.Bytes...)
	}

	// bufio.Reader{}
	img, err := imaging.Decode(bytes.NewReader(fileDatas))
	if err != nil {
		glog.Error("Decode error: {%v}", err)
		return
	}

	imgSz := MakeResizeInfo(img)
	for i, sz := range sizeList {
		var dst *image.NRGBA
		if imgSz.isWidth {
			dst = imaging.Resize(img, sz, 0, imaging.Lanczos)
		} else {
			dst = imaging.Resize(img, 0, sz, imaging.Lanczos)
		}

		err := imaging.Save(dst, fmt.Sprintf("/tmp/telegramd/%d.jpg", i))
		if err != nil {
			glog.Error("Save error: {%v}", err)
		}
	}
}
