// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package trace

import (
	"context"
	"testing"
)

// go test -v -cover -run=^TestNew$
func TestNew(t *testing.T) {
	for i := 0; i < 10; i++ {
		traceID := New()
		if traceID == "" {
			t.Error("traceID is wrong")
		}

		t.Log(traceID)
	}
}

// go test -v -cover -run=^TestNewContext$
func TestNewContext(t *testing.T) {
	traceID := New()
	ctx := NewContext(context.Background(), traceID)

	value := ctx.Value(contextKey)
	if value == nil {
		t.Error("ctx.Value returns nil")
	}

	if value.(string) != traceID {
		t.Errorf("value %+v != traceID %s", value, traceID)
	}

	t.Log("traceID:", traceID)
}

// go test -v -cover -run=^TestNewContext$
func TestNewContextWithID(t *testing.T) {
	ctx := context.Background()

	ctx, traceID := NewContextWithID(ctx)
	if traceID == "" {
		t.Error("traceID == ''")
	}

	value := ctx.Value(contextKey)
	if value == nil {
		t.Error("ctx.Value returns nil")
	}

	if value.(string) != traceID {
		t.Errorf("value %+v != traceID %s", value, traceID)
	}

	t.Log("traceID:", traceID)
}

// go test -v -cover -run=^TestFromContext$
func TestFromContext(t *testing.T) {
	ctx := context.Background()

	traceIDInCtx := FromContext(ctx)
	if traceIDInCtx != "" {
		t.Errorf("traceIDInCtx %s != ''", traceIDInCtx)
	}

	traceID := New()
	ctx = context.WithValue(ctx, contextKey, traceID)

	traceIDInCtx = FromContext(ctx)
	if traceIDInCtx != traceID {
		t.Errorf("traceIDInCtx %s != traceID %s", traceIDInCtx, traceID)
	}
}
