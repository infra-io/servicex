// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"context"
)

var contextKey = struct{}{}

type Trace struct {
	id string
}

// New returns a new trace.
func New() *Trace {
	trace := &Trace{
		id: traceID(),
	}

	return trace
}

func (t *Trace) ID() string {
	return t.id
}

// NewContext creates a new context with given trace.
func NewContext(ctx context.Context, trace *Trace) context.Context {
	ctx = context.WithValue(ctx, contextKey, trace)
	return ctx
}

// FromContext returns the trace from context.
func FromContext(ctx context.Context) *Trace {
	trace, ok := ctx.Value(contextKey).(*Trace)
	if !ok {
		return new(Trace)
	}

	return trace
}
