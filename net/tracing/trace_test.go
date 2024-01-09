// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestTraceID$
func TestTraceID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(traceID())
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNew$
func TestNew(t *testing.T) {
	for i := 0; i < 10; i++ {
		trace := New()
		if trace == nil {
			t.Error("trace is nil")
		}

		t.Log(trace.ID())
	}
}

// go test -bench=^BenchmarkNew$
// BenchmarkNew-2   	 3660554	       331.2 ns/op	      24 B/op	       1 allocs/op
func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		New()
	}
}
