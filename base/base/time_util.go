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

package base

import (
	"database/sql/driver"
	"time"
)

// Time be used to MySql timestamp converting.
type Time int64

func (t *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case time.Time:
		*t = Time(sc.Unix())
	case string:
		var i int64
		i, err = StringToInt64(sc)
		*t = Time(i)
	}
	return
}

func (t Time) Value() (driver.Value, error) {
	return time.Unix(int64(t), 0), nil
}

func (t Time) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}


/////////////////////////////////////////////////////////////
func NowFormatYMDHMS() string {
	return time.Now().Format("2006-01-02 15:04:05")
}