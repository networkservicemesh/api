// Copyright (c) 2020-2021 Nordix Foundation
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

package vxlan_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/vxlan"
)

type evenVNITest struct {
	parameters map[string]string
	expected   bool
}

var evenVNITests = []evenVNITest{
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "172.18.0.1",
			common.DstOriginalIP: "172.18.0.2",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			vxlan.SrcIP: "172.18.0.1",
			vxlan.DstIP: "172.18.0.2",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "172.28.0.1",
			common.DstOriginalIP: "172.28.0.2",
			vxlan.SrcIP:          "172.18.0.2",
			vxlan.DstIP:          "172.18.0.1",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "172.28.0.1",
			common.DstOriginalIP: "172.28.0.2",
			vxlan.SrcIP:          "172.18.0.1",
			vxlan.DstIP:          "172.18.0.2",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "fd00::1",
			common.DstOriginalIP: "fd00::2",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			vxlan.SrcIP: "fd00::1",
			vxlan.DstIP: "fd00::2",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "fd00::1",
			common.DstOriginalIP: "fd00::2",
			vxlan.SrcIP:          "fd00:1::1",
			vxlan.DstIP:          "fd00:1::2",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "fd00::1",
			common.DstOriginalIP: "fd00::2",
			vxlan.SrcIP:          "fd00:1::2",
			vxlan.DstIP:          "fd00:1::1",
		},
		expected: true,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "172.18.0.2",
			common.DstOriginalIP: "172.18.0.1",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			vxlan.SrcIP: "172.18.0.2",
			vxlan.DstIP: "172.18.0.1",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "172.28.0.2",
			common.DstOriginalIP: "172.28.0.1",
			vxlan.SrcIP:          "172.18.0.2",
			vxlan.DstIP:          "172.18.0.1",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "172.28.0.2",
			common.DstOriginalIP: "172.28.0.1",
			vxlan.SrcIP:          "172.18.0.1",
			vxlan.DstIP:          "172.18.0.2",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "fd00::2",
			common.DstOriginalIP: "fd00::1",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			vxlan.SrcIP: "fd00::2",
			vxlan.DstIP: "fd00::1",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "fd00::2",
			common.DstOriginalIP: "fd00::1",
			vxlan.SrcIP:          "fd00:1::2",
			vxlan.DstIP:          "fd00:1::1",
		},
		expected: false,
	},
	{
		parameters: map[string]string{
			common.SrcOriginalIP: "fd00::2",
			common.DstOriginalIP: "fd00::1",
			vxlan.SrcIP:          "fd00:1::1",
			vxlan.DstIP:          "fd00:1::2",
		},
		expected: false,
	},
}

func Test_EvenVNI(t *testing.T) {
	for _, evenVNI := range evenVNITests {
		networkServiceMechanism := &networkservice.Mechanism{
			Type:       vxlan.MECHANISM,
			Parameters: evenVNI.parameters,
		}
		mechanism := vxlan.ToMechanism(networkServiceMechanism)
		require.Equal(t, mechanism.EvenVNI(), evenVNI.expected)
	}
}
