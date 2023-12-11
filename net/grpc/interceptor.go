// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"path"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/servicex/net/tracing"
	"github.com/infra-io/servicex/runtime"
	"google.golang.org/grpc"
)

func shortMethod(info *grpc.UnaryServerInfo) string {
	if base := path.Base(info.FullMethod); base != "" {
		return base
	}

	return info.FullMethod
}

// TraceInterceptor sets a trace id to context.
func TraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		trace := tracing.New()
		ctx = tracing.NewContext(ctx, trace)

		return handler(ctx, req)
	}
}

// RecoveryInterceptor protects goroutine from panic.
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logit.ErrorContext(ctx, "recovery from panic", "r", r, "stack", runtime.Stack())
			}
		}()

		return handler(ctx, req)
	}
}

// CostInterceptor records the cost of method.
func CostInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		begin := time.Now()
		method := shortMethod(info)

		logit.InfoContext(ctx, method+" begin", "request", req)
		defer func() {
			cost := time.Since(begin)
			logit.InfoContext(ctx, method+" end", "response", resp, "cost", cost)
		}()

		return handler(ctx, req)
	}
}

// TimeoutInterceptor sets timeout to context.
func TimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		newCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return handler(newCtx, req)
	}
}
