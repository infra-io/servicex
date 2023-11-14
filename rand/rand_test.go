// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import (
	"encoding/base64"
	"testing"
)

// go test -v -cover -run=^TestGenerateString$
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

// go test -v -cover -run=^TestGenerateToken$
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
