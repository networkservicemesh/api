// Copyright (c) 2018-2020 VMware, Inc.
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

package connection

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/networkservicemesh/api/pkg/api/connectioncontext"
)

// IsRemote returns if connection is remote
func (m *Connection) IsRemote() bool {
	if m == nil {
		return false
	}
	// If we have two or more, it is remote
	return len(m.GetPath().GetPathSegments()) > 1
}

// GetSourceNetworkServiceManagerName - return source network service manager name
func (m *Connection) GetSourceNetworkServiceManagerName() string {
	if m == nil {
		return ""
	}
	if len(m.GetPath().GetPathSegments()) > 0 {
		return m.GetPath().GetPathSegments()[0].GetName()
	}
	return ""
}

// GetDestinationNetworkServiceManagerName - return destination network service manager name
func (m *Connection) GetDestinationNetworkServiceManagerName() string {
	if m == nil {
		return ""
	}
	if len(m.GetPath().GetPathSegments()) >= 2 {
		return m.GetPath().GetPathSegments()[1].GetName()
	}
	return ""
}

// Equals returns if connection equals given connection
func (m *Connection) Equals(connection *Connection) bool {
	// use as proto.Message
	return proto.Equal(m, connection)
}

// Clone clones connection
func (m *Connection) Clone() *Connection {
	// use as proto.Message
	return proto.Clone(m).(*Connection)
}

// UpdateContext checks and tries to set connection context
func (m *Connection) UpdateContext(context *connectioncontext.ConnectionContext) error {
	if err := context.MeetsRequirements(m.Context); err != nil {
		return err
	}

	oldContext := m.Context
	m.Context = context

	if err := m.IsValid(); err != nil {
		m.Context = oldContext
		return err
	}

	return nil
}

// IsValid checks if connection is minimally valid
func (m *Connection) IsValid() error {
	if m == nil {
		return errors.New("connection cannot be nil")
	}

	if m.GetNetworkService() == "" {
		return errors.Errorf("connection.NetworkService cannot be empty: %v", m)
	}

	if m.GetMechanism() != nil {
		if err := m.GetMechanism().IsValid(); err != nil {
			return errors.Wrapf(err, "invalid Mechanism in %v", m)
		}
	}

	if err := m.GetPath().IsValid(); err != nil {
		return err
	}

	return nil
}

// IsComplete checks if connection is complete valid
func (m *Connection) IsComplete() error {
	if err := m.IsValid(); err != nil {
		return err
	}

	if m.GetId() == "" {
		return errors.Errorf("connection.Id cannot be empty: %v", m)
	}

	if err := m.GetContext().IsValid(); err != nil {
		return err
	}

	return nil
}

// MatchesMonitorScopeSelector - Returns true of the connection matches the selector
func (m *Connection) MatchesMonitorScopeSelector(selector *MonitorScopeSelector) bool {
	if m == nil {
		return false
	}
	// Note: Empty selector always matches a non-nil Connection
	if len(selector.GetPathSegments()) == 0 {
		return true
	}
	// Iterate through the Connection.NetworkServiceManagers array looking for a subarray that matches
	// the selector.NetworkServiceManagers array, treating "" in the selector.NetworkServiceManagers array
	// as a wildcard
	for i := range m.GetPath().GetPathSegments() {
		// If there aren't enough elements left in the Connection.NetworkServiceManagers array to match
		// all of the elements in the select.NetworkServiceManager array...clearly we can't match
		if i+len(selector.GetPathSegments()) > len(m.GetPath().GetPathSegments()) {
			return false
		}
		// Iterate through the selector.NetworkServiceManagers array to see is the subarray starting at
		// Connection.NetworkServiceManagers[i] matches selector.NetworkServiceManagers
		for j := range selector.GetPathSegments() {
			// "" matches as a wildcard... failure to match either as wildcard or exact match means the subarray
			// starting at Connection.NetworkServiceManagers[i] doesn't match selectors.NetworkServiceManagers
			if selector.GetPathSegments()[j].GetName() != "" && m.GetPath().GetPathSegments()[i+j].GetName() != selector.GetPathSegments()[j].GetName() {
				break
			}
			// If this is the last element in the selector.NetworkServiceManagers array and we still are matching...
			// return true
			if j == len(selector.GetPathSegments())-1 {
				return true
			}
		}
	}
	return false
}

// FilterMapOnManagerScopeSelector - Filters out of map[string]*Connection Connections not matching the selector
func FilterMapOnManagerScopeSelector(c map[string]*Connection, selector *MonitorScopeSelector) map[string]*Connection {
	rv := make(map[string]*Connection)
	for k, v := range c {
		if v != nil && v.MatchesMonitorScopeSelector(selector) {
			rv[k] = v
		}
	}
	return rv
}
