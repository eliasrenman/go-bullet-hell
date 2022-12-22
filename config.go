package main

const (
	SCREEN_WIDTH              int     = 900
	SCREEN_HEIGHT             int     = 700
	PLAYFIELD_OFFSET          uint8   = 50
	PLAYFIELD_WIDTH                   = 500
	PLAYFIELD_HEIGHT                  = 600
	INITIAL_PLAYER_X                  = PLAYFIELD_WIDTH / 2
	INITIAL_PLAYER_Y                  = (PLAYFIELD_HEIGHT / 5) * 4
	CONTROLLER_DEADZONE       float32 = 0.2
	STANDARD_BACKGROUND_SPEED         = 3
)

// Player speeds
const PLAYER_SPEED_SLOW float64 = 2
const PLAYER_SPEED float64 = 4

// Player hitbox
const hitboxDimension = 8

const (
	LEFT        = "Left"
	RIGHT       = "Right"
	DOWN        = "Down"
	UP          = "Up"
	SLOW        = "Slow"
	REGULAR_GUN = "REGULAR_GUN"
)

var KeyboardBindings = map[string]string{
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

var ControllerBindings = map[string]string{
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
