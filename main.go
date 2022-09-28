package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {

	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {

	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Revenge of the golang")
	if err := ebiten.RunGame(InitalizeGame()); err != nil {
		log.Fatal(err)
	}
}
