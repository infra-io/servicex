// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package runtime

import "runtime"

var (
	// StackSize is the max size of stack.
	StackSize = 16 << 10 // 16KB
)

// Stack returns the stack information.
func Stack() string {
	stack := make([]byte, StackSize)
	n := runtime.Stack(stack, false)
	return string(stack[:n])
}
