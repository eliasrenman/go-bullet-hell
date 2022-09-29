package main

const (
	SCREEN_WIDTH        int     = 900
	SCREEN_HEIGHT       int     = 700
	PLAYFIELD_OFFSET    uint8   = 50
	PLAYFIELD_X_MAX             = 350
	PLAYFIELD_Y_MAX             = 550
	INITAL_PLAYER_X             = (PLAYFIELD_X_MAX / 2) - playerSize/4
	INITAL_PLAYER_Y             = (PLAYFIELD_Y_MAX / 5) * 4
	CONTROLLER_DEADZONE float32 = 0.2
)

// Player speeds
const playerSlowDelta int8 = 2
const playerFastDelta int8 = 4

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
