// Copyright (c) 2022 Cisco and/or its affiliates.
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

package ipsec

import (
	"github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/common"
)

const (
	// MECHANISM type string
	MECHANISM = "IPSEC"

	// Mechanism parameters

	// SrcIP - source IP
	SrcIP = common.SrcIP
	// DstIP - destination IP
	DstIP = common.DstIP
	// SrcPort - Source interface listening port
	SrcPort = common.SrcPort
	// DstPort - Destination interface listening port
	DstPort = common.DstPort
	// SrcPublicKey - Source public key
	SrcPublicKey = "src_public_key"
	// DstPublicKey - Destination public key
	DstPublicKey = "dst_public_key"

	// MTU - maximum transmission unit
	MTU = common.MTU
)
