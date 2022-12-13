package assets

import (
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	Image    *Image
	Velocity geometry.Vector
	offset   geometry.Vector
}

func (background *Background) Update() {
	offset := background.Velocity.ScaledBy((1. / 100) / 60)
	background.offset.Add(offset)
	// background.offset.Scale(util.CurrentSeconds())
}

func (background *Background) Draw(screen *ebiten.Image) {

	// background.offset.Multiply(background.Velocity)

	// background.offset.Add(background.Velocity)
	background.Image.DrawTiled(screen, geometry.Point{}, geometry.Size{Width: 2, Height: 2}, 0, background.offset)

}