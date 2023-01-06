package main

import (
	"math"
	"time"

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

var schedule = Schedule{
	Patterns: []Pattern{
		{
			Type: "shoot_arc",
			Options: map[string]interface{}{
				"count": 30,
				"speed": 1.0,
				"from":  0.25 * math.Pi,
				"to":    .75 * math.Pi,
			},
			Duration:   1 * time.Second,
			Cooldown:   1 * time.Second,
			BulletType: BulletSmallYellow,
		},
		{
			Type: "move_to",
			Options: map[string]interface{}{
				"target": Vector{X: 100, Y: 100},
				"speed":  100.0,
				"easing": "quad",
			},
			Duration: 2 * time.Second,
		},
		{
			Type: "shoot_arc",
			Options: map[string]interface{}{
				"count":   30,
				"speed":   1.0,
				"from":    0 * math.Pi,
				"to":      2 * math.Pi,
				"stagger": 1.0 / 30,
			},
			BulletType: BulletSmallYellow,
		},
		{
			Type: "move_to",
			Options: map[string]interface{}{
				"target": Vector{X: 400, Y: 100},
				"speed":  100.0,
				"easing": "quad",
			},
			Duration: 1 * time.Second,
		},
		{
			Type: "shoot_arc",
			Options: map[string]interface{}{
				"count":   30,
				"speed":   1.0,
				"from":    1 * math.Pi,
				"to":      3 * math.Pi,
				"stagger": 1.0 / 30,
			},
			Duration:   2 * time.Second,
			Cooldown:   1 * time.Second,
			BulletType: BulletSmallYellow,
		},
		{
			Type: "move_to",
			Options: map[string]interface{}{
				"target": Vector{X: 250, Y: 100},
				"speed":  100.0,
				"easing": "quad",
			},
			Duration: 1 * time.Second,
		},
	},
}

func (boss *BossOne) Start() {
}

func (boss *BossOne) Update(game *Game) {
	schedule.Update(boss.Entity)
	boss.checkBulletCollision(game.player)
}

func (boss *BossOne) checkBulletCollision(player *Player) {

	for b := range BulletObjects {
		bullet, ok := b.(*Bullet)

		// This should have to make sure that the bullet is owned by the player. otherwise this could result in friendly fire from other enemies
		if ok && *bullet.Owner == *player.Entity {

			if CollidesAt(boss.Hitbox, boss.Position, bullet.Hitbox, bullet.Position) {

				Destroy(bullet)
				// Lower the hp of the boss by the value of the bullet's damage
			}
		}
	}

}

func (boss *BossOne) Die() {

}
