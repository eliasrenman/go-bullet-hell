package main

var (
	// ScreenSize is the size of the game screen
	ScreenSize = Vector{X: 900, Y: 700}
	// PlayfieldOffset is the offset of the playfield from the top left corner of the screen
	PlayfieldOffset = Vector{X: 50, Y: 50}
	// PlayfieldSize is the size of the playfield
	PlayfieldSize = Vector{X: 500, Y: 600}
	// PlayerStart is the starting position of the player
	PlayerStart = PlayfieldSize.Dot(OriginBottom).Plus(Vector{Y: -100})
	// BackgroundSpeed is the default speed at which the background moves
	BackgroundSpeed = 3.0
	// PlayerSpeed is the default speed at which the player moves
	PlayerSpeed = 4.0
	// PlayerSpeedSlow is the speed at which the player moves when the slow key is pressed
	PlayerSpeedSlow = 2.0
)
