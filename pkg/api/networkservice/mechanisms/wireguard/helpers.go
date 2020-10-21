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

package wireguard

import (
	"net"
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a wireguard mechanism utility wrapper
type Mechanism interface {
	// SrcIP -  src ip
	SrcIP() net.IP
	// DstIP - dst ip
	DstIP() net.IP
	// SrcPublicKey - source public key
	SrcPublicKey() string
	// DstPublicKey - destination public key
	DstPublicKey() string
	// SrcPort - Source interface listening port
	SrcPort() int
	// SrcPort - Destination interface listening port
	DstPort() int
}

type mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) Mechanism {
	if m.Type == MECHANISM {
		if m.GetParameters() == nil {
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

// SrcPublicKey returns the SrcPublicKey parameter of the Mechanism
func (m *mechanism) SrcPublicKey() string {
	return m.GetParameters()[SrcPublicKey]
}

// DstPublicKey returns the DstPublicKey parameter of the Mechanism
func (m *mechanism) DstPublicKey() string {
	return m.GetParameters()[DstPublicKey]
}

// SrcPort - Source interface listening port
func (m *mechanism) SrcPort() int {
	return atoi(m.GetParameters()[SrcPort])
}

// DstPort - Destination interface listening port
func (m *mechanism) DstPort() int {
	return atoi(m.GetParameters()[DstPort])
}

// GetPort - returns unique port by connection ID for wireguard connection
func GetPort(connID string) string {
	id, err := strconv.ParseUint(connID, 16, 0)
	if err != nil {
		id = 0
	}
	return strconv.FormatUint(BasePort+id, 10)
}

func atoi(a string) int {
	i, err := strconv.ParseInt(a, 10, strconv.IntSize)
	if err != nil {
		return 0
	}
	return int(i)
}
