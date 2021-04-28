// Copyright (c) 2020-2021 Cisco and/or its affiliates.
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

	"google.golang.org/protobuf/proto"
)

// GetSrcIPNet - GetSrcIpAddr() converted to *net.IPNet or nil if empty or cannot be parsed
func (i *IPContext) GetSrcIPNet() []*net.IPNet {
	return strsToIPNets(i.GetSrcIpAddr())
}

// GetDstIPNet - GetDstIpAddr() converted to *net.IPNet or nil if empty or cannot be parsed
func (i *IPContext) GetDstIPNet() []*net.IPNet {
	return strsToIPNets(i.GetDstIpAddr())
}

// GetExcludedPrefixesIPNet - GetExcludedPrefixes() converted to []*net.IPNet prefixes that are empty or cannot be parsed are omitted
func (i *IPContext) GetExcludedPrefixesIPNet() []*net.IPNet {
	var prefixes []*net.IPNet
	for _, prefixStr := range i.GetExcludedPrefixes() {
		prefixes = append(prefixes, strToIPNet(prefixStr))
	}
	return prefixes
}

// GetGetExtraPrefixesIPNet - GetExtraPrefixes() converted to []*net.IPNet prefixes that are empty or cannot be parsed are omitted
func (i *IPContext) GetGetExtraPrefixesIPNet() []*net.IPNet {
	var prefixes []*net.IPNet
	for _, prefixStr := range i.GetExtraPrefixes() {
		prefixes = append(prefixes, strToIPNet(prefixStr))
	}
	return prefixes
}

// GetIP - GetIp() - converted to *net.IP or nil if empty or cannot be parsed
func (n *IpNeighbor) GetIP() net.IP {
	return net.ParseIP(n.GetIp())
}

// GetPrefixIPNet - GetPrefix() converted to *net.IPNet or nil if empty or cannot be parsed
func (r *Route) GetPrefixIPNet() *net.IPNet {
	return strToIPNet(r.GetPrefix())
}

// GetNextHopIP - GetNextHop() converted to net.IP or nil if empty or cannot be parsed
func (r *Route) GetNextHopIP() net.IP {
	return net.ParseIP(r.GetNextHop())
}

// Clone clones route
func (r *Route) Clone() *Route {
	// use as proto.Message
	return proto.Clone(r).(*Route)
}

func strsToIPNets(in []string) (out []*net.IPNet) {
	for _, str := range in {
		if ip := strToIPNet(str); ip != nil {
			out = append(out, ip)
		}
	}
	return out
}

func strToIPNet(in string) *net.IPNet {
	if in == "" {
		return nil
	}
	ip, ipNet, err := net.ParseCIDR(in)
	if err != nil {
		return nil
	}
	ipNet.IP = ip
	return ipNet
}
