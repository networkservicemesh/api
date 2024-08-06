// Copyright (c) 2020-2024 Cisco and/or its affiliates.
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
)

// Clone clones request
func (x *NetworkServiceRequest) Clone() *NetworkServiceRequest {
	return proto.Clone(x).(*NetworkServiceRequest)
}

// GetRequestConnection returns request connection
func (x *NetworkServiceRequest) GetRequestConnection() *Connection {
	return x.GetConnection()
}

// GetRequestMechanismPreferences returns request mechanism preferences
func (x *NetworkServiceRequest) GetRequestMechanismPreferences() []*Mechanism {
	preferences := make([]*Mechanism, 0, len(x.MechanismPreferences))
	preferences = append(preferences, x.MechanismPreferences...)

	return preferences
}

// ServiceNames - returns grpc ServiceNames implemented by impl
func ServiceNames(impl interface{}, existingServiceNames ...string) []string {
	if _, ok := impl.(NetworkServiceServer); ok {
		existingServiceNames = append(existingServiceNames, _NetworkService_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceClient); ok {
		existingServiceNames = append(existingServiceNames, _NetworkService_serviceDesc.ServiceName)
	}
	if _, ok := impl.(MonitorConnectionServer); ok {
		existingServiceNames = append(existingServiceNames, _MonitorConnection_serviceDesc.ServiceName)
	}
	if _, ok := impl.(MonitorConnectionClient); ok {
		existingServiceNames = append(existingServiceNames, _MonitorConnection_serviceDesc.ServiceName)
	}
	return existingServiceNames
}
