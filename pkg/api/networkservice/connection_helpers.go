// Copyright (c) 2018-2021 VMware, Inc.
//
// Copyright (c) 2021 Doc.ai and/or its affiliates.
//
// Copyright (c) 2023-2024 Cisco and/or its affiliates.
//
// Copyright (c) 2024 Nordix Foundation.
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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

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

// Checks if the connection matches any of the network services in the given
// network service selector.
// Treats an empty NetworkServices in the selector as a wildcard.
func (x *Connection) matchesNetworkServices(selector *MonitorScopeSelector) bool {
	selectorServices := selector.GetNetworkServices()
	if len(selectorServices) == 0 {
		return true // Wildcard match
	}

	for _, service := range selectorServices {
		if service == x.GetNetworkService() {
			return true // Match found
		}
	}

	return false
}

// Iterates through the connection's PathSegments array looking for a subarray
// that matches the selector's PathSegments array.
// Treats an empty PathSegments in the selector or empty strings in the selector
// as wildcards.
func (x *Connection) matchesPathSegments(selector *MonitorScopeSelector) bool {
	if len(selector.GetPathSegments()) == 0 {
		return true // Wildcard match
	}

	for i := range x.GetPath().GetPathSegments() {
		// If there aren't enough elements left in the Connection.PathSegments array to match
		// all of the elements in the selector.PathSegments array...clearly we can't match
		if i+len(selector.GetPathSegments()) > len(x.GetPath().GetPathSegments()) {
			return false
		}
		// Iterate through the selector.PathSegments array to see is the subarray starting at
		// Connection.PathSegments[i] matches selector.PathSegments
		for j := range selector.GetPathSegments() {
			// "" matches as a wildcard... failure to match either as wildcard or exact match means the subarray
			// starting at Connection.PathSegments[i] doesn't match selectors.PathSegments
			if selector.GetPathSegments()[j].GetName() != "" && x.GetPath().GetPathSegments()[i+j].GetName() != selector.GetPathSegments()[j].GetName() {
				break
			}

			if selector.GetPathSegments()[j].GetId() != "" && x.GetPath().GetPathSegments()[i+j].GetId() != selector.GetPathSegments()[j].GetId() {
				break
			}

			if selector.GetPathSegments()[j].GetToken() != "" && x.GetPath().GetPathSegments()[i+j].GetToken() != selector.GetPathSegments()[j].GetToken() {
				break
			}

			// If this is the last element in the selector.PathSegments array and we still are matching...
			// return true
			if j == len(selector.GetPathSegments())-1 {
				return true
			}
		}
	}

	return false
}

// MatchesMonitorScopeSelector - Returns true of the connection matches the selector
func (x *Connection) MatchesMonitorScopeSelector(selector *MonitorScopeSelector) bool {
	if x == nil {
		return false
	}
	// Empty selector matches any non-nil connection
	if len(selector.GetPathSegments()) == 0 && len(selector.GetNetworkServices()) == 0 {
		return true
	}
	// Check network service match first
	if !x.matchesNetworkServices(selector) {
		return false
	}
	// Check path segments match
	return x.matchesPathSegments(selector)
}

// GetCurrentPathSegment - Get the current path segment of the connection
func (x *Connection) GetCurrentPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if int(x.GetPath().GetIndex()) > len(x.GetPath().GetPathSegments())-1 {
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
