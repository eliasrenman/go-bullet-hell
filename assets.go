package main

import (
	"image/color"
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
	playerBullet    *ebiten.Image = InitalizePlayerBullet()
)

func InitalizeHitbox() *ebiten.Image {
	dc := gg.NewContext(hitboxDimension, hitboxDimension)
	dc.SetRGBA255(0, 255, 255, 255)
	dc.DrawCircle(hitboxDimension/2, hitboxDimension/2, hitboxDimension/2)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func InitalizePlayerBullet() *ebiten.Image {
	dc := gg.NewContext(regularBulletSize, regularBulletSize)
	halfBulletSize := float64(regularBulletSize / 2)
	grad := gg.NewRadialGradient(halfBulletSize, halfBulletSize, halfBulletSize, halfBulletSize, halfBulletSize, regularBulletSize*0.002)

	grad.AddColorStop(0, color.RGBA{140, 20, 252, 255})
	grad.AddColorStop(1, color.RGBA{255, 255, 255, 255})

	dc.SetFillStyle(grad)
	dc.DrawCircle(halfBulletSize, halfBulletSize, halfBulletSize)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
