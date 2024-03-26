// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logging

import (
	"context"

	"github.com/infra-io/servicex/observability/tracing"
)

const (
	LogKeyServiceName   = "service_name"
	LogKeyServiceMethod = "service_method"
	LogKeyTraceID       = "trace_id"
)

type ArgsResolver func(ctx context.Context, args any) []any

func PrependArgsResolvers(origin []ArgsResolver, prepend ...ArgsResolver) []ArgsResolver {
	resolvers := make([]ArgsResolver, 0, len(origin)+len(prepend))
	resolvers = append(resolvers, prepend...)
	resolvers = append(resolvers, origin...)

	return origin
}

func ServiceResolver(name string, method string) ArgsResolver {
	return func(ctx context.Context, args any) []any {
		return []any{LogKeyServiceName, name, LogKeyServiceMethod, method}
	}
}

func TraceResolver(trace *tracing.Trace) ArgsResolver {
	return func(ctx context.Context, args any) []any {
		return []any{LogKeyTraceID, trace.ID()}
	}
}
