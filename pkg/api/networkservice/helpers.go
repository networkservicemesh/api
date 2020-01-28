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

package networkservice

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// Clone clones request
func (m *NetworkServiceRequest) Clone() *NetworkServiceRequest {
	return proto.Clone(m).(*NetworkServiceRequest)
}

// GetRequestConnection returns request connection
func (m *NetworkServiceRequest) GetRequestConnection() *Connection {
	return m.GetConnection()
}

// SetRequestConnection sets request connection
func (m *NetworkServiceRequest) SetRequestConnection(conn *Connection) {
	m.Connection = conn
}

// GetRequestMechanismPreferences returns request mechanism preferences
func (m *NetworkServiceRequest) GetRequestMechanismPreferences() []*Mechanism {
	preferences := make([]*Mechanism, 0, len(m.MechanismPreferences))
	preferences = append(preferences, m.MechanismPreferences...)

	return preferences
}

// SetRequestMechanismPreferences sets request mechanism preferences
func (m *NetworkServiceRequest) SetRequestMechanismPreferences(mechanismPreferences []*Mechanism) {
	m.MechanismPreferences = mechanismPreferences
}

// IsValid returns if request is valid
func (m *NetworkServiceRequest) IsValid() error {
	if m == nil {
		return errors.New("request cannot be nil")
	}

	if m.GetConnection() == nil {
		return errors.Errorf("request.Connection cannot be nil %v", m)
	}

	if err := m.GetConnection().IsValid(); err != nil {
		return errors.Errorf("request.Connection is invalid: %s: %v", err, m)
	}

	if m.GetMechanismPreferences() == nil {
		return errors.Errorf("request.MechanismPreferences cannot be nil: %v", m)
	}

	if len(m.GetMechanismPreferences()) < 1 {
		return errors.Errorf("request.MechanismPreferences must have at least one entry: %v", m)
	}

	return nil
}
