// Copyright 2013 Artur Grabowski. All rights reserved.
// Use of this source code is governed by a ISC-style
// license that can be found in the LICENSE file.
package stopwatch

import (
	"testing"
	"time"
)

func TestSecond(t *testing.T) {
	d, _ := time.ParseDuration("1s")
	s := New()
	s.Start()
	select {
	case <- time.After(d):
	}
	s.Stop()
	if s.Seconds() < 1.0 || s.Seconds() > 1.1 {
		t.Fatalf("wrong duration: %v", s)
	}
}