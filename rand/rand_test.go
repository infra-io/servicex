// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import (
	"encoding/base64"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestAppendBytes$
func TestAppendBytes(t *testing.T) {
	n := 32

	for i := 0; i < 10; i++ {
		bs := make([]byte, 0, n)
		bs = AppendBytes(bs, n)

		if len(bs) != n {
			t.Errorf("bs length %d is wrong", len(bs))
		}

		t.Log(string(bs))
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestGenerateString$
func TestGenerateString(t *testing.T) {
	n := 32

	for i := 0; i < 10; i++ {
		str := GenerateString(n)
		if str == "" {
			t.Error("str is wrong")
		}

		if len(str) != n {
			t.Errorf("str length %d is wrong", len(str))
		}

		t.Log(str)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestGenerateToken$
func TestGenerateToken(t *testing.T) {
	n := 48

	for i := 0; i < 10; i++ {
		token := GenerateToken(n)
		if token == "" {
			t.Error("token is wrong")
		}

		raw, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			t.Error(err)
		}

		if len(raw) != n {
			t.Errorf("token length %d is wrong", len(raw))
		}

		t.Log(token)
	}
}
