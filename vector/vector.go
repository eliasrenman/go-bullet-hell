package vector

import "math"

type Vector struct {
	X float64
	Y float64
}

type Point = Vector

// Calculates the magnitude (length) of the vector
func (vec *Vector) Magnitude() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

// Calculates the direction of the vector
func (vec *Vector) Direction() float64 {
	return math.Tan(vec.Y / vec.X)
}

// Normalizes the vector, setting its magnitude to 1
func (vec *Vector) Normalize() {
	magnitude := vec.Magnitude()

	if magnitude == 0 {
		return // Avoid divide-by-zero overflow
	}

	vec.X /= magnitude
	vec.Y /= magnitude
}

// Creates a normalized copy of the vector
func (vec Vector) Normalized() Vector {
	v := vec.Copy()
	v.Normalize()

	return v
}

// Creates a copy of the vector
func (vec Vector) Copy() Vector {
	v := Vector{vec.X, vec.Y}
	return v
}
