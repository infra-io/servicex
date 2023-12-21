// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logging

import (
	"context"
	"fmt"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithTraceID$
func TestWithTraceID(t *testing.T) {
	conf := &Config{WithTraceID: false}
	WithTraceID().ApplyTo(conf)

	if !conf.WithTraceID {
		t.Fatal("conf.WithTraceID is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithServiceMethod$
func TestWithServiceMethod(t *testing.T) {
	conf := &Config{WithServiceMethod: false}
	WithServiceMethod().ApplyTo(conf)

	if !conf.WithServiceMethod {
		t.Fatal("conf.WithServiceMethod is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithRequestResolver$
func TestWithRequestResolver(t *testing.T) {
	args := []any{"key", "value"}

	resolver := func(ctx context.Context, request any) []any {
		return args
	}

	conf := &Config{Resolvers: nil}
	WithRequestResolver(resolver).ApplyTo(conf)

	if len(conf.Resolvers) != 1 {
		t.Fatal("len(conf.Resolvers) is wrong")
	}

	got := conf.Resolvers[0](context.Background(), nil)
	if fmt.Sprintf("%p", got) != fmt.Sprintf("%p", args) {
		t.Fatal("got wrong args")
	}
}
