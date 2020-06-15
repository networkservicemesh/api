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

// ReadNetworkServiceList read list of NetworkServices from passed stream
func ReadNetworkServiceList(stream NetworkServiceRegistry_FindClient) []*NetworkService {
	var result []*NetworkService
	for msg, err := stream.Recv(); true; msg, err = stream.Recv() {
		if err != nil {
			break
		}
		result = append(result, msg)
	}
	return result
}

// ReadNetworkServiceEndpointList read list of NetworkServiceEndpoints from passed stream
func ReadNetworkServiceEndpointList(stream NetworkServiceEndpointRegistry_FindClient) []*NetworkServiceEndpoint {
	var result []*NetworkServiceEndpoint
	for msg, err := stream.Recv(); true; msg, err = stream.Recv() {
		if err != nil {
			break
		}
		result = append(result, msg)
	}
	return result
}
