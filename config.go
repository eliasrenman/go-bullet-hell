package main

import (
	"github.com/eliasrenman/go-bullet-hell/input"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH        int     = 900
	SCREEN_HEIGHT       int     = 700
	PLAYFIELD_OFFSET    uint8   = 50
	PLAYFIELD_WIDTH             = 350
	PLAYFIELD_HEIGHT            = 550
	INITAL_PLAYER_X             = (PLAYFIELD_WIDTH / 2) - playerSize/4
	INITAL_PLAYER_Y             = (PLAYFIELD_HEIGHT / 5) * 4
	CONTROLLER_DEADZONE float32 = 0.2
)

// Player speeds
const PLAYER_SPEED_SLOW float64 = 2
const PLAYER_SPEED float64 = 4

// Player hitbox
const hitboxDimension = 8

// Regular player Bullets
const regularBulletSize = 8
const regularBulletFramesPerBullet = 5
const regularBulletDelta = 6

const (
	LEFT        = "Left"
	RIGHT       = "Right"
	DOWN        = "Down"
	UP          = "Up"
	SLOW        = "Slow"
	REGULAR_GUN = "REGULAR_GUN"
)

// TODO: Convert these to a struct and import via a json or toml marshaler
var (
	AXIS_HORIZONTAL = input.Axis{
		Axis:        0,
		KeyPositive: ebiten.KeyD,
		KeyNegative: ebiten.KeyA,
	}
	AXIS_VERTICAL = input.Axis{
		Axis:        1,
		KeyPositive: ebiten.KeyW,
		KeyNegative: ebiten.KeyS,
	}
	BUTTON_SLOW = input.Button{
		Button: ebiten.GamepadButton6,
		Key:    ebiten.KeyShift,
	}
)

var keyboardBindings = map[string]string{
	"W":          UP,
	"S":          DOWN,
	"A":          LEFT,
	"D":          RIGHT,
	"ArrowUp":    UP,
	"ArrowDown":  DOWN,
	"ArrowLeft":  LEFT,
	"ArrowRight": RIGHT,
	"Shift":      SLOW,
	"Z":          REGULAR_GUN,
}

var controllerBindings = map[string]string{
	// "W":          UP,
	// "S":          DOWN,
	// "A":          LEFT,
	// "D":          RIGHT,
	"15": UP,
	"17": DOWN,
	"18": LEFT,
	"16": RIGHT,
	"6":  SLOW,
	"7":  SLOW,
	"8":  SLOW,
	"9":  SLOW,
	"0":  REGULAR_GUN,
}
