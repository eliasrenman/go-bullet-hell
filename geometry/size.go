package geometry

type Size struct {
	Width  float64
	Height float64
}

func (s *Size) AsVector() Vector {
	return Vector{
		X: s.Width,
		Y: s.Height,
	}
}
