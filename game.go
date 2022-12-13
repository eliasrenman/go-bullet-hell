package main

import (
	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/entity"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player     *entity.Player
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
			Velocity: geometry.Up.ScaledBy(3),
		},
	}

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
}

var gameView = ebiten.NewImage(PLAYFIELD_WIDTH, PLAYFIELD_HEIGHT)

func DrawGameView(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	position := geometry.Point{X: float64(PLAYFIELD_OFFSET), Y: float64(PLAYFIELD_OFFSET)}

	assets.TranslateScaleAndRotateImage(&op.GeoM, position, geometry.Size{Width: 1, Height: 1}, 0)

	screen.DrawImage(gameView, op)
}

func (game *Game) Update() error {
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
