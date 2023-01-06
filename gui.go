package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type GuiText struct {
	font font.Face
	*Vector
	text string
}

func NewGuiText(font font.Face, x float64, y float64) *GuiText {
	return &GuiText{
		font: LoadFont("fonts/FiraCode.ttf", opentype.FaceOptions{}),
	}
}

func (menu *GuiText) Draw(screen *ebiten.Image) {

	text.Draw(screen, "debugText", menu.font, int(menu.X), int(menu.Y), color.White)

}

func (menu *GuiText) Update() {
}

type GuiHealthBar struct {
	*GuiText
	Health *Health
}

func (healthBar *GuiHealthBar) Update() {
	precentage := int((float64(healthBar.Health.HitPoints) / float64(healthBar.Health.MaxHitPoints)) * 100)
	healthBar.text = fmt.Sprintf("Health: %d %%, %d/%d", precentage, healthBar.Health.HitPoints, healthBar.Health.MaxHitPoints)
}
