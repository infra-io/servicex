// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"os"
	"strconv"
	"sync/atomic"

	"github.com/infra-io/servicex/rand"
	"github.com/infra-io/servicex/time"
)

const (
	timeFormat = "050415020106"
)

var (
	pid = strconv.Itoa(os.Getpid())
	seq = uint64(0)
)

func nextSequence() uint64 {
	return atomic.AddUint64(&seq, 1)
}

func traceID() string {
	now := time.Now().Format(timeFormat)
	str := rand.GenerateString(16)
	seq := strconv.FormatUint(nextSequence(), 10)
	traceID := str[:4] + now[:6] + str[4:8] + now[6:] + str[:12] + pid + str[12:] + seq
	return traceID
}
