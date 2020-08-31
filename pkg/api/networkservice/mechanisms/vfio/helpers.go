// Copyright (c) 2020 Intel Corporation. All Rights Reserved.
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

package vfio

import (
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - vfio mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism converts unified mechanism to helper
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == MECHANISM {
		return &Mechanism{
			m,
		}
	}
	return nil
}

// GetCgroupDir returns client on host cgroup directory
func (m *Mechanism) GetCgroupDir() string {
	if m.Parameters == nil {
		return ""
	}
	return m.GetParameters()[CgroupDirKey]
}

// SetCgroupDir sets client on host cgroup directory
func (m *Mechanism) SetCgroupDir(cgroupDir string) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[CgroupDirKey] = cgroupDir
}

// GetIommuGroup returns IOMMU group id
func (m *Mechanism) GetIommuGroup() uint {
	if m.Parameters == nil {
		return 0
	}
	return atou(m.Parameters[IommuGroupKey])
}

// SetIommuGroup sets IOMMU group id
func (m *Mechanism) SetIommuGroup(iommuGroup uint) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[IommuGroupKey] = utoa(iommuGroup)
}

// GetPCIAddress returns PCI address
func (m *Mechanism) GetPCIAddress() string {
	if m.Parameters == nil {
		return ""
	}
	return m.GetParameters()[PCIAddressKey]
}

// SetPCIAddress sets PCI address
func (m *Mechanism) SetPCIAddress(pciAddress string) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[PCIAddressKey] = pciAddress
}

// GetVfioMajor returns /dev/vfio major number
func (m *Mechanism) GetVfioMajor() uint32 {
	if m.Parameters == nil {
		return 0
	}
	return uint32(atou(m.Parameters[VfioMajorKey]))
}

// SetVfioMajor sets /dev/vfio major number
func (m *Mechanism) SetVfioMajor(vfioMajor uint32) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[VfioMajorKey] = utoa(uint(vfioMajor))
}

// GetVfioMinor returns /dev/vfio minor number
func (m *Mechanism) GetVfioMinor() uint32 {
	if m.Parameters == nil {
		return 0
	}
	return uint32(atou(m.Parameters[VfioMinorKey]))
}

// SetVfioMinor sets /dev/vfio minor number
func (m *Mechanism) SetVfioMinor(vfioMinor uint32) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[VfioMinorKey] = utoa(uint(vfioMinor))
}

// GetDeviceMajor returns /dev/${igid} major number
func (m *Mechanism) GetDeviceMajor() uint32 {
	if m.Parameters == nil {
		return 0
	}
	return uint32(atou(m.Parameters[DeviceMajorKey]))
}

// SetDeviceMajor sets /dev/${igid} major number
func (m *Mechanism) SetDeviceMajor(deviceMajor uint32) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[DeviceMajorKey] = utoa(uint(deviceMajor))
}

// GetDeviceMinor returns /dev/${igid} minor number
func (m *Mechanism) GetDeviceMinor() uint32 {
	if m.Parameters == nil {
		return 0
	}
	return uint32(atou(m.Parameters[DeviceMinorKey]))
}

// SetDeviceMinor sets /dev/${igid} minor number
func (m *Mechanism) SetDeviceMinor(deviceMinor uint32) {
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	m.Parameters[DeviceMinorKey] = utoa(uint(deviceMinor))
}

func atou(a string) uint {
	u, err := strconv.ParseUint(a, 10, 0)
	if err != nil {
		return 0
	}
	return uint(u)
}

func utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
