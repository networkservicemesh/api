// Copyright (c) 2020-2023 Cisco and/or its affiliates.
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

import "google.golang.org/protobuf/proto"

// Clone clones ConnectionEvents
func (x *ConnectionEvent) Clone() *ConnectionEvent {
	return proto.Clone(x).(*ConnectionEvent)
}

// GetEventSenderName returns name of the segment from the path who send the event
func (x *ConnectionEvent) GetEventSenderName() string {
	for v, k := range x.GetConnections() {
		if segment := k.GetPathSegmentByID(v); segment != nil {
			return segment.GetName()
		}
	}
	return "unknown"
}
