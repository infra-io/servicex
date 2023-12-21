// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tracing

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/infra-io/servicex/rand"
	"github.com/infra-io/servicex/time"
)

func numHex(num uint32) string {
	var bs [4]byte
	binary.BigEndian.PutUint32(bs[:], num)

	return hex.EncodeToString(bs[:])
}

func now() string {
	now := time.Now().Unix()
	now = now / 86400

	return numHex(uint32(now))
}

func traceID() string {
	now := now()
	str := rand.GenerateString(20)
	traceID := str[:8] + now[4:6] + str[8:16] + now[6:] + str[16:]

	return traceID
}
