package entity

import (
	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

// Bullets are Entities with additional values for Damage, Size, Speed and Direction
type Bullet struct {
	Entity
	Damage int

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
	Bullets map[*Bullet]struct{}
}

func NewBulletOwner() BulletOwner {
	return BulletOwner{
		Bullets: make(map[*Bullet]struct{}),
	}
}

func (owner *BulletOwner) Shoot(position geometry.Point, direction float64, speed float64) {
	bullet := Spawn(&Bullet{
		Entity: Entity{Position: position},
	})
	bullet.SetAngularVelocity(speed, direction)

	// Add a reference to the bullet in the owner's bullet set
	owner.Bullets[bullet] = struct{}{}
}

func (b *Bullet) Start() {}

func (b *Bullet) Update() {
	b.Move(b.Velocity)
}

var bulletImage = assets.LoadImage("bullets/bullet.png", assets.OriginCenter)

func (b *Bullet) Draw(screen *ebiten.Image) {
	bulletImage.Draw(screen, b.Position, geometry.Size{Width: 1, Height: 1}, 0)
}

func (b *Bullet) Die() {}
