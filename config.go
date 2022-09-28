package main

const (
	SCREEN_WIDTH     uint16 = 900
	SCREEN_HEIGHT    uint16 = 700
	PLAYFIELD_OFFSET uint8  = 50
)

const playerSlowSpeed int8 = 2
const playerFastSpeed int8 = 4

const (
	LEFT  = "Left"
	RIGHT = "Right"
	DOWN  = "Down"
	UP    = "Up"
	SLOW  = "Slow"
)

var keyboardBindings = map[string]string{
	"W":     UP,
	"S":     DOWN,
	"A":     LEFT,
	"D":     RIGHT,
	"Shift": SLOW,
}
