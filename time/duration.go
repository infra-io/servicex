// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package time

import "time"

// Duration is an alias of time.Duration for implementing method UnmarshalText.
type Duration time.Duration

// FromDuration creates duration from time.Duration.
func FromDuration(d time.Duration) Duration {
	return Duration(d)
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

// Std returns d as time.Duration.
func (d *Duration) Std() time.Duration {
	return time.Duration(*d)
}
