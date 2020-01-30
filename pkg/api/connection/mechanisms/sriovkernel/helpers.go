// Copyright (c) 2020 Intel Corporation. All Rights Reserved.
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

package sriovkernel

import (
	"github.com/networkservicemesh/api/pkg/api/connection"
	"github.com/networkservicemesh/api/pkg/api/connection/mechanisms/common"
)

// Mechanism represents SRIOV kernel stack interface mechanism helper
type Mechanism interface {
	GetNetNsInode() string
	GetParameters() map[string]string
	GetPCIAddress() string
}

type mechanism struct {
	*connection.Mechanism
}

// ToMechanism converts unified mechanism to helper
func ToMechanism(m *connection.Mechanism) Mechanism {
	if m.GetType() == MECHANISM {
		return &mechanism{
			m,
		}
	}
	return nil
}

// GetParameters returns map of mechanism parameters
func (m *mechanism) GetParameters() map[string]string {
	if m == nil {
		return nil
	}
	return m.Parameters
}

// GetParameters returns network namespace Inode
func (m *mechanism) GetNetNsInode() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[common.NetNsInodeKey]
}

// GetPCIAddress returns PCI address of the VF device
func (m *mechanism) GetPCIAddress() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[PCIAddress]
}
