package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type GuiElement interface {
	Draw(screen *ebiten.Image)
	Update()
	Start()
}

var GuiElements = make(map[GuiElement]struct{})

type GuiText struct {
	font font.Face
	*Vector
	text         string
	displayLabel string
}

func (menu *GuiText) Draw(screen *ebiten.Image) {
	text.Draw(screen, menu.text, menu.font, int(menu.X), int(menu.Y), color.White)

}

func (menu *GuiText) Update() {
}

func (menu *GuiText) Start() {
}

func NewGuiHealthBar(health *Health, X float64, Y float64, label string) *GuiHealthBar {
	return &GuiHealthBar{
		GuiText: &GuiText{
			font: LoadFont("fonts/FiraCode.ttf", opentype.FaceOptions{}),
			text: "Health: 100%",
			Vector: &Vector{
				X: X,
				Y: Y,
			},
			displayLabel: label,
		},

		Health: health}
}

type GuiHealthBar struct {
	*GuiText
	Health *Health
}

func (healthBar *GuiHealthBar) Update() {
	precentage := int((float64(healthBar.Health.HitPoints) / float64(healthBar.Health.MaxHitPoints)) * 100)
	healthBar.text = fmt.Sprintf("%s Health: %d %%, %d/%d", healthBar.displayLabel, precentage, healthBar.Health.HitPoints, healthBar.Health.MaxHitPoints)
}

func SpawnHealthBar[T GuiElement](obj T) T {
	GuiElements[obj] = struct{}{}
	return obj
}
