package main

import (
	"image/color"

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
}

// Draw function that will draw the CircleHitboxes borders
func (hb *CircleHitbox) Draw(screen *ebiten.Image, position Vector) {
	Image := ebiten.NewImage(int(hb.Radius*2), int(hb.Radius*2))

	op := &ebiten.DrawImageOptions{}
	position.Add(hb.Position)
	op.GeoM.Translate(position.X, position.Y)

	purpleCol := color.RGBA{255, 0, 255, 255}
	// Draw the borders around the min and max of the rectangle in purple
	drawLineX(Image, 0, 0, hb.Radius*2, purpleCol)
	drawLineX(Image, 0, hb.Radius*2-1, hb.Radius*2, purpleCol)
	drawLineY(Image, 0, 0, hb.Radius*2, purpleCol)
	drawLineY(Image, hb.Radius*2-1, 0, hb.Radius*2, purpleCol)
	screen.DrawImage(Image, op)
}

// RectangleHitbox is a rectangular hitbox
type RectangleHitbox struct {
	Hitbox
	Size Vector
}

// Draw function that will draw the RectangleHitboxes borders
func (hb *RectangleHitbox) Draw(screen *ebiten.Image, position Vector) {

	Image := ebiten.NewImage(int(hb.Size.X), int(hb.Size.Y))

	op := &ebiten.DrawImageOptions{}
	position.Add(hb.Position)
	op.GeoM.Translate(position.X, position.Y)

	purpleCol := color.RGBA{255, 0, 255, 255}
	// Draw the borders around the min and max of the rectangle in purple
	drawLineX(Image, 0, 0, hb.Size.X, purpleCol)
	drawLineX(Image, 0, hb.Size.Y-1, hb.Size.X, purpleCol)
	drawLineY(Image, 0, 0, hb.Size.Y, purpleCol)
	drawLineY(Image, hb.Size.X-1, 0, hb.Size.Y, purpleCol)
	screen.DrawImage(Image, op)
}

func drawLineX(screen *ebiten.Image, x, y, length float64, col color.RGBA) {
	for x := 0.; x < length; x++ {
		screen.Set(int(x), int(y), col)
	}
}
func drawLineY(screen *ebiten.Image, x, y, length float64, col color.RGBA) {
	for y := 0.; y < length; y++ {
		screen.Set(int(x), int(y), col)
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
	return aPos.Distance(bPos) < a.Radius+b.Radius
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
