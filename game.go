package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is the main game instance
type Game struct {
	player     *Player
	debugger   *Debugger
	background Background
}

var backgroundImage = LoadImage("bg/img1.png", OriginTopLeft)

// NewGame creates a new game instance
func NewGame() *Game {
	player := Spawn(NewPlayer(PlayerStart), CharacterQueue)

	SpawnHealthBar(NewGuiHealthBar(player.Health, PlayfieldSize.X+100, 250, "Player"))
	// Spawn boss
	bossOne := Spawn(NewBossOne(PlayfieldSize.Dot(OriginTop).Plus(Vector{Y: 100})), CharacterQueue)
	SpawnHealthBar(NewGuiHealthBar(bossOne.Health, PlayfieldSize.X+100, 265, "Boss 1"))
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

// Draw is the main draw loop, called every frame
func (game *Game) Draw(screen *ebiten.Image) {
	// Draw background
	game.background.Draw(gameView)

	// Draw game objects
	for obj := range BackgroundObjects {
		obj.Draw(gameView)
	}
	for obj := range CharacterObjects {
		obj.Draw(gameView)
	}

	for obj := range BulletObjects {
		obj.Draw(gameView)
	}

	for obj := range GuiElements {
		obj.Draw(screen)
	}
	drawGameView(screen)
	game.debugger.Draw(screen)
}

var gameView = ebiten.NewImage(int(PlayfieldSize.X), int(PlayfieldSize.Y))

func drawGameView(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	position := PlayfieldOffset

	translateScaleAndRotateImage(&op.GeoM, position, Vector{X: 1, Y: 1}, 0)

	screen.DrawImage(gameView, op)
}

// Update is the main update loop, called every game tick
func (game *Game) Update() error {
	game.background.Update()
	game.debugger.Update()

	for obj := range BackgroundObjects {
		obj.Update()
	}
	for obj := range BulletObjects {
		obj.Update()
	}
	for obj := range CharacterObjects {
		obj.Update()
	}

	for obj := range GuiElements {
		obj.Update()
	}
	SpawnObjects()
	return nil
}
