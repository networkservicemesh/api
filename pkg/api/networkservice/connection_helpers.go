// Copyright (c) 2018-2021 VMware, Inc.
//
// Copyright (c) 2021 Doc.ai and/or its affiliates.
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
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// IsRemote returns if connection is remote
func (x *Connection) IsRemote() bool {
	if x == nil {
		return false
	}
	// If we have two or more, it is remote
	return len(x.GetPath().GetPathSegments()) > 1
}

// GetSourceNetworkServiceManagerName - return source network service manager name
func (x *Connection) GetSourceNetworkServiceManagerName() string {
	if x == nil {
		return ""
	}
	if len(x.GetPath().GetPathSegments()) > 0 {
		return x.GetPath().GetPathSegments()[0].GetName()
	}
	return ""
}

// GetDestinationNetworkServiceManagerName - return destination network service manager name
func (x *Connection) GetDestinationNetworkServiceManagerName() string {
	if x == nil {
		return ""
	}
	if len(x.GetPath().GetPathSegments()) >= 2 {
		return x.GetPath().GetPathSegments()[1].GetName()
	}
	return ""
}

// Equals returns if connection equals given connection
func (x *Connection) Equals(connection protoreflect.ProtoMessage) bool {
	// use as proto.Message
	return proto.Equal(x, connection)
}

// Clone clones connection
func (x *Connection) Clone() *Connection {
	// use as proto.Message
	return proto.Clone(x).(*Connection)
}

// UpdateContext checks and tries to set connection context
func (x *Connection) UpdateContext(context *ConnectionContext) error {
	if err := context.MeetsRequirements(x.Context); err != nil {
		return err
	}

	oldContext := x.Context
	x.Context = context

	if err := x.IsValid(); err != nil {
		x.Context = oldContext
		return err
	}

	return nil
}

// IsValid checks if connection is minimally valid
func (x *Connection) IsValid() error {
	if x == nil {
		return errors.New("connection cannot be nil")
	}

	if x.GetNetworkService() == "" {
		return errors.Errorf("NetworkService cannot be empty: %v", x)
	}

	if x.GetMechanism() != nil {
		if err := x.GetMechanism().IsValid(); err != nil {
			return errors.Wrapf(err, "invalid Mechanism in %v", x)
		}
	}

	if err := x.GetPath().IsValid(); err != nil {
		return err
	}

	return nil
}

// IsComplete checks if connection is complete valid
func (x *Connection) IsComplete() error {
	if err := x.IsValid(); err != nil {
		return err
	}

	if x.GetId() == "" {
		return errors.Errorf("Id cannot be empty: %v", x)
	}

	if err := x.GetContext().IsValid(); err != nil {
		return err
	}

	return nil
}

// MatchesMonitorScopeSelector - Returns true of the connection matches the selector
func (x *Connection) MatchesMonitorScopeSelector(selector *MonitorScopeSelector) bool {
	if x == nil {
		return false
	}
	// Note: Empty selector always matches a non-nil Connection
	if len(selector.GetPathSegments()) == 0 {
		return true
	}
	// Iterate through the Connection.NetworkServiceManagers array looking for a subarray that matches
	// the selector.NetworkServiceManagers array, treating "" in the selector.NetworkServiceManagers array
	// as a wildcard
	for i := range x.GetPath().GetPathSegments() {
		// If there aren't enough elements left in the Connection.NetworkServiceManagers array to match
		// all of the elements in the select.NetworkServiceManager array...clearly we can't match
		if i+len(selector.GetPathSegments()) > len(x.GetPath().GetPathSegments()) {
			return false
		}
		// Iterate through the selector.NetworkServiceManagers array to see is the subarray starting at
		// Connection.NetworkServiceManagers[i] matches selector.NetworkServiceManagers
		for j := range selector.GetPathSegments() {
			// "" matches as a wildcard... failure to match either as wildcard or exact match means the subarray
			// starting at Connection.NetworkServiceManagers[i] doesn't match selectors.NetworkServiceManagers
			if selector.GetPathSegments()[j].GetName() != "" && x.GetPath().GetPathSegments()[i+j].GetName() != selector.GetPathSegments()[j].GetName() {
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

// GetCurrentPathSegment - Get the current path segment of the connection
func (x *Connection) GetCurrentPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if len(x.GetPath().GetPathSegments())-1 < int(x.GetPath().GetIndex()) {
		return nil
	}
	return x.GetPath().GetPathSegments()[x.GetPath().GetIndex()]
}

// GetPrevPathSegment - Get the previous path segment of the connection
func (x *Connection) GetPrevPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if int(x.GetPath().GetIndex()) == 0 {
		return nil
	}
	if int(x.GetPath().GetIndex())-1 > len(x.GetPath().GetPathSegments()) {
		return nil
	}
	return x.GetPath().GetPathSegments()[x.GetPath().GetIndex()-1]
}

// GetNextPathSegment - Get the next path segment of the connection
func (x *Connection) GetNextPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if len(x.GetPath().GetPathSegments())-1 < int(x.GetPath().GetIndex())+1 {
		return nil
	}
	return x.GetPath().GetPathSegments()[x.GetPath().GetIndex()+1]
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
