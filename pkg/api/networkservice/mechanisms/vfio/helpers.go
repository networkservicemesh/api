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

package vfio

import (
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
)

// Mechanism is a vfio mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// New returns *networkservice.Mechanism of type vfio with the given cgroupDir
func New(cgroupDir string) *networkservice.Mechanism {
	return &networkservice.Mechanism{
		Cls:  cls.LOCAL,
		Type: MECHANISM,
		Parameters: map[string]string{
			CgroupDirKey: cgroupDir,
		},
	}
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

// GetParameters returns the map of all parameters to the mechanism
func (m *Mechanism) GetParameters() map[string]string {
	if m == nil {
		return map[string]string{}
	}
	if m.Parameters == nil {
		m.Parameters = map[string]string{}
	}
	return m.Parameters
}

// GetCgroupDir returns client on host cgroup directory
func (m *Mechanism) GetCgroupDir() string {
	return m.GetParameters()[CgroupDirKey]
}

// SetCgroupDir sets client on host cgroup directory
func (m *Mechanism) SetCgroupDir(cgroupDir string) {
	m.GetParameters()[CgroupDirKey] = cgroupDir
}

// GetIommuGroup returns IOMMU group id
func (m *Mechanism) GetIommuGroup() uint {
	return atou(m.GetParameters()[IommuGroupKey])
}

// SetIommuGroup sets IOMMU group id
func (m *Mechanism) SetIommuGroup(iommuGroup uint) {
	m.GetParameters()[IommuGroupKey] = utoa(iommuGroup)
}

// GetPCIAddress returns PCI address
func (m *Mechanism) GetPCIAddress() string {
	return m.GetParameters()[PCIAddressKey]
}

// SetPCIAddress sets PCI address
func (m *Mechanism) SetPCIAddress(pciAddress string) {
	m.GetParameters()[PCIAddressKey] = pciAddress
}

// GetDeviceTokenID returns device token ID
func (m *Mechanism) GetDeviceTokenID() string {
	return m.Parameters[DeviceTokenIDKey]
}

// SetDeviceTokenID sets device token ID
func (m *Mechanism) SetDeviceTokenID(tokenID string) {
	m.Parameters[DeviceTokenIDKey] = tokenID
}

// GetVfioMajor returns /dev/vfio major number
func (m *Mechanism) GetVfioMajor() uint32 {
	return atou32(m.GetParameters()[VfioMajorKey])
}

// SetVfioMajor sets /dev/vfio major number
func (m *Mechanism) SetVfioMajor(vfioMajor uint32) {
	m.GetParameters()[VfioMajorKey] = u32toa(vfioMajor)
}

// GetVfioMinor returns /dev/vfio minor number
func (m *Mechanism) GetVfioMinor() uint32 {
	return atou32(m.GetParameters()[VfioMinorKey])
}

// SetVfioMinor sets /dev/vfio minor number
func (m *Mechanism) SetVfioMinor(vfioMinor uint32) {
	m.GetParameters()[VfioMinorKey] = u32toa(vfioMinor)
}

// GetDeviceMajor returns /dev/${iommuGroup} major number
func (m *Mechanism) GetDeviceMajor() uint32 {
	return atou32(m.GetParameters()[DeviceMajorKey])
}

// SetDeviceMajor sets /dev/${iommuGroup} major number
func (m *Mechanism) SetDeviceMajor(deviceMajor uint32) {
	m.GetParameters()[DeviceMajorKey] = u32toa(deviceMajor)
}

// GetDeviceMinor returns /dev/${iommuGroup} minor number
func (m *Mechanism) GetDeviceMinor() uint32 {
	return atou32(m.GetParameters()[DeviceMinorKey])
}

// SetDeviceMinor sets /dev/${iommuGroup} minor number
func (m *Mechanism) SetDeviceMinor(deviceMinor uint32) {
	m.GetParameters()[DeviceMinorKey] = u32toa(deviceMinor)
}

func atou(a string) uint {
	u, err := strconv.ParseUint(a, 10, strconv.IntSize)
	if err != nil {
		return 0
	}
	return uint(u)
}

func utoa(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func atou32(a string) uint32 {
	u, err := strconv.ParseUint(a, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(u)
}

func u32toa(u uint32) string {
	return strconv.FormatUint(uint64(u), 10)
}
