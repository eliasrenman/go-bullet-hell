package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	movingSlow         bool
	directionLeft      bool
	directionRight     bool
	directionUp        bool
	directionDown      bool
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
	var directionLeft, directionRight, directionUp, directionDown = false, false, false, false

	for _, key := range iC.keys {
		// Checks if the key pressed is within the bounds of the keybindings
		if val, ok := keyboardBindings[key.String()]; ok {
			// Switch case to check the inputs
			switch val {
			case Left:
				{
					directionLeft = true
					break
				}
			case Right:
				{
					directionRight = true
					break
				}
			case Up:
				{
					directionUp = true
					break
				}
			case Down:
				{
					directionDown = true
					break
				}
			}

		}
	}

	// Make Sure to cancel out the input if both up and down is pressed
	if directionUp && directionDown {
		directionDown = false
		directionUp = false
	}
	// Make Sure to cancel out the input if both left and right is pressed
	if directionUp && directionDown {
		directionLeft = false
		directionRight = false
	}

	return &Input{
		movingSlow:         false,
		directionLeft:      directionLeft,
		directionRight:     directionRight,
		directionUp:        directionUp,
		directionDown:      directionDown,
		shootingMainGun:    false,
		shootingSpecialGun: false,
		guarding:           false,
	}
}

var gameInput = InputController{}
