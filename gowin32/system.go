/*
 * Copyright (c) 2014-2019 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gowin32

import (
	"syscall"

	"github.com/winlabs/gowin32/wrappers"
)

func InitiateSystemShutdown(machineName string, message string, timeout int, forceAppsClosed bool, rebootAfterShutdown bool) error {
	if err := wrappers.InitiateSystemShutdown(
		syscall.StringToUTF16Ptr(machineName),
		syscall.StringToUTF16Ptr(message),
		uint32(timeout),
		forceAppsClosed,
		rebootAfterShutdown); err != nil {
		return NewWindowsError("InitiateSystemShutdown", err)
	}
	return nil
}
