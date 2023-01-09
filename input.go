package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type axis struct {
	axis        int
	keyPositive ebiten.Key
	keyNegative ebiten.Key
	deadzone    float64
}

type button struct {
	button ebiten.GamepadButton
	key    ebiten.Key

	value bool
}

// TODO: Convert these to a struct and import via a json or toml marshaler
var (
	AxisHorizontal = axis{
		axis:        0,
		keyPositive: ebiten.KeyD,
		keyNegative: ebiten.KeyA,
	}
	AxisVertical = axis{
		axis:        1,
		keyPositive: ebiten.KeyW,
		keyNegative: ebiten.KeyS,
	}
	ButtonSlow = button{
		button: ebiten.GamepadButton6,
		key:    ebiten.KeyShift,
	}
	ButtonShoot = button{
		button: ebiten.GamepadButton0,
		key:    ebiten.KeySpace,
	}
	ButtonDebug = button{
		key: ebiten.KeyF3,
	}
	ButtonDebugHitbox = button{
		key: ebiten.KeyF2,
	}
)

func (axis *axis) Get(gp ebiten.GamepadID) float64 {
	keyValue := float64(0)
	if ebiten.IsKeyPressed(axis.keyPositive) {
		keyValue++
	}
	if ebiten.IsKeyPressed(axis.keyNegative) {
		keyValue--
	}

	value := ClampFloat(-1, ebiten.GamepadAxisValue(gp, axis.axis)+keyValue, 1)
	if value > -axis.deadzone && value < axis.deadzone {
		return 0
	}
	return value
}

func (button *button) Get(gp ebiten.GamepadID) bool {
	button.value = ebiten.IsGamepadButtonPressed(gp, button.button) || ebiten.IsKeyPressed(button.key)
	return button.value
}

// GetPressed returns true if the button was pressed this frame
func (button *button) GetPressed(gp ebiten.GamepadID) bool {
	pv := button.value
	return button.Get(gp) && !pv
}
