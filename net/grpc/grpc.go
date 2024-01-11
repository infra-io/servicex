// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"fmt"
	"path"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func serviceMethod(info *grpc.UnaryServerInfo) string {
	if method := path.Base(info.FullMethod); method != "" {
		return method
	}

	return info.FullMethod
}

func jsonify(v any) string {
	jpb := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:     true,
			EmitUnpopulated:   true,
			EmitDefaultValues: true,
		},
	}

	marshaled, err := jpb.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}

	return string(marshaled)
}
