// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func GetMetadata(md metadata.MD, key string) string {
	strs := md.Get(key)
	if len(strs) != 1 {
		return ""
	}

	return strs[0]
}

func GetContextMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	return GetMetadata(md, key)
}
