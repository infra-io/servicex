// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/servicex/net/tracing"
)

type RequestResolver func(ctx context.Context, request any) []any

func newLogger(ctx context.Context, method string, trace *tracing.Trace, req any, resolvers ...RequestResolver) *logit.Logger {
	args := []any{"service.trace_id", trace.ID(), "service.method", method}

	for _, resolve := range resolvers {
		resolved := resolve(ctx, req)
		args = append(args, resolved...)
	}

	return logit.FromContext(ctx).With(args...)
}
