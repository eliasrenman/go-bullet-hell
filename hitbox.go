package main

type Collider interface {
	CollidesWith(compare Collider) bool
	CollidesAt(p Vector, compare Collider) bool
}

type BaseHitbox struct {
	Position Vector
	Owner    *Entity
}

type CircleHitbox struct {
	BaseHitbox
	Radius float64
}

func (a *CircleHitbox) CollidesWithCircle(p Vector, b *CircleHitbox) bool {
	aCenter := a.Position.Plus(p)
	bCenter := b.Position.Plus(b.Owner.Position)

	return aCenter.Distance(bCenter) < a.Radius+b.Radius
}

func (a *CircleHitbox) CollidesWithRectangle(p Vector, b *RectangleHitbox) bool {
	aCenter := a.Position.Plus(p)
	bMin, bMax := b.GetRectangle(b.Owner.Position)

	// Find the closest point to the circle within the rectangle
	x := ClampFloat(aCenter.X, bMin.X, bMax.X)
	y := ClampFloat(aCenter.Y, bMin.Y, bMax.Y)

	// Calculate the distance between the circle's center and this closest point
	distance := aCenter.Distance(Vector{X: x, Y: y})

	// If the distance is less than the circle's radius, an intersection occurs
	return distance < a.Radius
}

func (hb CircleHitbox) CollidesWith(other Collider) bool {
	return other.CollidesAt(hb.Owner.Position, &hb)
}

func (a CircleHitbox) CollidesAt(p Vector, b Collider) bool {
	switch other := b.(type) {
	case *CircleHitbox:
		return a.CollidesWithCircle(p, other)
	default:
		return false
	}
}

type RectangleHitbox struct {
	BaseHitbox
	Size Vector
}

func (a *RectangleHitbox) CollidesWithRectangle(p Vector, b *RectangleHitbox) bool {
	aMin, aMax := a.GetRectangle(p)
	bMin, bMax := b.GetRectangle(Vector{})
	if b.Owner != nil {
		bMin, bMax = b.GetRectangle(b.Owner.Position)
	}

	return aMin.X < bMax.X && aMax.X > bMin.X && aMin.Y < bMax.Y && aMax.Y > bMin.Y
}

func (a *RectangleHitbox) CollidesWithCircle(p Vector, b *CircleHitbox) bool {
	return b.CollidesWithRectangle(p, a)
}

func (hb *RectangleHitbox) GetRectangle(p Vector) (Vector, Vector) {
	if hb.Owner == nil {
		return hb.Position, hb.Position.Plus(hb.Size)
	}

	return hb.Position.Plus(p),
		hb.Position.Plus(p).Plus(hb.Size)
}

func (hb *RectangleHitbox) CollidesWith(other Collider) bool {
	return other.CollidesAt(hb.Owner.Position, hb)
}

func (a *RectangleHitbox) CollidesAt(p Vector, b Collider) bool {
	switch other := b.(type) {
	case *RectangleHitbox:
		return a.CollidesWithRectangle(p, other)
	default:
		return false
	}
}
