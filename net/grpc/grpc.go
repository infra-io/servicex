// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"encoding/json"
	"fmt"
	"path"

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
