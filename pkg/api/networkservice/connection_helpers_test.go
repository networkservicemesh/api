// Copyright (c) 2023 Cisco and/or its affiliates.
//
// Copyright (c) 2024 Nordix Foundation.
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

package networkservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
)

// nolint: funlen
func TestMonitorScopeSelector(t *testing.T) {
	cases := []struct {
		testname         string
		connSegments     []*networkservice.PathSegment
		selectorSegments []*networkservice.PathSegment
		matches          bool
	}{
		{
			testname:         "IdenticalPaths",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			matches:          true,
		},
		{
			testname:         "DifferentNames",
			connSegments:     []*networkservice.PathSegment{{Name: "s15", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			matches:          false,
		},
		{
			testname:         "DifferentIds",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id15", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			matches:          false,
		},
		{
			testname:         "DifferentTokens",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t15"}, {Name: "s2", Id: "id2", Token: "t2"}},
			matches:          false,
		},
		{
			testname:         "SelectorPathIsLonger",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: "id2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: "id2"}, {Name: "s3", Id: "id3"}},
			matches:          false,
		},
		{
			testname:         "ConnPathContainsSelectorPath",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: "id2"}, {Name: "s3", Id: "id3"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: "id2"}},
			matches:          true,
		},
		{
			testname:         "EmptyID",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: "id2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: ""}},
			matches:          true,
		},
		{
			testname:         "EmptyName",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "s2", Id: "id2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1"}, {Name: "", Id: "id2"}},
			matches:          true,
		},
		{
			testname:         "EmptyToken",
			connSegments:     []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
			selectorSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "", Id: "id2", Token: ""}},
			matches:          true,
		},
	}

	for _, testCase := range cases {
		c := testCase
		t.Run(c.testname, func(t *testing.T) {
			path := &networkservice.Path{PathSegments: c.connSegments}
			conn := networkservice.Connection{Path: path}
			selector := &networkservice.MonitorScopeSelector{PathSegments: c.selectorSegments}

			require.Equal(t, conn.MatchesMonitorScopeSelector(selector), c.matches)
		})
	}
}

// nolint: funlen
func TestNetworkServiceMonitorScopeSelector(t *testing.T) {
	cases := []struct {
		testname string
		conn     *networkservice.Connection
		selector *networkservice.MonitorScopeSelector
		matches  bool
	}{
		{
			testname: "EmptySelector",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{},
			matches:  true,
		},
		{
			testname: "IdenticalPathsAndNetworkService",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				PathSegments:    []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				NetworkServices: []string{"ns1"},
			},
			matches: true,
		},
		{
			testname: "IdenticalPathsAndDifferentNetworkService",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				PathSegments:    []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				NetworkServices: []string{"ns2"},
			},
			matches: false,
		},
		{
			testname: "IdenticalPathsAndMatchingNetworkServiceList",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				PathSegments:    []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				NetworkServices: []string{"ns2", "ns1", "ns3"},
			},
			matches: true,
		},
		{
			testname: "IdenticalPathsAndNonMatchingNetworkServiceList",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				PathSegments:    []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				NetworkServices: []string{"ns2", "ns3", "ns0"},
			},
			matches: false,
		},
		{
			testname: "IdenticalNetworkService",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				NetworkServices: []string{"ns1"},
			},
			matches: true,
		},
		{
			testname: "IdenticalNetworkServiceEmptyConnectionPath",
			conn: &networkservice.Connection{
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				NetworkServices: []string{"ns1"},
			},
			matches: true,
		},
		{
			testname: "DifferentNetworkService",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				NetworkServices: []string{"ns2"},
			},
			matches: false,
		},
		{
			testname: "DifferentNetworkServiceEmptyConnectionPath",
			conn: &networkservice.Connection{
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				NetworkServices: []string{"ns2"},
			},
			matches: false,
		},
		{
			testname: "MatchingNetworkServiceList",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				NetworkServices: []string{"ns2", "ns1", "ns3"},
			},
			matches: true,
		},
		{
			testname: "NonMatchingNetworkServiceList",
			conn: &networkservice.Connection{
				Path: &networkservice.Path{
					PathSegments: []*networkservice.PathSegment{{Name: "s1", Id: "id1", Token: "t1"}, {Name: "s2", Id: "id2", Token: "t2"}},
				},
				NetworkService: "ns1",
			},
			selector: &networkservice.MonitorScopeSelector{
				NetworkServices: []string{"ns2", "ns3", "ns0"},
			},
			matches: false,
		},
	}

	for _, testCase := range cases {
		c := testCase
		t.Run(c.testname, func(t *testing.T) {
			require.Equal(t, c.conn.MatchesMonitorScopeSelector(c.selector), c.matches)
		})
	}
}
