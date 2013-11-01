// Copyright 2013 Artur Grabowski. All rights reserved.
// Use of this source code is governed by a ISC-style
// license that can be found in the LICENSE file.
package stopwatch

import (
	"time"
)

// A Stopwatch is used for keeping track of elapsed wall-clock time.
// Stopwatches don't keep track if they are running or not,
// it's left to the user to deal with this. Unmatched Start/Stop
// calls will lead to the stopwatch having an undefined value.
//
// Stopwatches are currently limited by the data type of time.Duration
// and will not measure longer durations than 290 years.
type Stopwatch struct {
	acc time.Duration
	t time.Time
}

// Returns a new Stopwatch.
func New() Stopwatch {
	return Stopwatch{}
}

// Starts counting elapsed wall-clock time 
func (s *Stopwatch) Start() {
	s.t = time.Now()
}

// Stops counting elapsed wall-clock time
func (s *Stopwatch) Stop() time.Duration {
	d := time.Since(s.t)
	s.acc += d
	return d
}

// Equivalent to s.Stop(); s2.Start() with no time elapsed between,
// just using one time stamp.
func (s *Stopwatch) Handover(s2 *Stopwatch) time.Duration {
	s2.t = time.Now()
	d := s.t.Sub(s2.t)
	s.acc += d
	return d
}

// Takes a snapshot of a currently running stopwatch. If the stopwatch
// wasn't running the result is undefined.
func (s *Stopwatch) Snapshot() Stopwatch {
	ret := *s
	ret.Stop()
	return ret
}

// Reset the accumulated time in the stopwatch and 
func (s *Stopwatch) Reset() {
	*s = Stopwatch{}
}

// Returns the accumulated sanoseconds of the stopwatch. If the stopwatch
// is currently running that will not be accounted for.
func (s Stopwatch) Nanoseconds() int64 {
	return s.acc.Nanoseconds()
}

// Returns the accumulated seconds of the stopwatch. If the stopwatch is
// currently running that will not be accounted for.
func (s Stopwatch) Seconds() float64 {
	return s.acc.Seconds()
}

func (s Stopwatch) String() string {
	return s.acc.String()
}
