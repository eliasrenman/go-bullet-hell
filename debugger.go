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

// Debugger is a debug overlay that displays information about the game
type Debugger struct {
	Visible         bool
	Game            *Game
	font            font.Face
	fps             float64
	deltaTime       time.Duration
	totalTime       time.Duration
	startTime       time.Time
	graphicsLibrary ebiten.GraphicsLibrary
	cursorPosition  Vector
}

// NewDebugger creates a new debugger given a game instance
func NewDebugger(game *Game) *Debugger {

	return &Debugger{
		Visible: false,
		Game:    game,

		font:      LoadFont("fonts/FiraCode.ttf", opentype.FaceOptions{}),
		fps:       0,
		startTime: time.Now(),
	}
}

// Update updates the debugger information
func (debugger *Debugger) Update() error {
	// Toggle debug mode
	if ButtonDebug.GetPressed(0) {
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
	debugger.cursorPosition = Vector{X: float64(x), Y: float64(y)}

	return nil
}

var crosshairImage = LoadImage("crosshair.png", OriginCenter)

// Start is called when the debugger is enabled
func (debugger *Debugger) Start() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	fmt.Println("Debug mode enabled")
}

// Stop is called when the debugger is disabled
func (debugger *Debugger) Stop() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	fmt.Println("Debug mode disabled")
}

// Draw draws the debugger onto the screen
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
		len(GameObjects),
		getGraphicsLibraryName(int(debugger.graphicsLibrary)))

	text.Draw(screen, debugText, debugger.font, int(PlayfieldSize.X)+100, 5, color.White)

	// Draw crosshair
	crosshairImage.Draw(screen, debugger.cursorPosition, Vector{X: 1, Y: 1}, 0)
	crosshairText := fmt.Sprintf(
		"%v, %v",
		debugger.cursorPosition.X-PlayfieldOffset.X,
		debugger.cursorPosition.Y-PlayfieldOffset.Y,
	)
	text.Draw(screen, crosshairText, debugger.font, int(debugger.cursorPosition.X)+4, int(debugger.cursorPosition.Y)-4, color.White)
}

func getGraphicsLibraryName(gl int) string {
	switch gl {
	case 1:
		return "OpenGL"
	case 2:
		return "DirectX"
	case 3:
		return "Metal"
	default:
		return "Unknown"
	}
}
