package main

import (
	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/entity"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *entity.Player
}

func InitalizeGame() *Game {
	player := entity.Spawn(&entity.Player{})
	player.Position.X = INITIAL_PLAYER_X
	player.Position.Y = INITIAL_PLAYER_Y

	game := Game{player: player}
	return &game
}

var backgroundImage = assets.LoadImage("bg/playfield.png", assets.OriginTopLeft)

func (game Game) Draw(screen *ebiten.Image) {
	backgroundImage.Draw(screen, geometry.Point{X: 0, Y: 0}, geometry.Size{Width: 1, Height: 1}, 0)

	// Draw game objects
	for obj := range entity.GameObjects {
		obj.Draw(screen)
	}
}
func (game Game) Update() error {
	for obj := range entity.GameObjects {
		obj.Update()
	}
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
