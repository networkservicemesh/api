// Copyright (c) 2019-2023 Cisco Systems, Inc.
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
	"net/url"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
)

// Mechanism provides helper methods for mechanisms of type memif
type Mechanism struct {
	*networkservice.Mechanism
}

// New returns *networkservice.Mechanism of type memif using the given socketPath
func New(socketPath string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			SocketFileURL: (&url.URL{Scheme: FileScheme, Path: socketPath}).String(),
		},
	}
}

// NewAbstract returns *networkservice.Mechanism of type memif using the given netNSPath
func NewAbstract(netNSPath string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			NetNSURL: (&url.URL{Scheme: FileScheme, Path: netNSPath}).String(),
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

// GetSocketFilename [abstract sockets case] returns memif mechanism socket filename
func (m *Mechanism) GetSocketFilename() string {
	return m.GetParameters()[SocketFilename]
}

// SetSocketFilename [abstract sockets case] sets memif mechanism socket filename
func (m *Mechanism) SetSocketFilename(filename string) {
	m.GetParameters()[SocketFilename] = filename
}

// GetNetNSURL [abstract sockets case] returns the NetNS URL, it can be either:
// * file:///proc/${pid}/ns/net - ${pid} process net NS
// * inode://${dev}/${ino} - while transferring file between processes using grpcfd
func (m *Mechanism) GetNetNSURL() string {
	return m.GetParameters()[NetNSURL]
}

// SetNetNSURL [abstract sockets case] sets the NetNS URL - file:///proc/${pid}/ns/net
func (m *Mechanism) SetNetNSURL(urlString string) {
	m.GetParameters()[NetNSURL] = urlString
}

// GetSocketFileURL [FS sockets case] returns the memif socketfile URL, it can be either:
//
//	// * file://${path} - memif socketfile
//	// * inode://${dev}/${ino} - while transferring file between processes using grpcfd
func (m *Mechanism) GetSocketFileURL() string {
	return m.GetParameters()[SocketFileURL]
}

// SetSocketFileURL [FS sockets case] sets the memif socketfile URL.
func (m *Mechanism) SetSocketFileURL(urlString string) {
	m.GetParameters()[NetNSURL] = urlString
}
