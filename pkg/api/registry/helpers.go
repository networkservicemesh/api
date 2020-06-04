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

// NewEndpointNSMName - construct an NewEndpointNSMName from endpoint and manager
func NewEndpointNSMName(endpoint, manager *NetworkServiceEndpoint) EndpointNSMName {
	return EndpointNSMName(endpoint.Name + ":" + manager.Url)
}

// ServiceNames - returns grpc ServiceNames implemented by impl
func ServiceNames(impl interface{}, existingServiceNames ...string) []string {
	if _, ok := impl.(NetworkServiceRegistryServer); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceRegistry_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceRegistryClient); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceRegistry_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceEndpointRegistryServer); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceEndpointRegistry_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceEndpointRegistryClient); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceEndpointRegistry_serviceDesc.ServiceName)
	}
	return existingServiceNames
}
