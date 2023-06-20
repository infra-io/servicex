// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rand

import "testing"

// go test -v -cover -run=^TestString$
func TestString(t *testing.T) {
	str := String(32)
	if str == "" {
		t.Error("str is wrong")
	}

	t.Log(str)
}

// go test -v -cover -run=^TestUUID$
func TestUUID(t *testing.T) {
	str := UUID()
	if str == "" {
		t.Error("str is wrong")
	}

	t.Log(str)
}
