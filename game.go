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
	player := Spawn(NewPlayer(Point{
		X: INITIAL_PLAYER_X,
		Y: INITIAL_PLAYER_Y,
	}))

	// Spawn boss
	Spawn(NewBossOne(Point{
		X: INITIAL_PLAYER_X,
		Y: 200,
	}))

	game := Game{
		player: player,
		background: Background{
			Image:    backgroundImage,
			Velocity: Up.ScaledBy(STANDARD_BACKGROUND_SPEED),
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

var gameView = ebiten.NewImage(PLAYFIELD_WIDTH, PLAYFIELD_HEIGHT)

func DrawGameView(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	position := Point{X: float64(PLAYFIELD_OFFSET), Y: float64(PLAYFIELD_OFFSET)}

	TranslateScaleAndRotateImage(&op.GeoM, position, Size{Width: 1, Height: 1}, 0)

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
		offsetVelocity := ((game.player.Velocity.Y * -1) + PLAYER_SPEED) / 2
		game.background.Velocity = Up.ScaledBy(offsetVelocity + 1)
	} else {
		game.background.Velocity = Up.ScaledBy(STANDARD_BACKGROUND_SPEED)
	}
}
