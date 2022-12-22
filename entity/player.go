package entity

import (
	"time"

	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/constant"
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

	MoveHitbox   Collider
	DamageHitbox Collider
}

func NewPlayer(position geometry.Point) *Player {
	player := &Player{
		Entity: &Entity{
			Position: position,
		},
		ShootSpeed: 10,
		CanShoot:   true,
	}

	player.MoveHitbox = &RectangleHitbox{
		Size: geometry.Size{Width: 32, Height: 32},
		BaseHitbox: BaseHitbox{
			Position: geometry.Vector{X: -16, Y: -16},
			Owner:    player.Entity,
		},
	}

	player.DamageHitbox = &RectangleHitbox{
		Size: geometry.Size{Width: 16, Height: 16},
		BaseHitbox: BaseHitbox{
			Position: geometry.Vector{},
			Owner:    player.Entity,
		},
	}

	return player
}

func (player *Player) Start() {}

var gameFieldHitbox = &RectangleHitbox{
	BaseHitbox: BaseHitbox{
		Position: geometry.Vector{X: 32, Y: 32},
	},
	Size: geometry.Size{Width: constant.PLAYFIELD_WIDTH - 64, Height: constant.PLAYFIELD_HEIGHT - 64},
}

func (player *Player) Update() {
	// Handle movement
	moveInput := geometry.Vector{
		X: input.AxisHorizontal.Get(0),
		Y: -input.AxisVertical.Get(0),
	}
	direction := moveInput.Angle()
	speed := 0.
	if moveInput.X != 0 || moveInput.Y != 0 {
		if input.ButtonSlow.Get(0) {
			speed = moveSpeedSlow
		} else {
			speed = moveSpeed
		}
	}

	// Allow sliding against walls
	for i := 0.; i < 60 && i > -60; i = -(i + util.Sign(i)) {
		mv := geometry.VectorFromAngle(direction + util.DegToRad(i)).ScaledBy(speed)
		if player.MoveHitbox.CollidesAt(player.Position.Plus(mv), gameFieldHitbox) {
			player.Velocity = mv
			player.Move(mv)
			break
		}
	}

	// Handle shooting
	if player.CanShoot && input.ButtonShoot.Get(0) {
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
	EachGameObject(func(obj GameObject) {
		bullet, ok := obj.(*Bullet)
		if ok && bullet.Entity == *player.Entity {
			Destroy(bullet)
		}
	})
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
