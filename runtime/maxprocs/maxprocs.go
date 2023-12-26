// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maxprocs

import (
	"github.com/FishGoddess/logit"
	"go.uber.org/automaxprocs/maxprocs"
)

func Setup() {
	undo, err := maxprocs.Set(maxprocs.Logger(logit.Printf))
	if err != nil {
		logit.Error("set maxprocs failed", "err", err)
		undo()
	}
}
