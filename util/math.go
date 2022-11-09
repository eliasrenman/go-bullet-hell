package util

import "math"

func ClampFloat(min float64, value float64, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

func ClampInt(min int, value int, max int) int {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}
