package util

import "math"

func Clamp(min float64, value float64, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}
