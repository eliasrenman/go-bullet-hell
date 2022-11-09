package input

import (
	"github.com/eliasrenman/go-bullet-hell/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Axis struct {
	Axis        int
	KeyPositive ebiten.Key
	KeyNegative ebiten.Key
	Deadzone    float64
}

type Button struct {
	Button ebiten.GamepadButton
	Key    ebiten.Key
}

func (axis *Axis) Get(gp ebiten.GamepadID) float64 {
	keyValue := float64(0)
	if ebiten.IsKeyPressed(axis.KeyPositive) {
		keyValue++
	}
	if ebiten.IsKeyPressed(axis.KeyNegative) {
		keyValue--
	}

	value := util.ClampFloat(-1, ebiten.GamepadAxisValue(gp, axis.Axis)+keyValue, 1)
	if value > -axis.Deadzone && value < axis.Deadzone {
		return 0
	}
	return value
}

func (button *Button) Get(gp ebiten.GamepadID) bool {
	return ebiten.IsGamepadButtonPressed(gp, button.Button) || ebiten.IsKeyPressed(button.Key)
}
