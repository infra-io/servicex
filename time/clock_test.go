// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package time

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

// go test -bench=^BenchmarkTimeNow$ -benchtime=1s ./clock.go ./clock_test.go
func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

// go test -bench=^BenchmarkClockNow$ -benchtime=1s ./clock.go ./clock_test.go
func BenchmarkClockNow(b *testing.B) {
	clock := NewClock()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		clock.Now()
	}
}

// go test -v -cover -run=^TestNewClock$
func TestNewClock(t *testing.T) {
	var clocks []*Clock

	for i := 0; i < 100; i++ {
		clocks = append(clocks, NewClock())
	}

	for i := 0; i < 100; i++ {
		if clocks[i] != clock {
			t.Errorf("clocks[i] %p != clock %p", clocks[i], clock)
		}
	}
}

// go test -v -cover -run=^TestClock$
func TestClock(t *testing.T) {
	testClock := NewClock()

	for i := 0; i < 100; i++ {
		got := testClock.Now()
		gap := time.Since(got)
		t.Logf("got: %v, gap: %v", got, gap)

		if math.Abs(float64(gap.Nanoseconds())) > float64(duration)*1.5 {
			t.Errorf("now %v is wrong", got)
		}

		time.Sleep(time.Duration(rand.Int63n(int64(duration))))
	}
}
