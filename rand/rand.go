// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import (
	"math/rand"
	"strings"
	"time"

	"github.com/FishGoddess/logit"
	uuidlib "github.com/google/uuid"
)

const (
	uuidLength = 32
)

var (
	letters = [62]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

// String returns a random string including 0-9/a-z/A-Z not longer than length.
func String(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		index := random.Intn(len(letters))
		b[i] = letters[index]
	}

	return string(b)
}

// UUID returns an uuid of version 4 or a random string if failed.
func UUID() string {
	uuid, err := uuidlib.NewRandom()
	if err != nil {
		logit.Error(err, "new uuid failed").Log()
		return String(uuidLength)
	}

	return strings.ReplaceAll(uuid.String(), "-", "")
}
