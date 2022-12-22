package main

type Collidable = any

type Hitbox struct {
	Position Vector
	Owner    *Entity
}

type CircleHitbox struct {
	Hitbox
	Radius float64
}

type RectangleHitbox struct {
	Hitbox
	Size Vector
}

func CollisionRectangleRectangle(a *RectangleHitbox, aPos Vector, b *RectangleHitbox, bPos Vector) bool {
	aMin, aMax := a.GetRectangle(aPos)
	bMin, bMax := b.GetRectangle(bPos)

	return aMin.X < bMax.X && aMax.X > bMin.X && aMin.Y < bMax.Y && aMax.Y > bMin.Y
}

func CollisionRectangleCircle(a *RectangleHitbox, aPos Vector, b *CircleHitbox, bPos Vector) bool {
	aMin, aMax := a.GetRectangle(aPos)

	x := ClampFloat(bPos.X, aMin.X, aMax.X)
	y := ClampFloat(bPos.Y, aMin.Y, aMax.Y)

	return bPos.Distance(Vector{X: x, Y: y}) < b.Radius
}

func CollisionCircleCircle(a *CircleHitbox, aPos Vector, b *CircleHitbox, bPos Vector) bool {
	return aPos.Distance(bPos) < a.Radius+b.Radius
}

func CollidesAt(a Collidable, aPos Vector, b Collidable, bPos Vector) bool {
	switch aBox := a.(type) {
	case *CircleHitbox:
		switch bBox := b.(type) {
		case *CircleHitbox:
			return CollisionCircleCircle(aBox, aPos, bBox, bPos)
		case *RectangleHitbox:
			return CollisionRectangleCircle(bBox, bPos, aBox, aPos)
		}

	case *RectangleHitbox:
		switch bBox := b.(type) {
		case *CircleHitbox:
			return CollisionRectangleCircle(aBox, aPos, bBox, bPos)
		case *RectangleHitbox:
			return CollisionRectangleRectangle(aBox, aPos, bBox, bPos)
		}
	}

	return false
}

func (hb *RectangleHitbox) GetRectangle(p Vector) (Vector, Vector) {
	if hb.Owner == nil {
		return hb.Position, hb.Position.Plus(hb.Size)
	}

	return hb.Position.Plus(p),
		hb.Position.Plus(p).Plus(hb.Size)
}
