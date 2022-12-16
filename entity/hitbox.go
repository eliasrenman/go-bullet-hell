package entity

import (
	"github.com/eliasrenman/go-bullet-hell/geometry"
)

type Collider interface {
	CollidesWith(compare Collider) bool
	CollidesAt(p geometry.Point, compare Collider) bool
}

type BaseHitbox struct {
	Position geometry.Point
	Owner    *Entity
}
