package main

var (
	ScreenSize      = Vector{X: 900, Y: 700}
	PlayfieldOffset = Vector{X: 50, Y: 50}
	PlayfieldSize   = Vector{X: 500, Y: 600}
	PlayerStart     = PlayfieldSize.Dot(OriginBottom).Plus(Vector{Y: -100})
	BackgroundSpeed = 3.0
	PlayerSpeed     = 4.0
	PlayerSpeedSlow = 2.0
)
