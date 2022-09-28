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
	yDirection int8
	xDirection int8
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

	if player.xDirection < 0 {
		screen.DrawImage(playerLeftImage, op)

	} else if player.xDirection > 0 {
		screen.DrawImage(playerRightImage, op)
	} else {
		screen.DrawImage(playerForwardImage, op)
	}

}

func (player *Player) Update(input *Input) {
	player.updateDirections(input)

	player.showHitbox = input.movingSlow

	player.Move(player.x+int16(player.xDirection), player.y+int16(player.yDirection))
}

func (player *Player) updateDirections(input *Input) {

	if input.directions[0] == 0 {
		// Make sure to reset directions if no buttons are pressed
		player.xDirection = 0
	}
	if input.directions[1] == 0 {
		// Make sure to reset directions if no y direction buttons are pressed
		player.yDirection = 0
	}

	// Set the apporpriate delta depending on if the slow movement is enabled
	delta := playerFastSpeed
	if input.movingSlow {
		delta = playerSlowSpeed
	}

	// Check X direction
	if input.directions[0] < 0 {
		player.xDirection = -delta
	} else if input.directions[0] > 0 {
		player.xDirection = delta
	}
	// Check Y Direction
	if input.directions[1] < 0 {
		player.yDirection = -delta
	} else if input.directions[1] > 0 {
		player.yDirection = delta
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
		x:          0,
		y:          0,
		xDirection: 0,
		yDirection: 0,
	}
	return &player
}
