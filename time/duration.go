// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package time

import "time"

// Duration is an alias of time.Duration for implementing method UnmarshalText.
type Duration time.Duration

// NewDuration returns a duration created from time.Duration.
func NewDuration(d time.Duration) Duration {
	return Duration(d)
}

// Standard returns d as time.Duration in standard library.
func (d *Duration) Standard() time.Duration {
	return time.Duration(*d)
}

// UnmarshalText unmarshal text to Duration.
func (d *Duration) UnmarshalText(text []byte) error {
	parsed, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}

	*d = Duration(parsed)
	return nil
}
