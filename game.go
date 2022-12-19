package main

import (
	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/constant"
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
		X: constant.INITIAL_PLAYER_X,
		Y: constant.INITIAL_PLAYER_Y,
	}))

	// Spawn boss
	entity.Spawn(entity.NewBossOne(geometry.Point{
		X: constant.INITIAL_PLAYER_X,
		Y: 200,
	}))

	game := Game{
		player: player,
		background: assets.Background{
			Image:    backgroundImage,
			Velocity: geometry.Up.ScaledBy(constant.STANDARD_BACKGROUND_SPEED),
		},
		debugger: nil,
	}

	game.debugger = NewDebugger(&game)

	return &game
}

func (game *Game) Draw(screen *ebiten.Image) {
	// Draw background
	game.background.Draw(gameView)

	// Draw game objects
	for obj := range entity.GameObjects {
		obj.Draw(gameView)
	}
	DrawGameView(screen)
	game.debugger.Draw(screen)
}

var gameView = ebiten.NewImage(constant.PLAYFIELD_WIDTH, constant.PLAYFIELD_HEIGHT)

func DrawGameView(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	position := geometry.Point{X: float64(constant.PLAYFIELD_OFFSET), Y: float64(constant.PLAYFIELD_OFFSET)}

	assets.TranslateScaleAndRotateImage(&op.GeoM, position, geometry.Size{Width: 1, Height: 1}, 0)

	screen.DrawImage(gameView, op)
}

func (game *Game) Update() error {
	updateGameBackgroundSpeed(game)
	game.background.Update()
	game.debugger.Update()
	for obj := range entity.GameObjects {
		obj.Update()
	}
	return nil
}

func updateGameBackgroundSpeed(game *Game) {
	if game.player.Velocity.Y != 0 {
		offsetVelocity := ((game.player.Velocity.Y * -1) + constant.PLAYER_SPEED) / 2
		game.background.Velocity = geometry.Up.ScaledBy(offsetVelocity + 1)
	} else {
		game.background.Velocity = geometry.Up.ScaledBy(constant.STANDARD_BACKGROUND_SPEED)
	}
}
