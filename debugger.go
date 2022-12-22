package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Debugger struct {
	Visible     bool
	Game        *Game
	font        font.Face
	fps         float64
	deltaTime   time.Duration
	totalTime   time.Duration
	startTime   time.Time
	graphicslib ebiten.GraphicsLibrary
}

func NewDebugger(game *Game) *Debugger {

	return &Debugger{
		Visible: false,
		Game:    game,

		font:      LoadFont("fonts/FiraCode.ttf", opentype.FaceOptions{}),
		fps:       0,
		startTime: time.Now(),
	}
}

func (debugger *Debugger) Update() error {
	// Toggle debug mode
	if ButtonDebug.GetPressed(0) {
		debugger.Visible = !debugger.Visible
		fmt.Printf("Debug mode: %v\n", debugger.Visible)
	}

	debugger.fps = ebiten.ActualFPS()

	totalTime := time.Since(debugger.startTime)
	debugger.deltaTime = totalTime - debugger.totalTime
	debugger.totalTime = totalTime

	var debugInfo ebiten.DebugInfo
	ebiten.ReadDebugInfo(&debugInfo)
	debugger.graphicslib = debugInfo.GraphicsLibrary

	return nil
}

func init() {
}

func (debugger *Debugger) Draw(screen *ebiten.Image) {
	if !debugger.Visible {
		return
	}

	debugText := fmt.Sprintf(`
FPS: %.2f
Playing since: %v
Total play time: %v
Frame time: %v,
Total Game Objects: %v`,
		debugger.fps,
		debugger.startTime.Format("January 2 15:04:05"),
		debugger.totalTime.Truncate(time.Second),
		debugger.deltaTime.Truncate(time.Millisecond/100),
		len(GameObjects))

	text.Draw(screen, debugText, debugger.font, int(PlayfieldSize.X)+100, 5, color.White)
}
