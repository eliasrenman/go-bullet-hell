package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	playerForwardImage *ebiten.Image = LoadImage("./data/characters/player_forward.png")
	playerLeftImage    *ebiten.Image = LoadImage("./data/characters/player_left.png")
	playerRightImage   *ebiten.Image = LoadImage("./data/characters/player_right.png")
)

type Player struct {
	y          int16
	x          int16
	yDirection int8
	xDirection int8
}

func (player *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.x)+float64(PLAYFIELD_OFFSET), float64(player.y)+float64(PLAYFIELD_OFFSET))

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

	player.Move(player.x+int16(player.xDirection), player.y+int16(player.yDirection))
}

func (player *Player) updateDirections(input *Input) {

	if input.directions[1] == 0 {
		// Make sure to reset directions if no buttons are pressed
		player.yDirection = 0
	}
	if input.directions[0] == 0 {
		// Make sure to reset directions if no buttons are pressed
		player.xDirection = 0
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

	// Make sure to reset directions if no buttons are pressed

}

func (player *Player) Move(x int16, y int16) {

	player.x = x
	player.y = y

	if player.y <= -25 {
		player.y = -25
	} else if player.y >= 575 {
		player.y = 575
	}

	if player.x <= -25 {
		player.x = -25
	} else if player.x >= 375 {
		player.x = 375
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
