// Copyright (c) 2019-2020 Cisco Systems, Inc and/or its affiliates.
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

// Package kernel - describe kernel mechanism
package kernel

import (
	"fmt"

	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - kernel mechanism helper
type Mechanism interface {
	// GetNetNSInode - return net ns inode
	GetNetNSInode() string
	// GetParameters() returns the map of all parameters for the mechanism
	GetParameters() map[string]string
	// GetPCIAddress() returns the PCI address to be used to back the kernel interface if set.
	// If the PCIAddress is set, the mechanism should be backed by that PCI device
	// If "" this is not a kernel interface backed by a PCI Device
	GetPCIAddress() string
	// GetInterfaceName - returns the Kernel Interface Name
	//                    this is Mechanism.Parameters[InterfaceNameKey] if set
	//                    otherwise returns a name computed from networkservice.Connection 'conn'
	GetInterfaceName(conn *networkservice.Connection) string
	// GetNetNSURL - returns the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
	GetNetNSURL() string
	SetNetNSURL(urlString string)
}

type mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified mechanism to helper
func ToMechanism(m *networkservice.Mechanism) Mechanism {
	if m.GetType() == MECHANISM {
		return &mechanism{
			m,
		}
	}
	return nil
}

// New - return *networkservice.Mechanism of type kernel using the given netnsURL (inode://${dev}/${ino})
func New(netnsURL string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			NetNSURL: netnsURL,
		},
	}
}

func (m *mechanism) GetParameters() map[string]string {
	if m == nil {
		return nil
	}
	return m.Parameters
}

func (m *mechanism) GetNetNSInode() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[common.NetNSInodeKey]
}

// GetPCIAddress returns PCI address of the device
func (m *mechanism) GetPCIAddress() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[PCIAddress]
}

// IsPCIDevice - true if this mechanism is for a PCI device
func (m *mechanism) IsPCIDevice() bool {
	return m.GetPCIAddress() == ""
}

// GetInterfaceName - returns the Kernel Interface Name
//                    this is Mechanism.Parameters[InterfaceNameKey] if set
//                    otherwise returns a name computed from networkservice.Connection 'conn'
func (m *mechanism) GetInterfaceName(conn *networkservice.Connection) string {
	if m == nil || m.GetParameters()[InterfaceNameKey] == "" {
		ns := conn.GetNetworkService()
		nsMaxLength := LinuxIfMaxLength - 5
		if len(ns) > nsMaxLength {
			ns = ns[:nsMaxLength]
		}
		name := fmt.Sprintf("%s-%s", ns, conn.GetId())
		if len(name) > LinuxIfMaxLength {
			name = name[:LinuxIfMaxLength]
		}
		return name
	}
	name := m.GetParameters()[InterfaceNameKey]
	if len(name) > LinuxIfMaxLength {
		name = name[:LinuxIfMaxLength]
	}
	return name
}

// GetNetNSURL - returns the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *mechanism) GetNetNSURL() string {
	if m == nil || m.GetParameters() == nil {
		return ""
	}
	return m.GetParameters()[NetNSURL]
}

func (m *mechanism) SetNetNSURL(urlString string) {
	if m == nil {
		return
	}
	if m.Parameters == nil {
		m.Parameters = make(map[string]string)
	}
	m.GetParameters()[NetNSURL] = urlString
}
