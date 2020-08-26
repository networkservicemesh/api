// Copyright (c) 2020 Doc.ai and/or its affiliates.
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

package sriov

import (
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/kernel"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/vfio"
)

// Mechanism is a SR-IOV mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism converts unified mechanism to helper
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == kernel.MECHANISM || m.GetType() == vfio.MECHANISM {
		return &Mechanism{
			m,
		}
	}
	return nil
}

// GetEndpointVfPCIAddress returns endpoint VF PCI address
func (m *Mechanism) GetEndpointVfPCIAddress() string {
	if m.Parameters == nil {
		return ""
	}
	return m.Parameters[EndpointVfPCIAddressKey]
}

// SetEndpointVfPCIAddress sets endpoint VF PCI address
func (m *Mechanism) SetEndpointVfPCIAddress(endpointVfPCIAddress string) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[EndpointVfPCIAddressKey] = endpointVfPCIAddress
}

// GetClientPfPCIAddress returns client PF PCI address
func (m *Mechanism) GetClientPfPCIAddress() string {
	if m.Parameters == nil {
		return ""
	}
	return m.Parameters[ClientPfPCIAddressKey]
}

// SetClientPfPCIAddress sets client PF PCI address
func (m *Mechanism) SetClientPfPCIAddress(clientPfPCIAddress string) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[ClientPfPCIAddressKey] = clientPfPCIAddress
}

// GetClientVfPCIAddress returns client VF PCI address
func (m *Mechanism) GetClientVfPCIAddress() string {
	if m.Parameters == nil {
		return ""
	}
	return m.Parameters[ClientVfPCIAddressKey]
}

// SetClientVfPCIAddress sets client VF PCI address
func (m *Mechanism) SetClientVfPCIAddress(clientVfPCIAddress string) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[ClientVfPCIAddressKey] = clientVfPCIAddress
}

// GetIommuGroup returns IOMMU group id
func (m *Mechanism) GetIommuGroup() uint {
	if m.Parameters == nil {
		return 0
	}
	return atou(m.Parameters[IommuGroupKey])
}

// SetIommuGroup sets IOMMU group id
func (m *Mechanism) SetIommuGroup(iommuGroup uint) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[IommuGroupKey] = utoa(iommuGroup)
}

func atou(a string) uint {
	u, err := strconv.ParseUint(a, 10, 0)
	if err != nil {
		return 0
	}
	return uint(u)
}

func utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
