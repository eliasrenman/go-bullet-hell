package main

const (
	screenWidth  = 900
	screenHeight = 700
)

const playerSlowSpeed = 2
const playerFastSpeed = 4

const (
	Left  = "Left"
	Right = "Right"
	Down  = "Down"
	Up    = "Up"
)

var keyboardBindings = map[string]string{
	"W": Up,
	"S": Down,
	"A": Left,
	"D": Right,
}
