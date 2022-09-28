package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player *Player
}

func InitalizeGame() *Game {
	game := Game{
		player: InitalizePlayer(),
	}
	return &game
}

func (game Game) Draw(screen *ebiten.Image) {
	// Draw background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(backgroundImage, op)

	// op.GeoM.Translate(150, 150)
	// screen.DrawImage(playerBullet, op)

	// Draw characters
	game.player.Draw(screen)
	// Draw bullets
}
func (game Game) Update() error {
	//Append the keys
	gameInput.keys = inpututil.AppendPressedKeys(gameInput.keys[:0])

	// get gameInput
	input := gameInput.TranslateInput()
	// update characters
	game.player.Update(input)
	// update bullets
	// Return no error
	return nil
}
