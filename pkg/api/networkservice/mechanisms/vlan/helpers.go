// Copyright (c) 2021 Nordix Foundation.
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

import "github.com/networkservicemesh/api/pkg/api/networkservice"

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

// GetInterfaceName returns the Kernel Interface Name
func (m *Mechanism) GetInterfaceName() string {
	return m.GetParameters()[InterfaceNameKey]
}

// SetInterfaceName sets the Kernel Interface Name
func (m *Mechanism) SetInterfaceName(interfaceName string) {
	m.GetParameters()[InterfaceNameKey] = interfaceName
}

// GetNetNSURL returns the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *Mechanism) GetNetNSURL() string {
	return m.GetParameters()[NetNSURL]
}

// SetNetNSURL sets the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *Mechanism) SetNetNSURL(urlString string) {
	m.GetParameters()[NetNSURL] = urlString
}
