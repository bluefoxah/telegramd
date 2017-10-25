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

package crypto

import (
	"fmt"
)

const (
	KEY = "12345678901234561234567890123456"
	IV = "1234567890123456"
)


func main() {
	MyString := "This is my string and I want to protect it with encryption"

	fmt.Println("We start with a plain text: %s \n", MyString)
	MyStringByte := []byte(MyString)
	encryptor, _ := NewAesCTR128Encrypt([]byte(KEY), []byte(IV))
	Encrypted := encryptor.Encrypt(MyStringByte);
	fmt.Println("We encrypted the string this way %s \n", string(Encrypted))

	decryptor, _ := NewAesCTR128Encrypt([]byte(KEY), []byte(IV))
	Decrypted := decryptor.Encrypt(MyStringByte);
	fmt.Println("Than we have the plain text again %s \n", string(Decrypted))
}

