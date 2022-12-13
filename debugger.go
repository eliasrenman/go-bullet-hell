package main

import (
	"fmt"
	"image/color"

	"github.com/eliasrenman/go-bullet-hell/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Debugger struct {
	Visible     bool
	Game        *Game
	fps         float64
	deltaTime   float64
	totalTime   float64
	graphicslib ebiten.GraphicsLibrary
}

func NewDebugger(game *Game) *Debugger {
	return &Debugger{
		Visible: false,
		Game:    game,

		fps: 0,
	}
}

var lastButtonDebugState bool

func (debugger *Debugger) Update() error {
	// Toggle debug mode
	if input.ButtonDebug.Get(0) && !lastButtonDebugState {
		debugger.Visible = true
		fmt.Println("Debug mode toggled: ", debugger.Visible)
		lastButtonDebugState = true
	} else {
		lastButtonDebugState = false
	}

	debugger.fps = ebiten.ActualFPS()
	debugger.deltaTime = 1 / ebiten.ActualFPS()

	var debugInfo ebiten.DebugInfo
	ebiten.ReadDebugInfo(&debugInfo)
	debugger.graphicslib = debugInfo.GraphicsLibrary

	return nil
}

func (debugger *Debugger) Draw(screen *ebiten.Image) {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	font, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	text.Draw(screen, fmt.Sprintf("FPS: %2f", debugger.fps), font, 0, 0, color.White)
}
