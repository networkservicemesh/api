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
	"strconv"

	"github.com/pkg/errors"

	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a vxlan mechanism utility wrapper
type Mechanism interface {
	// SrcIP -  src ip
	SrcIP() (string, error)
	// DstIP - dst ip
	DstIP() (string, error)
	// VNI - vni
	VNI() (uint32, error)
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

func (m *mechanism) SrcIP() (string, error) {
	return common.GetSrcIP(m.Mechanism)
}

func (m *mechanism) DstIP() (string, error) {
	return common.GetDstIP(m.Mechanism)
}

// VNI returns the VNI parameter of the Mechanism
func (m *mechanism) VNI() (uint32, error) {
	if m == nil {
		return 0, errors.New("mechanism cannot be nil")
	}

	if m.GetParameters() == nil {
		return 0, errors.Errorf("mechanism.Parameters cannot be nil: %v", m)
	}

	vxlanvni, ok := m.Parameters[VNI]
	if !ok {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), VNI)
	}

	vni, err := strconv.ParseUint(vxlanvni, 10, 24)

	if err != nil {
		return 0, errors.Wrapf(err, "mechanism.Parameters[%s] must be a valid 24-bit unsigned integer, instead was: %s: %v", VNI, vxlanvni, m)
	}

	return uint32(vni), nil
}
