package main

import (
	"fmt"

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
	for _, key := range iC.keys {
		fmt.Println(key)

	}
	return &Input{
		movingSlow:         false,
		directionLeft:      false,
		directionRight:     true,
		shootingMainGun:    false,
		shootingSpecialGun: false,
		guarding:           false,
	}
}

var gameInput = InputController{}
