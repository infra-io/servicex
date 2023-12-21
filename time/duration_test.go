// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package time

import (
	"testing"
	"time"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewDuration$
func TestNewDuration(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want Duration
	}{
		{name: "0", d: time.Duration(0), want: Duration(0)},
		{name: "second", d: time.Second, want: Duration(time.Second)},
		{name: "minute", d: time.Minute, want: Duration(time.Minute)},
		{name: "hour", d: time.Hour, want: Duration(time.Hour)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDuration(tt.d); got != tt.want {
				t.Errorf("Duration.Standard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDurationStandard(t *testing.T) {
	tests := []struct {
		name string
		d    Duration
		want time.Duration
	}{
		{name: "0", d: Duration(0), want: time.Duration(0)},
		{name: "second", d: Duration(time.Second), want: time.Second},
		{name: "minute", d: Duration(time.Minute), want: time.Minute},
		{name: "hour", d: Duration(time.Hour), want: time.Hour},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Standard(); got != tt.want {
				t.Errorf("Duration.Standard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDurationUnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}

	tests := []struct {
		name    string
		d       Duration
		text    string
		want    Duration
		wantErr error
	}{
		{name: "0", text: "0", want: Duration(0), wantErr: nil},
		{name: "3s", text: "3s", want: Duration(3 * time.Second), wantErr: nil},
		{name: "2m40s", text: "2m40s", want: Duration(2*time.Minute + 40*time.Second), wantErr: nil},
		{name: "1h30m55s", text: "1h30m55s", want: Duration(time.Hour + 30*time.Minute + 55*time.Second), wantErr: nil},
		{name: "24h", text: "24h", want: Duration(24 * time.Hour), wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalText([]byte(tt.text)); err != tt.wantErr {
				t.Errorf("Duration.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got := tt.d; got != tt.want {
				t.Errorf("Duration.UnmarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}
