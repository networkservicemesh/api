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

package vlan

import (
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

const (
	tagBitSize = 12
)

// Mechanism - vlan mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
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

// GetTag returns VLAN tag
func (m *Mechanism) GetTag() uint16 {
	if m.Parameters == nil {
		return 0
	}
	return atou(m.Parameters[TagKey])
}

// SetTag sets VLAN tag
func (m *Mechanism) SetTag(tag uint16) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[TagKey] = utoa(tag)
}

func atou(a string) uint16 {
	u, err := strconv.ParseUint(a, 10, tagBitSize)
	if err != nil {
		return 0
	}
	return uint16(u)
}

func utoa(u uint16) string {
	return strconv.FormatUint(uint64(u), 10)
}
