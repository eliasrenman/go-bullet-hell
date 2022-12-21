package geometry

import "math"

type Vector struct {
	X float64
	Y float64
}

var (
	Zero  = Vector{X: 0, Y: 0}
	Up    = Vector{X: 0, Y: -1}
	Down  = Vector{X: 0, Y: 1}
	Left  = Vector{X: -1, Y: 0}
	Right = Vector{X: 1, Y: 0}
)

type Point = Vector

func VectorFromAngle(angle float64) Vector {
	return Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}

// Calculates the magnitude (length) of the vector
func (vec Vector) Magnitude() float64 {
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

func (vec Vector) Angle() float64 {
	return math.Atan2(vec.Y, vec.X)
}

// Creates a copy of the vector
func (vec Vector) Copy() Vector {
	v := Vector{vec.X, vec.Y}
	return v
}

// Calculates the distance between two vectors
func (a Vector) Distance(b Vector) float64 {
	return a.Minus(b).Magnitude()
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
