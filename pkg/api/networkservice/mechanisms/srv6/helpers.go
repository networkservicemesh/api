// Copyright (c) 2020 Cisco and/or its affiliates.
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

package srv6

import (
	"net"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a vxlan mechanism utility wrapper
type Mechanism interface {
	// SrcHostIP -  src localsid of mgmt interface
	SrcHostIP() string
	// DstHostIP - dst localsid of mgmt interface
	DstHostIP() string
	// SrcBSID -  src BSID
	SrcBSID() string
	// DstBSID - dst BSID
	DstBSID() string
	// SrcLocalSID -  src LocalSID
	SrcLocalSID() string
	// DstLocalSID - dst LocalSID
	DstLocalSID() string
	// SrcHostLocalSID -  src host unique LocalSID
	SrcHostLocalSID() string
	// DstHostLocalSID - dst host unique LocalSID
	DstHostLocalSID() string
	// SrcHardwareAddress -  src hw address
	SrcHardwareAddress() string
	// DstHardwareAddress - dst hw address
	DstHardwareAddress() string
}

type mechanism struct {
	*networkservice.Mechanism
}

func (m mechanism) SrcHostIP() string {
	return getIPParameter(m.Mechanism, SrcHostIP)
}

func (m mechanism) DstHostIP() string {
	return getIPParameter(m.Mechanism, DstHostIP)
}

func (m mechanism) SrcBSID() string {
	return getIPParameter(m.Mechanism, SrcBSID)
}

func (m mechanism) DstBSID() string {
	return getIPParameter(m.Mechanism, DstBSID)
}

func (m mechanism) SrcLocalSID() string {
	return getIPParameter(m.Mechanism, SrcLocalSID)
}

func (m mechanism) DstLocalSID() string {
	return getIPParameter(m.Mechanism, DstLocalSID)
}

func (m mechanism) SrcHostLocalSID() string {
	return getIPParameter(m.Mechanism, SrcHostLocalSID)
}

func (m mechanism) DstHostLocalSID() string {
	return getIPParameter(m.Mechanism, DstHostLocalSID)
}

func (m mechanism) SrcHardwareAddress() string {
	return m.Mechanism.GetParameters()[SrcHardwareAddress]
}

func (m mechanism) DstHardwareAddress() string {
	return m.Mechanism.GetParameters()[DstHardwareAddress]
}

// ToMechanism - convert unified mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) Mechanism {
	if m.GetType() == MECHANISM {
		if m.GetParameters() == nil {
			m.Parameters = map[string]string{}
		}
		return &mechanism{
			m,
		}
	}
	return nil
}

func getIPParameter(m *networkservice.Mechanism, name string) string {
	ip := m.GetParameters()[name]
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return ""
	}
	return ip
}
