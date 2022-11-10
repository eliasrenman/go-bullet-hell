package entity

import "github.com/eliasrenman/go-bullet-hell/geometry"

// Bullets are Entities with additional values for Damage, Size, Speed and Direction
type Bullet struct {
	Entity
	Damage int
	geometry.Size

	// Read-only! Use SetAngularVelocity to set the speed
	Speed float64
	// Read-only! Use SetAngularVelocity to set the direction
	Direction float64
}

func (b *Bullet) SetAngularVelocity(speed float64, direction float64) {
	b.Speed = speed
	b.Direction = direction
	b.Velocity = geometry.VectorFromAngle(direction).ScaledBy(speed)
}

// A Bullet Owner owns a number of bullets,
// they also have backreferences to their owner
type BulletOwner struct {
	Bullets []Bullet
}

func (owner *BulletOwner) Spawn(position geometry.Point, size geometry.Size, direction float64, speed float64) {
	bullet := Bullet{
		Entity: Entity{Position: position},
		Size:   size,
	}
	bullet.SetAngularVelocity(speed, direction)

	owner.Bullets = append(owner.Bullets, bullet)
}
