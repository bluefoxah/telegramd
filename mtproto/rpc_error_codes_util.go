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

package mtproto

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/golang/glog"
)

// FILE_MIGRATE_X = 303000;
// PHONE_MIGRATE_X = 303001;
// NETWORK_MIGRATE_X = 303002;
// USER_MIGRATE_X = 303003;
//
// ERROR_SEE_OTHER code has _X is dc number, We use custom NewXXXX()
func NewFileMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{
		ErrorCode: int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("FILE_MIGRATE_%d: %s", dc, message),
	}
}

func NewPhoneMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{
		ErrorCode: int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("PHONE_MIGRATE_%d: %s", dc, message),
	}
}

func NewNetworkMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{
		ErrorCode: int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("NETWORK_MIGRATE_%d: %s", dc, message),
	}
}

func NewUserMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{
		ErrorCode: int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("USER_MIGRATE_%d: %s", dc, message),
	}
}

// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
//
func NewFloodWaitX(second int32, message string) *TLRpcError {
	return &TLRpcError{
		ErrorCode: int32(TLRpcErrorCodes_FLOOD),
		ErrorMessage: fmt.Sprintf("FLOOD_WAIT_%d: %s", second, message),
	}
}

// normal code NewXXX
func NewRpcError(code int32, message string) (err *TLRpcError) {
	if name, ok := TLRpcErrorCodes_name[int32(code)]; ok {
		if code <= int32(TLRpcErrorCodes_OTHER2) {
			err = &TLRpcError{
				ErrorCode: code,
				ErrorMessage: fmt.Sprintf("%s: %s", name, message),
			}
		} else {
			switch code {
			// Not
			case int32(TLRpcErrorCodes_FILE_MIGRATE_X),
				 int32(TLRpcErrorCodes_NETWORK_MIGRATE_X),
				 int32(TLRpcErrorCodes_PHONE_MIGRATE_X),
				 int32(TLRpcErrorCodes_USER_MIGRATE_X):
				err = &TLRpcError{
					ErrorCode: int32(TLRpcErrorCodes_OTHER2),
					ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use New%s(dc, %s), ", name, name, message),
				}
				glog.Error(err)

			case int32(TLRpcErrorCodes_FLOOD_WAIT_X):
				err = &TLRpcError{
					ErrorCode: int32(TLRpcErrorCodes_FLOOD),
					ErrorMessage: fmt.Sprintf("FLOOD_WAIT_%d: %s", name, name),
				}
				glog.Error(err)
			default:
				err = &TLRpcError{
					// subcode = code * 10000 + i
					ErrorCode: code / 10000,
					ErrorMessage: fmt.Sprintf("%s: %s", name, message),
				}
			}
		}
	} else {
		err = &TLRpcError{
			// subcode = code * 10000 + i
			ErrorCode: int32(TLRpcErrorCodes_INTERNAL),
			ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: code = %d, message = %s", code, message),
		}
	}

	return
}

// Impl error interface
func (e *TLRpcError) Error() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s", e.ErrorCode, e.ErrorMessage)
}

// Impl error interface
func (e *TLRpcError) ToGrpcStatus() *status.Status {
	return status.New(codes.Internal, e.Error())
}

/*
// Impl error interface
func (e *TLRpcError) ToMetadata() (metadata.MD) {
	// return status.New(codes.Internal, e.Error())
	if name2, ok := TLRpcErrorCodes_name[e.ErrorCode]; ok {
		return metadata.Pairs(
			"rpc_error_code", name2,
			"rpc_error_message", e.ErrorMessage)
	}

	return metadata.Pairs(
		"rpc_error_code", "OTHER2",
		"rpc_error_message", fmt.Sprintf("INTERNAL_SERVER_ERROR: %s", e.ErrorMessage))
}

func NewRpcErrorFromMetadata(md metadata.MD) (*TLRpcError, error) {
	e := &TLRpcError{}

	if v, ok := getFirstKeyVal(md, "rpc_error_code"); ok {
		if code, ok := TLRpcErrorCodes_value[v]; !ok {
			return nil, fmt.Errorf("Invalid rpc_error_code: %s", v)
		} else {
			e.ErrorCode = code
		}
	} else {
		return nil, fmt.Errorf("Not found metadata's key: rpc_error_code")
	}

	if v, ok := getFirstKeyVal(md, "rpc_error_message"); !ok {
		e.ErrorMessage = v
	} else {
		return nil, fmt.Errorf("Not found metadata's key: rpc_error_message")
	}

	return e, nil
}
*/
