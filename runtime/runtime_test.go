// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import "testing"

// go test -v -cover -run=^TestStack$
func TestStack(t *testing.T) {
	stack := Stack()
	t.Log(stack)
}
