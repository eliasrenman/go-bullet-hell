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

// Convert from radians to degrees
func RadToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

// Convert from degrees to radians
func DegToRad(deg float64) float64 {
	return deg / (180 / math.Pi)
}
