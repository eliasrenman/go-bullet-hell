package main

import (
	_ "image/png"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

var (
	backgroundImage *ebiten.Image = LoadImage("./data/bg/playfield.png")
	hitbox          *ebiten.Image = InitalizeHitbox()
)

func InitalizeHitbox() *ebiten.Image {
	dc := gg.NewContext(hitboxDimension, hitboxDimension)
	dc.SetRGB255(255, 0, 0)
	dc.DrawCircle(hitboxDimension/2, hitboxDimension/2, hitboxDimension/2)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
