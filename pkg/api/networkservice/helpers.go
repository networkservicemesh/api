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
	"github.com/pkg/errors"
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

// SetRequestConnection sets request connection
func (x *NetworkServiceRequest) SetRequestConnection(conn *Connection) *NetworkServiceRequest {
	if x != nil {
		x.Connection = conn
	}
	return x
}

// GetRequestMechanismPreferences returns request mechanism preferences
func (x *NetworkServiceRequest) GetRequestMechanismPreferences() []*Mechanism {
	preferences := make([]*Mechanism, 0, len(x.MechanismPreferences))
	preferences = append(preferences, x.MechanismPreferences...)

	return preferences
}

// SetRequestMechanismPreferences sets request mechanism preferences
func (x *NetworkServiceRequest) SetRequestMechanismPreferences(mechanismPreferences []*Mechanism) {
	x.MechanismPreferences = mechanismPreferences
}

// IsValid returns if request is valid
func (x *NetworkServiceRequest) IsValid() error {
	if x == nil {
		return errors.New("request cannot be nil")
	}

	if x.GetConnection() == nil {
		return errors.Errorf("request.Connection cannot be nil %v", x)
	}

	if err := x.GetConnection().IsValid(); err != nil {
		return errors.Errorf("request.Connection is invalid: %s: %v", err, x)
	}

	if x.GetMechanismPreferences() == nil {
		return errors.Errorf("request.MechanismPreferences cannot be nil: %v", x)
	}

	if len(x.GetMechanismPreferences()) < 1 {
		return errors.Errorf("request.MechanismPreferences must have at least one entry: %v", x)
	}

	return nil
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
