package entity

import (
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/util"
)

type CircleHitbox struct {
	BaseHitbox
	Radius float64
}

func (a *CircleHitbox) CollidesWithCircle(p geometry.Point, b *CircleHitbox) bool {
	aCenter := a.Position.Plus(p)
	bCenter := b.Position.Plus(b.Owner.Position)

	return aCenter.Distance(bCenter) < a.Radius+b.Radius
}

func (a *CircleHitbox) CollidesWithRectangle(p geometry.Point, b *RectangleHitbox) bool {
	aCenter := a.Position.Plus(p)
	bMin, bMax := b.GetRectangle(b.Owner.Position)

	// Find the closest point to the circle within the rectangle
	x := util.ClampFloat(aCenter.X, bMin.X, bMax.X)
	y := util.ClampFloat(aCenter.Y, bMin.Y, bMax.Y)

	// Calculate the distance between the circle's center and this closest point
	distance := aCenter.Distance(geometry.Point{X: x, Y: y})

	// If the distance is less than the circle's radius, an intersection occurs
	return distance < a.Radius
}

func (hb CircleHitbox) CollidesWith(other Collider) bool {
	return other.CollidesAt(hb.Owner.Position, &hb)
}

func (a CircleHitbox) CollidesAt(p geometry.Point, b Collider) bool {
	switch other := b.(type) {
	case *CircleHitbox:
		return a.CollidesWithCircle(p, other)
	default:
		return false
	}
}
