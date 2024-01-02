// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"encoding/json"
	"path"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/servicex/net/logging"
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

// LogInterceptor sets logging attributes to logger.
func LogInterceptor(opts ...logging.Option) grpc.UnaryServerInterceptor {
	conf := logging.NewDefaultConfig()

	for _, opt := range opts {
		opt.ApplyTo(conf)
	}

	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		args := make([]any, 0, 4)

		if conf.WithTraceID {
			trace := tracing.FromContext(ctx)
			args = append(args, "service.trace_id", trace.ID())
		}

		if conf.WithServiceMethod {
			method := shortMethod(info)
			args = append(args, "service.method", method)
		}

		for _, resolve := range conf.Resolvers {
			resolved := resolve(ctx, req)
			args = append(args, resolved...)
		}

		logger := logit.FromContext(ctx).With(args...)
		ctx = logit.NewContext(ctx, logger)

		return handler(ctx, req)
	}
}

// RecoveryInterceptor protects goroutine from panic.
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logit.FromContext(ctx).Error("recovery from panic", "r", r, "stack", runtime.Stack())
			}
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

// CostInterceptor records the cost of method.
func CostInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		begin := time.Now()
		logger := logit.FromContext(ctx)

		marshaled, jsonErr := json.Marshal(req)
		if err != nil {
			logger.Debug("service method begin", "request", req, "json_err", jsonErr)
		} else {
			logger.Debug("service method begin", "request", string(marshaled))
		}

		defer func() {
			cost := time.Since(begin)

			marshaled, jsonErr = json.Marshal(resp)
			if err != nil {
				logger.Debug("service method end", "response", resp, "err", err, "cost", cost, "jsonErr", jsonErr)
			} else {
				logger.Debug("service method end", "response", string(marshaled), "err", err, "cost", cost)
			}
		}()

		return handler(ctx, req)
	}
}

func BaseInterceptors(timeout time.Duration, opts ...logging.Option) []grpc.UnaryServerInterceptor {
	opts = append(opts, logging.WithTraceID(), logging.WithServiceMethod())

	interceptors := []grpc.UnaryServerInterceptor{
		TraceInterceptor(), LogInterceptor(opts...), RecoveryInterceptor(), TimeoutInterceptor(timeout), CostInterceptor(),
	}

	return interceptors
}
