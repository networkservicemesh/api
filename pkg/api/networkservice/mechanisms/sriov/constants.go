// Copyright (c) 2020 Doc.ai and/or its affiliates.
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

// Package sriov provides common parameters for SR-IOV mechanisms
package sriov

const (
	// EndpointVfPCIAddressKey is an endpoint VF PCI address property key
	EndpointVfPCIAddressKey = "endpointVfPCIAddress"

	// ClientPfPCIAddressKey is a client PF PCI address property key
	ClientPfPCIAddressKey = "clientPfPCIAddress"

	// ClientVfPCIAddressKey is a client VF PCI address property key
	ClientVfPCIAddressKey = "clientVfPCIAddress"

	// IommuGroupKey is a IOMMU group id property key
	IommuGroupKey = "iommuGroup"
)
