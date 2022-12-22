package main

type CircleHitbox struct {
	BaseHitbox
	Radius float64
}

func (a *CircleHitbox) CollidesWithCircle(p Point, b *CircleHitbox) bool {
	aCenter := a.Position.Plus(p)
	bCenter := b.Position.Plus(b.Owner.Position)

	return aCenter.Distance(bCenter) < a.Radius+b.Radius
}

func (a *CircleHitbox) CollidesWithRectangle(p Point, b *RectangleHitbox) bool {
	aCenter := a.Position.Plus(p)
	bMin, bMax := b.GetRectangle(b.Owner.Position)

	// Find the closest point to the circle within the rectangle
	x := ClampFloat(aCenter.X, bMin.X, bMax.X)
	y := ClampFloat(aCenter.Y, bMin.Y, bMax.Y)

	// Calculate the distance between the circle's center and this closest point
	distance := aCenter.Distance(Point{X: x, Y: y})

	// If the distance is less than the circle's radius, an intersection occurs
	return distance < a.Radius
}

func (hb CircleHitbox) CollidesWith(other Collider) bool {
	return other.CollidesAt(hb.Owner.Position, &hb)
}

func (a CircleHitbox) CollidesAt(p Point, b Collider) bool {
	switch other := b.(type) {
	case *CircleHitbox:
		return a.CollidesWithCircle(p, other)
	default:
		return false
	}
}
