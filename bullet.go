package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Bullets are Entities with additional values for Damage, Size, Speed and Direction
type Bullet struct {
	Entity
	Owner  *Entity
	Damage int

	// Read-only! Use SetAngularVelocity to set the speed
	Speed float64
	// Read-only! Use SetAngularVelocity to set the direction
	Direction float64
}

func (b *Bullet) SetAngularVelocity(speed float64, direction float64) {
	b.Speed = speed
	b.Direction = direction
	b.Velocity = VectorFromAngle(direction).ScaledBy(speed)
}

func (owner *Entity) Shoot(position Point, direction float64, speed float64, offset float64) {

	// This offests the inital position based on the direction of the bullet.
	position.Add(VectorFromAngle(direction).ScaledBy(offset))

	bullet := Spawn(&Bullet{
		Entity: Entity{Position: position},
		Owner:  owner,
	})
	bullet.SetAngularVelocity(speed, direction)

}

func (b *Bullet) Start() {}

func (b *Bullet) Update() {
	b.Move(b.Velocity)
	if b.Position.Y < 0 || b.Position.Y > float64(SCREEN_HEIGHT) {

		Destroy(b)
	}
}

var bulletImage = LoadImage("bullets/bullet.png", OriginCenter)

func (b *Bullet) Draw(screen *ebiten.Image) {
	bulletImage.Draw(screen, b.Position, Size{Width: 1, Height: 1}, 0)
}

func (b *Bullet) Die() {
}
