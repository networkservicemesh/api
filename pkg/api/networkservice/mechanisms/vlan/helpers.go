package vlan

import (
	"fmt"
	"strconv"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// Mechanism - a vlan mechanism helper
type Mechanism struct {
	*networkservice.Mechanism
}

// ToMechanism - convert unified Mechanism to useful wrapper
func ToMechanism(m *networkservice.Mechanism) *Mechanism {
	if m.GetType() == MECHANISM {
		if m.Parameters == nil {
			m.Parameters = map[string]string{}
		}
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

func UpdateRequestParameters(umechanism *networkservice.Mechanism, ulabels map[string]string) {
	if baseName, ok := ulabels[BaseInterfaceNameKey]; ok {
		umechanism.Parameters[BaseInterfaceNameKey] = baseName
		delete(ulabels, BaseInterfaceNameKey)
	}
	if vlanId, ok := ulabels[VlanID]; ok {
		umechanism.Parameters[VlanID] = vlanId
		delete(ulabels, VlanID)
	}
}

// GetInterfaceName returns the Kernel Interface Name
//                  this is Mechanism.Parameters[InterfaceNameKey] if set
//                  otherwise returns a name computed from networkservice.Connection 'conn'
func (m *Mechanism) GetInterfaceName(conn *networkservice.Connection) string {
	return getIfName(conn, m.GetParameters()[InterfaceNameKey])
}

// SetInterfaceName sets the Kernel Interface Name
func (m *Mechanism) SetInterfaceName(interfaceName string) {
	m.GetParameters()[InterfaceNameKey] = interfaceName
}

// GetBaseInterfaceName returns the Kernel Interface Name
//                  this is Mechanism.Parameters[BaseInterfaceNameKey] if set
//                  otherwise returns a name computed from networkservice.Connection 'conn'
func (m *Mechanism) GetBaseInterfaceName(conn *networkservice.Connection) string {
	return getIfName(conn, m.GetParameters()[BaseInterfaceNameKey])
}

// SetInterfaceName sets the Kernel Base Interface Name
func (m *Mechanism) SetBaseInterfaceName(interfaceName string) {
	m.GetParameters()[BaseInterfaceNameKey] = interfaceName
}

func getIfName(conn *networkservice.Connection, storedValue string) string {
	// name not set; generate a name based on connection ID
	if storedValue == "" {
		ns := conn.GetNetworkService()
		nsMaxLength := LinuxIfMaxLength - 5
		if len(ns) > nsMaxLength {
			ns = ns[:nsMaxLength]
		}
		name := fmt.Sprintf("%s-%s", ns, conn.GetId())
		if len(name) > LinuxIfMaxLength {
			name = name[:LinuxIfMaxLength]
		}
		return name
	}

	// name is set; truncate to linux length limit
	name := storedValue
	if len(name) > LinuxIfMaxLength {
		name = name[:LinuxIfMaxLength]
	}
	return name
}

// GetNetNSURL returns the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *Mechanism) GetNetNSURL() string {
	return m.GetParameters()[NetNSURL]
}

// SetNetNSURL sets the NetNS URL - fmt.Sprintf("inode://%d/%d",dev,ino)
func (m *Mechanism) SetNetNSURL(urlString string) {
	m.GetParameters()[NetNSURL] = urlString
}

// VlanID returns the VlanID parameter of the Mechanism
func (m *Mechanism) VlanID() int {
	if m == nil {
		return 0
	}

	if m.GetParameters() == nil {
		return 0
	}

	vid := m.Parameters[VlanID]
	// vlan ID range is 0 to 4,095 stored in 12 bit
	vlanid, err := strconv.ParseInt(vid, 10, 12)

	if err != nil {
		return 0
	}

	return int(vlanid)
}

// SetVlanID - set the VNI for the vxlan tunnel and return *vxlan.Mechanism
func (m *Mechanism) SetVlanID(vlanid uint32) *Mechanism {
	if m == nil {
		return nil
	}
	m.GetParameters()[VlanID] = strconv.FormatUint(uint64(vlanid), 10)
	return m
}
