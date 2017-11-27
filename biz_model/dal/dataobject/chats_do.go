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

package dataobject

type ChatsDO struct {
	Id                   int32  `db:"id"`
	CreatorUserId        int32  `db:"creator_user_id"`
	CreateRandomId       int64  `db:"create_random_id"`
	AccessHash           int64  `db:"access_hash"`
	ParticipantCount     int32  `db:"participant_count"`
	Title                string `db:"title"`
	TitleChangerUserId   int32  `db:"title_changer_user_id"`
	TitleChangedAt       string `db:"title_changed_at"`
	TitleChangeRandomId  int64  `db:"title_change_random_id"`
	AvatarChangerUserId  int32  `db:"avatar_changer_user_id"`
	AvatarChangedAt      string `db:"avatar_changed_at"`
	AvatarChangeRandomId int64  `db:"avatar_change_random_id"`
	IsPublic             int8   `db:"is_public"`
	About                string `db:"about"`
	Topic                string `db:"topic"`
	IsHidden             int8   `db:"is_hidden"`
	Version              int32  `db:"version"`
	CreatedAt            string `db:"created_at"`
	UpdatedAt            string `db:"updated_at"`
}
