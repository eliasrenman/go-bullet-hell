package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is the main game instance
type Game struct {
	player     *Player
	debugger   *Debugger
	background Background
	gameOver   bool
}

var backgroundImage = LoadImage("bg/img1.png", OriginTopLeft)

var testLevel = Level{
	Events: []Event{
		{
			Type: "spawn",
			Options: EventOptions{
				Enemies: []EnemySpawnOptions{
					{
						Enemy:    "testenemy",
						Position: Vector{X: 100, Y: -20},
					},
					{
						Enemy:    "testenemy",
						Position: Vector{X: 300, Y: -20},
					},
				},
			},
			Cooldown: 5,
		},
		{
			Type: "spawn",
			Options: EventOptions{
				Enemies: []EnemySpawnOptions{
					{
						Enemy:    "testenemy",
						Position: Vector{X: 50, Y: -20},
					},
					{
						Enemy:    "testenemy",
						Position: Vector{X: 225, Y: -20},
					},
					{
						Enemy:    "testenemy",
						Position: Vector{X: 400, Y: -20},
					},
				},
			},
			Cooldown: 5,
		},
		{
			Type: "spawn",
			Options: EventOptions{
				Enemies: []EnemySpawnOptions{
					{
						Enemy:    "testenemy",
						Position: Vector{X: 100, Y: -20},
					},
					{
						Enemy:    "testenemy",
						Position: Vector{X: 300, Y: -20},
					},
				},
			},
			Cooldown: 5,
		},
	},
}

// NewGame creates a new game instance
func NewGame() *Game {
	player := Spawn(NewPlayer(PlayerStart), 1)

	SpawnHealthBar(NewGuiHealthBar(player.Health, PlayfieldSize.X+100, 250, "Player"))
	game := Game{
		player: player,
		background: Background{
			Image:    backgroundImage,
			Velocity: Up.ScaledBy(BackgroundSpeed),
		},
		debugger: nil,
	}

	game.debugger = NewDebugger(&game)

	go testLevel.Start()

	return &game
}

// Draw is the main draw loop, called every frame
func (game *Game) Draw(screen *ebiten.Image) {
	// Draw background
	game.background.Draw(gameView)
	EachGameObject(func(obj GameObject, layer int) {
		obj.Draw(gameView)
	})

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
	if game.gameOver {
		return nil
	}

	game.background.Update()
	game.debugger.Update()

	EachGameObject(func(obj GameObject, layer int) {
		obj.Update(game)
	})

	for obj := range GuiElements {
		obj.Update()
	}
	SpawnObjects()
	return nil
}

func (game *Game) GameOver() {
	game.gameOver = true
}
