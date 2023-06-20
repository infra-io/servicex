// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package interceptors

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/servicex/runtime"
	"github.com/FishGoddess/servicex/trace"
	"google.golang.org/grpc"
)

func shortMethod(info *grpc.UnaryServerInfo) string {
	if base := path.Base(info.FullMethod); base != "" {
		return base
	}

	return info.FullMethod
}

// Trace sets a trace id to context.
func Trace() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		ctx, _ = trace.WithID(ctx)
		return handler(ctx, req)
	}
}

// Recovery protects goroutine from panic.
func Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %+v", r)
				logit.Error(err, "recovery from panic").Any("r", r).String("stack", runtime.Stack()).LogX(ctx)
			}
		}()

		return handler(ctx, req)
	}
}

// Cost records the cost of method.
func Cost() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		begin := time.Now()
		method := shortMethod(info)

		logit.Info("%s begin", method).Json("request", req).LogX(ctx)
		defer func() {
			cost := time.Since(begin)
			logit.Info("%s end", method).Json("response", req).Duration("cost", cost).LogX(ctx)
		}()

		return handler(ctx, req)
	}
}

// Timeout sets timeout to context.
func Timeout(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		newCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return handler(newCtx, req)
	}
}
