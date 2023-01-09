package main

import (
	"math"
	"time"
)

// ClampFloat returns a copy of the value clamped between min and max
func ClampFloat(min float64, value float64, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

// ClampInt returns a copy of the value clamped between min and max
func ClampInt(min int, value int, max int) int {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

// RadToDeg converts from radians to degrees
func RadToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

// DegToRad converts from degrees to radians
func DegToRad(deg float64) float64 {
	return deg / (180 / math.Pi)
}

var startTime = time.Now()

// TimeSinceStart returns the current
func TimeSinceStart() time.Duration {
	return time.Since(startTime)
}

// Sign returns -1 if the value is negative, 1 otherwise
func Sign(value float64) float64 {
	if value < 0 {
		return -1
	}
	return 1
}
