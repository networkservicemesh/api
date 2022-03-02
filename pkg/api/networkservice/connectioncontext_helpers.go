// Copyright (c) 2020-2022 Cisco and/or its affiliates.
// Copyright (c) 2022 Nordix Foundation
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

package networkservice

import (
	"net"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// IsEthernetContextEmtpy returns true if ethernet config is empty
func (x *ConnectionContext) IsEthernetContextEmtpy() bool {
	return x.EthernetContext == nil || (x.EthernetContext.SrcMac == "" && x.EthernetContext.DstMac == "")
}

// IsValid - checks ConnectionContext validation
func (x *ConnectionContext) IsValid() error {
	if x == nil {
		return errors.New("ConnectionContext should not be nil")
	}
	ip := x.GetIpContext()
	for _, route := range append(ip.GetSrcRoutes(), ip.GetDstRoutes()...) {
		if route.GetPrefix() == "" {
			return errors.Errorf("ConnectionContext.Route.Prefix is required and cannot be empty/nil: %v", ip)
		}
		_, _, err := net.ParseCIDR(route.GetPrefix())
		if err != nil {
			return errors.Errorf("ConnectionContext.Route.Prefix should be a valid CIDR address: %v", ip)
		}
	}

	for _, neighbor := range ip.GetIpNeighbors() {
		if neighbor.GetIp() == "" {
			return errors.Errorf("ConnectionContext.IpNeighbors.Ip is required and cannot be empty/nil: %v", ip)
		}
		if neighbor.GetHardwareAddress() == "" {
			return errors.Errorf("ConnectionContext.IpNeighbors.HardwareAddress is required and cannot be empty/nil: %v", ip)
		}
	}
	return nil
}

// MeetsRequirements - checks required context parameters have bin set
func (x *ConnectionContext) MeetsRequirements(original *ConnectionContext) error {
	if x == nil {
		return errors.New("ConnectionContext should not be nil")
	}

	err := x.IsValid()
	if err != nil {
		return err
	}
	if original.GetIpContext().GetDstIpRequired() && len(x.GetIpContext().GetDstIpAddrs()) > 0 {
		return errors.Errorf("ConnectionContext.DestIp is required and cannot be empty/nil: %v", x)
	}
	if original.GetIpContext().GetSrcIpRequired() && len(x.GetIpContext().GetSrcIpAddrs()) > 0 {
		return errors.Errorf("ConnectionContext.SrcIp is required cannot be empty/nil: %v", x)
	}

	return nil
}

// Validate - checks DNSConfig and returns error if DNSConfig is not valid
func (c *DNSConfig) Validate() error {
	if c == nil {
		return errors.New(DNSConfigShouldNotBeNil)
	}
	if len(c.DnsServerIps) == 0 {
		return errors.New(DNSServerIpsShouldHaveRecords)
	}
	return nil
}

// IsValid - checks ExtraPrefixRequest validation
func (c *ExtraPrefixRequest) IsValid() error {
	if c == nil {
		return errors.New("ExtraPrefixRequest should not be nil")
	}

	if c.RequiredNumber < 1 {
		return errors.Errorf("ExtraPrefixRequest.RequiredNumber should be positive number >=1: %v", c)
	}
	if c.RequestedNumber < 1 {
		return errors.Errorf("ExtraPrefixRequest.RequestedNumber should be positive number >=1: %v", c)
	}

	if c.RequiredNumber > c.RequestedNumber {
		return errors.Errorf("ExtraPrefixRequest.RequiredNumber should be less or equal to ExtraPrefixRequest.RequestedNumber >=1: %v", c)
	}

	if c.PrefixLen < 1 {
		return errors.Errorf("ExtraPrefixRequest.PrefixLen should be positive number >=1: %v", c)
	}

	// Check protocols
	if c.AddrFamily == nil {
		return errors.Errorf("ExtraPrefixRequest.AfFamily should not be nil: %v", c)
	}

	switch c.AddrFamily.Family {
	case IpFamily_IPV4:
		if c.PrefixLen > 32 {
			return errors.Errorf("ExtraPrefixRequest.PrefixLen should be positive number >=1 and <=32 for IPv4 %v", c)
		}
	case IpFamily_IPV6:
		if c.PrefixLen > 128 {
			return errors.Errorf("ExtraPrefixRequest.PrefixLen should be positive number >=1 and <=32 for IPv6 %v", c)
		}
	}

	return nil
}

// PortRange represents source port / destination port range.
type PortRange struct {
	Start uint16
	End   uint16
}

// ParsePortRange - parses port range in format "start-end" or "port".
func ParsePortRange(portRange string) (*PortRange, error) {
	if portRange == "" {
		return nil, nil
	}
	ports := strings.Split(portRange, "-")
	if len(ports) > 2 {
		return nil, errors.Errorf("port range should be in format start-end: %v", portRange)
	}
	start, err := strconv.ParseUint(ports[0], 10, 16)
	if err != nil {
		return nil, err
	}
	endString := ports[0]
	if len(ports) == 2 {
		endString = ports[1]
	}
	end, err := strconv.ParseUint(endString, 10, 16)
	if err != nil {
		return nil, err
	}

	return &PortRange{
		Start: uint16(start),
		End:   uint16(end),
	}, nil
}
