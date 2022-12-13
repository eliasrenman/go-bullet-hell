package main

import (
	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/entity"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player     *entity.Player
	debugger   *Debugger
	background assets.Background
}

var backgroundImage = assets.LoadImage("bg/img1.png", assets.OriginTopLeft)

func InitalizeGame() *Game {
	player := entity.Spawn(entity.NewPlayer(geometry.Point{
		X: INITIAL_PLAYER_X,
		Y: INITIAL_PLAYER_Y,
	}))

	game := Game{
		player: player,
		background: assets.Background{
			Image:    backgroundImage,
			Velocity: geometry.Up.ScaledBy(15),
		},
		debugger: nil,
	}

	game.debugger = NewDebugger(&game)

	return &game
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.background.Draw(screen)
	game.debugger.Draw(screen)

	// Draw game objects
	for obj := range entity.GameObjects {
		obj.Draw(screen)
	}
}

func (game *Game) Update() error {
	game.debugger.Update()

	for obj := range entity.GameObjects {
		obj.Update()
	}
	return nil
}
