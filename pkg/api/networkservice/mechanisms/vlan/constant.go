// Package vlan provides helper methods for the Mechanism vlan
package vlan

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
)

const (
	// MECHANISM string
	MECHANISM = "VLAN"

	// Mechanism parameters

	// LinuxIfMaxLength - Linux has a limit of 15 characters for an interface name
	LinuxIfMaxLength = 15

	// InterfaceNameKey - interface name mechanism property key
	InterfaceNameKey = common.InterfaceNameKey

	// NetNSURL - url representing an inode - fmt.Sprintf("inode://%d/%d",dev,ino)
	NetNSURL = common.InodeURL

	// NetNSURLScheme - expected scheme of NetNSURLs
	NetNSURLScheme = "file"

	// VlanID - vlan ID
	VlanID = "vlan-id"

	// MTU - maximum transmission unit
	BaseInterfaceNameKey = "base"
)
