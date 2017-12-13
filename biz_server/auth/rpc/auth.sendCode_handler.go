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
	"github.com/nebulaim/telegramd/frontend/id"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"github.com/ttacon/libphonenumber"
	"golang.org/x/net/context"
	"time"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	"github.com/nebulaim/telegramd/biz_model/dal/dataobject"
)

// auth.sendCode#86aef0ec flags:# allow_flashcall:flags.0?true phone_number:string current_number:flags.0?Bool api_id:int api_hash:string = auth.SentCode;
func (s *AuthServiceImpl) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthSendCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// Check TLAuthSendCode
	// CurrentNumber: 是否为本机电话号码
	// 检查数据是否合法
	//switch request.CurrentNumber.(type) {
	//case *mtproto.Bool_BoolFalse:
	//	// 本机电话号码，AllowFlashcall为false
	//	if request.AllowFlashcall == false {
	//		// TODO(@benqi): 数据包非法
	//	}
	//}

	// TODO(@benqi): 独立出统一消息推送系统
	// 检查phpne是否存在，若存在是否在线决定是否通过短信发送或通过其他客户端发送
	// 透传AuthId，UserId，终端类型等
	// 检查满足条件的TransactionHash是否存在，可能的条件：
	//  1. is_deleted !=0 and now - created_at < 15 分钟

	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumer := libphonenumber.NormalizeDigitsOnly(request.PhoneNumber)

	do := dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).SelectByPhoneAndApiIdAndHash(phoneNumer, request.ApiId, request.ApiHash)
	if do == nil {
	    do = &dataobject.AuthPhoneTransactionsDO{}
	    do.ApiId = request.ApiId
	    do.ApiHash = request.ApiHash
	    do.PhoneNumber = phoneNumer
	    do.Code = "123456"
	    do.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	    // TODO(@benqi): 生成一个32字节的随机字串
	    do.TransactionHash = fmt.Sprintf("%20d", id.NextId())

	    dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).Insert(do)
	} else {
	    // TODO(@benqi): 检查是否已经过了失效期
	}

	authSentCode := mtproto.NewTLAuthSentCode()
	authSentCodeType := mtproto.NewTLAuthSentCodeTypeApp()
	authSentCodeType.SetLength(6)
	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	authSentCode.SetPhoneCodeHash(do.TransactionHash)

	// reply := mtproto.MakeAuth_SentCode(authSentCode)
	glog.Infof("AuthSendCode - reply: %s", logger.JsonDebugData(authSentCode))
	return authSentCode.To_Auth_SentCode(), nil
}
