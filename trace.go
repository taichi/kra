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
)

func FindCaller(skip int, bufferSize int, filter func(frame *runtime.Frame) (found bool)) *runtime.Frame {
	for { // buffering loop
		traced := 0
		callers := make([]uintptr, bufferSize)
		filled := runtime.Callers(skip, callers)
		frames := runtime.CallersFrames(callers)
		for ; ; traced++ {
			frame, more := frames.Next()
			if more {
				if filter(&frame) {
					return &frame
				}
			} else {
				break
			}
		}
		if filled < bufferSize {
			break // reached to the top of stack frame
		}
		skip += traced
	}
	return nil // not found
}

func ToPackageName(name string) string {
	dot := 0
	for i := len(name) - 1; 0 < i; i-- {
		if name[i] == '.' {
			dot = i
		} else if name[i] == '/' {
			break
		}
	}
	return name[:dot]
}
