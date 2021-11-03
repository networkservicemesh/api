// Copyright (c) 2020-2021 Cisco Systems, Inc.
//
// Copyright (c) 2021 Doc.ai and/or its affiliates.
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

// Package common constants for Mechanism parameters shared between multiple Mechanism types
package common

const (
	// SrcIP - key for SrcIP parameters
	SrcIP = "src_ip"
	// DstIP - key for DstIP parameters
	DstIP = "dst_ip"

	// SrcPort - key for SrcPort parameters
	SrcPort = "src_port"
	// DstPort - ley for DstPort parameters
	DstPort = "dst_port"

	// SrcOriginalIP - original src IP
	SrcOriginalIP = "orig_src_ip"

	// DstOriginalIP - original destination ip
	DstOriginalIP = "orig_dst_ip"

	// NetNSInodeKey - netns inode mechanism property key
	NetNSInodeKey = "netnsInode"

	// InterfaceNameKey - interface name mechanism property key
	InterfaceNameKey = "name"

	// InodeURL - file:// or inode:// URL representing some file shared between processes with grpcfd (e.g. netns file)
	InodeURL = "inodeURL"

	// PCIAddressKey - PCI address of the device for the SR-IOV supported mechanisms
	PCIAddressKey = "pciAddress"

	// DeviceTokenIDKey - Client/Endpoint device token ID
	// nolint:gosec
	DeviceTokenIDKey = "tokenID"

	// MTU - Maximum Transmission Unit
	MTU = "MTU"
)
