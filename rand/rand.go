// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import (
	"encoding/base64"
	"math/rand"
	"time"
)

var random *rand.Rand

var words = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func init() {
	now := time.Now().Unix()
	source := rand.NewSource(now)
	random = rand.New(source)
}

func AppendBytes(bs []byte, n int) []byte {
	length := len(words)

	for i := 0; i < n; i++ {
		index := random.Intn(length)
		bs = append(bs, words[index])
	}

	return bs
}

// GenerateString generates a string including n words in random.
func GenerateString(n int) string {
	bs := make([]byte, 0, n)
	bs = AppendBytes(bs, n)

	return string(bs)
}

// GenerateToken generates a string in base64 for token usages.
func GenerateToken(n int) string {
	raw := GenerateString(n)
	token := base64.StdEncoding.EncodeToString([]byte(raw))

	return token
}
