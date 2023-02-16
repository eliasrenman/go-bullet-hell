// The main package
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Layout returns the window layout, for ebiten
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return int(ScreenSize.X), int(ScreenSize.Y)
}

var game Game

func GetGame() *Game {
	return &game
}

func main() {
	game = *NewGame()
	ebiten.SetWindowSize(int(ScreenSize.X), int(ScreenSize.Y))
	ebiten.SetWindowTitle("Revenge of the golang")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
