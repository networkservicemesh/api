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

package vxlan

import (
	"math/rand"
	"net"
	"strconv"

	"github.com/pkg/errors"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
)

// Mechanism - a vxlan Mechanism utility wrapper
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

// SrcIP - Source net.IP for the VXLAN tunnel
func (m *Mechanism) SrcIP() net.IP {
	return net.ParseIP(m.GetParameters()[SrcIP])
}

// SetSrcIP - sets the SrcIP for the VXLAN and returns the *vxlan.Mechanism
func (m *Mechanism) SetSrcIP(ip net.IP) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[SrcIP] = ip.String()
	return m
}

// DstIP - returns the net.IP for the DstIP of the vxlan tunnel
func (m *Mechanism) DstIP() net.IP {
	return net.ParseIP(m.GetParameters()[DstIP])
}

// SetDstIP - sets the DstIP for the VXLAN and returns the *vxlan.Mechanism
func (m *Mechanism) SetDstIP(ip net.IP) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[DstIP] = ip.String()
	return m
}

// VNI returns the VNI parameter of the Mechanism
func (m *Mechanism) VNI() uint32 {
	if m == nil {
		return 0
	}

	if m.GetParameters() == nil {
		return 0
	}

	vxlanvni := m.Parameters[VNI]

	vni, err := strconv.ParseUint(vxlanvni, 10, 24)

	if err != nil {
		return 0
	}

	return uint32(vni)
}

// SetVNI - set the VNI for the vxlan tunnel and return *vxlan.Mechanism
func (m *Mechanism) SetVNI(vni uint32) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[VNI] = strconv.FormatUint(uint64(vni), 10)
	return m
}

// EvenVNI - true if the VNI issues by the NSE should be even, false otherwise
func (m *Mechanism) EvenVNI() bool {
	srcStr, ok := m.GetParameters()[common.SrcOriginalIP]
	if !ok {
		srcStr = m.GetParameters()[SrcIP]
	}
	src := net.ParseIP(srcStr)
	if src == nil {
		return true
	}
	dstStr, ok := m.GetParameters()[common.DstOriginalIP]
	if !ok {
		dstStr = m.GetParameters()[DstIP]
	}
	dst := net.ParseIP(dstStr)
	if dst == nil {
		return false
	}
	return compareIps(src, dst) <= 0
}

// GenerateRandomVNI - generates a random VNI for the mechanism, even or odd as determined by EvenVNI()
func (m *Mechanism) GenerateRandomVNI() (uint32, error) {
	if m.SrcIP() != nil && m.DstIP() != nil {
		vni := rand.Uint32() // #nosec
		if m.EvenVNI() {
			vni = (2 * vni) & 0x00FFFFFF
		} else {
			vni = (2*vni + 1) & 0x00FFFFFF
		}
		return vni, nil
	}
	return 0, errors.Errorf("both srcIP(%s) and dstIP(%s) must be non-nil", m.SrcIP(), m.DstIP())
}

func compareIps(ip1, ip2 net.IP) int {
	for index, value := range ip1 {
		if value < ip2[index] {
			return -1
		}
		if value > ip2[index] {
			return 1
		}
	}
	return 0
}

// MTU - return MTU value - 0 if unset
func (m *Mechanism) MTU() uint32 {
	mtu, err := strconv.ParseUint(m.GetParameters()[MTU], 10, 32)
	if err != nil {
		return 0
	}

	return uint32(mtu)
}

// SetMTU - set the MTU value
func (m *Mechanism) SetMTU(mtu uint32) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[MTU] = strconv.FormatUint(uint64(mtu), 10)

	return m
}

// SrcPort - Source vxlan listening port
func (m *Mechanism) SrcPort() uint16 {
	port, err := strconv.ParseUint(m.GetParameters()[SrcPort], 10, 16)
	if err != nil {
		return 0
	}
	return uint16(port)
}

// SetSrcPort - Set source vxlan listening port
func (m *Mechanism) SetSrcPort(port uint16) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[SrcPort] = strconv.FormatUint(uint64(port), 10)
	return m
}

// DstPort - Destination vxlan listening port
func (m *Mechanism) DstPort() uint16 {
	port, err := strconv.ParseUint(m.GetParameters()[DstPort], 10, 16)
	if err != nil {
		return 0
	}
	return uint16(port)
}

// SetDstPort - Set destination vxlan listening port
func (m *Mechanism) SetDstPort(port uint16) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[DstPort] = strconv.FormatUint(uint64(port), 10)
	return m
}
