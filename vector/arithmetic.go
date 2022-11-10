package vector

// Pointer arithmetic
// modifies the original variable

func (a *Vector) Add(b Vector) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Vector) Subtract(b Vector) {
	a.X -= b.X
	a.Y -= b.Y
}

func (a *Vector) Multiply(b Vector) {
	a.X *= b.X
	a.Y *= b.Y
}

func (a *Vector) Divide(b Vector) {
	a.X /= b.X
	a.Y /= b.Y
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
