package main

import (
	"log"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	movingSlow         bool
	directions         []int8
	shootingRegularGun bool
	shootingSpecialGun bool
	guarding           bool
}

type InputController struct {
	gamepads       []ebiten.GamepadID
	keys           []ebiten.Key
	gamepadIDsBuf  []ebiten.GamepadID
	gamepadIDs     map[ebiten.GamepadID]struct{}
	axes           map[ebiten.GamepadID][]float32
	pressedButtons map[ebiten.GamepadID][]string
	gamepadActive  bool
}

func (iC InputController) TranslateInput() *Input {

	// Check if the keyboard has retaken Control
	if len(iC.keys) > 0 {
		iC.gamepadActive = false
	}

	if iC.gamepadActive {

		return iC.translateControllerInput()
	}
	return iC.translateKeyboardInput()
}

func (iC InputController) translateKeyboardInput() *Input {
	directions := []int8{0, 0}
	movingSlow := false
	shootingRegularGun := false
	for _, key := range iC.keys {
		// Checks if the key pressed is within the bounds of the keybindings
		if val, ok := keyboardBindings[key.String()]; ok {
			// Switch case to check the inputs
			directions, movingSlow, shootingRegularGun = translateButtonInputs(val, directions, movingSlow, shootingRegularGun)
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

func (iC InputController) translateControllerInput() *Input {
	directions := []int8{0, 0}
	movingSlow := false
	shootingRegularGun := false
	// Check for each gamepad
	for _, axis := range iC.axes {
		if len(axis) > 0 {
			// Check axis for x direction
			xAxis := float32(axis[0])
			if xAxis > CONTROLLER_DEADZONE {
				directions[0] = 1
			} else if xAxis < -CONTROLLER_DEADZONE {
				directions[0] = -1
			}

			// Check axis for y direction
			yAxis := float32(axis[1])

			if yAxis > CONTROLLER_DEADZONE {
				directions[1] = 1
			} else if yAxis < -CONTROLLER_DEADZONE {
				directions[1] = -1
			}
		}
	}

	// Check for each gamepad
	for _, buttons := range iC.pressedButtons {
		if len(buttons) > 0 {
			for _, button := range buttons {
				// This is potentionally a bad practice since it takes the last gamepad with input on it.
				if val, ok := controllerBindings[button]; ok {
					directions, movingSlow, shootingRegularGun = translateButtonInputs(val, directions, movingSlow, shootingRegularGun)
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

func translateButtonInputs(val string, directions []int8, movingSlow bool, shootingRegularGun bool) ([]int8, bool, bool) {

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
	return directions, movingSlow, shootingRegularGun
}

func (ic *InputController) Update() error {

	if ic.gamepadIDs == nil {
		ic.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	// Log the gamepad connection events.
	ic.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(ic.gamepadIDsBuf[:0])
	for _, id := range ic.gamepadIDsBuf {
		log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
		ic.gamepadIDs[id] = struct{}{}
	}
	for id := range ic.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("gamepad disconnected: id: %d", id)
			delete(ic.gamepadIDs, id)
		}
	}

	ic.axes = map[ebiten.GamepadID][]float32{}
	ic.pressedButtons = map[ebiten.GamepadID][]string{}
	for id := range ic.gamepadIDs {
		maxAxis := ebiten.GamepadAxisCount(id)
		for a := 0; a < maxAxis; a++ {
			v := ebiten.GamepadAxisValue(id, a)

			// sets the gamePad to active
			if v != 0 {
				ic.gamepadActive = true
			}
			ic.axes[id] = append(ic.axes[id], float32(v))
		}

		maxButton := ebiten.GamepadButton(ebiten.GamepadButtonCount(id))
		for b := ebiten.GamepadButton(id); b < maxButton; b++ {
			if ebiten.IsGamepadButtonPressed(id, b) {
				ic.pressedButtons[id] = append(ic.pressedButtons[id], strconv.Itoa(int(b)))
			}
		}

		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				// Log button events.
				if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
					var strong float64
					var weak float64
					switch b {
					case ebiten.StandardGamepadButtonLeftTop,
						ebiten.StandardGamepadButtonLeftLeft,
						ebiten.StandardGamepadButtonLeftRight,
						ebiten.StandardGamepadButtonLeftBottom:
						weak = 0.5
					case ebiten.StandardGamepadButtonRightTop,
						ebiten.StandardGamepadButtonRightLeft,
						ebiten.StandardGamepadButtonRightRight,
						ebiten.StandardGamepadButtonRightBottom:
						strong = 0.5
					}
					if strong > 0 || weak > 0 {
						op := &ebiten.VibrateGamepadOptions{
							Duration:        200 * time.Millisecond,
							StrongMagnitude: strong,
							WeakMagnitude:   weak,
						}
						ebiten.VibrateGamepad(id, op)
					}
					log.Printf("standard button pressed: id: %d, button: %d", id, b)
				}
				if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
					log.Printf("standard button released: id: %d, button: %d", id, b)
				}
			}
		}
	}
	for _, v := range ic.pressedButtons {
		if len(v) > 0 {
			ic.gamepadActive = true
		}
	}

	return nil
}

var gameInput = InputController{}
