package entity

import (
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/util"
)

type CircleHitbox struct {
	BaseHitbox
	Radius float64
}

func (a *CircleHitbox) CollidesWithCircle(b CircleHitbox) bool {
	aCenter := a.Position.Plus(a.Owner.Position)
	bCenter := b.Position.Plus(b.Owner.Position)

	return aCenter.Distance(bCenter) < a.Radius+b.Radius
}

func (a *CircleHitbox) CollidesWithRectangle(b RectangleHitbox) bool {
	aCenter := a.Position.Plus(a.Owner.Position)
	bMin, bMax := b.GetRectangle()

	// Find the closest point to the circle within the rectangle
	x := util.ClampFloat(aCenter.X, bMin.X, bMax.X)
	y := util.ClampFloat(aCenter.Y, bMin.Y, bMax.Y)

	// Calculate the distance between the circle's center and this closest point
	distance := aCenter.Distance(geometry.Point{X: x, Y: y})

	// If the distance is less than the circle's radius, an intersection occurs
	return distance < a.Radius
}

func (hb CircleHitbox) CollidesWith(other Collider) bool {
	switch other := other.(type) {
	case CircleHitbox:
		return hb.CollidesWithCircle(other)
	default:
		return false
	}
}

func (hb CircleHitbox) CollidesAt(p geometry.Point, other Collider) bool {
	hb.Position.Add(p)
	defer hb.Position.Subtract(p)
	return hb.CollidesWith(other)
}
