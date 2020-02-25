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

	"github.com/pkg/errors"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a wireguard mechanism utility wrapper
type Mechanism interface {
	// SrcIP -  src ip
	SrcIP() net.IP
	// DstIP - dst ip
	DstIP() net.IP
	// SrcPublicKey - source public key
	SrcPublicKey() (string, error)
	// DstPublicKey - destination public key
	DstPublicKey() (string, error)
	// SrcPrivateKey - source private key
	SrcPrivateKey() (string, error)
	// dstPrivateKey - destination private key
	DstPrivateKey() (string, error)
	// SrcPort - Source interface listening port
	SrcPort() (int, error)
	// SrcPort - Destination interface listening port
	DstPort() (int, error)
}

type mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) Mechanism {
	if m.Type == MECHANISM {
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

func (m *mechanism) stringValue(parameter string) (string, error) {
	if m == nil {
		return "", errors.New("mechanism cannot be nil")
	}

	if m.GetParameters() == nil {
		return "", errors.Errorf("mechanism.Parameters cannot be nil: %v", m)
	}

	value, ok := m.Parameters[parameter]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), parameter)
	}

	return value, nil
}

// SrcPublicKey returns the SrcPublicKey parameter of the Mechanism
func (m *mechanism) SrcPublicKey() (string, error) {
	return m.stringValue(SrcPublicKey)
}

// DstPublicKey returns the DstPublicKey parameter of the Mechanism
func (m *mechanism) DstPublicKey() (string, error) {
	return m.stringValue(DstPublicKey)
}

// SrcPrivateKey returns the SrcPrivateKey parameter of the Mechanism
func (m *mechanism) SrcPrivateKey() (string, error) {
	return m.stringValue(SrcPrivateKey)
}

// DstPrivateKey returns the DstPrivateKey parameter of the Mechanism
func (m *mechanism) DstPrivateKey() (string, error) {
	return m.stringValue(DstPrivateKey)
}

// SrcPort - Source interface listening port
func (m *mechanism) SrcPort() (int, error) {
	srcPortStr, err := m.stringValue(SrcPort)
	if err != nil {
		return 0, err
	}

	srcPort, err := strconv.ParseInt(srcPortStr, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "cannot parse mechanism.Parameters[%s]=%v value", SrcPort, srcPortStr)
	}

	return int(srcPort), nil
}

// DstPort - Destination interface listening port
func (m *mechanism) DstPort() (int, error) {
	dstPortStr, err := m.stringValue(DstPort)
	if err != nil {
		return 0, err
	}

	dstPort, err := strconv.ParseInt(dstPortStr, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "cannot parse mechanism.Parameters[%s]=%v value", DstPort, dstPortStr)
	}

	return int(dstPort), nil
}

// AssignPort - generate unique port by connection ID for wireguard connection
func AssignPort(connID string) string {
	id, err := strconv.ParseUint(connID, 16, 64)
	if err != nil {
		id = 0
	}
	return strconv.FormatUint(BasePort+id, 10)
}
