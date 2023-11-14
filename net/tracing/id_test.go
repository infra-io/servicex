// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"testing"
)

func TestTraceID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(traceID())
	}
}
