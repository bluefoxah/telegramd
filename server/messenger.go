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
	account "github.com/nebulaim/telegramd/server/account/rpc"
	auth "github.com/nebulaim/telegramd/server/auth/rpc"
	bots "github.com/nebulaim/telegramd/server/bots/rpc"
	channels "github.com/nebulaim/telegramd/server/channels/rpc"
	contacts "github.com/nebulaim/telegramd/server/contacts/rpc"
	help "github.com/nebulaim/telegramd/server/help/rpc"
	langpack "github.com/nebulaim/telegramd/server/langpack/rpc"
	messages "github.com/nebulaim/telegramd/server/messages/rpc"
	payments "github.com/nebulaim/telegramd/server/payments/rpc"
	phone "github.com/nebulaim/telegramd/server/phone/rpc"
	photos "github.com/nebulaim/telegramd/server/photos/rpc"
	stickers "github.com/nebulaim/telegramd/server/stickers/rpc"
	updates "github.com/nebulaim/telegramd/server/updates/rpc"
	upload "github.com/nebulaim/telegramd/server/upload/rpc"
	users "github.com/nebulaim/telegramd/server/users/rpc"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", "0.0.0.0:10001")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	mtproto.RegisterRPCAccountServer(grpcServer, &account.AccountServiceImpl{})
	mtproto.RegisterRPCAuthServer(grpcServer, &auth.AuthServiceImpl{})
	mtproto.RegisterRPCBotsServer(grpcServer, &bots.BotsServiceImpl{})
	mtproto.RegisterRPCChannelsServer(grpcServer, &channels.ChannelsServiceImpl{})
	mtproto.RegisterRPCContactsServer(grpcServer, &contacts.ContactsServiceImpl{})
	mtproto.RegisterRPCHelpServer(grpcServer, &help.HelpServiceImpl{})
	mtproto.RegisterRPCLangpackServer(grpcServer, &langpack.LangpackServiceImpl{})
	mtproto.RegisterRPCMessagesServer(grpcServer, &messages.MessagesServiceImpl{})
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
