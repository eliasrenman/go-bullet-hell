package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	movingSlow         bool
	directions         []int8
	shootingMainGun    bool
	shootingSpecialGun bool
	guarding           bool
}

type InputController struct {
	keys []ebiten.Key
}

func (iC InputController) TranslateInput() *Input {
	return iC.translateKeyboardInput()
}
func (iC InputController) translateKeyboardInput() *Input {
	directions := []int8{0, 0}

	for _, key := range iC.keys {
		// Checks if the key pressed is within the bounds of the keybindings
		if val, ok := keyboardBindings[key.String()]; ok {
			// Switch case to check the inputs
			switch val {
			case Left:
				{
					directions[0] += -1
					break
				}
			case Right:
				{
					directions[0] += 1
					break
				}
			case Up:
				{
					directions[1] += -1
					break
				}
			case Down:
				{
					directions[1] += 1
					break
				}
			}

		}
	}

	return &Input{
		movingSlow:         false,
		directions:         directions,
		shootingMainGun:    false,
		shootingSpecialGun: false,
		guarding:           false,
	}
}

var gameInput = InputController{}
