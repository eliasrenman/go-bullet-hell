package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {

	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(InitalizeGame()); err != nil {
		log.Fatal(err)
	}
}
