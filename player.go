package main

import (
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
	y             *int16
	x             *int16
	input         *Input
	normalBullets *Bullets
}

func (player *Player) Draw(screen *ebiten.Image) {
	player.drawPlayer(screen)

	// Draw hitbox
	if player.input.movingSlow {
		hitboxOffset := int16(playerSize/2 - hitboxDimension/2)
		x, y := normalizeXCoord(*player.x+hitboxOffset), normalizeYCoord(*player.y+hitboxOffset)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)

		screen.DrawImage(hitbox, op)
	}

	// Draw regular bullets
	player.normalBullets.Draw(screen)
}

func (player *Player) drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x, y := normalizeXCoord(*player.x), normalizeYCoord(*player.y)
	op.GeoM.Translate(x, y)

	if player.input.directions[0] < 0 {
		screen.DrawImage(playerLeftImage, op)

	} else if player.input.directions[0] > 0 {
		screen.DrawImage(playerRightImage, op)
	} else {
		screen.DrawImage(playerForwardImage, op)
	}

}

func (player *Player) Update(input *Input) {
	player.input = input
	player.updateLocation()

	player.move(player.x, player.y)

	player.normalBullets.Update(input)
}

func (player *Player) updateLocation() {

	// Set the apporpriate delta depending on if the slow movement is enabled
	delta := int16(playerFastDelta)
	if player.input.movingSlow {
		delta = int16(playerSlowDelta)
	}

	// Check X direction
	if player.input.directions[0] < 0 {
		*player.x += -delta
	} else if player.input.directions[0] > 0 {
		*player.x += delta
	}
	// Check Y Direction
	if player.input.directions[1] < 0 {
		*player.y += -delta
	} else if player.input.directions[1] > 0 {
		*player.y += delta
	}
}

func (player *Player) move(x *int16, y *int16) {
	playerX := *x
	playerY := *y

	const halfPlayerSize = int16(playerSize / 2)

	const maxX = PLAYFIELD_X_MAX + halfPlayerSize

	// Limit the player from going out of bound on the x axis
	if playerX <= -halfPlayerSize {
		playerX = -halfPlayerSize
	} else if playerX >= maxX {
		playerX = maxX
	}

	const maxY = PLAYFIELD_Y_MAX + halfPlayerSize

	// Limit the player from going out of bound on the y axis
	if playerY <= -halfPlayerSize {
		playerY = -halfPlayerSize
	} else if playerY >= maxY {
		playerY = maxY
	}

	*player.x = playerX
	*player.y = playerY
}

func InitalizePlayer() *Player {
	var x, y int16 = 0, 0
	player := Player{
		x: &x,
		y: &y,
		normalBullets: &Bullets{
			framesPerBullet:  regularBulletFramesPerBullet,
			cooldown:         0,
			image:            playerRegularBullet,
			bulletSize:       regularBulletSize,
			playerX:          &x,
			playerY:          &y,
			defaultDirection: []int8{0, -1},
			defaultDelta:     regularBulletDelta,
		},
	}
	return &player
}
