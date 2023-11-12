// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"context"
	"encoding/base64"
	"os"
	"strconv"
	"sync/atomic"

	"github.com/infra-io/servicex/rand"
	"github.com/infra-io/servicex/time"
)

const (
	timeFormat = "20060102"
)

var (
	clock = time.NewClock()
	pid   = strconv.Itoa(os.Getpid())
	seq   = uint64(0)

	contextKey struct{}
)

func nextSequence() uint64 {
	return atomic.AddUint64(&seq, 1)
}

// New returns a new trace id.
func New() string {
	prefix := rand.GenerateString(8)
	date := clock.Now().Format(timeFormat)
	suffix := strconv.FormatUint(nextSequence(), 10)
	traceID := prefix + "_" + date + "_" + pid + "_" + suffix
	return base64.StdEncoding.EncodeToString([]byte(traceID))
}

// NewContext creates a new context with given traceID.
func NewContext(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, contextKey, traceID)
	return ctx
}

// NewContextWithID creates a new context with a new trace id.
func NewContextWithID(ctx context.Context) (context.Context, string) {
	traceID := New()
	ctx = NewContext(ctx, traceID)
	return ctx, traceID
}

// FromContext returns the trace id from context.
func FromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(contextKey).(string)
	if !ok {
		return ""
	}

	return traceID
}
