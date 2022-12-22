package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var bossOneImage = LoadImage("characters/boss_one.png", OriginCenter)

func NewBossOne(position Point) *BossOne {
	e := &Entity{
		Position: position,
	}
	boss := &BossOne{
		Entity: e,
		Hitbox: &RectangleHitbox{
			Size: Size{Width: 32, Height: 32},
			BaseHitbox: BaseHitbox{
				Position: Vector{X: -16, Y: -16},
				Owner:    e,
			},
		},
	}
	return boss
}

type BossOne struct {
	*Entity
	Hitbox *RectangleHitbox
}

func (boss *BossOne) Draw(screen *ebiten.Image) {
	bossOneImage.Draw(screen, boss.Position, Size{Width: 1, Height: 1}, 0)
}

var schedule = Schedule{
	Patterns: []Pattern{
		{
			Type: "arc",
			Options: map[string]interface{}{
				"count": 30,
				"speed": 1.0,
				"from":  0.25 * math.Pi,
				"to":    0.75 * math.Pi,
			},
			Duration: 1 * time.Second,
			Cooldown: 1 * time.Second,
		},
		{
			Type: "staggeredCircle",
			Options: map[string]interface{}{
				"count": 30,
				"speed": 1.0,
				"from":  0 * math.Pi,
				"to":    2 * math.Pi,
			},
			Duration: 1 * time.Second,
			Cooldown: 1 * time.Second,
		},
	},
}

func (boss *BossOne) Start() {
}

func (boss *BossOne) Update() {
	schedule.Update(boss.Entity)
}

func (boss *BossOne) Die() {

}
