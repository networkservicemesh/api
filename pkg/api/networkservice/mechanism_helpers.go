// Copyright (c) 2018-2020 Cisco Systems, Inc.
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
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Equals returns if mechanism equals given mechanism
func (x *Mechanism) Equals(mechanism protoreflect.ProtoMessage) bool {
	// use as an proto.Message
	return proto.Equal(x, mechanism)
}

// Clone clones mechanism
func (x *Mechanism) Clone() *Mechanism {
	return proto.Clone(x).(*Mechanism)
}

var mechanismValidators map[string]func(*Mechanism) error
var mechanismValidatorsMutex sync.Mutex

// AddMechanism adds a Mechanism
func AddMechanism(mtype string, validator func(*Mechanism) error) {
	mechanismValidatorsMutex.Lock()
	defer mechanismValidatorsMutex.Unlock()
	mechanismValidators[mtype] = validator
}

// IsValid - is the Mechanism Valid?
func (x *Mechanism) IsValid() error {
	if x == nil {
		return errors.New("mechanism cannot be nil")
	}
	validator, ok := mechanismValidators[x.GetType()]
	if ok {
		return validator(x)
	}
	// NOTE: this means that we intentionally decide that Mechanisms are valid
	// unless we have a Validator that says otherwise
	return nil
}
