// Copyright (c) 2020 Intel Corporation. All Rights Reserved.
//
// Copyright (c) 2020-2021 Doc.ai and/or its affiliates.
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

// Package vfio provides a Mechanism for using vfio devices.
package vfio

import "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"

const (
	// MECHANISM string
	MECHANISM = "VFIO"

	// Parameters

	// CgroupDirKey is a client on host cgroup directory property key
	CgroupDirKey = "cgroupDir"

	// IommuGroupKey is a IOMMU group id property key
	IommuGroupKey = "iommuGroup"

	// PCIAddressKey is a PCI address property key
	PCIAddressKey = common.PCIAddressKey

	// DeviceTokenIDKey is a device token ID property key
	DeviceTokenIDKey = common.DeviceTokenIDKey

	// VfioMajorKey is a /dev/vfio major number property key
	VfioMajorKey = "vfioMajor"

	// VfioMinorKey is a /dev/vfio minor number property key
	VfioMinorKey = "vfioMinor"

	// DeviceMajorKey is a /dev/${IOMMU group id} major number property key
	DeviceMajorKey = "deviceMajor"

	// DeviceMinorKey is a /dev/${IOMMU group id} minor number property key
	DeviceMinorKey = "deviceMinor"
)
