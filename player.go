package main

import (
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/util"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	playerForwardImage *ebiten.Image = LoadImage("./data/characters/player_forward.png")
	playerLeftImage    *ebiten.Image = LoadImage("./data/characters/player_left.png")
	playerRightImage   *ebiten.Image = LoadImage("./data/characters/player_right.png")
)

const (
	playerSize = 50
)

type Player struct {
	y          int
	x          int
	movingSlow bool
	bullets    *Bullets
}

func (player *Player) Draw(screen *ebiten.Image) {
	player.drawPlayer(screen)

	// Draw hitbox
	if player.movingSlow {
		hitboxOffset := playerSize/2 - hitboxDimension/2
		x, y := normalizeCoords(player.x+hitboxOffset, player.y+hitboxOffset)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)

		screen.DrawImage(hitbox, op)
	}

	// Draw regular bullets
	player.bullets.Draw(screen)
}

func (player *Player) drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x, y := normalizeCoords(player.x, player.y)
	op.GeoM.Translate(x, y)

	hInput := AXIS_HORIZONTAL.Get(0)
	image := playerForwardImage
	if hInput < 0 {
		image = playerLeftImage
	}
	if hInput > 0 {
		image = playerRightImage
	}

	screen.DrawImage(image, op)
}

func (player *Player) Update(input *Input) {
	player.updateLocation()
	player.move(player.x, player.y)
	player.bullets.Update(input)
}

func (player *Player) updateLocation() {
	// Set the apporpriate speed depending on if the slow movement is enabled
	speed := PLAYER_SPEED
	if player.movingSlow {
		speed = PLAYER_SPEED_SLOW
	}

	// Normalizing the movement vector means keyboard users
	// cannot move faster than gamepad users in diagonal directions.
	// However the game feels a bit sluggish when going diagonally
	// because of it.
	// TODO: Figure out a solution to this issue
	move := geometry.Vector{
		X: AXIS_HORIZONTAL.Get(0),
		Y: AXIS_VERTICAL.Get(0),
	}.Normalized()

	player.x += int(move.X * speed)

	// World origin is bottom-left, screen origin is top-left.
	// Invert Y input axis to account for it.
	player.y += -int(move.Y * speed)

	player.bullets.playerX = player.x
	player.bullets.playerY = player.y
}

func (player *Player) move(x int, y int) {
	const leftBound = 0
	const rightBound = PLAYFIELD_WIDTH
	const topBound = 0
	const bottomBound = PLAYFIELD_HEIGHT

	x = util.ClampInt(leftBound, x, rightBound)
	y = util.ClampInt(topBound, y, bottomBound)

	player.x = x
	player.y = y
}

func NewPlayer() *Player {
	var x, y = INITAL_PLAYER_X, INITAL_PLAYER_Y

	player := Player{
		x: INITAL_PLAYER_X,
		y: INITAL_PLAYER_Y,

		bullets: &Bullets{
			framesPerBullet:  regularBulletFramesPerBullet,
			cooldown:         0,
			image:            playerRegularBullet,
			bulletSize:       regularBulletSize,
			playerX:          x,
			playerY:          y,
			defaultDirection: util.DegToRad(-90),
			defaultDelta:     regularBulletDelta,
		},
	}
	return &player
}
