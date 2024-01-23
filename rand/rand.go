// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import (
	"math/rand"
	"time"
	"unsafe"
)

const words = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var source = rand.NewSource(time.Now().Unix())

// AppendBytes appends n bytes to bs in random.
func AppendBytes(bs []byte, n int) []byte {
	length := int64(len(words))

	for i := 0; i < n; i++ {
		index := source.Int63() % length
		bs = append(bs, words[index])
	}

	return bs
}

// GenerateBytes generates a byte slice including n words in random.
func GenerateBytes(n int) []byte {
	bs := make([]byte, 0, n)
	bs = AppendBytes(bs, n)

	return bs
}

// GenerateString generates a string including n words in random.
func GenerateString(n int) string {
	bs := GenerateBytes(n)

	// Learn from strings.Builder.String()
	return *(*string)(unsafe.Pointer(&bs))
}
