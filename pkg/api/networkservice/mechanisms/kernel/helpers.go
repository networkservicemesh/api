// Copyright (c) 2019-2021 Cisco Systems, Inc and/or its affiliates.
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

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
)

// Mechanism is a kernel mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// New returns *networkservice.Mechanism of type kernel using the given netnsURL (inode://${dev}/${ino})
func New(netnsURL string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			NetNSURL: netnsURL,
		},
	}
}

// ToMechanism converts unified mechanism to helper
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == MECHANISM {
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

// GetNetNSInode returns the NetNS inode
func (m *Mechanism) GetNetNSInode() string {
	return m.GetParameters()[NetNSInodeKey]
}

// SetNetNSInode sets the NetNS inode
func (m *Mechanism) SetNetNSInode(netNSInode string) {
	m.GetParameters()[NetNSInodeKey] = netNSInode
}

// GetPCIAddress returns the PCI address of the device
func (m *Mechanism) GetPCIAddress() string {
	return m.GetParameters()[PCIAddressKey]
}

// SetPCIAddress sets the PCI address of the device
func (m *Mechanism) SetPCIAddress(pciAddress string) {
	m.GetParameters()[PCIAddressKey] = pciAddress
}

// IsPCIDevice returns if this mechanism is for a PCI device
func (m *Mechanism) IsPCIDevice() bool {
	return m.GetPCIAddress() != ""
}

// ToInterfaceName - create interface name from conn for client or server side for forwarder.
//                   Note: Don't use this in a non-forwarder context
func (m *Mechanism) ToInterfaceName(conn *networkservice.Connection, isClient bool) string {
	// Naming is tricky.  We want to name based on either the next or prev connection id depending on whether we
	// are on the client or server side.  Since this chain element is designed for use in a Forwarder,
	// if we are on the client side, we want to name based on the connection id from the NSE that is Next
	// if we are not the client, we want to name for the connection of of the client addressing us, which is Prev
	namingConn := conn.Clone()
	namingConn.Id = namingConn.GetPrevPathSegment().GetId()
	if isClient {
		namingConn.Id = namingConn.GetNextPathSegment().GetId()
	}
	return m.GetInterfaceName(namingConn)
}

// GetInterfaceName returns the Kernel Interface Name
//                  this is Mechanism.Parameters[InterfaceNameKey] if set
//                  otherwise returns a name computed from networkservice.Connection 'conn'
func (m *Mechanism) GetInterfaceName(conn *networkservice.Connection) string {
	if m.GetParameters()[InterfaceNameKey] == "" {
		ns := conn.GetNetworkService()
		nsMaxLength := LinuxIfMaxLength - 5
		if len(ns) > nsMaxLength {
			ns = ns[:nsMaxLength]
		}
		name := fmt.Sprintf("%s-%s", ns, conn.GetId())
		if len(name) > LinuxIfMaxLength {
			name = name[:LinuxIfMaxLength]
		}
		m.GetParameters()[InterfaceNameKey] = name
	}
	name := m.GetParameters()[InterfaceNameKey]
	if len(name) > LinuxIfMaxLength {
		name = name[:LinuxIfMaxLength]
	}
	return name
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
