package main

import (
	"time"

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

func NewPlayer(position Vector) *Player {
	player := &Player{
		Entity: &Entity{
			Position: position,
		},
		ShootSpeed: 10,
		CanShoot:   true,
	}

	player.MoveHitbox = &RectangleHitbox{
		Size: Vector{X: 32, Y: 32},
		BaseHitbox: BaseHitbox{
			Position: Vector{X: -16, Y: -16},
			Owner:    player.Entity,
		},
	}

	player.DamageHitbox = &RectangleHitbox{
		Size: Vector{X: 16, Y: 16},
		BaseHitbox: BaseHitbox{
			Position: Vector{},
			Owner:    player.Entity,
		},
	}

	return player
}

func (player *Player) Start() {}

var gameFieldHitbox = &RectangleHitbox{
	BaseHitbox: BaseHitbox{
		Position: Vector{X: 32, Y: 32},
	},
	Size: PlayfieldSize.Minus(Vector{X: 64, Y: 64}),
}

func (player *Player) Update() {
	// Handle movement
	moveInput := Vector{
		X: AxisHorizontal.Get(0),
		Y: -AxisVertical.Get(0),
	}
	direction := moveInput.Angle()
	speed := 0.
	if moveInput.X != 0 || moveInput.Y != 0 {
		if ButtonSlow.Get(0) {
			speed = moveSpeedSlow
		} else {
			speed = moveSpeed
		}
	}

	// Allow sliding against walls
	for i := 0.; i < 60 && i > -60; i = -(i + Sign(i)) {
		mv := VectorFromAngle(direction + DegToRad(i)).ScaledBy(speed)
		if player.MoveHitbox.CollidesAt(player.Position.Plus(mv), gameFieldHitbox) {
			player.Velocity = mv
			player.Move(mv)
			break
		}
	}

	// Handle shooting
	if player.CanShoot && ButtonShoot.Get(0) {
		if time.Since(player.lastShootTime) > time.Second/time.Duration(player.ShootSpeed) {
			player.Shoot(
				player.Position.Copy().Minus(Vector{X: 0, Y: 0}),
				DegToRad(-90),
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
	playerImage      = LoadImage("characters/player_forward.png", OriginCenter)
	playerLeftImage  = LoadImage("characters/player_left.png", OriginCenter)
	playerRightImage = LoadImage("characters/player_right.png", OriginCenter)
)

func (player *Player) Draw(screen *ebiten.Image) {
	hAxis := AxisHorizontal.Get(0)
	image := playerImage

	if hAxis < 0 {
		image = playerLeftImage
	} else if hAxis > 0 {
		image = playerRightImage
	}

	image.Draw(screen, player.Position, Vector{X: 1, Y: 1}, 0)
}
