// Copyright (c) 2020 Cisco Systems, Inc.
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
	"net"
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a vxlan mechanism utility wrapper
type Mechanism interface {
	// SrcIP -  src ip
	SrcIP() net.IP
	// DstIP - dst ip
	DstIP() net.IP
	// VNI - vni
	VNI() uint32
}

type mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) Mechanism {
	if m.GetType() == MECHANISM {
		if m.Parameters == nil {
			m.Parameters = map[string]string{}
		}
		return &mechanism{
			m,
		}
	}
	return nil
}

func (m *mechanism) SrcIP() net.IP {
	return net.ParseIP(m.GetParameters()[SrcIP])
}

func (m *mechanism) DstIP() net.IP {
	return net.ParseIP(m.GetParameters()[DstIP])
}

// VNI returns the VNI parameter of the Mechanism
func (m *mechanism) VNI() uint32 {
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
