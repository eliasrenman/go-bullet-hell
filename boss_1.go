package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var bossOneImage = LoadImage("characters/boss_one.png", OriginCenter)

func NewBossOne(position Vector) *BossOne {
	e := &Entity{
		Position: position,
	}
	boss := &BossOne{
		Entity: e,
		Health: &Health{
			MaxHitPoints: 100,
			HitPoints:    100,
		},
		Hitbox: &RectangleHitbox{
			Size: Vector{X: 32, Y: 32},
			Hitbox: Hitbox{
				Position: Vector{X: -16, Y: -16},
				Owner:    e,
			},
		},
	}
	return boss
}

type BossOne struct {
	*Entity
	*Health
	Hitbox *RectangleHitbox
}

func (boss *BossOne) Draw(screen *ebiten.Image) {
	bossOneImage.Draw(screen, boss.Position, Vector{X: 1, Y: 1}, 0)
	if HitboxesVisible {
		boss.Hitbox.Draw(screen, boss.Position)
	}
}

var schedule = Schedule{}

func (boss *BossOne) Start() {
}

func (boss *BossOne) Update(game *Game) {
	schedule.Update(boss.Entity)
	boss.checkBulletCollision(game.player)
	if boss.Health.HitPoints == 0 {
		Destroy(boss)
	}
}

func (boss *BossOne) checkBulletCollision(player *Player) {
	EachGameObject(func(obj GameObject, layer int) {
		bullet, ok := obj.(*Bullet)
		if ok && bullet.Owner == player.Entity && CollidesAt(boss.Hitbox, boss.Position, bullet.Hitbox, bullet.Position) {
			boss.Health.TakeDamage(bullet)
			Destroy(bullet)
		}
	}, BulletLayer)
}

func (boss *BossOne) Die() {
	EachGameObject(func(obj GameObject, layer int) {
		// Cleanup bullets
		bullet, ok := obj.(*Bullet)
		if ok && bullet.Owner == boss.Entity {
			Destroy(bullet)
		}
	}, BulletLayer)

	println("Boss One died")
}
