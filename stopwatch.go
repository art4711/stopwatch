// Copyright 2013 Artur Grabowski. All rights reserved.
// Use of this source code is governed by a ISC-style
// license that can be found in the LICENSE file.
package stopwatch

import (
	"time"
)

type Stopwatch struct {
	acc time.Duration
	t time.Time
}

func New() Stopwatch {
	return Stopwatch{}
}

func (s *Stopwatch) Start() {
	s.t = time.Now()
}

func (s *Stopwatch) Stop() {
	s.acc = time.Since(s.t)
}

func (s *Stopwatch) Handover(s2 *Stopwatch) {
	s2.t = time.Now()
	s.acc = s.t.Sub(s2.t)
}

func (s *Stopwatch) Snapshot() Stopwatch {
	ret := *s
	ret.Stop()
	return ret
}

func (s *Stopwatch) Reset() {
	*s = Stopwatch{}
}

func (s Stopwatch) Nanoseconds() int64 {
	return s.acc.Nanoseconds()
}

func (s Stopwatch) Seconds() float64 {
	return s.acc.Seconds()
}

func (s Stopwatch) String() string {
	return s.acc.String()
}
