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
		player: NewPlayer(),
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
	//Append the keys
	gameInput.keys = inpututil.AppendPressedKeys(gameInput.keys[:0])
	gameInput.Update()
	// get gameInput
	input := gameInput.TranslateInput()
	// update characters
	game.player.Update(input)
	// update bullets
	// Return no error
	return nil
}

func normalizeXCoord(x int) float64 {
	return float64(x) + float64(PLAYFIELD_OFFSET)
}

func normalizeYCoord(y int) float64 {
	return float64(y) + float64(PLAYFIELD_OFFSET)
}

func normalizeCoords(x int, y int) (float64, float64) {
	return normalizeXCoord(x), normalizeYCoord(y)
}
