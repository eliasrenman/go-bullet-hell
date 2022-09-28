package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	movingSlow         bool
	directions         []int8
	shootingRegularGun bool
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
	movingSlow := false
	shootingRegularGun := false
	for _, key := range iC.keys {
		// fmt.Println(key)
		// Checks if the key pressed is within the bounds of the keybindings
		if val, ok := keyboardBindings[key.String()]; ok {
			// Switch case to check the inputs
			switch val {
			case LEFT:
				{
					directions[0] += -1
					break
				}
			case RIGHT:
				{
					directions[0] += 1
					break
				}
			case UP:
				{
					directions[1] += -1
					break
				}
			case DOWN:
				{
					directions[1] += 1
					break
				}
			case SLOW:
				{
					movingSlow = true
					break
				}
			case REGULAR_GUN:
				{
					shootingRegularGun = true
					break
				}
			}

		}
	}

	return &Input{
		movingSlow:         movingSlow,
		directions:         directions,
		shootingRegularGun: shootingRegularGun,
		shootingSpecialGun: false,
		guarding:           false,
	}
}

var gameInput = InputController{}
