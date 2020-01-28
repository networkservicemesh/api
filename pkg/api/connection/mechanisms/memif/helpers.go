// Copyright (c) 2019-2020 Cisco Systems, Inc.
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

// Package memif provides helper methods for the Mechanism memif
package memif

import (
	"github.com/networkservicemesh/api/pkg/api/connection"
	"github.com/networkservicemesh/api/pkg/api/connection/mechanisms/common"
)

// Mechanism provides helper methods for mechanisms of type memif
type Mechanism interface {
	GetSocketFilename() string
	GetWorkspace() string
	GetNetNSInode() string
}

type mechanism struct {
	*connection.Mechanism
}

// ToMechanism turns a connection.Mechanism into a version with helper methods for memif
// If Mechanism m is *not* of type memif.MECHANISM, it returns nil
func ToMechanism(m *connection.Mechanism) Mechanism {
	if m.GetType() == MECHANISM {
		return &mechanism{
			m,
		}
	}
	return nil
}

// GetWorkspace get the name of the Workspace directory
func (m *mechanism) GetWorkspace() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[common.Workspace]
}

// GetSocketFilename returns memif mechanism socket filename
func (m *mechanism) GetSocketFilename() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[SocketFilename]
}

// GetNetNsInode get the name of the Netns Inode
func (m *mechanism) GetNetNSInode() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[common.NetNSInodeKey]
}
