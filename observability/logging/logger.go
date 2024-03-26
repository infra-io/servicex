// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logging

import (
	"context"

	"github.com/FishGoddess/logit"
)

func NewLogger(ctx context.Context, args any, resolvers ...ArgsResolver) *logit.Logger {
	var allResolved []any
	for _, resolve := range resolvers {
		resolved := resolve(ctx, args)
		allResolved = append(allResolved, resolved...)
	}

	return logit.FromContext(ctx).With(allResolved...)
}
