// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sequence

import (
	"strconv"
	"sync/atomic"
	"testing"
)

// go test -v -cover -run=^TestNext$
func TestNext(t *testing.T) {
	for i := atomic.LoadUint64(&seq) + 1; i <= 100; i++ {
		if next := Next(); next != i {
			t.Errorf("next %d != i %d", next, i)
		}
	}
}

// go test -v -cover -run=^TestNextString$
func TestNextString(t *testing.T) {
	for i := atomic.LoadUint64(&seq) + 1; i <= 100; i++ {
		iStr := strconv.FormatUint(i, 10)
		if next := NextString(); next != iStr {
			t.Errorf("next %s != i %s", next, iStr)
		}
	}
}
