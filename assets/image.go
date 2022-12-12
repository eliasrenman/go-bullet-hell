package assets

import (
	"image"
	_ "image/png"
	"io"
	"log"

	"github.com/eliasrenman/go-bullet-hell/constant"
	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/util"
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

	dataImg, _, err := image.Decode(data)
	if err != nil {
		log.Fatal(err)
	}

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

func TranslateScaleAndRotateImage(geom *ebiten.GeoM, position geometry.Point, scale geometry.Size, rotation float64) {
	geom.Translate(position.X, position.Y)
	geom.Scale(float64(scale.Width), float64(scale.Height))
	geom.Rotate(rotation)
}

func (image *Image) Draw(target *ebiten.Image, position geometry.Point, scale geometry.Size, rotation float64) {
	op := &ebiten.DrawImageOptions{}

	// This moves the inital drawing offset.
	position.Add(constant.WORLD_ORIGIN)
	position.Subtract(image.Origin.Dot(geometry.Vector{
		X: image.Size.Width,
		Y: image.Size.Height,
	}))

	TranslateScaleAndRotateImage(&op.GeoM, position, scale, rotation)

	target.DrawImage(image.Image, op)
}

func (image *Image) DrawWithShader(target *ebiten.Image, position geometry.Point, scale geometry.Size, rotation float64, shader *ebiten.Shader, uniforms map[string]any) {
	op := &ebiten.DrawRectShaderOptions{}

	width, height := image.Image.Size()

	// Setting information for the shader to read.
	op.Uniforms = map[string]any{
		"Time":       float32(util.CurrentSeconds()),
		"Resolution": []float32{float32(width), float32(height)},
	}

	// Add any additional uniforms passed to this method.
	for key, value := range uniforms {
		op.Uniforms[key] = value
	}

	op.Images[0] = image.Image
	TranslateScaleAndRotateImage(&op.GeoM, position, scale, rotation)
	target.DrawRectShader(width, height, shader, op)
}

func (image *Image) DrawTiled(target *ebiten.Image, position geometry.Point, scale geometry.Size, rotation float64, offset geometry.Vector) {
	//  TODO: Move the initalisation of these assets and move them to a shader struct and pass a pointer.
	data, _ := Assets.Open("data/shaders/tile.go")
	bytes, _ := io.ReadAll(data)
	shader, _ := ebiten.NewShader(bytes)

	image.DrawWithShader(target, position, scale, rotation, shader, map[string]any{
		"Offset": []float32{float32(offset.X), float32(offset.Y)},
	})
}
