package main

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

// Collidable structs can be checked for collisions
type Collidable interface {
	Draw(image *ebiten.Image, position Vector)
}

// Hitbox is a base struct for all hitboxes
type Hitbox struct {
	Position Vector
	Owner    *Entity
}

// CircleHitbox is a perfectly circular hitbox
type CircleHitbox struct {
	Hitbox
	Radius float64
	image  *ebiten.Image
}

// Draw function that will draw the CircleHitboxes borders
func (hb CircleHitbox) Draw(screen *ebiten.Image, position Vector) {

	op := &ebiten.DrawImageOptions{}
	position.Add(hb.Position)
	op.GeoM.Translate(position.X, position.Y)
	if hb.image == nil {
		hb.image = generateCircleHitboxImage(&hb)
	}
	screen.DrawImage(hb.image, op)
}
func generateCircleHitboxImage(hb *CircleHitbox) *ebiten.Image {
	dc := gg.NewContext(int(hb.Radius*2), int(hb.Radius*2))
	dc.SetRGB255(255, 0, 255)
	// Draw the borders around the min and max of the rectangle in purple
	drawLineX(dc, 0, 0, hb.Radius*2)
	drawLineX(dc, 0, hb.Radius*2-1, hb.Radius*2)
	drawLineY(dc, 0, 0, hb.Radius*2)
	drawLineY(dc, hb.Radius*2-1, 0, hb.Radius*2)

	return ebiten.NewImageFromImage(dc.Image())
}

// RectangleHitbox is a rectangular hitbox
type RectangleHitbox struct {
	Hitbox
	Size  Vector
	image *ebiten.Image
}

// Draw function that will draw the RectangleHitboxes borders
func (hb RectangleHitbox) Draw(screen *ebiten.Image, position Vector) {

	op := &ebiten.DrawImageOptions{}
	position.Add(hb.Position)
	op.GeoM.Translate(position.X, position.Y)
	if hb.image == nil {
		hb.image = generateRectHitboxImage(&hb)
	}
	screen.DrawImage(hb.image, op)
}

func generateRectHitboxImage(hb *RectangleHitbox) *ebiten.Image {
	dc := gg.NewContext(int(hb.Size.X), int(hb.Size.Y))
	dc.SetRGB255(255, 0, 255)
	drawLineX(dc, 0, 0, hb.Size.X)
	drawLineX(dc, 0, hb.Size.Y-1, hb.Size.X)
	drawLineY(dc, 0, 0, hb.Size.Y)
	drawLineY(dc, hb.Size.X-1, 0, hb.Size.Y)
	return ebiten.NewImageFromImage(dc.Image())
}

func drawLineX(screen *gg.Context, x, y, length float64) {
	for x := 0.; x < length; x++ {
		screen.SetPixel(int(x), int(y))
	}
}
func drawLineY(screen *gg.Context, x, y, length float64) {
	for y := 0.; y < length; y++ {
		screen.SetPixel(int(x), int(y))
	}
}
func collisionRectangleRectangle(a *RectangleHitbox, aPos Vector, b *RectangleHitbox, bPos Vector) bool {
	aMin, aMax := a.getRectangle(aPos)
	bMin, bMax := b.getRectangle(bPos)

	return aMin.X < bMax.X && aMax.X > bMin.X && aMin.Y < bMax.Y && aMax.Y > bMin.Y
}

func collisionRectangleCircle(a *RectangleHitbox, aPos Vector, b *CircleHitbox, bPos Vector) bool {
	aMin, aMax := a.getRectangle(aPos)

	x := ClampFloat(bPos.X, aMin.X, aMax.X)
	y := ClampFloat(bPos.Y, aMin.Y, aMax.Y)

	return bPos.Distance(Vector{X: x, Y: y}) < b.Radius
}

func collisionCircleCircle(a *CircleHitbox, aPos Vector, b *CircleHitbox, bPos Vector) bool {
	return aPos.DistanceSquared(bPos) < (a.Radius+b.Radius)*(a.Radius+b.Radius)
}

// CollidesAt checks if two Collidables collide at the given positions
func CollidesAt(a Collidable, aPos Vector, b Collidable, bPos Vector) bool {
	switch aBox := a.(type) {
	case *CircleHitbox:
		switch bBox := b.(type) {
		case *CircleHitbox:
			return collisionCircleCircle(aBox, aPos, bBox, bPos)
		case *RectangleHitbox:
			return collisionRectangleCircle(bBox, bPos, aBox, aPos)
		}

	case *RectangleHitbox:
		switch bBox := b.(type) {
		case *CircleHitbox:
			return collisionRectangleCircle(aBox, aPos, bBox, bPos)
		case *RectangleHitbox:
			return collisionRectangleRectangle(aBox, aPos, bBox, bPos)
		}
	}

	return false
}

func (hb *RectangleHitbox) getRectangle(p Vector) (Vector, Vector) {

	if hb.Owner == nil {
		return hb.Position, hb.Position.Plus(hb.Size)
	}

	return hb.Position.Plus(p),
		hb.Position.Plus(p).Plus(hb.Size)
}
