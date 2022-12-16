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
	*Entity

	// Bullets shot per second
	ShootSpeed    float64
	CanShoot      bool
	lastShootTime time.Time

	Hitbox Hitbox
}

func NewPlayer(position geometry.Point) *Player {
	entity := &Entity{
		Position: position,
	}
	player := &Player{
		Entity: entity,
		Hitbox: Hitbox{
			MinPoint: geometry.Point{X: 0, Y: 0},
			MaxPoint: geometry.Point{X: 1, Y: 1},
			Entity:   entity,
		},
		ShootSpeed: 10,
		CanShoot:   true,
	}
	return player
}

func (player *Player) Start() {}

var gameFieldHitbox = NewFieldHitbox()

func (player *Player) Update() {
	// Handle movement
	move := geometry.Vector{
		X: input.AxisHorizontal.Get(0),
		Y: -input.AxisVertical.Get(0),
	}

	// Make sure to check border colision and cancel out movement.
	borderColisionVector := gameFieldHitbox.Inside(player.Hitbox)
	move.Add(borderColisionVector)

	// Make sure to normalise the movement.
	move.Normalized()

	speed := moveSpeed
	if input.ButtonSlow.Get(0) {
		speed = moveSpeedSlow
	}

	move.Scale(speed)
	player.Move(move)
	player.Velocity = move

	// println(colision.X, colision.Y)
	// Handle shooting
	if player.CanShoot && input.ButtonShoot.Get(0) {
		// Normalize the amount of bullets being shot.
		if time.Since(player.lastShootTime) > time.Second/time.Duration(player.ShootSpeed) {
			player.Shoot(
				player.Position.Copy().Minus(geometry.Vector{X: 0, Y: 0}),
				util.DegToRad(-90),
				6,
				25,
			)

			player.lastShootTime = time.Now()
		}
	}
}

func (player *Player) Die() {
	// Make sure to clean up all the players bullets
	for entity := range GameObjects {
		bullet, ok := entity.(*Bullet)
		if ok {
			Destroy(bullet)
		}
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
