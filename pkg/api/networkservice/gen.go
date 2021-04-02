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

// Package networkservice provides the Network Service Mesh API: NetworkService{Server,Client}.{Request,Close}
package networkservice

// Run with protoc and proto-gen-go matching the versions found in .github/workflows/ci.yaml
// Please also note that you need a 'batteries included' version of protoc such as the one installed
// with brew rather than the 'single binary' install to insure you get the correct *.proto files for imports
//go:generate bash -c "protoc -I .  *.proto --go_out=plugins=grpc,paths=source_relative:."
