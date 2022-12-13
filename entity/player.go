package entity

import (
	"time"

	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/input"
	"github.com/eliasrenman/go-bullet-hell/util"
	"github.com/hajimehoshi/ebiten/v2"
)

const moveSpeed float64 = 4
const moveSpeedSlow float64 = 2

type Player struct {
	BulletOwner
	Entity

	// Bullets shot per second
	ShootSpeed    float64
	CanShoot      bool
	lastShootTime time.Time
}

func NewPlayer(position geometry.Point) *Player {
	return &Player{
		Entity: Entity{
			Position: position,
		},
		BulletOwner: NewBulletOwner(),
		ShootSpeed:  10,
		CanShoot:    true,
	}
}

func (player *Player) Start() {}

func (player *Player) Update() {
	// Handle movement
	move := geometry.Vector{
		X: input.AxisHorizontal.Get(0),
		Y: -input.AxisVertical.Get(0),
	}.Normalized()

	speed := moveSpeed
	if input.ButtonSlow.Get(0) {
		speed = moveSpeedSlow
	}

	move.Scale(speed)
	player.Move(move)
	player.Velocity = move

	// Handle shooting
	if player.CanShoot && input.ButtonShoot.Get(0) {
		// Normalize the amount of bullets being shot.
		if time.Since(player.lastShootTime) > time.Second/time.Duration(player.ShootSpeed) {
			player.Shoot(
				player.Position,
				util.DegToRad(-90),
				6,
			)

			player.lastShootTime = time.Now()
		}
	}
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

	// Draw all bullets on the player
	for obj := range player.Bullets {
		obj.Draw(screen)
	}
}
