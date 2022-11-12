package util

import "math"

func NormalizeVector(x float64, y float64) (float64, float64) {
	if x == 0 && y == 0 {
		return x, y // Avoid divide-by-zero overflow
	}

	magnitude := math.Sqrt(x*x + y*y)
	return x / magnitude, y / magnitude
}
