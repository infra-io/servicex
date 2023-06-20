// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package time

import (
	"testing"
	"time"
)

// go test -v -cover -run=^TestFromDuration$
func TestFromDuration(t *testing.T) {
	d := FromDuration(time.Second)
	if int64(d) != int64(time.Second) {
		t.Errorf("d %d is wrong", d)
	}
}

// go test -v -cover -run=^TestDuration$
func TestDuration(t *testing.T) {
	d := Duration(0)
	if err := d.UnmarshalText([]byte("1s")); err != nil {
		t.Error(err)
	}

	if int64(d) != int64(time.Second) {
		t.Errorf("d %d is wrong", d)
	}

	if d.Std() != time.Second {
		t.Errorf("d.Std() %d is wrong", d.Std())
	}
}
