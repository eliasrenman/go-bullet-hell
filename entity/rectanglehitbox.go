package entity

import (
	"image/color"

	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

type RectangleHitbox struct {
	BaseHitbox
	Size geometry.Size
}

func (a *RectangleHitbox) CollidesWithRectangle(p geometry.Point, b *RectangleHitbox) bool {
	aMin, aMax := a.GetRectangle(p)
	bMin, bMax := b.GetRectangle(geometry.Vector{})
	if b.Owner != nil {
		bMin, bMax = b.GetRectangle(b.Owner.Position)
	}

	return aMin.X < bMax.X && aMax.X > bMin.X && aMin.Y < bMax.Y && aMax.Y > bMin.Y
}

func (a *RectangleHitbox) CollidesWithCircle(p geometry.Point, b *CircleHitbox) bool {
	return b.CollidesWithRectangle(p, a)
}

func (hb *RectangleHitbox) GetRectangle(p geometry.Point) (geometry.Point, geometry.Point) {
	if hb.Owner == nil {
		return hb.Position, hb.Position.Plus(hb.Size.AsVector())
	}

	return hb.Position.Plus(p),
		hb.Position.Plus(p).Plus(hb.Size.AsVector())
}

func (hb *RectangleHitbox) CollidesWith(other Collider) bool {
	return other.CollidesAt(hb.Owner.Position, hb)
}

func (a *RectangleHitbox) CollidesAt(p geometry.Point, b Collider) bool {
	switch other := b.(type) {
	case *RectangleHitbox:
		return a.CollidesWithRectangle(p, other)
	default:
		return false
	}
}

func (hb *RectangleHitbox) Draw(s *ebiten.Image) {
	img := ebiten.NewImage(int(hb.Size.Width), int(hb.Size.Height))
	img.Fill(color.RGBA{255, 0, 0, 25})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(hb.Position.X, hb.Position.Y)
	if hb.Owner != nil {
		op.GeoM.Translate(hb.Owner.Position.X, hb.Owner.Position.Y)
	}

	s.DrawImage(img, op)
}
