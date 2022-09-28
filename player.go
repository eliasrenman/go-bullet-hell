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
	op.GeoM.Translate(float64(player.x)+50, float64(player.y)+50)

	if player.xDirection < 0 {
		screen.DrawImage(playerLeftImage, op)

	} else if player.xDirection > 0 {
		screen.DrawImage(playerRightImage, op)
	} else {
		screen.DrawImage(playerForwardImage, op)
	}
}

func (player *Player) Update(input *Input) {
	x, y := GetNewCoordinates(player, input)

	player.Move(x, y)
}

func GetNewCoordinates(player *Player, input *Input) (int16, int16) {
	newPosX := player.x
	newPosY := player.y

	if input.directionLeft {
		if input.movingSlow {
			newPosX -= playerSlowSpeed
		} else {
			newPosX -= playerFastSpeed
		}
	} else if input.directionRight {
		if input.movingSlow {
			newPosX += playerSlowSpeed
		} else {
			newPosX += playerFastSpeed
		}
	}
	if input.directionUp {
		if input.movingSlow {
			newPosY -= playerSlowSpeed
		} else {
			newPosY -= playerFastSpeed
		}
	} else if input.directionDown {
		if input.movingSlow {
			newPosY += playerSlowSpeed
		} else {
			newPosY += playerFastSpeed
		}
	}
	return newPosX, newPosY
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
