package main

import "github.com/hajimehoshi/ebiten/v2"

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

	// Draw characters
	game.player.Draw(screen)
	// Draw bullets
}
func (game Game) Update() error {

	// update characters
	game.player.Update()
	// update bullets
	// Return no error
	return nil
}
