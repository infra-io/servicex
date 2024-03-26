// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"context"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewContext$
func TestNewContext(t *testing.T) {
	trace := New()
	ctx := NewContext(context.Background(), trace)

	value := ctx.Value(contextKey)
	if value == nil {
		t.Error("ctx.Value returns nil")
	}

	if value.(*Trace) != trace {
		t.Errorf("value %+v != trace %s", value, trace)
	}

	t.Log("traceID:", trace.ID())
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestFromContext$
func TestFromContext(t *testing.T) {
	ctx := context.Background()

	traceInCtx := FromContext(ctx)
	if traceInCtx.id != "" {
		t.Errorf("traceInCtx.id %s != ''", traceInCtx.id)
	}

	trace := New()
	ctx = context.WithValue(ctx, contextKey, trace)

	traceInCtx = FromContext(ctx)
	if traceInCtx != trace {
		t.Errorf("traceInCtx %+v != trace %+v", traceInCtx, trace)
	}

	t.Log("traceID:", trace.ID())
}
