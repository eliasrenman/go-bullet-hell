package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player     *Player
	debugger   *Debugger
	background Background
}

var backgroundImage = LoadImage("bg/img1.png", OriginTopLeft)

func InitalizeGame() *Game {
	player := Spawn(NewPlayer(PlayerStart))

	// Spawn boss
	Spawn(NewBossOne(PlayfieldSize.Dot(OriginTop).Plus(Vector{Y: 100})))

	game := Game{
		player: player,
		background: Background{
			Image:    backgroundImage,
			Velocity: Up.ScaledBy(BackgroundSpeed),
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
	EachGameObject(func(obj GameObject) {
		obj.Draw(gameView)
	})
	DrawGameView(screen)
	game.debugger.Draw(screen)
}

var gameView = ebiten.NewImage(int(PlayfieldSize.X), int(PlayfieldSize.Y))

func DrawGameView(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	position := PlayfieldOffset

	TranslateScaleAndRotateImage(&op.GeoM, position, Vector{X: 1, Y: 1}, 0)

	screen.DrawImage(gameView, op)
}

func (game *Game) Update() error {
	updateGameBackgroundSpeed(game)
	game.background.Update()
	game.debugger.Update()

	EachGameObject(func(obj GameObject) {
		obj.Update()
	})

	SpawnGameObjects()
	return nil
}

func updateGameBackgroundSpeed(game *Game) {
	if game.player.Velocity.Y != 0 {
		offsetVelocity := game.player.Velocity.Y*-0.5 + PlayerSpeed/2
		game.background.Velocity = Up.ScaledBy(offsetVelocity + 1)
	} else {
		game.background.Velocity = Up.ScaledBy(BackgroundSpeed)
	}
}
