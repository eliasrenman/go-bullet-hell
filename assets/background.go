package assets

import (
	"fmt"

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

	fmt.Println(offset)

	background.Image.DrawTiled(screen, geometry.Point{}, geometry.Size{Width: 1, Height: 1}, 0, offset)

}
