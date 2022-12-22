package main

type Collider interface {
	CollidesWith(compare Collider) bool
	CollidesAt(p Point, compare Collider) bool
}

type BaseHitbox struct {
	Position Point
	Owner    *Entity
}
