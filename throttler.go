package main

import "time"

type Throttler struct {
	lastCalled time.Time
	// How many times the function can be called per second
	RatePerSecond float64
}

func (t *Throttler) CanCall() bool {
	return time.Since(t.lastCalled) > time.Second/time.Duration(t.RatePerSecond)
}
func (t *Throttler) Call() {
	t.lastCalled = time.Now()
}
