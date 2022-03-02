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
	"reflect"
	"testing"
)

func TestParsePortRange(t *testing.T) {
	type args struct {
		portRange string
	}
	tests := []struct {
		name    string
		args    args
		want    *PortRange
		wantErr bool
	}{
		{
			name: "empty port range",
			args: args{
				portRange: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "single port",
			args: args{
				portRange: "80",
			},
			want: &PortRange{
				Start: 80,
				End:   80,
			},
			wantErr: false,
		},
		{
			name: "valid port range",
			args: args{
				portRange: "80-90",
			},
			want: &PortRange{
				Start: 80,
				End:   90,
			},
			wantErr: false,
		},
		{
			name: "invalid port",
			args: args{
				portRange: "abc",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid port range (first port)",
			args: args{
				portRange: "abc-80",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid port range (second port)",
			args: args{
				portRange: "80-abc",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePortRange(tt.args.portRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePortRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePortRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
