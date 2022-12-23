package main

import "math"

// Vector represents a 2D vector
type Vector struct {
	X float64
	Y float64
}

var (
	// Zero is a vector with both components set to 0
	Zero = Vector{X: 0, Y: 0}
	// Up is a vector pointing up
	Up = Vector{X: 0, Y: -1}
	// Down is a vector pointing down
	Down = Vector{X: 0, Y: 1}
	// Left is a vector pointing left
	Left = Vector{X: -1, Y: 0}
	// Right is a vector pointing right
	Right = Vector{X: 1, Y: 0}
)

// VectorFromAngle creates a vector given an angle in radians
func VectorFromAngle(angle float64) Vector {
	return Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}

// Magnitude calculates the magnitude (length) of the vector
func (vector Vector) Magnitude() float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y)
}

// Direction calculates the direction of the vector
func (vector *Vector) Direction() float64 {
	return math.Tan(vector.Y / vector.X)
}

// Normalize normalizes the vector, setting its magnitude to 1
func (vector *Vector) Normalize() {
	magnitude := vector.Magnitude()

	if magnitude == 0 {
		return // Avoid divide-by-zero overflow
	}

	vector.X /= magnitude
	vector.Y /= magnitude
}

// Normalized creates a normalized copy of the vector
func (vector Vector) Normalized() Vector {
	v := vector.Copy()
	v.Normalize()

	return v
}

// Angle returns the angle of the vector in radians
func (vector Vector) Angle() float64 {
	return math.Atan2(vector.Y, vector.X)
}

// Copy creates a copy of the vector
func (vector Vector) Copy() Vector {
	v := Vector{vector.X, vector.Y}
	return v
}

// Distance calculates the distance between two vectors
func (vector Vector) Distance(other Vector) float64 {
	return vector.Minus(other).Magnitude()
}

// Pointer arithmetic
// modifies the original variable

// Add adds another vector to the vector
func (vector *Vector) Add(other Vector) *Vector {
	vector.X += other.X
	vector.Y += other.Y
	return vector
}

// Subtract subtracts another vector from the vector
func (vector *Vector) Subtract(other Vector) *Vector {
	vector.X -= other.X
	vector.Y -= other.Y
	return vector
}

// Multiply multiplies the vector by another vector
func (vector *Vector) Multiply(other Vector) *Vector {
	vector.X *= other.X
	vector.Y *= other.Y
	return vector
}

// Divide divides the vector by another vector
func (vector *Vector) Divide(other Vector) *Vector {
	vector.X /= other.X
	vector.Y /= other.Y
	return vector
}

// Scale scales the vector by a scalar value
func (vector *Vector) Scale(value float64) *Vector {
	vector.X *= value
	vector.Y *= value
	return vector
}

// Value arithmetic
// creates a new variable

// Plus creates a new vector representing the sum of two vectors
func (vector Vector) Plus(other Vector) Vector {
	v := vector.Copy()
	v.Add(other)
	return v
}

// Minus creates a new vector representing the difference between two vectors
func (vector Vector) Minus(other Vector) Vector {
	v := vector.Copy()
	v.Subtract(other)
	return v
}

// Dot creates a new vector representing the dot product of two vectors
func (vector Vector) Dot(other Vector) Vector {
	v := vector.Copy()
	v.Multiply(other)
	return v
}

// DividedBy creates a new vector representing the quotient between two vectors
func (vector Vector) DividedBy(other Vector) Vector {
	v := vector.Copy()
	v.Divide(other)
	return v
}

// ScaledBy creates a new vector representing the product of a vector and a scalar
func (vector Vector) ScaledBy(value float64) Vector {
	v := vector.Copy()
	v.Scale(value)
	return v
}
