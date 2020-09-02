// Copyright (c) 2020 Cisco Systems, Inc.
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

	// NetNSInodeKey - netns inode mechanism property key
	NetNSInodeKey = "netnsInode"
	// Workspace - NSM workspace location mechanism property key
	Workspace = "workspace"
	// InterfaceNameKey - interface name mechanism property key
	InterfaceNameKey = "name"

	// InodeURL - url of the for inode://${dev}/${ino} to represent various netns, memif socket files etc
	InodeURL = "inodeURL"

	// PCIAddressKey - PCI address of the device for the SR-IOV supported mechanisms
	PCIAddressKey = "pciAddress"
)
