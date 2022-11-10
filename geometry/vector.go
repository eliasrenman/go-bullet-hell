package geometry

import "math"

type Vector struct {
	X float64
	Y float64
}

type Point = Vector

func VectorFromAngle(angle float64) Vector {
	return Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}

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

// Pointer arithmetic
// modifies the original variable

func (a *Vector) Add(b Vector) *Vector {
	a.X += b.X
	a.Y += b.Y
	return a
}

func (a *Vector) Subtract(b Vector) *Vector {
	a.X -= b.X
	a.Y -= b.Y
	return a
}

func (a *Vector) Multiply(b Vector) *Vector {
	a.X *= b.X
	a.Y *= b.Y
	return a
}

func (a *Vector) Divide(b Vector) *Vector {
	a.X /= b.X
	a.Y /= b.Y
	return a
}

func (a *Vector) Scale(b float64) *Vector {
	a.X *= b
	a.Y *= b
	return a
}

// Value arithmetic
// creates a new variable

func (a Vector) Plus(b Vector) Vector {
	v := a.Copy()
	v.Add(b)
	return v
}

func (a Vector) Minus(b Vector) Vector {
	v := a.Copy()
	v.Subtract(b)
	return v
}

func (a Vector) Dot(b Vector) Vector {
	v := a.Copy()
	v.Multiply(b)
	return v
}

func (a Vector) DividedBy(b Vector) Vector {
	v := a.Copy()
	v.Divide(b)
	return v
}

func (a Vector) ScaledBy(b float64) Vector {
	v := a.Copy()
	v.Scale(b)
	return v
}
