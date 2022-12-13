package main

import (
	"log"

	"github.com/eliasrenman/go-bullet-hell/constant"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {

	return constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT
}

func main() {

	ebiten.SetWindowSize(constant.SCREEN_WIDTH, constant.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Revenge of the golang")
	if err := ebiten.RunGame(InitalizeGame()); err != nil {
		log.Fatal(err)
	}
}
