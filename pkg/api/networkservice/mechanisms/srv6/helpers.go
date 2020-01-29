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

	"github.com/pkg/errors"


	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a vxlan mechanism utility wrapper
type Mechanism interface {
	// SrcHostIP -  src localsid of mgmt interface
	SrcHostIP() (string, error)
	// DstHostIP - dst localsid of mgmt interface
	DstHostIP() (string, error)
	// SrcBSID -  src BSID
	SrcBSID() (string, error)
	// DstBSID - dst BSID
	DstBSID() (string, error)
	// SrcLocalSID -  src LocalSID
	SrcLocalSID() (string, error)
	// DstLocalSID - dst LocalSID
	DstLocalSID() (string, error)
	// SrcHostLocalSID -  src host unique LocalSID
	SrcHostLocalSID() (string, error)
	// DstHostLocalSID - dst host unique LocalSID
	DstHostLocalSID() (string, error)
	// SrcHardwareAddress -  src hw address
	SrcHardwareAddress() (string, error)
	// DstHardwareAddress - dst hw address
	DstHardwareAddress() (string, error)
}

type mechanism struct {
	*networkservice.Mechanism
}

func (m mechanism) SrcHostIP() (string, error) {
	return getIPParameter(m.Mechanism, SrcHostIP)
}

func (m mechanism) DstHostIP() (string, error) {
	return getIPParameter(m.Mechanism, DstHostIP)
}

func (m mechanism) SrcBSID() (string, error) {
	return getIPParameter(m.Mechanism, SrcBSID)
}

func (m mechanism) DstBSID() (string, error) {
	return getIPParameter(m.Mechanism, DstBSID)
}

func (m mechanism) SrcLocalSID() (string, error) {
	return getIPParameter(m.Mechanism, SrcLocalSID)
}

func (m mechanism) DstLocalSID() (string, error) {
	return getIPParameter(m.Mechanism, DstLocalSID)
}

func (m mechanism) SrcHostLocalSID() (string, error) {
	return getIPParameter(m.Mechanism, SrcHostLocalSID)
}

func (m mechanism) DstHostLocalSID() (string, error) {
	return getIPParameter(m.Mechanism, DstHostLocalSID)
}

func (m mechanism) SrcHardwareAddress() (string, error) {
	return getStringParameter(m.Mechanism, SrcHardwareAddress)
}

func (m mechanism) DstHardwareAddress() (string, error) {
	return getStringParameter(m.Mechanism, DstHardwareAddress)
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

func getIPParameter(m *networkservice.Mechanism, name string) (string, error) {
	ip, err := getStringParameter(m, name)
	if err != nil {
		return "", err
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", errors.Errorf("mechanism.Parameters[%s] must be a valid IPv4 or IPv6 address, instead was: %s: %v", name, ip, m)
	}

	return ip, nil
}

func getStringParameter(m *networkservice.Mechanism, name string) (string, error) {
	if m == nil {
		return "", errors.New("mechanism cannot be nil")
	}

	if m.GetParameters() == nil {
		return "", errors.Errorf("mechanism.Parameters cannot be nil: %v", m)
	}

	v, ok := m.Parameters[name]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), name)
	}

	return v, nil
}
