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

package registry

// EndpointNSMName -  - a type to hold endpoint and nsm url composite type.
type EndpointNSMName string

// GetEndpointNSMName - return a Endpoint.Name + ":" + NetworkServiceManager.Url
func (nse *NSERegistration) GetEndpointNSMName() EndpointNSMName {
	if nse == nil {
		return ""
	}
	return NewEndpointNSMName(nse.NetworkServiceEndpoint, nse.NetworkServiceManager)
}

// NewEndpointNSMName - construct an NewEndpointNSMName from endpoint and manager
func NewEndpointNSMName(endpoint *NetworkServiceEndpoint, manager *NetworkServiceManager) EndpointNSMName {
	return EndpointNSMName(endpoint.Name + ":" + manager.Url)
}
