package assets

import (
  "github.com/eliasrenman/go-bullet-hell/geometry"
)

type Background struct {
	image *Image
  vector geometry.Vector
}

func (background *Background) TileDraw() {}

func NewBackground(img *Image) *Background {
  return &Background{
    image: img,
    vector: geometry.Vector{
      X: 0,
      Y: 1,
    },
  };
}