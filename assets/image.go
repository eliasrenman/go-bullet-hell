package assets

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/eliasrenman/go-bullet-hell/constant"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	OriginTop         = geometry.Point{X: 0.5, Y: 0}
	OriginTopLeft     = geometry.Point{X: 0, Y: 0}
	OriginTopRight    = geometry.Point{X: 1, Y: 0}
	OriginCenter      = geometry.Point{X: 0.5, Y: 0.5}
	OriginLeft        = geometry.Point{X: 0, Y: 0.5}
	OriginRight       = geometry.Point{X: 1, Y: 0.5}
	OriginBottom      = geometry.Point{X: 0.5, Y: 1}
	OriginBottomLeft  = geometry.Point{X: 0, Y: 1}
	OriginBottomRight = geometry.Point{X: 1, Y: 1}
)

type Image struct {
	*ebiten.Image
	Size   geometry.Size
	Origin geometry.Point
}

func LoadImage(path string, origin geometry.Point) *Image {
	data, err := Assets.Open("data/" + path)
	if err != nil {
		log.Fatal(err)
	}

	dataImg, format, err := image.Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(format)

	image := ebiten.NewImageFromImage(dataImg)

	width, height := image.Size()
	return &Image{
		Image: image,
		Size: geometry.Size{
			Width:  float64(width),
			Height: float64(height),
		},
		Origin: origin,
	}
}

func (image *Image) Draw(target *ebiten.Image, position geometry.Point, scale geometry.Size, rotation float64) {
	op := &ebiten.DrawImageOptions{}

	position.Subtract(image.Origin).Add(constant.WORLD_ORIGIN)
	op.GeoM.Translate(position.X, position.Y)
	op.GeoM.Scale(float64(scale.Width), float64(scale.Height))
	op.GeoM.Rotate(rotation)

	target.DrawImage(image.Image, op)
}
