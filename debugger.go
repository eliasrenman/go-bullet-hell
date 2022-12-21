package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/eliasrenman/go-bullet-hell/assets"
	"github.com/eliasrenman/go-bullet-hell/constant"
	"github.com/eliasrenman/go-bullet-hell/entity"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Debugger struct {
	Visible         bool
	Game            *Game
	font            font.Face
	fps             float64
	deltaTime       time.Duration
	totalTime       time.Duration
	startTime       time.Time
	graphicsLibrary ebiten.GraphicsLibrary
	cursorPosition  geometry.Point
}

func NewDebugger(game *Game) *Debugger {

	return &Debugger{
		Visible: false,
		Game:    game,

		font:      assets.LoadFont("fonts/FiraCode.ttf", opentype.FaceOptions{}),
		fps:       0,
		startTime: time.Now(),
	}
}

func (debugger *Debugger) Update() error {
	// Toggle debug mode
	if input.ButtonDebug.GetPressed(0) {
		debugger.Visible = !debugger.Visible

		if debugger.Visible {
			debugger.Start()
		} else {
			debugger.Stop()
		}
	}

	debugger.fps = ebiten.ActualFPS()

	totalTime := time.Since(debugger.startTime)
	debugger.deltaTime = totalTime - debugger.totalTime
	debugger.totalTime = totalTime

	var debugInfo ebiten.DebugInfo
	ebiten.ReadDebugInfo(&debugInfo)
	debugger.graphicsLibrary = debugInfo.GraphicsLibrary

	x, y := ebiten.CursorPosition()
	debugger.cursorPosition = geometry.Point{X: float64(x), Y: float64(y)}

	return nil
}

var crosshairImage = assets.LoadImage("crosshair.png", assets.OriginCenter)

func (debugger *Debugger) Start() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	fmt.Println("Debug mode enabled")
}

func (debugger *Debugger) Stop() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	fmt.Println("Debug mode disabled")
}

func (debugger *Debugger) Draw(screen *ebiten.Image) {
	if !debugger.Visible {
		return
	}

	// Draw debug text
	debugText := fmt.Sprintf(`
FPS: %.2f
Playing since: %v
Total play time: %v
Frame time: %v,
Total Game Objects: %v
Running on %v`,
		debugger.fps,
		debugger.startTime.Format("January 2 15:04:05"),
		debugger.totalTime.Truncate(time.Second),
		debugger.deltaTime.Truncate(time.Millisecond/100),
		len(entity.GameObjects),
		getGraphicsLibraryName(int(debugger.graphicsLibrary)))

	text.Draw(screen, debugText, debugger.font, constant.PLAYFIELD_WIDTH+100, 5, color.White)

	// Draw crosshair
	crosshairImage.Draw(screen, debugger.cursorPosition, geometry.Size{Width: 1, Height: 1}, 0)
	crosshairText := fmt.Sprintf(
		"%v, %v",
		debugger.cursorPosition.X-float64(constant.PLAYFIELD_OFFSET),
		debugger.cursorPosition.Y-float64(constant.PLAYFIELD_OFFSET),
	)
	text.Draw(screen, crosshairText, debugger.font, int(debugger.cursorPosition.X)+4, int(debugger.cursorPosition.Y)-4, color.White)
}

func getGraphicsLibraryName(gl int) string {
	switch gl {
	case 1:
		return "OpenGL"
	case 2:
		return "Metal"
	case 3:
		return "DirectX"
	default:
		return "Unknown"
	}
}
