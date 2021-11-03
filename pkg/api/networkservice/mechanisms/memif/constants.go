// Copyright (c) 2020-2021 Cisco Systems, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memif

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
)

const (
	// MECHANISM string
	MECHANISM = "MEMIF"

	// Mechanism parameters

	// SocketFilename - name of the memif socketfile
	SocketFilename = "socketfile"

	// NetNSURL - url representing a inode - fmt.Sprintf("inode://%d/%d",dev,ino)
	NetNSURL = common.InodeURL

	// SocketFileScheme - expected scheme of the NetNSURL
	SocketFileScheme = "file"
)
