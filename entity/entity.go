package entity

import "github.com/eliasrenman/go-bullet-hell/geometry"

// An Entity is an object in space
// it has a position and a velocity
type Entity struct {
	Position geometry.Point
	Velocity geometry.Vector
}
