// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"context"

	"github.com/FishGoddess/servicex/rand"
	"github.com/FishGoddess/servicex/sequence"
)

var (
	idContextKey struct{}
)

// ID returns a new trace id.
func ID() string {
	return rand.UUID() + "_" + sequence.NextString()
}

// WithID creates a new trace id and sets to ctx.
func WithID(ctx context.Context) (context.Context, string) {
	traceID := ID()
	ctx = NewContext(ctx, traceID)
	return ctx, traceID
}

// NewContext creates a new context with trace id.
func NewContext(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, idContextKey, traceID)
	return ctx
}

// FromContext returns the trace id from context.
func FromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(idContextKey).(string)
	if !ok {
		return ""
	}

	return traceID
}
