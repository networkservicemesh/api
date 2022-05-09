// Copyright (c) 2020-2021 Cisco Systems, Inc.
//
// Copyright (c) 2021 Doc.ai and/or its affiliates.
//
// Copyright (c) 2022 Xored Software Inc and others.
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

package kernel

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
)

const (

	// MECHANISM string
	MECHANISM = "KERNEL"

	// Parameters

	// LinuxIfMaxLength - Linux has a limit of 15 characters for an interface name
	LinuxIfMaxLength = 15

	// PCIAddressKey - device PCI address property key
	PCIAddressKey = common.PCIAddressKey

	// DeviceTokenIDKey is a device token ID property key
	DeviceTokenIDKey = common.DeviceTokenIDKey

	// NetNSInodeKey - netns inode mechanism property key
	NetNSInodeKey = common.NetNSInodeKey

	// InterfaceNameKey - interface name mechanism property key
	InterfaceNameKey = common.InterfaceNameKey

	// NetNSURL - NetNS URL, it can be either:
	// * file:///proc/${pid}/ns/net - ${pid} process net NS
	// * inode://${dev}/${ino} - while transferring file between processes using grpcfd
	NetNSURL = common.InodeURL

	// NetNSURLScheme - expected scheme of NetNSURLs
	NetNSURLScheme = "file"

	// SupportsVLAN - flag set if the forwarder supports VLAN trunking
	SupportsVLAN = "supportsVlan"

	// VLAN - VLAN ID
	VLAN = "vlan"

	// RouteLocalNet - flag set if route_localnet enabled
	RouteLocalNet = "routeLocalNet"

	// IPTables4NatTemplate - IP Tables ipv4 chains/rules template
	IPTables4NatTemplate = "IPTables4NatTemplate"
)
