package entity

import (
	"time"

	"github.com/eliasrenman/go-bullet-hell/assets"
)

type Player struct {
	BulletOwner
	Entity

	movingSlowly bool
	shooting     bool
	shootTimer   time.Timer
	// Bullets shot per second
	shootSpeed float64
}

func (player *Player) Update() {

}

var (
	playerImage = assets.LoadImage("characters/player_forward.png", assets.OriginCenter)
)

func (player *Player) Draw() {
	// playerImage.Draw()
}
