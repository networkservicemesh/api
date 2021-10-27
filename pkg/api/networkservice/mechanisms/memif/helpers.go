// Copyright (c) 2019-2021 Cisco Systems, Inc.
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

// Package memif provides helper methods for the Mechanism memif
package memif

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
)

// Mechanism provides helper methods for mechanisms of type memif
type Mechanism struct {
	*networkservice.Mechanism
}

// New returns *networkservice.Mechanism of type  memif using the given netnsURL (inode://${dev}/${ino})
func New(netnsURL string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			NetNSURL: netnsURL,
		},
	}
}

// ToMechanism turns a networkservice.Mechanism into a version with helper methods for memif
// If Mechanism m is *not* of type memif.MECHANISM, it returns nil
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

// GetSocketFilename returns memif mechanism socket filename
func (m *Mechanism) GetSocketFilename() string {
	return m.GetParameters()[SocketFilename]
}

// SetSocketFilename sets memif mechanism socket filename
func (m *Mechanism) SetSocketFilename(filename string) {
	m.GetParameters()[SocketFilename] = filename
}

// GetNetNSURL returns the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *Mechanism) GetNetNSURL() string {
	return m.GetParameters()[NetNSURL]
}

// SetNetNSURL sets the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *Mechanism) SetNetNSURL(urlString string) {
	m.GetParameters()[NetNSURL] = urlString
}
