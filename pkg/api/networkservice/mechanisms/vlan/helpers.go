// Copyright (c) 2021-2022 Nordix Foundation.
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

// Package vlan provides helper methods for the Mechanism vlan
package vlan

import (
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a vlan mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified Mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == MECHANISM {
		if m.Parameters == nil {
			m.Parameters = map[string]string{}
		}
		return &Mechanism{
			m,
		}
	}
	return nil
}

// GetParameters returns the map of all parameters to the mechanism
func (m *Mechanism) GetParameters() map[string]string {
	if m == nil {
		return map[string]string{}
	}
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	return m.Parameters
}

// GetVlanID returns the VlanID parameter of the Mechanism
func (m *Mechanism) GetVlanID() uint32 {
	if m == nil {
		return 0
	}

	if m.GetParameters() == nil {
		return 0
	}

	vid := m.Parameters[ID]
	// vlan ID range is 0 to 4,095 stored in 12 bit
	vlanid, err := strconv.ParseUint(vid, 10, 12)

	if err != nil {
		return 0
	}

	return uint32(vlanid)
}

// SetVlanID - set the VLAN ID and return the *vlan.Mechanism
func (m *Mechanism) SetVlanID(vlanid uint32) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[ID] = strconv.FormatUint(uint64(vlanid), 10)
	return m
}
