// Copyright 2013 Artur Grabowski. All rights reserved.
// Use of this source code is governed by a ISC-style
// license that can be found in the LICENSE file.
package stopwatch

import (
	"testing"
	"time"
)

func w1() {
	d, _ := time.ParseDuration("1s")
	select {
	case <- time.After(d):
	}
}

func af(t *testing.T, h, e float64, name string) {
	if h < e || h > e + 0.1 {
		t.Errorf("bad %v: %v !~= %v", name, h, e)
	}
}

func TestSecond(t *testing.T) {
	s := New()
	s.Start()
	w1()
	s.Stop()
	af(t, s.Seconds(), 1.0, "duration")
}

func TestSecondTwice(t *testing.T) {
	s := New()
	s.Start()
	w1()
	s.Stop()
	s.Start()
	w1()
	d2 := s.Stop()
	af(t, s.Seconds(), 2.0, "duration")
	af(t, d2.Seconds(), 1.0, "stop duration")
}

func TestHandOver(t *testing.T) {
	s1, s2 := New(), New()
	s1.Start()
	w1()
	d1 := s1.Handover(&s2)
	w1()
	d2 := s2.Stop()
	af(t, s1.Seconds(), 1.0, "s1 duration")
	af(t, s2.Seconds(), 1.0, "s2 duration")
	af(t, d1.Seconds(), 1.0, "d1 duration")
	af(t, d2.Seconds(), 1.0, "d2 duration")
}
