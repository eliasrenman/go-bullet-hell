package entity

import (
	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

var bossOneImage = assets.LoadImage("characters/boss_one.png", assets.OriginCenter)

func NewBossOne(position geometry.Point) *BossOne {
	e := &Entity{
		Position: position,
	}
	boss := &BossOne{
		Entity: e,
		Hitbox: Hitbox{
			MinPoint: geometry.Point{X: 0, Y: 0},
			MaxPoint: geometry.Point{X: 1, Y: 1},
			Entity:   e,
		},
	}
	return boss
}

type BossOne struct {
	*Entity
	Hitbox Hitbox
}

func (boss *BossOne) Draw(screen *ebiten.Image) {
	bossOneImage.Draw(screen, boss.Position, geometry.Size{Width: 1, Height: 1}, 0)
}

func (boss *BossOne) Start() {
	// CirclePattern.Start(boss.Entity)
	StaggeredCirclePatternInstance.Start(boss.Entity)

}

func (boss *BossOne) Update() {
	StaggeredCirclePatternInstance.Update()
}

func (boss *BossOne) Die() {

}
