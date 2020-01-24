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

// Package common containers helpers and constants for Mechanism parameters shared between multiple Mechanism types
package common

import (
	"net"

	"github.com/pkg/errors"

	"github.com/networkservicemesh/api/pkg/api/connection"
)

// GetSrcIP returns the source IP parameter of the Mechanism
func GetSrcIP(m *connection.Mechanism) (string, error) {
	return getIPParameter(m, SrcIP)
}

// GetDstIP returns the destination IP parameter of the Mechanism
func GetDstIP(m *connection.Mechanism) (string, error) {
	return getIPParameter(m, DstIP)
}

func getIPParameter(m *connection.Mechanism, name string) (string, error) {
	if m == nil {
		return "", errors.New("mechanism cannot be nil")
	}

	if m.GetParameters() == nil {
		return "", errors.Errorf("mechanism.Parameters cannot be nil: %v", m)
	}

	ip, ok := m.Parameters[name]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), name)
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", errors.Errorf("mechanism.Parameters[%s] must be a valid IPv4 or IPv6 address, instead was: %s: %v", name, ip, m)
	}

	return ip, nil
}
