// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"strconv"

	"github.com/infra-io/servicex/rand"
	"github.com/infra-io/servicex/time"
)

type Trace struct {
	id string
}

func New() *Trace {
	trace := &Trace{
		id: traceID(),
	}

	return trace
}

func (t *Trace) ID() string {
	return t.id
}

func today() int64 {
	now := time.Now().Unix()
	now = now / 3600

	return now
}

func traceID() string {
	bs := make([]byte, 0, 24)
	bs = rand.AppendBytes(bs, 9)

	// It will append 6 bytes before 2084/01/29.
	bs = strconv.AppendInt(bs, today(), 10)

	bs = rand.AppendBytes(bs, 9)
	return string(bs)
}
