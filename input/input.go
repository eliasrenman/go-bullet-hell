package input

import (
	"github.com/eliasrenman/go-bullet-hell/util"
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
)

func (axis *axis) Get(gp ebiten.GamepadID) float64 {
	keyValue := float64(0)
	if ebiten.IsKeyPressed(axis.keyPositive) {
		keyValue++
	}
	if ebiten.IsKeyPressed(axis.keyNegative) {
		keyValue--
	}

	value := util.ClampFloat(-1, ebiten.GamepadAxisValue(gp, axis.axis)+keyValue, 1)
	if value > -axis.deadzone && value < axis.deadzone {
		return 0
	}
	return value
}

func (button *button) Get(gp ebiten.GamepadID) bool {
	return ebiten.IsGamepadButtonPressed(gp, button.button) || ebiten.IsKeyPressed(button.key)
}
