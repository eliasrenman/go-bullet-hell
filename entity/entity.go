package entity

import "github.com/eliasrenman/go-bullet-hell/geometry"

// An Entity is an object in space
// it has a position and a velocity
type Entity struct {
	Position geometry.Point
	Velocity geometry.Vector
}

func (entity *Entity) Move(vector geometry.Vector) {
	entity.Position.X += vector.X
	entity.Position.Y += vector.Y

	// TODO: Implement some kind of collision system here
}
