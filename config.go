package main

const (
	SCREEN_WIDTH     int   = 900
	SCREEN_HEIGHT    int   = 700
	PLAYFIELD_OFFSET uint8 = 50
	PLAYFIELD_X_MAX        = 350
	PLAYFIELD_Y_MAX        = 550
)

const playerSlowSpeed int8 = 2
const playerFastSpeed int8 = 4
const hitboxDimension = 8
const regularBulletSize = 8

const (
	LEFT  = "Left"
	RIGHT = "Right"
	DOWN  = "Down"
	UP    = "Up"
	SLOW  = "Slow"
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
}
