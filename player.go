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
	y          int16
	x          int16
	input      *Input
	showHitbox bool
}

func (player *Player) Draw(screen *ebiten.Image) {
	player.drawPlayer(screen)

	// Draw hitbox
	if player.showHitbox {
		x, y := float64(player.x)+float64(PLAYFIELD_OFFSET)+(25-4), float64(player.y)+float64(PLAYFIELD_OFFSET)+(25-4)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)

		screen.DrawImage(hitbox, op)
	}
}

func (player *Player) drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x, y := float64(player.x)+float64(PLAYFIELD_OFFSET), float64(player.y)+float64(PLAYFIELD_OFFSET)
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
	player.updateDirections()

	player.showHitbox = input.movingSlow

	player.Move(player.x, player.y)
}

func (player *Player) updateDirections() {

	// Set the apporpriate delta depending on if the slow movement is enabled
	delta := int16(playerFastSpeed)
	if player.input.movingSlow {
		delta = int16(playerSlowSpeed)
	}

	// Check X direction
	if player.input.directions[0] < 0 {
		player.x += -delta
	} else if player.input.directions[0] > 0 {
		player.x += delta
	}
	// Check Y Direction
	if player.input.directions[1] < 0 {
		player.y += -delta
	} else if player.input.directions[1] > 0 {
		player.y += delta
	}
}

func (player *Player) Move(x int16, y int16) {

	player.x = x
	player.y = y

	const halfPlayerSize = int16(playerSize / 2)

	const maxX = PLAYFIELD_X_MAX + halfPlayerSize

	// Limit the player from going out of bound on the x axis
	if player.x <= -halfPlayerSize {
		player.x = -halfPlayerSize
	} else if player.x >= maxX {
		player.x = maxX
	}

	const maxY = PLAYFIELD_Y_MAX + halfPlayerSize

	// Limit the player from going out of bound on the y axis
	if player.y <= -halfPlayerSize {
		player.y = -halfPlayerSize
	} else if player.y >= maxY {
		player.y = maxY
	}
}

func InitalizePlayer() *Player {
	player := Player{
		x: 0,
		y: 0,
	}
	return &player
}
