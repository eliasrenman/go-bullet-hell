package main

// An Entity is an object in space
// it has a position and a velocity
type Entity struct {
	Position Point
	Velocity Vector
}

func (entity *Entity) Move(vector Vector) {
	entity.Position.X += vector.X
	entity.Position.Y += vector.Y
}
