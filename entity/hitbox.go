package entity

import (
	"github.com/eliasrenman/go-bullet-hell/constant"
	"github.com/eliasrenman/go-bullet-hell/geometry"
)

type Hitbox struct {
	MinPoint geometry.Point
	MaxPoint geometry.Point
	Entity   *Entity
}

/*
This checks that this Hitbox has the compare hitbox inside its box
@param compare is inside the hitbox
*/
func (hitbox *Hitbox) Inside(compare Hitbox) geometry.Vector {
	x, y := 0., 0.

	hitboxMin := hitbox.GetMinPoint()
	hitboxMax := hitbox.GetMaxPoint()

	compareMin := compare.GetMinPoint()
	compareMax := compare.GetMaxPoint()
	if hitboxMin.X >= compareMin.X || hitboxMax.X >= compareMin.X {
		x += 1
	}
	if hitboxMin.X <= compareMax.X || hitboxMax.X <= compareMax.X {
		x -= 1
	}

	if hitboxMin.Y >= compareMin.Y || hitboxMax.Y >= compareMin.Y {
		y += 1
	}
	if hitboxMin.Y <= compareMax.Y || hitboxMax.Y <= compareMax.Y {
		y -= 1
	}

	return geometry.Vector{
		X: x,
		Y: y,
	}

}

func (hitbox *Hitbox) GetMinPoint() geometry.Point {
	return *hitbox.MinPoint.Add(hitbox.Entity.Position)
}
func (hitbox *Hitbox) GetMaxPoint() geometry.Point {
	return *hitbox.MaxPoint.Add(hitbox.Entity.Position)
}

func NewFieldHitbox() *Hitbox {
	return &Hitbox{
		MinPoint: geometry.Point{X: 0, Y: 0},
		MaxPoint: geometry.Point{X: constant.PLAYFIELD_WIDTH, Y: constant.PLAYFIELD_HEIGHT},
		Entity: &Entity{
			Velocity: geometry.Zero,
			Position: geometry.Point{X: 0, Y: 0},
		},
	}

}
