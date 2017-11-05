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

package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"net"
	"google.golang.org/grpc"
	account "github.com/nebulaim/telegramd/biz_server/account/rpc"
	auth "github.com/nebulaim/telegramd/biz_server/auth/rpc"
	bots "github.com/nebulaim/telegramd/biz_server/bots/rpc"
	channels "github.com/nebulaim/telegramd/biz_server/channels/rpc"
	contacts "github.com/nebulaim/telegramd/biz_server/contacts/rpc"
	help "github.com/nebulaim/telegramd/biz_server/help/rpc"
	langpack "github.com/nebulaim/telegramd/biz_server/langpack/rpc"
	messages "github.com/nebulaim/telegramd/biz_server/messages/rpc"
	payments "github.com/nebulaim/telegramd/biz_server/payments/rpc"
	phone "github.com/nebulaim/telegramd/biz_server/phone/rpc"
	photos "github.com/nebulaim/telegramd/biz_server/photos/rpc"
	stickers "github.com/nebulaim/telegramd/biz_server/stickers/rpc"
	updates "github.com/nebulaim/telegramd/biz_server/updates/rpc"
	upload "github.com/nebulaim/telegramd/biz_server/upload/rpc"
	users "github.com/nebulaim/telegramd/biz_server/users/rpc"
	"github.com/jmoiron/sqlx"
	"github.com/nebulaim/telegramd/biz_model/dal/dao"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")

}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	// dsl ==> root:@/nebulaim?charset=utf8
	mysqlDsn := "root:@/nebulaim?charset=utf8"

	db, err := sqlx.Connect("mysql", mysqlDsn)
	if err != nil {
		glog.Fatalf("Connect mysql %s error: %s", mysqlDsn, err)
		return
	}

	lis, err := net.Listen("tcp", "0.0.0.0:10001")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	usersDAO := dao.NewUsersDAO(db)
	devicesDAO := dao.NewDevicesDAO(db)
	// masterKeysDAO := dao.NewMasterKeysDAO(db)
	authUsersDAO := dao.NewAuthUsersDAO(db)
	authPhoneTransactionsDAO := dao.NewAuthPhoneTransactionsDAO(db)
	// authsDAO := dao.NewAuthsDAO(db)
	// authSaltsDAO := dao.NewAuthSaltsDAO(db)
	// appsDAO := dao.NewAppsDAO(db)
	userDialogsDAO := dao.NewUserDialogsDAO(db)

	// AccountServiceImpl
	mtproto.RegisterRPCAccountServer(grpcServer, &account.AccountServiceImpl{
		UsersDAO: usersDAO,
		DeviceDAO: devicesDAO,
	})

	// AuthServiceImpl
	mtproto.RegisterRPCAuthServer(grpcServer, &auth.AuthServiceImpl{
		UsersDAO: usersDAO,
		AuthPhoneTransactionsDAO:	authPhoneTransactionsDAO,
	})

	mtproto.RegisterRPCBotsServer(grpcServer, &bots.BotsServiceImpl{})
	mtproto.RegisterRPCChannelsServer(grpcServer, &channels.ChannelsServiceImpl{})
	mtproto.RegisterRPCContactsServer(grpcServer, &contacts.ContactsServiceImpl{})
	mtproto.RegisterRPCHelpServer(grpcServer, &help.HelpServiceImpl{})
	mtproto.RegisterRPCLangpackServer(grpcServer, &langpack.LangpackServiceImpl{})

	// MessagesServiceImpl
	mtproto.RegisterRPCMessagesServer(grpcServer, &messages.MessagesServiceImpl{
		AuthUsersDAO: authUsersDAO,
		UserDialogsDAO:	userDialogsDAO,
	})

	mtproto.RegisterRPCPaymentsServer(grpcServer, &payments.PaymentsServiceImpl{})
	mtproto.RegisterRPCPhoneServer(grpcServer, &phone.PhoneServiceImpl{})
	mtproto.RegisterRPCPhotosServer(grpcServer, &photos.PhotosServiceImpl{})
	mtproto.RegisterRPCStickersServer(grpcServer, &stickers.StickersServiceImpl{})
	mtproto.RegisterRPCUpdatesServer(grpcServer, &updates.UpdatesServiceImpl{})
	mtproto.RegisterRPCUploadServer(grpcServer, &upload.UploadServiceImpl{})
	mtproto.RegisterRPCUsersServer(grpcServer, &users.UsersServiceImpl{})

	glog.Info("NewRPCServer in 0.0.0.0:10001.")

	grpcServer.Serve(lis)
}
