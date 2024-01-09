// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/infra-io/servicex/net/tracing"
	"github.com/infra-io/servicex/runtime"
	"google.golang.org/grpc"
)

func serviceMethod(info *grpc.UnaryServerInfo) string {
	if method := path.Base(info.FullMethod); method != "" {
		return method
	}

	return info.FullMethod
}

func jsonify(v any) string {
	marshaled, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}

	return string(marshaled)
}

func Interceptor(timeout time.Duration, resolvers ...RequestResolver) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				logit.FromContext(ctx).Error("recovery from panic", "r", r, "stack", runtime.Stack())
			}
		}()

		beginTime := time.Now()
		method := serviceMethod(info)

		trace := tracing.New()
		ctx = tracing.NewContext(ctx, trace)

		logger := newLogger(ctx, method, trace, req, resolvers...)
		ctx = logit.NewContext(ctx, logger)

		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()

		reqJson := jsonify(req)
		logger.Debug("service method end", "request", reqJson)

		defer func() {
			cost := time.Since(beginTime)
			respJson := jsonify(resp)

			if err == nil {
				logger.Debug("service method end", "response", respJson, "cost", cost)
			} else {
				logger.Error("service method end", "response", respJson, "err", err, "cost", cost)
			}
		}()

		return handler(ctx, req)
	}
}
