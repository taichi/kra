// Copyright 2022 taichi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kra

import (
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPackageName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"runtime", args{"runtime.Callers"}, "runtime"},
		{"pkg", args{"github.com/taichi/kra.FindCaller"}, "github.com/taichi/kra"},
		{"main", args{"main.main.func1"}, "main"},
		{"receiver", args{"github.com/taichi/kra.(*CoreParse).Proceed"}, "github.com/taichi/kra"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPackageName(tt.args.name); got != tt.want {
				t.Errorf("ToPackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCaller(t *testing.T) {
	frame := FindCaller(0, 2, func(frame *runtime.Frame) (found bool) {
		return strings.HasSuffix(frame.Function, "TestFindCaller")
	})
	assert.Equal(t, "github.com/taichi/kra.TestFindCaller", frame.Function)
}
