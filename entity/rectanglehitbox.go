package entity

import (
	"fmt"

	"github.com/eliasrenman/go-bullet-hell/geometry"
)

type RectangleHitbox struct {
	BaseHitbox
	Size geometry.Size
}

func (a *RectangleHitbox) CollidesWithRectangle(b RectangleHitbox) bool {
	aMin, aMax := a.GetRectangle()
	bMin, bMax := b.GetRectangle()

	fmt.Println(aMin, aMax, bMin, bMax)

	return aMin.X > bMin.X && aMin.Y > bMin.Y && aMax.X < bMax.X && aMax.Y < bMax.Y
}

func (a *RectangleHitbox) CollidesWithCircle(b CircleHitbox) bool {
	return b.CollidesWithRectangle(*a)
}

func (hb *RectangleHitbox) GetRectangle() (geometry.Point, geometry.Point) {
	if hb.Owner == nil {
		return hb.Position, hb.Position.Plus(hb.Size.AsVector())
	}

	return hb.Position.Plus(hb.Owner.Position),
		hb.Position.Plus(hb.Owner.Position).Plus(hb.Size.AsVector())
}

func (hb RectangleHitbox) CollidesWith(other Collider) bool {
	switch other := other.(type) {
	case RectangleHitbox:
		return hb.CollidesWithRectangle(other)
	default:
		return false
	}
}

func (hb RectangleHitbox) CollidesAt(p geometry.Point, other Collider) bool {
	d := p.Minus(hb.Position)
	hb.Position.Add(d)
	defer hb.Position.Subtract(d)
	return hb.CollidesWith(other)
}
