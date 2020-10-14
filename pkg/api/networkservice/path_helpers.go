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

package networkservice

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// Clone clones
func (x *Path) Clone() *Path {
	return proto.Clone(x).(*Path)
}

// IsValid returns true if Path p is Valid
func (x *Path) IsValid() error {
	if x == nil {
		return nil
	}
	if int(x.GetIndex()) >= len(x.GetPathSegments()) {
		return errors.New("Path.Index >= len(Path.PathSegments)")
	}
	return nil
}
