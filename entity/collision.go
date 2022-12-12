package entity

import "github.com/eliasrenman/go-bullet-hell/geometry"

func Move(entity *Entity, vector geometry.Vector) {
	entity.Position.X += vector.X
	entity.Position.Y += vector.Y
	//if (
	//	rect1.x < rect2.x + rect2.w &&
	//	rect1.x + rect1.w > rect2.x &&
	//	rect1.y < rect2.y + rect2.h &&
	//	rect1.h + rect1.y > rect2.y
	//
	//  ) {
	// TODO: Implement some kind of collision system here
}
