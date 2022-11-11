package entity

import (
	"time"

	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/input"
	"github.com/hajimehoshi/ebiten/v2"
)

const moveSpeed float64 = 4
const moveSpeedSlow float64 = 2

type Player struct {
	BulletOwner
	Entity

	shooting   bool
	shootTimer time.Timer
	// Bullets shot per second
	shootSpeed float64
}

func (player *Player) Start() {}

func (player *Player) Update() {
	move := geometry.Vector{
		X: input.AxisHorizontal.Get(0),
		Y: -input.AxisVertical.Get(0),
	}

	speed := moveSpeed
	if input.ButtonSlow.Get(0) {
		speed = moveSpeedSlow
	}

	move.Normalize()
	move.Scale(speed)

	player.Move(move)
}

func (player *Player) Die() {
	for bullet := range player.Bullets {
		Destroy(bullet)
	}
}

var (
	playerImage      = assets.LoadImage("characters/player_forward.png", assets.OriginCenter)
	playerLeftImage  = assets.LoadImage("characters/player_left.png", assets.OriginCenter)
	playerRightImage = assets.LoadImage("characters/player_right.png", assets.OriginCenter)
)

func (player *Player) Draw(screen *ebiten.Image) {
	hAxis := input.AxisHorizontal.Get(0)
	image := playerImage

	if hAxis < 0 {
		image = playerLeftImage
	} else if hAxis > 0 {
		image = playerRightImage
	}

	image.Draw(screen, player.Position, geometry.Size{Width: 1, Height: 1}, 0)
}
