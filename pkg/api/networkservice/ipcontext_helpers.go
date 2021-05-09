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

// GetSrcIPNets - GetSrcIpAddrs() converted to *net.IPNet or nil if empty or cannot be parsed
func (i *IPContext) GetSrcIPNets() []*net.IPNet {
	return strsToIPNets(i.GetSrcIpAddrs())
}

// GetDstIPNets - GetDstIpAddrs() converted to *net.IPNet or nil if empty or cannot be parsed
func (i *IPContext) GetDstIPNets() []*net.IPNet {
	return strsToIPNets(i.GetDstIpAddrs())
}

// GetDstRoutesWithExplicitNextHop - returns routes with the Route.NextHop explicitly set to the first IP address
// of the same address family (IPv4 or IPv6) from SrcIPAddrs if and only if Route.NextHop was initially nil and
// the Route.NextHop is not in DstIPAddrs
func (i *IPContext) GetDstRoutesWithExplicitNextHop() (routes []*Route) {
	srcIPs := i.GetSrcIPNets()
	ipv6Nets := filterIPNetByFamily(IpFamily_IPV6, srcIPs)
	ipv4Nets := filterIPNetByFamily(IpFamily_IPV4, srcIPs)
	return getRoutesWithExplicitNextHop(i.GetDstRoutes(), ipv4Nets, ipv6Nets)
}

// GetSrcRoutesWithExplicitNextHop - returns routes with the Route.NextHop explicitly set to the first IP address
// of the same address family (IPv4 or IPv6) from DstIPAddrs if and only if Route.NextHop was initially nil
func (i *IPContext) GetSrcRoutesWithExplicitNextHop() (routes []*Route) {
	// Set nextHop for any Route that is missing them
	dstIPs := i.GetDstIPNets()
	ipv6Nets := filterIPNetByFamily(IpFamily_IPV6, dstIPs)
	ipv4Nets := filterIPNetByFamily(IpFamily_IPV4, dstIPs)
	return getRoutesWithExplicitNextHop(i.GetSrcRoutes(), ipv4Nets, ipv6Nets)
}

// GetSrcIPRoutes - returns routes for any SrcIPs that are not contained in the prefixes of at least one DstIP
func (i *IPContext) GetSrcIPRoutes() (routes []*Route) {
	for _, srcIPNet := range i.GetSrcIPNets() {
		if srcIPNet == nil {
			continue
		}
		if contains(i.GetDstIPNets(), srcIPNet.IP) {
			continue
		}
		routes = append(routes, &Route{
			Prefix: srcIPNet.String(),
		})
	}
	return routes
}

// GetDstIPRoutes - returns routes for any DstIPs that are not contained in the prefixes of at least one SrcIP
func (i *IPContext) GetDstIPRoutes() (routes []*Route) {
	for _, dstIPNet := range i.GetDstIPNets() {
		if dstIPNet == nil {
			continue
		}
		if contains(i.GetSrcIPNets(), dstIPNet.IP) {
			continue
		}
		routes = append(routes, &Route{
			Prefix: dstIPNet.String(),
		})
	}
	return routes
}

func getRoutesWithExplicitNextHop(inRoutes []*Route, toIPv4Nets, toIPv6Nets []*net.IPNet) (routes []*Route) {
	ipv6NextHop := getNextHop(toIPv6Nets)
	ipv4NextHop := getNextHop(toIPv4Nets)
	for _, route := range inRoutes {
		if route.GetPrefixIPNet() != nil && route.GetNextHopIP() == nil {
			if ipv4NextHop != nil && route.GetPrefixIPNet().IP.To4() != nil && !contains(toIPv4Nets, route.GetPrefixIPNet().IP) {
				route = route.Clone()
				route.NextHop = ipv4NextHop.String()
			}
			if ipv6NextHop != nil && route.GetPrefixIPNet().IP.To4() == nil && !contains(toIPv6Nets, route.GetPrefixIPNet().IP) {
				route = route.Clone()
				route.NextHop = ipv6NextHop.String()
			}
		}
		routes = append(routes, route)
	}
	return routes
}

func getNextHop(ipNets []*net.IPNet) net.IP {
	if len(ipNets) > 0 && ipNets[0] != nil {
		return ipNets[0].IP
	}
	return nil
}

func filterIPNetByFamily(family IpFamily_Family, ipNets []*net.IPNet) []*net.IPNet {
	var rv []*net.IPNet
	for _, ipNet := range ipNets {
		if ipNet != nil && family == IpFamily_IPV4 && ipNet.IP.To4() != nil {
			rv = append(rv, ipNet)
		}
		if ipNet != nil && family == IpFamily_IPV6 && ipNet.IP.To4() == nil {
			rv = append(rv, ipNet)
		}
	}
	return rv
}

// contains - returns true if ip is contained any any of the supplied prefixes
func contains(prefixes []*net.IPNet, ip net.IP) bool {
	for _, prefix := range prefixes {
		if prefix != nil && prefix.Contains(ip) {
			return true
		}
	}
	return false
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
