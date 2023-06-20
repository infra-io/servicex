// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sequence

import (
	"strconv"
	"sync/atomic"
)

var (
	seq uint64 = 0
)

// Next returns next sequence.
func Next() uint64 {
	return atomic.AddUint64(&seq, 1)
}

// NextString returns next sequence in string.
func NextString() string {
	return strconv.FormatUint(Next(), 10)
}
