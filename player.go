package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const moveSpeed float64 = 4
const moveSpeedSlow float64 = 2

// Player is the player character
type Player struct {
	*Entity
	*Health
	// Bullets shot per second
	ShootSpeedThrottler Throttler
	CanShoot            bool

	MoveHitbox          *RectangleHitbox
	DamageHitbox        *CircleHitbox
	showHitbox          bool
	hit                 bool
	hitCleanupCounter   int
	hitCleanupThrottler Throttler
}

// NewPlayer creates a new player instance
func NewPlayer(position Vector) *Player {
	player := &Player{
		Entity: &Entity{
			Position: position,
		},
		Health: &Health{
			MaxHitPoints: 5,
			HitPoints:    5,
		},
		ShootSpeedThrottler: Throttler{
			RatePerSecond: 10,
		},
		CanShoot: true,
		hitCleanupThrottler: Throttler{
			RatePerSecond: 30,
		},
	}

	player.MoveHitbox = &RectangleHitbox{
		Size: Vector{X: 32, Y: 32},
		Hitbox: Hitbox{
			Position: Vector{X: -16, Y: -16},
			Owner:    player.Entity,
		},
	}

	player.DamageHitbox = &CircleHitbox{
		Radius: 4,
		Hitbox: Hitbox{
			Position: Vector{X: -4, Y: -4},
			Owner:    player.Entity,
		},
	}

	return player
}

// Start is called when the player is added to the game
func (player *Player) Start() {}

var gameFieldHitbox = &RectangleHitbox{
	Hitbox: Hitbox{
		Position: Vector{X: 32, Y: 32},
	},
	Size: PlayfieldSize.Minus(Vector{X: 64, Y: 64}),
}

// Update is called every game tick, and handles player behavior
func (player *Player) Update(game *Game) {
	// Handle movement
	moveInput := Vector{
		X: AxisHorizontal.Get(0),
		Y: -AxisVertical.Get(0),
	}
	direction := moveInput.Angle()
	speed := 0.
	moveSlow := ButtonSlow.Get(0)
	if moveInput.X != 0 || moveInput.Y != 0 {
		if moveSlow {

			speed = moveSpeedSlow
		} else {
			speed = moveSpeed
		}
	}
	player.showHitbox = moveSlow

	// Allow sliding against walls
	for i := 0.; i < 60 && i > -60; i = -(i + Sign(i)) {
		mv := VectorFromAngle(direction + DegToRad(i)).ScaledBy(speed)
		if CollidesAt(player.MoveHitbox, player.Position.Plus(mv), gameFieldHitbox, Vector{}) {
			player.Velocity = mv
			player.Move(mv)
			break
		}
	}
	player.checkBulletCollision()

	if player.Health.HitPoints == 0 {
		game.GameOver()
		return
	}

	// Handle shooting
	if player.CanShoot && ButtonShoot.Get(0) {
		if player.ShootSpeedThrottler.CanCall() {
			player.Shoot(
				player.Position.Copy().Minus(Vector{X: 0, Y: 0}),
				DegToRad(-90),
				6,
				25,
				0,
				2,
			)

			player.ShootSpeedThrottler.Call()
		}
	}

	if player.hit && player.hitCleanupThrottler.CanCall() {
		if player.hitCleanupCounter > 23 {
			player.hit = false
			player.hitCleanupCounter = 0
		} else {
			player.hitCleanupCounter++
		}
		player.hitCleanupThrottler.Call()
	}
}

func (player *Player) checkBulletCollision() {
	bulletsInMoveHitbox := make(map[*Bullet]struct{})
	cleanupHitbox := &CircleHitbox{
		Radius: float64(player.hitCleanupCounter),
	}
	for _, objects := range GameObjects {
		for obj := range objects {
			bullet, ok := obj.(*Bullet)

			if ok && bullet.Owner != player.Entity {
				if CollidesAt(cleanupHitbox, player.Position, bullet.Hitbox, bullet.Position) {
					bulletsInMoveHitbox[bullet] = struct{}{}
				}

				if !player.hit && CollidesAt(player.DamageHitbox, player.Position, bullet.Hitbox, bullet.Position) {
					player.hit = true
					player.hitCleanupCounter = 0

					player.Health.TakeDamage(bullet)
					Destroy(bullet)
				}
			}
		}
	}
	if player.hit {
		for b := range bulletsInMoveHitbox {
			Destroy(b)
		}
	}
}

// Die is called when the player dies
func (player *Player) Die() {
	// Make sure to clean up all the players bullets
	EachGameObject(func(obj GameObject, layer int) {
		bullet, ok := obj.(*Bullet)
		if ok && bullet.Owner == player.Entity {
			Destroy(obj)
		}
	}, BulletLayer)
}

var (
	playerImage      = LoadImage("characters/player_forward.png", OriginCenter)
	playerLeftImage  = LoadImage("characters/player_left.png", OriginCenter)
	playerRightImage = LoadImage("characters/player_right.png", OriginCenter)
	playerHitbox     = LoadImage("characters/player_hitbox.png", OriginCenter)
)

// Draw is called every frame to draw the player
func (player *Player) Draw(screen *ebiten.Image) {
	hAxis := AxisHorizontal.Get(0)
	image := playerImage

	if hAxis < 0 {
		image = playerLeftImage
	} else if hAxis > 0 {
		image = playerRightImage
	}

	image.Draw(screen, player.Position, Vector{X: 1, Y: 1}, 0)
	if player.showHitbox {
		playerHitbox.Draw(screen, player.Position, Vector{X: 1, Y: 1}, 0)

	}
	if HitboxesVisible {
		player.MoveHitbox.Draw(screen, player.Position)
		player.DamageHitbox.Draw(screen, player.Position)
	}
}
