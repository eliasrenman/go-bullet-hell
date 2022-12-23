package main

// Collidable structs can be checked for collisions
type Collidable = any

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

// RectangleHitbox is a rectangular hitbox
type RectangleHitbox struct {
	Hitbox
	Size Vector
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
