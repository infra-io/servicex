// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import "testing"

// go test -v -cover -run=^TestGenerateString$
func TestGenerateString(t *testing.T) {
	str := GenerateString(32)
	if str == "" {
		t.Error("str is wrong")
	}

	t.Log(str)
}

// go test -v -cover -run=^TestGenerateToken$
func TestGenerateToken(t *testing.T) {
	str := GenerateToken(48)
	if str == "" {
		t.Error("str is wrong")
	}

	t.Log(str)
}
