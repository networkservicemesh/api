// Copyright (c) 2022 Cisco and/or its affiliates.
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

package ipsec

import (
	"net"
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism is an ipsec mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == MECHANISM {
		if m.GetParameters() == nil {
			m.Parameters = map[string]string{}
		}
		return &Mechanism{
			m,
		}
	}
	return nil
}

// SrcIP returns source ip
func (m *Mechanism) SrcIP() net.IP {
	return net.ParseIP(m.GetParameters()[SrcIP])
}

// SetSrcIP sets source ip
func (m *Mechanism) SetSrcIP(ip net.IP) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[SrcIP] = ip.String()
	return m
}

// DstIP returns destination ip
func (m *Mechanism) DstIP() net.IP {
	return net.ParseIP(m.GetParameters()[DstIP])
}

// SetDstIP sets destination ip
func (m *Mechanism) SetDstIP(ip net.IP) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[DstIP] = ip.String()
	return m
}

// SrcPublicKey returns the SrcPublicKey parameter of the Mechanism
func (m *Mechanism) SrcPublicKey() string {
	return m.GetParameters()[SrcPublicKey]
}

// SetSrcPublicKey sets new source public key
func (m *Mechanism) SetSrcPublicKey(key string) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[SrcPublicKey] = key
	return m
}

// DstPublicKey returns the DstPublicKey parameter of the Mechanism
func (m *Mechanism) DstPublicKey() string {
	return m.GetParameters()[DstPublicKey]
}

// SetDstPublicKey sets new destination public key
func (m *Mechanism) SetDstPublicKey(key string) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[DstPublicKey] = key
	return m
}

// SrcPort - Source interface listening port
func (m *Mechanism) SrcPort() uint16 {
	return atou16(m.GetParameters()[SrcPort])
}

// SetSrcPort sets source udp port
func (m *Mechanism) SetSrcPort(port uint16) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[SrcPort] = strconv.FormatUint(uint64(port), 10)
	return m
}

// DstPort - Destination interface listening port
func (m *Mechanism) DstPort() uint16 {
	return atou16(m.GetParameters()[DstPort])
}

// SetDstPort sets destination udp port
func (m *Mechanism) SetDstPort(port uint16) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[DstPort] = strconv.FormatUint(uint64(port), 10)
	return m
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

func atou16(a string) uint16 {
	u, err := strconv.ParseUint(a, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(u)
}
