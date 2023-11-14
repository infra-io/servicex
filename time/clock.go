// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package time

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	duration = 100 * time.Millisecond
)

var (
	clock     *Clock
	clockOnce sync.Once
)

// Clock is a fast clock for getting current time.
// It caches time and updates it in fixed duration which may return an "incorrect" time compared with time.Now().UnixNano().
// In fact, we don't recommend you to use it unless you have to...
// According to our benchmarks, it does run faster than time.Now:
//
// In my win10 pc:
// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 5800X 8-Core Processor
// BenchmarkTimeNow-16             338466458                3.523 ns/op           0 B/op          0 allocs/op
// BenchmarkClockNow-16            1000000000               0.2165 ns/op          0 B/op          0 allocs/op
//
// In my macbook with charging:
// goos: darwin
// goarch: amd64
// pkg: github.com/FishGoddess/cachego/pkg/clock
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkTimeNow-12             17841402                67.05 ns/op            0 B/op          0 allocs/op
// BenchmarkClockNow-12            1000000000               0.2528 ns/op          0 B/op          0 allocs/op
//
// In my cloud server using 2 cores and running benchmarks in docker container:
// goos: linux
// goarch: amd64
// pkg: github.com/FishGoddess/cachego/pkg/clock
// cpu: AMD EPYC 7K62 48-Core Processor
// BenchmarkTimeNow-2      17946441                65.62 ns/op            0 B/op          0 allocs/op
// BenchmarkClockNow-2     1000000000               0.3162 ns/op          0 B/op          0 allocs/op
//
// PS: All benchmarks are ran with "go test -bench=. -benchtime=1s".
//
// However, the performance of time.Now is faster enough in many os and is enough for 99.9% situations.
// The another reason choosing to use it is .
// So, better performance should not be the first reason to use it.
// The first reason to use it is reducing gc objects, but we hope you never use it :)
type Clock struct {
	now int64
}

// NewClock creates a new clock which caches time and updates it in fixed duration.
func NewClock() *Clock {
	clockOnce.Do(func() {
		clock = &Clock{
			now: time.Now().UnixNano(),
		}

		go clock.start()
	})

	return clock
}

func (c *Clock) start() {
	for {
		for i := 0; i < 9; i++ {
			time.Sleep(duration)
			atomic.AddInt64(&c.now, int64(duration))
		}

		time.Sleep(duration)
		atomic.StoreInt64(&c.now, time.Now().UnixNano())
	}
}

// Now returns the current time.
func (c *Clock) Now() time.Time {
	nanos := atomic.LoadInt64(&c.now)
	return time.Unix(0, nanos)
}

// Now returns the current time.
func Now() time.Time {
	return NewClock().Now()
}
