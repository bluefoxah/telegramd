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

package chat_test

import (
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"context"
	"math/rand"
	"github.com/nebulaim/telegramd/base/base"
	"io"
	"time"
)

func DoChatTestClient() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("TestRPCClient...")
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := NewChatTestClient(conn)

	sess := &Session{base.Int64ToString(rand.Int63())}
	fmt.Println("sessionId : ", sess.SessionId)

	stream, err := client.Connect(context.Background(), &Session{base.Int64ToString(rand.Int63())})
	if err != nil {
		glog.Fatalln("connect:", err)
	}

	chatMessages := make(chan *ChatMessage, 1000)
	go func() {
		defer func() { close(chatMessages) }()
		for {
			chatMessage, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				glog.Fatalln("stream.Recv", err)
			}
			chatMessages <- chatMessage
		}
	}()


	go func() {
		for {
			select {
			case chat := <-chatMessages:
				fmt.Printf("Recv chat_message: {session_id: %s, message_data: %s}\n", chat.SenderSessionId, chat.MessageData)
			}
		}
	}()

	var message string
	for {
		fmt.Print("> ")
		if n, err := fmt.Scanln(&message); err == io.EOF {
			return
		} else if n > 0 {
			_, err := client.SendChat(context.Background(), &ChatMessage{SenderSessionId: sess.SessionId, MessageData: message})
			if err != nil {
				glog.Fatalln("sendChat:", err)
			}
		}
	}
}
