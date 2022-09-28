package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	playerForwardImage *ebiten.Image = LoadImage("./data/characters/player_forward.png")
	// playerLeftImage    *ebiten.Image = LoadImage("./data/characters/player_left.png")
	// playerRightImage   *ebiten.Image = LoadImage("./data/characters/player_right.png")
)

// //Border checking
// int borderXY = 50;
// if (coordinates.y <= +borderXY)
// 		coordinates.y = borderXY;
// if (coordinates.y >= 590 + borderXY)
// 		coordinates.y = 590 + borderXY;
// if (coordinates.x <= borderXY)
// 		coordinates.x = borderXY;
// if (coordinates.x >= 390 + borderXY)
// 		coordinates.x = 390 + borderXY;

type Player struct {
	y          int8
	x          int8
	yDirection int8
	xDirection int8
}

func (player Player) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.x)+50, float64(player.y)+50)
	screen.DrawImage(playerForwardImage, op)
}

func InitalizePlayer() Player {
	player := Player{
		x:          0,
		y:          0,
		xDirection: 0,
		yDirection: 0,
	}
	return player
}
