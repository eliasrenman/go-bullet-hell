package assets

import (
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	Image    *Image
	Velocity geometry.Vector
}

func (background *Background) Draw(screen *ebiten.Image) {
	offset := geometry.Vector{X: 1, Y: 1}
	offset.Multiply(background.Velocity)
	offset.Scale(util.CurrentSeconds())
	offset.Scale(1. / 100)

	background.Image.DrawTiled(screen, geometry.Point{}, geometry.Size{Width: 2, Height: 2}, 0, offset)

}
